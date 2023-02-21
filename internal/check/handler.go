// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package check

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/pkg/errors"
	"google.golang.org/grpc/metadata"

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
	_ *checkPermission       = nil
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

func (h *Handler) RegisterReadGRPCGateway(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts ...grpc.DialOption) error {
	return rts.RegisterCheckServiceHandlerFromEndpoint(ctx, mux, endpoint, opts)
}
func (h *Handler) RegisterReadGRPCGatewayConn(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return rts.RegisterCheckServiceHandler(ctx, mux, conn)
}

// Check Permission Result
//
// The content of the allowed field is mirrored in the HTTP status code.
//
// swagger:model checkPermissionResult
type CheckPermissionResult struct {
	// whether the relation tuple is allowed
	//
	// required: true
	Allowed bool `json:"allowed"`
}

// Check Permission Request Parameters
//
// swagger:parameters checkPermission
type checkPermission struct {
	// in: query
	MaxDepth int `json:"max-depth"`
}

// swagger:route GET /relation-tuples/check/openapi permission checkPermission
//
// # Check a permission
//
// To learn how relationship tuples and the check works, head over to [the documentation](https://www.ory.sh/docs/keto/concepts/api-overview).
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
//	  200: checkPermissionResult
//	  400: errorGeneric
//	  default: errorGeneric
func (h *Handler) getCheckNoStatus(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	allowed, err := h.getCheck(r.Context(), r.URL.Query())
	if err != nil {
		h.d.Writer().WriteError(w, r, err)
		return
	}
	h.d.Writer().Write(w, r, &CheckPermissionResult{Allowed: allowed})
}

// Check Permission Or Error Request Parameters
//
// swagger:parameters checkPermissionOrError
// nolint:deadcode,unused
type checkPermissionOrError struct {
	// in: query
	MaxDepth int `json:"max-depth"`
}

// swagger:route GET /relation-tuples/check permission checkPermissionOrError
//
// # Check a permission
//
// To learn how relationship tuples and the check works, head over to [the documentation](https://www.ory.sh/docs/keto/concepts/api-overview).
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
//	  200: checkPermissionResult
//	  400: errorGeneric
//	  403: checkPermissionResult
//	  default: errorGeneric
func (h *Handler) getCheckMirrorStatus(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	allowed, err := h.getCheck(r.Context(), r.URL.Query())
	if err != nil {
		h.d.Writer().WriteError(w, r, err)
		return
	}

	if allowed {
		h.d.Writer().Write(w, r, &CheckPermissionResult{Allowed: allowed})
		return
	}

	h.d.Writer().WriteCode(w, r, http.StatusForbidden, &CheckPermissionResult{Allowed: allowed})
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

	it, err := h.d.ReadOnlyMapper().FromTuple(ctx, tuple)
	// herodot.ErrNotFound occurs when the namespace is unknown
	if errors.Is(err, herodot.ErrNotFound) {
		return false, nil
	} else if err != nil {
		return false, err
	}

	return h.d.PermissionEngine().CheckIsMember(ctx, it[0], maxDepth)
}

// Check Permission using Post Request Parameters
//
// swagger:parameters postCheckPermission
// nolint:deadcode,unused
type postCheckPermission struct {
	// in: query
	MaxDepth int `json:"max-depth"`

	// in: body
	Payload postCheckPermissionBody
}

// Check Permission using Post Request Body
//
// swagger:model postCheckPermissionBody
// nolint:deadcode,unused
type postCheckPermissionBody struct {
	ketoapi.RelationQuery
}

// swagger:route POST /relation-tuples/check/openapi permission postCheckPermission
//
// # Check a permission
//
// To learn how relationship tuples and the check works, head over to [the documentation](https://www.ory.sh/docs/keto/concepts/api-overview).
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
//	  200: checkPermissionResult
//	  400: errorGeneric
//	  default: errorGeneric
func (h *Handler) postCheckNoStatus(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	allowed, err := h.postCheck(r.Context(), r.Body, r.URL.Query())
	if err != nil {
		h.d.Writer().WriteError(w, r, err)
		return
	}
	h.d.Writer().Write(w, r, &CheckPermissionResult{Allowed: allowed})
}

// Post Check Permission Or Error Request Parameters
//
// swagger:parameters postCheckPermissionOrError
// nolint:deadcode,unused
type postCheckPermissionOrError struct {
	// nolint:deadcode,unused
	// in: query
	MaxDepth int `json:"max-depth"`

	// in: body
	Body postCheckPermissionOrErrorBody
}

// Post Check Permission Or Error Body
//
// swagger:model postCheckPermissionOrErrorBody
// nolint:deadcode,unused
type postCheckPermissionOrErrorBody struct {
	ketoapi.RelationQuery
}

// swagger:route POST /relation-tuples/check permission postCheckPermissionOrError
//
// # Check a permission
//
// To learn how relationship tuples and the check works, head over to [the documentation](https://www.ory.sh/docs/keto/concepts/api-overview).
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
//	  200: checkPermissionResult
//	  400: errorGeneric
//	  403: checkPermissionResult
//	  default: errorGeneric
func (h *Handler) postCheckMirrorStatus(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	allowed, err := h.postCheck(r.Context(), r.Body, r.URL.Query())
	if err != nil {
		h.d.Writer().WriteError(w, r, err)
		return
	}

	if allowed {
		h.d.Writer().Write(w, r, &CheckPermissionResult{Allowed: allowed})
		return
	}

	h.d.Writer().WriteCode(w, r, http.StatusForbidden, &CheckPermissionResult{Allowed: allowed})
}

func (h *Handler) postCheck(ctx context.Context, body io.Reader, query url.Values) (bool, error) {
	maxDepth, err := x.GetMaxDepthFromQuery(query)
	if err != nil {
		return false, err
	}

	var tuple ketoapi.RelationTuple
	if err := json.NewDecoder(body).Decode(&tuple); err != nil {
		return false, errors.WithStack(herodot.ErrBadRequest.WithErrorf("could not unmarshal json: %s", err.Error()))
	}
	t, err := h.d.ReadOnlyMapper().FromTuple(ctx, &tuple)
	// herodot.ErrNotFound occurs when the namespace is unknown
	if errors.Is(err, herodot.ErrNotFound) {
		return false, nil
	} else if err != nil {
		return false, err
	}

	return h.d.PermissionEngine().CheckIsMember(ctx, t[0], maxDepth)
}

func (h *Handler) Check(ctx context.Context, req *rts.CheckRequest) (res *rts.CheckResponse, err error) {
	tuple := (&ketoapi.RelationTuple{}).FromCheckRequest(req)

	// Check if we should set the HTTP status code to 403 instead of 200 if the check fails.
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		path := md["path"]
		if len(path) > 0 && path[0] == RouteBase {
			defer func() {
				if res != nil && !res.Allowed {
					_ = grpc.SetHeader(ctx, metadata.Pairs("x-http-code", "403"))
				}
			}()
		}
	}

	if tuple.SubjectID == nil && tuple.SubjectSet == nil {
		return nil, ketoapi.ErrNilSubject
	}

	internalTuple, err := h.d.ReadOnlyMapper().FromTuple(ctx, tuple)
	if errors.Is(err, herodot.ErrNotFound) {
		res = &rts.CheckResponse{Allowed: false}
		return res, nil
	}
	if err != nil {
		return nil, err
	}
	allowed, err := h.d.PermissionEngine().CheckIsMember(ctx, internalTuple[0], int(req.MaxDepth))
	if err != nil {
		return nil, err
	}

	res = &rts.CheckResponse{Allowed: allowed}
	return res, nil
}
