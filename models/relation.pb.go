// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: models/relation.proto

package models

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Query by partial keys of tuples
type ReadRelationTuplesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// All conditions are concatenated with
	// an OR operator. If any query matches
	// the relation will be in the response
	TupleSets []*ReadRelationTuplesRequest_Query `protobuf:"bytes,1,rep,name=tuple_sets,json=tupleSets,proto3" json:"tuple_sets,omitempty"`
	Page      int32                              `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	PerPage   int32                              `protobuf:"varint,3,opt,name=per_page,json=perPage,proto3" json:"per_page,omitempty"`
}

func (x *ReadRelationTuplesRequest) Reset() {
	*x = ReadRelationTuplesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_relation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadRelationTuplesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadRelationTuplesRequest) ProtoMessage() {}

func (x *ReadRelationTuplesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_models_relation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadRelationTuplesRequest.ProtoReflect.Descriptor instead.
func (*ReadRelationTuplesRequest) Descriptor() ([]byte, []int) {
	return file_models_relation_proto_rawDescGZIP(), []int{0}
}

func (x *ReadRelationTuplesRequest) GetTupleSets() []*ReadRelationTuplesRequest_Query {
	if x != nil {
		return x.TupleSets
	}
	return nil
}

func (x *ReadRelationTuplesRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ReadRelationTuplesRequest) GetPerPage() int32 {
	if x != nil {
		return x.PerPage
	}
	return 0
}

type ReadRelationTuplesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tuples []*RelationTuple `protobuf:"bytes,1,rep,name=tuples,proto3" json:"tuples,omitempty"`
}

func (x *ReadRelationTuplesResponse) Reset() {
	*x = ReadRelationTuplesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_relation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadRelationTuplesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadRelationTuplesResponse) ProtoMessage() {}

func (x *ReadRelationTuplesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_models_relation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadRelationTuplesResponse.ProtoReflect.Descriptor instead.
func (*ReadRelationTuplesResponse) Descriptor() ([]byte, []int) {
	return file_models_relation_proto_rawDescGZIP(), []int{1}
}

func (x *ReadRelationTuplesResponse) GetTuples() []*RelationTuple {
	if x != nil {
		return x.Tuples
	}
	return nil
}

type WriteRelationTupleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tuple *RelationTuple `protobuf:"bytes,1,opt,name=tuple,proto3" json:"tuple,omitempty"`
}

func (x *WriteRelationTupleRequest) Reset() {
	*x = WriteRelationTupleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_relation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteRelationTupleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteRelationTupleRequest) ProtoMessage() {}

func (x *WriteRelationTupleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_models_relation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteRelationTupleRequest.ProtoReflect.Descriptor instead.
func (*WriteRelationTupleRequest) Descriptor() ([]byte, []int) {
	return file_models_relation_proto_rawDescGZIP(), []int{2}
}

func (x *WriteRelationTupleRequest) GetTuple() *RelationTuple {
	if x != nil {
		return x.Tuple
	}
	return nil
}

type WriteRelationTupleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *WriteRelationTupleResponse) Reset() {
	*x = WriteRelationTupleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_relation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteRelationTupleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteRelationTupleResponse) ProtoMessage() {}

func (x *WriteRelationTupleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_models_relation_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteRelationTupleResponse.ProtoReflect.Descriptor instead.
func (*WriteRelationTupleResponse) Descriptor() ([]byte, []int) {
	return file_models_relation_proto_rawDescGZIP(), []int{3}
}

// Represents a relation between
// an object and an user.
type RelationTuple struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Object *RelationObject `protobuf:"bytes,1,opt,name=object,proto3" json:"object,omitempty"`
	// Specifies the relation between an `object` and an `user`
	Relation string `protobuf:"bytes,2,opt,name=relation,proto3" json:"relation,omitempty"`
	// The user of the tuple can either be
	// a single user or a userset
	//
	// Types that are assignable to Subject:
	//	*RelationTuple_UserId
	//	*RelationTuple_UserSet
	Subject isRelationTuple_Subject `protobuf_oneof:"subject"`
}

func (x *RelationTuple) Reset() {
	*x = RelationTuple{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_relation_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelationTuple) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelationTuple) ProtoMessage() {}

func (x *RelationTuple) ProtoReflect() protoreflect.Message {
	mi := &file_models_relation_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelationTuple.ProtoReflect.Descriptor instead.
func (*RelationTuple) Descriptor() ([]byte, []int) {
	return file_models_relation_proto_rawDescGZIP(), []int{4}
}

func (x *RelationTuple) GetObject() *RelationObject {
	if x != nil {
		return x.Object
	}
	return nil
}

func (x *RelationTuple) GetRelation() string {
	if x != nil {
		return x.Relation
	}
	return ""
}

func (m *RelationTuple) GetSubject() isRelationTuple_Subject {
	if m != nil {
		return m.Subject
	}
	return nil
}

func (x *RelationTuple) GetUserId() string {
	if x, ok := x.GetSubject().(*RelationTuple_UserId); ok {
		return x.UserId
	}
	return ""
}

func (x *RelationTuple) GetUserSet() *RelationUserSet {
	if x, ok := x.GetSubject().(*RelationTuple_UserSet); ok {
		return x.UserSet
	}
	return nil
}

type isRelationTuple_Subject interface {
	isRelationTuple_Subject()
}

type RelationTuple_UserId struct {
	UserId string `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3,oneof"`
}

