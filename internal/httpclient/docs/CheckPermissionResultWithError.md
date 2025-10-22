# CheckPermissionResultWithError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Allowed** | **bool** | Whether the specified subject (id) is related to the requested object.  It is false by default if no ACL matches. | 
**Error** | Pointer to **string** | If there was an error checking the tuple, this will contain the error message.  If the check was performed successfully, this will be empty. | [optional] 

## Methods

### NewCheckPermissionResultWithError

`func NewCheckPermissionResultWithError(allowed bool, ) *CheckPermissionResultWithError`

NewCheckPermissionResultWithError instantiates a new CheckPermissionResultWithError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckPermissionResultWithErrorWithDefaults

`func NewCheckPermissionResultWithErrorWithDefaults() *CheckPermissionResultWithError`

NewCheckPermissionResultWithErrorWithDefaults instantiates a new CheckPermissionResultWithError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowed

`func (o *CheckPermissionResultWithError) GetAllowed() bool`

GetAllowed returns the Allowed field if non-nil, zero value otherwise.

### GetAllowedOk

`func (o *CheckPermissionResultWithError) GetAllowedOk() (*bool, bool)`

GetAllowedOk returns a tuple with the Allowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowed

`func (o *CheckPermissionResultWithError) SetAllowed(v bool)`

SetAllowed sets Allowed field to given value.


### GetError

`func (o *CheckPermissionResultWithError) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *CheckPermissionResultWithError) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *CheckPermissionResultWithError) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *CheckPermissionResultWithError) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


