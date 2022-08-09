# ExpandTree

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | Pointer to [**[]ExpandTree**](ExpandTree.md) | The children of the node, possibly none. | [optional] 
**Tuple** | Pointer to [**RelationTuple**](RelationTuple.md) |  | [optional] 
**Type** | **string** | The type of the node. union TreeNodeUnion exclusion TreeNodeExclusion intersection TreeNodeIntersection leaf TreeNodeLeaf tuple_to_subject_set TreeNodeTupleToSubjectSet computed_subject_set TreeNodeComputedSubjectSet not TreeNodeNot unspecified TreeNodeUnspecified | 

## Methods

### NewExpandTree

`func NewExpandTree(type_ string, ) *ExpandTree`

NewExpandTree instantiates a new ExpandTree object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpandTreeWithDefaults

`func NewExpandTreeWithDefaults() *ExpandTree`

NewExpandTreeWithDefaults instantiates a new ExpandTree object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChildren

`func (o *ExpandTree) GetChildren() []ExpandTree`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *ExpandTree) GetChildrenOk() (*[]ExpandTree, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *ExpandTree) SetChildren(v []ExpandTree)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *ExpandTree) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetTuple

`func (o *ExpandTree) GetTuple() RelationTuple`

GetTuple returns the Tuple field if non-nil, zero value otherwise.

### GetTupleOk

`func (o *ExpandTree) GetTupleOk() (*RelationTuple, bool)`

GetTupleOk returns a tuple with the Tuple field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTuple

`func (o *ExpandTree) SetTuple(v RelationTuple)`

SetTuple sets Tuple field to given value.

### HasTuple

`func (o *ExpandTree) HasTuple() bool`

HasTuple returns a boolean if a field has been set.

### GetType

`func (o *ExpandTree) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExpandTree) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExpandTree) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


