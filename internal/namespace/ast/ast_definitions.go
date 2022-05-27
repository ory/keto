package ast

type (
	Relation struct {
		Name           string
		UsersetRewrite *UsersetRewrite
	}

	UsersetRewrite struct {
		Operation SetOperation
		Children  []Child
	}

	Children = []Child

	// Define interface to restrict the child types of userset rewrites.
	Child interface{ onlyComputedUserSetOrTupleToUserset() }
	child struct{}

	ComputedUserset struct {
		Relation string
		child
	}

	TupleToUserset struct {
		Relation                string
		ComputedUsersetRelation string
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
