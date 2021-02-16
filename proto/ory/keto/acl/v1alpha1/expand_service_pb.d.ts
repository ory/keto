// package: ory.keto.acl.v1alpha1
// file: ory/keto/acl/v1alpha1/expand_service.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";
import * as ory_keto_acl_v1alpha1_acl_pb from "../../../../ory/keto/acl/v1alpha1/acl_pb";

export class ExpandRequest extends jspb.Message { 

    hasSubject(): boolean;
    clearSubject(): void;
    getSubject(): ory_keto_acl_v1alpha1_acl_pb.Subject | undefined;
    setSubject(value?: ory_keto_acl_v1alpha1_acl_pb.Subject): ExpandRequest;

    getMaxDepth(): number;
    setMaxDepth(value: number): ExpandRequest;

    getSnaptoken(): string;
    setSnaptoken(value: string): ExpandRequest;


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
        subject?: ory_keto_acl_v1alpha1_acl_pb.Subject.AsObject,
        maxDepth: number,
        snaptoken: string,
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
    getSubject(): ory_keto_acl_v1alpha1_acl_pb.Subject | undefined;
    setSubject(value?: ory_keto_acl_v1alpha1_acl_pb.Subject): SubjectTree;

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
        subject?: ory_keto_acl_v1alpha1_acl_pb.Subject.AsObject,
        childrenList: Array<SubjectTree.AsObject>,
    }
}

export enum NodeType {
    NODE_TYPE_UNSPECIFIED = 0,
    NODE_TYPE_UNION = 1,
    NODE_TYPE_EXCLUSION = 2,
    NODE_TYPE_INTERSECTION = 3,
    NODE_TYPE_LEAF = 4,
}
