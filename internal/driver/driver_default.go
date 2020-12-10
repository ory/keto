package driver

import (
	"testing"

	"github.com/ory/keto/internal/namespace"

	"github.com/ory/x/configx"
	"github.com/ory/x/logrusx"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/require"

	"github.com/ory/keto/internal/driver/config"
)

type DefaultDriver struct {
	c config.Provider
	r Registry
}

func NewDefaultRegistry(l *logrusx.Logger, flags *pflag.FlagSet, version, build, date string) Registry {
	c, err := config.New(flags, l)
	if err != nil {
		l.WithError(err).Fatal("Unable to initialize config provider.")
	}

	r, err := NewRegistry(c)
	if err != nil {
		l.WithError(err).Fatal("Unable to instantiate service registry.")
	}

	r.
		WithConfig(c).
		WithLogger(l).
		WithBuildInfo(version, build, date)

	if err = r.Init(); err != nil {
		l.WithError(err).Fatal("Unable to initialize service registry.")
	}

	return r
}

func NewMemoryTestRegistry(t *testing.T, namespaces []*namespace.Namespace) Registry {
	l := logrusx.New("keto", "test")

	flags := pflag.NewFlagSet("test flags", pflag.ContinueOnError)
	configx.RegisterFlags(flags)
	require.NoError(t, flags.Set("config", ""))

	c, err := config.New(flags, l)
	require.NoError(t, err)
	c.Set(config.KeyDSN, config.DSNMemory)
	c.Set(config.KeyNamespaces, namespaces)

	r, err := NewRegistry(c)
	require.NoError(t, err)

	require.NoError(t, r.Init())

	return r
}

func (r *DefaultDriver) Configuration() config.Provider {
	return r.c
}

func (r *DefaultDriver) Registry() Registry {
	return r.r
}
