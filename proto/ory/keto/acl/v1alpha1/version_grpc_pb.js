// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var ory_keto_acl_v1alpha1_version_pb = require('../../../../ory/keto/acl/v1alpha1/version_pb.js');

function serialize_ory_keto_acl_v1alpha1_GetVersionRequest(arg) {
  if (!(arg instanceof ory_keto_acl_v1alpha1_version_pb.GetVersionRequest)) {
    throw new Error('Expected argument of type ory.keto.acl.v1alpha1.GetVersionRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_ory_keto_acl_v1alpha1_GetVersionRequest(buffer_arg) {
  return ory_keto_acl_v1alpha1_version_pb.GetVersionRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_ory_keto_acl_v1alpha1_GetVersionResponse(arg) {
  if (!(arg instanceof ory_keto_acl_v1alpha1_version_pb.GetVersionResponse)) {
    throw new Error('Expected argument of type ory.keto.acl.v1alpha1.GetVersionResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_ory_keto_acl_v1alpha1_GetVersionResponse(buffer_arg) {
  return ory_keto_acl_v1alpha1_version_pb.GetVersionResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


// The service returning the specific Ory Keto instance version.
//
// This service is part of the [read-APIs](../concepts/api-overview.mdx#read-apis) and [write-APIs](../concepts/api-overview.mdx#write-apis).
var VersionServiceService = exports.VersionServiceService = {
  // Returns the version of the Ory Keto instance.
getVersion: {
    path: '/ory.keto.acl.v1alpha1.VersionService/GetVersion',
    requestStream: false,
    responseStream: false,
    requestType: ory_keto_acl_v1alpha1_version_pb.GetVersionRequest,
    responseType: ory_keto_acl_v1alpha1_version_pb.GetVersionResponse,
    requestSerialize: serialize_ory_keto_acl_v1alpha1_GetVersionRequest,
    requestDeserialize: deserialize_ory_keto_acl_v1alpha1_GetVersionRequest,
    responseSerialize: serialize_ory_keto_acl_v1alpha1_GetVersionResponse,
    responseDeserialize: deserialize_ory_keto_acl_v1alpha1_GetVersionResponse,
  },
};

exports.VersionServiceClient = grpc.makeGenericClientConstructor(VersionServiceService);
