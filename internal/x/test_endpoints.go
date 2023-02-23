// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package x

import (
	"context"
	"net"
	"net/http/httptest"
	"testing"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

type (
	TestEndpoints struct {
		GRPC *grpc.ClientConn
		HTTP *httptest.Server
	}
	readHandler interface {
		RegisterReadGRPC(s *grpc.Server)
		RegisterReadGRPCGatewayConn(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error
	}
	writeHandler interface {
		RegisterWriteGRPC(s *grpc.Server)
		RegisterWriteGRPCGatewayConn(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error
	}
	oplSyntaxHandler interface {
		RegisterSyntaxGRPC(s *grpc.Server)
		RegisterSyntaxGRPCGatewayConn(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error
	}
)

func NewTestEndpoints(
	t *testing.T,
	handler any,
) *TestEndpoints {
	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)

	l := bufconn.Listen(1024 * 1024)
	t.Cleanup(func() { _ = l.Close() })
	s := grpc.NewServer(
		grpc.ChainUnaryInterceptor(GlobalGRPCUnaryServerInterceptors...),
	)

	if h, ok := handler.(readHandler); ok {
		h.RegisterReadGRPC(s)
	}
	if h, ok := handler.(writeHandler); ok {
		h.RegisterWriteGRPC(s)
	}
	if h, ok := handler.(oplSyntaxHandler); ok {
		h.RegisterSyntaxGRPC(s)
	}

	go func() {
		if err := s.Serve(l); err != nil {
			t.Logf("Server exited with error: %v", err)
		}
	}()
	t.Cleanup(s.Stop)

	conn, err := grpc.Dial("bufnet",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) { return l.Dial() }),
	)
	require.NoError(t, err)
	t.Cleanup(func() { _ = conn.Close() })

	mux := runtime.NewServeMux(GRPCGatewayMuxOptions...)
	if h, ok := handler.(readHandler); ok {
		require.NoError(t, h.RegisterReadGRPCGatewayConn(ctx, mux, conn))
	}
	if h, ok := handler.(writeHandler); ok {
		require.NoError(t, h.RegisterWriteGRPCGatewayConn(ctx, mux, conn))
	}
	if h, ok := handler.(oplSyntaxHandler); ok {
		require.NoError(t, h.RegisterSyntaxGRPCGatewayConn(ctx, mux, conn))
	}
	ts := httptest.NewServer(mux)
	t.Cleanup(ts.Close)

	return &TestEndpoints{
		HTTP: ts,
		GRPC: conn,
	}
}
