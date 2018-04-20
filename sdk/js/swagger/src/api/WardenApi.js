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
    define(
      [
        'ApiClient',
        'model/InlineResponse401',
        'model/WardenOAuth2AuthorizationRequest',
        'model/WardenOAuth2AuthorizationResponse',
        'model/WardenSubjectAuthorizationRequest',
        'model/WardenSubjectAuthorizationResponse'
      ],
      factory
    )
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(
      require('../ApiClient'),
      require('../model/InlineResponse401'),
      require('../model/WardenOAuth2AuthorizationRequest'),
      require('../model/WardenOAuth2AuthorizationResponse'),
      require('../model/WardenSubjectAuthorizationRequest'),
      require('../model/WardenSubjectAuthorizationResponse')
    )
  } else {
    // Browser globals (root is window)
    if (!root.SwaggerJsClient) {
      root.SwaggerJsClient = {}
    }
    root.SwaggerJsClient.WardenApi = factory(
      root.SwaggerJsClient.ApiClient,
      root.SwaggerJsClient.InlineResponse401,
      root.SwaggerJsClient.WardenOAuth2AuthorizationRequest,
      root.SwaggerJsClient.WardenOAuth2AuthorizationResponse,
      root.SwaggerJsClient.WardenSubjectAuthorizationRequest,
      root.SwaggerJsClient.WardenSubjectAuthorizationResponse
    )
  }
})(this, function(
  ApiClient,
  InlineResponse401,
  WardenOAuth2AuthorizationRequest,
  WardenOAuth2AuthorizationResponse,
  WardenSubjectAuthorizationRequest,
  WardenSubjectAuthorizationResponse
) {
  'use strict'

  /**
   * Warden service.
   * @module api/WardenApi
   * @version Latest
   */

  /**
   * Constructs a new WardenApi. 
   * @alias module:api/WardenApi
   * @class
   * @param {module:ApiClient} apiClient Optional API client implementation to use,
   * default to {@link module:ApiClient#instance} if unspecified.
   */
  var exports = function(apiClient) {
    this.apiClient = apiClient || ApiClient.instance

    /**
     * Callback function to receive the result of the isOAuth2AccessTokenAuthorized operation.
     * @callback module:api/WardenApi~isOAuth2AccessTokenAuthorizedCallback
     * @param {String} error Error message, if any.
     * @param {module:model/WardenOAuth2AuthorizationResponse} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Check if an OAuth 2.0 access token is authorized to access a resource
     * Checks if a token is valid and if the token subject is allowed to perform an action on a resource. This endpoint requires a token, a scope, a resource name, an action name and a context.   If a token is expired/invalid, has not been granted the requested scope or the subject is not allowed to perform the action on the resource, this endpoint returns a 200 response with &#x60;{ \&quot;allowed\&quot;: false }&#x60;.   This endpoint passes all data from the upstream OAuth 2.0 token introspection endpoint. If you use ORY Hydra as an upstream OAuth 2.0 provider, data set through the &#x60;accessTokenExtra&#x60; field in the consent flow will be included in this response as well.
     * @param {Object} opts Optional parameters
     * @param {module:model/WardenOAuth2AuthorizationRequest} opts.body 
     * @param {module:api/WardenApi~isOAuth2AccessTokenAuthorizedCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/WardenOAuth2AuthorizationResponse}
     */
    this.isOAuth2AccessTokenAuthorized = function(opts, callback) {
      opts = opts || {}
      var postBody = opts['body']

      var pathParams = {}
      var queryParams = {}
      var headerParams = {}
      var formParams = {}

      var authNames = []
      var contentTypes = ['application/json']
      var accepts = ['application/json']
      var returnType = WardenOAuth2AuthorizationResponse

      return this.apiClient.callApi(
        '/warden/oauth2/authorize',
        'POST',
        pathParams,
        queryParams,
        headerParams,
        formParams,
        postBody,
        authNames,
        contentTypes,
        accepts,
        returnType,
        callback
      )
    }

    /**
     * Callback function to receive the result of the isSubjectAuthorized operation.
     * @callback module:api/WardenApi~isSubjectAuthorizedCallback
     * @param {String} error Error message, if any.
     * @param {module:model/WardenSubjectAuthorizationResponse} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Check if a subject is authorized to access a resource
     * Checks if a subject (e.g. user ID, API key, ...) is allowed to perform a certain action on a resource.
     * @param {Object} opts Optional parameters
     * @param {module:model/WardenSubjectAuthorizationRequest} opts.body 
     * @param {module:api/WardenApi~isSubjectAuthorizedCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/WardenSubjectAuthorizationResponse}
     */
    this.isSubjectAuthorized = function(opts, callback) {
      opts = opts || {}
      var postBody = opts['body']

      var pathParams = {}
      var queryParams = {}
      var headerParams = {}
      var formParams = {}

      var authNames = []
      var contentTypes = ['application/json']
      var accepts = ['application/json']
      var returnType = WardenSubjectAuthorizationResponse

      return this.apiClient.callApi(
        '/warden/subjects/authorize',
        'POST',
        pathParams,
        queryParams,
        headerParams,
        formParams,
        postBody,
        authNames,
        contentTypes,
        accepts,
        returnType,
        callback
      )
    }
  }

  return exports
})
