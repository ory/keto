/* eslint-disable */
import type { CallContext, CallOptions } from "nice-grpc-common";
import _m0 from "protobufjs/minimal.js";
import { RelationQuery, RelationTuple, Subject } from "./relation_tuples.js";

export const protobufPackage = "ory.keto.relation_tuples.v1alpha2";

/** The request of a WriteService.TransactRelationTuples RPC. */
export interface TransactRelationTuplesRequest {
  /**
   * The write delta for the relationships operated in one single transaction.
   * Either all actions succeed or no change takes effect on error.
   */
  relationTupleDeltas: RelationTupleDelta[];
}

/** Write-delta for a TransactRelationTuplesRequest. */
export interface RelationTupleDelta {
  /** The action to do on the RelationTuple. */
  action: RelationTupleDelta_Action;
  /** The target RelationTuple. */
  relationTuple: RelationTuple | undefined;
}

export enum RelationTupleDelta_Action {
  /**
   * ACTION_UNSPECIFIED - Unspecified.
   * The `TransactRelationTuples` RPC ignores this
   * RelationTupleDelta if an action was unspecified.
   */
  ACTION_UNSPECIFIED = 0,
  /**
   * ACTION_INSERT - Insertion of a new RelationTuple.
   * It is ignored if already existing.
   */
  ACTION_INSERT = 1,
  /**
   * ACTION_DELETE - Deletion of the RelationTuple.
   * It is ignored if it does not exist.
   */
  ACTION_DELETE = 2,
  UNRECOGNIZED = -1,
}

export function relationTupleDelta_ActionFromJSON(object: any): RelationTupleDelta_Action {
  switch (object) {
    case 0:
    case "ACTION_UNSPECIFIED":
      return RelationTupleDelta_Action.ACTION_UNSPECIFIED;
    case 1:
    case "ACTION_INSERT":
      return RelationTupleDelta_Action.ACTION_INSERT;
    case 2:
    case "ACTION_DELETE":
      return RelationTupleDelta_Action.ACTION_DELETE;
    case -1:
    case "UNRECOGNIZED":
    default:
      return RelationTupleDelta_Action.UNRECOGNIZED;
  }
}

