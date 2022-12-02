// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var ory_keto_relation_tuples_v1alpha2_write_service_pb = require('../../../../ory/keto/relation_tuples/v1alpha2/write_service_pb.js');
var google_api_annotations_pb = require('../../../../google/api/annotations_pb.js');
var google_api_visibility_pb = require('../../../../google/api/visibility_pb.js');
var protoc$gen$openapiv2_options_annotations_pb = require('../../../../protoc-gen-openapiv2/options/annotations_pb.js');
var ory_keto_relation_tuples_v1alpha2_relation_tuples_pb = require('../../../../ory/keto/relation_tuples/v1alpha2/relation_tuples_pb.js');

function serialize_ory_keto_relation_tuples_v1alpha2_CreateRelationTupleRequest(arg) {
  if (!(arg instanceof ory_keto_relation_tuples_v1alpha2_write_service_pb.CreateRelationTupleRequest)) {
    throw new Error('Expected argument of type ory.keto.relation_tuples.v1alpha2.CreateRelationTupleRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_ory_keto_relation_tuples_v1alpha2_CreateRelationTupleRequest(buffer_arg) {
  return ory_keto_relation_tuples_v1alpha2_write_service_pb.CreateRelationTupleRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_ory_keto_relation_tuples_v1alpha2_CreateRelationTupleResponse(arg) {
  if (!(arg instanceof ory_keto_relation_tuples_v1alpha2_write_service_pb.CreateRelationTupleResponse)) {
    throw new Error('Expected argument of type ory.keto.relation_tuples.v1alpha2.CreateRelationTupleResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_ory_keto_relation_tuples_v1alpha2_CreateRelationTupleResponse(buffer_arg) {
  return ory_keto_relation_tuples_v1alpha2_write_service_pb.CreateRelationTupleResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_ory_keto_relation_tuples_v1alpha2_DeleteRelationTuplesRequest(arg) {
  if (!(arg instanceof ory_keto_relation_tuples_v1alpha2_write_service_pb.DeleteRelationTuplesRequest)) {
    throw new Error('Expected argument of type ory.keto.relation_tuples.v1alpha2.DeleteRelationTuplesRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_ory_keto_relation_tuples_v1alpha2_DeleteRelationTuplesRequest(buffer_arg) {
  return ory_keto_relation_tuples_v1alpha2_write_service_pb.DeleteRelationTuplesRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_ory_keto_relation_tuples_v1alpha2_DeleteRelationTuplesResponse(arg) {
  if (!(arg instanceof ory_keto_relation_tuples_v1alpha2_write_service_pb.DeleteRelationTuplesResponse)) {
    throw new Error('Expected argument of type ory.keto.relation_tuples.v1alpha2.DeleteRelationTuplesResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_ory_keto_relation_tuples_v1alpha2_DeleteRelationTuplesResponse(buffer_arg) {
  return ory_keto_relation_tuples_v1alpha2_write_service_pb.DeleteRelationTuplesResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_ory_keto_relation_tuples_v1alpha2_TransactRelationTuplesRequest(arg) {
  if (!(arg instanceof ory_keto_relation_tuples_v1alpha2_write_service_pb.TransactRelationTuplesRequest)) {
    throw new Error('Expected argument of type ory.keto.relation_tuples.v1alpha2.TransactRelationTuplesRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_ory_keto_relation_tuples_v1alpha2_TransactRelationTuplesRequest(buffer_arg) {
  return ory_keto_relation_tuples_v1alpha2_write_service_pb.TransactRelationTuplesRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_ory_keto_relation_tuples_v1alpha2_TransactRelationTuplesResponse(arg) {
  if (!(arg instanceof ory_keto_relation_tuples_v1alpha2_write_service_pb.TransactRelationTuplesResponse)) {
    throw new Error('Expected argument of type ory.keto.relation_tuples.v1alpha2.TransactRelationTuplesResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_ory_keto_relation_tuples_v1alpha2_TransactRelationTuplesResponse(buffer_arg) {
  return ory_keto_relation_tuples_v1alpha2_write_service_pb.TransactRelationTuplesResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


// The write service to create and delete Access Control Lists.
//
// This service is part of the [write-APIs](../concepts/api-overview.mdx#write-apis).
var WriteServiceService = exports.WriteServiceService = {
  // Writes one or more relationships in a single transaction.
transactRelationTuples: {
    path: '/ory.keto.relation_tuples.v1alpha2.WriteService/TransactRelationTuples',
    requestStream: false,
    responseStream: false,
    requestType: ory_keto_relation_tuples_v1alpha2_write_service_pb.TransactRelationTuplesRequest,
    responseType: ory_keto_relation_tuples_v1alpha2_write_service_pb.TransactRelationTuplesResponse,
    requestSerialize: serialize_ory_keto_relation_tuples_v1alpha2_TransactRelationTuplesRequest,
    requestDeserialize: deserialize_ory_keto_relation_tuples_v1alpha2_TransactRelationTuplesRequest,
    responseSerialize: serialize_ory_keto_relation_tuples_v1alpha2_TransactRelationTuplesResponse,
    responseDeserialize: deserialize_ory_keto_relation_tuples_v1alpha2_TransactRelationTuplesResponse,
  },
  // Creates a relationship
createRelationTuple: {
    path: '/ory.keto.relation_tuples.v1alpha2.WriteService/CreateRelationTuple',
    requestStream: false,
    responseStream: false,
    requestType: ory_keto_relation_tuples_v1alpha2_write_service_pb.CreateRelationTupleRequest,
    responseType: ory_keto_relation_tuples_v1alpha2_write_service_pb.CreateRelationTupleResponse,
    requestSerialize: serialize_ory_keto_relation_tuples_v1alpha2_CreateRelationTupleRequest,
    requestDeserialize: deserialize_ory_keto_relation_tuples_v1alpha2_CreateRelationTupleRequest,
    responseSerialize: serialize_ory_keto_relation_tuples_v1alpha2_CreateRelationTupleResponse,
    responseDeserialize: deserialize_ory_keto_relation_tuples_v1alpha2_CreateRelationTupleResponse,
  },
  // Deletes relationships based on relation query
deleteRelationTuples: {
    path: '/ory.keto.relation_tuples.v1alpha2.WriteService/DeleteRelationTuples',
    requestStream: false,
    responseStream: false,
    requestType: ory_keto_relation_tuples_v1alpha2_write_service_pb.DeleteRelationTuplesRequest,
    responseType: ory_keto_relation_tuples_v1alpha2_write_service_pb.DeleteRelationTuplesResponse,
    requestSerialize: serialize_ory_keto_relation_tuples_v1alpha2_DeleteRelationTuplesRequest,
    requestDeserialize: deserialize_ory_keto_relation_tuples_v1alpha2_DeleteRelationTuplesRequest,
    responseSerialize: serialize_ory_keto_relation_tuples_v1alpha2_DeleteRelationTuplesResponse,
    responseDeserialize: deserialize_ory_keto_relation_tuples_v1alpha2_DeleteRelationTuplesResponse,
  },
};

exports.WriteServiceClient = grpc.makeGenericClientConstructor(WriteServiceService);
