// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package ketoctx

import (
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"

	"github.com/stretchr/testify/assert"
)

func TestOptions(t *testing.T) {
	t.Run("case=has default contextualizer", func(t *testing.T) {
		assert.Equal(t, &DefaultContextualizer{}, Options().Contextualizer())
	})

	t.Run("case=overwrites contextualizer", func(t *testing.T) {
		ctxer := &struct {
			DefaultContextualizer
			x string
		}{x: "x"}

		opts := Options(WithContextualizer(ctxer))
		assert.Equal(t, ctxer, opts.Contextualizer())
	})
	t.Run("case=overwrites grpcServerOpts", func(t *testing.T) {
		sp := keepalive.ServerParameters{
			MaxConnectionAge:      time.Second * 30,
			MaxConnectionAgeGrace: time.Second * 10,
		}
		opts := Options(WithGRPCServerOptions(grpc.KeepaliveParams(sp)))
		assert.NotNil(t, opts.grpcServerOptions)
	})
}
