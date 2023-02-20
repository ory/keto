// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package x

import (
	"context"
	"net"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/ory/herodot"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/test/bufconn"
	"google.golang.org/protobuf/proto"
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

func HttpResponseModifier(ctx context.Context, w http.ResponseWriter, p proto.Message) error {
	md, ok := runtime.ServerMetadataFromContext(ctx)
	if !ok {
		return nil
	}

	delete(w.Header(), "Grpc-Metadata-Content-Type")

	if vals := md.HeaderMD.Get("x-http-location"); len(vals) > 0 {
		w.Header().Add("location", vals[0])
	}

	// set http status code
	if vals := md.HeaderMD.Get("x-http-code"); len(vals) > 0 {
		code, err := strconv.Atoi(vals[0])
		if err != nil {
			return err
		}
		// delete the headers to not expose any grpc-metadata in http response
		delete(md.HeaderMD, "x-http-code")
		delete(w.Header(), "Grpc-Metadata-X-Http-Code")
		w.WriteHeader(code)
	}

	return nil
}

func NewTestEndpoints(
	t *testing.T,
	handler any,
) *TestEndpoints {
	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)

	l := bufconn.Listen(1024 * 1024)
	s := grpc.NewServer(
		grpc.ChainUnaryInterceptor(herodot.UnaryErrorUnwrapInterceptor),
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

	// TODO: Sync with router setup in daemon.go
	mux := runtime.NewServeMux(
		runtime.WithForwardResponseOption(HttpResponseModifier),
		runtime.WithMetadata(func(ctx context.Context, req *http.Request) metadata.MD {
			md := make(metadata.MD)
			contentLength, _ := strconv.Atoi(req.Header.Get("Content-Length"))
			md.Set("hasbody", strconv.FormatBool(contentLength > 0))

			return md
		}),
	)
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
