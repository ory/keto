// package: ory.keto.acl.v1alpha1
// file: ory/keto/acl/v1alpha1/write_service.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "grpc";
import * as ory_keto_acl_v1alpha1_write_service_pb from "../../../../ory/keto/acl/v1alpha1/write_service_pb";
import * as ory_keto_acl_v1alpha1_acl_pb from "../../../../ory/keto/acl/v1alpha1/acl_pb";

interface IWriteServiceService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    transactRelationTuples: IWriteServiceService_ITransactRelationTuples;
}

interface IWriteServiceService_ITransactRelationTuples extends grpc.MethodDefinition<ory_keto_acl_v1alpha1_write_service_pb.TransactRelationTuplesRequest, ory_keto_acl_v1alpha1_write_service_pb.TransactRelationTuplesResponse> {
    path: "/ory.keto.acl.v1alpha1.WriteService/TransactRelationTuples";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<ory_keto_acl_v1alpha1_write_service_pb.TransactRelationTuplesRequest>;
    requestDeserialize: grpc.deserialize<ory_keto_acl_v1alpha1_write_service_pb.TransactRelationTuplesRequest>;
    responseSerialize: grpc.serialize<ory_keto_acl_v1alpha1_write_service_pb.TransactRelationTuplesResponse>;
    responseDeserialize: grpc.deserialize<ory_keto_acl_v1alpha1_write_service_pb.TransactRelationTuplesResponse>;
}

export const WriteServiceService: IWriteServiceService;

export interface IWriteServiceServer {
    transactRelationTuples: grpc.handleUnaryCall<ory_keto_acl_v1alpha1_write_service_pb.TransactRelationTuplesRequest, ory_keto_acl_v1alpha1_write_service_pb.TransactRelationTuplesResponse>;
}

export interface IWriteServiceClient {
    transactRelationTuples(request: ory_keto_acl_v1alpha1_write_service_pb.TransactRelationTuplesRequest, callback: (error: grpc.ServiceError | null, response: ory_keto_acl_v1alpha1_write_service_pb.TransactRelationTuplesResponse) => void): grpc.ClientUnaryCall;
    transactRelationTuples(request: ory_keto_acl_v1alpha1_write_service_pb.TransactRelationTuplesRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: ory_keto_acl_v1alpha1_write_service_pb.TransactRelationTuplesResponse) => void): grpc.ClientUnaryCall;
    transactRelationTuples(request: ory_keto_acl_v1alpha1_write_service_pb.TransactRelationTuplesRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: ory_keto_acl_v1alpha1_write_service_pb.TransactRelationTuplesResponse) => void): grpc.ClientUnaryCall;
}

export class WriteServiceClient extends grpc.Client implements IWriteServiceClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
    public transactRelationTuples(request: ory_keto_acl_v1alpha1_write_service_pb.TransactRelationTuplesRequest, callback: (error: grpc.ServiceError | null, response: ory_keto_acl_v1alpha1_write_service_pb.TransactRelationTuplesResponse) => void): grpc.ClientUnaryCall;
    public transactRelationTuples(request: ory_keto_acl_v1alpha1_write_service_pb.TransactRelationTuplesRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: ory_keto_acl_v1alpha1_write_service_pb.TransactRelationTuplesResponse) => void): grpc.ClientUnaryCall;
    public transactRelationTuples(request: ory_keto_acl_v1alpha1_write_service_pb.TransactRelationTuplesRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: ory_keto_acl_v1alpha1_write_service_pb.TransactRelationTuplesResponse) => void): grpc.ClientUnaryCall;
}
