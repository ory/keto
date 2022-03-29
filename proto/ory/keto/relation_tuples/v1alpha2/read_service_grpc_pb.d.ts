// package: ory.keto.relation_tuples.v1alpha2
// file: ory/keto/relation_tuples/v1alpha2/read_service.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "grpc";
import * as ory_keto_relation_tuples_v1alpha2_read_service_pb from "../../../../ory/keto/relation_tuples/v1alpha2/read_service_pb";
import * as ory_keto_relation_tuples_v1alpha2_relation_tuples_pb from "../../../../ory/keto/relation_tuples/v1alpha2/relation_tuples_pb";
import * as google_protobuf_field_mask_pb from "google-protobuf/google/protobuf/field_mask_pb";

interface IReadServiceService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    listRelationTuples: IReadServiceService_IListRelationTuples;
}

interface IReadServiceService_IListRelationTuples extends grpc.MethodDefinition<ory_keto_relation_tuples_v1alpha2_read_service_pb.ListRelationTuplesRequest, ory_keto_relation_tuples_v1alpha2_read_service_pb.ListRelationTuplesResponse> {
    path: "/ory.keto.relation_tuples.v1alpha2.ReadService/ListRelationTuples";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<ory_keto_relation_tuples_v1alpha2_read_service_pb.ListRelationTuplesRequest>;
    requestDeserialize: grpc.deserialize<ory_keto_relation_tuples_v1alpha2_read_service_pb.ListRelationTuplesRequest>;
    responseSerialize: grpc.serialize<ory_keto_relation_tuples_v1alpha2_read_service_pb.ListRelationTuplesResponse>;
    responseDeserialize: grpc.deserialize<ory_keto_relation_tuples_v1alpha2_read_service_pb.ListRelationTuplesResponse>;
}

export const ReadServiceService: IReadServiceService;

export interface IReadServiceServer {
    listRelationTuples: grpc.handleUnaryCall<ory_keto_relation_tuples_v1alpha2_read_service_pb.ListRelationTuplesRequest, ory_keto_relation_tuples_v1alpha2_read_service_pb.ListRelationTuplesResponse>;
}

export interface IReadServiceClient {
    listRelationTuples(request: ory_keto_relation_tuples_v1alpha2_read_service_pb.ListRelationTuplesRequest, callback: (error: grpc.ServiceError | null, response: ory_keto_relation_tuples_v1alpha2_read_service_pb.ListRelationTuplesResponse) => void): grpc.ClientUnaryCall;
    listRelationTuples(request: ory_keto_relation_tuples_v1alpha2_read_service_pb.ListRelationTuplesRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: ory_keto_relation_tuples_v1alpha2_read_service_pb.ListRelationTuplesResponse) => void): grpc.ClientUnaryCall;
    listRelationTuples(request: ory_keto_relation_tuples_v1alpha2_read_service_pb.ListRelationTuplesRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: ory_keto_relation_tuples_v1alpha2_read_service_pb.ListRelationTuplesResponse) => void): grpc.ClientUnaryCall;
}

export class ReadServiceClient extends grpc.Client implements IReadServiceClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
    public listRelationTuples(request: ory_keto_relation_tuples_v1alpha2_read_service_pb.ListRelationTuplesRequest, callback: (error: grpc.ServiceError | null, response: ory_keto_relation_tuples_v1alpha2_read_service_pb.ListRelationTuplesResponse) => void): grpc.ClientUnaryCall;
    public listRelationTuples(request: ory_keto_relation_tuples_v1alpha2_read_service_pb.ListRelationTuplesRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: ory_keto_relation_tuples_v1alpha2_read_service_pb.ListRelationTuplesResponse) => void): grpc.ClientUnaryCall;
    public listRelationTuples(request: ory_keto_relation_tuples_v1alpha2_read_service_pb.ListRelationTuplesRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: ory_keto_relation_tuples_v1alpha2_read_service_pb.ListRelationTuplesResponse) => void): grpc.ClientUnaryCall;
}
