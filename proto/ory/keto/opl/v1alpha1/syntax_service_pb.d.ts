// package: ory.keto.opl.v1alpha1
// file: ory/keto/opl/v1alpha1/syntax_service.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";
import * as protoc_gen_openapiv2_options_annotations_pb from "../../../../protoc-gen-openapiv2/options/annotations_pb";

export class CheckRequest extends jspb.Message { 
    getContent(): string;
    setContent(value: string): CheckRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): CheckRequest.AsObject;
    static toObject(includeInstance: boolean, msg: CheckRequest): CheckRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: CheckRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): CheckRequest;
    static deserializeBinaryFromReader(message: CheckRequest, reader: jspb.BinaryReader): CheckRequest;
}

export namespace CheckRequest {
    export type AsObject = {
        content: string,
    }
}

export class CheckResponse extends jspb.Message { 
    clearErrorsList(): void;
    getErrorsList(): Array<ParseError>;
    setErrorsList(value: Array<ParseError>): CheckResponse;
    addErrors(value?: ParseError, index?: number): ParseError;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): CheckResponse.AsObject;
    static toObject(includeInstance: boolean, msg: CheckResponse): CheckResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: CheckResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): CheckResponse;
    static deserializeBinaryFromReader(message: CheckResponse, reader: jspb.BinaryReader): CheckResponse;
}

export namespace CheckResponse {
    export type AsObject = {
        errorsList: Array<ParseError.AsObject>,
    }
}

export class ParseError extends jspb.Message { 
    getMessage(): string;
    setMessage(value: string): ParseError;

    hasStart(): boolean;
    clearStart(): void;
    getStart(): SourcePosition | undefined;
    setStart(value?: SourcePosition): ParseError;

    hasEnd(): boolean;
    clearEnd(): void;
    getEnd(): SourcePosition | undefined;
    setEnd(value?: SourcePosition): ParseError;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ParseError.AsObject;
    static toObject(includeInstance: boolean, msg: ParseError): ParseError.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ParseError, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ParseError;
    static deserializeBinaryFromReader(message: ParseError, reader: jspb.BinaryReader): ParseError;
}

export namespace ParseError {
    export type AsObject = {
        message: string,
        start?: SourcePosition.AsObject,
        end?: SourcePosition.AsObject,
    }
}

export class SourcePosition extends jspb.Message { 
    getLine(): number;
    setLine(value: number): SourcePosition;
    getColumn(): number;
    setColumn(value: number): SourcePosition;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): SourcePosition.AsObject;
    static toObject(includeInstance: boolean, msg: SourcePosition): SourcePosition.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: SourcePosition, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): SourcePosition;
    static deserializeBinaryFromReader(message: SourcePosition, reader: jspb.BinaryReader): SourcePosition;
}

export namespace SourcePosition {
    export type AsObject = {
        line: number,
        column: number,
    }
}
