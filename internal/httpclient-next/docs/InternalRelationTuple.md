# InternalRelationTuple

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Namespace** | **string** | Namespace of the Relation Tuple | 
**Object** | **string** | Object of the Relation Tuple | 
**Relation** | **string** | Relation of the Relation Tuple | 
**SubjectId** | Pointer to **string** | SubjectID of the Relation Tuple  Either SubjectSet or SubjectID are required. | [optional] 
**SubjectSet** | Pointer to [**SubjectSet**](SubjectSet.md) |  | [optional] 

## Methods

### NewInternalRelationTuple

`func NewInternalRelationTuple(namespace string, object string, relation string, ) *InternalRelationTuple`

NewInternalRelationTuple instantiates a new InternalRelationTuple object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternalRelationTupleWithDefaults

`func NewInternalRelationTupleWithDefaults() *InternalRelationTuple`

NewInternalRelationTupleWithDefaults instantiates a new InternalRelationTuple object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNamespace

`func (o *InternalRelationTuple) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *InternalRelationTuple) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *InternalRelationTuple) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.


### GetObject

`func (o *InternalRelationTuple) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *InternalRelationTuple) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *InternalRelationTuple) SetObject(v string)`

SetObject sets Object field to given value.


### GetRelation

`func (o *InternalRelationTuple) GetRelation() string`

GetRelation returns the Relation field if non-nil, zero value otherwise.

### GetRelationOk

`func (o *InternalRelationTuple) GetRelationOk() (*string, bool)`

GetRelationOk returns a tuple with the Relation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelation

`func (o *InternalRelationTuple) SetRelation(v string)`

SetRelation sets Relation field to given value.


### GetSubjectId

`func (o *InternalRelationTuple) GetSubjectId() string`

GetSubjectId returns the SubjectId field if non-nil, zero value otherwise.

### GetSubjectIdOk

`func (o *InternalRelationTuple) GetSubjectIdOk() (*string, bool)`

GetSubjectIdOk returns a tuple with the SubjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectId

`func (o *InternalRelationTuple) SetSubjectId(v string)`

SetSubjectId sets SubjectId field to given value.

### HasSubjectId

`func (o *InternalRelationTuple) HasSubjectId() bool`

HasSubjectId returns a boolean if a field has been set.

### GetSubjectSet

`func (o *InternalRelationTuple) GetSubjectSet() SubjectSet`

GetSubjectSet returns the SubjectSet field if non-nil, zero value otherwise.

### GetSubjectSetOk

`func (o *InternalRelationTuple) GetSubjectSetOk() (*SubjectSet, bool)`

GetSubjectSetOk returns a tuple with the SubjectSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectSet

`func (o *InternalRelationTuple) SetSubjectSet(v SubjectSet)`

SetSubjectSet sets SubjectSet field to given value.

### HasSubjectSet

`func (o *InternalRelationTuple) HasSubjectSet() bool`

HasSubjectSet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


