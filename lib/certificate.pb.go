// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.19.1
// source: certificate.proto

package lib

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// *****************************************************************************************************
// This file is auto-generated from source files in `/lib/.proto/*` using Protocol Buffers (protobuf)
//
// Protobuf is a language-neutral, platform-neutral serialization format. It allows users
// to define objects in a way that’s both efficient to store and fast to transmit over the network.
// These definitions are compiled into code that *enables different systems and programming languages
// to communicate in a byte-perfect manner*
//
// To update these structures, make changes to the source .proto files, then recompile
// to regenerate this file.
// These auto-generated files are easily recognized by checking for a `.pb.go` ending
// *****************************************************************************************************
// _
// _
// _
// A QuorumCertificate is a collection of signatures from a super-majority of validators that confirms consensus on a
// particular block and results. It serves as proof that enough validators have agreed on the block & result’s validity,
// ensuring its acceptance and security in the blockchain.
type QuorumCertificate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// header: is the view of the quorum certificate
	Header *View `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// results: is the certificate result that Canopy uses to process payments, evidence, swaps, and checkpoints
	Results *CertificateResult `protobuf:"bytes,2,opt,name=results,proto3" json:"results,omitempty"`
	// results_hash: is the cryptographic integrity bytes for results, results hash may be used to confirm the validator
	// quorum signed off on the results
	ResultsHash []byte `protobuf:"bytes,3,opt,name=results_hash,json=resultsHash,proto3" json:"results_hash,omitempty"`
	// block: the proposed block to be added to the blockchain
	Block []byte `protobuf:"bytes,4,opt,name=block,proto3" json:"block,omitempty"`
	// block_hash: is the cryptographic integrity bytes for block, block hash may be used to confirm the validator quorum
	// signed off on the block
	BlockHash []byte `protobuf:"bytes,5,opt,name=block_hash,json=blockHash,proto3" json:"block_hash,omitempty"`
	// proposer_key: is the public key of the block proposer
	ProposerKey []byte `protobuf:"bytes,6,opt,name=proposer_key,json=proposerKey,proto3" json:"proposer_key,omitempty"`
	// (aggregate) signature: the compact signature created by combining multiple individual signatures from replica
	// validators. This signature serves as a justification that a super-majority quorum signed off on the certificate
	Signature *AggregateSignature `protobuf:"bytes,7,opt,name=signature,proto3" json:"signature,omitempty"` // aggregate signature from the current proposer message
}

func (x *QuorumCertificate) Reset() {
	*x = QuorumCertificate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_certificate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuorumCertificate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuorumCertificate) ProtoMessage() {}

func (x *QuorumCertificate) ProtoReflect() protoreflect.Message {
	mi := &file_certificate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuorumCertificate.ProtoReflect.Descriptor instead.
func (*QuorumCertificate) Descriptor() ([]byte, []int) {
	return file_certificate_proto_rawDescGZIP(), []int{0}
}

func (x *QuorumCertificate) GetHeader() *View {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *QuorumCertificate) GetResults() *CertificateResult {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *QuorumCertificate) GetResultsHash() []byte {
	if x != nil {
		return x.ResultsHash
	}
	return nil
}

func (x *QuorumCertificate) GetBlock() []byte {
	if x != nil {
		return x.Block
	}
	return nil
}

func (x *QuorumCertificate) GetBlockHash() []byte {
	if x != nil {
		return x.BlockHash
	}
	return nil
}

func (x *QuorumCertificate) GetProposerKey() []byte {
	if x != nil {
		return x.ProposerKey
	}
	return nil
}

func (x *QuorumCertificate) GetSignature() *AggregateSignature {
	if x != nil {
		return x.Signature
	}
	return nil
}

// CertificateResult contains the outcome of a Quorum Certificate as it pertains to Canopy's base chain
type CertificateResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// reward_recipients: the recipients who are rewarded based on the quorum certificate
	// specifically who the committee agreed to reward from the committee treasury
	RewardRecipients *RewardRecipients `protobuf:"bytes,1,opt,name=reward_recipients,json=rewardRecipients,proto3" json:"reward_recipients,omitempty"`
	// slash_recipients: the recipients who are penalized (slashed) based on the quorum certificate
	// specifically who the committee agreed to slash due to evidence of bad behavior
	SlashRecipients *SlashRecipients `protobuf:"bytes,2,opt,name=slash_recipients,json=slashRecipients,proto3" json:"slash_recipients,omitempty"`
	// orders: contains information regarding the 'buying side' of sell orders
	// including actions like 'buy/reserve order' or 'close/complete order'
	Orders *Orders `protobuf:"bytes,3,opt,name=orders,proto3" json:"orders,omitempty"`
	// checkpoint: contains information from the 3rd party chain in order for Canopy to provide Checkpoint-as-a-Service
	Checkpoint *Checkpoint `protobuf:"bytes,4,opt,name=checkpoint,proto3" json:"checkpoint,omitempty"`
}

func (x *CertificateResult) Reset() {
	*x = CertificateResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_certificate_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CertificateResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CertificateResult) ProtoMessage() {}

func (x *CertificateResult) ProtoReflect() protoreflect.Message {
	mi := &file_certificate_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CertificateResult.ProtoReflect.Descriptor instead.
func (*CertificateResult) Descriptor() ([]byte, []int) {
	return file_certificate_proto_rawDescGZIP(), []int{1}
}

func (x *CertificateResult) GetRewardRecipients() *RewardRecipients {
	if x != nil {
		return x.RewardRecipients
	}
	return nil
}

func (x *CertificateResult) GetSlashRecipients() *SlashRecipients {
	if x != nil {
		return x.SlashRecipients
	}
	return nil
}

func (x *CertificateResult) GetOrders() *Orders {
	if x != nil {
		return x.Orders
	}
	return nil
}

func (x *CertificateResult) GetCheckpoint() *Checkpoint {
	if x != nil {
		return x.Checkpoint
	}
	return nil
}

// RewardRecipients is the list of recipients who will receive rewards from the committee's treasury pool,
// based on decisions confirmed in a Quorum Certificate
type RewardRecipients struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// payment_percents: the percentage of rewards allocated to each recipient
	PaymentPercents []*PaymentPercents `protobuf:"bytes,1,rep,name=payment_percents,json=paymentPercents,proto3" json:"payment_percents,omitempty"`
	// number_of_samples: (internal processing only) the number of samples used to determine reward distribution
	NumberOfSamples uint64 `protobuf:"varint,2,opt,name=number_of_samples,json=numberOfSamples,proto3" json:"number_of_samples,omitempty"`
}

func (x *RewardRecipients) Reset() {
	*x = RewardRecipients{}
	if protoimpl.UnsafeEnabled {
		mi := &file_certificate_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RewardRecipients) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RewardRecipients) ProtoMessage() {}

func (x *RewardRecipients) ProtoReflect() protoreflect.Message {
	mi := &file_certificate_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RewardRecipients.ProtoReflect.Descriptor instead.
func (*RewardRecipients) Descriptor() ([]byte, []int) {
	return file_certificate_proto_rawDescGZIP(), []int{2}
}

func (x *RewardRecipients) GetPaymentPercents() []*PaymentPercents {
	if x != nil {
		return x.PaymentPercents
	}
	return nil
}

func (x *RewardRecipients) GetNumberOfSamples() uint64 {
	if x != nil {
		return x.NumberOfSamples
	}
	return 0
}

// SlashRecipients is the list of recipients who are penalized based on misbehavior, like double signing or bad
// proposing based on a committee agreement confirmed in a Quorum Certificate
type SlashRecipients struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// double_signers: a list of actors who the committee agreed double-signed based on evidence
	DoubleSigners []*DoubleSigner `protobuf:"bytes,1,rep,name=double_signers,json=doubleSigners,proto3" json:"double_signers,omitempty"`
}

func (x *SlashRecipients) Reset() {
	*x = SlashRecipients{}
	if protoimpl.UnsafeEnabled {
		mi := &file_certificate_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SlashRecipients) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SlashRecipients) ProtoMessage() {}

func (x *SlashRecipients) ProtoReflect() protoreflect.Message {
	mi := &file_certificate_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SlashRecipients.ProtoReflect.Descriptor instead.
func (*SlashRecipients) Descriptor() ([]byte, []int) {
	return file_certificate_proto_rawDescGZIP(), []int{3}
}

func (x *SlashRecipients) GetDoubleSigners() []*DoubleSigner {
	if x != nil {
		return x.DoubleSigners
	}
	return nil
}

// Orders: tracks actions related to 'buyer side' activities for sell orders
// The committee monitors the 3rd party chain for actions such as intent to buy, funds sent,
// and funds not sent, and communicates these states to the Canopy chain
type Orders struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// buy_orders: a list of actions where a buyer expresses an intent to purchase an order,
	// often referred to as 'claiming' the order
	BuyOrders []*BuyOrder `protobuf:"bytes,1,rep,name=buy_orders,json=buyOrders,proto3" json:"buy_orders,omitempty"`
	// reset_orders: a list of orders where no funds were sent before the deadline,
	// signaling to Canopy to 'un-claim' the order
	ResetOrders []uint64 `protobuf:"varint,2,rep,packed,name=reset_orders,json=resetOrders,proto3" json:"reset_orders,omitempty"`
	// close_orders: a list of orders where funds were sent,
	// signaling Canopy to transfer escrowed tokens to the buyer's Canopy address
	CloseOrders []uint64 `protobuf:"varint,3,rep,packed,name=close_orders,json=closeOrders,proto3" json:"close_orders,omitempty"`
}

func (x *Orders) Reset() {
	*x = Orders{}
	if protoimpl.UnsafeEnabled {
		mi := &file_certificate_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Orders) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Orders) ProtoMessage() {}

func (x *Orders) ProtoReflect() protoreflect.Message {
	mi := &file_certificate_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Orders.ProtoReflect.Descriptor instead.
func (*Orders) Descriptor() ([]byte, []int) {
	return file_certificate_proto_rawDescGZIP(), []int{4}
}

func (x *Orders) GetBuyOrders() []*BuyOrder {
	if x != nil {
		return x.BuyOrders
	}
	return nil
}

func (x *Orders) GetResetOrders() []uint64 {
	if x != nil {
		return x.ResetOrders
	}
	return nil
}

func (x *Orders) GetCloseOrders() []uint64 {
	if x != nil {
		return x.CloseOrders
	}
	return nil
}

// BuyOrder is a buyer expressing an intent to purchase an order, often referred to as 'claiming' the order
type BuyOrder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// order_id: is the number id that is unique to this committee to identify the order
	OrderId uint64 `protobuf:"varint,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	// buyer_receive_address: the Canopy address where the tokens may be received
	BuyerReceiveAddress []byte `protobuf:"bytes,2,opt,name=buyer_receive_address,json=buyerReceiveAddress,proto3" json:"buyer_receive_address,omitempty"`
	// buyer_send_address: the 'counter asset' address where the tokens will be sent from
	BuyerSendAddress []byte `protobuf:"bytes,3,opt,name=buyer_send_address,json=buyerSendAddress,proto3" json:"buyer_send_address,omitempty"`
	// buyer_chain_deadline: the 'counter asset' chain height at which the buyer must send the 'counter asset' by
	// or the 'intent to buy' will be voided
	BuyerChainDeadline uint64 `protobuf:"varint,4,opt,name=buyer_chain_deadline,json=buyerChainDeadline,proto3" json:"buyer_chain_deadline,omitempty"`
}

