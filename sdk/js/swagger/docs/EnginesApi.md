# SwaggerJsClient.EnginesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**addOryAccessControlPolicyRoleMembers**](EnginesApi.md#addOryAccessControlPolicyRoleMembers) | **PUT** /engines/acp/ory/{flavor}/roles/{id}/members | Add a member to an ORY Access Control Policy Role
[**deleteOryAccessControlPolicy**](EnginesApi.md#deleteOryAccessControlPolicy) | **DELETE** /engines/acp/ory/{flavor}/policies/{id} | 
[**deleteOryAccessControlPolicyRole**](EnginesApi.md#deleteOryAccessControlPolicyRole) | **DELETE** /engines/acp/ory/{flavor}/roles/{id} | Delete an ORY Access Control Policy Role
[**doOryAccessControlPoliciesAllow**](EnginesApi.md#doOryAccessControlPoliciesAllow) | **GET** /engines/acp/ory/{flavor}/allowed | Check if a request is allowed
[**getOryAccessControlPolicy**](EnginesApi.md#getOryAccessControlPolicy) | **GET** /engines/acp/ory/{flavor}/policies/{id} | 
[**getOryAccessControlPolicyRole**](EnginesApi.md#getOryAccessControlPolicyRole) | **GET** /engines/acp/ory/{flavor}/roles/{id} | Get an ORY Access Control Policy Role
[**listOryAccessControlPolicies**](EnginesApi.md#listOryAccessControlPolicies) | **GET** /engines/acp/ory/{flavor}/policies | 
[**listOryAccessControlPolicyRoles**](EnginesApi.md#listOryAccessControlPolicyRoles) | **GET** /engines/acp/ory/{flavor}/roles | List ORY Access Control Policy Roles
[**removeOryAccessControlPolicyRoleMembers**](EnginesApi.md#removeOryAccessControlPolicyRoleMembers) | **DELETE** /engines/acp/ory/{flavor}/roles/{id}/members | Remove a member from an ORY Access Control Policy Role
[**upsertOryAccessControlPolicy**](EnginesApi.md#upsertOryAccessControlPolicy) | **PUT** /engines/acp/ory/{flavor}/policies | 
[**upsertOryAccessControlPolicyRole**](EnginesApi.md#upsertOryAccessControlPolicyRole) | **PUT** /engines/acp/ory/{flavor}/roles | Upsert an ORY Access Control Policy Role


<a name="addOryAccessControlPolicyRoleMembers"></a>
# **addOryAccessControlPolicyRoleMembers**
> OryAccessControlPolicyRole addOryAccessControlPolicyRoleMembers(id, flavor, opts)

Add a member to an ORY Access Control Policy Role

Roles group several subjects into one. Rules can be assigned to ORY Access Control Policy (OACP) by using the Role ID as subject in the OACP.

### Example
```javascript
var SwaggerJsClient = require('swagger-js-client');

var apiInstance = new SwaggerJsClient.EnginesApi();

var id = "id_example"; // String | The ID of the ORY Access Control Policy Role.

var flavor = "flavor_example"; // String | The ORY Access Control Policy flavor. Can be \"regex\" and \"exact\".

var opts = { 
  'body': new SwaggerJsClient.AddOryAccessControlPolicyRoleMembersBody() // AddOryAccessControlPolicyRoleMembersBody | 
};

var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
};
apiInstance.addOryAccessControlPolicyRoleMembers(id, flavor, opts, callback);
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **String**| The ID of the ORY Access Control Policy Role. | 
 **flavor** | **String**| The ORY Access Control Policy flavor. Can be \&quot;regex\&quot; and \&quot;exact\&quot;. | 
 **body** | [**AddOryAccessControlPolicyRoleMembersBody**](AddOryAccessControlPolicyRoleMembersBody.md)|  | [optional] 

### Return type

[**OryAccessControlPolicyRole**](OryAccessControlPolicyRole.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="deleteOryAccessControlPolicy"></a>
# **deleteOryAccessControlPolicy**
> deleteOryAccessControlPolicy(flavor, id)



Delete an ORY Access Control Policy

### Example
```javascript
var SwaggerJsClient = require('swagger-js-client');

var apiInstance = new SwaggerJsClient.EnginesApi();

var flavor = "flavor_example"; // String | The ORY Access Control Policy flavor. Can be \"regex\" and \"exact\".

var id = "id_example"; // String | The ID of the ORY Access Control Policy Role.


var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully.');
  }
};
apiInstance.deleteOryAccessControlPolicy(flavor, id, callback);
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **flavor** | **String**| The ORY Access Control Policy flavor. Can be \&quot;regex\&quot; and \&quot;exact\&quot;. | 
 **id** | **String**| The ID of the ORY Access Control Policy Role. | 

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="deleteOryAccessControlPolicyRole"></a>
# **deleteOryAccessControlPolicyRole**
> deleteOryAccessControlPolicyRole(flavor, id)

Delete an ORY Access Control Policy Role

Roles group several subjects into one. Rules can be assigned to ORY Access Control Policy (OACP) by using the Role ID as subject in the OACP.

### Example
```javascript
var SwaggerJsClient = require('swagger-js-client');

var apiInstance = new SwaggerJsClient.EnginesApi();

var flavor = "flavor_example"; // String | The ORY Access Control Policy flavor. Can be \"regex\" and \"exact\".

var id = "id_example"; // String | The ID of the ORY Access Control Policy Role.


var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully.');
  }
};
apiInstance.deleteOryAccessControlPolicyRole(flavor, id, callback);
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **flavor** | **String**| The ORY Access Control Policy flavor. Can be \&quot;regex\&quot; and \&quot;exact\&quot;. | 
 **id** | **String**| The ID of the ORY Access Control Policy Role. | 

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="doOryAccessControlPoliciesAllow"></a>
# **doOryAccessControlPoliciesAllow**
> AuthorizationResult doOryAccessControlPoliciesAllow(flavor, opts)

Check if a request is allowed

Use this endpoint to check if a request is allowed or not.

### Example
```javascript
var SwaggerJsClient = require('swagger-js-client');

var apiInstance = new SwaggerJsClient.EnginesApi();

var flavor = "flavor_example"; // String | The ORY Access Control Policy flavor. Can be \"regex\" and \"exact\".

var opts = { 
  'body': new SwaggerJsClient.OryAccessControlPolicyAllowedInput() // OryAccessControlPolicyAllowedInput | 
};

var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
};
apiInstance.doOryAccessControlPoliciesAllow(flavor, opts, callback);
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **flavor** | **String**| The ORY Access Control Policy flavor. Can be \&quot;regex\&quot; and \&quot;exact\&quot;. | 
 **body** | [**OryAccessControlPolicyAllowedInput**](OryAccessControlPolicyAllowedInput.md)|  | [optional] 

### Return type

[**AuthorizationResult**](AuthorizationResult.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="getOryAccessControlPolicy"></a>
# **getOryAccessControlPolicy**
> OryAccessControlPolicy getOryAccessControlPolicy(flavor, id)



Get an ORY Access Control Policy

### Example
```javascript
var SwaggerJsClient = require('swagger-js-client');

var apiInstance = new SwaggerJsClient.EnginesApi();

var flavor = "flavor_example"; // String | The ORY Access Control Policy flavor. Can be \"regex\" and \"exact\".

var id = "id_example"; // String | The ID of the ORY Access Control Policy Role.


var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
};
apiInstance.getOryAccessControlPolicy(flavor, id, callback);
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **flavor** | **String**| The ORY Access Control Policy flavor. Can be \&quot;regex\&quot; and \&quot;exact\&quot;. | 
 **id** | **String**| The ID of the ORY Access Control Policy Role. | 

### Return type

[**OryAccessControlPolicy**](OryAccessControlPolicy.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="getOryAccessControlPolicyRole"></a>
# **getOryAccessControlPolicyRole**
> OryAccessControlPolicyRole getOryAccessControlPolicyRole(flavor, id)

Get an ORY Access Control Policy Role

Roles group several subjects into one. Rules can be assigned to ORY Access Control Policy (OACP) by using the Role ID as subject in the OACP.

### Example
```javascript
var SwaggerJsClient = require('swagger-js-client');

var apiInstance = new SwaggerJsClient.EnginesApi();

var flavor = "flavor_example"; // String | The ORY Access Control Policy flavor. Can be \"regex\" and \"exact\".

var id = "id_example"; // String | The ID of the ORY Access Control Policy Role.


var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
};
apiInstance.getOryAccessControlPolicyRole(flavor, id, callback);
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **flavor** | **String**| The ORY Access Control Policy flavor. Can be \&quot;regex\&quot; and \&quot;exact\&quot;. | 
 **id** | **String**| The ID of the ORY Access Control Policy Role. | 

### Return type

[**OryAccessControlPolicyRole**](OryAccessControlPolicyRole.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="listOryAccessControlPolicies"></a>
# **listOryAccessControlPolicies**
> OryAccessControlPolicies listOryAccessControlPolicies(flavor, opts)



List ORY Access Control Policies

### Example
```javascript
var SwaggerJsClient = require('swagger-js-client');

var apiInstance = new SwaggerJsClient.EnginesApi();

var flavor = "flavor_example"; // String | The ORY Access Control Policy flavor. Can be \"regex\" and \"exact\"

var opts = { 
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
apiInstance.listOryAccessControlPolicies(flavor, opts, callback);
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **flavor** | **String**| The ORY Access Control Policy flavor. Can be \&quot;regex\&quot; and \&quot;exact\&quot; | 
 **limit** | **Number**| The maximum amount of policies returned. | [optional] 
 **offset** | **Number**| The offset from where to start looking. | [optional] 

### Return type

[**OryAccessControlPolicies**](OryAccessControlPolicies.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="listOryAccessControlPolicyRoles"></a>
# **listOryAccessControlPolicyRoles**
> OryAccessControlPolicyRoles listOryAccessControlPolicyRoles(flavor, opts)

List ORY Access Control Policy Roles

Roles group several subjects into one. Rules can be assigned to ORY Access Control Policy (OACP) by using the Role ID as subject in the OACP.

### Example
```javascript
var SwaggerJsClient = require('swagger-js-client');

var apiInstance = new SwaggerJsClient.EnginesApi();

var flavor = "flavor_example"; // String | The ORY Access Control Policy flavor. Can be \"regex\" and \"exact\"

var opts = { 
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
apiInstance.listOryAccessControlPolicyRoles(flavor, opts, callback);
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **flavor** | **String**| The ORY Access Control Policy flavor. Can be \&quot;regex\&quot; and \&quot;exact\&quot; | 
 **limit** | **Number**| The maximum amount of policies returned. | [optional] 
 **offset** | **Number**| The offset from where to start looking. | [optional] 

### Return type

[**OryAccessControlPolicyRoles**](OryAccessControlPolicyRoles.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="removeOryAccessControlPolicyRoleMembers"></a>
# **removeOryAccessControlPolicyRoleMembers**
> removeOryAccessControlPolicyRoleMembers(id, flavor, opts)

Remove a member from an ORY Access Control Policy Role

Roles group several subjects into one. Rules can be assigned to ORY Access Control Policy (OACP) by using the Role ID as subject in the OACP.

### Example
```javascript
var SwaggerJsClient = require('swagger-js-client');

var apiInstance = new SwaggerJsClient.EnginesApi();

var id = "id_example"; // String | The ID of the ORY Access Control Policy Role.

var flavor = "flavor_example"; // String | The ORY Access Control Policy flavor. Can be \"regex\" and \"exact\".

var opts = { 
  'body': new SwaggerJsClient.RemoveOryAccessControlPolicyRoleMembersBody() // RemoveOryAccessControlPolicyRoleMembersBody | 
};

var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully.');
  }
};
apiInstance.removeOryAccessControlPolicyRoleMembers(id, flavor, opts, callback);
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **String**| The ID of the ORY Access Control Policy Role. | 
 **flavor** | **String**| The ORY Access Control Policy flavor. Can be \&quot;regex\&quot; and \&quot;exact\&quot;. | 
 **body** | [**RemoveOryAccessControlPolicyRoleMembersBody**](RemoveOryAccessControlPolicyRoleMembersBody.md)|  | [optional] 

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="upsertOryAccessControlPolicy"></a>
# **upsertOryAccessControlPolicy**
> upsertOryAccessControlPolicy(id, flavor, opts)



Upsert an ORY Access Control Policy

### Example
```javascript
var SwaggerJsClient = require('swagger-js-client');

var apiInstance = new SwaggerJsClient.EnginesApi();

var id = "id_example"; // String | The ID of the ORY Access Control Policy.

var flavor = "flavor_example"; // String | The ORY Access Control Policy flavor. Can be \"regex\" and \"exact\".

var opts = { 
  'body': new SwaggerJsClient.OryAccessControlPolicy() // OryAccessControlPolicy | 
};

var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully.');
  }
};
apiInstance.upsertOryAccessControlPolicy(id, flavor, opts, callback);
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **String**| The ID of the ORY Access Control Policy. | 
 **flavor** | **String**| The ORY Access Control Policy flavor. Can be \&quot;regex\&quot; and \&quot;exact\&quot;. | 
 **body** | [**OryAccessControlPolicy**](OryAccessControlPolicy.md)|  | [optional] 

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="upsertOryAccessControlPolicyRole"></a>
# **upsertOryAccessControlPolicyRole**
> OryAccessControlPolicyRole upsertOryAccessControlPolicyRole(id, flavor, opts)

Upsert an ORY Access Control Policy Role

Roles group several subjects into one. Rules can be assigned to ORY Access Control Policy (OACP) by using the Role ID as subject in the OACP.

### Example
```javascript
var SwaggerJsClient = require('swagger-js-client');

var apiInstance = new SwaggerJsClient.EnginesApi();

var id = "id_example"; // String | The ID of the ORY Access Control Policy Role.

var flavor = "flavor_example"; // String | The ORY Access Control Policy flavor. Can be \"regex\" and \"exact\".

var opts = { 
  'body': new SwaggerJsClient.OryAccessControlPolicyRole() // OryAccessControlPolicyRole | 
};

var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
};
apiInstance.upsertOryAccessControlPolicyRole(id, flavor, opts, callback);
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **String**| The ID of the ORY Access Control Policy Role. | 
 **flavor** | **String**| The ORY Access Control Policy flavor. Can be \&quot;regex\&quot; and \&quot;exact\&quot;. | 
 **body** | [**OryAccessControlPolicyRole**](OryAccessControlPolicyRole.md)|  | [optional] 

### Return type

[**OryAccessControlPolicyRole**](OryAccessControlPolicyRole.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

