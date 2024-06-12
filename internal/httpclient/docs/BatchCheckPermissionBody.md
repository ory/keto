# BatchCheckPermissionBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tuples** | Pointer to [**[]Relationship**](Relationship.md) |  | [optional] 

## Methods

### NewBatchCheckPermissionBody

`func NewBatchCheckPermissionBody() *BatchCheckPermissionBody`

NewBatchCheckPermissionBody instantiates a new BatchCheckPermissionBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchCheckPermissionBodyWithDefaults

`func NewBatchCheckPermissionBodyWithDefaults() *BatchCheckPermissionBody`

NewBatchCheckPermissionBodyWithDefaults instantiates a new BatchCheckPermissionBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTuples

`func (o *BatchCheckPermissionBody) GetTuples() []Relationship`

GetTuples returns the Tuples field if non-nil, zero value otherwise.

### GetTuplesOk

`func (o *BatchCheckPermissionBody) GetTuplesOk() (*[]Relationship, bool)`

GetTuplesOk returns a tuple with the Tuples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTuples

`func (o *BatchCheckPermissionBody) SetTuples(v []Relationship)`

SetTuples sets Tuples field to given value.

### HasTuples

`func (o *BatchCheckPermissionBody) HasTuples() bool`

HasTuples returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


