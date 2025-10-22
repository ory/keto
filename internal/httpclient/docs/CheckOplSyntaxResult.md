# CheckOplSyntaxResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to [**[]OryKetoOplV1alpha1ParseError**](OryKetoOplV1alpha1ParseError.md) |  | [optional] 
**ParseErrors** | Pointer to [**[]OryKetoOplV1alpha1ParseError**](OryKetoOplV1alpha1ParseError.md) |  | [optional] 

## Methods

### NewCheckOplSyntaxResult

`func NewCheckOplSyntaxResult() *CheckOplSyntaxResult`

NewCheckOplSyntaxResult instantiates a new CheckOplSyntaxResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckOplSyntaxResultWithDefaults

`func NewCheckOplSyntaxResultWithDefaults() *CheckOplSyntaxResult`

NewCheckOplSyntaxResultWithDefaults instantiates a new CheckOplSyntaxResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *CheckOplSyntaxResult) GetErrors() []OryKetoOplV1alpha1ParseError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *CheckOplSyntaxResult) GetErrorsOk() (*[]OryKetoOplV1alpha1ParseError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *CheckOplSyntaxResult) SetErrors(v []OryKetoOplV1alpha1ParseError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *CheckOplSyntaxResult) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetParseErrors

`func (o *CheckOplSyntaxResult) GetParseErrors() []OryKetoOplV1alpha1ParseError`

GetParseErrors returns the ParseErrors field if non-nil, zero value otherwise.

### GetParseErrorsOk

`func (o *CheckOplSyntaxResult) GetParseErrorsOk() (*[]OryKetoOplV1alpha1ParseError, bool)`

GetParseErrorsOk returns a tuple with the ParseErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParseErrors

`func (o *CheckOplSyntaxResult) SetParseErrors(v []OryKetoOplV1alpha1ParseError)`

SetParseErrors sets ParseErrors field to given value.

### HasParseErrors

`func (o *CheckOplSyntaxResult) HasParseErrors() bool`

HasParseErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


