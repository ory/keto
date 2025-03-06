// source: ory/keto/relation_tuples/v1alpha2/openapi.proto
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
var global =
    (typeof globalThis !== 'undefined' && globalThis) ||
    (typeof window !== 'undefined' && window) ||
    (typeof global !== 'undefined' && global) ||
    (typeof self !== 'undefined' && self) ||
    (function () { return this; }).call(null) ||
    Function('return this')();

var google_api_field_behavior_pb = require('../../../../google/api/field_behavior_pb.js');
goog.object.extend(proto, google_api_field_behavior_pb);
var protoc$gen$openapiv2_options_annotations_pb = require('../../../../protoc-gen-openapiv2/options/annotations_pb.js');
goog.object.extend(proto, protoc$gen$openapiv2_options_annotations_pb);
goog.exportSymbol('proto.ory.keto.relation_tuples.v1alpha2.ErrorResponse', null, global);
goog.exportSymbol('proto.ory.keto.relation_tuples.v1alpha2.ErrorResponse.Error', null, global);
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
proto.ory.keto.relation_tuples.v1alpha2.ErrorResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.ory.keto.relation_tuples.v1alpha2.ErrorResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.ory.keto.relation_tuples.v1alpha2.ErrorResponse.displayName = 'proto.ory.keto.relation_tuples.v1alpha2.ErrorResponse';
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
proto.ory.keto.relation_tuples.v1alpha2.ErrorResponse.Error = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.ory.keto.relation_tuples.v1alpha2.ErrorResponse.Error, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.ory.keto.relation_tuples.v1alpha2.ErrorResponse.Error.displayName = 'proto.ory.keto.relation_tuples.v1alpha2.ErrorResponse.Error';
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
proto.ory.keto.relation_tuples.v1alpha2.ErrorResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.ory.keto.relation_tuples.v1alpha2.ErrorResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.ory.keto.relation_tuples.v1alpha2.ErrorResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.ory.keto.relation_tuples.v1alpha2.ErrorResponse.toObject = function(includeInstance, msg) {
  var f, obj = {
error: (f = msg.getError()) && proto.ory.keto.relation_tuples.v1alpha2.ErrorResponse.Error.toObject(includeInstance, f)
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
 * @return {!proto.ory.keto.relation_tuples.v1alpha2.ErrorResponse}
 */
proto.ory.keto.relation_tuples.v1alpha2.ErrorResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.ory.keto.relation_tuples.v1alpha2.ErrorResponse;
  return proto.ory.keto.relation_tuples.v1alpha2.ErrorResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.ory.keto.relation_tuples.v1alpha2.ErrorResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.ory.keto.relation_tuples.v1alpha2.ErrorResponse}
 */
proto.ory.keto.relation_tuples.v1alpha2.ErrorResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.ory.keto.relation_tuples.v1alpha2.ErrorResponse.Error;
      reader.readMessage(value,proto.ory.keto.relation_tuples.v1alpha2.ErrorResponse.Error.deserializeBinaryFromReader);
      msg.setError(value);
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
proto.ory.keto.relation_tuples.v1alpha2.ErrorResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.ory.keto.relation_tuples.v1alpha2.ErrorResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.ory.keto.relation_tuples.v1alpha2.ErrorResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.ory.keto.relation_tuples.v1alpha2.ErrorResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getError();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      proto.ory.keto.relation_tuples.v1alpha2.ErrorResponse.Error.serializeBinaryToWriter
    );
  }
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
proto.ory.keto.relation_tuples.v1alpha2.ErrorResponse.Error.prototype.toObject = function(opt_includeInstance) {
  return proto.ory.keto.relation_tuples.v1alpha2.ErrorResponse.Error.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.ory.keto.relation_tuples.v1alpha2.ErrorResponse.Error} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.ory.keto.relation_tuples.v1alpha2.ErrorResponse.Error.toObject = function(includeInstance, msg) {
  var f, obj = {
code: jspb.Message.getFieldWithDefault(msg, 1, 0),
debug: jspb.Message.getFieldWithDefault(msg, 2, ""),
detailsMap: (f = msg.getDetailsMap()) ? f.toObject(includeInstance, undefined) : [],
id: jspb.Message.getFieldWithDefault(msg, 4, ""),
message: jspb.Message.getFieldWithDefault(msg, 5, ""),
reason: jspb.Message.getFieldWithDefault(msg, 6, ""),
request: jspb.Message.getFieldWithDefault(msg, 7, ""),
status: jspb.Message.getFieldWithDefault(msg, 8, "")
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
 * @return {!proto.ory.keto.relation_tuples.v1alpha2.ErrorResponse.Error}
 */
proto.ory.keto.relation_tuples.v1alpha2.ErrorResponse.Error.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.ory.keto.relation_tuples.v1alpha2.ErrorResponse.Error;
  return proto.ory.keto.relation_tuples.v1alpha2.ErrorResponse.Error.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.ory.keto.relation_tuples.v1alpha2.ErrorResponse.Error} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.ory.keto.relation_tuples.v1alpha2.ErrorResponse.Error}
 */
