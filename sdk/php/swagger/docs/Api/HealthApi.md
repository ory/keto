# keto\SDK\HealthApi
Client for keto

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**isInstanceAlive**](HealthApi.md#isInstanceAlive) | **GET** /health/alive | Check the Alive Status
[**isInstanceReady**](HealthApi.md#isInstanceReady) | **GET** /health/ready | Check the Readiness Status


# **isInstanceAlive**
> \keto\SDK\Model\HealthStatus isInstanceAlive()

Check the Alive Status

This endpoint returns a 200 status code when the HTTP server is up running. This status does currently not include checks whether the database connection is working. This endpoint does not require the `X-Forwarded-Proto` header when TLS termination is set.  Be aware that if you are running multiple nodes of ORY Keto, the health status will never refer to the cluster state, only to a single instance.

### Example
```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');

$api_instance = new keto\SDK\Api\HealthApi();

try {
    $result = $api_instance->isInstanceAlive();
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling HealthApi->isInstanceAlive: ', $e->getMessage(), PHP_EOL;
}
?>
```

### Parameters
This endpoint does not need any parameter.

### Return type

[**\keto\SDK\Model\HealthStatus**](../Model/HealthStatus.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **isInstanceReady**
> \keto\SDK\Model\HealthStatus isInstanceReady()

Check the Readiness Status

This endpoint returns a 200 status code when the HTTP server is up running and the environment dependencies (e.g. the database) are responsive as well.  This status does currently not include checks whether the database connection is working. This endpoint does not require the `X-Forwarded-Proto` header when TLS termination is set.  Be aware that if you are running multiple nodes of ORY Keto, the health status will never refer to the cluster state, only to a single instance.

### Example
```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');

$api_instance = new keto\SDK\Api\HealthApi();

try {
    $result = $api_instance->isInstanceReady();
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling HealthApi->isInstanceReady: ', $e->getMessage(), PHP_EOL;
}
?>
```

### Parameters
This endpoint does not need any parameter.

### Return type

[**\keto\SDK\Model\HealthStatus**](../Model/HealthStatus.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

