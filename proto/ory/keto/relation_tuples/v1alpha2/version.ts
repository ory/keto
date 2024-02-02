/* eslint-disable */
import type { CallContext, CallOptions } from "nice-grpc-common";
import _m0 from "protobufjs/minimal.js";

export const protobufPackage = "ory.keto.relation_tuples.v1alpha2";

/** Request for the VersionService.GetVersion RPC. */
export interface GetVersionRequest {
}

/** Response of the VersionService.GetVersion RPC. */
export interface GetVersionResponse {
  /** The version string of the Ory Keto instance. */
  version: string;
}

function createBaseGetVersionRequest(): GetVersionRequest {
  return {};
}

export const GetVersionRequest = {
  encode(_: GetVersionRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetVersionRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetVersionRequest();
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

  fromJSON(_: any): GetVersionRequest {
    return {};
  },

  toJSON(_: GetVersionRequest): unknown {
    const obj: any = {};
    return obj;
  },

  create<I extends Exact<DeepPartial<GetVersionRequest>, I>>(base?: I): GetVersionRequest {
    return GetVersionRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<GetVersionRequest>, I>>(_: I): GetVersionRequest {
    const message = createBaseGetVersionRequest();
    return message;
  },
};

function createBaseGetVersionResponse(): GetVersionResponse {
  return { version: "" };
}

export const GetVersionResponse = {
  encode(message: GetVersionResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.version !== "") {
      writer.uint32(10).string(message.version);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GetVersionResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetVersionResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.version = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetVersionResponse {
    return { version: isSet(object.version) ? globalThis.String(object.version) : "" };
  },

  toJSON(message: GetVersionResponse): unknown {
    const obj: any = {};
    if (message.version !== "") {
      obj.version = message.version;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<GetVersionResponse>, I>>(base?: I): GetVersionResponse {
    return GetVersionResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<GetVersionResponse>, I>>(object: I): GetVersionResponse {
    const message = createBaseGetVersionResponse();
    message.version = object.version ?? "";
    return message;
  },
};

/**
 * The service returning the specific Ory Keto instance version.
 *
 * This service is part of the [read-APIs](../concepts/25_api-overview.mdx#read-apis) and [write-APIs](../concepts/25_api-overview.mdx#write-apis).
 */
export type VersionServiceDefinition = typeof VersionServiceDefinition;
export const VersionServiceDefinition = {
  name: "VersionService",
  fullName: "ory.keto.relation_tuples.v1alpha2.VersionService",
  methods: {
    /** Returns the version of the Ory Keto instance. */
    getVersion: {
      name: "GetVersion",
      requestType: GetVersionRequest,
      requestStream: false,
      responseType: GetVersionResponse,
      responseStream: false,
      options: {},
    },
  },
} as const;

export interface VersionServiceImplementation<CallContextExt = {}> {
  /** Returns the version of the Ory Keto instance. */
  getVersion(
    request: GetVersionRequest,
    context: CallContext & CallContextExt,
  ): Promise<DeepPartial<GetVersionResponse>>;
}

export interface VersionServiceClient<CallOptionsExt = {}> {
  /** Returns the version of the Ory Keto instance. */
  getVersion(
    request: DeepPartial<GetVersionRequest>,
    options?: CallOptions & CallOptionsExt,
  ): Promise<GetVersionResponse>;
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
