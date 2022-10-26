# GetNamespacesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Namespaces** | Pointer to [**[]Namespace**](Namespace.md) |  | [optional] 

## Methods

### NewGetNamespacesResponse

`func NewGetNamespacesResponse() *GetNamespacesResponse`

NewGetNamespacesResponse instantiates a new GetNamespacesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNamespacesResponseWithDefaults

`func NewGetNamespacesResponseWithDefaults() *GetNamespacesResponse`

NewGetNamespacesResponseWithDefaults instantiates a new GetNamespacesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNamespaces

`func (o *GetNamespacesResponse) GetNamespaces() []Namespace`

GetNamespaces returns the Namespaces field if non-nil, zero value otherwise.

### GetNamespacesOk

`func (o *GetNamespacesResponse) GetNamespacesOk() (*[]Namespace, bool)`

GetNamespacesOk returns a tuple with the Namespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaces

`func (o *GetNamespacesResponse) SetNamespaces(v []Namespace)`

SetNamespaces sets Namespaces field to given value.

### HasNamespaces

`func (o *GetNamespacesResponse) HasNamespaces() bool`

HasNamespaces returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


