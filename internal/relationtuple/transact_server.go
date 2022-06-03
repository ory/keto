package relationtuple

import (
	"context"
	"encoding/json"
	"github.com/ory/keto/ketoapi"
	"net/http"

	rts "github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2"

	"github.com/julienschmidt/httprouter"
	"github.com/ory/herodot"
	"github.com/pkg/errors"
)

var (
	_ rts.WriteServiceServer = (*handler)(nil)
	_                        = (*bodyRelationTuple)(nil)
	_                        = (*queryRelationTuple)(nil)
)

func protoTuplesWithAction(deltas []*rts.RelationTupleDelta, action rts.RelationTupleDelta_Action) (filtered []*InternalRelationTuple, err error) {
	for _, d := range deltas {
		if d.Action == action {
			it, err := (&InternalRelationTuple{}).FromDataProvider(d.RelationTuple)
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

	if err = h.d.UUIDMappingManager().MapFieldsToUUID(ctx, InternalRelationTuples(append(insertTuples, deleteTuples...))); err != nil {
		return nil, err
	}
	err = h.d.RelationTupleManager().TransactRelationTuples(ctx, insertTuples, deleteTuples)
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
	if req.Query == nil {
		return nil, errors.WithStack(herodot.ErrBadRequest.WithReason("invalid request"))
	}

	q, err := (&RelationQuery{}).FromProto(req.Query)
	if err != nil {
		return nil, errors.WithStack(herodot.ErrBadRequest.WithError(err.Error()))
	}

	if err := h.d.UUIDMappingManager().MapFieldsToUUID(ctx, q); err != nil {
		return nil, err
	}
	if err := h.d.RelationTupleManager().DeleteAllRelationTuples(ctx, q); err != nil {
		return nil, errors.WithStack(herodot.ErrInternalServerError.WithError(err.Error()))
	}

	return &rts.DeleteRelationTuplesResponse{}, nil
}

// The basic ACL relation tuple
//
// swagger:parameters postCheck createRelationTuple
type bodyRelationTuple struct {
	// in: body
	Payload RelationQuery
}

// The basic ACL relation tuple
//
// swagger:parameters getCheck deleteRelationTuples
type queryRelationTuple struct {
	// Namespace of the Relation Tuple
	//
	// in: query
	Namespace string `json:"namespace"`

	// Object of the Relation Tuple
	//
	// in: query
	Object string `json:"object"`

	// Relation of the Relation Tuple
	//
	// in: query
	Relation string `json:"relation"`

	// SubjectID of the Relation Tuple
	//
	// in: query
	// Either subject_set.* or subject_id are required.
	SubjectID string `json:"subject_id"`

	// Namespace of the Subject Set
	//
	// in: query
	// Either subject_set.* or subject_id are required.
	SNamespace string `json:"subject_set.namespace"`

	// Object of the Subject Set
	//
	// in: query
	// Either subject_set.* or subject_id are required.
	SObject string `json:"subject_set.object"`

	// Relation of the Subject Set
	//
	// in: query
	// Either subject_set.* or subject_id are required.
	SRelation string `json:"subject_set.relation"`
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
//       201: RelationQuery
//       400: genericError
//       500: genericError
func (h *handler) createRelation(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var rt ketoapi.RelationTuple

	if err := json.NewDecoder(r.Body).Decode(&rt); err != nil {
		h.d.Writer().WriteError(w, r, errors.WithStack(herodot.ErrBadRequest.WithError(err.Error())))
		return
	}

	h.d.Logger().WithFields(rt.ToLoggerFields()).Debug("creating relation tuple")

	if err := h.d.UUIDMappingManager().MapFieldsToUUID(r.Context(), &rel); err != nil {
		h.d.Logger().WithError(err).WithFields(rel.ToLoggerFields()).Errorf("got an error while mapping fields to UUID")
		h.d.Writer().WriteError(w, r, err)
		return
	}
	if err := h.d.RelationTupleManager().WriteRelationTuples(r.Context(), &rel); err != nil {
		h.d.Logger().WithError(err).WithFields(rel.ToLoggerFields()).Errorf("got an error while creating the relation tuple")
		h.d.Writer().WriteError(w, r, err)
		return
	}
	if err := h.d.UUIDMappingManager().MapFieldsFromUUID(r.Context(), &rel); err != nil {
		h.d.Logger().WithError(err).WithFields(rel.ToLoggerFields()).Errorf("got an error while mapping fields from UUID")
		h.d.Writer().WriteError(w, r, err)
		return
	}

	q, err := rel.ToURLQuery()
	if err != nil {
		h.d.Writer().WriteError(w, r, err)
		return
	}

	h.d.Writer().WriteCreated(w, r, ReadRouteBase+"?"+q.Encode(), &rel)
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

	if err := h.d.UUIDMappingManager().MapFieldsToUUID(r.Context(), query); err != nil {
		h.d.Logger().WithError(err).Errorf("got an error while mapping fields to UUID")
		h.d.Writer().WriteError(w, r, err)
		return
	}
	if err := h.d.RelationTupleManager().DeleteAllRelationTuples(r.Context(), query); err != nil {
		l.WithError(err).Errorf("got an error while deleting relation tuples")
		h.d.Writer().WriteError(w, r, herodot.ErrInternalServerError.WithError(err.Error()))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func internalTuplesWithAction(deltas []*PatchDelta, action patchAction) (filtered []*InternalRelationTuple) {
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
func (h *handler) patchRelations(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var deltas []*PatchDelta
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

	if err := h.d.UUIDMappingManager().MapFieldsToUUID(r.Context(), PatchDeltas(deltas)); err != nil {
		h.d.Logger().WithError(err).Errorf("got an error while mapping fields to UUID")
		h.d.Writer().WriteError(w, r, err)
		return
	}
	if err := h.d.RelationTupleManager().
		TransactRelationTuples(
			r.Context(),
			internalTuplesWithAction(deltas, ActionInsert),
			internalTuplesWithAction(deltas, ActionDelete)); err != nil {

		h.d.Writer().WriteError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
