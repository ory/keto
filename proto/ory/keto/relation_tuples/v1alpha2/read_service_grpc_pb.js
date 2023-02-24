// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var ory_keto_relation_tuples_v1alpha2_read_service_pb = require('../../../../ory/keto/relation_tuples/v1alpha2/read_service_pb.js');
var ory_keto_relation_tuples_v1alpha2_relation_tuples_pb = require('../../../../ory/keto/relation_tuples/v1alpha2/relation_tuples_pb.js');
var google_protobuf_field_mask_pb = require('google-protobuf/google/protobuf/field_mask_pb.js');

function serialize_ory_keto_relation_tuples_v1alpha2_ListRelationTuplesRequest(arg) {
  if (!(arg instanceof ory_keto_relation_tuples_v1alpha2_read_service_pb.ListRelationTuplesRequest)) {
    throw new Error('Expected argument of type ory.keto.relation_tuples.v1alpha2.ListRelationTuplesRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_ory_keto_relation_tuples_v1alpha2_ListRelationTuplesRequest(buffer_arg) {
  return ory_keto_relation_tuples_v1alpha2_read_service_pb.ListRelationTuplesRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_ory_keto_relation_tuples_v1alpha2_ListRelationTuplesResponse(arg) {
  if (!(arg instanceof ory_keto_relation_tuples_v1alpha2_read_service_pb.ListRelationTuplesResponse)) {
    throw new Error('Expected argument of type ory.keto.relation_tuples.v1alpha2.ListRelationTuplesResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_ory_keto_relation_tuples_v1alpha2_ListRelationTuplesResponse(buffer_arg) {
  return ory_keto_relation_tuples_v1alpha2_read_service_pb.ListRelationTuplesResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


// The service to query relationships.
//
// This service is part of the [read-APIs](../concepts/25_api-overview.mdx#read-apis).
var ReadServiceService = exports.ReadServiceService = {
  // Lists ACL relationships.
listRelationTuples: {
    path: '/ory.keto.relation_tuples.v1alpha2.ReadService/ListRelationTuples',
    requestStream: false,
    responseStream: false,
    requestType: ory_keto_relation_tuples_v1alpha2_read_service_pb.ListRelationTuplesRequest,
    responseType: ory_keto_relation_tuples_v1alpha2_read_service_pb.ListRelationTuplesResponse,
    requestSerialize: serialize_ory_keto_relation_tuples_v1alpha2_ListRelationTuplesRequest,
    requestDeserialize: deserialize_ory_keto_relation_tuples_v1alpha2_ListRelationTuplesRequest,
    responseSerialize: serialize_ory_keto_relation_tuples_v1alpha2_ListRelationTuplesResponse,
    responseDeserialize: deserialize_ory_keto_relation_tuples_v1alpha2_ListRelationTuplesResponse,
  },
};

exports.ReadServiceClient = grpc.makeGenericClientConstructor(ReadServiceService);
