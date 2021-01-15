package relationtuple

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/pkg/errors"

	"github.com/ory/herodot"

	"github.com/ory/keto/internal/x"
)

type (
	handlerDeps interface {
		ManagerProvider
		x.LoggerProvider
		x.WriterProvider
	}
	restHandler struct {
		d handlerDeps
	}
	grpcHandler struct {
		d handlerDeps
	}
)

const (
	RouteBase = "/relationtuple"
)

func NewHandler(d handlerDeps) *restHandler {
	return &restHandler{
		d: d,
	}
}

func NewGRPCServer(d handlerDeps) *grpcHandler {
	return &grpcHandler{
		d: d,
	}
}

func (h *restHandler) RegisterPublicRoutes(router *httprouter.Router) {
	router.GET(RouteBase, h.getRelations)
	router.PUT(RouteBase, h.createRelation)
}

func (h *restHandler) getRelations(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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

func (h *restHandler) createRelation(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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
