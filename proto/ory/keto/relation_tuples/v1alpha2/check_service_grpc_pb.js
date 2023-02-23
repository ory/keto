// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var ory_keto_relation_tuples_v1alpha2_check_service_pb = require('../../../../ory/keto/relation_tuples/v1alpha2/check_service_pb.js');
var ory_keto_relation_tuples_v1alpha2_relation_tuples_pb = require('../../../../ory/keto/relation_tuples/v1alpha2/relation_tuples_pb.js');

function serialize_ory_keto_relation_tuples_v1alpha2_CheckRequest(arg) {
  if (!(arg instanceof ory_keto_relation_tuples_v1alpha2_check_service_pb.CheckRequest)) {
    throw new Error('Expected argument of type ory.keto.relation_tuples.v1alpha2.CheckRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_ory_keto_relation_tuples_v1alpha2_CheckRequest(buffer_arg) {
  return ory_keto_relation_tuples_v1alpha2_check_service_pb.CheckRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_ory_keto_relation_tuples_v1alpha2_CheckResponse(arg) {
  if (!(arg instanceof ory_keto_relation_tuples_v1alpha2_check_service_pb.CheckResponse)) {
    throw new Error('Expected argument of type ory.keto.relation_tuples.v1alpha2.CheckResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_ory_keto_relation_tuples_v1alpha2_CheckResponse(buffer_arg) {
  return ory_keto_relation_tuples_v1alpha2_check_service_pb.CheckResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


// The service that performs authorization checks
// based on the stored Access Control Lists.
//
// This service is part of the [read-APIs](../concepts/25_api-overview.mdx#read-apis).
var CheckServiceService = exports.CheckServiceService = {
  // Performs an authorization check.
check: {
    path: '/ory.keto.relation_tuples.v1alpha2.CheckService/Check',
    requestStream: false,
    responseStream: false,
    requestType: ory_keto_relation_tuples_v1alpha2_check_service_pb.CheckRequest,
    responseType: ory_keto_relation_tuples_v1alpha2_check_service_pb.CheckResponse,
    requestSerialize: serialize_ory_keto_relation_tuples_v1alpha2_CheckRequest,
    requestDeserialize: deserialize_ory_keto_relation_tuples_v1alpha2_CheckRequest,
    responseSerialize: serialize_ory_keto_relation_tuples_v1alpha2_CheckResponse,
    responseDeserialize: deserialize_ory_keto_relation_tuples_v1alpha2_CheckResponse,
  },
};

exports.CheckServiceClient = grpc.makeGenericClientConstructor(CheckServiceService);
