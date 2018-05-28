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
    define(['ApiClient', 'model/InlineResponse401', 'model/Policy'], factory)
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(
      require('../ApiClient'),
      require('../model/InlineResponse401'),
      require('../model/Policy')
    )
  } else {
    // Browser globals (root is window)
    if (!root.SwaggerJsClient) {
      root.SwaggerJsClient = {}
    }
    root.SwaggerJsClient.PolicyApi = factory(
      root.SwaggerJsClient.ApiClient,
      root.SwaggerJsClient.InlineResponse401,
      root.SwaggerJsClient.Policy
    )
  }
})(this, function(ApiClient, InlineResponse401, Policy) {
  'use strict'

  /**
   * Policy service.
   * @module api/PolicyApi
   * @version Latest
   */

  /**
   * Constructs a new PolicyApi. 
   * @alias module:api/PolicyApi
   * @class
   * @param {module:ApiClient} apiClient Optional API client implementation to use,
   * default to {@link module:ApiClient#instance} if unspecified.
   */
  var exports = function(apiClient) {
    this.apiClient = apiClient || ApiClient.instance

    /**
     * Callback function to receive the result of the createPolicy operation.
     * @callback module:api/PolicyApi~createPolicyCallback
     * @param {String} error Error message, if any.
     * @param {module:model/Policy} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Create an Access Control Policy
     * @param {Object} opts Optional parameters
     * @param {module:model/Policy} opts.body 
     * @param {module:api/PolicyApi~createPolicyCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/Policy}
     */
    this.createPolicy = function(opts, callback) {
      opts = opts || {}
      var postBody = opts['body']

      var pathParams = {}
      var queryParams = {}
      var headerParams = {}
      var formParams = {}

      var authNames = []
      var contentTypes = ['application/json']
      var accepts = ['application/json']
      var returnType = Policy

      return this.apiClient.callApi(
        '/policies',
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
     * Callback function to receive the result of the deletePolicy operation.
     * @callback module:api/PolicyApi~deletePolicyCallback
     * @param {String} error Error message, if any.
     * @param data This operation does not return a value.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Delete an Access Control Policy
     * @param {String} id The id of the policy.
     * @param {module:api/PolicyApi~deletePolicyCallback} callback The callback function, accepting three arguments: error, data, response
     */
    this.deletePolicy = function(id, callback) {
      var postBody = null

      // verify the required parameter 'id' is set
      if (id === undefined || id === null) {
        throw new Error(
          "Missing the required parameter 'id' when calling deletePolicy"
        )
      }

      var pathParams = {
        id: id
      }
      var queryParams = {}
      var headerParams = {}
      var formParams = {}

      var authNames = []
      var contentTypes = ['application/json']
      var accepts = ['application/json']
      var returnType = null

      return this.apiClient.callApi(
        '/policies/{id}',
        'DELETE',
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
     * Callback function to receive the result of the getPolicy operation.
     * @callback module:api/PolicyApi~getPolicyCallback
     * @param {String} error Error message, if any.
     * @param {module:model/Policy} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Get an Access Control Policy
     * @param {String} id The id of the policy.
     * @param {module:api/PolicyApi~getPolicyCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/Policy}
     */
    this.getPolicy = function(id, callback) {
      var postBody = null

      // verify the required parameter 'id' is set
      if (id === undefined || id === null) {
        throw new Error(
          "Missing the required parameter 'id' when calling getPolicy"
        )
      }

      var pathParams = {
        id: id
      }
      var queryParams = {}
      var headerParams = {}
      var formParams = {}

      var authNames = []
      var contentTypes = ['application/json']
      var accepts = ['application/json']
      var returnType = Policy

      return this.apiClient.callApi(
        '/policies/{id}',
        'GET',
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
     * Callback function to receive the result of the listPolicies operation.
     * @callback module:api/PolicyApi~listPoliciesCallback
     * @param {String} error Error message, if any.
     * @param {Array.<module:model/Policy>} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * List Access Control Policies
     * @param {Object} opts Optional parameters
     * @param {Number} opts.offset The offset from where to start looking.
     * @param {Number} opts.limit The maximum amount of policies returned.
     * @param {module:api/PolicyApi~listPoliciesCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link Array.<module:model/Policy>}
     */
    this.listPolicies = function(opts, callback) {
      opts = opts || {}
      var postBody = null

      var pathParams = {}
      var queryParams = {
        offset: opts['offset'],
        limit: opts['limit']
      }
      var headerParams = {}
      var formParams = {}

      var authNames = []
      var contentTypes = ['application/json']
      var accepts = ['application/json']
      var returnType = [Policy]

      return this.apiClient.callApi(
        '/policies',
        'GET',
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
     * Callback function to receive the result of the updatePolicy operation.
     * @callback module:api/PolicyApi~updatePolicyCallback
     * @param {String} error Error message, if any.
     * @param {module:model/Policy} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Update an Access Control Policy
     * @param {String} id The id of the policy.
     * @param {Object} opts Optional parameters
     * @param {module:model/Policy} opts.body 
     * @param {module:api/PolicyApi~updatePolicyCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/Policy}
     */
    this.updatePolicy = function(id, opts, callback) {
      opts = opts || {}
      var postBody = opts['body']

      // verify the required parameter 'id' is set
      if (id === undefined || id === null) {
        throw new Error(
          "Missing the required parameter 'id' when calling updatePolicy"
        )
      }

      var pathParams = {
        id: id
      }
      var queryParams = {}
      var headerParams = {}
      var formParams = {}

      var authNames = []
      var contentTypes = ['application/json']
      var accepts = ['application/json']
      var returnType = Policy

      return this.apiClient.callApi(
        '/policies/{id}',
        'PUT',
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