export function relationTupleDelta_ActionToJSON(object: RelationTupleDelta_Action): string {
  switch (object) {
    case RelationTupleDelta_Action.ACTION_UNSPECIFIED:
      return "ACTION_UNSPECIFIED";
    case RelationTupleDelta_Action.ACTION_INSERT:
      return "ACTION_INSERT";
    case RelationTupleDelta_Action.ACTION_DELETE:
      return "ACTION_DELETE";
    case RelationTupleDelta_Action.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

/** The response of a WriteService.TransactRelationTuples rpc. */
export interface TransactRelationTuplesResponse {
  /**
   * This field is not implemented yet and has no effect.
   * <!--
   * The list of the new latest snapshot tokens of the affected RelationTuple,
   * with the same index as specified in the `relation_tuple_deltas` field of
   * the TransactRelationTuplesRequest request.
   *
   * If the RelationTupleDelta_Action was DELETE
   * the snaptoken is empty at the same index.
   * -->
   */
  snaptokens: string[];
}

export interface DeleteRelationTuplesRequest {
  /** @deprecated */
  query: DeleteRelationTuplesRequest_Query | undefined;
  relationQuery: RelationQuery | undefined;
}

/** The query for deleting relationships */
export interface DeleteRelationTuplesRequest_Query {
  /** Optional. The namespace to query. */
  namespace: string;
  /** Optional. The object to query for. */
  object: string;
  /** Optional. The relation to query for. */
  relation: string;
  /** Optional. The subject to query for. */
  subject: Subject | undefined;
}

export interface DeleteRelationTuplesResponse {
}

function createBaseTransactRelationTuplesRequest(): TransactRelationTuplesRequest {
  return { relationTupleDeltas: [] };
}

export const TransactRelationTuplesRequest = {
  encode(message: TransactRelationTuplesRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.relationTupleDeltas) {
      RelationTupleDelta.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): TransactRelationTuplesRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseTransactRelationTuplesRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.relationTupleDeltas.push(RelationTupleDelta.decode(reader, reader.uint32()));
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): TransactRelationTuplesRequest {
    return {
      relationTupleDeltas: globalThis.Array.isArray(object?.relationTupleDeltas)
        ? object.relationTupleDeltas.map((e: any) => RelationTupleDelta.fromJSON(e))
        : [],
    };
  },

  toJSON(message: TransactRelationTuplesRequest): unknown {
    const obj: any = {};
    if (message.relationTupleDeltas?.length) {
      obj.relationTupleDeltas = message.relationTupleDeltas.map((e) => RelationTupleDelta.toJSON(e));
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<TransactRelationTuplesRequest>, I>>(base?: I): TransactRelationTuplesRequest {
    return TransactRelationTuplesRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<TransactRelationTuplesRequest>, I>>(
    object: I,
  ): TransactRelationTuplesRequest {
    const message = createBaseTransactRelationTuplesRequest();
    message.relationTupleDeltas = object.relationTupleDeltas?.map((e) => RelationTupleDelta.fromPartial(e)) || [];
    return message;
  },
};

function createBaseRelationTupleDelta(): RelationTupleDelta {
  return { action: 0, relationTuple: undefined };
}

export const RelationTupleDelta = {
  encode(message: RelationTupleDelta, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.action !== 0) {
      writer.uint32(8).int32(message.action);
    }
    if (message.relationTuple !== undefined) {
      RelationTuple.encode(message.relationTuple, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): RelationTupleDelta {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseRelationTupleDelta();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.action = reader.int32() as any;
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.relationTuple = RelationTuple.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): RelationTupleDelta {
    return {
      action: isSet(object.action) ? relationTupleDelta_ActionFromJSON(object.action) : 0,
      relationTuple: isSet(object.relationTuple) ? RelationTuple.fromJSON(object.relationTuple) : undefined,
    };
  },

  toJSON(message: RelationTupleDelta): unknown {
    const obj: any = {};
    if (message.action !== 0) {
      obj.action = relationTupleDelta_ActionToJSON(message.action);
    }
    if (message.relationTuple !== undefined) {
      obj.relationTuple = RelationTuple.toJSON(message.relationTuple);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<RelationTupleDelta>, I>>(base?: I): RelationTupleDelta {
    return RelationTupleDelta.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<RelationTupleDelta>, I>>(object: I): RelationTupleDelta {
    const message = createBaseRelationTupleDelta();
    message.action = object.action ?? 0;
    message.relationTuple = (object.relationTuple !== undefined && object.relationTuple !== null)
      ? RelationTuple.fromPartial(object.relationTuple)
      : undefined;
    return message;
  },
};

function createBaseTransactRelationTuplesResponse(): TransactRelationTuplesResponse {
  return { snaptokens: [] };
}

export const TransactRelationTuplesResponse = {
  encode(message: TransactRelationTuplesResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.snaptokens) {
      writer.uint32(10).string(v!);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): TransactRelationTuplesResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseTransactRelationTuplesResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.snaptokens.push(reader.string());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): TransactRelationTuplesResponse {
    return {
      snaptokens: globalThis.Array.isArray(object?.snaptokens)
        ? object.snaptokens.map((e: any) => globalThis.String(e))
        : [],
    };
  },

  toJSON(message: TransactRelationTuplesResponse): unknown {
    const obj: any = {};
    if (message.snaptokens?.length) {
      obj.snaptokens = message.snaptokens;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<TransactRelationTuplesResponse>, I>>(base?: I): TransactRelationTuplesResponse {
    return TransactRelationTuplesResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<TransactRelationTuplesResponse>, I>>(
    object: I,
  ): TransactRelationTuplesResponse {
    const message = createBaseTransactRelationTuplesResponse();
    message.snaptokens = object.snaptokens?.map((e) => e) || [];
    return message;
  },
};

function createBaseDeleteRelationTuplesRequest(): DeleteRelationTuplesRequest {
  return { query: undefined, relationQuery: undefined };
}

export const DeleteRelationTuplesRequest = {
  encode(message: DeleteRelationTuplesRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.query !== undefined) {
      DeleteRelationTuplesRequest_Query.encode(message.query, writer.uint32(10).fork()).ldelim();
    }
    if (message.relationQuery !== undefined) {
      RelationQuery.encode(message.relationQuery, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeleteRelationTuplesRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteRelationTuplesRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.query = DeleteRelationTuplesRequest_Query.decode(reader, reader.uint32());
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.relationQuery = RelationQuery.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): DeleteRelationTuplesRequest {
    return {
      query: isSet(object.query) ? DeleteRelationTuplesRequest_Query.fromJSON(object.query) : undefined,
      relationQuery: isSet(object.relationQuery) ? RelationQuery.fromJSON(object.relationQuery) : undefined,
    };
  },

  toJSON(message: DeleteRelationTuplesRequest): unknown {
    const obj: any = {};
    if (message.query !== undefined) {
      obj.query = DeleteRelationTuplesRequest_Query.toJSON(message.query);
    }
    if (message.relationQuery !== undefined) {
      obj.relationQuery = RelationQuery.toJSON(message.relationQuery);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<DeleteRelationTuplesRequest>, I>>(base?: I): DeleteRelationTuplesRequest {
    return DeleteRelationTuplesRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<DeleteRelationTuplesRequest>, I>>(object: I): DeleteRelationTuplesRequest {
    const message = createBaseDeleteRelationTuplesRequest();
    message.query = (object.query !== undefined && object.query !== null)
      ? DeleteRelationTuplesRequest_Query.fromPartial(object.query)
      : undefined;
    message.relationQuery = (object.relationQuery !== undefined && object.relationQuery !== null)
      ? RelationQuery.fromPartial(object.relationQuery)
      : undefined;
    return message;
  },
};

function createBaseDeleteRelationTuplesRequest_Query(): DeleteRelationTuplesRequest_Query {
  return { namespace: "", object: "", relation: "", subject: undefined };
}

export const DeleteRelationTuplesRequest_Query = {
  encode(message: DeleteRelationTuplesRequest_Query, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.namespace !== "") {
      writer.uint32(10).string(message.namespace);
    }
    if (message.object !== "") {
      writer.uint32(18).string(message.object);
    }
    if (message.relation !== "") {
      writer.uint32(26).string(message.relation);
    }
    if (message.subject !== undefined) {
      Subject.encode(message.subject, writer.uint32(34).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeleteRelationTuplesRequest_Query {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteRelationTuplesRequest_Query();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.namespace = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.object = reader.string();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.relation = reader.string();
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.subject = Subject.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): DeleteRelationTuplesRequest_Query {
    return {
      namespace: isSet(object.namespace) ? globalThis.String(object.namespace) : "",
      object: isSet(object.object) ? globalThis.String(object.object) : "",
      relation: isSet(object.relation) ? globalThis.String(object.relation) : "",
      subject: isSet(object.subject) ? Subject.fromJSON(object.subject) : undefined,
    };
  },

  toJSON(message: DeleteRelationTuplesRequest_Query): unknown {
    const obj: any = {};
    if (message.namespace !== "") {
      obj.namespace = message.namespace;
    }
    if (message.object !== "") {
      obj.object = message.object;
    }
    if (message.relation !== "") {
      obj.relation = message.relation;
    }
    if (message.subject !== undefined) {
      obj.subject = Subject.toJSON(message.subject);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<DeleteRelationTuplesRequest_Query>, I>>(
    base?: I,
  ): DeleteRelationTuplesRequest_Query {
    return DeleteRelationTuplesRequest_Query.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<DeleteRelationTuplesRequest_Query>, I>>(
    object: I,
  ): DeleteRelationTuplesRequest_Query {
    const message = createBaseDeleteRelationTuplesRequest_Query();
    message.namespace = object.namespace ?? "";
    message.object = object.object ?? "";
    message.relation = object.relation ?? "";
    message.subject = (object.subject !== undefined && object.subject !== null)
      ? Subject.fromPartial(object.subject)
      : undefined;
    return message;
  },
};

function createBaseDeleteRelationTuplesResponse(): DeleteRelationTuplesResponse {
  return {};
}

export const DeleteRelationTuplesResponse = {
  encode(_: DeleteRelationTuplesResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeleteRelationTuplesResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteRelationTuplesResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(_: any): DeleteRelationTuplesResponse {
    return {};
  },

  toJSON(_: DeleteRelationTuplesResponse): unknown {
    const obj: any = {};
    return obj;
  },

  create<I extends Exact<DeepPartial<DeleteRelationTuplesResponse>, I>>(base?: I): DeleteRelationTuplesResponse {
    return DeleteRelationTuplesResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<DeleteRelationTuplesResponse>, I>>(_: I): DeleteRelationTuplesResponse {
    const message = createBaseDeleteRelationTuplesResponse();
    return message;
  },
};

/**
 * The write service to create and delete Access Control Lists.
 *
 * This service is part of the [write-APIs](../concepts/25_api-overview.mdx#write-apis).
 */
export type WriteServiceDefinition = typeof WriteServiceDefinition;
export const WriteServiceDefinition = {
  name: "WriteService",
  fullName: "ory.keto.relation_tuples.v1alpha2.WriteService",
  methods: {
    /** Writes one or more relationships in a single transaction. */
    transactRelationTuples: {
      name: "TransactRelationTuples",
      requestType: TransactRelationTuplesRequest,
      requestStream: false,
      responseType: TransactRelationTuplesResponse,
      responseStream: false,
      options: {},
    },
    /** Deletes relationships based on relation query */
    deleteRelationTuples: {
      name: "DeleteRelationTuples",
      requestType: DeleteRelationTuplesRequest,
      requestStream: false,
      responseType: DeleteRelationTuplesResponse,
      responseStream: false,
      options: {},
    },
  },
} as const;

export interface WriteServiceImplementation<CallContextExt = {}> {
  /** Writes one or more relationships in a single transaction. */
  transactRelationTuples(
    request: TransactRelationTuplesRequest,
    context: CallContext & CallContextExt,
  ): Promise<DeepPartial<TransactRelationTuplesResponse>>;
  /** Deletes relationships based on relation query */
  deleteRelationTuples(
    request: DeleteRelationTuplesRequest,
    context: CallContext & CallContextExt,
  ): Promise<DeepPartial<DeleteRelationTuplesResponse>>;
}

export interface WriteServiceClient<CallOptionsExt = {}> {
  /** Writes one or more relationships in a single transaction. */
  transactRelationTuples(
    request: DeepPartial<TransactRelationTuplesRequest>,
    options?: CallOptions & CallOptionsExt,
  ): Promise<TransactRelationTuplesResponse>;
  /** Deletes relationships based on relation query */
  deleteRelationTuples(
    request: DeepPartial<DeleteRelationTuplesRequest>,
    options?: CallOptions & CallOptionsExt,
  ): Promise<DeleteRelationTuplesResponse>;
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
