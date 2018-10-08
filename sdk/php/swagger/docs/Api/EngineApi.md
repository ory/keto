# keto\SDK\EngineApi
Client for keto

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**listOryAccessControlPolicyRoles**](EngineApi.md#listOryAccessControlPolicyRoles) | **GET** /engines/ory/{flavor}/roles | List ORY Access Control Policy Roles


# **listOryAccessControlPolicyRoles**
> \keto\SDK\Model\OryAccessControlPolicyRoles listOryAccessControlPolicyRoles($flavor, $limit, $offset)

List ORY Access Control Policy Roles

Roles group several subjects into one. Rules can be assigned to ORY Access Control Policy (OACP) by using the Role ID as subject in the OACP.

### Example
```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');

$api_instance = new keto\SDK\Api\EngineApi();
$flavor = "flavor_example"; // string | The ORY Access Control Policy flavor. Can be \"regex\" and \"exact\"
$limit = 789; // int | The maximum amount of policies returned.
$offset = 789; // int | The offset from where to start looking.

try {
    $result = $api_instance->listOryAccessControlPolicyRoles($flavor, $limit, $offset);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling EngineApi->listOryAccessControlPolicyRoles: ', $e->getMessage(), PHP_EOL;
}
?>
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **flavor** | **string**| The ORY Access Control Policy flavor. Can be \&quot;regex\&quot; and \&quot;exact\&quot; |
 **limit** | **int**| The maximum amount of policies returned. | [optional]
 **offset** | **int**| The offset from where to start looking. | [optional]

### Return type

[**\keto\SDK\Model\OryAccessControlPolicyRoles**](../Model/OryAccessControlPolicyRoles.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

