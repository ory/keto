package driver

import (
	"context"
	"fmt"
	"net"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ory/keto/internal/driver/config"
)

type GetAddr = func(t testing.TB, endpoint string) (host string, port string, fullAddr string)

func UseDynamicPorts(ctx context.Context, t testing.TB, r Registry) GetAddr {
	t.Helper()

	listenDir := t.TempDir()
	readListenFile := fmt.Sprintf("%s/read.addr", listenDir)
	writeListenFile := fmt.Sprintf("%s/write.addr", listenDir)
	metricsListenFile := fmt.Sprintf("%s/metrics.addr", listenDir)
	oplListenFile := fmt.Sprintf("%s/opl.addr", listenDir)

	require.NoError(t, r.Config(ctx).Set(config.KeyReadAPIPort, 0))
	require.NoError(t, r.Config(ctx).Set(config.KeyReadAPIListenFile, "file://"+readListenFile))
	require.NoError(t, r.Config(ctx).Set(config.KeyWriteAPIPort, 0))
	require.NoError(t, r.Config(ctx).Set(config.KeyWriteAPIListenFile, "file://"+writeListenFile))
	require.NoError(t, r.Config(ctx).Set(config.KeyMetricsPort, 0))
	require.NoError(t, r.Config(ctx).Set(config.KeyMetricsListenFile, "file://"+metricsListenFile))
	require.NoError(t, r.Config(ctx).Set(config.KeyOPLSyntaxAPIPort, 0))
	require.NoError(t, r.Config(ctx).Set(config.KeyOPLSyntaxListenFile, "file://"+oplListenFile))

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
