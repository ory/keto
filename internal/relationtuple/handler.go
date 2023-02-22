// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package relationtuple

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	"github.com/ory/keto/internal/x"
	rts "github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2"
)

type (
	handlerDeps interface {
		ManagerProvider
		MapperProvider
		x.LoggerProvider
		x.WriterProvider
		x.TracingProvider
		x.NetworkIDProvider
	}
	handler struct {
		d handlerDeps
	}
)

const (
	ReadRouteBase  = "/relation-tuples"
	WriteRouteBase = "/admin/relation-tuples"
)

func NewHandler(d handlerDeps) *handler {
	return &handler{
		d: d,
	}
}

func (h *handler) RegisterReadGRPC(s *grpc.Server) {
	rts.RegisterReadServiceServer(s, h)
}

func (h *handler) RegisterWriteGRPC(s *grpc.Server) {
	rts.RegisterWriteServiceServer(s, h)
}

func (h *handler) RegisterReadGRPCGateway(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts ...grpc.DialOption) error {
	return rts.RegisterReadServiceHandlerFromEndpoint(ctx, mux, endpoint, opts)
}

func (h *handler) RegisterReadGRPCGatewayConn(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return rts.RegisterReadServiceHandler(ctx, mux, conn)
}

func (h *handler) RegisterWriteGRPCGateway(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts ...grpc.DialOption) error {
	return rts.RegisterWriteServiceHandlerFromEndpoint(ctx, mux, endpoint, opts)
}
func (h *handler) RegisterWriteGRPCGatewayConn(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return rts.RegisterWriteServiceHandler(ctx, mux, conn)
}
