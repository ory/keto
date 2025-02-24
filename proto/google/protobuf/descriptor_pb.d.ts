// package: google.protobuf
// file: google/protobuf/descriptor.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";

export class FileDescriptorSet extends jspb.Message { 
    clearFileList(): void;
    getFileList(): Array<FileDescriptorProto>;
    setFileList(value: Array<FileDescriptorProto>): FileDescriptorSet;
    addFile(value?: FileDescriptorProto, index?: number): FileDescriptorProto;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): FileDescriptorSet.AsObject;
    static toObject(includeInstance: boolean, msg: FileDescriptorSet): FileDescriptorSet.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: FileDescriptorSet, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): FileDescriptorSet;
    static deserializeBinaryFromReader(message: FileDescriptorSet, reader: jspb.BinaryReader): FileDescriptorSet;
}

export namespace FileDescriptorSet {
    export type AsObject = {
        fileList: Array<FileDescriptorProto.AsObject>,
    }
}

export class FileDescriptorProto extends jspb.Message { 

    hasName(): boolean;
    clearName(): void;
    getName(): string | undefined;
    setName(value: string): FileDescriptorProto;

    hasPackage(): boolean;
    clearPackage(): void;
    getPackage(): string | undefined;
    setPackage(value: string): FileDescriptorProto;
    clearDependencyList(): void;
    getDependencyList(): Array<string>;
    setDependencyList(value: Array<string>): FileDescriptorProto;
    addDependency(value: string, index?: number): string;
    clearPublicDependencyList(): void;
    getPublicDependencyList(): Array<number>;
    setPublicDependencyList(value: Array<number>): FileDescriptorProto;
    addPublicDependency(value: number, index?: number): number;
    clearWeakDependencyList(): void;
    getWeakDependencyList(): Array<number>;
    setWeakDependencyList(value: Array<number>): FileDescriptorProto;
    addWeakDependency(value: number, index?: number): number;
    clearMessageTypeList(): void;
    getMessageTypeList(): Array<DescriptorProto>;
    setMessageTypeList(value: Array<DescriptorProto>): FileDescriptorProto;
    addMessageType(value?: DescriptorProto, index?: number): DescriptorProto;
    clearEnumTypeList(): void;
    getEnumTypeList(): Array<EnumDescriptorProto>;
    setEnumTypeList(value: Array<EnumDescriptorProto>): FileDescriptorProto;
    addEnumType(value?: EnumDescriptorProto, index?: number): EnumDescriptorProto;
    clearServiceList(): void;
    getServiceList(): Array<ServiceDescriptorProto>;
    setServiceList(value: Array<ServiceDescriptorProto>): FileDescriptorProto;
    addService(value?: ServiceDescriptorProto, index?: number): ServiceDescriptorProto;
    clearExtensionList(): void;
    getExtensionList(): Array<FieldDescriptorProto>;
    setExtensionList(value: Array<FieldDescriptorProto>): FileDescriptorProto;
    addExtension$(value?: FieldDescriptorProto, index?: number): FieldDescriptorProto;

    hasOptions(): boolean;
    clearOptions(): void;
    getOptions(): FileOptions | undefined;
    setOptions(value?: FileOptions): FileDescriptorProto;

    hasSourceCodeInfo(): boolean;
    clearSourceCodeInfo(): void;
    getSourceCodeInfo(): SourceCodeInfo | undefined;
    setSourceCodeInfo(value?: SourceCodeInfo): FileDescriptorProto;

    hasSyntax(): boolean;
    clearSyntax(): void;
    getSyntax(): string | undefined;
    setSyntax(value: string): FileDescriptorProto;

    hasEdition(): boolean;
    clearEdition(): void;
    getEdition(): Edition | undefined;
    setEdition(value: Edition): FileDescriptorProto;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): FileDescriptorProto.AsObject;
    static toObject(includeInstance: boolean, msg: FileDescriptorProto): FileDescriptorProto.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: FileDescriptorProto, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): FileDescriptorProto;
    static deserializeBinaryFromReader(message: FileDescriptorProto, reader: jspb.BinaryReader): FileDescriptorProto;
}

export namespace FileDescriptorProto {
    export type AsObject = {
        name?: string,
        pb_package?: string,
        dependencyList: Array<string>,
        publicDependencyList: Array<number>,
        weakDependencyList: Array<number>,
        messageTypeList: Array<DescriptorProto.AsObject>,
        enumTypeList: Array<EnumDescriptorProto.AsObject>,
        serviceList: Array<ServiceDescriptorProto.AsObject>,
        extensionList: Array<FieldDescriptorProto.AsObject>,
        options?: FileOptions.AsObject,
        sourceCodeInfo?: SourceCodeInfo.AsObject,
        syntax?: string,
        edition?: Edition,
    }
}

export class DescriptorProto extends jspb.Message { 

    hasName(): boolean;
    clearName(): void;
    getName(): string | undefined;
    setName(value: string): DescriptorProto;
    clearFieldList(): void;
    getFieldList(): Array<FieldDescriptorProto>;
    setFieldList(value: Array<FieldDescriptorProto>): DescriptorProto;
    addField(value?: FieldDescriptorProto, index?: number): FieldDescriptorProto;
    clearExtensionList(): void;
    getExtensionList(): Array<FieldDescriptorProto>;
    setExtensionList(value: Array<FieldDescriptorProto>): DescriptorProto;
    addExtension$(value?: FieldDescriptorProto, index?: number): FieldDescriptorProto;
    clearNestedTypeList(): void;
    getNestedTypeList(): Array<DescriptorProto>;
    setNestedTypeList(value: Array<DescriptorProto>): DescriptorProto;
    addNestedType(value?: DescriptorProto, index?: number): DescriptorProto;
    clearEnumTypeList(): void;
    getEnumTypeList(): Array<EnumDescriptorProto>;
    setEnumTypeList(value: Array<EnumDescriptorProto>): DescriptorProto;
    addEnumType(value?: EnumDescriptorProto, index?: number): EnumDescriptorProto;
    clearExtensionRangeList(): void;
    getExtensionRangeList(): Array<DescriptorProto.ExtensionRange>;
    setExtensionRangeList(value: Array<DescriptorProto.ExtensionRange>): DescriptorProto;
    addExtensionRange(value?: DescriptorProto.ExtensionRange, index?: number): DescriptorProto.ExtensionRange;
    clearOneofDeclList(): void;
    getOneofDeclList(): Array<OneofDescriptorProto>;
    setOneofDeclList(value: Array<OneofDescriptorProto>): DescriptorProto;
    addOneofDecl(value?: OneofDescriptorProto, index?: number): OneofDescriptorProto;

    hasOptions(): boolean;
    clearOptions(): void;
    getOptions(): MessageOptions | undefined;
    setOptions(value?: MessageOptions): DescriptorProto;
    clearReservedRangeList(): void;
    getReservedRangeList(): Array<DescriptorProto.ReservedRange>;
    setReservedRangeList(value: Array<DescriptorProto.ReservedRange>): DescriptorProto;
    addReservedRange(value?: DescriptorProto.ReservedRange, index?: number): DescriptorProto.ReservedRange;
    clearReservedNameList(): void;
    getReservedNameList(): Array<string>;
    setReservedNameList(value: Array<string>): DescriptorProto;
    addReservedName(value: string, index?: number): string;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): DescriptorProto.AsObject;
    static toObject(includeInstance: boolean, msg: DescriptorProto): DescriptorProto.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: DescriptorProto, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): DescriptorProto;
    static deserializeBinaryFromReader(message: DescriptorProto, reader: jspb.BinaryReader): DescriptorProto;
}

