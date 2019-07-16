// Package ladon
package ladon

// swagger:parameters doOryAccessControlPoliciesAllow
type doOryAccessControlPoliciesAllow struct {
	// The ORY Access Control Policy flavor. Can be "regex", "glob", and "exact".
	//
	// in: path
	// required: true
	Flavor string `json:"flavor"`

	// in: body
	Body oryAccessControlPolicyAllowedInput
}

// Input for checking if a request is allowed or not.
//
// swagger:model oryAccessControlPolicyAllowedInput
type oryAccessControlPolicyAllowedInput struct {
	// Resource is the resource that access is requested to.
	Resource string `json:"resource"`

	// Action is the action that is requested on the resource.
	Action string `json:"action"`

	// Subject is the subject that is requesting access.
	Subject string `json:"subject"`

	// Context is the request's environmental context.
	Context map[string]interface{} `json:"context"`
}

// swagger:parameters upsertOryAccessControlPolicy
type upsertOryAccessControlPolicy struct {
	// The ORY Access Control Policy flavor. Can be "regex", "glob", and "exact".
	//
	// in: path
	// required: true
	Flavor string `json:"flavor"`

	// in: body
	Body oryAccessControlPolicy
}

// swagger:parameters listOryAccessControlPolicies
type listOryAccessControlPolicies struct {
	// The ORY Access Control Policy flavor. Can be "regex", "glob", and "exact"
	//
	// in: path
	// required: true
	Flavor string `json:"flavor"`

	// The maximum amount of policies returned.
	//
	// in: query
	Limit int `json:"limit"`

	// The offset from where to start looking.
	//
	// in: query
	Offset int `json:"offset"`
}

// swagger:parameters getOryAccessControlPolicy
type getOryAccessControlPolicy struct {
	// The ORY Access Control Policy flavor. Can be "regex", "glob", and "exact".
	//
	// in: path
	// required: true
	Flavor string `json:"flavor"`

	// The ID of the ORY Access Control Policy Role.
	//
	// in: path
	// required: true
	ID string `json:"id"`
}

// swagger:parameters deleteOryAccessControlPolicy
type deleteOryAccessControlPolicy struct {
	// The ORY Access Control Policy flavor. Can be "regex", "glob", and "exact".
	//
	// in: path
	// required: true
	Flavor string `json:"flavor"`

	// The ID of the ORY Access Control Policy Role.
	//
	// in: path
	// required: true
	ID string `json:"id"`
}

// swagger:parameters getOryAccessControlPolicyRole
type getOryAccessControlPolicyRole struct {
	// The ORY Access Control Policy flavor. Can be "regex", "glob", and "exact".
	//
	// in: path
	// required: true
	Flavor string `json:"flavor"`

	// The ID of the ORY Access Control Policy Role.
	//
	// in: path
	// required: true
	ID string `json:"id"`
}

// swagger:parameters deleteOryAccessControlPolicyRole
type deleteOryAccessControlPolicyRole struct {
	// The ORY Access Control Policy flavor. Can be "regex", "glob", and "exact".
	//
	// in: path
	// required: true
	Flavor string `json:"flavor"`

	// The ID of the ORY Access Control Policy Role.
	// in: path
	// required: true
	ID string `json:"id"`
}

// swagger:parameters upsertOryAccessControlPolicyRole
type upsertOryAccessControlPolicyRole struct {
	// The ORY Access Control Policy flavor. Can be "regex", "glob", and "exact".
	//
	// in: path
	// required: true
	Flavor string `json:"flavor"`

	// in: body
	Body oryAccessControlPolicyRole
}

// swagger:model addOryAccessControlPolicyRoleMembersBody
type addOryAccessControlPolicyRoleMembersBody struct {
	// The members to be added.
	Members []string `json:"members"`
}

// swagger:parameters addOryAccessControlPolicyRoleMembers
type addOryAccessControlPolicyRoleMembers struct {
	// The ORY Access Control Policy flavor. Can be "regex", "glob", and "exact".
	//
	// in: path
	// required: true
	Flavor string `json:"flavor"`

	// The ID of the ORY Access Control Policy Role.
	//
	// in: path
	// required: true
	ID string `json:"id"`

	// in: body
	Body addOryAccessControlPolicyRoleMembersBody
}

// swagger:parameters removeOryAccessControlPolicyRoleMembers
type removeOryAccessControlPolicyRoleMembers struct {
	// The ORY Access Control Policy flavor. Can be "regex", "glob", and "exact".
	//
	// in: path
	// required: true
	Flavor string `json:"flavor"`

	// The ID of the ORY Access Control Policy Role.
	//
	// in: path
	// required: true
	ID string `json:"id"`

	// The member to be removed.
	//
	// in: path
	// required: true
	Member string `json:"member"`
}

// Policies is an array of policies.
//
// swagger:response oryAccessControlPolicies
type oryAccessControlPolicies struct {
	// The request body.
	//
	// in: body
	// type: array
	Body []oryAccessControlPolicy
}

// Roles is an array of roles.
//
// swagger:response oryAccessControlPolicyRoles
type oryAccessControlPolicyRoles struct {
	// The request body.
	//
	// in: body
	// type: array
	Body []oryAccessControlPolicyRole
}

// oryAccessControlPolicyRole represents a group of users that share the same role. A role could be an administrator, a moderator, a regular
// user or some other sort of role.
//
// swagger:model oryAccessControlPolicyRole
type oryAccessControlPolicyRole struct {
	// ID is the role's unique id.
	ID string `json:"id"`

	// Members is who belongs to the role.
	Members []string `json:"members"`
}

// oryAccessControlPolicy specifies an ORY Access Policy document.
//
// swagger:model oryAccessControlPolicy
type oryAccessControlPolicy struct {
	// ID is the unique identifier of the ORY Access Policy. It is used to query, update, and remove the ORY Access Policy.
	ID string `json:"id"`

	// Description is an optional, human-readable description.
	Description string `json:"description"`

	// Subjects is an array representing all the subjects this ORY Access Policy applies to.
	Subjects []string `json:"subjects"`

	// Resources is an array representing all the resources this ORY Access Policy applies to.
	Resources []string `json:"resources"`

	// Actions is an array representing all the actions this ORY Access Policy applies to.
	Actions []string `json:"actions"`

	// Effect is the effect of this ORY Access Policy. It can be "allow" or "deny".
	Effect string `json:"effect"`

	// Conditions represents a keyed object of conditions under which this ORY Access Policy is active.
	Conditions map[string]interface{} `json:"conditions"`
}

// swagger:parameters listOryAccessControlPolicyRoles
type listOryAccessControlPolicyRoles struct {
	// The ORY Access Control Policy flavor. Can be "regex", "glob", and "exact"
	//
	// in: path
	// required: true
	Flavor string `json:"flavor"`

	// The Member (ID) for which the roles are to be listed (Optional).
	// in: query
	// required: false
	Member string `json:"member"`

	// The maximum amount of policies returned.
	//
	// in: query
	Limit int `json:"limit"`

	// The offset from where to start looking.
	//
	// in: query
	Offset int `json:"offset"`
}
