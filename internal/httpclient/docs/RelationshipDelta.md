# RelationshipDelta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | [**RelationshipDeltaAction**](RelationshipDeltaAction.md) |  | [default to RELATIONSHIPDELTAACTION_ACTION_UNSPECIFIED]
**RelationTuple** | [**Relationship**](Relationship.md) |  | 

## Methods

### NewRelationshipDelta

`func NewRelationshipDelta(action RelationshipDeltaAction, relationTuple Relationship, ) *RelationshipDelta`

NewRelationshipDelta instantiates a new RelationshipDelta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelationshipDeltaWithDefaults

`func NewRelationshipDeltaWithDefaults() *RelationshipDelta`

NewRelationshipDeltaWithDefaults instantiates a new RelationshipDelta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *RelationshipDelta) GetAction() RelationshipDeltaAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *RelationshipDelta) GetActionOk() (*RelationshipDeltaAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *RelationshipDelta) SetAction(v RelationshipDeltaAction)`

SetAction sets Action field to given value.


### GetRelationTuple

`func (o *RelationshipDelta) GetRelationTuple() Relationship`

GetRelationTuple returns the RelationTuple field if non-nil, zero value otherwise.

### GetRelationTupleOk

`func (o *RelationshipDelta) GetRelationTupleOk() (*Relationship, bool)`

GetRelationTupleOk returns a tuple with the RelationTuple field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationTuple

`func (o *RelationshipDelta) SetRelationTuple(v Relationship)`

SetRelationTuple sets RelationTuple field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


