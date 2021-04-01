// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var ory_keto_acl_v1alpha1_check_service_pb = require('../../../../ory/keto/acl/v1alpha1/check_service_pb.js');
var ory_keto_acl_v1alpha1_acl_pb = require('../../../../ory/keto/acl/v1alpha1/acl_pb.js');

function serialize_ory_keto_acl_v1alpha1_CheckRequest(arg) {
  if (!(arg instanceof ory_keto_acl_v1alpha1_check_service_pb.CheckRequest)) {
    throw new Error('Expected argument of type ory.keto.acl.v1alpha1.CheckRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_ory_keto_acl_v1alpha1_CheckRequest(buffer_arg) {
  return ory_keto_acl_v1alpha1_check_service_pb.CheckRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_ory_keto_acl_v1alpha1_CheckResponse(arg) {
  if (!(arg instanceof ory_keto_acl_v1alpha1_check_service_pb.CheckResponse)) {
    throw new Error('Expected argument of type ory.keto.acl.v1alpha1.CheckResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_ory_keto_acl_v1alpha1_CheckResponse(buffer_arg) {
  return ory_keto_acl_v1alpha1_check_service_pb.CheckResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


// The service that performs authorization checks
// based on the stored Access Control Lists.
//
// This service is part of the [read-APIs](../concepts/api-overview.mdx#read-apis).
var CheckServiceService = exports.CheckServiceService = {
  // Performs an authorization check.
check: {
    path: '/ory.keto.acl.v1alpha1.CheckService/Check',
    requestStream: false,
    responseStream: false,
    requestType: ory_keto_acl_v1alpha1_check_service_pb.CheckRequest,
    responseType: ory_keto_acl_v1alpha1_check_service_pb.CheckResponse,
    requestSerialize: serialize_ory_keto_acl_v1alpha1_CheckRequest,
    requestDeserialize: deserialize_ory_keto_acl_v1alpha1_CheckRequest,
    responseSerialize: serialize_ory_keto_acl_v1alpha1_CheckResponse,
    responseDeserialize: deserialize_ory_keto_acl_v1alpha1_CheckResponse,
  },
};

exports.CheckServiceClient = grpc.makeGenericClientConstructor(CheckServiceService);
