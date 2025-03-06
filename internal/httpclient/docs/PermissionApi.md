# \PermissionApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchCheckPermission**](PermissionApi.md#BatchCheckPermission) | **Post** /relation-tuples/batch/check | Performs an authorization check for a batch of tuples.
[**CheckPermission**](PermissionApi.md#CheckPermission) | **Get** /relation-tuples/check/openapi | Performs an authorization check.
[**CheckPermissionOrError**](PermissionApi.md#CheckPermissionOrError) | **Get** /relation-tuples/check | Performs an authorization check.
[**ExpandPermissions**](PermissionApi.md#ExpandPermissions) | **Get** /relation-tuples/expand | Expands the subject set into a tree of subjects.
[**PostCheckPermission**](PermissionApi.md#PostCheckPermission) | **Post** /relation-tuples/check/openapi | Performs an authorization check.
[**PostCheckPermissionOrError**](PermissionApi.md#PostCheckPermissionOrError) | **Post** /relation-tuples/check | Performs an authorization check.



## BatchCheckPermission

> BatchCheckPermissionResult BatchCheckPermission(ctx).BatchCheckPermissionBody(batchCheckPermissionBody).MaxDepth(maxDepth).Execute()

Performs an authorization check for a batch of tuples.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    batchCheckPermissionBody := *openapiclient.NewBatchCheckPermissionBody() // BatchCheckPermissionBody | Batch Check Permission Body.
    maxDepth := int32(56) // int32 | The maximum depth to search for a relation.  If the value is less than 1 or greater than the global max-depth then the global max-depth will be used instead. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PermissionApi.BatchCheckPermission(context.Background()).BatchCheckPermissionBody(batchCheckPermissionBody).MaxDepth(maxDepth).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PermissionApi.BatchCheckPermission``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BatchCheckPermission`: BatchCheckPermissionResult
    fmt.Fprintf(os.Stdout, "Response from `PermissionApi.BatchCheckPermission`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchCheckPermissionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchCheckPermissionBody** | [**BatchCheckPermissionBody**](BatchCheckPermissionBody.md) | Batch Check Permission Body. | 
 **maxDepth** | **int32** | The maximum depth to search for a relation.  If the value is less than 1 or greater than the global max-depth then the global max-depth will be used instead. | 

### Return type

[**BatchCheckPermissionResult**](BatchCheckPermissionResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckPermission

> CheckPermissionResult CheckPermission(ctx).Namespace(namespace).Object(object).Relation(relation).SubjectId(subjectId).SubjectSetNamespace(subjectSetNamespace).SubjectSetObject(subjectSetObject).SubjectSetRelation(subjectSetRelation).MaxDepth(maxDepth).Execute()

Performs an authorization check.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    namespace := "namespace_example" // string | The namespace to evaluate the check.  Note: If you use the expand-API and the check evaluates a RelationTuple specifying a SubjectSet as subject or due to a rewrite rule in a namespace config this check request may involve other namespaces automatically. (optional)
    object := "object_example" // string | The related object in this check. (optional)
    relation := "relation_example" // string | The relation between the Object and the Subject. (optional)
    subjectId := "subjectId_example" // string | A concrete id of the subject. (optional)
    subjectSetNamespace := "subjectSetNamespace_example" // string | The namespace of the object and relation referenced in this subject set. (optional)
    subjectSetObject := "subjectSetObject_example" // string | The object related by this subject set. (optional)
    subjectSetRelation := "subjectSetRelation_example" // string | The relation between the object and the subjects. (optional)
    maxDepth := int32(56) // int32 | The maximum depth to search for a relation.  If the value is less than 1 or greater than the global max-depth then the global max-depth will be used instead. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PermissionApi.CheckPermission(context.Background()).Namespace(namespace).Object(object).Relation(relation).SubjectId(subjectId).SubjectSetNamespace(subjectSetNamespace).SubjectSetObject(subjectSetObject).SubjectSetRelation(subjectSetRelation).MaxDepth(maxDepth).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PermissionApi.CheckPermission``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckPermission`: CheckPermissionResult
    fmt.Fprintf(os.Stdout, "Response from `PermissionApi.CheckPermission`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckPermissionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **namespace** | **string** | The namespace to evaluate the check.  Note: If you use the expand-API and the check evaluates a RelationTuple specifying a SubjectSet as subject or due to a rewrite rule in a namespace config this check request may involve other namespaces automatically. | 
 **object** | **string** | The related object in this check. | 
 **relation** | **string** | The relation between the Object and the Subject. | 
 **subjectId** | **string** | A concrete id of the subject. | 
 **subjectSetNamespace** | **string** | The namespace of the object and relation referenced in this subject set. | 
 **subjectSetObject** | **string** | The object related by this subject set. | 
 **subjectSetRelation** | **string** | The relation between the object and the subjects. | 
 **maxDepth** | **int32** | The maximum depth to search for a relation.  If the value is less than 1 or greater than the global max-depth then the global max-depth will be used instead. | 

### Return type

[**CheckPermissionResult**](CheckPermissionResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckPermissionOrError

> CheckPermissionResult CheckPermissionOrError(ctx).Namespace(namespace).Object(object).Relation(relation).SubjectId(subjectId).SubjectSetNamespace(subjectSetNamespace).SubjectSetObject(subjectSetObject).SubjectSetRelation(subjectSetRelation).MaxDepth(maxDepth).Execute()

Performs an authorization check.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    namespace := "namespace_example" // string | The namespace to evaluate the check.  Note: If you use the expand-API and the check evaluates a RelationTuple specifying a SubjectSet as subject or due to a rewrite rule in a namespace config this check request may involve other namespaces automatically. (optional)
    object := "object_example" // string | The related object in this check. (optional)
    relation := "relation_example" // string | The relation between the Object and the Subject. (optional)
    subjectId := "subjectId_example" // string | A concrete id of the subject. (optional)
    subjectSetNamespace := "subjectSetNamespace_example" // string | The namespace of the object and relation referenced in this subject set. (optional)
    subjectSetObject := "subjectSetObject_example" // string | The object related by this subject set. (optional)
    subjectSetRelation := "subjectSetRelation_example" // string | The relation between the object and the subjects. (optional)
    maxDepth := int32(56) // int32 | The maximum depth to search for a relation.  If the value is less than 1 or greater than the global max-depth then the global max-depth will be used instead. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PermissionApi.CheckPermissionOrError(context.Background()).Namespace(namespace).Object(object).Relation(relation).SubjectId(subjectId).SubjectSetNamespace(subjectSetNamespace).SubjectSetObject(subjectSetObject).SubjectSetRelation(subjectSetRelation).MaxDepth(maxDepth).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PermissionApi.CheckPermissionOrError``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckPermissionOrError`: CheckPermissionResult
    fmt.Fprintf(os.Stdout, "Response from `PermissionApi.CheckPermissionOrError`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckPermissionOrErrorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **namespace** | **string** | The namespace to evaluate the check.  Note: If you use the expand-API and the check evaluates a RelationTuple specifying a SubjectSet as subject or due to a rewrite rule in a namespace config this check request may involve other namespaces automatically. | 
 **object** | **string** | The related object in this check. | 
 **relation** | **string** | The relation between the Object and the Subject. | 
 **subjectId** | **string** | A concrete id of the subject. | 
 **subjectSetNamespace** | **string** | The namespace of the object and relation referenced in this subject set. | 
 **subjectSetObject** | **string** | The object related by this subject set. | 
 **subjectSetRelation** | **string** | The relation between the object and the subjects. | 
 **maxDepth** | **int32** | The maximum depth to search for a relation.  If the value is less than 1 or greater than the global max-depth then the global max-depth will be used instead. | 

### Return type

[**CheckPermissionResult**](CheckPermissionResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExpandPermissions

> ExpandedPermissionTree ExpandPermissions(ctx).Namespace(namespace).Object(object).Relation(relation).MaxDepth(maxDepth).Execute()

Expands the subject set into a tree of subjects.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    namespace := "namespace_example" // string | The namespace of the object and relation referenced in this subject set.
    object := "object_example" // string | The object related by this subject set.
    relation := "relation_example" // string | The relation between the object and the subjects.
    maxDepth := int32(56) // int32 | The maximum depth of tree to build.  If the value is less than 1 or greater than the global max-depth then the global max-depth will be used instead.  It is important to set this parameter to a meaningful value. Ponder how deep you really want to display this. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PermissionApi.ExpandPermissions(context.Background()).Namespace(namespace).Object(object).Relation(relation).MaxDepth(maxDepth).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PermissionApi.ExpandPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExpandPermissions`: ExpandedPermissionTree
    fmt.Fprintf(os.Stdout, "Response from `PermissionApi.ExpandPermissions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExpandPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **namespace** | **string** | The namespace of the object and relation referenced in this subject set. | 
 **object** | **string** | The object related by this subject set. | 
 **relation** | **string** | The relation between the object and the subjects. | 
 **maxDepth** | **int32** | The maximum depth of tree to build.  If the value is less than 1 or greater than the global max-depth then the global max-depth will be used instead.  It is important to set this parameter to a meaningful value. Ponder how deep you really want to display this. | 

### Return type

[**ExpandedPermissionTree**](ExpandedPermissionTree.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCheckPermission

> CheckPermissionResult PostCheckPermission(ctx).PostCheckPermissionBody(postCheckPermissionBody).Execute()

Performs an authorization check.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    postCheckPermissionBody := *openapiclient.NewPostCheckPermissionBody() // PostCheckPermissionBody | The request for a CheckService.Check RPC. Checks whether a specific subject is related to an object.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PermissionApi.PostCheckPermission(context.Background()).PostCheckPermissionBody(postCheckPermissionBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PermissionApi.PostCheckPermission``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCheckPermission`: CheckPermissionResult
    fmt.Fprintf(os.Stdout, "Response from `PermissionApi.PostCheckPermission`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCheckPermissionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postCheckPermissionBody** | [**PostCheckPermissionBody**](PostCheckPermissionBody.md) | The request for a CheckService.Check RPC. Checks whether a specific subject is related to an object. | 

### Return type

[**CheckPermissionResult**](CheckPermissionResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCheckPermissionOrError

> CheckPermissionResult PostCheckPermissionOrError(ctx).PostCheckPermissionBody(postCheckPermissionBody).Execute()

Performs an authorization check.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    postCheckPermissionBody := *openapiclient.NewPostCheckPermissionBody() // PostCheckPermissionBody | The request for a CheckService.Check RPC. Checks whether a specific subject is related to an object.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PermissionApi.PostCheckPermissionOrError(context.Background()).PostCheckPermissionBody(postCheckPermissionBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PermissionApi.PostCheckPermissionOrError``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCheckPermissionOrError`: CheckPermissionResult
    fmt.Fprintf(os.Stdout, "Response from `PermissionApi.PostCheckPermissionOrError`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCheckPermissionOrErrorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postCheckPermissionBody** | [**PostCheckPermissionBody**](PostCheckPermissionBody.md) | The request for a CheckService.Check RPC. Checks whether a specific subject is related to an object. | 

### Return type

[**CheckPermissionResult**](CheckPermissionResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

