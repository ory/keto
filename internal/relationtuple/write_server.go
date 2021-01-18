package relationtuple

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/ory/herodot"
	"github.com/pkg/errors"

	acl "github.com/ory/keto/api/keto/acl/v1alpha1"
)

var _ acl.WriteServiceServer = (*Handler)(nil)

func tuplesWithAction(deltas []*acl.RelationTupleWriteDelta, action acl.RelationTupleWriteDelta_Action) (filtered []*InternalRelationTuple) {
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

func (h *Handler) WriteRelationTuples(ctx context.Context, req *acl.WriteRelationTuplesRequest) (*acl.WriteRelationTuplesResponse, error) {
	insertTuples := tuplesWithAction(req.RelationTupleDeltas, acl.RelationTupleWriteDelta_INSERT)

	err := h.d.RelationTupleManager().WriteRelationTuples(ctx, insertTuples...)
	if err != nil {
		return nil, err
	}

	snaptokens := make([]string, len(insertTuples))
	for i := range insertTuples {
		snaptokens[i] = "not yet implemented"
	}
	return &acl.WriteRelationTuplesResponse{
		Snaptokens: snaptokens,
	}, nil
}

func (h *Handler) createRelation(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var rel InternalRelationTuple

	if err := json.NewDecoder(r.Body).Decode(&rel); err != nil {
		h.d.Writer().WriteError(w, r, errors.WithStack(herodot.ErrBadRequest))
		return
	}

	if err := h.d.RelationTupleManager().WriteRelationTuples(r.Context(), &rel); err != nil {
		h.d.Logger().WithError(err).WithField("relationtuple", rel).Errorf("got an error while creating the relation tuple")
		h.d.Writer().WriteError(w, r, errors.WithStack(herodot.ErrInternalServerError))
		return
	}

	w.WriteHeader(http.StatusCreated)
}