export namespace DescriptorProto {
    export type AsObject = {
        name?: string,
        fieldList: Array<FieldDescriptorProto.AsObject>,
        extensionList: Array<FieldDescriptorProto.AsObject>,
        nestedTypeList: Array<DescriptorProto.AsObject>,
        enumTypeList: Array<EnumDescriptorProto.AsObject>,
        extensionRangeList: Array<DescriptorProto.ExtensionRange.AsObject>,
        oneofDeclList: Array<OneofDescriptorProto.AsObject>,
        options?: MessageOptions.AsObject,
        reservedRangeList: Array<DescriptorProto.ReservedRange.AsObject>,
        reservedNameList: Array<string>,
    }


    export class ExtensionRange extends jspb.Message { 

        hasStart(): boolean;
        clearStart(): void;
        getStart(): number | undefined;
        setStart(value: number): ExtensionRange;

        hasEnd(): boolean;
        clearEnd(): void;
        getEnd(): number | undefined;
        setEnd(value: number): ExtensionRange;

        hasOptions(): boolean;
        clearOptions(): void;
        getOptions(): ExtensionRangeOptions | undefined;
        setOptions(value?: ExtensionRangeOptions): ExtensionRange;

        serializeBinary(): Uint8Array;
        toObject(includeInstance?: boolean): ExtensionRange.AsObject;
        static toObject(includeInstance: boolean, msg: ExtensionRange): ExtensionRange.AsObject;
        static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
        static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
        static serializeBinaryToWriter(message: ExtensionRange, writer: jspb.BinaryWriter): void;
        static deserializeBinary(bytes: Uint8Array): ExtensionRange;
        static deserializeBinaryFromReader(message: ExtensionRange, reader: jspb.BinaryReader): ExtensionRange;
    }

    export namespace ExtensionRange {
        export type AsObject = {
            start?: number,
            end?: number,
            options?: ExtensionRangeOptions.AsObject,
        }
    }

    export class ReservedRange extends jspb.Message { 

        hasStart(): boolean;
        clearStart(): void;
        getStart(): number | undefined;
        setStart(value: number): ReservedRange;

        hasEnd(): boolean;
        clearEnd(): void;
        getEnd(): number | undefined;
        setEnd(value: number): ReservedRange;

        serializeBinary(): Uint8Array;
        toObject(includeInstance?: boolean): ReservedRange.AsObject;
        static toObject(includeInstance: boolean, msg: ReservedRange): ReservedRange.AsObject;
        static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
        static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
        static serializeBinaryToWriter(message: ReservedRange, writer: jspb.BinaryWriter): void;
        static deserializeBinary(bytes: Uint8Array): ReservedRange;
        static deserializeBinaryFromReader(message: ReservedRange, reader: jspb.BinaryReader): ReservedRange;
    }

    export namespace ReservedRange {
        export type AsObject = {
            start?: number,
            end?: number,
        }
    }

}

export class ExtensionRangeOptions extends jspb.Message { 
    clearUninterpretedOptionList(): void;
    getUninterpretedOptionList(): Array<UninterpretedOption>;
    setUninterpretedOptionList(value: Array<UninterpretedOption>): ExtensionRangeOptions;
    addUninterpretedOption(value?: UninterpretedOption, index?: number): UninterpretedOption;
    clearDeclarationList(): void;
    getDeclarationList(): Array<ExtensionRangeOptions.Declaration>;
    setDeclarationList(value: Array<ExtensionRangeOptions.Declaration>): ExtensionRangeOptions;
    addDeclaration(value?: ExtensionRangeOptions.Declaration, index?: number): ExtensionRangeOptions.Declaration;

    hasFeatures(): boolean;
    clearFeatures(): void;
    getFeatures(): FeatureSet | undefined;
    setFeatures(value?: FeatureSet): ExtensionRangeOptions;

    hasVerification(): boolean;
    clearVerification(): void;
    getVerification(): ExtensionRangeOptions.VerificationState | undefined;
    setVerification(value: ExtensionRangeOptions.VerificationState): ExtensionRangeOptions;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ExtensionRangeOptions.AsObject;
    static toObject(includeInstance: boolean, msg: ExtensionRangeOptions): ExtensionRangeOptions.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ExtensionRangeOptions, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ExtensionRangeOptions;
    static deserializeBinaryFromReader(message: ExtensionRangeOptions, reader: jspb.BinaryReader): ExtensionRangeOptions;
}

export namespace ExtensionRangeOptions {
    export type AsObject = {
        uninterpretedOptionList: Array<UninterpretedOption.AsObject>,
        declarationList: Array<ExtensionRangeOptions.Declaration.AsObject>,
        features?: FeatureSet.AsObject,
        verification?: ExtensionRangeOptions.VerificationState,
    }


    export class Declaration extends jspb.Message { 

        hasNumber(): boolean;
        clearNumber(): void;
        getNumber(): number | undefined;
        setNumber(value: number): Declaration;

        hasFullName(): boolean;
        clearFullName(): void;
        getFullName(): string | undefined;
        setFullName(value: string): Declaration;

        hasType(): boolean;
        clearType(): void;
        getType(): string | undefined;
        setType(value: string): Declaration;

        hasReserved(): boolean;
        clearReserved(): void;
        getReserved(): boolean | undefined;
        setReserved(value: boolean): Declaration;

        hasRepeated(): boolean;
        clearRepeated(): void;
        getRepeated(): boolean | undefined;
        setRepeated(value: boolean): Declaration;

        serializeBinary(): Uint8Array;
        toObject(includeInstance?: boolean): Declaration.AsObject;
        static toObject(includeInstance: boolean, msg: Declaration): Declaration.AsObject;
        static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
        static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
        static serializeBinaryToWriter(message: Declaration, writer: jspb.BinaryWriter): void;
        static deserializeBinary(bytes: Uint8Array): Declaration;
        static deserializeBinaryFromReader(message: Declaration, reader: jspb.BinaryReader): Declaration;
    }

    export namespace Declaration {
        export type AsObject = {
            number?: number,
            fullName?: string,
            type?: string,
            reserved?: boolean,
            repeated?: boolean,
        }
    }


    export enum VerificationState {
    DECLARATION = 0,
    UNVERIFIED = 1,
    }

}

export class FieldDescriptorProto extends jspb.Message { 

    hasName(): boolean;
    clearName(): void;
    getName(): string | undefined;
    setName(value: string): FieldDescriptorProto;

    hasNumber(): boolean;
    clearNumber(): void;
    getNumber(): number | undefined;
    setNumber(value: number): FieldDescriptorProto;

    hasLabel(): boolean;
    clearLabel(): void;
    getLabel(): FieldDescriptorProto.Label | undefined;
    setLabel(value: FieldDescriptorProto.Label): FieldDescriptorProto;

    hasType(): boolean;
    clearType(): void;
    getType(): FieldDescriptorProto.Type | undefined;
    setType(value: FieldDescriptorProto.Type): FieldDescriptorProto;

    hasTypeName(): boolean;
    clearTypeName(): void;
    getTypeName(): string | undefined;
    setTypeName(value: string): FieldDescriptorProto;

    hasExtendee(): boolean;
    clearExtendee(): void;
    getExtendee(): string | undefined;
    setExtendee(value: string): FieldDescriptorProto;

    hasDefaultValue(): boolean;
    clearDefaultValue(): void;
    getDefaultValue(): string | undefined;
    setDefaultValue(value: string): FieldDescriptorProto;

    hasOneofIndex(): boolean;
    clearOneofIndex(): void;
    getOneofIndex(): number | undefined;
    setOneofIndex(value: number): FieldDescriptorProto;

    hasJsonName(): boolean;
    clearJsonName(): void;
    getJsonName(): string | undefined;
    setJsonName(value: string): FieldDescriptorProto;

    hasOptions(): boolean;
    clearOptions(): void;
    getOptions(): FieldOptions | undefined;
    setOptions(value?: FieldOptions): FieldDescriptorProto;

