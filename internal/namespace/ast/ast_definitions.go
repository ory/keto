// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package ast

import "encoding/json"

type (
	Relation struct {
		Name              string             `json:"name"`
		Types             []RelationType     `json:"types,omitempty"`
		SubjectSetRewrite *SubjectSetRewrite `json:"rewrite,omitempty"`
	}

	RelationType struct {
		Namespace string `json:"namespace"`
		Relation  string `json:"relation,omitempty"` // optional
	}

	SubjectSetRewrite struct {
		Operation Operator `json:"operator"`
		Children  Children `json:"children"`
	}

	Children = []Child

	// Child are all possible types of subject-set rewrites.
	Child interface {
		// AsRewrite returns the child as a subject-set rewrite, as relations
		// require a top-level rewrite, even if just one child was parsed.
		AsRewrite() *SubjectSetRewrite
	}

	ComputedSubjectSet struct {
		Relation string `json:"relation"`
	}

	TupleToSubjectSet struct {
		Relation                   string `json:"relation"`
		ComputedSubjectSetRelation string `json:"computed_subject_set_relation"`
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

func (i Operator) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())

}

func (r *SubjectSetRewrite) AsRewrite() *SubjectSetRewrite { return r }
func (c *ComputedSubjectSet) AsRewrite() *SubjectSetRewrite {
	return &SubjectSetRewrite{Children: []Child{c}}
}
func (t *TupleToSubjectSet) AsRewrite() *SubjectSetRewrite {
	return &SubjectSetRewrite{Children: []Child{t}}
}
func (i *InvertResult) AsRewrite() *SubjectSetRewrite {
	return &SubjectSetRewrite{Children: []Child{i}}
}
