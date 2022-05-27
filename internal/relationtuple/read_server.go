package relationtuple

import (
	"context"
	"net/http"
	"strconv"

	rts "github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2"

	"github.com/ory/herodot"

	"github.com/pkg/errors"

	"github.com/julienschmidt/httprouter"

	"github.com/ory/keto/internal/x"
)

var (
	_ rts.ReadServiceServer = (*handler)(nil)
	_                       = (*getRelationsParams)(nil)
)

func (h *handler) ListRelationTuples(ctx context.Context, req *rts.ListRelationTuplesRequest) (*rts.ListRelationTuplesResponse, error) {
	if req.Query == nil {
		return nil, errors.New("invalid request")
	}

	q, err := (&RelationQuery{}).FromProto(req.Query)
	if err != nil {
		return nil, err
	}

	rels, nextPage, err := h.d.RelationTupleManager().GetRelationTuples(ctx, q,
		x.WithSize(int(req.PageSize)),
		x.WithToken(req.PageToken),
	)
	if err != nil {
		return nil, err
	}

	resp := &rts.ListRelationTuplesResponse{
		RelationTuples: make([]*rts.RelationTuple, len(rels)),
		NextPageToken:  nextPage,
	}
	for i, r := range rels {
		resp.RelationTuples[i] = r.ToProto()
	}

	return resp, nil
}

// swagger:parameters getRelationTuples
type getRelationsParams struct {
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

	// swagger:allOf
	x.PaginationOptions
}

// swagger:route GET /relation-tuples read getRelationTuples
//
// Query relation tuples
//
// Get all relation tuples that match the query. Only the namespace field is required.
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
//       200: getRelationTuplesResponse
//       404: genericError
//       500: genericError
func (h *handler) getRelations(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	q := r.URL.Query()
	query, err := (&RelationQuery{}).FromURLQuery(q)
	if err != nil {
		h.d.Writer().WriteError(w, r, herodot.ErrBadRequest.WithError(err.Error()))
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
			h.d.Writer().WriteError(w, r, herodot.ErrBadRequest.WithError(err.Error()))
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
	}

	h.d.Writer().Write(w, r, resp)
}
