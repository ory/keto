/* eslint-disable */
import type { CallContext, CallOptions } from "nice-grpc-common";
import _m0 from "protobufjs/minimal.js";

export const protobufPackage = "ory.keto.opl.v1alpha1";

export interface CheckRequest {
  content: Uint8Array;
}

export interface CheckResponse {
  parseErrors: ParseError[];
}

export interface ParseError {
  message: string;
  start: SourcePosition | undefined;
  end: SourcePosition | undefined;
}

export interface SourcePosition {
  line: number;
  column: number;
}

function createBaseCheckRequest(): CheckRequest {
  return { content: new Uint8Array(0) };
}

export const CheckRequest = {
  encode(message: CheckRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.content.length !== 0) {
      writer.uint32(10).bytes(message.content);
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

          message.content = reader.bytes();
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
    return { content: isSet(object.content) ? bytesFromBase64(object.content) : new Uint8Array(0) };
  },

  toJSON(message: CheckRequest): unknown {
    const obj: any = {};
    if (message.content.length !== 0) {
      obj.content = base64FromBytes(message.content);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<CheckRequest>, I>>(base?: I): CheckRequest {
    return CheckRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<CheckRequest>, I>>(object: I): CheckRequest {
    const message = createBaseCheckRequest();
    message.content = object.content ?? new Uint8Array(0);
    return message;
  },
};

function createBaseCheckResponse(): CheckResponse {
  return { parseErrors: [] };
}

export const CheckResponse = {
  encode(message: CheckResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.parseErrors) {
      ParseError.encode(v!, writer.uint32(10).fork()).ldelim();
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
          if (tag !== 10) {
            break;
          }

          message.parseErrors.push(ParseError.decode(reader, reader.uint32()));
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
      parseErrors: globalThis.Array.isArray(object?.parseErrors)
        ? object.parseErrors.map((e: any) => ParseError.fromJSON(e))
        : [],
    };
  },

  toJSON(message: CheckResponse): unknown {
    const obj: any = {};
    if (message.parseErrors?.length) {
      obj.parseErrors = message.parseErrors.map((e) => ParseError.toJSON(e));
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<CheckResponse>, I>>(base?: I): CheckResponse {
    return CheckResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<CheckResponse>, I>>(object: I): CheckResponse {
    const message = createBaseCheckResponse();
    message.parseErrors = object.parseErrors?.map((e) => ParseError.fromPartial(e)) || [];
    return message;
  },
};

function createBaseParseError(): ParseError {
  return { message: "", start: undefined, end: undefined };
}

export const ParseError = {
  encode(message: ParseError, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.message !== "") {
      writer.uint32(10).string(message.message);
    }
    if (message.start !== undefined) {
      SourcePosition.encode(message.start, writer.uint32(18).fork()).ldelim();
    }
    if (message.end !== undefined) {
      SourcePosition.encode(message.end, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ParseError {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseParseError();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.message = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.start = SourcePosition.decode(reader, reader.uint32());
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.end = SourcePosition.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ParseError {
    return {
      message: isSet(object.message) ? globalThis.String(object.message) : "",
      start: isSet(object.start) ? SourcePosition.fromJSON(object.start) : undefined,
      end: isSet(object.end) ? SourcePosition.fromJSON(object.end) : undefined,
    };
  },

  toJSON(message: ParseError): unknown {
    const obj: any = {};
    if (message.message !== "") {
      obj.message = message.message;
    }
    if (message.start !== undefined) {
      obj.start = SourcePosition.toJSON(message.start);
    }
    if (message.end !== undefined) {
      obj.end = SourcePosition.toJSON(message.end);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<ParseError>, I>>(base?: I): ParseError {
    return ParseError.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<ParseError>, I>>(object: I): ParseError {
    const message = createBaseParseError();
    message.message = object.message ?? "";
    message.start = (object.start !== undefined && object.start !== null)
      ? SourcePosition.fromPartial(object.start)
      : undefined;
    message.end = (object.end !== undefined && object.end !== null)
      ? SourcePosition.fromPartial(object.end)
      : undefined;
    return message;
  },
};

function createBaseSourcePosition(): SourcePosition {
  return { line: 0, column: 0 };
}

export const SourcePosition = {
  encode(message: SourcePosition, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.line !== 0) {
      writer.uint32(8).uint32(message.line);
    }
    if (message.column !== 0) {
      writer.uint32(16).uint32(message.column);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SourcePosition {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSourcePosition();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.line = reader.uint32();
          continue;
        case 2:
          if (tag !== 16) {
            break;
          }

          message.column = reader.uint32();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): SourcePosition {
    return {
      line: isSet(object.line) ? globalThis.Number(object.line) : 0,
      column: isSet(object.column) ? globalThis.Number(object.column) : 0,
    };
  },

  toJSON(message: SourcePosition): unknown {
    const obj: any = {};
    if (message.line !== 0) {
      obj.line = Math.round(message.line);
    }
    if (message.column !== 0) {
      obj.column = Math.round(message.column);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<SourcePosition>, I>>(base?: I): SourcePosition {
    return SourcePosition.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<SourcePosition>, I>>(object: I): SourcePosition {
    const message = createBaseSourcePosition();
    message.line = object.line ?? 0;
    message.column = object.column ?? 0;
    return message;
  },
};

/** The service that checks the syntax of an OPL file. */
export type SyntaxServiceDefinition = typeof SyntaxServiceDefinition;
export const SyntaxServiceDefinition = {
  name: "SyntaxService",
  fullName: "ory.keto.opl.v1alpha1.SyntaxService",
  methods: {
    /** Performs a syntax check request. */
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

export interface SyntaxServiceImplementation<CallContextExt = {}> {
  /** Performs a syntax check request. */
  check(request: CheckRequest, context: CallContext & CallContextExt): Promise<DeepPartial<CheckResponse>>;
}

export interface SyntaxServiceClient<CallOptionsExt = {}> {
  /** Performs a syntax check request. */
  check(request: DeepPartial<CheckRequest>, options?: CallOptions & CallOptionsExt): Promise<CheckResponse>;
}

function bytesFromBase64(b64: string): Uint8Array {
  if (globalThis.Buffer) {
    return Uint8Array.from(globalThis.Buffer.from(b64, "base64"));
  } else {
    const bin = globalThis.atob(b64);
    const arr = new Uint8Array(bin.length);
    for (let i = 0; i < bin.length; ++i) {
      arr[i] = bin.charCodeAt(i);
    }
    return arr;
  }
}

function base64FromBytes(arr: Uint8Array): string {
  if (globalThis.Buffer) {
    return globalThis.Buffer.from(arr).toString("base64");
  } else {
    const bin: string[] = [];
    arr.forEach((byte) => {
      bin.push(globalThis.String.fromCharCode(byte));
    });
    return globalThis.btoa(bin.join(""));
  }
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
