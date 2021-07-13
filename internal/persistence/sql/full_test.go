package sql_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/gofrs/uuid"

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
	setup := func(t *testing.T, dsn *dbx.DsnT) (p *sql.Persister, r *driver.RegistryDefault, hook *test.Hook) {
		r = driver.NewTestRegistry(t, dsn)

		p, err := sql.NewPersister(r, uuid.Must(uuid.NewV4()))
		require.NoError(t, err)

		require.NoError(t, r.MigrateUp(context.Background()))

		t.Cleanup(func() {
			require.NoError(t, r.MigrateDown(context.Background()))
		})

		return
	}

	addNamespace := func(r driver.Registry, nspaces []*namespace.Namespace) func(context.Context, *testing.T, string) {
		return func(ctx context.Context, t *testing.T, name string) {
			n := &namespace.Namespace{
				Name: name,
				ID:   int64(len(nspaces)),
			}
			nspaces = append(nspaces, n)

			require.NoError(t, r.Config().Set(config.KeyNamespaces, nspaces))
		}
	}

	for _, dsn := range dbx.GetDSNs(t, false) {
		t.Run(fmt.Sprintf("dsn=%s", dsn.Name), func(t *testing.T) {
			t.Run("relationtuple.ManagerTest", func(t *testing.T) {
				var nspaces []*namespace.Namespace
				p, r, _ := setup(t, dsn)

				relationtuple.ManagerTest(t, p, addNamespace(r, nspaces))
			})

			t.Run("relationtuple.IsolationTest", func(t *testing.T) {
				var nspaces []*namespace.Namespace
				p0, r, _ := setup(t, dsn)
				p1, err := sql.NewPersister(r, uuid.Must(uuid.NewV4()))
				require.NoError(t, err)

				// same registry, but different persisters only differing in the network ID
				relationtuple.IsolationTest(t, p0, p1, addNamespace(r, nspaces))
			})
		})
	}
}
