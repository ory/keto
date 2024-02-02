/* eslint-disable */
import type { CallContext, CallOptions } from "nice-grpc-common";
import _m0 from "protobufjs/minimal.js";
import { FieldMask } from "../../../../google/protobuf/field_mask.js";
import { RelationQuery, RelationTuple, Subject } from "./relation_tuples.js";

export const protobufPackage = "ory.keto.relation_tuples.v1alpha2";

/**
 * Request for ReadService.ListRelationTuples RPC.
 * See `ListRelationTuplesRequest_Query` for how to filter the query.
 */
export interface ListRelationTuplesRequest {
  /**
   * All query constraints are concatenated
   * with a logical AND operator.
   *
   * The RelationTuple list from ListRelationTuplesResponse
   * is ordered from the newest RelationTuple to the oldest.
   *
   * @deprecated
   */
  query: ListRelationTuplesRequest_Query | undefined;
  relationQuery:
    | RelationQuery
    | undefined;
  /**
   * This field is not implemented yet and has no effect.
   * <!--
   * Optional. The list of fields to be expanded
   * in the RelationTuple list returned in `ListRelationTuplesResponse`.
   * Leaving this field unspecified means all fields are expanded.
   *
   * Available fields:
   * "object", "relation", "subject",
   * "namespace", "subject.id", "subject.namespace",
   * "subject.object", "subject.relation"
   * -->
   */
  expandMask:
    | string[]
    | undefined;
  /**
   * This field is not implemented yet and has no effect.
   * <!--
   * Optional. The snapshot token for this read.
   * -->
   */
  snaptoken: string;
  /**
   * Optional. The maximum number of
   * RelationTuples to return in the response.
   *
   * Default: 100
   */
  pageSize: number;
  /**
   * Optional. An opaque pagination token returned from
   * a previous call to `ListRelationTuples` that
   * indicates where the page should start at.
   *
   * An empty token denotes the first page. All successive
   * pages require the token from the previous page.
   */
  pageToken: string;
}

/**
 * The query for listing relationships.
 * Clients can specify any optional field to
 * partially filter for specific relationships.
 *
 * Example use cases (namespace is always required):
 *  - object only: display a list of all permissions referring to a specific object
 *  - relation only: get all groups that have members; get all directories that have content
 *  - object & relation: display all subjects that have a specific permission relation
 *  - subject & relation: display all groups a subject belongs to; display all objects a subject has access to
 *  - object & relation & subject: check whether the relation tuple already exists
 */
export interface ListRelationTuplesRequest_Query {
  /** Required. The namespace to query. */
  namespace: string;
  /** Optional. The object to query for. */
  object: string;
  /** Optional. The relation to query for. */
  relation: string;
  /** Optional. The subject to query for. */
  subject: Subject | undefined;
}

/** The response of a ReadService.ListRelationTuples RPC. */
export interface ListRelationTuplesResponse {
  /** The relationships matching the list request. */
  relationTuples: RelationTuple[];
  /**
   * The token required to get the next page.
   * If this is the last page, the token will be the empty string.
   */
  nextPageToken: string;
}

function createBaseListRelationTuplesRequest(): ListRelationTuplesRequest {
  return {
    query: undefined,
    relationQuery: undefined,
    expandMask: undefined,
    snaptoken: "",
    pageSize: 0,
    pageToken: "",
  };
}

