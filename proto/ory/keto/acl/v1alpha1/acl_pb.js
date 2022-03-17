// source: ory/keto/acl/v1alpha1/acl.proto
/**
 * @fileoverview
 * @enhanceable
 * @suppress {missingRequire} reports error on implicit type usages.
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!
/* eslint-disable */
// @ts-nocheck

var jspb = require('google-protobuf');
var goog = jspb;
var global = (function() {
  if (this) { return this; }
  if (typeof window !== 'undefined') { return window; }
  if (typeof global !== 'undefined') { return global; }
  if (typeof self !== 'undefined') { return self; }
  return Function('return this')();
}.call(null));

goog.exportSymbol('proto.ory.keto.acl.v1alpha1.RelationTuple', null, global);
goog.exportSymbol('proto.ory.keto.acl.v1alpha1.Subject', null, global);
goog.exportSymbol('proto.ory.keto.acl.v1alpha1.Subject.RefCase', null, global);
goog.exportSymbol('proto.ory.keto.acl.v1alpha1.SubjectSet', null, global);
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.ory.keto.acl.v1alpha1.RelationTuple = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.ory.keto.acl.v1alpha1.RelationTuple, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.ory.keto.acl.v1alpha1.RelationTuple.displayName = 'proto.ory.keto.acl.v1alpha1.RelationTuple';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.ory.keto.acl.v1alpha1.Subject = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, proto.ory.keto.acl.v1alpha1.Subject.oneofGroups_);
};
goog.inherits(proto.ory.keto.acl.v1alpha1.Subject, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.ory.keto.acl.v1alpha1.Subject.displayName = 'proto.ory.keto.acl.v1alpha1.Subject';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.ory.keto.acl.v1alpha1.SubjectSet = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.ory.keto.acl.v1alpha1.SubjectSet, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.ory.keto.acl.v1alpha1.SubjectSet.displayName = 'proto.ory.keto.acl.v1alpha1.SubjectSet';
}



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.ory.keto.acl.v1alpha1.RelationTuple.prototype.toObject = function(opt_includeInstance) {
  return proto.ory.keto.acl.v1alpha1.RelationTuple.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.ory.keto.acl.v1alpha1.RelationTuple} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.ory.keto.acl.v1alpha1.RelationTuple.toObject = function(includeInstance, msg) {
  var f, obj = {
    namespace: jspb.Message.getFieldWithDefault(msg, 1, ""),
    object: jspb.Message.getFieldWithDefault(msg, 2, ""),
    relation: jspb.Message.getFieldWithDefault(msg, 3, ""),
    subject: (f = msg.getSubject()) && proto.ory.keto.acl.v1alpha1.Subject.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.ory.keto.acl.v1alpha1.RelationTuple}
 */
proto.ory.keto.acl.v1alpha1.RelationTuple.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.ory.keto.acl.v1alpha1.RelationTuple;
  return proto.ory.keto.acl.v1alpha1.RelationTuple.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.ory.keto.acl.v1alpha1.RelationTuple} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.ory.keto.acl.v1alpha1.RelationTuple}
 */
proto.ory.keto.acl.v1alpha1.RelationTuple.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setNamespace(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setObject(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setRelation(value);
      break;
    case 4:
      var value = new proto.ory.keto.acl.v1alpha1.Subject;
      reader.readMessage(value,proto.ory.keto.acl.v1alpha1.Subject.deserializeBinaryFromReader);
      msg.setSubject(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.ory.keto.acl.v1alpha1.RelationTuple.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.ory.keto.acl.v1alpha1.RelationTuple.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.ory.keto.acl.v1alpha1.RelationTuple} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.ory.keto.acl.v1alpha1.RelationTuple.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getNamespace();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getObject();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getRelation();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getSubject();
  if (f != null) {
    writer.writeMessage(
      4,
      f,
      proto.ory.keto.acl.v1alpha1.Subject.serializeBinaryToWriter
    );
  }
};


/**
 * optional string namespace = 1;
 * @return {string}
 */
proto.ory.keto.acl.v1alpha1.RelationTuple.prototype.getNamespace = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.ory.keto.acl.v1alpha1.RelationTuple} returns this
 */
proto.ory.keto.acl.v1alpha1.RelationTuple.prototype.setNamespace = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string object = 2;
 * @return {string}
 */
proto.ory.keto.acl.v1alpha1.RelationTuple.prototype.getObject = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.ory.keto.acl.v1alpha1.RelationTuple} returns this
 */