proto.ory.keto.relation_tuples.v1alpha2.ErrorResponse.Error.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setCode(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setDebug(value);
      break;
    case 3:
      var value = msg.getDetailsMap();
      reader.readMessage(value, function(message, reader) {
        jspb.Map.deserializeBinary(message, reader, jspb.BinaryReader.prototype.readString, jspb.BinaryReader.prototype.readString, null, "", "");
         });
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setId(value);
      break;
    case 5:
      var value = /** @type {string} */ (reader.readString());
      msg.setMessage(value);
      break;
    case 6:
      var value = /** @type {string} */ (reader.readString());
      msg.setReason(value);
      break;
    case 7:
      var value = /** @type {string} */ (reader.readString());
      msg.setRequest(value);
      break;
    case 8:
      var value = /** @type {string} */ (reader.readString());
      msg.setStatus(value);
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
proto.ory.keto.relation_tuples.v1alpha2.ErrorResponse.Error.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.ory.keto.relation_tuples.v1alpha2.ErrorResponse.Error.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.ory.keto.relation_tuples.v1alpha2.ErrorResponse.Error} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.ory.keto.relation_tuples.v1alpha2.ErrorResponse.Error.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getCode();
  if (f !== 0) {
    writer.writeInt64(
      1,
      f
    );
  }
  f = message.getDebug();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getDetailsMap(true);
  if (f && f.getLength() > 0) {
    f.serializeBinary(3, writer, jspb.BinaryWriter.prototype.writeString, jspb.BinaryWriter.prototype.writeString);
  }
  f = message.getId();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
  f = message.getMessage();
  if (f.length > 0) {
    writer.writeString(
      5,
      f
    );
  }
  f = message.getReason();
  if (f.length > 0) {
    writer.writeString(
      6,
      f
    );
  }
  f = message.getRequest();
  if (f.length > 0) {
    writer.writeString(
      7,
      f
    );
  }
  f = message.getStatus();
  if (f.length > 0) {
    writer.writeString(
      8,
      f
    );
  }
};


/**
 * optional int64 code = 1;
 * @return {number}
 */
proto.ory.keto.relation_tuples.v1alpha2.ErrorResponse.Error.prototype.getCode = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/**
 * @param {number} value
 * @return {!proto.ory.keto.relation_tuples.v1alpha2.ErrorResponse.Error} returns this
 */
proto.ory.keto.relation_tuples.v1alpha2.ErrorResponse.Error.prototype.setCode = function(value) {
  return jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * optional string debug = 2;
 * @return {string}
 */
proto.ory.keto.relation_tuples.v1alpha2.ErrorResponse.Error.prototype.getDebug = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.ory.keto.relation_tuples.v1alpha2.ErrorResponse.Error} returns this
 */
proto.ory.keto.relation_tuples.v1alpha2.ErrorResponse.Error.prototype.setDebug = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * map<string, string> details = 3;
 * @param {boolean=} opt_noLazyCreate Do not create the map if
 * empty, instead returning `undefined`
 * @return {!jspb.Map<string,string>}
 */
proto.ory.keto.relation_tuples.v1alpha2.ErrorResponse.Error.prototype.getDetailsMap = function(opt_noLazyCreate) {
  return /** @type {!jspb.Map<string,string>} */ (
      jspb.Message.getMapField(this, 3, opt_noLazyCreate,
      null));
};


/**
 * Clears values from the map. The map will be non-null.
 * @return {!proto.ory.keto.relation_tuples.v1alpha2.ErrorResponse.Error} returns this
 */
