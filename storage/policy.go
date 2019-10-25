package storage

// Policies is an array of policies.
//
// swagger:ignore
type Policies []Policy

// Policy specifies an ORY Access Policy document.
//
// swagger:ignore
type Policy struct {
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

func (p *Policy) withSubjects(subjects []string) *Policy {
	if p == nil || len(subjects) == 0 || containsAll(subjects, p.Subjects) {
		return p
	}
	return nil
}

func (p *Policy) withResources(resources []string) *Policy {
	if p == nil || len(resources) == 0 || containsAll(resources, p.Resources) {
		return p
	}
	return nil
}

func (p *Policy) withActions(actions []string) *Policy {
	if p == nil || len(actions) == 0 || containsAll(actions, p.Actions) {
		return p
	}
	return nil
}

func (p *Policy) withIDs(ids []string) *Policy {
	if p == nil || len(ids) == 0 || contains(p.ID, ids) {
		return p
	}
	return nil
}
