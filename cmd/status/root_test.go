// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package status

import (
	"bytes"
	"context"
	"errors"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"connectrpc.com/connect"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/sync/errgroup"
	grpcHealthV1 "google.golang.org/grpc/health/grpc_health_v1"

	"github.com/ory/x/cmdx"

	"github.com/ory/keto/cmd/client"
	"github.com/ory/keto/internal/driver"
	"github.com/ory/keto/internal/namespace"
)

func TestStatusCmd(t *testing.T) {
	ts := client.NewTestServer(t, []*namespace.Namespace{{Name: t.Name()}}, newStatusCmd)

	for _, serverType := range []client.ServerType{client.ReadServer, client.WriteServer} {
		t.Run("server_type="+string(serverType), func(t *testing.T) {
			ts.Cmd.PersistentArgs = append(ts.Cmd.PersistentArgs, "--"+cmdx.FlagQuiet, "--"+FlagEndpoint, string(serverType))

			t.Run("case=timeout,noblock", func(t *testing.T) {
				ctx, cancel := context.WithDeadline(t.Context(), time.Now().Add(-time.Second))
				defer cancel()

				stdErr := cmdx.ExecExpectedErrCtx(ctx, t, newStatusCmd(), "--"+FlagEndpoint, string(serverType))
				assert.Contains(t, stdErr, "context deadline exceeded")
			})

			t.Run("case=noblock", func(t *testing.T) {
				stdOut := ts.Cmd.ExecNoErr(t)
				assert.Equal(t, grpcHealthV1.HealthCheckResponse_SERVING.String()+"\n", stdOut)
			})

			t.Run("case=block", func(t *testing.T) {
				ctx := context.WithValue(t.Context(), client.ContextKeyTimeout, 100*time.Millisecond)

				s := httptest.NewUnstartedServer(ts.NewHandler(t, serverType))
				s.EnableHTTP2 = true

				startServe := make(chan struct{})

				serveErr := &errgroup.Group{}
				serveErr.Go(func() error {
					// wait until we get the signal to start
					<-startServe
					s.StartTLS()
					return nil
				})

				var stdIn, stdErr bytes.Buffer
				stdOut := cmdx.CallbackWriter{
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
					cmdx.ExecBackgroundCtx(ctx, newStatusCmd(), &stdIn, &stdOut, &stdErr,
						"--"+FlagEndpoint, string(serverType),
						"--"+serverType.FlagName(), s.Listener.Addr().String(),
						"--insecure-skip-hostname-verification=true",
						"--"+client.FlagBlock,
					).Wait(),
				)

				fullStdOut := stdOut.String()
				assert.True(t, strings.HasSuffix(fullStdOut, "\n"+grpcHealthV1.HealthCheckResponse_SERVING.String()+"\n"), fullStdOut)

				require.NoError(t, serveErr.Wait())
			})
		})
	}
}

func authInterceptor(header, validValue string) connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			val := req.Header().Get(header)
			if val == "" {
				return nil, errors.New("not authorized, no header values")
			}
			if val != validValue {
				return nil, errors.New("not authorized, incorrect value")
			}
			return next(ctx, req)
		}
	}
}

func TestAuthorizedRequest(t *testing.T) {
	ts := client.NewTestServer(
		t, []*namespace.Namespace{{Name: t.Name()}}, newStatusCmd,
		driver.WithHandlerOptions(connect.WithInterceptors(authInterceptor("authorization", "Bearer secret"))),
	)

	t.Run("case=not authorized", func(t *testing.T) {
		out := ts.Cmd.ExecExpectedErr(t)
		assert.Contains(t, out, "not authorized")
	})

	t.Run("case=authorized", func(t *testing.T) {
		t.Setenv("KETO_BEARER_TOKEN", "secret")
		out := ts.Cmd.ExecNoErr(t)
		assert.Contains(t, out, "SERVING")
	})
}

func TestAuthorityRequest(t *testing.T) {
	ts := client.NewTestServer(
		t, []*namespace.Namespace{{Name: t.Name()}}, newStatusCmd,
		driver.WithHandlerOptions(connect.WithInterceptors(authInterceptor("Host", "example.com"))),
	)

	t.Run("case=no authority", func(t *testing.T) {
		out := ts.Cmd.ExecExpectedErr(t)
		assert.Contains(t, out, "not authorized")
	})

	t.Run("case=env authority", func(t *testing.T) {
		t.Setenv("KETO_AUTHORITY", "example.com")
		out := ts.Cmd.ExecNoErr(t)
		assert.Contains(t, out, "SERVING")
	})

	t.Run("case=flag authority", func(t *testing.T) {
		out := ts.Cmd.ExecNoErr(t, "--"+client.FlagAuthority, "example.com")
		assert.Contains(t, out, "SERVING")
	})
}
