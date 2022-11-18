# \RelationshipApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckOplSyntax**](RelationshipApi.md#CheckOplSyntax) | **Post** /opl/syntax/check | Check the syntax of an OPL file
[**CreateRelationships**](RelationshipApi.md#CreateRelationships) | **Put** /admin/relation-tuples | Create a Relationship
[**DeleteRelationships**](RelationshipApi.md#DeleteRelationships) | **Delete** /admin/relation-tuples | Delete Relationships
[**GetRelationshipNamespaces**](RelationshipApi.md#GetRelationshipNamespaces) | **Get** /namespaces | Query namespaces
[**GetRelationships**](RelationshipApi.md#GetRelationships) | **Get** /relation-tuples | Query relationships
[**PatchRelationships**](RelationshipApi.md#PatchRelationships) | **Patch** /admin/relation-tuples | Patch Multiple Relationships



## CheckOplSyntax

> PostCheckOplSyntaxResponse CheckOplSyntax(ctx).Body(body).Execute()

Check the syntax of an OPL file



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
    body := "body_example" // string | the OPL content to check

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RelationshipApi.CheckOplSyntax(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RelationshipApi.CheckOplSyntax``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckOplSyntax`: PostCheckOplSyntaxResponse
    fmt.Fprintf(os.Stdout, "Response from `RelationshipApi.CheckOplSyntax`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckOplSyntaxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **string** | the OPL content to check | 

### Return type

[**PostCheckOplSyntaxResponse**](PostCheckOplSyntaxResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: text/plain
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRelationships

> RelationQuery CreateRelationships(ctx).RelationQuery(relationQuery).Execute()

Create a Relationship



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
    relationQuery := *openapiclient.NewRelationQuery() // RelationQuery |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RelationshipApi.CreateRelationships(context.Background()).RelationQuery(relationQuery).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RelationshipApi.CreateRelationships``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRelationships`: RelationQuery
    fmt.Fprintf(os.Stdout, "Response from `RelationshipApi.CreateRelationships`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRelationshipsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **relationQuery** | [**RelationQuery**](RelationQuery.md) |  | 

### Return type

[**RelationQuery**](RelationQuery.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRelationships

> DeleteRelationships(ctx).Namespace(namespace).Object(object).Relation(relation).SubjectId(subjectId).SubjectSetNamespace(subjectSetNamespace).SubjectSetObject(subjectSetObject).SubjectSetRelation(subjectSetRelation).Execute()

Delete Relationships



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RelationshipApi.DeleteRelationships(context.Background()).Namespace(namespace).Object(object).Relation(relation).SubjectId(subjectId).SubjectSetNamespace(subjectSetNamespace).SubjectSetObject(subjectSetObject).SubjectSetRelation(subjectSetRelation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RelationshipApi.DeleteRelationships``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRelationshipsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **namespace** | **string** | Namespace of the Relationship | 
 **object** | **string** | Object of the Relationship | 
 **relation** | **string** | Relation of the Relationship | 
 **subjectId** | **string** | SubjectID of the Relationship | 
 **subjectSetNamespace** | **string** | Namespace of the Subject Set | 
 **subjectSetObject** | **string** | Object of the Subject Set | 
 **subjectSetRelation** | **string** | Relation of the Subject Set | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRelationshipNamespaces

> GetRelationshipNamespacesResponse GetRelationshipNamespaces(ctx).Execute()

Query namespaces



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
    resp, r, err := apiClient.RelationshipApi.GetRelationshipNamespaces(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RelationshipApi.GetRelationshipNamespaces``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRelationshipNamespaces`: GetRelationshipNamespacesResponse
    fmt.Fprintf(os.Stdout, "Response from `RelationshipApi.GetRelationshipNamespaces`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRelationshipNamespacesRequest struct via the builder pattern


### Return type

[**GetRelationshipNamespacesResponse**](GetRelationshipNamespacesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRelationships

> GetRelationshipsResponse GetRelationships(ctx).PageToken(pageToken).PageSize(pageSize).Namespace(namespace).Object(object).Relation(relation).SubjectId(subjectId).SubjectSetNamespace(subjectSetNamespace).SubjectSetObject(subjectSetObject).SubjectSetRelation(subjectSetRelation).Execute()

Query relationships



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
    namespace := "namespace_example" // string | Namespace of the Relationship (optional)
    object := "object_example" // string | Object of the Relationship (optional)
    relation := "relation_example" // string | Relation of the Relationship (optional)
    subjectId := "subjectId_example" // string | SubjectID of the Relationship (optional)
    subjectSetNamespace := "subjectSetNamespace_example" // string | Namespace of the Subject Set (optional)
    subjectSetObject := "subjectSetObject_example" // string | Object of the Subject Set (optional)
    subjectSetRelation := "subjectSetRelation_example" // string | Relation of the Subject Set (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RelationshipApi.GetRelationships(context.Background()).PageToken(pageToken).PageSize(pageSize).Namespace(namespace).Object(object).Relation(relation).SubjectId(subjectId).SubjectSetNamespace(subjectSetNamespace).SubjectSetObject(subjectSetObject).SubjectSetRelation(subjectSetRelation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RelationshipApi.GetRelationships``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRelationships`: GetRelationshipsResponse
    fmt.Fprintf(os.Stdout, "Response from `RelationshipApi.GetRelationships`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRelationshipsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageToken** | **string** |  | 
 **pageSize** | **int64** |  | 
 **namespace** | **string** | Namespace of the Relationship | 
 **object** | **string** | Object of the Relationship | 
 **relation** | **string** | Relation of the Relationship | 
 **subjectId** | **string** | SubjectID of the Relationship | 
 **subjectSetNamespace** | **string** | Namespace of the Subject Set | 
 **subjectSetObject** | **string** | Object of the Subject Set | 
 **subjectSetRelation** | **string** | Relation of the Subject Set | 

### Return type

[**GetRelationshipsResponse**](GetRelationshipsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchRelationships

> PatchRelationships(ctx).PatchDelta(patchDelta).Execute()

Patch Multiple Relationships



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
    patchDelta := []openapiclient.PatchDelta{*openapiclient.NewPatchDelta()} // []PatchDelta |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RelationshipApi.PatchRelationships(context.Background()).PatchDelta(patchDelta).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RelationshipApi.PatchRelationships``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPatchRelationshipsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchDelta** | [**[]PatchDelta**](PatchDelta.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

