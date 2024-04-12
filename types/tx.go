package types

import (
	"github.com/ginchuco/ginchu/types/crypto"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

var _ TransactionI = &Transaction{}

type TransactionI interface {
	proto.Message
	GetMsg() *anypb.Any
	GetSig() SignatureI
	GetSequence() uint64
	GetBytes() ([]byte, error)
	GetSignBytes() ([]byte, error)
	GetHash() ([]byte, error)
}

func (x *Transaction) GetHash() ([]byte, error) {
	bz, err := x.GetBytes()
	if err != nil {
		return nil, err
	}
	return crypto.Hash(bz), nil
}

func (x *Transaction) GetBytes() ([]byte, error) {
	return cdc.Marshal(x)
}

func (x *Transaction) GetSig() SignatureI {
	return x.Signature
}

func (x *Transaction) GetSignBytes() ([]byte, error) {
	return cdc.Marshal(Transaction{
		Msg:       x.Msg,
		Signature: nil,
		Sequence:  x.Sequence,
	})
}

var _ SignatureI = &Signature{}

type SignatureI interface {
	proto.Message
	GetPublicKey() []byte
	GetSignature() []byte
}

var _ TxResultI = &TxResult{}

func (x *TxResult) GetTx() TransactionI        { return x.Transaction }
func (x *TxResult) GetBytes() ([]byte, ErrorI) { return Marshal(x) }

type TxResultI interface {
	proto.Message
	GetSender() []byte
	GetRecipient() []byte
	GetMessageType() string
	GetHeight() uint64
	GetIndex() uint64
	GetBytes() ([]byte, ErrorI)
	GetTxHash() string
	GetTx() TransactionI
}

func (x *Signature) Copy() *Signature {
	return &Signature{
		PublicKey: CopyBytes(x.PublicKey),
		Signature: CopyBytes(x.Signature),
	}
}
