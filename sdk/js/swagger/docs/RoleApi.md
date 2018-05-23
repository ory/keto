# SwaggerJsClient.RoleApi

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


<a name="addMembersToRole"></a>
# **addMembersToRole**
> addMembersToRole(id, opts)

Add members to a role

A Role represents a group of users that share the same role and thus permissions. A role could be an administrator, a moderator, a regular user or some other sort of role.  This endpoint allows you to add members (users, applications, ...) to a specific role. You have to know the role&#39;s ID.

### Example
```javascript
var SwaggerJsClient = require('swagger-js-client');

var apiInstance = new SwaggerJsClient.RoleApi();

var id = "id_example"; // String | The id of the role to modify.

var opts = { 
  'body': new SwaggerJsClient.RoleMembers() // RoleMembers | 
};

var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully.');
  }
};
apiInstance.addMembersToRole(id, opts, callback);
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **String**| The id of the role to modify. | 
 **body** | [**RoleMembers**](RoleMembers.md)|  | [optional] 

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="createRole"></a>
# **createRole**
> Role createRole(opts)

Create a role

A Role represents a group of users that share the same role and thus permissions. A role could be an administrator, a moderator, a regular user or some other sort of role.  This endpoint allows you to create a new role. You may define members as well but you don&#39;t have to.

### Example
```javascript
var SwaggerJsClient = require('swagger-js-client');

var apiInstance = new SwaggerJsClient.RoleApi();

var opts = { 
  'body': new SwaggerJsClient.Role() // Role | 
};

var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
};
apiInstance.createRole(opts, callback);
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Role**](Role.md)|  | [optional] 

### Return type

[**Role**](Role.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="deleteRole"></a>
# **deleteRole**
> deleteRole(id)

Get a role by its ID

A Role represents a group of users that share the same role and thus permissions. A role could be an administrator, a moderator, a regular user or some other sort of role.  This endpoint allows you to delete an existing role. You have to know the role&#39;s ID.

### Example
```javascript
var SwaggerJsClient = require('swagger-js-client');

var apiInstance = new SwaggerJsClient.RoleApi();

var id = "id_example"; // String | The id of the role to look up.


var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully.');
  }
};
apiInstance.deleteRole(id, callback);
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **String**| The id of the role to look up. | 

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="getRole"></a>
# **getRole**
> Role getRole(id)

Get a role by its ID

A Role represents a group of users that share the same role and thus permissions. A role could be an administrator, a moderator, a regular user or some other sort of role.  This endpoint allows you to retrieve an existing role. You have to know the role&#39;s ID.

### Example
```javascript
var SwaggerJsClient = require('swagger-js-client');

var apiInstance = new SwaggerJsClient.RoleApi();

var id = "id_example"; // String | The id of the role to look up.


var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
};
apiInstance.getRole(id, callback);
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **String**| The id of the role to look up. | 

### Return type

[**Role**](Role.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="listRoles"></a>
# **listRoles**
> [Role] listRoles(opts)

List all roles

A Role represents a group of users that share the same role and thus permissions. A role could be an administrator, a moderator, a regular user or some other sort of role.  This endpoint allows you to retrieve all roles that are stored in the system.

### Example
```javascript
var SwaggerJsClient = require('swagger-js-client');

var apiInstance = new SwaggerJsClient.RoleApi();

var opts = { 
  'member': "member_example", // String | The id of the member to look up.
  'limit': 789, // Number | The maximum amount of policies returned.
  'offset': 789 // Number | The offset from where to start looking.
};

var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
};
apiInstance.listRoles(opts, callback);
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **member** | **String**| The id of the member to look up. | [optional] 
 **limit** | **Number**| The maximum amount of policies returned. | [optional] 
 **offset** | **Number**| The offset from where to start looking. | [optional] 

### Return type

[**[Role]**](Role.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="removeMembersFromRole"></a>
# **removeMembersFromRole**
> removeMembersFromRole(id, opts)

Remove members from a role

A Role represents a group of users that share the same role and thus permissions. A role could be an administrator, a moderator, a regular user or some other sort of role.  This endpoint allows you to remove members (users, applications, ...) from a specific role. You have to know the role&#39;s ID.

### Example
```javascript
var SwaggerJsClient = require('swagger-js-client');

var apiInstance = new SwaggerJsClient.RoleApi();

var id = "id_example"; // String | The id of the role to modify.

var opts = { 
  'body': new SwaggerJsClient.RoleMembers() // RoleMembers | 
};

var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully.');
  }
};
apiInstance.removeMembersFromRole(id, opts, callback);
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **String**| The id of the role to modify. | 
 **body** | [**RoleMembers**](RoleMembers.md)|  | [optional] 

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="setRole"></a>
# **setRole**
> setRole()

A Role represents a group of users that share the same role and thus permissions. A role could be an administrator, a moderator, a regular user or some other sort of role.

This endpoint allows you to overwrite a role. You have to know the role&#39;s ID.

### Example
```javascript
var SwaggerJsClient = require('swagger-js-client');

var apiInstance = new SwaggerJsClient.RoleApi();

var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully.');
  }
};
apiInstance.setRole(callback);
```

### Parameters
This endpoint does not need any parameter.

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

