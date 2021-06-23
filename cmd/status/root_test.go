package status

import (
	"bytes"
	"context"
	"net"
	"strings"
	"testing"
	"time"

	"github.com/ory/x/cmdx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/sync/errgroup"
	grpcHealthV1 "google.golang.org/grpc/health/grpc_health_v1"

	"github.com/ory/keto/cmd/client"
	"github.com/ory/keto/internal/namespace"
)

func TestStatusCmd(t *testing.T) {
	for _, serverType := range []client.ServerType{client.ReadServer, client.WriteServer} {
		t.Run("server_type="+string(serverType), func(t *testing.T) {
			ts := client.NewTestServer(t, serverType, []*namespace.Namespace{{Name: t.Name()}}, newStatusCmd)
			defer ts.Shutdown(t)
			ts.Cmd.PersistentArgs = append(ts.Cmd.PersistentArgs, "--"+cmdx.FlagQuiet, "--"+FlagEndpoint, string(serverType))

			t.Run("case=timeout,noblock", func(t *testing.T) {
				ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
				defer cancel()

				stdOut := cmdx.ExecNoErrCtx(ctx, t, newStatusCmd(), "--"+FlagEndpoint, string(serverType), "--"+ts.FlagRemote, ts.Addr+"0")
				assert.Equal(t, grpcHealthV1.HealthCheckResponse_NOT_SERVING.String()+"\n", stdOut)
			})

			t.Run("case=noblock", func(t *testing.T) {
				stdOut := ts.Cmd.ExecNoErr(t)
				assert.Equal(t, grpcHealthV1.HealthCheckResponse_SERVING.String()+"\n", stdOut)
			})

			t.Run("case=block", func(t *testing.T) {
				l, err := net.Listen("tcp", "127.0.0.1:0")
				require.NoError(t, err)
				s := ts.NewServer()

				startServe := make(chan struct{})

				serveErr := &errgroup.Group{}
				serveErr.Go(func() error {
					// wait until we get the signal to start
					<-startServe
					return s.Serve(l)
				})

				ctx := context.WithValue(context.Background(), client.ContextKeyTimeout, time.Millisecond)

				stdIn, stdErr := &bytes.Buffer{}, &bytes.Buffer{}
				stdOut := &cmdx.CallbackWriter{
					Callbacks: map[string]func([]byte) error{
						// once we get the first retry message, we want to start serving
						"Context deadline exceeded, going to retry.": func([]byte) error {
							// select ensures we only call this if the chan is not already closed
							select {
							case <-startServe:
							default:
								close(startServe)
							}
							return nil
						},
					},
				}

				require.NoError(t,
					cmdx.ExecBackgroundCtx(ctx, newStatusCmd(), stdIn, stdOut, stdErr,
						"--"+FlagEndpoint, string(serverType),
						"--"+ts.FlagRemote, l.Addr().String(),
						"--"+FlagBlock,
					).Wait(),
				)

				fullStdOut := stdOut.String()
				assert.True(t, strings.HasSuffix(fullStdOut, grpcHealthV1.HealthCheckResponse_SERVING.String()+"\n"), fullStdOut)

				s.GracefulStop()
				require.NoError(t, serveErr.Wait())
			})
		})
	}
}
