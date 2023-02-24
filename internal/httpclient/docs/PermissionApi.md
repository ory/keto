# \PermissionApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckPermission**](PermissionApi.md#CheckPermission) | **Get** /relation-tuples/check/openapi | Check a permission
[**CheckPermissionOrError**](PermissionApi.md#CheckPermissionOrError) | **Get** /relation-tuples/check | Check a permission
[**ExpandPermissions**](PermissionApi.md#ExpandPermissions) | **Get** /relation-tuples/expand | Expand a Relationship into permissions.
[**PostCheckPermission**](PermissionApi.md#PostCheckPermission) | **Post** /relation-tuples/check/openapi | Check a permission
[**PostCheckPermissionOrError**](PermissionApi.md#PostCheckPermissionOrError) | **Post** /relation-tuples/check | Check a permission



## CheckPermission

> CheckPermissionResult CheckPermission(ctx).Namespace(namespace).Object(object).Relation(relation).SubjectId(subjectId).SubjectSetNamespace(subjectSetNamespace).SubjectSetObject(subjectSetObject).SubjectSetRelation(subjectSetRelation).MaxDepth(maxDepth).Execute()

Check a permission



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
    namespace := "namespace_example" // string | Namespace of the Relationship (optional)
    object := "object_example" // string | Object of the Relationship (optional)
    relation := "relation_example" // string | Relation of the Relationship (optional)
    subjectId := "subjectId_example" // string | SubjectID of the Relationship (optional)
    subjectSetNamespace := "subjectSetNamespace_example" // string | Namespace of the Subject Set (optional)
    subjectSetObject := "subjectSetObject_example" // string | Object of the Subject Set (optional)
    subjectSetRelation := "subjectSetRelation_example" // string | Relation of the Subject Set (optional)
    maxDepth := int64(789) // int64 |  (optional)

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
 **namespace** | **string** | Namespace of the Relationship | 
 **object** | **string** | Object of the Relationship | 
 **relation** | **string** | Relation of the Relationship | 
 **subjectId** | **string** | SubjectID of the Relationship | 
 **subjectSetNamespace** | **string** | Namespace of the Subject Set | 
 **subjectSetObject** | **string** | Object of the Subject Set | 
 **subjectSetRelation** | **string** | Relation of the Subject Set | 
 **maxDepth** | **int64** |  | 

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

Check a permission



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
    namespace := "namespace_example" // string | Namespace of the Relationship (optional)
    object := "object_example" // string | Object of the Relationship (optional)
    relation := "relation_example" // string | Relation of the Relationship (optional)
    subjectId := "subjectId_example" // string | SubjectID of the Relationship (optional)
    subjectSetNamespace := "subjectSetNamespace_example" // string | Namespace of the Subject Set (optional)
    subjectSetObject := "subjectSetObject_example" // string | Object of the Subject Set (optional)
    subjectSetRelation := "subjectSetRelation_example" // string | Relation of the Subject Set (optional)
    maxDepth := int64(789) // int64 |  (optional)

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
 **namespace** | **string** | Namespace of the Relationship | 
 **object** | **string** | Object of the Relationship | 
 **relation** | **string** | Relation of the Relationship | 
 **subjectId** | **string** | SubjectID of the Relationship | 
 **subjectSetNamespace** | **string** | Namespace of the Subject Set | 
 **subjectSetObject** | **string** | Object of the Subject Set | 
 **subjectSetRelation** | **string** | Relation of the Subject Set | 
 **maxDepth** | **int64** |  | 

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

Expand a Relationship into permissions.



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
    namespace := "namespace_example" // string | Namespace of the Subject Set
    object := "object_example" // string | Object of the Subject Set
    relation := "relation_example" // string | Relation of the Subject Set
    maxDepth := int64(789) // int64 |  (optional)

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
 **namespace** | **string** | Namespace of the Subject Set | 
 **object** | **string** | Object of the Subject Set | 
 **relation** | **string** | Relation of the Subject Set | 
 **maxDepth** | **int64** |  | 

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

> CheckPermissionResult PostCheckPermission(ctx).MaxDepth(maxDepth).PostCheckPermissionBody(postCheckPermissionBody).Execute()

Check a permission



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
    maxDepth := int64(789) // int64 |  (optional)
    postCheckPermissionBody := *openapiclient.NewPostCheckPermissionBody() // PostCheckPermissionBody |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PermissionApi.PostCheckPermission(context.Background()).MaxDepth(maxDepth).PostCheckPermissionBody(postCheckPermissionBody).Execute()
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
 **maxDepth** | **int64** |  | 
 **postCheckPermissionBody** | [**PostCheckPermissionBody**](PostCheckPermissionBody.md) |  | 

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

> CheckPermissionResult PostCheckPermissionOrError(ctx).MaxDepth(maxDepth).PostCheckPermissionOrErrorBody(postCheckPermissionOrErrorBody).Execute()

Check a permission



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
    maxDepth := int64(789) // int64 | nolint:deadcode,unused (optional)
    postCheckPermissionOrErrorBody := *openapiclient.NewPostCheckPermissionOrErrorBody() // PostCheckPermissionOrErrorBody |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PermissionApi.PostCheckPermissionOrError(context.Background()).MaxDepth(maxDepth).PostCheckPermissionOrErrorBody(postCheckPermissionOrErrorBody).Execute()
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
 **maxDepth** | **int64** | nolint:deadcode,unused | 
 **postCheckPermissionOrErrorBody** | [**PostCheckPermissionOrErrorBody**](PostCheckPermissionOrErrorBody.md) |  | 

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

