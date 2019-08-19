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

func (r *Role) withMembers(members []string) *Role {
	if r == nil || len(members) == 0 || contains(members[0], r.Members) {
		return r
	}
	return nil
}

func (r *Role) withIDs(ids []string) *Role {
	if r == nil || len(ids) == 0 || contains(r.ID, ids) {
		return r
	}
	return nil
}
