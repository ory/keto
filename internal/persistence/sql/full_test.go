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
	setup := func(t *testing.T, dsn *x.DsnT, nspaces []*namespace.Namespace) (p *Persister, hook *test.Hook) {
		c, err := pop.NewConnection(&pop.ConnectionDetails{
			URL: dsn.Conn,
		})
		require.NoError(t, err)

		hook = &test.Hook{}
		lx := logrusx.New("", "", logrusx.WithHook(hook))

		p, err = NewPersister(c, lx, config.NewMemoryNamespaceManager(nspaces...))
		require.NoError(t, err)

		require.NoError(t, p.MigrateUp(context.Background()))
		for _, n := range nspaces {
			require.NoError(t, p.MigrateNamespaceUp(context.Background(), n))
		}

		return
	}

	for _, dsn := range x.GetDSNs(t) {
		t.Run(fmt.Sprintf("dsn=%s", dsn.Name), func(t *testing.T) {
			nspace := &namespace.Namespace{Name: "test"}
			p, _ := setup(t, dsn, []*namespace.Namespace{nspace})

			t.Run("relationtuple.ManagerTest", func(t *testing.T) {
				relationtuple.ManagerTest(t, p, nspace.Name)
			})
		})
	}
}