func (x *BuyOrder) Reset() {
	*x = BuyOrder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_certificate_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuyOrder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuyOrder) ProtoMessage() {}

func (x *BuyOrder) ProtoReflect() protoreflect.Message {
	mi := &file_certificate_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuyOrder.ProtoReflect.Descriptor instead.
func (*BuyOrder) Descriptor() ([]byte, []int) {
	return file_certificate_proto_rawDescGZIP(), []int{5}
}

func (x *BuyOrder) GetOrderId() uint64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *BuyOrder) GetBuyerReceiveAddress() []byte {
	if x != nil {
		return x.BuyerReceiveAddress
	}
	return nil
}

func (x *BuyOrder) GetBuyerSendAddress() []byte {
	if x != nil {
		return x.BuyerSendAddress
	}
	return nil
}

func (x *BuyOrder) GetBuyerChainDeadline() uint64 {
	if x != nil {
		return x.BuyerChainDeadline
	}
	return 0
}

// Checkpoint is 3rd party chain information that allows Canopy to provide Checkpointing-as-a-Service for the 3rd party
// checkpointing is important to prevent `long-range-attacks` in proof of stake blockchains and is currently the
// secure standard
type Checkpoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// height: the height of the third party chain
	Height uint64 `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	// block_hash: the cryptographic hash of the third party chain block for the height
	BlockHash []byte `protobuf:"bytes,2,opt,name=block_hash,json=blockHash,proto3" json:"block_hash,omitempty"`
}

func (x *Checkpoint) Reset() {
	*x = Checkpoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_certificate_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Checkpoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Checkpoint) ProtoMessage() {}

func (x *Checkpoint) ProtoReflect() protoreflect.Message {
	mi := &file_certificate_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Checkpoint.ProtoReflect.Descriptor instead.
func (*Checkpoint) Descriptor() ([]byte, []int) {
	return file_certificate_proto_rawDescGZIP(), []int{6}
}

func (x *Checkpoint) GetHeight() uint64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *Checkpoint) GetBlockHash() []byte {
	if x != nil {
		return x.BlockHash
	}
	return nil
}

// PaymentPercents represents the distribution of rewards to recipients from the committee treasury pool
// Each recipient is identified by their address and the percentage of the reward they will receive
// Percents are diluted based on how many samples are in the Committee Data
type PaymentPercents struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// address: the address where the tokens will be received
	Address []byte `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// percent: the dilutable share of the committee treasury pool
	Percent uint64 `protobuf:"varint,2,opt,name=percent,proto3" json:"percent,omitempty"`
}