proto.ory.keto.acl.v1alpha1.RelationTuple.prototype.setObject = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string relation = 3;
 * @return {string}
 */
proto.ory.keto.acl.v1alpha1.RelationTuple.prototype.getRelation = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/**
 * @param {string} value
 * @return {!proto.ory.keto.acl.v1alpha1.RelationTuple} returns this
 */
proto.ory.keto.acl.v1alpha1.RelationTuple.prototype.setRelation = function(value) {
  return jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional Subject subject = 4;
 * @return {?proto.ory.keto.acl.v1alpha1.Subject}
 */
proto.ory.keto.acl.v1alpha1.RelationTuple.prototype.getSubject = function() {
  return /** @type{?proto.ory.keto.acl.v1alpha1.Subject} */ (
    jspb.Message.getWrapperField(this, proto.ory.keto.acl.v1alpha1.Subject, 4));
};


/**
 * @param {?proto.ory.keto.acl.v1alpha1.Subject|undefined} value
 * @return {!proto.ory.keto.acl.v1alpha1.RelationTuple} returns this
*/
proto.ory.keto.acl.v1alpha1.RelationTuple.prototype.setSubject = function(value) {
  return jspb.Message.setWrapperField(this, 4, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.ory.keto.acl.v1alpha1.RelationTuple} returns this
 */
proto.ory.keto.acl.v1alpha1.RelationTuple.prototype.clearSubject = function() {
  return this.setSubject(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.ory.keto.acl.v1alpha1.RelationTuple.prototype.hasSubject = function() {
  return jspb.Message.getField(this, 4) != null;
};



/**
 * Oneof group definitions for this message. Each group defines the field
 * numbers belonging to that group. When of these fields' value is set, all
 * other fields in the group are cleared. During deserialization, if multiple
 * fields are encountered for a group, only the last value seen will be kept.
 * @private {!Array<!Array<number>>}
 * @const
 */
proto.ory.keto.acl.v1alpha1.Subject.oneofGroups_ = [[1,2]];

/**
 * @enum {number}
 */
proto.ory.keto.acl.v1alpha1.Subject.RefCase = {
  REF_NOT_SET: 0,
  ID: 1,
  SET: 2
};

/**
 * @return {proto.ory.keto.acl.v1alpha1.Subject.RefCase}
 */
proto.ory.keto.acl.v1alpha1.Subject.prototype.getRefCase = function() {
  return /** @type {proto.ory.keto.acl.v1alpha1.Subject.RefCase} */(jspb.Message.computeOneofCase(this, proto.ory.keto.acl.v1alpha1.Subject.oneofGroups_[0]));
};



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.ory.keto.acl.v1alpha1.Subject.prototype.toObject = function(opt_includeInstance) {
  return proto.ory.keto.acl.v1alpha1.Subject.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.ory.keto.acl.v1alpha1.Subject} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.ory.keto.acl.v1alpha1.Subject.toObject = function(includeInstance, msg) {
  var f, obj = {
    id: jspb.Message.getFieldWithDefault(msg, 1, ""),
    set: (f = msg.getSet()) && proto.ory.keto.acl.v1alpha1.SubjectSet.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.ory.keto.acl.v1alpha1.Subject}
 */
proto.ory.keto.acl.v1alpha1.Subject.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.ory.keto.acl.v1alpha1.Subject;
  return proto.ory.keto.acl.v1alpha1.Subject.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.ory.keto.acl.v1alpha1.Subject} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.ory.keto.acl.v1alpha1.Subject}
 */
proto.ory.keto.acl.v1alpha1.Subject.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setId(value);
      break;
    case 2:
      var value = new proto.ory.keto.acl.v1alpha1.SubjectSet;
      reader.readMessage(value,proto.ory.keto.acl.v1alpha1.SubjectSet.deserializeBinaryFromReader);
      msg.setSet(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.ory.keto.acl.v1alpha1.Subject.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.ory.keto.acl.v1alpha1.Subject.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.ory.keto.acl.v1alpha1.Subject} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.ory.keto.acl.v1alpha1.Subject.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = /** @type {string} */ (jspb.Message.getField(message, 1));
  if (f != null) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getSet();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      proto.ory.keto.acl.v1alpha1.SubjectSet.serializeBinaryToWriter
    );
  }
};


/**
 * optional string id = 1;
 * @return {string}
 */
