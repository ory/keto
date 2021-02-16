// package: ory.keto.acl.v1alpha1
// file: ory/keto/acl/v1alpha1/expand_service.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "grpc";
import * as ory_keto_acl_v1alpha1_expand_service_pb from "../../../../ory/keto/acl/v1alpha1/expand_service_pb";
import * as ory_keto_acl_v1alpha1_acl_pb from "../../../../ory/keto/acl/v1alpha1/acl_pb";

interface IExpandServiceService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    expand: IExpandServiceService_IExpand;
}

interface IExpandServiceService_IExpand extends grpc.MethodDefinition<ory_keto_acl_v1alpha1_expand_service_pb.ExpandRequest, ory_keto_acl_v1alpha1_expand_service_pb.ExpandResponse> {
    path: "/ory.keto.acl.v1alpha1.ExpandService/Expand";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<ory_keto_acl_v1alpha1_expand_service_pb.ExpandRequest>;
    requestDeserialize: grpc.deserialize<ory_keto_acl_v1alpha1_expand_service_pb.ExpandRequest>;
    responseSerialize: grpc.serialize<ory_keto_acl_v1alpha1_expand_service_pb.ExpandResponse>;
    responseDeserialize: grpc.deserialize<ory_keto_acl_v1alpha1_expand_service_pb.ExpandResponse>;
}

export const ExpandServiceService: IExpandServiceService;

export interface IExpandServiceServer {
    expand: grpc.handleUnaryCall<ory_keto_acl_v1alpha1_expand_service_pb.ExpandRequest, ory_keto_acl_v1alpha1_expand_service_pb.ExpandResponse>;
}

export interface IExpandServiceClient {
    expand(request: ory_keto_acl_v1alpha1_expand_service_pb.ExpandRequest, callback: (error: grpc.ServiceError | null, response: ory_keto_acl_v1alpha1_expand_service_pb.ExpandResponse) => void): grpc.ClientUnaryCall;
    expand(request: ory_keto_acl_v1alpha1_expand_service_pb.ExpandRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: ory_keto_acl_v1alpha1_expand_service_pb.ExpandResponse) => void): grpc.ClientUnaryCall;
    expand(request: ory_keto_acl_v1alpha1_expand_service_pb.ExpandRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: ory_keto_acl_v1alpha1_expand_service_pb.ExpandResponse) => void): grpc.ClientUnaryCall;
}

export class ExpandServiceClient extends grpc.Client implements IExpandServiceClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
    public expand(request: ory_keto_acl_v1alpha1_expand_service_pb.ExpandRequest, callback: (error: grpc.ServiceError | null, response: ory_keto_acl_v1alpha1_expand_service_pb.ExpandResponse) => void): grpc.ClientUnaryCall;
    public expand(request: ory_keto_acl_v1alpha1_expand_service_pb.ExpandRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: ory_keto_acl_v1alpha1_expand_service_pb.ExpandResponse) => void): grpc.ClientUnaryCall;
    public expand(request: ory_keto_acl_v1alpha1_expand_service_pb.ExpandRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: ory_keto_acl_v1alpha1_expand_service_pb.ExpandResponse) => void): grpc.ClientUnaryCall;
}
