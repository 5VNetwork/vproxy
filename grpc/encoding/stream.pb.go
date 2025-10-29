package encoding

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Hunk struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Data          []byte                 `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Hunk) Reset() {
	*x = Hunk{}
	mi := &file_transport_protocols_grpc_encoding_stream_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Hunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hunk) ProtoMessage() {}

func (x *Hunk) ProtoReflect() protoreflect.Message {
	mi := &file_transport_protocols_grpc_encoding_stream_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hunk.ProtoReflect.Descriptor instead.
func (*Hunk) Descriptor() ([]byte, []int) {
	return file_transport_protocols_grpc_encoding_stream_proto_rawDescGZIP(), []int{0}
}

func (x *Hunk) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type MultiHunk struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Data          [][]byte               `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MultiHunk) Reset() {
	*x = MultiHunk{}
	mi := &file_transport_protocols_grpc_encoding_stream_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MultiHunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultiHunk) ProtoMessage() {}

func (x *MultiHunk) ProtoReflect() protoreflect.Message {
	mi := &file_transport_protocols_grpc_encoding_stream_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultiHunk.ProtoReflect.Descriptor instead.
func (*MultiHunk) Descriptor() ([]byte, []int) {
	return file_transport_protocols_grpc_encoding_stream_proto_rawDescGZIP(), []int{1}
}

func (x *MultiHunk) GetData() [][]byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_transport_protocols_grpc_encoding_stream_proto protoreflect.FileDescriptor

const file_transport_protocols_grpc_encoding_stream_proto_rawDesc = "" +
	"\n" +
	".transport/protocols/grpc/encoding/stream.proto\x12#x.transport.protocols.grpc.encoding\"\x1a\n" +
	"\x04Hunk\x12\x12\n" +
	"\x04data\x18\x01 \x01(\fR\x04data\"\x1f\n" +
	"\tMultiHunk\x12\x12\n" +
	"\x04data\x18\x01 \x03(\fR\x04data2\xde\x01\n" +
	"\vGRPCService\x12_\n" +
	"\x03Tun\x12).x.transport.protocols.grpc.encoding.Hunk\x1a).x.transport.protocols.grpc.encoding.Hunk(\x010\x01\x12n\n" +
	"\bTunMulti\x12..x.transport.protocols.grpc.encoding.MultiHunk\x1a..x.transport.protocols.grpc.encoding.MultiHunk(\x010\x01B:Z8github.com/5vnetwork/x/transport/protocols/grpc/encodingb\x06proto3"

var (
	file_transport_protocols_grpc_encoding_stream_proto_rawDescOnce sync.Once
	file_transport_protocols_grpc_encoding_stream_proto_rawDescData []byte
)

func file_transport_protocols_grpc_encoding_stream_proto_rawDescGZIP() []byte {
	file_transport_protocols_grpc_encoding_stream_proto_rawDescOnce.Do(func() {
		file_transport_protocols_grpc_encoding_stream_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_transport_protocols_grpc_encoding_stream_proto_rawDesc), len(file_transport_protocols_grpc_encoding_stream_proto_rawDesc)))
	})
	return file_transport_protocols_grpc_encoding_stream_proto_rawDescData
}

var file_transport_protocols_grpc_encoding_stream_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_transport_protocols_grpc_encoding_stream_proto_goTypes = []any{
	(*Hunk)(nil),      // 0: x.transport.protocols.grpc.encoding.Hunk
	(*MultiHunk)(nil), // 1: x.transport.protocols.grpc.encoding.MultiHunk
}
var file_transport_protocols_grpc_encoding_stream_proto_depIdxs = []int32{
	0, // 0: x.transport.protocols.grpc.encoding.GRPCService.Tun:input_type -> x.transport.protocols.grpc.encoding.Hunk
	1, // 1: x.transport.protocols.grpc.encoding.GRPCService.TunMulti:input_type -> x.transport.protocols.grpc.encoding.MultiHunk
	0, // 2: x.transport.protocols.grpc.encoding.GRPCService.Tun:output_type -> x.transport.protocols.grpc.encoding.Hunk
	1, // 3: x.transport.protocols.grpc.encoding.GRPCService.TunMulti:output_type -> x.transport.protocols.grpc.encoding.MultiHunk
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_transport_protocols_grpc_encoding_stream_proto_init() }
func file_transport_protocols_grpc_encoding_stream_proto_init() {
	if File_transport_protocols_grpc_encoding_stream_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_transport_protocols_grpc_encoding_stream_proto_rawDesc), len(file_transport_protocols_grpc_encoding_stream_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_transport_protocols_grpc_encoding_stream_proto_goTypes,
		DependencyIndexes: file_transport_protocols_grpc_encoding_stream_proto_depIdxs,
		MessageInfos:      file_transport_protocols_grpc_encoding_stream_proto_msgTypes,
	}.Build()
	File_transport_protocols_grpc_encoding_stream_proto = out.File
	file_transport_protocols_grpc_encoding_stream_proto_goTypes = nil
	file_transport_protocols_grpc_encoding_stream_proto_depIdxs = nil
}
