package ast

import "encoding/json"

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
		Operation Operator `json:"operator"`
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

	// InvertResult inverts the check result of the child.
	InvertResult struct {
		Child Child `json:"inverted"`
	}
)

type Operator int

//go:generate stringer -type=Operator -linecomment
const (
	OperatorOr  Operator = iota // or
	OperatorAnd                 // and
)

func (o Operator) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.String())

}

func (r *UsersetRewrite) AsRewrite() *UsersetRewrite  { return r }
func (c *ComputedUserset) AsRewrite() *UsersetRewrite { return &UsersetRewrite{Children: []Child{c}} }
func (t *TupleToUserset) AsRewrite() *UsersetRewrite  { return &UsersetRewrite{Children: []Child{t}} }
func (i *InvertResult) AsRewrite() *UsersetRewrite    { return &UsersetRewrite{Children: []Child{i}} }
