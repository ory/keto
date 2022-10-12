// package: ory.keto.relation_tuples.v1alpha2
// file: ory/keto/relation_tuples/v1alpha2/namespaces_service.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "grpc";
import * as ory_keto_relation_tuples_v1alpha2_namespaces_service_pb from "../../../../ory/keto/relation_tuples/v1alpha2/namespaces_service_pb";

interface INamespacesServiceService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    listNamespaces: INamespacesServiceService_IListNamespaces;
}

interface INamespacesServiceService_IListNamespaces extends grpc.MethodDefinition<ory_keto_relation_tuples_v1alpha2_namespaces_service_pb.ListNamespacesRequest, ory_keto_relation_tuples_v1alpha2_namespaces_service_pb.ListNamespacesResponse> {
    path: "/ory.keto.relation_tuples.v1alpha2.NamespacesService/ListNamespaces";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<ory_keto_relation_tuples_v1alpha2_namespaces_service_pb.ListNamespacesRequest>;
    requestDeserialize: grpc.deserialize<ory_keto_relation_tuples_v1alpha2_namespaces_service_pb.ListNamespacesRequest>;
    responseSerialize: grpc.serialize<ory_keto_relation_tuples_v1alpha2_namespaces_service_pb.ListNamespacesResponse>;
    responseDeserialize: grpc.deserialize<ory_keto_relation_tuples_v1alpha2_namespaces_service_pb.ListNamespacesResponse>;
}

export const NamespacesServiceService: INamespacesServiceService;

export interface INamespacesServiceServer {
    listNamespaces: grpc.handleUnaryCall<ory_keto_relation_tuples_v1alpha2_namespaces_service_pb.ListNamespacesRequest, ory_keto_relation_tuples_v1alpha2_namespaces_service_pb.ListNamespacesResponse>;
}

export interface INamespacesServiceClient {
    listNamespaces(request: ory_keto_relation_tuples_v1alpha2_namespaces_service_pb.ListNamespacesRequest, callback: (error: grpc.ServiceError | null, response: ory_keto_relation_tuples_v1alpha2_namespaces_service_pb.ListNamespacesResponse) => void): grpc.ClientUnaryCall;
    listNamespaces(request: ory_keto_relation_tuples_v1alpha2_namespaces_service_pb.ListNamespacesRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: ory_keto_relation_tuples_v1alpha2_namespaces_service_pb.ListNamespacesResponse) => void): grpc.ClientUnaryCall;
    listNamespaces(request: ory_keto_relation_tuples_v1alpha2_namespaces_service_pb.ListNamespacesRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: ory_keto_relation_tuples_v1alpha2_namespaces_service_pb.ListNamespacesResponse) => void): grpc.ClientUnaryCall;
}

export class NamespacesServiceClient extends grpc.Client implements INamespacesServiceClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
    public listNamespaces(request: ory_keto_relation_tuples_v1alpha2_namespaces_service_pb.ListNamespacesRequest, callback: (error: grpc.ServiceError | null, response: ory_keto_relation_tuples_v1alpha2_namespaces_service_pb.ListNamespacesResponse) => void): grpc.ClientUnaryCall;
    public listNamespaces(request: ory_keto_relation_tuples_v1alpha2_namespaces_service_pb.ListNamespacesRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: ory_keto_relation_tuples_v1alpha2_namespaces_service_pb.ListNamespacesResponse) => void): grpc.ClientUnaryCall;
    public listNamespaces(request: ory_keto_relation_tuples_v1alpha2_namespaces_service_pb.ListNamespacesRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: ory_keto_relation_tuples_v1alpha2_namespaces_service_pb.ListNamespacesResponse) => void): grpc.ClientUnaryCall;
}
