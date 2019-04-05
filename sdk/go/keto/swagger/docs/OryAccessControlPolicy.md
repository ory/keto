# OryAccessControlPolicy

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | **[]string** | Actions is an array representing all the actions this ORY Access Policy applies to. | [optional] [default to null]
**Conditions** | [**map[string]interface{}**](interface{}.md) | Conditions represents a keyed object of conditions under which this ORY Access Policy is active. | [optional] [default to null]
**Description** | **string** | Description is an optional, human-readable description. | [optional] [default to null]
**Effect** | **string** | Effect is the effect of this ORY Access Policy. It can be \&quot;allow\&quot; or \&quot;deny\&quot;. | [optional] [default to null]
**Id** | **string** | ID is the unique identifier of the ORY Access Policy. It is used to query, update, and remove the ORY Access Policy. | [optional] [default to null]
**Resources** | **[]string** | Resources is an array representing all the resources this ORY Access Policy applies to. | [optional] [default to null]
**Subjects** | **[]string** | Subjects is an array representing all the subjects this ORY Access Policy applies to. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


