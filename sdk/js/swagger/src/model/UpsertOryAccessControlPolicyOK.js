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
    define(['ApiClient', 'model/Policy'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'), require('./Policy'));
  } else {
    // Browser globals (root is window)
    if (!root.OryKeto) {
      root.OryKeto = {};
    }
    root.OryKeto.UpsertOryAccessControlPolicyOK = factory(root.OryKeto.ApiClient, root.OryKeto.Policy);
  }
}(this, function(ApiClient, Policy) {
  'use strict';




  /**
   * The UpsertOryAccessControlPolicyOK model module.
   * @module model/UpsertOryAccessControlPolicyOK
   * @version Latest
   */

  /**
   * Constructs a new <code>UpsertOryAccessControlPolicyOK</code>.
   * oryAccessControlPolicy
   * @alias module:model/UpsertOryAccessControlPolicyOK
   * @class
   */
  var exports = function() {
    var _this = this;


  };

  /**
   * Constructs a <code>UpsertOryAccessControlPolicyOK</code> from a plain JavaScript object, optionally creating a new instance.
   * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
   * @param {Object} data The plain JavaScript object bearing properties of interest.
   * @param {module:model/UpsertOryAccessControlPolicyOK} obj Optional instance to populate.
   * @return {module:model/UpsertOryAccessControlPolicyOK} The populated <code>UpsertOryAccessControlPolicyOK</code> instance.
   */
  exports.constructFromObject = function(data, obj) {
    if (data) {
      obj = obj || new exports();

      if (data.hasOwnProperty('Payload')) {
        obj['Payload'] = Policy.constructFromObject(data['Payload']);
      }
    }
    return obj;
  }

  /**
   * @member {module:model/Policy} Payload
   */
  exports.prototype['Payload'] = undefined;



  return exports;
}));


