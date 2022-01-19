package config

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/ory/x/logrusx"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ory/keto/internal/namespace"
)

func TestKoanfNamespaceManager(t *testing.T) {
	setup := func(t *testing.T) (*test.Hook, *Config) {
		hook := test.Hook{}
		l := logrusx.New("test", "today", logrusx.WithHook(&hook))

		ctx, cancel := context.WithCancel(context.Background())
		t.Cleanup(cancel)

		p, err := NewDefault(ctx, pflag.NewFlagSet("test", pflag.ContinueOnError), l)
		require.NoError(t, err)

		return &hook, p
	}

	assertNamespaces := func(t *testing.T, p *Config, nn ...*namespace.Namespace) {
		nm, err := p.NamespaceManager()
		require.NoError(t, err)

		actualNamespaces, err := nm.Namespaces(context.Background())
		require.NoError(t, err)
		assert.Equal(t, len(nn), len(actualNamespaces))

		for _, n := range nn {
			assert.Contains(t, actualNamespaces, n)
		}
	}

	t.Run("case=creates memory namespace manager when namespaces are set", func(t *testing.T) {
		run := func(namespaces []*namespace.Namespace, value interface{}) func(*testing.T) {
			return func(t *testing.T) {
				_, p := setup(t)

				require.NoError(t, p.Set(KeyNamespaces, value))

				assertNamespaces(t, p, namespaces...)

				nm, err := p.NamespaceManager()
				require.NoError(t, err)
				_, ok := nm.(*memoryNamespaceManager)
				assert.True(t, ok)
			}

		}

		nn := []*namespace.Namespace{
			{
				ID:   0,
				Name: "n0",
			},
			{
				ID:   1,
				Name: "n1",
			},
			{
				ID:   2,
				Name: "n2",
			},
		}
		nnJson, err := json.Marshal(nn)
		require.NoError(t, err)
		nnValue := make([]interface{}, 0)
		require.NoError(t, json.Unmarshal(nnJson, &nnValue))

		t.Run(
			"type=[]*namespace.Namespace",
			run(nn, nn),
		)

		t.Run(
			"type=[]interface{}",
			run(nn, nnValue),
		)
	})

	t.Run("case=reloads namespace manager when namespaces are updated using Set()", func(t *testing.T) {
		_, p := setup(t)

		n0 := &namespace.Namespace{
			ID:   0,
			Name: "n0",
		}
		n1 := &namespace.Namespace{
			ID:   1,
			Name: "n1",
		}

		require.NoError(t, p.Set(KeyNamespaces, []*namespace.Namespace{n0}))
		assertNamespaces(t, p, n0)

		require.NoError(t, p.Set(KeyNamespaces, []*namespace.Namespace{n1}))
		assertNamespaces(t, p, n1)
	})

	t.Run("case=creates watcher manager when namespaces is string URL", func(t *testing.T) {
		_, p := setup(t)

		require.NoError(t, p.Set(KeyNamespaces, "file://"+t.TempDir()))

		nm, err := p.NamespaceManager()
		require.NoError(t, err)
		_, ok := nm.(*NamespaceWatcher)
		assert.True(t, ok)
	})
}
