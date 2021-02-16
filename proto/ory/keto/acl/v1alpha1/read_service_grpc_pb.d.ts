// package: ory.keto.acl.v1alpha1
// file: ory/keto/acl/v1alpha1/read_service.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "grpc";
import * as ory_keto_acl_v1alpha1_read_service_pb from "../../../../ory/keto/acl/v1alpha1/read_service_pb";
import * as ory_keto_acl_v1alpha1_acl_pb from "../../../../ory/keto/acl/v1alpha1/acl_pb";
import * as google_protobuf_field_mask_pb from "google-protobuf/google/protobuf/field_mask_pb";

interface IReadServiceService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    listRelationTuples: IReadServiceService_IListRelationTuples;
}

interface IReadServiceService_IListRelationTuples extends grpc.MethodDefinition<ory_keto_acl_v1alpha1_read_service_pb.ListRelationTuplesRequest, ory_keto_acl_v1alpha1_read_service_pb.ListRelationTuplesResponse> {
    path: "/ory.keto.acl.v1alpha1.ReadService/ListRelationTuples";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<ory_keto_acl_v1alpha1_read_service_pb.ListRelationTuplesRequest>;
    requestDeserialize: grpc.deserialize<ory_keto_acl_v1alpha1_read_service_pb.ListRelationTuplesRequest>;
    responseSerialize: grpc.serialize<ory_keto_acl_v1alpha1_read_service_pb.ListRelationTuplesResponse>;
    responseDeserialize: grpc.deserialize<ory_keto_acl_v1alpha1_read_service_pb.ListRelationTuplesResponse>;
}

export const ReadServiceService: IReadServiceService;

export interface IReadServiceServer {
    listRelationTuples: grpc.handleUnaryCall<ory_keto_acl_v1alpha1_read_service_pb.ListRelationTuplesRequest, ory_keto_acl_v1alpha1_read_service_pb.ListRelationTuplesResponse>;
}

export interface IReadServiceClient {
    listRelationTuples(request: ory_keto_acl_v1alpha1_read_service_pb.ListRelationTuplesRequest, callback: (error: grpc.ServiceError | null, response: ory_keto_acl_v1alpha1_read_service_pb.ListRelationTuplesResponse) => void): grpc.ClientUnaryCall;
    listRelationTuples(request: ory_keto_acl_v1alpha1_read_service_pb.ListRelationTuplesRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: ory_keto_acl_v1alpha1_read_service_pb.ListRelationTuplesResponse) => void): grpc.ClientUnaryCall;
    listRelationTuples(request: ory_keto_acl_v1alpha1_read_service_pb.ListRelationTuplesRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: ory_keto_acl_v1alpha1_read_service_pb.ListRelationTuplesResponse) => void): grpc.ClientUnaryCall;
}

export class ReadServiceClient extends grpc.Client implements IReadServiceClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
    public listRelationTuples(request: ory_keto_acl_v1alpha1_read_service_pb.ListRelationTuplesRequest, callback: (error: grpc.ServiceError | null, response: ory_keto_acl_v1alpha1_read_service_pb.ListRelationTuplesResponse) => void): grpc.ClientUnaryCall;
    public listRelationTuples(request: ory_keto_acl_v1alpha1_read_service_pb.ListRelationTuplesRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: ory_keto_acl_v1alpha1_read_service_pb.ListRelationTuplesResponse) => void): grpc.ClientUnaryCall;
    public listRelationTuples(request: ory_keto_acl_v1alpha1_read_service_pb.ListRelationTuplesRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: ory_keto_acl_v1alpha1_read_service_pb.ListRelationTuplesResponse) => void): grpc.ClientUnaryCall;
}
