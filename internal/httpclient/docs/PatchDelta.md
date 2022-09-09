# PatchDelta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** |  | [optional] 
**RelationTuple** | Pointer to [**RelationTuple**](RelationTuple.md) |  | [optional] 

## Methods

### NewPatchDelta

`func NewPatchDelta() *PatchDelta`

NewPatchDelta instantiates a new PatchDelta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchDeltaWithDefaults

`func NewPatchDeltaWithDefaults() *PatchDelta`

NewPatchDeltaWithDefaults instantiates a new PatchDelta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *PatchDelta) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *PatchDelta) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *PatchDelta) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *PatchDelta) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetRelationTuple

`func (o *PatchDelta) GetRelationTuple() RelationTuple`

GetRelationTuple returns the RelationTuple field if non-nil, zero value otherwise.

### GetRelationTupleOk

`func (o *PatchDelta) GetRelationTupleOk() (*RelationTuple, bool)`

GetRelationTupleOk returns a tuple with the RelationTuple field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationTuple

`func (o *PatchDelta) SetRelationTuple(v RelationTuple)`

SetRelationTuple sets RelationTuple field to given value.

### HasRelationTuple

`func (o *PatchDelta) HasRelationTuple() bool`

HasRelationTuple returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


