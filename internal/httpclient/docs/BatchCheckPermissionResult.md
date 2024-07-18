# BatchCheckPermissionResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | [**[]CheckPermissionResultWithError**](CheckPermissionResultWithError.md) | An array of check results. The order aligns with the input order. | 

## Methods

### NewBatchCheckPermissionResult

`func NewBatchCheckPermissionResult(results []CheckPermissionResultWithError, ) *BatchCheckPermissionResult`

NewBatchCheckPermissionResult instantiates a new BatchCheckPermissionResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchCheckPermissionResultWithDefaults

`func NewBatchCheckPermissionResultWithDefaults() *BatchCheckPermissionResult`

NewBatchCheckPermissionResultWithDefaults instantiates a new BatchCheckPermissionResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *BatchCheckPermissionResult) GetResults() []CheckPermissionResultWithError`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *BatchCheckPermissionResult) GetResultsOk() (*[]CheckPermissionResultWithError, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *BatchCheckPermissionResult) SetResults(v []CheckPermissionResultWithError)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


