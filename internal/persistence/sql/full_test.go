package sql_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/ory/x/networkx"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/require"

	"github.com/ory/keto/internal/driver"
	"github.com/ory/keto/internal/persistence/sql"
	"github.com/ory/keto/internal/relationtuple"
	"github.com/ory/keto/internal/x/dbx"
)

func TestPersister(t *testing.T) {
	t.Parallel()

	setup := func(t *testing.T, dsn *dbx.DsnT) (p *sql.Persister, r *driver.RegistryDefault, hook *test.Hook) {
		r = driver.NewTestRegistry(t, dsn)

		p, ok := r.Persister().(*sql.Persister)
		require.True(t, ok)

		require.NoError(t, r.MigrateUp(context.Background()))

		t.Cleanup(func() {
			require.NoError(t, r.MigrateDown(context.Background()))
		})

		return
	}

	for _, dsn := range dbx.GetDSNs(t, false) {
		dsn := dsn
		t.Run(fmt.Sprintf("dsn=%s", dsn.Name), func(t *testing.T) {
			t.Parallel()

			t.Run("relationtuple.ManagerTest", func(t *testing.T) {
				p, _, _ := setup(t, dsn)

				relationtuple.ManagerTest(t, p)
			})

			t.Run("relationtuple.IsolationTest", func(t *testing.T) {
				p0, r, _ := setup(t, dsn)
				n1 := networkx.NewNetwork()
				conn, err := r.PopConnection(context.Background())
				require.NoError(t, err)
				require.NoError(t, conn.Create(n1))
				p1, err := sql.NewPersister(context.Background(), r, n1.ID)
				require.NoError(t, err)

				// same registry, but different persisters only differing in the network ID
				relationtuple.IsolationTest(t, p0, p1)
			})

			t.Run("relationtuple.UUIDMappingManagerTest", func(t *testing.T) {
				p, _, _ := setup(t, dsn)
				relationtuple.UUIDMappingManagerTest(t, p)
			})
		})
	}
}