    hasProto3Optional(): boolean;
    clearProto3Optional(): void;
    getProto3Optional(): boolean | undefined;
    setProto3Optional(value: boolean): FieldDescriptorProto;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): FieldDescriptorProto.AsObject;
    static toObject(includeInstance: boolean, msg: FieldDescriptorProto): FieldDescriptorProto.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: FieldDescriptorProto, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): FieldDescriptorProto;
    static deserializeBinaryFromReader(message: FieldDescriptorProto, reader: jspb.BinaryReader): FieldDescriptorProto;
}

export namespace FieldDescriptorProto {
    export type AsObject = {
        name?: string,
        number?: number,
        label?: FieldDescriptorProto.Label,
        type?: FieldDescriptorProto.Type,
        typeName?: string,
        extendee?: string,
        defaultValue?: string,
        oneofIndex?: number,
        jsonName?: string,
        options?: FieldOptions.AsObject,
        proto3Optional?: boolean,
    }

    export enum Type {
    TYPE_DOUBLE = 1,
    TYPE_FLOAT = 2,
    TYPE_INT64 = 3,
    TYPE_UINT64 = 4,
    TYPE_INT32 = 5,
    TYPE_FIXED64 = 6,
    TYPE_FIXED32 = 7,
    TYPE_BOOL = 8,
    TYPE_STRING = 9,
    TYPE_GROUP = 10,
    TYPE_MESSAGE = 11,
    TYPE_BYTES = 12,
    TYPE_UINT32 = 13,
    TYPE_ENUM = 14,
    TYPE_SFIXED32 = 15,
    TYPE_SFIXED64 = 16,
    TYPE_SINT32 = 17,
    TYPE_SINT64 = 18,
    }

    export enum Label {
    LABEL_OPTIONAL = 1,
    LABEL_REPEATED = 3,
    LABEL_REQUIRED = 2,
    }

}

export class OneofDescriptorProto extends jspb.Message { 

    hasName(): boolean;
    clearName(): void;
    getName(): string | undefined;
    setName(value: string): OneofDescriptorProto;

    hasOptions(): boolean;
    clearOptions(): void;
    getOptions(): OneofOptions | undefined;
    setOptions(value?: OneofOptions): OneofDescriptorProto;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): OneofDescriptorProto.AsObject;
    static toObject(includeInstance: boolean, msg: OneofDescriptorProto): OneofDescriptorProto.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: OneofDescriptorProto, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): OneofDescriptorProto;
    static deserializeBinaryFromReader(message: OneofDescriptorProto, reader: jspb.BinaryReader): OneofDescriptorProto;
}

export namespace OneofDescriptorProto {
    export type AsObject = {
        name?: string,
        options?: OneofOptions.AsObject,
    }
}

export class EnumDescriptorProto extends jspb.Message { 

    hasName(): boolean;
    clearName(): void;
    getName(): string | undefined;
    setName(value: string): EnumDescriptorProto;
    clearValueList(): void;
    getValueList(): Array<EnumValueDescriptorProto>;
    setValueList(value: Array<EnumValueDescriptorProto>): EnumDescriptorProto;
    addValue(value?: EnumValueDescriptorProto, index?: number): EnumValueDescriptorProto;

    hasOptions(): boolean;
    clearOptions(): void;
    getOptions(): EnumOptions | undefined;
    setOptions(value?: EnumOptions): EnumDescriptorProto;
    clearReservedRangeList(): void;
    getReservedRangeList(): Array<EnumDescriptorProto.EnumReservedRange>;
    setReservedRangeList(value: Array<EnumDescriptorProto.EnumReservedRange>): EnumDescriptorProto;
    addReservedRange(value?: EnumDescriptorProto.EnumReservedRange, index?: number): EnumDescriptorProto.EnumReservedRange;
    clearReservedNameList(): void;
    getReservedNameList(): Array<string>;
    setReservedNameList(value: Array<string>): EnumDescriptorProto;
    addReservedName(value: string, index?: number): string;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): EnumDescriptorProto.AsObject;
    static toObject(includeInstance: boolean, msg: EnumDescriptorProto): EnumDescriptorProto.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: EnumDescriptorProto, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): EnumDescriptorProto;
    static deserializeBinaryFromReader(message: EnumDescriptorProto, reader: jspb.BinaryReader): EnumDescriptorProto;
}

export namespace EnumDescriptorProto {
    export type AsObject = {
        name?: string,
        valueList: Array<EnumValueDescriptorProto.AsObject>,
        options?: EnumOptions.AsObject,
        reservedRangeList: Array<EnumDescriptorProto.EnumReservedRange.AsObject>,
        reservedNameList: Array<string>,
    }


    export class EnumReservedRange extends jspb.Message { 

        hasStart(): boolean;
        clearStart(): void;
        getStart(): number | undefined;
        setStart(value: number): EnumReservedRange;

        hasEnd(): boolean;
        clearEnd(): void;
        getEnd(): number | undefined;
        setEnd(value: number): EnumReservedRange;

        serializeBinary(): Uint8Array;
        toObject(includeInstance?: boolean): EnumReservedRange.AsObject;
        static toObject(includeInstance: boolean, msg: EnumReservedRange): EnumReservedRange.AsObject;
        static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
        static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
        static serializeBinaryToWriter(message: EnumReservedRange, writer: jspb.BinaryWriter): void;
        static deserializeBinary(bytes: Uint8Array): EnumReservedRange;
        static deserializeBinaryFromReader(message: EnumReservedRange, reader: jspb.BinaryReader): EnumReservedRange;
    }

    export namespace EnumReservedRange {
        export type AsObject = {
            start?: number,
            end?: number,
        }
    }

}

export class EnumValueDescriptorProto extends jspb.Message { 

    hasName(): boolean;
    clearName(): void;
    getName(): string | undefined;
    setName(value: string): EnumValueDescriptorProto;

    hasNumber(): boolean;
    clearNumber(): void;
    getNumber(): number | undefined;
    setNumber(value: number): EnumValueDescriptorProto;

    hasOptions(): boolean;
    clearOptions(): void;
    getOptions(): EnumValueOptions | undefined;
    setOptions(value?: EnumValueOptions): EnumValueDescriptorProto;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): EnumValueDescriptorProto.AsObject;
    static toObject(includeInstance: boolean, msg: EnumValueDescriptorProto): EnumValueDescriptorProto.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: EnumValueDescriptorProto, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): EnumValueDescriptorProto;
    static deserializeBinaryFromReader(message: EnumValueDescriptorProto, reader: jspb.BinaryReader): EnumValueDescriptorProto;
}

export namespace EnumValueDescriptorProto {
    export type AsObject = {
        name?: string,
        number?: number,
        options?: EnumValueOptions.AsObject,
    }
}

export class ServiceDescriptorProto extends jspb.Message { 

    hasName(): boolean;
    clearName(): void;
    getName(): string | undefined;
    setName(value: string): ServiceDescriptorProto;
    clearMethodList(): void;
    getMethodList(): Array<MethodDescriptorProto>;
    setMethodList(value: Array<MethodDescriptorProto>): ServiceDescriptorProto;
    addMethod(value?: MethodDescriptorProto, index?: number): MethodDescriptorProto;

    hasOptions(): boolean;
    clearOptions(): void;
    getOptions(): ServiceOptions | undefined;
    setOptions(value?: ServiceOptions): ServiceDescriptorProto;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ServiceDescriptorProto.AsObject;
    static toObject(includeInstance: boolean, msg: ServiceDescriptorProto): ServiceDescriptorProto.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ServiceDescriptorProto, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ServiceDescriptorProto;
    static deserializeBinaryFromReader(message: ServiceDescriptorProto, reader: jspb.BinaryReader): ServiceDescriptorProto;
}

export namespace ServiceDescriptorProto {
    export type AsObject = {
        name?: string,
        methodList: Array<MethodDescriptorProto.AsObject>,
        options?: ServiceOptions.AsObject,
    }
}

export class MethodDescriptorProto extends jspb.Message { 

    hasName(): boolean;
    clearName(): void;
    getName(): string | undefined;
    setName(value: string): MethodDescriptorProto;

