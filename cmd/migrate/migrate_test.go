package migrate

import (
	"bytes"
	"context"
	"strings"
	"testing"

	"github.com/ory/x/cmdx"
	"github.com/ory/x/configx"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ory/keto/internal/driver"
	"github.com/ory/keto/internal/driver/config"
	"github.com/ory/keto/internal/namespace"
	"github.com/ory/keto/internal/x"
)

func assertAllApplied(t *testing.T, status string) {
	assert.NotContains(t, status, "Pending")
	assert.Contains(t, status, "Applied")
}

func assertNoneApplied(t *testing.T, status string) {
	assert.Contains(t, status, "Pending")
	assert.NotContains(t, status, "Applied")
}

func TestMigrate(t *testing.T) {
	nspaces := []*namespace.Namespace{
		{
			Name: "default",
			ID:   0,
		},
		{
			Name: "other",
			ID:   1,
		},
	}

	newCmd := func(ctx context.Context, persistentArgs ...string) *cmdx.CommandExecuter {
		return &cmdx.CommandExecuter{
			New: func() *cobra.Command {
				cmd := newMigrateCmd()
				configx.RegisterFlags(cmd.PersistentFlags())
				return cmd
			},
			Ctx:            ctx,
			PersistentArgs: persistentArgs,
		}
	}

	for _, dsn := range x.GetDSNs(t) {
		if dsn.Name == "memory" {
			t.Run("dsn=memory", func(t *testing.T) {
				t.Run("case=auto migrates", func(t *testing.T) {
					hook := &test.Hook{}
					ctx := context.WithValue(context.Background(), driver.LogrusHookContextKey, hook)

					cf := x.ConfigFile(t, map[string]interface{}{
						config.KeyDSN:        dsn.Conn,
						config.KeyNamespaces: nspaces,
						"log.level":          "debug",
					})

					cmd := newCmd(ctx, "-c", cf)

					out := cmd.ExecNoErr(t, "up", "--"+FlagYes)
					assert.Contains(t, out, "All migrations are already applied, there is nothing to do.")
				})
			})
		} else {
			t.Run("dsn="+dsn.Name, func(t *testing.T) {
				hook := &test.Hook{}
				ctx := context.WithValue(context.Background(), driver.LogrusHookContextKey, hook)

				cf := x.ConfigFile(t, map[string]interface{}{
					config.KeyDSN:        dsn.Conn,
					config.KeyNamespaces: nspaces,
					"log.level":          "debug",
				})

				cmd := newCmd(ctx, "-c", cf)

				t.Run("case=aborts on no", func(t *testing.T) {
					stdOut, stdErr, err := cmd.Exec(bytes.NewBufferString("n\n"), "up")
					require.NoError(t, err, "%s %s", stdOut, stdErr)

					assert.Contains(t, stdOut, "Pending", "%s %s", stdOut, stdErr)
					assert.NotContains(t, stdOut, "Applied", "%s %s", stdOut, stdErr)
					assert.Contains(t, stdOut, "Aborting", "%s %s", stdOut, stdErr)
				})

				t.Run("case=applies on yes input", func(t *testing.T) {
					stdOut, stdErr, err := cmd.Exec(bytes.NewBufferString("y\n"), "up")
					require.NoError(t, err, "%s %s", stdOut, stdErr)

					t.Cleanup(func() {
						// migrate all down
						t.Log(cmd.ExecNoErr(t, "down", "0", "--"+FlagYes))
					})

					parts := strings.Split(stdOut, "Are you sure that you want to apply this migration?")
					require.Len(t, parts, 2)

					assertNoneApplied(t, parts[0])
					assertAllApplied(t, parts[1])
				})

				t.Run("case=applies on yes flag", func(t *testing.T) {
					out := cmd.ExecNoErr(t, "up", "--"+FlagYes)

					t.Cleanup(func() {
						// migrate all down
						t.Log(cmd.ExecNoErr(t, "down", "0", "--"+FlagYes))
					})

					parts := strings.Split(out, "Applying migrations...")
					require.Len(t, parts, 2)

					assertNoneApplied(t, parts[0])
					assertAllApplied(t, parts[1])
				})

				t.Run("case=applies namespace migrations on flag", func(t *testing.T) {
					out := cmd.ExecNoErr(t, "up", "--"+FlagYes, "--"+FlagAllNamespace)

					t.Cleanup(func() {
						// migrate all down
						t.Log(cmd.ExecNoErr(t, "down", "0", "--"+FlagYes))
					})

					parts := strings.Split(out, "Applying migrations...")
					require.Len(t, parts, 2)
					innerParts := strings.Split(parts[1], "Going to migrate namespaces")
					require.Len(t, innerParts, 2)

					assertNoneApplied(t, parts[0])
					assertAllApplied(t, innerParts[0])

					assert.Contains(t, innerParts[1], "Successfully migrated namespace "+nspaces[0].Name)
					assert.Contains(t, innerParts[1], "Successfully migrated namespace "+nspaces[1].Name)
				})
			})
		}
	}
}
