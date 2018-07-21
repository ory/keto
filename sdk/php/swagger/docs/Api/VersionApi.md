# keto\SDK\VersionApi
Client for keto

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**getVersion**](VersionApi.md#getVersion) | **GET** /version | Get the version of Keto


# **getVersion**
> \keto\SDK\Model\Version getVersion()

Get the version of Keto

This endpoint returns the version as `{ \"version\": \"VERSION\" }`. The version is only correct with the prebuilt binary and not custom builds.

### Example
```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');

$api_instance = new keto\SDK\Api\VersionApi();

try {
    $result = $api_instance->getVersion();
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling VersionApi->getVersion: ', $e->getMessage(), PHP_EOL;
}
?>
```

### Parameters
This endpoint does not need any parameter.

### Return type

[**\keto\SDK\Model\Version**](../Model/Version.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

