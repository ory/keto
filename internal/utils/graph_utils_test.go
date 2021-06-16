package utils

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEngineUtilsProvider_CheckVisited(t *testing.T) {
	t.Run("case=finds cycle", func(t *testing.T) {
		ep := &EngineUtilsProvider{}

		linkedList := []string{"A", "B", "C", "B", "D"}

		ctx := context.Background()
		var isThereACycle bool
		for _, n := range linkedList {
			ctx, isThereACycle = ep.CheckVisited(ctx, n)
			if isThereACycle {
				break
			}
		}

		assert.Equal(t, isThereACycle, true)
	})

	t.Run("case=ignores if no cycle", func(t *testing.T) {
		ep := &EngineUtilsProvider{}

		list := []string{"A", "B", "C", "D"}

		ctx := context.Background()
		var isThereACycle bool
		for _, n := range list {
			ctx, isThereACycle = ep.CheckVisited(ctx, n)
			if isThereACycle {
				break
			}
		}

		assert.Equal(t, isThereACycle, false)
	})
}
