// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var ory_keto_opl_v1alpha1_syntax_service_pb = require('../../../../ory/keto/opl/v1alpha1/syntax_service_pb.js');

function serialize_ory_keto_opl_v1alpha1_CheckRequest(arg) {
  if (!(arg instanceof ory_keto_opl_v1alpha1_syntax_service_pb.CheckRequest)) {
    throw new Error('Expected argument of type ory.keto.opl.v1alpha1.CheckRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_ory_keto_opl_v1alpha1_CheckRequest(buffer_arg) {
  return ory_keto_opl_v1alpha1_syntax_service_pb.CheckRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_ory_keto_opl_v1alpha1_CheckResponse(arg) {
  if (!(arg instanceof ory_keto_opl_v1alpha1_syntax_service_pb.CheckResponse)) {
    throw new Error('Expected argument of type ory.keto.opl.v1alpha1.CheckResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_ory_keto_opl_v1alpha1_CheckResponse(buffer_arg) {
  return ory_keto_opl_v1alpha1_syntax_service_pb.CheckResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


// The service that checks the syntax of an OPL file.
var SyntaxServiceService = exports.SyntaxServiceService = {
  // Performs a syntax check request.
check: {
    path: '/ory.keto.opl.v1alpha1.SyntaxService/Check',
    requestStream: false,
    responseStream: false,
    requestType: ory_keto_opl_v1alpha1_syntax_service_pb.CheckRequest,
    responseType: ory_keto_opl_v1alpha1_syntax_service_pb.CheckResponse,
    requestSerialize: serialize_ory_keto_opl_v1alpha1_CheckRequest,
    requestDeserialize: deserialize_ory_keto_opl_v1alpha1_CheckRequest,
    responseSerialize: serialize_ory_keto_opl_v1alpha1_CheckResponse,
    responseDeserialize: deserialize_ory_keto_opl_v1alpha1_CheckResponse,
  },
};

exports.SyntaxServiceClient = grpc.makeGenericClientConstructor(SyntaxServiceService);
