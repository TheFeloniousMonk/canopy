// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.19.1
// source: genesis.proto

package types

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GenesisState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Time       *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	Pools      []*Pool                `protobuf:"bytes,2,rep,name=pools,proto3" json:"pools,omitempty"`
	Accounts   []*Account             `protobuf:"bytes,3,rep,name=accounts,proto3" json:"accounts,omitempty"`
	Validators []*Validator           `protobuf:"bytes,4,rep,name=validators,proto3" json:"validators,omitempty"`
	NonSigners []*NonSigner           `protobuf:"bytes,5,rep,name=non_signers,json=nonSigners,proto3" json:"non_signers,omitempty"` // only used for export debugging
	Params     *Params                `protobuf:"bytes,6,opt,name=params,proto3" json:"params,omitempty"`
	Supply     *Supply                `protobuf:"bytes,7,opt,name=supply,proto3" json:"supply,omitempty"` // only used for export debugging
}

func (x *GenesisState) Reset() {
	*x = GenesisState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_genesis_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenesisState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenesisState) ProtoMessage() {}

func (x *GenesisState) ProtoReflect() protoreflect.Message {
	mi := &file_genesis_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenesisState.ProtoReflect.Descriptor instead.
func (*GenesisState) Descriptor() ([]byte, []int) {
	return file_genesis_proto_rawDescGZIP(), []int{0}
}

func (x *GenesisState) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *GenesisState) GetPools() []*Pool {
	if x != nil {
		return x.Pools
	}
	return nil
}

func (x *GenesisState) GetAccounts() []*Account {
	if x != nil {
		return x.Accounts
	}
	return nil
}

func (x *GenesisState) GetValidators() []*Validator {
	if x != nil {
		return x.Validators
	}
	return nil
}

func (x *GenesisState) GetNonSigners() []*NonSigner {
	if x != nil {
		return x.NonSigners
	}
	return nil
}

func (x *GenesisState) GetParams() *Params {
	if x != nil {
		return x.Params
	}
	return nil
}

func (x *GenesisState) GetSupply() *Supply {
	if x != nil {
		return x.Supply
	}
	return nil
}

var File_genesis_proto protoreflect.FileDescriptor

var file_genesis_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x67, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x74, 0x79, 0x70, 0x65, 0x73, 0x1a, 0x0d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x09, 0x67, 0x6f, 0x76, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xc0, 0x02, 0x0a, 0x0c, 0x47, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x74,
	0x69, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x05, 0x70, 0x6f, 0x6f, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x50, 0x6f, 0x6f, 0x6c, 0x52,
	0x05, 0x70, 0x6f, 0x6f, 0x6c, 0x73, 0x12, 0x2a, 0x0a, 0x08, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x08, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x73, 0x12, 0x30, 0x0a, 0x0a, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x56,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x0a, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x6f, 0x72, 0x73, 0x12, 0x31, 0x0a, 0x0b, 0x6e, 0x6f, 0x6e, 0x5f, 0x73, 0x69, 0x67, 0x6e,
	0x65, 0x72, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x4e, 0x6f, 0x6e, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x52, 0x0a, 0x6e, 0x6f, 0x6e,
	0x53, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x73, 0x12, 0x25, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x25,
	0x0a, 0x06, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x52, 0x06, 0x73,
	0x75, 0x70, 0x70, 0x6c, 0x79, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x69, 0x6e, 0x63, 0x68, 0x75, 0x63, 0x6f, 0x2f, 0x67, 0x69, 0x6e,
	0x63, 0x68, 0x75, 0x2f, 0x66, 0x73, 0x6d, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_genesis_proto_rawDescOnce sync.Once
	file_genesis_proto_rawDescData = file_genesis_proto_rawDesc
)

func file_genesis_proto_rawDescGZIP() []byte {
	file_genesis_proto_rawDescOnce.Do(func() {
		file_genesis_proto_rawDescData = protoimpl.X.CompressGZIP(file_genesis_proto_rawDescData)
	})
	return file_genesis_proto_rawDescData
}

var file_genesis_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_genesis_proto_goTypes = []interface{}{
	(*GenesisState)(nil),          // 0: types.GenesisState
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
	(*Pool)(nil),                  // 2: types.Pool
	(*Account)(nil),               // 3: types.Account
	(*Validator)(nil),             // 4: types.Validator
	(*NonSigner)(nil),             // 5: types.NonSigner
	(*Params)(nil),                // 6: types.Params
	(*Supply)(nil),                // 7: types.Supply
}
var file_genesis_proto_depIdxs = []int32{
	1, // 0: types.GenesisState.time:type_name -> google.protobuf.Timestamp
	2, // 1: types.GenesisState.pools:type_name -> types.Pool
	3, // 2: types.GenesisState.accounts:type_name -> types.Account
	4, // 3: types.GenesisState.validators:type_name -> types.Validator
	5, // 4: types.GenesisState.non_signers:type_name -> types.NonSigner
	6, // 5: types.GenesisState.params:type_name -> types.Params
	7, // 6: types.GenesisState.supply:type_name -> types.Supply
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_genesis_proto_init() }
func file_genesis_proto_init() {
	if File_genesis_proto != nil {
		return
	}
	file_account_proto_init()
	file_validator_proto_init()
	file_gov_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_genesis_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenesisState); i {
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
			RawDescriptor: file_genesis_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_genesis_proto_goTypes,
		DependencyIndexes: file_genesis_proto_depIdxs,
		MessageInfos:      file_genesis_proto_msgTypes,
	}.Build()
	File_genesis_proto = out.File
	file_genesis_proto_rawDesc = nil
	file_genesis_proto_goTypes = nil
	file_genesis_proto_depIdxs = nil
}
