// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.19.1
// source: validator.proto

package types

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
//
// A Validator is an actor in the blockchain network responsible for verifying transactions, creating new blocks,
// and maintaining the blockchain's integrity and security. In Canopy, Validators provide this service for multiple
// chains by restaking their tokens across multiple 'committees' and run a blockchain client for each
// Both the Operator and the Output private key may sign transactions on behalf of the Validator, but only the
// Output private key may change the Output address.
type Validator struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// address: the short version of the public key of the operator of the service, corresponding to the 2
	Address []byte `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// public_key: the public cryptographic identity of the operator of the service. This key must be an aggregable BLS
	// key for efficiency in the BFT process.
	PublicKey []byte `protobuf:"bytes,2,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	// net_address: the tcp peer-to-peer address of the node to enable easy discovery of the peer for multi-consensus
	NetAddress string `protobuf:"bytes,3,opt,name=net_address,json=netAddress,proto3" json:"net_address,omitempty"`
	// staked_amount: the amount of tokens locked as a surety bond against malicious behavior. These tokens may be
	// increased by auto-compounding rewards or by an edit-stake command. This bond is returned to the output address
	// after executing an unstake command and waiting the unstaking period.
	StakedAmount uint64 `protobuf:"varint,4,opt,name=staked_amount,json=stakedAmount,proto3" json:"staked_amount,omitempty"`
	// committees: a list of ids of the committees the validator is offering Validation or Delegation services for
	Committees []uint64 `protobuf:"varint,5,rep,packed,name=committees,proto3" json:"committees,omitempty"`
	// max_paused_height: if the Validator is paused, this value tracks the maximum height it may be paused before it
	// automatically begins unstaking
	MaxPausedHeight uint64 `protobuf:"varint,6,opt,name=max_paused_height,json=maxPausedHeight,proto3" json:"max_paused_height,omitempty"`
	// unstaking_height: if the Validator is unstaking, this value tracks the future block height a Validator's surety
	// bond will be returned
	UnstakingHeight uint64 `protobuf:"varint,7,opt,name=unstaking_height,json=unstakingHeight,proto3" json:"unstaking_height,omitempty"`
	// output: the address where early-withdrawal rewards and the unstaking surety bond are transferred to
	Output []byte `protobuf:"bytes,8,opt,name=output,proto3" json:"output,omitempty"`
	// delegate: signals whether the Validator is a Delegate or not. If true, the Validator only passively participates
	// in Validation-as-a-Service by dedicating their funds to help the chain qualify for protocol subsidisation
	Delegate bool `protobuf:"varint,9,opt,name=delegate,proto3" json:"delegate,omitempty"`
	// compound: signals whether the Validator is auto-compounding (increasing their stake with new rewards) or
	// withdrawing rewards early to their output address as they come in
	Compound bool `protobuf:"varint,10,opt,name=compound,proto3" json:"compound,omitempty"`
}

func (x *Validator) Reset() {
	*x = Validator{}
	if protoimpl.UnsafeEnabled {
		mi := &file_validator_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Validator) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Validator) ProtoMessage() {}