type RelationTuple_UserSet struct {
	UserSet *RelationUserSet `protobuf:"bytes,4,opt,name=user_set,json=userSet,proto3,oneof"`
}

func (*RelationTuple_UserId) isRelationTuple_Subject() {}

func (*RelationTuple_UserSet) isRelationTuple_Subject() {}

// Refers to all users which have
// a `relation` with an `object`
type RelationUserSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Object   *RelationObject `protobuf:"bytes,1,opt,name=object,proto3" json:"object,omitempty"`
	Relation string          `protobuf:"bytes,2,opt,name=relation,proto3" json:"relation,omitempty"`
}

func (x *RelationUserSet) Reset() {
	*x = RelationUserSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_relation_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelationUserSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelationUserSet) ProtoMessage() {}

func (x *RelationUserSet) ProtoReflect() protoreflect.Message {
	mi := &file_models_relation_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelationUserSet.ProtoReflect.Descriptor instead.
func (*RelationUserSet) Descriptor() ([]byte, []int) {
	return file_models_relation_proto_rawDescGZIP(), []int{5}
}

func (x *RelationUserSet) GetObject() *RelationObject {
	if x != nil {
		return x.Object
	}
	return nil
}

func (x *RelationUserSet) GetRelation() string {
	if x != nil {
		return x.Relation
	}
	return ""
}

// Represents a "resource"
type RelationObject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	ObjectId  string `protobuf:"bytes,2,opt,name=object_id,json=objectId,proto3" json:"object_id,omitempty"`
}

func (x *RelationObject) Reset() {
	*x = RelationObject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_relation_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelationObject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelationObject) ProtoMessage() {}

func (x *RelationObject) ProtoReflect() protoreflect.Message {
	mi := &file_models_relation_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelationObject.ProtoReflect.Descriptor instead.
func (*RelationObject) Descriptor() ([]byte, []int) {
	return file_models_relation_proto_rawDescGZIP(), []int{6}
}

func (x *RelationObject) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *RelationObject) GetObjectId() string {
	if x != nil {
		return x.ObjectId
	}
	return ""
}

type ReadRelationTuplesRequest_Query struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional
	Object *RelationObject `protobuf:"bytes,1,opt,name=object,proto3" json:"object,omitempty"`
	// Optional
	Relation string `protobuf:"bytes,2,opt,name=relation,proto3" json:"relation,omitempty"`
	// Optional
	//
	// Types that are assignable to Subject:
	//	*ReadRelationTuplesRequest_Query_UserId
	//	*ReadRelationTuplesRequest_Query_UserSet
	Subject isReadRelationTuplesRequest_Query_Subject `protobuf_oneof:"subject"`
}

func (x *ReadRelationTuplesRequest_Query) Reset() {
	*x = ReadRelationTuplesRequest_Query{}
	if protoimpl.UnsafeEnabled {
		mi := &file_models_relation_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadRelationTuplesRequest_Query) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadRelationTuplesRequest_Query) ProtoMessage() {}

