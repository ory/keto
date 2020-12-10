package config

import (
	"context"
	"testing"

	"github.com/ory/x/logrusx"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ory/keto/internal/namespace"
)

func TestKoanfNamespaceManager(t *testing.T) {
	t.Run("case=creates static namespace manager when they are set", func(t *testing.T) {
		hook := test.Hook{}
		l := logrusx.New("test", "today", logrusx.WithHook(&hook))
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

		p, err := New(pflag.NewFlagSet("test", pflag.ContinueOnError), l)
		require.NoError(t, err)

		p.Set(KeyNamespaces, nn)

		actualNamespaces, err := p.Namespaces(context.Background())
		require.NoError(t, err)

		for _, n := range nn {
			assert.Contains(t, actualNamespaces, n)
		}
	})
}
