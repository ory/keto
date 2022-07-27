// package: ory.keto.relation_tuples.v1alpha2
// file: ory/keto/relation_tuples/v1alpha2/read_service.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";
import * as ory_keto_relation_tuples_v1alpha2_relation_tuples_pb from "../../../../ory/keto/relation_tuples/v1alpha2/relation_tuples_pb";
import * as google_protobuf_field_mask_pb from "google-protobuf/google/protobuf/field_mask_pb";

export class ListRelationTuplesRequest extends jspb.Message { 

    hasQuery(): boolean;
    clearQuery(): void;
    getQuery(): ListRelationTuplesRequest.Query | undefined;
    setQuery(value?: ListRelationTuplesRequest.Query): ListRelationTuplesRequest;

    hasRelationQuery(): boolean;
    clearRelationQuery(): void;
    getRelationQuery(): ory_keto_relation_tuples_v1alpha2_relation_tuples_pb.RelationQuery | undefined;
    setRelationQuery(value?: ory_keto_relation_tuples_v1alpha2_relation_tuples_pb.RelationQuery): ListRelationTuplesRequest;

    hasExpandMask(): boolean;
    clearExpandMask(): void;
    getExpandMask(): google_protobuf_field_mask_pb.FieldMask | undefined;
    setExpandMask(value?: google_protobuf_field_mask_pb.FieldMask): ListRelationTuplesRequest;
    getSnaptoken(): string;
    setSnaptoken(value: string): ListRelationTuplesRequest;
    getPageSize(): number;
    setPageSize(value: number): ListRelationTuplesRequest;
    getPageToken(): string;
    setPageToken(value: string): ListRelationTuplesRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ListRelationTuplesRequest.AsObject;
    static toObject(includeInstance: boolean, msg: ListRelationTuplesRequest): ListRelationTuplesRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ListRelationTuplesRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ListRelationTuplesRequest;
    static deserializeBinaryFromReader(message: ListRelationTuplesRequest, reader: jspb.BinaryReader): ListRelationTuplesRequest;
}

export namespace ListRelationTuplesRequest {
    export type AsObject = {
        query?: ListRelationTuplesRequest.Query.AsObject,
        relationQuery?: ory_keto_relation_tuples_v1alpha2_relation_tuples_pb.RelationQuery.AsObject,
        expandMask?: google_protobuf_field_mask_pb.FieldMask.AsObject,
        snaptoken: string,
        pageSize: number,
        pageToken: string,
    }


    export class Query extends jspb.Message { 
        getNamespace(): string;
        setNamespace(value: string): Query;
        getObject(): string;
        setObject(value: string): Query;
        getRelation(): string;
        setRelation(value: string): Query;

        hasSubject(): boolean;
        clearSubject(): void;
        getSubject(): ory_keto_relation_tuples_v1alpha2_relation_tuples_pb.Subject | undefined;
        setSubject(value?: ory_keto_relation_tuples_v1alpha2_relation_tuples_pb.Subject): Query;

        serializeBinary(): Uint8Array;
        toObject(includeInstance?: boolean): Query.AsObject;
        static toObject(includeInstance: boolean, msg: Query): Query.AsObject;
        static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
        static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
        static serializeBinaryToWriter(message: Query, writer: jspb.BinaryWriter): void;
        static deserializeBinary(bytes: Uint8Array): Query;
        static deserializeBinaryFromReader(message: Query, reader: jspb.BinaryReader): Query;
    }

    export namespace Query {
        export type AsObject = {
            namespace: string,
            object: string,
            relation: string,
            subject?: ory_keto_relation_tuples_v1alpha2_relation_tuples_pb.Subject.AsObject,
        }
    }

}

export class ListRelationTuplesResponse extends jspb.Message { 
    clearRelationTuplesList(): void;
    getRelationTuplesList(): Array<ory_keto_relation_tuples_v1alpha2_relation_tuples_pb.RelationTuple>;
    setRelationTuplesList(value: Array<ory_keto_relation_tuples_v1alpha2_relation_tuples_pb.RelationTuple>): ListRelationTuplesResponse;
    addRelationTuples(value?: ory_keto_relation_tuples_v1alpha2_relation_tuples_pb.RelationTuple, index?: number): ory_keto_relation_tuples_v1alpha2_relation_tuples_pb.RelationTuple;
    getNextPageToken(): string;
    setNextPageToken(value: string): ListRelationTuplesResponse;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ListRelationTuplesResponse.AsObject;
    static toObject(includeInstance: boolean, msg: ListRelationTuplesResponse): ListRelationTuplesResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ListRelationTuplesResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ListRelationTuplesResponse;
    static deserializeBinaryFromReader(message: ListRelationTuplesResponse, reader: jspb.BinaryReader): ListRelationTuplesResponse;
}

export namespace ListRelationTuplesResponse {
    export type AsObject = {
        relationTuplesList: Array<ory_keto_relation_tuples_v1alpha2_relation_tuples_pb.RelationTuple.AsObject>,
        nextPageToken: string,
    }
}
