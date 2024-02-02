/* eslint-disable */
import type { CallContext, CallOptions } from "nice-grpc-common";
import _m0 from "protobufjs/minimal.js";
import { RelationTuple, Subject } from "./relation_tuples.js";

export const protobufPackage = "ory.keto.relation_tuples.v1alpha2";

export enum NodeType {
  NODE_TYPE_UNSPECIFIED = 0,
  /** NODE_TYPE_UNION - This node expands to a union of all children. */
  NODE_TYPE_UNION = 1,
  /** NODE_TYPE_EXCLUSION - Not implemented yet. */
  NODE_TYPE_EXCLUSION = 2,
  /** NODE_TYPE_INTERSECTION - Not implemented yet. */
  NODE_TYPE_INTERSECTION = 3,
  /**
   * NODE_TYPE_LEAF - This node is a leaf and contains no children.
   * Its subject is a `SubjectID` unless `max_depth` was reached.
   */
  NODE_TYPE_LEAF = 4,
  UNRECOGNIZED = -1,
}

export function nodeTypeFromJSON(object: any): NodeType {
  switch (object) {
    case 0:
    case "NODE_TYPE_UNSPECIFIED":
      return NodeType.NODE_TYPE_UNSPECIFIED;
    case 1:
    case "NODE_TYPE_UNION":
      return NodeType.NODE_TYPE_UNION;
    case 2:
    case "NODE_TYPE_EXCLUSION":
      return NodeType.NODE_TYPE_EXCLUSION;
    case 3:
    case "NODE_TYPE_INTERSECTION":
      return NodeType.NODE_TYPE_INTERSECTION;
    case 4:
    case "NODE_TYPE_LEAF":
      return NodeType.NODE_TYPE_LEAF;
    case -1:
    case "UNRECOGNIZED":
    default:
      return NodeType.UNRECOGNIZED;
  }
}