    hasInputType(): boolean;
    clearInputType(): void;
    getInputType(): string | undefined;
    setInputType(value: string): MethodDescriptorProto;

    hasOutputType(): boolean;
    clearOutputType(): void;
    getOutputType(): string | undefined;
    setOutputType(value: string): MethodDescriptorProto;

    hasOptions(): boolean;
    clearOptions(): void;
    getOptions(): MethodOptions | undefined;
    setOptions(value?: MethodOptions): MethodDescriptorProto;

    hasClientStreaming(): boolean;
    clearClientStreaming(): void;
    getClientStreaming(): boolean | undefined;
    setClientStreaming(value: boolean): MethodDescriptorProto;

    hasServerStreaming(): boolean;
    clearServerStreaming(): void;
    getServerStreaming(): boolean | undefined;
    setServerStreaming(value: boolean): MethodDescriptorProto;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): MethodDescriptorProto.AsObject;
    static toObject(includeInstance: boolean, msg: MethodDescriptorProto): MethodDescriptorProto.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: MethodDescriptorProto, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): MethodDescriptorProto;
    static deserializeBinaryFromReader(message: MethodDescriptorProto, reader: jspb.BinaryReader): MethodDescriptorProto;
}

export namespace MethodDescriptorProto {
    export type AsObject = {
        name?: string,
        inputType?: string,
        outputType?: string,
        options?: MethodOptions.AsObject,
        clientStreaming?: boolean,
        serverStreaming?: boolean,
    }
}

export class FileOptions extends jspb.Message { 

    hasJavaPackage(): boolean;
    clearJavaPackage(): void;
    getJavaPackage(): string | undefined;
    setJavaPackage(value: string): FileOptions;

    hasJavaOuterClassname(): boolean;
    clearJavaOuterClassname(): void;
    getJavaOuterClassname(): string | undefined;
    setJavaOuterClassname(value: string): FileOptions;

    hasJavaMultipleFiles(): boolean;
    clearJavaMultipleFiles(): void;
    getJavaMultipleFiles(): boolean | undefined;
    setJavaMultipleFiles(value: boolean): FileOptions;

    hasJavaGenerateEqualsAndHash(): boolean;
    clearJavaGenerateEqualsAndHash(): void;
    getJavaGenerateEqualsAndHash(): boolean | undefined;
    setJavaGenerateEqualsAndHash(value: boolean): FileOptions;

    hasJavaStringCheckUtf8(): boolean;
    clearJavaStringCheckUtf8(): void;
    getJavaStringCheckUtf8(): boolean | undefined;
    setJavaStringCheckUtf8(value: boolean): FileOptions;

    hasOptimizeFor(): boolean;
    clearOptimizeFor(): void;
    getOptimizeFor(): FileOptions.OptimizeMode | undefined;
    setOptimizeFor(value: FileOptions.OptimizeMode): FileOptions;

    hasGoPackage(): boolean;
    clearGoPackage(): void;
    getGoPackage(): string | undefined;
    setGoPackage(value: string): FileOptions;

    hasCcGenericServices(): boolean;
    clearCcGenericServices(): void;
    getCcGenericServices(): boolean | undefined;
    setCcGenericServices(value: boolean): FileOptions;

    hasJavaGenericServices(): boolean;
    clearJavaGenericServices(): void;
    getJavaGenericServices(): boolean | undefined;
    setJavaGenericServices(value: boolean): FileOptions;

    hasPyGenericServices(): boolean;
    clearPyGenericServices(): void;
    getPyGenericServices(): boolean | undefined;
    setPyGenericServices(value: boolean): FileOptions;

    hasDeprecated(): boolean;
    clearDeprecated(): void;
    getDeprecated(): boolean | undefined;
    setDeprecated(value: boolean): FileOptions;

    hasCcEnableArenas(): boolean;
    clearCcEnableArenas(): void;
    getCcEnableArenas(): boolean | undefined;
    setCcEnableArenas(value: boolean): FileOptions;

    hasObjcClassPrefix(): boolean;
    clearObjcClassPrefix(): void;
    getObjcClassPrefix(): string | undefined;
    setObjcClassPrefix(value: string): FileOptions;

    hasCsharpNamespace(): boolean;
    clearCsharpNamespace(): void;
    getCsharpNamespace(): string | undefined;
    setCsharpNamespace(value: string): FileOptions;

    hasSwiftPrefix(): boolean;
    clearSwiftPrefix(): void;
    getSwiftPrefix(): string | undefined;
    setSwiftPrefix(value: string): FileOptions;

    hasPhpClassPrefix(): boolean;
    clearPhpClassPrefix(): void;
    getPhpClassPrefix(): string | undefined;
    setPhpClassPrefix(value: string): FileOptions;

    hasPhpNamespace(): boolean;
    clearPhpNamespace(): void;
    getPhpNamespace(): string | undefined;
    setPhpNamespace(value: string): FileOptions;

    hasPhpMetadataNamespace(): boolean;
    clearPhpMetadataNamespace(): void;
    getPhpMetadataNamespace(): string | undefined;
    setPhpMetadataNamespace(value: string): FileOptions;

    hasRubyPackage(): boolean;
    clearRubyPackage(): void;
    getRubyPackage(): string | undefined;
    setRubyPackage(value: string): FileOptions;

    hasFeatures(): boolean;
    clearFeatures(): void;
    getFeatures(): FeatureSet | undefined;
    setFeatures(value?: FeatureSet): FileOptions;
    clearUninterpretedOptionList(): void;
    getUninterpretedOptionList(): Array<UninterpretedOption>;
    setUninterpretedOptionList(value: Array<UninterpretedOption>): FileOptions;
    addUninterpretedOption(value?: UninterpretedOption, index?: number): UninterpretedOption;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): FileOptions.AsObject;
    static toObject(includeInstance: boolean, msg: FileOptions): FileOptions.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: FileOptions, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): FileOptions;
    static deserializeBinaryFromReader(message: FileOptions, reader: jspb.BinaryReader): FileOptions;
}

export namespace FileOptions {
    export type AsObject = {
        javaPackage?: string,
        javaOuterClassname?: string,
        javaMultipleFiles?: boolean,
        javaGenerateEqualsAndHash?: boolean,
        javaStringCheckUtf8?: boolean,
        optimizeFor?: FileOptions.OptimizeMode,
        goPackage?: string,
        ccGenericServices?: boolean,
        javaGenericServices?: boolean,
        pyGenericServices?: boolean,
        deprecated?: boolean,
        ccEnableArenas?: boolean,
        objcClassPrefix?: string,
        csharpNamespace?: string,
        swiftPrefix?: string,
        phpClassPrefix?: string,
        phpNamespace?: string,
        phpMetadataNamespace?: string,
        rubyPackage?: string,
        features?: FeatureSet.AsObject,
        uninterpretedOptionList: Array<UninterpretedOption.AsObject>,
    }

    export enum OptimizeMode {
    SPEED = 1,
    CODE_SIZE = 2,
    LITE_RUNTIME = 3,
    }

}

export class MessageOptions extends jspb.Message { 

    hasMessageSetWireFormat(): boolean;
    clearMessageSetWireFormat(): void;
    getMessageSetWireFormat(): boolean | undefined;
    setMessageSetWireFormat(value: boolean): MessageOptions;

    hasNoStandardDescriptorAccessor(): boolean;
    clearNoStandardDescriptorAccessor(): void;
    getNoStandardDescriptorAccessor(): boolean | undefined;
    setNoStandardDescriptorAccessor(value: boolean): MessageOptions;

    hasDeprecated(): boolean;
    clearDeprecated(): void;
    getDeprecated(): boolean | undefined;
    setDeprecated(value: boolean): MessageOptions;

    hasMapEntry(): boolean;
    clearMapEntry(): void;
    getMapEntry(): boolean | undefined;
    setMapEntry(value: boolean): MessageOptions;

