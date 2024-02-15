# ExpandedPermissionTree

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | Pointer to [**[]ExpandedPermissionTree**](ExpandedPermissionTree.md) | The children of this node.  This is never set if &#x60;node_type&#x60; &#x3D;&#x3D; &#x60;NODE_TYPE_LEAF&#x60;. | [optional] 
**Subject** | Pointer to [**OryKetoRelationTuplesV1alpha2Subject**](OryKetoRelationTuplesV1alpha2Subject.md) |  | [optional] 
**Tuple** | Pointer to [**Relationship**](Relationship.md) |  | [optional] 
**Type** | [**OryKetoRelationTuplesV1alpha2NodeType**](OryKetoRelationTuplesV1alpha2NodeType.md) |  | [default to ORYKETORELATIONTUPLESV1ALPHA2NODETYPE_UNSPECIFIED]

## Methods

### NewExpandedPermissionTree

`func NewExpandedPermissionTree(type_ OryKetoRelationTuplesV1alpha2NodeType, ) *ExpandedPermissionTree`

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

### GetSubject

`func (o *ExpandedPermissionTree) GetSubject() OryKetoRelationTuplesV1alpha2Subject`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *ExpandedPermissionTree) GetSubjectOk() (*OryKetoRelationTuplesV1alpha2Subject, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *ExpandedPermissionTree) SetSubject(v OryKetoRelationTuplesV1alpha2Subject)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *ExpandedPermissionTree) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

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

`func (o *ExpandedPermissionTree) GetType() OryKetoRelationTuplesV1alpha2NodeType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExpandedPermissionTree) GetTypeOk() (*OryKetoRelationTuplesV1alpha2NodeType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExpandedPermissionTree) SetType(v OryKetoRelationTuplesV1alpha2NodeType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


