// package: ory.keto.relation_tuples.v1alpha2
// file: ory/keto/relation_tuples/v1alpha2/expand_service.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";
import * as google_api_field_behavior_pb from "../../../../google/api/field_behavior_pb";
import * as google_api_visibility_pb from "../../../../google/api/visibility_pb";
import * as ory_keto_relation_tuples_v1alpha2_relation_tuples_pb from "../../../../ory/keto/relation_tuples/v1alpha2/relation_tuples_pb";
import * as protoc_gen_openapiv2_options_annotations_pb from "../../../../protoc-gen-openapiv2/options/annotations_pb";

export class ExpandRequest extends jspb.Message { 

    hasSubject(): boolean;
    clearSubject(): void;
    getSubject(): ory_keto_relation_tuples_v1alpha2_relation_tuples_pb.Subject | undefined;
    setSubject(value?: ory_keto_relation_tuples_v1alpha2_relation_tuples_pb.Subject): ExpandRequest;
    getMaxDepth(): number;
    setMaxDepth(value: number): ExpandRequest;
    getSnaptoken(): string;
    setSnaptoken(value: string): ExpandRequest;
    getNamespace(): string;
    setNamespace(value: string): ExpandRequest;
    getObject(): string;
    setObject(value: string): ExpandRequest;
    getRelation(): string;
    setRelation(value: string): ExpandRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ExpandRequest.AsObject;
    static toObject(includeInstance: boolean, msg: ExpandRequest): ExpandRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ExpandRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ExpandRequest;
    static deserializeBinaryFromReader(message: ExpandRequest, reader: jspb.BinaryReader): ExpandRequest;
}

export namespace ExpandRequest {
    export type AsObject = {
        subject?: ory_keto_relation_tuples_v1alpha2_relation_tuples_pb.Subject.AsObject,
        maxDepth: number,
        snaptoken: string,
        namespace: string,
        object: string,
        relation: string,
    }
}

export class ExpandResponse extends jspb.Message { 

    hasTree(): boolean;
    clearTree(): void;
    getTree(): SubjectTree | undefined;
    setTree(value?: SubjectTree): ExpandResponse;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ExpandResponse.AsObject;
    static toObject(includeInstance: boolean, msg: ExpandResponse): ExpandResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ExpandResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ExpandResponse;
    static deserializeBinaryFromReader(message: ExpandResponse, reader: jspb.BinaryReader): ExpandResponse;
}

export namespace ExpandResponse {
    export type AsObject = {
        tree?: SubjectTree.AsObject,
    }
}

export class SubjectTree extends jspb.Message { 
    getNodeType(): NodeType;
    setNodeType(value: NodeType): SubjectTree;

    hasSubject(): boolean;
    clearSubject(): void;
    getSubject(): ory_keto_relation_tuples_v1alpha2_relation_tuples_pb.Subject | undefined;
    setSubject(value?: ory_keto_relation_tuples_v1alpha2_relation_tuples_pb.Subject): SubjectTree;

    hasTuple(): boolean;
    clearTuple(): void;
    getTuple(): ory_keto_relation_tuples_v1alpha2_relation_tuples_pb.RelationTuple | undefined;
    setTuple(value?: ory_keto_relation_tuples_v1alpha2_relation_tuples_pb.RelationTuple): SubjectTree;
    clearChildrenList(): void;
    getChildrenList(): Array<SubjectTree>;
    setChildrenList(value: Array<SubjectTree>): SubjectTree;
    addChildren(value?: SubjectTree, index?: number): SubjectTree;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): SubjectTree.AsObject;
    static toObject(includeInstance: boolean, msg: SubjectTree): SubjectTree.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: SubjectTree, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): SubjectTree;
    static deserializeBinaryFromReader(message: SubjectTree, reader: jspb.BinaryReader): SubjectTree;
}

export namespace SubjectTree {
    export type AsObject = {
        nodeType: NodeType,
        subject?: ory_keto_relation_tuples_v1alpha2_relation_tuples_pb.Subject.AsObject,
        tuple?: ory_keto_relation_tuples_v1alpha2_relation_tuples_pb.RelationTuple.AsObject,
        childrenList: Array<SubjectTree.AsObject>,
    }
}

export enum NodeType {
    UNSPECIFIED = 0,
    NODE_TYPE_UNSPECIFIED = 0,
    UNION = 1,
    NODE_TYPE_UNION = 1,
    EXCLUSION = 2,
    NODE_TYPE_EXCLUSION = 2,
    INTERSECTION = 3,
    NODE_TYPE_INTERSECTION = 3,
    LEAF = 4,
    NODE_TYPE_LEAF = 4,
    TUPLE_TO_SUBJECT_SET = 5,
    NODE_TYPE_TUPLE_TO_SUBJECT_SET = 5,
    COMPUTED_SUBJECT_SET = 6,
    NODE_TYPE_COMPUTED_SUBJECT_SET = 6,
    NOT = 7,
    NODE_TYPE_NOT = 7,
}
