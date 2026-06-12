// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package sql_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/ory/x/networkx"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/require"

	"github.com/ory/keto/internal/driver"
	"github.com/ory/keto/internal/persistence"
	"github.com/ory/keto/internal/persistence/sql"
	relationtupletest "github.com/ory/keto/internal/relationtuple/test"
	"github.com/ory/keto/internal/x/dbx"
)

func TestPersister(t *testing.T) {
	t.Parallel()

	setup := func(t *testing.T, dsn *dbx.DsnT) (p persistence.Persister, r *driver.RegistryDefault, hook *test.Hook) {
		r = driver.NewTestRegistry(t, dsn)
		p = r.Persister()
		require.NotNil(t, p)
		require.NoError(t, r.MigrateUp(context.Background()))
		return
	}

	for _, dsn := range dbx.GetDSNs(t) {
		t.Run(fmt.Sprintf("dsn=%s", dsn.Name), func(t *testing.T) {
			t.Parallel()

			t.Run("relationtuple.ManagerTest", func(t *testing.T) {
				p, _, _ := setup(t, dsn)

				relationtupletest.ManagerTest(t, p)
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
				relationtupletest.IsolationTest(t, p0, p1)
			})

			t.Run("relationtuple.UUIDMappingManagerTest", func(t *testing.T) {
				p, _, _ := setup(t, dsn)
				relationtupletest.MappingManagerTest(t, p)
			})
		})
	}
}
