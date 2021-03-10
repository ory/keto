package check

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/ory/herodot"
	"github.com/pkg/errors"

	acl "github.com/ory/keto/proto/ory/keto/acl/v1alpha1"

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

var _ acl.CheckServiceServer = (*Handler)(nil)

func NewHandler(d handlerDependencies) *Handler {
	return &Handler{d: d}
}

const RouteBase = "/check"

func (h *Handler) RegisterReadRoutes(r *x.ReadRouter) {
	r.GET(RouteBase, h.getCheck)
	r.POST(RouteBase, h.postCheck)
}

func (h *Handler) RegisterWriteRoutes(_ *x.WriteRouter) {}

func (h *Handler) RegisterReadGRPC(s *grpc.Server) {
	acl.RegisterCheckServiceServer(s, h)
}

func (h *Handler) RegisterWriteGRPC(_ *grpc.Server) {}

// Represents the response for a check request.
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

// swagger:parameters getCheck
type getCheckParams struct {
	// swagger:allOf
	// in: query
	relationtuple.InternalRelationTuple
}

// swagger:route GET /check read getCheck
//
// Check a relation tuple
//
// To learn how relation tuples and the check works, head over to [the documentation](/TODO).
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
//		 403: getCheckResponse
//       500: genericError
func (h *Handler) getCheck(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	tuple, err := (&relationtuple.InternalRelationTuple{}).FromURLQuery(r.URL.Query())
	if err != nil {
		h.d.Writer().WriteError(w, r, herodot.ErrBadRequest.WithError(err.Error()))
	}

	allowed, err := h.d.PermissionEngine().SubjectIsAllowed(r.Context(), tuple)
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

// swagger:route POST /check read postCheck
//
// Check a relation tuple
//
// To learn how relation tuples and the check works, head over to [the documentation](/TODO).
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
//		 403: getCheckResponse
//       500: genericError
func (h *Handler) postCheck(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var tuple relationtuple.InternalRelationTuple
	if err := json.NewDecoder(r.Body).Decode(&tuple); err != nil {
		h.d.Writer().WriteError(w, r, errors.WithStack(herodot.ErrBadRequest.WithReasonf("Unable to decode JSON payload: %s", err)))
	}

	allowed, err := h.d.PermissionEngine().SubjectIsAllowed(r.Context(), &tuple)
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

func (h *Handler) Check(ctx context.Context, req *acl.CheckRequest) (*acl.CheckResponse, error) {
	tuple, err := (&relationtuple.InternalRelationTuple{}).FromDataProvider(req)
	if err != nil {
		return nil, err
	}

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
