package config

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ory/keto/internal/namespace"
)

func TestNamespaceMemoryManager(t *testing.T) {
	t.Run("method=should reload", func(t *testing.T) {
		n := &namespace.Namespace{Name: "2"}
		nm := NewMemoryNamespaceManager(n)

		assert.False(t, nm.ShouldReload([]*namespace.Namespace{n}))
		assert.False(t, nm.ShouldReload([]*namespace.Namespace{{Name: "2"}}))

		assert.True(t, nm.ShouldReload([]*namespace.Namespace{{Name: "3"}}))
		assert.True(t, nm.ShouldReload("foo"))
	})
}
