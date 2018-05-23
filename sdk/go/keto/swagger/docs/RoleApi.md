# \RoleApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddMembersToRole**](RoleApi.md#AddMembersToRole) | **Post** /roles/{id}/members | Add members to a role
[**CreateRole**](RoleApi.md#CreateRole) | **Post** /roles | Create a role
[**DeleteRole**](RoleApi.md#DeleteRole) | **Delete** /roles/{id} | Get a role by its ID
[**GetRole**](RoleApi.md#GetRole) | **Get** /roles/{id} | Get a role by its ID
[**ListRoles**](RoleApi.md#ListRoles) | **Get** /roles | List all roles
[**RemoveMembersFromRole**](RoleApi.md#RemoveMembersFromRole) | **Delete** /roles/{id}/members | Remove members from a role
[**SetRole**](RoleApi.md#SetRole) | **Put** /roles/{id} | A Role represents a group of users that share the same role and thus permissions. A role could be an administrator, a moderator, a regular user or some other sort of role.


# **AddMembersToRole**
> AddMembersToRole($id, $body)

Add members to a role

A Role represents a group of users that share the same role and thus permissions. A role could be an administrator, a moderator, a regular user or some other sort of role.  This endpoint allows you to add members (users, applications, ...) to a specific role. You have to know the role's ID.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the role to modify. | 
 **body** | [**RoleMembers**](RoleMembers.md)|  | [optional] 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateRole**
> Role CreateRole($body)

Create a role

A Role represents a group of users that share the same role and thus permissions. A role could be an administrator, a moderator, a regular user or some other sort of role.  This endpoint allows you to create a new role. You may define members as well but you don't have to.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Role**](Role.md)|  | [optional] 

### Return type

[**Role**](role.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRole**
> DeleteRole($id)

Get a role by its ID

A Role represents a group of users that share the same role and thus permissions. A role could be an administrator, a moderator, a regular user or some other sort of role.  This endpoint allows you to delete an existing role. You have to know the role's ID.


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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRole**
> Role GetRole($id)

Get a role by its ID

A Role represents a group of users that share the same role and thus permissions. A role could be an administrator, a moderator, a regular user or some other sort of role.  This endpoint allows you to retrieve an existing role. You have to know the role's ID.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the role to look up. | 

### Return type

[**Role**](role.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListRoles**
> []Role ListRoles($member, $limit, $offset)

List all roles

A Role represents a group of users that share the same role and thus permissions. A role could be an administrator, a moderator, a regular user or some other sort of role.  This endpoint allows you to retrieve all roles that are stored in the system.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **member** | **string**| The id of the member to look up. | [optional] 
 **limit** | **int64**| The maximum amount of policies returned. | [optional] 
 **offset** | **int64**| The offset from where to start looking. | [optional] 

### Return type

[**[]Role**](role.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveMembersFromRole**
> RemoveMembersFromRole($id, $body)

Remove members from a role

A Role represents a group of users that share the same role and thus permissions. A role could be an administrator, a moderator, a regular user or some other sort of role.  This endpoint allows you to remove members (users, applications, ...) from a specific role. You have to know the role's ID.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The id of the role to modify. | 
 **body** | [**RoleMembers**](RoleMembers.md)|  | [optional] 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetRole**
> SetRole()

A Role represents a group of users that share the same role and thus permissions. A role could be an administrator, a moderator, a regular user or some other sort of role.

This endpoint allows you to overwrite a role. You have to know the role's ID.


### Parameters
This endpoint does not need any parameter.

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