    hasDeprecatedLegacyJsonFieldConflicts(): boolean;
    clearDeprecatedLegacyJsonFieldConflicts(): void;
    getDeprecatedLegacyJsonFieldConflicts(): boolean | undefined;
    setDeprecatedLegacyJsonFieldConflicts(value: boolean): MessageOptions;

    hasFeatures(): boolean;
    clearFeatures(): void;
    getFeatures(): FeatureSet | undefined;
    setFeatures(value?: FeatureSet): MessageOptions;
    clearUninterpretedOptionList(): void;
    getUninterpretedOptionList(): Array<UninterpretedOption>;
    setUninterpretedOptionList(value: Array<UninterpretedOption>): MessageOptions;
    addUninterpretedOption(value?: UninterpretedOption, index?: number): UninterpretedOption;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): MessageOptions.AsObject;
    static toObject(includeInstance: boolean, msg: MessageOptions): MessageOptions.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: MessageOptions, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): MessageOptions;
    static deserializeBinaryFromReader(message: MessageOptions, reader: jspb.BinaryReader): MessageOptions;
}

export namespace MessageOptions {
    export type AsObject = {
        messageSetWireFormat?: boolean,
        noStandardDescriptorAccessor?: boolean,
        deprecated?: boolean,
        mapEntry?: boolean,
        deprecatedLegacyJsonFieldConflicts?: boolean,
        features?: FeatureSet.AsObject,
        uninterpretedOptionList: Array<UninterpretedOption.AsObject>,
    }
}

export class FieldOptions extends jspb.Message { 

    hasCtype(): boolean;
    clearCtype(): void;
    getCtype(): FieldOptions.CType | undefined;
    setCtype(value: FieldOptions.CType): FieldOptions;

    hasPacked(): boolean;
    clearPacked(): void;
    getPacked(): boolean | undefined;
    setPacked(value: boolean): FieldOptions;

    hasJstype(): boolean;
    clearJstype(): void;
    getJstype(): FieldOptions.JSType | undefined;
    setJstype(value: FieldOptions.JSType): FieldOptions;

    hasLazy(): boolean;
    clearLazy(): void;
    getLazy(): boolean | undefined;
    setLazy(value: boolean): FieldOptions;

    hasUnverifiedLazy(): boolean;
    clearUnverifiedLazy(): void;
    getUnverifiedLazy(): boolean | undefined;
    setUnverifiedLazy(value: boolean): FieldOptions;

    hasDeprecated(): boolean;
    clearDeprecated(): void;
    getDeprecated(): boolean | undefined;
    setDeprecated(value: boolean): FieldOptions;

    hasWeak(): boolean;
    clearWeak(): void;
    getWeak(): boolean | undefined;
    setWeak(value: boolean): FieldOptions;

    hasDebugRedact(): boolean;
    clearDebugRedact(): void;
    getDebugRedact(): boolean | undefined;
    setDebugRedact(value: boolean): FieldOptions;

    hasRetention(): boolean;
    clearRetention(): void;
    getRetention(): FieldOptions.OptionRetention | undefined;
    setRetention(value: FieldOptions.OptionRetention): FieldOptions;
    clearTargetsList(): void;
    getTargetsList(): Array<FieldOptions.OptionTargetType>;
    setTargetsList(value: Array<FieldOptions.OptionTargetType>): FieldOptions;
    addTargets(value: FieldOptions.OptionTargetType, index?: number): FieldOptions.OptionTargetType;
    clearEditionDefaultsList(): void;
    getEditionDefaultsList(): Array<FieldOptions.EditionDefault>;
    setEditionDefaultsList(value: Array<FieldOptions.EditionDefault>): FieldOptions;
    addEditionDefaults(value?: FieldOptions.EditionDefault, index?: number): FieldOptions.EditionDefault;

    hasFeatures(): boolean;
    clearFeatures(): void;
    getFeatures(): FeatureSet | undefined;
    setFeatures(value?: FeatureSet): FieldOptions;

    hasFeatureSupport(): boolean;
    clearFeatureSupport(): void;
    getFeatureSupport(): FieldOptions.FeatureSupport | undefined;
    setFeatureSupport(value?: FieldOptions.FeatureSupport): FieldOptions;
    clearUninterpretedOptionList(): void;
    getUninterpretedOptionList(): Array<UninterpretedOption>;
    setUninterpretedOptionList(value: Array<UninterpretedOption>): FieldOptions;
    addUninterpretedOption(value?: UninterpretedOption, index?: number): UninterpretedOption;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): FieldOptions.AsObject;
    static toObject(includeInstance: boolean, msg: FieldOptions): FieldOptions.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: FieldOptions, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): FieldOptions;
    static deserializeBinaryFromReader(message: FieldOptions, reader: jspb.BinaryReader): FieldOptions;
}

export namespace FieldOptions {
    export type AsObject = {
        ctype?: FieldOptions.CType,
        packed?: boolean,
        jstype?: FieldOptions.JSType,
        lazy?: boolean,
        unverifiedLazy?: boolean,
        deprecated?: boolean,
        weak?: boolean,
        debugRedact?: boolean,
        retention?: FieldOptions.OptionRetention,
        targetsList: Array<FieldOptions.OptionTargetType>,
        editionDefaultsList: Array<FieldOptions.EditionDefault.AsObject>,
        features?: FeatureSet.AsObject,
        featureSupport?: FieldOptions.FeatureSupport.AsObject,
        uninterpretedOptionList: Array<UninterpretedOption.AsObject>,
    }


    export class EditionDefault extends jspb.Message { 

        hasEdition(): boolean;
        clearEdition(): void;
        getEdition(): Edition | undefined;
        setEdition(value: Edition): EditionDefault;

        hasValue(): boolean;
        clearValue(): void;
        getValue(): string | undefined;
        setValue(value: string): EditionDefault;

        serializeBinary(): Uint8Array;
        toObject(includeInstance?: boolean): EditionDefault.AsObject;
        static toObject(includeInstance: boolean, msg: EditionDefault): EditionDefault.AsObject;
        static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
        static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
        static serializeBinaryToWriter(message: EditionDefault, writer: jspb.BinaryWriter): void;
        static deserializeBinary(bytes: Uint8Array): EditionDefault;
        static deserializeBinaryFromReader(message: EditionDefault, reader: jspb.BinaryReader): EditionDefault;
    }

    export namespace EditionDefault {
        export type AsObject = {
            edition?: Edition,
            value?: string,
        }
    }

    export class FeatureSupport extends jspb.Message { 

        hasEditionIntroduced(): boolean;
        clearEditionIntroduced(): void;
        getEditionIntroduced(): Edition | undefined;
        setEditionIntroduced(value: Edition): FeatureSupport;

        hasEditionDeprecated(): boolean;
        clearEditionDeprecated(): void;
        getEditionDeprecated(): Edition | undefined;
        setEditionDeprecated(value: Edition): FeatureSupport;

        hasDeprecationWarning(): boolean;
        clearDeprecationWarning(): void;
        getDeprecationWarning(): string | undefined;
        setDeprecationWarning(value: string): FeatureSupport;

        hasEditionRemoved(): boolean;
        clearEditionRemoved(): void;
        getEditionRemoved(): Edition | undefined;
        setEditionRemoved(value: Edition): FeatureSupport;

        serializeBinary(): Uint8Array;
        toObject(includeInstance?: boolean): FeatureSupport.AsObject;
        static toObject(includeInstance: boolean, msg: FeatureSupport): FeatureSupport.AsObject;
        static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
        static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
        static serializeBinaryToWriter(message: FeatureSupport, writer: jspb.BinaryWriter): void;
        static deserializeBinary(bytes: Uint8Array): FeatureSupport;
        static deserializeBinaryFromReader(message: FeatureSupport, reader: jspb.BinaryReader): FeatureSupport;
    }

