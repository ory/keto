/* eslint-disable */
import type { CallContext, CallOptions } from "nice-grpc-common";
import _m0 from "protobufjs/minimal.js";
import { RelationTuple, Subject } from "./relation_tuples.js";

export const protobufPackage = "ory.keto.relation_tuples.v1alpha2";

/**
 * The request for a CheckService.Check RPC.
 * Checks whether a specific subject is related to an object.
 */
export interface CheckRequest {
  /**
   * The namespace to evaluate the check.
   *
   * Note: If you use the expand-API and the check
   * evaluates a RelationTuple specifying a SubjectSet as
   * subject or due to a rewrite rule in a namespace config
   * this check request may involve other namespaces automatically.
   *
   * @deprecated
   */
  namespace: string;
  /**
   * The related object in this check.
   *
   * @deprecated
   */
  object: string;
  /**
   * The relation between the Object and the Subject.
   *
   * @deprecated
   */
  relation: string;
  /**
   * The related subject in this check.
   *
   * @deprecated
   */
  subject: Subject | undefined;
  tuple:
    | RelationTuple
    | undefined;
  /**
   * This field is not implemented yet and has no effect.
   * <!--
   * Set this field to `true` in case your application
   * needs to authorize depending on up to date ACLs,
   * also called a "content-change check".
   *
   * If set to `true` the `snaptoken` field is ignored,
   * the check is evaluated at the latest snapshot
   * (globally consistent) and the response includes a
   * snaptoken for clients to store along with object
   * contents that can be used for subsequent checks
   * of the same content version.
   *
   * Example use case:
   *  - You need to authorize a user to modify/delete some resource
   *    and it is unacceptable that if the permission to do that had
   *    just been revoked some seconds ago so that the change had not
   *    yet been fully replicated to all availability zones.
   * -->
   */
  latest: boolean;
  /**
   * This field is not implemented yet and has no effect.
   * <!--
   * Optional. Like reads, a check is always evaluated at a
   * consistent snapshot no earlier than the given snaptoken.
   *
   * Leave this field blank if you want to evaluate the check
   * based on eventually consistent ACLs, benefiting from very
   * low latency, but possibly slightly stale results.
   *
   * If the specified token is too old and no longer known,
   * the server falls back as if no snaptoken had been specified.
   *
   * If not specified the server tries to evaluate the check
   * on the best snapshot version where it is very likely that
   * ACLs had already been replicated to all availability zones.
   * -->
   */
  snaptoken: string;
  /**
   * The maximum depth to search for a relation.
   *
   * If the value is less than 1 or greater than the global
   * max-depth then the global max-depth will be used instead.
   */
  maxDepth: number;
}

/** The response for a CheckService.Check rpc. */
export interface CheckResponse {
  /**
   * Whether the specified subject (id)
   * is related to the requested object.
   *
   * It is false by default if no ACL matches.
   */
  allowed: boolean;
  /**
   * This field is not implemented yet and has no effect.
   * <!--
   * The last known snapshot token ONLY specified if
   * the request had not specified a snaptoken,
   * since this performed a "content-change request"
   * and consistently fetched the last known snapshot token.
   *
   * This field is not set if the request had specified a snaptoken!
   *
   * If set, clients should cache and use this token
   * for subsequent requests to have minimal latency,
   * but allow slightly stale responses (only some milliseconds or seconds).
   * -->
   */
  snaptoken: string;
}

function createBaseCheckRequest(): CheckRequest {
  return {
    namespace: "",
    object: "",
    relation: "",
    subject: undefined,
    tuple: undefined,
    latest: false,
    snaptoken: "",
    maxDepth: 0,
  };
}

