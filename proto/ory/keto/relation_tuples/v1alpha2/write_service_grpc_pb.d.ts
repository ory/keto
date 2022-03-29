// package: ory.keto.relation_tuples.v1alpha2
// file: ory/keto/relation_tuples/v1alpha2/write_service.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "grpc";
import * as ory_keto_relation_tuples_v1alpha2_write_service_pb from "../../../../ory/keto/relation_tuples/v1alpha2/write_service_pb";
import * as ory_keto_relation_tuples_v1alpha2_relation_tuples_pb from "../../../../ory/keto/relation_tuples/v1alpha2/relation_tuples_pb";

interface IWriteServiceService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    transactRelationTuples: IWriteServiceService_ITransactRelationTuples;
    deleteRelationTuples: IWriteServiceService_IDeleteRelationTuples;
}

interface IWriteServiceService_ITransactRelationTuples extends grpc.MethodDefinition<ory_keto_relation_tuples_v1alpha2_write_service_pb.TransactRelationTuplesRequest, ory_keto_relation_tuples_v1alpha2_write_service_pb.TransactRelationTuplesResponse> {
    path: "/ory.keto.relation_tuples.v1alpha2.WriteService/TransactRelationTuples";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<ory_keto_relation_tuples_v1alpha2_write_service_pb.TransactRelationTuplesRequest>;
    requestDeserialize: grpc.deserialize<ory_keto_relation_tuples_v1alpha2_write_service_pb.TransactRelationTuplesRequest>;
    responseSerialize: grpc.serialize<ory_keto_relation_tuples_v1alpha2_write_service_pb.TransactRelationTuplesResponse>;
    responseDeserialize: grpc.deserialize<ory_keto_relation_tuples_v1alpha2_write_service_pb.TransactRelationTuplesResponse>;
}
interface IWriteServiceService_IDeleteRelationTuples extends grpc.MethodDefinition<ory_keto_relation_tuples_v1alpha2_write_service_pb.DeleteRelationTuplesRequest, ory_keto_relation_tuples_v1alpha2_write_service_pb.DeleteRelationTuplesResponse> {
    path: "/ory.keto.relation_tuples.v1alpha2.WriteService/DeleteRelationTuples";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<ory_keto_relation_tuples_v1alpha2_write_service_pb.DeleteRelationTuplesRequest>;
    requestDeserialize: grpc.deserialize<ory_keto_relation_tuples_v1alpha2_write_service_pb.DeleteRelationTuplesRequest>;
    responseSerialize: grpc.serialize<ory_keto_relation_tuples_v1alpha2_write_service_pb.DeleteRelationTuplesResponse>;
    responseDeserialize: grpc.deserialize<ory_keto_relation_tuples_v1alpha2_write_service_pb.DeleteRelationTuplesResponse>;
}

export const WriteServiceService: IWriteServiceService;

export interface IWriteServiceServer {
    transactRelationTuples: grpc.handleUnaryCall<ory_keto_relation_tuples_v1alpha2_write_service_pb.TransactRelationTuplesRequest, ory_keto_relation_tuples_v1alpha2_write_service_pb.TransactRelationTuplesResponse>;
    deleteRelationTuples: grpc.handleUnaryCall<ory_keto_relation_tuples_v1alpha2_write_service_pb.DeleteRelationTuplesRequest, ory_keto_relation_tuples_v1alpha2_write_service_pb.DeleteRelationTuplesResponse>;
}

export interface IWriteServiceClient {
    transactRelationTuples(request: ory_keto_relation_tuples_v1alpha2_write_service_pb.TransactRelationTuplesRequest, callback: (error: grpc.ServiceError | null, response: ory_keto_relation_tuples_v1alpha2_write_service_pb.TransactRelationTuplesResponse) => void): grpc.ClientUnaryCall;
    transactRelationTuples(request: ory_keto_relation_tuples_v1alpha2_write_service_pb.TransactRelationTuplesRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: ory_keto_relation_tuples_v1alpha2_write_service_pb.TransactRelationTuplesResponse) => void): grpc.ClientUnaryCall;
    transactRelationTuples(request: ory_keto_relation_tuples_v1alpha2_write_service_pb.TransactRelationTuplesRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: ory_keto_relation_tuples_v1alpha2_write_service_pb.TransactRelationTuplesResponse) => void): grpc.ClientUnaryCall;
    deleteRelationTuples(request: ory_keto_relation_tuples_v1alpha2_write_service_pb.DeleteRelationTuplesRequest, callback: (error: grpc.ServiceError | null, response: ory_keto_relation_tuples_v1alpha2_write_service_pb.DeleteRelationTuplesResponse) => void): grpc.ClientUnaryCall;
    deleteRelationTuples(request: ory_keto_relation_tuples_v1alpha2_write_service_pb.DeleteRelationTuplesRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: ory_keto_relation_tuples_v1alpha2_write_service_pb.DeleteRelationTuplesResponse) => void): grpc.ClientUnaryCall;
    deleteRelationTuples(request: ory_keto_relation_tuples_v1alpha2_write_service_pb.DeleteRelationTuplesRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: ory_keto_relation_tuples_v1alpha2_write_service_pb.DeleteRelationTuplesResponse) => void): grpc.ClientUnaryCall;
}

export class WriteServiceClient extends grpc.Client implements IWriteServiceClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
    public transactRelationTuples(request: ory_keto_relation_tuples_v1alpha2_write_service_pb.TransactRelationTuplesRequest, callback: (error: grpc.ServiceError | null, response: ory_keto_relation_tuples_v1alpha2_write_service_pb.TransactRelationTuplesResponse) => void): grpc.ClientUnaryCall;
    public transactRelationTuples(request: ory_keto_relation_tuples_v1alpha2_write_service_pb.TransactRelationTuplesRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: ory_keto_relation_tuples_v1alpha2_write_service_pb.TransactRelationTuplesResponse) => void): grpc.ClientUnaryCall;
    public transactRelationTuples(request: ory_keto_relation_tuples_v1alpha2_write_service_pb.TransactRelationTuplesRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: ory_keto_relation_tuples_v1alpha2_write_service_pb.TransactRelationTuplesResponse) => void): grpc.ClientUnaryCall;
    public deleteRelationTuples(request: ory_keto_relation_tuples_v1alpha2_write_service_pb.DeleteRelationTuplesRequest, callback: (error: grpc.ServiceError | null, response: ory_keto_relation_tuples_v1alpha2_write_service_pb.DeleteRelationTuplesResponse) => void): grpc.ClientUnaryCall;
    public deleteRelationTuples(request: ory_keto_relation_tuples_v1alpha2_write_service_pb.DeleteRelationTuplesRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: ory_keto_relation_tuples_v1alpha2_write_service_pb.DeleteRelationTuplesResponse) => void): grpc.ClientUnaryCall;
    public deleteRelationTuples(request: ory_keto_relation_tuples_v1alpha2_write_service_pb.DeleteRelationTuplesRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: ory_keto_relation_tuples_v1alpha2_write_service_pb.DeleteRelationTuplesResponse) => void): grpc.ClientUnaryCall;
}
