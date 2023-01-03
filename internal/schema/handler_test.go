// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package schema_test

import (
	"bytes"
	"context"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tidwall/gjson"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"

	"github.com/ory/keto/internal/driver"
	"github.com/ory/keto/internal/schema"
	"github.com/ory/keto/internal/x"
	opl "github.com/ory/keto/proto/ory/keto/opl/v1alpha1"
)

func TestNewHandler(t *testing.T) {
	ctx := context.Background()

	r := &x.OPLSyntaxRouter{Router: httprouter.New()}
	reg := driver.NewSqliteTestRegistry(t, false)
	h := schema.NewHandler(reg)
	h.RegisterSyntaxRoutes(r)
	ts := httptest.NewServer(r)
	t.Cleanup(ts.Close)

	t.Run("proto=REST", func(t *testing.T) {
		t.Run("method=POST /opl/syntax/check", func(t *testing.T) {
			response, err := ts.Client().Post(
				ts.URL+"/opl/syntax/check",
				"text/plain",
				bytes.NewBufferString("/* comment???"))
			require.NoError(t, err)
			require.Equal(t, http.StatusOK, response.StatusCode)
			body, err := io.ReadAll(response.Body)
			require.NoError(t, err)
			t.Log(string(body))
			assert.Contains(t, gjson.GetBytes(body, "errors.0.message").String(), "unclosed comment")
		})
	})

	t.Run("proto=gRPC", func(t *testing.T) {
		l := bufconn.Listen(1024 * 1024)
		s := grpc.NewServer()
		h.RegisterSyntaxGRPC(s)
		go func() {
			if err := s.Serve(l); err != nil {
				t.Logf("Server exited with error: %v", err)
			}
		}()
		t.Cleanup(s.Stop)

		conn, err := grpc.Dial("bufnet",
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) { return l.Dial() }),
		)
		require.NoError(t, err)

		client := opl.NewSyntaxServiceClient(conn)

		t.Run("method=Syntax.Check", func(t *testing.T) {
			response, err := client.Check(ctx, &opl.CheckRequest{
				Content: []byte("/* comment???"),
			})
			require.NoError(t, err)
			assert.Contains(t, response.ParseErrors[0].Message, "unclosed comment")
		})
	})
}
