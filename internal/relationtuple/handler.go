package relationtuple

import (
	"github.com/julienschmidt/httprouter"

	"github.com/ory/keto/internal/x"
)

type (
	handlerDeps interface {
		ManagerProvider
		x.LoggerProvider
		x.WriterProvider
	}
	Handler struct {
		d handlerDeps
	}
)

const (
	RouteBase = "/relationtuple"
)

func NewHandler(d handlerDeps) *Handler {
	return &Handler{
		d: d,
	}
}

func (h *Handler) registerBasicRoutes(r *httprouter.Router) {
	r.GET(RouteBase, h.getRelations)
}

func (h *Handler) registerPrivilegedRoutes(r *httprouter.Router) {
	h.registerBasicRoutes(r)
	r.PUT(RouteBase, h.createRelation)
}

func (h *Handler) RegisterBasicRoutes(r *x.BasicRouter) {
	h.registerBasicRoutes(r.Router)
}

func (h *Handler) RegisterPrivilegedRoutes(r *x.PrivilegedRouter) {
	h.registerPrivilegedRoutes(r.Router)
}