proto.ory.keto.relation_tuples.v1alpha2.ErrorResponse.Error.prototype.clearDetailsMap = function() {
  this.getDetailsMap().clear();
  return this;
};


/**
 * optional string id = 4;
 * @return {string}
 */
proto.ory.keto.relation_tuples.v1alpha2.ErrorResponse.Error.prototype.getId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/**
 * @param {string} value
 * @return {!proto.ory.keto.relation_tuples.v1alpha2.ErrorResponse.Error} returns this
 */
proto.ory.keto.relation_tuples.v1alpha2.ErrorResponse.Error.prototype.setId = function(value) {
  return jspb.Message.setProto3StringField(this, 4, value);
};


/**
 * optional string message = 5;
 * @return {string}
 */
proto.ory.keto.relation_tuples.v1alpha2.ErrorResponse.Error.prototype.getMessage = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 5, ""));
};


/**
 * @param {string} value
 * @return {!proto.ory.keto.relation_tuples.v1alpha2.ErrorResponse.Error} returns this
 */
proto.ory.keto.relation_tuples.v1alpha2.ErrorResponse.Error.prototype.setMessage = function(value) {
  return jspb.Message.setProto3StringField(this, 5, value);
};


/**
 * optional string reason = 6;
 * @return {string}
 */
proto.ory.keto.relation_tuples.v1alpha2.ErrorResponse.Error.prototype.getReason = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 6, ""));
};


/**
 * @param {string} value
 * @return {!proto.ory.keto.relation_tuples.v1alpha2.ErrorResponse.Error} returns this
 */
proto.ory.keto.relation_tuples.v1alpha2.ErrorResponse.Error.prototype.setReason = function(value) {
  return jspb.Message.setProto3StringField(this, 6, value);
};


/**
 * optional string request = 7;
 * @return {string}
 */
proto.ory.keto.relation_tuples.v1alpha2.ErrorResponse.Error.prototype.getRequest = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 7, ""));
};


/**
 * @param {string} value
 * @return {!proto.ory.keto.relation_tuples.v1alpha2.ErrorResponse.Error} returns this
 */
proto.ory.keto.relation_tuples.v1alpha2.ErrorResponse.Error.prototype.setRequest = function(value) {
  return jspb.Message.setProto3StringField(this, 7, value);
};


/**
 * optional string status = 8;
 * @return {string}
 */
proto.ory.keto.relation_tuples.v1alpha2.ErrorResponse.Error.prototype.getStatus = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 8, ""));
};


/**
 * @param {string} value
 * @return {!proto.ory.keto.relation_tuples.v1alpha2.ErrorResponse.Error} returns this
 */
proto.ory.keto.relation_tuples.v1alpha2.ErrorResponse.Error.prototype.setStatus = function(value) {
  return jspb.Message.setProto3StringField(this, 8, value);
};


/**
 * optional Error error = 1;
 * @return {?proto.ory.keto.relation_tuples.v1alpha2.ErrorResponse.Error}
 */
proto.ory.keto.relation_tuples.v1alpha2.ErrorResponse.prototype.getError = function() {
  return /** @type{?proto.ory.keto.relation_tuples.v1alpha2.ErrorResponse.Error} */ (
    jspb.Message.getWrapperField(this, proto.ory.keto.relation_tuples.v1alpha2.ErrorResponse.Error, 1));
};


/**
 * @param {?proto.ory.keto.relation_tuples.v1alpha2.ErrorResponse.Error|undefined} value
 * @return {!proto.ory.keto.relation_tuples.v1alpha2.ErrorResponse} returns this
*/
proto.ory.keto.relation_tuples.v1alpha2.ErrorResponse.prototype.setError = function(value) {
  return jspb.Message.setWrapperField(this, 1, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.ory.keto.relation_tuples.v1alpha2.ErrorResponse} returns this
 */
proto.ory.keto.relation_tuples.v1alpha2.ErrorResponse.prototype.clearError = function() {
  return this.setError(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.ory.keto.relation_tuples.v1alpha2.ErrorResponse.prototype.hasError = function() {
  return jspb.Message.getField(this, 1) != null;
};


goog.object.extend(exports, proto.ory.keto.relation_tuples.v1alpha2);
