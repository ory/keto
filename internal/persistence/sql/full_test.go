package sql

import (
	"context"
	"fmt"
	"testing"

	"github.com/gobuffalo/pop/v5"
	"github.com/ory/x/logrusx"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/require"

	"github.com/ory/keto/internal/driver/config"
	"github.com/ory/keto/internal/namespace"
	"github.com/ory/keto/internal/relationtuple"
	"github.com/ory/keto/internal/x"
)

func TestPersister(t *testing.T) {
	setup := func(t *testing.T, dsn *x.DsnT) (p *Persister, hook *test.Hook) {
		c, err := pop.NewConnection(&pop.ConnectionDetails{
			URL: dsn.Conn,
		})
		require.NoError(t, err)

		hook = &test.Hook{}
		lx := logrusx.New("", "", logrusx.WithHook(hook))

		p, err = NewPersister(c, lx, config.NewMemoryNamespaceManager())
		require.NoError(t, err)

		require.NoError(t, p.MigrateUp(context.Background()))

		return
	}

	addNamespace := func(p *Persister, nspaces []*namespace.Namespace) func(context.Context, *testing.T, string) {
		return func(ctx context.Context, t *testing.T, name string) {
			n := &namespace.Namespace{
				Name: name,
				ID:   len(nspaces),
			}
			nspaces = append(nspaces, n)

			p.namespaces = config.NewMemoryNamespaceManager(nspaces...)
			require.NoError(t, p.MigrateNamespaceUp(ctx, n))
		}
	}

	for _, dsn := range x.GetDSNs(t) {
		t.Run(fmt.Sprintf("dsn=%s", dsn.Name), func(t *testing.T) {
			var nspaces []*namespace.Namespace
			p, _ := setup(t, dsn)

			t.Run("relationtuple.ManagerTest", func(t *testing.T) {
				relationtuple.ManagerTest(t, p, addNamespace(p, nspaces))
			})
		})
	}
}
