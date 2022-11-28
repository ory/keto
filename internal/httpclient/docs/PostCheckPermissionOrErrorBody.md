# PostCheckPermissionOrErrorBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Namespace** | Pointer to **string** | Namespace to query | [optional] 
**Object** | Pointer to **string** | Object to query | [optional] 
**Relation** | Pointer to **string** | Relation to query | [optional] 
**SubjectId** | Pointer to **string** | SubjectID to query  Either SubjectSet or SubjectID can be provided. | [optional] 
**SubjectSet** | Pointer to [**SubjectSet**](SubjectSet.md) |  | [optional] 

## Methods

### NewPostCheckPermissionOrErrorBody

`func NewPostCheckPermissionOrErrorBody() *PostCheckPermissionOrErrorBody`

NewPostCheckPermissionOrErrorBody instantiates a new PostCheckPermissionOrErrorBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostCheckPermissionOrErrorBodyWithDefaults

`func NewPostCheckPermissionOrErrorBodyWithDefaults() *PostCheckPermissionOrErrorBody`

NewPostCheckPermissionOrErrorBodyWithDefaults instantiates a new PostCheckPermissionOrErrorBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNamespace

`func (o *PostCheckPermissionOrErrorBody) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *PostCheckPermissionOrErrorBody) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *PostCheckPermissionOrErrorBody) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *PostCheckPermissionOrErrorBody) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetObject

`func (o *PostCheckPermissionOrErrorBody) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *PostCheckPermissionOrErrorBody) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *PostCheckPermissionOrErrorBody) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *PostCheckPermissionOrErrorBody) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetRelation

`func (o *PostCheckPermissionOrErrorBody) GetRelation() string`

GetRelation returns the Relation field if non-nil, zero value otherwise.

### GetRelationOk

`func (o *PostCheckPermissionOrErrorBody) GetRelationOk() (*string, bool)`

GetRelationOk returns a tuple with the Relation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelation

`func (o *PostCheckPermissionOrErrorBody) SetRelation(v string)`

SetRelation sets Relation field to given value.

### HasRelation

`func (o *PostCheckPermissionOrErrorBody) HasRelation() bool`

HasRelation returns a boolean if a field has been set.

### GetSubjectId

`func (o *PostCheckPermissionOrErrorBody) GetSubjectId() string`

GetSubjectId returns the SubjectId field if non-nil, zero value otherwise.

### GetSubjectIdOk

`func (o *PostCheckPermissionOrErrorBody) GetSubjectIdOk() (*string, bool)`

GetSubjectIdOk returns a tuple with the SubjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectId

`func (o *PostCheckPermissionOrErrorBody) SetSubjectId(v string)`

SetSubjectId sets SubjectId field to given value.

### HasSubjectId

`func (o *PostCheckPermissionOrErrorBody) HasSubjectId() bool`

HasSubjectId returns a boolean if a field has been set.

### GetSubjectSet

`func (o *PostCheckPermissionOrErrorBody) GetSubjectSet() SubjectSet`

GetSubjectSet returns the SubjectSet field if non-nil, zero value otherwise.

### GetSubjectSetOk

`func (o *PostCheckPermissionOrErrorBody) GetSubjectSetOk() (*SubjectSet, bool)`

GetSubjectSetOk returns a tuple with the SubjectSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectSet

`func (o *PostCheckPermissionOrErrorBody) SetSubjectSet(v SubjectSet)`

SetSubjectSet sets SubjectSet field to given value.

### HasSubjectSet

`func (o *PostCheckPermissionOrErrorBody) HasSubjectSet() bool`

HasSubjectSet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


