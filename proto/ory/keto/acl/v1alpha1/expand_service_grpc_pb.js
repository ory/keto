// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var ory_keto_acl_v1alpha1_expand_service_pb = require('../../../../ory/keto/acl/v1alpha1/expand_service_pb.js');
var ory_keto_acl_v1alpha1_acl_pb = require('../../../../ory/keto/acl/v1alpha1/acl_pb.js');

function serialize_ory_keto_acl_v1alpha1_ExpandRequest(arg) {
  if (!(arg instanceof ory_keto_acl_v1alpha1_expand_service_pb.ExpandRequest)) {
    throw new Error('Expected argument of type ory.keto.acl.v1alpha1.ExpandRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_ory_keto_acl_v1alpha1_ExpandRequest(buffer_arg) {
  return ory_keto_acl_v1alpha1_expand_service_pb.ExpandRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_ory_keto_acl_v1alpha1_ExpandResponse(arg) {
  if (!(arg instanceof ory_keto_acl_v1alpha1_expand_service_pb.ExpandResponse)) {
    throw new Error('Expected argument of type ory.keto.acl.v1alpha1.ExpandResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_ory_keto_acl_v1alpha1_ExpandResponse(buffer_arg) {
  return ory_keto_acl_v1alpha1_expand_service_pb.ExpandResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


// The service that performs subject set expansion
// based on the stored Access Control Lists.
//
// This service is part of the [read-APIs](../concepts/api-overview.mdx#read-apis).
var ExpandServiceService = exports.ExpandServiceService = {
  // Expands the subject set into a tree of subjects.
expand: {
    path: '/ory.keto.acl.v1alpha1.ExpandService/Expand',
    requestStream: false,
    responseStream: false,
    requestType: ory_keto_acl_v1alpha1_expand_service_pb.ExpandRequest,
    responseType: ory_keto_acl_v1alpha1_expand_service_pb.ExpandResponse,
    requestSerialize: serialize_ory_keto_acl_v1alpha1_ExpandRequest,
    requestDeserialize: deserialize_ory_keto_acl_v1alpha1_ExpandRequest,
    responseSerialize: serialize_ory_keto_acl_v1alpha1_ExpandResponse,
    responseDeserialize: deserialize_ory_keto_acl_v1alpha1_ExpandResponse,
  },
};

exports.ExpandServiceClient = grpc.makeGenericClientConstructor(ExpandServiceService);
