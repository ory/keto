// package: ory.keto.acl.v1alpha1
// file: ory/keto/acl/v1alpha1/write_service.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";
import * as ory_keto_acl_v1alpha1_acl_pb from "../../../../ory/keto/acl/v1alpha1/acl_pb";

export class TransactRelationTuplesRequest extends jspb.Message { 
    clearRelationTupleDeltasList(): void;
    getRelationTupleDeltasList(): Array<RelationTupleDelta>;
    setRelationTupleDeltasList(value: Array<RelationTupleDelta>): TransactRelationTuplesRequest;
    addRelationTupleDeltas(value?: RelationTupleDelta, index?: number): RelationTupleDelta;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): TransactRelationTuplesRequest.AsObject;
    static toObject(includeInstance: boolean, msg: TransactRelationTuplesRequest): TransactRelationTuplesRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: TransactRelationTuplesRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): TransactRelationTuplesRequest;
    static deserializeBinaryFromReader(message: TransactRelationTuplesRequest, reader: jspb.BinaryReader): TransactRelationTuplesRequest;
}

export namespace TransactRelationTuplesRequest {
    export type AsObject = {
        relationTupleDeltasList: Array<RelationTupleDelta.AsObject>,
    }
}

export class RelationTupleDelta extends jspb.Message { 
    getAction(): RelationTupleDelta.Action;
    setAction(value: RelationTupleDelta.Action): RelationTupleDelta;

    hasRelationTuple(): boolean;
    clearRelationTuple(): void;
    getRelationTuple(): ory_keto_acl_v1alpha1_acl_pb.RelationTuple | undefined;
    setRelationTuple(value?: ory_keto_acl_v1alpha1_acl_pb.RelationTuple): RelationTupleDelta;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): RelationTupleDelta.AsObject;
    static toObject(includeInstance: boolean, msg: RelationTupleDelta): RelationTupleDelta.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: RelationTupleDelta, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): RelationTupleDelta;
    static deserializeBinaryFromReader(message: RelationTupleDelta, reader: jspb.BinaryReader): RelationTupleDelta;
}

export namespace RelationTupleDelta {
    export type AsObject = {
        action: RelationTupleDelta.Action,
        relationTuple?: ory_keto_acl_v1alpha1_acl_pb.RelationTuple.AsObject,
    }

    export enum Action {
    ACTION_UNSPECIFIED = 0,
    INSERT = 1,
    DELETE = 2,
    }

}

export class TransactRelationTuplesResponse extends jspb.Message { 
    clearSnaptokensList(): void;
    getSnaptokensList(): Array<string>;
    setSnaptokensList(value: Array<string>): TransactRelationTuplesResponse;
    addSnaptokens(value: string, index?: number): string;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): TransactRelationTuplesResponse.AsObject;
    static toObject(includeInstance: boolean, msg: TransactRelationTuplesResponse): TransactRelationTuplesResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: TransactRelationTuplesResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): TransactRelationTuplesResponse;
    static deserializeBinaryFromReader(message: TransactRelationTuplesResponse, reader: jspb.BinaryReader): TransactRelationTuplesResponse;
}

export namespace TransactRelationTuplesResponse {
    export type AsObject = {
        snaptokensList: Array<string>,
    }
}

export class DeleteRelationTuplesRequest extends jspb.Message { 

    hasQuery(): boolean;
    clearQuery(): void;
    getQuery(): DeleteRelationTuplesRequest.Query | undefined;
    setQuery(value?: DeleteRelationTuplesRequest.Query): DeleteRelationTuplesRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): DeleteRelationTuplesRequest.AsObject;
    static toObject(includeInstance: boolean, msg: DeleteRelationTuplesRequest): DeleteRelationTuplesRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: DeleteRelationTuplesRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): DeleteRelationTuplesRequest;
    static deserializeBinaryFromReader(message: DeleteRelationTuplesRequest, reader: jspb.BinaryReader): DeleteRelationTuplesRequest;
}

export namespace DeleteRelationTuplesRequest {
    export type AsObject = {
        query?: DeleteRelationTuplesRequest.Query.AsObject,
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
        getSubject(): ory_keto_acl_v1alpha1_acl_pb.Subject | undefined;
        setSubject(value?: ory_keto_acl_v1alpha1_acl_pb.Subject): Query;

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
            subject?: ory_keto_acl_v1alpha1_acl_pb.Subject.AsObject,
        }
    }

}

export class DeleteRelationTuplesResponse extends jspb.Message { 

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): DeleteRelationTuplesResponse.AsObject;
    static toObject(includeInstance: boolean, msg: DeleteRelationTuplesResponse): DeleteRelationTuplesResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: DeleteRelationTuplesResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): DeleteRelationTuplesResponse;
    static deserializeBinaryFromReader(message: DeleteRelationTuplesResponse, reader: jspb.BinaryReader): DeleteRelationTuplesResponse;
}

export namespace DeleteRelationTuplesResponse {
    export type AsObject = {
    }
}
