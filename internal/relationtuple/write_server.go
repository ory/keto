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

func tuplesWithAction(deltas []*acl.RelationTupleDelta, action acl.RelationTupleDelta_Action) (filtered []*InternalRelationTuple) {
	for _, d := range deltas {
		if d.Action == action {
			filtered = append(
				filtered,
				(&InternalRelationTuple{}).FromDataProvider(d.RelationTuple),
			)
		}
	}
	return
}

func (h *handler) TransactRelationTuples(ctx context.Context, req *acl.TransactRelationTuplesRequest) (*acl.TransactRelationTuplesResponse, error) {
	insertTuples := tuplesWithAction(req.RelationTupleDeltas, acl.RelationTupleDelta_INSERT)

	err := h.d.RelationTupleManager().WriteRelationTuples(ctx, insertTuples...)
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

func (h *handler) createRelation(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var rel InternalRelationTuple

	if err := json.NewDecoder(r.Body).Decode(&rel); err != nil {
		h.d.Writer().WriteError(w, r, errors.WithStack(herodot.ErrBadRequest))
		return
	}

	h.d.Logger().WithFields(rel.ToLoggerFields()).Debug("creating relation tuple")

	if err := h.d.RelationTupleManager().WriteRelationTuples(r.Context(), &rel); err != nil {
		h.d.Logger().WithError(err).WithField("relationtuple", rel).Errorf("got an error while creating the relation tuple")
		h.d.Writer().WriteError(w, r, errors.WithStack(herodot.ErrInternalServerError))
		return
	}

	w.WriteHeader(http.StatusCreated)
}
