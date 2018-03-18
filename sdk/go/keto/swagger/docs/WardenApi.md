# \WardenApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IsOAuth2AccessTokenAuthorized**](WardenApi.md#IsOAuth2AccessTokenAuthorized) | **Post** /warden/oauth2/authorize | Check if an OAuth 2.0 access token is authorized to access a resource
[**IsSubjectAuthorized**](WardenApi.md#IsSubjectAuthorized) | **Post** /warden/subjects/authorize | Check if a subject is authorized to access a resource


# **IsOAuth2AccessTokenAuthorized**
> IsOAuth2AccessTokenAuthorized($body)

Check if an OAuth 2.0 access token is authorized to access a resource

Checks if a token is valid and if the token subject is allowed to perform an action on a resource. This endpoint requires a token, a scope, a resource name, an action name and a context.   If a token is expired/invalid, has not been granted the requested scope or the subject is not allowed to perform the action on the resource, this endpoint returns a 200 response with `{ \"allowed\": false }`.   This endpoint passes all data from the upstream OAuth 2.0 token introspection endpoint. If you use ORY Hydra as an upstream OAuth 2.0 provider, data set through the `accessTokenExtra` field in the consent flow will be included in this response as well.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**WardenOAuth2AuthorizationRequest**](WardenOAuth2AuthorizationRequest.md)|  | [optional] 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IsSubjectAuthorized**
> IsSubjectAuthorized($body)

Check if a subject is authorized to access a resource

Checks if a subject (e.g. user ID, API key, ...) is allowed to perform a certain action on a resource.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**WardenSubjectAuthorizationRequest**](WardenSubjectAuthorizationRequest.md)|  | [optional] 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

