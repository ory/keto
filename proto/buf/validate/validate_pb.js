// source: buf/validate/validate.proto
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

var google_protobuf_descriptor_pb = require('google-protobuf/google/protobuf/descriptor_pb.js');
goog.object.extend(proto, google_protobuf_descriptor_pb);
var google_protobuf_duration_pb = require('google-protobuf/google/protobuf/duration_pb.js');
goog.object.extend(proto, google_protobuf_duration_pb);
var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js');
goog.object.extend(proto, google_protobuf_timestamp_pb);
goog.exportSymbol('proto.buf.validate.AnyRules', null, global);
goog.exportSymbol('proto.buf.validate.BoolRules', null, global);
goog.exportSymbol('proto.buf.validate.BytesRules', null, global);
goog.exportSymbol('proto.buf.validate.BytesRules.WellKnownCase', null, global);
goog.exportSymbol('proto.buf.validate.Constraint', null, global);
goog.exportSymbol('proto.buf.validate.DoubleRules', null, global);
goog.exportSymbol('proto.buf.validate.DoubleRules.GreaterThanCase', null, global);
goog.exportSymbol('proto.buf.validate.DoubleRules.LessThanCase', null, global);
goog.exportSymbol('proto.buf.validate.DurationRules', null, global);
goog.exportSymbol('proto.buf.validate.DurationRules.GreaterThanCase', null, global);
goog.exportSymbol('proto.buf.validate.DurationRules.LessThanCase', null, global);
goog.exportSymbol('proto.buf.validate.EnumRules', null, global);
goog.exportSymbol('proto.buf.validate.FieldConstraints', null, global);
goog.exportSymbol('proto.buf.validate.FieldConstraints.TypeCase', null, global);
goog.exportSymbol('proto.buf.validate.FieldPath', null, global);
goog.exportSymbol('proto.buf.validate.FieldPathElement', null, global);
goog.exportSymbol('proto.buf.validate.FieldPathElement.SubscriptCase', null, global);
goog.exportSymbol('proto.buf.validate.Fixed32Rules', null, global);
goog.exportSymbol('proto.buf.validate.Fixed32Rules.GreaterThanCase', null, global);
goog.exportSymbol('proto.buf.validate.Fixed32Rules.LessThanCase', null, global);
goog.exportSymbol('proto.buf.validate.Fixed64Rules', null, global);
goog.exportSymbol('proto.buf.validate.Fixed64Rules.GreaterThanCase', null, global);
goog.exportSymbol('proto.buf.validate.Fixed64Rules.LessThanCase', null, global);
goog.exportSymbol('proto.buf.validate.FloatRules', null, global);
goog.exportSymbol('proto.buf.validate.FloatRules.GreaterThanCase', null, global);
goog.exportSymbol('proto.buf.validate.FloatRules.LessThanCase', null, global);
goog.exportSymbol('proto.buf.validate.Ignore', null, global);
goog.exportSymbol('proto.buf.validate.Int32Rules', null, global);
goog.exportSymbol('proto.buf.validate.Int32Rules.GreaterThanCase', null, global);
goog.exportSymbol('proto.buf.validate.Int32Rules.LessThanCase', null, global);
goog.exportSymbol('proto.buf.validate.Int64Rules', null, global);
goog.exportSymbol('proto.buf.validate.Int64Rules.GreaterThanCase', null, global);
goog.exportSymbol('proto.buf.validate.Int64Rules.LessThanCase', null, global);
goog.exportSymbol('proto.buf.validate.KnownRegex', null, global);
goog.exportSymbol('proto.buf.validate.MapRules', null, global);
goog.exportSymbol('proto.buf.validate.MessageConstraints', null, global);
goog.exportSymbol('proto.buf.validate.OneofConstraints', null, global);
goog.exportSymbol('proto.buf.validate.PredefinedConstraints', null, global);
goog.exportSymbol('proto.buf.validate.RepeatedRules', null, global);
goog.exportSymbol('proto.buf.validate.SFixed32Rules', null, global);
goog.exportSymbol('proto.buf.validate.SFixed32Rules.GreaterThanCase', null, global);
goog.exportSymbol('proto.buf.validate.SFixed32Rules.LessThanCase', null, global);
goog.exportSymbol('proto.buf.validate.SFixed64Rules', null, global);
goog.exportSymbol('proto.buf.validate.SFixed64Rules.GreaterThanCase', null, global);
goog.exportSymbol('proto.buf.validate.SFixed64Rules.LessThanCase', null, global);
goog.exportSymbol('proto.buf.validate.SInt32Rules', null, global);
goog.exportSymbol('proto.buf.validate.SInt32Rules.GreaterThanCase', null, global);
goog.exportSymbol('proto.buf.validate.SInt32Rules.LessThanCase', null, global);
goog.exportSymbol('proto.buf.validate.SInt64Rules', null, global);
goog.exportSymbol('proto.buf.validate.SInt64Rules.GreaterThanCase', null, global);
goog.exportSymbol('proto.buf.validate.SInt64Rules.LessThanCase', null, global);
goog.exportSymbol('proto.buf.validate.StringRules', null, global);
goog.exportSymbol('proto.buf.validate.StringRules.WellKnownCase', null, global);
goog.exportSymbol('proto.buf.validate.TimestampRules', null, global);
goog.exportSymbol('proto.buf.validate.TimestampRules.GreaterThanCase', null, global);
goog.exportSymbol('proto.buf.validate.TimestampRules.LessThanCase', null, global);
goog.exportSymbol('proto.buf.validate.UInt32Rules', null, global);
goog.exportSymbol('proto.buf.validate.UInt32Rules.GreaterThanCase', null, global);
goog.exportSymbol('proto.buf.validate.UInt32Rules.LessThanCase', null, global);
goog.exportSymbol('proto.buf.validate.UInt64Rules', null, global);
goog.exportSymbol('proto.buf.validate.UInt64Rules.GreaterThanCase', null, global);
goog.exportSymbol('proto.buf.validate.UInt64Rules.LessThanCase', null, global);
goog.exportSymbol('proto.buf.validate.Violation', null, global);
goog.exportSymbol('proto.buf.validate.Violations', null, global);
goog.exportSymbol('proto.buf.validate.field', null, global);
goog.exportSymbol('proto.buf.validate.message', null, global);
goog.exportSymbol('proto.buf.validate.oneof', null, global);
goog.exportSymbol('proto.buf.validate.predefined', null, global);
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
proto.buf.validate.Constraint = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.buf.validate.Constraint, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.buf.validate.Constraint.displayName = 'proto.buf.validate.Constraint';
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
proto.buf.validate.MessageConstraints = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.buf.validate.MessageConstraints.repeatedFields_, null);
};
goog.inherits(proto.buf.validate.MessageConstraints, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.buf.validate.MessageConstraints.displayName = 'proto.buf.validate.MessageConstraints';
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
proto.buf.validate.OneofConstraints = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.buf.validate.OneofConstraints, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.buf.validate.OneofConstraints.displayName = 'proto.buf.validate.OneofConstraints';
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
proto.buf.validate.FieldConstraints = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.buf.validate.FieldConstraints.repeatedFields_, proto.buf.validate.FieldConstraints.oneofGroups_);
};
goog.inherits(proto.buf.validate.FieldConstraints, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.buf.validate.FieldConstraints.displayName = 'proto.buf.validate.FieldConstraints';
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
proto.buf.validate.PredefinedConstraints = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.buf.validate.PredefinedConstraints.repeatedFields_, null);
};
goog.inherits(proto.buf.validate.PredefinedConstraints, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.buf.validate.PredefinedConstraints.displayName = 'proto.buf.validate.PredefinedConstraints';
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
proto.buf.validate.FloatRules = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, 10, proto.buf.validate.FloatRules.repeatedFields_, proto.buf.validate.FloatRules.oneofGroups_);
};
goog.inherits(proto.buf.validate.FloatRules, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.buf.validate.FloatRules.displayName = 'proto.buf.validate.FloatRules';
}

/**
 * The extensions registered with this message class. This is a map of
 * extension field number to fieldInfo object.
 *
 * For example:
 *     { 123: {fieldIndex: 123, fieldName: {my_field_name: 0}, ctor: proto.example.MyMessage} }
 *
 * fieldName contains the JsCompiler renamed field name property so that it
 * works in OPTIMIZED mode.
 *
 * @type {!Object<number, jspb.ExtensionFieldInfo>}
 */
proto.buf.validate.FloatRules.extensions = {};


/**
 * The extensions registered with this message class. This is a map of
 * extension field number to fieldInfo object.
 *
 * For example:
 *     { 123: {fieldIndex: 123, fieldName: {my_field_name: 0}, ctor: proto.example.MyMessage} }
 *
 * fieldName contains the JsCompiler renamed field name property so that it
 * works in OPTIMIZED mode.
 *
 * @type {!Object<number, jspb.ExtensionFieldBinaryInfo>}
 */
proto.buf.validate.FloatRules.extensionsBinary = {};

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
proto.buf.validate.DoubleRules = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, 10, proto.buf.validate.DoubleRules.repeatedFields_, proto.buf.validate.DoubleRules.oneofGroups_);
};
goog.inherits(proto.buf.validate.DoubleRules, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.buf.validate.DoubleRules.displayName = 'proto.buf.validate.DoubleRules';
}

/**
 * The extensions registered with this message class. This is a map of
 * extension field number to fieldInfo object.
 *
 * For example:
 *     { 123: {fieldIndex: 123, fieldName: {my_field_name: 0}, ctor: proto.example.MyMessage} }
 *
 * fieldName contains the JsCompiler renamed field name property so that it
 * works in OPTIMIZED mode.
 *
 * @type {!Object<number, jspb.ExtensionFieldInfo>}
 */
proto.buf.validate.DoubleRules.extensions = {};


/**
 * The extensions registered with this message class. This is a map of
 * extension field number to fieldInfo object.
 *
 * For example:
 *     { 123: {fieldIndex: 123, fieldName: {my_field_name: 0}, ctor: proto.example.MyMessage} }
 *
 * fieldName contains the JsCompiler renamed field name property so that it
 * works in OPTIMIZED mode.
 *
 * @type {!Object<number, jspb.ExtensionFieldBinaryInfo>}
 */
proto.buf.validate.DoubleRules.extensionsBinary = {};

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
proto.buf.validate.Int32Rules = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, 9, proto.buf.validate.Int32Rules.repeatedFields_, proto.buf.validate.Int32Rules.oneofGroups_);
};
goog.inherits(proto.buf.validate.Int32Rules, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.buf.validate.Int32Rules.displayName = 'proto.buf.validate.Int32Rules';
}

/**
 * The extensions registered with this message class. This is a map of
 * extension field number to fieldInfo object.
 *
 * For example:
 *     { 123: {fieldIndex: 123, fieldName: {my_field_name: 0}, ctor: proto.example.MyMessage} }
 *
 * fieldName contains the JsCompiler renamed field name property so that it
 * works in OPTIMIZED mode.
 *
 * @type {!Object<number, jspb.ExtensionFieldInfo>}
 */
proto.buf.validate.Int32Rules.extensions = {};


/**
 * The extensions registered with this message class. This is a map of
 * extension field number to fieldInfo object.
 *
 * For example:
 *     { 123: {fieldIndex: 123, fieldName: {my_field_name: 0}, ctor: proto.example.MyMessage} }
 *
 * fieldName contains the JsCompiler renamed field name property so that it
 * works in OPTIMIZED mode.
 *
 * @type {!Object<number, jspb.ExtensionFieldBinaryInfo>}
 */
proto.buf.validate.Int32Rules.extensionsBinary = {};

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
proto.buf.validate.Int64Rules = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, 10, proto.buf.validate.Int64Rules.repeatedFields_, proto.buf.validate.Int64Rules.oneofGroups_);
};
goog.inherits(proto.buf.validate.Int64Rules, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.buf.validate.Int64Rules.displayName = 'proto.buf.validate.Int64Rules';
}

/**
 * The extensions registered with this message class. This is a map of
 * extension field number to fieldInfo object.
 *
 * For example:
 *     { 123: {fieldIndex: 123, fieldName: {my_field_name: 0}, ctor: proto.example.MyMessage} }
 *
 * fieldName contains the JsCompiler renamed field name property so that it
 * works in OPTIMIZED mode.
 *
 * @type {!Object<number, jspb.ExtensionFieldInfo>}
 */
proto.buf.validate.Int64Rules.extensions = {};


/**
 * The extensions registered with this message class. This is a map of
 * extension field number to fieldInfo object.
 *
 * For example:
 *     { 123: {fieldIndex: 123, fieldName: {my_field_name: 0}, ctor: proto.example.MyMessage} }
 *
 * fieldName contains the JsCompiler renamed field name property so that it
 * works in OPTIMIZED mode.
 *
 * @type {!Object<number, jspb.ExtensionFieldBinaryInfo>}
 */
proto.buf.validate.Int64Rules.extensionsBinary = {};

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
proto.buf.validate.UInt32Rules = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, 9, proto.buf.validate.UInt32Rules.repeatedFields_, proto.buf.validate.UInt32Rules.oneofGroups_);
};
goog.inherits(proto.buf.validate.UInt32Rules, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.buf.validate.UInt32Rules.displayName = 'proto.buf.validate.UInt32Rules';
}

/**
 * The extensions registered with this message class. This is a map of
 * extension field number to fieldInfo object.
 *
 * For example:
 *     { 123: {fieldIndex: 123, fieldName: {my_field_name: 0}, ctor: proto.example.MyMessage} }
 *
 * fieldName contains the JsCompiler renamed field name property so that it
 * works in OPTIMIZED mode.
 *
 * @type {!Object<number, jspb.ExtensionFieldInfo>}
 */
proto.buf.validate.UInt32Rules.extensions = {};


/**
 * The extensions registered with this message class. This is a map of
 * extension field number to fieldInfo object.
 *
 * For example:
 *     { 123: {fieldIndex: 123, fieldName: {my_field_name: 0}, ctor: proto.example.MyMessage} }
 *
 * fieldName contains the JsCompiler renamed field name property so that it
 * works in OPTIMIZED mode.
 *
 * @type {!Object<number, jspb.ExtensionFieldBinaryInfo>}
 */
proto.buf.validate.UInt32Rules.extensionsBinary = {};

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
proto.buf.validate.UInt64Rules = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, 9, proto.buf.validate.UInt64Rules.repeatedFields_, proto.buf.validate.UInt64Rules.oneofGroups_);
};
goog.inherits(proto.buf.validate.UInt64Rules, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.buf.validate.UInt64Rules.displayName = 'proto.buf.validate.UInt64Rules';
}

/**
 * The extensions registered with this message class. This is a map of
 * extension field number to fieldInfo object.
 *
 * For example:
 *     { 123: {fieldIndex: 123, fieldName: {my_field_name: 0}, ctor: proto.example.MyMessage} }
 *
 * fieldName contains the JsCompiler renamed field name property so that it
 * works in OPTIMIZED mode.
 *
 * @type {!Object<number, jspb.ExtensionFieldInfo>}
 */
proto.buf.validate.UInt64Rules.extensions = {};


/**
 * The extensions registered with this message class. This is a map of
 * extension field number to fieldInfo object.
 *
 * For example:
 *     { 123: {fieldIndex: 123, fieldName: {my_field_name: 0}, ctor: proto.example.MyMessage} }
 *
 * fieldName contains the JsCompiler renamed field name property so that it
 * works in OPTIMIZED mode.
 *
 * @type {!Object<number, jspb.ExtensionFieldBinaryInfo>}
 */
proto.buf.validate.UInt64Rules.extensionsBinary = {};

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
proto.buf.validate.SInt32Rules = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, 9, proto.buf.validate.SInt32Rules.repeatedFields_, proto.buf.validate.SInt32Rules.oneofGroups_);
};
goog.inherits(proto.buf.validate.SInt32Rules, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.buf.validate.SInt32Rules.displayName = 'proto.buf.validate.SInt32Rules';
}

/**
 * The extensions registered with this message class. This is a map of
 * extension field number to fieldInfo object.
 *
 * For example:
 *     { 123: {fieldIndex: 123, fieldName: {my_field_name: 0}, ctor: proto.example.MyMessage} }
 *
 * fieldName contains the JsCompiler renamed field name property so that it
 * works in OPTIMIZED mode.
 *
 * @type {!Object<number, jspb.ExtensionFieldInfo>}
 */
proto.buf.validate.SInt32Rules.extensions = {};


/**
 * The extensions registered with this message class. This is a map of
 * extension field number to fieldInfo object.
 *
 * For example:
 *     { 123: {fieldIndex: 123, fieldName: {my_field_name: 0}, ctor: proto.example.MyMessage} }
 *
 * fieldName contains the JsCompiler renamed field name property so that it
 * works in OPTIMIZED mode.
 *
 * @type {!Object<number, jspb.ExtensionFieldBinaryInfo>}
 */
proto.buf.validate.SInt32Rules.extensionsBinary = {};

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
proto.buf.validate.SInt64Rules = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, 9, proto.buf.validate.SInt64Rules.repeatedFields_, proto.buf.validate.SInt64Rules.oneofGroups_);
};
goog.inherits(proto.buf.validate.SInt64Rules, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.buf.validate.SInt64Rules.displayName = 'proto.buf.validate.SInt64Rules';
}

/**
 * The extensions registered with this message class. This is a map of
 * extension field number to fieldInfo object.
 *
 * For example:
 *     { 123: {fieldIndex: 123, fieldName: {my_field_name: 0}, ctor: proto.example.MyMessage} }
 *
 * fieldName contains the JsCompiler renamed field name property so that it
 * works in OPTIMIZED mode.
 *
 * @type {!Object<number, jspb.ExtensionFieldInfo>}
 */
proto.buf.validate.SInt64Rules.extensions = {};


/**
 * The extensions registered with this message class. This is a map of
 * extension field number to fieldInfo object.
 *
 * For example:
 *     { 123: {fieldIndex: 123, fieldName: {my_field_name: 0}, ctor: proto.example.MyMessage} }
 *
 * fieldName contains the JsCompiler renamed field name property so that it
 * works in OPTIMIZED mode.
 *
 * @type {!Object<number, jspb.ExtensionFieldBinaryInfo>}
 */
proto.buf.validate.SInt64Rules.extensionsBinary = {};

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
proto.buf.validate.Fixed32Rules = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, 9, proto.buf.validate.Fixed32Rules.repeatedFields_, proto.buf.validate.Fixed32Rules.oneofGroups_);
};
goog.inherits(proto.buf.validate.Fixed32Rules, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.buf.validate.Fixed32Rules.displayName = 'proto.buf.validate.Fixed32Rules';
}

/**
 * The extensions registered with this message class. This is a map of
 * extension field number to fieldInfo object.
 *
 * For example:
 *     { 123: {fieldIndex: 123, fieldName: {my_field_name: 0}, ctor: proto.example.MyMessage} }
 *
 * fieldName contains the JsCompiler renamed field name property so that it
 * works in OPTIMIZED mode.
 *
 * @type {!Object<number, jspb.ExtensionFieldInfo>}
 */
proto.buf.validate.Fixed32Rules.extensions = {};


/**
 * The extensions registered with this message class. This is a map of
 * extension field number to fieldInfo object.
 *
 * For example:
 *     { 123: {fieldIndex: 123, fieldName: {my_field_name: 0}, ctor: proto.example.MyMessage} }
 *
 * fieldName contains the JsCompiler renamed field name property so that it
 * works in OPTIMIZED mode.
 *
 * @type {!Object<number, jspb.ExtensionFieldBinaryInfo>}
 */
proto.buf.validate.Fixed32Rules.extensionsBinary = {};

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
proto.buf.validate.Fixed64Rules = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, 9, proto.buf.validate.Fixed64Rules.repeatedFields_, proto.buf.validate.Fixed64Rules.oneofGroups_);
};
goog.inherits(proto.buf.validate.Fixed64Rules, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.buf.validate.Fixed64Rules.displayName = 'proto.buf.validate.Fixed64Rules';
}

/**
 * The extensions registered with this message class. This is a map of
 * extension field number to fieldInfo object.
 *
 * For example:
 *     { 123: {fieldIndex: 123, fieldName: {my_field_name: 0}, ctor: proto.example.MyMessage} }
 *
 * fieldName contains the JsCompiler renamed field name property so that it
 * works in OPTIMIZED mode.
 *
 * @type {!Object<number, jspb.ExtensionFieldInfo>}
 */
proto.buf.validate.Fixed64Rules.extensions = {};


/**
 * The extensions registered with this message class. This is a map of
 * extension field number to fieldInfo object.
 *
 * For example:
 *     { 123: {fieldIndex: 123, fieldName: {my_field_name: 0}, ctor: proto.example.MyMessage} }
 *
 * fieldName contains the JsCompiler renamed field name property so that it
 * works in OPTIMIZED mode.
 *
 * @type {!Object<number, jspb.ExtensionFieldBinaryInfo>}
 */
proto.buf.validate.Fixed64Rules.extensionsBinary = {};

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
proto.buf.validate.SFixed32Rules = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, 9, proto.buf.validate.SFixed32Rules.repeatedFields_, proto.buf.validate.SFixed32Rules.oneofGroups_);
};
goog.inherits(proto.buf.validate.SFixed32Rules, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.buf.validate.SFixed32Rules.displayName = 'proto.buf.validate.SFixed32Rules';
}

/**
 * The extensions registered with this message class. This is a map of
 * extension field number to fieldInfo object.
 *
 * For example:
 *     { 123: {fieldIndex: 123, fieldName: {my_field_name: 0}, ctor: proto.example.MyMessage} }
 *
 * fieldName contains the JsCompiler renamed field name property so that it
 * works in OPTIMIZED mode.
 *
 * @type {!Object<number, jspb.ExtensionFieldInfo>}
 */
proto.buf.validate.SFixed32Rules.extensions = {};


/**
 * The extensions registered with this message class. This is a map of
 * extension field number to fieldInfo object.
 *
 * For example:
 *     { 123: {fieldIndex: 123, fieldName: {my_field_name: 0}, ctor: proto.example.MyMessage} }
 *
 * fieldName contains the JsCompiler renamed field name property so that it
 * works in OPTIMIZED mode.
 *
 * @type {!Object<number, jspb.ExtensionFieldBinaryInfo>}
 */
proto.buf.validate.SFixed32Rules.extensionsBinary = {};

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
proto.buf.validate.SFixed64Rules = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, 9, proto.buf.validate.SFixed64Rules.repeatedFields_, proto.buf.validate.SFixed64Rules.oneofGroups_);
};
goog.inherits(proto.buf.validate.SFixed64Rules, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.buf.validate.SFixed64Rules.displayName = 'proto.buf.validate.SFixed64Rules';
}

/**
 * The extensions registered with this message class. This is a map of
 * extension field number to fieldInfo object.
 *
 * For example:
 *     { 123: {fieldIndex: 123, fieldName: {my_field_name: 0}, ctor: proto.example.MyMessage} }
 *
 * fieldName contains the JsCompiler renamed field name property so that it
 * works in OPTIMIZED mode.
 *
 * @type {!Object<number, jspb.ExtensionFieldInfo>}
 */
proto.buf.validate.SFixed64Rules.extensions = {};


/**
 * The extensions registered with this message class. This is a map of
 * extension field number to fieldInfo object.
 *
 * For example:
 *     { 123: {fieldIndex: 123, fieldName: {my_field_name: 0}, ctor: proto.example.MyMessage} }
 *
 * fieldName contains the JsCompiler renamed field name property so that it
 * works in OPTIMIZED mode.
 *
 * @type {!Object<number, jspb.ExtensionFieldBinaryInfo>}
 */
proto.buf.validate.SFixed64Rules.extensionsBinary = {};

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
proto.buf.validate.BoolRules = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, 3, proto.buf.validate.BoolRules.repeatedFields_, null);
};
goog.inherits(proto.buf.validate.BoolRules, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.buf.validate.BoolRules.displayName = 'proto.buf.validate.BoolRules';
}

/**
 * The extensions registered with this message class. This is a map of
 * extension field number to fieldInfo object.
 *
 * For example:
 *     { 123: {fieldIndex: 123, fieldName: {my_field_name: 0}, ctor: proto.example.MyMessage} }
 *
 * fieldName contains the JsCompiler renamed field name property so that it
 * works in OPTIMIZED mode.
 *
 * @type {!Object<number, jspb.ExtensionFieldInfo>}
 */
proto.buf.validate.BoolRules.extensions = {};


/**
 * The extensions registered with this message class. This is a map of
 * extension field number to fieldInfo object.
 *
 * For example:
 *     { 123: {fieldIndex: 123, fieldName: {my_field_name: 0}, ctor: proto.example.MyMessage} }
 *
 * fieldName contains the JsCompiler renamed field name property so that it
 * works in OPTIMIZED mode.
 *
 * @type {!Object<number, jspb.ExtensionFieldBinaryInfo>}
 */
proto.buf.validate.BoolRules.extensionsBinary = {};

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
proto.buf.validate.StringRules = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, 35, proto.buf.validate.StringRules.repeatedFields_, proto.buf.validate.StringRules.oneofGroups_);
};
goog.inherits(proto.buf.validate.StringRules, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.buf.validate.StringRules.displayName = 'proto.buf.validate.StringRules';
}

/**
 * The extensions registered with this message class. This is a map of
 * extension field number to fieldInfo object.
 *
 * For example:
 *     { 123: {fieldIndex: 123, fieldName: {my_field_name: 0}, ctor: proto.example.MyMessage} }
 *
 * fieldName contains the JsCompiler renamed field name property so that it
 * works in OPTIMIZED mode.
 *
 * @type {!Object<number, jspb.ExtensionFieldInfo>}
 */
proto.buf.validate.StringRules.extensions = {};


/**
 * The extensions registered with this message class. This is a map of
 * extension field number to fieldInfo object.
 *
 * For example:
 *     { 123: {fieldIndex: 123, fieldName: {my_field_name: 0}, ctor: proto.example.MyMessage} }
 *
 * fieldName contains the JsCompiler renamed field name property so that it
 * works in OPTIMIZED mode.
 *
 * @type {!Object<number, jspb.ExtensionFieldBinaryInfo>}
 */
proto.buf.validate.StringRules.extensionsBinary = {};

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
proto.buf.validate.BytesRules = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, 15, proto.buf.validate.BytesRules.repeatedFields_, proto.buf.validate.BytesRules.oneofGroups_);
};
goog.inherits(proto.buf.validate.BytesRules, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.buf.validate.BytesRules.displayName = 'proto.buf.validate.BytesRules';
}

/**
 * The extensions registered with this message class. This is a map of
 * extension field number to fieldInfo object.
 *
 * For example:
 *     { 123: {fieldIndex: 123, fieldName: {my_field_name: 0}, ctor: proto.example.MyMessage} }
 *
 * fieldName contains the JsCompiler renamed field name property so that it
 * works in OPTIMIZED mode.
 *
 * @type {!Object<number, jspb.ExtensionFieldInfo>}
 */
proto.buf.validate.BytesRules.extensions = {};


/**
 * The extensions registered with this message class. This is a map of
 * extension field number to fieldInfo object.
 *
 * For example:
 *     { 123: {fieldIndex: 123, fieldName: {my_field_name: 0}, ctor: proto.example.MyMessage} }
 *
 * fieldName contains the JsCompiler renamed field name property so that it
 * works in OPTIMIZED mode.
 *
 * @type {!Object<number, jspb.ExtensionFieldBinaryInfo>}
 */
proto.buf.validate.BytesRules.extensionsBinary = {};

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
proto.buf.validate.EnumRules = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, 6, proto.buf.validate.EnumRules.repeatedFields_, null);
};
goog.inherits(proto.buf.validate.EnumRules, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.buf.validate.EnumRules.displayName = 'proto.buf.validate.EnumRules';
}

/**
 * The extensions registered with this message class. This is a map of
 * extension field number to fieldInfo object.
 *
 * For example:
 *     { 123: {fieldIndex: 123, fieldName: {my_field_name: 0}, ctor: proto.example.MyMessage} }
 *
 * fieldName contains the JsCompiler renamed field name property so that it
 * works in OPTIMIZED mode.
 *
 * @type {!Object<number, jspb.ExtensionFieldInfo>}
 */
proto.buf.validate.EnumRules.extensions = {};


/**
 * The extensions registered with this message class. This is a map of
 * extension field number to fieldInfo object.
 *
 * For example:
 *     { 123: {fieldIndex: 123, fieldName: {my_field_name: 0}, ctor: proto.example.MyMessage} }
 *
 * fieldName contains the JsCompiler renamed field name property so that it
 * works in OPTIMIZED mode.
 *
 * @type {!Object<number, jspb.ExtensionFieldBinaryInfo>}
 */
proto.buf.validate.EnumRules.extensionsBinary = {};

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
proto.buf.validate.RepeatedRules = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, 5, null, null);
};
goog.inherits(proto.buf.validate.RepeatedRules, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.buf.validate.RepeatedRules.displayName = 'proto.buf.validate.RepeatedRules';
}

/**
 * The extensions registered with this message class. This is a map of
 * extension field number to fieldInfo object.
 *
 * For example:
 *     { 123: {fieldIndex: 123, fieldName: {my_field_name: 0}, ctor: proto.example.MyMessage} }
 *
 * fieldName contains the JsCompiler renamed field name property so that it
 * works in OPTIMIZED mode.
 *
 * @type {!Object<number, jspb.ExtensionFieldInfo>}
 */
proto.buf.validate.RepeatedRules.extensions = {};


/**
 * The extensions registered with this message class. This is a map of
 * extension field number to fieldInfo object.
 *
 * For example:
 *     { 123: {fieldIndex: 123, fieldName: {my_field_name: 0}, ctor: proto.example.MyMessage} }
 *
 * fieldName contains the JsCompiler renamed field name property so that it
 * works in OPTIMIZED mode.
 *
 * @type {!Object<number, jspb.ExtensionFieldBinaryInfo>}
 */
proto.buf.validate.RepeatedRules.extensionsBinary = {};

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
proto.buf.validate.MapRules = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, 6, null, null);
};
goog.inherits(proto.buf.validate.MapRules, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.buf.validate.MapRules.displayName = 'proto.buf.validate.MapRules';
}

/**
 * The extensions registered with this message class. This is a map of
 * extension field number to fieldInfo object.
 *
 * For example:
 *     { 123: {fieldIndex: 123, fieldName: {my_field_name: 0}, ctor: proto.example.MyMessage} }
 *
 * fieldName contains the JsCompiler renamed field name property so that it
 * works in OPTIMIZED mode.
 *
 * @type {!Object<number, jspb.ExtensionFieldInfo>}
 */
proto.buf.validate.MapRules.extensions = {};


/**
 * The extensions registered with this message class. This is a map of
 * extension field number to fieldInfo object.
 *
 * For example:
 *     { 123: {fieldIndex: 123, fieldName: {my_field_name: 0}, ctor: proto.example.MyMessage} }
 *
 * fieldName contains the JsCompiler renamed field name property so that it
 * works in OPTIMIZED mode.
 *
 * @type {!Object<number, jspb.ExtensionFieldBinaryInfo>}
 */
proto.buf.validate.MapRules.extensionsBinary = {};

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
proto.buf.validate.AnyRules = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.buf.validate.AnyRules.repeatedFields_, null);
};
goog.inherits(proto.buf.validate.AnyRules, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.buf.validate.AnyRules.displayName = 'proto.buf.validate.AnyRules';
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
proto.buf.validate.DurationRules = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, 10, proto.buf.validate.DurationRules.repeatedFields_, proto.buf.validate.DurationRules.oneofGroups_);
};
goog.inherits(proto.buf.validate.DurationRules, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.buf.validate.DurationRules.displayName = 'proto.buf.validate.DurationRules';
}

/**
 * The extensions registered with this message class. This is a map of
 * extension field number to fieldInfo object.
 *
 * For example:
 *     { 123: {fieldIndex: 123, fieldName: {my_field_name: 0}, ctor: proto.example.MyMessage} }
 *
 * fieldName contains the JsCompiler renamed field name property so that it
 * works in OPTIMIZED mode.
 *
 * @type {!Object<number, jspb.ExtensionFieldInfo>}
 */
proto.buf.validate.DurationRules.extensions = {};


/**
 * The extensions registered with this message class. This is a map of
 * extension field number to fieldInfo object.
 *
 * For example:
 *     { 123: {fieldIndex: 123, fieldName: {my_field_name: 0}, ctor: proto.example.MyMessage} }
 *
 * fieldName contains the JsCompiler renamed field name property so that it
 * works in OPTIMIZED mode.
 *
 * @type {!Object<number, jspb.ExtensionFieldBinaryInfo>}
 */
proto.buf.validate.DurationRules.extensionsBinary = {};

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
proto.buf.validate.TimestampRules = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, 11, proto.buf.validate.TimestampRules.repeatedFields_, proto.buf.validate.TimestampRules.oneofGroups_);
};
goog.inherits(proto.buf.validate.TimestampRules, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.buf.validate.TimestampRules.displayName = 'proto.buf.validate.TimestampRules';
}

/**
 * The extensions registered with this message class. This is a map of
 * extension field number to fieldInfo object.
 *
 * For example:
 *     { 123: {fieldIndex: 123, fieldName: {my_field_name: 0}, ctor: proto.example.MyMessage} }
 *
 * fieldName contains the JsCompiler renamed field name property so that it
 * works in OPTIMIZED mode.
 *
 * @type {!Object<number, jspb.ExtensionFieldInfo>}
 */
proto.buf.validate.TimestampRules.extensions = {};


/**
 * The extensions registered with this message class. This is a map of
 * extension field number to fieldInfo object.
 *
 * For example:
 *     { 123: {fieldIndex: 123, fieldName: {my_field_name: 0}, ctor: proto.example.MyMessage} }
 *
 * fieldName contains the JsCompiler renamed field name property so that it
 * works in OPTIMIZED mode.
 *
 * @type {!Object<number, jspb.ExtensionFieldBinaryInfo>}
 */
proto.buf.validate.TimestampRules.extensionsBinary = {};

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
proto.buf.validate.Violations = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.buf.validate.Violations.repeatedFields_, null);
};
goog.inherits(proto.buf.validate.Violations, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.buf.validate.Violations.displayName = 'proto.buf.validate.Violations';
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
proto.buf.validate.Violation = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.buf.validate.Violation, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.buf.validate.Violation.displayName = 'proto.buf.validate.Violation';
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
proto.buf.validate.FieldPath = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.buf.validate.FieldPath.repeatedFields_, null);
};
goog.inherits(proto.buf.validate.FieldPath, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.buf.validate.FieldPath.displayName = 'proto.buf.validate.FieldPath';
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
proto.buf.validate.FieldPathElement = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, proto.buf.validate.FieldPathElement.oneofGroups_);
};
goog.inherits(proto.buf.validate.FieldPathElement, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.buf.validate.FieldPathElement.displayName = 'proto.buf.validate.FieldPathElement';
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
proto.buf.validate.Constraint.prototype.toObject = function(opt_includeInstance) {
  return proto.buf.validate.Constraint.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.buf.validate.Constraint} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.buf.validate.Constraint.toObject = function(includeInstance, msg) {
  var f, obj = {
id: (f = jspb.Message.getField(msg, 1)) == null ? undefined : f,
message: (f = jspb.Message.getField(msg, 2)) == null ? undefined : f,
expression: (f = jspb.Message.getField(msg, 3)) == null ? undefined : f
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
 * @return {!proto.buf.validate.Constraint}
 */
proto.buf.validate.Constraint.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.buf.validate.Constraint;
  return proto.buf.validate.Constraint.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.buf.validate.Constraint} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.buf.validate.Constraint}
 */
proto.buf.validate.Constraint.deserializeBinaryFromReader = function(msg, reader) {
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
      var value = /** @type {string} */ (reader.readString());
      msg.setMessage(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setExpression(value);
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
proto.buf.validate.Constraint.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.buf.validate.Constraint.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.buf.validate.Constraint} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.buf.validate.Constraint.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = /** @type {string} */ (jspb.Message.getField(message, 1));
  if (f != null) {
    writer.writeString(
      1,
      f
    );
  }
  f = /** @type {string} */ (jspb.Message.getField(message, 2));
  if (f != null) {
    writer.writeString(
      2,
      f
    );
  }
  f = /** @type {string} */ (jspb.Message.getField(message, 3));
  if (f != null) {
    writer.writeString(
      3,
      f
    );
  }
};


/**
 * optional string id = 1;
 * @return {string}
 */
proto.buf.validate.Constraint.prototype.getId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.buf.validate.Constraint} returns this
 */
proto.buf.validate.Constraint.prototype.setId = function(value) {
  return jspb.Message.setField(this, 1, value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.Constraint} returns this
 */
proto.buf.validate.Constraint.prototype.clearId = function() {
  return jspb.Message.setField(this, 1, undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.Constraint.prototype.hasId = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional string message = 2;
 * @return {string}
 */
proto.buf.validate.Constraint.prototype.getMessage = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.buf.validate.Constraint} returns this
 */
proto.buf.validate.Constraint.prototype.setMessage = function(value) {
  return jspb.Message.setField(this, 2, value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.Constraint} returns this
 */
proto.buf.validate.Constraint.prototype.clearMessage = function() {
  return jspb.Message.setField(this, 2, undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.Constraint.prototype.hasMessage = function() {
  return jspb.Message.getField(this, 2) != null;
};


/**
 * optional string expression = 3;
 * @return {string}
 */
proto.buf.validate.Constraint.prototype.getExpression = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/**
 * @param {string} value
 * @return {!proto.buf.validate.Constraint} returns this
 */
proto.buf.validate.Constraint.prototype.setExpression = function(value) {
  return jspb.Message.setField(this, 3, value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.Constraint} returns this
 */
proto.buf.validate.Constraint.prototype.clearExpression = function() {
  return jspb.Message.setField(this, 3, undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.Constraint.prototype.hasExpression = function() {
  return jspb.Message.getField(this, 3) != null;
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.buf.validate.MessageConstraints.repeatedFields_ = [3];



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
proto.buf.validate.MessageConstraints.prototype.toObject = function(opt_includeInstance) {
  return proto.buf.validate.MessageConstraints.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.buf.validate.MessageConstraints} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.buf.validate.MessageConstraints.toObject = function(includeInstance, msg) {
  var f, obj = {
disabled: (f = jspb.Message.getBooleanField(msg, 1)) == null ? undefined : f,
celList: jspb.Message.toObjectList(msg.getCelList(),
    proto.buf.validate.Constraint.toObject, includeInstance)
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
 * @return {!proto.buf.validate.MessageConstraints}
 */
proto.buf.validate.MessageConstraints.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.buf.validate.MessageConstraints;
  return proto.buf.validate.MessageConstraints.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.buf.validate.MessageConstraints} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.buf.validate.MessageConstraints}
 */
proto.buf.validate.MessageConstraints.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setDisabled(value);
      break;
    case 3:
      var value = new proto.buf.validate.Constraint;
      reader.readMessage(value,proto.buf.validate.Constraint.deserializeBinaryFromReader);
      msg.addCel(value);
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
proto.buf.validate.MessageConstraints.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.buf.validate.MessageConstraints.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.buf.validate.MessageConstraints} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.buf.validate.MessageConstraints.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = /** @type {boolean} */ (jspb.Message.getField(message, 1));
  if (f != null) {
    writer.writeBool(
      1,
      f
    );
  }
  f = message.getCelList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      3,
      f,
      proto.buf.validate.Constraint.serializeBinaryToWriter
    );
  }
};


/**
 * optional bool disabled = 1;
 * @return {boolean}
 */
proto.buf.validate.MessageConstraints.prototype.getDisabled = function() {
  return /** @type {boolean} */ (jspb.Message.getBooleanFieldWithDefault(this, 1, false));
};


/**
 * @param {boolean} value
 * @return {!proto.buf.validate.MessageConstraints} returns this
 */
proto.buf.validate.MessageConstraints.prototype.setDisabled = function(value) {
  return jspb.Message.setField(this, 1, value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.MessageConstraints} returns this
 */
proto.buf.validate.MessageConstraints.prototype.clearDisabled = function() {
  return jspb.Message.setField(this, 1, undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.MessageConstraints.prototype.hasDisabled = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * repeated Constraint cel = 3;
 * @return {!Array<!proto.buf.validate.Constraint>}
 */
proto.buf.validate.MessageConstraints.prototype.getCelList = function() {
  return /** @type{!Array<!proto.buf.validate.Constraint>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.buf.validate.Constraint, 3));
};


/**
 * @param {!Array<!proto.buf.validate.Constraint>} value
 * @return {!proto.buf.validate.MessageConstraints} returns this
*/
proto.buf.validate.MessageConstraints.prototype.setCelList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 3, value);
};


/**
 * @param {!proto.buf.validate.Constraint=} opt_value
 * @param {number=} opt_index
 * @return {!proto.buf.validate.Constraint}
 */
proto.buf.validate.MessageConstraints.prototype.addCel = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 3, opt_value, proto.buf.validate.Constraint, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.buf.validate.MessageConstraints} returns this
 */
proto.buf.validate.MessageConstraints.prototype.clearCelList = function() {
  return this.setCelList([]);
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
proto.buf.validate.OneofConstraints.prototype.toObject = function(opt_includeInstance) {
  return proto.buf.validate.OneofConstraints.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.buf.validate.OneofConstraints} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.buf.validate.OneofConstraints.toObject = function(includeInstance, msg) {
  var f, obj = {
required: (f = jspb.Message.getBooleanField(msg, 1)) == null ? undefined : f
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
 * @return {!proto.buf.validate.OneofConstraints}
 */
proto.buf.validate.OneofConstraints.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.buf.validate.OneofConstraints;
  return proto.buf.validate.OneofConstraints.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.buf.validate.OneofConstraints} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.buf.validate.OneofConstraints}
 */
proto.buf.validate.OneofConstraints.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setRequired(value);
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
proto.buf.validate.OneofConstraints.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.buf.validate.OneofConstraints.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.buf.validate.OneofConstraints} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.buf.validate.OneofConstraints.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = /** @type {boolean} */ (jspb.Message.getField(message, 1));
  if (f != null) {
    writer.writeBool(
      1,
      f
    );
  }
};


/**
 * optional bool required = 1;
 * @return {boolean}
 */
proto.buf.validate.OneofConstraints.prototype.getRequired = function() {
  return /** @type {boolean} */ (jspb.Message.getBooleanFieldWithDefault(this, 1, false));
};


/**
 * @param {boolean} value
 * @return {!proto.buf.validate.OneofConstraints} returns this
 */
proto.buf.validate.OneofConstraints.prototype.setRequired = function(value) {
  return jspb.Message.setField(this, 1, value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.OneofConstraints} returns this
 */
proto.buf.validate.OneofConstraints.prototype.clearRequired = function() {
  return jspb.Message.setField(this, 1, undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.OneofConstraints.prototype.hasRequired = function() {
  return jspb.Message.getField(this, 1) != null;
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.buf.validate.FieldConstraints.repeatedFields_ = [23];

/**
 * Oneof group definitions for this message. Each group defines the field
 * numbers belonging to that group. When of these fields' value is set, all
 * other fields in the group are cleared. During deserialization, if multiple
 * fields are encountered for a group, only the last value seen will be kept.
 * @private {!Array<!Array<number>>}
 * @const
 */
proto.buf.validate.FieldConstraints.oneofGroups_ = [[1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,18,19,20,21,22]];

/**
 * @enum {number}
 */
proto.buf.validate.FieldConstraints.TypeCase = {
  TYPE_NOT_SET: 0,
  FLOAT: 1,
  DOUBLE: 2,
  INT32: 3,
  INT64: 4,
  UINT32: 5,
  UINT64: 6,
  SINT32: 7,
  SINT64: 8,
  FIXED32: 9,
  FIXED64: 10,
  SFIXED32: 11,
  SFIXED64: 12,
  BOOL: 13,
  STRING: 14,
  BYTES: 15,
  ENUM: 16,
  REPEATED: 18,
  MAP: 19,
  ANY: 20,
  DURATION: 21,
  TIMESTAMP: 22
};

/**
 * @return {proto.buf.validate.FieldConstraints.TypeCase}
 */
proto.buf.validate.FieldConstraints.prototype.getTypeCase = function() {
  return /** @type {proto.buf.validate.FieldConstraints.TypeCase} */(jspb.Message.computeOneofCase(this, proto.buf.validate.FieldConstraints.oneofGroups_[0]));
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
proto.buf.validate.FieldConstraints.prototype.toObject = function(opt_includeInstance) {
  return proto.buf.validate.FieldConstraints.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.buf.validate.FieldConstraints} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.buf.validate.FieldConstraints.toObject = function(includeInstance, msg) {
  var f, obj = {
celList: jspb.Message.toObjectList(msg.getCelList(),
    proto.buf.validate.Constraint.toObject, includeInstance),
required: (f = jspb.Message.getBooleanField(msg, 25)) == null ? undefined : f,
ignore: (f = jspb.Message.getField(msg, 27)) == null ? undefined : f,
pb_float: (f = msg.getFloat()) && proto.buf.validate.FloatRules.toObject(includeInstance, f),
pb_double: (f = msg.getDouble()) && proto.buf.validate.DoubleRules.toObject(includeInstance, f),
int32: (f = msg.getInt32()) && proto.buf.validate.Int32Rules.toObject(includeInstance, f),
int64: (f = msg.getInt64()) && proto.buf.validate.Int64Rules.toObject(includeInstance, f),
uint32: (f = msg.getUint32()) && proto.buf.validate.UInt32Rules.toObject(includeInstance, f),
uint64: (f = msg.getUint64()) && proto.buf.validate.UInt64Rules.toObject(includeInstance, f),
sint32: (f = msg.getSint32()) && proto.buf.validate.SInt32Rules.toObject(includeInstance, f),
sint64: (f = msg.getSint64()) && proto.buf.validate.SInt64Rules.toObject(includeInstance, f),
fixed32: (f = msg.getFixed32()) && proto.buf.validate.Fixed32Rules.toObject(includeInstance, f),
fixed64: (f = msg.getFixed64()) && proto.buf.validate.Fixed64Rules.toObject(includeInstance, f),
sfixed32: (f = msg.getSfixed32()) && proto.buf.validate.SFixed32Rules.toObject(includeInstance, f),
sfixed64: (f = msg.getSfixed64()) && proto.buf.validate.SFixed64Rules.toObject(includeInstance, f),
bool: (f = msg.getBool()) && proto.buf.validate.BoolRules.toObject(includeInstance, f),
string: (f = msg.getString()) && proto.buf.validate.StringRules.toObject(includeInstance, f),
bytes: (f = msg.getBytes()) && proto.buf.validate.BytesRules.toObject(includeInstance, f),
pb_enum: (f = msg.getEnum()) && proto.buf.validate.EnumRules.toObject(includeInstance, f),
repeated: (f = msg.getRepeated()) && proto.buf.validate.RepeatedRules.toObject(includeInstance, f),
map: (f = msg.getMap()) && proto.buf.validate.MapRules.toObject(includeInstance, f),
any: (f = msg.getAny()) && proto.buf.validate.AnyRules.toObject(includeInstance, f),
duration: (f = msg.getDuration()) && proto.buf.validate.DurationRules.toObject(includeInstance, f),
timestamp: (f = msg.getTimestamp()) && proto.buf.validate.TimestampRules.toObject(includeInstance, f)
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
 * @return {!proto.buf.validate.FieldConstraints}
 */
proto.buf.validate.FieldConstraints.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.buf.validate.FieldConstraints;
  return proto.buf.validate.FieldConstraints.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.buf.validate.FieldConstraints} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.buf.validate.FieldConstraints}
 */
proto.buf.validate.FieldConstraints.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 23:
      var value = new proto.buf.validate.Constraint;
      reader.readMessage(value,proto.buf.validate.Constraint.deserializeBinaryFromReader);
      msg.addCel(value);
      break;
    case 25:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setRequired(value);
      break;
    case 27:
      var value = /** @type {!proto.buf.validate.Ignore} */ (reader.readEnum());
      msg.setIgnore(value);
      break;
    case 1:
      var value = new proto.buf.validate.FloatRules;
      reader.readMessage(value,proto.buf.validate.FloatRules.deserializeBinaryFromReader);
      msg.setFloat(value);
      break;
    case 2:
      var value = new proto.buf.validate.DoubleRules;
      reader.readMessage(value,proto.buf.validate.DoubleRules.deserializeBinaryFromReader);
      msg.setDouble(value);
      break;
    case 3:
      var value = new proto.buf.validate.Int32Rules;
      reader.readMessage(value,proto.buf.validate.Int32Rules.deserializeBinaryFromReader);
      msg.setInt32(value);
      break;
    case 4:
      var value = new proto.buf.validate.Int64Rules;
      reader.readMessage(value,proto.buf.validate.Int64Rules.deserializeBinaryFromReader);
      msg.setInt64(value);
      break;
    case 5:
      var value = new proto.buf.validate.UInt32Rules;
      reader.readMessage(value,proto.buf.validate.UInt32Rules.deserializeBinaryFromReader);
      msg.setUint32(value);
      break;
    case 6:
      var value = new proto.buf.validate.UInt64Rules;
      reader.readMessage(value,proto.buf.validate.UInt64Rules.deserializeBinaryFromReader);
      msg.setUint64(value);
      break;
    case 7:
      var value = new proto.buf.validate.SInt32Rules;
      reader.readMessage(value,proto.buf.validate.SInt32Rules.deserializeBinaryFromReader);
      msg.setSint32(value);
      break;
    case 8:
      var value = new proto.buf.validate.SInt64Rules;
      reader.readMessage(value,proto.buf.validate.SInt64Rules.deserializeBinaryFromReader);
      msg.setSint64(value);
      break;
    case 9:
      var value = new proto.buf.validate.Fixed32Rules;
      reader.readMessage(value,proto.buf.validate.Fixed32Rules.deserializeBinaryFromReader);
      msg.setFixed32(value);
      break;
    case 10:
      var value = new proto.buf.validate.Fixed64Rules;
      reader.readMessage(value,proto.buf.validate.Fixed64Rules.deserializeBinaryFromReader);
      msg.setFixed64(value);
      break;
    case 11:
      var value = new proto.buf.validate.SFixed32Rules;
      reader.readMessage(value,proto.buf.validate.SFixed32Rules.deserializeBinaryFromReader);
      msg.setSfixed32(value);
      break;
    case 12:
      var value = new proto.buf.validate.SFixed64Rules;
      reader.readMessage(value,proto.buf.validate.SFixed64Rules.deserializeBinaryFromReader);
      msg.setSfixed64(value);
      break;
    case 13:
      var value = new proto.buf.validate.BoolRules;
      reader.readMessage(value,proto.buf.validate.BoolRules.deserializeBinaryFromReader);
      msg.setBool(value);
      break;
    case 14:
      var value = new proto.buf.validate.StringRules;
      reader.readMessage(value,proto.buf.validate.StringRules.deserializeBinaryFromReader);
      msg.setString(value);
      break;
    case 15:
      var value = new proto.buf.validate.BytesRules;
      reader.readMessage(value,proto.buf.validate.BytesRules.deserializeBinaryFromReader);
      msg.setBytes(value);
      break;
    case 16:
      var value = new proto.buf.validate.EnumRules;
      reader.readMessage(value,proto.buf.validate.EnumRules.deserializeBinaryFromReader);
      msg.setEnum(value);
      break;
    case 18:
      var value = new proto.buf.validate.RepeatedRules;
      reader.readMessage(value,proto.buf.validate.RepeatedRules.deserializeBinaryFromReader);
      msg.setRepeated(value);
      break;
    case 19:
      var value = new proto.buf.validate.MapRules;
      reader.readMessage(value,proto.buf.validate.MapRules.deserializeBinaryFromReader);
      msg.setMap(value);
      break;
    case 20:
      var value = new proto.buf.validate.AnyRules;
      reader.readMessage(value,proto.buf.validate.AnyRules.deserializeBinaryFromReader);
      msg.setAny(value);
      break;
    case 21:
      var value = new proto.buf.validate.DurationRules;
      reader.readMessage(value,proto.buf.validate.DurationRules.deserializeBinaryFromReader);
      msg.setDuration(value);
      break;
    case 22:
      var value = new proto.buf.validate.TimestampRules;
      reader.readMessage(value,proto.buf.validate.TimestampRules.deserializeBinaryFromReader);
      msg.setTimestamp(value);
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
proto.buf.validate.FieldConstraints.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.buf.validate.FieldConstraints.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.buf.validate.FieldConstraints} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.buf.validate.FieldConstraints.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getCelList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      23,
      f,
      proto.buf.validate.Constraint.serializeBinaryToWriter
    );
  }
  f = /** @type {boolean} */ (jspb.Message.getField(message, 25));
  if (f != null) {
    writer.writeBool(
      25,
      f
    );
  }
  f = /** @type {!proto.buf.validate.Ignore} */ (jspb.Message.getField(message, 27));
  if (f != null) {
    writer.writeEnum(
      27,
      f
    );
  }
  f = message.getFloat();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      proto.buf.validate.FloatRules.serializeBinaryToWriter
    );
  }
  f = message.getDouble();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      proto.buf.validate.DoubleRules.serializeBinaryToWriter
    );
  }
  f = message.getInt32();
  if (f != null) {
    writer.writeMessage(
      3,
      f,
      proto.buf.validate.Int32Rules.serializeBinaryToWriter
    );
  }
  f = message.getInt64();
  if (f != null) {
    writer.writeMessage(
      4,
      f,
      proto.buf.validate.Int64Rules.serializeBinaryToWriter
    );
  }
  f = message.getUint32();
  if (f != null) {
    writer.writeMessage(
      5,
      f,
      proto.buf.validate.UInt32Rules.serializeBinaryToWriter
    );
  }
  f = message.getUint64();
  if (f != null) {
    writer.writeMessage(
      6,
      f,
      proto.buf.validate.UInt64Rules.serializeBinaryToWriter
    );
  }
  f = message.getSint32();
  if (f != null) {
    writer.writeMessage(
      7,
      f,
      proto.buf.validate.SInt32Rules.serializeBinaryToWriter
    );
  }
  f = message.getSint64();
  if (f != null) {
    writer.writeMessage(
      8,
      f,
      proto.buf.validate.SInt64Rules.serializeBinaryToWriter
    );
  }
  f = message.getFixed32();
  if (f != null) {
    writer.writeMessage(
      9,
      f,
      proto.buf.validate.Fixed32Rules.serializeBinaryToWriter
    );
  }
  f = message.getFixed64();
  if (f != null) {
    writer.writeMessage(
      10,
      f,
      proto.buf.validate.Fixed64Rules.serializeBinaryToWriter
    );
  }
  f = message.getSfixed32();
  if (f != null) {
    writer.writeMessage(
      11,
      f,
      proto.buf.validate.SFixed32Rules.serializeBinaryToWriter
    );
  }
  f = message.getSfixed64();
  if (f != null) {
    writer.writeMessage(
      12,
      f,
      proto.buf.validate.SFixed64Rules.serializeBinaryToWriter
    );
  }
  f = message.getBool();
  if (f != null) {
    writer.writeMessage(
      13,
      f,
      proto.buf.validate.BoolRules.serializeBinaryToWriter
    );
  }
  f = message.getString();
  if (f != null) {
    writer.writeMessage(
      14,
      f,
      proto.buf.validate.StringRules.serializeBinaryToWriter
    );
  }
  f = message.getBytes();
  if (f != null) {
    writer.writeMessage(
      15,
      f,
      proto.buf.validate.BytesRules.serializeBinaryToWriter
    );
  }
  f = message.getEnum();
  if (f != null) {
    writer.writeMessage(
      16,
      f,
      proto.buf.validate.EnumRules.serializeBinaryToWriter
    );
  }
  f = message.getRepeated();
  if (f != null) {
    writer.writeMessage(
      18,
      f,
      proto.buf.validate.RepeatedRules.serializeBinaryToWriter
    );
  }
  f = message.getMap();
  if (f != null) {
    writer.writeMessage(
      19,
      f,
      proto.buf.validate.MapRules.serializeBinaryToWriter
    );
  }
  f = message.getAny();
  if (f != null) {
    writer.writeMessage(
      20,
      f,
      proto.buf.validate.AnyRules.serializeBinaryToWriter
    );
  }
  f = message.getDuration();
  if (f != null) {
    writer.writeMessage(
      21,
      f,
      proto.buf.validate.DurationRules.serializeBinaryToWriter
    );
  }
  f = message.getTimestamp();
  if (f != null) {
    writer.writeMessage(
      22,
      f,
      proto.buf.validate.TimestampRules.serializeBinaryToWriter
    );
  }
};


/**
 * repeated Constraint cel = 23;
 * @return {!Array<!proto.buf.validate.Constraint>}
 */
proto.buf.validate.FieldConstraints.prototype.getCelList = function() {
  return /** @type{!Array<!proto.buf.validate.Constraint>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.buf.validate.Constraint, 23));
};


/**
 * @param {!Array<!proto.buf.validate.Constraint>} value
 * @return {!proto.buf.validate.FieldConstraints} returns this
*/
proto.buf.validate.FieldConstraints.prototype.setCelList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 23, value);
};


/**
 * @param {!proto.buf.validate.Constraint=} opt_value
 * @param {number=} opt_index
 * @return {!proto.buf.validate.Constraint}
 */
proto.buf.validate.FieldConstraints.prototype.addCel = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 23, opt_value, proto.buf.validate.Constraint, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.buf.validate.FieldConstraints} returns this
 */
proto.buf.validate.FieldConstraints.prototype.clearCelList = function() {
  return this.setCelList([]);
};


/**
 * optional bool required = 25;
 * @return {boolean}
 */
proto.buf.validate.FieldConstraints.prototype.getRequired = function() {
  return /** @type {boolean} */ (jspb.Message.getBooleanFieldWithDefault(this, 25, false));
};


/**
 * @param {boolean} value
 * @return {!proto.buf.validate.FieldConstraints} returns this
 */
proto.buf.validate.FieldConstraints.prototype.setRequired = function(value) {
  return jspb.Message.setField(this, 25, value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.FieldConstraints} returns this
 */
proto.buf.validate.FieldConstraints.prototype.clearRequired = function() {
  return jspb.Message.setField(this, 25, undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.FieldConstraints.prototype.hasRequired = function() {
  return jspb.Message.getField(this, 25) != null;
};


/**
 * optional Ignore ignore = 27;
 * @return {!proto.buf.validate.Ignore}
 */
proto.buf.validate.FieldConstraints.prototype.getIgnore = function() {
  return /** @type {!proto.buf.validate.Ignore} */ (jspb.Message.getFieldWithDefault(this, 27, 0));
};


/**
 * @param {!proto.buf.validate.Ignore} value
 * @return {!proto.buf.validate.FieldConstraints} returns this
 */
proto.buf.validate.FieldConstraints.prototype.setIgnore = function(value) {
  return jspb.Message.setField(this, 27, value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.FieldConstraints} returns this
 */
proto.buf.validate.FieldConstraints.prototype.clearIgnore = function() {
  return jspb.Message.setField(this, 27, undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.FieldConstraints.prototype.hasIgnore = function() {
  return jspb.Message.getField(this, 27) != null;
};


/**
 * optional FloatRules float = 1;
 * @return {?proto.buf.validate.FloatRules}
 */
proto.buf.validate.FieldConstraints.prototype.getFloat = function() {
  return /** @type{?proto.buf.validate.FloatRules} */ (
    jspb.Message.getWrapperField(this, proto.buf.validate.FloatRules, 1));
};


/**
 * @param {?proto.buf.validate.FloatRules|undefined} value
 * @return {!proto.buf.validate.FieldConstraints} returns this
*/
proto.buf.validate.FieldConstraints.prototype.setFloat = function(value) {
  return jspb.Message.setOneofWrapperField(this, 1, proto.buf.validate.FieldConstraints.oneofGroups_[0], value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.buf.validate.FieldConstraints} returns this
 */
proto.buf.validate.FieldConstraints.prototype.clearFloat = function() {
  return this.setFloat(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.FieldConstraints.prototype.hasFloat = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional DoubleRules double = 2;
 * @return {?proto.buf.validate.DoubleRules}
 */
proto.buf.validate.FieldConstraints.prototype.getDouble = function() {
  return /** @type{?proto.buf.validate.DoubleRules} */ (
    jspb.Message.getWrapperField(this, proto.buf.validate.DoubleRules, 2));
};


/**
 * @param {?proto.buf.validate.DoubleRules|undefined} value
 * @return {!proto.buf.validate.FieldConstraints} returns this
*/
proto.buf.validate.FieldConstraints.prototype.setDouble = function(value) {
  return jspb.Message.setOneofWrapperField(this, 2, proto.buf.validate.FieldConstraints.oneofGroups_[0], value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.buf.validate.FieldConstraints} returns this
 */
proto.buf.validate.FieldConstraints.prototype.clearDouble = function() {
  return this.setDouble(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.FieldConstraints.prototype.hasDouble = function() {
  return jspb.Message.getField(this, 2) != null;
};


/**
 * optional Int32Rules int32 = 3;
 * @return {?proto.buf.validate.Int32Rules}
 */
proto.buf.validate.FieldConstraints.prototype.getInt32 = function() {
  return /** @type{?proto.buf.validate.Int32Rules} */ (
    jspb.Message.getWrapperField(this, proto.buf.validate.Int32Rules, 3));
};


/**
 * @param {?proto.buf.validate.Int32Rules|undefined} value
 * @return {!proto.buf.validate.FieldConstraints} returns this
*/
proto.buf.validate.FieldConstraints.prototype.setInt32 = function(value) {
  return jspb.Message.setOneofWrapperField(this, 3, proto.buf.validate.FieldConstraints.oneofGroups_[0], value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.buf.validate.FieldConstraints} returns this
 */
proto.buf.validate.FieldConstraints.prototype.clearInt32 = function() {
  return this.setInt32(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.FieldConstraints.prototype.hasInt32 = function() {
  return jspb.Message.getField(this, 3) != null;
};


/**
 * optional Int64Rules int64 = 4;
 * @return {?proto.buf.validate.Int64Rules}
 */
proto.buf.validate.FieldConstraints.prototype.getInt64 = function() {
  return /** @type{?proto.buf.validate.Int64Rules} */ (
    jspb.Message.getWrapperField(this, proto.buf.validate.Int64Rules, 4));
};


/**
 * @param {?proto.buf.validate.Int64Rules|undefined} value
 * @return {!proto.buf.validate.FieldConstraints} returns this
*/
proto.buf.validate.FieldConstraints.prototype.setInt64 = function(value) {
  return jspb.Message.setOneofWrapperField(this, 4, proto.buf.validate.FieldConstraints.oneofGroups_[0], value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.buf.validate.FieldConstraints} returns this
 */
proto.buf.validate.FieldConstraints.prototype.clearInt64 = function() {
  return this.setInt64(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.FieldConstraints.prototype.hasInt64 = function() {
  return jspb.Message.getField(this, 4) != null;
};


/**
 * optional UInt32Rules uint32 = 5;
 * @return {?proto.buf.validate.UInt32Rules}
 */
proto.buf.validate.FieldConstraints.prototype.getUint32 = function() {
  return /** @type{?proto.buf.validate.UInt32Rules} */ (
    jspb.Message.getWrapperField(this, proto.buf.validate.UInt32Rules, 5));
};


/**
 * @param {?proto.buf.validate.UInt32Rules|undefined} value
 * @return {!proto.buf.validate.FieldConstraints} returns this
*/
proto.buf.validate.FieldConstraints.prototype.setUint32 = function(value) {
  return jspb.Message.setOneofWrapperField(this, 5, proto.buf.validate.FieldConstraints.oneofGroups_[0], value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.buf.validate.FieldConstraints} returns this
 */
proto.buf.validate.FieldConstraints.prototype.clearUint32 = function() {
  return this.setUint32(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.FieldConstraints.prototype.hasUint32 = function() {
  return jspb.Message.getField(this, 5) != null;
};


/**
 * optional UInt64Rules uint64 = 6;
 * @return {?proto.buf.validate.UInt64Rules}
 */
proto.buf.validate.FieldConstraints.prototype.getUint64 = function() {
  return /** @type{?proto.buf.validate.UInt64Rules} */ (
    jspb.Message.getWrapperField(this, proto.buf.validate.UInt64Rules, 6));
};


/**
 * @param {?proto.buf.validate.UInt64Rules|undefined} value
 * @return {!proto.buf.validate.FieldConstraints} returns this
*/
proto.buf.validate.FieldConstraints.prototype.setUint64 = function(value) {
  return jspb.Message.setOneofWrapperField(this, 6, proto.buf.validate.FieldConstraints.oneofGroups_[0], value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.buf.validate.FieldConstraints} returns this
 */
proto.buf.validate.FieldConstraints.prototype.clearUint64 = function() {
  return this.setUint64(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.FieldConstraints.prototype.hasUint64 = function() {
  return jspb.Message.getField(this, 6) != null;
};


/**
 * optional SInt32Rules sint32 = 7;
 * @return {?proto.buf.validate.SInt32Rules}
 */
proto.buf.validate.FieldConstraints.prototype.getSint32 = function() {
  return /** @type{?proto.buf.validate.SInt32Rules} */ (
    jspb.Message.getWrapperField(this, proto.buf.validate.SInt32Rules, 7));
};


/**
 * @param {?proto.buf.validate.SInt32Rules|undefined} value
 * @return {!proto.buf.validate.FieldConstraints} returns this
*/
proto.buf.validate.FieldConstraints.prototype.setSint32 = function(value) {
  return jspb.Message.setOneofWrapperField(this, 7, proto.buf.validate.FieldConstraints.oneofGroups_[0], value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.buf.validate.FieldConstraints} returns this
 */
proto.buf.validate.FieldConstraints.prototype.clearSint32 = function() {
  return this.setSint32(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.FieldConstraints.prototype.hasSint32 = function() {
  return jspb.Message.getField(this, 7) != null;
};


/**
 * optional SInt64Rules sint64 = 8;
 * @return {?proto.buf.validate.SInt64Rules}
 */
proto.buf.validate.FieldConstraints.prototype.getSint64 = function() {
  return /** @type{?proto.buf.validate.SInt64Rules} */ (
    jspb.Message.getWrapperField(this, proto.buf.validate.SInt64Rules, 8));
};


/**
 * @param {?proto.buf.validate.SInt64Rules|undefined} value
 * @return {!proto.buf.validate.FieldConstraints} returns this
*/
proto.buf.validate.FieldConstraints.prototype.setSint64 = function(value) {
  return jspb.Message.setOneofWrapperField(this, 8, proto.buf.validate.FieldConstraints.oneofGroups_[0], value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.buf.validate.FieldConstraints} returns this
 */
proto.buf.validate.FieldConstraints.prototype.clearSint64 = function() {
  return this.setSint64(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.FieldConstraints.prototype.hasSint64 = function() {
  return jspb.Message.getField(this, 8) != null;
};


/**
 * optional Fixed32Rules fixed32 = 9;
 * @return {?proto.buf.validate.Fixed32Rules}
 */
proto.buf.validate.FieldConstraints.prototype.getFixed32 = function() {
  return /** @type{?proto.buf.validate.Fixed32Rules} */ (
    jspb.Message.getWrapperField(this, proto.buf.validate.Fixed32Rules, 9));
};


/**
 * @param {?proto.buf.validate.Fixed32Rules|undefined} value
 * @return {!proto.buf.validate.FieldConstraints} returns this
*/
proto.buf.validate.FieldConstraints.prototype.setFixed32 = function(value) {
  return jspb.Message.setOneofWrapperField(this, 9, proto.buf.validate.FieldConstraints.oneofGroups_[0], value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.buf.validate.FieldConstraints} returns this
 */
proto.buf.validate.FieldConstraints.prototype.clearFixed32 = function() {
  return this.setFixed32(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.FieldConstraints.prototype.hasFixed32 = function() {
  return jspb.Message.getField(this, 9) != null;
};


/**
 * optional Fixed64Rules fixed64 = 10;
 * @return {?proto.buf.validate.Fixed64Rules}
 */
proto.buf.validate.FieldConstraints.prototype.getFixed64 = function() {
  return /** @type{?proto.buf.validate.Fixed64Rules} */ (
    jspb.Message.getWrapperField(this, proto.buf.validate.Fixed64Rules, 10));
};


/**
 * @param {?proto.buf.validate.Fixed64Rules|undefined} value
 * @return {!proto.buf.validate.FieldConstraints} returns this
*/
proto.buf.validate.FieldConstraints.prototype.setFixed64 = function(value) {
  return jspb.Message.setOneofWrapperField(this, 10, proto.buf.validate.FieldConstraints.oneofGroups_[0], value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.buf.validate.FieldConstraints} returns this
 */
proto.buf.validate.FieldConstraints.prototype.clearFixed64 = function() {
  return this.setFixed64(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.FieldConstraints.prototype.hasFixed64 = function() {
  return jspb.Message.getField(this, 10) != null;
};


/**
 * optional SFixed32Rules sfixed32 = 11;
 * @return {?proto.buf.validate.SFixed32Rules}
 */
proto.buf.validate.FieldConstraints.prototype.getSfixed32 = function() {
  return /** @type{?proto.buf.validate.SFixed32Rules} */ (
    jspb.Message.getWrapperField(this, proto.buf.validate.SFixed32Rules, 11));
};


/**
 * @param {?proto.buf.validate.SFixed32Rules|undefined} value
 * @return {!proto.buf.validate.FieldConstraints} returns this
*/
proto.buf.validate.FieldConstraints.prototype.setSfixed32 = function(value) {
  return jspb.Message.setOneofWrapperField(this, 11, proto.buf.validate.FieldConstraints.oneofGroups_[0], value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.buf.validate.FieldConstraints} returns this
 */
proto.buf.validate.FieldConstraints.prototype.clearSfixed32 = function() {
  return this.setSfixed32(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.FieldConstraints.prototype.hasSfixed32 = function() {
  return jspb.Message.getField(this, 11) != null;
};


/**
 * optional SFixed64Rules sfixed64 = 12;
 * @return {?proto.buf.validate.SFixed64Rules}
 */
proto.buf.validate.FieldConstraints.prototype.getSfixed64 = function() {
  return /** @type{?proto.buf.validate.SFixed64Rules} */ (
    jspb.Message.getWrapperField(this, proto.buf.validate.SFixed64Rules, 12));
};


/**
 * @param {?proto.buf.validate.SFixed64Rules|undefined} value
 * @return {!proto.buf.validate.FieldConstraints} returns this
*/
proto.buf.validate.FieldConstraints.prototype.setSfixed64 = function(value) {
  return jspb.Message.setOneofWrapperField(this, 12, proto.buf.validate.FieldConstraints.oneofGroups_[0], value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.buf.validate.FieldConstraints} returns this
 */
proto.buf.validate.FieldConstraints.prototype.clearSfixed64 = function() {
  return this.setSfixed64(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.FieldConstraints.prototype.hasSfixed64 = function() {
  return jspb.Message.getField(this, 12) != null;
};


/**
 * optional BoolRules bool = 13;
 * @return {?proto.buf.validate.BoolRules}
 */
proto.buf.validate.FieldConstraints.prototype.getBool = function() {
  return /** @type{?proto.buf.validate.BoolRules} */ (
    jspb.Message.getWrapperField(this, proto.buf.validate.BoolRules, 13));
};


/**
 * @param {?proto.buf.validate.BoolRules|undefined} value
 * @return {!proto.buf.validate.FieldConstraints} returns this
*/
proto.buf.validate.FieldConstraints.prototype.setBool = function(value) {
  return jspb.Message.setOneofWrapperField(this, 13, proto.buf.validate.FieldConstraints.oneofGroups_[0], value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.buf.validate.FieldConstraints} returns this
 */
proto.buf.validate.FieldConstraints.prototype.clearBool = function() {
  return this.setBool(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.FieldConstraints.prototype.hasBool = function() {
  return jspb.Message.getField(this, 13) != null;
};


/**
 * optional StringRules string = 14;
 * @return {?proto.buf.validate.StringRules}
 */
proto.buf.validate.FieldConstraints.prototype.getString = function() {
  return /** @type{?proto.buf.validate.StringRules} */ (
    jspb.Message.getWrapperField(this, proto.buf.validate.StringRules, 14));
};


/**
 * @param {?proto.buf.validate.StringRules|undefined} value
 * @return {!proto.buf.validate.FieldConstraints} returns this
*/
proto.buf.validate.FieldConstraints.prototype.setString = function(value) {
  return jspb.Message.setOneofWrapperField(this, 14, proto.buf.validate.FieldConstraints.oneofGroups_[0], value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.buf.validate.FieldConstraints} returns this
 */
proto.buf.validate.FieldConstraints.prototype.clearString = function() {
  return this.setString(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.FieldConstraints.prototype.hasString = function() {
  return jspb.Message.getField(this, 14) != null;
};


/**
 * optional BytesRules bytes = 15;
 * @return {?proto.buf.validate.BytesRules}
 */
proto.buf.validate.FieldConstraints.prototype.getBytes = function() {
  return /** @type{?proto.buf.validate.BytesRules} */ (
    jspb.Message.getWrapperField(this, proto.buf.validate.BytesRules, 15));
};


/**
 * @param {?proto.buf.validate.BytesRules|undefined} value
 * @return {!proto.buf.validate.FieldConstraints} returns this
*/
proto.buf.validate.FieldConstraints.prototype.setBytes = function(value) {
  return jspb.Message.setOneofWrapperField(this, 15, proto.buf.validate.FieldConstraints.oneofGroups_[0], value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.buf.validate.FieldConstraints} returns this
 */
proto.buf.validate.FieldConstraints.prototype.clearBytes = function() {
  return this.setBytes(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.FieldConstraints.prototype.hasBytes = function() {
  return jspb.Message.getField(this, 15) != null;
};


/**
 * optional EnumRules enum = 16;
 * @return {?proto.buf.validate.EnumRules}
 */
proto.buf.validate.FieldConstraints.prototype.getEnum = function() {
  return /** @type{?proto.buf.validate.EnumRules} */ (
    jspb.Message.getWrapperField(this, proto.buf.validate.EnumRules, 16));
};


/**
 * @param {?proto.buf.validate.EnumRules|undefined} value
 * @return {!proto.buf.validate.FieldConstraints} returns this
*/
proto.buf.validate.FieldConstraints.prototype.setEnum = function(value) {
  return jspb.Message.setOneofWrapperField(this, 16, proto.buf.validate.FieldConstraints.oneofGroups_[0], value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.buf.validate.FieldConstraints} returns this
 */
proto.buf.validate.FieldConstraints.prototype.clearEnum = function() {
  return this.setEnum(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.FieldConstraints.prototype.hasEnum = function() {
  return jspb.Message.getField(this, 16) != null;
};


/**
 * optional RepeatedRules repeated = 18;
 * @return {?proto.buf.validate.RepeatedRules}
 */
proto.buf.validate.FieldConstraints.prototype.getRepeated = function() {
  return /** @type{?proto.buf.validate.RepeatedRules} */ (
    jspb.Message.getWrapperField(this, proto.buf.validate.RepeatedRules, 18));
};


/**
 * @param {?proto.buf.validate.RepeatedRules|undefined} value
 * @return {!proto.buf.validate.FieldConstraints} returns this
*/
proto.buf.validate.FieldConstraints.prototype.setRepeated = function(value) {
  return jspb.Message.setOneofWrapperField(this, 18, proto.buf.validate.FieldConstraints.oneofGroups_[0], value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.buf.validate.FieldConstraints} returns this
 */
proto.buf.validate.FieldConstraints.prototype.clearRepeated = function() {
  return this.setRepeated(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.FieldConstraints.prototype.hasRepeated = function() {
  return jspb.Message.getField(this, 18) != null;
};


/**
 * optional MapRules map = 19;
 * @return {?proto.buf.validate.MapRules}
 */
proto.buf.validate.FieldConstraints.prototype.getMap = function() {
  return /** @type{?proto.buf.validate.MapRules} */ (
    jspb.Message.getWrapperField(this, proto.buf.validate.MapRules, 19));
};


/**
 * @param {?proto.buf.validate.MapRules|undefined} value
 * @return {!proto.buf.validate.FieldConstraints} returns this
*/
proto.buf.validate.FieldConstraints.prototype.setMap = function(value) {
  return jspb.Message.setOneofWrapperField(this, 19, proto.buf.validate.FieldConstraints.oneofGroups_[0], value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.buf.validate.FieldConstraints} returns this
 */
proto.buf.validate.FieldConstraints.prototype.clearMap = function() {
  return this.setMap(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.FieldConstraints.prototype.hasMap = function() {
  return jspb.Message.getField(this, 19) != null;
};


/**
 * optional AnyRules any = 20;
 * @return {?proto.buf.validate.AnyRules}
 */
proto.buf.validate.FieldConstraints.prototype.getAny = function() {
  return /** @type{?proto.buf.validate.AnyRules} */ (
    jspb.Message.getWrapperField(this, proto.buf.validate.AnyRules, 20));
};


/**
 * @param {?proto.buf.validate.AnyRules|undefined} value
 * @return {!proto.buf.validate.FieldConstraints} returns this
*/
proto.buf.validate.FieldConstraints.prototype.setAny = function(value) {
  return jspb.Message.setOneofWrapperField(this, 20, proto.buf.validate.FieldConstraints.oneofGroups_[0], value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.buf.validate.FieldConstraints} returns this
 */
proto.buf.validate.FieldConstraints.prototype.clearAny = function() {
  return this.setAny(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.FieldConstraints.prototype.hasAny = function() {
  return jspb.Message.getField(this, 20) != null;
};


/**
 * optional DurationRules duration = 21;
 * @return {?proto.buf.validate.DurationRules}
 */
proto.buf.validate.FieldConstraints.prototype.getDuration = function() {
  return /** @type{?proto.buf.validate.DurationRules} */ (
    jspb.Message.getWrapperField(this, proto.buf.validate.DurationRules, 21));
};


/**
 * @param {?proto.buf.validate.DurationRules|undefined} value
 * @return {!proto.buf.validate.FieldConstraints} returns this
*/
proto.buf.validate.FieldConstraints.prototype.setDuration = function(value) {
  return jspb.Message.setOneofWrapperField(this, 21, proto.buf.validate.FieldConstraints.oneofGroups_[0], value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.buf.validate.FieldConstraints} returns this
 */
proto.buf.validate.FieldConstraints.prototype.clearDuration = function() {
  return this.setDuration(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.FieldConstraints.prototype.hasDuration = function() {
  return jspb.Message.getField(this, 21) != null;
};


/**
 * optional TimestampRules timestamp = 22;
 * @return {?proto.buf.validate.TimestampRules}
 */
proto.buf.validate.FieldConstraints.prototype.getTimestamp = function() {
  return /** @type{?proto.buf.validate.TimestampRules} */ (
    jspb.Message.getWrapperField(this, proto.buf.validate.TimestampRules, 22));
};


/**
 * @param {?proto.buf.validate.TimestampRules|undefined} value
 * @return {!proto.buf.validate.FieldConstraints} returns this
*/
proto.buf.validate.FieldConstraints.prototype.setTimestamp = function(value) {
  return jspb.Message.setOneofWrapperField(this, 22, proto.buf.validate.FieldConstraints.oneofGroups_[0], value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.buf.validate.FieldConstraints} returns this
 */
proto.buf.validate.FieldConstraints.prototype.clearTimestamp = function() {
  return this.setTimestamp(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.FieldConstraints.prototype.hasTimestamp = function() {
  return jspb.Message.getField(this, 22) != null;
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.buf.validate.PredefinedConstraints.repeatedFields_ = [1];



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
proto.buf.validate.PredefinedConstraints.prototype.toObject = function(opt_includeInstance) {
  return proto.buf.validate.PredefinedConstraints.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.buf.validate.PredefinedConstraints} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.buf.validate.PredefinedConstraints.toObject = function(includeInstance, msg) {
  var f, obj = {
celList: jspb.Message.toObjectList(msg.getCelList(),
    proto.buf.validate.Constraint.toObject, includeInstance)
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
 * @return {!proto.buf.validate.PredefinedConstraints}
 */
proto.buf.validate.PredefinedConstraints.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.buf.validate.PredefinedConstraints;
  return proto.buf.validate.PredefinedConstraints.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.buf.validate.PredefinedConstraints} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.buf.validate.PredefinedConstraints}
 */
proto.buf.validate.PredefinedConstraints.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.buf.validate.Constraint;
      reader.readMessage(value,proto.buf.validate.Constraint.deserializeBinaryFromReader);
      msg.addCel(value);
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
proto.buf.validate.PredefinedConstraints.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.buf.validate.PredefinedConstraints.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.buf.validate.PredefinedConstraints} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.buf.validate.PredefinedConstraints.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getCelList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      1,
      f,
      proto.buf.validate.Constraint.serializeBinaryToWriter
    );
  }
};


/**
 * repeated Constraint cel = 1;
 * @return {!Array<!proto.buf.validate.Constraint>}
 */
proto.buf.validate.PredefinedConstraints.prototype.getCelList = function() {
  return /** @type{!Array<!proto.buf.validate.Constraint>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.buf.validate.Constraint, 1));
};


/**
 * @param {!Array<!proto.buf.validate.Constraint>} value
 * @return {!proto.buf.validate.PredefinedConstraints} returns this
*/
proto.buf.validate.PredefinedConstraints.prototype.setCelList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 1, value);
};


/**
 * @param {!proto.buf.validate.Constraint=} opt_value
 * @param {number=} opt_index
 * @return {!proto.buf.validate.Constraint}
 */
proto.buf.validate.PredefinedConstraints.prototype.addCel = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 1, opt_value, proto.buf.validate.Constraint, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.buf.validate.PredefinedConstraints} returns this
 */
proto.buf.validate.PredefinedConstraints.prototype.clearCelList = function() {
  return this.setCelList([]);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.buf.validate.FloatRules.repeatedFields_ = [6,7,9];

/**
 * Oneof group definitions for this message. Each group defines the field
 * numbers belonging to that group. When of these fields' value is set, all
 * other fields in the group are cleared. During deserialization, if multiple
 * fields are encountered for a group, only the last value seen will be kept.
 * @private {!Array<!Array<number>>}
 * @const
 */
proto.buf.validate.FloatRules.oneofGroups_ = [[2,3],[4,5]];

/**
 * @enum {number}
 */
proto.buf.validate.FloatRules.LessThanCase = {
  LESS_THAN_NOT_SET: 0,
  LT: 2,
  LTE: 3
};

/**
 * @return {proto.buf.validate.FloatRules.LessThanCase}
 */
proto.buf.validate.FloatRules.prototype.getLessThanCase = function() {
  return /** @type {proto.buf.validate.FloatRules.LessThanCase} */(jspb.Message.computeOneofCase(this, proto.buf.validate.FloatRules.oneofGroups_[0]));
};

/**
 * @enum {number}
 */
proto.buf.validate.FloatRules.GreaterThanCase = {
  GREATER_THAN_NOT_SET: 0,
  GT: 4,
  GTE: 5
};

/**
 * @return {proto.buf.validate.FloatRules.GreaterThanCase}
 */
proto.buf.validate.FloatRules.prototype.getGreaterThanCase = function() {
  return /** @type {proto.buf.validate.FloatRules.GreaterThanCase} */(jspb.Message.computeOneofCase(this, proto.buf.validate.FloatRules.oneofGroups_[1]));
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
proto.buf.validate.FloatRules.prototype.toObject = function(opt_includeInstance) {
  return proto.buf.validate.FloatRules.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.buf.validate.FloatRules} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.buf.validate.FloatRules.toObject = function(includeInstance, msg) {
  var f, obj = {
pb_const: (f = jspb.Message.getOptionalFloatingPointField(msg, 1)) == null ? undefined : f,
lt: (f = jspb.Message.getOptionalFloatingPointField(msg, 2)) == null ? undefined : f,
lte: (f = jspb.Message.getOptionalFloatingPointField(msg, 3)) == null ? undefined : f,
gt: (f = jspb.Message.getOptionalFloatingPointField(msg, 4)) == null ? undefined : f,
gte: (f = jspb.Message.getOptionalFloatingPointField(msg, 5)) == null ? undefined : f,
inList: (f = jspb.Message.getRepeatedFloatingPointField(msg, 6)) == null ? undefined : f,
notInList: (f = jspb.Message.getRepeatedFloatingPointField(msg, 7)) == null ? undefined : f,
finite: (f = jspb.Message.getBooleanField(msg, 8)) == null ? undefined : f,
exampleList: (f = jspb.Message.getRepeatedFloatingPointField(msg, 9)) == null ? undefined : f
  };

  jspb.Message.toObjectExtension(/** @type {!jspb.Message} */ (msg), obj,
      proto.buf.validate.FloatRules.extensions, proto.buf.validate.FloatRules.prototype.getExtension,
      includeInstance);
  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.buf.validate.FloatRules}
 */
proto.buf.validate.FloatRules.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.buf.validate.FloatRules;
  return proto.buf.validate.FloatRules.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.buf.validate.FloatRules} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.buf.validate.FloatRules}
 */
proto.buf.validate.FloatRules.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readFloat());
      msg.setConst(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readFloat());
      msg.setLt(value);
      break;
    case 3:
      var value = /** @type {number} */ (reader.readFloat());
      msg.setLte(value);
      break;
    case 4:
      var value = /** @type {number} */ (reader.readFloat());
      msg.setGt(value);
      break;
    case 5:
      var value = /** @type {number} */ (reader.readFloat());
      msg.setGte(value);
      break;
    case 6:
      var values = /** @type {!Array<number>} */ (reader.isDelimited() ? reader.readPackedFloat() : [reader.readFloat()]);
      for (var i = 0; i < values.length; i++) {
        msg.addIn(values[i]);
      }
      break;
    case 7:
      var values = /** @type {!Array<number>} */ (reader.isDelimited() ? reader.readPackedFloat() : [reader.readFloat()]);
      for (var i = 0; i < values.length; i++) {
        msg.addNotIn(values[i]);
      }
      break;
    case 8:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setFinite(value);
      break;
    case 9:
      var values = /** @type {!Array<number>} */ (reader.isDelimited() ? reader.readPackedFloat() : [reader.readFloat()]);
      for (var i = 0; i < values.length; i++) {
        msg.addExample(values[i]);
      }
      break;
    default:
      jspb.Message.readBinaryExtension(msg, reader,
        proto.buf.validate.FloatRules.extensionsBinary,
        proto.buf.validate.FloatRules.prototype.getExtension,
        proto.buf.validate.FloatRules.prototype.setExtension);
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.buf.validate.FloatRules.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.buf.validate.FloatRules.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.buf.validate.FloatRules} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.buf.validate.FloatRules.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = /** @type {number} */ (jspb.Message.getField(message, 1));
  if (f != null) {
    writer.writeFloat(
      1,
      f
    );
  }
  f = /** @type {number} */ (jspb.Message.getField(message, 2));
  if (f != null) {
    writer.writeFloat(
      2,
      f
    );
  }
  f = /** @type {number} */ (jspb.Message.getField(message, 3));
  if (f != null) {
    writer.writeFloat(
      3,
      f
    );
  }
  f = /** @type {number} */ (jspb.Message.getField(message, 4));
  if (f != null) {
    writer.writeFloat(
      4,
      f
    );
  }
  f = /** @type {number} */ (jspb.Message.getField(message, 5));
  if (f != null) {
    writer.writeFloat(
      5,
      f
    );
  }
  f = message.getInList();
  if (f.length > 0) {
    writer.writeRepeatedFloat(
      6,
      f
    );
  }
  f = message.getNotInList();
  if (f.length > 0) {
    writer.writeRepeatedFloat(
      7,
      f
    );
  }
  f = /** @type {boolean} */ (jspb.Message.getField(message, 8));
  if (f != null) {
    writer.writeBool(
      8,
      f
    );
  }
  f = message.getExampleList();
  if (f.length > 0) {
    writer.writeRepeatedFloat(
      9,
      f
    );
  }
  jspb.Message.serializeBinaryExtensions(message, writer,
    proto.buf.validate.FloatRules.extensionsBinary, proto.buf.validate.FloatRules.prototype.getExtension);
};


/**
 * optional float const = 1;
 * @return {number}
 */
proto.buf.validate.FloatRules.prototype.getConst = function() {
  return /** @type {number} */ (jspb.Message.getFloatingPointFieldWithDefault(this, 1, 0.0));
};


/**
 * @param {number} value
 * @return {!proto.buf.validate.FloatRules} returns this
 */
proto.buf.validate.FloatRules.prototype.setConst = function(value) {
  return jspb.Message.setField(this, 1, value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.FloatRules} returns this
 */
proto.buf.validate.FloatRules.prototype.clearConst = function() {
  return jspb.Message.setField(this, 1, undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.FloatRules.prototype.hasConst = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional float lt = 2;
 * @return {number}
 */
proto.buf.validate.FloatRules.prototype.getLt = function() {
  return /** @type {number} */ (jspb.Message.getFloatingPointFieldWithDefault(this, 2, 0.0));
};


/**
 * @param {number} value
 * @return {!proto.buf.validate.FloatRules} returns this
 */
proto.buf.validate.FloatRules.prototype.setLt = function(value) {
  return jspb.Message.setOneofField(this, 2, proto.buf.validate.FloatRules.oneofGroups_[0], value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.FloatRules} returns this
 */
proto.buf.validate.FloatRules.prototype.clearLt = function() {
  return jspb.Message.setOneofField(this, 2, proto.buf.validate.FloatRules.oneofGroups_[0], undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.FloatRules.prototype.hasLt = function() {
  return jspb.Message.getField(this, 2) != null;
};


/**
 * optional float lte = 3;
 * @return {number}
 */
proto.buf.validate.FloatRules.prototype.getLte = function() {
  return /** @type {number} */ (jspb.Message.getFloatingPointFieldWithDefault(this, 3, 0.0));
};


/**
 * @param {number} value
 * @return {!proto.buf.validate.FloatRules} returns this
 */
proto.buf.validate.FloatRules.prototype.setLte = function(value) {
  return jspb.Message.setOneofField(this, 3, proto.buf.validate.FloatRules.oneofGroups_[0], value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.FloatRules} returns this
 */
proto.buf.validate.FloatRules.prototype.clearLte = function() {
  return jspb.Message.setOneofField(this, 3, proto.buf.validate.FloatRules.oneofGroups_[0], undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.FloatRules.prototype.hasLte = function() {
  return jspb.Message.getField(this, 3) != null;
};


/**
 * optional float gt = 4;
 * @return {number}
 */
proto.buf.validate.FloatRules.prototype.getGt = function() {
  return /** @type {number} */ (jspb.Message.getFloatingPointFieldWithDefault(this, 4, 0.0));
};


/**
 * @param {number} value
 * @return {!proto.buf.validate.FloatRules} returns this
 */
proto.buf.validate.FloatRules.prototype.setGt = function(value) {
  return jspb.Message.setOneofField(this, 4, proto.buf.validate.FloatRules.oneofGroups_[1], value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.FloatRules} returns this
 */
proto.buf.validate.FloatRules.prototype.clearGt = function() {
  return jspb.Message.setOneofField(this, 4, proto.buf.validate.FloatRules.oneofGroups_[1], undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.FloatRules.prototype.hasGt = function() {
  return jspb.Message.getField(this, 4) != null;
};


/**
 * optional float gte = 5;
 * @return {number}
 */
proto.buf.validate.FloatRules.prototype.getGte = function() {
  return /** @type {number} */ (jspb.Message.getFloatingPointFieldWithDefault(this, 5, 0.0));
};


/**
 * @param {number} value
 * @return {!proto.buf.validate.FloatRules} returns this
 */
proto.buf.validate.FloatRules.prototype.setGte = function(value) {
  return jspb.Message.setOneofField(this, 5, proto.buf.validate.FloatRules.oneofGroups_[1], value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.FloatRules} returns this
 */
proto.buf.validate.FloatRules.prototype.clearGte = function() {
  return jspb.Message.setOneofField(this, 5, proto.buf.validate.FloatRules.oneofGroups_[1], undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.FloatRules.prototype.hasGte = function() {
  return jspb.Message.getField(this, 5) != null;
};


/**
 * repeated float in = 6;
 * @return {!Array<number>}
 */
proto.buf.validate.FloatRules.prototype.getInList = function() {
  return /** @type {!Array<number>} */ (jspb.Message.getRepeatedFloatingPointField(this, 6));
};


/**
 * @param {!Array<number>} value
 * @return {!proto.buf.validate.FloatRules} returns this
 */
proto.buf.validate.FloatRules.prototype.setInList = function(value) {
  return jspb.Message.setField(this, 6, value || []);
};


/**
 * @param {number} value
 * @param {number=} opt_index
 * @return {!proto.buf.validate.FloatRules} returns this
 */
proto.buf.validate.FloatRules.prototype.addIn = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 6, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.buf.validate.FloatRules} returns this
 */
proto.buf.validate.FloatRules.prototype.clearInList = function() {
  return this.setInList([]);
};


/**
 * repeated float not_in = 7;
 * @return {!Array<number>}
 */
proto.buf.validate.FloatRules.prototype.getNotInList = function() {
  return /** @type {!Array<number>} */ (jspb.Message.getRepeatedFloatingPointField(this, 7));
};


/**
 * @param {!Array<number>} value
 * @return {!proto.buf.validate.FloatRules} returns this
 */
proto.buf.validate.FloatRules.prototype.setNotInList = function(value) {
  return jspb.Message.setField(this, 7, value || []);
};


/**
 * @param {number} value
 * @param {number=} opt_index
 * @return {!proto.buf.validate.FloatRules} returns this
 */
proto.buf.validate.FloatRules.prototype.addNotIn = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 7, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.buf.validate.FloatRules} returns this
 */
proto.buf.validate.FloatRules.prototype.clearNotInList = function() {
  return this.setNotInList([]);
};


/**
 * optional bool finite = 8;
 * @return {boolean}
 */
proto.buf.validate.FloatRules.prototype.getFinite = function() {
  return /** @type {boolean} */ (jspb.Message.getBooleanFieldWithDefault(this, 8, false));
};


/**
 * @param {boolean} value
 * @return {!proto.buf.validate.FloatRules} returns this
 */
proto.buf.validate.FloatRules.prototype.setFinite = function(value) {
  return jspb.Message.setField(this, 8, value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.FloatRules} returns this
 */
proto.buf.validate.FloatRules.prototype.clearFinite = function() {
  return jspb.Message.setField(this, 8, undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.FloatRules.prototype.hasFinite = function() {
  return jspb.Message.getField(this, 8) != null;
};


/**
 * repeated float example = 9;
 * @return {!Array<number>}
 */
proto.buf.validate.FloatRules.prototype.getExampleList = function() {
  return /** @type {!Array<number>} */ (jspb.Message.getRepeatedFloatingPointField(this, 9));
};


/**
 * @param {!Array<number>} value
 * @return {!proto.buf.validate.FloatRules} returns this
 */
proto.buf.validate.FloatRules.prototype.setExampleList = function(value) {
  return jspb.Message.setField(this, 9, value || []);
};


/**
 * @param {number} value
 * @param {number=} opt_index
 * @return {!proto.buf.validate.FloatRules} returns this
 */
proto.buf.validate.FloatRules.prototype.addExample = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 9, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.buf.validate.FloatRules} returns this
 */
proto.buf.validate.FloatRules.prototype.clearExampleList = function() {
  return this.setExampleList([]);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.buf.validate.DoubleRules.repeatedFields_ = [6,7,9];

/**
 * Oneof group definitions for this message. Each group defines the field
 * numbers belonging to that group. When of these fields' value is set, all
 * other fields in the group are cleared. During deserialization, if multiple
 * fields are encountered for a group, only the last value seen will be kept.
 * @private {!Array<!Array<number>>}
 * @const
 */
proto.buf.validate.DoubleRules.oneofGroups_ = [[2,3],[4,5]];

/**
 * @enum {number}
 */
proto.buf.validate.DoubleRules.LessThanCase = {
  LESS_THAN_NOT_SET: 0,
  LT: 2,
  LTE: 3
};

/**
 * @return {proto.buf.validate.DoubleRules.LessThanCase}
 */
proto.buf.validate.DoubleRules.prototype.getLessThanCase = function() {
  return /** @type {proto.buf.validate.DoubleRules.LessThanCase} */(jspb.Message.computeOneofCase(this, proto.buf.validate.DoubleRules.oneofGroups_[0]));
};

/**
 * @enum {number}
 */
proto.buf.validate.DoubleRules.GreaterThanCase = {
  GREATER_THAN_NOT_SET: 0,
  GT: 4,
  GTE: 5
};

/**
 * @return {proto.buf.validate.DoubleRules.GreaterThanCase}
 */
proto.buf.validate.DoubleRules.prototype.getGreaterThanCase = function() {
  return /** @type {proto.buf.validate.DoubleRules.GreaterThanCase} */(jspb.Message.computeOneofCase(this, proto.buf.validate.DoubleRules.oneofGroups_[1]));
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
proto.buf.validate.DoubleRules.prototype.toObject = function(opt_includeInstance) {
  return proto.buf.validate.DoubleRules.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.buf.validate.DoubleRules} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.buf.validate.DoubleRules.toObject = function(includeInstance, msg) {
  var f, obj = {
pb_const: (f = jspb.Message.getOptionalFloatingPointField(msg, 1)) == null ? undefined : f,
lt: (f = jspb.Message.getOptionalFloatingPointField(msg, 2)) == null ? undefined : f,
lte: (f = jspb.Message.getOptionalFloatingPointField(msg, 3)) == null ? undefined : f,
gt: (f = jspb.Message.getOptionalFloatingPointField(msg, 4)) == null ? undefined : f,
gte: (f = jspb.Message.getOptionalFloatingPointField(msg, 5)) == null ? undefined : f,
inList: (f = jspb.Message.getRepeatedFloatingPointField(msg, 6)) == null ? undefined : f,
notInList: (f = jspb.Message.getRepeatedFloatingPointField(msg, 7)) == null ? undefined : f,
finite: (f = jspb.Message.getBooleanField(msg, 8)) == null ? undefined : f,
exampleList: (f = jspb.Message.getRepeatedFloatingPointField(msg, 9)) == null ? undefined : f
  };

  jspb.Message.toObjectExtension(/** @type {!jspb.Message} */ (msg), obj,
      proto.buf.validate.DoubleRules.extensions, proto.buf.validate.DoubleRules.prototype.getExtension,
      includeInstance);
  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.buf.validate.DoubleRules}
 */
proto.buf.validate.DoubleRules.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.buf.validate.DoubleRules;
  return proto.buf.validate.DoubleRules.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.buf.validate.DoubleRules} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.buf.validate.DoubleRules}
 */
proto.buf.validate.DoubleRules.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readDouble());
      msg.setConst(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readDouble());
      msg.setLt(value);
      break;
    case 3:
      var value = /** @type {number} */ (reader.readDouble());
      msg.setLte(value);
      break;
    case 4:
      var value = /** @type {number} */ (reader.readDouble());
      msg.setGt(value);
      break;
    case 5:
      var value = /** @type {number} */ (reader.readDouble());
      msg.setGte(value);
      break;
    case 6:
      var values = /** @type {!Array<number>} */ (reader.isDelimited() ? reader.readPackedDouble() : [reader.readDouble()]);
      for (var i = 0; i < values.length; i++) {
        msg.addIn(values[i]);
      }
      break;
    case 7:
      var values = /** @type {!Array<number>} */ (reader.isDelimited() ? reader.readPackedDouble() : [reader.readDouble()]);
      for (var i = 0; i < values.length; i++) {
        msg.addNotIn(values[i]);
      }
      break;
    case 8:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setFinite(value);
      break;
    case 9:
      var values = /** @type {!Array<number>} */ (reader.isDelimited() ? reader.readPackedDouble() : [reader.readDouble()]);
      for (var i = 0; i < values.length; i++) {
        msg.addExample(values[i]);
      }
      break;
    default:
      jspb.Message.readBinaryExtension(msg, reader,
        proto.buf.validate.DoubleRules.extensionsBinary,
        proto.buf.validate.DoubleRules.prototype.getExtension,
        proto.buf.validate.DoubleRules.prototype.setExtension);
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.buf.validate.DoubleRules.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.buf.validate.DoubleRules.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.buf.validate.DoubleRules} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.buf.validate.DoubleRules.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = /** @type {number} */ (jspb.Message.getField(message, 1));
  if (f != null) {
    writer.writeDouble(
      1,
      f
    );
  }
  f = /** @type {number} */ (jspb.Message.getField(message, 2));
  if (f != null) {
    writer.writeDouble(
      2,
      f
    );
  }
  f = /** @type {number} */ (jspb.Message.getField(message, 3));
  if (f != null) {
    writer.writeDouble(
      3,
      f
    );
  }
  f = /** @type {number} */ (jspb.Message.getField(message, 4));
  if (f != null) {
    writer.writeDouble(
      4,
      f
    );
  }
  f = /** @type {number} */ (jspb.Message.getField(message, 5));
  if (f != null) {
    writer.writeDouble(
      5,
      f
    );
  }
  f = message.getInList();
  if (f.length > 0) {
    writer.writeRepeatedDouble(
      6,
      f
    );
  }
  f = message.getNotInList();
  if (f.length > 0) {
    writer.writeRepeatedDouble(
      7,
      f
    );
  }
  f = /** @type {boolean} */ (jspb.Message.getField(message, 8));
  if (f != null) {
    writer.writeBool(
      8,
      f
    );
  }
  f = message.getExampleList();
  if (f.length > 0) {
    writer.writeRepeatedDouble(
      9,
      f
    );
  }
  jspb.Message.serializeBinaryExtensions(message, writer,
    proto.buf.validate.DoubleRules.extensionsBinary, proto.buf.validate.DoubleRules.prototype.getExtension);
};


/**
 * optional double const = 1;
 * @return {number}
 */
proto.buf.validate.DoubleRules.prototype.getConst = function() {
  return /** @type {number} */ (jspb.Message.getFloatingPointFieldWithDefault(this, 1, 0.0));
};


/**
 * @param {number} value
 * @return {!proto.buf.validate.DoubleRules} returns this
 */
proto.buf.validate.DoubleRules.prototype.setConst = function(value) {
  return jspb.Message.setField(this, 1, value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.DoubleRules} returns this
 */
proto.buf.validate.DoubleRules.prototype.clearConst = function() {
  return jspb.Message.setField(this, 1, undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.DoubleRules.prototype.hasConst = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional double lt = 2;
 * @return {number}
 */
proto.buf.validate.DoubleRules.prototype.getLt = function() {
  return /** @type {number} */ (jspb.Message.getFloatingPointFieldWithDefault(this, 2, 0.0));
};


/**
 * @param {number} value
 * @return {!proto.buf.validate.DoubleRules} returns this
 */
proto.buf.validate.DoubleRules.prototype.setLt = function(value) {
  return jspb.Message.setOneofField(this, 2, proto.buf.validate.DoubleRules.oneofGroups_[0], value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.DoubleRules} returns this
 */
proto.buf.validate.DoubleRules.prototype.clearLt = function() {
  return jspb.Message.setOneofField(this, 2, proto.buf.validate.DoubleRules.oneofGroups_[0], undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.DoubleRules.prototype.hasLt = function() {
  return jspb.Message.getField(this, 2) != null;
};


/**
 * optional double lte = 3;
 * @return {number}
 */
proto.buf.validate.DoubleRules.prototype.getLte = function() {
  return /** @type {number} */ (jspb.Message.getFloatingPointFieldWithDefault(this, 3, 0.0));
};


/**
 * @param {number} value
 * @return {!proto.buf.validate.DoubleRules} returns this
 */
proto.buf.validate.DoubleRules.prototype.setLte = function(value) {
  return jspb.Message.setOneofField(this, 3, proto.buf.validate.DoubleRules.oneofGroups_[0], value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.DoubleRules} returns this
 */
proto.buf.validate.DoubleRules.prototype.clearLte = function() {
  return jspb.Message.setOneofField(this, 3, proto.buf.validate.DoubleRules.oneofGroups_[0], undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.DoubleRules.prototype.hasLte = function() {
  return jspb.Message.getField(this, 3) != null;
};


/**
 * optional double gt = 4;
 * @return {number}
 */
proto.buf.validate.DoubleRules.prototype.getGt = function() {
  return /** @type {number} */ (jspb.Message.getFloatingPointFieldWithDefault(this, 4, 0.0));
};


/**
 * @param {number} value
 * @return {!proto.buf.validate.DoubleRules} returns this
 */
proto.buf.validate.DoubleRules.prototype.setGt = function(value) {
  return jspb.Message.setOneofField(this, 4, proto.buf.validate.DoubleRules.oneofGroups_[1], value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.DoubleRules} returns this
 */
proto.buf.validate.DoubleRules.prototype.clearGt = function() {
  return jspb.Message.setOneofField(this, 4, proto.buf.validate.DoubleRules.oneofGroups_[1], undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.DoubleRules.prototype.hasGt = function() {
  return jspb.Message.getField(this, 4) != null;
};


/**
 * optional double gte = 5;
 * @return {number}
 */
proto.buf.validate.DoubleRules.prototype.getGte = function() {
  return /** @type {number} */ (jspb.Message.getFloatingPointFieldWithDefault(this, 5, 0.0));
};


/**
 * @param {number} value
 * @return {!proto.buf.validate.DoubleRules} returns this
 */
proto.buf.validate.DoubleRules.prototype.setGte = function(value) {
  return jspb.Message.setOneofField(this, 5, proto.buf.validate.DoubleRules.oneofGroups_[1], value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.DoubleRules} returns this
 */
proto.buf.validate.DoubleRules.prototype.clearGte = function() {
  return jspb.Message.setOneofField(this, 5, proto.buf.validate.DoubleRules.oneofGroups_[1], undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.DoubleRules.prototype.hasGte = function() {
  return jspb.Message.getField(this, 5) != null;
};


/**
 * repeated double in = 6;
 * @return {!Array<number>}
 */
proto.buf.validate.DoubleRules.prototype.getInList = function() {
  return /** @type {!Array<number>} */ (jspb.Message.getRepeatedFloatingPointField(this, 6));
};


/**
 * @param {!Array<number>} value
 * @return {!proto.buf.validate.DoubleRules} returns this
 */
proto.buf.validate.DoubleRules.prototype.setInList = function(value) {
  return jspb.Message.setField(this, 6, value || []);
};


/**
 * @param {number} value
 * @param {number=} opt_index
 * @return {!proto.buf.validate.DoubleRules} returns this
 */
proto.buf.validate.DoubleRules.prototype.addIn = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 6, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.buf.validate.DoubleRules} returns this
 */
proto.buf.validate.DoubleRules.prototype.clearInList = function() {
  return this.setInList([]);
};


/**
 * repeated double not_in = 7;
 * @return {!Array<number>}
 */
proto.buf.validate.DoubleRules.prototype.getNotInList = function() {
  return /** @type {!Array<number>} */ (jspb.Message.getRepeatedFloatingPointField(this, 7));
};


/**
 * @param {!Array<number>} value
 * @return {!proto.buf.validate.DoubleRules} returns this
 */
proto.buf.validate.DoubleRules.prototype.setNotInList = function(value) {
  return jspb.Message.setField(this, 7, value || []);
};


/**
 * @param {number} value
 * @param {number=} opt_index
 * @return {!proto.buf.validate.DoubleRules} returns this
 */
proto.buf.validate.DoubleRules.prototype.addNotIn = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 7, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.buf.validate.DoubleRules} returns this
 */
proto.buf.validate.DoubleRules.prototype.clearNotInList = function() {
  return this.setNotInList([]);
};


/**
 * optional bool finite = 8;
 * @return {boolean}
 */
proto.buf.validate.DoubleRules.prototype.getFinite = function() {
  return /** @type {boolean} */ (jspb.Message.getBooleanFieldWithDefault(this, 8, false));
};


/**
 * @param {boolean} value
 * @return {!proto.buf.validate.DoubleRules} returns this
 */
proto.buf.validate.DoubleRules.prototype.setFinite = function(value) {
  return jspb.Message.setField(this, 8, value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.DoubleRules} returns this
 */
proto.buf.validate.DoubleRules.prototype.clearFinite = function() {
  return jspb.Message.setField(this, 8, undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.DoubleRules.prototype.hasFinite = function() {
  return jspb.Message.getField(this, 8) != null;
};


/**
 * repeated double example = 9;
 * @return {!Array<number>}
 */
proto.buf.validate.DoubleRules.prototype.getExampleList = function() {
  return /** @type {!Array<number>} */ (jspb.Message.getRepeatedFloatingPointField(this, 9));
};


/**
 * @param {!Array<number>} value
 * @return {!proto.buf.validate.DoubleRules} returns this
 */
proto.buf.validate.DoubleRules.prototype.setExampleList = function(value) {
  return jspb.Message.setField(this, 9, value || []);
};


/**
 * @param {number} value
 * @param {number=} opt_index
 * @return {!proto.buf.validate.DoubleRules} returns this
 */
proto.buf.validate.DoubleRules.prototype.addExample = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 9, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.buf.validate.DoubleRules} returns this
 */
proto.buf.validate.DoubleRules.prototype.clearExampleList = function() {
  return this.setExampleList([]);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.buf.validate.Int32Rules.repeatedFields_ = [6,7,8];

/**
 * Oneof group definitions for this message. Each group defines the field
 * numbers belonging to that group. When of these fields' value is set, all
 * other fields in the group are cleared. During deserialization, if multiple
 * fields are encountered for a group, only the last value seen will be kept.
 * @private {!Array<!Array<number>>}
 * @const
 */
proto.buf.validate.Int32Rules.oneofGroups_ = [[2,3],[4,5]];

/**
 * @enum {number}
 */
proto.buf.validate.Int32Rules.LessThanCase = {
  LESS_THAN_NOT_SET: 0,
  LT: 2,
  LTE: 3
};

/**
 * @return {proto.buf.validate.Int32Rules.LessThanCase}
 */
proto.buf.validate.Int32Rules.prototype.getLessThanCase = function() {
  return /** @type {proto.buf.validate.Int32Rules.LessThanCase} */(jspb.Message.computeOneofCase(this, proto.buf.validate.Int32Rules.oneofGroups_[0]));
};

/**
 * @enum {number}
 */
proto.buf.validate.Int32Rules.GreaterThanCase = {
  GREATER_THAN_NOT_SET: 0,
  GT: 4,
  GTE: 5
};

/**
 * @return {proto.buf.validate.Int32Rules.GreaterThanCase}
 */
proto.buf.validate.Int32Rules.prototype.getGreaterThanCase = function() {
  return /** @type {proto.buf.validate.Int32Rules.GreaterThanCase} */(jspb.Message.computeOneofCase(this, proto.buf.validate.Int32Rules.oneofGroups_[1]));
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
proto.buf.validate.Int32Rules.prototype.toObject = function(opt_includeInstance) {
  return proto.buf.validate.Int32Rules.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.buf.validate.Int32Rules} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.buf.validate.Int32Rules.toObject = function(includeInstance, msg) {
  var f, obj = {
pb_const: (f = jspb.Message.getField(msg, 1)) == null ? undefined : f,
lt: (f = jspb.Message.getField(msg, 2)) == null ? undefined : f,
lte: (f = jspb.Message.getField(msg, 3)) == null ? undefined : f,
gt: (f = jspb.Message.getField(msg, 4)) == null ? undefined : f,
gte: (f = jspb.Message.getField(msg, 5)) == null ? undefined : f,
inList: (f = jspb.Message.getRepeatedField(msg, 6)) == null ? undefined : f,
notInList: (f = jspb.Message.getRepeatedField(msg, 7)) == null ? undefined : f,
exampleList: (f = jspb.Message.getRepeatedField(msg, 8)) == null ? undefined : f
  };

  jspb.Message.toObjectExtension(/** @type {!jspb.Message} */ (msg), obj,
      proto.buf.validate.Int32Rules.extensions, proto.buf.validate.Int32Rules.prototype.getExtension,
      includeInstance);
  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.buf.validate.Int32Rules}
 */
proto.buf.validate.Int32Rules.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.buf.validate.Int32Rules;
  return proto.buf.validate.Int32Rules.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.buf.validate.Int32Rules} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.buf.validate.Int32Rules}
 */
proto.buf.validate.Int32Rules.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setConst(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setLt(value);
      break;
    case 3:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setLte(value);
      break;
    case 4:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setGt(value);
      break;
    case 5:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setGte(value);
      break;
    case 6:
      var values = /** @type {!Array<number>} */ (reader.isDelimited() ? reader.readPackedInt32() : [reader.readInt32()]);
      for (var i = 0; i < values.length; i++) {
        msg.addIn(values[i]);
      }
      break;
    case 7:
      var values = /** @type {!Array<number>} */ (reader.isDelimited() ? reader.readPackedInt32() : [reader.readInt32()]);
      for (var i = 0; i < values.length; i++) {
        msg.addNotIn(values[i]);
      }
      break;
    case 8:
      var values = /** @type {!Array<number>} */ (reader.isDelimited() ? reader.readPackedInt32() : [reader.readInt32()]);
      for (var i = 0; i < values.length; i++) {
        msg.addExample(values[i]);
      }
      break;
    default:
      jspb.Message.readBinaryExtension(msg, reader,
        proto.buf.validate.Int32Rules.extensionsBinary,
        proto.buf.validate.Int32Rules.prototype.getExtension,
        proto.buf.validate.Int32Rules.prototype.setExtension);
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.buf.validate.Int32Rules.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.buf.validate.Int32Rules.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.buf.validate.Int32Rules} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.buf.validate.Int32Rules.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = /** @type {number} */ (jspb.Message.getField(message, 1));
  if (f != null) {
    writer.writeInt32(
      1,
      f
    );
  }
  f = /** @type {number} */ (jspb.Message.getField(message, 2));
  if (f != null) {
    writer.writeInt32(
      2,
      f
    );
  }
  f = /** @type {number} */ (jspb.Message.getField(message, 3));
  if (f != null) {
    writer.writeInt32(
      3,
      f
    );
  }
  f = /** @type {number} */ (jspb.Message.getField(message, 4));
  if (f != null) {
    writer.writeInt32(
      4,
      f
    );
  }
  f = /** @type {number} */ (jspb.Message.getField(message, 5));
  if (f != null) {
    writer.writeInt32(
      5,
      f
    );
  }
  f = message.getInList();
  if (f.length > 0) {
    writer.writeRepeatedInt32(
      6,
      f
    );
  }
  f = message.getNotInList();
  if (f.length > 0) {
    writer.writeRepeatedInt32(
      7,
      f
    );
  }
  f = message.getExampleList();
  if (f.length > 0) {
    writer.writeRepeatedInt32(
      8,
      f
    );
  }
  jspb.Message.serializeBinaryExtensions(message, writer,
    proto.buf.validate.Int32Rules.extensionsBinary, proto.buf.validate.Int32Rules.prototype.getExtension);
};


/**
 * optional int32 const = 1;
 * @return {number}
 */
proto.buf.validate.Int32Rules.prototype.getConst = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/**
 * @param {number} value
 * @return {!proto.buf.validate.Int32Rules} returns this
 */
proto.buf.validate.Int32Rules.prototype.setConst = function(value) {
  return jspb.Message.setField(this, 1, value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.Int32Rules} returns this
 */
proto.buf.validate.Int32Rules.prototype.clearConst = function() {
  return jspb.Message.setField(this, 1, undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.Int32Rules.prototype.hasConst = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional int32 lt = 2;
 * @return {number}
 */
proto.buf.validate.Int32Rules.prototype.getLt = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/**
 * @param {number} value
 * @return {!proto.buf.validate.Int32Rules} returns this
 */
proto.buf.validate.Int32Rules.prototype.setLt = function(value) {
  return jspb.Message.setOneofField(this, 2, proto.buf.validate.Int32Rules.oneofGroups_[0], value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.Int32Rules} returns this
 */
proto.buf.validate.Int32Rules.prototype.clearLt = function() {
  return jspb.Message.setOneofField(this, 2, proto.buf.validate.Int32Rules.oneofGroups_[0], undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.Int32Rules.prototype.hasLt = function() {
  return jspb.Message.getField(this, 2) != null;
};


/**
 * optional int32 lte = 3;
 * @return {number}
 */
proto.buf.validate.Int32Rules.prototype.getLte = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 3, 0));
};


/**
 * @param {number} value
 * @return {!proto.buf.validate.Int32Rules} returns this
 */
proto.buf.validate.Int32Rules.prototype.setLte = function(value) {
  return jspb.Message.setOneofField(this, 3, proto.buf.validate.Int32Rules.oneofGroups_[0], value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.Int32Rules} returns this
 */
proto.buf.validate.Int32Rules.prototype.clearLte = function() {
  return jspb.Message.setOneofField(this, 3, proto.buf.validate.Int32Rules.oneofGroups_[0], undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.Int32Rules.prototype.hasLte = function() {
  return jspb.Message.getField(this, 3) != null;
};


/**
 * optional int32 gt = 4;
 * @return {number}
 */
proto.buf.validate.Int32Rules.prototype.getGt = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 4, 0));
};


/**
 * @param {number} value
 * @return {!proto.buf.validate.Int32Rules} returns this
 */
proto.buf.validate.Int32Rules.prototype.setGt = function(value) {
  return jspb.Message.setOneofField(this, 4, proto.buf.validate.Int32Rules.oneofGroups_[1], value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.Int32Rules} returns this
 */
proto.buf.validate.Int32Rules.prototype.clearGt = function() {
  return jspb.Message.setOneofField(this, 4, proto.buf.validate.Int32Rules.oneofGroups_[1], undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.Int32Rules.prototype.hasGt = function() {
  return jspb.Message.getField(this, 4) != null;
};


/**
 * optional int32 gte = 5;
 * @return {number}
 */
proto.buf.validate.Int32Rules.prototype.getGte = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 5, 0));
};


/**
 * @param {number} value
 * @return {!proto.buf.validate.Int32Rules} returns this
 */
proto.buf.validate.Int32Rules.prototype.setGte = function(value) {
  return jspb.Message.setOneofField(this, 5, proto.buf.validate.Int32Rules.oneofGroups_[1], value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.Int32Rules} returns this
 */
proto.buf.validate.Int32Rules.prototype.clearGte = function() {
  return jspb.Message.setOneofField(this, 5, proto.buf.validate.Int32Rules.oneofGroups_[1], undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.Int32Rules.prototype.hasGte = function() {
  return jspb.Message.getField(this, 5) != null;
};


/**
 * repeated int32 in = 6;
 * @return {!Array<number>}
 */
proto.buf.validate.Int32Rules.prototype.getInList = function() {
  return /** @type {!Array<number>} */ (jspb.Message.getRepeatedField(this, 6));
};


/**
 * @param {!Array<number>} value
 * @return {!proto.buf.validate.Int32Rules} returns this
 */
proto.buf.validate.Int32Rules.prototype.setInList = function(value) {
  return jspb.Message.setField(this, 6, value || []);
};


/**
 * @param {number} value
 * @param {number=} opt_index
 * @return {!proto.buf.validate.Int32Rules} returns this
 */
proto.buf.validate.Int32Rules.prototype.addIn = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 6, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.buf.validate.Int32Rules} returns this
 */
proto.buf.validate.Int32Rules.prototype.clearInList = function() {
  return this.setInList([]);
};


/**
 * repeated int32 not_in = 7;
 * @return {!Array<number>}
 */
proto.buf.validate.Int32Rules.prototype.getNotInList = function() {
  return /** @type {!Array<number>} */ (jspb.Message.getRepeatedField(this, 7));
};


/**
 * @param {!Array<number>} value
 * @return {!proto.buf.validate.Int32Rules} returns this
 */
proto.buf.validate.Int32Rules.prototype.setNotInList = function(value) {
  return jspb.Message.setField(this, 7, value || []);
};


/**
 * @param {number} value
 * @param {number=} opt_index
 * @return {!proto.buf.validate.Int32Rules} returns this
 */
proto.buf.validate.Int32Rules.prototype.addNotIn = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 7, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.buf.validate.Int32Rules} returns this
 */
proto.buf.validate.Int32Rules.prototype.clearNotInList = function() {
  return this.setNotInList([]);
};


/**
 * repeated int32 example = 8;
 * @return {!Array<number>}
 */
proto.buf.validate.Int32Rules.prototype.getExampleList = function() {
  return /** @type {!Array<number>} */ (jspb.Message.getRepeatedField(this, 8));
};


/**
 * @param {!Array<number>} value
 * @return {!proto.buf.validate.Int32Rules} returns this
 */
proto.buf.validate.Int32Rules.prototype.setExampleList = function(value) {
  return jspb.Message.setField(this, 8, value || []);
};


/**
 * @param {number} value
 * @param {number=} opt_index
 * @return {!proto.buf.validate.Int32Rules} returns this
 */
proto.buf.validate.Int32Rules.prototype.addExample = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 8, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.buf.validate.Int32Rules} returns this
 */
proto.buf.validate.Int32Rules.prototype.clearExampleList = function() {
  return this.setExampleList([]);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.buf.validate.Int64Rules.repeatedFields_ = [6,7,9];

/**
 * Oneof group definitions for this message. Each group defines the field
 * numbers belonging to that group. When of these fields' value is set, all
 * other fields in the group are cleared. During deserialization, if multiple
 * fields are encountered for a group, only the last value seen will be kept.
 * @private {!Array<!Array<number>>}
 * @const
 */
proto.buf.validate.Int64Rules.oneofGroups_ = [[2,3],[4,5]];

/**
 * @enum {number}
 */
proto.buf.validate.Int64Rules.LessThanCase = {
  LESS_THAN_NOT_SET: 0,
  LT: 2,
  LTE: 3
};

/**
 * @return {proto.buf.validate.Int64Rules.LessThanCase}
 */
proto.buf.validate.Int64Rules.prototype.getLessThanCase = function() {
  return /** @type {proto.buf.validate.Int64Rules.LessThanCase} */(jspb.Message.computeOneofCase(this, proto.buf.validate.Int64Rules.oneofGroups_[0]));
};

/**
 * @enum {number}
 */
proto.buf.validate.Int64Rules.GreaterThanCase = {
  GREATER_THAN_NOT_SET: 0,
  GT: 4,
  GTE: 5
};

/**
 * @return {proto.buf.validate.Int64Rules.GreaterThanCase}
 */
proto.buf.validate.Int64Rules.prototype.getGreaterThanCase = function() {
  return /** @type {proto.buf.validate.Int64Rules.GreaterThanCase} */(jspb.Message.computeOneofCase(this, proto.buf.validate.Int64Rules.oneofGroups_[1]));
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
proto.buf.validate.Int64Rules.prototype.toObject = function(opt_includeInstance) {
  return proto.buf.validate.Int64Rules.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.buf.validate.Int64Rules} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.buf.validate.Int64Rules.toObject = function(includeInstance, msg) {
  var f, obj = {
pb_const: (f = jspb.Message.getField(msg, 1)) == null ? undefined : f,
lt: (f = jspb.Message.getField(msg, 2)) == null ? undefined : f,
lte: (f = jspb.Message.getField(msg, 3)) == null ? undefined : f,
gt: (f = jspb.Message.getField(msg, 4)) == null ? undefined : f,
gte: (f = jspb.Message.getField(msg, 5)) == null ? undefined : f,
inList: (f = jspb.Message.getRepeatedField(msg, 6)) == null ? undefined : f,
notInList: (f = jspb.Message.getRepeatedField(msg, 7)) == null ? undefined : f,
exampleList: (f = jspb.Message.getRepeatedField(msg, 9)) == null ? undefined : f
  };

  jspb.Message.toObjectExtension(/** @type {!jspb.Message} */ (msg), obj,
      proto.buf.validate.Int64Rules.extensions, proto.buf.validate.Int64Rules.prototype.getExtension,
      includeInstance);
  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.buf.validate.Int64Rules}
 */
proto.buf.validate.Int64Rules.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.buf.validate.Int64Rules;
  return proto.buf.validate.Int64Rules.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.buf.validate.Int64Rules} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.buf.validate.Int64Rules}
 */
proto.buf.validate.Int64Rules.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setConst(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setLt(value);
      break;
    case 3:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setLte(value);
      break;
    case 4:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setGt(value);
      break;
    case 5:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setGte(value);
      break;
    case 6:
      var values = /** @type {!Array<number>} */ (reader.isDelimited() ? reader.readPackedInt64() : [reader.readInt64()]);
      for (var i = 0; i < values.length; i++) {
        msg.addIn(values[i]);
      }
      break;
    case 7:
      var values = /** @type {!Array<number>} */ (reader.isDelimited() ? reader.readPackedInt64() : [reader.readInt64()]);
      for (var i = 0; i < values.length; i++) {
        msg.addNotIn(values[i]);
      }
      break;
    case 9:
      var values = /** @type {!Array<number>} */ (reader.isDelimited() ? reader.readPackedInt64() : [reader.readInt64()]);
      for (var i = 0; i < values.length; i++) {
        msg.addExample(values[i]);
      }
      break;
    default:
      jspb.Message.readBinaryExtension(msg, reader,
        proto.buf.validate.Int64Rules.extensionsBinary,
        proto.buf.validate.Int64Rules.prototype.getExtension,
        proto.buf.validate.Int64Rules.prototype.setExtension);
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.buf.validate.Int64Rules.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.buf.validate.Int64Rules.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.buf.validate.Int64Rules} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.buf.validate.Int64Rules.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = /** @type {number} */ (jspb.Message.getField(message, 1));
  if (f != null) {
    writer.writeInt64(
      1,
      f
    );
  }
  f = /** @type {number} */ (jspb.Message.getField(message, 2));
  if (f != null) {
    writer.writeInt64(
      2,
      f
    );
  }
  f = /** @type {number} */ (jspb.Message.getField(message, 3));
  if (f != null) {
    writer.writeInt64(
      3,
      f
    );
  }
  f = /** @type {number} */ (jspb.Message.getField(message, 4));
  if (f != null) {
    writer.writeInt64(
      4,
      f
    );
  }
  f = /** @type {number} */ (jspb.Message.getField(message, 5));
  if (f != null) {
    writer.writeInt64(
      5,
      f
    );
  }
  f = message.getInList();
  if (f.length > 0) {
    writer.writeRepeatedInt64(
      6,
      f
    );
  }
  f = message.getNotInList();
  if (f.length > 0) {
    writer.writeRepeatedInt64(
      7,
      f
    );
  }
  f = message.getExampleList();
  if (f.length > 0) {
    writer.writeRepeatedInt64(
      9,
      f
    );
  }
  jspb.Message.serializeBinaryExtensions(message, writer,
    proto.buf.validate.Int64Rules.extensionsBinary, proto.buf.validate.Int64Rules.prototype.getExtension);
};


/**
 * optional int64 const = 1;
 * @return {number}
 */
proto.buf.validate.Int64Rules.prototype.getConst = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/**
 * @param {number} value
 * @return {!proto.buf.validate.Int64Rules} returns this
 */
proto.buf.validate.Int64Rules.prototype.setConst = function(value) {
  return jspb.Message.setField(this, 1, value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.Int64Rules} returns this
 */
proto.buf.validate.Int64Rules.prototype.clearConst = function() {
  return jspb.Message.setField(this, 1, undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.Int64Rules.prototype.hasConst = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional int64 lt = 2;
 * @return {number}
 */
proto.buf.validate.Int64Rules.prototype.getLt = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/**
 * @param {number} value
 * @return {!proto.buf.validate.Int64Rules} returns this
 */
proto.buf.validate.Int64Rules.prototype.setLt = function(value) {
  return jspb.Message.setOneofField(this, 2, proto.buf.validate.Int64Rules.oneofGroups_[0], value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.Int64Rules} returns this
 */
proto.buf.validate.Int64Rules.prototype.clearLt = function() {
  return jspb.Message.setOneofField(this, 2, proto.buf.validate.Int64Rules.oneofGroups_[0], undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.Int64Rules.prototype.hasLt = function() {
  return jspb.Message.getField(this, 2) != null;
};


/**
 * optional int64 lte = 3;
 * @return {number}
 */
proto.buf.validate.Int64Rules.prototype.getLte = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 3, 0));
};


/**
 * @param {number} value
 * @return {!proto.buf.validate.Int64Rules} returns this
 */
proto.buf.validate.Int64Rules.prototype.setLte = function(value) {
  return jspb.Message.setOneofField(this, 3, proto.buf.validate.Int64Rules.oneofGroups_[0], value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.Int64Rules} returns this
 */
proto.buf.validate.Int64Rules.prototype.clearLte = function() {
  return jspb.Message.setOneofField(this, 3, proto.buf.validate.Int64Rules.oneofGroups_[0], undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.Int64Rules.prototype.hasLte = function() {
  return jspb.Message.getField(this, 3) != null;
};


/**
 * optional int64 gt = 4;
 * @return {number}
 */
proto.buf.validate.Int64Rules.prototype.getGt = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 4, 0));
};


/**
 * @param {number} value
 * @return {!proto.buf.validate.Int64Rules} returns this
 */
proto.buf.validate.Int64Rules.prototype.setGt = function(value) {
  return jspb.Message.setOneofField(this, 4, proto.buf.validate.Int64Rules.oneofGroups_[1], value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.Int64Rules} returns this
 */
proto.buf.validate.Int64Rules.prototype.clearGt = function() {
  return jspb.Message.setOneofField(this, 4, proto.buf.validate.Int64Rules.oneofGroups_[1], undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.Int64Rules.prototype.hasGt = function() {
  return jspb.Message.getField(this, 4) != null;
};


/**
 * optional int64 gte = 5;
 * @return {number}
 */
proto.buf.validate.Int64Rules.prototype.getGte = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 5, 0));
};


/**
 * @param {number} value
 * @return {!proto.buf.validate.Int64Rules} returns this
 */
proto.buf.validate.Int64Rules.prototype.setGte = function(value) {
  return jspb.Message.setOneofField(this, 5, proto.buf.validate.Int64Rules.oneofGroups_[1], value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.Int64Rules} returns this
 */
proto.buf.validate.Int64Rules.prototype.clearGte = function() {
  return jspb.Message.setOneofField(this, 5, proto.buf.validate.Int64Rules.oneofGroups_[1], undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.Int64Rules.prototype.hasGte = function() {
  return jspb.Message.getField(this, 5) != null;
};


/**
 * repeated int64 in = 6;
 * @return {!Array<number>}
 */
proto.buf.validate.Int64Rules.prototype.getInList = function() {
  return /** @type {!Array<number>} */ (jspb.Message.getRepeatedField(this, 6));
};


/**
 * @param {!Array<number>} value
 * @return {!proto.buf.validate.Int64Rules} returns this
 */
proto.buf.validate.Int64Rules.prototype.setInList = function(value) {
  return jspb.Message.setField(this, 6, value || []);
};


/**
 * @param {number} value
 * @param {number=} opt_index
 * @return {!proto.buf.validate.Int64Rules} returns this
 */
proto.buf.validate.Int64Rules.prototype.addIn = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 6, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.buf.validate.Int64Rules} returns this
 */
proto.buf.validate.Int64Rules.prototype.clearInList = function() {
  return this.setInList([]);
};


/**
 * repeated int64 not_in = 7;
 * @return {!Array<number>}
 */
proto.buf.validate.Int64Rules.prototype.getNotInList = function() {
  return /** @type {!Array<number>} */ (jspb.Message.getRepeatedField(this, 7));
};


/**
 * @param {!Array<number>} value
 * @return {!proto.buf.validate.Int64Rules} returns this
 */
proto.buf.validate.Int64Rules.prototype.setNotInList = function(value) {
  return jspb.Message.setField(this, 7, value || []);
};


/**
 * @param {number} value
 * @param {number=} opt_index
 * @return {!proto.buf.validate.Int64Rules} returns this
 */
proto.buf.validate.Int64Rules.prototype.addNotIn = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 7, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.buf.validate.Int64Rules} returns this
 */
proto.buf.validate.Int64Rules.prototype.clearNotInList = function() {
  return this.setNotInList([]);
};


/**
 * repeated int64 example = 9;
 * @return {!Array<number>}
 */
proto.buf.validate.Int64Rules.prototype.getExampleList = function() {
  return /** @type {!Array<number>} */ (jspb.Message.getRepeatedField(this, 9));
};


/**
 * @param {!Array<number>} value
 * @return {!proto.buf.validate.Int64Rules} returns this
 */
proto.buf.validate.Int64Rules.prototype.setExampleList = function(value) {
  return jspb.Message.setField(this, 9, value || []);
};


/**
 * @param {number} value
 * @param {number=} opt_index
 * @return {!proto.buf.validate.Int64Rules} returns this
 */
proto.buf.validate.Int64Rules.prototype.addExample = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 9, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.buf.validate.Int64Rules} returns this
 */
proto.buf.validate.Int64Rules.prototype.clearExampleList = function() {
  return this.setExampleList([]);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.buf.validate.UInt32Rules.repeatedFields_ = [6,7,8];

/**
 * Oneof group definitions for this message. Each group defines the field
 * numbers belonging to that group. When of these fields' value is set, all
 * other fields in the group are cleared. During deserialization, if multiple
 * fields are encountered for a group, only the last value seen will be kept.
 * @private {!Array<!Array<number>>}
 * @const
 */
proto.buf.validate.UInt32Rules.oneofGroups_ = [[2,3],[4,5]];

/**
 * @enum {number}
 */
proto.buf.validate.UInt32Rules.LessThanCase = {
  LESS_THAN_NOT_SET: 0,
  LT: 2,
  LTE: 3
};

/**
 * @return {proto.buf.validate.UInt32Rules.LessThanCase}
 */
proto.buf.validate.UInt32Rules.prototype.getLessThanCase = function() {
  return /** @type {proto.buf.validate.UInt32Rules.LessThanCase} */(jspb.Message.computeOneofCase(this, proto.buf.validate.UInt32Rules.oneofGroups_[0]));
};

/**
 * @enum {number}
 */
proto.buf.validate.UInt32Rules.GreaterThanCase = {
  GREATER_THAN_NOT_SET: 0,
  GT: 4,
  GTE: 5
};

/**
 * @return {proto.buf.validate.UInt32Rules.GreaterThanCase}
 */
proto.buf.validate.UInt32Rules.prototype.getGreaterThanCase = function() {
  return /** @type {proto.buf.validate.UInt32Rules.GreaterThanCase} */(jspb.Message.computeOneofCase(this, proto.buf.validate.UInt32Rules.oneofGroups_[1]));
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
proto.buf.validate.UInt32Rules.prototype.toObject = function(opt_includeInstance) {
  return proto.buf.validate.UInt32Rules.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.buf.validate.UInt32Rules} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.buf.validate.UInt32Rules.toObject = function(includeInstance, msg) {
  var f, obj = {
pb_const: (f = jspb.Message.getField(msg, 1)) == null ? undefined : f,
lt: (f = jspb.Message.getField(msg, 2)) == null ? undefined : f,
lte: (f = jspb.Message.getField(msg, 3)) == null ? undefined : f,
gt: (f = jspb.Message.getField(msg, 4)) == null ? undefined : f,
gte: (f = jspb.Message.getField(msg, 5)) == null ? undefined : f,
inList: (f = jspb.Message.getRepeatedField(msg, 6)) == null ? undefined : f,
notInList: (f = jspb.Message.getRepeatedField(msg, 7)) == null ? undefined : f,
exampleList: (f = jspb.Message.getRepeatedField(msg, 8)) == null ? undefined : f
  };

  jspb.Message.toObjectExtension(/** @type {!jspb.Message} */ (msg), obj,
      proto.buf.validate.UInt32Rules.extensions, proto.buf.validate.UInt32Rules.prototype.getExtension,
      includeInstance);
  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.buf.validate.UInt32Rules}
 */
proto.buf.validate.UInt32Rules.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.buf.validate.UInt32Rules;
  return proto.buf.validate.UInt32Rules.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.buf.validate.UInt32Rules} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.buf.validate.UInt32Rules}
 */
proto.buf.validate.UInt32Rules.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readUint32());
      msg.setConst(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readUint32());
      msg.setLt(value);
      break;
    case 3:
      var value = /** @type {number} */ (reader.readUint32());
      msg.setLte(value);
      break;
    case 4:
      var value = /** @type {number} */ (reader.readUint32());
      msg.setGt(value);
      break;
    case 5:
      var value = /** @type {number} */ (reader.readUint32());
      msg.setGte(value);
      break;
    case 6:
      var values = /** @type {!Array<number>} */ (reader.isDelimited() ? reader.readPackedUint32() : [reader.readUint32()]);
      for (var i = 0; i < values.length; i++) {
        msg.addIn(values[i]);
      }
      break;
    case 7:
      var values = /** @type {!Array<number>} */ (reader.isDelimited() ? reader.readPackedUint32() : [reader.readUint32()]);
      for (var i = 0; i < values.length; i++) {
        msg.addNotIn(values[i]);
      }
      break;
    case 8:
      var values = /** @type {!Array<number>} */ (reader.isDelimited() ? reader.readPackedUint32() : [reader.readUint32()]);
      for (var i = 0; i < values.length; i++) {
        msg.addExample(values[i]);
      }
      break;
    default:
      jspb.Message.readBinaryExtension(msg, reader,
        proto.buf.validate.UInt32Rules.extensionsBinary,
        proto.buf.validate.UInt32Rules.prototype.getExtension,
        proto.buf.validate.UInt32Rules.prototype.setExtension);
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.buf.validate.UInt32Rules.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.buf.validate.UInt32Rules.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.buf.validate.UInt32Rules} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.buf.validate.UInt32Rules.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = /** @type {number} */ (jspb.Message.getField(message, 1));
  if (f != null) {
    writer.writeUint32(
      1,
      f
    );
  }
  f = /** @type {number} */ (jspb.Message.getField(message, 2));
  if (f != null) {
    writer.writeUint32(
      2,
      f
    );
  }
  f = /** @type {number} */ (jspb.Message.getField(message, 3));
  if (f != null) {
    writer.writeUint32(
      3,
      f
    );
  }
  f = /** @type {number} */ (jspb.Message.getField(message, 4));
  if (f != null) {
    writer.writeUint32(
      4,
      f
    );
  }
  f = /** @type {number} */ (jspb.Message.getField(message, 5));
  if (f != null) {
    writer.writeUint32(
      5,
      f
    );
  }
  f = message.getInList();
  if (f.length > 0) {
    writer.writeRepeatedUint32(
      6,
      f
    );
  }
  f = message.getNotInList();
  if (f.length > 0) {
    writer.writeRepeatedUint32(
      7,
      f
    );
  }
  f = message.getExampleList();
  if (f.length > 0) {
    writer.writeRepeatedUint32(
      8,
      f
    );
  }
  jspb.Message.serializeBinaryExtensions(message, writer,
    proto.buf.validate.UInt32Rules.extensionsBinary, proto.buf.validate.UInt32Rules.prototype.getExtension);
};


/**
 * optional uint32 const = 1;
 * @return {number}
 */
proto.buf.validate.UInt32Rules.prototype.getConst = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/**
 * @param {number} value
 * @return {!proto.buf.validate.UInt32Rules} returns this
 */
proto.buf.validate.UInt32Rules.prototype.setConst = function(value) {
  return jspb.Message.setField(this, 1, value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.UInt32Rules} returns this
 */
proto.buf.validate.UInt32Rules.prototype.clearConst = function() {
  return jspb.Message.setField(this, 1, undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.UInt32Rules.prototype.hasConst = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional uint32 lt = 2;
 * @return {number}
 */
proto.buf.validate.UInt32Rules.prototype.getLt = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/**
 * @param {number} value
 * @return {!proto.buf.validate.UInt32Rules} returns this
 */
proto.buf.validate.UInt32Rules.prototype.setLt = function(value) {
  return jspb.Message.setOneofField(this, 2, proto.buf.validate.UInt32Rules.oneofGroups_[0], value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.UInt32Rules} returns this
 */
proto.buf.validate.UInt32Rules.prototype.clearLt = function() {
  return jspb.Message.setOneofField(this, 2, proto.buf.validate.UInt32Rules.oneofGroups_[0], undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.UInt32Rules.prototype.hasLt = function() {
  return jspb.Message.getField(this, 2) != null;
};


/**
 * optional uint32 lte = 3;
 * @return {number}
 */
proto.buf.validate.UInt32Rules.prototype.getLte = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 3, 0));
};


/**
 * @param {number} value
 * @return {!proto.buf.validate.UInt32Rules} returns this
 */
proto.buf.validate.UInt32Rules.prototype.setLte = function(value) {
  return jspb.Message.setOneofField(this, 3, proto.buf.validate.UInt32Rules.oneofGroups_[0], value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.UInt32Rules} returns this
 */
proto.buf.validate.UInt32Rules.prototype.clearLte = function() {
  return jspb.Message.setOneofField(this, 3, proto.buf.validate.UInt32Rules.oneofGroups_[0], undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.UInt32Rules.prototype.hasLte = function() {
  return jspb.Message.getField(this, 3) != null;
};


/**
 * optional uint32 gt = 4;
 * @return {number}
 */
proto.buf.validate.UInt32Rules.prototype.getGt = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 4, 0));
};


/**
 * @param {number} value
 * @return {!proto.buf.validate.UInt32Rules} returns this
 */
proto.buf.validate.UInt32Rules.prototype.setGt = function(value) {
  return jspb.Message.setOneofField(this, 4, proto.buf.validate.UInt32Rules.oneofGroups_[1], value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.UInt32Rules} returns this
 */
proto.buf.validate.UInt32Rules.prototype.clearGt = function() {
  return jspb.Message.setOneofField(this, 4, proto.buf.validate.UInt32Rules.oneofGroups_[1], undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.UInt32Rules.prototype.hasGt = function() {
  return jspb.Message.getField(this, 4) != null;
};


/**
 * optional uint32 gte = 5;
 * @return {number}
 */
proto.buf.validate.UInt32Rules.prototype.getGte = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 5, 0));
};


/**
 * @param {number} value
 * @return {!proto.buf.validate.UInt32Rules} returns this
 */
proto.buf.validate.UInt32Rules.prototype.setGte = function(value) {
  return jspb.Message.setOneofField(this, 5, proto.buf.validate.UInt32Rules.oneofGroups_[1], value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.UInt32Rules} returns this
 */
proto.buf.validate.UInt32Rules.prototype.clearGte = function() {
  return jspb.Message.setOneofField(this, 5, proto.buf.validate.UInt32Rules.oneofGroups_[1], undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.UInt32Rules.prototype.hasGte = function() {
  return jspb.Message.getField(this, 5) != null;
};


/**
 * repeated uint32 in = 6;
 * @return {!Array<number>}
 */
proto.buf.validate.UInt32Rules.prototype.getInList = function() {
  return /** @type {!Array<number>} */ (jspb.Message.getRepeatedField(this, 6));
};


/**
 * @param {!Array<number>} value
 * @return {!proto.buf.validate.UInt32Rules} returns this
 */
proto.buf.validate.UInt32Rules.prototype.setInList = function(value) {
  return jspb.Message.setField(this, 6, value || []);
};


/**
 * @param {number} value
 * @param {number=} opt_index
 * @return {!proto.buf.validate.UInt32Rules} returns this
 */
proto.buf.validate.UInt32Rules.prototype.addIn = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 6, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.buf.validate.UInt32Rules} returns this
 */
proto.buf.validate.UInt32Rules.prototype.clearInList = function() {
  return this.setInList([]);
};


/**
 * repeated uint32 not_in = 7;
 * @return {!Array<number>}
 */
proto.buf.validate.UInt32Rules.prototype.getNotInList = function() {
  return /** @type {!Array<number>} */ (jspb.Message.getRepeatedField(this, 7));
};


/**
 * @param {!Array<number>} value
 * @return {!proto.buf.validate.UInt32Rules} returns this
 */
proto.buf.validate.UInt32Rules.prototype.setNotInList = function(value) {
  return jspb.Message.setField(this, 7, value || []);
};


/**
 * @param {number} value
 * @param {number=} opt_index
 * @return {!proto.buf.validate.UInt32Rules} returns this
 */
proto.buf.validate.UInt32Rules.prototype.addNotIn = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 7, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.buf.validate.UInt32Rules} returns this
 */
proto.buf.validate.UInt32Rules.prototype.clearNotInList = function() {
  return this.setNotInList([]);
};


/**
 * repeated uint32 example = 8;
 * @return {!Array<number>}
 */
proto.buf.validate.UInt32Rules.prototype.getExampleList = function() {
  return /** @type {!Array<number>} */ (jspb.Message.getRepeatedField(this, 8));
};


/**
 * @param {!Array<number>} value
 * @return {!proto.buf.validate.UInt32Rules} returns this
 */
proto.buf.validate.UInt32Rules.prototype.setExampleList = function(value) {
  return jspb.Message.setField(this, 8, value || []);
};


/**
 * @param {number} value
 * @param {number=} opt_index
 * @return {!proto.buf.validate.UInt32Rules} returns this
 */
proto.buf.validate.UInt32Rules.prototype.addExample = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 8, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.buf.validate.UInt32Rules} returns this
 */
proto.buf.validate.UInt32Rules.prototype.clearExampleList = function() {
  return this.setExampleList([]);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.buf.validate.UInt64Rules.repeatedFields_ = [6,7,8];

/**
 * Oneof group definitions for this message. Each group defines the field
 * numbers belonging to that group. When of these fields' value is set, all
 * other fields in the group are cleared. During deserialization, if multiple
 * fields are encountered for a group, only the last value seen will be kept.
 * @private {!Array<!Array<number>>}
 * @const
 */
proto.buf.validate.UInt64Rules.oneofGroups_ = [[2,3],[4,5]];

/**
 * @enum {number}
 */
proto.buf.validate.UInt64Rules.LessThanCase = {
  LESS_THAN_NOT_SET: 0,
  LT: 2,
  LTE: 3
};

/**
 * @return {proto.buf.validate.UInt64Rules.LessThanCase}
 */
proto.buf.validate.UInt64Rules.prototype.getLessThanCase = function() {
  return /** @type {proto.buf.validate.UInt64Rules.LessThanCase} */(jspb.Message.computeOneofCase(this, proto.buf.validate.UInt64Rules.oneofGroups_[0]));
};

/**
 * @enum {number}
 */
proto.buf.validate.UInt64Rules.GreaterThanCase = {
  GREATER_THAN_NOT_SET: 0,
  GT: 4,
  GTE: 5
};

/**
 * @return {proto.buf.validate.UInt64Rules.GreaterThanCase}
 */
proto.buf.validate.UInt64Rules.prototype.getGreaterThanCase = function() {
  return /** @type {proto.buf.validate.UInt64Rules.GreaterThanCase} */(jspb.Message.computeOneofCase(this, proto.buf.validate.UInt64Rules.oneofGroups_[1]));
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
proto.buf.validate.UInt64Rules.prototype.toObject = function(opt_includeInstance) {
  return proto.buf.validate.UInt64Rules.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.buf.validate.UInt64Rules} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.buf.validate.UInt64Rules.toObject = function(includeInstance, msg) {
  var f, obj = {
pb_const: (f = jspb.Message.getField(msg, 1)) == null ? undefined : f,
lt: (f = jspb.Message.getField(msg, 2)) == null ? undefined : f,
lte: (f = jspb.Message.getField(msg, 3)) == null ? undefined : f,
gt: (f = jspb.Message.getField(msg, 4)) == null ? undefined : f,
gte: (f = jspb.Message.getField(msg, 5)) == null ? undefined : f,
inList: (f = jspb.Message.getRepeatedField(msg, 6)) == null ? undefined : f,
notInList: (f = jspb.Message.getRepeatedField(msg, 7)) == null ? undefined : f,
exampleList: (f = jspb.Message.getRepeatedField(msg, 8)) == null ? undefined : f
  };

  jspb.Message.toObjectExtension(/** @type {!jspb.Message} */ (msg), obj,
      proto.buf.validate.UInt64Rules.extensions, proto.buf.validate.UInt64Rules.prototype.getExtension,
      includeInstance);
  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.buf.validate.UInt64Rules}
 */
proto.buf.validate.UInt64Rules.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.buf.validate.UInt64Rules;
  return proto.buf.validate.UInt64Rules.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.buf.validate.UInt64Rules} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.buf.validate.UInt64Rules}
 */
proto.buf.validate.UInt64Rules.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readUint64());
      msg.setConst(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readUint64());
      msg.setLt(value);
      break;
    case 3:
      var value = /** @type {number} */ (reader.readUint64());
      msg.setLte(value);
      break;
    case 4:
      var value = /** @type {number} */ (reader.readUint64());
      msg.setGt(value);
      break;
    case 5:
      var value = /** @type {number} */ (reader.readUint64());
      msg.setGte(value);
      break;
    case 6:
      var values = /** @type {!Array<number>} */ (reader.isDelimited() ? reader.readPackedUint64() : [reader.readUint64()]);
      for (var i = 0; i < values.length; i++) {
        msg.addIn(values[i]);
      }
      break;
    case 7:
      var values = /** @type {!Array<number>} */ (reader.isDelimited() ? reader.readPackedUint64() : [reader.readUint64()]);
      for (var i = 0; i < values.length; i++) {
        msg.addNotIn(values[i]);
      }
      break;
    case 8:
      var values = /** @type {!Array<number>} */ (reader.isDelimited() ? reader.readPackedUint64() : [reader.readUint64()]);
      for (var i = 0; i < values.length; i++) {
        msg.addExample(values[i]);
      }
      break;
    default:
      jspb.Message.readBinaryExtension(msg, reader,
        proto.buf.validate.UInt64Rules.extensionsBinary,
        proto.buf.validate.UInt64Rules.prototype.getExtension,
        proto.buf.validate.UInt64Rules.prototype.setExtension);
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.buf.validate.UInt64Rules.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.buf.validate.UInt64Rules.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.buf.validate.UInt64Rules} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.buf.validate.UInt64Rules.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = /** @type {number} */ (jspb.Message.getField(message, 1));
  if (f != null) {
    writer.writeUint64(
      1,
      f
    );
  }
  f = /** @type {number} */ (jspb.Message.getField(message, 2));
  if (f != null) {
    writer.writeUint64(
      2,
      f
    );
  }
  f = /** @type {number} */ (jspb.Message.getField(message, 3));
  if (f != null) {
    writer.writeUint64(
      3,
      f
    );
  }
  f = /** @type {number} */ (jspb.Message.getField(message, 4));
  if (f != null) {
    writer.writeUint64(
      4,
      f
    );
  }
  f = /** @type {number} */ (jspb.Message.getField(message, 5));
  if (f != null) {
    writer.writeUint64(
      5,
      f
    );
  }
  f = message.getInList();
  if (f.length > 0) {
    writer.writeRepeatedUint64(
      6,
      f
    );
  }
  f = message.getNotInList();
  if (f.length > 0) {
    writer.writeRepeatedUint64(
      7,
      f
    );
  }
  f = message.getExampleList();
  if (f.length > 0) {
    writer.writeRepeatedUint64(
      8,
      f
    );
  }
  jspb.Message.serializeBinaryExtensions(message, writer,
    proto.buf.validate.UInt64Rules.extensionsBinary, proto.buf.validate.UInt64Rules.prototype.getExtension);
};


/**
 * optional uint64 const = 1;
 * @return {number}
 */
proto.buf.validate.UInt64Rules.prototype.getConst = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/**
 * @param {number} value
 * @return {!proto.buf.validate.UInt64Rules} returns this
 */
proto.buf.validate.UInt64Rules.prototype.setConst = function(value) {
  return jspb.Message.setField(this, 1, value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.UInt64Rules} returns this
 */
proto.buf.validate.UInt64Rules.prototype.clearConst = function() {
  return jspb.Message.setField(this, 1, undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.UInt64Rules.prototype.hasConst = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional uint64 lt = 2;
 * @return {number}
 */
proto.buf.validate.UInt64Rules.prototype.getLt = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/**
 * @param {number} value
 * @return {!proto.buf.validate.UInt64Rules} returns this
 */
proto.buf.validate.UInt64Rules.prototype.setLt = function(value) {
  return jspb.Message.setOneofField(this, 2, proto.buf.validate.UInt64Rules.oneofGroups_[0], value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.UInt64Rules} returns this
 */
proto.buf.validate.UInt64Rules.prototype.clearLt = function() {
  return jspb.Message.setOneofField(this, 2, proto.buf.validate.UInt64Rules.oneofGroups_[0], undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.UInt64Rules.prototype.hasLt = function() {
  return jspb.Message.getField(this, 2) != null;
};


/**
 * optional uint64 lte = 3;
 * @return {number}
 */
proto.buf.validate.UInt64Rules.prototype.getLte = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 3, 0));
};


/**
 * @param {number} value
 * @return {!proto.buf.validate.UInt64Rules} returns this
 */
proto.buf.validate.UInt64Rules.prototype.setLte = function(value) {
  return jspb.Message.setOneofField(this, 3, proto.buf.validate.UInt64Rules.oneofGroups_[0], value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.UInt64Rules} returns this
 */
proto.buf.validate.UInt64Rules.prototype.clearLte = function() {
  return jspb.Message.setOneofField(this, 3, proto.buf.validate.UInt64Rules.oneofGroups_[0], undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.UInt64Rules.prototype.hasLte = function() {
  return jspb.Message.getField(this, 3) != null;
};


/**
 * optional uint64 gt = 4;
 * @return {number}
 */
proto.buf.validate.UInt64Rules.prototype.getGt = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 4, 0));
};


/**
 * @param {number} value
 * @return {!proto.buf.validate.UInt64Rules} returns this
 */
proto.buf.validate.UInt64Rules.prototype.setGt = function(value) {
  return jspb.Message.setOneofField(this, 4, proto.buf.validate.UInt64Rules.oneofGroups_[1], value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.UInt64Rules} returns this
 */
proto.buf.validate.UInt64Rules.prototype.clearGt = function() {
  return jspb.Message.setOneofField(this, 4, proto.buf.validate.UInt64Rules.oneofGroups_[1], undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.UInt64Rules.prototype.hasGt = function() {
  return jspb.Message.getField(this, 4) != null;
};


/**
 * optional uint64 gte = 5;
 * @return {number}
 */
proto.buf.validate.UInt64Rules.prototype.getGte = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 5, 0));
};


/**
 * @param {number} value
 * @return {!proto.buf.validate.UInt64Rules} returns this
 */
proto.buf.validate.UInt64Rules.prototype.setGte = function(value) {
  return jspb.Message.setOneofField(this, 5, proto.buf.validate.UInt64Rules.oneofGroups_[1], value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.UInt64Rules} returns this
 */
proto.buf.validate.UInt64Rules.prototype.clearGte = function() {
  return jspb.Message.setOneofField(this, 5, proto.buf.validate.UInt64Rules.oneofGroups_[1], undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.UInt64Rules.prototype.hasGte = function() {
  return jspb.Message.getField(this, 5) != null;
};


/**
 * repeated uint64 in = 6;
 * @return {!Array<number>}
 */
proto.buf.validate.UInt64Rules.prototype.getInList = function() {
  return /** @type {!Array<number>} */ (jspb.Message.getRepeatedField(this, 6));
};


/**
 * @param {!Array<number>} value
 * @return {!proto.buf.validate.UInt64Rules} returns this
 */
proto.buf.validate.UInt64Rules.prototype.setInList = function(value) {
  return jspb.Message.setField(this, 6, value || []);
};


/**
 * @param {number} value
 * @param {number=} opt_index
 * @return {!proto.buf.validate.UInt64Rules} returns this
 */
proto.buf.validate.UInt64Rules.prototype.addIn = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 6, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.buf.validate.UInt64Rules} returns this
 */
proto.buf.validate.UInt64Rules.prototype.clearInList = function() {
  return this.setInList([]);
};


/**
 * repeated uint64 not_in = 7;
 * @return {!Array<number>}
 */
proto.buf.validate.UInt64Rules.prototype.getNotInList = function() {
  return /** @type {!Array<number>} */ (jspb.Message.getRepeatedField(this, 7));
};


/**
 * @param {!Array<number>} value
 * @return {!proto.buf.validate.UInt64Rules} returns this
 */
proto.buf.validate.UInt64Rules.prototype.setNotInList = function(value) {
  return jspb.Message.setField(this, 7, value || []);
};


/**
 * @param {number} value
 * @param {number=} opt_index
 * @return {!proto.buf.validate.UInt64Rules} returns this
 */
proto.buf.validate.UInt64Rules.prototype.addNotIn = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 7, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.buf.validate.UInt64Rules} returns this
 */
proto.buf.validate.UInt64Rules.prototype.clearNotInList = function() {
  return this.setNotInList([]);
};


/**
 * repeated uint64 example = 8;
 * @return {!Array<number>}
 */
proto.buf.validate.UInt64Rules.prototype.getExampleList = function() {
  return /** @type {!Array<number>} */ (jspb.Message.getRepeatedField(this, 8));
};


/**
 * @param {!Array<number>} value
 * @return {!proto.buf.validate.UInt64Rules} returns this
 */
proto.buf.validate.UInt64Rules.prototype.setExampleList = function(value) {
  return jspb.Message.setField(this, 8, value || []);
};


/**
 * @param {number} value
 * @param {number=} opt_index
 * @return {!proto.buf.validate.UInt64Rules} returns this
 */
proto.buf.validate.UInt64Rules.prototype.addExample = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 8, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.buf.validate.UInt64Rules} returns this
 */
proto.buf.validate.UInt64Rules.prototype.clearExampleList = function() {
  return this.setExampleList([]);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.buf.validate.SInt32Rules.repeatedFields_ = [6,7,8];

/**
 * Oneof group definitions for this message. Each group defines the field
 * numbers belonging to that group. When of these fields' value is set, all
 * other fields in the group are cleared. During deserialization, if multiple
 * fields are encountered for a group, only the last value seen will be kept.
 * @private {!Array<!Array<number>>}
 * @const
 */
proto.buf.validate.SInt32Rules.oneofGroups_ = [[2,3],[4,5]];

/**
 * @enum {number}
 */
proto.buf.validate.SInt32Rules.LessThanCase = {
  LESS_THAN_NOT_SET: 0,
  LT: 2,
  LTE: 3
};

/**
 * @return {proto.buf.validate.SInt32Rules.LessThanCase}
 */
proto.buf.validate.SInt32Rules.prototype.getLessThanCase = function() {
  return /** @type {proto.buf.validate.SInt32Rules.LessThanCase} */(jspb.Message.computeOneofCase(this, proto.buf.validate.SInt32Rules.oneofGroups_[0]));
};

/**
 * @enum {number}
 */
proto.buf.validate.SInt32Rules.GreaterThanCase = {
  GREATER_THAN_NOT_SET: 0,
  GT: 4,
  GTE: 5
};

/**
 * @return {proto.buf.validate.SInt32Rules.GreaterThanCase}
 */
proto.buf.validate.SInt32Rules.prototype.getGreaterThanCase = function() {
  return /** @type {proto.buf.validate.SInt32Rules.GreaterThanCase} */(jspb.Message.computeOneofCase(this, proto.buf.validate.SInt32Rules.oneofGroups_[1]));
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
proto.buf.validate.SInt32Rules.prototype.toObject = function(opt_includeInstance) {
  return proto.buf.validate.SInt32Rules.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.buf.validate.SInt32Rules} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.buf.validate.SInt32Rules.toObject = function(includeInstance, msg) {
  var f, obj = {
pb_const: (f = jspb.Message.getField(msg, 1)) == null ? undefined : f,
lt: (f = jspb.Message.getField(msg, 2)) == null ? undefined : f,
lte: (f = jspb.Message.getField(msg, 3)) == null ? undefined : f,
gt: (f = jspb.Message.getField(msg, 4)) == null ? undefined : f,
gte: (f = jspb.Message.getField(msg, 5)) == null ? undefined : f,
inList: (f = jspb.Message.getRepeatedField(msg, 6)) == null ? undefined : f,
notInList: (f = jspb.Message.getRepeatedField(msg, 7)) == null ? undefined : f,
exampleList: (f = jspb.Message.getRepeatedField(msg, 8)) == null ? undefined : f
  };

  jspb.Message.toObjectExtension(/** @type {!jspb.Message} */ (msg), obj,
      proto.buf.validate.SInt32Rules.extensions, proto.buf.validate.SInt32Rules.prototype.getExtension,
      includeInstance);
  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.buf.validate.SInt32Rules}
 */
proto.buf.validate.SInt32Rules.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.buf.validate.SInt32Rules;
  return proto.buf.validate.SInt32Rules.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.buf.validate.SInt32Rules} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.buf.validate.SInt32Rules}
 */
proto.buf.validate.SInt32Rules.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readSint32());
      msg.setConst(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readSint32());
      msg.setLt(value);
      break;
    case 3:
      var value = /** @type {number} */ (reader.readSint32());
      msg.setLte(value);
      break;
    case 4:
      var value = /** @type {number} */ (reader.readSint32());
      msg.setGt(value);
      break;
    case 5:
      var value = /** @type {number} */ (reader.readSint32());
      msg.setGte(value);
      break;
    case 6:
      var values = /** @type {!Array<number>} */ (reader.isDelimited() ? reader.readPackedSint32() : [reader.readSint32()]);
      for (var i = 0; i < values.length; i++) {
        msg.addIn(values[i]);
      }
      break;
    case 7:
      var values = /** @type {!Array<number>} */ (reader.isDelimited() ? reader.readPackedSint32() : [reader.readSint32()]);
      for (var i = 0; i < values.length; i++) {
        msg.addNotIn(values[i]);
      }
      break;
    case 8:
      var values = /** @type {!Array<number>} */ (reader.isDelimited() ? reader.readPackedSint32() : [reader.readSint32()]);
      for (var i = 0; i < values.length; i++) {
        msg.addExample(values[i]);
      }
      break;
    default:
      jspb.Message.readBinaryExtension(msg, reader,
        proto.buf.validate.SInt32Rules.extensionsBinary,
        proto.buf.validate.SInt32Rules.prototype.getExtension,
        proto.buf.validate.SInt32Rules.prototype.setExtension);
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.buf.validate.SInt32Rules.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.buf.validate.SInt32Rules.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.buf.validate.SInt32Rules} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.buf.validate.SInt32Rules.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = /** @type {number} */ (jspb.Message.getField(message, 1));
  if (f != null) {
    writer.writeSint32(
      1,
      f
    );
  }
  f = /** @type {number} */ (jspb.Message.getField(message, 2));
  if (f != null) {
    writer.writeSint32(
      2,
      f
    );
  }
  f = /** @type {number} */ (jspb.Message.getField(message, 3));
  if (f != null) {
    writer.writeSint32(
      3,
      f
    );
  }
  f = /** @type {number} */ (jspb.Message.getField(message, 4));
  if (f != null) {
    writer.writeSint32(
      4,
      f
    );
  }
  f = /** @type {number} */ (jspb.Message.getField(message, 5));
  if (f != null) {
    writer.writeSint32(
      5,
      f
    );
  }
  f = message.getInList();
  if (f.length > 0) {
    writer.writeRepeatedSint32(
      6,
      f
    );
  }
  f = message.getNotInList();
  if (f.length > 0) {
    writer.writeRepeatedSint32(
      7,
      f
    );
  }
  f = message.getExampleList();
  if (f.length > 0) {
    writer.writeRepeatedSint32(
      8,
      f
    );
  }
  jspb.Message.serializeBinaryExtensions(message, writer,
    proto.buf.validate.SInt32Rules.extensionsBinary, proto.buf.validate.SInt32Rules.prototype.getExtension);
};


/**
 * optional sint32 const = 1;
 * @return {number}
 */
proto.buf.validate.SInt32Rules.prototype.getConst = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/**
 * @param {number} value
 * @return {!proto.buf.validate.SInt32Rules} returns this
 */
proto.buf.validate.SInt32Rules.prototype.setConst = function(value) {
  return jspb.Message.setField(this, 1, value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.SInt32Rules} returns this
 */
proto.buf.validate.SInt32Rules.prototype.clearConst = function() {
  return jspb.Message.setField(this, 1, undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.SInt32Rules.prototype.hasConst = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional sint32 lt = 2;
 * @return {number}
 */
proto.buf.validate.SInt32Rules.prototype.getLt = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/**
 * @param {number} value
 * @return {!proto.buf.validate.SInt32Rules} returns this
 */
proto.buf.validate.SInt32Rules.prototype.setLt = function(value) {
  return jspb.Message.setOneofField(this, 2, proto.buf.validate.SInt32Rules.oneofGroups_[0], value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.SInt32Rules} returns this
 */
proto.buf.validate.SInt32Rules.prototype.clearLt = function() {
  return jspb.Message.setOneofField(this, 2, proto.buf.validate.SInt32Rules.oneofGroups_[0], undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.SInt32Rules.prototype.hasLt = function() {
  return jspb.Message.getField(this, 2) != null;
};


/**
 * optional sint32 lte = 3;
 * @return {number}
 */
proto.buf.validate.SInt32Rules.prototype.getLte = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 3, 0));
};


/**
 * @param {number} value
 * @return {!proto.buf.validate.SInt32Rules} returns this
 */
proto.buf.validate.SInt32Rules.prototype.setLte = function(value) {
  return jspb.Message.setOneofField(this, 3, proto.buf.validate.SInt32Rules.oneofGroups_[0], value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.SInt32Rules} returns this
 */
proto.buf.validate.SInt32Rules.prototype.clearLte = function() {
  return jspb.Message.setOneofField(this, 3, proto.buf.validate.SInt32Rules.oneofGroups_[0], undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.SInt32Rules.prototype.hasLte = function() {
  return jspb.Message.getField(this, 3) != null;
};


/**
 * optional sint32 gt = 4;
 * @return {number}
 */
proto.buf.validate.SInt32Rules.prototype.getGt = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 4, 0));
};


/**
 * @param {number} value
 * @return {!proto.buf.validate.SInt32Rules} returns this
 */
proto.buf.validate.SInt32Rules.prototype.setGt = function(value) {
  return jspb.Message.setOneofField(this, 4, proto.buf.validate.SInt32Rules.oneofGroups_[1], value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.SInt32Rules} returns this
 */
proto.buf.validate.SInt32Rules.prototype.clearGt = function() {
  return jspb.Message.setOneofField(this, 4, proto.buf.validate.SInt32Rules.oneofGroups_[1], undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.SInt32Rules.prototype.hasGt = function() {
  return jspb.Message.getField(this, 4) != null;
};


/**
 * optional sint32 gte = 5;
 * @return {number}
 */
proto.buf.validate.SInt32Rules.prototype.getGte = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 5, 0));
};


/**
 * @param {number} value
 * @return {!proto.buf.validate.SInt32Rules} returns this
 */
proto.buf.validate.SInt32Rules.prototype.setGte = function(value) {
  return jspb.Message.setOneofField(this, 5, proto.buf.validate.SInt32Rules.oneofGroups_[1], value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.SInt32Rules} returns this
 */
proto.buf.validate.SInt32Rules.prototype.clearGte = function() {
  return jspb.Message.setOneofField(this, 5, proto.buf.validate.SInt32Rules.oneofGroups_[1], undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.SInt32Rules.prototype.hasGte = function() {
  return jspb.Message.getField(this, 5) != null;
};


/**
 * repeated sint32 in = 6;
 * @return {!Array<number>}
 */
proto.buf.validate.SInt32Rules.prototype.getInList = function() {
  return /** @type {!Array<number>} */ (jspb.Message.getRepeatedField(this, 6));
};


/**
 * @param {!Array<number>} value
 * @return {!proto.buf.validate.SInt32Rules} returns this
 */
proto.buf.validate.SInt32Rules.prototype.setInList = function(value) {
  return jspb.Message.setField(this, 6, value || []);
};


/**
 * @param {number} value
 * @param {number=} opt_index
 * @return {!proto.buf.validate.SInt32Rules} returns this
 */
proto.buf.validate.SInt32Rules.prototype.addIn = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 6, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.buf.validate.SInt32Rules} returns this
 */
proto.buf.validate.SInt32Rules.prototype.clearInList = function() {
  return this.setInList([]);
};


/**
 * repeated sint32 not_in = 7;
 * @return {!Array<number>}
 */
proto.buf.validate.SInt32Rules.prototype.getNotInList = function() {
  return /** @type {!Array<number>} */ (jspb.Message.getRepeatedField(this, 7));
};


/**
 * @param {!Array<number>} value
 * @return {!proto.buf.validate.SInt32Rules} returns this
 */
proto.buf.validate.SInt32Rules.prototype.setNotInList = function(value) {
  return jspb.Message.setField(this, 7, value || []);
};


/**
 * @param {number} value
 * @param {number=} opt_index
 * @return {!proto.buf.validate.SInt32Rules} returns this
 */
proto.buf.validate.SInt32Rules.prototype.addNotIn = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 7, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.buf.validate.SInt32Rules} returns this
 */
proto.buf.validate.SInt32Rules.prototype.clearNotInList = function() {
  return this.setNotInList([]);
};


/**
 * repeated sint32 example = 8;
 * @return {!Array<number>}
 */
proto.buf.validate.SInt32Rules.prototype.getExampleList = function() {
  return /** @type {!Array<number>} */ (jspb.Message.getRepeatedField(this, 8));
};


/**
 * @param {!Array<number>} value
 * @return {!proto.buf.validate.SInt32Rules} returns this
 */
proto.buf.validate.SInt32Rules.prototype.setExampleList = function(value) {
  return jspb.Message.setField(this, 8, value || []);
};


/**
 * @param {number} value
 * @param {number=} opt_index
 * @return {!proto.buf.validate.SInt32Rules} returns this
 */
proto.buf.validate.SInt32Rules.prototype.addExample = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 8, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.buf.validate.SInt32Rules} returns this
 */
proto.buf.validate.SInt32Rules.prototype.clearExampleList = function() {
  return this.setExampleList([]);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.buf.validate.SInt64Rules.repeatedFields_ = [6,7,8];

/**
 * Oneof group definitions for this message. Each group defines the field
 * numbers belonging to that group. When of these fields' value is set, all
 * other fields in the group are cleared. During deserialization, if multiple
 * fields are encountered for a group, only the last value seen will be kept.
 * @private {!Array<!Array<number>>}
 * @const
 */
proto.buf.validate.SInt64Rules.oneofGroups_ = [[2,3],[4,5]];

/**
 * @enum {number}
 */
proto.buf.validate.SInt64Rules.LessThanCase = {
  LESS_THAN_NOT_SET: 0,
  LT: 2,
  LTE: 3
};

/**
 * @return {proto.buf.validate.SInt64Rules.LessThanCase}
 */
proto.buf.validate.SInt64Rules.prototype.getLessThanCase = function() {
  return /** @type {proto.buf.validate.SInt64Rules.LessThanCase} */(jspb.Message.computeOneofCase(this, proto.buf.validate.SInt64Rules.oneofGroups_[0]));
};

/**
 * @enum {number}
 */
proto.buf.validate.SInt64Rules.GreaterThanCase = {
  GREATER_THAN_NOT_SET: 0,
  GT: 4,
  GTE: 5
};

/**
 * @return {proto.buf.validate.SInt64Rules.GreaterThanCase}
 */
proto.buf.validate.SInt64Rules.prototype.getGreaterThanCase = function() {
  return /** @type {proto.buf.validate.SInt64Rules.GreaterThanCase} */(jspb.Message.computeOneofCase(this, proto.buf.validate.SInt64Rules.oneofGroups_[1]));
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
proto.buf.validate.SInt64Rules.prototype.toObject = function(opt_includeInstance) {
  return proto.buf.validate.SInt64Rules.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.buf.validate.SInt64Rules} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.buf.validate.SInt64Rules.toObject = function(includeInstance, msg) {
  var f, obj = {
pb_const: (f = jspb.Message.getField(msg, 1)) == null ? undefined : f,
lt: (f = jspb.Message.getField(msg, 2)) == null ? undefined : f,
lte: (f = jspb.Message.getField(msg, 3)) == null ? undefined : f,
gt: (f = jspb.Message.getField(msg, 4)) == null ? undefined : f,
gte: (f = jspb.Message.getField(msg, 5)) == null ? undefined : f,
inList: (f = jspb.Message.getRepeatedField(msg, 6)) == null ? undefined : f,
notInList: (f = jspb.Message.getRepeatedField(msg, 7)) == null ? undefined : f,
exampleList: (f = jspb.Message.getRepeatedField(msg, 8)) == null ? undefined : f
  };

  jspb.Message.toObjectExtension(/** @type {!jspb.Message} */ (msg), obj,
      proto.buf.validate.SInt64Rules.extensions, proto.buf.validate.SInt64Rules.prototype.getExtension,
      includeInstance);
  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.buf.validate.SInt64Rules}
 */
proto.buf.validate.SInt64Rules.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.buf.validate.SInt64Rules;
  return proto.buf.validate.SInt64Rules.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.buf.validate.SInt64Rules} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.buf.validate.SInt64Rules}
 */
proto.buf.validate.SInt64Rules.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readSint64());
      msg.setConst(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readSint64());
      msg.setLt(value);
      break;
    case 3:
      var value = /** @type {number} */ (reader.readSint64());
      msg.setLte(value);
      break;
    case 4:
      var value = /** @type {number} */ (reader.readSint64());
      msg.setGt(value);
      break;
    case 5:
      var value = /** @type {number} */ (reader.readSint64());
      msg.setGte(value);
      break;
    case 6:
      var values = /** @type {!Array<number>} */ (reader.isDelimited() ? reader.readPackedSint64() : [reader.readSint64()]);
      for (var i = 0; i < values.length; i++) {
        msg.addIn(values[i]);
      }
      break;
    case 7:
      var values = /** @type {!Array<number>} */ (reader.isDelimited() ? reader.readPackedSint64() : [reader.readSint64()]);
      for (var i = 0; i < values.length; i++) {
        msg.addNotIn(values[i]);
      }
      break;
    case 8:
      var values = /** @type {!Array<number>} */ (reader.isDelimited() ? reader.readPackedSint64() : [reader.readSint64()]);
      for (var i = 0; i < values.length; i++) {
        msg.addExample(values[i]);
      }
      break;
    default:
      jspb.Message.readBinaryExtension(msg, reader,
        proto.buf.validate.SInt64Rules.extensionsBinary,
        proto.buf.validate.SInt64Rules.prototype.getExtension,
        proto.buf.validate.SInt64Rules.prototype.setExtension);
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.buf.validate.SInt64Rules.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.buf.validate.SInt64Rules.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.buf.validate.SInt64Rules} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.buf.validate.SInt64Rules.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = /** @type {number} */ (jspb.Message.getField(message, 1));
  if (f != null) {
    writer.writeSint64(
      1,
      f
    );
  }
  f = /** @type {number} */ (jspb.Message.getField(message, 2));
  if (f != null) {
    writer.writeSint64(
      2,
      f
    );
  }
  f = /** @type {number} */ (jspb.Message.getField(message, 3));
  if (f != null) {
    writer.writeSint64(
      3,
      f
    );
  }
  f = /** @type {number} */ (jspb.Message.getField(message, 4));
  if (f != null) {
    writer.writeSint64(
      4,
      f
    );
  }
  f = /** @type {number} */ (jspb.Message.getField(message, 5));
  if (f != null) {
    writer.writeSint64(
      5,
      f
    );
  }
  f = message.getInList();
  if (f.length > 0) {
    writer.writeRepeatedSint64(
      6,
      f
    );
  }
  f = message.getNotInList();
  if (f.length > 0) {
    writer.writeRepeatedSint64(
      7,
      f
    );
  }
  f = message.getExampleList();
  if (f.length > 0) {
    writer.writeRepeatedSint64(
      8,
      f
    );
  }
  jspb.Message.serializeBinaryExtensions(message, writer,
    proto.buf.validate.SInt64Rules.extensionsBinary, proto.buf.validate.SInt64Rules.prototype.getExtension);
};


/**
 * optional sint64 const = 1;
 * @return {number}
 */
proto.buf.validate.SInt64Rules.prototype.getConst = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/**
 * @param {number} value
 * @return {!proto.buf.validate.SInt64Rules} returns this
 */
proto.buf.validate.SInt64Rules.prototype.setConst = function(value) {
  return jspb.Message.setField(this, 1, value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.SInt64Rules} returns this
 */
proto.buf.validate.SInt64Rules.prototype.clearConst = function() {
  return jspb.Message.setField(this, 1, undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.SInt64Rules.prototype.hasConst = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional sint64 lt = 2;
 * @return {number}
 */
proto.buf.validate.SInt64Rules.prototype.getLt = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/**
 * @param {number} value
 * @return {!proto.buf.validate.SInt64Rules} returns this
 */
proto.buf.validate.SInt64Rules.prototype.setLt = function(value) {
  return jspb.Message.setOneofField(this, 2, proto.buf.validate.SInt64Rules.oneofGroups_[0], value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.SInt64Rules} returns this
 */
proto.buf.validate.SInt64Rules.prototype.clearLt = function() {
  return jspb.Message.setOneofField(this, 2, proto.buf.validate.SInt64Rules.oneofGroups_[0], undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.SInt64Rules.prototype.hasLt = function() {
  return jspb.Message.getField(this, 2) != null;
};


/**
 * optional sint64 lte = 3;
 * @return {number}
 */
proto.buf.validate.SInt64Rules.prototype.getLte = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 3, 0));
};


/**
 * @param {number} value
 * @return {!proto.buf.validate.SInt64Rules} returns this
 */
proto.buf.validate.SInt64Rules.prototype.setLte = function(value) {
  return jspb.Message.setOneofField(this, 3, proto.buf.validate.SInt64Rules.oneofGroups_[0], value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.SInt64Rules} returns this
 */
proto.buf.validate.SInt64Rules.prototype.clearLte = function() {
  return jspb.Message.setOneofField(this, 3, proto.buf.validate.SInt64Rules.oneofGroups_[0], undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.SInt64Rules.prototype.hasLte = function() {
  return jspb.Message.getField(this, 3) != null;
};


/**
 * optional sint64 gt = 4;
 * @return {number}
 */
proto.buf.validate.SInt64Rules.prototype.getGt = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 4, 0));
};


/**
 * @param {number} value
 * @return {!proto.buf.validate.SInt64Rules} returns this
 */
proto.buf.validate.SInt64Rules.prototype.setGt = function(value) {
  return jspb.Message.setOneofField(this, 4, proto.buf.validate.SInt64Rules.oneofGroups_[1], value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.SInt64Rules} returns this
 */
proto.buf.validate.SInt64Rules.prototype.clearGt = function() {
  return jspb.Message.setOneofField(this, 4, proto.buf.validate.SInt64Rules.oneofGroups_[1], undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.SInt64Rules.prototype.hasGt = function() {
  return jspb.Message.getField(this, 4) != null;
};


/**
 * optional sint64 gte = 5;
 * @return {number}
 */
proto.buf.validate.SInt64Rules.prototype.getGte = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 5, 0));
};


/**
 * @param {number} value
 * @return {!proto.buf.validate.SInt64Rules} returns this
 */
proto.buf.validate.SInt64Rules.prototype.setGte = function(value) {
  return jspb.Message.setOneofField(this, 5, proto.buf.validate.SInt64Rules.oneofGroups_[1], value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.SInt64Rules} returns this
 */
proto.buf.validate.SInt64Rules.prototype.clearGte = function() {
  return jspb.Message.setOneofField(this, 5, proto.buf.validate.SInt64Rules.oneofGroups_[1], undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.SInt64Rules.prototype.hasGte = function() {
  return jspb.Message.getField(this, 5) != null;
};


/**
 * repeated sint64 in = 6;
 * @return {!Array<number>}
 */
proto.buf.validate.SInt64Rules.prototype.getInList = function() {
  return /** @type {!Array<number>} */ (jspb.Message.getRepeatedField(this, 6));
};


/**
 * @param {!Array<number>} value
 * @return {!proto.buf.validate.SInt64Rules} returns this
 */
proto.buf.validate.SInt64Rules.prototype.setInList = function(value) {
  return jspb.Message.setField(this, 6, value || []);
};


/**
 * @param {number} value
 * @param {number=} opt_index
 * @return {!proto.buf.validate.SInt64Rules} returns this
 */
proto.buf.validate.SInt64Rules.prototype.addIn = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 6, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.buf.validate.SInt64Rules} returns this
 */
proto.buf.validate.SInt64Rules.prototype.clearInList = function() {
  return this.setInList([]);
};


/**
 * repeated sint64 not_in = 7;
 * @return {!Array<number>}
 */
proto.buf.validate.SInt64Rules.prototype.getNotInList = function() {
  return /** @type {!Array<number>} */ (jspb.Message.getRepeatedField(this, 7));
};


/**
 * @param {!Array<number>} value
 * @return {!proto.buf.validate.SInt64Rules} returns this
 */
proto.buf.validate.SInt64Rules.prototype.setNotInList = function(value) {
  return jspb.Message.setField(this, 7, value || []);
};


/**
 * @param {number} value
 * @param {number=} opt_index
 * @return {!proto.buf.validate.SInt64Rules} returns this
 */
proto.buf.validate.SInt64Rules.prototype.addNotIn = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 7, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.buf.validate.SInt64Rules} returns this
 */
proto.buf.validate.SInt64Rules.prototype.clearNotInList = function() {
  return this.setNotInList([]);
};


/**
 * repeated sint64 example = 8;
 * @return {!Array<number>}
 */
proto.buf.validate.SInt64Rules.prototype.getExampleList = function() {
  return /** @type {!Array<number>} */ (jspb.Message.getRepeatedField(this, 8));
};


/**
 * @param {!Array<number>} value
 * @return {!proto.buf.validate.SInt64Rules} returns this
 */
proto.buf.validate.SInt64Rules.prototype.setExampleList = function(value) {
  return jspb.Message.setField(this, 8, value || []);
};


/**
 * @param {number} value
 * @param {number=} opt_index
 * @return {!proto.buf.validate.SInt64Rules} returns this
 */
proto.buf.validate.SInt64Rules.prototype.addExample = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 8, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.buf.validate.SInt64Rules} returns this
 */
proto.buf.validate.SInt64Rules.prototype.clearExampleList = function() {
  return this.setExampleList([]);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.buf.validate.Fixed32Rules.repeatedFields_ = [6,7,8];

/**
 * Oneof group definitions for this message. Each group defines the field
 * numbers belonging to that group. When of these fields' value is set, all
 * other fields in the group are cleared. During deserialization, if multiple
 * fields are encountered for a group, only the last value seen will be kept.
 * @private {!Array<!Array<number>>}
 * @const
 */
proto.buf.validate.Fixed32Rules.oneofGroups_ = [[2,3],[4,5]];

/**
 * @enum {number}
 */
proto.buf.validate.Fixed32Rules.LessThanCase = {
  LESS_THAN_NOT_SET: 0,
  LT: 2,
  LTE: 3
};

/**
 * @return {proto.buf.validate.Fixed32Rules.LessThanCase}
 */
proto.buf.validate.Fixed32Rules.prototype.getLessThanCase = function() {
  return /** @type {proto.buf.validate.Fixed32Rules.LessThanCase} */(jspb.Message.computeOneofCase(this, proto.buf.validate.Fixed32Rules.oneofGroups_[0]));
};

/**
 * @enum {number}
 */
proto.buf.validate.Fixed32Rules.GreaterThanCase = {
  GREATER_THAN_NOT_SET: 0,
  GT: 4,
  GTE: 5
};

/**
 * @return {proto.buf.validate.Fixed32Rules.GreaterThanCase}
 */
proto.buf.validate.Fixed32Rules.prototype.getGreaterThanCase = function() {
  return /** @type {proto.buf.validate.Fixed32Rules.GreaterThanCase} */(jspb.Message.computeOneofCase(this, proto.buf.validate.Fixed32Rules.oneofGroups_[1]));
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
proto.buf.validate.Fixed32Rules.prototype.toObject = function(opt_includeInstance) {
  return proto.buf.validate.Fixed32Rules.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.buf.validate.Fixed32Rules} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.buf.validate.Fixed32Rules.toObject = function(includeInstance, msg) {
  var f, obj = {
pb_const: (f = jspb.Message.getField(msg, 1)) == null ? undefined : f,
lt: (f = jspb.Message.getField(msg, 2)) == null ? undefined : f,
lte: (f = jspb.Message.getField(msg, 3)) == null ? undefined : f,
gt: (f = jspb.Message.getField(msg, 4)) == null ? undefined : f,
gte: (f = jspb.Message.getField(msg, 5)) == null ? undefined : f,
inList: (f = jspb.Message.getRepeatedField(msg, 6)) == null ? undefined : f,
notInList: (f = jspb.Message.getRepeatedField(msg, 7)) == null ? undefined : f,
exampleList: (f = jspb.Message.getRepeatedField(msg, 8)) == null ? undefined : f
  };

  jspb.Message.toObjectExtension(/** @type {!jspb.Message} */ (msg), obj,
      proto.buf.validate.Fixed32Rules.extensions, proto.buf.validate.Fixed32Rules.prototype.getExtension,
      includeInstance);
  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.buf.validate.Fixed32Rules}
 */
proto.buf.validate.Fixed32Rules.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.buf.validate.Fixed32Rules;
  return proto.buf.validate.Fixed32Rules.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.buf.validate.Fixed32Rules} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.buf.validate.Fixed32Rules}
 */
proto.buf.validate.Fixed32Rules.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readFixed32());
      msg.setConst(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readFixed32());
      msg.setLt(value);
      break;
    case 3:
      var value = /** @type {number} */ (reader.readFixed32());
      msg.setLte(value);
      break;
    case 4:
      var value = /** @type {number} */ (reader.readFixed32());
      msg.setGt(value);
      break;
    case 5:
      var value = /** @type {number} */ (reader.readFixed32());
      msg.setGte(value);
      break;
    case 6:
      var values = /** @type {!Array<number>} */ (reader.isDelimited() ? reader.readPackedFixed32() : [reader.readFixed32()]);
      for (var i = 0; i < values.length; i++) {
        msg.addIn(values[i]);
      }
      break;
    case 7:
      var values = /** @type {!Array<number>} */ (reader.isDelimited() ? reader.readPackedFixed32() : [reader.readFixed32()]);
      for (var i = 0; i < values.length; i++) {
        msg.addNotIn(values[i]);
      }
      break;
    case 8:
      var values = /** @type {!Array<number>} */ (reader.isDelimited() ? reader.readPackedFixed32() : [reader.readFixed32()]);
      for (var i = 0; i < values.length; i++) {
        msg.addExample(values[i]);
      }
      break;
    default:
      jspb.Message.readBinaryExtension(msg, reader,
        proto.buf.validate.Fixed32Rules.extensionsBinary,
        proto.buf.validate.Fixed32Rules.prototype.getExtension,
        proto.buf.validate.Fixed32Rules.prototype.setExtension);
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.buf.validate.Fixed32Rules.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.buf.validate.Fixed32Rules.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.buf.validate.Fixed32Rules} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.buf.validate.Fixed32Rules.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = /** @type {number} */ (jspb.Message.getField(message, 1));
  if (f != null) {
    writer.writeFixed32(
      1,
      f
    );
  }
  f = /** @type {number} */ (jspb.Message.getField(message, 2));
  if (f != null) {
    writer.writeFixed32(
      2,
      f
    );
  }
  f = /** @type {number} */ (jspb.Message.getField(message, 3));
  if (f != null) {
    writer.writeFixed32(
      3,
      f
    );
  }
  f = /** @type {number} */ (jspb.Message.getField(message, 4));
  if (f != null) {
    writer.writeFixed32(
      4,
      f
    );
  }
  f = /** @type {number} */ (jspb.Message.getField(message, 5));
  if (f != null) {
    writer.writeFixed32(
      5,
      f
    );
  }
  f = message.getInList();
  if (f.length > 0) {
    writer.writeRepeatedFixed32(
      6,
      f
    );
  }
  f = message.getNotInList();
  if (f.length > 0) {
    writer.writeRepeatedFixed32(
      7,
      f
    );
  }
  f = message.getExampleList();
  if (f.length > 0) {
    writer.writeRepeatedFixed32(
      8,
      f
    );
  }
  jspb.Message.serializeBinaryExtensions(message, writer,
    proto.buf.validate.Fixed32Rules.extensionsBinary, proto.buf.validate.Fixed32Rules.prototype.getExtension);
};


/**
 * optional fixed32 const = 1;
 * @return {number}
 */
proto.buf.validate.Fixed32Rules.prototype.getConst = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/**
 * @param {number} value
 * @return {!proto.buf.validate.Fixed32Rules} returns this
 */
proto.buf.validate.Fixed32Rules.prototype.setConst = function(value) {
  return jspb.Message.setField(this, 1, value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.Fixed32Rules} returns this
 */
proto.buf.validate.Fixed32Rules.prototype.clearConst = function() {
  return jspb.Message.setField(this, 1, undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.Fixed32Rules.prototype.hasConst = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional fixed32 lt = 2;
 * @return {number}
 */
proto.buf.validate.Fixed32Rules.prototype.getLt = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/**
 * @param {number} value
 * @return {!proto.buf.validate.Fixed32Rules} returns this
 */
proto.buf.validate.Fixed32Rules.prototype.setLt = function(value) {
  return jspb.Message.setOneofField(this, 2, proto.buf.validate.Fixed32Rules.oneofGroups_[0], value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.Fixed32Rules} returns this
 */
proto.buf.validate.Fixed32Rules.prototype.clearLt = function() {
  return jspb.Message.setOneofField(this, 2, proto.buf.validate.Fixed32Rules.oneofGroups_[0], undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.Fixed32Rules.prototype.hasLt = function() {
  return jspb.Message.getField(this, 2) != null;
};


/**
 * optional fixed32 lte = 3;
 * @return {number}
 */
proto.buf.validate.Fixed32Rules.prototype.getLte = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 3, 0));
};


/**
 * @param {number} value
 * @return {!proto.buf.validate.Fixed32Rules} returns this
 */
proto.buf.validate.Fixed32Rules.prototype.setLte = function(value) {
  return jspb.Message.setOneofField(this, 3, proto.buf.validate.Fixed32Rules.oneofGroups_[0], value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.Fixed32Rules} returns this
 */
proto.buf.validate.Fixed32Rules.prototype.clearLte = function() {
  return jspb.Message.setOneofField(this, 3, proto.buf.validate.Fixed32Rules.oneofGroups_[0], undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.Fixed32Rules.prototype.hasLte = function() {
  return jspb.Message.getField(this, 3) != null;
};


/**
 * optional fixed32 gt = 4;
 * @return {number}
 */
proto.buf.validate.Fixed32Rules.prototype.getGt = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 4, 0));
};


/**
 * @param {number} value
 * @return {!proto.buf.validate.Fixed32Rules} returns this
 */
proto.buf.validate.Fixed32Rules.prototype.setGt = function(value) {
  return jspb.Message.setOneofField(this, 4, proto.buf.validate.Fixed32Rules.oneofGroups_[1], value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.Fixed32Rules} returns this
 */
proto.buf.validate.Fixed32Rules.prototype.clearGt = function() {
  return jspb.Message.setOneofField(this, 4, proto.buf.validate.Fixed32Rules.oneofGroups_[1], undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.Fixed32Rules.prototype.hasGt = function() {
  return jspb.Message.getField(this, 4) != null;
};


/**
 * optional fixed32 gte = 5;
 * @return {number}
 */
proto.buf.validate.Fixed32Rules.prototype.getGte = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 5, 0));
};


/**
 * @param {number} value
 * @return {!proto.buf.validate.Fixed32Rules} returns this
 */
proto.buf.validate.Fixed32Rules.prototype.setGte = function(value) {
  return jspb.Message.setOneofField(this, 5, proto.buf.validate.Fixed32Rules.oneofGroups_[1], value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.Fixed32Rules} returns this
 */
proto.buf.validate.Fixed32Rules.prototype.clearGte = function() {
  return jspb.Message.setOneofField(this, 5, proto.buf.validate.Fixed32Rules.oneofGroups_[1], undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.Fixed32Rules.prototype.hasGte = function() {
  return jspb.Message.getField(this, 5) != null;
};


/**
 * repeated fixed32 in = 6;
 * @return {!Array<number>}
 */
proto.buf.validate.Fixed32Rules.prototype.getInList = function() {
  return /** @type {!Array<number>} */ (jspb.Message.getRepeatedField(this, 6));
};


/**
 * @param {!Array<number>} value
 * @return {!proto.buf.validate.Fixed32Rules} returns this
 */
proto.buf.validate.Fixed32Rules.prototype.setInList = function(value) {
  return jspb.Message.setField(this, 6, value || []);
};


/**
 * @param {number} value
 * @param {number=} opt_index
 * @return {!proto.buf.validate.Fixed32Rules} returns this
 */
proto.buf.validate.Fixed32Rules.prototype.addIn = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 6, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.buf.validate.Fixed32Rules} returns this
 */
proto.buf.validate.Fixed32Rules.prototype.clearInList = function() {
  return this.setInList([]);
};


/**
 * repeated fixed32 not_in = 7;
 * @return {!Array<number>}
 */
proto.buf.validate.Fixed32Rules.prototype.getNotInList = function() {
  return /** @type {!Array<number>} */ (jspb.Message.getRepeatedField(this, 7));
};


/**
 * @param {!Array<number>} value
 * @return {!proto.buf.validate.Fixed32Rules} returns this
 */
proto.buf.validate.Fixed32Rules.prototype.setNotInList = function(value) {
  return jspb.Message.setField(this, 7, value || []);
};


/**
 * @param {number} value
 * @param {number=} opt_index
 * @return {!proto.buf.validate.Fixed32Rules} returns this
 */
proto.buf.validate.Fixed32Rules.prototype.addNotIn = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 7, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.buf.validate.Fixed32Rules} returns this
 */
proto.buf.validate.Fixed32Rules.prototype.clearNotInList = function() {
  return this.setNotInList([]);
};


/**
 * repeated fixed32 example = 8;
 * @return {!Array<number>}
 */
proto.buf.validate.Fixed32Rules.prototype.getExampleList = function() {
  return /** @type {!Array<number>} */ (jspb.Message.getRepeatedField(this, 8));
};


/**
 * @param {!Array<number>} value
 * @return {!proto.buf.validate.Fixed32Rules} returns this
 */
proto.buf.validate.Fixed32Rules.prototype.setExampleList = function(value) {
  return jspb.Message.setField(this, 8, value || []);
};


/**
 * @param {number} value
 * @param {number=} opt_index
 * @return {!proto.buf.validate.Fixed32Rules} returns this
 */
proto.buf.validate.Fixed32Rules.prototype.addExample = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 8, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.buf.validate.Fixed32Rules} returns this
 */
proto.buf.validate.Fixed32Rules.prototype.clearExampleList = function() {
  return this.setExampleList([]);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.buf.validate.Fixed64Rules.repeatedFields_ = [6,7,8];

/**
 * Oneof group definitions for this message. Each group defines the field
 * numbers belonging to that group. When of these fields' value is set, all
 * other fields in the group are cleared. During deserialization, if multiple
 * fields are encountered for a group, only the last value seen will be kept.
 * @private {!Array<!Array<number>>}
 * @const
 */
proto.buf.validate.Fixed64Rules.oneofGroups_ = [[2,3],[4,5]];

/**
 * @enum {number}
 */
proto.buf.validate.Fixed64Rules.LessThanCase = {
  LESS_THAN_NOT_SET: 0,
  LT: 2,
  LTE: 3
};

/**
 * @return {proto.buf.validate.Fixed64Rules.LessThanCase}
 */
proto.buf.validate.Fixed64Rules.prototype.getLessThanCase = function() {
  return /** @type {proto.buf.validate.Fixed64Rules.LessThanCase} */(jspb.Message.computeOneofCase(this, proto.buf.validate.Fixed64Rules.oneofGroups_[0]));
};

/**
 * @enum {number}
 */
proto.buf.validate.Fixed64Rules.GreaterThanCase = {
  GREATER_THAN_NOT_SET: 0,
  GT: 4,
  GTE: 5
};

/**
 * @return {proto.buf.validate.Fixed64Rules.GreaterThanCase}
 */
proto.buf.validate.Fixed64Rules.prototype.getGreaterThanCase = function() {
  return /** @type {proto.buf.validate.Fixed64Rules.GreaterThanCase} */(jspb.Message.computeOneofCase(this, proto.buf.validate.Fixed64Rules.oneofGroups_[1]));
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
proto.buf.validate.Fixed64Rules.prototype.toObject = function(opt_includeInstance) {
  return proto.buf.validate.Fixed64Rules.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.buf.validate.Fixed64Rules} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.buf.validate.Fixed64Rules.toObject = function(includeInstance, msg) {
  var f, obj = {
pb_const: (f = jspb.Message.getField(msg, 1)) == null ? undefined : f,
lt: (f = jspb.Message.getField(msg, 2)) == null ? undefined : f,
lte: (f = jspb.Message.getField(msg, 3)) == null ? undefined : f,
gt: (f = jspb.Message.getField(msg, 4)) == null ? undefined : f,
gte: (f = jspb.Message.getField(msg, 5)) == null ? undefined : f,
inList: (f = jspb.Message.getRepeatedField(msg, 6)) == null ? undefined : f,
notInList: (f = jspb.Message.getRepeatedField(msg, 7)) == null ? undefined : f,
exampleList: (f = jspb.Message.getRepeatedField(msg, 8)) == null ? undefined : f
  };

  jspb.Message.toObjectExtension(/** @type {!jspb.Message} */ (msg), obj,
      proto.buf.validate.Fixed64Rules.extensions, proto.buf.validate.Fixed64Rules.prototype.getExtension,
      includeInstance);
  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.buf.validate.Fixed64Rules}
 */
proto.buf.validate.Fixed64Rules.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.buf.validate.Fixed64Rules;
  return proto.buf.validate.Fixed64Rules.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.buf.validate.Fixed64Rules} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.buf.validate.Fixed64Rules}
 */
proto.buf.validate.Fixed64Rules.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readFixed64());
      msg.setConst(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readFixed64());
      msg.setLt(value);
      break;
    case 3:
      var value = /** @type {number} */ (reader.readFixed64());
      msg.setLte(value);
      break;
    case 4:
      var value = /** @type {number} */ (reader.readFixed64());
      msg.setGt(value);
      break;
    case 5:
      var value = /** @type {number} */ (reader.readFixed64());
      msg.setGte(value);
      break;
    case 6:
      var values = /** @type {!Array<number>} */ (reader.isDelimited() ? reader.readPackedFixed64() : [reader.readFixed64()]);
      for (var i = 0; i < values.length; i++) {
        msg.addIn(values[i]);
      }
      break;
    case 7:
      var values = /** @type {!Array<number>} */ (reader.isDelimited() ? reader.readPackedFixed64() : [reader.readFixed64()]);
      for (var i = 0; i < values.length; i++) {
        msg.addNotIn(values[i]);
      }
      break;
    case 8:
      var values = /** @type {!Array<number>} */ (reader.isDelimited() ? reader.readPackedFixed64() : [reader.readFixed64()]);
      for (var i = 0; i < values.length; i++) {
        msg.addExample(values[i]);
      }
      break;
    default:
      jspb.Message.readBinaryExtension(msg, reader,
        proto.buf.validate.Fixed64Rules.extensionsBinary,
        proto.buf.validate.Fixed64Rules.prototype.getExtension,
        proto.buf.validate.Fixed64Rules.prototype.setExtension);
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.buf.validate.Fixed64Rules.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.buf.validate.Fixed64Rules.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.buf.validate.Fixed64Rules} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.buf.validate.Fixed64Rules.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = /** @type {number} */ (jspb.Message.getField(message, 1));
  if (f != null) {
    writer.writeFixed64(
      1,
      f
    );
  }
  f = /** @type {number} */ (jspb.Message.getField(message, 2));
  if (f != null) {
    writer.writeFixed64(
      2,
      f
    );
  }
  f = /** @type {number} */ (jspb.Message.getField(message, 3));
  if (f != null) {
    writer.writeFixed64(
      3,
      f
    );
  }
  f = /** @type {number} */ (jspb.Message.getField(message, 4));
  if (f != null) {
    writer.writeFixed64(
      4,
      f
    );
  }
  f = /** @type {number} */ (jspb.Message.getField(message, 5));
  if (f != null) {
    writer.writeFixed64(
      5,
      f
    );
  }
  f = message.getInList();
  if (f.length > 0) {
    writer.writeRepeatedFixed64(
      6,
      f
    );
  }
  f = message.getNotInList();
  if (f.length > 0) {
    writer.writeRepeatedFixed64(
      7,
      f
    );
  }
  f = message.getExampleList();
  if (f.length > 0) {
    writer.writeRepeatedFixed64(
      8,
      f
    );
  }
  jspb.Message.serializeBinaryExtensions(message, writer,
    proto.buf.validate.Fixed64Rules.extensionsBinary, proto.buf.validate.Fixed64Rules.prototype.getExtension);
};


/**
 * optional fixed64 const = 1;
 * @return {number}
 */
proto.buf.validate.Fixed64Rules.prototype.getConst = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/**
 * @param {number} value
 * @return {!proto.buf.validate.Fixed64Rules} returns this
 */
proto.buf.validate.Fixed64Rules.prototype.setConst = function(value) {
  return jspb.Message.setField(this, 1, value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.Fixed64Rules} returns this
 */
proto.buf.validate.Fixed64Rules.prototype.clearConst = function() {
  return jspb.Message.setField(this, 1, undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.Fixed64Rules.prototype.hasConst = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional fixed64 lt = 2;
 * @return {number}
 */
proto.buf.validate.Fixed64Rules.prototype.getLt = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/**
 * @param {number} value
 * @return {!proto.buf.validate.Fixed64Rules} returns this
 */
proto.buf.validate.Fixed64Rules.prototype.setLt = function(value) {
  return jspb.Message.setOneofField(this, 2, proto.buf.validate.Fixed64Rules.oneofGroups_[0], value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.Fixed64Rules} returns this
 */
proto.buf.validate.Fixed64Rules.prototype.clearLt = function() {
  return jspb.Message.setOneofField(this, 2, proto.buf.validate.Fixed64Rules.oneofGroups_[0], undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.Fixed64Rules.prototype.hasLt = function() {
  return jspb.Message.getField(this, 2) != null;
};


/**
 * optional fixed64 lte = 3;
 * @return {number}
 */
proto.buf.validate.Fixed64Rules.prototype.getLte = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 3, 0));
};


/**
 * @param {number} value
 * @return {!proto.buf.validate.Fixed64Rules} returns this
 */
proto.buf.validate.Fixed64Rules.prototype.setLte = function(value) {
  return jspb.Message.setOneofField(this, 3, proto.buf.validate.Fixed64Rules.oneofGroups_[0], value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.Fixed64Rules} returns this
 */
proto.buf.validate.Fixed64Rules.prototype.clearLte = function() {
  return jspb.Message.setOneofField(this, 3, proto.buf.validate.Fixed64Rules.oneofGroups_[0], undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.Fixed64Rules.prototype.hasLte = function() {
  return jspb.Message.getField(this, 3) != null;
};


/**
 * optional fixed64 gt = 4;
 * @return {number}
 */
proto.buf.validate.Fixed64Rules.prototype.getGt = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 4, 0));
};


/**
 * @param {number} value
 * @return {!proto.buf.validate.Fixed64Rules} returns this
 */
proto.buf.validate.Fixed64Rules.prototype.setGt = function(value) {
  return jspb.Message.setOneofField(this, 4, proto.buf.validate.Fixed64Rules.oneofGroups_[1], value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.Fixed64Rules} returns this
 */
proto.buf.validate.Fixed64Rules.prototype.clearGt = function() {
  return jspb.Message.setOneofField(this, 4, proto.buf.validate.Fixed64Rules.oneofGroups_[1], undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.Fixed64Rules.prototype.hasGt = function() {
  return jspb.Message.getField(this, 4) != null;
};


/**
 * optional fixed64 gte = 5;
 * @return {number}
 */
proto.buf.validate.Fixed64Rules.prototype.getGte = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 5, 0));
};


/**
 * @param {number} value
 * @return {!proto.buf.validate.Fixed64Rules} returns this
 */
proto.buf.validate.Fixed64Rules.prototype.setGte = function(value) {
  return jspb.Message.setOneofField(this, 5, proto.buf.validate.Fixed64Rules.oneofGroups_[1], value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.Fixed64Rules} returns this
 */
proto.buf.validate.Fixed64Rules.prototype.clearGte = function() {
  return jspb.Message.setOneofField(this, 5, proto.buf.validate.Fixed64Rules.oneofGroups_[1], undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.Fixed64Rules.prototype.hasGte = function() {
  return jspb.Message.getField(this, 5) != null;
};


/**
 * repeated fixed64 in = 6;
 * @return {!Array<number>}
 */
proto.buf.validate.Fixed64Rules.prototype.getInList = function() {
  return /** @type {!Array<number>} */ (jspb.Message.getRepeatedField(this, 6));
};


/**
 * @param {!Array<number>} value
 * @return {!proto.buf.validate.Fixed64Rules} returns this
 */
proto.buf.validate.Fixed64Rules.prototype.setInList = function(value) {
  return jspb.Message.setField(this, 6, value || []);
};


/**
 * @param {number} value
 * @param {number=} opt_index
 * @return {!proto.buf.validate.Fixed64Rules} returns this
 */
proto.buf.validate.Fixed64Rules.prototype.addIn = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 6, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.buf.validate.Fixed64Rules} returns this
 */
proto.buf.validate.Fixed64Rules.prototype.clearInList = function() {
  return this.setInList([]);
};


/**
 * repeated fixed64 not_in = 7;
 * @return {!Array<number>}
 */
proto.buf.validate.Fixed64Rules.prototype.getNotInList = function() {
  return /** @type {!Array<number>} */ (jspb.Message.getRepeatedField(this, 7));
};


/**
 * @param {!Array<number>} value
 * @return {!proto.buf.validate.Fixed64Rules} returns this
 */
proto.buf.validate.Fixed64Rules.prototype.setNotInList = function(value) {
  return jspb.Message.setField(this, 7, value || []);
};


/**
 * @param {number} value
 * @param {number=} opt_index
 * @return {!proto.buf.validate.Fixed64Rules} returns this
 */
proto.buf.validate.Fixed64Rules.prototype.addNotIn = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 7, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.buf.validate.Fixed64Rules} returns this
 */
proto.buf.validate.Fixed64Rules.prototype.clearNotInList = function() {
  return this.setNotInList([]);
};


/**
 * repeated fixed64 example = 8;
 * @return {!Array<number>}
 */
proto.buf.validate.Fixed64Rules.prototype.getExampleList = function() {
  return /** @type {!Array<number>} */ (jspb.Message.getRepeatedField(this, 8));
};


/**
 * @param {!Array<number>} value
 * @return {!proto.buf.validate.Fixed64Rules} returns this
 */
proto.buf.validate.Fixed64Rules.prototype.setExampleList = function(value) {
  return jspb.Message.setField(this, 8, value || []);
};


/**
 * @param {number} value
 * @param {number=} opt_index
 * @return {!proto.buf.validate.Fixed64Rules} returns this
 */
proto.buf.validate.Fixed64Rules.prototype.addExample = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 8, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.buf.validate.Fixed64Rules} returns this
 */
proto.buf.validate.Fixed64Rules.prototype.clearExampleList = function() {
  return this.setExampleList([]);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.buf.validate.SFixed32Rules.repeatedFields_ = [6,7,8];

/**
 * Oneof group definitions for this message. Each group defines the field
 * numbers belonging to that group. When of these fields' value is set, all
 * other fields in the group are cleared. During deserialization, if multiple
 * fields are encountered for a group, only the last value seen will be kept.
 * @private {!Array<!Array<number>>}
 * @const
 */
proto.buf.validate.SFixed32Rules.oneofGroups_ = [[2,3],[4,5]];

/**
 * @enum {number}
 */
proto.buf.validate.SFixed32Rules.LessThanCase = {
  LESS_THAN_NOT_SET: 0,
  LT: 2,
  LTE: 3
};

/**
 * @return {proto.buf.validate.SFixed32Rules.LessThanCase}
 */
proto.buf.validate.SFixed32Rules.prototype.getLessThanCase = function() {
  return /** @type {proto.buf.validate.SFixed32Rules.LessThanCase} */(jspb.Message.computeOneofCase(this, proto.buf.validate.SFixed32Rules.oneofGroups_[0]));
};

/**
 * @enum {number}
 */
proto.buf.validate.SFixed32Rules.GreaterThanCase = {
  GREATER_THAN_NOT_SET: 0,
  GT: 4,
  GTE: 5
};

/**
 * @return {proto.buf.validate.SFixed32Rules.GreaterThanCase}
 */
proto.buf.validate.SFixed32Rules.prototype.getGreaterThanCase = function() {
  return /** @type {proto.buf.validate.SFixed32Rules.GreaterThanCase} */(jspb.Message.computeOneofCase(this, proto.buf.validate.SFixed32Rules.oneofGroups_[1]));
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
proto.buf.validate.SFixed32Rules.prototype.toObject = function(opt_includeInstance) {
  return proto.buf.validate.SFixed32Rules.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.buf.validate.SFixed32Rules} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.buf.validate.SFixed32Rules.toObject = function(includeInstance, msg) {
  var f, obj = {
pb_const: (f = jspb.Message.getField(msg, 1)) == null ? undefined : f,
lt: (f = jspb.Message.getField(msg, 2)) == null ? undefined : f,
lte: (f = jspb.Message.getField(msg, 3)) == null ? undefined : f,
gt: (f = jspb.Message.getField(msg, 4)) == null ? undefined : f,
gte: (f = jspb.Message.getField(msg, 5)) == null ? undefined : f,
inList: (f = jspb.Message.getRepeatedField(msg, 6)) == null ? undefined : f,
notInList: (f = jspb.Message.getRepeatedField(msg, 7)) == null ? undefined : f,
exampleList: (f = jspb.Message.getRepeatedField(msg, 8)) == null ? undefined : f
  };

  jspb.Message.toObjectExtension(/** @type {!jspb.Message} */ (msg), obj,
      proto.buf.validate.SFixed32Rules.extensions, proto.buf.validate.SFixed32Rules.prototype.getExtension,
      includeInstance);
  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.buf.validate.SFixed32Rules}
 */
proto.buf.validate.SFixed32Rules.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.buf.validate.SFixed32Rules;
  return proto.buf.validate.SFixed32Rules.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.buf.validate.SFixed32Rules} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.buf.validate.SFixed32Rules}
 */
proto.buf.validate.SFixed32Rules.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readSfixed32());
      msg.setConst(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readSfixed32());
      msg.setLt(value);
      break;
    case 3:
      var value = /** @type {number} */ (reader.readSfixed32());
      msg.setLte(value);
      break;
    case 4:
      var value = /** @type {number} */ (reader.readSfixed32());
      msg.setGt(value);
      break;
    case 5:
      var value = /** @type {number} */ (reader.readSfixed32());
      msg.setGte(value);
      break;
    case 6:
      var values = /** @type {!Array<number>} */ (reader.isDelimited() ? reader.readPackedSfixed32() : [reader.readSfixed32()]);
      for (var i = 0; i < values.length; i++) {
        msg.addIn(values[i]);
      }
      break;
    case 7:
      var values = /** @type {!Array<number>} */ (reader.isDelimited() ? reader.readPackedSfixed32() : [reader.readSfixed32()]);
      for (var i = 0; i < values.length; i++) {
        msg.addNotIn(values[i]);
      }
      break;
    case 8:
      var values = /** @type {!Array<number>} */ (reader.isDelimited() ? reader.readPackedSfixed32() : [reader.readSfixed32()]);
      for (var i = 0; i < values.length; i++) {
        msg.addExample(values[i]);
      }
      break;
    default:
      jspb.Message.readBinaryExtension(msg, reader,
        proto.buf.validate.SFixed32Rules.extensionsBinary,
        proto.buf.validate.SFixed32Rules.prototype.getExtension,
        proto.buf.validate.SFixed32Rules.prototype.setExtension);
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.buf.validate.SFixed32Rules.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.buf.validate.SFixed32Rules.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.buf.validate.SFixed32Rules} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.buf.validate.SFixed32Rules.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = /** @type {number} */ (jspb.Message.getField(message, 1));
  if (f != null) {
    writer.writeSfixed32(
      1,
      f
    );
  }
  f = /** @type {number} */ (jspb.Message.getField(message, 2));
  if (f != null) {
    writer.writeSfixed32(
      2,
      f
    );
  }
  f = /** @type {number} */ (jspb.Message.getField(message, 3));
  if (f != null) {
    writer.writeSfixed32(
      3,
      f
    );
  }
  f = /** @type {number} */ (jspb.Message.getField(message, 4));
  if (f != null) {
    writer.writeSfixed32(
      4,
      f
    );
  }
  f = /** @type {number} */ (jspb.Message.getField(message, 5));
  if (f != null) {
    writer.writeSfixed32(
      5,
      f
    );
  }
  f = message.getInList();
  if (f.length > 0) {
    writer.writeRepeatedSfixed32(
      6,
      f
    );
  }
  f = message.getNotInList();
  if (f.length > 0) {
    writer.writeRepeatedSfixed32(
      7,
      f
    );
  }
  f = message.getExampleList();
  if (f.length > 0) {
    writer.writeRepeatedSfixed32(
      8,
      f
    );
  }
  jspb.Message.serializeBinaryExtensions(message, writer,
    proto.buf.validate.SFixed32Rules.extensionsBinary, proto.buf.validate.SFixed32Rules.prototype.getExtension);
};


/**
 * optional sfixed32 const = 1;
 * @return {number}
 */
proto.buf.validate.SFixed32Rules.prototype.getConst = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/**
 * @param {number} value
 * @return {!proto.buf.validate.SFixed32Rules} returns this
 */
proto.buf.validate.SFixed32Rules.prototype.setConst = function(value) {
  return jspb.Message.setField(this, 1, value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.SFixed32Rules} returns this
 */
proto.buf.validate.SFixed32Rules.prototype.clearConst = function() {
  return jspb.Message.setField(this, 1, undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.SFixed32Rules.prototype.hasConst = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional sfixed32 lt = 2;
 * @return {number}
 */
proto.buf.validate.SFixed32Rules.prototype.getLt = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/**
 * @param {number} value
 * @return {!proto.buf.validate.SFixed32Rules} returns this
 */
proto.buf.validate.SFixed32Rules.prototype.setLt = function(value) {
  return jspb.Message.setOneofField(this, 2, proto.buf.validate.SFixed32Rules.oneofGroups_[0], value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.SFixed32Rules} returns this
 */
proto.buf.validate.SFixed32Rules.prototype.clearLt = function() {
  return jspb.Message.setOneofField(this, 2, proto.buf.validate.SFixed32Rules.oneofGroups_[0], undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.SFixed32Rules.prototype.hasLt = function() {
  return jspb.Message.getField(this, 2) != null;
};


/**
 * optional sfixed32 lte = 3;
 * @return {number}
 */
proto.buf.validate.SFixed32Rules.prototype.getLte = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 3, 0));
};


/**
 * @param {number} value
 * @return {!proto.buf.validate.SFixed32Rules} returns this
 */
proto.buf.validate.SFixed32Rules.prototype.setLte = function(value) {
  return jspb.Message.setOneofField(this, 3, proto.buf.validate.SFixed32Rules.oneofGroups_[0], value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.SFixed32Rules} returns this
 */
proto.buf.validate.SFixed32Rules.prototype.clearLte = function() {
  return jspb.Message.setOneofField(this, 3, proto.buf.validate.SFixed32Rules.oneofGroups_[0], undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.SFixed32Rules.prototype.hasLte = function() {
  return jspb.Message.getField(this, 3) != null;
};


/**
 * optional sfixed32 gt = 4;
 * @return {number}
 */
proto.buf.validate.SFixed32Rules.prototype.getGt = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 4, 0));
};


/**
 * @param {number} value
 * @return {!proto.buf.validate.SFixed32Rules} returns this
 */
proto.buf.validate.SFixed32Rules.prototype.setGt = function(value) {
  return jspb.Message.setOneofField(this, 4, proto.buf.validate.SFixed32Rules.oneofGroups_[1], value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.SFixed32Rules} returns this
 */
proto.buf.validate.SFixed32Rules.prototype.clearGt = function() {
  return jspb.Message.setOneofField(this, 4, proto.buf.validate.SFixed32Rules.oneofGroups_[1], undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.SFixed32Rules.prototype.hasGt = function() {
  return jspb.Message.getField(this, 4) != null;
};


/**
 * optional sfixed32 gte = 5;
 * @return {number}
 */
proto.buf.validate.SFixed32Rules.prototype.getGte = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 5, 0));
};


/**
 * @param {number} value
 * @return {!proto.buf.validate.SFixed32Rules} returns this
 */
proto.buf.validate.SFixed32Rules.prototype.setGte = function(value) {
  return jspb.Message.setOneofField(this, 5, proto.buf.validate.SFixed32Rules.oneofGroups_[1], value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.SFixed32Rules} returns this
 */
proto.buf.validate.SFixed32Rules.prototype.clearGte = function() {
  return jspb.Message.setOneofField(this, 5, proto.buf.validate.SFixed32Rules.oneofGroups_[1], undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.SFixed32Rules.prototype.hasGte = function() {
  return jspb.Message.getField(this, 5) != null;
};


/**
 * repeated sfixed32 in = 6;
 * @return {!Array<number>}
 */
proto.buf.validate.SFixed32Rules.prototype.getInList = function() {
  return /** @type {!Array<number>} */ (jspb.Message.getRepeatedField(this, 6));
};


/**
 * @param {!Array<number>} value
 * @return {!proto.buf.validate.SFixed32Rules} returns this
 */
proto.buf.validate.SFixed32Rules.prototype.setInList = function(value) {
  return jspb.Message.setField(this, 6, value || []);
};


/**
 * @param {number} value
 * @param {number=} opt_index
 * @return {!proto.buf.validate.SFixed32Rules} returns this
 */
proto.buf.validate.SFixed32Rules.prototype.addIn = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 6, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.buf.validate.SFixed32Rules} returns this
 */
proto.buf.validate.SFixed32Rules.prototype.clearInList = function() {
  return this.setInList([]);
};


/**
 * repeated sfixed32 not_in = 7;
 * @return {!Array<number>}
 */
proto.buf.validate.SFixed32Rules.prototype.getNotInList = function() {
  return /** @type {!Array<number>} */ (jspb.Message.getRepeatedField(this, 7));
};


/**
 * @param {!Array<number>} value
 * @return {!proto.buf.validate.SFixed32Rules} returns this
 */
proto.buf.validate.SFixed32Rules.prototype.setNotInList = function(value) {
  return jspb.Message.setField(this, 7, value || []);
};


/**
 * @param {number} value
 * @param {number=} opt_index
 * @return {!proto.buf.validate.SFixed32Rules} returns this
 */
proto.buf.validate.SFixed32Rules.prototype.addNotIn = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 7, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.buf.validate.SFixed32Rules} returns this
 */
proto.buf.validate.SFixed32Rules.prototype.clearNotInList = function() {
  return this.setNotInList([]);
};


/**
 * repeated sfixed32 example = 8;
 * @return {!Array<number>}
 */
proto.buf.validate.SFixed32Rules.prototype.getExampleList = function() {
  return /** @type {!Array<number>} */ (jspb.Message.getRepeatedField(this, 8));
};


/**
 * @param {!Array<number>} value
 * @return {!proto.buf.validate.SFixed32Rules} returns this
 */
proto.buf.validate.SFixed32Rules.prototype.setExampleList = function(value) {
  return jspb.Message.setField(this, 8, value || []);
};


/**
 * @param {number} value
 * @param {number=} opt_index
 * @return {!proto.buf.validate.SFixed32Rules} returns this
 */
proto.buf.validate.SFixed32Rules.prototype.addExample = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 8, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.buf.validate.SFixed32Rules} returns this
 */
proto.buf.validate.SFixed32Rules.prototype.clearExampleList = function() {
  return this.setExampleList([]);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.buf.validate.SFixed64Rules.repeatedFields_ = [6,7,8];

/**
 * Oneof group definitions for this message. Each group defines the field
 * numbers belonging to that group. When of these fields' value is set, all
 * other fields in the group are cleared. During deserialization, if multiple
 * fields are encountered for a group, only the last value seen will be kept.
 * @private {!Array<!Array<number>>}
 * @const
 */
proto.buf.validate.SFixed64Rules.oneofGroups_ = [[2,3],[4,5]];

/**
 * @enum {number}
 */
proto.buf.validate.SFixed64Rules.LessThanCase = {
  LESS_THAN_NOT_SET: 0,
  LT: 2,
  LTE: 3
};

/**
 * @return {proto.buf.validate.SFixed64Rules.LessThanCase}
 */
proto.buf.validate.SFixed64Rules.prototype.getLessThanCase = function() {
  return /** @type {proto.buf.validate.SFixed64Rules.LessThanCase} */(jspb.Message.computeOneofCase(this, proto.buf.validate.SFixed64Rules.oneofGroups_[0]));
};

/**
 * @enum {number}
 */
proto.buf.validate.SFixed64Rules.GreaterThanCase = {
  GREATER_THAN_NOT_SET: 0,
  GT: 4,
  GTE: 5
};

/**
 * @return {proto.buf.validate.SFixed64Rules.GreaterThanCase}
 */
proto.buf.validate.SFixed64Rules.prototype.getGreaterThanCase = function() {
  return /** @type {proto.buf.validate.SFixed64Rules.GreaterThanCase} */(jspb.Message.computeOneofCase(this, proto.buf.validate.SFixed64Rules.oneofGroups_[1]));
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
proto.buf.validate.SFixed64Rules.prototype.toObject = function(opt_includeInstance) {
  return proto.buf.validate.SFixed64Rules.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.buf.validate.SFixed64Rules} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.buf.validate.SFixed64Rules.toObject = function(includeInstance, msg) {
  var f, obj = {
pb_const: (f = jspb.Message.getField(msg, 1)) == null ? undefined : f,
lt: (f = jspb.Message.getField(msg, 2)) == null ? undefined : f,
lte: (f = jspb.Message.getField(msg, 3)) == null ? undefined : f,
gt: (f = jspb.Message.getField(msg, 4)) == null ? undefined : f,
gte: (f = jspb.Message.getField(msg, 5)) == null ? undefined : f,
inList: (f = jspb.Message.getRepeatedField(msg, 6)) == null ? undefined : f,
notInList: (f = jspb.Message.getRepeatedField(msg, 7)) == null ? undefined : f,
exampleList: (f = jspb.Message.getRepeatedField(msg, 8)) == null ? undefined : f
  };

  jspb.Message.toObjectExtension(/** @type {!jspb.Message} */ (msg), obj,
      proto.buf.validate.SFixed64Rules.extensions, proto.buf.validate.SFixed64Rules.prototype.getExtension,
      includeInstance);
  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.buf.validate.SFixed64Rules}
 */
proto.buf.validate.SFixed64Rules.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.buf.validate.SFixed64Rules;
  return proto.buf.validate.SFixed64Rules.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.buf.validate.SFixed64Rules} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.buf.validate.SFixed64Rules}
 */
proto.buf.validate.SFixed64Rules.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readSfixed64());
      msg.setConst(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readSfixed64());
      msg.setLt(value);
      break;
    case 3:
      var value = /** @type {number} */ (reader.readSfixed64());
      msg.setLte(value);
      break;
    case 4:
      var value = /** @type {number} */ (reader.readSfixed64());
      msg.setGt(value);
      break;
    case 5:
      var value = /** @type {number} */ (reader.readSfixed64());
      msg.setGte(value);
      break;
    case 6:
      var values = /** @type {!Array<number>} */ (reader.isDelimited() ? reader.readPackedSfixed64() : [reader.readSfixed64()]);
      for (var i = 0; i < values.length; i++) {
        msg.addIn(values[i]);
      }
      break;
    case 7:
      var values = /** @type {!Array<number>} */ (reader.isDelimited() ? reader.readPackedSfixed64() : [reader.readSfixed64()]);
      for (var i = 0; i < values.length; i++) {
        msg.addNotIn(values[i]);
      }
      break;
    case 8:
      var values = /** @type {!Array<number>} */ (reader.isDelimited() ? reader.readPackedSfixed64() : [reader.readSfixed64()]);
      for (var i = 0; i < values.length; i++) {
        msg.addExample(values[i]);
      }
      break;
    default:
      jspb.Message.readBinaryExtension(msg, reader,
        proto.buf.validate.SFixed64Rules.extensionsBinary,
        proto.buf.validate.SFixed64Rules.prototype.getExtension,
        proto.buf.validate.SFixed64Rules.prototype.setExtension);
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.buf.validate.SFixed64Rules.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.buf.validate.SFixed64Rules.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.buf.validate.SFixed64Rules} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.buf.validate.SFixed64Rules.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = /** @type {number} */ (jspb.Message.getField(message, 1));
  if (f != null) {
    writer.writeSfixed64(
      1,
      f
    );
  }
  f = /** @type {number} */ (jspb.Message.getField(message, 2));
  if (f != null) {
    writer.writeSfixed64(
      2,
      f
    );
  }
  f = /** @type {number} */ (jspb.Message.getField(message, 3));
  if (f != null) {
    writer.writeSfixed64(
      3,
      f
    );
  }
  f = /** @type {number} */ (jspb.Message.getField(message, 4));
  if (f != null) {
    writer.writeSfixed64(
      4,
      f
    );
  }
  f = /** @type {number} */ (jspb.Message.getField(message, 5));
  if (f != null) {
    writer.writeSfixed64(
      5,
      f
    );
  }
  f = message.getInList();
  if (f.length > 0) {
    writer.writeRepeatedSfixed64(
      6,
      f
    );
  }
  f = message.getNotInList();
  if (f.length > 0) {
    writer.writeRepeatedSfixed64(
      7,
      f
    );
  }
  f = message.getExampleList();
  if (f.length > 0) {
    writer.writeRepeatedSfixed64(
      8,
      f
    );
  }
  jspb.Message.serializeBinaryExtensions(message, writer,
    proto.buf.validate.SFixed64Rules.extensionsBinary, proto.buf.validate.SFixed64Rules.prototype.getExtension);
};


/**
 * optional sfixed64 const = 1;
 * @return {number}
 */
proto.buf.validate.SFixed64Rules.prototype.getConst = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/**
 * @param {number} value
 * @return {!proto.buf.validate.SFixed64Rules} returns this
 */
proto.buf.validate.SFixed64Rules.prototype.setConst = function(value) {
  return jspb.Message.setField(this, 1, value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.SFixed64Rules} returns this
 */
proto.buf.validate.SFixed64Rules.prototype.clearConst = function() {
  return jspb.Message.setField(this, 1, undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.SFixed64Rules.prototype.hasConst = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional sfixed64 lt = 2;
 * @return {number}
 */
proto.buf.validate.SFixed64Rules.prototype.getLt = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/**
 * @param {number} value
 * @return {!proto.buf.validate.SFixed64Rules} returns this
 */
proto.buf.validate.SFixed64Rules.prototype.setLt = function(value) {
  return jspb.Message.setOneofField(this, 2, proto.buf.validate.SFixed64Rules.oneofGroups_[0], value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.SFixed64Rules} returns this
 */
proto.buf.validate.SFixed64Rules.prototype.clearLt = function() {
  return jspb.Message.setOneofField(this, 2, proto.buf.validate.SFixed64Rules.oneofGroups_[0], undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.SFixed64Rules.prototype.hasLt = function() {
  return jspb.Message.getField(this, 2) != null;
};


/**
 * optional sfixed64 lte = 3;
 * @return {number}
 */
proto.buf.validate.SFixed64Rules.prototype.getLte = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 3, 0));
};


/**
 * @param {number} value
 * @return {!proto.buf.validate.SFixed64Rules} returns this
 */
proto.buf.validate.SFixed64Rules.prototype.setLte = function(value) {
  return jspb.Message.setOneofField(this, 3, proto.buf.validate.SFixed64Rules.oneofGroups_[0], value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.SFixed64Rules} returns this
 */
proto.buf.validate.SFixed64Rules.prototype.clearLte = function() {
  return jspb.Message.setOneofField(this, 3, proto.buf.validate.SFixed64Rules.oneofGroups_[0], undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.SFixed64Rules.prototype.hasLte = function() {
  return jspb.Message.getField(this, 3) != null;
};


/**
 * optional sfixed64 gt = 4;
 * @return {number}
 */
proto.buf.validate.SFixed64Rules.prototype.getGt = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 4, 0));
};


/**
 * @param {number} value
 * @return {!proto.buf.validate.SFixed64Rules} returns this
 */
proto.buf.validate.SFixed64Rules.prototype.setGt = function(value) {
  return jspb.Message.setOneofField(this, 4, proto.buf.validate.SFixed64Rules.oneofGroups_[1], value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.SFixed64Rules} returns this
 */
proto.buf.validate.SFixed64Rules.prototype.clearGt = function() {
  return jspb.Message.setOneofField(this, 4, proto.buf.validate.SFixed64Rules.oneofGroups_[1], undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.SFixed64Rules.prototype.hasGt = function() {
  return jspb.Message.getField(this, 4) != null;
};


/**
 * optional sfixed64 gte = 5;
 * @return {number}
 */
proto.buf.validate.SFixed64Rules.prototype.getGte = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 5, 0));
};


/**
 * @param {number} value
 * @return {!proto.buf.validate.SFixed64Rules} returns this
 */
proto.buf.validate.SFixed64Rules.prototype.setGte = function(value) {
  return jspb.Message.setOneofField(this, 5, proto.buf.validate.SFixed64Rules.oneofGroups_[1], value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.SFixed64Rules} returns this
 */
proto.buf.validate.SFixed64Rules.prototype.clearGte = function() {
  return jspb.Message.setOneofField(this, 5, proto.buf.validate.SFixed64Rules.oneofGroups_[1], undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.SFixed64Rules.prototype.hasGte = function() {
  return jspb.Message.getField(this, 5) != null;
};


/**
 * repeated sfixed64 in = 6;
 * @return {!Array<number>}
 */
proto.buf.validate.SFixed64Rules.prototype.getInList = function() {
  return /** @type {!Array<number>} */ (jspb.Message.getRepeatedField(this, 6));
};


/**
 * @param {!Array<number>} value
 * @return {!proto.buf.validate.SFixed64Rules} returns this
 */
proto.buf.validate.SFixed64Rules.prototype.setInList = function(value) {
  return jspb.Message.setField(this, 6, value || []);
};


/**
 * @param {number} value
 * @param {number=} opt_index
 * @return {!proto.buf.validate.SFixed64Rules} returns this
 */
proto.buf.validate.SFixed64Rules.prototype.addIn = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 6, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.buf.validate.SFixed64Rules} returns this
 */
proto.buf.validate.SFixed64Rules.prototype.clearInList = function() {
  return this.setInList([]);
};


/**
 * repeated sfixed64 not_in = 7;
 * @return {!Array<number>}
 */
proto.buf.validate.SFixed64Rules.prototype.getNotInList = function() {
  return /** @type {!Array<number>} */ (jspb.Message.getRepeatedField(this, 7));
};


/**
 * @param {!Array<number>} value
 * @return {!proto.buf.validate.SFixed64Rules} returns this
 */
proto.buf.validate.SFixed64Rules.prototype.setNotInList = function(value) {
  return jspb.Message.setField(this, 7, value || []);
};


/**
 * @param {number} value
 * @param {number=} opt_index
 * @return {!proto.buf.validate.SFixed64Rules} returns this
 */
proto.buf.validate.SFixed64Rules.prototype.addNotIn = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 7, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.buf.validate.SFixed64Rules} returns this
 */
proto.buf.validate.SFixed64Rules.prototype.clearNotInList = function() {
  return this.setNotInList([]);
};


/**
 * repeated sfixed64 example = 8;
 * @return {!Array<number>}
 */
proto.buf.validate.SFixed64Rules.prototype.getExampleList = function() {
  return /** @type {!Array<number>} */ (jspb.Message.getRepeatedField(this, 8));
};


/**
 * @param {!Array<number>} value
 * @return {!proto.buf.validate.SFixed64Rules} returns this
 */
proto.buf.validate.SFixed64Rules.prototype.setExampleList = function(value) {
  return jspb.Message.setField(this, 8, value || []);
};


/**
 * @param {number} value
 * @param {number=} opt_index
 * @return {!proto.buf.validate.SFixed64Rules} returns this
 */
proto.buf.validate.SFixed64Rules.prototype.addExample = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 8, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.buf.validate.SFixed64Rules} returns this
 */
proto.buf.validate.SFixed64Rules.prototype.clearExampleList = function() {
  return this.setExampleList([]);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.buf.validate.BoolRules.repeatedFields_ = [2];



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
proto.buf.validate.BoolRules.prototype.toObject = function(opt_includeInstance) {
  return proto.buf.validate.BoolRules.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.buf.validate.BoolRules} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.buf.validate.BoolRules.toObject = function(includeInstance, msg) {
  var f, obj = {
pb_const: (f = jspb.Message.getBooleanField(msg, 1)) == null ? undefined : f,
exampleList: (f = jspb.Message.getRepeatedBooleanField(msg, 2)) == null ? undefined : f
  };

  jspb.Message.toObjectExtension(/** @type {!jspb.Message} */ (msg), obj,
      proto.buf.validate.BoolRules.extensions, proto.buf.validate.BoolRules.prototype.getExtension,
      includeInstance);
  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.buf.validate.BoolRules}
 */
proto.buf.validate.BoolRules.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.buf.validate.BoolRules;
  return proto.buf.validate.BoolRules.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.buf.validate.BoolRules} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.buf.validate.BoolRules}
 */
proto.buf.validate.BoolRules.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setConst(value);
      break;
    case 2:
      var values = /** @type {!Array<boolean>} */ (reader.isDelimited() ? reader.readPackedBool() : [reader.readBool()]);
      for (var i = 0; i < values.length; i++) {
        msg.addExample(values[i]);
      }
      break;
    default:
      jspb.Message.readBinaryExtension(msg, reader,
        proto.buf.validate.BoolRules.extensionsBinary,
        proto.buf.validate.BoolRules.prototype.getExtension,
        proto.buf.validate.BoolRules.prototype.setExtension);
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.buf.validate.BoolRules.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.buf.validate.BoolRules.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.buf.validate.BoolRules} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.buf.validate.BoolRules.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = /** @type {boolean} */ (jspb.Message.getField(message, 1));
  if (f != null) {
    writer.writeBool(
      1,
      f
    );
  }
  f = message.getExampleList();
  if (f.length > 0) {
    writer.writeRepeatedBool(
      2,
      f
    );
  }
  jspb.Message.serializeBinaryExtensions(message, writer,
    proto.buf.validate.BoolRules.extensionsBinary, proto.buf.validate.BoolRules.prototype.getExtension);
};


/**
 * optional bool const = 1;
 * @return {boolean}
 */
proto.buf.validate.BoolRules.prototype.getConst = function() {
  return /** @type {boolean} */ (jspb.Message.getBooleanFieldWithDefault(this, 1, false));
};


/**
 * @param {boolean} value
 * @return {!proto.buf.validate.BoolRules} returns this
 */
proto.buf.validate.BoolRules.prototype.setConst = function(value) {
  return jspb.Message.setField(this, 1, value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.BoolRules} returns this
 */
proto.buf.validate.BoolRules.prototype.clearConst = function() {
  return jspb.Message.setField(this, 1, undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.BoolRules.prototype.hasConst = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * repeated bool example = 2;
 * @return {!Array<boolean>}
 */
proto.buf.validate.BoolRules.prototype.getExampleList = function() {
  return /** @type {!Array<boolean>} */ (jspb.Message.getRepeatedBooleanField(this, 2));
};


/**
 * @param {!Array<boolean>} value
 * @return {!proto.buf.validate.BoolRules} returns this
 */
proto.buf.validate.BoolRules.prototype.setExampleList = function(value) {
  return jspb.Message.setField(this, 2, value || []);
};


/**
 * @param {boolean} value
 * @param {number=} opt_index
 * @return {!proto.buf.validate.BoolRules} returns this
 */
proto.buf.validate.BoolRules.prototype.addExample = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 2, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.buf.validate.BoolRules} returns this
 */
proto.buf.validate.BoolRules.prototype.clearExampleList = function() {
  return this.setExampleList([]);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.buf.validate.StringRules.repeatedFields_ = [10,11,34];

/**
 * Oneof group definitions for this message. Each group defines the field
 * numbers belonging to that group. When of these fields' value is set, all
 * other fields in the group are cleared. During deserialization, if multiple
 * fields are encountered for a group, only the last value seen will be kept.
 * @private {!Array<!Array<number>>}
 * @const
 */
proto.buf.validate.StringRules.oneofGroups_ = [[12,13,14,15,16,17,18,21,22,33,26,27,28,29,30,31,32,24]];

/**
 * @enum {number}
 */
proto.buf.validate.StringRules.WellKnownCase = {
  WELL_KNOWN_NOT_SET: 0,
  EMAIL: 12,
  HOSTNAME: 13,
  IP: 14,
  IPV4: 15,
  IPV6: 16,
  URI: 17,
  URI_REF: 18,
  ADDRESS: 21,
  UUID: 22,
  TUUID: 33,
  IP_WITH_PREFIXLEN: 26,
  IPV4_WITH_PREFIXLEN: 27,
  IPV6_WITH_PREFIXLEN: 28,
  IP_PREFIX: 29,
  IPV4_PREFIX: 30,
  IPV6_PREFIX: 31,
  HOST_AND_PORT: 32,
  WELL_KNOWN_REGEX: 24
};

/**
 * @return {proto.buf.validate.StringRules.WellKnownCase}
 */
proto.buf.validate.StringRules.prototype.getWellKnownCase = function() {
  return /** @type {proto.buf.validate.StringRules.WellKnownCase} */(jspb.Message.computeOneofCase(this, proto.buf.validate.StringRules.oneofGroups_[0]));
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
proto.buf.validate.StringRules.prototype.toObject = function(opt_includeInstance) {
  return proto.buf.validate.StringRules.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.buf.validate.StringRules} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.buf.validate.StringRules.toObject = function(includeInstance, msg) {
  var f, obj = {
pb_const: (f = jspb.Message.getField(msg, 1)) == null ? undefined : f,
len: (f = jspb.Message.getField(msg, 19)) == null ? undefined : f,
minLen: (f = jspb.Message.getField(msg, 2)) == null ? undefined : f,
maxLen: (f = jspb.Message.getField(msg, 3)) == null ? undefined : f,
lenBytes: (f = jspb.Message.getField(msg, 20)) == null ? undefined : f,
minBytes: (f = jspb.Message.getField(msg, 4)) == null ? undefined : f,
maxBytes: (f = jspb.Message.getField(msg, 5)) == null ? undefined : f,
pattern: (f = jspb.Message.getField(msg, 6)) == null ? undefined : f,
prefix: (f = jspb.Message.getField(msg, 7)) == null ? undefined : f,
suffix: (f = jspb.Message.getField(msg, 8)) == null ? undefined : f,
contains: (f = jspb.Message.getField(msg, 9)) == null ? undefined : f,
notContains: (f = jspb.Message.getField(msg, 23)) == null ? undefined : f,
inList: (f = jspb.Message.getRepeatedField(msg, 10)) == null ? undefined : f,
notInList: (f = jspb.Message.getRepeatedField(msg, 11)) == null ? undefined : f,
email: (f = jspb.Message.getBooleanField(msg, 12)) == null ? undefined : f,
hostname: (f = jspb.Message.getBooleanField(msg, 13)) == null ? undefined : f,
ip: (f = jspb.Message.getBooleanField(msg, 14)) == null ? undefined : f,
ipv4: (f = jspb.Message.getBooleanField(msg, 15)) == null ? undefined : f,
ipv6: (f = jspb.Message.getBooleanField(msg, 16)) == null ? undefined : f,
uri: (f = jspb.Message.getBooleanField(msg, 17)) == null ? undefined : f,
uriRef: (f = jspb.Message.getBooleanField(msg, 18)) == null ? undefined : f,
address: (f = jspb.Message.getBooleanField(msg, 21)) == null ? undefined : f,
uuid: (f = jspb.Message.getBooleanField(msg, 22)) == null ? undefined : f,
tuuid: (f = jspb.Message.getBooleanField(msg, 33)) == null ? undefined : f,
ipWithPrefixlen: (f = jspb.Message.getBooleanField(msg, 26)) == null ? undefined : f,
ipv4WithPrefixlen: (f = jspb.Message.getBooleanField(msg, 27)) == null ? undefined : f,
ipv6WithPrefixlen: (f = jspb.Message.getBooleanField(msg, 28)) == null ? undefined : f,
ipPrefix: (f = jspb.Message.getBooleanField(msg, 29)) == null ? undefined : f,
ipv4Prefix: (f = jspb.Message.getBooleanField(msg, 30)) == null ? undefined : f,
ipv6Prefix: (f = jspb.Message.getBooleanField(msg, 31)) == null ? undefined : f,
hostAndPort: (f = jspb.Message.getBooleanField(msg, 32)) == null ? undefined : f,
wellKnownRegex: (f = jspb.Message.getField(msg, 24)) == null ? undefined : f,
strict: (f = jspb.Message.getBooleanField(msg, 25)) == null ? undefined : f,
exampleList: (f = jspb.Message.getRepeatedField(msg, 34)) == null ? undefined : f
  };

  jspb.Message.toObjectExtension(/** @type {!jspb.Message} */ (msg), obj,
      proto.buf.validate.StringRules.extensions, proto.buf.validate.StringRules.prototype.getExtension,
      includeInstance);
  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.buf.validate.StringRules}
 */
proto.buf.validate.StringRules.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.buf.validate.StringRules;
  return proto.buf.validate.StringRules.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.buf.validate.StringRules} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.buf.validate.StringRules}
 */
proto.buf.validate.StringRules.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setConst(value);
      break;
    case 19:
      var value = /** @type {number} */ (reader.readUint64());
      msg.setLen(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readUint64());
      msg.setMinLen(value);
      break;
    case 3:
      var value = /** @type {number} */ (reader.readUint64());
      msg.setMaxLen(value);
      break;
    case 20:
      var value = /** @type {number} */ (reader.readUint64());
      msg.setLenBytes(value);
      break;
    case 4:
      var value = /** @type {number} */ (reader.readUint64());
      msg.setMinBytes(value);
      break;
    case 5:
      var value = /** @type {number} */ (reader.readUint64());
      msg.setMaxBytes(value);
      break;
    case 6:
      var value = /** @type {string} */ (reader.readString());
      msg.setPattern(value);
      break;
    case 7:
      var value = /** @type {string} */ (reader.readString());
      msg.setPrefix(value);
      break;
    case 8:
      var value = /** @type {string} */ (reader.readString());
      msg.setSuffix(value);
      break;
    case 9:
      var value = /** @type {string} */ (reader.readString());
      msg.setContains(value);
      break;
    case 23:
      var value = /** @type {string} */ (reader.readString());
      msg.setNotContains(value);
      break;
    case 10:
      var value = /** @type {string} */ (reader.readString());
      msg.addIn(value);
      break;
    case 11:
      var value = /** @type {string} */ (reader.readString());
      msg.addNotIn(value);
      break;
    case 12:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setEmail(value);
      break;
    case 13:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setHostname(value);
      break;
    case 14:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setIp(value);
      break;
    case 15:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setIpv4(value);
      break;
    case 16:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setIpv6(value);
      break;
    case 17:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setUri(value);
      break;
    case 18:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setUriRef(value);
      break;
    case 21:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setAddress(value);
      break;
    case 22:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setUuid(value);
      break;
    case 33:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setTuuid(value);
      break;
    case 26:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setIpWithPrefixlen(value);
      break;
    case 27:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setIpv4WithPrefixlen(value);
      break;
    case 28:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setIpv6WithPrefixlen(value);
      break;
    case 29:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setIpPrefix(value);
      break;
    case 30:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setIpv4Prefix(value);
      break;
    case 31:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setIpv6Prefix(value);
      break;
    case 32:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setHostAndPort(value);
      break;
    case 24:
      var value = /** @type {!proto.buf.validate.KnownRegex} */ (reader.readEnum());
      msg.setWellKnownRegex(value);
      break;
    case 25:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setStrict(value);
      break;
    case 34:
      var value = /** @type {string} */ (reader.readString());
      msg.addExample(value);
      break;
    default:
      jspb.Message.readBinaryExtension(msg, reader,
        proto.buf.validate.StringRules.extensionsBinary,
        proto.buf.validate.StringRules.prototype.getExtension,
        proto.buf.validate.StringRules.prototype.setExtension);
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.buf.validate.StringRules.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.buf.validate.StringRules.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.buf.validate.StringRules} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.buf.validate.StringRules.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = /** @type {string} */ (jspb.Message.getField(message, 1));
  if (f != null) {
    writer.writeString(
      1,
      f
    );
  }
  f = /** @type {number} */ (jspb.Message.getField(message, 19));
  if (f != null) {
    writer.writeUint64(
      19,
      f
    );
  }
  f = /** @type {number} */ (jspb.Message.getField(message, 2));
  if (f != null) {
    writer.writeUint64(
      2,
      f
    );
  }
  f = /** @type {number} */ (jspb.Message.getField(message, 3));
  if (f != null) {
    writer.writeUint64(
      3,
      f
    );
  }
  f = /** @type {number} */ (jspb.Message.getField(message, 20));
  if (f != null) {
    writer.writeUint64(
      20,
      f
    );
  }
  f = /** @type {number} */ (jspb.Message.getField(message, 4));
  if (f != null) {
    writer.writeUint64(
      4,
      f
    );
  }
  f = /** @type {number} */ (jspb.Message.getField(message, 5));
  if (f != null) {
    writer.writeUint64(
      5,
      f
    );
  }
  f = /** @type {string} */ (jspb.Message.getField(message, 6));
  if (f != null) {
    writer.writeString(
      6,
      f
    );
  }
  f = /** @type {string} */ (jspb.Message.getField(message, 7));
  if (f != null) {
    writer.writeString(
      7,
      f
    );
  }
  f = /** @type {string} */ (jspb.Message.getField(message, 8));
  if (f != null) {
    writer.writeString(
      8,
      f
    );
  }
  f = /** @type {string} */ (jspb.Message.getField(message, 9));
  if (f != null) {
    writer.writeString(
      9,
      f
    );
  }
  f = /** @type {string} */ (jspb.Message.getField(message, 23));
  if (f != null) {
    writer.writeString(
      23,
      f
    );
  }
  f = message.getInList();
  if (f.length > 0) {
    writer.writeRepeatedString(
      10,
      f
    );
  }
  f = message.getNotInList();
  if (f.length > 0) {
    writer.writeRepeatedString(
      11,
      f
    );
  }
  f = /** @type {boolean} */ (jspb.Message.getField(message, 12));
  if (f != null) {
    writer.writeBool(
      12,
      f
    );
  }
  f = /** @type {boolean} */ (jspb.Message.getField(message, 13));
  if (f != null) {
    writer.writeBool(
      13,
      f
    );
  }
  f = /** @type {boolean} */ (jspb.Message.getField(message, 14));
  if (f != null) {
    writer.writeBool(
      14,
      f
    );
  }
  f = /** @type {boolean} */ (jspb.Message.getField(message, 15));
  if (f != null) {
    writer.writeBool(
      15,
      f
    );
  }
  f = /** @type {boolean} */ (jspb.Message.getField(message, 16));
  if (f != null) {
    writer.writeBool(
      16,
      f
    );
  }
  f = /** @type {boolean} */ (jspb.Message.getField(message, 17));
  if (f != null) {
    writer.writeBool(
      17,
      f
    );
  }
  f = /** @type {boolean} */ (jspb.Message.getField(message, 18));
  if (f != null) {
    writer.writeBool(
      18,
      f
    );
  }
  f = /** @type {boolean} */ (jspb.Message.getField(message, 21));
  if (f != null) {
    writer.writeBool(
      21,
      f
    );
  }
  f = /** @type {boolean} */ (jspb.Message.getField(message, 22));
  if (f != null) {
    writer.writeBool(
      22,
      f
    );
  }
  f = /** @type {boolean} */ (jspb.Message.getField(message, 33));
  if (f != null) {
    writer.writeBool(
      33,
      f
    );
  }
  f = /** @type {boolean} */ (jspb.Message.getField(message, 26));
  if (f != null) {
    writer.writeBool(
      26,
      f
    );
  }
  f = /** @type {boolean} */ (jspb.Message.getField(message, 27));
  if (f != null) {
    writer.writeBool(
      27,
      f
    );
  }
  f = /** @type {boolean} */ (jspb.Message.getField(message, 28));
  if (f != null) {
    writer.writeBool(
      28,
      f
    );
  }
  f = /** @type {boolean} */ (jspb.Message.getField(message, 29));
  if (f != null) {
    writer.writeBool(
      29,
      f
    );
  }
  f = /** @type {boolean} */ (jspb.Message.getField(message, 30));
  if (f != null) {
    writer.writeBool(
      30,
      f
    );
  }
  f = /** @type {boolean} */ (jspb.Message.getField(message, 31));
  if (f != null) {
    writer.writeBool(
      31,
      f
    );
  }
  f = /** @type {boolean} */ (jspb.Message.getField(message, 32));
  if (f != null) {
    writer.writeBool(
      32,
      f
    );
  }
  f = /** @type {!proto.buf.validate.KnownRegex} */ (jspb.Message.getField(message, 24));
  if (f != null) {
    writer.writeEnum(
      24,
      f
    );
  }
  f = /** @type {boolean} */ (jspb.Message.getField(message, 25));
  if (f != null) {
    writer.writeBool(
      25,
      f
    );
  }
  f = message.getExampleList();
  if (f.length > 0) {
    writer.writeRepeatedString(
      34,
      f
    );
  }
  jspb.Message.serializeBinaryExtensions(message, writer,
    proto.buf.validate.StringRules.extensionsBinary, proto.buf.validate.StringRules.prototype.getExtension);
};


/**
 * optional string const = 1;
 * @return {string}
 */
proto.buf.validate.StringRules.prototype.getConst = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.buf.validate.StringRules} returns this
 */
proto.buf.validate.StringRules.prototype.setConst = function(value) {
  return jspb.Message.setField(this, 1, value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.StringRules} returns this
 */
proto.buf.validate.StringRules.prototype.clearConst = function() {
  return jspb.Message.setField(this, 1, undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.StringRules.prototype.hasConst = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional uint64 len = 19;
 * @return {number}
 */
proto.buf.validate.StringRules.prototype.getLen = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 19, 0));
};


/**
 * @param {number} value
 * @return {!proto.buf.validate.StringRules} returns this
 */
proto.buf.validate.StringRules.prototype.setLen = function(value) {
  return jspb.Message.setField(this, 19, value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.StringRules} returns this
 */
proto.buf.validate.StringRules.prototype.clearLen = function() {
  return jspb.Message.setField(this, 19, undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.StringRules.prototype.hasLen = function() {
  return jspb.Message.getField(this, 19) != null;
};


/**
 * optional uint64 min_len = 2;
 * @return {number}
 */
proto.buf.validate.StringRules.prototype.getMinLen = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/**
 * @param {number} value
 * @return {!proto.buf.validate.StringRules} returns this
 */
proto.buf.validate.StringRules.prototype.setMinLen = function(value) {
  return jspb.Message.setField(this, 2, value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.StringRules} returns this
 */
proto.buf.validate.StringRules.prototype.clearMinLen = function() {
  return jspb.Message.setField(this, 2, undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.StringRules.prototype.hasMinLen = function() {
  return jspb.Message.getField(this, 2) != null;
};


/**
 * optional uint64 max_len = 3;
 * @return {number}
 */
proto.buf.validate.StringRules.prototype.getMaxLen = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 3, 0));
};


/**
 * @param {number} value
 * @return {!proto.buf.validate.StringRules} returns this
 */
proto.buf.validate.StringRules.prototype.setMaxLen = function(value) {
  return jspb.Message.setField(this, 3, value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.StringRules} returns this
 */
proto.buf.validate.StringRules.prototype.clearMaxLen = function() {
  return jspb.Message.setField(this, 3, undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.StringRules.prototype.hasMaxLen = function() {
  return jspb.Message.getField(this, 3) != null;
};


/**
 * optional uint64 len_bytes = 20;
 * @return {number}
 */
proto.buf.validate.StringRules.prototype.getLenBytes = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 20, 0));
};


/**
 * @param {number} value
 * @return {!proto.buf.validate.StringRules} returns this
 */
proto.buf.validate.StringRules.prototype.setLenBytes = function(value) {
  return jspb.Message.setField(this, 20, value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.StringRules} returns this
 */
proto.buf.validate.StringRules.prototype.clearLenBytes = function() {
  return jspb.Message.setField(this, 20, undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.StringRules.prototype.hasLenBytes = function() {
  return jspb.Message.getField(this, 20) != null;
};


/**
 * optional uint64 min_bytes = 4;
 * @return {number}
 */
proto.buf.validate.StringRules.prototype.getMinBytes = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 4, 0));
};


/**
 * @param {number} value
 * @return {!proto.buf.validate.StringRules} returns this
 */
proto.buf.validate.StringRules.prototype.setMinBytes = function(value) {
  return jspb.Message.setField(this, 4, value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.StringRules} returns this
 */
proto.buf.validate.StringRules.prototype.clearMinBytes = function() {
  return jspb.Message.setField(this, 4, undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.StringRules.prototype.hasMinBytes = function() {
  return jspb.Message.getField(this, 4) != null;
};


/**
 * optional uint64 max_bytes = 5;
 * @return {number}
 */
proto.buf.validate.StringRules.prototype.getMaxBytes = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 5, 0));
};


/**
 * @param {number} value
 * @return {!proto.buf.validate.StringRules} returns this
 */
proto.buf.validate.StringRules.prototype.setMaxBytes = function(value) {
  return jspb.Message.setField(this, 5, value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.StringRules} returns this
 */
proto.buf.validate.StringRules.prototype.clearMaxBytes = function() {
  return jspb.Message.setField(this, 5, undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.StringRules.prototype.hasMaxBytes = function() {
  return jspb.Message.getField(this, 5) != null;
};


/**
 * optional string pattern = 6;
 * @return {string}
 */
proto.buf.validate.StringRules.prototype.getPattern = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 6, ""));
};


/**
 * @param {string} value
 * @return {!proto.buf.validate.StringRules} returns this
 */
proto.buf.validate.StringRules.prototype.setPattern = function(value) {
  return jspb.Message.setField(this, 6, value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.StringRules} returns this
 */
proto.buf.validate.StringRules.prototype.clearPattern = function() {
  return jspb.Message.setField(this, 6, undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.StringRules.prototype.hasPattern = function() {
  return jspb.Message.getField(this, 6) != null;
};


/**
 * optional string prefix = 7;
 * @return {string}
 */
proto.buf.validate.StringRules.prototype.getPrefix = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 7, ""));
};


/**
 * @param {string} value
 * @return {!proto.buf.validate.StringRules} returns this
 */
proto.buf.validate.StringRules.prototype.setPrefix = function(value) {
  return jspb.Message.setField(this, 7, value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.StringRules} returns this
 */
proto.buf.validate.StringRules.prototype.clearPrefix = function() {
  return jspb.Message.setField(this, 7, undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.StringRules.prototype.hasPrefix = function() {
  return jspb.Message.getField(this, 7) != null;
};


/**
 * optional string suffix = 8;
 * @return {string}
 */
proto.buf.validate.StringRules.prototype.getSuffix = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 8, ""));
};


/**
 * @param {string} value
 * @return {!proto.buf.validate.StringRules} returns this
 */
proto.buf.validate.StringRules.prototype.setSuffix = function(value) {
  return jspb.Message.setField(this, 8, value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.StringRules} returns this
 */
proto.buf.validate.StringRules.prototype.clearSuffix = function() {
  return jspb.Message.setField(this, 8, undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.StringRules.prototype.hasSuffix = function() {
  return jspb.Message.getField(this, 8) != null;
};


/**
 * optional string contains = 9;
 * @return {string}
 */
proto.buf.validate.StringRules.prototype.getContains = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 9, ""));
};


/**
 * @param {string} value
 * @return {!proto.buf.validate.StringRules} returns this
 */
proto.buf.validate.StringRules.prototype.setContains = function(value) {
  return jspb.Message.setField(this, 9, value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.StringRules} returns this
 */
proto.buf.validate.StringRules.prototype.clearContains = function() {
  return jspb.Message.setField(this, 9, undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.StringRules.prototype.hasContains = function() {
  return jspb.Message.getField(this, 9) != null;
};


/**
 * optional string not_contains = 23;
 * @return {string}
 */
proto.buf.validate.StringRules.prototype.getNotContains = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 23, ""));
};


/**
 * @param {string} value
 * @return {!proto.buf.validate.StringRules} returns this
 */
proto.buf.validate.StringRules.prototype.setNotContains = function(value) {
  return jspb.Message.setField(this, 23, value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.StringRules} returns this
 */
proto.buf.validate.StringRules.prototype.clearNotContains = function() {
  return jspb.Message.setField(this, 23, undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.StringRules.prototype.hasNotContains = function() {
  return jspb.Message.getField(this, 23) != null;
};


/**
 * repeated string in = 10;
 * @return {!Array<string>}
 */
proto.buf.validate.StringRules.prototype.getInList = function() {
  return /** @type {!Array<string>} */ (jspb.Message.getRepeatedField(this, 10));
};


/**
 * @param {!Array<string>} value
 * @return {!proto.buf.validate.StringRules} returns this
 */
proto.buf.validate.StringRules.prototype.setInList = function(value) {
  return jspb.Message.setField(this, 10, value || []);
};


/**
 * @param {string} value
 * @param {number=} opt_index
 * @return {!proto.buf.validate.StringRules} returns this
 */
proto.buf.validate.StringRules.prototype.addIn = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 10, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.buf.validate.StringRules} returns this
 */
proto.buf.validate.StringRules.prototype.clearInList = function() {
  return this.setInList([]);
};


/**
 * repeated string not_in = 11;
 * @return {!Array<string>}
 */
proto.buf.validate.StringRules.prototype.getNotInList = function() {
  return /** @type {!Array<string>} */ (jspb.Message.getRepeatedField(this, 11));
};


/**
 * @param {!Array<string>} value
 * @return {!proto.buf.validate.StringRules} returns this
 */
proto.buf.validate.StringRules.prototype.setNotInList = function(value) {
  return jspb.Message.setField(this, 11, value || []);
};


/**
 * @param {string} value
 * @param {number=} opt_index
 * @return {!proto.buf.validate.StringRules} returns this
 */
proto.buf.validate.StringRules.prototype.addNotIn = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 11, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.buf.validate.StringRules} returns this
 */
proto.buf.validate.StringRules.prototype.clearNotInList = function() {
  return this.setNotInList([]);
};


/**
 * optional bool email = 12;
 * @return {boolean}
 */
proto.buf.validate.StringRules.prototype.getEmail = function() {
  return /** @type {boolean} */ (jspb.Message.getBooleanFieldWithDefault(this, 12, false));
};


/**
 * @param {boolean} value
 * @return {!proto.buf.validate.StringRules} returns this
 */
proto.buf.validate.StringRules.prototype.setEmail = function(value) {
  return jspb.Message.setOneofField(this, 12, proto.buf.validate.StringRules.oneofGroups_[0], value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.StringRules} returns this
 */
proto.buf.validate.StringRules.prototype.clearEmail = function() {
  return jspb.Message.setOneofField(this, 12, proto.buf.validate.StringRules.oneofGroups_[0], undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.StringRules.prototype.hasEmail = function() {
  return jspb.Message.getField(this, 12) != null;
};


/**
 * optional bool hostname = 13;
 * @return {boolean}
 */
proto.buf.validate.StringRules.prototype.getHostname = function() {
  return /** @type {boolean} */ (jspb.Message.getBooleanFieldWithDefault(this, 13, false));
};


/**
 * @param {boolean} value
 * @return {!proto.buf.validate.StringRules} returns this
 */
proto.buf.validate.StringRules.prototype.setHostname = function(value) {
  return jspb.Message.setOneofField(this, 13, proto.buf.validate.StringRules.oneofGroups_[0], value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.StringRules} returns this
 */
proto.buf.validate.StringRules.prototype.clearHostname = function() {
  return jspb.Message.setOneofField(this, 13, proto.buf.validate.StringRules.oneofGroups_[0], undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.StringRules.prototype.hasHostname = function() {
  return jspb.Message.getField(this, 13) != null;
};


/**
 * optional bool ip = 14;
 * @return {boolean}
 */
proto.buf.validate.StringRules.prototype.getIp = function() {
  return /** @type {boolean} */ (jspb.Message.getBooleanFieldWithDefault(this, 14, false));
};


/**
 * @param {boolean} value
 * @return {!proto.buf.validate.StringRules} returns this
 */
proto.buf.validate.StringRules.prototype.setIp = function(value) {
  return jspb.Message.setOneofField(this, 14, proto.buf.validate.StringRules.oneofGroups_[0], value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.StringRules} returns this
 */
proto.buf.validate.StringRules.prototype.clearIp = function() {
  return jspb.Message.setOneofField(this, 14, proto.buf.validate.StringRules.oneofGroups_[0], undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.StringRules.prototype.hasIp = function() {
  return jspb.Message.getField(this, 14) != null;
};


/**
 * optional bool ipv4 = 15;
 * @return {boolean}
 */
proto.buf.validate.StringRules.prototype.getIpv4 = function() {
  return /** @type {boolean} */ (jspb.Message.getBooleanFieldWithDefault(this, 15, false));
};


/**
 * @param {boolean} value
 * @return {!proto.buf.validate.StringRules} returns this
 */
proto.buf.validate.StringRules.prototype.setIpv4 = function(value) {
  return jspb.Message.setOneofField(this, 15, proto.buf.validate.StringRules.oneofGroups_[0], value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.StringRules} returns this
 */
proto.buf.validate.StringRules.prototype.clearIpv4 = function() {
  return jspb.Message.setOneofField(this, 15, proto.buf.validate.StringRules.oneofGroups_[0], undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.StringRules.prototype.hasIpv4 = function() {
  return jspb.Message.getField(this, 15) != null;
};


/**
 * optional bool ipv6 = 16;
 * @return {boolean}
 */
proto.buf.validate.StringRules.prototype.getIpv6 = function() {
  return /** @type {boolean} */ (jspb.Message.getBooleanFieldWithDefault(this, 16, false));
};


/**
 * @param {boolean} value
 * @return {!proto.buf.validate.StringRules} returns this
 */
proto.buf.validate.StringRules.prototype.setIpv6 = function(value) {
  return jspb.Message.setOneofField(this, 16, proto.buf.validate.StringRules.oneofGroups_[0], value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.StringRules} returns this
 */
proto.buf.validate.StringRules.prototype.clearIpv6 = function() {
  return jspb.Message.setOneofField(this, 16, proto.buf.validate.StringRules.oneofGroups_[0], undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.StringRules.prototype.hasIpv6 = function() {
  return jspb.Message.getField(this, 16) != null;
};


/**
 * optional bool uri = 17;
 * @return {boolean}
 */
proto.buf.validate.StringRules.prototype.getUri = function() {
  return /** @type {boolean} */ (jspb.Message.getBooleanFieldWithDefault(this, 17, false));
};


/**
 * @param {boolean} value
 * @return {!proto.buf.validate.StringRules} returns this
 */
proto.buf.validate.StringRules.prototype.setUri = function(value) {
  return jspb.Message.setOneofField(this, 17, proto.buf.validate.StringRules.oneofGroups_[0], value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.StringRules} returns this
 */
proto.buf.validate.StringRules.prototype.clearUri = function() {
  return jspb.Message.setOneofField(this, 17, proto.buf.validate.StringRules.oneofGroups_[0], undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.StringRules.prototype.hasUri = function() {
  return jspb.Message.getField(this, 17) != null;
};


/**
 * optional bool uri_ref = 18;
 * @return {boolean}
 */
proto.buf.validate.StringRules.prototype.getUriRef = function() {
  return /** @type {boolean} */ (jspb.Message.getBooleanFieldWithDefault(this, 18, false));
};


/**
 * @param {boolean} value
 * @return {!proto.buf.validate.StringRules} returns this
 */
proto.buf.validate.StringRules.prototype.setUriRef = function(value) {
  return jspb.Message.setOneofField(this, 18, proto.buf.validate.StringRules.oneofGroups_[0], value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.StringRules} returns this
 */
proto.buf.validate.StringRules.prototype.clearUriRef = function() {
  return jspb.Message.setOneofField(this, 18, proto.buf.validate.StringRules.oneofGroups_[0], undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.StringRules.prototype.hasUriRef = function() {
  return jspb.Message.getField(this, 18) != null;
};


/**
 * optional bool address = 21;
 * @return {boolean}
 */
proto.buf.validate.StringRules.prototype.getAddress = function() {
  return /** @type {boolean} */ (jspb.Message.getBooleanFieldWithDefault(this, 21, false));
};


/**
 * @param {boolean} value
 * @return {!proto.buf.validate.StringRules} returns this
 */
proto.buf.validate.StringRules.prototype.setAddress = function(value) {
  return jspb.Message.setOneofField(this, 21, proto.buf.validate.StringRules.oneofGroups_[0], value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.StringRules} returns this
 */
proto.buf.validate.StringRules.prototype.clearAddress = function() {
  return jspb.Message.setOneofField(this, 21, proto.buf.validate.StringRules.oneofGroups_[0], undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.StringRules.prototype.hasAddress = function() {
  return jspb.Message.getField(this, 21) != null;
};


/**
 * optional bool uuid = 22;
 * @return {boolean}
 */
proto.buf.validate.StringRules.prototype.getUuid = function() {
  return /** @type {boolean} */ (jspb.Message.getBooleanFieldWithDefault(this, 22, false));
};


/**
 * @param {boolean} value
 * @return {!proto.buf.validate.StringRules} returns this
 */
proto.buf.validate.StringRules.prototype.setUuid = function(value) {
  return jspb.Message.setOneofField(this, 22, proto.buf.validate.StringRules.oneofGroups_[0], value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.StringRules} returns this
 */
proto.buf.validate.StringRules.prototype.clearUuid = function() {
  return jspb.Message.setOneofField(this, 22, proto.buf.validate.StringRules.oneofGroups_[0], undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.StringRules.prototype.hasUuid = function() {
  return jspb.Message.getField(this, 22) != null;
};


/**
 * optional bool tuuid = 33;
 * @return {boolean}
 */
proto.buf.validate.StringRules.prototype.getTuuid = function() {
  return /** @type {boolean} */ (jspb.Message.getBooleanFieldWithDefault(this, 33, false));
};


/**
 * @param {boolean} value
 * @return {!proto.buf.validate.StringRules} returns this
 */
proto.buf.validate.StringRules.prototype.setTuuid = function(value) {
  return jspb.Message.setOneofField(this, 33, proto.buf.validate.StringRules.oneofGroups_[0], value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.StringRules} returns this
 */
proto.buf.validate.StringRules.prototype.clearTuuid = function() {
  return jspb.Message.setOneofField(this, 33, proto.buf.validate.StringRules.oneofGroups_[0], undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.StringRules.prototype.hasTuuid = function() {
  return jspb.Message.getField(this, 33) != null;
};


/**
 * optional bool ip_with_prefixlen = 26;
 * @return {boolean}
 */
proto.buf.validate.StringRules.prototype.getIpWithPrefixlen = function() {
  return /** @type {boolean} */ (jspb.Message.getBooleanFieldWithDefault(this, 26, false));
};


/**
 * @param {boolean} value
 * @return {!proto.buf.validate.StringRules} returns this
 */
proto.buf.validate.StringRules.prototype.setIpWithPrefixlen = function(value) {
  return jspb.Message.setOneofField(this, 26, proto.buf.validate.StringRules.oneofGroups_[0], value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.StringRules} returns this
 */
proto.buf.validate.StringRules.prototype.clearIpWithPrefixlen = function() {
  return jspb.Message.setOneofField(this, 26, proto.buf.validate.StringRules.oneofGroups_[0], undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.StringRules.prototype.hasIpWithPrefixlen = function() {
  return jspb.Message.getField(this, 26) != null;
};


/**
 * optional bool ipv4_with_prefixlen = 27;
 * @return {boolean}
 */
proto.buf.validate.StringRules.prototype.getIpv4WithPrefixlen = function() {
  return /** @type {boolean} */ (jspb.Message.getBooleanFieldWithDefault(this, 27, false));
};


/**
 * @param {boolean} value
 * @return {!proto.buf.validate.StringRules} returns this
 */
proto.buf.validate.StringRules.prototype.setIpv4WithPrefixlen = function(value) {
  return jspb.Message.setOneofField(this, 27, proto.buf.validate.StringRules.oneofGroups_[0], value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.StringRules} returns this
 */
proto.buf.validate.StringRules.prototype.clearIpv4WithPrefixlen = function() {
  return jspb.Message.setOneofField(this, 27, proto.buf.validate.StringRules.oneofGroups_[0], undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.StringRules.prototype.hasIpv4WithPrefixlen = function() {
  return jspb.Message.getField(this, 27) != null;
};


/**
 * optional bool ipv6_with_prefixlen = 28;
 * @return {boolean}
 */
proto.buf.validate.StringRules.prototype.getIpv6WithPrefixlen = function() {
  return /** @type {boolean} */ (jspb.Message.getBooleanFieldWithDefault(this, 28, false));
};


/**
 * @param {boolean} value
 * @return {!proto.buf.validate.StringRules} returns this
 */
proto.buf.validate.StringRules.prototype.setIpv6WithPrefixlen = function(value) {
  return jspb.Message.setOneofField(this, 28, proto.buf.validate.StringRules.oneofGroups_[0], value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.StringRules} returns this
 */
proto.buf.validate.StringRules.prototype.clearIpv6WithPrefixlen = function() {
  return jspb.Message.setOneofField(this, 28, proto.buf.validate.StringRules.oneofGroups_[0], undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.StringRules.prototype.hasIpv6WithPrefixlen = function() {
  return jspb.Message.getField(this, 28) != null;
};


/**
 * optional bool ip_prefix = 29;
 * @return {boolean}
 */
proto.buf.validate.StringRules.prototype.getIpPrefix = function() {
  return /** @type {boolean} */ (jspb.Message.getBooleanFieldWithDefault(this, 29, false));
};


/**
 * @param {boolean} value
 * @return {!proto.buf.validate.StringRules} returns this
 */
proto.buf.validate.StringRules.prototype.setIpPrefix = function(value) {
  return jspb.Message.setOneofField(this, 29, proto.buf.validate.StringRules.oneofGroups_[0], value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.StringRules} returns this
 */
proto.buf.validate.StringRules.prototype.clearIpPrefix = function() {
  return jspb.Message.setOneofField(this, 29, proto.buf.validate.StringRules.oneofGroups_[0], undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.StringRules.prototype.hasIpPrefix = function() {
  return jspb.Message.getField(this, 29) != null;
};


/**
 * optional bool ipv4_prefix = 30;
 * @return {boolean}
 */
proto.buf.validate.StringRules.prototype.getIpv4Prefix = function() {
  return /** @type {boolean} */ (jspb.Message.getBooleanFieldWithDefault(this, 30, false));
};


/**
 * @param {boolean} value
 * @return {!proto.buf.validate.StringRules} returns this
 */
proto.buf.validate.StringRules.prototype.setIpv4Prefix = function(value) {
  return jspb.Message.setOneofField(this, 30, proto.buf.validate.StringRules.oneofGroups_[0], value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.StringRules} returns this
 */
proto.buf.validate.StringRules.prototype.clearIpv4Prefix = function() {
  return jspb.Message.setOneofField(this, 30, proto.buf.validate.StringRules.oneofGroups_[0], undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.StringRules.prototype.hasIpv4Prefix = function() {
  return jspb.Message.getField(this, 30) != null;
};


/**
 * optional bool ipv6_prefix = 31;
 * @return {boolean}
 */
proto.buf.validate.StringRules.prototype.getIpv6Prefix = function() {
  return /** @type {boolean} */ (jspb.Message.getBooleanFieldWithDefault(this, 31, false));
};


/**
 * @param {boolean} value
 * @return {!proto.buf.validate.StringRules} returns this
 */
proto.buf.validate.StringRules.prototype.setIpv6Prefix = function(value) {
  return jspb.Message.setOneofField(this, 31, proto.buf.validate.StringRules.oneofGroups_[0], value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.StringRules} returns this
 */
proto.buf.validate.StringRules.prototype.clearIpv6Prefix = function() {
  return jspb.Message.setOneofField(this, 31, proto.buf.validate.StringRules.oneofGroups_[0], undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.StringRules.prototype.hasIpv6Prefix = function() {
  return jspb.Message.getField(this, 31) != null;
};


/**
 * optional bool host_and_port = 32;
 * @return {boolean}
 */
proto.buf.validate.StringRules.prototype.getHostAndPort = function() {
  return /** @type {boolean} */ (jspb.Message.getBooleanFieldWithDefault(this, 32, false));
};


/**
 * @param {boolean} value
 * @return {!proto.buf.validate.StringRules} returns this
 */
proto.buf.validate.StringRules.prototype.setHostAndPort = function(value) {
  return jspb.Message.setOneofField(this, 32, proto.buf.validate.StringRules.oneofGroups_[0], value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.StringRules} returns this
 */
proto.buf.validate.StringRules.prototype.clearHostAndPort = function() {
  return jspb.Message.setOneofField(this, 32, proto.buf.validate.StringRules.oneofGroups_[0], undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.StringRules.prototype.hasHostAndPort = function() {
  return jspb.Message.getField(this, 32) != null;
};


/**
 * optional KnownRegex well_known_regex = 24;
 * @return {!proto.buf.validate.KnownRegex}
 */
proto.buf.validate.StringRules.prototype.getWellKnownRegex = function() {
  return /** @type {!proto.buf.validate.KnownRegex} */ (jspb.Message.getFieldWithDefault(this, 24, 0));
};


/**
 * @param {!proto.buf.validate.KnownRegex} value
 * @return {!proto.buf.validate.StringRules} returns this
 */
proto.buf.validate.StringRules.prototype.setWellKnownRegex = function(value) {
  return jspb.Message.setOneofField(this, 24, proto.buf.validate.StringRules.oneofGroups_[0], value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.StringRules} returns this
 */
proto.buf.validate.StringRules.prototype.clearWellKnownRegex = function() {
  return jspb.Message.setOneofField(this, 24, proto.buf.validate.StringRules.oneofGroups_[0], undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.StringRules.prototype.hasWellKnownRegex = function() {
  return jspb.Message.getField(this, 24) != null;
};


/**
 * optional bool strict = 25;
 * @return {boolean}
 */
proto.buf.validate.StringRules.prototype.getStrict = function() {
  return /** @type {boolean} */ (jspb.Message.getBooleanFieldWithDefault(this, 25, false));
};


/**
 * @param {boolean} value
 * @return {!proto.buf.validate.StringRules} returns this
 */
proto.buf.validate.StringRules.prototype.setStrict = function(value) {
  return jspb.Message.setField(this, 25, value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.StringRules} returns this
 */
proto.buf.validate.StringRules.prototype.clearStrict = function() {
  return jspb.Message.setField(this, 25, undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.StringRules.prototype.hasStrict = function() {
  return jspb.Message.getField(this, 25) != null;
};


/**
 * repeated string example = 34;
 * @return {!Array<string>}
 */
proto.buf.validate.StringRules.prototype.getExampleList = function() {
  return /** @type {!Array<string>} */ (jspb.Message.getRepeatedField(this, 34));
};


/**
 * @param {!Array<string>} value
 * @return {!proto.buf.validate.StringRules} returns this
 */
proto.buf.validate.StringRules.prototype.setExampleList = function(value) {
  return jspb.Message.setField(this, 34, value || []);
};


/**
 * @param {string} value
 * @param {number=} opt_index
 * @return {!proto.buf.validate.StringRules} returns this
 */
proto.buf.validate.StringRules.prototype.addExample = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 34, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.buf.validate.StringRules} returns this
 */
proto.buf.validate.StringRules.prototype.clearExampleList = function() {
  return this.setExampleList([]);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.buf.validate.BytesRules.repeatedFields_ = [8,9,14];

/**
 * Oneof group definitions for this message. Each group defines the field
 * numbers belonging to that group. When of these fields' value is set, all
 * other fields in the group are cleared. During deserialization, if multiple
 * fields are encountered for a group, only the last value seen will be kept.
 * @private {!Array<!Array<number>>}
 * @const
 */
proto.buf.validate.BytesRules.oneofGroups_ = [[10,11,12]];

/**
 * @enum {number}
 */
proto.buf.validate.BytesRules.WellKnownCase = {
  WELL_KNOWN_NOT_SET: 0,
  IP: 10,
  IPV4: 11,
  IPV6: 12
};

/**
 * @return {proto.buf.validate.BytesRules.WellKnownCase}
 */
proto.buf.validate.BytesRules.prototype.getWellKnownCase = function() {
  return /** @type {proto.buf.validate.BytesRules.WellKnownCase} */(jspb.Message.computeOneofCase(this, proto.buf.validate.BytesRules.oneofGroups_[0]));
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
proto.buf.validate.BytesRules.prototype.toObject = function(opt_includeInstance) {
  return proto.buf.validate.BytesRules.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.buf.validate.BytesRules} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.buf.validate.BytesRules.toObject = function(includeInstance, msg) {
  var f, obj = {
pb_const: msg.getConst_asB64(),
len: (f = jspb.Message.getField(msg, 13)) == null ? undefined : f,
minLen: (f = jspb.Message.getField(msg, 2)) == null ? undefined : f,
maxLen: (f = jspb.Message.getField(msg, 3)) == null ? undefined : f,
pattern: (f = jspb.Message.getField(msg, 4)) == null ? undefined : f,
prefix: msg.getPrefix_asB64(),
suffix: msg.getSuffix_asB64(),
contains: msg.getContains_asB64(),
inList: msg.getInList_asB64(),
notInList: msg.getNotInList_asB64(),
ip: (f = jspb.Message.getBooleanField(msg, 10)) == null ? undefined : f,
ipv4: (f = jspb.Message.getBooleanField(msg, 11)) == null ? undefined : f,
ipv6: (f = jspb.Message.getBooleanField(msg, 12)) == null ? undefined : f,
exampleList: msg.getExampleList_asB64()
  };

  jspb.Message.toObjectExtension(/** @type {!jspb.Message} */ (msg), obj,
      proto.buf.validate.BytesRules.extensions, proto.buf.validate.BytesRules.prototype.getExtension,
      includeInstance);
  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.buf.validate.BytesRules}
 */
proto.buf.validate.BytesRules.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.buf.validate.BytesRules;
  return proto.buf.validate.BytesRules.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.buf.validate.BytesRules} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.buf.validate.BytesRules}
 */
proto.buf.validate.BytesRules.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {!Uint8Array} */ (reader.readBytes());
      msg.setConst(value);
      break;
    case 13:
      var value = /** @type {number} */ (reader.readUint64());
      msg.setLen(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readUint64());
      msg.setMinLen(value);
      break;
    case 3:
      var value = /** @type {number} */ (reader.readUint64());
      msg.setMaxLen(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setPattern(value);
      break;
    case 5:
      var value = /** @type {!Uint8Array} */ (reader.readBytes());
      msg.setPrefix(value);
      break;
    case 6:
      var value = /** @type {!Uint8Array} */ (reader.readBytes());
      msg.setSuffix(value);
      break;
    case 7:
      var value = /** @type {!Uint8Array} */ (reader.readBytes());
      msg.setContains(value);
      break;
    case 8:
      var value = /** @type {!Uint8Array} */ (reader.readBytes());
      msg.addIn(value);
      break;
    case 9:
      var value = /** @type {!Uint8Array} */ (reader.readBytes());
      msg.addNotIn(value);
      break;
    case 10:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setIp(value);
      break;
    case 11:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setIpv4(value);
      break;
    case 12:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setIpv6(value);
      break;
    case 14:
      var value = /** @type {!Uint8Array} */ (reader.readBytes());
      msg.addExample(value);
      break;
    default:
      jspb.Message.readBinaryExtension(msg, reader,
        proto.buf.validate.BytesRules.extensionsBinary,
        proto.buf.validate.BytesRules.prototype.getExtension,
        proto.buf.validate.BytesRules.prototype.setExtension);
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.buf.validate.BytesRules.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.buf.validate.BytesRules.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.buf.validate.BytesRules} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.buf.validate.BytesRules.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = /** @type {!(string|Uint8Array)} */ (jspb.Message.getField(message, 1));
  if (f != null) {
    writer.writeBytes(
      1,
      f
    );
  }
  f = /** @type {number} */ (jspb.Message.getField(message, 13));
  if (f != null) {
    writer.writeUint64(
      13,
      f
    );
  }
  f = /** @type {number} */ (jspb.Message.getField(message, 2));
  if (f != null) {
    writer.writeUint64(
      2,
      f
    );
  }
  f = /** @type {number} */ (jspb.Message.getField(message, 3));
  if (f != null) {
    writer.writeUint64(
      3,
      f
    );
  }
  f = /** @type {string} */ (jspb.Message.getField(message, 4));
  if (f != null) {
    writer.writeString(
      4,
      f
    );
  }
  f = /** @type {!(string|Uint8Array)} */ (jspb.Message.getField(message, 5));
  if (f != null) {
    writer.writeBytes(
      5,
      f
    );
  }
  f = /** @type {!(string|Uint8Array)} */ (jspb.Message.getField(message, 6));
  if (f != null) {
    writer.writeBytes(
      6,
      f
    );
  }
  f = /** @type {!(string|Uint8Array)} */ (jspb.Message.getField(message, 7));
  if (f != null) {
    writer.writeBytes(
      7,
      f
    );
  }
  f = message.getInList_asU8();
  if (f.length > 0) {
    writer.writeRepeatedBytes(
      8,
      f
    );
  }
  f = message.getNotInList_asU8();
  if (f.length > 0) {
    writer.writeRepeatedBytes(
      9,
      f
    );
  }
  f = /** @type {boolean} */ (jspb.Message.getField(message, 10));
  if (f != null) {
    writer.writeBool(
      10,
      f
    );
  }
  f = /** @type {boolean} */ (jspb.Message.getField(message, 11));
  if (f != null) {
    writer.writeBool(
      11,
      f
    );
  }
  f = /** @type {boolean} */ (jspb.Message.getField(message, 12));
  if (f != null) {
    writer.writeBool(
      12,
      f
    );
  }
  f = message.getExampleList_asU8();
  if (f.length > 0) {
    writer.writeRepeatedBytes(
      14,
      f
    );
  }
  jspb.Message.serializeBinaryExtensions(message, writer,
    proto.buf.validate.BytesRules.extensionsBinary, proto.buf.validate.BytesRules.prototype.getExtension);
};


/**
 * optional bytes const = 1;
 * @return {!(string|Uint8Array)}
 */
proto.buf.validate.BytesRules.prototype.getConst = function() {
  return /** @type {!(string|Uint8Array)} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * optional bytes const = 1;
 * This is a type-conversion wrapper around `getConst()`
 * @return {string}
 */
proto.buf.validate.BytesRules.prototype.getConst_asB64 = function() {
  return /** @type {string} */ (jspb.Message.bytesAsB64(
      this.getConst()));
};


/**
 * optional bytes const = 1;
 * Note that Uint8Array is not supported on all browsers.
 * @see http://caniuse.com/Uint8Array
 * This is a type-conversion wrapper around `getConst()`
 * @return {!Uint8Array}
 */
proto.buf.validate.BytesRules.prototype.getConst_asU8 = function() {
  return /** @type {!Uint8Array} */ (jspb.Message.bytesAsU8(
      this.getConst()));
};


/**
 * @param {!(string|Uint8Array)} value
 * @return {!proto.buf.validate.BytesRules} returns this
 */
proto.buf.validate.BytesRules.prototype.setConst = function(value) {
  return jspb.Message.setField(this, 1, value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.BytesRules} returns this
 */
proto.buf.validate.BytesRules.prototype.clearConst = function() {
  return jspb.Message.setField(this, 1, undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.BytesRules.prototype.hasConst = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional uint64 len = 13;
 * @return {number}
 */
proto.buf.validate.BytesRules.prototype.getLen = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 13, 0));
};


/**
 * @param {number} value
 * @return {!proto.buf.validate.BytesRules} returns this
 */
proto.buf.validate.BytesRules.prototype.setLen = function(value) {
  return jspb.Message.setField(this, 13, value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.BytesRules} returns this
 */
proto.buf.validate.BytesRules.prototype.clearLen = function() {
  return jspb.Message.setField(this, 13, undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.BytesRules.prototype.hasLen = function() {
  return jspb.Message.getField(this, 13) != null;
};


/**
 * optional uint64 min_len = 2;
 * @return {number}
 */
proto.buf.validate.BytesRules.prototype.getMinLen = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/**
 * @param {number} value
 * @return {!proto.buf.validate.BytesRules} returns this
 */
proto.buf.validate.BytesRules.prototype.setMinLen = function(value) {
  return jspb.Message.setField(this, 2, value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.BytesRules} returns this
 */
proto.buf.validate.BytesRules.prototype.clearMinLen = function() {
  return jspb.Message.setField(this, 2, undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.BytesRules.prototype.hasMinLen = function() {
  return jspb.Message.getField(this, 2) != null;
};


/**
 * optional uint64 max_len = 3;
 * @return {number}
 */
proto.buf.validate.BytesRules.prototype.getMaxLen = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 3, 0));
};


/**
 * @param {number} value
 * @return {!proto.buf.validate.BytesRules} returns this
 */
proto.buf.validate.BytesRules.prototype.setMaxLen = function(value) {
  return jspb.Message.setField(this, 3, value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.BytesRules} returns this
 */
proto.buf.validate.BytesRules.prototype.clearMaxLen = function() {
  return jspb.Message.setField(this, 3, undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.BytesRules.prototype.hasMaxLen = function() {
  return jspb.Message.getField(this, 3) != null;
};


/**
 * optional string pattern = 4;
 * @return {string}
 */
proto.buf.validate.BytesRules.prototype.getPattern = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/**
 * @param {string} value
 * @return {!proto.buf.validate.BytesRules} returns this
 */
proto.buf.validate.BytesRules.prototype.setPattern = function(value) {
  return jspb.Message.setField(this, 4, value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.BytesRules} returns this
 */
proto.buf.validate.BytesRules.prototype.clearPattern = function() {
  return jspb.Message.setField(this, 4, undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.BytesRules.prototype.hasPattern = function() {
  return jspb.Message.getField(this, 4) != null;
};


/**
 * optional bytes prefix = 5;
 * @return {!(string|Uint8Array)}
 */
proto.buf.validate.BytesRules.prototype.getPrefix = function() {
  return /** @type {!(string|Uint8Array)} */ (jspb.Message.getFieldWithDefault(this, 5, ""));
};


/**
 * optional bytes prefix = 5;
 * This is a type-conversion wrapper around `getPrefix()`
 * @return {string}
 */
proto.buf.validate.BytesRules.prototype.getPrefix_asB64 = function() {
  return /** @type {string} */ (jspb.Message.bytesAsB64(
      this.getPrefix()));
};


/**
 * optional bytes prefix = 5;
 * Note that Uint8Array is not supported on all browsers.
 * @see http://caniuse.com/Uint8Array
 * This is a type-conversion wrapper around `getPrefix()`
 * @return {!Uint8Array}
 */
proto.buf.validate.BytesRules.prototype.getPrefix_asU8 = function() {
  return /** @type {!Uint8Array} */ (jspb.Message.bytesAsU8(
      this.getPrefix()));
};


/**
 * @param {!(string|Uint8Array)} value
 * @return {!proto.buf.validate.BytesRules} returns this
 */
proto.buf.validate.BytesRules.prototype.setPrefix = function(value) {
  return jspb.Message.setField(this, 5, value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.BytesRules} returns this
 */
proto.buf.validate.BytesRules.prototype.clearPrefix = function() {
  return jspb.Message.setField(this, 5, undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.BytesRules.prototype.hasPrefix = function() {
  return jspb.Message.getField(this, 5) != null;
};


/**
 * optional bytes suffix = 6;
 * @return {!(string|Uint8Array)}
 */
proto.buf.validate.BytesRules.prototype.getSuffix = function() {
  return /** @type {!(string|Uint8Array)} */ (jspb.Message.getFieldWithDefault(this, 6, ""));
};


/**
 * optional bytes suffix = 6;
 * This is a type-conversion wrapper around `getSuffix()`
 * @return {string}
 */
proto.buf.validate.BytesRules.prototype.getSuffix_asB64 = function() {
  return /** @type {string} */ (jspb.Message.bytesAsB64(
      this.getSuffix()));
};


/**
 * optional bytes suffix = 6;
 * Note that Uint8Array is not supported on all browsers.
 * @see http://caniuse.com/Uint8Array
 * This is a type-conversion wrapper around `getSuffix()`
 * @return {!Uint8Array}
 */
proto.buf.validate.BytesRules.prototype.getSuffix_asU8 = function() {
  return /** @type {!Uint8Array} */ (jspb.Message.bytesAsU8(
      this.getSuffix()));
};


/**
 * @param {!(string|Uint8Array)} value
 * @return {!proto.buf.validate.BytesRules} returns this
 */
proto.buf.validate.BytesRules.prototype.setSuffix = function(value) {
  return jspb.Message.setField(this, 6, value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.BytesRules} returns this
 */
proto.buf.validate.BytesRules.prototype.clearSuffix = function() {
  return jspb.Message.setField(this, 6, undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.BytesRules.prototype.hasSuffix = function() {
  return jspb.Message.getField(this, 6) != null;
};


/**
 * optional bytes contains = 7;
 * @return {!(string|Uint8Array)}
 */
proto.buf.validate.BytesRules.prototype.getContains = function() {
  return /** @type {!(string|Uint8Array)} */ (jspb.Message.getFieldWithDefault(this, 7, ""));
};


/**
 * optional bytes contains = 7;
 * This is a type-conversion wrapper around `getContains()`
 * @return {string}
 */
proto.buf.validate.BytesRules.prototype.getContains_asB64 = function() {
  return /** @type {string} */ (jspb.Message.bytesAsB64(
      this.getContains()));
};


/**
 * optional bytes contains = 7;
 * Note that Uint8Array is not supported on all browsers.
 * @see http://caniuse.com/Uint8Array
 * This is a type-conversion wrapper around `getContains()`
 * @return {!Uint8Array}
 */
proto.buf.validate.BytesRules.prototype.getContains_asU8 = function() {
  return /** @type {!Uint8Array} */ (jspb.Message.bytesAsU8(
      this.getContains()));
};


/**
 * @param {!(string|Uint8Array)} value
 * @return {!proto.buf.validate.BytesRules} returns this
 */
proto.buf.validate.BytesRules.prototype.setContains = function(value) {
  return jspb.Message.setField(this, 7, value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.BytesRules} returns this
 */
proto.buf.validate.BytesRules.prototype.clearContains = function() {
  return jspb.Message.setField(this, 7, undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.BytesRules.prototype.hasContains = function() {
  return jspb.Message.getField(this, 7) != null;
};


/**
 * repeated bytes in = 8;
 * @return {!(Array<!Uint8Array>|Array<string>)}
 */
proto.buf.validate.BytesRules.prototype.getInList = function() {
  return /** @type {!(Array<!Uint8Array>|Array<string>)} */ (jspb.Message.getRepeatedField(this, 8));
};


/**
 * repeated bytes in = 8;
 * This is a type-conversion wrapper around `getInList()`
 * @return {!Array<string>}
 */
proto.buf.validate.BytesRules.prototype.getInList_asB64 = function() {
  return /** @type {!Array<string>} */ (jspb.Message.bytesListAsB64(
      this.getInList()));
};


/**
 * repeated bytes in = 8;
 * Note that Uint8Array is not supported on all browsers.
 * @see http://caniuse.com/Uint8Array
 * This is a type-conversion wrapper around `getInList()`
 * @return {!Array<!Uint8Array>}
 */
proto.buf.validate.BytesRules.prototype.getInList_asU8 = function() {
  return /** @type {!Array<!Uint8Array>} */ (jspb.Message.bytesListAsU8(
      this.getInList()));
};


/**
 * @param {!(Array<!Uint8Array>|Array<string>)} value
 * @return {!proto.buf.validate.BytesRules} returns this
 */
proto.buf.validate.BytesRules.prototype.setInList = function(value) {
  return jspb.Message.setField(this, 8, value || []);
};


/**
 * @param {!(string|Uint8Array)} value
 * @param {number=} opt_index
 * @return {!proto.buf.validate.BytesRules} returns this
 */
proto.buf.validate.BytesRules.prototype.addIn = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 8, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.buf.validate.BytesRules} returns this
 */
proto.buf.validate.BytesRules.prototype.clearInList = function() {
  return this.setInList([]);
};


/**
 * repeated bytes not_in = 9;
 * @return {!(Array<!Uint8Array>|Array<string>)}
 */
proto.buf.validate.BytesRules.prototype.getNotInList = function() {
  return /** @type {!(Array<!Uint8Array>|Array<string>)} */ (jspb.Message.getRepeatedField(this, 9));
};


/**
 * repeated bytes not_in = 9;
 * This is a type-conversion wrapper around `getNotInList()`
 * @return {!Array<string>}
 */
proto.buf.validate.BytesRules.prototype.getNotInList_asB64 = function() {
  return /** @type {!Array<string>} */ (jspb.Message.bytesListAsB64(
      this.getNotInList()));
};


/**
 * repeated bytes not_in = 9;
 * Note that Uint8Array is not supported on all browsers.
 * @see http://caniuse.com/Uint8Array
 * This is a type-conversion wrapper around `getNotInList()`
 * @return {!Array<!Uint8Array>}
 */
proto.buf.validate.BytesRules.prototype.getNotInList_asU8 = function() {
  return /** @type {!Array<!Uint8Array>} */ (jspb.Message.bytesListAsU8(
      this.getNotInList()));
};


/**
 * @param {!(Array<!Uint8Array>|Array<string>)} value
 * @return {!proto.buf.validate.BytesRules} returns this
 */
proto.buf.validate.BytesRules.prototype.setNotInList = function(value) {
  return jspb.Message.setField(this, 9, value || []);
};


/**
 * @param {!(string|Uint8Array)} value
 * @param {number=} opt_index
 * @return {!proto.buf.validate.BytesRules} returns this
 */
proto.buf.validate.BytesRules.prototype.addNotIn = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 9, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.buf.validate.BytesRules} returns this
 */
proto.buf.validate.BytesRules.prototype.clearNotInList = function() {
  return this.setNotInList([]);
};


/**
 * optional bool ip = 10;
 * @return {boolean}
 */
proto.buf.validate.BytesRules.prototype.getIp = function() {
  return /** @type {boolean} */ (jspb.Message.getBooleanFieldWithDefault(this, 10, false));
};


/**
 * @param {boolean} value
 * @return {!proto.buf.validate.BytesRules} returns this
 */
proto.buf.validate.BytesRules.prototype.setIp = function(value) {
  return jspb.Message.setOneofField(this, 10, proto.buf.validate.BytesRules.oneofGroups_[0], value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.BytesRules} returns this
 */
proto.buf.validate.BytesRules.prototype.clearIp = function() {
  return jspb.Message.setOneofField(this, 10, proto.buf.validate.BytesRules.oneofGroups_[0], undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.BytesRules.prototype.hasIp = function() {
  return jspb.Message.getField(this, 10) != null;
};


/**
 * optional bool ipv4 = 11;
 * @return {boolean}
 */
proto.buf.validate.BytesRules.prototype.getIpv4 = function() {
  return /** @type {boolean} */ (jspb.Message.getBooleanFieldWithDefault(this, 11, false));
};


/**
 * @param {boolean} value
 * @return {!proto.buf.validate.BytesRules} returns this
 */
proto.buf.validate.BytesRules.prototype.setIpv4 = function(value) {
  return jspb.Message.setOneofField(this, 11, proto.buf.validate.BytesRules.oneofGroups_[0], value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.BytesRules} returns this
 */
proto.buf.validate.BytesRules.prototype.clearIpv4 = function() {
  return jspb.Message.setOneofField(this, 11, proto.buf.validate.BytesRules.oneofGroups_[0], undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.BytesRules.prototype.hasIpv4 = function() {
  return jspb.Message.getField(this, 11) != null;
};


/**
 * optional bool ipv6 = 12;
 * @return {boolean}
 */
proto.buf.validate.BytesRules.prototype.getIpv6 = function() {
  return /** @type {boolean} */ (jspb.Message.getBooleanFieldWithDefault(this, 12, false));
};


/**
 * @param {boolean} value
 * @return {!proto.buf.validate.BytesRules} returns this
 */
proto.buf.validate.BytesRules.prototype.setIpv6 = function(value) {
  return jspb.Message.setOneofField(this, 12, proto.buf.validate.BytesRules.oneofGroups_[0], value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.BytesRules} returns this
 */
proto.buf.validate.BytesRules.prototype.clearIpv6 = function() {
  return jspb.Message.setOneofField(this, 12, proto.buf.validate.BytesRules.oneofGroups_[0], undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.BytesRules.prototype.hasIpv6 = function() {
  return jspb.Message.getField(this, 12) != null;
};


/**
 * repeated bytes example = 14;
 * @return {!(Array<!Uint8Array>|Array<string>)}
 */
proto.buf.validate.BytesRules.prototype.getExampleList = function() {
  return /** @type {!(Array<!Uint8Array>|Array<string>)} */ (jspb.Message.getRepeatedField(this, 14));
};


/**
 * repeated bytes example = 14;
 * This is a type-conversion wrapper around `getExampleList()`
 * @return {!Array<string>}
 */
proto.buf.validate.BytesRules.prototype.getExampleList_asB64 = function() {
  return /** @type {!Array<string>} */ (jspb.Message.bytesListAsB64(
      this.getExampleList()));
};


/**
 * repeated bytes example = 14;
 * Note that Uint8Array is not supported on all browsers.
 * @see http://caniuse.com/Uint8Array
 * This is a type-conversion wrapper around `getExampleList()`
 * @return {!Array<!Uint8Array>}
 */
proto.buf.validate.BytesRules.prototype.getExampleList_asU8 = function() {
  return /** @type {!Array<!Uint8Array>} */ (jspb.Message.bytesListAsU8(
      this.getExampleList()));
};


/**
 * @param {!(Array<!Uint8Array>|Array<string>)} value
 * @return {!proto.buf.validate.BytesRules} returns this
 */
proto.buf.validate.BytesRules.prototype.setExampleList = function(value) {
  return jspb.Message.setField(this, 14, value || []);
};


/**
 * @param {!(string|Uint8Array)} value
 * @param {number=} opt_index
 * @return {!proto.buf.validate.BytesRules} returns this
 */
proto.buf.validate.BytesRules.prototype.addExample = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 14, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.buf.validate.BytesRules} returns this
 */
proto.buf.validate.BytesRules.prototype.clearExampleList = function() {
  return this.setExampleList([]);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.buf.validate.EnumRules.repeatedFields_ = [3,4,5];



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
proto.buf.validate.EnumRules.prototype.toObject = function(opt_includeInstance) {
  return proto.buf.validate.EnumRules.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.buf.validate.EnumRules} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.buf.validate.EnumRules.toObject = function(includeInstance, msg) {
  var f, obj = {
pb_const: (f = jspb.Message.getField(msg, 1)) == null ? undefined : f,
definedOnly: (f = jspb.Message.getBooleanField(msg, 2)) == null ? undefined : f,
inList: (f = jspb.Message.getRepeatedField(msg, 3)) == null ? undefined : f,
notInList: (f = jspb.Message.getRepeatedField(msg, 4)) == null ? undefined : f,
exampleList: (f = jspb.Message.getRepeatedField(msg, 5)) == null ? undefined : f
  };

  jspb.Message.toObjectExtension(/** @type {!jspb.Message} */ (msg), obj,
      proto.buf.validate.EnumRules.extensions, proto.buf.validate.EnumRules.prototype.getExtension,
      includeInstance);
  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.buf.validate.EnumRules}
 */
proto.buf.validate.EnumRules.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.buf.validate.EnumRules;
  return proto.buf.validate.EnumRules.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.buf.validate.EnumRules} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.buf.validate.EnumRules}
 */
proto.buf.validate.EnumRules.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setConst(value);
      break;
    case 2:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setDefinedOnly(value);
      break;
    case 3:
      var values = /** @type {!Array<number>} */ (reader.isDelimited() ? reader.readPackedInt32() : [reader.readInt32()]);
      for (var i = 0; i < values.length; i++) {
        msg.addIn(values[i]);
      }
      break;
    case 4:
      var values = /** @type {!Array<number>} */ (reader.isDelimited() ? reader.readPackedInt32() : [reader.readInt32()]);
      for (var i = 0; i < values.length; i++) {
        msg.addNotIn(values[i]);
      }
      break;
    case 5:
      var values = /** @type {!Array<number>} */ (reader.isDelimited() ? reader.readPackedInt32() : [reader.readInt32()]);
      for (var i = 0; i < values.length; i++) {
        msg.addExample(values[i]);
      }
      break;
    default:
      jspb.Message.readBinaryExtension(msg, reader,
        proto.buf.validate.EnumRules.extensionsBinary,
        proto.buf.validate.EnumRules.prototype.getExtension,
        proto.buf.validate.EnumRules.prototype.setExtension);
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.buf.validate.EnumRules.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.buf.validate.EnumRules.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.buf.validate.EnumRules} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.buf.validate.EnumRules.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = /** @type {number} */ (jspb.Message.getField(message, 1));
  if (f != null) {
    writer.writeInt32(
      1,
      f
    );
  }
  f = /** @type {boolean} */ (jspb.Message.getField(message, 2));
  if (f != null) {
    writer.writeBool(
      2,
      f
    );
  }
  f = message.getInList();
  if (f.length > 0) {
    writer.writeRepeatedInt32(
      3,
      f
    );
  }
  f = message.getNotInList();
  if (f.length > 0) {
    writer.writeRepeatedInt32(
      4,
      f
    );
  }
  f = message.getExampleList();
  if (f.length > 0) {
    writer.writeRepeatedInt32(
      5,
      f
    );
  }
  jspb.Message.serializeBinaryExtensions(message, writer,
    proto.buf.validate.EnumRules.extensionsBinary, proto.buf.validate.EnumRules.prototype.getExtension);
};


/**
 * optional int32 const = 1;
 * @return {number}
 */
proto.buf.validate.EnumRules.prototype.getConst = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/**
 * @param {number} value
 * @return {!proto.buf.validate.EnumRules} returns this
 */
proto.buf.validate.EnumRules.prototype.setConst = function(value) {
  return jspb.Message.setField(this, 1, value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.EnumRules} returns this
 */
proto.buf.validate.EnumRules.prototype.clearConst = function() {
  return jspb.Message.setField(this, 1, undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.EnumRules.prototype.hasConst = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional bool defined_only = 2;
 * @return {boolean}
 */
proto.buf.validate.EnumRules.prototype.getDefinedOnly = function() {
  return /** @type {boolean} */ (jspb.Message.getBooleanFieldWithDefault(this, 2, false));
};


/**
 * @param {boolean} value
 * @return {!proto.buf.validate.EnumRules} returns this
 */
proto.buf.validate.EnumRules.prototype.setDefinedOnly = function(value) {
  return jspb.Message.setField(this, 2, value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.EnumRules} returns this
 */
proto.buf.validate.EnumRules.prototype.clearDefinedOnly = function() {
  return jspb.Message.setField(this, 2, undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.EnumRules.prototype.hasDefinedOnly = function() {
  return jspb.Message.getField(this, 2) != null;
};


/**
 * repeated int32 in = 3;
 * @return {!Array<number>}
 */
proto.buf.validate.EnumRules.prototype.getInList = function() {
  return /** @type {!Array<number>} */ (jspb.Message.getRepeatedField(this, 3));
};


/**
 * @param {!Array<number>} value
 * @return {!proto.buf.validate.EnumRules} returns this
 */
proto.buf.validate.EnumRules.prototype.setInList = function(value) {
  return jspb.Message.setField(this, 3, value || []);
};


/**
 * @param {number} value
 * @param {number=} opt_index
 * @return {!proto.buf.validate.EnumRules} returns this
 */
proto.buf.validate.EnumRules.prototype.addIn = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 3, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.buf.validate.EnumRules} returns this
 */
proto.buf.validate.EnumRules.prototype.clearInList = function() {
  return this.setInList([]);
};


/**
 * repeated int32 not_in = 4;
 * @return {!Array<number>}
 */
proto.buf.validate.EnumRules.prototype.getNotInList = function() {
  return /** @type {!Array<number>} */ (jspb.Message.getRepeatedField(this, 4));
};


/**
 * @param {!Array<number>} value
 * @return {!proto.buf.validate.EnumRules} returns this
 */
proto.buf.validate.EnumRules.prototype.setNotInList = function(value) {
  return jspb.Message.setField(this, 4, value || []);
};


/**
 * @param {number} value
 * @param {number=} opt_index
 * @return {!proto.buf.validate.EnumRules} returns this
 */
proto.buf.validate.EnumRules.prototype.addNotIn = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 4, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.buf.validate.EnumRules} returns this
 */
proto.buf.validate.EnumRules.prototype.clearNotInList = function() {
  return this.setNotInList([]);
};


/**
 * repeated int32 example = 5;
 * @return {!Array<number>}
 */
proto.buf.validate.EnumRules.prototype.getExampleList = function() {
  return /** @type {!Array<number>} */ (jspb.Message.getRepeatedField(this, 5));
};


/**
 * @param {!Array<number>} value
 * @return {!proto.buf.validate.EnumRules} returns this
 */
proto.buf.validate.EnumRules.prototype.setExampleList = function(value) {
  return jspb.Message.setField(this, 5, value || []);
};


/**
 * @param {number} value
 * @param {number=} opt_index
 * @return {!proto.buf.validate.EnumRules} returns this
 */
proto.buf.validate.EnumRules.prototype.addExample = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 5, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.buf.validate.EnumRules} returns this
 */
proto.buf.validate.EnumRules.prototype.clearExampleList = function() {
  return this.setExampleList([]);
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
proto.buf.validate.RepeatedRules.prototype.toObject = function(opt_includeInstance) {
  return proto.buf.validate.RepeatedRules.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.buf.validate.RepeatedRules} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.buf.validate.RepeatedRules.toObject = function(includeInstance, msg) {
  var f, obj = {
minItems: (f = jspb.Message.getField(msg, 1)) == null ? undefined : f,
maxItems: (f = jspb.Message.getField(msg, 2)) == null ? undefined : f,
unique: (f = jspb.Message.getBooleanField(msg, 3)) == null ? undefined : f,
items: (f = msg.getItems()) && proto.buf.validate.FieldConstraints.toObject(includeInstance, f)
  };

  jspb.Message.toObjectExtension(/** @type {!jspb.Message} */ (msg), obj,
      proto.buf.validate.RepeatedRules.extensions, proto.buf.validate.RepeatedRules.prototype.getExtension,
      includeInstance);
  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.buf.validate.RepeatedRules}
 */
proto.buf.validate.RepeatedRules.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.buf.validate.RepeatedRules;
  return proto.buf.validate.RepeatedRules.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.buf.validate.RepeatedRules} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.buf.validate.RepeatedRules}
 */
proto.buf.validate.RepeatedRules.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readUint64());
      msg.setMinItems(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readUint64());
      msg.setMaxItems(value);
      break;
    case 3:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setUnique(value);
      break;
    case 4:
      var value = new proto.buf.validate.FieldConstraints;
      reader.readMessage(value,proto.buf.validate.FieldConstraints.deserializeBinaryFromReader);
      msg.setItems(value);
      break;
    default:
      jspb.Message.readBinaryExtension(msg, reader,
        proto.buf.validate.RepeatedRules.extensionsBinary,
        proto.buf.validate.RepeatedRules.prototype.getExtension,
        proto.buf.validate.RepeatedRules.prototype.setExtension);
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.buf.validate.RepeatedRules.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.buf.validate.RepeatedRules.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.buf.validate.RepeatedRules} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.buf.validate.RepeatedRules.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = /** @type {number} */ (jspb.Message.getField(message, 1));
  if (f != null) {
    writer.writeUint64(
      1,
      f
    );
  }
  f = /** @type {number} */ (jspb.Message.getField(message, 2));
  if (f != null) {
    writer.writeUint64(
      2,
      f
    );
  }
  f = /** @type {boolean} */ (jspb.Message.getField(message, 3));
  if (f != null) {
    writer.writeBool(
      3,
      f
    );
  }
  f = message.getItems();
  if (f != null) {
    writer.writeMessage(
      4,
      f,
      proto.buf.validate.FieldConstraints.serializeBinaryToWriter
    );
  }
  jspb.Message.serializeBinaryExtensions(message, writer,
    proto.buf.validate.RepeatedRules.extensionsBinary, proto.buf.validate.RepeatedRules.prototype.getExtension);
};


/**
 * optional uint64 min_items = 1;
 * @return {number}
 */
proto.buf.validate.RepeatedRules.prototype.getMinItems = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/**
 * @param {number} value
 * @return {!proto.buf.validate.RepeatedRules} returns this
 */
proto.buf.validate.RepeatedRules.prototype.setMinItems = function(value) {
  return jspb.Message.setField(this, 1, value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.RepeatedRules} returns this
 */
proto.buf.validate.RepeatedRules.prototype.clearMinItems = function() {
  return jspb.Message.setField(this, 1, undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.RepeatedRules.prototype.hasMinItems = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional uint64 max_items = 2;
 * @return {number}
 */
proto.buf.validate.RepeatedRules.prototype.getMaxItems = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/**
 * @param {number} value
 * @return {!proto.buf.validate.RepeatedRules} returns this
 */
proto.buf.validate.RepeatedRules.prototype.setMaxItems = function(value) {
  return jspb.Message.setField(this, 2, value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.RepeatedRules} returns this
 */
proto.buf.validate.RepeatedRules.prototype.clearMaxItems = function() {
  return jspb.Message.setField(this, 2, undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.RepeatedRules.prototype.hasMaxItems = function() {
  return jspb.Message.getField(this, 2) != null;
};


/**
 * optional bool unique = 3;
 * @return {boolean}
 */
proto.buf.validate.RepeatedRules.prototype.getUnique = function() {
  return /** @type {boolean} */ (jspb.Message.getBooleanFieldWithDefault(this, 3, false));
};


/**
 * @param {boolean} value
 * @return {!proto.buf.validate.RepeatedRules} returns this
 */
proto.buf.validate.RepeatedRules.prototype.setUnique = function(value) {
  return jspb.Message.setField(this, 3, value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.RepeatedRules} returns this
 */
proto.buf.validate.RepeatedRules.prototype.clearUnique = function() {
  return jspb.Message.setField(this, 3, undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.RepeatedRules.prototype.hasUnique = function() {
  return jspb.Message.getField(this, 3) != null;
};


/**
 * optional FieldConstraints items = 4;
 * @return {?proto.buf.validate.FieldConstraints}
 */
proto.buf.validate.RepeatedRules.prototype.getItems = function() {
  return /** @type{?proto.buf.validate.FieldConstraints} */ (
    jspb.Message.getWrapperField(this, proto.buf.validate.FieldConstraints, 4));
};


/**
 * @param {?proto.buf.validate.FieldConstraints|undefined} value
 * @return {!proto.buf.validate.RepeatedRules} returns this
*/
proto.buf.validate.RepeatedRules.prototype.setItems = function(value) {
  return jspb.Message.setWrapperField(this, 4, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.buf.validate.RepeatedRules} returns this
 */
proto.buf.validate.RepeatedRules.prototype.clearItems = function() {
  return this.setItems(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.RepeatedRules.prototype.hasItems = function() {
  return jspb.Message.getField(this, 4) != null;
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
proto.buf.validate.MapRules.prototype.toObject = function(opt_includeInstance) {
  return proto.buf.validate.MapRules.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.buf.validate.MapRules} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.buf.validate.MapRules.toObject = function(includeInstance, msg) {
  var f, obj = {
minPairs: (f = jspb.Message.getField(msg, 1)) == null ? undefined : f,
maxPairs: (f = jspb.Message.getField(msg, 2)) == null ? undefined : f,
keys: (f = msg.getKeys()) && proto.buf.validate.FieldConstraints.toObject(includeInstance, f),
values: (f = msg.getValues()) && proto.buf.validate.FieldConstraints.toObject(includeInstance, f)
  };

  jspb.Message.toObjectExtension(/** @type {!jspb.Message} */ (msg), obj,
      proto.buf.validate.MapRules.extensions, proto.buf.validate.MapRules.prototype.getExtension,
      includeInstance);
  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.buf.validate.MapRules}
 */
proto.buf.validate.MapRules.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.buf.validate.MapRules;
  return proto.buf.validate.MapRules.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.buf.validate.MapRules} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.buf.validate.MapRules}
 */
proto.buf.validate.MapRules.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readUint64());
      msg.setMinPairs(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readUint64());
      msg.setMaxPairs(value);
      break;
    case 4:
      var value = new proto.buf.validate.FieldConstraints;
      reader.readMessage(value,proto.buf.validate.FieldConstraints.deserializeBinaryFromReader);
      msg.setKeys(value);
      break;
    case 5:
      var value = new proto.buf.validate.FieldConstraints;
      reader.readMessage(value,proto.buf.validate.FieldConstraints.deserializeBinaryFromReader);
      msg.setValues(value);
      break;
    default:
      jspb.Message.readBinaryExtension(msg, reader,
        proto.buf.validate.MapRules.extensionsBinary,
        proto.buf.validate.MapRules.prototype.getExtension,
        proto.buf.validate.MapRules.prototype.setExtension);
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.buf.validate.MapRules.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.buf.validate.MapRules.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.buf.validate.MapRules} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.buf.validate.MapRules.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = /** @type {number} */ (jspb.Message.getField(message, 1));
  if (f != null) {
    writer.writeUint64(
      1,
      f
    );
  }
  f = /** @type {number} */ (jspb.Message.getField(message, 2));
  if (f != null) {
    writer.writeUint64(
      2,
      f
    );
  }
  f = message.getKeys();
  if (f != null) {
    writer.writeMessage(
      4,
      f,
      proto.buf.validate.FieldConstraints.serializeBinaryToWriter
    );
  }
  f = message.getValues();
  if (f != null) {
    writer.writeMessage(
      5,
      f,
      proto.buf.validate.FieldConstraints.serializeBinaryToWriter
    );
  }
  jspb.Message.serializeBinaryExtensions(message, writer,
    proto.buf.validate.MapRules.extensionsBinary, proto.buf.validate.MapRules.prototype.getExtension);
};


/**
 * optional uint64 min_pairs = 1;
 * @return {number}
 */
proto.buf.validate.MapRules.prototype.getMinPairs = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/**
 * @param {number} value
 * @return {!proto.buf.validate.MapRules} returns this
 */
proto.buf.validate.MapRules.prototype.setMinPairs = function(value) {
  return jspb.Message.setField(this, 1, value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.MapRules} returns this
 */
proto.buf.validate.MapRules.prototype.clearMinPairs = function() {
  return jspb.Message.setField(this, 1, undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.MapRules.prototype.hasMinPairs = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional uint64 max_pairs = 2;
 * @return {number}
 */
proto.buf.validate.MapRules.prototype.getMaxPairs = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/**
 * @param {number} value
 * @return {!proto.buf.validate.MapRules} returns this
 */
proto.buf.validate.MapRules.prototype.setMaxPairs = function(value) {
  return jspb.Message.setField(this, 2, value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.MapRules} returns this
 */
proto.buf.validate.MapRules.prototype.clearMaxPairs = function() {
  return jspb.Message.setField(this, 2, undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.MapRules.prototype.hasMaxPairs = function() {
  return jspb.Message.getField(this, 2) != null;
};


/**
 * optional FieldConstraints keys = 4;
 * @return {?proto.buf.validate.FieldConstraints}
 */
proto.buf.validate.MapRules.prototype.getKeys = function() {
  return /** @type{?proto.buf.validate.FieldConstraints} */ (
    jspb.Message.getWrapperField(this, proto.buf.validate.FieldConstraints, 4));
};


/**
 * @param {?proto.buf.validate.FieldConstraints|undefined} value
 * @return {!proto.buf.validate.MapRules} returns this
*/
proto.buf.validate.MapRules.prototype.setKeys = function(value) {
  return jspb.Message.setWrapperField(this, 4, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.buf.validate.MapRules} returns this
 */
proto.buf.validate.MapRules.prototype.clearKeys = function() {
  return this.setKeys(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.MapRules.prototype.hasKeys = function() {
  return jspb.Message.getField(this, 4) != null;
};


/**
 * optional FieldConstraints values = 5;
 * @return {?proto.buf.validate.FieldConstraints}
 */
proto.buf.validate.MapRules.prototype.getValues = function() {
  return /** @type{?proto.buf.validate.FieldConstraints} */ (
    jspb.Message.getWrapperField(this, proto.buf.validate.FieldConstraints, 5));
};


/**
 * @param {?proto.buf.validate.FieldConstraints|undefined} value
 * @return {!proto.buf.validate.MapRules} returns this
*/
proto.buf.validate.MapRules.prototype.setValues = function(value) {
  return jspb.Message.setWrapperField(this, 5, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.buf.validate.MapRules} returns this
 */
proto.buf.validate.MapRules.prototype.clearValues = function() {
  return this.setValues(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.MapRules.prototype.hasValues = function() {
  return jspb.Message.getField(this, 5) != null;
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.buf.validate.AnyRules.repeatedFields_ = [2,3];



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
proto.buf.validate.AnyRules.prototype.toObject = function(opt_includeInstance) {
  return proto.buf.validate.AnyRules.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.buf.validate.AnyRules} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.buf.validate.AnyRules.toObject = function(includeInstance, msg) {
  var f, obj = {
inList: (f = jspb.Message.getRepeatedField(msg, 2)) == null ? undefined : f,
notInList: (f = jspb.Message.getRepeatedField(msg, 3)) == null ? undefined : f
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
 * @return {!proto.buf.validate.AnyRules}
 */
proto.buf.validate.AnyRules.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.buf.validate.AnyRules;
  return proto.buf.validate.AnyRules.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.buf.validate.AnyRules} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.buf.validate.AnyRules}
 */
proto.buf.validate.AnyRules.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.addIn(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.addNotIn(value);
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
proto.buf.validate.AnyRules.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.buf.validate.AnyRules.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.buf.validate.AnyRules} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.buf.validate.AnyRules.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getInList();
  if (f.length > 0) {
    writer.writeRepeatedString(
      2,
      f
    );
  }
  f = message.getNotInList();
  if (f.length > 0) {
    writer.writeRepeatedString(
      3,
      f
    );
  }
};


/**
 * repeated string in = 2;
 * @return {!Array<string>}
 */
proto.buf.validate.AnyRules.prototype.getInList = function() {
  return /** @type {!Array<string>} */ (jspb.Message.getRepeatedField(this, 2));
};


/**
 * @param {!Array<string>} value
 * @return {!proto.buf.validate.AnyRules} returns this
 */
proto.buf.validate.AnyRules.prototype.setInList = function(value) {
  return jspb.Message.setField(this, 2, value || []);
};


/**
 * @param {string} value
 * @param {number=} opt_index
 * @return {!proto.buf.validate.AnyRules} returns this
 */
proto.buf.validate.AnyRules.prototype.addIn = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 2, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.buf.validate.AnyRules} returns this
 */
proto.buf.validate.AnyRules.prototype.clearInList = function() {
  return this.setInList([]);
};


/**
 * repeated string not_in = 3;
 * @return {!Array<string>}
 */
proto.buf.validate.AnyRules.prototype.getNotInList = function() {
  return /** @type {!Array<string>} */ (jspb.Message.getRepeatedField(this, 3));
};


/**
 * @param {!Array<string>} value
 * @return {!proto.buf.validate.AnyRules} returns this
 */
proto.buf.validate.AnyRules.prototype.setNotInList = function(value) {
  return jspb.Message.setField(this, 3, value || []);
};


/**
 * @param {string} value
 * @param {number=} opt_index
 * @return {!proto.buf.validate.AnyRules} returns this
 */
proto.buf.validate.AnyRules.prototype.addNotIn = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 3, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.buf.validate.AnyRules} returns this
 */
proto.buf.validate.AnyRules.prototype.clearNotInList = function() {
  return this.setNotInList([]);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.buf.validate.DurationRules.repeatedFields_ = [7,8,9];

/**
 * Oneof group definitions for this message. Each group defines the field
 * numbers belonging to that group. When of these fields' value is set, all
 * other fields in the group are cleared. During deserialization, if multiple
 * fields are encountered for a group, only the last value seen will be kept.
 * @private {!Array<!Array<number>>}
 * @const
 */
proto.buf.validate.DurationRules.oneofGroups_ = [[3,4],[5,6]];

/**
 * @enum {number}
 */
proto.buf.validate.DurationRules.LessThanCase = {
  LESS_THAN_NOT_SET: 0,
  LT: 3,
  LTE: 4
};

/**
 * @return {proto.buf.validate.DurationRules.LessThanCase}
 */
proto.buf.validate.DurationRules.prototype.getLessThanCase = function() {
  return /** @type {proto.buf.validate.DurationRules.LessThanCase} */(jspb.Message.computeOneofCase(this, proto.buf.validate.DurationRules.oneofGroups_[0]));
};

/**
 * @enum {number}
 */
proto.buf.validate.DurationRules.GreaterThanCase = {
  GREATER_THAN_NOT_SET: 0,
  GT: 5,
  GTE: 6
};

/**
 * @return {proto.buf.validate.DurationRules.GreaterThanCase}
 */
proto.buf.validate.DurationRules.prototype.getGreaterThanCase = function() {
  return /** @type {proto.buf.validate.DurationRules.GreaterThanCase} */(jspb.Message.computeOneofCase(this, proto.buf.validate.DurationRules.oneofGroups_[1]));
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
proto.buf.validate.DurationRules.prototype.toObject = function(opt_includeInstance) {
  return proto.buf.validate.DurationRules.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.buf.validate.DurationRules} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.buf.validate.DurationRules.toObject = function(includeInstance, msg) {
  var f, obj = {
pb_const: (f = msg.getConst()) && google_protobuf_duration_pb.Duration.toObject(includeInstance, f),
lt: (f = msg.getLt()) && google_protobuf_duration_pb.Duration.toObject(includeInstance, f),
lte: (f = msg.getLte()) && google_protobuf_duration_pb.Duration.toObject(includeInstance, f),
gt: (f = msg.getGt()) && google_protobuf_duration_pb.Duration.toObject(includeInstance, f),
gte: (f = msg.getGte()) && google_protobuf_duration_pb.Duration.toObject(includeInstance, f),
inList: jspb.Message.toObjectList(msg.getInList(),
    google_protobuf_duration_pb.Duration.toObject, includeInstance),
notInList: jspb.Message.toObjectList(msg.getNotInList(),
    google_protobuf_duration_pb.Duration.toObject, includeInstance),
exampleList: jspb.Message.toObjectList(msg.getExampleList(),
    google_protobuf_duration_pb.Duration.toObject, includeInstance)
  };

  jspb.Message.toObjectExtension(/** @type {!jspb.Message} */ (msg), obj,
      proto.buf.validate.DurationRules.extensions, proto.buf.validate.DurationRules.prototype.getExtension,
      includeInstance);
  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.buf.validate.DurationRules}
 */
proto.buf.validate.DurationRules.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.buf.validate.DurationRules;
  return proto.buf.validate.DurationRules.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.buf.validate.DurationRules} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.buf.validate.DurationRules}
 */
proto.buf.validate.DurationRules.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 2:
      var value = new google_protobuf_duration_pb.Duration;
      reader.readMessage(value,google_protobuf_duration_pb.Duration.deserializeBinaryFromReader);
      msg.setConst(value);
      break;
    case 3:
      var value = new google_protobuf_duration_pb.Duration;
      reader.readMessage(value,google_protobuf_duration_pb.Duration.deserializeBinaryFromReader);
      msg.setLt(value);
      break;
    case 4:
      var value = new google_protobuf_duration_pb.Duration;
      reader.readMessage(value,google_protobuf_duration_pb.Duration.deserializeBinaryFromReader);
      msg.setLte(value);
      break;
    case 5:
      var value = new google_protobuf_duration_pb.Duration;
      reader.readMessage(value,google_protobuf_duration_pb.Duration.deserializeBinaryFromReader);
      msg.setGt(value);
      break;
    case 6:
      var value = new google_protobuf_duration_pb.Duration;
      reader.readMessage(value,google_protobuf_duration_pb.Duration.deserializeBinaryFromReader);
      msg.setGte(value);
      break;
    case 7:
      var value = new google_protobuf_duration_pb.Duration;
      reader.readMessage(value,google_protobuf_duration_pb.Duration.deserializeBinaryFromReader);
      msg.addIn(value);
      break;
    case 8:
      var value = new google_protobuf_duration_pb.Duration;
      reader.readMessage(value,google_protobuf_duration_pb.Duration.deserializeBinaryFromReader);
      msg.addNotIn(value);
      break;
    case 9:
      var value = new google_protobuf_duration_pb.Duration;
      reader.readMessage(value,google_protobuf_duration_pb.Duration.deserializeBinaryFromReader);
      msg.addExample(value);
      break;
    default:
      jspb.Message.readBinaryExtension(msg, reader,
        proto.buf.validate.DurationRules.extensionsBinary,
        proto.buf.validate.DurationRules.prototype.getExtension,
        proto.buf.validate.DurationRules.prototype.setExtension);
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.buf.validate.DurationRules.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.buf.validate.DurationRules.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.buf.validate.DurationRules} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.buf.validate.DurationRules.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getConst();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      google_protobuf_duration_pb.Duration.serializeBinaryToWriter
    );
  }
  f = message.getLt();
  if (f != null) {
    writer.writeMessage(
      3,
      f,
      google_protobuf_duration_pb.Duration.serializeBinaryToWriter
    );
  }
  f = message.getLte();
  if (f != null) {
    writer.writeMessage(
      4,
      f,
      google_protobuf_duration_pb.Duration.serializeBinaryToWriter
    );
  }
  f = message.getGt();
  if (f != null) {
    writer.writeMessage(
      5,
      f,
      google_protobuf_duration_pb.Duration.serializeBinaryToWriter
    );
  }
  f = message.getGte();
  if (f != null) {
    writer.writeMessage(
      6,
      f,
      google_protobuf_duration_pb.Duration.serializeBinaryToWriter
    );
  }
  f = message.getInList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      7,
      f,
      google_protobuf_duration_pb.Duration.serializeBinaryToWriter
    );
  }
  f = message.getNotInList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      8,
      f,
      google_protobuf_duration_pb.Duration.serializeBinaryToWriter
    );
  }
  f = message.getExampleList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      9,
      f,
      google_protobuf_duration_pb.Duration.serializeBinaryToWriter
    );
  }
  jspb.Message.serializeBinaryExtensions(message, writer,
    proto.buf.validate.DurationRules.extensionsBinary, proto.buf.validate.DurationRules.prototype.getExtension);
};


/**
 * optional google.protobuf.Duration const = 2;
 * @return {?proto.google.protobuf.Duration}
 */
proto.buf.validate.DurationRules.prototype.getConst = function() {
  return /** @type{?proto.google.protobuf.Duration} */ (
    jspb.Message.getWrapperField(this, google_protobuf_duration_pb.Duration, 2));
};


/**
 * @param {?proto.google.protobuf.Duration|undefined} value
 * @return {!proto.buf.validate.DurationRules} returns this
*/
proto.buf.validate.DurationRules.prototype.setConst = function(value) {
  return jspb.Message.setWrapperField(this, 2, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.buf.validate.DurationRules} returns this
 */
proto.buf.validate.DurationRules.prototype.clearConst = function() {
  return this.setConst(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.DurationRules.prototype.hasConst = function() {
  return jspb.Message.getField(this, 2) != null;
};


/**
 * optional google.protobuf.Duration lt = 3;
 * @return {?proto.google.protobuf.Duration}
 */
proto.buf.validate.DurationRules.prototype.getLt = function() {
  return /** @type{?proto.google.protobuf.Duration} */ (
    jspb.Message.getWrapperField(this, google_protobuf_duration_pb.Duration, 3));
};


/**
 * @param {?proto.google.protobuf.Duration|undefined} value
 * @return {!proto.buf.validate.DurationRules} returns this
*/
proto.buf.validate.DurationRules.prototype.setLt = function(value) {
  return jspb.Message.setOneofWrapperField(this, 3, proto.buf.validate.DurationRules.oneofGroups_[0], value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.buf.validate.DurationRules} returns this
 */
proto.buf.validate.DurationRules.prototype.clearLt = function() {
  return this.setLt(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.DurationRules.prototype.hasLt = function() {
  return jspb.Message.getField(this, 3) != null;
};


/**
 * optional google.protobuf.Duration lte = 4;
 * @return {?proto.google.protobuf.Duration}
 */
proto.buf.validate.DurationRules.prototype.getLte = function() {
  return /** @type{?proto.google.protobuf.Duration} */ (
    jspb.Message.getWrapperField(this, google_protobuf_duration_pb.Duration, 4));
};


/**
 * @param {?proto.google.protobuf.Duration|undefined} value
 * @return {!proto.buf.validate.DurationRules} returns this
*/
proto.buf.validate.DurationRules.prototype.setLte = function(value) {
  return jspb.Message.setOneofWrapperField(this, 4, proto.buf.validate.DurationRules.oneofGroups_[0], value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.buf.validate.DurationRules} returns this
 */
proto.buf.validate.DurationRules.prototype.clearLte = function() {
  return this.setLte(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.DurationRules.prototype.hasLte = function() {
  return jspb.Message.getField(this, 4) != null;
};


/**
 * optional google.protobuf.Duration gt = 5;
 * @return {?proto.google.protobuf.Duration}
 */
proto.buf.validate.DurationRules.prototype.getGt = function() {
  return /** @type{?proto.google.protobuf.Duration} */ (
    jspb.Message.getWrapperField(this, google_protobuf_duration_pb.Duration, 5));
};


/**
 * @param {?proto.google.protobuf.Duration|undefined} value
 * @return {!proto.buf.validate.DurationRules} returns this
*/
proto.buf.validate.DurationRules.prototype.setGt = function(value) {
  return jspb.Message.setOneofWrapperField(this, 5, proto.buf.validate.DurationRules.oneofGroups_[1], value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.buf.validate.DurationRules} returns this
 */
proto.buf.validate.DurationRules.prototype.clearGt = function() {
  return this.setGt(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.DurationRules.prototype.hasGt = function() {
  return jspb.Message.getField(this, 5) != null;
};


/**
 * optional google.protobuf.Duration gte = 6;
 * @return {?proto.google.protobuf.Duration}
 */
proto.buf.validate.DurationRules.prototype.getGte = function() {
  return /** @type{?proto.google.protobuf.Duration} */ (
    jspb.Message.getWrapperField(this, google_protobuf_duration_pb.Duration, 6));
};


/**
 * @param {?proto.google.protobuf.Duration|undefined} value
 * @return {!proto.buf.validate.DurationRules} returns this
*/
proto.buf.validate.DurationRules.prototype.setGte = function(value) {
  return jspb.Message.setOneofWrapperField(this, 6, proto.buf.validate.DurationRules.oneofGroups_[1], value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.buf.validate.DurationRules} returns this
 */
proto.buf.validate.DurationRules.prototype.clearGte = function() {
  return this.setGte(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.DurationRules.prototype.hasGte = function() {
  return jspb.Message.getField(this, 6) != null;
};


/**
 * repeated google.protobuf.Duration in = 7;
 * @return {!Array<!proto.google.protobuf.Duration>}
 */
proto.buf.validate.DurationRules.prototype.getInList = function() {
  return /** @type{!Array<!proto.google.protobuf.Duration>} */ (
    jspb.Message.getRepeatedWrapperField(this, google_protobuf_duration_pb.Duration, 7));
};


/**
 * @param {!Array<!proto.google.protobuf.Duration>} value
 * @return {!proto.buf.validate.DurationRules} returns this
*/
proto.buf.validate.DurationRules.prototype.setInList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 7, value);
};


/**
 * @param {!proto.google.protobuf.Duration=} opt_value
 * @param {number=} opt_index
 * @return {!proto.google.protobuf.Duration}
 */
proto.buf.validate.DurationRules.prototype.addIn = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 7, opt_value, proto.google.protobuf.Duration, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.buf.validate.DurationRules} returns this
 */
proto.buf.validate.DurationRules.prototype.clearInList = function() {
  return this.setInList([]);
};


/**
 * repeated google.protobuf.Duration not_in = 8;
 * @return {!Array<!proto.google.protobuf.Duration>}
 */
proto.buf.validate.DurationRules.prototype.getNotInList = function() {
  return /** @type{!Array<!proto.google.protobuf.Duration>} */ (
    jspb.Message.getRepeatedWrapperField(this, google_protobuf_duration_pb.Duration, 8));
};


/**
 * @param {!Array<!proto.google.protobuf.Duration>} value
 * @return {!proto.buf.validate.DurationRules} returns this
*/
proto.buf.validate.DurationRules.prototype.setNotInList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 8, value);
};


/**
 * @param {!proto.google.protobuf.Duration=} opt_value
 * @param {number=} opt_index
 * @return {!proto.google.protobuf.Duration}
 */
proto.buf.validate.DurationRules.prototype.addNotIn = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 8, opt_value, proto.google.protobuf.Duration, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.buf.validate.DurationRules} returns this
 */
proto.buf.validate.DurationRules.prototype.clearNotInList = function() {
  return this.setNotInList([]);
};


/**
 * repeated google.protobuf.Duration example = 9;
 * @return {!Array<!proto.google.protobuf.Duration>}
 */
proto.buf.validate.DurationRules.prototype.getExampleList = function() {
  return /** @type{!Array<!proto.google.protobuf.Duration>} */ (
    jspb.Message.getRepeatedWrapperField(this, google_protobuf_duration_pb.Duration, 9));
};


/**
 * @param {!Array<!proto.google.protobuf.Duration>} value
 * @return {!proto.buf.validate.DurationRules} returns this
*/
proto.buf.validate.DurationRules.prototype.setExampleList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 9, value);
};


/**
 * @param {!proto.google.protobuf.Duration=} opt_value
 * @param {number=} opt_index
 * @return {!proto.google.protobuf.Duration}
 */
proto.buf.validate.DurationRules.prototype.addExample = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 9, opt_value, proto.google.protobuf.Duration, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.buf.validate.DurationRules} returns this
 */
proto.buf.validate.DurationRules.prototype.clearExampleList = function() {
  return this.setExampleList([]);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.buf.validate.TimestampRules.repeatedFields_ = [10];

/**
 * Oneof group definitions for this message. Each group defines the field
 * numbers belonging to that group. When of these fields' value is set, all
 * other fields in the group are cleared. During deserialization, if multiple
 * fields are encountered for a group, only the last value seen will be kept.
 * @private {!Array<!Array<number>>}
 * @const
 */
proto.buf.validate.TimestampRules.oneofGroups_ = [[3,4,7],[5,6,8]];

/**
 * @enum {number}
 */
proto.buf.validate.TimestampRules.LessThanCase = {
  LESS_THAN_NOT_SET: 0,
  LT: 3,
  LTE: 4,
  LT_NOW: 7
};

/**
 * @return {proto.buf.validate.TimestampRules.LessThanCase}
 */
proto.buf.validate.TimestampRules.prototype.getLessThanCase = function() {
  return /** @type {proto.buf.validate.TimestampRules.LessThanCase} */(jspb.Message.computeOneofCase(this, proto.buf.validate.TimestampRules.oneofGroups_[0]));
};

/**
 * @enum {number}
 */
proto.buf.validate.TimestampRules.GreaterThanCase = {
  GREATER_THAN_NOT_SET: 0,
  GT: 5,
  GTE: 6,
  GT_NOW: 8
};

/**
 * @return {proto.buf.validate.TimestampRules.GreaterThanCase}
 */
proto.buf.validate.TimestampRules.prototype.getGreaterThanCase = function() {
  return /** @type {proto.buf.validate.TimestampRules.GreaterThanCase} */(jspb.Message.computeOneofCase(this, proto.buf.validate.TimestampRules.oneofGroups_[1]));
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
proto.buf.validate.TimestampRules.prototype.toObject = function(opt_includeInstance) {
  return proto.buf.validate.TimestampRules.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.buf.validate.TimestampRules} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.buf.validate.TimestampRules.toObject = function(includeInstance, msg) {
  var f, obj = {
pb_const: (f = msg.getConst()) && google_protobuf_timestamp_pb.Timestamp.toObject(includeInstance, f),
lt: (f = msg.getLt()) && google_protobuf_timestamp_pb.Timestamp.toObject(includeInstance, f),
lte: (f = msg.getLte()) && google_protobuf_timestamp_pb.Timestamp.toObject(includeInstance, f),
ltNow: (f = jspb.Message.getBooleanField(msg, 7)) == null ? undefined : f,
gt: (f = msg.getGt()) && google_protobuf_timestamp_pb.Timestamp.toObject(includeInstance, f),
gte: (f = msg.getGte()) && google_protobuf_timestamp_pb.Timestamp.toObject(includeInstance, f),
gtNow: (f = jspb.Message.getBooleanField(msg, 8)) == null ? undefined : f,
within: (f = msg.getWithin()) && google_protobuf_duration_pb.Duration.toObject(includeInstance, f),
exampleList: jspb.Message.toObjectList(msg.getExampleList(),
    google_protobuf_timestamp_pb.Timestamp.toObject, includeInstance)
  };

  jspb.Message.toObjectExtension(/** @type {!jspb.Message} */ (msg), obj,
      proto.buf.validate.TimestampRules.extensions, proto.buf.validate.TimestampRules.prototype.getExtension,
      includeInstance);
  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.buf.validate.TimestampRules}
 */
proto.buf.validate.TimestampRules.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.buf.validate.TimestampRules;
  return proto.buf.validate.TimestampRules.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.buf.validate.TimestampRules} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.buf.validate.TimestampRules}
 */
proto.buf.validate.TimestampRules.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 2:
      var value = new google_protobuf_timestamp_pb.Timestamp;
      reader.readMessage(value,google_protobuf_timestamp_pb.Timestamp.deserializeBinaryFromReader);
      msg.setConst(value);
      break;
    case 3:
      var value = new google_protobuf_timestamp_pb.Timestamp;
      reader.readMessage(value,google_protobuf_timestamp_pb.Timestamp.deserializeBinaryFromReader);
      msg.setLt(value);
      break;
    case 4:
      var value = new google_protobuf_timestamp_pb.Timestamp;
      reader.readMessage(value,google_protobuf_timestamp_pb.Timestamp.deserializeBinaryFromReader);
      msg.setLte(value);
      break;
    case 7:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setLtNow(value);
      break;
    case 5:
      var value = new google_protobuf_timestamp_pb.Timestamp;
      reader.readMessage(value,google_protobuf_timestamp_pb.Timestamp.deserializeBinaryFromReader);
      msg.setGt(value);
      break;
    case 6:
      var value = new google_protobuf_timestamp_pb.Timestamp;
      reader.readMessage(value,google_protobuf_timestamp_pb.Timestamp.deserializeBinaryFromReader);
      msg.setGte(value);
      break;
    case 8:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setGtNow(value);
      break;
    case 9:
      var value = new google_protobuf_duration_pb.Duration;
      reader.readMessage(value,google_protobuf_duration_pb.Duration.deserializeBinaryFromReader);
      msg.setWithin(value);
      break;
    case 10:
      var value = new google_protobuf_timestamp_pb.Timestamp;
      reader.readMessage(value,google_protobuf_timestamp_pb.Timestamp.deserializeBinaryFromReader);
      msg.addExample(value);
      break;
    default:
      jspb.Message.readBinaryExtension(msg, reader,
        proto.buf.validate.TimestampRules.extensionsBinary,
        proto.buf.validate.TimestampRules.prototype.getExtension,
        proto.buf.validate.TimestampRules.prototype.setExtension);
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.buf.validate.TimestampRules.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.buf.validate.TimestampRules.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.buf.validate.TimestampRules} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.buf.validate.TimestampRules.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getConst();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      google_protobuf_timestamp_pb.Timestamp.serializeBinaryToWriter
    );
  }
  f = message.getLt();
  if (f != null) {
    writer.writeMessage(
      3,
      f,
      google_protobuf_timestamp_pb.Timestamp.serializeBinaryToWriter
    );
  }
  f = message.getLte();
  if (f != null) {
    writer.writeMessage(
      4,
      f,
      google_protobuf_timestamp_pb.Timestamp.serializeBinaryToWriter
    );
  }
  f = /** @type {boolean} */ (jspb.Message.getField(message, 7));
  if (f != null) {
    writer.writeBool(
      7,
      f
    );
  }
  f = message.getGt();
  if (f != null) {
    writer.writeMessage(
      5,
      f,
      google_protobuf_timestamp_pb.Timestamp.serializeBinaryToWriter
    );
  }
  f = message.getGte();
  if (f != null) {
    writer.writeMessage(
      6,
      f,
      google_protobuf_timestamp_pb.Timestamp.serializeBinaryToWriter
    );
  }
  f = /** @type {boolean} */ (jspb.Message.getField(message, 8));
  if (f != null) {
    writer.writeBool(
      8,
      f
    );
  }
  f = message.getWithin();
  if (f != null) {
    writer.writeMessage(
      9,
      f,
      google_protobuf_duration_pb.Duration.serializeBinaryToWriter
    );
  }
  f = message.getExampleList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      10,
      f,
      google_protobuf_timestamp_pb.Timestamp.serializeBinaryToWriter
    );
  }
  jspb.Message.serializeBinaryExtensions(message, writer,
    proto.buf.validate.TimestampRules.extensionsBinary, proto.buf.validate.TimestampRules.prototype.getExtension);
};


/**
 * optional google.protobuf.Timestamp const = 2;
 * @return {?proto.google.protobuf.Timestamp}
 */
proto.buf.validate.TimestampRules.prototype.getConst = function() {
  return /** @type{?proto.google.protobuf.Timestamp} */ (
    jspb.Message.getWrapperField(this, google_protobuf_timestamp_pb.Timestamp, 2));
};


/**
 * @param {?proto.google.protobuf.Timestamp|undefined} value
 * @return {!proto.buf.validate.TimestampRules} returns this
*/
proto.buf.validate.TimestampRules.prototype.setConst = function(value) {
  return jspb.Message.setWrapperField(this, 2, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.buf.validate.TimestampRules} returns this
 */
proto.buf.validate.TimestampRules.prototype.clearConst = function() {
  return this.setConst(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.TimestampRules.prototype.hasConst = function() {
  return jspb.Message.getField(this, 2) != null;
};


/**
 * optional google.protobuf.Timestamp lt = 3;
 * @return {?proto.google.protobuf.Timestamp}
 */
proto.buf.validate.TimestampRules.prototype.getLt = function() {
  return /** @type{?proto.google.protobuf.Timestamp} */ (
    jspb.Message.getWrapperField(this, google_protobuf_timestamp_pb.Timestamp, 3));
};


/**
 * @param {?proto.google.protobuf.Timestamp|undefined} value
 * @return {!proto.buf.validate.TimestampRules} returns this
*/
proto.buf.validate.TimestampRules.prototype.setLt = function(value) {
  return jspb.Message.setOneofWrapperField(this, 3, proto.buf.validate.TimestampRules.oneofGroups_[0], value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.buf.validate.TimestampRules} returns this
 */
proto.buf.validate.TimestampRules.prototype.clearLt = function() {
  return this.setLt(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.TimestampRules.prototype.hasLt = function() {
  return jspb.Message.getField(this, 3) != null;
};


/**
 * optional google.protobuf.Timestamp lte = 4;
 * @return {?proto.google.protobuf.Timestamp}
 */
proto.buf.validate.TimestampRules.prototype.getLte = function() {
  return /** @type{?proto.google.protobuf.Timestamp} */ (
    jspb.Message.getWrapperField(this, google_protobuf_timestamp_pb.Timestamp, 4));
};


/**
 * @param {?proto.google.protobuf.Timestamp|undefined} value
 * @return {!proto.buf.validate.TimestampRules} returns this
*/
proto.buf.validate.TimestampRules.prototype.setLte = function(value) {
  return jspb.Message.setOneofWrapperField(this, 4, proto.buf.validate.TimestampRules.oneofGroups_[0], value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.buf.validate.TimestampRules} returns this
 */
proto.buf.validate.TimestampRules.prototype.clearLte = function() {
  return this.setLte(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.TimestampRules.prototype.hasLte = function() {
  return jspb.Message.getField(this, 4) != null;
};


/**
 * optional bool lt_now = 7;
 * @return {boolean}
 */
proto.buf.validate.TimestampRules.prototype.getLtNow = function() {
  return /** @type {boolean} */ (jspb.Message.getBooleanFieldWithDefault(this, 7, false));
};


/**
 * @param {boolean} value
 * @return {!proto.buf.validate.TimestampRules} returns this
 */
proto.buf.validate.TimestampRules.prototype.setLtNow = function(value) {
  return jspb.Message.setOneofField(this, 7, proto.buf.validate.TimestampRules.oneofGroups_[0], value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.TimestampRules} returns this
 */
proto.buf.validate.TimestampRules.prototype.clearLtNow = function() {
  return jspb.Message.setOneofField(this, 7, proto.buf.validate.TimestampRules.oneofGroups_[0], undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.TimestampRules.prototype.hasLtNow = function() {
  return jspb.Message.getField(this, 7) != null;
};


/**
 * optional google.protobuf.Timestamp gt = 5;
 * @return {?proto.google.protobuf.Timestamp}
 */
proto.buf.validate.TimestampRules.prototype.getGt = function() {
  return /** @type{?proto.google.protobuf.Timestamp} */ (
    jspb.Message.getWrapperField(this, google_protobuf_timestamp_pb.Timestamp, 5));
};


/**
 * @param {?proto.google.protobuf.Timestamp|undefined} value
 * @return {!proto.buf.validate.TimestampRules} returns this
*/
proto.buf.validate.TimestampRules.prototype.setGt = function(value) {
  return jspb.Message.setOneofWrapperField(this, 5, proto.buf.validate.TimestampRules.oneofGroups_[1], value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.buf.validate.TimestampRules} returns this
 */
proto.buf.validate.TimestampRules.prototype.clearGt = function() {
  return this.setGt(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.TimestampRules.prototype.hasGt = function() {
  return jspb.Message.getField(this, 5) != null;
};


/**
 * optional google.protobuf.Timestamp gte = 6;
 * @return {?proto.google.protobuf.Timestamp}
 */
proto.buf.validate.TimestampRules.prototype.getGte = function() {
  return /** @type{?proto.google.protobuf.Timestamp} */ (
    jspb.Message.getWrapperField(this, google_protobuf_timestamp_pb.Timestamp, 6));
};


/**
 * @param {?proto.google.protobuf.Timestamp|undefined} value
 * @return {!proto.buf.validate.TimestampRules} returns this
*/
proto.buf.validate.TimestampRules.prototype.setGte = function(value) {
  return jspb.Message.setOneofWrapperField(this, 6, proto.buf.validate.TimestampRules.oneofGroups_[1], value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.buf.validate.TimestampRules} returns this
 */
proto.buf.validate.TimestampRules.prototype.clearGte = function() {
  return this.setGte(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.TimestampRules.prototype.hasGte = function() {
  return jspb.Message.getField(this, 6) != null;
};


/**
 * optional bool gt_now = 8;
 * @return {boolean}
 */
proto.buf.validate.TimestampRules.prototype.getGtNow = function() {
  return /** @type {boolean} */ (jspb.Message.getBooleanFieldWithDefault(this, 8, false));
};


/**
 * @param {boolean} value
 * @return {!proto.buf.validate.TimestampRules} returns this
 */
proto.buf.validate.TimestampRules.prototype.setGtNow = function(value) {
  return jspb.Message.setOneofField(this, 8, proto.buf.validate.TimestampRules.oneofGroups_[1], value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.TimestampRules} returns this
 */
proto.buf.validate.TimestampRules.prototype.clearGtNow = function() {
  return jspb.Message.setOneofField(this, 8, proto.buf.validate.TimestampRules.oneofGroups_[1], undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.TimestampRules.prototype.hasGtNow = function() {
  return jspb.Message.getField(this, 8) != null;
};


/**
 * optional google.protobuf.Duration within = 9;
 * @return {?proto.google.protobuf.Duration}
 */
proto.buf.validate.TimestampRules.prototype.getWithin = function() {
  return /** @type{?proto.google.protobuf.Duration} */ (
    jspb.Message.getWrapperField(this, google_protobuf_duration_pb.Duration, 9));
};


/**
 * @param {?proto.google.protobuf.Duration|undefined} value
 * @return {!proto.buf.validate.TimestampRules} returns this
*/
proto.buf.validate.TimestampRules.prototype.setWithin = function(value) {
  return jspb.Message.setWrapperField(this, 9, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.buf.validate.TimestampRules} returns this
 */
proto.buf.validate.TimestampRules.prototype.clearWithin = function() {
  return this.setWithin(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.TimestampRules.prototype.hasWithin = function() {
  return jspb.Message.getField(this, 9) != null;
};


/**
 * repeated google.protobuf.Timestamp example = 10;
 * @return {!Array<!proto.google.protobuf.Timestamp>}
 */
proto.buf.validate.TimestampRules.prototype.getExampleList = function() {
  return /** @type{!Array<!proto.google.protobuf.Timestamp>} */ (
    jspb.Message.getRepeatedWrapperField(this, google_protobuf_timestamp_pb.Timestamp, 10));
};


/**
 * @param {!Array<!proto.google.protobuf.Timestamp>} value
 * @return {!proto.buf.validate.TimestampRules} returns this
*/
proto.buf.validate.TimestampRules.prototype.setExampleList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 10, value);
};


/**
 * @param {!proto.google.protobuf.Timestamp=} opt_value
 * @param {number=} opt_index
 * @return {!proto.google.protobuf.Timestamp}
 */
proto.buf.validate.TimestampRules.prototype.addExample = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 10, opt_value, proto.google.protobuf.Timestamp, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.buf.validate.TimestampRules} returns this
 */
proto.buf.validate.TimestampRules.prototype.clearExampleList = function() {
  return this.setExampleList([]);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.buf.validate.Violations.repeatedFields_ = [1];



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
proto.buf.validate.Violations.prototype.toObject = function(opt_includeInstance) {
  return proto.buf.validate.Violations.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.buf.validate.Violations} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.buf.validate.Violations.toObject = function(includeInstance, msg) {
  var f, obj = {
violationsList: jspb.Message.toObjectList(msg.getViolationsList(),
    proto.buf.validate.Violation.toObject, includeInstance)
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
 * @return {!proto.buf.validate.Violations}
 */
proto.buf.validate.Violations.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.buf.validate.Violations;
  return proto.buf.validate.Violations.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.buf.validate.Violations} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.buf.validate.Violations}
 */
proto.buf.validate.Violations.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.buf.validate.Violation;
      reader.readMessage(value,proto.buf.validate.Violation.deserializeBinaryFromReader);
      msg.addViolations(value);
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
proto.buf.validate.Violations.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.buf.validate.Violations.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.buf.validate.Violations} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.buf.validate.Violations.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getViolationsList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      1,
      f,
      proto.buf.validate.Violation.serializeBinaryToWriter
    );
  }
};


/**
 * repeated Violation violations = 1;
 * @return {!Array<!proto.buf.validate.Violation>}
 */
proto.buf.validate.Violations.prototype.getViolationsList = function() {
  return /** @type{!Array<!proto.buf.validate.Violation>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.buf.validate.Violation, 1));
};


/**
 * @param {!Array<!proto.buf.validate.Violation>} value
 * @return {!proto.buf.validate.Violations} returns this
*/
proto.buf.validate.Violations.prototype.setViolationsList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 1, value);
};


/**
 * @param {!proto.buf.validate.Violation=} opt_value
 * @param {number=} opt_index
 * @return {!proto.buf.validate.Violation}
 */
proto.buf.validate.Violations.prototype.addViolations = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 1, opt_value, proto.buf.validate.Violation, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.buf.validate.Violations} returns this
 */
proto.buf.validate.Violations.prototype.clearViolationsList = function() {
  return this.setViolationsList([]);
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
proto.buf.validate.Violation.prototype.toObject = function(opt_includeInstance) {
  return proto.buf.validate.Violation.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.buf.validate.Violation} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.buf.validate.Violation.toObject = function(includeInstance, msg) {
  var f, obj = {
field: (f = msg.getField()) && proto.buf.validate.FieldPath.toObject(includeInstance, f),
rule: (f = msg.getRule()) && proto.buf.validate.FieldPath.toObject(includeInstance, f),
constraintId: (f = jspb.Message.getField(msg, 2)) == null ? undefined : f,
message: (f = jspb.Message.getField(msg, 3)) == null ? undefined : f,
forKey: (f = jspb.Message.getBooleanField(msg, 4)) == null ? undefined : f
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
 * @return {!proto.buf.validate.Violation}
 */
proto.buf.validate.Violation.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.buf.validate.Violation;
  return proto.buf.validate.Violation.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.buf.validate.Violation} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.buf.validate.Violation}
 */
proto.buf.validate.Violation.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 5:
      var value = new proto.buf.validate.FieldPath;
      reader.readMessage(value,proto.buf.validate.FieldPath.deserializeBinaryFromReader);
      msg.setField(value);
      break;
    case 6:
      var value = new proto.buf.validate.FieldPath;
      reader.readMessage(value,proto.buf.validate.FieldPath.deserializeBinaryFromReader);
      msg.setRule(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setConstraintId(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setMessage(value);
      break;
    case 4:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setForKey(value);
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
proto.buf.validate.Violation.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.buf.validate.Violation.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.buf.validate.Violation} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.buf.validate.Violation.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getField();
  if (f != null) {
    writer.writeMessage(
      5,
      f,
      proto.buf.validate.FieldPath.serializeBinaryToWriter
    );
  }
  f = message.getRule();
  if (f != null) {
    writer.writeMessage(
      6,
      f,
      proto.buf.validate.FieldPath.serializeBinaryToWriter
    );
  }
  f = /** @type {string} */ (jspb.Message.getField(message, 2));
  if (f != null) {
    writer.writeString(
      2,
      f
    );
  }
  f = /** @type {string} */ (jspb.Message.getField(message, 3));
  if (f != null) {
    writer.writeString(
      3,
      f
    );
  }
  f = /** @type {boolean} */ (jspb.Message.getField(message, 4));
  if (f != null) {
    writer.writeBool(
      4,
      f
    );
  }
};


/**
 * optional FieldPath field = 5;
 * @return {?proto.buf.validate.FieldPath}
 */
proto.buf.validate.Violation.prototype.getField = function() {
  return /** @type{?proto.buf.validate.FieldPath} */ (
    jspb.Message.getWrapperField(this, proto.buf.validate.FieldPath, 5));
};


/**
 * @param {?proto.buf.validate.FieldPath|undefined} value
 * @return {!proto.buf.validate.Violation} returns this
*/
proto.buf.validate.Violation.prototype.setField = function(value) {
  return jspb.Message.setWrapperField(this, 5, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.buf.validate.Violation} returns this
 */
proto.buf.validate.Violation.prototype.clearField = function() {
  return this.setField(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.Violation.prototype.hasField = function() {
  return jspb.Message.getField(this, 5) != null;
};


/**
 * optional FieldPath rule = 6;
 * @return {?proto.buf.validate.FieldPath}
 */
proto.buf.validate.Violation.prototype.getRule = function() {
  return /** @type{?proto.buf.validate.FieldPath} */ (
    jspb.Message.getWrapperField(this, proto.buf.validate.FieldPath, 6));
};


/**
 * @param {?proto.buf.validate.FieldPath|undefined} value
 * @return {!proto.buf.validate.Violation} returns this
*/
proto.buf.validate.Violation.prototype.setRule = function(value) {
  return jspb.Message.setWrapperField(this, 6, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.buf.validate.Violation} returns this
 */
proto.buf.validate.Violation.prototype.clearRule = function() {
  return this.setRule(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.Violation.prototype.hasRule = function() {
  return jspb.Message.getField(this, 6) != null;
};


/**
 * optional string constraint_id = 2;
 * @return {string}
 */
proto.buf.validate.Violation.prototype.getConstraintId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.buf.validate.Violation} returns this
 */
proto.buf.validate.Violation.prototype.setConstraintId = function(value) {
  return jspb.Message.setField(this, 2, value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.Violation} returns this
 */
proto.buf.validate.Violation.prototype.clearConstraintId = function() {
  return jspb.Message.setField(this, 2, undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.Violation.prototype.hasConstraintId = function() {
  return jspb.Message.getField(this, 2) != null;
};


/**
 * optional string message = 3;
 * @return {string}
 */
proto.buf.validate.Violation.prototype.getMessage = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/**
 * @param {string} value
 * @return {!proto.buf.validate.Violation} returns this
 */
proto.buf.validate.Violation.prototype.setMessage = function(value) {
  return jspb.Message.setField(this, 3, value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.Violation} returns this
 */
proto.buf.validate.Violation.prototype.clearMessage = function() {
  return jspb.Message.setField(this, 3, undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.Violation.prototype.hasMessage = function() {
  return jspb.Message.getField(this, 3) != null;
};


/**
 * optional bool for_key = 4;
 * @return {boolean}
 */
proto.buf.validate.Violation.prototype.getForKey = function() {
  return /** @type {boolean} */ (jspb.Message.getBooleanFieldWithDefault(this, 4, false));
};


/**
 * @param {boolean} value
 * @return {!proto.buf.validate.Violation} returns this
 */
proto.buf.validate.Violation.prototype.setForKey = function(value) {
  return jspb.Message.setField(this, 4, value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.Violation} returns this
 */
proto.buf.validate.Violation.prototype.clearForKey = function() {
  return jspb.Message.setField(this, 4, undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.Violation.prototype.hasForKey = function() {
  return jspb.Message.getField(this, 4) != null;
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.buf.validate.FieldPath.repeatedFields_ = [1];



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
proto.buf.validate.FieldPath.prototype.toObject = function(opt_includeInstance) {
  return proto.buf.validate.FieldPath.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.buf.validate.FieldPath} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.buf.validate.FieldPath.toObject = function(includeInstance, msg) {
  var f, obj = {
elementsList: jspb.Message.toObjectList(msg.getElementsList(),
    proto.buf.validate.FieldPathElement.toObject, includeInstance)
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
 * @return {!proto.buf.validate.FieldPath}
 */
proto.buf.validate.FieldPath.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.buf.validate.FieldPath;
  return proto.buf.validate.FieldPath.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.buf.validate.FieldPath} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.buf.validate.FieldPath}
 */
proto.buf.validate.FieldPath.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.buf.validate.FieldPathElement;
      reader.readMessage(value,proto.buf.validate.FieldPathElement.deserializeBinaryFromReader);
      msg.addElements(value);
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
proto.buf.validate.FieldPath.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.buf.validate.FieldPath.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.buf.validate.FieldPath} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.buf.validate.FieldPath.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getElementsList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      1,
      f,
      proto.buf.validate.FieldPathElement.serializeBinaryToWriter
    );
  }
};


/**
 * repeated FieldPathElement elements = 1;
 * @return {!Array<!proto.buf.validate.FieldPathElement>}
 */
proto.buf.validate.FieldPath.prototype.getElementsList = function() {
  return /** @type{!Array<!proto.buf.validate.FieldPathElement>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.buf.validate.FieldPathElement, 1));
};


/**
 * @param {!Array<!proto.buf.validate.FieldPathElement>} value
 * @return {!proto.buf.validate.FieldPath} returns this
*/
proto.buf.validate.FieldPath.prototype.setElementsList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 1, value);
};


/**
 * @param {!proto.buf.validate.FieldPathElement=} opt_value
 * @param {number=} opt_index
 * @return {!proto.buf.validate.FieldPathElement}
 */
proto.buf.validate.FieldPath.prototype.addElements = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 1, opt_value, proto.buf.validate.FieldPathElement, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.buf.validate.FieldPath} returns this
 */
proto.buf.validate.FieldPath.prototype.clearElementsList = function() {
  return this.setElementsList([]);
};



/**
 * Oneof group definitions for this message. Each group defines the field
 * numbers belonging to that group. When of these fields' value is set, all
 * other fields in the group are cleared. During deserialization, if multiple
 * fields are encountered for a group, only the last value seen will be kept.
 * @private {!Array<!Array<number>>}
 * @const
 */
proto.buf.validate.FieldPathElement.oneofGroups_ = [[6,7,8,9,10]];

/**
 * @enum {number}
 */
proto.buf.validate.FieldPathElement.SubscriptCase = {
  SUBSCRIPT_NOT_SET: 0,
  INDEX: 6,
  BOOL_KEY: 7,
  INT_KEY: 8,
  UINT_KEY: 9,
  STRING_KEY: 10
};

/**
 * @return {proto.buf.validate.FieldPathElement.SubscriptCase}
 */
proto.buf.validate.FieldPathElement.prototype.getSubscriptCase = function() {
  return /** @type {proto.buf.validate.FieldPathElement.SubscriptCase} */(jspb.Message.computeOneofCase(this, proto.buf.validate.FieldPathElement.oneofGroups_[0]));
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
proto.buf.validate.FieldPathElement.prototype.toObject = function(opt_includeInstance) {
  return proto.buf.validate.FieldPathElement.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.buf.validate.FieldPathElement} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.buf.validate.FieldPathElement.toObject = function(includeInstance, msg) {
  var f, obj = {
fieldNumber: (f = jspb.Message.getField(msg, 1)) == null ? undefined : f,
fieldName: (f = jspb.Message.getField(msg, 2)) == null ? undefined : f,
fieldType: (f = jspb.Message.getField(msg, 3)) == null ? undefined : f,
keyType: (f = jspb.Message.getField(msg, 4)) == null ? undefined : f,
valueType: (f = jspb.Message.getField(msg, 5)) == null ? undefined : f,
index: (f = jspb.Message.getField(msg, 6)) == null ? undefined : f,
boolKey: (f = jspb.Message.getBooleanField(msg, 7)) == null ? undefined : f,
intKey: (f = jspb.Message.getField(msg, 8)) == null ? undefined : f,
uintKey: (f = jspb.Message.getField(msg, 9)) == null ? undefined : f,
stringKey: (f = jspb.Message.getField(msg, 10)) == null ? undefined : f
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
 * @return {!proto.buf.validate.FieldPathElement}
 */
proto.buf.validate.FieldPathElement.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.buf.validate.FieldPathElement;
  return proto.buf.validate.FieldPathElement.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.buf.validate.FieldPathElement} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.buf.validate.FieldPathElement}
 */
proto.buf.validate.FieldPathElement.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setFieldNumber(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setFieldName(value);
      break;
    case 3:
      var value = /** @type {!proto.google.protobuf.FieldDescriptorProto.Type} */ (reader.readEnum());
      msg.setFieldType(value);
      break;
    case 4:
      var value = /** @type {!proto.google.protobuf.FieldDescriptorProto.Type} */ (reader.readEnum());
      msg.setKeyType(value);
      break;
    case 5:
      var value = /** @type {!proto.google.protobuf.FieldDescriptorProto.Type} */ (reader.readEnum());
      msg.setValueType(value);
      break;
    case 6:
      var value = /** @type {number} */ (reader.readUint64());
      msg.setIndex(value);
      break;
    case 7:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setBoolKey(value);
      break;
    case 8:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setIntKey(value);
      break;
    case 9:
      var value = /** @type {number} */ (reader.readUint64());
      msg.setUintKey(value);
      break;
    case 10:
      var value = /** @type {string} */ (reader.readString());
      msg.setStringKey(value);
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
proto.buf.validate.FieldPathElement.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.buf.validate.FieldPathElement.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.buf.validate.FieldPathElement} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.buf.validate.FieldPathElement.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = /** @type {number} */ (jspb.Message.getField(message, 1));
  if (f != null) {
    writer.writeInt32(
      1,
      f
    );
  }
  f = /** @type {string} */ (jspb.Message.getField(message, 2));
  if (f != null) {
    writer.writeString(
      2,
      f
    );
  }
  f = /** @type {!proto.google.protobuf.FieldDescriptorProto.Type} */ (jspb.Message.getField(message, 3));
  if (f != null) {
    writer.writeEnum(
      3,
      f
    );
  }
  f = /** @type {!proto.google.protobuf.FieldDescriptorProto.Type} */ (jspb.Message.getField(message, 4));
  if (f != null) {
    writer.writeEnum(
      4,
      f
    );
  }
  f = /** @type {!proto.google.protobuf.FieldDescriptorProto.Type} */ (jspb.Message.getField(message, 5));
  if (f != null) {
    writer.writeEnum(
      5,
      f
    );
  }
  f = /** @type {number} */ (jspb.Message.getField(message, 6));
  if (f != null) {
    writer.writeUint64(
      6,
      f
    );
  }
  f = /** @type {boolean} */ (jspb.Message.getField(message, 7));
  if (f != null) {
    writer.writeBool(
      7,
      f
    );
  }
  f = /** @type {number} */ (jspb.Message.getField(message, 8));
  if (f != null) {
    writer.writeInt64(
      8,
      f
    );
  }
  f = /** @type {number} */ (jspb.Message.getField(message, 9));
  if (f != null) {
    writer.writeUint64(
      9,
      f
    );
  }
  f = /** @type {string} */ (jspb.Message.getField(message, 10));
  if (f != null) {
    writer.writeString(
      10,
      f
    );
  }
};


/**
 * optional int32 field_number = 1;
 * @return {number}
 */
proto.buf.validate.FieldPathElement.prototype.getFieldNumber = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/**
 * @param {number} value
 * @return {!proto.buf.validate.FieldPathElement} returns this
 */
proto.buf.validate.FieldPathElement.prototype.setFieldNumber = function(value) {
  return jspb.Message.setField(this, 1, value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.FieldPathElement} returns this
 */
proto.buf.validate.FieldPathElement.prototype.clearFieldNumber = function() {
  return jspb.Message.setField(this, 1, undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.FieldPathElement.prototype.hasFieldNumber = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional string field_name = 2;
 * @return {string}
 */
proto.buf.validate.FieldPathElement.prototype.getFieldName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.buf.validate.FieldPathElement} returns this
 */
proto.buf.validate.FieldPathElement.prototype.setFieldName = function(value) {
  return jspb.Message.setField(this, 2, value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.FieldPathElement} returns this
 */
proto.buf.validate.FieldPathElement.prototype.clearFieldName = function() {
  return jspb.Message.setField(this, 2, undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.FieldPathElement.prototype.hasFieldName = function() {
  return jspb.Message.getField(this, 2) != null;
};


/**
 * optional google.protobuf.FieldDescriptorProto.Type field_type = 3;
 * @return {!proto.google.protobuf.FieldDescriptorProto.Type}
 */
proto.buf.validate.FieldPathElement.prototype.getFieldType = function() {
  return /** @type {!proto.google.protobuf.FieldDescriptorProto.Type} */ (jspb.Message.getFieldWithDefault(this, 3, 1));
};


/**
 * @param {!proto.google.protobuf.FieldDescriptorProto.Type} value
 * @return {!proto.buf.validate.FieldPathElement} returns this
 */
proto.buf.validate.FieldPathElement.prototype.setFieldType = function(value) {
  return jspb.Message.setField(this, 3, value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.FieldPathElement} returns this
 */
proto.buf.validate.FieldPathElement.prototype.clearFieldType = function() {
  return jspb.Message.setField(this, 3, undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.FieldPathElement.prototype.hasFieldType = function() {
  return jspb.Message.getField(this, 3) != null;
};


/**
 * optional google.protobuf.FieldDescriptorProto.Type key_type = 4;
 * @return {!proto.google.protobuf.FieldDescriptorProto.Type}
 */
proto.buf.validate.FieldPathElement.prototype.getKeyType = function() {
  return /** @type {!proto.google.protobuf.FieldDescriptorProto.Type} */ (jspb.Message.getFieldWithDefault(this, 4, 1));
};


/**
 * @param {!proto.google.protobuf.FieldDescriptorProto.Type} value
 * @return {!proto.buf.validate.FieldPathElement} returns this
 */
proto.buf.validate.FieldPathElement.prototype.setKeyType = function(value) {
  return jspb.Message.setField(this, 4, value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.FieldPathElement} returns this
 */
proto.buf.validate.FieldPathElement.prototype.clearKeyType = function() {
  return jspb.Message.setField(this, 4, undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.FieldPathElement.prototype.hasKeyType = function() {
  return jspb.Message.getField(this, 4) != null;
};


/**
 * optional google.protobuf.FieldDescriptorProto.Type value_type = 5;
 * @return {!proto.google.protobuf.FieldDescriptorProto.Type}
 */
proto.buf.validate.FieldPathElement.prototype.getValueType = function() {
  return /** @type {!proto.google.protobuf.FieldDescriptorProto.Type} */ (jspb.Message.getFieldWithDefault(this, 5, 1));
};


/**
 * @param {!proto.google.protobuf.FieldDescriptorProto.Type} value
 * @return {!proto.buf.validate.FieldPathElement} returns this
 */
proto.buf.validate.FieldPathElement.prototype.setValueType = function(value) {
  return jspb.Message.setField(this, 5, value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.FieldPathElement} returns this
 */
proto.buf.validate.FieldPathElement.prototype.clearValueType = function() {
  return jspb.Message.setField(this, 5, undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.FieldPathElement.prototype.hasValueType = function() {
  return jspb.Message.getField(this, 5) != null;
};


/**
 * optional uint64 index = 6;
 * @return {number}
 */
proto.buf.validate.FieldPathElement.prototype.getIndex = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 6, 0));
};


/**
 * @param {number} value
 * @return {!proto.buf.validate.FieldPathElement} returns this
 */
proto.buf.validate.FieldPathElement.prototype.setIndex = function(value) {
  return jspb.Message.setOneofField(this, 6, proto.buf.validate.FieldPathElement.oneofGroups_[0], value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.FieldPathElement} returns this
 */
proto.buf.validate.FieldPathElement.prototype.clearIndex = function() {
  return jspb.Message.setOneofField(this, 6, proto.buf.validate.FieldPathElement.oneofGroups_[0], undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.FieldPathElement.prototype.hasIndex = function() {
  return jspb.Message.getField(this, 6) != null;
};


/**
 * optional bool bool_key = 7;
 * @return {boolean}
 */
proto.buf.validate.FieldPathElement.prototype.getBoolKey = function() {
  return /** @type {boolean} */ (jspb.Message.getBooleanFieldWithDefault(this, 7, false));
};


/**
 * @param {boolean} value
 * @return {!proto.buf.validate.FieldPathElement} returns this
 */
proto.buf.validate.FieldPathElement.prototype.setBoolKey = function(value) {
  return jspb.Message.setOneofField(this, 7, proto.buf.validate.FieldPathElement.oneofGroups_[0], value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.FieldPathElement} returns this
 */
proto.buf.validate.FieldPathElement.prototype.clearBoolKey = function() {
  return jspb.Message.setOneofField(this, 7, proto.buf.validate.FieldPathElement.oneofGroups_[0], undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.FieldPathElement.prototype.hasBoolKey = function() {
  return jspb.Message.getField(this, 7) != null;
};


/**
 * optional int64 int_key = 8;
 * @return {number}
 */
proto.buf.validate.FieldPathElement.prototype.getIntKey = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 8, 0));
};


/**
 * @param {number} value
 * @return {!proto.buf.validate.FieldPathElement} returns this
 */
proto.buf.validate.FieldPathElement.prototype.setIntKey = function(value) {
  return jspb.Message.setOneofField(this, 8, proto.buf.validate.FieldPathElement.oneofGroups_[0], value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.FieldPathElement} returns this
 */
proto.buf.validate.FieldPathElement.prototype.clearIntKey = function() {
  return jspb.Message.setOneofField(this, 8, proto.buf.validate.FieldPathElement.oneofGroups_[0], undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.FieldPathElement.prototype.hasIntKey = function() {
  return jspb.Message.getField(this, 8) != null;
};


/**
 * optional uint64 uint_key = 9;
 * @return {number}
 */
proto.buf.validate.FieldPathElement.prototype.getUintKey = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 9, 0));
};


/**
 * @param {number} value
 * @return {!proto.buf.validate.FieldPathElement} returns this
 */
proto.buf.validate.FieldPathElement.prototype.setUintKey = function(value) {
  return jspb.Message.setOneofField(this, 9, proto.buf.validate.FieldPathElement.oneofGroups_[0], value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.FieldPathElement} returns this
 */
proto.buf.validate.FieldPathElement.prototype.clearUintKey = function() {
  return jspb.Message.setOneofField(this, 9, proto.buf.validate.FieldPathElement.oneofGroups_[0], undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.FieldPathElement.prototype.hasUintKey = function() {
  return jspb.Message.getField(this, 9) != null;
};


/**
 * optional string string_key = 10;
 * @return {string}
 */
proto.buf.validate.FieldPathElement.prototype.getStringKey = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 10, ""));
};


/**
 * @param {string} value
 * @return {!proto.buf.validate.FieldPathElement} returns this
 */
proto.buf.validate.FieldPathElement.prototype.setStringKey = function(value) {
  return jspb.Message.setOneofField(this, 10, proto.buf.validate.FieldPathElement.oneofGroups_[0], value);
};


/**
 * Clears the field making it undefined.
 * @return {!proto.buf.validate.FieldPathElement} returns this
 */
proto.buf.validate.FieldPathElement.prototype.clearStringKey = function() {
  return jspb.Message.setOneofField(this, 10, proto.buf.validate.FieldPathElement.oneofGroups_[0], undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.buf.validate.FieldPathElement.prototype.hasStringKey = function() {
  return jspb.Message.getField(this, 10) != null;
};


/**
 * @enum {number}
 */
proto.buf.validate.Ignore = {
  IGNORE_UNSPECIFIED: 0,
  IGNORE_IF_UNPOPULATED: 1,
  IGNORE_IF_DEFAULT_VALUE: 2,
  IGNORE_ALWAYS: 3
};

/**
 * @enum {number}
 */
proto.buf.validate.KnownRegex = {
  KNOWN_REGEX_UNSPECIFIED: 0,
  KNOWN_REGEX_HTTP_HEADER_NAME: 1,
  KNOWN_REGEX_HTTP_HEADER_VALUE: 2
};


/**
 * A tuple of {field number, class constructor} for the extension
 * field named `message`.
 * @type {!jspb.ExtensionFieldInfo<!proto.buf.validate.MessageConstraints>}
 */
proto.buf.validate.message = new jspb.ExtensionFieldInfo(
    1159,
    {message: 0},
    proto.buf.validate.MessageConstraints,
     /** @type {?function((boolean|undefined),!jspb.Message=): !Object} */ (
         proto.buf.validate.MessageConstraints.toObject),
    0);

google_protobuf_descriptor_pb.MessageOptions.extensionsBinary[1159] = new jspb.ExtensionFieldBinaryInfo(
    proto.buf.validate.message,
    jspb.BinaryReader.prototype.readMessage,
    jspb.BinaryWriter.prototype.writeMessage,
    proto.buf.validate.MessageConstraints.serializeBinaryToWriter,
    proto.buf.validate.MessageConstraints.deserializeBinaryFromReader,
    false);
// This registers the extension field with the extended class, so that
// toObject() will function correctly.
google_protobuf_descriptor_pb.MessageOptions.extensions[1159] = proto.buf.validate.message;


/**
 * A tuple of {field number, class constructor} for the extension
 * field named `oneof`.
 * @type {!jspb.ExtensionFieldInfo<!proto.buf.validate.OneofConstraints>}
 */
proto.buf.validate.oneof = new jspb.ExtensionFieldInfo(
    1159,
    {oneof: 0},
    proto.buf.validate.OneofConstraints,
     /** @type {?function((boolean|undefined),!jspb.Message=): !Object} */ (
         proto.buf.validate.OneofConstraints.toObject),
    0);

google_protobuf_descriptor_pb.OneofOptions.extensionsBinary[1159] = new jspb.ExtensionFieldBinaryInfo(
    proto.buf.validate.oneof,
    jspb.BinaryReader.prototype.readMessage,
    jspb.BinaryWriter.prototype.writeMessage,
    proto.buf.validate.OneofConstraints.serializeBinaryToWriter,
    proto.buf.validate.OneofConstraints.deserializeBinaryFromReader,
    false);
// This registers the extension field with the extended class, so that
// toObject() will function correctly.
google_protobuf_descriptor_pb.OneofOptions.extensions[1159] = proto.buf.validate.oneof;


/**
 * A tuple of {field number, class constructor} for the extension
 * field named `field`.
 * @type {!jspb.ExtensionFieldInfo<!proto.buf.validate.FieldConstraints>}
 */
proto.buf.validate.field = new jspb.ExtensionFieldInfo(
    1159,
    {field: 0},
    proto.buf.validate.FieldConstraints,
     /** @type {?function((boolean|undefined),!jspb.Message=): !Object} */ (
         proto.buf.validate.FieldConstraints.toObject),
    0);

google_protobuf_descriptor_pb.FieldOptions.extensionsBinary[1159] = new jspb.ExtensionFieldBinaryInfo(
    proto.buf.validate.field,
    jspb.BinaryReader.prototype.readMessage,
    jspb.BinaryWriter.prototype.writeMessage,
    proto.buf.validate.FieldConstraints.serializeBinaryToWriter,
    proto.buf.validate.FieldConstraints.deserializeBinaryFromReader,
    false);
// This registers the extension field with the extended class, so that
// toObject() will function correctly.
google_protobuf_descriptor_pb.FieldOptions.extensions[1159] = proto.buf.validate.field;


/**
 * A tuple of {field number, class constructor} for the extension
 * field named `predefined`.
 * @type {!jspb.ExtensionFieldInfo<!proto.buf.validate.PredefinedConstraints>}
 */
proto.buf.validate.predefined = new jspb.ExtensionFieldInfo(
    1160,
    {predefined: 0},
    proto.buf.validate.PredefinedConstraints,
     /** @type {?function((boolean|undefined),!jspb.Message=): !Object} */ (
         proto.buf.validate.PredefinedConstraints.toObject),
    0);

google_protobuf_descriptor_pb.FieldOptions.extensionsBinary[1160] = new jspb.ExtensionFieldBinaryInfo(
    proto.buf.validate.predefined,
    jspb.BinaryReader.prototype.readMessage,
    jspb.BinaryWriter.prototype.writeMessage,
    proto.buf.validate.PredefinedConstraints.serializeBinaryToWriter,
    proto.buf.validate.PredefinedConstraints.deserializeBinaryFromReader,
    false);
// This registers the extension field with the extended class, so that
// toObject() will function correctly.
google_protobuf_descriptor_pb.FieldOptions.extensions[1160] = proto.buf.validate.predefined;

goog.object.extend(exports, proto.buf.validate);
