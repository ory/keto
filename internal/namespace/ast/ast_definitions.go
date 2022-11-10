// Copyright © 2022 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package ast

import (
	"context"
	"encoding/json"
)

type (
	Relation struct {
		Name              string             `json:"name"`
		Types             []RelationType     `json:"types,omitempty"`
		SubjectSetRewrite *SubjectSetRewrite `json:"rewrite,omitempty"`
		Params            []string           `json:"params,omitempty"`
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

	Arg interface {
		Value(ctx context.Context) string
	}

	ComputedSubjectSet struct {
		Relation string `json:"relation"`
		Args     []Arg  `json:"args,omitempty"`
	}

	TupleToSubjectSet struct {
		Relation                   string `json:"relation"`
		ComputedSubjectSetRelation string `json:"computed_subject_set_relation"`
		Args                       []Arg  `json:"args,omitempty"`
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

// concrete argument types

type NamedArg string

func (p NamedArg) Value(ctx context.Context) string {
	return string(p)
}

type StringLiteralArg string

func (p StringLiteralArg) Value(ctx context.Context) string {
	return string(p)
}

var ContextArg = contextArg(0)
var CtxSubjectArg = contextArg(1)

type contextArg int

func (p contextArg) Value(ctx context.Context) string {
	panic("should not reach here")
}
