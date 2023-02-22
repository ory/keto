// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package relationtuple

import (
	"context"

	"github.com/ory/x/pointerx"

	"github.com/ory/keto/ketoapi"

	rts "github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2"

	"github.com/ory/keto/internal/x"
)

var (
	_ rts.ReadServiceServer = (*handler)(nil)
)

type (
	queryWrapper struct {
		*rts.RelationQuery
	}
	deprecatedQueryWrapper struct {
		*rts.ListRelationTuplesRequest_Query
	}
	openAPIQueryWrapper struct {
		wrapped interface {
			GetObject() string
			GetRelation() string
			GetNamespace() string
			GetSubjectSet() *rts.SubjectSetQuery
			GetSubjectId() string
		}
	}
)

func stringPtr(s string) *string {
	if s == "" {
		return nil
	}
	return pointerx.Ptr(s)
}

func (q *queryWrapper) GetObject() *string    { return q.Object }
func (q *queryWrapper) GetNamespace() *string { return q.Namespace }
func (q *queryWrapper) GetRelation() *string  { return q.Relation }

func (q *deprecatedQueryWrapper) GetObject() *string    { return stringPtr(q.Object) }
func (q *deprecatedQueryWrapper) GetNamespace() *string { return stringPtr(q.Namespace) }
func (q *deprecatedQueryWrapper) GetRelation() *string  { return stringPtr(q.Relation) }

func (q *openAPIQueryWrapper) GetObject() *string { return stringPtr(q.wrapped.GetObject()) }
func (q *openAPIQueryWrapper) GetNamespace() *string {
	return stringPtr(q.wrapped.GetNamespace())
}
func (q *openAPIQueryWrapper) GetRelation() *string { return stringPtr(q.wrapped.GetRelation()) }
func (q *openAPIQueryWrapper) GetSubject() *rts.Subject {
	if set := q.wrapped.GetSubjectSet(); set != nil {
		return rts.NewSubjectSet(set.Namespace, set.Object, set.Relation)
	}
	if sID := q.wrapped.GetSubjectId(); sID != "" {
		return rts.NewSubjectID(q.wrapped.GetSubjectId())
	}
	return nil
}

func (h *handler) ListRelationTuples(ctx context.Context, req *rts.ListRelationTuplesRequest) (*rts.ListRelationTuplesResponse, error) {
	var q ketoapi.RelationQuery

	switch {
	case req.RelationQuery != nil:
		q.FromDataProvider(&queryWrapper{req.RelationQuery})
	case req.Query != nil: // nolint
		q.FromDataProvider(&deprecatedQueryWrapper{req.Query}) // nolint
	default:
		q.FromDataProvider(&openAPIQueryWrapper{req})
	}

	iq, err := h.d.ReadOnlyMapper().FromQuery(ctx, &q)
	if err != nil {
		return nil, err
	}
	ir, nextPage, err := h.d.RelationTupleManager().GetRelationTuples(ctx, iq,
		x.WithSize(int(req.PageSize)),
		x.WithToken(req.PageToken),
	)
	if err != nil {
		return nil, err
	}
	relations, err := h.d.ReadOnlyMapper().ToTuple(ctx, ir...)
	if err != nil {
		return nil, err
	}

	resp := &rts.ListRelationTuplesResponse{
		RelationTuples: make([]*rts.RelationTuple, len(ir)),
		NextPageToken:  nextPage,
	}
	for i, r := range relations {
		resp.RelationTuples[i] = r.ToProto()
	}

	return resp, nil
}