proto.ory.keto.acl.v1alpha1.Subject.prototype.getId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.ory.keto.acl.v1alpha1.Subject} returns this
 */
proto.ory.keto.acl.v1alpha1.Subject.prototype.setId = function(value) {
  return jspb.Message.setOneofField(this, 1, proto.ory.keto.acl.v1alpha1.Subject.oneofGroups_[0], value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.ory.keto.acl.v1alpha1.Subject} returns this
 */
proto.ory.keto.acl.v1alpha1.Subject.prototype.clearId = function() {
  return jspb.Message.setOneofField(this, 1, proto.ory.keto.acl.v1alpha1.Subject.oneofGroups_[0], undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.ory.keto.acl.v1alpha1.Subject.prototype.hasId = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional SubjectSet set = 2;
 * @return {?proto.ory.keto.acl.v1alpha1.SubjectSet}
 */
proto.ory.keto.acl.v1alpha1.Subject.prototype.getSet = function() {
  return /** @type{?proto.ory.keto.acl.v1alpha1.SubjectSet} */ (
    jspb.Message.getWrapperField(this, proto.ory.keto.acl.v1alpha1.SubjectSet, 2));
};


/**
 * @param {?proto.ory.keto.acl.v1alpha1.SubjectSet|undefined} value
 * @return {!proto.ory.keto.acl.v1alpha1.Subject} returns this
*/
proto.ory.keto.acl.v1alpha1.Subject.prototype.setSet = function(value) {
  return jspb.Message.setOneofWrapperField(this, 2, proto.ory.keto.acl.v1alpha1.Subject.oneofGroups_[0], value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.ory.keto.acl.v1alpha1.Subject} returns this
 */
proto.ory.keto.acl.v1alpha1.Subject.prototype.clearSet = function() {
  return this.setSet(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.ory.keto.acl.v1alpha1.Subject.prototype.hasSet = function() {
  return jspb.Message.getField(this, 2) != null;
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.ory.keto.acl.v1alpha1.SubjectSet.prototype.toObject = function(opt_includeInstance) {
  return proto.ory.keto.acl.v1alpha1.SubjectSet.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.ory.keto.acl.v1alpha1.SubjectSet} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.ory.keto.acl.v1alpha1.SubjectSet.toObject = function(includeInstance, msg) {
  var f, obj = {
    namespace: jspb.Message.getFieldWithDefault(msg, 1, ""),
    object: jspb.Message.getFieldWithDefault(msg, 2, ""),
    relation: jspb.Message.getFieldWithDefault(msg, 3, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.ory.keto.acl.v1alpha1.SubjectSet}
 */
proto.ory.keto.acl.v1alpha1.SubjectSet.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.ory.keto.acl.v1alpha1.SubjectSet;
  return proto.ory.keto.acl.v1alpha1.SubjectSet.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.ory.keto.acl.v1alpha1.SubjectSet} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.ory.keto.acl.v1alpha1.SubjectSet}
 */
proto.ory.keto.acl.v1alpha1.SubjectSet.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setNamespace(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setObject(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setRelation(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.ory.keto.acl.v1alpha1.SubjectSet.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.ory.keto.acl.v1alpha1.SubjectSet.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.ory.keto.acl.v1alpha1.SubjectSet} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.ory.keto.acl.v1alpha1.SubjectSet.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getNamespace();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getObject();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getRelation();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
};


/**
 * optional string namespace = 1;
 * @return {string}
 */
proto.ory.keto.acl.v1alpha1.SubjectSet.prototype.getNamespace = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.ory.keto.acl.v1alpha1.SubjectSet} returns this
 */
proto.ory.keto.acl.v1alpha1.SubjectSet.prototype.setNamespace = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string object = 2;
 * @return {string}
 */
proto.ory.keto.acl.v1alpha1.SubjectSet.prototype.getObject = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.ory.keto.acl.v1alpha1.SubjectSet} returns this
 */
proto.ory.keto.acl.v1alpha1.SubjectSet.prototype.setObject = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string relation = 3;
 * @return {string}
 */
proto.ory.keto.acl.v1alpha1.SubjectSet.prototype.getRelation = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/**
 * @param {string} value
 * @return {!proto.ory.keto.acl.v1alpha1.SubjectSet} returns this
 */
proto.ory.keto.acl.v1alpha1.SubjectSet.prototype.setRelation = function(value) {
  return jspb.Message.setProto3StringField(this, 3, value);
};


goog.object.extend(exports, proto.ory.keto.acl.v1alpha1);