export const ListRelationTuplesRequest = {
  encode(message: ListRelationTuplesRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.query !== undefined) {
      ListRelationTuplesRequest_Query.encode(message.query, writer.uint32(10).fork()).ldelim();
    }
    if (message.relationQuery !== undefined) {
      RelationQuery.encode(message.relationQuery, writer.uint32(50).fork()).ldelim();
    }
    if (message.expandMask !== undefined) {
      FieldMask.encode(FieldMask.wrap(message.expandMask), writer.uint32(18).fork()).ldelim();
    }
    if (message.snaptoken !== "") {
      writer.uint32(26).string(message.snaptoken);
    }
    if (message.pageSize !== 0) {
      writer.uint32(32).int32(message.pageSize);
    }
    if (message.pageToken !== "") {
      writer.uint32(42).string(message.pageToken);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListRelationTuplesRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListRelationTuplesRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.query = ListRelationTuplesRequest_Query.decode(reader, reader.uint32());
          continue;
        case 6:
          if (tag !== 50) {
            break;
          }

          message.relationQuery = RelationQuery.decode(reader, reader.uint32());
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.expandMask = FieldMask.unwrap(FieldMask.decode(reader, reader.uint32()));
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.snaptoken = reader.string();
          continue;
        case 4:
          if (tag !== 32) {
            break;
          }

          message.pageSize = reader.int32();
          continue;
        case 5:
          if (tag !== 42) {
            break;
          }

          message.pageToken = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ListRelationTuplesRequest {
    return {
      query: isSet(object.query) ? ListRelationTuplesRequest_Query.fromJSON(object.query) : undefined,
      relationQuery: isSet(object.relationQuery) ? RelationQuery.fromJSON(object.relationQuery) : undefined,
      expandMask: isSet(object.expandMask) ? FieldMask.unwrap(FieldMask.fromJSON(object.expandMask)) : undefined,
      snaptoken: isSet(object.snaptoken) ? globalThis.String(object.snaptoken) : "",
      pageSize: isSet(object.pageSize) ? globalThis.Number(object.pageSize) : 0,
      pageToken: isSet(object.pageToken) ? globalThis.String(object.pageToken) : "",
    };
  },

  toJSON(message: ListRelationTuplesRequest): unknown {
    const obj: any = {};
    if (message.query !== undefined) {
      obj.query = ListRelationTuplesRequest_Query.toJSON(message.query);
    }
    if (message.relationQuery !== undefined) {
      obj.relationQuery = RelationQuery.toJSON(message.relationQuery);
    }
    if (message.expandMask !== undefined) {
      obj.expandMask = FieldMask.toJSON(FieldMask.wrap(message.expandMask));
    }
    if (message.snaptoken !== "") {
      obj.snaptoken = message.snaptoken;
    }
    if (message.pageSize !== 0) {
      obj.pageSize = Math.round(message.pageSize);
    }
    if (message.pageToken !== "") {
      obj.pageToken = message.pageToken;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<ListRelationTuplesRequest>, I>>(base?: I): ListRelationTuplesRequest {
    return ListRelationTuplesRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<ListRelationTuplesRequest>, I>>(object: I): ListRelationTuplesRequest {
    const message = createBaseListRelationTuplesRequest();
    message.query = (object.query !== undefined && object.query !== null)
      ? ListRelationTuplesRequest_Query.fromPartial(object.query)
      : undefined;
    message.relationQuery = (object.relationQuery !== undefined && object.relationQuery !== null)
      ? RelationQuery.fromPartial(object.relationQuery)
      : undefined;
    message.expandMask = object.expandMask ?? undefined;
    message.snaptoken = object.snaptoken ?? "";
    message.pageSize = object.pageSize ?? 0;
    message.pageToken = object.pageToken ?? "";
    return message;
  },
};

function createBaseListRelationTuplesRequest_Query(): ListRelationTuplesRequest_Query {
  return { namespace: "", object: "", relation: "", subject: undefined };
}

export const ListRelationTuplesRequest_Query = {
  encode(message: ListRelationTuplesRequest_Query, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
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

  decode(input: _m0.Reader | Uint8Array, length?: number): ListRelationTuplesRequest_Query {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListRelationTuplesRequest_Query();
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

  fromJSON(object: any): ListRelationTuplesRequest_Query {
    return {
      namespace: isSet(object.namespace) ? globalThis.String(object.namespace) : "",
      object: isSet(object.object) ? globalThis.String(object.object) : "",
      relation: isSet(object.relation) ? globalThis.String(object.relation) : "",
      subject: isSet(object.subject) ? Subject.fromJSON(object.subject) : undefined,
    };
  },

  toJSON(message: ListRelationTuplesRequest_Query): unknown {
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

  create<I extends Exact<DeepPartial<ListRelationTuplesRequest_Query>, I>>(base?: I): ListRelationTuplesRequest_Query {
    return ListRelationTuplesRequest_Query.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<ListRelationTuplesRequest_Query>, I>>(
    object: I,
  ): ListRelationTuplesRequest_Query {
    const message = createBaseListRelationTuplesRequest_Query();
    message.namespace = object.namespace ?? "";
    message.object = object.object ?? "";
    message.relation = object.relation ?? "";
    message.subject = (object.subject !== undefined && object.subject !== null)
      ? Subject.fromPartial(object.subject)
      : undefined;
    return message;
  },
};

function createBaseListRelationTuplesResponse(): ListRelationTuplesResponse {
  return { relationTuples: [], nextPageToken: "" };
}

export const ListRelationTuplesResponse = {
  encode(message: ListRelationTuplesResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.relationTuples) {
      RelationTuple.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.nextPageToken !== "") {
      writer.uint32(18).string(message.nextPageToken);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListRelationTuplesResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListRelationTuplesResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.relationTuples.push(RelationTuple.decode(reader, reader.uint32()));
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.nextPageToken = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ListRelationTuplesResponse {
    return {
      relationTuples: globalThis.Array.isArray(object?.relationTuples)
        ? object.relationTuples.map((e: any) => RelationTuple.fromJSON(e))
        : [],
      nextPageToken: isSet(object.nextPageToken) ? globalThis.String(object.nextPageToken) : "",
    };
  },

  toJSON(message: ListRelationTuplesResponse): unknown {
    const obj: any = {};
    if (message.relationTuples?.length) {
      obj.relationTuples = message.relationTuples.map((e) => RelationTuple.toJSON(e));
    }
    if (message.nextPageToken !== "") {
      obj.nextPageToken = message.nextPageToken;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<ListRelationTuplesResponse>, I>>(base?: I): ListRelationTuplesResponse {
    return ListRelationTuplesResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<ListRelationTuplesResponse>, I>>(object: I): ListRelationTuplesResponse {
    const message = createBaseListRelationTuplesResponse();
    message.relationTuples = object.relationTuples?.map((e) => RelationTuple.fromPartial(e)) || [];
    message.nextPageToken = object.nextPageToken ?? "";
    return message;
  },
};

/**
 * The service to query relationships.
 *
 * This service is part of the [read-APIs](../concepts/25_api-overview.mdx#read-apis).
 */
export type ReadServiceDefinition = typeof ReadServiceDefinition;
export const ReadServiceDefinition = {
  name: "ReadService",
  fullName: "ory.keto.relation_tuples.v1alpha2.ReadService",
  methods: {
    /** Lists ACL relationships. */
    listRelationTuples: {
      name: "ListRelationTuples",
      requestType: ListRelationTuplesRequest,
      requestStream: false,
      responseType: ListRelationTuplesResponse,
      responseStream: false,
      options: {},
    },
  },
} as const;

export interface ReadServiceImplementation<CallContextExt = {}> {
  /** Lists ACL relationships. */
  listRelationTuples(
    request: ListRelationTuplesRequest,
    context: CallContext & CallContextExt,
  ): Promise<DeepPartial<ListRelationTuplesResponse>>;
}

export interface ReadServiceClient<CallOptionsExt = {}> {
  /** Lists ACL relationships. */
  listRelationTuples(
    request: DeepPartial<ListRelationTuplesRequest>,
    options?: CallOptions & CallOptionsExt,
  ): Promise<ListRelationTuplesResponse>;
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
