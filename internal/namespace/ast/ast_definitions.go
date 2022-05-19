package ast

type (
	Relation struct {
		Name           string
		UsersetRewrite *UsersetRewrite
	}

	UsersetRewrite struct {
		Operation SetOperation
		Children  Children
	}

	Children struct {
		ComputedUsersets []ComputedUserset
		TupleToUsersets  []TupleToUserset
	}

	ComputedUserset struct {
		Relation string
	}

	TupleToUserset struct {
		Relation                string
		ComputedUsersetRelation string
	}
)

type SetOperation int

const (
	SetOperationUnion SetOperation = iota
	SetOperationIntersection
	SetOperationExclusion
)
