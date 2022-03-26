# \WriteApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRelationTuple**](WriteApi.md#CreateRelationTuple) | **Put** /admin/relation-tuples | Create a Relation Tuple
[**DeleteRelationTuples**](WriteApi.md#DeleteRelationTuples) | **Delete** /admin/relation-tuples | Delete Relation Tuples
[**PatchRelationTuples**](WriteApi.md#PatchRelationTuples) | **Patch** /admin/relation-tuples | Patch Multiple Relation Tuples



## CreateRelationTuple

> RelationQuery CreateRelationTuple(ctx).RelationQuery(relationQuery).Execute()

Create a Relation Tuple



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
    resp, r, err := apiClient.WriteApi.CreateRelationTuple(context.Background()).RelationQuery(relationQuery).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WriteApi.CreateRelationTuple``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRelationTuple`: RelationQuery
    fmt.Fprintf(os.Stdout, "Response from `WriteApi.CreateRelationTuple`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRelationTupleRequest struct via the builder pattern


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


## DeleteRelationTuples

> DeleteRelationTuples(ctx).Namespace(namespace).Object(object).Relation(relation).SubjectId(subjectId).SubjectSetNamespace(subjectSetNamespace).SubjectSetObject(subjectSetObject).SubjectSetRelation(subjectSetRelation).Execute()

Delete Relation Tuples



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WriteApi.DeleteRelationTuples(context.Background()).Namespace(namespace).Object(object).Relation(relation).SubjectId(subjectId).SubjectSetNamespace(subjectSetNamespace).SubjectSetObject(subjectSetObject).SubjectSetRelation(subjectSetRelation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WriteApi.DeleteRelationTuples``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRelationTuplesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **namespace** | **string** | Namespace of the Relation Tuple | 
 **object** | **string** | Object of the Relation Tuple | 
 **relation** | **string** | Relation of the Relation Tuple | 
 **subjectId** | **string** | SubjectID of the Relation Tuple | 
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


## PatchRelationTuples

> PatchRelationTuples(ctx).PatchDelta(patchDelta).Execute()

Patch Multiple Relation Tuples



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
    resp, r, err := apiClient.WriteApi.PatchRelationTuples(context.Background()).PatchDelta(patchDelta).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WriteApi.PatchRelationTuples``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPatchRelationTuplesRequest struct via the builder pattern


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

