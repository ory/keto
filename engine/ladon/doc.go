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
	Body Input
}

// swagger:parameters upsertOryAccessControlPolicy
type upsertOryAccessControlPolicy struct {
	// The ORY Access Control Policy flavor. Can be "regex", "glob", and "exact".
	//
	// in: path
	// required: true
	Flavor string `json:"flavor"`

	// in: body
	Body Policy
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
	Body Role
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

// swagger:model removeOryAccessControlPolicyRoleMembersBody
type removeOryAccessControlPolicyRoleMembersBody struct {
	// The members to be removed.
	Members []string `json:"members"`
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

	// in: body
	Body removeOryAccessControlPolicyRoleMembersBody
}

// Policies is an array of policies.
//
// swagger:response oryAccessControlPolicies
type oryAccessControlPolicies struct {
	// in: body
	Body []Policy
}

// Roles is an array of roles.
//
// swagger:response oryAccessControlPolicyRoles
type oryAccessControlPolicyRoles struct {
	// in: body
	Body []Role
}

// swagger:parameters listOryAccessControlPolicyRoles
type listOryAccessControlPolicyRoles struct {
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
