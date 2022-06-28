package ast

type (
	Relation struct {
		Name           string          `json:"name"`
		UsersetRewrite *UsersetRewrite `json:"rewrite,omitempty"`
	}

	UsersetRewrite struct {
		Operation SetOperation `json:"set_operation"`
		Children  []Child      `json:"children"`
	}

	Children = []Child

	// Define interface to restrict the child types of userset rewrites.
	Child interface{ onlyComputedUserSetOrTupleToUserset() }
	child struct{}

	ComputedUserset struct {
		Relation string `json:"relation"`
		child
	}

	TupleToUserset struct {
		Relation                string `json:"relation"`
		ComputedUsersetRelation string `json:"computed_userset_relation"`
		child
	}
)

type SetOperation int

const (
	SetOperationUnion SetOperation = iota
	SetOperationIntersection
	SetOperationDifference
)

func (child) onlyComputedUserSetOrTupleToUserset() {}
