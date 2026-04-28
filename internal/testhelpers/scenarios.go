package testhelpers

import (
	"testing"

	"github.com/ory/keto/internal/driver"
	"github.com/ory/keto/internal/driver/config"
)

type Scenario struct {
	Name        string
	Strict      bool
	Opl         string
	InputTuples []string
}

func (s Scenario) WithStrict(v bool) Scenario {
	s.Strict = v
	return s
}

func (s Scenario) Run(t *testing.T, f func(t *testing.T, reg driver.Registry)) {
	name := s.Name
	if s.Strict {
		name += " (strict)"
	}
	t.Run(name, func(t *testing.T) {
		reg := driver.NewSqliteTestRegistry(t, false,
			driver.WithOPL(s.Opl),
			driver.WithMapperNamespace(CustomMapperNamespace),
			driver.WithConfig(config.KeyNamespacesExperimentalStrictMode, s.Strict),
		)

		MapAndInsertTuplesFromString(t, reg, s.InputTuples)
		f(t, reg)
	})
}
