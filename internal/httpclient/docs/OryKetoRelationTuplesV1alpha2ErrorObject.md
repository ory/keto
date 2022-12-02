# OryKetoRelationTuplesV1alpha2ErrorObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **int64** | The status code | [optional] 
**Debug** | Pointer to **string** | This field is often not exposed to protect against leaking sensitive information | [optional] 
**Details** | Pointer to **map[string]string** |  | [optional] 
**Id** | Pointer to **string** | Useful when trying to identify various errors in application logic. | [optional] 
**Message** | Pointer to **string** | Response message | [optional] 
**Reason** | Pointer to **string** | A human-readable reason for the error | [optional] 
**Request** | Pointer to **string** | The request ID is often exposed internally in order to trace errors across service architectures. This is often a UUID. | [optional] 
**Status** | Pointer to **string** | The human-readable description of the code. | [optional] 

## Methods

### NewOryKetoRelationTuplesV1alpha2ErrorObject

`func NewOryKetoRelationTuplesV1alpha2ErrorObject() *OryKetoRelationTuplesV1alpha2ErrorObject`

NewOryKetoRelationTuplesV1alpha2ErrorObject instantiates a new OryKetoRelationTuplesV1alpha2ErrorObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOryKetoRelationTuplesV1alpha2ErrorObjectWithDefaults

`func NewOryKetoRelationTuplesV1alpha2ErrorObjectWithDefaults() *OryKetoRelationTuplesV1alpha2ErrorObject`

NewOryKetoRelationTuplesV1alpha2ErrorObjectWithDefaults instantiates a new OryKetoRelationTuplesV1alpha2ErrorObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *OryKetoRelationTuplesV1alpha2ErrorObject) GetCode() int64`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *OryKetoRelationTuplesV1alpha2ErrorObject) GetCodeOk() (*int64, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *OryKetoRelationTuplesV1alpha2ErrorObject) SetCode(v int64)`

SetCode sets Code field to given value.

### HasCode

`func (o *OryKetoRelationTuplesV1alpha2ErrorObject) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDebug

`func (o *OryKetoRelationTuplesV1alpha2ErrorObject) GetDebug() string`

GetDebug returns the Debug field if non-nil, zero value otherwise.

### GetDebugOk

`func (o *OryKetoRelationTuplesV1alpha2ErrorObject) GetDebugOk() (*string, bool)`

GetDebugOk returns a tuple with the Debug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebug

`func (o *OryKetoRelationTuplesV1alpha2ErrorObject) SetDebug(v string)`

SetDebug sets Debug field to given value.

### HasDebug

`func (o *OryKetoRelationTuplesV1alpha2ErrorObject) HasDebug() bool`

HasDebug returns a boolean if a field has been set.

### GetDetails

`func (o *OryKetoRelationTuplesV1alpha2ErrorObject) GetDetails() map[string]string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *OryKetoRelationTuplesV1alpha2ErrorObject) GetDetailsOk() (*map[string]string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *OryKetoRelationTuplesV1alpha2ErrorObject) SetDetails(v map[string]string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *OryKetoRelationTuplesV1alpha2ErrorObject) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetId

`func (o *OryKetoRelationTuplesV1alpha2ErrorObject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OryKetoRelationTuplesV1alpha2ErrorObject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OryKetoRelationTuplesV1alpha2ErrorObject) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OryKetoRelationTuplesV1alpha2ErrorObject) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMessage

`func (o *OryKetoRelationTuplesV1alpha2ErrorObject) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *OryKetoRelationTuplesV1alpha2ErrorObject) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *OryKetoRelationTuplesV1alpha2ErrorObject) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *OryKetoRelationTuplesV1alpha2ErrorObject) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetReason

`func (o *OryKetoRelationTuplesV1alpha2ErrorObject) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *OryKetoRelationTuplesV1alpha2ErrorObject) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *OryKetoRelationTuplesV1alpha2ErrorObject) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *OryKetoRelationTuplesV1alpha2ErrorObject) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetRequest

`func (o *OryKetoRelationTuplesV1alpha2ErrorObject) GetRequest() string`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *OryKetoRelationTuplesV1alpha2ErrorObject) GetRequestOk() (*string, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *OryKetoRelationTuplesV1alpha2ErrorObject) SetRequest(v string)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *OryKetoRelationTuplesV1alpha2ErrorObject) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetStatus

`func (o *OryKetoRelationTuplesV1alpha2ErrorObject) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OryKetoRelationTuplesV1alpha2ErrorObject) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OryKetoRelationTuplesV1alpha2ErrorObject) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OryKetoRelationTuplesV1alpha2ErrorObject) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


