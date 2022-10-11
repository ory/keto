# PostCheckOplSyntaxResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to [**[]ParseError**](ParseError.md) | The list of syntax errors | [optional] 

## Methods

### NewPostCheckOplSyntaxResponse

`func NewPostCheckOplSyntaxResponse() *PostCheckOplSyntaxResponse`

NewPostCheckOplSyntaxResponse instantiates a new PostCheckOplSyntaxResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostCheckOplSyntaxResponseWithDefaults

`func NewPostCheckOplSyntaxResponseWithDefaults() *PostCheckOplSyntaxResponse`

NewPostCheckOplSyntaxResponseWithDefaults instantiates a new PostCheckOplSyntaxResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *PostCheckOplSyntaxResponse) GetErrors() []ParseError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *PostCheckOplSyntaxResponse) GetErrorsOk() (*[]ParseError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *PostCheckOplSyntaxResponse) SetErrors(v []ParseError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *PostCheckOplSyntaxResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


