// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var ory_keto_relation_tuples_v1alpha2_namespaces_service_pb = require('../../../../ory/keto/relation_tuples/v1alpha2/namespaces_service_pb.js');

function serialize_ory_keto_relation_tuples_v1alpha2_ListNamespacesRequest(arg) {
  if (!(arg instanceof ory_keto_relation_tuples_v1alpha2_namespaces_service_pb.ListNamespacesRequest)) {
    throw new Error('Expected argument of type ory.keto.relation_tuples.v1alpha2.ListNamespacesRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_ory_keto_relation_tuples_v1alpha2_ListNamespacesRequest(buffer_arg) {
  return ory_keto_relation_tuples_v1alpha2_namespaces_service_pb.ListNamespacesRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_ory_keto_relation_tuples_v1alpha2_ListNamespacesResponse(arg) {
  if (!(arg instanceof ory_keto_relation_tuples_v1alpha2_namespaces_service_pb.ListNamespacesResponse)) {
    throw new Error('Expected argument of type ory.keto.relation_tuples.v1alpha2.ListNamespacesResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_ory_keto_relation_tuples_v1alpha2_ListNamespacesResponse(buffer_arg) {
  return ory_keto_relation_tuples_v1alpha2_namespaces_service_pb.ListNamespacesResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


// The service to query namespaces.
//
// This service is part of the [read-APIs](../concepts/25_api-overview.mdx#read-apis).
var NamespacesServiceService = exports.NamespacesServiceService = {
  // Lists Namespaces
listNamespaces: {
    path: '/ory.keto.relation_tuples.v1alpha2.NamespacesService/ListNamespaces',
    requestStream: false,
    responseStream: false,
    requestType: ory_keto_relation_tuples_v1alpha2_namespaces_service_pb.ListNamespacesRequest,
    responseType: ory_keto_relation_tuples_v1alpha2_namespaces_service_pb.ListNamespacesResponse,
    requestSerialize: serialize_ory_keto_relation_tuples_v1alpha2_ListNamespacesRequest,
    requestDeserialize: deserialize_ory_keto_relation_tuples_v1alpha2_ListNamespacesRequest,
    responseSerialize: serialize_ory_keto_relation_tuples_v1alpha2_ListNamespacesResponse,
    responseDeserialize: deserialize_ory_keto_relation_tuples_v1alpha2_ListNamespacesResponse,
  },
};

exports.NamespacesServiceClient = grpc.makeGenericClientConstructor(NamespacesServiceService);
