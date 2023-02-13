// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package schema_test

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tidwall/gjson"

	"github.com/ory/keto/internal/driver"
	"github.com/ory/keto/internal/schema"
	"github.com/ory/keto/internal/x"
	opl "github.com/ory/keto/proto/ory/keto/opl/v1alpha1"
)

func TestNewHandler(t *testing.T) {
	ctx := context.Background()
	reg := driver.NewSqliteTestRegistry(t, false)

	endpoints := x.NewTestEndpoints(t, schema.NewHandler(reg))

	t.Run("proto=REST", func(t *testing.T) {
		t.Run("method=POST /opl/syntax/check", func(t *testing.T) {
			response, err := endpoints.HTTP.Client().Post(
				endpoints.HTTP.URL+"/opl/syntax/check",
				"text/plain",
				bytes.NewBufferString(`"/* comment???"`))
			require.NoError(t, err)
			body, err := io.ReadAll(response.Body)
			require.NoError(t, err)
			t.Log(string(body))
			require.Equal(t, http.StatusOK, response.StatusCode)
			assert.Contains(t, gjson.GetBytes(body, "errors.0.message").String(), "unclosed comment")
		})
	})

	t.Run("proto=gRPC", func(t *testing.T) {
		client := opl.NewSyntaxServiceClient(endpoints.GRPC)

		t.Run("method=Syntax.Check", func(t *testing.T) {
			response, err := client.Check(ctx, &opl.CheckRequest{
				Content: "/* comment???",
			})
			require.NoError(t, err)
			assert.Contains(t, response.Errors[0].Message, "unclosed comment")
		})
	})
}
