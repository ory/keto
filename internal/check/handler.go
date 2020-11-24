package check

import (
	"context"
	"net/http"

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
	restHandler struct {
		d handlerDependencies
	}
	grpcHandler struct {
		d handlerDependencies
	}
)

var _ acl.CheckServiceServer = &grpcHandler{}

func NewHandler(d handlerDependencies) *restHandler {
	return &restHandler{d: d}
}

const routeBase = "/check"

func (h *restHandler) RegisterPublicRoutes(router *httprouter.Router) {
	router.GET(routeBase, h.getCheck)
}

func (h *restHandler) getCheck(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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

func (g grpcHandler) Check(ctx context.Context, req *acl.CheckRequest) (*acl.CheckResponse, error) {
	tuple := &relationtuple.InternalRelationTuple{
		Namespace: req.Namespace,
		Object:    req.Object,
		Relation:  req.Relation,
		Subject:   relationtuple.SubjectFromGRPC(req.Subject),
	}

	allowed, err := g.d.PermissionEngine().SubjectIsAllowed(ctx, tuple)
	// TODO add content change handling
	if err != nil {
		return nil, err
	}

	return &acl.CheckResponse{
		Allowed:   allowed,
		Snaptoken: "not yet implemented",
	}, nil
}
