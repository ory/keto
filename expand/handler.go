package expand

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"

	"github.com/ory/keto/models"
	"github.com/ory/keto/x"
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

const routeBase = "/expand"

func NewHandler(d handlerDependencies) *handler {
	return &handler{d: d}
}

func (h *handler) RegisterPublicRoutes(router *httprouter.Router) {
	router.GET(routeBase, h.getCheck)
}

func (h *handler) getCheck(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	objectID := r.URL.Query().Get("object-id")
	relationName := r.URL.Query().Get("relation-name")
	depth, err := strconv.ParseInt(r.URL.Query().Get("depth"), 0, 0)
	if err != nil {
		h.d.Writer().WriteError(w, r, err)
		return
	}

	res, err := h.d.ExpandEngine().BuildTree(r.Context(), &models.UserSet{
		Relation: relationName,
		Object:   (&models.Object{}).FromString(objectID),
	}, int(depth))
	if err != nil {
		h.d.Writer().WriteError(w, r, err)
		return
	}

	h.d.Writer().Write(w, r, res)
}
