// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var ory_keto_relation_tuples_v1alpha2_version_pb = require('../../../../ory/keto/relation_tuples/v1alpha2/version_pb.js');

function serialize_ory_keto_relation_tuples_v1alpha2_GetVersionRequest(arg) {
  if (!(arg instanceof ory_keto_relation_tuples_v1alpha2_version_pb.GetVersionRequest)) {
    throw new Error('Expected argument of type ory.keto.relation_tuples.v1alpha2.GetVersionRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_ory_keto_relation_tuples_v1alpha2_GetVersionRequest(buffer_arg) {
  return ory_keto_relation_tuples_v1alpha2_version_pb.GetVersionRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_ory_keto_relation_tuples_v1alpha2_GetVersionResponse(arg) {
  if (!(arg instanceof ory_keto_relation_tuples_v1alpha2_version_pb.GetVersionResponse)) {
    throw new Error('Expected argument of type ory.keto.relation_tuples.v1alpha2.GetVersionResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_ory_keto_relation_tuples_v1alpha2_GetVersionResponse(buffer_arg) {
  return ory_keto_relation_tuples_v1alpha2_version_pb.GetVersionResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


// The service returning the specific Ory Keto instance version.
//
// This service is part of the [read-APIs](../concepts/25_api-overview.mdx#read-apis) and [write-APIs](../concepts/25_api-overview.mdx#write-apis).
var VersionServiceService = exports.VersionServiceService = {
  // Returns the version of the Ory Keto instance.
getVersion: {
    path: '/ory.keto.relation_tuples.v1alpha2.VersionService/GetVersion',
    requestStream: false,
    responseStream: false,
    requestType: ory_keto_relation_tuples_v1alpha2_version_pb.GetVersionRequest,
    responseType: ory_keto_relation_tuples_v1alpha2_version_pb.GetVersionResponse,
    requestSerialize: serialize_ory_keto_relation_tuples_v1alpha2_GetVersionRequest,
    requestDeserialize: deserialize_ory_keto_relation_tuples_v1alpha2_GetVersionRequest,
    responseSerialize: serialize_ory_keto_relation_tuples_v1alpha2_GetVersionResponse,
    responseDeserialize: deserialize_ory_keto_relation_tuples_v1alpha2_GetVersionResponse,
  },
};

exports.VersionServiceClient = grpc.makeGenericClientConstructor(VersionServiceService);
