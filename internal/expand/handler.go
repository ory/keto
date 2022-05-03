package expand

import (
	"context"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/ory/herodot"
	"google.golang.org/grpc"

	"github.com/ory/keto/internal/relationtuple"
	"github.com/ory/keto/internal/x"
	rts "github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2"
)

type (
	handlerDependencies interface {
		EngineProvider
		relationtuple.ManagerProvider
		x.LoggerProvider
		x.WriterProvider
	}
	handler struct {
		d handlerDependencies
	}
)

var (
	_ rts.ExpandServiceServer = (*handler)(nil)
	_ *getExpandRequest       = nil
)

const RouteBase = "/relation-tuples/expand"

func NewHandler(d handlerDependencies) *handler {
	return &handler{d: d}
}

func (h *handler) RegisterReadRoutes(r *x.ReadRouter) {
	r.GET(RouteBase, h.getExpand)
}

func (h *handler) RegisterWriteRoutes(_ *x.WriteRouter) {}

func (h *handler) RegisterReadGRPC(s *grpc.Server) {
	rts.RegisterExpandServiceServer(s, h)
}

func (h *handler) RegisterWriteGRPC(s *grpc.Server) {}

// swagger:parameters getExpand
type getExpandRequest struct {
	// in:query
	MaxDepth int `json:"max-depth"`
}

// swagger:route GET /relation-tuples/expand read getExpand
//
// Expand a Relation Tuple
//
// Use this endpoint to expand a relation tuple.
//
//     Consumes:
//     -  application/x-www-form-urlencoded
//
//     Produces:
//     - application/json
//
//     Schemes: http, https
//
//     Responses:
//       200: expandTree
//       400: genericError
//       404: genericError
//       500: genericError
func (h *handler) getExpand(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	maxDepth, err := x.GetMaxDepthFromQuery(r.URL.Query())
	if err != nil {
		h.d.Writer().WriteError(w, r, herodot.ErrBadRequest.WithError(err.Error()))
		return
	}

	subject := (&relationtuple.SubjectSet{}).FromURLQuery(r.URL.Query())

	if err := h.d.UUIDMappingManager().MapFieldsToUUID(r.Context(), subject); err != nil {
		h.d.Writer().WriteError(w, r, err)
		return
	}
	res, err := h.d.ExpandEngine().BuildTree(r.Context(), subject, maxDepth)
	if err != nil {
		h.d.Writer().WriteError(w, r, err)
		return
	}
	if err := h.d.UUIDMappingManager().MapFieldsFromUUID(r.Context(), res); err != nil {
		h.d.Writer().WriteError(w, r, err)
		return
	}

	h.d.Writer().Write(w, r, res)
}

func (h *handler) Expand(ctx context.Context, req *rts.ExpandRequest) (*rts.ExpandResponse, error) {
	subject, err := relationtuple.SubjectFromProto(req.Subject)
	if err != nil {
		return nil, err
	}

	if err := h.d.UUIDMappingManager().MapFieldsToUUID(ctx, subject); err != nil {
		return nil, err
	}
	res, err := h.d.ExpandEngine().BuildTree(ctx, subject, int(req.MaxDepth))
	if err != nil {
		return nil, err
	}
	if err := h.d.UUIDMappingManager().MapFieldsFromUUID(ctx, res); err != nil {
		return nil, err
	}

	return &rts.ExpandResponse{Tree: res.ToProto()}, nil
}
