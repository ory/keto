package driver

import (
	"github.com/ory/viper"
	"github.com/ory/x/logrusx"
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/ory/keto/internal/driver/configuration"
)

type DefaultDriver struct {
	c configuration.Provider
	r Registry
}

func NewDefaultDriver(l *logrusx.Logger, version, build, date string) Driver {
	c := configuration.NewViperProvider(l)
	configuration.MustValidate(l, c)

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

	return &DefaultDriver{r: r, c: c}
}

func NewMemoryTestDriver(t *testing.T) Driver {
	l := logrusx.New("keto", "test")

	c := configuration.NewViperProvider(l)
	viper.Set(configuration.ViperKeyDSN, configuration.DSNMemory)

	r, err := NewRegistry(c)
	require.NoError(t, err)

	require.NoError(t, r.Init())

	return &DefaultDriver{r: r, c: c}
}

func (r *DefaultDriver) Configuration() configuration.Provider {
	return r.c
}

func (r *DefaultDriver) Registry() Registry {
	return r.r
}
