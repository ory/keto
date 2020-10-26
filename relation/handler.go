package relation

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/pkg/errors"

	"github.com/ory/herodot"

	"github.com/ory/keto/models"
	"github.com/ory/keto/x"
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
	userID := r.URL.Query().Get("user-id")
	objectID := r.URL.Query().Get("object-id")

	var rels []*models.Relation
	var err error

	if userID != "" {
		rels, err = h.d.RelationManager().GetRelationsBySubject(r.Context(), userID)
	} else if objectID != "" {
		rels, err = h.d.RelationManager().GetRelationsByObject(r.Context(), objectID)
	} else {
		h.d.Writer().WriteError(w, r, errors.WithStack(herodot.ErrBadRequest))
		return
	}

	if err != nil {
		h.d.Writer().WriteError(w, r, err)
		return
	}
	h.d.Writer().Write(w, r, rels)
}

func (h *handler) createRelation(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var rel models.Relation

	if err := json.NewDecoder(r.Body).Decode(&rel); err != nil {
		h.d.Writer().WriteError(w, r, errors.WithStack(herodot.ErrBadRequest))
		return
	}

	if err := h.d.RelationManager().WriteRelation(r.Context(), &rel); err != nil {
		h.d.Writer().WriteError(w, r, errors.WithStack(herodot.ErrInternalServerError))
		return
	}

	w.WriteHeader(http.StatusCreated)
}
