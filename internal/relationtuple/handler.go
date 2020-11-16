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
	handlerDependencies interface {
		ManagerProvider
		x.LoggerProvider
		x.WriterProvider
	}
	handler struct {
		d handlerDependencies
	}
)

const (
	routeBase = "/relations"
)

func NewHandler(d handlerDependencies) *handler {
	return &handler{
		d: d,
	}
}

func (h *handler) RegisterPublicRoutes(router *httprouter.Router) {
	router.GET(routeBase, h.getRelations)
	router.PUT(routeBase, h.createRelation)
}

func (h *handler) getRelations(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	params := r.URL.Query()
	res, err := h.d.RelationTupleManager().GetRelationTuples(r.Context(), &RelationQuery{
		Relation: params.Get("relation"),
		Object:   (&Object{}).FromString(params.Get("object")),
		Subject:  SubjectFromString(params.Get("subject")),
	})

	if err != nil {
		h.d.Writer().WriteError(w, r, err)
		return
	}

	h.d.Writer().Write(w, r, res)
}

func (h *handler) createRelation(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var rel InternalRelationTuple

	if err := json.NewDecoder(r.Body).Decode(&rel); err != nil {
		h.d.Writer().WriteError(w, r, errors.WithStack(herodot.ErrBadRequest))
		return
	}

	if err := h.d.RelationTupleManager().WriteRelationTuples(r.Context(), &rel); err != nil {
		h.d.Writer().WriteError(w, r, errors.WithStack(herodot.ErrInternalServerError))
		return
	}

	w.WriteHeader(http.StatusCreated)
}