func (x *PaymentPercents) Reset() {
	*x = PaymentPercents{}
	if protoimpl.UnsafeEnabled {
		mi := &file_certificate_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentPercents) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentPercents) ProtoMessage() {}

func (x *PaymentPercents) ProtoReflect() protoreflect.Message {
	mi := &file_certificate_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentPercents.ProtoReflect.Descriptor instead.
func (*PaymentPercents) Descriptor() ([]byte, []int) {
	return file_certificate_proto_rawDescGZIP(), []int{7}
}

func (x *PaymentPercents) GetAddress() []byte {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *PaymentPercents) GetPercent() uint64 {
	if x != nil {
		return x.Percent
	}
	return 0
}

// DoubleSigner identifies a validator who has been caught double signing and should be slashed
// The structure includes the validator's public key and a list of block heights where the double signing occurred
type DoubleSigner struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// pubkey: the cryptographic identifier of the malicious actor
	PubKey []byte `protobuf:"bytes,1,opt,name=pubKey,proto3" json:"pubKey,omitempty"`
	// heights: the list of heights when the infractions occurred
	Heights []uint64 `protobuf:"varint,2,rep,packed,name=heights,proto3" json:"heights,omitempty"`
}

func (x *DoubleSigner) Reset() {
	*x = DoubleSigner{}
	if protoimpl.UnsafeEnabled {
		mi := &file_certificate_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoubleSigner) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoubleSigner) ProtoMessage() {}

