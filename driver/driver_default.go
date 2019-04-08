package driver

import (
	"github.com/sirupsen/logrus"

	"github.com/ory/keto/driver/configuration"
)

type DefaultDriver struct {
	c configuration.Provider
	r Registry
}

func NewDefaultDriver(l logrus.FieldLogger, version, build, date string) Driver {
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

func (r *DefaultDriver) Configuration() configuration.Provider {
	return r.c
}

func (r *DefaultDriver) Registry() Registry {
	return r.r
}
