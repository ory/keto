package expand

import (
	"context"
	"net/http"
	"strconv"

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

var _ acl.ExpandServiceServer = &grpcHandler{}

const RouteBase = "/expand"

func NewHandler(d handlerDependencies) *restHandler {
	return &restHandler{d: d}
}

func NewGRPCServer(d handlerDependencies) *grpcHandler {
	return &grpcHandler{d: d}
}

func (h *restHandler) RegisterPublicRoutes(router *httprouter.Router) {
	router.GET(RouteBase, h.getCheck)
}

func (h *restHandler) getCheck(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	depth, err := strconv.ParseInt(r.URL.Query().Get("depth"), 0, 0)
	if err != nil {
		h.d.Writer().WriteError(w, r, err)
		return
	}

	res, err := h.d.ExpandEngine().BuildTree(r.Context(), (&relationtuple.SubjectSet{}).FromURLQuery(r.URL.Query()), int(depth))
	if err != nil {
		h.d.Writer().WriteError(w, r, err)
		return
	}

	h.d.Writer().Write(w, r, res)
}

func (g *grpcHandler) Expand(ctx context.Context, req *acl.ExpandRequest) (*acl.ExpandResponse, error) {
	tree, err := g.d.ExpandEngine().BuildTree(ctx,
		relationtuple.SubjectFromProto(req.Subject),
		int(req.MaxDepth))
	if err != nil {
		return nil, err
	}

	return &acl.ExpandResponse{Tree: tree.ToProto()}, nil
}
