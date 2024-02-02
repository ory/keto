/* eslint-disable */
import type { CallContext, CallOptions } from "nice-grpc-common";
import _m0 from "protobufjs/minimal.js";

export const protobufPackage = "ory.keto.relation_tuples.v1alpha2";

/** Request for ReadService.ListNamespaces RPC. */
export interface ListNamespacesRequest {
}

export interface ListNamespacesResponse {
  namespaces: Namespace[];
}

export interface Namespace {
  name: string;
}

function createBaseListNamespacesRequest(): ListNamespacesRequest {
  return {};
}

export const ListNamespacesRequest = {
  encode(_: ListNamespacesRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListNamespacesRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListNamespacesRequest();
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

  fromJSON(_: any): ListNamespacesRequest {
    return {};
  },

  toJSON(_: ListNamespacesRequest): unknown {
    const obj: any = {};
    return obj;
  },

  create<I extends Exact<DeepPartial<ListNamespacesRequest>, I>>(base?: I): ListNamespacesRequest {
    return ListNamespacesRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<ListNamespacesRequest>, I>>(_: I): ListNamespacesRequest {
    const message = createBaseListNamespacesRequest();
    return message;
  },
};

function createBaseListNamespacesResponse(): ListNamespacesResponse {
  return { namespaces: [] };
}

export const ListNamespacesResponse = {
  encode(message: ListNamespacesResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.namespaces) {
      Namespace.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ListNamespacesResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListNamespacesResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.namespaces.push(Namespace.decode(reader, reader.uint32()));
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ListNamespacesResponse {
    return {
      namespaces: globalThis.Array.isArray(object?.namespaces)
        ? object.namespaces.map((e: any) => Namespace.fromJSON(e))
        : [],
    };
  },

  toJSON(message: ListNamespacesResponse): unknown {
    const obj: any = {};
    if (message.namespaces?.length) {
      obj.namespaces = message.namespaces.map((e) => Namespace.toJSON(e));
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<ListNamespacesResponse>, I>>(base?: I): ListNamespacesResponse {
    return ListNamespacesResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<ListNamespacesResponse>, I>>(object: I): ListNamespacesResponse {
    const message = createBaseListNamespacesResponse();
    message.namespaces = object.namespaces?.map((e) => Namespace.fromPartial(e)) || [];
    return message;
  },
};

function createBaseNamespace(): Namespace {
  return { name: "" };
}

export const Namespace = {
  encode(message: Namespace, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Namespace {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseNamespace();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.name = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): Namespace {
    return { name: isSet(object.name) ? globalThis.String(object.name) : "" };
  },

  toJSON(message: Namespace): unknown {
    const obj: any = {};
    if (message.name !== "") {
      obj.name = message.name;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<Namespace>, I>>(base?: I): Namespace {
    return Namespace.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<Namespace>, I>>(object: I): Namespace {
    const message = createBaseNamespace();
    message.name = object.name ?? "";
    return message;
  },
};

/**
 * The service to query namespaces.
 *
 * This service is part of the [read-APIs](../concepts/25_api-overview.mdx#read-apis).
 */
export type NamespacesServiceDefinition = typeof NamespacesServiceDefinition;
export const NamespacesServiceDefinition = {
  name: "NamespacesService",
  fullName: "ory.keto.relation_tuples.v1alpha2.NamespacesService",
  methods: {
    /** Lists Namespaces */
    listNamespaces: {
      name: "ListNamespaces",
      requestType: ListNamespacesRequest,
      requestStream: false,
      responseType: ListNamespacesResponse,
      responseStream: false,
      options: {},
    },
  },
} as const;

export interface NamespacesServiceImplementation<CallContextExt = {}> {
  /** Lists Namespaces */
  listNamespaces(
    request: ListNamespacesRequest,
    context: CallContext & CallContextExt,
  ): Promise<DeepPartial<ListNamespacesResponse>>;
}

export interface NamespacesServiceClient<CallOptionsExt = {}> {
  /** Lists Namespaces */
  listNamespaces(
    request: DeepPartial<ListNamespacesRequest>,
    options?: CallOptions & CallOptionsExt,
  ): Promise<ListNamespacesResponse>;
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
