package driver

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/credentials"

	"github.com/ory/keto/internal/driver/config"
)

type GetAddr = func(t testing.TB, endpoint string) (host string, port string, fullAddr string)

func UseDynamicPorts(t testing.TB, r Registry) GetAddr {
	t.Helper()

	listenDir := t.TempDir()
	readListenFile := fmt.Sprintf("%s/read.addr", listenDir)
	writeListenFile := fmt.Sprintf("%s/write.addr", listenDir)
	metricsListenFile := fmt.Sprintf("%s/metrics.addr", listenDir)
	oplListenFile := fmt.Sprintf("%s/opl.addr", listenDir)

	require.NoError(t, r.Config(t.Context()).Set(config.KeyReadAPIPort, 0))
	require.NoError(t, r.Config(t.Context()).Set(config.KeyReadAPIListenFile, "file://"+readListenFile))
	require.NoError(t, r.Config(t.Context()).Set(config.KeyWriteAPIPort, 0))
	require.NoError(t, r.Config(t.Context()).Set(config.KeyWriteAPIListenFile, "file://"+writeListenFile))
	require.NoError(t, r.Config(t.Context()).Set(config.KeyMetricsPort, 0))
	require.NoError(t, r.Config(t.Context()).Set(config.KeyMetricsListenFile, "file://"+metricsListenFile))
	require.NoError(t, r.Config(t.Context()).Set(config.KeyOPLSyntaxAPIPort, 0))
	require.NoError(t, r.Config(t.Context()).Set(config.KeyOPLSyntaxListenFile, "file://"+oplListenFile))

	return func(t testing.TB, endpoint string) (string, string, string) {
		fp := ""
		switch endpoint {
		case "read":
			fp = readListenFile
		case "write":
			fp = writeListenFile
		case "metrics":
			fp = metricsListenFile
		case "opl":
			fp = oplListenFile
		default:
			t.Fatalf("unknown endpoint: %q", endpoint)
		}

		var addr []byte
		var host, port string

		require.EventuallyWithT(t, func(t *assert.CollectT) {
			var err error
			addr, err = os.ReadFile(fp)
			require.NotEmpty(t, addr)
			require.NoError(t, err)
			host, port, err = net.SplitHostPort(string(addr))
			require.NoError(t, err)
			require.NotEmpty(t, host)
			require.NotEmpty(t, port)
		}, 2*time.Second, 10*time.Millisecond)

		return host, port, string(addr)
	}
}

func HTTP2TestServer(t testing.TB, handler http.Handler) (*httptest.Server, credentials.TransportCredentials) {
	ts := httptest.NewUnstartedServer(handler)
	ts.EnableHTTP2 = true
	ts.StartTLS()
	ts.Client()
	t.Cleanup(ts.Close)

	certPool := x509.NewCertPool()
	certPool.AddCert(ts.Certificate())
	return ts, credentials.NewTLS(&tls.Config{RootCAs: certPool})
}
