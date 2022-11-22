# CreateRelationshipBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Namespace** | Pointer to **string** | Namespace to query | [optional] 
**Object** | Pointer to **string** | Object to query | [optional] 
**Relation** | Pointer to **string** | Relation to query | [optional] 
**SubjectId** | Pointer to **string** | SubjectID to query  Either SubjectSet or SubjectID can be provided. | [optional] 
**SubjectSet** | Pointer to [**SubjectSet**](SubjectSet.md) |  | [optional] 

## Methods

### NewCreateRelationshipBody

`func NewCreateRelationshipBody() *CreateRelationshipBody`

NewCreateRelationshipBody instantiates a new CreateRelationshipBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRelationshipBodyWithDefaults

`func NewCreateRelationshipBodyWithDefaults() *CreateRelationshipBody`

NewCreateRelationshipBodyWithDefaults instantiates a new CreateRelationshipBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNamespace

`func (o *CreateRelationshipBody) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *CreateRelationshipBody) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *CreateRelationshipBody) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *CreateRelationshipBody) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetObject

`func (o *CreateRelationshipBody) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *CreateRelationshipBody) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *CreateRelationshipBody) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *CreateRelationshipBody) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetRelation

`func (o *CreateRelationshipBody) GetRelation() string`

GetRelation returns the Relation field if non-nil, zero value otherwise.

### GetRelationOk

`func (o *CreateRelationshipBody) GetRelationOk() (*string, bool)`

GetRelationOk returns a tuple with the Relation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelation

`func (o *CreateRelationshipBody) SetRelation(v string)`

SetRelation sets Relation field to given value.

### HasRelation

`func (o *CreateRelationshipBody) HasRelation() bool`

HasRelation returns a boolean if a field has been set.

### GetSubjectId

`func (o *CreateRelationshipBody) GetSubjectId() string`

GetSubjectId returns the SubjectId field if non-nil, zero value otherwise.

### GetSubjectIdOk

`func (o *CreateRelationshipBody) GetSubjectIdOk() (*string, bool)`

GetSubjectIdOk returns a tuple with the SubjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectId

`func (o *CreateRelationshipBody) SetSubjectId(v string)`

SetSubjectId sets SubjectId field to given value.

### HasSubjectId

`func (o *CreateRelationshipBody) HasSubjectId() bool`

HasSubjectId returns a boolean if a field has been set.

### GetSubjectSet

`func (o *CreateRelationshipBody) GetSubjectSet() SubjectSet`

GetSubjectSet returns the SubjectSet field if non-nil, zero value otherwise.

### GetSubjectSetOk

`func (o *CreateRelationshipBody) GetSubjectSetOk() (*SubjectSet, bool)`

GetSubjectSetOk returns a tuple with the SubjectSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectSet

`func (o *CreateRelationshipBody) SetSubjectSet(v SubjectSet)`

SetSubjectSet sets SubjectSet field to given value.

### HasSubjectSet

`func (o *CreateRelationshipBody) HasSubjectSet() bool`

HasSubjectSet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


