// package: ory.keto.acl.v1alpha1
// file: ory/keto/acl/v1alpha1/version.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";

export class GetVersionRequest extends jspb.Message { 

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): GetVersionRequest.AsObject;
    static toObject(includeInstance: boolean, msg: GetVersionRequest): GetVersionRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: GetVersionRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): GetVersionRequest;
    static deserializeBinaryFromReader(message: GetVersionRequest, reader: jspb.BinaryReader): GetVersionRequest;
}

export namespace GetVersionRequest {
    export type AsObject = {
    }
}

export class GetVersionResponse extends jspb.Message { 
    getVersion(): string;
    setVersion(value: string): GetVersionResponse;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): GetVersionResponse.AsObject;
    static toObject(includeInstance: boolean, msg: GetVersionResponse): GetVersionResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: GetVersionResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): GetVersionResponse;
    static deserializeBinaryFromReader(message: GetVersionResponse, reader: jspb.BinaryReader): GetVersionResponse;
}

export namespace GetVersionResponse {
    export type AsObject = {
        version: string,
    }
}
