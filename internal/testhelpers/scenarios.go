package testhelpers

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ory/keto/internal/driver"
	"github.com/ory/keto/internal/driver/config"
	"github.com/ory/keto/schema"
)

type Scenario struct {
	Name           string
	Strict         bool
	OrderedShardID bool
	Opl            string
	InputTuples    []string
}

func (s Scenario) WithStrict(v bool) Scenario {
	s.Strict = v
	return s
}

func (s Scenario) WithOrderedShardID() Scenario {
	s.OrderedShardID = true
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
		opts := []driver.TestRegistryOption{
			driver.WithOPL(s.Opl),
			driver.WithMapperNamespace(CustomMapperNamespace),
			driver.WithConfig(config.KeyFeatureFlagStrictMode, s.Strict),
		}
		if s.OrderedShardID {
			opts = append(opts, driver.WithOrderedShardID())
		}
		reg := driver.NewSqliteTestRegistry(t, opts...)

		MapAndInsertTuplesFromString(t, reg, s.InputTuples)
		f(t, reg)
	})
}
