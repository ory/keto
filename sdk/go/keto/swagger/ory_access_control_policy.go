/*
 * Package main ORY Keto
 *
 * OpenAPI spec version: Latest
 * Contact: hi@ory.sh
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package swagger

type OryAccessControlPolicy struct {

	// Actions is an array representing all the actions this ORY Access Policy applies to.
	Actions []string `json:"actions,omitempty"`

	// Conditions represents a keyed object of conditions under which this ORY Access Policy is active.
	Conditions map[string]interface{} `json:"conditions,omitempty"`

	// Description is an optional, human-readable description.
	Description string `json:"description,omitempty"`

	// Effect is the effect of this ORY Access Policy. It can be \"allow\" or \"deny\".
	Effect string `json:"effect,omitempty"`

	// ID is the unique identifier of the ORY Access Policy. It is used to query, update, and remove the ORY Access Policy.
	Id string `json:"id,omitempty"`

	// Resources is an array representing all the resources this ORY Access Policy applies to.
	Resources []string `json:"resources,omitempty"`

	// Subjects is an array representing all the subjects this ORY Access Policy applies to.
	Subjects []string `json:"subjects,omitempty"`
}
