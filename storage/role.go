package storage

// A list of roles.
//
// swagger:ignore
type Roles []Role

// Role represents a group of users that share the same role. A role could be an administrator, a moderator, a regular
// user or some other sort of role.
//
// swagger:ignore
type Role struct {
	// ID is the role's unique id.
	ID string `json:"id"`

	// Members is who belongs to the role.
	Members []string `json:"members"`
}

func (r Role) withMembers(members []string) *Role {
	res := common(members, r.Members)
	return &res
}

func (r Role) withID(id string) *Role {
	if r.ID == id {
		return &r
	}
	return &Role{}
}
