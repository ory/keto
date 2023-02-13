// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package ketoctx

import (
	"context"
	"testing"

	"github.com/gofrs/uuid"
	"github.com/ory/x/configx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDefaultContextualizer(t *testing.T) {
	ctx := context.Background()
	ctxer := &DefaultContextualizer{}

	network := uuid.Must(uuid.NewV4())
	config, err := configx.New(ctx, []byte("{}"))
	require.NoError(t, err)

	assert.Equal(t, network, ctxer.Network(ctx, network))
	assert.Same(t, config, ctxer.Config(ctx, config))
}