export function nodeTypeToJSON(object: NodeType): string {
  switch (object) {
    case NodeType.NODE_TYPE_UNSPECIFIED:
      return "NODE_TYPE_UNSPECIFIED";
    case NodeType.NODE_TYPE_UNION:
      return "NODE_TYPE_UNION";
    case NodeType.NODE_TYPE_EXCLUSION:
      return "NODE_TYPE_EXCLUSION";
    case NodeType.NODE_TYPE_INTERSECTION:
      return "NODE_TYPE_INTERSECTION";
    case NodeType.NODE_TYPE_LEAF:
      return "NODE_TYPE_LEAF";
    case NodeType.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

/**
 * The request for an ExpandService.Expand RPC.
 * Expands the given subject set.
 */
export interface ExpandRequest {
  /** The subject to expand. */
  subject:
    | Subject
    | undefined;
  /**
   * The maximum depth of tree to build.
   *
   * If the value is less than 1 or greater than the global
   * max-depth then the global max-depth will be used instead.
   *
   * It is important to set this parameter to a meaningful
   * value. Ponder how deep you really want to display this.
   */
  maxDepth: number;
  /**
   * This field is not implemented yet and has no effect.
   * <!--
   * Optional. Like reads, a expand is always evaluated at a
   * consistent snapshot no earlier than the given snaptoken.
   *
   * Leave this field blank if you want to expand
   * based on eventually consistent ACLs, benefiting from very
   * low latency, but possibly slightly stale results.
   *
   * If the specified token is too old and no longer known,
   * the server falls back as if no snaptoken had been specified.
   *
   * If not specified the server tries to build the tree
   * on the best snapshot version where it is very likely that
   * ACLs had already been replicated to all availability zones.
   * -->
   */
  snaptoken: string;
}

/** The response for a ExpandService.Expand RPC. */
export interface ExpandResponse {
  /**
   * The tree the requested subject set expands to.
   * The requested subject set is the subject of the root.
   *
   * This field can be nil in some circumstances.
   */
  tree: SubjectTree | undefined;
}

export interface SubjectTree {
  /** The type of the node. */
  nodeType: NodeType;
  /**
   * The subject this node represents.
   * Deprecated: More information is now available in the tuple field.
   *
   * @deprecated
   */
  subject:
    | Subject
    | undefined;
  /** The relation tuple this node represents. */
  tuple:
    | RelationTuple
    | undefined;
  /**
   * The children of this node.
   *
   * This is never set if `node_type` == `NODE_TYPE_LEAF`.
   */
  children: SubjectTree[];
}

function createBaseExpandRequest(): ExpandRequest {
  return { subject: undefined, maxDepth: 0, snaptoken: "" };
}

export const ExpandRequest = {
  encode(message: ExpandRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.subject !== undefined) {
      Subject.encode(message.subject, writer.uint32(10).fork()).ldelim();
    }
    if (message.maxDepth !== 0) {
      writer.uint32(16).int32(message.maxDepth);
    }
    if (message.snaptoken !== "") {
      writer.uint32(26).string(message.snaptoken);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ExpandRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseExpandRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.subject = Subject.decode(reader, reader.uint32());
          continue;
        case 2:
          if (tag !== 16) {
            break;
          }

          message.maxDepth = reader.int32();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.snaptoken = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ExpandRequest {
    return {
      subject: isSet(object.subject) ? Subject.fromJSON(object.subject) : undefined,
      maxDepth: isSet(object.maxDepth) ? globalThis.Number(object.maxDepth) : 0,
      snaptoken: isSet(object.snaptoken) ? globalThis.String(object.snaptoken) : "",
    };
  },

  toJSON(message: ExpandRequest): unknown {
    const obj: any = {};
    if (message.subject !== undefined) {
      obj.subject = Subject.toJSON(message.subject);
    }
    if (message.maxDepth !== 0) {
      obj.maxDepth = Math.round(message.maxDepth);
    }
    if (message.snaptoken !== "") {
      obj.snaptoken = message.snaptoken;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<ExpandRequest>, I>>(base?: I): ExpandRequest {
    return ExpandRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<ExpandRequest>, I>>(object: I): ExpandRequest {
    const message = createBaseExpandRequest();
    message.subject = (object.subject !== undefined && object.subject !== null)
      ? Subject.fromPartial(object.subject)
      : undefined;
    message.maxDepth = object.maxDepth ?? 0;
    message.snaptoken = object.snaptoken ?? "";
    return message;
  },
};

function createBaseExpandResponse(): ExpandResponse {
  return { tree: undefined };
}

export const ExpandResponse = {
  encode(message: ExpandResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.tree !== undefined) {
      SubjectTree.encode(message.tree, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ExpandResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseExpandResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.tree = SubjectTree.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ExpandResponse {
    return { tree: isSet(object.tree) ? SubjectTree.fromJSON(object.tree) : undefined };
  },

  toJSON(message: ExpandResponse): unknown {
    const obj: any = {};
    if (message.tree !== undefined) {
      obj.tree = SubjectTree.toJSON(message.tree);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<ExpandResponse>, I>>(base?: I): ExpandResponse {
    return ExpandResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<ExpandResponse>, I>>(object: I): ExpandResponse {
    const message = createBaseExpandResponse();
    message.tree = (object.tree !== undefined && object.tree !== null)
      ? SubjectTree.fromPartial(object.tree)
      : undefined;
    return message;
  },
};

function createBaseSubjectTree(): SubjectTree {
  return { nodeType: 0, subject: undefined, tuple: undefined, children: [] };
}

export const SubjectTree = {
  encode(message: SubjectTree, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.nodeType !== 0) {
      writer.uint32(8).int32(message.nodeType);
    }
    if (message.subject !== undefined) {
      Subject.encode(message.subject, writer.uint32(18).fork()).ldelim();
    }
    if (message.tuple !== undefined) {
      RelationTuple.encode(message.tuple, writer.uint32(34).fork()).ldelim();
    }
    for (const v of message.children) {
      SubjectTree.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SubjectTree {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSubjectTree();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.nodeType = reader.int32() as any;
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.subject = Subject.decode(reader, reader.uint32());
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.tuple = RelationTuple.decode(reader, reader.uint32());
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.children.push(SubjectTree.decode(reader, reader.uint32()));
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): SubjectTree {
    return {
      nodeType: isSet(object.nodeType) ? nodeTypeFromJSON(object.nodeType) : 0,
      subject: isSet(object.subject) ? Subject.fromJSON(object.subject) : undefined,
      tuple: isSet(object.tuple) ? RelationTuple.fromJSON(object.tuple) : undefined,
      children: globalThis.Array.isArray(object?.children)
        ? object.children.map((e: any) => SubjectTree.fromJSON(e))
        : [],
    };
  },

  toJSON(message: SubjectTree): unknown {
    const obj: any = {};
    if (message.nodeType !== 0) {
      obj.nodeType = nodeTypeToJSON(message.nodeType);
    }
    if (message.subject !== undefined) {
      obj.subject = Subject.toJSON(message.subject);
    }
    if (message.tuple !== undefined) {
      obj.tuple = RelationTuple.toJSON(message.tuple);
    }
    if (message.children?.length) {
      obj.children = message.children.map((e) => SubjectTree.toJSON(e));
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<SubjectTree>, I>>(base?: I): SubjectTree {
    return SubjectTree.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<SubjectTree>, I>>(object: I): SubjectTree {
    const message = createBaseSubjectTree();
    message.nodeType = object.nodeType ?? 0;
    message.subject = (object.subject !== undefined && object.subject !== null)
      ? Subject.fromPartial(object.subject)
      : undefined;
    message.tuple = (object.tuple !== undefined && object.tuple !== null)
      ? RelationTuple.fromPartial(object.tuple)
      : undefined;
    message.children = object.children?.map((e) => SubjectTree.fromPartial(e)) || [];
    return message;
  },
};

/**
 * The service that performs subject set expansion
 * based on the stored Access Control Lists.
 *
 * This service is part of the [read-APIs](../concepts/25_api-overview.mdx#read-apis).
 */
export type ExpandServiceDefinition = typeof ExpandServiceDefinition;
export const ExpandServiceDefinition = {
  name: "ExpandService",
  fullName: "ory.keto.relation_tuples.v1alpha2.ExpandService",
  methods: {
    /** Expands the subject set into a tree of subjects. */
    expand: {
      name: "Expand",
      requestType: ExpandRequest,
      requestStream: false,
      responseType: ExpandResponse,
      responseStream: false,
      options: {},
    },
  },
} as const;

export interface ExpandServiceImplementation<CallContextExt = {}> {
  /** Expands the subject set into a tree of subjects. */
  expand(request: ExpandRequest, context: CallContext & CallContextExt): Promise<DeepPartial<ExpandResponse>>;
}

export interface ExpandServiceClient<CallOptionsExt = {}> {
  /** Expands the subject set into a tree of subjects. */
  expand(request: DeepPartial<ExpandRequest>, options?: CallOptions & CallOptionsExt): Promise<ExpandResponse>;
}

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends globalThis.Array<infer U> ? globalThis.Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
