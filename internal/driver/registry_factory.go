package driver

import (
	"context"
	"fmt"
	"net/url"
	"testing"

	"github.com/sirupsen/logrus"

	"github.com/pkg/errors"

	"github.com/ory/x/logrusx"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/require"

	"github.com/ory/keto/internal/driver/config"
)

func NewDefaultRegistry(ctx context.Context, flags *pflag.FlagSet) (Registry, error) {
	hook, ok := ctx.Value(LogrusHookContextKey).(logrus.Hook)

	var opts []logrusx.Option
	if ok {
		opts = append(opts, logrusx.WithHook(hook))
	}

	l := logrusx.New("ORY Keto", config.Version, opts...)

	c, err := config.New(ctx, flags, l)
	if err != nil {
		return nil, errors.Wrap(err, "unable to initialize config provider")
	}

	r := &RegistryDefault{
		c: c,
		l: l,
	}

	if err = r.Init(ctx); err != nil {
		return nil, errors.Wrap(err, "unable to initialize service registry")
	}

	return r, nil
}

func SqliteTestDSN(t testing.TB, debugDatabaseOnDisk bool) string {
	dsn := fmt.Sprintf("sqlite://file:%s.sqlite?_fk=true&cache=shared", url.PathEscape(t.Name()))
	if !debugDatabaseOnDisk {
		dsn += "&mode=memory"
	}
	return dsn
}

func NewSqliteTestRegistry(t *testing.T, debugDatabaseOnDisk bool) Registry {
	l := logrusx.New("ORY Keto", "testing")
	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)

	c, err := config.New(ctx, nil, l)
	require.NoError(t, err)

	require.NoError(t, c.Set(config.KeyDSN, SqliteTestDSN(t, debugDatabaseOnDisk)))
	require.NoError(t, c.Set("log.level", "debug"))

	r := &RegistryDefault{
		c: c,
		l: l,
	}

	require.NoError(t, r.Init(ctx))
	if debugDatabaseOnDisk {
		mb, err := r.Migrator().MigrationBox(ctx)
		require.NoError(t, err)
		require.NoError(t, mb.Up(ctx))
	}

	t.Cleanup(func() {
		if debugDatabaseOnDisk {
			return
		}

		mb, err := r.Migrator().MigrationBox(ctx)
		require.NoError(t, err)
		require.NoError(t, mb.Down(ctx, -1))
	})

	return r
}
