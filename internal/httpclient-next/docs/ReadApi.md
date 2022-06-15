# \ReadApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCheck**](ReadApi.md#GetCheck) | **Get** /relation-tuples/check/openapi | Check a relation tuple
[**GetCheckMirrorStatus**](ReadApi.md#GetCheckMirrorStatus) | **Get** /relation-tuples/check | Check a relation tuple
[**GetExpand**](ReadApi.md#GetExpand) | **Get** /relation-tuples/expand | Expand a Relation Tuple
[**GetRelationTuples**](ReadApi.md#GetRelationTuples) | **Get** /relation-tuples | Query relation tuples
[**PostCheck**](ReadApi.md#PostCheck) | **Post** /relation-tuples/check/openapi | Check a relation tuple
[**PostCheckMirrorStatus**](ReadApi.md#PostCheckMirrorStatus) | **Post** /relation-tuples/check | Check a relation tuple



## GetCheck

> GetCheckResponse GetCheck(ctx).Namespace(namespace).Object(object).Relation(relation).SubjectId(subjectId).SubjectSetNamespace(subjectSetNamespace).SubjectSetObject(subjectSetObject).SubjectSetRelation(subjectSetRelation).MaxDepth(maxDepth).Execute()

Check a relation tuple



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
    namespace := "namespace_example" // string | Namespace of the Relation Tuple (optional)
    object := "object_example" // string | Object of the Relation Tuple (optional)
    relation := "relation_example" // string | Relation of the Relation Tuple (optional)
    subjectId := "subjectId_example" // string | SubjectID of the Relation Tuple (optional)
    subjectSetNamespace := "subjectSetNamespace_example" // string | Namespace of the Subject Set (optional)
    subjectSetObject := "subjectSetObject_example" // string | Object of the Subject Set (optional)
    subjectSetRelation := "subjectSetRelation_example" // string | Relation of the Subject Set (optional)
    maxDepth := int64(789) // int64 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReadApi.GetCheck(context.Background()).Namespace(namespace).Object(object).Relation(relation).SubjectId(subjectId).SubjectSetNamespace(subjectSetNamespace).SubjectSetObject(subjectSetObject).SubjectSetRelation(subjectSetRelation).MaxDepth(maxDepth).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReadApi.GetCheck``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCheck`: GetCheckResponse
    fmt.Fprintf(os.Stdout, "Response from `ReadApi.GetCheck`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **namespace** | **string** | Namespace of the Relation Tuple | 
 **object** | **string** | Object of the Relation Tuple | 
 **relation** | **string** | Relation of the Relation Tuple | 
 **subjectId** | **string** | SubjectID of the Relation Tuple | 
 **subjectSetNamespace** | **string** | Namespace of the Subject Set | 
 **subjectSetObject** | **string** | Object of the Subject Set | 
 **subjectSetRelation** | **string** | Relation of the Subject Set | 
 **maxDepth** | **int64** |  | 

### Return type

[**GetCheckResponse**](GetCheckResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCheckMirrorStatus

> GetCheckResponse GetCheckMirrorStatus(ctx).Execute()

Check a relation tuple



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReadApi.GetCheckMirrorStatus(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReadApi.GetCheckMirrorStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCheckMirrorStatus`: GetCheckResponse
    fmt.Fprintf(os.Stdout, "Response from `ReadApi.GetCheckMirrorStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCheckMirrorStatusRequest struct via the builder pattern


### Return type

[**GetCheckResponse**](GetCheckResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExpand

> ExpandTree GetExpand(ctx).Namespace(namespace).Object(object).Relation(relation).MaxDepth(maxDepth).Execute()

Expand a Relation Tuple



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
    resp, r, err := apiClient.ReadApi.GetExpand(context.Background()).Namespace(namespace).Object(object).Relation(relation).MaxDepth(maxDepth).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReadApi.GetExpand``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExpand`: ExpandTree
    fmt.Fprintf(os.Stdout, "Response from `ReadApi.GetExpand`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetExpandRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **namespace** | **string** | Namespace of the Subject Set | 
 **object** | **string** | Object of the Subject Set | 
 **relation** | **string** | Relation of the Subject Set | 
 **maxDepth** | **int64** |  | 

### Return type

[**ExpandTree**](ExpandTree.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRelationTuples

> GetRelationTuplesResponse GetRelationTuples(ctx).PageToken(pageToken).PageSize(pageSize).Namespace(namespace).Object(object).Relation(relation).SubjectId(subjectId).SubjectSetNamespace(subjectSetNamespace).SubjectSetObject(subjectSetObject).SubjectSetRelation(subjectSetRelation).Execute()

Query relation tuples



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
    pageToken := "pageToken_example" // string |  (optional)
    pageSize := int64(789) // int64 |  (optional)
    namespace := "namespace_example" // string | Namespace of the Relation Tuple (optional)
    object := "object_example" // string | Object of the Relation Tuple (optional)
    relation := "relation_example" // string | Relation of the Relation Tuple (optional)
    subjectId := "subjectId_example" // string | SubjectID of the Relation Tuple (optional)
    subjectSetNamespace := "subjectSetNamespace_example" // string | Namespace of the Subject Set (optional)
    subjectSetObject := "subjectSetObject_example" // string | Object of the Subject Set (optional)
    subjectSetRelation := "subjectSetRelation_example" // string | Relation of the Subject Set (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReadApi.GetRelationTuples(context.Background()).PageToken(pageToken).PageSize(pageSize).Namespace(namespace).Object(object).Relation(relation).SubjectId(subjectId).SubjectSetNamespace(subjectSetNamespace).SubjectSetObject(subjectSetObject).SubjectSetRelation(subjectSetRelation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReadApi.GetRelationTuples``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRelationTuples`: GetRelationTuplesResponse
    fmt.Fprintf(os.Stdout, "Response from `ReadApi.GetRelationTuples`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRelationTuplesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageToken** | **string** |  | 
 **pageSize** | **int64** |  | 
 **namespace** | **string** | Namespace of the Relation Tuple | 
 **object** | **string** | Object of the Relation Tuple | 
 **relation** | **string** | Relation of the Relation Tuple | 
 **subjectId** | **string** | SubjectID of the Relation Tuple | 
 **subjectSetNamespace** | **string** | Namespace of the Subject Set | 
 **subjectSetObject** | **string** | Object of the Subject Set | 
 **subjectSetRelation** | **string** | Relation of the Subject Set | 

### Return type

[**GetRelationTuplesResponse**](GetRelationTuplesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCheck

> GetCheckResponse PostCheck(ctx).MaxDepth(maxDepth).RelationQuery(relationQuery).Execute()

Check a relation tuple



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
    relationQuery := *openapiclient.NewRelationQuery() // RelationQuery |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReadApi.PostCheck(context.Background()).MaxDepth(maxDepth).RelationQuery(relationQuery).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReadApi.PostCheck``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCheck`: GetCheckResponse
    fmt.Fprintf(os.Stdout, "Response from `ReadApi.PostCheck`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **maxDepth** | **int64** |  | 
 **relationQuery** | [**RelationQuery**](RelationQuery.md) |  | 

### Return type

[**GetCheckResponse**](GetCheckResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCheckMirrorStatus

> GetCheckResponse PostCheckMirrorStatus(ctx).Execute()

Check a relation tuple



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReadApi.PostCheckMirrorStatus(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReadApi.PostCheckMirrorStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostCheckMirrorStatus`: GetCheckResponse
    fmt.Fprintf(os.Stdout, "Response from `ReadApi.PostCheckMirrorStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostCheckMirrorStatusRequest struct via the builder pattern


### Return type

[**GetCheckResponse**](GetCheckResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