export const CheckRequest = {
  encode(message: CheckRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
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
    if (message.tuple !== undefined) {
      RelationTuple.encode(message.tuple, writer.uint32(66).fork()).ldelim();
    }
    if (message.latest === true) {
      writer.uint32(40).bool(message.latest);
    }
    if (message.snaptoken !== "") {
      writer.uint32(50).string(message.snaptoken);
    }
    if (message.maxDepth !== 0) {
      writer.uint32(56).int32(message.maxDepth);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CheckRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCheckRequest();
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
        case 8:
          if (tag !== 66) {
            break;
          }

          message.tuple = RelationTuple.decode(reader, reader.uint32());
          continue;
        case 5:
          if (tag !== 40) {
            break;
          }

          message.latest = reader.bool();
          continue;
        case 6:
          if (tag !== 50) {
            break;
          }

          message.snaptoken = reader.string();
          continue;
        case 7:
          if (tag !== 56) {
            break;
          }

          message.maxDepth = reader.int32();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): CheckRequest {
    return {
      namespace: isSet(object.namespace) ? globalThis.String(object.namespace) : "",
      object: isSet(object.object) ? globalThis.String(object.object) : "",
      relation: isSet(object.relation) ? globalThis.String(object.relation) : "",
      subject: isSet(object.subject) ? Subject.fromJSON(object.subject) : undefined,
      tuple: isSet(object.tuple) ? RelationTuple.fromJSON(object.tuple) : undefined,
      latest: isSet(object.latest) ? globalThis.Boolean(object.latest) : false,
      snaptoken: isSet(object.snaptoken) ? globalThis.String(object.snaptoken) : "",
      maxDepth: isSet(object.maxDepth) ? globalThis.Number(object.maxDepth) : 0,
    };
  },

  toJSON(message: CheckRequest): unknown {
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
    if (message.tuple !== undefined) {
      obj.tuple = RelationTuple.toJSON(message.tuple);
    }
    if (message.latest === true) {
      obj.latest = message.latest;
    }
    if (message.snaptoken !== "") {
      obj.snaptoken = message.snaptoken;
    }
    if (message.maxDepth !== 0) {
      obj.maxDepth = Math.round(message.maxDepth);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<CheckRequest>, I>>(base?: I): CheckRequest {
    return CheckRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<CheckRequest>, I>>(object: I): CheckRequest {
    const message = createBaseCheckRequest();
    message.namespace = object.namespace ?? "";
    message.object = object.object ?? "";
    message.relation = object.relation ?? "";
    message.subject = (object.subject !== undefined && object.subject !== null)
      ? Subject.fromPartial(object.subject)
      : undefined;
    message.tuple = (object.tuple !== undefined && object.tuple !== null)
      ? RelationTuple.fromPartial(object.tuple)
      : undefined;
    message.latest = object.latest ?? false;
    message.snaptoken = object.snaptoken ?? "";
    message.maxDepth = object.maxDepth ?? 0;
    return message;
  },
};

function createBaseCheckResponse(): CheckResponse {
  return { allowed: false, snaptoken: "" };
}

export const CheckResponse = {
  encode(message: CheckResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.allowed === true) {
      writer.uint32(8).bool(message.allowed);
    }
    if (message.snaptoken !== "") {
      writer.uint32(18).string(message.snaptoken);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CheckResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCheckResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.allowed = reader.bool();
          continue;
        case 2:
          if (tag !== 18) {
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

  fromJSON(object: any): CheckResponse {
    return {
      allowed: isSet(object.allowed) ? globalThis.Boolean(object.allowed) : false,
      snaptoken: isSet(object.snaptoken) ? globalThis.String(object.snaptoken) : "",
    };
  },

  toJSON(message: CheckResponse): unknown {
    const obj: any = {};
    if (message.allowed === true) {
      obj.allowed = message.allowed;
    }
    if (message.snaptoken !== "") {
      obj.snaptoken = message.snaptoken;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<CheckResponse>, I>>(base?: I): CheckResponse {
    return CheckResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<CheckResponse>, I>>(object: I): CheckResponse {
    const message = createBaseCheckResponse();
    message.allowed = object.allowed ?? false;
    message.snaptoken = object.snaptoken ?? "";
    return message;
  },
};

/**
 * The service that performs authorization checks
 * based on the stored Access Control Lists.
 *
 * This service is part of the [read-APIs](../concepts/25_api-overview.mdx#read-apis).
 */
export type CheckServiceDefinition = typeof CheckServiceDefinition;
export const CheckServiceDefinition = {
  name: "CheckService",
  fullName: "ory.keto.relation_tuples.v1alpha2.CheckService",
  methods: {
    /** Performs an authorization check. */
    check: {
      name: "Check",
      requestType: CheckRequest,
      requestStream: false,
      responseType: CheckResponse,
      responseStream: false,
      options: {},
    },
  },
} as const;

export interface CheckServiceImplementation<CallContextExt = {}> {
  /** Performs an authorization check. */
  check(request: CheckRequest, context: CallContext & CallContextExt): Promise<DeepPartial<CheckResponse>>;
}

export interface CheckServiceClient<CallOptionsExt = {}> {
  /** Performs an authorization check. */
  check(request: DeepPartial<CheckRequest>, options?: CallOptions & CallOptionsExt): Promise<CheckResponse>;
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
