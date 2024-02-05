"use strict";
exports.__esModule = true;
exports.FieldMask = exports.protobufPackage = void 0;
/* eslint-disable */
var minimal_js_1 = require("protobufjs/minimal.js");
exports.protobufPackage = "google.protobuf";
function createBaseFieldMask() {
    return { paths: [] };
}
exports.FieldMask = {
    encode: function (message, writer) {
        if (writer === void 0) { writer = minimal_js_1["default"].Writer.create(); }
        for (var _i = 0, _a = message.paths; _i < _a.length; _i++) {
            var v = _a[_i];
            writer.uint32(10).string(v);
        }
        return writer;
    },
    decode: function (input, length) {
        var reader = input instanceof minimal_js_1["default"].Reader ? input : minimal_js_1["default"].Reader.create(input);
        var end = length === undefined ? reader.len : reader.pos + length;
        var message = createBaseFieldMask();
        while (reader.pos < end) {
            var tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    if (tag !== 10) {
                        break;
                    }
                    message.paths.push(reader.string());
                    continue;
            }
            if ((tag & 7) === 4 || tag === 0) {
                break;
            }
            reader.skipType(tag & 7);
        }
        return message;
    },
    fromJSON: function (object) {
        return {
            paths: typeof object === "string"
                ? object.split(",").filter(globalThis.Boolean)
                : globalThis.Array.isArray(object === null || object === void 0 ? void 0 : object.paths)
                    ? object.paths.map(globalThis.String)
                    : []
        };
    },
    toJSON: function (message) {
        return message.paths.join(",");
    },
    create: function (base) {
        return exports.FieldMask.fromPartial(base !== null && base !== void 0 ? base : {});
    },
    fromPartial: function (object) {
        var _a;
        var message = createBaseFieldMask();
        message.paths = ((_a = object.paths) === null || _a === void 0 ? void 0 : _a.map(function (e) { return e; })) || [];
        return message;
    },
    wrap: function (paths) {
        var result = createBaseFieldMask();
        result.paths = paths;
        return result;
    },
    unwrap: function (message) {
        return message.paths;
    }
};
