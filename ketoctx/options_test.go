// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package ketoctx

import (
	"testing"

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
}
