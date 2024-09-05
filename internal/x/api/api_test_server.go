package api

import (
	"net/http/httptest"
	"strings"
	"testing"

	"connectrpc.com/vanguard"
	"connectrpc.com/vanguard/vanguardgrpc"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/encoding"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/ory/keto/internal/x"
)

type (
	TestServer struct {
		GRPC *grpc.ClientConn
		HTTP *httptest.Server
	}
	readHandler interface {
		RegisterReadGRPC(s *grpc.Server)
	}
	writeHandler interface {
		RegisterWriteGRPC(s *grpc.Server)
	}
	oplSyntaxHandler interface {
		RegisterSyntaxGRPC(s *grpc.Server)
	}
)

func init() {
	encoding.RegisterCodec(vanguardgrpc.NewCodec(&vanguard.JSONCodec{
		MarshalOptions:   protojson.MarshalOptions{EmitUnpopulated: true},
		UnmarshalOptions: protojson.UnmarshalOptions{DiscardUnknown: true},
	}))
}

func NewTestServer(t *testing.T, handler any) *TestServer {
	apiServer := NewServer(WithGRPCOption(grpc.ChainUnaryInterceptor(x.GlobalGRPCUnaryServerInterceptors...)))

	if h, ok := handler.(readHandler); ok {
		h.RegisterReadGRPC(apiServer.GRPCServer)
	}
	if h, ok := handler.(writeHandler); ok {
		h.RegisterWriteGRPC(apiServer.GRPCServer)
	}
	if h, ok := handler.(oplSyntaxHandler); ok {
		h.RegisterSyntaxGRPC(apiServer.GRPCServer)
	}

	h, err := apiServer.Handler()
	require.NoError(t, err)
	ts := httptest.NewUnstartedServer(h)
	ts.EnableHTTP2 = true
	ts.StartTLS()
	ts.TLS.InsecureSkipVerify = true
	t.Cleanup(ts.Close)

	conn, err := grpc.Dial(
		strings.TrimPrefix(ts.URL, "https://"),
		grpc.WithTransportCredentials(credentials.NewTLS(ts.TLS)),
	)
	require.NoError(t, err)
	t.Cleanup(func() { _ = conn.Close() })

	return &TestServer{
		GRPC: conn,
		HTTP: ts,
	}
}
