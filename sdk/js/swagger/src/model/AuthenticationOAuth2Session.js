/**
 * 
 * Package main ORY Keto
 *
 * OpenAPI spec version: Latest
 * Contact: hi@ory.am
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 *
 * Swagger Codegen version: 2.2.3
 *
 * Do not edit the class manually.
 *
 */

;(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD. Register as an anonymous module.
    define(['ApiClient'], factory)
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'))
  } else {
    // Browser globals (root is window)
    if (!root.SwaggerJsClient) {
      root.SwaggerJsClient = {}
    }
    root.SwaggerJsClient.AuthenticationOAuth2Session = factory(
      root.SwaggerJsClient.ApiClient
    )
  }
})(this, function(ApiClient) {
  'use strict'

  /**
   * The AuthenticationOAuth2Session model module.
   * @module model/AuthenticationOAuth2Session
   * @version Latest
   */

  /**
   * Constructs a new <code>AuthenticationOAuth2Session</code>.
   * @alias module:model/AuthenticationOAuth2Session
   * @class
   */
  var exports = function() {
    var _this = this
  }

  /**
   * Constructs a <code>AuthenticationOAuth2Session</code> from a plain JavaScript object, optionally creating a new instance.
   * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
   * @param {Object} data The plain JavaScript object bearing properties of interest.
   * @param {module:model/AuthenticationOAuth2Session} obj Optional instance to populate.
   * @return {module:model/AuthenticationOAuth2Session} The populated <code>AuthenticationOAuth2Session</code> instance.
   */
  exports.constructFromObject = function(data, obj) {
    if (data) {
      obj = obj || new exports()

      if (data.hasOwnProperty('allowed')) {
        obj['allowed'] = ApiClient.convertToType(data['allowed'], 'Boolean')
      }
      if (data.hasOwnProperty('aud')) {
        obj['aud'] = ApiClient.convertToType(data['aud'], ['String'])
      }
      if (data.hasOwnProperty('client_id')) {
        obj['client_id'] = ApiClient.convertToType(data['client_id'], 'String')
      }
      if (data.hasOwnProperty('exp')) {
        obj['exp'] = ApiClient.convertToType(data['exp'], 'Date')
      }
      if (data.hasOwnProperty('iat')) {
        obj['iat'] = ApiClient.convertToType(data['iat'], 'Date')
      }
      if (data.hasOwnProperty('iss')) {
        obj['iss'] = ApiClient.convertToType(data['iss'], 'String')
      }
      if (data.hasOwnProperty('nbf')) {
        obj['nbf'] = ApiClient.convertToType(data['nbf'], 'Date')
      }
      if (data.hasOwnProperty('scope')) {
        obj['scope'] = ApiClient.convertToType(data['scope'], 'String')
      }
      if (data.hasOwnProperty('session')) {
        obj['session'] = ApiClient.convertToType(data['session'], {
          String: Object
        })
      }
      if (data.hasOwnProperty('sub')) {
        obj['sub'] = ApiClient.convertToType(data['sub'], 'String')
      }
      if (data.hasOwnProperty('username')) {
        obj['username'] = ApiClient.convertToType(data['username'], 'String')
      }
    }
    return obj
  }

  /**
   * Allowed is true if the request is allowed and false otherwise.
   * @member {Boolean} allowed
   */
  exports.prototype['allowed'] = undefined
  /**
   * @member {Array.<String>} aud
   */
  exports.prototype['aud'] = undefined
  /**
   * ClientID is the id of the OAuth2 client that requested the token.
   * @member {String} client_id
   */
  exports.prototype['client_id'] = undefined
  /**
   * ExpiresAt is the expiry timestamp.
   * @member {Date} exp
   */
  exports.prototype['exp'] = undefined
  /**
   * IssuedAt is the token creation time stamp.
   * @member {Date} iat
   */
  exports.prototype['iat'] = undefined
  /**
   * Issuer is the id of the issuer, typically an hydra instance.
   * @member {String} iss
   */
  exports.prototype['iss'] = undefined
  /**
   * @member {Date} nbf
   */
  exports.prototype['nbf'] = undefined
  /**
   * GrantedScopes is a list of scopes that the subject authorized when asked for consent.
   * @member {String} scope
   */
  exports.prototype['scope'] = undefined
  /**
   * Session represents arbitrary session data.
   * @member {Object.<String, Object>} session
   */
  exports.prototype['session'] = undefined
  /**
   * Subject is the identity that authorized issuing the token, for example a user or an OAuth2 app. This is usually a uuid but you can choose a urn or some other id too.
   * @member {String} sub
   */
  exports.prototype['sub'] = undefined
  /**
   * @member {String} username
   */
  exports.prototype['username'] = undefined

  return exports
})
