package store

import (
	"fmt"
	"github.com/dgraph-io/badger/v4"
	"github.com/ginchuco/ginchu/lib"
	"github.com/ginchuco/ginchu/lib/crypto"
	"math"
	"path/filepath"
)

const (
	stateStorePrefix      = "s/"
	stateCommitmentPrefix = "c/"
	stateCommitIDPrefix   = "x/"
	transactionPrefix     = "t/"

	maxKeyBytes = 128
)

var (
	lastCIDKey = []byte("xl/")

	_ lib.StoreI = &Store{}
)

type Store struct {
	version uint64
	log     lib.LoggerI
	db      *badger.DB
	writer  *badger.Txn
	ss      *TxnWrapper
	sc      *SMTWrapper
	ix      *Indexer
	root    []byte
}

func New(config lib.Config, l lib.LoggerI) (lib.StoreI, lib.ErrorI) {
	if config.StoreConfig.InMemory {
		return NewStoreInMemory(l)
	}
	return NewStore(filepath.Join(config.DataDirPath, config.DBName), l)
}

func NewStore(path string, log lib.LoggerI) (lib.StoreI, lib.ErrorI) {
	db, err := badger.OpenManaged(badger.DefaultOptions(path).
		WithLoggingLevel(badger.ERROR).WithMemTableSize(1000000000))
	if err != nil {
		return nil, ErrOpenDB(err)
	}
	return newStore(db, log)
}

func NewStoreInMemory(log lib.LoggerI) (lib.StoreI, lib.ErrorI) {
	db, err := badger.OpenManaged(badger.DefaultOptions("").
		WithInMemory(true).WithLoggingLevel(badger.ERROR))
	if err != nil {
		return nil, ErrOpenDB(err)
	}
	return newStore(db, log)
}

func newStore(db *badger.DB, log lib.LoggerI) (*Store, lib.ErrorI) {
	id := getLatestCommitID(db, log)
	writer := db.NewTransactionAt(id.Height, true)
	return &Store{
		version: id.Height,
		log:     log,
		db:      db,
		writer:  writer,
		ss:      NewTxnWrapper(writer, log, stateStorePrefix),
		sc:      NewSMTWrapper(NewTxnWrapper(writer, log, stateCommitmentPrefix), id.Root, log),
		ix:      &Indexer{NewTxnWrapper(writer, log, transactionPrefix)},
		root:    id.Root,
	}, nil
}

func (s *Store) NewReadOnly(version uint64) (lib.StoreI, lib.ErrorI) {
	id, err := s.getCommitID(version)
	if err != nil {
		return nil, err
	}
	reader := s.db.NewTransactionAt(version, false)
	return &Store{
		version: version,
		log:     s.log,
		db:      s.db,
		ss:      NewTxnWrapper(reader, s.log, stateStorePrefix),
		sc:      NewSMTWrapper(NewTxnWrapper(reader, s.log, stateCommitmentPrefix), id.Root, s.log),
		ix:      s.ix,
	}, nil
}

func (s *Store) Copy() (lib.StoreI, lib.ErrorI) {
	writer := s.db.NewTransactionAt(s.version, true)
	return &Store{
		version: s.version,
		log:     s.log,
		db:      s.db,
		writer:  writer,
		ss:      NewTxnWrapper(writer, s.log, stateStorePrefix),
		sc:      NewSMTWrapper(NewTxnWrapper(writer, s.log, stateCommitmentPrefix), s.root, s.log),
		ix:      &Indexer{NewTxnWrapper(writer, s.log, transactionPrefix)},
		root:    s.root,
	}, nil
}

func (s *Store) Get(key []byte) ([]byte, lib.ErrorI) { return s.ss.Get(key) }
func (s *Store) Set(k, v []byte) lib.ErrorI {
	if err := s.ss.Set(k, v); err != nil {
		return err
	}
	return s.sc.Set(k, v)
}

func (s *Store) Delete(k []byte) lib.ErrorI {
	if err := s.ss.Delete(k); err != nil {
		return err
	}
	return s.sc.Delete(k)
}

func (s *Store) GetProof(k []byte) ([]byte, []byte, lib.ErrorI) {
	val, err := s.ss.Get(k)
	if err != nil {
		return nil, nil, err
	}
	proof, _, err := s.sc.GetProof(k)
	return proof, val, err
}

func (s *Store) VerifyProof(k, v, p []byte) bool                  { return s.sc.VerifyProof(k, v, p) }
func (s *Store) Iterator(p []byte) (lib.IteratorI, lib.ErrorI)    { return s.ss.Iterator(p) }
func (s *Store) RevIterator(p []byte) (lib.IteratorI, lib.ErrorI) { return s.ss.RevIterator(p) }
func (s *Store) Version() uint64                                  { return s.version }
func (s *Store) NewTxn() lib.StoreTxnI                            { return NewTxn(s) }
func (s *Store) Root() (root []byte, err lib.ErrorI)              { return s.sc.Root() }
func (s *Store) Commit() (root []byte, err lib.ErrorI) {
	s.root, err = s.sc.Commit()
	if err != nil {
		return nil, err
	}
	if err = s.setCommitID(s.version, s.root); err != nil {
		return nil, err
	}
	if err := s.writer.CommitAt(s.version, nil); err != nil {
		return nil, ErrCommitDB(err)
	}
	s.version++
	s.resetWriter(s.root)
	return lib.CopyBytes(s.root), nil
}

