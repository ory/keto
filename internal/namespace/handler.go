package namespace

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"

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

const routeBase = "/namespaces"

func NewHandler(d handlerDependencies) *handler {
	return &handler{d: d}
}

func (h *handler) RegisterPublicRoutes(router *httprouter.Router) {
	router.PUT(routeBase, h.createNamespace)
}

func (h *handler) createNamespace(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	n := &Namespace{
		ID: r.URL.Query().Get("name"),
	}

	if err := h.d.NamespaceManagerProvider().NewNamespace(r.Context(), n); err != nil {
		h.d.Writer().WriteError(w, r, err)
		return
	}

	h.d.Writer().WriteCreated(w, r, fmt.Sprintf("%s/%s", routeBase, n.ID), n)
}
