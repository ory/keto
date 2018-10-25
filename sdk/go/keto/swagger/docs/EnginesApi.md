# \EnginesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddOryAccessControlPolicyRoleMembers**](EnginesApi.md#AddOryAccessControlPolicyRoleMembers) | **Put** /engines/acp/ory/{flavor}/roles/{id}/members | Add a member to an ORY Access Control Policy Role
[**DeleteOryAccessControlPolicy**](EnginesApi.md#DeleteOryAccessControlPolicy) | **Delete** /engines/acp/ory/{flavor}/policies/{id} | 
[**DeleteOryAccessControlPolicyRole**](EnginesApi.md#DeleteOryAccessControlPolicyRole) | **Delete** /engines/acp/ory/{flavor}/roles/{id} | Delete an ORY Access Control Policy Role
[**DoOryAccessControlPoliciesAllow**](EnginesApi.md#DoOryAccessControlPoliciesAllow) | **Get** /engines/acp/ory/{flavor}/allowed | Check if a request is allowed
[**GetOryAccessControlPolicy**](EnginesApi.md#GetOryAccessControlPolicy) | **Get** /engines/acp/ory/{flavor}/policies/{id} | 
[**GetOryAccessControlPolicyRole**](EnginesApi.md#GetOryAccessControlPolicyRole) | **Get** /engines/acp/ory/{flavor}/roles/{id} | Get an ORY Access Control Policy Role
[**ListOryAccessControlPolicies**](EnginesApi.md#ListOryAccessControlPolicies) | **Get** /engines/acp/ory/{flavor}/policies | 
[**ListOryAccessControlPolicyRoles**](EnginesApi.md#ListOryAccessControlPolicyRoles) | **Get** /engines/acp/ory/{flavor}/roles | List ORY Access Control Policy Roles
[**RemoveOryAccessControlPolicyRoleMembers**](EnginesApi.md#RemoveOryAccessControlPolicyRoleMembers) | **Delete** /engines/acp/ory/{flavor}/roles/{id}/members | Remove a member from an ORY Access Control Policy Role
[**UpsertOryAccessControlPolicy**](EnginesApi.md#UpsertOryAccessControlPolicy) | **Put** /engines/acp/ory/{flavor}/policies | 
[**UpsertOryAccessControlPolicyRole**](EnginesApi.md#UpsertOryAccessControlPolicyRole) | **Put** /engines/acp/ory/{flavor}/roles | Upsert an ORY Access Control Policy Role


# **AddOryAccessControlPolicyRoleMembers**
> OryAccessControlPolicyRole AddOryAccessControlPolicyRoleMembers($id, $flavor, $body)

Add a member to an ORY Access Control Policy Role

Roles group several subjects into one. Rules can be assigned to ORY Access Control Policy (OACP) by using the Role ID as subject in the OACP.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The ID of the ORY Access Control Policy Role. | 
 **flavor** | **string**| The ORY Access Control Policy flavor. Can be \&quot;regex\&quot; and \&quot;exact\&quot;. | 
 **body** | [**AddOryAccessControlPolicyRoleMembersBody**](AddOryAccessControlPolicyRoleMembersBody.md)|  | [optional] 

### Return type

[**OryAccessControlPolicyRole**](oryAccessControlPolicyRole.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOryAccessControlPolicy**
> DeleteOryAccessControlPolicy($flavor, $id)



Delete an ORY Access Control Policy


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **flavor** | **string**| The ORY Access Control Policy flavor. Can be \&quot;regex\&quot; and \&quot;exact\&quot;. | 
 **id** | **string**| The ID of the ORY Access Control Policy Role. | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOryAccessControlPolicyRole**
> DeleteOryAccessControlPolicyRole($flavor, $id)

Delete an ORY Access Control Policy Role

Roles group several subjects into one. Rules can be assigned to ORY Access Control Policy (OACP) by using the Role ID as subject in the OACP.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **flavor** | **string**| The ORY Access Control Policy flavor. Can be \&quot;regex\&quot; and \&quot;exact\&quot;. | 
 **id** | **string**| The ID of the ORY Access Control Policy Role. | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DoOryAccessControlPoliciesAllow**
> AuthorizationResult DoOryAccessControlPoliciesAllow($flavor, $body)

Check if a request is allowed

Use this endpoint to check if a request is allowed or not.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **flavor** | **string**| The ORY Access Control Policy flavor. Can be \&quot;regex\&quot; and \&quot;exact\&quot;. | 
 **body** | [**OryAccessControlPolicyAllowedInput**](OryAccessControlPolicyAllowedInput.md)|  | [optional] 

### Return type

[**AuthorizationResult**](authorizationResult.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOryAccessControlPolicy**
> OryAccessControlPolicy GetOryAccessControlPolicy($flavor, $id)



Get an ORY Access Control Policy


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **flavor** | **string**| The ORY Access Control Policy flavor. Can be \&quot;regex\&quot; and \&quot;exact\&quot;. | 
 **id** | **string**| The ID of the ORY Access Control Policy Role. | 

### Return type

[**OryAccessControlPolicy**](oryAccessControlPolicy.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOryAccessControlPolicyRole**
> OryAccessControlPolicyRole GetOryAccessControlPolicyRole($flavor, $id)

Get an ORY Access Control Policy Role

Roles group several subjects into one. Rules can be assigned to ORY Access Control Policy (OACP) by using the Role ID as subject in the OACP.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **flavor** | **string**| The ORY Access Control Policy flavor. Can be \&quot;regex\&quot; and \&quot;exact\&quot;. | 
 **id** | **string**| The ID of the ORY Access Control Policy Role. | 

### Return type

[**OryAccessControlPolicyRole**](oryAccessControlPolicyRole.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOryAccessControlPolicies**
> OryAccessControlPolicies ListOryAccessControlPolicies($flavor, $limit, $offset)



List ORY Access Control Policies


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **flavor** | **string**| The ORY Access Control Policy flavor. Can be \&quot;regex\&quot; and \&quot;exact\&quot; | 
 **limit** | **int64**| The maximum amount of policies returned. | [optional] 
 **offset** | **int64**| The offset from where to start looking. | [optional] 

### Return type

[**OryAccessControlPolicies**](oryAccessControlPolicies.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOryAccessControlPolicyRoles**
> OryAccessControlPolicyRoles ListOryAccessControlPolicyRoles($flavor, $limit, $offset)

List ORY Access Control Policy Roles

Roles group several subjects into one. Rules can be assigned to ORY Access Control Policy (OACP) by using the Role ID as subject in the OACP.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **flavor** | **string**| The ORY Access Control Policy flavor. Can be \&quot;regex\&quot; and \&quot;exact\&quot; | 
 **limit** | **int64**| The maximum amount of policies returned. | [optional] 
 **offset** | **int64**| The offset from where to start looking. | [optional] 

### Return type

[**OryAccessControlPolicyRoles**](oryAccessControlPolicyRoles.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveOryAccessControlPolicyRoleMembers**
> RemoveOryAccessControlPolicyRoleMembers($id, $flavor, $body)

Remove a member from an ORY Access Control Policy Role

Roles group several subjects into one. Rules can be assigned to ORY Access Control Policy (OACP) by using the Role ID as subject in the OACP.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The ID of the ORY Access Control Policy Role. | 
 **flavor** | **string**| The ORY Access Control Policy flavor. Can be \&quot;regex\&quot; and \&quot;exact\&quot;. | 
 **body** | [**RemoveOryAccessControlPolicyRoleMembersBody**](RemoveOryAccessControlPolicyRoleMembersBody.md)|  | [optional] 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpsertOryAccessControlPolicy**
> UpsertOryAccessControlPolicy($id, $flavor, $body)



Upsert an ORY Access Control Policy


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The ID of the ORY Access Control Policy. | 
 **flavor** | **string**| The ORY Access Control Policy flavor. Can be \&quot;regex\&quot; and \&quot;exact\&quot;. | 
 **body** | [**OryAccessControlPolicy**](OryAccessControlPolicy.md)|  | [optional] 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpsertOryAccessControlPolicyRole**
> OryAccessControlPolicyRole UpsertOryAccessControlPolicyRole($id, $flavor, $body)

Upsert an ORY Access Control Policy Role

Roles group several subjects into one. Rules can be assigned to ORY Access Control Policy (OACP) by using the Role ID as subject in the OACP.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| The ID of the ORY Access Control Policy Role. | 
 **flavor** | **string**| The ORY Access Control Policy flavor. Can be \&quot;regex\&quot; and \&quot;exact\&quot;. | 
 **body** | [**OryAccessControlPolicyRole**](OryAccessControlPolicyRole.md)|  | [optional] 

### Return type

[**OryAccessControlPolicyRole**](oryAccessControlPolicyRole.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

