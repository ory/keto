# SwaggerJsClient.WardenApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**isOAuth2AccessTokenAuthorized**](WardenApi.md#isOAuth2AccessTokenAuthorized) | **POST** /warden/oauth2/access-tokens/authorize | Check if an OAuth 2.0 access token is authorized to access a resource
[**isOAuth2ClientAuthorized**](WardenApi.md#isOAuth2ClientAuthorized) | **POST** /warden/oauth2/clients/authorize | Check if an OAuth 2.0 Client is authorized to access a resource
[**isSubjectAuthorized**](WardenApi.md#isSubjectAuthorized) | **POST** /warden/subjects/authorize | Check if a subject is authorized to access a resource


<a name="isOAuth2AccessTokenAuthorized"></a>
# **isOAuth2AccessTokenAuthorized**
> WardenOAuth2AccessTokenAuthorizationResponse isOAuth2AccessTokenAuthorized(opts)

Check if an OAuth 2.0 access token is authorized to access a resource

Checks if a token is valid and if the token subject is allowed to perform an action on a resource. This endpoint requires a token, a scope, a resource name, an action name and a context.   If a token is expired/invalid, has not been granted the requested scope or the subject is not allowed to perform the action on the resource, this endpoint returns a 200 response with &#x60;{ \&quot;allowed\&quot;: false }&#x60;.   This endpoint passes all data from the upstream OAuth 2.0 token introspection endpoint. If you use ORY Hydra as an upstream OAuth 2.0 provider, data set through the &#x60;accessTokenExtra&#x60; field in the consent flow will be included in this response as well.

### Example
```javascript
var SwaggerJsClient = require('swagger-js-client');

var apiInstance = new SwaggerJsClient.WardenApi();

var opts = { 
  'body': new SwaggerJsClient.WardenOAuth2AccessTokenAuthorizationRequest() // WardenOAuth2AccessTokenAuthorizationRequest | 
};

var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
};
apiInstance.isOAuth2AccessTokenAuthorized(opts, callback);
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**WardenOAuth2AccessTokenAuthorizationRequest**](WardenOAuth2AccessTokenAuthorizationRequest.md)|  | [optional] 

### Return type

[**WardenOAuth2AccessTokenAuthorizationResponse**](WardenOAuth2AccessTokenAuthorizationResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="isOAuth2ClientAuthorized"></a>
# **isOAuth2ClientAuthorized**
> WardenOAuth2ClientAuthorizationResponse isOAuth2ClientAuthorized(opts)

Check if an OAuth 2.0 Client is authorized to access a resource

Checks if an OAuth 2.0 Client provided the correct access credentials and and if the client is allowed to perform an action on a resource. This endpoint requires a client id, a client secret, a scope, a resource name, an action name and a context.

### Example
```javascript
var SwaggerJsClient = require('swagger-js-client');

var apiInstance = new SwaggerJsClient.WardenApi();

var opts = { 
  'body': new SwaggerJsClient.WardenOAuth2ClientAuthorizationRequest() // WardenOAuth2ClientAuthorizationRequest | 
};

var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
};
apiInstance.isOAuth2ClientAuthorized(opts, callback);
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**WardenOAuth2ClientAuthorizationRequest**](WardenOAuth2ClientAuthorizationRequest.md)|  | [optional] 

### Return type

[**WardenOAuth2ClientAuthorizationResponse**](WardenOAuth2ClientAuthorizationResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="isSubjectAuthorized"></a>
# **isSubjectAuthorized**
> WardenSubjectAuthorizationResponse isSubjectAuthorized(opts)

Check if a subject is authorized to access a resource

Checks if a subject (e.g. user ID, API key, ...) is allowed to perform a certain action on a resource.

### Example
```javascript
var SwaggerJsClient = require('swagger-js-client');

var apiInstance = new SwaggerJsClient.WardenApi();

var opts = { 
  'body': new SwaggerJsClient.WardenSubjectAuthorizationRequest() // WardenSubjectAuthorizationRequest | 
};

var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
};
apiInstance.isSubjectAuthorized(opts, callback);
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**WardenSubjectAuthorizationRequest**](WardenSubjectAuthorizationRequest.md)|  | [optional] 

### Return type

[**WardenSubjectAuthorizationResponse**](WardenSubjectAuthorizationResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

