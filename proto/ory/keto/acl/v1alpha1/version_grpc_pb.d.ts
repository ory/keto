// package: ory.keto.acl.v1alpha1
// file: ory/keto/acl/v1alpha1/version.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "grpc";
import * as ory_keto_acl_v1alpha1_version_pb from "../../../../ory/keto/acl/v1alpha1/version_pb";

interface IVersionServiceService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    getVersion: IVersionServiceService_IGetVersion;
}

interface IVersionServiceService_IGetVersion extends grpc.MethodDefinition<ory_keto_acl_v1alpha1_version_pb.GetVersionRequest, ory_keto_acl_v1alpha1_version_pb.GetVersionResponse> {
    path: "/ory.keto.acl.v1alpha1.VersionService/GetVersion";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<ory_keto_acl_v1alpha1_version_pb.GetVersionRequest>;
    requestDeserialize: grpc.deserialize<ory_keto_acl_v1alpha1_version_pb.GetVersionRequest>;
    responseSerialize: grpc.serialize<ory_keto_acl_v1alpha1_version_pb.GetVersionResponse>;
    responseDeserialize: grpc.deserialize<ory_keto_acl_v1alpha1_version_pb.GetVersionResponse>;
}

export const VersionServiceService: IVersionServiceService;

export interface IVersionServiceServer {
    getVersion: grpc.handleUnaryCall<ory_keto_acl_v1alpha1_version_pb.GetVersionRequest, ory_keto_acl_v1alpha1_version_pb.GetVersionResponse>;
}

export interface IVersionServiceClient {
    getVersion(request: ory_keto_acl_v1alpha1_version_pb.GetVersionRequest, callback: (error: grpc.ServiceError | null, response: ory_keto_acl_v1alpha1_version_pb.GetVersionResponse) => void): grpc.ClientUnaryCall;
    getVersion(request: ory_keto_acl_v1alpha1_version_pb.GetVersionRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: ory_keto_acl_v1alpha1_version_pb.GetVersionResponse) => void): grpc.ClientUnaryCall;
    getVersion(request: ory_keto_acl_v1alpha1_version_pb.GetVersionRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: ory_keto_acl_v1alpha1_version_pb.GetVersionResponse) => void): grpc.ClientUnaryCall;
}

export class VersionServiceClient extends grpc.Client implements IVersionServiceClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
    public getVersion(request: ory_keto_acl_v1alpha1_version_pb.GetVersionRequest, callback: (error: grpc.ServiceError | null, response: ory_keto_acl_v1alpha1_version_pb.GetVersionResponse) => void): grpc.ClientUnaryCall;
    public getVersion(request: ory_keto_acl_v1alpha1_version_pb.GetVersionRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: ory_keto_acl_v1alpha1_version_pb.GetVersionResponse) => void): grpc.ClientUnaryCall;
    public getVersion(request: ory_keto_acl_v1alpha1_version_pb.GetVersionRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: ory_keto_acl_v1alpha1_version_pb.GetVersionResponse) => void): grpc.ClientUnaryCall;
}
