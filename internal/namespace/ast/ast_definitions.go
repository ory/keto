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
		Operation Operator `json:"set_operation"`
		Children  []Child  `json:"children"`
	}

	Children = []Child

	// Define interface to restrict the child types of userset rewrites.
	Child interface {
		// AsRewrite returns the child as a userset rewrite, as relations
		// require a top-level rewrite, even if just one child was parsed.
		AsRewrite() *UsersetRewrite
	}

	ComputedUserset struct {
		Relation string `json:"relation"`
	}

	TupleToUserset struct {
		Relation                string `json:"relation"`
		ComputedUsersetRelation string `json:"computed_userset_relation"`
	}
)

type Operator int

//go:generate stringer -type=Operator -trimprefix=Operator
const (
	OperatorOr Operator = iota
	OperatorAnd
	OperatorNot
)

func (r *UsersetRewrite) AsRewrite() *UsersetRewrite  { return r }
func (c *ComputedUserset) AsRewrite() *UsersetRewrite { return &UsersetRewrite{Children: []Child{c}} }
func (t *TupleToUserset) AsRewrite() *UsersetRewrite  { return &UsersetRewrite{Children: []Child{t}} }
