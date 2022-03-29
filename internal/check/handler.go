package check

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/ory/herodot"
	"github.com/pkg/errors"

	rts "github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2"

	"google.golang.org/grpc"

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

var _ rts.CheckServiceServer = (*Handler)(nil)

func NewHandler(d handlerDependencies) *Handler {
	return &Handler{d: d}
}

const RouteBase = "/relation-tuples/check"

func (h *Handler) RegisterReadRoutes(r *x.ReadRouter) {
	r.GET(RouteBase, h.getCheck)
	r.POST(RouteBase, h.postCheck)
}

func (h *Handler) RegisterWriteRoutes(_ *x.WriteRouter) {}

func (h *Handler) RegisterReadGRPC(s *grpc.Server) {
	rts.RegisterCheckServiceServer(s, h)
}

func (h *Handler) RegisterWriteGRPC(_ *grpc.Server) {}

// RESTResponse is the response for a check request.
//
// The content of the allowed field is mirrored in the HTTP status code.
//
// swagger:model getCheckResponse
type RESTResponse struct {
	// whether the relation tuple is allowed
	//
	// required: true
	Allowed bool `json:"allowed"`
}

// swagger:parameters getCheck postCheck
// nolint:deadcode,unused
type getCheckRequest struct {
	// in:query
	MaxDepth int `json:"max-depth"`
}

// swagger:route GET /relation-tuples/check read getCheck
//
// Check a relation tuple
//
// To learn how relation tuples and the check works, head over to [the documentation](../concepts/relation-tuples.mdx).
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
//       200: getCheckResponse
//       400: genericError
//       403: getCheckResponse
//       500: genericError
func (h *Handler) getCheck(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	maxDepth, err := x.GetMaxDepthFromQuery(r.URL.Query())
	if err != nil {
		h.d.Writer().WriteError(w, r, herodot.ErrBadRequest.WithError(err.Error()))
		return
	}

	tuple, err := (&relationtuple.InternalRelationTuple{}).FromURLQuery(r.URL.Query())
	if errors.Is(err, relationtuple.ErrNilSubject) {
		h.d.Writer().WriteError(w, r, herodot.ErrBadRequest.WithReason("Subject has to be specified."))
		return
	} else if err != nil {
		h.d.Writer().WriteError(w, r, herodot.ErrBadRequest.WithError(err.Error()))
		return
	}

	allowed, err := h.d.PermissionEngine().SubjectIsAllowed(r.Context(), tuple, maxDepth)
	if err != nil {
		h.d.Writer().WriteError(w, r, err)
		return
	}

	if allowed {
		h.d.Writer().WriteCode(w, r, http.StatusOK, &RESTResponse{Allowed: true})
		return
	}

	h.d.Writer().WriteCode(w, r, http.StatusForbidden, &RESTResponse{Allowed: false})
}

// swagger:route POST /relation-tuples/check read postCheck
//
// Check a relation tuple
//
// To learn how relation tuples and the check works, head over to [the documentation](../concepts/relation-tuples.mdx).
//
//     Consumes:
//     -  application/json
//
//     Produces:
//     - application/json
//
//     Schemes: http, https
//
//     Responses:
//       200: getCheckResponse
//       400: genericError
//       403: getCheckResponse
//       500: genericError
func (h *Handler) postCheck(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	maxDepth, err := x.GetMaxDepthFromQuery(r.URL.Query())
	if err != nil {
		h.d.Writer().WriteError(w, r, herodot.ErrBadRequest.WithError(err.Error()))
		return
	}

	var tuple relationtuple.InternalRelationTuple
	if err := json.NewDecoder(r.Body).Decode(&tuple); err != nil {
		h.d.Writer().WriteError(w, r, errors.WithStack(herodot.ErrBadRequest.WithReasonf("Unable to decode JSON payload: %s", err)))
		return
	}

	allowed, err := h.d.PermissionEngine().SubjectIsAllowed(r.Context(), &tuple, maxDepth)
	if err != nil {
		h.d.Writer().WriteError(w, r, err)
		return
	}

	if allowed {
		h.d.Writer().WriteCode(w, r, http.StatusOK, &RESTResponse{Allowed: true})
		return
	}

	h.d.Writer().WriteCode(w, r, http.StatusForbidden, &RESTResponse{Allowed: false})
}

func (h *Handler) Check(ctx context.Context, req *rts.CheckRequest) (*rts.CheckResponse, error) {
	tuple, err := (&relationtuple.InternalRelationTuple{}).FromDataProvider(req)
	if err != nil {
		return nil, err
	}

	allowed, err := h.d.PermissionEngine().SubjectIsAllowed(ctx, tuple, int(req.MaxDepth))
	// TODO add content change handling
	if err != nil {
		return nil, err
	}

	return &rts.CheckResponse{
		Allowed:   allowed,
		Snaptoken: "not yet implemented",
	}, nil
}
