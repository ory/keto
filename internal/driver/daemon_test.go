package driver

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"context"

	prometheus "github.com/ory/x/prometheusx"
	"github.com/stretchr/testify/require"

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

		body, err := ioutil.ReadAll(resp.Body)
		require.NoError(t, err)
		require.Contains(t, string(body), promLogLine)
	}
}
