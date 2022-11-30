# Relationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextPageToken** | Pointer to **string** | The opaque token to provide in a subsequent request to get the next page. It is the empty string iff this is the last page. | [optional] 
**RelationTuples** | Pointer to [**[]Relationship**](Relationship.md) |  | [optional] 

## Methods

### NewRelationships

`func NewRelationships() *Relationships`

NewRelationships instantiates a new Relationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelationshipsWithDefaults

`func NewRelationshipsWithDefaults() *Relationships`

NewRelationshipsWithDefaults instantiates a new Relationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextPageToken

`func (o *Relationships) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *Relationships) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *Relationships) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *Relationships) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetRelationTuples

`func (o *Relationships) GetRelationTuples() []Relationship`

GetRelationTuples returns the RelationTuples field if non-nil, zero value otherwise.

### GetRelationTuplesOk

`func (o *Relationships) GetRelationTuplesOk() (*[]Relationship, bool)`

GetRelationTuplesOk returns a tuple with the RelationTuples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationTuples

`func (o *Relationships) SetRelationTuples(v []Relationship)`

SetRelationTuples sets RelationTuples field to given value.

### HasRelationTuples

`func (o *Relationships) HasRelationTuples() bool`

HasRelationTuples returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


