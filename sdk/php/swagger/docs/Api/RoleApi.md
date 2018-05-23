# keto\SDK\RoleApi
Client for keto

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**addMembersToRole**](RoleApi.md#addMembersToRole) | **POST** /roles/{id}/members | Add members to a role
[**createRole**](RoleApi.md#createRole) | **POST** /roles | Create a role
[**deleteRole**](RoleApi.md#deleteRole) | **DELETE** /roles/{id} | Get a role by its ID
[**getRole**](RoleApi.md#getRole) | **GET** /roles/{id} | Get a role by its ID
[**listRoles**](RoleApi.md#listRoles) | **GET** /roles | List all roles
[**removeMembersFromRole**](RoleApi.md#removeMembersFromRole) | **DELETE** /roles/{id}/members | Remove members from a role
[**setRole**](RoleApi.md#setRole) | **PUT** /roles/{id} | A Role represents a group of users that share the same role and thus permissions. A role could be an administrator, a moderator, a regular user or some other sort of role.


# **addMembersToRole**
> addMembersToRole($id, $body)

Add members to a role

A Role represents a group of users that share the same role and thus permissions. A role could be an administrator, a moderator, a regular user or some other sort of role.  This endpoint allows you to add members (users, applications, ...) to a specific role. You have to know the role's ID.

### Example
```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');

$api_instance = new keto\SDK\Api\RoleApi();
$id = "id_example"; // string | The id of the role to modify.
$body = new \keto\SDK\Model\RoleMembers(); // \keto\SDK\Model\RoleMembers | 

try {
    $api_instance->addMembersToRole($id, $body);
} catch (Exception $e) {
    echo 'Exception when calling RoleApi->addMembersToRole: ', $e->getMessage(), PHP_EOL;
}
?>
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the role to modify. |
 **body** | [**\keto\SDK\Model\RoleMembers**](../Model/RoleMembers.md)|  | [optional]

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **createRole**
> \keto\SDK\Model\Role createRole($body)

Create a role

A Role represents a group of users that share the same role and thus permissions. A role could be an administrator, a moderator, a regular user or some other sort of role.  This endpoint allows you to create a new role. You may define members as well but you don't have to.

### Example
```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');

$api_instance = new keto\SDK\Api\RoleApi();
$body = new \keto\SDK\Model\Role(); // \keto\SDK\Model\Role | 

try {
    $result = $api_instance->createRole($body);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling RoleApi->createRole: ', $e->getMessage(), PHP_EOL;
}
?>
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**\keto\SDK\Model\Role**](../Model/Role.md)|  | [optional]

### Return type

[**\keto\SDK\Model\Role**](../Model/Role.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **deleteRole**
> deleteRole($id)

Get a role by its ID

A Role represents a group of users that share the same role and thus permissions. A role could be an administrator, a moderator, a regular user or some other sort of role.  This endpoint allows you to delete an existing role. You have to know the role's ID.

### Example
```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');

$api_instance = new keto\SDK\Api\RoleApi();
$id = "id_example"; // string | The id of the role to look up.

try {
    $api_instance->deleteRole($id);
} catch (Exception $e) {
    echo 'Exception when calling RoleApi->deleteRole: ', $e->getMessage(), PHP_EOL;
}
?>
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the role to look up. |

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **getRole**
> \keto\SDK\Model\Role getRole($id)

Get a role by its ID

A Role represents a group of users that share the same role and thus permissions. A role could be an administrator, a moderator, a regular user or some other sort of role.  This endpoint allows you to retrieve an existing role. You have to know the role's ID.

### Example
```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');

$api_instance = new keto\SDK\Api\RoleApi();
$id = "id_example"; // string | The id of the role to look up.

try {
    $result = $api_instance->getRole($id);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling RoleApi->getRole: ', $e->getMessage(), PHP_EOL;
}
?>
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the role to look up. |

### Return type

[**\keto\SDK\Model\Role**](../Model/Role.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **listRoles**
> \keto\SDK\Model\Role[] listRoles($member, $limit, $offset)

List all roles

A Role represents a group of users that share the same role and thus permissions. A role could be an administrator, a moderator, a regular user or some other sort of role.  This endpoint allows you to retrieve all roles that are stored in the system.

### Example
```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');

$api_instance = new keto\SDK\Api\RoleApi();
$member = "member_example"; // string | The id of the member to look up.
$limit = 789; // int | The maximum amount of policies returned.
$offset = 789; // int | The offset from where to start looking.

try {
    $result = $api_instance->listRoles($member, $limit, $offset);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling RoleApi->listRoles: ', $e->getMessage(), PHP_EOL;
}
?>
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **member** | **string**| The id of the member to look up. | [optional]
 **limit** | **int**| The maximum amount of policies returned. | [optional]
 **offset** | **int**| The offset from where to start looking. | [optional]

### Return type

[**\keto\SDK\Model\Role[]**](../Model/Role.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **removeMembersFromRole**
> removeMembersFromRole($id, $body)

Remove members from a role

A Role represents a group of users that share the same role and thus permissions. A role could be an administrator, a moderator, a regular user or some other sort of role.  This endpoint allows you to remove members (users, applications, ...) from a specific role. You have to know the role's ID.

### Example
```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');

$api_instance = new keto\SDK\Api\RoleApi();
$id = "id_example"; // string | The id of the role to modify.
$body = new \keto\SDK\Model\RoleMembers(); // \keto\SDK\Model\RoleMembers | 

try {
    $api_instance->removeMembersFromRole($id, $body);
} catch (Exception $e) {
    echo 'Exception when calling RoleApi->removeMembersFromRole: ', $e->getMessage(), PHP_EOL;
}
?>
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the role to modify. |
 **body** | [**\keto\SDK\Model\RoleMembers**](../Model/RoleMembers.md)|  | [optional]

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **setRole**
> setRole()

A Role represents a group of users that share the same role and thus permissions. A role could be an administrator, a moderator, a regular user or some other sort of role.

This endpoint allows you to overwrite a role. You have to know the role's ID.

### Example
```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');

$api_instance = new keto\SDK\Api\RoleApi();

try {
    $api_instance->setRole();
} catch (Exception $e) {
    echo 'Exception when calling RoleApi->setRole: ', $e->getMessage(), PHP_EOL;
}
?>
```

### Parameters
This endpoint does not need any parameter.

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

