package relationtuple

import (
	"context"
	"net/http"
	"strconv"

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
	q := r.URL.Query()
	query, err := (&RelationQuery{}).FromURLQuery(q)
	if err != nil {
		h.d.Writer().WriteError(w, r, err)
		return
	}

	l := h.d.Logger()
	for k := range q {
		l = l.WithField(k, q.Get(k))
	}
	l.Debug("querying relation tuples")

	var paginationOpts []x.PaginationOptionSetter
	if pageToken := q.Get("page_token"); pageToken != "" {
		paginationOpts = append(paginationOpts, x.WithToken(pageToken))
	}

	if pageSize := q.Get("page_size"); pageSize != "" {
		s, err := strconv.ParseInt(pageSize, 0, 0)
		if err != nil {
			h.d.Writer().WriteError(w, r, err)
			return
		}
		paginationOpts = append(paginationOpts, x.WithSize(int(s)))
	}

	rels, nextPage, err := h.d.RelationTupleManager().GetRelationTuples(r.Context(), query, paginationOpts...)
	if err != nil {
		h.d.Writer().WriteError(w, r, err)
		return
	}

	resp := &GetResponse{
		RelationTuples: rels,
		NextPageToken:  nextPage,
		IsLastPage:     nextPage == x.PageTokenEnd,
	}

	h.d.Writer().Write(w, r, resp)
}
