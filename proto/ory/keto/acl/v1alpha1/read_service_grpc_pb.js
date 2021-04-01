// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var ory_keto_acl_v1alpha1_read_service_pb = require('../../../../ory/keto/acl/v1alpha1/read_service_pb.js');
var ory_keto_acl_v1alpha1_acl_pb = require('../../../../ory/keto/acl/v1alpha1/acl_pb.js');
var google_protobuf_field_mask_pb = require('google-protobuf/google/protobuf/field_mask_pb.js');

function serialize_ory_keto_acl_v1alpha1_ListRelationTuplesRequest(arg) {
  if (!(arg instanceof ory_keto_acl_v1alpha1_read_service_pb.ListRelationTuplesRequest)) {
    throw new Error('Expected argument of type ory.keto.acl.v1alpha1.ListRelationTuplesRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_ory_keto_acl_v1alpha1_ListRelationTuplesRequest(buffer_arg) {
  return ory_keto_acl_v1alpha1_read_service_pb.ListRelationTuplesRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_ory_keto_acl_v1alpha1_ListRelationTuplesResponse(arg) {
  if (!(arg instanceof ory_keto_acl_v1alpha1_read_service_pb.ListRelationTuplesResponse)) {
    throw new Error('Expected argument of type ory.keto.acl.v1alpha1.ListRelationTuplesResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_ory_keto_acl_v1alpha1_ListRelationTuplesResponse(buffer_arg) {
  return ory_keto_acl_v1alpha1_read_service_pb.ListRelationTuplesResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


// The service to query relation tuples.
//
// This service is part of the [read-APIs](../concepts/api-overview.mdx#read-apis).
var ReadServiceService = exports.ReadServiceService = {
  // Lists ACL relation tuples.
listRelationTuples: {
    path: '/ory.keto.acl.v1alpha1.ReadService/ListRelationTuples',
    requestStream: false,
    responseStream: false,
    requestType: ory_keto_acl_v1alpha1_read_service_pb.ListRelationTuplesRequest,
    responseType: ory_keto_acl_v1alpha1_read_service_pb.ListRelationTuplesResponse,
    requestSerialize: serialize_ory_keto_acl_v1alpha1_ListRelationTuplesRequest,
    requestDeserialize: deserialize_ory_keto_acl_v1alpha1_ListRelationTuplesRequest,
    responseSerialize: serialize_ory_keto_acl_v1alpha1_ListRelationTuplesResponse,
    responseDeserialize: deserialize_ory_keto_acl_v1alpha1_ListRelationTuplesResponse,
  },
};

exports.ReadServiceClient = grpc.makeGenericClientConstructor(ReadServiceService);
