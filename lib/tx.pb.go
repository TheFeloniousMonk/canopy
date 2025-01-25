// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.19.1
// source: tx.proto

package lib

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
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
// Transaction represents a request or action submitted to the network like transfer assets or perform other operations
// within the blockchain system
type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// message_type: The type of the transaction like 'send' or 'stake'
	MessageType string `protobuf:"bytes,1,opt,name=message_type,json=messageType,proto3" json:"message_type,omitempty"`
	// msg: The actual transaction message payload, which is encapsulated in a generic message format
	Msg *anypb.Any `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	// signature: The cryptographic signature used to verify the authenticity of the transaction
	Signature *Signature `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
	// time: The timestamp when the transaction was created
	Time uint64 `protobuf:"varint,4,opt,name=time,proto3" json:"time,omitempty"`
	// fee: The fee associated with processing the transaction
	Fee uint64 `protobuf:"varint,5,opt,name=fee,proto3" json:"fee,omitempty"`
	// memo: An optional message or note attached to the transaction
	Memo string `protobuf:"bytes,6,opt,name=memo,proto3" json:"memo,omitempty"`
	// network_id: The identity of the network the transaction is intended for
	NetworkId uint64 `protobuf:"varint,7,opt,name=network_id,json=networkId,proto3" json:"network_id,omitempty"`
	// chain_id: The identity of the committee the transaction is intended for
	ChainId uint64 `protobuf:"varint,8,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tx_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_tx_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_tx_proto_rawDescGZIP(), []int{0}
}

func (x *Transaction) GetMessageType() string {
	if x != nil {
		return x.MessageType
	}
	return ""
}

func (x *Transaction) GetMsg() *anypb.Any {
	if x != nil {
		return x.Msg
	}
	return nil
}

func (x *Transaction) GetSignature() *Signature {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *Transaction) GetTime() uint64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *Transaction) GetFee() uint64 {
	if x != nil {
		return x.Fee
	}
	return 0
}

func (x *Transaction) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

func (x *Transaction) GetNetworkId() uint64 {
	if x != nil {
		return x.NetworkId
	}
	return 0
}

func (x *Transaction) GetChainId() uint64 {
	if x != nil {
		return x.ChainId
	}
	return 0
}

// TxResult represents the result of a processed transaction, including information about the sender, recipient,
// transaction hash, and the associated block height and index.
type TxResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// sender: The address of the user sending the transaction
	Sender []byte `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	// recipient: The address of the user receiving the transaction
	Recipient []byte `protobuf:"bytes,2,opt,name=recipient,proto3" json:"recipient,omitempty"`
	// message_type: The type of the transaction like 'send' or 'stake'
	MessageType string `protobuf:"bytes,3,opt,name=message_type,json=messageType,proto3" json:"message_type,omitempty"`
	// height: The block height at which the transaction was included
	Height uint64 `protobuf:"varint,4,opt,name=height,proto3" json:"height,omitempty"`
	// index: The position of the transaction within the block
	Index uint64 `protobuf:"varint,5,opt,name=index,proto3" json:"index,omitempty"`
	// transaction: The original transaction object
	Transaction *Transaction `protobuf:"bytes,6,opt,name=transaction,proto3" json:"transaction,omitempty"`
	// tx_hash: The unique hash that identifies the transaction
	TxHash string `protobuf:"bytes,7,opt,name=tx_hash,json=txHash,proto3" json:"tx_hash,omitempty"`
}

func (x *TxResult) Reset() {
	*x = TxResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tx_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TxResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TxResult) ProtoMessage() {}

func (x *TxResult) ProtoReflect() protoreflect.Message {
	mi := &file_tx_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TxResult.ProtoReflect.Descriptor instead.
func (*TxResult) Descriptor() ([]byte, []int) {
	return file_tx_proto_rawDescGZIP(), []int{1}
}

func (x *TxResult) GetSender() []byte {
	if x != nil {
		return x.Sender
	}
	return nil
}

func (x *TxResult) GetRecipient() []byte {
	if x != nil {
		return x.Recipient
	}
	return nil
}

func (x *TxResult) GetMessageType() string {
	if x != nil {
		return x.MessageType
	}
	return ""
}

func (x *TxResult) GetHeight() uint64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *TxResult) GetIndex() uint64 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *TxResult) GetTransaction() *Transaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

