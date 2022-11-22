# ExpandedPermissionTree

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | Pointer to [**[]ExpandedPermissionTree**](ExpandedPermissionTree.md) | The children of the node, possibly none. | [optional] 
**Tuple** | Pointer to [**Relationship**](Relationship.md) |  | [optional] 
**Type** | **string** | The type of the node. union TreeNodeUnion exclusion TreeNodeExclusion intersection TreeNodeIntersection leaf TreeNodeLeaf tuple_to_subject_set TreeNodeTupleToSubjectSet computed_subject_set TreeNodeComputedSubjectSet not TreeNodeNot unspecified TreeNodeUnspecified | 

## Methods

### NewExpandedPermissionTree

`func NewExpandedPermissionTree(type_ string, ) *ExpandedPermissionTree`

NewExpandedPermissionTree instantiates a new ExpandedPermissionTree object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpandedPermissionTreeWithDefaults

`func NewExpandedPermissionTreeWithDefaults() *ExpandedPermissionTree`

NewExpandedPermissionTreeWithDefaults instantiates a new ExpandedPermissionTree object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChildren

`func (o *ExpandedPermissionTree) GetChildren() []ExpandedPermissionTree`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *ExpandedPermissionTree) GetChildrenOk() (*[]ExpandedPermissionTree, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *ExpandedPermissionTree) SetChildren(v []ExpandedPermissionTree)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *ExpandedPermissionTree) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetTuple

`func (o *ExpandedPermissionTree) GetTuple() Relationship`

GetTuple returns the Tuple field if non-nil, zero value otherwise.

### GetTupleOk

`func (o *ExpandedPermissionTree) GetTupleOk() (*Relationship, bool)`

GetTupleOk returns a tuple with the Tuple field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTuple

`func (o *ExpandedPermissionTree) SetTuple(v Relationship)`

SetTuple sets Tuple field to given value.

### HasTuple

`func (o *ExpandedPermissionTree) HasTuple() bool`

HasTuple returns a boolean if a field has been set.

### GetType

`func (o *ExpandedPermissionTree) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExpandedPermissionTree) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExpandedPermissionTree) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


