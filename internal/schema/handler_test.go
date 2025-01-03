// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package schema_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ory/keto/internal/driver"
	"github.com/ory/keto/internal/schema"
	"github.com/ory/keto/internal/x/api"
	opl "github.com/ory/keto/proto/ory/keto/opl/v1alpha1"
)

func TestNewHandler(t *testing.T) {
	ctx := context.Background()
	reg := driver.NewSqliteTestRegistry(t, false)

	endpoints := api.NewTestServer(t, schema.NewHandler(reg))

	t.Run("proto=REST", func(t *testing.T) {
		t.Skip("already tested in E2E tests and tricky to unit test due to a special-case middleware for Content-Type: text/plain")
	})

	t.Run("proto=gRPC", func(t *testing.T) {
		client := opl.NewSyntaxServiceClient(endpoints.GRPC)

		t.Run("method=Syntax.Check", func(t *testing.T) {
			response, err := client.Check(ctx, &opl.CheckRequest{
				Content: []byte("/* comment???"),
			})
			require.NoError(t, err)
			assert.Contains(t, response.Errors[0].Message, "unclosed comment")
		})
	})
}
