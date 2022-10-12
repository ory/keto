# ParseError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**End** | Pointer to [**SourcePosition**](SourcePosition.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Start** | Pointer to [**SourcePosition**](SourcePosition.md) |  | [optional] 

## Methods

### NewParseError

`func NewParseError() *ParseError`

NewParseError instantiates a new ParseError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParseErrorWithDefaults

`func NewParseErrorWithDefaults() *ParseError`

NewParseErrorWithDefaults instantiates a new ParseError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnd

`func (o *ParseError) GetEnd() SourcePosition`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *ParseError) GetEndOk() (*SourcePosition, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *ParseError) SetEnd(v SourcePosition)`

SetEnd sets End field to given value.

### HasEnd

`func (o *ParseError) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetMessage

`func (o *ParseError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ParseError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ParseError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ParseError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStart

`func (o *ParseError) GetStart() SourcePosition`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ParseError) GetStartOk() (*SourcePosition, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ParseError) SetStart(v SourcePosition)`

SetStart sets Start field to given value.

### HasStart

`func (o *ParseError) HasStart() bool`

HasStart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


