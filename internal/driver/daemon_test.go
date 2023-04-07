// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package driver

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/phayes/freeport"
	"github.com/prometheus/common/expfmt"
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
	ioprometheusclient "github.com/prometheus/client_model/go"
	"github.com/stretchr/testify/require"
)

func TestScrapingEndpoint(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	port, err := freeport.GetFreePort()
	require.NoError(t, err)

	r := NewSqliteTestRegistry(t, false)
	require.NoError(t, r.Config(ctx).Set(config.KeyWriteAPIPort, port))

	//metrics port
	portMetrics, err := freeport.GetFreePort()
	require.NoError(t, err)
	require.NoError(t, r.Config(ctx).Set(config.KeyMetricsPort, portMetrics))

	eg := errgroup.Group{}
	doneShutdown := make(chan struct{})
	eg.Go(r.serveWrite(ctx, doneShutdown))
	eg.Go(r.serveMetrics(ctx, doneShutdown))

	conn, err := grpc.DialContext(ctx, fmt.Sprintf("127.0.0.1:%d", port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err)
	defer conn.Close()

	cl := grpcHealthV1.NewHealthClient(conn)
	watcher, err := cl.Watch(ctx, &grpcHealthV1.HealthCheckRequest{})
	require.NoError(t, err)
	require.NoError(t, watcher.CloseSend())
	for err := status.Error(codes.Unavailable, "init"); status.Code(err) != codes.Unavailable; _, err = watcher.Recv() {
	}

	promresp, err := http.Get(fmt.Sprintf("http://127.0.0.1:%d", portMetrics) + prometheus.MetricsPrometheusPath)
	require.NoError(t, err)
	require.EqualValues(t, http.StatusOK, promresp.StatusCode)

	textParser := expfmt.TextParser{}
	text, err := textParser.TextToMetricFamilies(promresp.Body)
	require.NoError(t, err)
	require.EqualValues(t, "grpc_server_handled_total", *text["grpc_server_handled_total"].Name)
	require.EqualValues(t, "Check", getLabelValue("grpc_method", text["grpc_server_handled_total"].Metric))
	require.EqualValues(t, "grpc.health.v1.Health", getLabelValue("grpc_service", text["grpc_server_handled_total"].Metric))

	require.EqualValues(t, "grpc_server_msg_sent_total", *text["grpc_server_msg_sent_total"].Name)
	require.EqualValues(t, "Check", getLabelValue("grpc_method", text["grpc_server_msg_sent_total"].Metric))
	require.EqualValues(t, "grpc.health.v1.Health", getLabelValue("grpc_service", text["grpc_server_msg_sent_total"].Metric))

	require.EqualValues(t, "grpc_server_msg_received_total", *text["grpc_server_msg_received_total"].Name)
	require.EqualValues(t, "Check", getLabelValue("grpc_method", text["grpc_server_msg_received_total"].Metric))
	require.EqualValues(t, "grpc.health.v1.Health", getLabelValue("grpc_service", text["grpc_server_msg_received_total"].Metric))

	cancel()
	<-doneShutdown
	<-doneShutdown
	require.NoError(t, eg.Wait())
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

func getLabelValue(name string, metric []*ioprometheusclient.Metric) string {
	for _, label := range metric[0].Label {
		if *label.Name == name {
			return *label.Value
		}
	}

	return ""
}
