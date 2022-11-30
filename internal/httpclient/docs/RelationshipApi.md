# \RelationshipApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckOplSyntax**](RelationshipApi.md#CheckOplSyntax) | **Post** /opl/syntax/check | Check the syntax of an OPL file
[**CreateRelationship**](RelationshipApi.md#CreateRelationship) | **Put** /admin/relation-tuples | Create a Relationship
[**DeleteRelationships**](RelationshipApi.md#DeleteRelationships) | **Delete** /admin/relation-tuples | Delete Relationships
[**GetRelationships**](RelationshipApi.md#GetRelationships) | **Get** /relation-tuples | Query relationships
[**ListRelationshipNamespaces**](RelationshipApi.md#ListRelationshipNamespaces) | **Get** /namespaces | Query namespaces
[**PatchRelationships**](RelationshipApi.md#PatchRelationships) | **Patch** /admin/relation-tuples | Patch Multiple Relationships



## CheckOplSyntax

> CheckOplSyntaxResult CheckOplSyntax(ctx).Body(body).Execute()

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
    body := "body_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RelationshipApi.CheckOplSyntax(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RelationshipApi.CheckOplSyntax``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckOplSyntax`: CheckOplSyntaxResult
    fmt.Fprintf(os.Stdout, "Response from `RelationshipApi.CheckOplSyntax`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckOplSyntaxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **string** |  | 

### Return type

[**CheckOplSyntaxResult**](CheckOplSyntaxResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: text/plain
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRelationship

> Relationship CreateRelationship(ctx).CreateRelationshipBody(createRelationshipBody).Execute()

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
    createRelationshipBody := *openapiclient.NewCreateRelationshipBody() // CreateRelationshipBody |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RelationshipApi.CreateRelationship(context.Background()).CreateRelationshipBody(createRelationshipBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RelationshipApi.CreateRelationship``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRelationship`: Relationship
    fmt.Fprintf(os.Stdout, "Response from `RelationshipApi.CreateRelationship`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRelationshipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createRelationshipBody** | [**CreateRelationshipBody**](CreateRelationshipBody.md) |  | 

### Return type

[**Relationship**](Relationship.md)

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


## GetRelationships

> Relationships GetRelationships(ctx).PageToken(pageToken).PageSize(pageSize).Namespace(namespace).Object(object).Relation(relation).SubjectId(subjectId).SubjectSetNamespace(subjectSetNamespace).SubjectSetObject(subjectSetObject).SubjectSetRelation(subjectSetRelation).Execute()

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
    // response from `GetRelationships`: Relationships
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

[**Relationships**](Relationships.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRelationshipNamespaces

> RelationshipNamespaces ListRelationshipNamespaces(ctx).Execute()

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
    resp, r, err := apiClient.RelationshipApi.ListRelationshipNamespaces(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RelationshipApi.ListRelationshipNamespaces``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRelationshipNamespaces`: RelationshipNamespaces
    fmt.Fprintf(os.Stdout, "Response from `RelationshipApi.ListRelationshipNamespaces`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListRelationshipNamespacesRequest struct via the builder pattern


### Return type

[**RelationshipNamespaces**](RelationshipNamespaces.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchRelationships

> PatchRelationships(ctx).RelationshipPatch(relationshipPatch).Execute()

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
    relationshipPatch := []openapiclient.RelationshipPatch{*openapiclient.NewRelationshipPatch()} // []RelationshipPatch |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RelationshipApi.PatchRelationships(context.Background()).RelationshipPatch(relationshipPatch).Execute()
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
 **relationshipPatch** | [**[]RelationshipPatch**](RelationshipPatch.md) |  | 

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

