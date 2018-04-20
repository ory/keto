# keto\SDK\WardenApi
Client for keto

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**isOAuth2AccessTokenAuthorized**](WardenApi.md#isOAuth2AccessTokenAuthorized) | **POST** /warden/oauth2/authorize | Check if an OAuth 2.0 access token is authorized to access a resource
[**isSubjectAuthorized**](WardenApi.md#isSubjectAuthorized) | **POST** /warden/subjects/authorize | Check if a subject is authorized to access a resource


# **isOAuth2AccessTokenAuthorized**
> \keto\SDK\Model\WardenOAuth2AuthorizationResponse isOAuth2AccessTokenAuthorized($body)

Check if an OAuth 2.0 access token is authorized to access a resource

Checks if a token is valid and if the token subject is allowed to perform an action on a resource. This endpoint requires a token, a scope, a resource name, an action name and a context.   If a token is expired/invalid, has not been granted the requested scope or the subject is not allowed to perform the action on the resource, this endpoint returns a 200 response with `{ \"allowed\": false }`.   This endpoint passes all data from the upstream OAuth 2.0 token introspection endpoint. If you use ORY Hydra as an upstream OAuth 2.0 provider, data set through the `accessTokenExtra` field in the consent flow will be included in this response as well.

### Example
```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');

$api_instance = new keto\SDK\Api\WardenApi();
$body = new \keto\SDK\Model\WardenOAuth2AuthorizationRequest(); // \keto\SDK\Model\WardenOAuth2AuthorizationRequest | 

try {
    $result = $api_instance->isOAuth2AccessTokenAuthorized($body);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling WardenApi->isOAuth2AccessTokenAuthorized: ', $e->getMessage(), PHP_EOL;
}
?>
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**\keto\SDK\Model\WardenOAuth2AuthorizationRequest**](../Model/WardenOAuth2AuthorizationRequest.md)|  | [optional]

### Return type

[**\keto\SDK\Model\WardenOAuth2AuthorizationResponse**](../Model/WardenOAuth2AuthorizationResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **isSubjectAuthorized**
> \keto\SDK\Model\WardenSubjectAuthorizationResponse isSubjectAuthorized($body)

Check if a subject is authorized to access a resource

Checks if a subject (e.g. user ID, API key, ...) is allowed to perform a certain action on a resource.

### Example
```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');

$api_instance = new keto\SDK\Api\WardenApi();
$body = new \keto\SDK\Model\WardenSubjectAuthorizationRequest(); // \keto\SDK\Model\WardenSubjectAuthorizationRequest | 

try {
    $result = $api_instance->isSubjectAuthorized($body);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling WardenApi->isSubjectAuthorized: ', $e->getMessage(), PHP_EOL;
}
?>
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**\keto\SDK\Model\WardenSubjectAuthorizationRequest**](../Model/WardenSubjectAuthorizationRequest.md)|  | [optional]

### Return type

[**\keto\SDK\Model\WardenSubjectAuthorizationResponse**](../Model/WardenSubjectAuthorizationResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

