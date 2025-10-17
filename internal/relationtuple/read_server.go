// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package relationtuple

import (
	"context"

	"github.com/ory/herodot"
	keysetpagination "github.com/ory/x/pagination/keysetpagination_v2"
	"github.com/ory/x/pointerx"

	"github.com/ory/keto/ketoapi"
	rts "github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2"
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
	case req.Query != nil: //nolint:staticcheck //lint:ignore SA1019 backwards compatibility
		q.FromDataProvider(&deprecatedQueryWrapper{req.Query}) //nolint:staticcheck //lint:ignore SA1019 backwards compatibility
	default:
		q.FromDataProvider(&openAPIQueryWrapper{req})
	}

	iq, err := h.d.ReadOnlyMapper().FromQuery(ctx, &q)
	if err != nil {
		return nil, err
	}

	paginationKeys := h.d.Config(ctx).PaginationEncryptionKeys()
	pageOpts := make([]keysetpagination.Option, 0, 2)
	if req.PageSize > 0 {
		pageOpts = append(pageOpts, keysetpagination.WithSize(int(req.PageSize)))
	}
	if req.PageToken != "" {
		token, err := keysetpagination.ParsePageToken(paginationKeys, req.PageToken)
		if err != nil {
			return nil, herodot.ErrBadRequest.WithError(err.Error())
		}
		pageOpts = append(pageOpts, keysetpagination.WithToken(token))
	}
	ir, nextPage, err := h.d.RelationTupleManager().GetRelationTuples(ctx, iq, pageOpts...)
	if err != nil {
		return nil, err
	}
	relations, err := h.d.ReadOnlyMapper().ToTuple(ctx, ir...)
	if err != nil {
		return nil, err
	}

	resp := &rts.ListRelationTuplesResponse{
		RelationTuples: make([]*rts.RelationTuple, len(ir)),
	}
	if !nextPage.IsLast() {
		resp.NextPageToken = nextPage.PageToken().Encrypt(paginationKeys)
	}
	for i, r := range relations {
		resp.RelationTuples[i] = r.ToProto()
	}

	return resp, nil
}
