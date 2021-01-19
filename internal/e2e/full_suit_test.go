package e2e

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/ory/keto/internal/expand"

	"github.com/ory/x/cmdx"
	"github.com/ory/x/healthx"
	"github.com/ory/x/sqlcon/dockertest"
	"github.com/stretchr/testify/assert"

	"github.com/ory/keto/cmd"
	cliclient "github.com/ory/keto/cmd/client"
	"github.com/ory/keto/internal/driver/config"
	"github.com/ory/keto/internal/namespace"
	"github.com/ory/keto/internal/relationtuple"
)

type (
	dsnT struct {
		name string
		conn string
		pre  func(*testing.T, *cmdx.CommandExecuter, []*namespace.Namespace)
	}
	client interface {
		createTuple(t *testing.T, r *relationtuple.InternalRelationTuple)
		queryTuple(t *testing.T, q *relationtuple.RelationQuery) []*relationtuple.InternalRelationTuple
		check(t *testing.T, r *relationtuple.InternalRelationTuple) bool
		expand(t *testing.T, r *relationtuple.SubjectSet, depth int) *expand.Tree
	}
)

func Test(t *testing.T) {
	// we use a slice of structs here to always have the same execution order
	dsns := []*dsnT{{
		name: "memory",
		conn: "memory",
		pre: func(t *testing.T, c *cmdx.CommandExecuter, nn []*namespace.Namespace) {
			// check if migrations are auto applied for dsn=memory
			out := c.ExecNoErr(t, "migrate", "status")
			assert.Contains(t, out, "Applied")
			assert.NotContains(t, out, "Pending")

			for _, n := range nn {
				out = c.ExecNoErr(t, "namespace", "migrate", "up", n.Name)
				assert.Contains(t, out, "already migrated")
			}
		},
	}}
	if !testing.Short() {
		dsns = append(dsns,
			&dsnT{
				name: "mysql",
				conn: dockertest.RunTestMySQL(t),
				pre:  migrateEverythingUp,
			},
			&dsnT{
				name: "postgres",
				conn: dockertest.RunTestPostgreSQL(t),
				pre:  migrateEverythingUp,
			},
			&dsnT{
				name: "cockroach",
				conn: dockertest.RunTestCockroachDB(t),
				pre:  migrateEverythingUp,
			},
		)
	}

	for _, dsn := range dsns {
		t.Run(fmt.Sprintf("dsn=%s", dsn.name), func(t *testing.T) {
			// We initialize and migrate everything for each DSN once
			_, ctx := setup(t)

			nspaces := []*namespace.Namespace{{
				Name: "dreams",
				ID:   0,
			}}

			c := &cmdx.CommandExecuter{
				New: cmd.NewRootCmd,
				Ctx: ctx,
				PersistentArgs: []string{"--config", configFile(t, map[string]interface{}{
					config.KeyDSN:        dsn.conn,
					config.KeyNamespaces: nspaces,
					"log.level":          "debug",
				})},
			}

			dsn.pre(t, c, nspaces)
			// Initialization done

			// Start the server
			serverCtx, serverCancel := context.WithCancel(ctx)
			serverErr := make(chan error)
			go func() {
				stdOut, stdErr, err := cmdx.ExecCtx(serverCtx, t, cmd.NewRootCmd(), nil, append(c.PersistentArgs, "serve")...)
				if err != nil {
					t.Logf("STD_OUT:\n%s\n\nSTD_ERR:\n%s", stdOut, stdErr)
				}
				serverErr <- err
			}()

			// defer this to make sure it is shutdown on test failure as well
			defer func() {
				// stop the server
				serverCancel()
				// wait for it to stop
				require.NoError(t, <-serverErr)
			}()

			var healthReady = func() error {
				ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
				defer cancel()

				r, err := http.NewRequestWithContext(ctx, "GET", "http://127.0.0.1:4466"+healthx.ReadyCheckPath, nil)
				if err != nil {
					return err
				}
				_, err = http.DefaultClient.Do(r)
				return err
			}
			// wait for /health/ready
			for err := healthReady(); err != nil; err = healthReady() {
				time.Sleep(10 * time.Millisecond)
			}

			// The test cases start here
			// We execute every test with the GRPC client (using the client commands) and REST client
			for _, cl := range []client{
				&grpcClient{c: &cmdx.CommandExecuter{
					New:            cmd.NewRootCmd,
					Ctx:            ctx,
					PersistentArgs: []string{"--" + cliclient.FlagBasicRemote, "127.0.0.1:4466", "--" + cliclient.FlagPrivilegedRemote, "127.0.0.1:4467", "--" + cmdx.FlagFormat, string(cmdx.FormatJSON)},
				}},
				&restClient{
					basicURL:      "http://127.0.0.1:4466",
					privilegedURL: "http://127.0.0.1:4467",
				},
			} {
				t.Run(fmt.Sprintf("client=%T", cl), runCases(cl, nspaces))
			}
		})
	}
}
