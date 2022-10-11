# SourcePosition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Line** | Pointer to **int64** |  | [optional] 
**Column** | Pointer to **int64** |  | [optional] 

## Methods

### NewSourcePosition

`func NewSourcePosition() *SourcePosition`

NewSourcePosition instantiates a new SourcePosition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourcePositionWithDefaults

`func NewSourcePositionWithDefaults() *SourcePosition`

NewSourcePositionWithDefaults instantiates a new SourcePosition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLine

`func (o *SourcePosition) GetLine() int64`

GetLine returns the Line field if non-nil, zero value otherwise.

### GetLineOk

`func (o *SourcePosition) GetLineOk() (*int64, bool)`

GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine

`func (o *SourcePosition) SetLine(v int64)`

SetLine sets Line field to given value.

### HasLine

`func (o *SourcePosition) HasLine() bool`

HasLine returns a boolean if a field has been set.

### GetColumn

`func (o *SourcePosition) GetColumn() int64`

GetColumn returns the Column field if non-nil, zero value otherwise.

### GetColumnOk

`func (o *SourcePosition) GetColumnOk() (*int64, bool)`

GetColumnOk returns a tuple with the Column field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumn

`func (o *SourcePosition) SetColumn(v int64)`

SetColumn sets Column field to given value.

### HasColumn

`func (o *SourcePosition) HasColumn() bool`

HasColumn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


