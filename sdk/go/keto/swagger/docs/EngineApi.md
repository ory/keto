# \EngineApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListOryAccessControlPolicyRoles**](EngineApi.md#ListOryAccessControlPolicyRoles) | **Get** /engines/ory/{flavor}/roles | List ORY Access Control Policy Roles


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

