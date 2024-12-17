// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        (unknown)
// source: ory/keto/relation_tuples/v1alpha2/read_service.proto

package rts

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "google.golang.org/genproto/googleapis/api/visibility"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Request for ReadService.ListRelationTuples RPC.
// See `ListRelationTuplesRequest_Query` for how to filter the query.
type ListRelationTuplesRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// All query constraints are concatenated
	// with a logical AND operator.
	//
	// The RelationTuple list from ListRelationTuplesResponse
	// is ordered from the newest RelationTuple to the oldest.
	//
	// Deprecated: Marked as deprecated in ory/keto/relation_tuples/v1alpha2/read_service.proto.
	Query         *ListRelationTuplesRequest_Query `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	RelationQuery *RelationQuery                   `protobuf:"bytes,6,opt,name=relation_query,json=relationQuery,proto3" json:"relation_query,omitempty"`
	// This field is not implemented yet and has no effect.
	// <!--
	// Optional. The list of fields to be expanded
	// in the RelationTuple list returned in `ListRelationTuplesResponse`.
	// Leaving this field unspecified means all fields are expanded.
	//
	// Available fields:
	// "object", "relation", "subject",
	// "namespace", "subject.id", "subject.namespace",
	// "subject.object", "subject.relation"
	// -->
	ExpandMask *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=expand_mask,json=expandMask,proto3" json:"expand_mask,omitempty"`
	// This field is not implemented yet and has no effect.
	// <!--
	// Optional. The snapshot token for this read.
	// -->
	Snaptoken string `protobuf:"bytes,3,opt,name=snaptoken,proto3" json:"snaptoken,omitempty"`
	// Optional. The maximum number of
	// RelationTuples to return in the response.
	//
	// Default: 100
	PageSize int32 `protobuf:"varint,4,opt,name=page_size,proto3" json:"page_size,omitempty"`
	// Optional. An opaque pagination token returned from
	// a previous call to `ListRelationTuples` that
	// indicates where the page should start at.
	//
	// An empty token denotes the first page. All successive
	// pages require the token from the previous page.
	PageToken string `protobuf:"bytes,5,opt,name=page_token,proto3" json:"page_token,omitempty"`
	// The namespace
	//
	// Deprecated: Marked as deprecated in ory/keto/relation_tuples/v1alpha2/read_service.proto.
	Namespace string `protobuf:"bytes,7,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// The related object in this check.
	//
	// Deprecated: Marked as deprecated in ory/keto/relation_tuples/v1alpha2/read_service.proto.
	Object string `protobuf:"bytes,8,opt,name=object,proto3" json:"object,omitempty"`
	// The relation between the Object and the Subject.
	//
	// Deprecated: Marked as deprecated in ory/keto/relation_tuples/v1alpha2/read_service.proto.
	Relation string `protobuf:"bytes,9,opt,name=relation,proto3" json:"relation,omitempty"`
	// Types that are valid to be assigned to RestApiSubject:
	//
	//	*ListRelationTuplesRequest_SubjectId
	//	*ListRelationTuplesRequest_SubjectSet
	RestApiSubject isListRelationTuplesRequest_RestApiSubject `protobuf_oneof:"rest_api_subject"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *ListRelationTuplesRequest) Reset() {
	*x = ListRelationTuplesRequest{}
	mi := &file_ory_keto_relation_tuples_v1alpha2_read_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListRelationTuplesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRelationTuplesRequest) ProtoMessage() {}

