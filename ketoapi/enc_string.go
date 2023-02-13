// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package ketoapi

import (
	"fmt"
	"strings"

	"github.com/ory/herodot"
	"github.com/pkg/errors"
)

var ErrMalformedInput = herodot.ErrBadRequest.WithError("malformed string input")

func (r *RelationTuple) String() string {
	if r == nil {
		return ""
	}
	sb := strings.Builder{}
	sb.WriteString(r.Namespace)
	sb.WriteRune(':')
	sb.WriteString(r.Object)
	sb.WriteRune('#')
	sb.WriteString(r.Relation)
	sb.WriteRune('@')

	if r.SubjectID != nil {
		sb.WriteString(*r.SubjectID)
	} else if r.SubjectSet != nil {
		sb.WriteString(r.SubjectSet.String())
	} else {
		sb.WriteString("<ERROR: no subject>")
	}
	return sb.String()
}

func (r *RelationTuple) FromString(s string) (*RelationTuple, error) {
	var (
		objectAndRelationAndSubject string
		relationAndSubject          string
		subject                     string
		ok                          bool
	)
	if r.Namespace, objectAndRelationAndSubject, ok = strings.Cut(s, ":"); !ok {
		return nil, errors.WithStack(ErrMalformedInput.WithDebug("expected input to contain ':'"))
	}

	if r.Object, relationAndSubject, ok = strings.Cut(objectAndRelationAndSubject, "#"); !ok {
		return nil, errors.WithStack(ErrMalformedInput.WithDebug("expected input to contain '#'"))
	}

	if r.Relation, subject, ok = strings.Cut(relationAndSubject, "@"); !ok {
		return nil, errors.WithStack(ErrMalformedInput.WithDebug("expected input to contain '@'"))
	}

	// remove optional brackets around the subject set
	subject = strings.Trim(subject, "()")
	if strings.Contains(subject, ":") {
		subSet, err := (&SubjectSet{}).FromString(subject)
		if err != nil {
			return nil, err
		}
		r.SubjectSet = subSet
	} else {
		r.SubjectID = &subject
	}

	return r, nil
}

func (s *SubjectSet) String() string {
	if s.Relation == "" {
		return fmt.Sprintf("%s:%s", s.Namespace, s.Object)
	}
	return fmt.Sprintf("%s:%s#%s", s.Namespace, s.Object, s.Relation)
}

func (s *SubjectSet) FromString(str string) (*SubjectSet, error) {
	// If there is no '#' we have a subject set without a relation, such as
	// Users:Bob, which just means that the relation is empty.
	namespaceAndObject, relation, _ := strings.Cut(str, "#")

	namespace, object, ok := strings.Cut(namespaceAndObject, ":")
	if !ok {
		return nil, errors.WithStack(ErrMalformedInput.WithDebug("expected subject set to contain ':'"))
	}

	return &SubjectSet{
		Namespace: namespace,
		Object:    object,
		Relation:  relation,
	}, nil
}

func (t TreeNodeType) String() string {
	return string(t)
}

func (t *Tree[NodeT]) Label() string {
	if t == nil {
		return ""
	}

	return t.Tuple.String()
}

func (t *Tree[NodeT]) String() string {
	if t == nil {
		return ""
	}

	nodeLabel := t.Label()

	if t.Type == TreeNodeLeaf {
		return fmt.Sprintf("∋ %s️", nodeLabel)
	}

	children := make([]string, len(t.Children))
	for i, c := range t.Children {
		var indent string
		if i == len(t.Children)-1 {
			indent = "   "
		} else {
			indent = "│  "
		}
		children[i] = strings.Join(strings.Split(c.String(), "\n"), "\n"+indent)
	}

	setOperation := ""
	switch t.Type {
	case TreeNodeIntersection:
		setOperation = "and"
	case TreeNodeUnion:
		setOperation = "or"
	case TreeNodeExclusion:
		setOperation = `\`
	case TreeNodeNot:
		setOperation = `not`
	case TreeNodeTupleToSubjectSet:
		setOperation = "┐ tuple to userset"
	case TreeNodeComputedSubjectSet:
		setOperation = "┐ computed userset"
	}

	boxSymbol := "├"
	if len(children) == 1 {
		boxSymbol = "└"
	}
	return fmt.Sprintf("%s %s\n%s──%s", setOperation, nodeLabel, boxSymbol, strings.Join(children, "\n└──"))
}
