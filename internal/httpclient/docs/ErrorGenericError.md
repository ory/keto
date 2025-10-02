# ErrorGenericError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** |  | [optional] 
**Debug** | Pointer to **string** | Debug information is often not exposed to protect against leaking sensitive information. | [optional] 
**Details** | Pointer to **map[string]string** | Further details about the error. | [optional] 
**Id** | Pointer to **string** | The error ID is useful when trying to identify various errors in application logic. | [optional] 
**Message** | **string** | The error&#39;s message (required). | 
**Reason** | Pointer to **string** | Reason holds a human-readable reason for the error. | [optional] 
**Request** | Pointer to **string** | The request ID is often exposed internally in order to trace errors across service architectures. This is often a UUID. | [optional] 
**Status** | Pointer to **string** | Status holds the human-readable HTTP status code. | [optional] 

## Methods

### NewErrorGenericError

`func NewErrorGenericError(message string, ) *ErrorGenericError`

NewErrorGenericError instantiates a new ErrorGenericError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorGenericErrorWithDefaults

`func NewErrorGenericErrorWithDefaults() *ErrorGenericError`

NewErrorGenericErrorWithDefaults instantiates a new ErrorGenericError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ErrorGenericError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ErrorGenericError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ErrorGenericError) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ErrorGenericError) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDebug

`func (o *ErrorGenericError) GetDebug() string`

GetDebug returns the Debug field if non-nil, zero value otherwise.

### GetDebugOk

`func (o *ErrorGenericError) GetDebugOk() (*string, bool)`

GetDebugOk returns a tuple with the Debug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebug

`func (o *ErrorGenericError) SetDebug(v string)`

SetDebug sets Debug field to given value.

### HasDebug

`func (o *ErrorGenericError) HasDebug() bool`

HasDebug returns a boolean if a field has been set.

### GetDetails

`func (o *ErrorGenericError) GetDetails() map[string]string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ErrorGenericError) GetDetailsOk() (*map[string]string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ErrorGenericError) SetDetails(v map[string]string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *ErrorGenericError) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetId

`func (o *ErrorGenericError) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ErrorGenericError) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ErrorGenericError) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ErrorGenericError) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMessage

`func (o *ErrorGenericError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ErrorGenericError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ErrorGenericError) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetReason

`func (o *ErrorGenericError) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ErrorGenericError) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ErrorGenericError) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ErrorGenericError) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetRequest

`func (o *ErrorGenericError) GetRequest() string`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *ErrorGenericError) GetRequestOk() (*string, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *ErrorGenericError) SetRequest(v string)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *ErrorGenericError) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetStatus

`func (o *ErrorGenericError) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ErrorGenericError) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ErrorGenericError) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ErrorGenericError) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


