// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var ory_keto_acl_v1alpha1_write_service_pb = require('../../../../ory/keto/acl/v1alpha1/write_service_pb.js');
var ory_keto_acl_v1alpha1_acl_pb = require('../../../../ory/keto/acl/v1alpha1/acl_pb.js');

function serialize_ory_keto_acl_v1alpha1_TransactRelationTuplesRequest(arg) {
  if (!(arg instanceof ory_keto_acl_v1alpha1_write_service_pb.TransactRelationTuplesRequest)) {
    throw new Error('Expected argument of type ory.keto.acl.v1alpha1.TransactRelationTuplesRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_ory_keto_acl_v1alpha1_TransactRelationTuplesRequest(buffer_arg) {
  return ory_keto_acl_v1alpha1_write_service_pb.TransactRelationTuplesRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_ory_keto_acl_v1alpha1_TransactRelationTuplesResponse(arg) {
  if (!(arg instanceof ory_keto_acl_v1alpha1_write_service_pb.TransactRelationTuplesResponse)) {
    throw new Error('Expected argument of type ory.keto.acl.v1alpha1.TransactRelationTuplesResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_ory_keto_acl_v1alpha1_TransactRelationTuplesResponse(buffer_arg) {
  return ory_keto_acl_v1alpha1_write_service_pb.TransactRelationTuplesResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


// The write service to create and delete Access Control Lists.
//
// This service is part of the [write-APIs](../concepts/api-overview.mdx#write-apis).
var WriteServiceService = exports.WriteServiceService = {
  // Writes one or more relation tuples in a single transaction.
transactRelationTuples: {
    path: '/ory.keto.acl.v1alpha1.WriteService/TransactRelationTuples',
    requestStream: false,
    responseStream: false,
    requestType: ory_keto_acl_v1alpha1_write_service_pb.TransactRelationTuplesRequest,
    responseType: ory_keto_acl_v1alpha1_write_service_pb.TransactRelationTuplesResponse,
    requestSerialize: serialize_ory_keto_acl_v1alpha1_TransactRelationTuplesRequest,
    requestDeserialize: deserialize_ory_keto_acl_v1alpha1_TransactRelationTuplesRequest,
    responseSerialize: serialize_ory_keto_acl_v1alpha1_TransactRelationTuplesResponse,
    responseDeserialize: deserialize_ory_keto_acl_v1alpha1_TransactRelationTuplesResponse,
  },
};

exports.WriteServiceClient = grpc.makeGenericClientConstructor(WriteServiceService);
