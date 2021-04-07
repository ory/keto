package sql

import (
	"context"
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"

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
		hook = &test.Hook{}
		lx := logrusx.New("", "", logrusx.WithHook(hook), logrusx.ForceLevel(logrus.TraceLevel))

		p, err := NewPersister(dsn.Conn, lx, config.NewMemoryNamespaceManager(), nil)
		require.NoError(t, err)

		mb, err := p.MigrationBox(context.Background())
		require.NoError(t, err)
		require.NoError(t, mb.Up(context.Background()))

		t.Cleanup(func() {
			require.NoError(t, mb.Down(context.Background(), 0))
		})

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
			mb, err := p.NamespaceMigrationBox(context.Background(), n)
			require.NoError(t, err)
			require.NoError(t, mb.Up(context.Background()))

			t.Cleanup(func() {
				require.NoError(t, mb.Down(context.Background(), 0))
			})
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
