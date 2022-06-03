package check

import (
	"context"
	"encoding/json"
	"github.com/ory/keto/ketoapi"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/ory/herodot"
	"github.com/pkg/errors"
	"google.golang.org/grpc"

	"github.com/ory/keto/internal/relationtuple"
	"github.com/ory/keto/internal/x"
	rts "github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2"
)

type (
	handlerDependencies interface {
		EngineProvider
		relationtuple.ManagerProvider
		relationtuple.MapperProvider
		x.LoggerProvider
		x.WriterProvider
	}
	Handler struct {
		d handlerDependencies
	}
)

var (
	_ rts.CheckServiceServer = (*Handler)(nil)
	_ *getCheckRequest       = nil
)

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

	tuple, err := (&ketoapi.RelationTuple{}).FromURLQuery(r.URL.Query())
	if errors.Is(err, ketoapi.ErrNilSubject) {
		h.d.Writer().WriteError(w, r, herodot.ErrBadRequest.WithReason("Subject has to be specified."))
		return
	} else if err != nil {
		h.d.Writer().WriteError(w, r, herodot.ErrBadRequest.WithError(err.Error()))
		return
	}

	internalTuple, err := h.d.UUIDMapper().FromTuple(r.Context(), tuple)
	if err != nil {
		h.d.Writer().WriteError(w, r, err)
		return
	}
	allowed, err := h.d.PermissionEngine().SubjectIsAllowed(r.Context(), internalTuple[0], maxDepth)
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

	var tuple ketoapi.RelationTuple
	if err := json.NewDecoder(r.Body).Decode(&tuple); err != nil {
		h.d.Writer().WriteError(w, r, errors.WithStack(herodot.ErrBadRequest.WithReasonf("Unable to decode JSON payload: %s", err)))
		return
	}

	internalTuple, err := h.d.UUIDMapper().FromTuple(r.Context(), &tuple)
	if err != nil {
		h.d.Writer().WriteError(w, r, err)
		return
	}
	allowed, err := h.d.PermissionEngine().SubjectIsAllowed(r.Context(), internalTuple[0], maxDepth)
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
	tuple := (&ketoapi.RelationTuple{}).FromDataProvider(req)

	internalTuple, err := h.d.UUIDMapper().FromTuple(ctx, tuple)
	if err != nil {
		return nil, err
	}
	allowed, err := h.d.PermissionEngine().SubjectIsAllowed(ctx, internalTuple[0], int(req.MaxDepth))
	// TODO add content change handling
	if err != nil {
		return nil, err
	}

	return &rts.CheckResponse{
		Allowed:   allowed,
		Snaptoken: "not yet implemented",
	}, nil
}
