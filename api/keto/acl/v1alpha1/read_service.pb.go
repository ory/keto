// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: keto/acl/v1alpha1/read_service.proto

package acl

import (
	field_mask "google.golang.org/genproto/protobuf/field_mask"
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

// Request for ReadService.ListRelationTuples rpc.
// See ListRelationTuplesRequest_Query for more querying details.
type ListRelationTuplesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// All field constraints are concatenated
	// with a logical AND operator.
	//
	// The RelationTuple list from ListRelationTuplesResponse
	// is ordered from the newest RelationTuple to the oldest.
	Query *ListRelationTuplesRequest_Query `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	// Optional. The list of fields to be expanded
	// in the RelationTuple list returned in `ListRelationTuplesResponse`.
	// Leaving this field unspecified means all fields are expanded.
	//
	// Available fields:
	// "object", "relation", "subject",
	// "namespace", "subject.id", "subject.namespace",
	// "subject.object", "subject.relation"
	ExpandMask *field_mask.FieldMask `protobuf:"bytes,2,opt,name=expand_mask,json=expandMask,proto3" json:"expand_mask,omitempty"`
	// Optional. The snapshot token for this read.
	Snaptoken string `protobuf:"bytes,3,opt,name=snaptoken,proto3" json:"snaptoken,omitempty"`
	// Optional. The maximum number of
	// RelationTuples to return in the response.
	PageSize int32 `protobuf:"varint,4,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Optional. A pagination token returned from
	// a previous call to `ListRelationTuples` that
	// indicates where the page should start at.
	PageToken string `protobuf:"bytes,5,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListRelationTuplesRequest) Reset() {
	*x = ListRelationTuplesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_keto_acl_v1alpha1_read_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRelationTuplesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRelationTuplesRequest) ProtoMessage() {}

func (x *ListRelationTuplesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_keto_acl_v1alpha1_read_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRelationTuplesRequest.ProtoReflect.Descriptor instead.
func (*ListRelationTuplesRequest) Descriptor() ([]byte, []int) {
	return file_keto_acl_v1alpha1_read_service_proto_rawDescGZIP(), []int{0}
}

func (x *ListRelationTuplesRequest) GetQuery() *ListRelationTuplesRequest_Query {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *ListRelationTuplesRequest) GetExpandMask() *field_mask.FieldMask {
	if x != nil {
		return x.ExpandMask
	}
	return nil
}

func (x *ListRelationTuplesRequest) GetSnaptoken() string {
	if x != nil {
		return x.Snaptoken
	}
	return ""
}

func (x *ListRelationTuplesRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListRelationTuplesRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

// The response of a ReadService.ListRelationTuples rpc.
type ListRelationTuplesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The relation tuples matching the list request.
	RelationTuples []*RelationTuple `protobuf:"bytes,1,rep,name=relation_tuples,json=relationTuples,proto3" json:"relation_tuples,omitempty"`
	// Optional. A pagination token returned from a previous call to `ListRelationTuples`
	// that indicates where this listing should continue from.
	//
	// All fields of the subsequent ListRelationTuplesRequest request
	// using this `next_page_token` as the `page_token` are ignored and
	// CAN be left blank, since the request's data is baked in this `next_page_token`.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListRelationTuplesResponse) Reset() {
	*x = ListRelationTuplesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_keto_acl_v1alpha1_read_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRelationTuplesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRelationTuplesResponse) ProtoMessage() {}

func (x *ListRelationTuplesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_keto_acl_v1alpha1_read_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRelationTuplesResponse.ProtoReflect.Descriptor instead.
func (*ListRelationTuplesResponse) Descriptor() ([]byte, []int) {
	return file_keto_acl_v1alpha1_read_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListRelationTuplesResponse) GetRelationTuples() []*RelationTuple {
	if x != nil {
		return x.RelationTuples
	}
	return nil
}

func (x *ListRelationTuplesResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

// The query for listing relation tuples.
// Clients can specify any optional field to
// partially filter for specific relation tuples.
//
// Example use cases:
//  - object only: display a list of all rules of one object
//  - relation only: get all groups that have members; e.g. get all directories that have content
//  - object & relation: display all subjects that have e.g. write relation
//  - subject & relation: display all groups a subject belongs to/display all objects a subject has access to
//  - object & relation & subject: check whether the relation tuple already exists, before writing it
//
type ListRelationTuplesRequest_Query struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The namespace to query.
	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// Optional.
	Object string `protobuf:"bytes,2,opt,name=object,proto3" json:"object,omitempty"`
	// Optional.
	Relation string `protobuf:"bytes,3,opt,name=relation,proto3" json:"relation,omitempty"`
	// Optional.
	Subject *Subject `protobuf:"bytes,4,opt,name=subject,proto3" json:"subject,omitempty"`
}

func (x *ListRelationTuplesRequest_Query) Reset() {
	*x = ListRelationTuplesRequest_Query{}
	if protoimpl.UnsafeEnabled {
		mi := &file_keto_acl_v1alpha1_read_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRelationTuplesRequest_Query) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRelationTuplesRequest_Query) ProtoMessage() {}

func (x *ListRelationTuplesRequest_Query) ProtoReflect() protoreflect.Message {
	mi := &file_keto_acl_v1alpha1_read_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRelationTuplesRequest_Query.ProtoReflect.Descriptor instead.
func (*ListRelationTuplesRequest_Query) Descriptor() ([]byte, []int) {
	return file_keto_acl_v1alpha1_read_service_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ListRelationTuplesRequest_Query) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *ListRelationTuplesRequest_Query) GetObject() string {
	if x != nil {
		return x.Object
	}
	return ""
}

func (x *ListRelationTuplesRequest_Query) GetRelation() string {
	if x != nil {
		return x.Relation
	}
	return ""
}

func (x *ListRelationTuplesRequest_Query) GetSubject() *Subject {
	if x != nil {
		return x.Subject
	}
	return nil
}

var File_keto_acl_v1alpha1_read_service_proto protoreflect.FileDescriptor

var file_keto_acl_v1alpha1_read_service_proto_rawDesc = []byte{
	0x0a, 0x24, 0x6b, 0x65, 0x74, 0x6f, 0x2f, 0x61, 0x63, 0x6c, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2f, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x6b, 0x65, 0x74, 0x6f, 0x2e, 0x61, 0x63, 0x6c,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x1b, 0x6b, 0x65, 0x74, 0x6f, 0x2f,
	0x61, 0x63, 0x6c, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x61, 0x63, 0x6c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61,
	0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8e, 0x03, 0x0a, 0x19, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x75, 0x70, 0x6c, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x48, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x6b, 0x65, 0x74, 0x6f, 0x2e, 0x61, 0x63, 0x6c,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x75, 0x70, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x12, 0x3b, 0x0a, 0x0b, 0x65, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73,
	0x6b, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x12, 0x1c, 0x0a,
	0x09, 0x73, 0x6e, 0x61, 0x70, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x73, 0x6e, 0x61, 0x70, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65,
	0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61,
	0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x1a, 0x8f, 0x01, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x34, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6b, 0x65, 0x74, 0x6f, 0x2e, 0x61, 0x63, 0x6c, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x8f, 0x01, 0x0a, 0x1a, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x75, 0x70, 0x6c, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x0f, 0x72, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x75, 0x70, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x20, 0x2e, 0x6b, 0x65, 0x74, 0x6f, 0x2e, 0x61, 0x63, 0x6c, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x75,
	0x70, 0x6c, 0x65, 0x52, 0x0e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x75, 0x70,
	0x6c, 0x65, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65,
	0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65,
	0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0x80, 0x01, 0x0a, 0x0b,
	0x52, 0x65, 0x61, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x71, 0x0a, 0x12, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x75, 0x70, 0x6c, 0x65,
	0x73, 0x12, 0x2c, 0x2e, 0x6b, 0x65, 0x74, 0x6f, 0x2e, 0x61, 0x63, 0x6c, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x75, 0x70, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2d, 0x2e, 0x6b, 0x65, 0x74, 0x6f, 0x2e, 0x61, 0x63, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x75, 0x70, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x8d,
	0x01, 0x0a, 0x18, 0x73, 0x68, 0x2e, 0x6f, 0x72, 0x79, 0x2e, 0x6b, 0x65, 0x74, 0x6f, 0x2e, 0x61,
	0x63, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x10, 0x52, 0x65, 0x61,
	0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x72, 0x79, 0x2f,
	0x6b, 0x65, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6b, 0x65, 0x74, 0x6f, 0x2f, 0x61, 0x63,
	0x6c, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x61, 0x63, 0x6c, 0xaa, 0x02,
	0x15, 0x4f, 0x72, 0x79, 0x2e, 0x4b, 0x65, 0x74, 0x6f, 0x2e, 0x41, 0x63, 0x6c, 0x2e, 0x56, 0x31,
	0x41, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x15, 0x4f, 0x72, 0x79, 0x5c, 0x4b, 0x65, 0x74,
	0x6f, 0x5c, 0x41, 0x63, 0x6c, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_keto_acl_v1alpha1_read_service_proto_rawDescOnce sync.Once
	file_keto_acl_v1alpha1_read_service_proto_rawDescData = file_keto_acl_v1alpha1_read_service_proto_rawDesc
)

func file_keto_acl_v1alpha1_read_service_proto_rawDescGZIP() []byte {
	file_keto_acl_v1alpha1_read_service_proto_rawDescOnce.Do(func() {
		file_keto_acl_v1alpha1_read_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_keto_acl_v1alpha1_read_service_proto_rawDescData)
	})
	return file_keto_acl_v1alpha1_read_service_proto_rawDescData
}

var file_keto_acl_v1alpha1_read_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_keto_acl_v1alpha1_read_service_proto_goTypes = []interface{}{
	(*ListRelationTuplesRequest)(nil),       // 0: keto.acl.v1alpha1.ListRelationTuplesRequest
	(*ListRelationTuplesResponse)(nil),      // 1: keto.acl.v1alpha1.ListRelationTuplesResponse
	(*ListRelationTuplesRequest_Query)(nil), // 2: keto.acl.v1alpha1.ListRelationTuplesRequest.Query
	(*field_mask.FieldMask)(nil),            // 3: google.protobuf.FieldMask
	(*RelationTuple)(nil),                   // 4: keto.acl.v1alpha1.RelationTuple
	(*Subject)(nil),                         // 5: keto.acl.v1alpha1.Subject
}
var file_keto_acl_v1alpha1_read_service_proto_depIdxs = []int32{
	2, // 0: keto.acl.v1alpha1.ListRelationTuplesRequest.query:type_name -> keto.acl.v1alpha1.ListRelationTuplesRequest.Query
	3, // 1: keto.acl.v1alpha1.ListRelationTuplesRequest.expand_mask:type_name -> google.protobuf.FieldMask
	4, // 2: keto.acl.v1alpha1.ListRelationTuplesResponse.relation_tuples:type_name -> keto.acl.v1alpha1.RelationTuple
	5, // 3: keto.acl.v1alpha1.ListRelationTuplesRequest.Query.subject:type_name -> keto.acl.v1alpha1.Subject
	0, // 4: keto.acl.v1alpha1.ReadService.ListRelationTuples:input_type -> keto.acl.v1alpha1.ListRelationTuplesRequest
	1, // 5: keto.acl.v1alpha1.ReadService.ListRelationTuples:output_type -> keto.acl.v1alpha1.ListRelationTuplesResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_keto_acl_v1alpha1_read_service_proto_init() }
func file_keto_acl_v1alpha1_read_service_proto_init() {
	if File_keto_acl_v1alpha1_read_service_proto != nil {
		return
	}
	file_keto_acl_v1alpha1_acl_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_keto_acl_v1alpha1_read_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRelationTuplesRequest); i {
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
		file_keto_acl_v1alpha1_read_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRelationTuplesResponse); i {
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
		file_keto_acl_v1alpha1_read_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRelationTuplesRequest_Query); i {
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
			RawDescriptor: file_keto_acl_v1alpha1_read_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_keto_acl_v1alpha1_read_service_proto_goTypes,
		DependencyIndexes: file_keto_acl_v1alpha1_read_service_proto_depIdxs,
		MessageInfos:      file_keto_acl_v1alpha1_read_service_proto_msgTypes,
	}.Build()
	File_keto_acl_v1alpha1_read_service_proto = out.File
	file_keto_acl_v1alpha1_read_service_proto_rawDesc = nil
	file_keto_acl_v1alpha1_read_service_proto_goTypes = nil
	file_keto_acl_v1alpha1_read_service_proto_depIdxs = nil
}
