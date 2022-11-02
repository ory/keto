// Copyright © 2022 Ory Corp

package ketoapi

import (
	"net/url"

	"github.com/ory/x/pointerx"
	"github.com/pkg/errors"
)

func (q *RelationQuery) FromURLQuery(query url.Values) (*RelationQuery, error) {
	if q == nil {
		q = &RelationQuery{}
	}

	if query.Has("subject") {
		return nil, ErrDroppedSubjectKey
	}

	// reset subject
	q.SubjectID = nil
	q.SubjectSet = nil

	switch {
	case !query.Has(SubjectIDKey) && !query.Has(SubjectSetNamespaceKey) && !query.Has(SubjectSetObjectKey) && !query.Has(SubjectSetRelationKey):
		// was not queried for the subject
	case query.Has(SubjectIDKey) && (query.Has(SubjectSetNamespaceKey) || query.Has(SubjectSetObjectKey) || query.Has(SubjectSetRelationKey)):
		return nil, ErrDuplicateSubject.WithDebugf("please provide either %s or all of %s, %s, and %s", SubjectIDKey, SubjectSetNamespaceKey, SubjectSetObjectKey, SubjectSetRelationKey)
	case query.Has(SubjectIDKey):
		q.SubjectID = pointerx.Ptr(query.Get(SubjectIDKey))
	case query.Has(SubjectSetNamespaceKey) && query.Has(SubjectSetObjectKey) && query.Has(SubjectSetRelationKey):
		q.SubjectSet = &SubjectSet{
			Namespace: query.Get(SubjectSetNamespaceKey),
			Object:    query.Get(SubjectSetObjectKey),
			Relation:  query.Get(SubjectSetRelationKey),
		}
	default:
		return nil, ErrIncompleteSubject
	}

	if query.Has("namespace") {
		q.Namespace = pointerx.Ptr(query.Get("namespace"))
	}
	if query.Has("object") {
		q.Object = pointerx.Ptr(query.Get("object"))
	}
	if query.Has("relation") {
		q.Relation = pointerx.Ptr(query.Get("relation"))
	}

	return q, nil
}

func (q *RelationQuery) ToURLQuery() url.Values {
	v := make(url.Values, 7)

	if q.Namespace != nil {
		v.Add("namespace", *q.Namespace)
	}
	if q.Relation != nil {
		v.Add("relation", *q.Relation)
	}
	if q.Object != nil {
		v.Add("object", *q.Object)
	}
	if q.SubjectID != nil {
		v.Add(SubjectIDKey, *q.SubjectID)
	} else if q.SubjectSet != nil {
		v.Add(SubjectSetNamespaceKey, q.SubjectSet.Namespace)
		v.Add(SubjectSetObjectKey, q.SubjectSet.Object)
		v.Add(SubjectSetRelationKey, q.SubjectSet.Relation)
	}

	return v
}

func (r *RelationTuple) FromURLQuery(query url.Values) (*RelationTuple, error) {
	q, err := (&RelationQuery{}).FromURLQuery(query)
	if err != nil {
		return nil, err
	}
	if q.SubjectID == nil && q.SubjectSet == nil {
		return nil, errors.WithStack(ErrNilSubject)
	}
	if q.Namespace == nil || q.Object == nil || q.Relation == nil {
		return nil, errors.WithStack(ErrIncompleteTuple)
	}

	r.Namespace = *q.Namespace
	r.Object = *q.Object
	r.Relation = *q.Relation
	r.SubjectID = q.SubjectID
	r.SubjectSet = q.SubjectSet

	return r, nil
}

func (r *RelationTuple) ToURLQuery() url.Values {
	return (&RelationQuery{
		Namespace:  &r.Namespace,
		Object:     &r.Object,
		Relation:   &r.Relation,
		SubjectID:  r.SubjectID,
		SubjectSet: r.SubjectSet,
	}).ToURLQuery()
}

func (s *SubjectSet) FromURLQuery(values url.Values) *SubjectSet {
	if s == nil {
		s = &SubjectSet{}
	}

	s.Namespace = values.Get("namespace")
	s.Relation = values.Get("relation")
	s.Object = values.Get("object")

	return s
}

func (s *SubjectSet) ToURLQuery() url.Values {
	return url.Values{
		"namespace": []string{s.Namespace},
		"object":    []string{s.Object},
		"relation":  []string{s.Relation},
	}
}
