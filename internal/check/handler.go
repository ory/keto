package check

import (
	"context"
	"net/http"

	"google.golang.org/grpc"

	acl "github.com/ory/keto/api/keto/acl/v1alpha1"

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
	Handler struct {
		d handlerDependencies
	}
)

var _ acl.CheckServiceServer = (*Handler)(nil)

func NewHandler(d handlerDependencies) *Handler {
	return &Handler{d: d}
}

const RouteBase = "/check"

func (h *Handler) RegisterReadRoutes(r *x.ReadRouter) {
	r.GET(RouteBase, h.getCheck)
}

func (h *Handler) RegisterWriteRoutes(_ *x.WriteRouter) {}

func (h *Handler) RegisterReadGRPC(s *grpc.Server) {
	acl.RegisterCheckServiceServer(s, h)
}

func (h *Handler) RegisterWriteGRPC(_ *grpc.Server) {}

func (h *Handler) getCheck(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	tuple, err := (&relationtuple.InternalRelationTuple{}).FromURLQuery(r.URL.Query())
	if err != nil {
		h.d.Writer().WriteError(w, r, err)
	}

	allowed, err := h.d.PermissionEngine().SubjectIsAllowed(r.Context(), tuple)
	if err != nil {
		h.d.Writer().WriteError(w, r, err)
		return
	}

	if allowed {
		h.d.Writer().WriteCode(w, r, http.StatusOK, "allowed")
		return
	}

	h.d.Writer().WriteCode(w, r, http.StatusForbidden, "rejected")
}

func (h *Handler) Check(ctx context.Context, req *acl.CheckRequest) (*acl.CheckResponse, error) {
	tuple := (&relationtuple.InternalRelationTuple{}).FromDataProvider(req)

	allowed, err := h.d.PermissionEngine().SubjectIsAllowed(ctx, tuple)
	// TODO add content change handling
	if err != nil {
		return nil, err
	}

	return &acl.CheckResponse{
		Allowed:   allowed,
		Snaptoken: "not yet implemented",
	}, nil
}
