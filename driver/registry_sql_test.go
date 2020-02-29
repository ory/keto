package driver

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegistrySQL_CanHandle(t *testing.T) {
	for k, tc := range []struct {
		dsn      string
		expected bool
	}{
		{dsn: "memory"},
		{dsn: "mysql://foo:bar@tcp(baz:1234)/db?foo=bar", expected: true},
		{dsn: "postgres://foo:bar@baz:1234/db?foo=bar", expected: true},
	} {
		t.Run(fmt.Sprintf("case=%d", k), func(t *testing.T) {
			r := RegistrySQL{}
			assert.Equal(t, tc.expected, r.CanHandle(tc.dsn))
		})
	}
}
