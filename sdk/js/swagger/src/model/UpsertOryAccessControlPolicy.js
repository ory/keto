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
    define(['ApiClient', 'model/OryAccessControlPolicy'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'), require('./OryAccessControlPolicy'));
  } else {
    // Browser globals (root is window)
    if (!root.SwaggerJsClient) {
      root.SwaggerJsClient = {};
    }
    root.SwaggerJsClient.UpsertOryAccessControlPolicy = factory(root.SwaggerJsClient.ApiClient, root.SwaggerJsClient.OryAccessControlPolicy);
  }
}(this, function(ApiClient, OryAccessControlPolicy) {
  'use strict';




  /**
   * The UpsertOryAccessControlPolicy model module.
   * @module model/UpsertOryAccessControlPolicy
   * @version Latest
   */

  /**
   * Constructs a new <code>UpsertOryAccessControlPolicy</code>.
   * @alias module:model/UpsertOryAccessControlPolicy
   * @class
   * @param flavor {String} The ORY Access Control Policy flavor. Can be \"regex\" and \"exact\".  in: path
   */
  var exports = function(flavor) {
    var _this = this;


    _this['flavor'] = flavor;
  };

  /**
   * Constructs a <code>UpsertOryAccessControlPolicy</code> from a plain JavaScript object, optionally creating a new instance.
   * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
   * @param {Object} data The plain JavaScript object bearing properties of interest.
   * @param {module:model/UpsertOryAccessControlPolicy} obj Optional instance to populate.
   * @return {module:model/UpsertOryAccessControlPolicy} The populated <code>UpsertOryAccessControlPolicy</code> instance.
   */
  exports.constructFromObject = function(data, obj) {
    if (data) {
      obj = obj || new exports();

      if (data.hasOwnProperty('Body')) {
        obj['Body'] = OryAccessControlPolicy.constructFromObject(data['Body']);
      }
      if (data.hasOwnProperty('flavor')) {
        obj['flavor'] = ApiClient.convertToType(data['flavor'], 'String');
      }
    }
    return obj;
  }

  /**
   * @member {module:model/OryAccessControlPolicy} Body
   */
  exports.prototype['Body'] = undefined;
  /**
   * The ORY Access Control Policy flavor. Can be \"regex\" and \"exact\".  in: path
   * @member {String} flavor
   */
  exports.prototype['flavor'] = undefined;



  return exports;
}));


