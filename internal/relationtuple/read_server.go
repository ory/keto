package relationtuple

import (
	"context"
	"net/http"

	acl "github.com/ory/keto/proto/ory/keto/acl/v1alpha1"

	"github.com/julienschmidt/httprouter"

	"github.com/ory/keto/internal/x"
)

var _ acl.ReadServiceServer = (*handler)(nil)

func (h *handler) ListRelationTuples(ctx context.Context, req *acl.ListRelationTuplesRequest) (*acl.ListRelationTuplesResponse, error) {
	rels, nextPage, err := h.d.RelationTupleManager().GetRelationTuples(ctx,
		&RelationQuery{
			Namespace: req.Query.Namespace,
			Object:    req.Query.Object,
			Relation:  req.Query.Relation,
			Subject:   SubjectFromProto(req.Query.Subject),
		},
		x.WithSize(int(req.PageSize)),
		x.WithToken(req.PageToken),
	)
	if err != nil {
		return nil, err
	}

	resp := &acl.ListRelationTuplesResponse{
		RelationTuples: make([]*acl.RelationTuple, len(rels)),
		NextPageToken:  nextPage,
		IsLastPage:     nextPage == x.PageTokenEnd,
	}
	for i, r := range rels {
		resp.RelationTuples[i] = r.ToProto()
	}

	return resp, nil
}

func (h *handler) getRelations(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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
