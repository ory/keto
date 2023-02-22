// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package schema

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	"github.com/ory/keto/internal/x"
	opl "github.com/ory/keto/proto/ory/keto/opl/v1alpha1"
)

type (
	handlerDependencies interface {
		x.LoggerProvider
		x.WriterProvider
	}
	Handler struct {
		d handlerDependencies
	}
)

const RouteBase = "/opl/syntax/check"

func NewHandler(d handlerDependencies) *Handler {
	return &Handler{d: d}
}

func (h *Handler) RegisterSyntaxGRPC(s *grpc.Server) {
	opl.RegisterSyntaxServiceServer(s, h)
}

func (h *Handler) RegisterSyntaxGRPCGateway(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts ...grpc.DialOption) error {
	return opl.RegisterSyntaxServiceHandlerFromEndpoint(ctx, mux, endpoint, opts)
}
func (h *Handler) RegisterSyntaxGRPCGatewayConn(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return opl.RegisterSyntaxServiceHandler(ctx, mux, conn)
}

func (h *Handler) Check(_ context.Context, request *opl.CheckRequest) (*opl.CheckResponse, error) {
	_, parseErrors := Parse(string(request.GetContent()))
	apiErrors := make([]*opl.ParseError, len(parseErrors))
	for i, e := range parseErrors {
		apiErrors[i] = e.ToProto()
	}
	return &opl.CheckResponse{Errors: apiErrors}, nil
}
