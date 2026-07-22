// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package driver

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"testing"
	"time"

	"connectrpc.com/connect"
	ioprometheusclient "github.com/prometheus/client_model/go"
	"github.com/prometheus/common/expfmt"
	"github.com/prometheus/common/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	grpcHealthV1 "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/status"

	prometheus "github.com/ory/x/prometheusx"

	rts "github.com/ory/keto/gen/go/ory/keto/relation_tuples/v1alpha2"
)

func TestScrapingEndpoint(t *testing.T) {
	t.Parallel()

	r := NewSqliteTestRegistry(t)
	getAddr := UseDynamicPorts(t, r)

	ctx, cancel := context.WithCancel(t.Context())
	defer cancel()

	eg := errgroup.Group{}
	doneShutdown := make(chan struct{})
	eg.Go(r.serveRead(ctx, doneShutdown))
	eg.Go(r.serveMetrics(ctx, doneShutdown))

	_, readPort, _ := getAddr(t, "read")
	_, metricsPort, _ := getAddr(t, "metrics")

	t.Logf("write port: %s, metrics port: %s", readPort, metricsPort)

	assert.EventuallyWithT(t, func(t *assert.CollectT) {
		conn, err := grpc.NewClient(fmt.Sprintf("127.0.0.1:%s", readPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
		require.NoError(t, err)
		defer func() { require.NoError(t, conn.Close()) }()

		cl := grpcHealthV1.NewHealthClient(conn)
		resp, err := cl.Check(ctx, &grpcHealthV1.HealthCheckRequest{})
		require.NoError(t, err)
		assert.Equal(t, grpcHealthV1.HealthCheckResponse_SERVING, resp.Status)
	}, 2*time.Second, 10*time.Millisecond)

	conn, err := grpc.NewClient(fmt.Sprintf("127.0.0.1:%s", readPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err)
	defer func() { require.NoError(t, conn.Close()) }()

	cl := rts.NewReadServiceClient(conn)
	resp, err := cl.ListRelationTuples(ctx, &rts.ListRelationTuplesRequest{RelationQuery: &rts.RelationQuery{}})
	require.NoError(t, err)
	require.Len(t, resp.RelationTuples, 0)

	// A dedicated transport keeps httptest.Server.Close in parallel tests from
	// breaking these requests via http.DefaultTransport.CloseIdleConnections.
	transport := http.DefaultTransport.(*http.Transport).Clone()
	t.Cleanup(transport.CloseIdleConnections)
	hc := &http.Client{Transport: transport}

	respx, err := hc.Get(fmt.Sprintf("http://127.0.0.1:%s/relation-tuples", readPort))
	require.NoError(t, err)
	t.Cleanup(func() { _ = respx.Body.Close() })
	require.EqualValues(t, http.StatusOK, respx.StatusCode)
	body, err := io.ReadAll(respx.Body)
	require.NoError(t, err)
	require.Contains(t, string(body), `"relation_tuples":[]`)

	promresp, err := hc.Get(fmt.Sprintf("http://127.0.0.1:%s%s", metricsPort, prometheus.MetricsPrometheusPath))
	require.NoError(t, err)
	t.Cleanup(func() { _ = promresp.Body.Close() })
	require.EqualValues(t, http.StatusOK, promresp.StatusCode)
	promrespBody, err := io.ReadAll(promresp.Body)
	require.NoError(t, err)

	textParser := expfmt.NewTextParser(model.UTF8Validation)
	text, err := textParser.TextToMetricFamilies(bytes.NewReader(promrespBody))
	require.NoError(t, err)

	require.EqualValuesf(t, "http_requests_total", *text["http_requests_total"].Name, "%s", promrespBody)
	require.ElementsMatch(t,
		[]string{"/grpc.health.v1.Health/", "/ory.keto.relation_tuples.v1alpha2.ReadService/ListRelationTuples", "/relation-tuples"},
		getLabelValues("endpoint", text["http_requests_total"].Metric),
	)

	cancel()
	<-doneShutdown
	<-doneShutdown
	require.NoError(t, eg.Wait())
}

type panicInterceptor struct{}

func (panicInterceptor) WrapUnary(connect.UnaryFunc) connect.UnaryFunc {
	return func(context.Context, connect.AnyRequest) (connect.AnyResponse, error) { panic("test panic") }
}
func (panicInterceptor) WrapStreamingClient(connect.StreamingClientFunc) connect.StreamingClientFunc {
	return func(context.Context, connect.Spec) connect.StreamingClientConn { panic("test panic") }
}
func (panicInterceptor) WrapStreamingHandler(connect.StreamingHandlerFunc) connect.StreamingHandlerFunc {
	return func(context.Context, connect.StreamingHandlerConn) error { panic("test panic") }
}

var _ connect.Interceptor = (*panicInterceptor)(nil)

func TestPanicRecovery(t *testing.T) {
	t.Parallel()

	r := NewSqliteTestRegistry(t, WithHandlerOptions(connect.WithInterceptors(panicInterceptor{})))
	getAddr := UseDynamicPorts(t, r)

	ctx, cancel := context.WithCancel(t.Context())
	defer cancel()

	eg := errgroup.Group{}
	doneShutdown := make(chan struct{})
	eg.Go(r.serveWrite(ctx, doneShutdown))

	_, port, _ := getAddr(t, "write")

	conn, err := grpc.NewClient(fmt.Sprintf("127.0.0.1:%s", port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err)
	defer func() { require.NoError(t, conn.Close()) }()

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
		require.Errorf(t, err, "%+v", resp)
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

func getLabelValues(name string, metric []*ioprometheusclient.Metric) (vals []string) {
	for _, m := range metric {
		for _, label := range m.Label {
			if *label.Name == name {
				vals = append(vals, *label.Value)
			}
		}
	}
	return vals
}
