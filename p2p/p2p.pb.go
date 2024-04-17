// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.19.1
// source: p2p.proto

package p2p

import (
	types "github.com/ginchuco/ginchu/types"
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

type Envelope struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload *anypb.Any `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *Envelope) Reset() {
	*x = Envelope{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p2p_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Envelope) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Envelope) ProtoMessage() {}

func (x *Envelope) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Envelope.ProtoReflect.Descriptor instead.
func (*Envelope) Descriptor() ([]byte, []int) {
	return file_p2p_proto_rawDescGZIP(), []int{0}
}

func (x *Envelope) GetPayload() *anypb.Any {
	if x != nil {
		return x.Payload
	}
	return nil
}

type Packet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StreamId types.Topic `protobuf:"varint,1,opt,name=stream_id,json=streamId,proto3,enum=types.Topic" json:"stream_id,omitempty"`
	Eof      bool        `protobuf:"varint,2,opt,name=eof,proto3" json:"eof,omitempty"`
	Bytes    []byte      `protobuf:"bytes,3,opt,name=bytes,proto3" json:"bytes,omitempty"`
}

func (x *Packet) Reset() {
	*x = Packet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p2p_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Packet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Packet) ProtoMessage() {}

func (x *Packet) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Packet.ProtoReflect.Descriptor instead.
func (*Packet) Descriptor() ([]byte, []int) {
	return file_p2p_proto_rawDescGZIP(), []int{1}
}

func (x *Packet) GetStreamId() types.Topic {
	if x != nil {
		return x.StreamId
	}
	return types.Topic(0)
}

func (x *Packet) GetEof() bool {
	if x != nil {
		return x.Eof
	}
	return false
}

func (x *Packet) GetBytes() []byte {
	if x != nil {
		return x.Bytes
	}
	return nil
}

type Ping struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Ping) Reset() {
	*x = Ping{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p2p_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ping) ProtoMessage() {}

func (x *Ping) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ping.ProtoReflect.Descriptor instead.
func (*Ping) Descriptor() ([]byte, []int) {
	return file_p2p_proto_rawDescGZIP(), []int{2}
}

type Pong struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Pong) Reset() {
	*x = Pong{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p2p_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pong) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pong) ProtoMessage() {}

func (x *Pong) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pong.ProtoReflect.Descriptor instead.
func (*Pong) Descriptor() ([]byte, []int) {
	return file_p2p_proto_rawDescGZIP(), []int{3}
}

type PeerBookRequestMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PeerBookRequestMessage) Reset() {
	*x = PeerBookRequestMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p2p_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeerBookRequestMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeerBookRequestMessage) ProtoMessage() {}

func (x *PeerBookRequestMessage) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeerBookRequestMessage.ProtoReflect.Descriptor instead.
func (*PeerBookRequestMessage) Descriptor() ([]byte, []int) {
	return file_p2p_proto_rawDescGZIP(), []int{4}
}

type PeerBookResponseMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Book []*BookPeer `protobuf:"bytes,1,rep,name=book,proto3" json:"book,omitempty"`
}

func (x *PeerBookResponseMessage) Reset() {
	*x = PeerBookResponseMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p2p_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeerBookResponseMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeerBookResponseMessage) ProtoMessage() {}

func (x *PeerBookResponseMessage) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeerBookResponseMessage.ProtoReflect.Descriptor instead.
func (*PeerBookResponseMessage) Descriptor() ([]byte, []int) {
	return file_p2p_proto_rawDescGZIP(), []int{5}
}

func (x *PeerBookResponseMessage) GetBook() []*BookPeer {
	if x != nil {
		return x.Book
	}
	return nil
}

type BookPeer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address               *types.PeerAddress `protobuf:"bytes,1,opt,name=Address,proto3" json:"Address,omitempty"`
	ConsecutiveFailedDial int32              `protobuf:"varint,2,opt,name=consecutive_failed_dial,json=consecutiveFailedDial,proto3" json:"consecutive_failed_dial,omitempty"`
}

func (x *BookPeer) Reset() {
	*x = BookPeer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p2p_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookPeer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookPeer) ProtoMessage() {}

func (x *BookPeer) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookPeer.ProtoReflect.Descriptor instead.
func (*BookPeer) Descriptor() ([]byte, []int) {
	return file_p2p_proto_rawDescGZIP(), []int{6}
}

func (x *BookPeer) GetAddress() *types.PeerAddress {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *BookPeer) GetConsecutiveFailedDial() int32 {
	if x != nil {
		return x.ConsecutiveFailedDial
	}
	return 0
}

var File_p2p_proto protoreflect.FileDescriptor

var file_p2p_proto_rawDesc = []byte{
	0x0a, 0x09, 0x70, 0x32, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x1a, 0x0a, 0x70, 0x65, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3a, 0x0a, 0x08, 0x45, 0x6e, 0x76,
	0x65, 0x6c, 0x6f, 0x70, 0x65, 0x12, 0x2e, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x07, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x5b, 0x0a, 0x06, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x12,
	0x29, 0x0a, 0x09, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x54, 0x6f, 0x70, 0x69, 0x63,
	0x52, 0x08, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6f,
	0x66, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x65, 0x6f, 0x66, 0x12, 0x14, 0x0a, 0x05,
	0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x62, 0x79, 0x74,
	0x65, 0x73, 0x22, 0x06, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x22, 0x06, 0x0a, 0x04, 0x50, 0x6f,
	0x6e, 0x67, 0x22, 0x18, 0x0a, 0x16, 0x50, 0x65, 0x65, 0x72, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x3e, 0x0a, 0x17,
	0x50, 0x65, 0x65, 0x72, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x23, 0x0a, 0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x42, 0x6f,
	0x6f, 0x6b, 0x50, 0x65, 0x65, 0x72, 0x52, 0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x22, 0x70, 0x0a, 0x08,
	0x42, 0x6f, 0x6f, 0x6b, 0x50, 0x65, 0x65, 0x72, 0x12, 0x2c, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x07, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x36, 0x0a, 0x17, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x63,
	0x75, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x5f, 0x64, 0x69, 0x61,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x15, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x63, 0x75,
	0x74, 0x69, 0x76, 0x65, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x44, 0x69, 0x61, 0x6c, 0x42, 0x20,
	0x5a, 0x1e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x69, 0x6e,
	0x63, 0x68, 0x75, 0x63, 0x6f, 0x2f, 0x67, 0x69, 0x6e, 0x63, 0x68, 0x75, 0x2f, 0x70, 0x32, 0x70,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_p2p_proto_rawDescOnce sync.Once
	file_p2p_proto_rawDescData = file_p2p_proto_rawDesc
)

func file_p2p_proto_rawDescGZIP() []byte {
	file_p2p_proto_rawDescOnce.Do(func() {
		file_p2p_proto_rawDescData = protoimpl.X.CompressGZIP(file_p2p_proto_rawDescData)
	})
	return file_p2p_proto_rawDescData
}

var file_p2p_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_p2p_proto_goTypes = []interface{}{
	(*Envelope)(nil),                // 0: types.Envelope
	(*Packet)(nil),                  // 1: types.Packet
	(*Ping)(nil),                    // 2: types.Ping
	(*Pong)(nil),                    // 3: types.Pong
	(*PeerBookRequestMessage)(nil),  // 4: types.PeerBookRequestMessage
	(*PeerBookResponseMessage)(nil), // 5: types.PeerBookResponseMessage
	(*BookPeer)(nil),                // 6: types.BookPeer
	(*anypb.Any)(nil),               // 7: google.protobuf.Any
	(types.Topic)(0),                // 8: types.Topic
	(*types.PeerAddress)(nil),       // 9: types.PeerAddress
}
var file_p2p_proto_depIdxs = []int32{
	7, // 0: types.Envelope.payload:type_name -> google.protobuf.Any
	8, // 1: types.Packet.stream_id:type_name -> types.Topic
	6, // 2: types.PeerBookResponseMessage.book:type_name -> types.BookPeer
	9, // 3: types.BookPeer.Address:type_name -> types.PeerAddress
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_p2p_proto_init() }
func file_p2p_proto_init() {
	if File_p2p_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_p2p_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Envelope); i {
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
		file_p2p_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Packet); i {
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
		file_p2p_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ping); i {
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
		file_p2p_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pong); i {
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
		file_p2p_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeerBookRequestMessage); i {
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
		file_p2p_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeerBookResponseMessage); i {
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
		file_p2p_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookPeer); i {
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
			RawDescriptor: file_p2p_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_p2p_proto_goTypes,
		DependencyIndexes: file_p2p_proto_depIdxs,
		MessageInfos:      file_p2p_proto_msgTypes,
	}.Build()
	File_p2p_proto = out.File
	file_p2p_proto_rawDesc = nil
	file_p2p_proto_goTypes = nil
	file_p2p_proto_depIdxs = nil
}
