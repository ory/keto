package check

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/ory/herodot"

	"github.com/ory/keto/ketoapi"

	"github.com/julienschmidt/httprouter"
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

const (
	RouteBase        = "/relation-tuples/check"
	OpenAPIRouteBase = RouteBase + "/openapi"
)

func (h *Handler) RegisterReadRoutes(r *x.ReadRouter) {
	r.GET(RouteBase, h.getCheckMirrorStatus)
	r.GET(OpenAPIRouteBase, h.getCheckNoStatus)
	r.POST(RouteBase, h.postCheckMirrorStatus)
	r.POST(OpenAPIRouteBase, h.postCheckNoStatus)
}

func (h *Handler) RegisterReadGRPC(s *grpc.Server) {
	rts.RegisterCheckServiceServer(s, h)
}

// RESTResponse represents the response for a check request.
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

// swagger:route GET /relation-tuples/check/openapi read getCheck
//
// # Check a relation tuple
//
// To learn how relation tuples and the check works, head over to [the documentation](../concepts/relation-tuples.mdx).
//
//	Consumes:
//	-  application/x-www-form-urlencoded
//
//	Produces:
//	- application/json
//
//	Schemes: http, https
//
//	Responses:
//	  200: getCheckResponse
//	  400: genericError
//	  500: genericError
func (h *Handler) getCheckNoStatus(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	allowed, err := h.getCheck(r.Context(), r.URL.Query())
	if err != nil {
		h.d.Writer().WriteError(w, r, err)
		return
	}
	h.d.Writer().Write(w, r, &RESTResponse{Allowed: allowed})
}

// swagger:route GET /relation-tuples/check read getCheckMirrorStatus
//
// # Check a relation tuple
//
// To learn how relation tuples and the check works, head over to [the documentation](../concepts/relation-tuples.mdx).
//
//	Consumes:
//	-  application/x-www-form-urlencoded
//
//	Produces:
//	- application/json
//
//	Schemes: http, https
//
//	Responses:
//	  200: getCheckResponse
//	  400: genericError
//	  403: getCheckResponse
//	  500: genericError
func (h *Handler) getCheckMirrorStatus(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	allowed, err := h.getCheck(r.Context(), r.URL.Query())
	if err != nil {
		h.d.Writer().WriteError(w, r, err)
		return
	}

	if allowed {
		h.d.Writer().Write(w, r, &RESTResponse{Allowed: allowed})
		return
	}

	h.d.Writer().WriteCode(w, r, http.StatusForbidden, &RESTResponse{Allowed: allowed})
}

func (h *Handler) getCheck(ctx context.Context, q url.Values) (bool, error) {
	maxDepth, err := x.GetMaxDepthFromQuery(q)
	if err != nil {
		return false, err
	}

	tuple, err := (&ketoapi.RelationTuple{}).FromURLQuery(q)
	if err != nil {
		return false, err
	}

	it, err := h.d.Mapper().FromTuple(ctx, tuple)
	// herodot.ErrNotFound occurs when the namespace is unknown
	if errors.Is(err, herodot.ErrNotFound) {
		return false, nil
	} else if err != nil {
		return false, err
	}

	return h.d.PermissionEngine().CheckIsMember(ctx, it[0], maxDepth)
}

// swagger:route POST /relation-tuples/check/openapi read postCheck
//
// # Check a relation tuple
//
// To learn how relation tuples and the check works, head over to [the documentation](../concepts/relation-tuples.mdx).
//
//	Consumes:
//	-  application/json
//
//	Produces:
//	- application/json
//
//	Schemes: http, https
//
//	Responses:
//	  200: getCheckResponse
//	  400: genericError
//	  500: genericError
func (h *Handler) postCheckNoStatus(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	allowed, err := h.postCheck(r.Context(), r.Body, r.URL.Query())
	if err != nil {
		h.d.Writer().WriteError(w, r, err)
		return
	}
	h.d.Writer().Write(w, r, &RESTResponse{Allowed: allowed})
}

// swagger:route POST /relation-tuples/check read postCheckMirrorStatus
//
// # Check a relation tuple
//
// To learn how relation tuples and the check works, head over to [the documentation](../concepts/relation-tuples.mdx).
//
//	Consumes:
//	-  application/json
//
//	Produces:
//	- application/json
//
//	Schemes: http, https
//
//	Responses:
//	  200: getCheckResponse
//	  400: genericError
//	  403: getCheckResponse
//	  500: genericError
func (h *Handler) postCheckMirrorStatus(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	allowed, err := h.postCheck(r.Context(), r.Body, r.URL.Query())
	if err != nil {
		h.d.Writer().WriteError(w, r, err)
		return
	}

	if allowed {
		h.d.Writer().Write(w, r, &RESTResponse{Allowed: allowed})
		return
	}

	h.d.Writer().WriteCode(w, r, http.StatusForbidden, &RESTResponse{Allowed: allowed})
}

func (h *Handler) postCheck(ctx context.Context, body io.Reader, query url.Values) (bool, error) {
	maxDepth, err := x.GetMaxDepthFromQuery(query)
	if err != nil {
		return false, err
	}

	var tuple ketoapi.RelationTuple
	if err := json.NewDecoder(body).Decode(&tuple); err != nil {
		return false, herodot.ErrBadRequest.WithErrorf("could not unmarshal json: %s", err.Error())
	}
	t, err := h.d.Mapper().FromTuple(ctx, &tuple)
	// herodot.ErrNotFound occurs when the namespace is unknown
	if errors.Is(err, herodot.ErrNotFound) {
		return false, nil
	} else if err != nil {
		return false, err
	}

	return h.d.PermissionEngine().CheckIsMember(ctx, t[0], maxDepth)
}

func (h *Handler) Check(ctx context.Context, req *rts.CheckRequest) (*rts.CheckResponse, error) {
	var src ketoapi.TupleData
	if req.Tuple != nil {
		src = req.Tuple
	} else {
		src = req
	}

	tuple, err := (&ketoapi.RelationTuple{}).FromDataProvider(src)
	if err != nil {
		return nil, err
	}

	internalTuple, err := h.d.Mapper().FromTuple(ctx, tuple)
	if err != nil {
		return nil, err
	}
	allowed, err := h.d.PermissionEngine().CheckIsMember(ctx, internalTuple[0], int(req.MaxDepth))
	// TODO add content change handling
	if err != nil {
		return nil, err
	}

	return &rts.CheckResponse{
		Allowed:   allowed,
		Snaptoken: "not yet implemented",
	}, nil
}
