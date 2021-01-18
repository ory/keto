package relationtuple

import (
	"context"
	"net/http"

	"github.com/julienschmidt/httprouter"

	acl "github.com/ory/keto/api/keto/acl/v1alpha1"
	"github.com/ory/keto/internal/x"
)

var _ acl.ReadServiceServer = (*Handler)(nil)

func (h *Handler) ListRelationTuples(ctx context.Context, req *acl.ListRelationTuplesRequest) (*acl.ListRelationTuplesResponse, error) {
	rels, nextPage, err := h.d.RelationTupleManager().GetRelationTuples(ctx,
		&RelationQuery{
			Namespace: req.Query.Namespace,
			Object:    req.Query.Object,
			Relation:  req.Query.Relation,
			Subject:   SubjectFromGRPC(req.Query.Subject),
		},
		x.WithSize(uint(req.PageSize)),
		x.WithToken(req.PageToken),
	)
	if err != nil {
		return nil, err
	}

	resp := &acl.ListRelationTuplesResponse{
		RelationTuples: make([]*acl.RelationTuple, len(rels)),
		NextPageToken:  nextPage,
	}
	for i, r := range rels {
		resp.RelationTuples[i] = r.ToGRPC()
	}

	return resp, nil
}

func (h *Handler) getRelations(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	query, err := (&RelationQuery{}).FromURLQuery(r.URL.Query())
	if err != nil {
		h.d.Writer().WriteError(w, r, err)
		return
	}

	rels, nextPage, err := h.d.RelationTupleManager().GetRelationTuples(r.Context(), query)
	if err != nil {
		h.d.Writer().WriteError(w, r, err)
		return
	}

	resp := map[string]interface{}{
		"relations": rels,
		"next_page": nextPage,
	}

	h.d.Writer().Write(w, r, resp)
}
