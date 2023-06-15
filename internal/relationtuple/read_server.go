// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package relationtuple

import (
	"context"
	"net/http"
	"strconv"

	"github.com/ory/x/pointerx"

	"github.com/ory/keto/ketoapi"

	rts "github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2"

	"github.com/ory/herodot"

	"github.com/julienschmidt/httprouter"

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
)

func (q *queryWrapper) GetObject() *string {
	return q.Object
}

func (q *queryWrapper) GetNamespace() *string {
	return q.Namespace
}

func (q *queryWrapper) GetRelation() *string {
	return q.Relation
}

func (q *deprecatedQueryWrapper) GetObject() *string {
	if q.Object == "" {
		return nil
	}
	return pointerx.Ptr(q.Object)
}

func (q *deprecatedQueryWrapper) GetNamespace() *string {
	if q.Namespace == "" {
		return nil
	}
	return pointerx.Ptr(q.Namespace)
}

func (q *deprecatedQueryWrapper) GetRelation() *string {
	if q.Relation == "" {
		return nil
	}
	return pointerx.Ptr(q.Relation)
}

func (h *handler) ListRelationTuples(ctx context.Context, req *rts.ListRelationTuplesRequest) (*rts.ListRelationTuplesResponse, error) {
	var q ketoapi.RelationQuery

	switch {
	case req.RelationQuery != nil:
		q.FromDataProvider(&queryWrapper{req.RelationQuery})
	case req.Query != nil: //nolint:staticcheck //lint:ignore SA1019 backwards compatibility
		q.FromDataProvider(&deprecatedQueryWrapper{req.Query}) //nolint:staticcheck //lint:ignore SA1019 backwards compatibility
	default:
		return nil, herodot.ErrBadRequest.WithError("you must provide a query")
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

// swagger:route GET /relation-tuples relationship getRelationships
//
// # Query relationships
//
// Get all relationships that match the query. Only the namespace field is required.
//
//	Consumes:
//	-  application/x-www-form-urlencoded
//
//	Produces:
//	- application/json
//
//	Schemes: http, https
//
//	Responses:
//	  200: relationships
//	  404: errorGeneric
//	  default: errorGeneric
func (h *handler) getRelations(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx := r.Context()

	q := r.URL.Query()
	query, err := (&ketoapi.RelationQuery{}).FromURLQuery(q)
	if err != nil {
		h.d.Writer().WriteError(w, r, herodot.ErrBadRequest.WithError(err.Error()))
		return
	}

	l := h.d.Logger()
	for k := range q {
		l = l.WithField(k, q.Get(k))
	}
	l.Debug("querying relationships")

	var paginationOpts []x.PaginationOptionSetter
	if pageToken := q.Get("page_token"); pageToken != "" {
		paginationOpts = append(paginationOpts, x.WithToken(pageToken))
	}

	if pageSize := q.Get("page_size"); pageSize != "" {
		s, err := strconv.ParseInt(pageSize, 0, 0)
		if err != nil {
			h.d.Writer().WriteError(w, r, herodot.ErrBadRequest.WithError(err.Error()))
			return
		}
		paginationOpts = append(paginationOpts, x.WithSize(int(s)))
	}

	iq, err := h.d.ReadOnlyMapper().FromQuery(ctx, query)
	if err != nil {
		h.d.Writer().WriteError(w, r, err)
		return
	}
	ir, nextPage, err := h.d.RelationTupleManager().GetRelationTuples(ctx, iq, paginationOpts...)
	if err != nil {
		h.d.Writer().WriteError(w, r, err)
		return
	}

	relations, err := h.d.ReadOnlyMapper().ToTuple(ctx, ir...)
	if err != nil {
		h.d.Writer().WriteError(w, r, err)
		return
	}

	resp := &ketoapi.GetResponse{
		RelationTuples: relations,
		NextPageToken:  nextPage,
	}

	h.d.Writer().Write(w, r, resp)
}
