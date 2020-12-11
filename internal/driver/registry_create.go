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

func NewDefaultRegistry(l *logrusx.Logger, flags *pflag.FlagSet, version, hash, date string) Registry {
	c, err := config.New(flags, l)
	if err != nil {
		l.WithError(err).Fatal("Unable to initialize config provider.")
	}

	r := &RegistryDefault{
		c:       c,
		l:       l,
		version: version,
		hash:    hash,
		date:    date,
	}

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

	r := &RegistryDefault{
		c: c,
		l: l,
	}

	require.NoError(t, r.Init())

	return r
}
