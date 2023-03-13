// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package namespace_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/ory/herodot"
	"github.com/ory/x/pointerx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ory/keto/internal/driver/config"
	"github.com/ory/keto/internal/namespace"
	"github.com/ory/keto/internal/namespace/ast"
)

func TestASTRelationFor(t *testing.T) {
	ctx := context.Background()
	nm := config.NewMemoryNamespaceManager(&namespace.Namespace{
		Name:      "test",
		Relations: []ast.Relation{{Name: "test"}},
	})

	rel, err := namespace.ASTRelationFor(ctx, nm, "test", "test")
	require.NoError(t, err)
	assert.Equal(t, "test", rel.Name)

	_, err = namespace.ASTRelationFor(ctx, nm, "test", "unknown")
	herodotErr := herodot.ErrBadRequest
	require.ErrorAs(t, err, pointerx.Ptr(&herodotErr))
	assert.Equal(t, http.StatusBadRequest, herodotErr.CodeField)

	rel, err = namespace.ASTRelationFor(ctx, nm, "unknown", "")
	require.NoError(t, err)
	assert.Nil(t, rel)
}
