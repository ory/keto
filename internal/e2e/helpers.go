package e2e

import (
	"context"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/require"
	"github.com/tidwall/sjson"

	"github.com/ory/keto/internal/driver"
)

func configFile(t *testing.T, values map[string]interface{}) string {
	dir := t.TempDir()
	fn := filepath.Join(dir, "keto.yml")

	c := []byte("{}")
	for key, val := range values {
		var err error
		c, err = sjson.SetBytes(c, key, val)
		require.NoError(t, err)
	}

	require.NoError(t, ioutil.WriteFile(fn, c, 0600))

	return fn
}

func setup(t *testing.T) (*test.Hook, context.Context) {
	hook := &test.Hook{}
	ctx, cancel := context.WithCancel(context.WithValue(context.Background(), driver.LogrusHookContextKey, hook))
	t.Cleanup(func() {
		cancel()

	})

	return hook, ctx
}
