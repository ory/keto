// package: ory.keto.relation_tuples.v1alpha2
// file: ory/keto/relation_tuples/v1alpha2/version.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "@grpc/grpc-js";
import * as ory_keto_relation_tuples_v1alpha2_version_pb from "../../../../ory/keto/relation_tuples/v1alpha2/version_pb";

interface IVersionServiceService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    getVersion: IVersionServiceService_IGetVersion;
}

interface IVersionServiceService_IGetVersion extends grpc.MethodDefinition<ory_keto_relation_tuples_v1alpha2_version_pb.GetVersionRequest, ory_keto_relation_tuples_v1alpha2_version_pb.GetVersionResponse> {
    path: "/ory.keto.relation_tuples.v1alpha2.VersionService/GetVersion";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<ory_keto_relation_tuples_v1alpha2_version_pb.GetVersionRequest>;
    requestDeserialize: grpc.deserialize<ory_keto_relation_tuples_v1alpha2_version_pb.GetVersionRequest>;
    responseSerialize: grpc.serialize<ory_keto_relation_tuples_v1alpha2_version_pb.GetVersionResponse>;
    responseDeserialize: grpc.deserialize<ory_keto_relation_tuples_v1alpha2_version_pb.GetVersionResponse>;
}

export const VersionServiceService: IVersionServiceService;

export interface IVersionServiceServer extends grpc.UntypedServiceImplementation {
    getVersion: grpc.handleUnaryCall<ory_keto_relation_tuples_v1alpha2_version_pb.GetVersionRequest, ory_keto_relation_tuples_v1alpha2_version_pb.GetVersionResponse>;
}

export interface IVersionServiceClient {
    getVersion(request: ory_keto_relation_tuples_v1alpha2_version_pb.GetVersionRequest, callback: (error: grpc.ServiceError | null, response: ory_keto_relation_tuples_v1alpha2_version_pb.GetVersionResponse) => void): grpc.ClientUnaryCall;
    getVersion(request: ory_keto_relation_tuples_v1alpha2_version_pb.GetVersionRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: ory_keto_relation_tuples_v1alpha2_version_pb.GetVersionResponse) => void): grpc.ClientUnaryCall;
    getVersion(request: ory_keto_relation_tuples_v1alpha2_version_pb.GetVersionRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: ory_keto_relation_tuples_v1alpha2_version_pb.GetVersionResponse) => void): grpc.ClientUnaryCall;
}

export class VersionServiceClient extends grpc.Client implements IVersionServiceClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: Partial<grpc.ClientOptions>);
    public getVersion(request: ory_keto_relation_tuples_v1alpha2_version_pb.GetVersionRequest, callback: (error: grpc.ServiceError | null, response: ory_keto_relation_tuples_v1alpha2_version_pb.GetVersionResponse) => void): grpc.ClientUnaryCall;
    public getVersion(request: ory_keto_relation_tuples_v1alpha2_version_pb.GetVersionRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: ory_keto_relation_tuples_v1alpha2_version_pb.GetVersionResponse) => void): grpc.ClientUnaryCall;
    public getVersion(request: ory_keto_relation_tuples_v1alpha2_version_pb.GetVersionRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: ory_keto_relation_tuples_v1alpha2_version_pb.GetVersionResponse) => void): grpc.ClientUnaryCall;
}
