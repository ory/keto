// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package driver

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	pbTestproto "github.com/grpc-ecosystem/go-grpc-prometheus/examples/testproto"
	"github.com/phayes/freeport"
	"github.com/stretchr/testify/assert"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	grpcHealthV1 "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/status"

	"github.com/ory/keto/internal/driver/config"

	"context"

	prometheus "github.com/ory/x/prometheusx"
	"github.com/stretchr/testify/require"

	"github.com/ory/keto/internal/x/dbx"
)

const (
	promLogLine        = "promhttp_metric_handler_requests_total"
	promGrpcLogLine    = "grpc_server_handled_total"
	pingDefaultValue   = "I like kittens."
	countListResponses = 42
)

func TestMetricsHandler(t *testing.T) {
	for _, dsn := range dbx.GetDSNs(t, false) {
		serverListener, err := net.Listen("tcp", "127.0.0.1:0")
		require.NoError(t, err, "must be able to allocate a port for serverListener")
		// This is the point where we hook up the interceptor
		server := grpc.NewServer(
			grpc.StreamInterceptor(prometheus.StreamServerInterceptor),
			grpc.UnaryInterceptor(prometheus.UnaryServerInterceptor),
		)
		pbTestproto.RegisterTestServiceServer(server, &testService{t})

		go func() {
			server.Serve(serverListener)
		}()

		clientConn, err := grpc.Dial(serverListener.Addr().String(), grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(2*time.Second))
		require.NoError(t, err, "must not error on client Dial")
		testClient := pbTestproto.NewTestServiceClient(clientConn)

		ctx, cancel := context.WithTimeout(context.TODO(), 2*time.Second)

		_, err = testClient.PingList(ctx, &pbTestproto.PingRequest{})
		prometheus.Register(server)

		r := NewTestRegistry(t, dsn)
		handler := r.metricsRouter(context.Background())
		serverHttp := httptest.NewServer(handler)
		defer serverHttp.Close()

		resp, err := http.Get(serverHttp.URL + prometheus.MetricsPrometheusPath)
		require.NoError(t, err)
		require.Equal(t, resp.StatusCode, http.StatusOK)

		body, err := io.ReadAll(resp.Body)
		require.NoError(t, err)
		require.Contains(t, string(body), promLogLine)
		require.Contains(t, string(body), promGrpcLogLine)

		cancel()
		server.Stop()
		serverListener.Close()
	}
}

func TestPanicRecovery(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	unaryPanicInterceptor := func(context.Context, interface{}, *grpc.UnaryServerInfo, grpc.UnaryHandler) (interface{}, error) {
		panic("test panic")
	}
	streamPanicInterceptor := func(context.Context, interface{}, *grpc.UnaryServerInfo, grpc.UnaryHandler) (interface{}, error) {
		panic("test panic")
	}
	port, err := freeport.GetFreePort()
	require.NoError(t, err)

	r := NewSqliteTestRegistry(t, false, WithGRPCUnaryInterceptors(unaryPanicInterceptor), WithGRPCUnaryInterceptors(streamPanicInterceptor))
	require.NoError(t, r.Config(ctx).Set(config.KeyWriteAPIPort, port))

	eg := errgroup.Group{}
	doneShutdown := make(chan struct{})
	eg.Go(r.serveWrite(ctx, doneShutdown))

	conn, err := grpc.DialContext(ctx, fmt.Sprintf("127.0.0.1:%d", port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err)
	defer conn.Close()

	cl := grpcHealthV1.NewHealthClient(conn)

	watcher, err := cl.Watch(ctx, &grpcHealthV1.HealthCheckRequest{})
	require.NoError(t, err)
	require.NoError(t, watcher.CloseSend())
	for err := status.Error(codes.Unavailable, "init"); status.Code(err) != codes.Unavailable; _, err = watcher.Recv() {
	}

	// we want to ensure the server is still running after the panic
	for i := 0; i < 10; i++ {
		// Unary call
		resp, err := cl.Check(ctx, &grpcHealthV1.HealthCheckRequest{})
		require.Error(t, err, "%+v", resp)
		assert.Equal(t, codes.Internal, status.Code(err))

		// Streaming call
		wResp, err := cl.Watch(ctx, &grpcHealthV1.HealthCheckRequest{})
		require.NoError(t, err)
		err = wResp.RecvMsg(nil)
		require.Error(t, err)
		assert.Equal(t, codes.Internal, status.Code(err))
	}

	cancel()
	<-doneShutdown
	require.NoError(t, eg.Wait())
}

type testService struct {
	t *testing.T
}

func (s *testService) PingEmpty(ctx context.Context, _ *pbTestproto.Empty) (*pbTestproto.PingResponse, error) {
	return &pbTestproto.PingResponse{Value: pingDefaultValue, Counter: 42}, nil
}

func (s *testService) Ping(ctx context.Context, ping *pbTestproto.PingRequest) (*pbTestproto.PingResponse, error) {
	// Send user trailers and headers.
	return &pbTestproto.PingResponse{Value: ping.Value, Counter: 42}, nil
}

func (s *testService) PingError(ctx context.Context, ping *pbTestproto.PingRequest) (*pbTestproto.Empty, error) {
	code := codes.Code(ping.ErrorCodeReturned)
	return nil, status.Errorf(code, "Userspace error.")
}

func (s *testService) PingList(ping *pbTestproto.PingRequest, stream pbTestproto.TestService_PingListServer) error {
	if ping.ErrorCodeReturned != 0 {
		return status.Errorf(codes.Code(ping.ErrorCodeReturned), "foobar")
	}
	// Send user trailers and headers.
	for i := 0; i < countListResponses; i++ {
		stream.Send(&pbTestproto.PingResponse{Value: ping.Value, Counter: int32(i)})
	}
	return nil
}
