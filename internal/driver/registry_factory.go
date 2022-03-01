package driver

import (
	"context"
	"testing"

	"github.com/ory/x/configx"

	"github.com/ory/keto/ketoctx"

	"github.com/ory/keto/internal/x/dbx"

	"github.com/sirupsen/logrus"

	"github.com/pkg/errors"

	"github.com/ory/x/logrusx"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/require"

	"github.com/ory/keto/internal/driver/config"
)

func NewDefaultRegistry(ctx context.Context, flags *pflag.FlagSet, withoutNetwork bool, opts ...ketoctx.Option) (Registry, error) {
	reg, ok := ctx.Value(RegistryContextKey).(Registry)
	if ok {
		return reg, nil
	}

	options := ketoctx.Options(opts...)

	l := options.Logger()
	if l == nil {
		l = newLogger(ctx)
	}

	c := config.New(ctx, l, nil)
	cp, err := config.NewProvider(ctx, flags, c)
	if err != nil {
		return nil, errors.Wrap(err, "unable to initialize config provider")
	}
	c.WithSource(options.Contextualizer().Config(ctx, cp))

	r := &RegistryDefault{
		c:                         c,
		l:                         l,
		ctxer:                     options.Contextualizer(),
		defaultUnaryInterceptors:  options.GRPCUnaryInterceptors(),
		defaultStreamInterceptors: options.GRPCStreamInterceptors(),
		defaultHttpMiddlewares:    options.HTTPMiddlewares(),
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
	l := logrusx.New("Ory Keto", "testing")
	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)

	ctx = configx.ContextWithConfigOptions(ctx, configx.WithValues(map[string]interface{}{
		config.KeyDSN: dsn.Conn,
		"log.level":   "debug",
	}))
	c, err := config.NewDefault(ctx, nil, l)
	require.NoError(t, err)

	r := &RegistryDefault{
		c:     c,
		l:     l,
		ctxer: &ketoctx.DefaultContextualizer{},
	}

	if dsn.MigrateUp {
		require.NoError(t, r.MigrateUp(ctx))
	}

	require.NoError(t, r.Init(ctx))

	t.Cleanup(func() {
		if !dsn.MigrateDown {
			t.Log("Skipping down migration")
			return
		}

		t.Log("Migrating down")
		require.NoError(t, r.MigrateDown(ctx))
	})

	return r
}

func newLogger(ctx context.Context) *logrusx.Logger {
	hook, ok := ctx.Value(LogrusHookContextKey).(logrus.Hook)

	var opts []logrusx.Option
	if ok {
		opts = append(opts, logrusx.WithHook(hook))
	}

	return logrusx.New("Ory Keto", config.Version, opts...)
}
