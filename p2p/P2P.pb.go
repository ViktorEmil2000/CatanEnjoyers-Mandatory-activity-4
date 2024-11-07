// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.12.4
// source: p2p/P2P.proto

package p2p

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

type FromClient struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FromClient) Reset() {
	*x = FromClient{}
	mi := &file_p2p_P2P_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FromClient) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FromClient) ProtoMessage() {}

func (x *FromClient) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_P2P_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FromClient.ProtoReflect.Descriptor instead.
func (*FromClient) Descriptor() ([]byte, []int) {
	return file_p2p_P2P_proto_rawDescGZIP(), []int{0}
}

type FromServer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Request  string `protobuf:"bytes,2,opt,name=request,proto3" json:"request,omitempty"`
	Response string `protobuf:"bytes,3,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *FromServer) Reset() {
	*x = FromServer{}
	mi := &file_p2p_P2P_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FromServer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FromServer) ProtoMessage() {}

func (x *FromServer) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_P2P_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FromServer.ProtoReflect.Descriptor instead.
func (*FromServer) Descriptor() ([]byte, []int) {
	return file_p2p_P2P_proto_rawDescGZIP(), []int{1}
}

func (x *FromServer) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FromServer) GetRequest() string {
	if x != nil {
		return x.Request
	}
	return ""
}

func (x *FromServer) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

type ResquestFromClient struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	IsResponse bool  `protobuf:"varint,2,opt,name=isResponse,proto3" json:"isResponse,omitempty"`
	Response   bool  `protobuf:"varint,3,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *ResquestFromClient) Reset() {
	*x = ResquestFromClient{}
	mi := &file_p2p_P2P_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResquestFromClient) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResquestFromClient) ProtoMessage() {}

func (x *ResquestFromClient) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_P2P_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResquestFromClient.ProtoReflect.Descriptor instead.
func (*ResquestFromClient) Descriptor() ([]byte, []int) {
	return file_p2p_P2P_proto_rawDescGZIP(), []int{2}
}

func (x *ResquestFromClient) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ResquestFromClient) GetIsResponse() bool {
	if x != nil {
		return x.IsResponse
	}
	return false
}

func (x *ResquestFromClient) GetResponse() bool {
	if x != nil {
		return x.Response
	}
	return false
}

type ResponseFromServer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ResponseFromServer) Reset() {
	*x = ResponseFromServer{}
	mi := &file_p2p_P2P_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResponseFromServer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseFromServer) ProtoMessage() {}

func (x *ResponseFromServer) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_P2P_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseFromServer.ProtoReflect.Descriptor instead.
func (*ResponseFromServer) Descriptor() ([]byte, []int) {
	return file_p2p_P2P_proto_rawDescGZIP(), []int{3}
}

var File_p2p_P2P_proto protoreflect.FileDescriptor

var file_p2p_P2P_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x32, 0x70, 0x2f, 0x50, 0x32, 0x50, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x50, 0x32, 0x50, 0x22, 0x0c, 0x0a, 0x0a, 0x46, 0x72, 0x6f, 0x6d, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x22, 0x52, 0x0a, 0x0a, 0x46, 0x72, 0x6f, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x60, 0x0a, 0x12, 0x52, 0x65, 0x73, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x46, 0x72, 0x6f, 0x6d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a,
	0x69, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0a, 0x69, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14, 0x0a, 0x12, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x32, 0x88,
	0x01, 0x0a, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x37, 0x0a, 0x0d, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x0f, 0x2e, 0x50, 0x32, 0x50,
	0x2e, 0x46, 0x72, 0x6f, 0x6d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x1a, 0x0f, 0x2e, 0x50, 0x32,
	0x50, 0x2e, 0x46, 0x72, 0x6f, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x22, 0x00, 0x28, 0x01,
	0x30, 0x01, 0x12, 0x45, 0x0a, 0x0f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x2e, 0x50, 0x32, 0x50, 0x2e, 0x52, 0x65, 0x73, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x46, 0x72, 0x6f, 0x6d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x1a, 0x17,
	0x2e, 0x50, 0x32, 0x50, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x46, 0x72, 0x6f,
	0x6d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x22, 0x00, 0x42, 0x42, 0x5a, 0x40, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x56, 0x69, 0x6b, 0x74, 0x6f, 0x72, 0x45, 0x6d,
	0x69, 0x6c, 0x32, 0x30, 0x30, 0x30, 0x2f, 0x43, 0x61, 0x74, 0x61, 0x6e, 0x45, 0x6e, 0x6a, 0x6f,
	0x79, 0x65, 0x72, 0x73, 0x2d, 0x4d, 0x61, 0x6e, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x79, 0x2d, 0x61,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2d, 0x34, 0x2f, 0x70, 0x32, 0x70, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_p2p_P2P_proto_rawDescOnce sync.Once
	file_p2p_P2P_proto_rawDescData = file_p2p_P2P_proto_rawDesc
)

func file_p2p_P2P_proto_rawDescGZIP() []byte {
	file_p2p_P2P_proto_rawDescOnce.Do(func() {
		file_p2p_P2P_proto_rawDescData = protoimpl.X.CompressGZIP(file_p2p_P2P_proto_rawDescData)
	})
	return file_p2p_P2P_proto_rawDescData
}

var file_p2p_P2P_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_p2p_P2P_proto_goTypes = []any{
	(*FromClient)(nil),         // 0: P2P.FromClient
	(*FromServer)(nil),         // 1: P2P.FromServer
	(*ResquestFromClient)(nil), // 2: P2P.ResquestFromClient
	(*ResponseFromServer)(nil), // 3: P2P.responseFromServer
}
var file_p2p_P2P_proto_depIdxs = []int32{
	0, // 0: P2P.client.ClientConnect:input_type -> P2P.FromClient
	2, // 1: P2P.client.RequestResponse:input_type -> P2P.ResquestFromClient
	1, // 2: P2P.client.ClientConnect:output_type -> P2P.FromServer
	3, // 3: P2P.client.RequestResponse:output_type -> P2P.responseFromServer
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_p2p_P2P_proto_init() }
func file_p2p_P2P_proto_init() {
	if File_p2p_P2P_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_p2p_P2P_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_p2p_P2P_proto_goTypes,
		DependencyIndexes: file_p2p_P2P_proto_depIdxs,
		MessageInfos:      file_p2p_P2P_proto_msgTypes,
	}.Build()
	File_p2p_P2P_proto = out.File
	file_p2p_P2P_proto_rawDesc = nil
	file_p2p_P2P_proto_goTypes = nil
	file_p2p_P2P_proto_depIdxs = nil
}
