package relationtuple

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/ory/keto/ketoapi"

	rts "github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2"

	"github.com/julienschmidt/httprouter"
	"github.com/ory/herodot"
	"github.com/pkg/errors"
)

var (
	_ rts.WriteServiceServer = (*handler)(nil)
)

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

func (h *handler) TransactRelationTuples(ctx context.Context, req *rts.TransactRelationTuplesRequest) (*rts.TransactRelationTuplesResponse, error) {
	insertTuples, err := protoTuplesWithAction(req.RelationTupleDeltas, rts.RelationTupleDelta_ACTION_INSERT)
	if err != nil {
		return nil, err
	}

	deleteTuples, err := protoTuplesWithAction(req.RelationTupleDeltas, rts.RelationTupleDelta_ACTION_DELETE)
	if err != nil {
		return nil, err
	}

	its, err := h.d.Mapper().FromTuple(ctx, append(insertTuples, deleteTuples...)...)
	if err != nil {
		return nil, err
	}

	err = h.d.RelationTupleManager().TransactRelationTuples(ctx, its[:len(insertTuples)], its[len(insertTuples):])
	if err != nil {
		return nil, err
	}

	snaptokens := make([]string, len(insertTuples))
	for i := range insertTuples {
		snaptokens[i] = "not yet implemented"
	}
	return &rts.TransactRelationTuplesResponse{
		Snaptokens: snaptokens,
	}, nil
}

func (h *handler) DeleteRelationTuples(ctx context.Context, req *rts.DeleteRelationTuplesRequest) (*rts.DeleteRelationTuplesResponse, error) {
	var q ketoapi.RelationQuery

	switch {
	case req.RelationQuery != nil:
		q.FromDataProvider(&queryWrapper{req.RelationQuery})
	case req.Query != nil:
		q.FromDataProvider(&deprecatedQueryWrapper{(*rts.ListRelationTuplesRequest_Query)(req.Query)})
	default:
		return nil, errors.WithStack(herodot.ErrBadRequest.WithReason("invalid request"))
	}

	iq, err := h.d.Mapper().FromQuery(ctx, &q)
	if err != nil {
		return nil, err
	}
	if err := h.d.RelationTupleManager().DeleteAllRelationTuples(ctx, iq); err != nil {
		return nil, errors.WithStack(herodot.ErrInternalServerError.WithError(err.Error()))
	}

	return &rts.DeleteRelationTuplesResponse{}, nil
}

// swagger:route PUT /admin/relation-tuples write createRelationTuple
//
// Create a Relation Tuple
//
// Use this endpoint to create a relation tuple.
//
//     Consumes:
//     -  application/json
//
//     Produces:
//     - application/json
//
//     Schemes: http, https
//
//     Responses:
//       201: relationQuery
//       400: genericError
//       500: genericError
func (h *handler) createRelation(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx := r.Context()

	var rt ketoapi.RelationTuple
	if err := json.NewDecoder(r.Body).Decode(&rt); err != nil {
		h.d.Writer().WriteError(w, r, errors.WithStack(herodot.ErrBadRequest.WithError(err.Error())))
		return
	}

	h.d.Logger().WithFields(rt.ToLoggerFields()).Debug("creating relation tuple")

	it, err := h.d.Mapper().FromTuple(ctx, &rt)
	if err != nil {
		h.d.Logger().WithError(err).WithFields(rt.ToLoggerFields()).Errorf("could not map relation tuple to UUIDs")
		h.d.Writer().WriteError(w, r, err)
		return
	}
	if err := h.d.RelationTupleManager().WriteRelationTuples(ctx, it...); err != nil {
		h.d.Logger().WithError(err).WithFields(rt.ToLoggerFields()).Errorf("got an error while creating the relation tuple")
		h.d.Writer().WriteError(w, r, err)
		return
	}

	h.d.Writer().WriteCreated(w, r,
		ReadRouteBase+"?"+rt.ToURLQuery().Encode(),
		&rt,
	)
}

// swagger:route DELETE /admin/relation-tuples write deleteRelationTuples
//
// Delete Relation Tuples
//
// Use this endpoint to delete relation tuples
//
//     Consumes:
//     -  application/x-www-form-urlencoded
//
//     Produces:
//     - application/json
//
//     Schemes: http, https
//
//     Responses:
//       204: emptyResponse
//       400: genericError
//       500: genericError
func (h *handler) deleteRelations(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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
	l.Debug("deleting relation tuples")

	iq, err := h.d.Mapper().FromQuery(ctx, query)
	if err != nil {
		h.d.Logger().WithError(err).Errorf("could not map fields to UUIDs")
		h.d.Writer().WriteError(w, r, err)
		return
	}
	if err := h.d.RelationTupleManager().DeleteAllRelationTuples(ctx, iq); err != nil {
		l.WithError(err).Errorf("got an error while deleting relation tuples")
		h.d.Writer().WriteError(w, r, herodot.ErrInternalServerError.WithError(err.Error()))
		return
	}

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

// swagger:route PATCH /admin/relation-tuples write patchRelationTuples
//
// Patch Multiple Relation Tuples
//
// Use this endpoint to patch one or more relation tuples.
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Schemes: http, https
//
//     Responses:
//       204: emptyResponse
//       400: genericError
//       404: genericError
//       500: genericError
func (h *handler) patchRelationTuples(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx := r.Context()

	var deltas []*ketoapi.PatchDelta
	if err := json.NewDecoder(r.Body).Decode(&deltas); err != nil {
		h.d.Writer().WriteError(w, r, herodot.ErrBadRequest.WithError(err.Error()))
		return
	}
	for _, d := range deltas {
		if d.RelationTuple == nil {
			h.d.Writer().WriteError(w, r, herodot.ErrBadRequest.WithError("relation_tuple is missing"))
			return
		}
		switch d.Action {
		case ketoapi.ActionInsert, ketoapi.ActionDelete:
		default:
			h.d.Writer().WriteError(w, r, herodot.ErrBadRequest.WithError("unknown action "+string(d.Action)))
			return
		}
	}

	insertTuples := internalTuplesWithAction(deltas, ketoapi.ActionInsert)
	deleteTuples := internalTuplesWithAction(deltas, ketoapi.ActionDelete)

	its, err := h.d.Mapper().FromTuple(ctx, append(insertTuples, deleteTuples...)...)
	if err != nil {
		h.d.Logger().WithError(err).Errorf("got an error while mapping fields to UUID")
		h.d.Writer().WriteError(w, r, err)
		return
	}
	if err := h.d.RelationTupleManager().
		TransactRelationTuples(
			ctx,
			its[:len(insertTuples)],
			its[len(insertTuples):]); err != nil {

		h.d.Writer().WriteError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