func (x *TxResult) GetTxHash() string {
	if x != nil {
		return x.TxHash
	}
	return ""
}

// A Signature is a digital signature is a cryptographic "fingerprint" created with a private key,
// allowing others to verify the authenticity and integrity of a message using the corresponding public key
type Signature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// public_key: is a cryptographic code shared openly, used to verify digital signatures
	PublicKey []byte `protobuf:"bytes,1,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	// signature: the bytes of the signature output from a private key which may be verified with the message and public
	Signature []byte `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *Signature) Reset() {
	*x = Signature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tx_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Signature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Signature) ProtoMessage() {}

func (x *Signature) ProtoReflect() protoreflect.Message {
	mi := &file_tx_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Signature.ProtoReflect.Descriptor instead.
func (*Signature) Descriptor() ([]byte, []int) {
	return file_tx_proto_rawDescGZIP(), []int{2}
}

func (x *Signature) GetPublicKey() []byte {
	if x != nil {
		return x.PublicKey
	}
	return nil
}

func (x *Signature) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

var File_tx_proto protoreflect.FileDescriptor

var file_tx_proto_rawDesc = []byte{
	0x0a, 0x08, 0x74, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfc, 0x01, 0x0a,
	0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x26, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41,
	0x6e, 0x79, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x2e, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x09, 0x73, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x66,
	0x65, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x66, 0x65, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x65, 0x6d,
	0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x69, 0x64, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x64,
	0x12, 0x19, 0x0a, 0x08, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x07, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x22, 0xe0, 0x01, 0x0a, 0x08,
	0x54, 0x78, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x09, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x21,
	0x0a, 0x0c, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12,
	0x34, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x78, 0x5f, 0x68, 0x61, 0x73, 0x68,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x78, 0x48, 0x61, 0x73, 0x68, 0x22, 0x48,
	0x0a, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73,
	0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x61, 0x6e, 0x6f, 0x70, 0x79, 0x2d, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x63, 0x61, 0x6e, 0x6f, 0x70, 0x79, 0x2f, 0x6c, 0x69, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tx_proto_rawDescOnce sync.Once
	file_tx_proto_rawDescData = file_tx_proto_rawDesc
)

func file_tx_proto_rawDescGZIP() []byte {
	file_tx_proto_rawDescOnce.Do(func() {
		file_tx_proto_rawDescData = protoimpl.X.CompressGZIP(file_tx_proto_rawDescData)
	})
	return file_tx_proto_rawDescData
}

var file_tx_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_tx_proto_goTypes = []interface{}{
	(*Transaction)(nil), // 0: types.Transaction
	(*TxResult)(nil),    // 1: types.TxResult
	(*Signature)(nil),   // 2: types.Signature
	(*anypb.Any)(nil),   // 3: google.protobuf.Any
}
var file_tx_proto_depIdxs = []int32{
	3, // 0: types.Transaction.msg:type_name -> google.protobuf.Any
	2, // 1: types.Transaction.signature:type_name -> types.Signature
	0, // 2: types.TxResult.transaction:type_name -> types.Transaction
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_tx_proto_init() }
func file_tx_proto_init() {
	if File_tx_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tx_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction); i {
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
		file_tx_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TxResult); i {
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
		file_tx_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Signature); i {
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
			RawDescriptor: file_tx_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tx_proto_goTypes,
		DependencyIndexes: file_tx_proto_depIdxs,
		MessageInfos:      file_tx_proto_msgTypes,
	}.Build()
	File_tx_proto = out.File
	file_tx_proto_rawDesc = nil
	file_tx_proto_goTypes = nil
	file_tx_proto_depIdxs = nil
}
