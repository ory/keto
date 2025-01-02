// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        (unknown)
// source: ory/keto/relation_tuples/v1alpha2/namespaces_service.proto

package rts

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Request for ReadService.ListNamespaces RPC.
type ListNamespacesRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListNamespacesRequest) Reset() {
	*x = ListNamespacesRequest{}
	mi := &file_ory_keto_relation_tuples_v1alpha2_namespaces_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListNamespacesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListNamespacesRequest) ProtoMessage() {}

func (x *ListNamespacesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ory_keto_relation_tuples_v1alpha2_namespaces_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListNamespacesRequest.ProtoReflect.Descriptor instead.
func (*ListNamespacesRequest) Descriptor() ([]byte, []int) {
	return file_ory_keto_relation_tuples_v1alpha2_namespaces_service_proto_rawDescGZIP(), []int{0}
}

type ListNamespacesResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Namespaces    []*Namespace           `protobuf:"bytes,1,rep,name=namespaces,proto3" json:"namespaces,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListNamespacesResponse) Reset() {
	*x = ListNamespacesResponse{}
	mi := &file_ory_keto_relation_tuples_v1alpha2_namespaces_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListNamespacesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListNamespacesResponse) ProtoMessage() {}

func (x *ListNamespacesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ory_keto_relation_tuples_v1alpha2_namespaces_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListNamespacesResponse.ProtoReflect.Descriptor instead.
func (*ListNamespacesResponse) Descriptor() ([]byte, []int) {
	return file_ory_keto_relation_tuples_v1alpha2_namespaces_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListNamespacesResponse) GetNamespaces() []*Namespace {
	if x != nil {
		return x.Namespaces
	}
	return nil
}

type Namespace struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Namespace) Reset() {
	*x = Namespace{}
	mi := &file_ory_keto_relation_tuples_v1alpha2_namespaces_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Namespace) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Namespace) ProtoMessage() {}

func (x *Namespace) ProtoReflect() protoreflect.Message {
	mi := &file_ory_keto_relation_tuples_v1alpha2_namespaces_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Namespace.ProtoReflect.Descriptor instead.
func (*Namespace) Descriptor() ([]byte, []int) {
	return file_ory_keto_relation_tuples_v1alpha2_namespaces_service_proto_rawDescGZIP(), []int{2}
}

func (x *Namespace) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_ory_keto_relation_tuples_v1alpha2_namespaces_service_proto protoreflect.FileDescriptor

var file_ory_keto_relation_tuples_v1alpha2_namespaces_service_proto_rawDesc = []byte{
	0x0a, 0x3a, 0x6f, 0x72, 0x79, 0x2f, 0x6b, 0x65, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x75, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x32, 0x2f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x6f, 0x72,
	0x79, 0x2e, 0x6b, 0x65, 0x74, 0x6f, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x74, 0x75, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70,
	0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x17, 0x0a,
	0x15, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x66, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x4c, 0x0a, 0x0a, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x6f, 0x72, 0x79, 0x2e, 0x6b, 0x65, 0x74, 0x6f, 0x2e,
	0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x75, 0x70, 0x6c, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x52, 0x0a, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x22, 0x1f,
	0x0a, 0x09, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x32,
	0xab, 0x03, 0x0a, 0x11, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x95, 0x03, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x12, 0x38, 0x2e, 0x6f, 0x72, 0x79, 0x2e, 0x6b,
	0x65, 0x74, 0x6f, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x75, 0x70,
	0x6c, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x39, 0x2e, 0x6f, 0x72, 0x79, 0x2e, 0x6b, 0x65, 0x74, 0x6f, 0x2e, 0x72, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x75, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x8d, 0x02,
	0x92, 0x41, 0xf6, 0x01, 0x0a, 0x0c, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68,
	0x69, 0x70, 0x2a, 0x1a, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x68, 0x69, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x32, 0x21,
	0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x78, 0x2d, 0x77, 0x77,
	0x77, 0x2d, 0x66, 0x6f, 0x72, 0x6d, 0x2d, 0x75, 0x72, 0x6c, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65,
	0x64, 0x4a, 0xa6, 0x01, 0x0a, 0x03, 0x32, 0x30, 0x30, 0x12, 0x9e, 0x01, 0x0a, 0x17, 0x54, 0x68,
	0x65, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x20, 0x6f, 0x66, 0x20, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x73, 0x2e, 0x12, 0x3d, 0x0a, 0x3b, 0x1a, 0x39, 0x2e, 0x6f, 0x72, 0x79, 0x2e,
	0x6b, 0x65, 0x74, 0x6f, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x75,
	0x70, 0x6c, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x44, 0x0a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x12, 0x30, 0x7b, 0x20, 0x22, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x22, 0x3a, 0x20, 0x5b, 0x20, 0x7b, 0x20, 0x22, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x3a, 0x20, 0x22, 0x6d, 0x79, 0x20, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x22, 0x20, 0x7d, 0x20, 0x5d, 0x20, 0x7d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d,
	0x12, 0x0b, 0x2f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x42, 0xc7, 0x01,
	0x0a, 0x24, 0x73, 0x68, 0x2e, 0x6f, 0x72, 0x79, 0x2e, 0x6b, 0x65, 0x74, 0x6f, 0x2e, 0x72, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x75, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x42, 0x16, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x72, 0x79,
	0x2f, 0x6b, 0x65, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x72, 0x79, 0x2f,
	0x6b, 0x65, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x75,
	0x70, 0x6c, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x3b, 0x72, 0x74,
	0x73, 0xaa, 0x02, 0x20, 0x4f, 0x72, 0x79, 0x2e, 0x4b, 0x65, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x75, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x32, 0xca, 0x02, 0x20, 0x4f, 0x72, 0x79, 0x5c, 0x4b, 0x65, 0x74, 0x6f, 0x5c,
	0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x75, 0x70, 0x6c, 0x65, 0x73, 0x5c, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ory_keto_relation_tuples_v1alpha2_namespaces_service_proto_rawDescOnce sync.Once
	file_ory_keto_relation_tuples_v1alpha2_namespaces_service_proto_rawDescData = file_ory_keto_relation_tuples_v1alpha2_namespaces_service_proto_rawDesc
)

func file_ory_keto_relation_tuples_v1alpha2_namespaces_service_proto_rawDescGZIP() []byte {
	file_ory_keto_relation_tuples_v1alpha2_namespaces_service_proto_rawDescOnce.Do(func() {
		file_ory_keto_relation_tuples_v1alpha2_namespaces_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_ory_keto_relation_tuples_v1alpha2_namespaces_service_proto_rawDescData)
	})
	return file_ory_keto_relation_tuples_v1alpha2_namespaces_service_proto_rawDescData
}

var file_ory_keto_relation_tuples_v1alpha2_namespaces_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_ory_keto_relation_tuples_v1alpha2_namespaces_service_proto_goTypes = []any{
	(*ListNamespacesRequest)(nil),  // 0: ory.keto.relation_tuples.v1alpha2.ListNamespacesRequest
	(*ListNamespacesResponse)(nil), // 1: ory.keto.relation_tuples.v1alpha2.ListNamespacesResponse
	(*Namespace)(nil),              // 2: ory.keto.relation_tuples.v1alpha2.Namespace
}
var file_ory_keto_relation_tuples_v1alpha2_namespaces_service_proto_depIdxs = []int32{
	2, // 0: ory.keto.relation_tuples.v1alpha2.ListNamespacesResponse.namespaces:type_name -> ory.keto.relation_tuples.v1alpha2.Namespace
	0, // 1: ory.keto.relation_tuples.v1alpha2.NamespacesService.ListNamespaces:input_type -> ory.keto.relation_tuples.v1alpha2.ListNamespacesRequest
	1, // 2: ory.keto.relation_tuples.v1alpha2.NamespacesService.ListNamespaces:output_type -> ory.keto.relation_tuples.v1alpha2.ListNamespacesResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ory_keto_relation_tuples_v1alpha2_namespaces_service_proto_init() }
func file_ory_keto_relation_tuples_v1alpha2_namespaces_service_proto_init() {
	if File_ory_keto_relation_tuples_v1alpha2_namespaces_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ory_keto_relation_tuples_v1alpha2_namespaces_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ory_keto_relation_tuples_v1alpha2_namespaces_service_proto_goTypes,
		DependencyIndexes: file_ory_keto_relation_tuples_v1alpha2_namespaces_service_proto_depIdxs,
		MessageInfos:      file_ory_keto_relation_tuples_v1alpha2_namespaces_service_proto_msgTypes,
	}.Build()
	File_ory_keto_relation_tuples_v1alpha2_namespaces_service_proto = out.File
	file_ory_keto_relation_tuples_v1alpha2_namespaces_service_proto_rawDesc = nil
	file_ory_keto_relation_tuples_v1alpha2_namespaces_service_proto_goTypes = nil
	file_ory_keto_relation_tuples_v1alpha2_namespaces_service_proto_depIdxs = nil
}
