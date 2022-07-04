package ast

type (
	Relation struct {
		Name           string          `json:"name"`
		Types          []RelationType  `json:"types,omitempty"`
		UsersetRewrite *UsersetRewrite `json:"rewrite,omitempty"`
	}

	RelationType struct {
		Namespace string `json:"namespace"`
		Relation  string `json:"relation,omitempty"` // optional
	}

	UsersetRewrite struct {
		Operation SetOperation `json:"set_operation"`
		Children  []Child      `json:"children"`
		child
	}

	Children = []Child

	// Define interface to restrict the child types of userset rewrites.
	Child interface {
		// AsRewrite returns the child as a userset rewrite, as relations
		// require a top-level rewrite, even if there just one child was parsed.
		AsRewrite() *UsersetRewrite
		onlyComputedUserSetOrTupleToUserset()
	}
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

//go:generate stringer -type=SetOperation -trimprefix=SetOperation
const (
	SetOperationUnion SetOperation = iota
	SetOperationIntersection
	SetOperationDifference
)

func (child) onlyComputedUserSetOrTupleToUserset() {}

func (r *UsersetRewrite) AsRewrite() *UsersetRewrite  { return r }
func (c *ComputedUserset) AsRewrite() *UsersetRewrite { return &UsersetRewrite{Children: []Child{c}} }
func (t *TupleToUserset) AsRewrite() *UsersetRewrite  { return &UsersetRewrite{Children: []Child{t}} }
