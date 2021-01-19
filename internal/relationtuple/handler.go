package relationtuple

import (
	"google.golang.org/grpc"

	acl "github.com/ory/keto/api/keto/acl/v1alpha1"
	"github.com/ory/keto/internal/x"
)

type (
	handlerDeps interface {
		ManagerProvider
		x.LoggerProvider
		x.WriterProvider
	}
	handler struct {
		d handlerDeps
	}
)

const (
	RouteBase = "/relationtuple"
)

func NewHandler(d handlerDeps) *handler {
	return &handler{
		d: d,
	}
}

func (h *handler) RegisterReadRoutes(r *x.ReadRouter) {
	r.GET(RouteBase, h.getRelations)
}

func (h *handler) RegisterWriteRoutes(r *x.WriteRouter) {
	r.PUT(RouteBase, h.createRelation)
}

func (h *handler) RegisterReadGRPC(s *grpc.Server) {
	acl.RegisterReadServiceServer(s, h)
}

func (h *handler) RegisterWriteGRPC(s *grpc.Server) {
	acl.RegisterWriteServiceServer(s, h)
}
