// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package relationtuple

import (
	"context"
	"encoding/json"
	"net/http"

	"connectrpc.com/connect"
	"github.com/pkg/errors"
	"go.opentelemetry.io/otel/trace"

	"github.com/ory/herodot"

	rts "github.com/ory/keto/gen/go/ory/keto/relation_tuples/v1alpha2"
	"github.com/ory/keto/gen/go/ory/keto/relation_tuples/v1alpha2/relationtuplesconnect"
	"github.com/ory/keto/internal/x/validate"
	"github.com/ory/keto/ketoapi"
	"github.com/ory/keto/x/events"
)

var _ relationtuplesconnect.WriteServiceHandler = (*WriteHandler)(nil)

func protoTuplesWithAction(deltas []*rts.RelationTupleDelta, action rts.RelationTupleDelta_Action) (filtered []*ketoapi.RelationTuple, err error) {
	for _, d := range deltas {
		if d.Action == action {
			it, err := (&ketoapi.RelationTuple{}).FromDataProvider(d.RelationTuple)
			if err != nil {
				return nil, err
			}
			filtered = append(filtered, it)
		}
	}
	return
}

func (h *WriteHandler) TransactRelationTuples(ctx context.Context, req *connect.Request[rts.TransactRelationTuplesRequest]) (*connect.Response[rts.TransactRelationTuplesResponse], error) {
	insertTuples, err := protoTuplesWithAction(req.Msg.RelationTupleDeltas, rts.RelationTupleDelta_ACTION_INSERT)
	if err != nil {
		return nil, err
	}

	deleteTuples, err := protoTuplesWithAction(req.Msg.RelationTupleDeltas, rts.RelationTupleDelta_ACTION_DELETE)
	if err != nil {
		return nil, err
	}

	err = h.d.RelationTupleManager().Transaction(ctx, func(ctx context.Context) error {
		its, err := h.d.Mapper().FromTuple(ctx, append(insertTuples, deleteTuples...)...)
		if err != nil {
			return err
		}
		return h.d.RelationTupleManager().TransactRelationTuples(ctx, its[:len(insertTuples)], its[len(insertTuples):])
	})

	if err != nil {
		return nil, err
	}

	trace.SpanFromContext(ctx).AddEvent(events.NewRelationtuplesChanged(ctx))

	snaptokens := make([]string, len(insertTuples))
	for i := range insertTuples {
		snaptokens[i] = "not yet implemented"
	}
	return connect.NewResponse(&rts.TransactRelationTuplesResponse{
		Snaptokens: snaptokens,
	}), nil
}

func (h *WriteHandler) DeleteRelationTuples(ctx context.Context, req *connect.Request[rts.DeleteRelationTuplesRequest]) (*connect.Response[rts.DeleteRelationTuplesResponse], error) {
	var q ketoapi.RelationQuery

	switch {
	case req.Msg.RelationQuery != nil:
		q.FromDataProvider(&queryWrapper{req.Msg.RelationQuery})
		//lint:ignore SA1019 required for compatibility
	case req.Msg.Query != nil: //nolint:staticcheck
		//lint:ignore SA1019 backwards compatibility
		q.FromDataProvider(&deprecatedQueryWrapper{(*rts.ListRelationTuplesRequest_Query)(req.Msg.Query)}) //nolint:staticcheck
	default:
		return nil, errors.WithStack(herodot.ErrBadRequest().WithReason("invalid request"))
	}

	iq, err := h.d.ReadOnlyMapper().FromQuery(ctx, &q)
	if err != nil {
		return nil, err
	}
	if err := h.d.RelationTupleManager().DeleteAllRelationTuples(ctx, iq); err != nil {
		return nil, errors.WithStack(herodot.ErrInternalServerError().WithError(err.Error()))
	}

	trace.SpanFromContext(ctx).AddEvent(events.NewRelationtuplesDeleted(ctx))

	return connect.NewResponse(&rts.DeleteRelationTuplesResponse{}), nil
}

// Create Relationship Request Parameters
//
// swagger:parameters createRelationship
//
//lint:ignore U1000 required for OpenAPI
type createRelationship struct {
	// in: body
	Body createRelationshipBody
}

// Create Relationship Request Body
//
// swagger:model createRelationshipBody
type createRelationshipBody struct {
	ketoapi.RelationQuery
}

// swagger:route PUT /admin/relation-tuples relationship createRelationship
//
// # Create a Relationship
//
// Use this endpoint to create a relationship.
//
//	Consumes:
//	-  application/json
//
//	Produces:
//	- application/json
//
//	Schemes: http, https
//
//	Responses:
//	  201: relationship
//	  400: errorGeneric
//	  default: errorGeneric
//
//	Extensions:
//	  x-ory-ratelimit-bucket: keto-admin-low
func (h *WriteHandler) createRelation(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var rt ketoapi.RelationTuple
	if err := json.NewDecoder(r.Body).Decode(&rt); err != nil {
		h.d.Writer().WriteError(w, r, errors.WithStack(herodot.ErrBadRequest().WithError(err.Error())))
		return
	}

	if err := rt.Validate(); err != nil {
		h.d.Writer().WriteError(w, r, err)
		return
	}

	h.d.Logger().WithFields(rt.ToLoggerFields()).Debug("creating relation tuple")

	err := h.d.RelationTupleManager().Transaction(ctx, func(ctx context.Context) error {
		it, err := h.d.Mapper().FromTuple(ctx, &rt)
		if err != nil {
			h.d.Logger().WithError(err).WithFields(rt.ToLoggerFields()).Errorf("could not map relation tuple to UUIDs")
			return err
		}
		if err := h.d.RelationTupleManager().WriteRelationTuples(ctx, it...); err != nil {
			h.d.Logger().WithError(err).WithFields(rt.ToLoggerFields()).Errorf("got an error while creating the relation tuple")
			return err
		}
		return nil
	})

	if err != nil {
		h.d.Writer().WriteError(w, r, err)
		return
	}

	trace.SpanFromContext(ctx).AddEvent(events.NewRelationtuplesCreated(ctx))

	h.d.Writer().WriteCreated(w, r,
		ReadRouteBase+"?"+rt.ToURLQuery().Encode(),
		&rt,
	)
}