    export namespace FeatureSupport {
        export type AsObject = {
            editionIntroduced?: Edition,
            editionDeprecated?: Edition,
            deprecationWarning?: string,
            editionRemoved?: Edition,
        }
    }


    export enum CType {
    STRING = 0,
    CORD = 1,
    STRING_PIECE = 2,
    }

    export enum JSType {
    JS_NORMAL = 0,
    JS_STRING = 1,
    JS_NUMBER = 2,
    }

    export enum OptionRetention {
    RETENTION_UNKNOWN = 0,
    RETENTION_RUNTIME = 1,
    RETENTION_SOURCE = 2,
    }

    export enum OptionTargetType {
    TARGET_TYPE_UNKNOWN = 0,
    TARGET_TYPE_FILE = 1,
    TARGET_TYPE_EXTENSION_RANGE = 2,
    TARGET_TYPE_MESSAGE = 3,
    TARGET_TYPE_FIELD = 4,
    TARGET_TYPE_ONEOF = 5,
    TARGET_TYPE_ENUM = 6,
    TARGET_TYPE_ENUM_ENTRY = 7,
    TARGET_TYPE_SERVICE = 8,
    TARGET_TYPE_METHOD = 9,
    }

}

export class OneofOptions extends jspb.Message { 

    hasFeatures(): boolean;
    clearFeatures(): void;
    getFeatures(): FeatureSet | undefined;
    setFeatures(value?: FeatureSet): OneofOptions;
    clearUninterpretedOptionList(): void;
    getUninterpretedOptionList(): Array<UninterpretedOption>;
    setUninterpretedOptionList(value: Array<UninterpretedOption>): OneofOptions;
    addUninterpretedOption(value?: UninterpretedOption, index?: number): UninterpretedOption;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): OneofOptions.AsObject;
    static toObject(includeInstance: boolean, msg: OneofOptions): OneofOptions.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: OneofOptions, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): OneofOptions;
    static deserializeBinaryFromReader(message: OneofOptions, reader: jspb.BinaryReader): OneofOptions;
}

export namespace OneofOptions {
    export type AsObject = {
        features?: FeatureSet.AsObject,
        uninterpretedOptionList: Array<UninterpretedOption.AsObject>,
    }
}

export class EnumOptions extends jspb.Message { 

    hasAllowAlias(): boolean;
    clearAllowAlias(): void;
    getAllowAlias(): boolean | undefined;
    setAllowAlias(value: boolean): EnumOptions;

    hasDeprecated(): boolean;
    clearDeprecated(): void;
    getDeprecated(): boolean | undefined;
    setDeprecated(value: boolean): EnumOptions;

    hasDeprecatedLegacyJsonFieldConflicts(): boolean;
    clearDeprecatedLegacyJsonFieldConflicts(): void;
    getDeprecatedLegacyJsonFieldConflicts(): boolean | undefined;
    setDeprecatedLegacyJsonFieldConflicts(value: boolean): EnumOptions;

    hasFeatures(): boolean;
    clearFeatures(): void;
    getFeatures(): FeatureSet | undefined;
    setFeatures(value?: FeatureSet): EnumOptions;
    clearUninterpretedOptionList(): void;
    getUninterpretedOptionList(): Array<UninterpretedOption>;
    setUninterpretedOptionList(value: Array<UninterpretedOption>): EnumOptions;
    addUninterpretedOption(value?: UninterpretedOption, index?: number): UninterpretedOption;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): EnumOptions.AsObject;
    static toObject(includeInstance: boolean, msg: EnumOptions): EnumOptions.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: EnumOptions, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): EnumOptions;
    static deserializeBinaryFromReader(message: EnumOptions, reader: jspb.BinaryReader): EnumOptions;
}

export namespace EnumOptions {
    export type AsObject = {
        allowAlias?: boolean,
        deprecated?: boolean,
        deprecatedLegacyJsonFieldConflicts?: boolean,
        features?: FeatureSet.AsObject,
        uninterpretedOptionList: Array<UninterpretedOption.AsObject>,
    }
}

export class EnumValueOptions extends jspb.Message { 

    hasDeprecated(): boolean;
    clearDeprecated(): void;
    getDeprecated(): boolean | undefined;
    setDeprecated(value: boolean): EnumValueOptions;

    hasFeatures(): boolean;
    clearFeatures(): void;
    getFeatures(): FeatureSet | undefined;
    setFeatures(value?: FeatureSet): EnumValueOptions;

    hasDebugRedact(): boolean;
    clearDebugRedact(): void;
    getDebugRedact(): boolean | undefined;
    setDebugRedact(value: boolean): EnumValueOptions;

    hasFeatureSupport(): boolean;
    clearFeatureSupport(): void;
    getFeatureSupport(): FieldOptions.FeatureSupport | undefined;
    setFeatureSupport(value?: FieldOptions.FeatureSupport): EnumValueOptions;
    clearUninterpretedOptionList(): void;
    getUninterpretedOptionList(): Array<UninterpretedOption>;
    setUninterpretedOptionList(value: Array<UninterpretedOption>): EnumValueOptions;
    addUninterpretedOption(value?: UninterpretedOption, index?: number): UninterpretedOption;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): EnumValueOptions.AsObject;
    static toObject(includeInstance: boolean, msg: EnumValueOptions): EnumValueOptions.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: EnumValueOptions, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): EnumValueOptions;
    static deserializeBinaryFromReader(message: EnumValueOptions, reader: jspb.BinaryReader): EnumValueOptions;
}

export namespace EnumValueOptions {
    export type AsObject = {
        deprecated?: boolean,
        features?: FeatureSet.AsObject,
        debugRedact?: boolean,
        featureSupport?: FieldOptions.FeatureSupport.AsObject,
        uninterpretedOptionList: Array<UninterpretedOption.AsObject>,
    }
}

export class ServiceOptions extends jspb.Message { 

    hasFeatures(): boolean;
    clearFeatures(): void;
    getFeatures(): FeatureSet | undefined;
    setFeatures(value?: FeatureSet): ServiceOptions;

    hasDeprecated(): boolean;
    clearDeprecated(): void;
    getDeprecated(): boolean | undefined;
    setDeprecated(value: boolean): ServiceOptions;
    clearUninterpretedOptionList(): void;
    getUninterpretedOptionList(): Array<UninterpretedOption>;
    setUninterpretedOptionList(value: Array<UninterpretedOption>): ServiceOptions;
    addUninterpretedOption(value?: UninterpretedOption, index?: number): UninterpretedOption;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ServiceOptions.AsObject;
    static toObject(includeInstance: boolean, msg: ServiceOptions): ServiceOptions.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ServiceOptions, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ServiceOptions;
    static deserializeBinaryFromReader(message: ServiceOptions, reader: jspb.BinaryReader): ServiceOptions;
}

export namespace ServiceOptions {
    export type AsObject = {
        features?: FeatureSet.AsObject,
        deprecated?: boolean,
        uninterpretedOptionList: Array<UninterpretedOption.AsObject>,
    }
}

export class MethodOptions extends jspb.Message { 

    hasDeprecated(): boolean;
    clearDeprecated(): void;
    getDeprecated(): boolean | undefined;
    setDeprecated(value: boolean): MethodOptions;

    hasIdempotencyLevel(): boolean;
    clearIdempotencyLevel(): void;
    getIdempotencyLevel(): MethodOptions.IdempotencyLevel | undefined;
    setIdempotencyLevel(value: MethodOptions.IdempotencyLevel): MethodOptions;

    hasFeatures(): boolean;
    clearFeatures(): void;
    getFeatures(): FeatureSet | undefined;
    setFeatures(value?: FeatureSet): MethodOptions;
    clearUninterpretedOptionList(): void;
    getUninterpretedOptionList(): Array<UninterpretedOption>;
    setUninterpretedOptionList(value: Array<UninterpretedOption>): MethodOptions;
    addUninterpretedOption(value?: UninterpretedOption, index?: number): UninterpretedOption;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): MethodOptions.AsObject;
    static toObject(includeInstance: boolean, msg: MethodOptions): MethodOptions.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: MethodOptions, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): MethodOptions;
    static deserializeBinaryFromReader(message: MethodOptions, reader: jspb.BinaryReader): MethodOptions;
}

