// package: ory.keto.relation_tuples.v1alpha2
// file: ory/keto/relation_tuples/v1alpha2/namespaces_service.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";

export class ListNamespacesRequest extends jspb.Message { 

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ListNamespacesRequest.AsObject;
    static toObject(includeInstance: boolean, msg: ListNamespacesRequest): ListNamespacesRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ListNamespacesRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ListNamespacesRequest;
    static deserializeBinaryFromReader(message: ListNamespacesRequest, reader: jspb.BinaryReader): ListNamespacesRequest;
}

export namespace ListNamespacesRequest {
    export type AsObject = {
    }
}

export class ListNamespacesResponse extends jspb.Message { 
    clearNamespacesList(): void;
    getNamespacesList(): Array<Namespace>;
    setNamespacesList(value: Array<Namespace>): ListNamespacesResponse;
    addNamespaces(value?: Namespace, index?: number): Namespace;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ListNamespacesResponse.AsObject;
    static toObject(includeInstance: boolean, msg: ListNamespacesResponse): ListNamespacesResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ListNamespacesResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ListNamespacesResponse;
    static deserializeBinaryFromReader(message: ListNamespacesResponse, reader: jspb.BinaryReader): ListNamespacesResponse;
}

export namespace ListNamespacesResponse {
    export type AsObject = {
        namespacesList: Array<Namespace.AsObject>,
    }
}

export class Namespace extends jspb.Message { 
    getName(): string;
    setName(value: string): Namespace;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Namespace.AsObject;
    static toObject(includeInstance: boolean, msg: Namespace): Namespace.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Namespace, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Namespace;
    static deserializeBinaryFromReader(message: Namespace, reader: jspb.BinaryReader): Namespace;
}

export namespace Namespace {
    export type AsObject = {
        name: string,
    }
}
