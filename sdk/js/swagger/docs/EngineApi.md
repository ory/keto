# SwaggerJsClient.EngineApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**listOryAccessControlPolicyRoles**](EngineApi.md#listOryAccessControlPolicyRoles) | **GET** /engines/ory/{flavor}/roles | List ORY Access Control Policy Roles


<a name="listOryAccessControlPolicyRoles"></a>
# **listOryAccessControlPolicyRoles**
> OryAccessControlPolicyRoles listOryAccessControlPolicyRoles(flavor, opts)

List ORY Access Control Policy Roles

Roles group several subjects into one. Rules can be assigned to ORY Access Control Policy (OACP) by using the Role ID as subject in the OACP.

### Example
```javascript
var SwaggerJsClient = require('swagger-js-client');

var apiInstance = new SwaggerJsClient.EngineApi();

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

