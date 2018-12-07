/**
 * 
 * Package main ORY Keto
 *
 * OpenAPI spec version: Latest
 * Contact: hi@ory.sh
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 *
 * Swagger Codegen version: 2.2.3
 *
 * Do not edit the class manually.
 *
 */

(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD. Register as an anonymous module.
    define(['ApiClient'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'));
  } else {
    // Browser globals (root is window)
    if (!root.SwaggerJsClient) {
      root.SwaggerJsClient = {};
    }
    root.SwaggerJsClient.DeleteOryAccessControlPolicyRole = factory(root.SwaggerJsClient.ApiClient);
  }
}(this, function(ApiClient) {
  'use strict';




  /**
   * The DeleteOryAccessControlPolicyRole model module.
   * @module model/DeleteOryAccessControlPolicyRole
   * @version Latest
   */

  /**
   * Constructs a new <code>DeleteOryAccessControlPolicyRole</code>.
   * @alias module:model/DeleteOryAccessControlPolicyRole
   * @class
   * @param flavor {String} The ORY Access Control Policy flavor. Can be \"regex\" and \"exact\".  in: path
   * @param id {String} The ID of the ORY Access Control Policy Role. in: path
   */
  var exports = function(flavor, id) {
    var _this = this;

    _this['flavor'] = flavor;
    _this['id'] = id;
  };

  /**
   * Constructs a <code>DeleteOryAccessControlPolicyRole</code> from a plain JavaScript object, optionally creating a new instance.
   * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
   * @param {Object} data The plain JavaScript object bearing properties of interest.
   * @param {module:model/DeleteOryAccessControlPolicyRole} obj Optional instance to populate.
   * @return {module:model/DeleteOryAccessControlPolicyRole} The populated <code>DeleteOryAccessControlPolicyRole</code> instance.
   */
  exports.constructFromObject = function(data, obj) {
    if (data) {
      obj = obj || new exports();

      if (data.hasOwnProperty('flavor')) {
        obj['flavor'] = ApiClient.convertToType(data['flavor'], 'String');
      }
      if (data.hasOwnProperty('id')) {
        obj['id'] = ApiClient.convertToType(data['id'], 'String');
      }
    }
    return obj;
  }

  /**
   * The ORY Access Control Policy flavor. Can be \"regex\" and \"exact\".  in: path
   * @member {String} flavor
   */
  exports.prototype['flavor'] = undefined;
  /**
   * The ID of the ORY Access Control Policy Role. in: path
   * @member {String} id
   */
  exports.prototype['id'] = undefined;



  return exports;
}));


