package driver

import (
	"context"
	"testing"

	"github.com/ory/keto/internal/x/dbx"

	"github.com/sirupsen/logrus"

	"github.com/pkg/errors"

	"github.com/ory/x/logrusx"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/require"

	"github.com/ory/keto/internal/driver/config"
)

func NewDefaultRegistry(ctx context.Context, flags *pflag.FlagSet, withoutNetwork bool) (Registry, error) {
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

	init := r.Init
	if withoutNetwork {
		init = r.InitWithoutNetworkID
	}
	if err := init(ctx); err != nil {
		return nil, errors.Wrap(err, "unable to initialize service registry")
	}

	return r, nil
}

func NewSqliteTestRegistry(t *testing.T, debugOnDisk bool) *RegistryDefault {
	mode := dbx.SQLiteMemory
	if debugOnDisk {
		mode = dbx.SQLiteDebug
	}
	return NewTestRegistry(t, dbx.GetSqlite(t, mode))
}

func NewTestRegistry(t *testing.T, dsn *dbx.DsnT) *RegistryDefault {
	l := logrusx.New("ORY Keto", "testing")
	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)

	c, err := config.New(ctx, nil, l)
	require.NoError(t, err)

	require.NoError(t, c.Set(config.KeyDSN, dsn.Conn))
	require.NoError(t, c.Set("log.level", "debug"))

	r := &RegistryDefault{
		c: c,
		l: l,
	}

	if dsn.MigrateUp {
		require.NoError(t, r.MigrateUp(ctx))
	}

	require.NoError(t, r.Init(ctx))

	t.Cleanup(func() {
		if dsn.MigrateDown {
			return
		}

		require.NoError(t, r.MigrateDown(ctx))
	})

	return r
}
