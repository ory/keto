/**
 * ORY Keto
 * A cloud native access control server providing best-practice patterns (RBAC, ABAC, ACL, AWS IAM Policies, Kubernetes Roles, ...) via REST APIs.
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
    define(['ApiClient', 'model/UpsertOryAccessControlPolicyRoleInternalServerErrorBody'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'), require('./UpsertOryAccessControlPolicyRoleInternalServerErrorBody'));
  } else {
    // Browser globals (root is window)
    if (!root.OryKeto) {
      root.OryKeto = {};
    }
    root.OryKeto.UpsertOryAccessControlPolicyRoleInternalServerError = factory(root.OryKeto.ApiClient, root.OryKeto.UpsertOryAccessControlPolicyRoleInternalServerErrorBody);
  }
}(this, function(ApiClient, UpsertOryAccessControlPolicyRoleInternalServerErrorBody) {
  'use strict';




  /**
   * The UpsertOryAccessControlPolicyRoleInternalServerError model module.
   * @module model/UpsertOryAccessControlPolicyRoleInternalServerError
   * @version Latest
   */

  /**
   * Constructs a new <code>UpsertOryAccessControlPolicyRoleInternalServerError</code>.
   * The standard error format
   * @alias module:model/UpsertOryAccessControlPolicyRoleInternalServerError
   * @class
   */
  var exports = function() {
    var _this = this;


  };

  /**
   * Constructs a <code>UpsertOryAccessControlPolicyRoleInternalServerError</code> from a plain JavaScript object, optionally creating a new instance.
   * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
   * @param {Object} data The plain JavaScript object bearing properties of interest.
   * @param {module:model/UpsertOryAccessControlPolicyRoleInternalServerError} obj Optional instance to populate.
   * @return {module:model/UpsertOryAccessControlPolicyRoleInternalServerError} The populated <code>UpsertOryAccessControlPolicyRoleInternalServerError</code> instance.
   */
  exports.constructFromObject = function(data, obj) {
    if (data) {
      obj = obj || new exports();

      if (data.hasOwnProperty('Payload')) {
        obj['Payload'] = UpsertOryAccessControlPolicyRoleInternalServerErrorBody.constructFromObject(data['Payload']);
      }
    }
    return obj;
  }

  /**
   * @member {module:model/UpsertOryAccessControlPolicyRoleInternalServerErrorBody} Payload
   */
  exports.prototype['Payload'] = undefined;



  return exports;
}));


