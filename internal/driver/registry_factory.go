package driver

import (
	"context"
	"testing"

	"github.com/pkg/errors"

	"github.com/sirupsen/logrus/hooks/test"

	"github.com/ory/keto/internal/namespace"

	"github.com/ory/x/logrusx"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/require"

	"github.com/ory/keto/internal/driver/config"
)

func NewDefaultRegistry(ctx context.Context, flags *pflag.FlagSet) (Registry, error) {
	hook, ok := ctx.Value(LogrusHookContextKey).(*test.Hook)

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

func NewMemoryTestRegistry(t *testing.T, namespaces []*namespace.Namespace) Registry {
	l := logrusx.New("ORY Keto", "testing")
	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)

	c, err := config.New(ctx, nil, l)
	require.NoError(t, err)
	require.NoError(t, c.Set(config.KeyDSN, config.DSNMemory))
	require.NoError(t, c.Set(config.KeyNamespaces, namespaces))

	r := &RegistryDefault{
		c: c,
		l: l,
	}

	require.NoError(t, r.Init(ctx))

	return r
}
