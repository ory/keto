package relation

import (
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
}

func (h *handler) getRelations(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	userID := r.URL.Query().Get("user-id")
	objectID := r.URL.Query().Get("object-id")

	var rels []*models.Relation
	var err error

	if userID != "" {
		rels, err = h.d.RelationManager().GetRelationsByUser(r.Context(), userID, 0, 100)
	} else if objectID != "" {
		rels, err = h.d.RelationManager().GetRelationsByObject(r.Context(), objectID, 0, 100)
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
