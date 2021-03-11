package relationtuple

import (
	"context"
	"encoding/json"
	"net/http"

	acl "github.com/ory/keto/proto/ory/keto/acl/v1alpha1"

	"github.com/julienschmidt/httprouter"
	"github.com/ory/herodot"
	"github.com/pkg/errors"
)

var _ acl.WriteServiceServer = (*handler)(nil)

func tuplesWithAction(deltas []*acl.RelationTupleDelta, action acl.RelationTupleDelta_Action) (filtered []*InternalRelationTuple, err error) {
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

func (h *handler) TransactRelationTuples(ctx context.Context, req *acl.TransactRelationTuplesRequest) (*acl.TransactRelationTuplesResponse, error) {
	insertTuples, err := tuplesWithAction(req.RelationTupleDeltas, acl.RelationTupleDelta_INSERT)
	if err != nil {
		return nil, err
	}

	deleteTuples, err := tuplesWithAction(req.RelationTupleDeltas, acl.RelationTupleDelta_DELETE)
	if err != nil {
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
	return &acl.TransactRelationTuplesResponse{
		Snaptokens: snaptokens,
	}, nil
}

// The basic ACL relation tuple
//
// swagger:parameters postCheck createRelationTuple
// nolint:deadcode,unused
type bodyRelationTuple struct {
	// in: body
	Payload InternalRelationTuple
}

// swagger:route PUT /relationtuple write createRelationTuple
//
// Create a relation tuple
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
//       201: InternalRelationTuple
//       400: genericError
//       500: genericError
func (h *handler) createRelation(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var rel InternalRelationTuple

	if err := json.NewDecoder(r.Body).Decode(&rel); err != nil {
		h.d.Writer().WriteError(w, r, errors.WithStack(herodot.ErrBadRequest.WithError(err.Error())))
		return
	}

	h.d.Logger().WithFields(rel.ToLoggerFields()).Debug("creating relation tuple")

	if err := h.d.RelationTupleManager().WriteRelationTuples(r.Context(), &rel); err != nil {
		h.d.Logger().WithError(err).WithFields(rel.ToLoggerFields()).Errorf("got an error while creating the relation tuple")
		h.d.Writer().WriteError(w, r, errors.WithStack(herodot.ErrInternalServerError.WithError(err.Error())))
		return
	}

	h.d.Writer().WriteCreated(w, r, RouteBase+"?"+rel.ToURLQuery().Encode(), rel)
}

// swagger:route DELETE /relationtuple write deleteRelationTuple
//
// Delete a relation tuple
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
//       500: genericError
func (h *handler) deleteRelation(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	rel, err := (&InternalRelationTuple{}).FromURLQuery(r.URL.Query())
	if err != nil {
		h.d.Writer().WriteError(w, r, herodot.ErrBadRequest.WithError(err.Error()))
		return
	}

	if err := h.d.RelationTupleManager().DeleteRelationTuples(r.Context(), rel); err != nil {
		h.d.Logger().WithError(err).WithFields(rel.ToLoggerFields()).Errorf("got an error while deleting the relation tuple")
		h.d.Writer().WriteError(w, r, herodot.ErrInternalServerError.WithError(err.Error()))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
