package sql

import (
	"context"
	"strings"
	"testing"

	"github.com/ory/x/logrusx"
	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ory/keto/internal/driver/config"
	"github.com/ory/keto/internal/namespace"
	"github.com/ory/keto/internal/x"
)

func TestNamespaceMigrations(t *testing.T) {
	setup := func(t *testing.T, dsn *x.DsnT, nn ...*namespace.Namespace) (*Persister, *test.Hook) {
		hook := test.Hook{}
		l := logrusx.New("", "", logrusx.ForceLevel(logrus.DebugLevel), logrusx.WithHook(&hook))
		nm := config.NewMemoryNamespaceManager(nn...)

		p, err := NewPersister(dsn.Conn, l, nm)
		require.NoError(t, err)
		return p, &hook
	}

	for _, dsn := range x.GetDSNs(t) {
		t.Run("dsn="+dsn.Name, func(t *testing.T) {
			t.Run("case=migrates up and down", func(t *testing.T) {
				n := &namespace.Namespace{
					ID:   0,
					Name: "some namespace",
				}

				p, hook := setup(t, dsn, n)

				require.NoError(t, p.MigrateNamespaceUp(context.Background(), n))
				// logs the time it took as a second message
				assert.True(t, strings.HasPrefix(hook.Entries[len(hook.Entries)-2].Message, "Successfully applied"))

				require.NoError(t, p.MigrateNamespaceDown(context.Background(), n, 0))
				// logs the time it took as a second message
				assert.True(t, strings.HasPrefix(hook.Entries[len(hook.Entries)-2].Message, "< "))
			})

			t.Run("case=migrates namespace again", func(t *testing.T) {
				n := &namespace.Namespace{
					ID:   0,
					Name: "some namespace",
				}

				p, _ := setup(t, dsn, n)

				require.NoError(t, p.MigrateNamespaceUp(context.Background(), n))
				require.NoError(t, p.MigrateNamespaceDown(context.Background(), n, 0))
				require.NoError(t, p.MigrateNamespaceUp(context.Background(), n))
			})
		})
	}
}
