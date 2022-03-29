// package: ory.keto.relation_tuples.v1alpha2
// file: ory/keto/relation_tuples/v1alpha2/expand_service.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "grpc";
import * as ory_keto_relation_tuples_v1alpha2_expand_service_pb from "../../../../ory/keto/relation_tuples/v1alpha2/expand_service_pb";
import * as ory_keto_relation_tuples_v1alpha2_relation_tuples_pb from "../../../../ory/keto/relation_tuples/v1alpha2/relation_tuples_pb";

interface IExpandServiceService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    expand: IExpandServiceService_IExpand;
}

interface IExpandServiceService_IExpand extends grpc.MethodDefinition<ory_keto_relation_tuples_v1alpha2_expand_service_pb.ExpandRequest, ory_keto_relation_tuples_v1alpha2_expand_service_pb.ExpandResponse> {
    path: "/ory.keto.relation_tuples.v1alpha2.ExpandService/Expand";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<ory_keto_relation_tuples_v1alpha2_expand_service_pb.ExpandRequest>;
    requestDeserialize: grpc.deserialize<ory_keto_relation_tuples_v1alpha2_expand_service_pb.ExpandRequest>;
    responseSerialize: grpc.serialize<ory_keto_relation_tuples_v1alpha2_expand_service_pb.ExpandResponse>;
    responseDeserialize: grpc.deserialize<ory_keto_relation_tuples_v1alpha2_expand_service_pb.ExpandResponse>;
}

export const ExpandServiceService: IExpandServiceService;

export interface IExpandServiceServer {
    expand: grpc.handleUnaryCall<ory_keto_relation_tuples_v1alpha2_expand_service_pb.ExpandRequest, ory_keto_relation_tuples_v1alpha2_expand_service_pb.ExpandResponse>;
}

export interface IExpandServiceClient {
    expand(request: ory_keto_relation_tuples_v1alpha2_expand_service_pb.ExpandRequest, callback: (error: grpc.ServiceError | null, response: ory_keto_relation_tuples_v1alpha2_expand_service_pb.ExpandResponse) => void): grpc.ClientUnaryCall;
    expand(request: ory_keto_relation_tuples_v1alpha2_expand_service_pb.ExpandRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: ory_keto_relation_tuples_v1alpha2_expand_service_pb.ExpandResponse) => void): grpc.ClientUnaryCall;
    expand(request: ory_keto_relation_tuples_v1alpha2_expand_service_pb.ExpandRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: ory_keto_relation_tuples_v1alpha2_expand_service_pb.ExpandResponse) => void): grpc.ClientUnaryCall;
}

export class ExpandServiceClient extends grpc.Client implements IExpandServiceClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
    public expand(request: ory_keto_relation_tuples_v1alpha2_expand_service_pb.ExpandRequest, callback: (error: grpc.ServiceError | null, response: ory_keto_relation_tuples_v1alpha2_expand_service_pb.ExpandResponse) => void): grpc.ClientUnaryCall;
    public expand(request: ory_keto_relation_tuples_v1alpha2_expand_service_pb.ExpandRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: ory_keto_relation_tuples_v1alpha2_expand_service_pb.ExpandResponse) => void): grpc.ClientUnaryCall;
    public expand(request: ory_keto_relation_tuples_v1alpha2_expand_service_pb.ExpandRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: ory_keto_relation_tuples_v1alpha2_expand_service_pb.ExpandResponse) => void): grpc.ClientUnaryCall;
}
