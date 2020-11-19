package check

import (
	"net/http"

	"github.com/ory/keto/internal/relationtuple"

	"github.com/julienschmidt/httprouter"

	"github.com/ory/keto/internal/x"
)

type (
	handlerDependencies interface {
		EngineProvider
		x.LoggerProvider
		x.WriterProvider
	}
	handler struct {
		d handlerDependencies
	}
)

const routeBase = "/check"

func NewHandler(d handlerDependencies) *handler {
	return &handler{d: d}
}

func (h *handler) RegisterPublicRoutes(router *httprouter.Router) {
	router.GET(routeBase, h.getCheck)
}

func (h *handler) getCheck(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	subjectID := r.URL.Query().Get("subject-id")
	objectID := r.URL.Query().Get("object-id")
	namespace := r.URL.Query().Get("namespace")
	relationName := r.URL.Query().Get("relation-name")

	res, err := h.d.PermissionEngine().SubjectIsAllowed(r.Context(), &relationtuple.InternalRelationTuple{
		Relation:  relationName,
		ObjectID:  objectID,
		Namespace: namespace,
		Subject:   relationtuple.SubjectFromString(subjectID),
	})
	if err != nil {
		h.d.Writer().WriteError(w, r, err)
		return
	}

	if res {
		h.d.Writer().WriteCode(w, r, http.StatusOK, "allowed")
		return
	}

	h.d.Writer().WriteCode(w, r, http.StatusForbidden, "rejected")
}