export namespace MethodOptions {
    export type AsObject = {
        deprecated?: boolean,
        idempotencyLevel?: MethodOptions.IdempotencyLevel,
        features?: FeatureSet.AsObject,
        uninterpretedOptionList: Array<UninterpretedOption.AsObject>,
    }

    export enum IdempotencyLevel {
    IDEMPOTENCY_UNKNOWN = 0,
    NO_SIDE_EFFECTS = 1,
    IDEMPOTENT = 2,
    }

}

export class UninterpretedOption extends jspb.Message { 
    clearNameList(): void;
    getNameList(): Array<UninterpretedOption.NamePart>;
    setNameList(value: Array<UninterpretedOption.NamePart>): UninterpretedOption;
    addName(value?: UninterpretedOption.NamePart, index?: number): UninterpretedOption.NamePart;

    hasIdentifierValue(): boolean;
    clearIdentifierValue(): void;
    getIdentifierValue(): string | undefined;
    setIdentifierValue(value: string): UninterpretedOption;

    hasPositiveIntValue(): boolean;
    clearPositiveIntValue(): void;
    getPositiveIntValue(): number | undefined;
    setPositiveIntValue(value: number): UninterpretedOption;

    hasNegativeIntValue(): boolean;
    clearNegativeIntValue(): void;
    getNegativeIntValue(): number | undefined;
    setNegativeIntValue(value: number): UninterpretedOption;

    hasDoubleValue(): boolean;
    clearDoubleValue(): void;
    getDoubleValue(): number | undefined;
    setDoubleValue(value: number): UninterpretedOption;

    hasStringValue(): boolean;
    clearStringValue(): void;
    getStringValue(): Uint8Array | string;
    getStringValue_asU8(): Uint8Array;
    getStringValue_asB64(): string;
    setStringValue(value: Uint8Array | string): UninterpretedOption;

    hasAggregateValue(): boolean;
    clearAggregateValue(): void;
    getAggregateValue(): string | undefined;
    setAggregateValue(value: string): UninterpretedOption;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): UninterpretedOption.AsObject;
    static toObject(includeInstance: boolean, msg: UninterpretedOption): UninterpretedOption.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: UninterpretedOption, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): UninterpretedOption;
    static deserializeBinaryFromReader(message: UninterpretedOption, reader: jspb.BinaryReader): UninterpretedOption;
}

export namespace UninterpretedOption {
    export type AsObject = {
        nameList: Array<UninterpretedOption.NamePart.AsObject>,
        identifierValue?: string,
        positiveIntValue?: number,
        negativeIntValue?: number,
        doubleValue?: number,
        stringValue: Uint8Array | string,
        aggregateValue?: string,
    }


    export class NamePart extends jspb.Message { 

        hasNamePart(): boolean;
        clearNamePart(): void;
        getNamePart(): string | undefined;
        setNamePart(value: string): NamePart;

        hasIsExtension(): boolean;
        clearIsExtension(): void;
        getIsExtension(): boolean | undefined;
        setIsExtension(value: boolean): NamePart;

        serializeBinary(): Uint8Array;
        toObject(includeInstance?: boolean): NamePart.AsObject;
        static toObject(includeInstance: boolean, msg: NamePart): NamePart.AsObject;
        static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
        static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
        static serializeBinaryToWriter(message: NamePart, writer: jspb.BinaryWriter): void;
        static deserializeBinary(bytes: Uint8Array): NamePart;
        static deserializeBinaryFromReader(message: NamePart, reader: jspb.BinaryReader): NamePart;
    }

    export namespace NamePart {
        export type AsObject = {
            namePart?: string,
            isExtension?: boolean,
        }
    }

}

export class FeatureSet extends jspb.Message { 

    hasFieldPresence(): boolean;
    clearFieldPresence(): void;
    getFieldPresence(): FeatureSet.FieldPresence | undefined;
    setFieldPresence(value: FeatureSet.FieldPresence): FeatureSet;

    hasEnumType(): boolean;
    clearEnumType(): void;
    getEnumType(): FeatureSet.EnumType | undefined;
    setEnumType(value: FeatureSet.EnumType): FeatureSet;

    hasRepeatedFieldEncoding(): boolean;
    clearRepeatedFieldEncoding(): void;
    getRepeatedFieldEncoding(): FeatureSet.RepeatedFieldEncoding | undefined;
    setRepeatedFieldEncoding(value: FeatureSet.RepeatedFieldEncoding): FeatureSet;

    hasUtf8Validation(): boolean;
    clearUtf8Validation(): void;
    getUtf8Validation(): FeatureSet.Utf8Validation | undefined;
    setUtf8Validation(value: FeatureSet.Utf8Validation): FeatureSet;

    hasMessageEncoding(): boolean;
    clearMessageEncoding(): void;
    getMessageEncoding(): FeatureSet.MessageEncoding | undefined;
    setMessageEncoding(value: FeatureSet.MessageEncoding): FeatureSet;

    hasJsonFormat(): boolean;
    clearJsonFormat(): void;
    getJsonFormat(): FeatureSet.JsonFormat | undefined;
    setJsonFormat(value: FeatureSet.JsonFormat): FeatureSet;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): FeatureSet.AsObject;
    static toObject(includeInstance: boolean, msg: FeatureSet): FeatureSet.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: FeatureSet, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): FeatureSet;
    static deserializeBinaryFromReader(message: FeatureSet, reader: jspb.BinaryReader): FeatureSet;
}

export namespace FeatureSet {
    export type AsObject = {
        fieldPresence?: FeatureSet.FieldPresence,
        enumType?: FeatureSet.EnumType,
        repeatedFieldEncoding?: FeatureSet.RepeatedFieldEncoding,
        utf8Validation?: FeatureSet.Utf8Validation,
        messageEncoding?: FeatureSet.MessageEncoding,
        jsonFormat?: FeatureSet.JsonFormat,
    }

    export enum FieldPresence {
    FIELD_PRESENCE_UNKNOWN = 0,
    EXPLICIT = 1,
    IMPLICIT = 2,
    LEGACY_REQUIRED = 3,
    }

    export enum EnumType {
    ENUM_TYPE_UNKNOWN = 0,
    OPEN = 1,
    CLOSED = 2,
    }

    export enum RepeatedFieldEncoding {
    REPEATED_FIELD_ENCODING_UNKNOWN = 0,
    PACKED = 1,
    EXPANDED = 2,
    }

    export enum Utf8Validation {
    UTF8_VALIDATION_UNKNOWN = 0,
    VERIFY = 2,
    NONE = 3,
    }

    export enum MessageEncoding {
    MESSAGE_ENCODING_UNKNOWN = 0,
    LENGTH_PREFIXED = 1,
    DELIMITED = 2,
    }

    export enum JsonFormat {
    JSON_FORMAT_UNKNOWN = 0,
    ALLOW = 1,
    LEGACY_BEST_EFFORT = 2,
    }

}

export class FeatureSetDefaults extends jspb.Message { 
    clearDefaultsList(): void;
    getDefaultsList(): Array<FeatureSetDefaults.FeatureSetEditionDefault>;
    setDefaultsList(value: Array<FeatureSetDefaults.FeatureSetEditionDefault>): FeatureSetDefaults;
    addDefaults(value?: FeatureSetDefaults.FeatureSetEditionDefault, index?: number): FeatureSetDefaults.FeatureSetEditionDefault;

    hasMinimumEdition(): boolean;
    clearMinimumEdition(): void;
    getMinimumEdition(): Edition | undefined;
    setMinimumEdition(value: Edition): FeatureSetDefaults;

