// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package driver

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	prometheus "github.com/ory/x/prometheusx"
	"github.com/phayes/freeport"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	grpcHealthV1 "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/status"

	"github.com/ory/keto/internal/driver/config"
	"github.com/ory/keto/internal/x/dbx"
)

const (
	promLogLine = "promhttp_metric_handler_requests_total"
)

func TestMetricsHandler(t *testing.T) {
	for _, dsn := range dbx.GetDSNs(t, false) {
		r := NewTestRegistry(t, dsn)
		handler := r.metricsRouter(context.Background())
		server := httptest.NewServer(handler)
		defer server.Close()

		resp, err := http.Get(server.URL + prometheus.MetricsPrometheusPath)
		require.NoError(t, err)
		require.Equal(t, resp.StatusCode, http.StatusOK)

		body, err := io.ReadAll(resp.Body)
		require.NoError(t, err)
		require.Contains(t, string(body), promLogLine)
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

	conn, err := grpc.DialContext(ctx, fmt.Sprintf("127.0.0.1:%d", port),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
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