func (x *ReadRelationTuplesRequest_Query) ProtoReflect() protoreflect.Message {
	mi := &file_models_relation_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadRelationTuplesRequest_Query.ProtoReflect.Descriptor instead.
func (*ReadRelationTuplesRequest_Query) Descriptor() ([]byte, []int) {
	return file_models_relation_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ReadRelationTuplesRequest_Query) GetObject() *RelationObject {
	if x != nil {
		return x.Object
	}
	return nil
}

func (x *ReadRelationTuplesRequest_Query) GetRelation() string {
	if x != nil {
		return x.Relation
	}
	return ""
}

func (m *ReadRelationTuplesRequest_Query) GetSubject() isReadRelationTuplesRequest_Query_Subject {
	if m != nil {
		return m.Subject
	}
	return nil
}

func (x *ReadRelationTuplesRequest_Query) GetUserId() string {
	if x, ok := x.GetSubject().(*ReadRelationTuplesRequest_Query_UserId); ok {
		return x.UserId
	}
	return ""
}

func (x *ReadRelationTuplesRequest_Query) GetUserSet() *RelationUserSet {
	if x, ok := x.GetSubject().(*ReadRelationTuplesRequest_Query_UserSet); ok {
		return x.UserSet
	}
	return nil
}

type isReadRelationTuplesRequest_Query_Subject interface {
	isReadRelationTuplesRequest_Query_Subject()
}

type ReadRelationTuplesRequest_Query_UserId struct {
	UserId string `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3,oneof"`
}

type ReadRelationTuplesRequest_Query_UserSet struct {
	UserSet *RelationUserSet `protobuf:"bytes,4,opt,name=user_set,json=userSet,proto3,oneof"`
}

func (*ReadRelationTuplesRequest_Query_UserId) isReadRelationTuplesRequest_Query_Subject() {}

func (*ReadRelationTuplesRequest_Query_UserSet) isReadRelationTuplesRequest_Query_Subject() {}

var File_models_relation_proto protoreflect.FileDescriptor

var file_models_relation_proto_rawDesc = []byte{
	0x0a, 0x15, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x22,
	0xc4, 0x02, 0x0a, 0x19, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x75, 0x70, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x46, 0x0a,
	0x0a, 0x74, 0x75, 0x70, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x27, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x52,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x75, 0x70, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x09, 0x74, 0x75, 0x70, 0x6c,
	0x65, 0x53, 0x65, 0x74, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x65, 0x72,
	0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x70, 0x65, 0x72,
	0x50, 0x61, 0x67, 0x65, 0x1a, 0xaf, 0x01, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x2e,
	0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x34, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73,
	0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x74,
	0x48, 0x00, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x53, 0x65, 0x74, 0x42, 0x09, 0x0a, 0x07, 0x73,
	0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x4b, 0x0a, 0x1a, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x75, 0x70, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x06, 0x74, 0x75, 0x70, 0x6c, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x52, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x75, 0x70, 0x6c, 0x65, 0x52, 0x06, 0x74, 0x75, 0x70,
	0x6c, 0x65, 0x73, 0x22, 0x48, 0x0a, 0x19, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x75, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x2b, 0x0a, 0x05, 0x74, 0x75, 0x70, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x75, 0x70, 0x6c, 0x65, 0x52, 0x05, 0x74, 0x75, 0x70, 0x6c, 0x65, 0x22, 0x1c, 0x0a,
	0x1a, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x75,
	0x70, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xb7, 0x01, 0x0a, 0x0d,
	0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x75, 0x70, 0x6c, 0x65, 0x12, 0x2e, 0x0a,
	0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x34, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e,
	0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x74, 0x48,
	0x00, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x53, 0x65, 0x74, 0x42, 0x09, 0x0a, 0x07, 0x73, 0x75,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x5d, 0x0a, 0x0f, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x74, 0x12, 0x2e, 0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x73, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x52, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x4b, 0x0a, 0x0e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49,
	0x64, 0x32, 0xd0, 0x01, 0x0a, 0x14, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x75,
	0x70, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5b, 0x0a, 0x12, 0x52, 0x65,
	0x61, 0x64, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x75, 0x70, 0x6c, 0x65, 0x73,
	0x12, 0x21, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x75, 0x70, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x52, 0x65, 0x61,
	0x64, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x75, 0x70, 0x6c, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5b, 0x0a, 0x12, 0x57, 0x72, 0x69, 0x74, 0x65,
	0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x75, 0x70, 0x6c, 0x65, 0x12, 0x21, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x75, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x22, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x75, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1c, 0x5a, 0x1a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6f, 0x72, 0x79, 0x2f, 0x6b, 0x65, 0x74, 0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_models_relation_proto_rawDescOnce sync.Once
	file_models_relation_proto_rawDescData = file_models_relation_proto_rawDesc
)

func file_models_relation_proto_rawDescGZIP() []byte {
	file_models_relation_proto_rawDescOnce.Do(func() {
		file_models_relation_proto_rawDescData = protoimpl.X.CompressGZIP(file_models_relation_proto_rawDescData)
	})
	return file_models_relation_proto_rawDescData
}

var file_models_relation_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_models_relation_proto_goTypes = []interface{}{
	(*ReadRelationTuplesRequest)(nil),       // 0: models.ReadRelationTuplesRequest
	(*ReadRelationTuplesResponse)(nil),      // 1: models.ReadRelationTuplesResponse
	(*WriteRelationTupleRequest)(nil),       // 2: models.WriteRelationTupleRequest
	(*WriteRelationTupleResponse)(nil),      // 3: models.WriteRelationTupleResponse
	(*RelationTuple)(nil),                   // 4: models.RelationTuple
	(*RelationUserSet)(nil),                 // 5: models.RelationUserSet
	(*RelationObject)(nil),                  // 6: models.RelationObject
	(*ReadRelationTuplesRequest_Query)(nil), // 7: models.ReadRelationTuplesRequest.Query
}
var file_models_relation_proto_depIdxs = []int32{
	7,  // 0: models.ReadRelationTuplesRequest.tuple_sets:type_name -> models.ReadRelationTuplesRequest.Query
	4,  // 1: models.ReadRelationTuplesResponse.tuples:type_name -> models.RelationTuple
	4,  // 2: models.WriteRelationTupleRequest.tuple:type_name -> models.RelationTuple
	6,  // 3: models.RelationTuple.object:type_name -> models.RelationObject
	5,  // 4: models.RelationTuple.user_set:type_name -> models.RelationUserSet
	6,  // 5: models.RelationUserSet.object:type_name -> models.RelationObject
	6,  // 6: models.ReadRelationTuplesRequest.Query.object:type_name -> models.RelationObject
	5,  // 7: models.ReadRelationTuplesRequest.Query.user_set:type_name -> models.RelationUserSet
	0,  // 8: models.RelationTupleService.ReadRelationTuples:input_type -> models.ReadRelationTuplesRequest
	2,  // 9: models.RelationTupleService.WriteRelationTuple:input_type -> models.WriteRelationTupleRequest
	1,  // 10: models.RelationTupleService.ReadRelationTuples:output_type -> models.ReadRelationTuplesResponse
	3,  // 11: models.RelationTupleService.WriteRelationTuple:output_type -> models.WriteRelationTupleResponse
	10, // [10:12] is the sub-list for method output_type
	8,  // [8:10] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_models_relation_proto_init() }
func file_models_relation_proto_init() {
	if File_models_relation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_models_relation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadRelationTuplesRequest); i {
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
		file_models_relation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadRelationTuplesResponse); i {
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
		file_models_relation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteRelationTupleRequest); i {
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
		file_models_relation_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteRelationTupleResponse); i {
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
		file_models_relation_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelationTuple); i {
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
		file_models_relation_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelationUserSet); i {
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
		file_models_relation_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelationObject); i {
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
		file_models_relation_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadRelationTuplesRequest_Query); i {
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
	file_models_relation_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*RelationTuple_UserId)(nil),
		(*RelationTuple_UserSet)(nil),
	}
	file_models_relation_proto_msgTypes[7].OneofWrappers = []interface{}{
		(*ReadRelationTuplesRequest_Query_UserId)(nil),
		(*ReadRelationTuplesRequest_Query_UserSet)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_models_relation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_models_relation_proto_goTypes,
		DependencyIndexes: file_models_relation_proto_depIdxs,
		MessageInfos:      file_models_relation_proto_msgTypes,
	}.Build()
	File_models_relation_proto = out.File
	file_models_relation_proto_rawDesc = nil
	file_models_relation_proto_goTypes = nil
	file_models_relation_proto_depIdxs = nil
}
