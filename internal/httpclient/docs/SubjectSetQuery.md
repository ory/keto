# SubjectSetQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Namespace** | Pointer to **string** | The namespace of the object and relation referenced in this subject set. | [optional] 
**Object** | Pointer to **string** | The object related by this subject set. | [optional] 
**Relation** | Pointer to **string** | The relation between the object and the subjects. | [optional] 

## Methods

### NewSubjectSetQuery

`func NewSubjectSetQuery() *SubjectSetQuery`

NewSubjectSetQuery instantiates a new SubjectSetQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubjectSetQueryWithDefaults

`func NewSubjectSetQueryWithDefaults() *SubjectSetQuery`

NewSubjectSetQueryWithDefaults instantiates a new SubjectSetQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNamespace

`func (o *SubjectSetQuery) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *SubjectSetQuery) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *SubjectSetQuery) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *SubjectSetQuery) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetObject

`func (o *SubjectSetQuery) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *SubjectSetQuery) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *SubjectSetQuery) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *SubjectSetQuery) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetRelation

`func (o *SubjectSetQuery) GetRelation() string`

GetRelation returns the Relation field if non-nil, zero value otherwise.

### GetRelationOk

`func (o *SubjectSetQuery) GetRelationOk() (*string, bool)`

GetRelationOk returns a tuple with the Relation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelation

`func (o *SubjectSetQuery) SetRelation(v string)`

SetRelation sets Relation field to given value.

### HasRelation

`func (o *SubjectSetQuery) HasRelation() bool`

HasRelation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