    hasMaximumEdition(): boolean;
    clearMaximumEdition(): void;
    getMaximumEdition(): Edition | undefined;
    setMaximumEdition(value: Edition): FeatureSetDefaults;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): FeatureSetDefaults.AsObject;
    static toObject(includeInstance: boolean, msg: FeatureSetDefaults): FeatureSetDefaults.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: FeatureSetDefaults, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): FeatureSetDefaults;
    static deserializeBinaryFromReader(message: FeatureSetDefaults, reader: jspb.BinaryReader): FeatureSetDefaults;
}

export namespace FeatureSetDefaults {
    export type AsObject = {
        defaultsList: Array<FeatureSetDefaults.FeatureSetEditionDefault.AsObject>,
        minimumEdition?: Edition,
        maximumEdition?: Edition,
    }


    export class FeatureSetEditionDefault extends jspb.Message { 

        hasEdition(): boolean;
        clearEdition(): void;
        getEdition(): Edition | undefined;
        setEdition(value: Edition): FeatureSetEditionDefault;

        hasOverridableFeatures(): boolean;
        clearOverridableFeatures(): void;
        getOverridableFeatures(): FeatureSet | undefined;
        setOverridableFeatures(value?: FeatureSet): FeatureSetEditionDefault;

        hasFixedFeatures(): boolean;
        clearFixedFeatures(): void;
        getFixedFeatures(): FeatureSet | undefined;
        setFixedFeatures(value?: FeatureSet): FeatureSetEditionDefault;

        serializeBinary(): Uint8Array;
        toObject(includeInstance?: boolean): FeatureSetEditionDefault.AsObject;
        static toObject(includeInstance: boolean, msg: FeatureSetEditionDefault): FeatureSetEditionDefault.AsObject;
        static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
        static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
        static serializeBinaryToWriter(message: FeatureSetEditionDefault, writer: jspb.BinaryWriter): void;
        static deserializeBinary(bytes: Uint8Array): FeatureSetEditionDefault;
        static deserializeBinaryFromReader(message: FeatureSetEditionDefault, reader: jspb.BinaryReader): FeatureSetEditionDefault;
    }

    export namespace FeatureSetEditionDefault {
        export type AsObject = {
            edition?: Edition,
            overridableFeatures?: FeatureSet.AsObject,
            fixedFeatures?: FeatureSet.AsObject,
        }
    }

}

export class SourceCodeInfo extends jspb.Message { 
    clearLocationList(): void;
    getLocationList(): Array<SourceCodeInfo.Location>;
    setLocationList(value: Array<SourceCodeInfo.Location>): SourceCodeInfo;
    addLocation(value?: SourceCodeInfo.Location, index?: number): SourceCodeInfo.Location;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): SourceCodeInfo.AsObject;
    static toObject(includeInstance: boolean, msg: SourceCodeInfo): SourceCodeInfo.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: SourceCodeInfo, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): SourceCodeInfo;
    static deserializeBinaryFromReader(message: SourceCodeInfo, reader: jspb.BinaryReader): SourceCodeInfo;
}

export namespace SourceCodeInfo {
    export type AsObject = {
        locationList: Array<SourceCodeInfo.Location.AsObject>,
    }


    export class Location extends jspb.Message { 
        clearPathList(): void;
        getPathList(): Array<number>;
        setPathList(value: Array<number>): Location;
        addPath(value: number, index?: number): number;
        clearSpanList(): void;
        getSpanList(): Array<number>;
        setSpanList(value: Array<number>): Location;
        addSpan(value: number, index?: number): number;

        hasLeadingComments(): boolean;
        clearLeadingComments(): void;
        getLeadingComments(): string | undefined;
        setLeadingComments(value: string): Location;

        hasTrailingComments(): boolean;
        clearTrailingComments(): void;
        getTrailingComments(): string | undefined;
        setTrailingComments(value: string): Location;
        clearLeadingDetachedCommentsList(): void;
        getLeadingDetachedCommentsList(): Array<string>;
        setLeadingDetachedCommentsList(value: Array<string>): Location;
        addLeadingDetachedComments(value: string, index?: number): string;

        serializeBinary(): Uint8Array;
        toObject(includeInstance?: boolean): Location.AsObject;
        static toObject(includeInstance: boolean, msg: Location): Location.AsObject;
        static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
        static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
        static serializeBinaryToWriter(message: Location, writer: jspb.BinaryWriter): void;
        static deserializeBinary(bytes: Uint8Array): Location;
        static deserializeBinaryFromReader(message: Location, reader: jspb.BinaryReader): Location;
    }

    export namespace Location {
        export type AsObject = {
            pathList: Array<number>,
            spanList: Array<number>,
            leadingComments?: string,
            trailingComments?: string,
            leadingDetachedCommentsList: Array<string>,
        }
    }

}

export class GeneratedCodeInfo extends jspb.Message { 
    clearAnnotationList(): void;
    getAnnotationList(): Array<GeneratedCodeInfo.Annotation>;
    setAnnotationList(value: Array<GeneratedCodeInfo.Annotation>): GeneratedCodeInfo;
    addAnnotation(value?: GeneratedCodeInfo.Annotation, index?: number): GeneratedCodeInfo.Annotation;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): GeneratedCodeInfo.AsObject;
    static toObject(includeInstance: boolean, msg: GeneratedCodeInfo): GeneratedCodeInfo.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: GeneratedCodeInfo, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): GeneratedCodeInfo;
    static deserializeBinaryFromReader(message: GeneratedCodeInfo, reader: jspb.BinaryReader): GeneratedCodeInfo;
}

export namespace GeneratedCodeInfo {
    export type AsObject = {
        annotationList: Array<GeneratedCodeInfo.Annotation.AsObject>,
    }


    export class Annotation extends jspb.Message { 
        clearPathList(): void;
        getPathList(): Array<number>;
        setPathList(value: Array<number>): Annotation;
        addPath(value: number, index?: number): number;

        hasSourceFile(): boolean;
        clearSourceFile(): void;
        getSourceFile(): string | undefined;
        setSourceFile(value: string): Annotation;

        hasBegin(): boolean;
        clearBegin(): void;
        getBegin(): number | undefined;
        setBegin(value: number): Annotation;

        hasEnd(): boolean;
        clearEnd(): void;
        getEnd(): number | undefined;
        setEnd(value: number): Annotation;

        hasSemantic(): boolean;
        clearSemantic(): void;
        getSemantic(): GeneratedCodeInfo.Annotation.Semantic | undefined;
        setSemantic(value: GeneratedCodeInfo.Annotation.Semantic): Annotation;

        serializeBinary(): Uint8Array;
        toObject(includeInstance?: boolean): Annotation.AsObject;
        static toObject(includeInstance: boolean, msg: Annotation): Annotation.AsObject;
        static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
        static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
        static serializeBinaryToWriter(message: Annotation, writer: jspb.BinaryWriter): void;
        static deserializeBinary(bytes: Uint8Array): Annotation;
        static deserializeBinaryFromReader(message: Annotation, reader: jspb.BinaryReader): Annotation;
    }

    export namespace Annotation {
        export type AsObject = {
            pathList: Array<number>,
            sourceFile?: string,
            begin?: number,
            end?: number,
            semantic?: GeneratedCodeInfo.Annotation.Semantic,
        }

        export enum Semantic {
    NONE = 0,
    SET = 1,
    ALIAS = 2,
        }

    }

}

export enum Edition {
    EDITION_UNKNOWN = 0,
    EDITION_LEGACY = 900,
    EDITION_PROTO2 = 998,
    EDITION_PROTO3 = 999,
    EDITION_2023 = 1000,
    EDITION_2024 = 1001,
    EDITION_1_TEST_ONLY = 1,
    EDITION_2_TEST_ONLY = 2,
    EDITION_99997_TEST_ONLY = 99997,
    EDITION_99998_TEST_ONLY = 99998,
    EDITION_99999_TEST_ONLY = 99999,
    EDITION_MAX = 2147483647,
}