// swagger:route DELETE /admin/relation-tuples relationship deleteRelationships
//
// # Delete Relationships
//
// Use this endpoint to delete relationships
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
//	  204: emptyResponse
//	  400: errorGeneric
//	  default: errorGeneric
//
//	Extensions:
//	  x-ory-ratelimit-bucket: keto-admin-low
func (h *WriteHandler) deleteRelations(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if err := validate.All(r,
		validate.NoExtraQueryParams(ketoapi.RelationQueryKeys...),
		validate.QueryParamsContainsOneOf(ketoapi.NamespaceKey),
		validate.HasEmptyBody(),
	); err != nil {
		h.d.Writer().WriteError(w, r, err)
		return
	}

	q := r.URL.Query()
	query, err := (&ketoapi.RelationQuery{}).FromURLQuery(q)
	if err != nil {
		h.d.Writer().WriteError(w, r, herodot.ErrBadRequest().WithError(err.Error()))
		return
	}

	l := h.d.Logger()
	for k := range q {
		l = l.WithField(k, q.Get(k))
	}
	l.Debug("deleting relationships")

	iq, err := h.d.ReadOnlyMapper().FromQuery(ctx, query)
	if err != nil {
		h.d.Logger().WithError(err).Errorf("could not map fields to UUIDs")
		h.d.Writer().WriteError(w, r, err)
		return
	}
	if err := h.d.RelationTupleManager().DeleteAllRelationTuples(ctx, iq); err != nil {
		l.WithError(err).Errorf("got an error while deleting relationships")
		h.d.Writer().WriteError(w, r, herodot.ErrInternalServerError().WithError(err.Error()))
		return
	}

	trace.SpanFromContext(ctx).AddEvent(events.NewRelationtuplesDeleted(ctx))

	w.WriteHeader(http.StatusNoContent)
}

func internalTuplesWithAction(deltas []*ketoapi.PatchDelta, action ketoapi.PatchAction) (filtered []*ketoapi.RelationTuple) {
	for _, d := range deltas {
		if d.Action == action {
			filtered = append(filtered, d.RelationTuple)
		}
	}
	return
}

// swagger:route PATCH /admin/relation-tuples relationship patchRelationships
//
// # Patch Multiple Relationships
//
// Use this endpoint to patch one or more relationships.
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Schemes: http, https
//
//	Responses:
//	  204: emptyResponse
//	  400: errorGeneric
//	  404: errorGeneric
//	  default: errorGeneric
//
//	Extensions:
//	  x-ory-ratelimit-bucket: keto-admin-low
func (h *WriteHandler) patchRelationTuples(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var deltas []*ketoapi.PatchDelta
	if err := json.NewDecoder(r.Body).Decode(&deltas); err != nil {
		h.d.Writer().WriteError(w, r, herodot.ErrBadRequest().WithError(err.Error()))
		return
	}
	for _, d := range deltas {
		if d.RelationTuple == nil {
			h.d.Writer().WriteError(w, r, herodot.ErrBadRequest().WithError("relation_tuple is missing"))
			return
		}
		if err := d.RelationTuple.Validate(); err != nil {
			h.d.Writer().WriteError(w, r, err)
			return
		}

		switch d.Action {
		case ketoapi.ActionInsert, ketoapi.ActionDelete:
		default:
			h.d.Writer().WriteError(w, r, herodot.ErrBadRequest().WithError("unknown action "+string(d.Action)))
			return
		}
	}

	insertTuples := internalTuplesWithAction(deltas, ketoapi.ActionInsert)
	deleteTuples := internalTuplesWithAction(deltas, ketoapi.ActionDelete)

	err := h.d.RelationTupleManager().Transaction(ctx, func(ctx context.Context) error {
		its, err := h.d.Mapper().FromTuple(ctx, append(insertTuples, deleteTuples...)...)
		if err != nil {
			h.d.Logger().WithError(err).Errorf("got an error while mapping fields to UUID")
			return err
		}
		return h.d.RelationTupleManager().TransactRelationTuples(ctx, its[:len(insertTuples)], its[len(insertTuples):])
	})

	if err != nil {
		h.d.Writer().WriteError(w, r, err)
		return
	}

	trace.SpanFromContext(ctx).AddEvent(events.NewRelationtuplesChanged(ctx))

	w.WriteHeader(http.StatusNoContent)
}