func (x *DoubleSigner) ProtoReflect() protoreflect.Message {
	mi := &file_certificate_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoubleSigner.ProtoReflect.Descriptor instead.
func (*DoubleSigner) Descriptor() ([]byte, []int) {
	return file_certificate_proto_rawDescGZIP(), []int{8}
}

func (x *DoubleSigner) GetPubKey() []byte {
	if x != nil {
		return x.PubKey
	}
	return nil
}

func (x *DoubleSigner) GetHeights() []uint64 {
	if x != nil {
		return x.Heights
	}
	return nil
}

var File_certificate_proto protoreflect.FileDescriptor

var file_certificate_proto_rawDesc = []byte{
	0x0a, 0x11, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x74, 0x79, 0x70, 0x65, 0x73, 0x1a, 0x0f, 0x63, 0x6f, 0x6e, 0x73,
	0x65, 0x6e, 0x73, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa0, 0x02, 0x0a, 0x11,
	0x51, 0x75, 0x6f, 0x72, 0x75, 0x6d, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x65, 0x12, 0x23, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0b, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x56, 0x69, 0x65, 0x77, 0x52, 0x06,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x32, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x73, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x0b, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x48, 0x61, 0x73, 0x68, 0x12, 0x14, 0x0a,
	0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x68, 0x61, 0x73,
	0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61,
	0x73, 0x68, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x65, 0x72, 0x5f, 0x6b,
	0x65, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73,
	0x65, 0x72, 0x4b, 0x65, 0x79, 0x12, 0x37, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0xf6,
	0x01, 0x0a, 0x11, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x12, 0x44, 0x0a, 0x11, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x72,
	0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x52, 0x65,
	0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x10, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64,
	0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x41, 0x0a, 0x10, 0x73, 0x6c,
	0x61, 0x73, 0x68, 0x5f, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x53, 0x6c, 0x61,
	0x73, 0x68, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x0f, 0x73, 0x6c,
	0x61, 0x73, 0x68, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x25, 0x0a,
	0x06, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x52, 0x06, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x73, 0x12, 0x31, 0x0a, 0x0a, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x0a, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x22, 0x81, 0x01, 0x0a, 0x10, 0x52, 0x65, 0x77, 0x61,
	0x72, 0x64, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x41, 0x0a, 0x10,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x50,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x0f,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x73, 0x12,
	0x2a, 0x0a, 0x11, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x6f, 0x66, 0x5f, 0x73, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0f, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x4f, 0x66, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x22, 0x4d, 0x0a, 0x0f, 0x53,
	0x6c, 0x61, 0x73, 0x68, 0x52, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x3a,
	0x0a, 0x0e, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x44,
	0x6f, 0x75, 0x62, 0x6c, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x52, 0x0d, 0x64, 0x6f, 0x75,
	0x62, 0x6c, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x73, 0x22, 0x7e, 0x0a, 0x06, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x73, 0x12, 0x2e, 0x0a, 0x0a, 0x62, 0x75, 0x79, 0x5f, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x42, 0x75, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x09, 0x62, 0x75, 0x79, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x73, 0x65, 0x74, 0x5f, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x04, 0x52, 0x0b, 0x72, 0x65, 0x73, 0x65,
	0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6c, 0x6f, 0x73, 0x65,
	0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x04, 0x52, 0x0b, 0x63,
	0x6c, 0x6f, 0x73, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x22, 0xb9, 0x01, 0x0a, 0x08, 0x42,
	0x75, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x32, 0x0a, 0x15, 0x62, 0x75, 0x79, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x63, 0x65,
	0x69, 0x76, 0x65, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x13, 0x62, 0x75, 0x79, 0x65, 0x72, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x2c, 0x0a, 0x12, 0x62, 0x75, 0x79, 0x65, 0x72, 0x5f,
	0x73, 0x65, 0x6e, 0x64, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x10, 0x62, 0x75, 0x79, 0x65, 0x72, 0x53, 0x65, 0x6e, 0x64, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x30, 0x0a, 0x14, 0x62, 0x75, 0x79, 0x65, 0x72, 0x5f, 0x63, 0x68,
	0x61, 0x69, 0x6e, 0x5f, 0x64, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x12, 0x62, 0x75, 0x79, 0x65, 0x72, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x44, 0x65,
	0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x22, 0x43, 0x0a, 0x0a, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x22, 0x45, 0x0a, 0x0f, 0x50,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x18,
	0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x65, 0x72, 0x63,
	0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x70, 0x65, 0x72, 0x63, 0x65,
	0x6e, 0x74, 0x22, 0x40, 0x0a, 0x0c, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x53, 0x69, 0x67, 0x6e,
	0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x06, 0x70, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x68, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x04, 0x52, 0x07, 0x68, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x73, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x61, 0x6e, 0x6f, 0x70, 0x79, 0x2d, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x2f, 0x63, 0x61, 0x6e, 0x6f, 0x70, 0x79, 0x2f, 0x6c, 0x69, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_certificate_proto_rawDescOnce sync.Once
	file_certificate_proto_rawDescData = file_certificate_proto_rawDesc
)

func file_certificate_proto_rawDescGZIP() []byte {
	file_certificate_proto_rawDescOnce.Do(func() {
		file_certificate_proto_rawDescData = protoimpl.X.CompressGZIP(file_certificate_proto_rawDescData)
	})
	return file_certificate_proto_rawDescData
}

var file_certificate_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_certificate_proto_goTypes = []interface{}{
	(*QuorumCertificate)(nil),  // 0: types.QuorumCertificate
	(*CertificateResult)(nil),  // 1: types.CertificateResult
	(*RewardRecipients)(nil),   // 2: types.RewardRecipients
	(*SlashRecipients)(nil),    // 3: types.SlashRecipients
	(*Orders)(nil),             // 4: types.Orders
	(*BuyOrder)(nil),           // 5: types.BuyOrder
	(*Checkpoint)(nil),         // 6: types.Checkpoint
	(*PaymentPercents)(nil),    // 7: types.PaymentPercents
	(*DoubleSigner)(nil),       // 8: types.DoubleSigner
	(*View)(nil),               // 9: types.View
	(*AggregateSignature)(nil), // 10: types.AggregateSignature
}
var file_certificate_proto_depIdxs = []int32{
	9,  // 0: types.QuorumCertificate.header:type_name -> types.View
	1,  // 1: types.QuorumCertificate.results:type_name -> types.CertificateResult
	10, // 2: types.QuorumCertificate.signature:type_name -> types.AggregateSignature
	2,  // 3: types.CertificateResult.reward_recipients:type_name -> types.RewardRecipients
	3,  // 4: types.CertificateResult.slash_recipients:type_name -> types.SlashRecipients
	4,  // 5: types.CertificateResult.orders:type_name -> types.Orders
	6,  // 6: types.CertificateResult.checkpoint:type_name -> types.Checkpoint
	7,  // 7: types.RewardRecipients.payment_percents:type_name -> types.PaymentPercents
	8,  // 8: types.SlashRecipients.double_signers:type_name -> types.DoubleSigner
	5,  // 9: types.Orders.buy_orders:type_name -> types.BuyOrder
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_certificate_proto_init() }
func file_certificate_proto_init() {
	if File_certificate_proto != nil {
		return
	}
	file_consensus_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_certificate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuorumCertificate); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_certificate_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CertificateResult); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_certificate_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RewardRecipients); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_certificate_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SlashRecipients); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_certificate_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Orders); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_certificate_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuyOrder); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_certificate_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Checkpoint); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_certificate_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaymentPercents); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_certificate_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoubleSigner); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_certificate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_certificate_proto_goTypes,
		DependencyIndexes: file_certificate_proto_depIdxs,
		MessageInfos:      file_certificate_proto_msgTypes,
	}.Build()
	File_certificate_proto = out.File
	file_certificate_proto_rawDesc = nil
	file_certificate_proto_goTypes = nil
	file_certificate_proto_depIdxs = nil
}
