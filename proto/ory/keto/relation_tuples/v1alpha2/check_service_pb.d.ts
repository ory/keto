// package: ory.keto.relation_tuples.v1alpha2
// file: ory/keto/relation_tuples/v1alpha2/check_service.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";
import * as ory_keto_relation_tuples_v1alpha2_relation_tuples_pb from "../../../../ory/keto/relation_tuples/v1alpha2/relation_tuples_pb";

export class CheckRequest extends jspb.Message { 
    getNamespace(): string;
    setNamespace(value: string): CheckRequest;
    getObject(): string;
    setObject(value: string): CheckRequest;
    getRelation(): string;
    setRelation(value: string): CheckRequest;

    hasSubject(): boolean;
    clearSubject(): void;
    getSubject(): ory_keto_relation_tuples_v1alpha2_relation_tuples_pb.Subject | undefined;
    setSubject(value?: ory_keto_relation_tuples_v1alpha2_relation_tuples_pb.Subject): CheckRequest;

    hasTuple(): boolean;
    clearTuple(): void;
    getTuple(): ory_keto_relation_tuples_v1alpha2_relation_tuples_pb.RelationTuple | undefined;
    setTuple(value?: ory_keto_relation_tuples_v1alpha2_relation_tuples_pb.RelationTuple): CheckRequest;
    getLatest(): boolean;
    setLatest(value: boolean): CheckRequest;
    getSnaptoken(): string;
    setSnaptoken(value: string): CheckRequest;
    getMaxDepth(): number;
    setMaxDepth(value: number): CheckRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): CheckRequest.AsObject;
    static toObject(includeInstance: boolean, msg: CheckRequest): CheckRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: CheckRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): CheckRequest;
    static deserializeBinaryFromReader(message: CheckRequest, reader: jspb.BinaryReader): CheckRequest;
}

export namespace CheckRequest {
    export type AsObject = {
        namespace: string,
        object: string,
        relation: string,
        subject?: ory_keto_relation_tuples_v1alpha2_relation_tuples_pb.Subject.AsObject,
        tuple?: ory_keto_relation_tuples_v1alpha2_relation_tuples_pb.RelationTuple.AsObject,
        latest: boolean,
        snaptoken: string,
        maxDepth: number,
    }
}

export class CheckResponse extends jspb.Message { 
    getAllowed(): boolean;
    setAllowed(value: boolean): CheckResponse;
    getSnaptoken(): string;
    setSnaptoken(value: string): CheckResponse;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): CheckResponse.AsObject;
    static toObject(includeInstance: boolean, msg: CheckResponse): CheckResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: CheckResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): CheckResponse;
    static deserializeBinaryFromReader(message: CheckResponse, reader: jspb.BinaryReader): CheckResponse;
}

export namespace CheckResponse {
    export type AsObject = {
        allowed: boolean,
        snaptoken: string,
    }
}
