# \RelationshipApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckOplSyntax**](RelationshipApi.md#CheckOplSyntax) | **Post** /opl/syntax/check | Performs a syntax check request.
[**CreateRelationship**](RelationshipApi.md#CreateRelationship) | **Put** /admin/relation-tuples | Creates a relationship
[**DeleteRelationships**](RelationshipApi.md#DeleteRelationships) | **Delete** /admin/relation-tuples | Deletes relationships based on relation query
[**GetRelationships**](RelationshipApi.md#GetRelationships) | **Get** /relation-tuples | Lists ACL relationships.
[**ListRelationshipNamespaces**](RelationshipApi.md#ListRelationshipNamespaces) | **Get** /namespaces | Lists Namespaces
[**PatchRelationships**](RelationshipApi.md#PatchRelationships) | **Patch** /admin/relation-tuples | Writes one or more relationships in a single transaction.



## CheckOplSyntax

> CheckOplSyntaxResult CheckOplSyntax(ctx).Body(body).Execute()

Performs a syntax check request.

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
    body := string(BYTE_ARRAY_DATA_HERE) // string | 

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

Creates a relationship

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
    createRelationshipBody := *openapiclient.NewCreateRelationshipBody() // CreateRelationshipBody | The relationship to create.

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
 **createRelationshipBody** | [**CreateRelationshipBody**](CreateRelationshipBody.md) | The relationship to create. | 

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

Deletes relationships based on relation query

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
    namespace := "namespace_example" // string | The namespace this relation tuple lives in. (optional)
    object := "object_example" // string | The object related by this tuple. It is an object in the namespace of the tuple. (optional)
    relation := "relation_example" // string | The relation between an Object and a Subject. (optional)
    subjectId := "subjectId_example" // string | A concrete id of the subject. (optional)
    subjectSetNamespace := "subjectSetNamespace_example" // string | The namespace of the object and relation referenced in this subject set. (optional)
    subjectSetObject := "subjectSetObject_example" // string | The object related by this subject set. (optional)
    subjectSetRelation := "subjectSetRelation_example" // string | The relation between the object and the subjects. (optional)

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
 **namespace** | **string** | The namespace this relation tuple lives in. | 
 **object** | **string** | The object related by this tuple. It is an object in the namespace of the tuple. | 
 **relation** | **string** | The relation between an Object and a Subject. | 
 **subjectId** | **string** | A concrete id of the subject. | 
 **subjectSetNamespace** | **string** | The namespace of the object and relation referenced in this subject set. | 
 **subjectSetObject** | **string** | The object related by this subject set. | 
 **subjectSetRelation** | **string** | The relation between the object and the subjects. | 

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

> Relationships GetRelationships(ctx).PageSize(pageSize).PageToken(pageToken).Namespace(namespace).Object(object).Relation(relation).SubjectId(subjectId).SubjectSetNamespace(subjectSetNamespace).SubjectSetObject(subjectSetObject).SubjectSetRelation(subjectSetRelation).Execute()

Lists ACL relationships.

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
    pageSize := int32(56) // int32 | Optional. The maximum number of RelationTuples to return in the response.  Default: 100 (optional)
    pageToken := "pageToken_example" // string | Optional. An opaque pagination token returned from a previous call to `ListRelationTuples` that indicates where the page should start at.  An empty token denotes the first page. All successive pages require the token from the previous page. (optional)
    namespace := "namespace_example" // string | The namespace (optional)
    object := "object_example" // string | The related object in this check. (optional)
    relation := "relation_example" // string | The relation between the Object and the Subject. (optional)
    subjectId := "subjectId_example" // string | A concrete id of the subject. (optional)
    subjectSetNamespace := "subjectSetNamespace_example" // string | The namespace of the object and relation referenced in this subject set. (optional)
    subjectSetObject := "subjectSetObject_example" // string | The object related by this subject set. (optional)
    subjectSetRelation := "subjectSetRelation_example" // string | The relation between the object and the subjects. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RelationshipApi.GetRelationships(context.Background()).PageSize(pageSize).PageToken(pageToken).Namespace(namespace).Object(object).Relation(relation).SubjectId(subjectId).SubjectSetNamespace(subjectSetNamespace).SubjectSetObject(subjectSetObject).SubjectSetRelation(subjectSetRelation).Execute()
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
 **pageSize** | **int32** | Optional. The maximum number of RelationTuples to return in the response.  Default: 100 | 
 **pageToken** | **string** | Optional. An opaque pagination token returned from a previous call to &#x60;ListRelationTuples&#x60; that indicates where the page should start at.  An empty token denotes the first page. All successive pages require the token from the previous page. | 
 **namespace** | **string** | The namespace | 
 **object** | **string** | The related object in this check. | 
 **relation** | **string** | The relation between the Object and the Subject. | 
 **subjectId** | **string** | A concrete id of the subject. | 
 **subjectSetNamespace** | **string** | The namespace of the object and relation referenced in this subject set. | 
 **subjectSetObject** | **string** | The object related by this subject set. | 
 **subjectSetRelation** | **string** | The relation between the object and the subjects. | 

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

Lists Namespaces



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

> PatchRelationships(ctx).RelationshipDelta(relationshipDelta).Execute()

Writes one or more relationships in a single transaction.

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
    relationshipDelta := []openapiclient.RelationshipDelta{*openapiclient.NewRelationshipDelta()} // []RelationshipDelta | The write delta for the relationships operated in one single transaction. Either all actions succeed or no change takes effect on error.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RelationshipApi.PatchRelationships(context.Background()).RelationshipDelta(relationshipDelta).Execute()
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
 **relationshipDelta** | [**[]RelationshipDelta**](RelationshipDelta.md) | The write delta for the relationships operated in one single transaction. Either all actions succeed or no change takes effect on error. | 

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

