# PostCheckPermissionBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxDepth** | Pointer to **int32** | The maximum depth to search for a relation.  If the value is less than 1 or greater than the global max-depth then the global max-depth will be used instead. | [optional] 
**Namespace** | Pointer to **string** | The namespace to evaluate the check.  Note: If you use the expand-API and the check evaluates a RelationTuple specifying a SubjectSet as subject or due to a rewrite rule in a namespace config this check request may involve other namespaces automatically. | [optional] 
**Object** | Pointer to **string** | The related object in this check. | [optional] 
**Relation** | Pointer to **string** | The relation between the Object and the Subject. | [optional] 
**SubjectId** | Pointer to **string** | A concrete id of the subject. | [optional] 
**SubjectSet** | Pointer to [**SubjectSetQuery**](SubjectSetQuery.md) |  | [optional] 

## Methods

### NewPostCheckPermissionBody

`func NewPostCheckPermissionBody() *PostCheckPermissionBody`

NewPostCheckPermissionBody instantiates a new PostCheckPermissionBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostCheckPermissionBodyWithDefaults

`func NewPostCheckPermissionBodyWithDefaults() *PostCheckPermissionBody`

NewPostCheckPermissionBodyWithDefaults instantiates a new PostCheckPermissionBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxDepth

`func (o *PostCheckPermissionBody) GetMaxDepth() int32`

GetMaxDepth returns the MaxDepth field if non-nil, zero value otherwise.

### GetMaxDepthOk

`func (o *PostCheckPermissionBody) GetMaxDepthOk() (*int32, bool)`

GetMaxDepthOk returns a tuple with the MaxDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDepth

`func (o *PostCheckPermissionBody) SetMaxDepth(v int32)`

SetMaxDepth sets MaxDepth field to given value.

### HasMaxDepth

`func (o *PostCheckPermissionBody) HasMaxDepth() bool`

HasMaxDepth returns a boolean if a field has been set.

### GetNamespace

`func (o *PostCheckPermissionBody) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *PostCheckPermissionBody) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *PostCheckPermissionBody) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *PostCheckPermissionBody) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetObject

`func (o *PostCheckPermissionBody) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *PostCheckPermissionBody) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *PostCheckPermissionBody) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *PostCheckPermissionBody) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetRelation

`func (o *PostCheckPermissionBody) GetRelation() string`

GetRelation returns the Relation field if non-nil, zero value otherwise.

### GetRelationOk

`func (o *PostCheckPermissionBody) GetRelationOk() (*string, bool)`

GetRelationOk returns a tuple with the Relation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelation

`func (o *PostCheckPermissionBody) SetRelation(v string)`

SetRelation sets Relation field to given value.

### HasRelation

`func (o *PostCheckPermissionBody) HasRelation() bool`

HasRelation returns a boolean if a field has been set.

### GetSubjectId

`func (o *PostCheckPermissionBody) GetSubjectId() string`

GetSubjectId returns the SubjectId field if non-nil, zero value otherwise.

### GetSubjectIdOk

`func (o *PostCheckPermissionBody) GetSubjectIdOk() (*string, bool)`

GetSubjectIdOk returns a tuple with the SubjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectId

`func (o *PostCheckPermissionBody) SetSubjectId(v string)`

SetSubjectId sets SubjectId field to given value.

### HasSubjectId

`func (o *PostCheckPermissionBody) HasSubjectId() bool`

HasSubjectId returns a boolean if a field has been set.

### GetSubjectSet

`func (o *PostCheckPermissionBody) GetSubjectSet() SubjectSetQuery`

GetSubjectSet returns the SubjectSet field if non-nil, zero value otherwise.

### GetSubjectSetOk

`func (o *PostCheckPermissionBody) GetSubjectSetOk() (*SubjectSetQuery, bool)`

GetSubjectSetOk returns a tuple with the SubjectSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectSet

`func (o *PostCheckPermissionBody) SetSubjectSet(v SubjectSetQuery)`

SetSubjectSet sets SubjectSet field to given value.

### HasSubjectSet

`func (o *PostCheckPermissionBody) HasSubjectSet() bool`

HasSubjectSet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


