package e2e

import (
	"context"
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"

	"github.com/ory/x/cmdx"

	"github.com/ory/keto/cmd/migrate"
	"github.com/ory/keto/internal/namespace"

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

func migrateEverythingUp(t *testing.T, c *cmdx.CommandExecuter, nn []*namespace.Namespace) {
	out := c.ExecNoErr(t, "migrate", "status")
	if strings.Contains(out, "Pending") {
		c.ExecNoErr(t, "migrate", "up", "--"+migrate.FlagYes)
	}

	for _, n := range nn {
		c.ExecNoErr(t, "namespace", "migrate", "up", n.Name)
	}

	t.Cleanup(func() {
		for _, n := range nn {
			c.ExecNoErr(t, "namespace", "migrate", "down", n.Name, "1")
		}

		c.ExecNoErr(t, "migrate", "down", "1")
	})
}
