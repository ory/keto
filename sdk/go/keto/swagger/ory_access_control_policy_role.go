/* 
 * Package main ORY Keto
 *
 * OpenAPI spec version: Latest
 * Contact: hi@ory.sh
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package swagger

// Role represents a group of users that share the same role. A role could be an administrator, a moderator, a regular user or some other sort of role.
type OryAccessControlPolicyRole struct {

	// ID is the role's unique id.
	Id string `json:"id,omitempty"`

	// Members is who belongs to the role.
	Members []string `json:"members,omitempty"`
}