func (x *ListRelationTuplesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ory_keto_relation_tuples_v1alpha2_read_service_proto_msgTypes[0]
	if x != nil {
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
	return file_ory_keto_relation_tuples_v1alpha2_read_service_proto_rawDescGZIP(), []int{0}
}

// Deprecated: Marked as deprecated in ory/keto/relation_tuples/v1alpha2/read_service.proto.
func (x *ListRelationTuplesRequest) GetQuery() *ListRelationTuplesRequest_Query {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *ListRelationTuplesRequest) GetRelationQuery() *RelationQuery {
	if x != nil {
		return x.RelationQuery
	}
	return nil
}

func (x *ListRelationTuplesRequest) GetExpandMask() *fieldmaskpb.FieldMask {
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

// Deprecated: Marked as deprecated in ory/keto/relation_tuples/v1alpha2/read_service.proto.
func (x *ListRelationTuplesRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

// Deprecated: Marked as deprecated in ory/keto/relation_tuples/v1alpha2/read_service.proto.
func (x *ListRelationTuplesRequest) GetObject() string {
	if x != nil {
		return x.Object
	}
	return ""
}

// Deprecated: Marked as deprecated in ory/keto/relation_tuples/v1alpha2/read_service.proto.
func (x *ListRelationTuplesRequest) GetRelation() string {
	if x != nil {
		return x.Relation
	}
	return ""
}

func (x *ListRelationTuplesRequest) GetRestApiSubject() isListRelationTuplesRequest_RestApiSubject {
	if x != nil {
		return x.RestApiSubject
	}
	return nil
}

func (x *ListRelationTuplesRequest) GetSubjectId() string {
	if x != nil {
		if x, ok := x.RestApiSubject.(*ListRelationTuplesRequest_SubjectId); ok {
			return x.SubjectId
		}
	}
	return ""
}

func (x *ListRelationTuplesRequest) GetSubjectSet() *SubjectSetQuery {
	if x != nil {
		if x, ok := x.RestApiSubject.(*ListRelationTuplesRequest_SubjectSet); ok {
			return x.SubjectSet
		}
	}
	return nil
}

type isListRelationTuplesRequest_RestApiSubject interface {
	isListRelationTuplesRequest_RestApiSubject()
}

type ListRelationTuplesRequest_SubjectId struct {
	// A concrete id of the subject.
	SubjectId string `protobuf:"bytes,10,opt,name=subject_id,proto3,oneof"`
}

type ListRelationTuplesRequest_SubjectSet struct {
	// A subject set that expands to more Subjects.
	// More information are available under [concepts](../concepts/subjects.mdx).
	SubjectSet *SubjectSetQuery `protobuf:"bytes,11,opt,name=subject_set,proto3,oneof"`
}

func (*ListRelationTuplesRequest_SubjectId) isListRelationTuplesRequest_RestApiSubject() {}

func (*ListRelationTuplesRequest_SubjectSet) isListRelationTuplesRequest_RestApiSubject() {}

// The response of a ReadService.ListRelationTuples RPC.
type ListRelationTuplesResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The relationships matching the list request.
	RelationTuples []*RelationTuple `protobuf:"bytes,1,rep,name=relation_tuples,proto3" json:"relation_tuples,omitempty"`
	// The token required to get the next page.
	// If this is the last page, the token will be the empty string.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,proto3" json:"next_page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListRelationTuplesResponse) Reset() {
	*x = ListRelationTuplesResponse{}
	mi := &file_ory_keto_relation_tuples_v1alpha2_read_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListRelationTuplesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRelationTuplesResponse) ProtoMessage() {}

func (x *ListRelationTuplesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ory_keto_relation_tuples_v1alpha2_read_service_proto_msgTypes[1]
	if x != nil {
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
	return file_ory_keto_relation_tuples_v1alpha2_read_service_proto_rawDescGZIP(), []int{1}
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

// The query for listing relationships.
// Clients can specify any optional field to
// partially filter for specific relationships.
//
// Example use cases (namespace is always required):
//   - object only: display a list of all permissions referring to a specific object
//   - relation only: get all groups that have members; get all directories that have content
//   - object & relation: display all subjects that have a specific permission relation
//   - subject & relation: display all groups a subject belongs to; display all objects a subject has access to
//   - object & relation & subject: check whether the relation tuple already exists
type ListRelationTuplesRequest_Query struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Required. The namespace to query.
	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// Optional. The object to query for.
	Object string `protobuf:"bytes,2,opt,name=object,proto3" json:"object,omitempty"`
	// Optional. The relation to query for.
	Relation string `protobuf:"bytes,3,opt,name=relation,proto3" json:"relation,omitempty"`
	// Optional. The subject to query for.
	Subject       *Subject `protobuf:"bytes,4,opt,name=subject,proto3" json:"subject,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListRelationTuplesRequest_Query) Reset() {
	*x = ListRelationTuplesRequest_Query{}
	mi := &file_ory_keto_relation_tuples_v1alpha2_read_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListRelationTuplesRequest_Query) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRelationTuplesRequest_Query) ProtoMessage() {}

func (x *ListRelationTuplesRequest_Query) ProtoReflect() protoreflect.Message {
	mi := &file_ory_keto_relation_tuples_v1alpha2_read_service_proto_msgTypes[2]
	if x != nil {
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
	return file_ory_keto_relation_tuples_v1alpha2_read_service_proto_rawDescGZIP(), []int{0, 0}
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

var File_ory_keto_relation_tuples_v1alpha2_read_service_proto protoreflect.FileDescriptor

var file_ory_keto_relation_tuples_v1alpha2_read_service_proto_rawDesc = []byte{
	0x0a, 0x34, 0x6f, 0x72, 0x79, 0x2f, 0x6b, 0x65, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x75, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x32, 0x2f, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x6f, 0x72, 0x79, 0x2e, 0x6b, 0x65, 0x74, 0x6f,
	0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x75, 0x70, 0x6c, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x37, 0x6f, 0x72, 0x79, 0x2f, 0x6b, 0x65, 0x74, 0x6f,
	0x2f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x75, 0x70, 0x6c, 0x65, 0x73,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x74, 0x75, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e,
	0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xc7, 0x06, 0x0a, 0x19, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x75, 0x70, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x6e, 0x0a,
	0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x42, 0x2e, 0x6f,
	0x72, 0x79, 0x2e, 0x6b, 0x65, 0x74, 0x6f, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x74, 0x75, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x75, 0x70,
	0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x42, 0x14, 0xfa, 0xd2, 0xe4, 0x93, 0x02, 0x0c, 0x12, 0x0a, 0x4e, 0x4f, 0x5f, 0x53, 0x57, 0x41,
	0x47, 0x47, 0x45, 0x52, 0x18, 0x01, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x6b, 0x0a,
	0x0e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x6f, 0x72, 0x79, 0x2e, 0x6b, 0x65, 0x74, 0x6f,
	0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x75, 0x70, 0x6c, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x42, 0x12, 0xfa, 0xd2, 0xe4, 0x93, 0x02, 0x0c, 0x12,
	0x0a, 0x4e, 0x4f, 0x5f, 0x53, 0x57, 0x41, 0x47, 0x47, 0x45, 0x52, 0x52, 0x0d, 0x72, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x4f, 0x0a, 0x0b, 0x65, 0x78,
	0x70, 0x61, 0x6e, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x42, 0x12, 0xfa, 0xd2, 0xe4,
	0x93, 0x02, 0x0c, 0x12, 0x0a, 0x4e, 0x4f, 0x5f, 0x53, 0x57, 0x41, 0x47, 0x47, 0x45, 0x52, 0x52,
	0x0a, 0x65, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x12, 0x30, 0x0a, 0x09, 0x73,
	0x6e, 0x61, 0x70, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x12,
	0xfa, 0xd2, 0xe4, 0x93, 0x02, 0x0c, 0x12, 0x0a, 0x4e, 0x4f, 0x5f, 0x53, 0x57, 0x41, 0x47, 0x47,
	0x45, 0x52, 0x52, 0x09, 0x73, 0x6e, 0x61, 0x70, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1c, 0x0a,
	0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x20, 0x0a, 0x09, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x02,
	0x18, 0x01, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x1a, 0x0a,
	0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x42, 0x02, 0x18,
	0x01, 0x52, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1e, 0x0a, 0x08, 0x72, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x42, 0x02, 0x18, 0x01, 0x52,
	0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0a, 0x73, 0x75, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x0a, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x12, 0x56, 0x0a, 0x0b, 0x73,
	0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x73, 0x65, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x32, 0x2e, 0x6f, 0x72, 0x79, 0x2e, 0x6b, 0x65, 0x74, 0x6f, 0x2e, 0x72, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x75, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x32, 0x2e, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x65, 0x74, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x48, 0x00, 0x52, 0x0b, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f,
	0x73, 0x65, 0x74, 0x1a, 0x9f, 0x01, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x1c, 0x0a,
	0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x44, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2a, 0x2e, 0x6f, 0x72, 0x79, 0x2e, 0x6b, 0x65, 0x74, 0x6f, 0x2e, 0x72, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x75, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x32, 0x2e, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x07, 0x73, 0x75,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x42, 0x12, 0x0a, 0x10, 0x72, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x70,
	0x69, 0x5f, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0xa2, 0x01, 0x0a, 0x1a, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x75, 0x70, 0x6c, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5a, 0x0a, 0x0f, 0x72, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x75, 0x70, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x30, 0x2e, 0x6f, 0x72, 0x79, 0x2e, 0x6b, 0x65, 0x74, 0x6f, 0x2e, 0x72, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x75, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x75,
	0x70, 0x6c, 0x65, 0x52, 0x0f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x75,
	0x70, 0x6c, 0x65, 0x73, 0x12, 0x28, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67,
	0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6e,
	0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0xeb,
	0x02, 0x0a, 0x0b, 0x52, 0x65, 0x61, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xdb,
	0x02, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x75, 0x70, 0x6c, 0x65, 0x73, 0x12, 0x3c, 0x2e, 0x6f, 0x72, 0x79, 0x2e, 0x6b, 0x65, 0x74, 0x6f,
	0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x75, 0x70, 0x6c, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x75, 0x70, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x3d, 0x2e, 0x6f, 0x72, 0x79, 0x2e, 0x6b, 0x65, 0x74, 0x6f, 0x2e, 0x72,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x75, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x75, 0x70, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0xc7, 0x01, 0x92, 0x41, 0xab, 0x01, 0x0a, 0x0c, 0x72, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x2a, 0x10, 0x67, 0x65, 0x74, 0x52, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x73, 0x32, 0x21, 0x61, 0x70, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x78, 0x2d, 0x77, 0x77, 0x77, 0x2d, 0x66, 0x6f, 0x72,
	0x6d, 0x2d, 0x75, 0x72, 0x6c, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x64, 0x4a, 0x66, 0x0a, 0x03,
	0x32, 0x30, 0x30, 0x12, 0x5f, 0x0a, 0x1a, 0x54, 0x68, 0x65, 0x20, 0x6c, 0x69, 0x73, 0x74, 0x20,
	0x6f, 0x66, 0x20, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x73,
	0x2e, 0x12, 0x41, 0x0a, 0x3f, 0x1a, 0x3d, 0x2e, 0x6f, 0x72, 0x79, 0x2e, 0x6b, 0x65, 0x74, 0x6f,
	0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x75, 0x70, 0x6c, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x75, 0x70, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x12, 0x10, 0x2f, 0x72, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x74, 0x75, 0x70, 0x6c, 0x65, 0x73, 0x42, 0xc1, 0x01, 0x0a,
	0x24, 0x73, 0x68, 0x2e, 0x6f, 0x72, 0x79, 0x2e, 0x6b, 0x65, 0x74, 0x6f, 0x2e, 0x72, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x75, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x32, 0x42, 0x10, 0x52, 0x65, 0x61, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x72, 0x79, 0x2f, 0x6b, 0x65, 0x74, 0x6f, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x72, 0x79, 0x2f, 0x6b, 0x65, 0x74, 0x6f, 0x2f, 0x72, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x75, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x3b, 0x72, 0x74, 0x73, 0xaa, 0x02, 0x20, 0x4f, 0x72, 0x79,
	0x2e, 0x4b, 0x65, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x75,
	0x70, 0x6c, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0xca, 0x02, 0x20,
	0x4f, 0x72, 0x79, 0x5c, 0x4b, 0x65, 0x74, 0x6f, 0x5c, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x75, 0x70, 0x6c, 0x65, 0x73, 0x5c, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ory_keto_relation_tuples_v1alpha2_read_service_proto_rawDescOnce sync.Once
	file_ory_keto_relation_tuples_v1alpha2_read_service_proto_rawDescData = file_ory_keto_relation_tuples_v1alpha2_read_service_proto_rawDesc
)

func file_ory_keto_relation_tuples_v1alpha2_read_service_proto_rawDescGZIP() []byte {
	file_ory_keto_relation_tuples_v1alpha2_read_service_proto_rawDescOnce.Do(func() {
		file_ory_keto_relation_tuples_v1alpha2_read_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_ory_keto_relation_tuples_v1alpha2_read_service_proto_rawDescData)
	})
	return file_ory_keto_relation_tuples_v1alpha2_read_service_proto_rawDescData
}

var file_ory_keto_relation_tuples_v1alpha2_read_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_ory_keto_relation_tuples_v1alpha2_read_service_proto_goTypes = []any{
	(*ListRelationTuplesRequest)(nil),       // 0: ory.keto.relation_tuples.v1alpha2.ListRelationTuplesRequest
	(*ListRelationTuplesResponse)(nil),      // 1: ory.keto.relation_tuples.v1alpha2.ListRelationTuplesResponse
	(*ListRelationTuplesRequest_Query)(nil), // 2: ory.keto.relation_tuples.v1alpha2.ListRelationTuplesRequest.Query
	(*RelationQuery)(nil),                   // 3: ory.keto.relation_tuples.v1alpha2.RelationQuery
	(*fieldmaskpb.FieldMask)(nil),           // 4: google.protobuf.FieldMask
	(*SubjectSetQuery)(nil),                 // 5: ory.keto.relation_tuples.v1alpha2.SubjectSetQuery
	(*RelationTuple)(nil),                   // 6: ory.keto.relation_tuples.v1alpha2.RelationTuple
	(*Subject)(nil),                         // 7: ory.keto.relation_tuples.v1alpha2.Subject
}
var file_ory_keto_relation_tuples_v1alpha2_read_service_proto_depIdxs = []int32{
	2, // 0: ory.keto.relation_tuples.v1alpha2.ListRelationTuplesRequest.query:type_name -> ory.keto.relation_tuples.v1alpha2.ListRelationTuplesRequest.Query
	3, // 1: ory.keto.relation_tuples.v1alpha2.ListRelationTuplesRequest.relation_query:type_name -> ory.keto.relation_tuples.v1alpha2.RelationQuery
	4, // 2: ory.keto.relation_tuples.v1alpha2.ListRelationTuplesRequest.expand_mask:type_name -> google.protobuf.FieldMask
	5, // 3: ory.keto.relation_tuples.v1alpha2.ListRelationTuplesRequest.subject_set:type_name -> ory.keto.relation_tuples.v1alpha2.SubjectSetQuery
	6, // 4: ory.keto.relation_tuples.v1alpha2.ListRelationTuplesResponse.relation_tuples:type_name -> ory.keto.relation_tuples.v1alpha2.RelationTuple
	7, // 5: ory.keto.relation_tuples.v1alpha2.ListRelationTuplesRequest.Query.subject:type_name -> ory.keto.relation_tuples.v1alpha2.Subject
	0, // 6: ory.keto.relation_tuples.v1alpha2.ReadService.ListRelationTuples:input_type -> ory.keto.relation_tuples.v1alpha2.ListRelationTuplesRequest
	1, // 7: ory.keto.relation_tuples.v1alpha2.ReadService.ListRelationTuples:output_type -> ory.keto.relation_tuples.v1alpha2.ListRelationTuplesResponse
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_ory_keto_relation_tuples_v1alpha2_read_service_proto_init() }
func file_ory_keto_relation_tuples_v1alpha2_read_service_proto_init() {
	if File_ory_keto_relation_tuples_v1alpha2_read_service_proto != nil {
		return
	}
	file_ory_keto_relation_tuples_v1alpha2_relation_tuples_proto_init()
	file_ory_keto_relation_tuples_v1alpha2_read_service_proto_msgTypes[0].OneofWrappers = []any{
		(*ListRelationTuplesRequest_SubjectId)(nil),
		(*ListRelationTuplesRequest_SubjectSet)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ory_keto_relation_tuples_v1alpha2_read_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ory_keto_relation_tuples_v1alpha2_read_service_proto_goTypes,
		DependencyIndexes: file_ory_keto_relation_tuples_v1alpha2_read_service_proto_depIdxs,
		MessageInfos:      file_ory_keto_relation_tuples_v1alpha2_read_service_proto_msgTypes,
	}.Build()
	File_ory_keto_relation_tuples_v1alpha2_read_service_proto = out.File
	file_ory_keto_relation_tuples_v1alpha2_read_service_proto_rawDesc = nil
	file_ory_keto_relation_tuples_v1alpha2_read_service_proto_goTypes = nil
	file_ory_keto_relation_tuples_v1alpha2_read_service_proto_depIdxs = nil
}
