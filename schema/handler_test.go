// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package schema_test

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	opl "github.com/ory/keto/gen/go/ory/keto/opl/v1alpha1"
	"github.com/ory/keto/internal/driver"
	"github.com/ory/keto/schema"
	"github.com/ory/x/httprouterx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tidwall/gjson"
	"google.golang.org/grpc"
)

func TestNewHandler(t *testing.T) {
	reg := driver.NewSqliteTestRegistry(t)
	h := schema.NewHandler(reg)

	r := httprouterx.NewRouterPublic()
	h.RegisterSyntaxRoutes(r)

	ts, clientTLS := driver.HTTP2TestServer(t, r)

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
		conn, err := grpc.NewClient(ts.Listener.Addr().String(), grpc.WithTransportCredentials(clientTLS))
		require.NoError(t, err)

		client := opl.NewSyntaxServiceClient(conn)

		t.Run("method=Syntax.Check", func(t *testing.T) {
			response, err := client.Check(t.Context(), &opl.CheckRequest{
				Content: []byte("/* comment???"),
			})
			require.NoError(t, err)
			assert.Contains(t, response.ParseErrors[0].Message, "unclosed comment")
		})
	})
}
