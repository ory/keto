package testhelpers

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ory/keto/internal/driver"
	"github.com/ory/keto/internal/driver/config"
	"github.com/ory/keto/schema"
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
		if s.Opl != "" {
			_, errs := schema.Parse(s.Opl)
			require.Len(t, errs, 0)
		}
		reg := driver.NewSqliteTestRegistry(t, driver.WithOPL(s.Opl), driver.WithMapperNamespace(CustomMapperNamespace), driver.WithConfig(config.KeyFeatureFlagStrictMode, s.Strict))

		MapAndInsertTuplesFromString(t, reg, s.InputTuples)
		f(t, reg)
	})
}
