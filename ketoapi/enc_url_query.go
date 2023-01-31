// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

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

	if query.Has(NamespaceKey) {
		q.Namespace = pointerx.Ptr(query.Get(NamespaceKey))
	}
	if query.Has(ObjectKey) {
		q.Object = pointerx.Ptr(query.Get(ObjectKey))
	}
	if query.Has(RelationKey) {
		q.Relation = pointerx.Ptr(query.Get(RelationKey))
	}

	return q, nil
}

func (q *RelationQuery) ToURLQuery() url.Values {
	v := make(url.Values, 7)

	if q.Namespace != nil {
		v.Add(NamespaceKey, *q.Namespace)
	}
	if q.Relation != nil {
		v.Add(RelationKey, *q.Relation)
	}
	if q.Object != nil {
		v.Add(ObjectKey, *q.Object)
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

	s.Namespace = values.Get(NamespaceKey)
	s.Relation = values.Get(RelationKey)
	s.Object = values.Get(ObjectKey)

	return s
}

func (s *SubjectSet) ToURLQuery() url.Values {
	return url.Values{
		NamespaceKey: []string{s.Namespace},
		ObjectKey:    []string{s.Object},
		RelationKey:  []string{s.Relation},
	}
}