func (x *Validator) ProtoReflect() protoreflect.Message {
	mi := &file_validator_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Validator.ProtoReflect.Descriptor instead.
func (*Validator) Descriptor() ([]byte, []int) {
	return file_validator_proto_rawDescGZIP(), []int{0}
}

func (x *Validator) GetAddress() []byte {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *Validator) GetPublicKey() []byte {
	if x != nil {
		return x.PublicKey
	}
	return nil
}

func (x *Validator) GetNetAddress() string {
	if x != nil {
		return x.NetAddress
	}
	return ""
}

func (x *Validator) GetStakedAmount() uint64 {
	if x != nil {
		return x.StakedAmount
	}
	return 0
}

func (x *Validator) GetCommittees() []uint64 {
	if x != nil {
		return x.Committees
	}
	return nil
}

func (x *Validator) GetMaxPausedHeight() uint64 {
	if x != nil {
		return x.MaxPausedHeight
	}
	return 0
}

func (x *Validator) GetUnstakingHeight() uint64 {
	if x != nil {
		return x.UnstakingHeight
	}
	return 0
}

func (x *Validator) GetOutput() []byte {
	if x != nil {
		return x.Output
	}
	return nil
}

func (x *Validator) GetDelegate() bool {
	if x != nil {
		return x.Delegate
	}
	return false
}

func (x *Validator) GetCompound() bool {
	if x != nil {
		return x.Compound
	}
	return false
}

// NonSignerInfo is information that tracks the number of blocks not signed by the Validator within the Non-Sign-Window
type NonSigner struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// address: shorter version of the operator public key of the non signer
	Address []byte `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// counter: increments when a Validator doesn't sign a block and resets every non-sign-window
	Counter uint64 `protobuf:"varint,2,opt,name=counter,proto3" json:"counter,omitempty"`
}

func (x *NonSigner) Reset() {
	*x = NonSigner{}
	if protoimpl.UnsafeEnabled {
		mi := &file_validator_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NonSigner) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NonSigner) ProtoMessage() {}

func (x *NonSigner) ProtoReflect() protoreflect.Message {
	mi := &file_validator_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NonSigner.ProtoReflect.Descriptor instead.
func (*NonSigner) Descriptor() ([]byte, []int) {
	return file_validator_proto_rawDescGZIP(), []int{1}
}

func (x *NonSigner) GetAddress() []byte {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *NonSigner) GetCounter() uint64 {
	if x != nil {
		return x.Counter
	}
	return 0
}

var File_validator_proto protoreflect.FileDescriptor

var file_validator_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x74, 0x79, 0x70, 0x65, 0x73, 0x22, 0xd1, 0x02, 0x0a, 0x09, 0x56, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12,
	0x1f, 0x0a, 0x0b, 0x6e, 0x65, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6e, 0x65, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x23, 0x0a, 0x0d, 0x73, 0x74, 0x61, 0x6b, 0x65, 0x64, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x73, 0x74, 0x61, 0x6b, 0x65, 0x64, 0x41,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74,
	0x65, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x04, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x69,
	0x74, 0x74, 0x65, 0x65, 0x73, 0x12, 0x2a, 0x0a, 0x11, 0x6d, 0x61, 0x78, 0x5f, 0x70, 0x61, 0x75,
	0x73, 0x65, 0x64, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0f, 0x6d, 0x61, 0x78, 0x50, 0x61, 0x75, 0x73, 0x65, 0x64, 0x48, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x12, 0x29, 0x0a, 0x10, 0x75, 0x6e, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x68,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0f, 0x75, 0x6e, 0x73,
	0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x6f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x75, 0x6e, 0x64, 0x22, 0x3f, 0x0a, 0x09,
	0x4e, 0x6f, 0x6e, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x42, 0x2c, 0x5a,
	0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x61, 0x6e, 0x6f,
	0x70, 0x79, 0x2d, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x63, 0x61, 0x6e, 0x6f, 0x70,
	0x79, 0x2f, 0x66, 0x73, 0x6d, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_validator_proto_rawDescOnce sync.Once
	file_validator_proto_rawDescData = file_validator_proto_rawDesc
)

func file_validator_proto_rawDescGZIP() []byte {
	file_validator_proto_rawDescOnce.Do(func() {
		file_validator_proto_rawDescData = protoimpl.X.CompressGZIP(file_validator_proto_rawDescData)
	})
	return file_validator_proto_rawDescData
}

var file_validator_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_validator_proto_goTypes = []interface{}{
	(*Validator)(nil), // 0: types.Validator
	(*NonSigner)(nil), // 1: types.NonSigner
}
var file_validator_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_validator_proto_init() }
func file_validator_proto_init() {
	if File_validator_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_validator_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Validator); i {
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
		file_validator_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NonSigner); i {
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
			RawDescriptor: file_validator_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_validator_proto_goTypes,
		DependencyIndexes: file_validator_proto_depIdxs,
		MessageInfos:      file_validator_proto_msgTypes,
	}.Build()
	File_validator_proto = out.File
	file_validator_proto_rawDesc = nil
	file_validator_proto_goTypes = nil
	file_validator_proto_depIdxs = nil
}
