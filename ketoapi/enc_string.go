package ketoapi

import (
	"fmt"
	"strings"

	"github.com/ory/herodot"
	"github.com/pkg/errors"
)

var ErrMalformedInput = herodot.ErrBadRequest.WithError("malformed string input")

func (r *RelationTuple) String() string {
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
		sb.WriteRune('(')
		sb.WriteString(r.SubjectSet.Namespace)
		sb.WriteRune(':')
		sb.WriteString(r.SubjectSet.Object)
		sb.WriteRune('#')
		sb.WriteString(r.SubjectSet.Relation)
		sb.WriteRune(')')
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
	if strings.Contains(subject, "#") {
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
	return fmt.Sprintf("%s:%s#%s", s.Namespace, s.Object, s.Relation)
}

func (s *SubjectSet) FromString(str string) (*SubjectSet, error) {
	namespaceAndObject, relation, ok := strings.Cut(str, "#")
	if !ok {
		return nil, errors.WithStack(ErrMalformedInput.WithDebug("expected subject set to contain '#'"))
	}

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

func (t *ExpandTree) String() string {
	if t == nil {
		return ""
	}

	sub := "<!--no subject-->"
	switch {
	case t.SubjectID != nil:
		sub = *t.SubjectID
	case t.SubjectSet != nil:
		sub = t.SubjectSet.String()
	}

	if t.Type == Leaf {
		return fmt.Sprintf("☘ %s️", sub)
	}

	children := make([]string, len(t.Children))
	for i, c := range t.Children {
		children[i] = strings.Join(strings.Split(c.String(), "\n"), "\n│  ")
	}

	return fmt.Sprintf("∪ %s\n├─ %s", sub, strings.Join(children, "\n├─ "))
}
