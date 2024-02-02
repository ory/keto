/* eslint-disable */
import _m0 from "protobufjs/minimal.js";

export const protobufPackage = "ory.keto.relation_tuples.v1alpha2";

/** RelationTuple defines a relation between an Object and a Subject. */
export interface RelationTuple {
  /** The namespace this relation tuple lives in. */
  namespace: string;
  /**
   * The object related by this tuple.
   * It is an object in the namespace of the tuple.
   */
  object: string;
  /** The relation between an Object and a Subject. */
  relation: string;
  /**
   * The subject related by this tuple.
   * A Subject either represents a concrete subject id or
   * a `SubjectSet` that expands to more Subjects.
   */
  subject: Subject | undefined;
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
export interface RelationQuery {
  /** The namespace this relation tuple lives in. */
  namespace?:
    | string
    | undefined;
  /**
   * The object related by this tuple.
   * It is an object in the namespace of the tuple.
   */
  object?:
    | string
    | undefined;
  /** The relation between an Object and a Subject. */
  relation?:
    | string
    | undefined;
  /**
   * The subject related by this tuple.
   * A Subject either represents a concrete subject id or
   * a `SubjectSet` that expands to more Subjects.
   */
  subject?: Subject | undefined;
}

/**
 * Subject is either a concrete subject id or
 * a `SubjectSet` expanding to more Subjects.
 */
export interface Subject {
  /** A concrete id of the subject. */
  id?:
    | string
    | undefined;
  /**
   * A subject set that expands to more Subjects.
   * More information are available under [concepts](../concepts/15_subjects.mdx).
   */
  set?: SubjectSet | undefined;
}

/**
 * SubjectSet refers to all subjects who have
 * the same `relation` on an `object`.
 */
export interface SubjectSet {
  /**
   * The namespace of the object and relation
   * referenced in this subject set.
   */
  namespace: string;
  /** The object related by this subject set. */
  object: string;
  /** The relation between the object and the subjects. */
  relation: string;
}

function createBaseRelationTuple(): RelationTuple {
  return { namespace: "", object: "", relation: "", subject: undefined };
}

export const RelationTuple = {
  encode(message: RelationTuple, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
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

  decode(input: _m0.Reader | Uint8Array, length?: number): RelationTuple {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseRelationTuple();
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

  fromJSON(object: any): RelationTuple {
    return {
      namespace: isSet(object.namespace) ? globalThis.String(object.namespace) : "",
      object: isSet(object.object) ? globalThis.String(object.object) : "",
      relation: isSet(object.relation) ? globalThis.String(object.relation) : "",
      subject: isSet(object.subject) ? Subject.fromJSON(object.subject) : undefined,
    };
  },

  toJSON(message: RelationTuple): unknown {
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

  create<I extends Exact<DeepPartial<RelationTuple>, I>>(base?: I): RelationTuple {
    return RelationTuple.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<RelationTuple>, I>>(object: I): RelationTuple {
    const message = createBaseRelationTuple();
    message.namespace = object.namespace ?? "";
    message.object = object.object ?? "";
    message.relation = object.relation ?? "";
    message.subject = (object.subject !== undefined && object.subject !== null)
      ? Subject.fromPartial(object.subject)
      : undefined;
    return message;
  },
};

function createBaseRelationQuery(): RelationQuery {
  return { namespace: undefined, object: undefined, relation: undefined, subject: undefined };
}

export const RelationQuery = {
  encode(message: RelationQuery, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.namespace !== undefined) {
      writer.uint32(10).string(message.namespace);
    }
    if (message.object !== undefined) {
      writer.uint32(18).string(message.object);
    }
    if (message.relation !== undefined) {
      writer.uint32(26).string(message.relation);
    }
    if (message.subject !== undefined) {
      Subject.encode(message.subject, writer.uint32(34).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): RelationQuery {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseRelationQuery();
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

  fromJSON(object: any): RelationQuery {
    return {
      namespace: isSet(object.namespace) ? globalThis.String(object.namespace) : undefined,
      object: isSet(object.object) ? globalThis.String(object.object) : undefined,
      relation: isSet(object.relation) ? globalThis.String(object.relation) : undefined,
      subject: isSet(object.subject) ? Subject.fromJSON(object.subject) : undefined,
    };
  },

  toJSON(message: RelationQuery): unknown {
    const obj: any = {};
    if (message.namespace !== undefined) {
      obj.namespace = message.namespace;
    }
    if (message.object !== undefined) {
      obj.object = message.object;
    }
    if (message.relation !== undefined) {
      obj.relation = message.relation;
    }
    if (message.subject !== undefined) {
      obj.subject = Subject.toJSON(message.subject);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<RelationQuery>, I>>(base?: I): RelationQuery {
    return RelationQuery.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<RelationQuery>, I>>(object: I): RelationQuery {
    const message = createBaseRelationQuery();
    message.namespace = object.namespace ?? undefined;
    message.object = object.object ?? undefined;
    message.relation = object.relation ?? undefined;
    message.subject = (object.subject !== undefined && object.subject !== null)
      ? Subject.fromPartial(object.subject)
      : undefined;
    return message;
  },
};

function createBaseSubject(): Subject {
  return { id: undefined, set: undefined };
}

export const Subject = {
  encode(message: Subject, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== undefined) {
      writer.uint32(10).string(message.id);
    }
    if (message.set !== undefined) {
      SubjectSet.encode(message.set, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Subject {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSubject();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.id = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.set = SubjectSet.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): Subject {
    return {
      id: isSet(object.id) ? globalThis.String(object.id) : undefined,
      set: isSet(object.set) ? SubjectSet.fromJSON(object.set) : undefined,
    };
  },

  toJSON(message: Subject): unknown {
    const obj: any = {};
    if (message.id !== undefined) {
      obj.id = message.id;
    }
    if (message.set !== undefined) {
      obj.set = SubjectSet.toJSON(message.set);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<Subject>, I>>(base?: I): Subject {
    return Subject.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<Subject>, I>>(object: I): Subject {
    const message = createBaseSubject();
    message.id = object.id ?? undefined;
    message.set = (object.set !== undefined && object.set !== null) ? SubjectSet.fromPartial(object.set) : undefined;
    return message;
  },
};

function createBaseSubjectSet(): SubjectSet {
  return { namespace: "", object: "", relation: "" };
}

export const SubjectSet = {
  encode(message: SubjectSet, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.namespace !== "") {
      writer.uint32(10).string(message.namespace);
    }
    if (message.object !== "") {
      writer.uint32(18).string(message.object);
    }
    if (message.relation !== "") {
      writer.uint32(26).string(message.relation);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SubjectSet {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSubjectSet();
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
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): SubjectSet {
    return {
      namespace: isSet(object.namespace) ? globalThis.String(object.namespace) : "",
      object: isSet(object.object) ? globalThis.String(object.object) : "",
      relation: isSet(object.relation) ? globalThis.String(object.relation) : "",
    };
  },

  toJSON(message: SubjectSet): unknown {
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
    return obj;
  },

  create<I extends Exact<DeepPartial<SubjectSet>, I>>(base?: I): SubjectSet {
    return SubjectSet.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<SubjectSet>, I>>(object: I): SubjectSet {
    const message = createBaseSubjectSet();
    message.namespace = object.namespace ?? "";
    message.object = object.object ?? "";
    message.relation = object.relation ?? "";
    return message;
  },
};

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
