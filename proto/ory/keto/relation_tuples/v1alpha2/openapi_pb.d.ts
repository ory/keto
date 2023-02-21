// package: ory.keto.relation_tuples.v1alpha2
// file: ory/keto/relation_tuples/v1alpha2/openapi.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";
import * as protoc_gen_openapiv2_options_annotations_pb from "../../../../protoc-gen-openapiv2/options/annotations_pb";
import * as google_api_field_behavior_pb from "../../../../google/api/field_behavior_pb";

export class ErrorResponse extends jspb.Message { 

    hasError(): boolean;
    clearError(): void;
    getError(): ErrorResponse.Error | undefined;
    setError(value?: ErrorResponse.Error): ErrorResponse;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ErrorResponse.AsObject;
    static toObject(includeInstance: boolean, msg: ErrorResponse): ErrorResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ErrorResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ErrorResponse;
    static deserializeBinaryFromReader(message: ErrorResponse, reader: jspb.BinaryReader): ErrorResponse;
}

export namespace ErrorResponse {
    export type AsObject = {
        error?: ErrorResponse.Error.AsObject,
    }


    export class Error extends jspb.Message { 
        getCode(): number;
        setCode(value: number): Error;
        getDebug(): string;
        setDebug(value: string): Error;

        getDetailsMap(): jspb.Map<string, string>;
        clearDetailsMap(): void;
        getId(): string;
        setId(value: string): Error;
        getMessage(): string;
        setMessage(value: string): Error;
        getReason(): string;
        setReason(value: string): Error;
        getRequest(): string;
        setRequest(value: string): Error;
        getStatus(): string;
        setStatus(value: string): Error;

        serializeBinary(): Uint8Array;
        toObject(includeInstance?: boolean): Error.AsObject;
        static toObject(includeInstance: boolean, msg: Error): Error.AsObject;
        static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
        static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
        static serializeBinaryToWriter(message: Error, writer: jspb.BinaryWriter): void;
        static deserializeBinary(bytes: Uint8Array): Error;
        static deserializeBinaryFromReader(message: Error, reader: jspb.BinaryReader): Error;
    }

    export namespace Error {
        export type AsObject = {
            code: number,
            debug: string,

            detailsMap: Array<[string, string]>,
            id: string,
            message: string,
            reason: string,
            request: string,
            status: string,
        }
    }

}
