// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package check

import (
	"context"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/pkg/errors"

	"github.com/ory/herodot"

	"github.com/ory/keto/internal/x/api"

	"github.com/ory/keto/ketoapi"

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
)

func NewHandler(d handlerDependencies) *Handler {
	return &Handler{d: d}
}

const (
	RouteBase        = "/relation-tuples/check"
	OpenAPIRouteBase = RouteBase + "/openapi"
)

func (h *Handler) RegisterReadGRPC(s *grpc.Server) {
	rts.RegisterCheckServiceServer(s, h)
}

func (h *Handler) RegisterReadGRPCGateway(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts ...grpc.DialOption) error {
	return rts.RegisterCheckServiceHandlerFromEndpoint(ctx, mux, endpoint, opts)
}
func (h *Handler) RegisterReadGRPCGatewayConn(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return rts.RegisterCheckServiceHandler(ctx, mux, conn)
}

func (h *Handler) Check(ctx context.Context, req *rts.CheckRequest) (res *rts.CheckResponse, err error) {
	tuple := (&ketoapi.RelationTuple{}).FromCheckRequest(req)

	// Check if we should set the HTTP status code to 403 instead of 200 if the check fails.

	if api.RequestPath(ctx) == RouteBase {
		defer func() {
			if res != nil && !res.Allowed {
				api.SetStatusCode(ctx, http.StatusForbidden)
			}
		}()
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
