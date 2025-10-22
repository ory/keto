// package: ory.keto.opl.v1alpha1
// file: ory/keto/opl/v1alpha1/syntax_service.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "@grpc/grpc-js";
import * as ory_keto_opl_v1alpha1_syntax_service_pb from "../../../../ory/keto/opl/v1alpha1/syntax_service_pb";
import * as protoc_gen_openapiv2_options_annotations_pb from "../../../../protoc-gen-openapiv2/options/annotations_pb";

interface ISyntaxServiceService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    check: ISyntaxServiceService_ICheck;
}

interface ISyntaxServiceService_ICheck extends grpc.MethodDefinition<ory_keto_opl_v1alpha1_syntax_service_pb.CheckRequest, ory_keto_opl_v1alpha1_syntax_service_pb.CheckResponse> {
    path: "/ory.keto.opl.v1alpha1.SyntaxService/Check";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<ory_keto_opl_v1alpha1_syntax_service_pb.CheckRequest>;
    requestDeserialize: grpc.deserialize<ory_keto_opl_v1alpha1_syntax_service_pb.CheckRequest>;
    responseSerialize: grpc.serialize<ory_keto_opl_v1alpha1_syntax_service_pb.CheckResponse>;
    responseDeserialize: grpc.deserialize<ory_keto_opl_v1alpha1_syntax_service_pb.CheckResponse>;
}

export const SyntaxServiceService: ISyntaxServiceService;

export interface ISyntaxServiceServer extends grpc.UntypedServiceImplementation {
    check: grpc.handleUnaryCall<ory_keto_opl_v1alpha1_syntax_service_pb.CheckRequest, ory_keto_opl_v1alpha1_syntax_service_pb.CheckResponse>;
}

export interface ISyntaxServiceClient {
    check(request: ory_keto_opl_v1alpha1_syntax_service_pb.CheckRequest, callback: (error: grpc.ServiceError | null, response: ory_keto_opl_v1alpha1_syntax_service_pb.CheckResponse) => void): grpc.ClientUnaryCall;
    check(request: ory_keto_opl_v1alpha1_syntax_service_pb.CheckRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: ory_keto_opl_v1alpha1_syntax_service_pb.CheckResponse) => void): grpc.ClientUnaryCall;
    check(request: ory_keto_opl_v1alpha1_syntax_service_pb.CheckRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: ory_keto_opl_v1alpha1_syntax_service_pb.CheckResponse) => void): grpc.ClientUnaryCall;
}

export class SyntaxServiceClient extends grpc.Client implements ISyntaxServiceClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: Partial<grpc.ClientOptions>);
    public check(request: ory_keto_opl_v1alpha1_syntax_service_pb.CheckRequest, callback: (error: grpc.ServiceError | null, response: ory_keto_opl_v1alpha1_syntax_service_pb.CheckResponse) => void): grpc.ClientUnaryCall;
    public check(request: ory_keto_opl_v1alpha1_syntax_service_pb.CheckRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: ory_keto_opl_v1alpha1_syntax_service_pb.CheckResponse) => void): grpc.ClientUnaryCall;
    public check(request: ory_keto_opl_v1alpha1_syntax_service_pb.CheckRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: ory_keto_opl_v1alpha1_syntax_service_pb.CheckResponse) => void): grpc.ClientUnaryCall;
}
