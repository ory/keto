// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package driver

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/prometheus/common/expfmt"
	"github.com/stretchr/testify/assert"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	grpcHealthV1 "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/status"

	"context"

	prometheus "github.com/ory/x/prometheusx"
	ioprometheusclient "github.com/prometheus/client_model/go"
	"github.com/stretchr/testify/require"
)

func TestScrapingEndpoint(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	r := NewSqliteTestRegistry(t, false)
	getAddr := UseDynamicPorts(ctx, t, r)

	eg := errgroup.Group{}
	doneShutdown := make(chan struct{})
	eg.Go(func() error {
		return r.serveWrite(ctx, doneShutdown)
	})
	eg.Go(func() error {
		return r.serveMetrics(ctx, doneShutdown)
	})

	_, writePort, _ := getAddr(t, "write")
	_, metricsPort, _ := getAddr(t, "metrics")

	t.Logf("write port: %s, metrics port: %s", writePort, metricsPort)

	assert.EventuallyWithT(t, func(t *assert.CollectT) {
		conn, err := grpc.DialContext(ctx, fmt.Sprintf("127.0.0.1:%s", writePort), grpc.WithTransportCredentials(insecure.NewCredentials()))
		require.NoError(t, err)
		defer conn.Close()

		cl := grpcHealthV1.NewHealthClient(conn)
		watcher, err := cl.Watch(ctx, &grpcHealthV1.HealthCheckRequest{})
		require.NoError(t, err)
		require.NoError(t, watcher.CloseSend())
		for err := status.Error(codes.Unavailable, "init"); status.Code(err) != codes.Unavailable; _, err = watcher.Recv() {
		}
	}, 2*time.Second, 10*time.Millisecond)

	promresp, err := http.Get(fmt.Sprintf("http://127.0.0.1:%s", metricsPort) + prometheus.MetricsPrometheusPath)
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
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	unaryPanicInterceptor := func(context.Context, interface{}, *grpc.UnaryServerInfo, grpc.UnaryHandler) (interface{}, error) {
		panic("test panic")
	}
	streamPanicInterceptor := func(context.Context, interface{}, *grpc.UnaryServerInfo, grpc.UnaryHandler) (interface{}, error) {
		panic("test panic")
	}

	r := NewSqliteTestRegistry(t, false, WithGRPCUnaryInterceptors(unaryPanicInterceptor), WithGRPCUnaryInterceptors(streamPanicInterceptor))
	getAddr := UseDynamicPorts(ctx, t, r)

	eg := errgroup.Group{}
	doneShutdown := make(chan struct{})
	eg.Go(func() error {
		return r.serveWrite(ctx, doneShutdown)
	})

	_, port, _ := getAddr(t, "write")

	conn, err := grpc.DialContext(ctx, fmt.Sprintf("127.0.0.1:%s", port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err)
	defer conn.Close()

	assert.EventuallyWithT(t, func(t *assert.CollectT) {
		cl := grpcHealthV1.NewHealthClient(conn)

		watcher, err := cl.Watch(ctx, &grpcHealthV1.HealthCheckRequest{})
		require.NoError(t, err)
		require.NoError(t, watcher.CloseSend())
		for err := status.Error(codes.Unavailable, "init"); status.Code(err) != codes.Unavailable; _, err = watcher.Recv() {
		}
	}, 2*time.Second, 10*time.Millisecond)

	cl := grpcHealthV1.NewHealthClient(conn)
	// we want to ensure the server is still running after the panic
	for range 10 {
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