func (s *Store) Reset() {
	s.resetWriter(s.root)
}

func (s *Store) Discard() {
	s.writer.Discard()
}

func (s *Store) resetWriter(root []byte) {
	s.writer.Discard()
	s.writer = s.db.NewTransactionAt(s.version, true)
	s.ss.setDB(s.writer)
	s.sc.setDB(NewTxnWrapper(s.writer, s.log, stateCommitmentPrefix), root)
	s.ix.setDB(NewTxnWrapper(s.writer, s.log, transactionPrefix))
}

func (s *Store) commitIDKey(version uint64) []byte {
	return []byte(fmt.Sprintf("%s%d", stateCommitIDPrefix, version))
}

func (s *Store) getCommitID(version uint64) (id CommitID, err lib.ErrorI) {
	var bz []byte
	bz, err = NewTxnWrapper(s.writer, s.log, "").Get(s.commitIDKey(version))
	if err != nil {
		return
	}
	if err = lib.Unmarshal(bz, &id); err != nil {
		return
	}
	return
}

func (s *Store) setCommitID(version uint64, root []byte) lib.ErrorI {
	w := NewTxnWrapper(s.writer, s.log, "")
	value, err := lib.Marshal(&CommitID{
		Height: version,
		Root:   root,
	})
	if err != nil {
		return err
	}
	if err = w.Set(lastCIDKey, value); err != nil {
		return err
	}
	key := s.commitIDKey(version)
	if err != nil {
		return err
	}
	return w.Set(key, value)
}

func getLatestCommitID(db *badger.DB, log lib.LoggerI) (id *CommitID) {
	tx := NewTxnWrapper(db.NewTransactionAt(math.MaxUint64, false), log, "")
	defer tx.Close()
	id = new(CommitID)
	bz, err := tx.Get(lastCIDKey)
	if err != nil {
		log.Fatalf("getLatestCommitID() failed with err: %s", err.Error())
	}
	if err = lib.Unmarshal(bz, id); err != nil {
		log.Fatalf("unmarshalCommitID() failed with err: %s", err.Error())
	}
	if id.Height == 0 {
		id.Height++
	}
	return
}

func (s *Store) IndexTx(result *lib.TxResult) lib.ErrorI { return s.ix.IndexTx(result) }
func (s *Store) DeleteTxsForHeight(height uint64) lib.ErrorI {
	return s.ix.DeleteTxsForHeight(height)
}
func (s *Store) GetTxByHash(h []byte) (*lib.TxResult, lib.ErrorI) { return s.ix.GetTxByHash(h) }

func (s *Store) GetTxsByHeight(height uint64, newestToOldest bool) ([]*lib.TxResult, lib.ErrorI) {
	return s.ix.GetTxsByHeight(height, newestToOldest)
}

func (s *Store) GetTxsBySender(address crypto.AddressI, newestToOldest bool) ([]*lib.TxResult, lib.ErrorI) {
	return s.ix.GetTxsBySender(address, newestToOldest)
}

func (s *Store) GetTxsByRecipient(address crypto.AddressI, newestToOldest bool) ([]*lib.TxResult, lib.ErrorI) {
	return s.ix.GetTxsByRecipient(address, newestToOldest)
}

func (s *Store) IndexBlock(b *lib.BlockResult) lib.ErrorI {
	return s.ix.IndexBlock(b)
}

func (s *Store) DeleteBlockForHeight(h uint64) lib.ErrorI {
	return s.ix.DeleteBlockForHeight(h)
}

func (s *Store) GetBlockByHash(h []byte) (*lib.BlockResult, lib.ErrorI) {
	return s.ix.GetBlockByHash(h)
}

func (s *Store) GetBlockByHeight(h uint64) (*lib.BlockResult, lib.ErrorI) {
	return s.ix.GetBlockByHeight(h)
}

func (s *Store) DeleteDoubleSignersForHeight(h uint64) lib.ErrorI {
	return s.ix.DeleteDoubleSignersForHeight(h)
}

func (s *Store) IndexQC(qc *lib.QuorumCertificate) lib.ErrorI {
	return s.ix.IndexQC(qc)
}

func (s *Store) GetQCByHeight(height uint64) (*lib.QuorumCertificate, lib.ErrorI) {
	return s.ix.GetQCByHeight(height)
}

func (s *Store) DeleteQCForHeight(height uint64) lib.ErrorI {
	return s.ix.DeleteQCForHeight(height)
}

func (s *Store) GetDoubleSigners(h uint64) (*lib.DoubleSigners, lib.ErrorI) {
	return s.ix.GetDoubleSigners(h)
}

func (s *Store) Close() lib.ErrorI {
	s.Discard()
	if err := s.db.Close(); err != nil {
		return ErrCloseDB(s.db.Close())
	}
	return nil
}
