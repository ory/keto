package sql_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/ory/keto/internal/x/dbx"

	"github.com/ory/keto/internal/driver"
	"github.com/ory/keto/internal/persistence/sql"

	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/require"

	"github.com/ory/keto/internal/driver/config"
	"github.com/ory/keto/internal/namespace"
	"github.com/ory/keto/internal/relationtuple"
)

func TestPersister(t *testing.T) {
	setup := func(t *testing.T, dsn *dbx.DsnT) (p *sql.Persister, d driver.Registry, hook *test.Hook) {
		d = driver.NewSqliteTestRegistry(t, false)

		p, err := sql.NewPersister(dsn.Conn, d, nil)
		require.NoError(t, err)

		mb, err := p.MigrationBox(context.Background())
		require.NoError(t, err)
		require.NoError(t, mb.Up(context.Background()))

		t.Cleanup(func() {
			require.NoError(t, mb.Down(context.Background(), 0))
		})

		return
	}

	addNamespace := func(d driver.Registry, nspaces []*namespace.Namespace) func(context.Context, *testing.T, string) {
		return func(ctx context.Context, t *testing.T, name string) {
			n := &namespace.Namespace{
				Name: name,
				ID:   int64(len(nspaces)),
			}
			nspaces = append(nspaces, n)

			require.NoError(t, d.Config().Set(config.KeyNamespaces, nspaces))
		}
	}

	for _, dsn := range dbx.GetDSNs(t, false) {
		t.Run(fmt.Sprintf("dsn=%s", dsn.Name), func(t *testing.T) {
			var nspaces []*namespace.Namespace
			p, d, _ := setup(t, dsn)

			t.Run("relationtuple.ManagerTest", func(t *testing.T) {
				relationtuple.ManagerTest(t, p, addNamespace(d, nspaces))
			})
		})
	}
}
