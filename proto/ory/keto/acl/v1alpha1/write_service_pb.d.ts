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
