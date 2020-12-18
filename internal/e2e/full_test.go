package e2e

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/ory/keto/cmd/migrate"

	"github.com/ory/x/cmdx"
	"github.com/ory/x/healthx"
	"github.com/ory/x/sqlcon/dockertest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tidwall/gjson"

	"github.com/ory/keto/cmd"
	"github.com/ory/keto/internal/driver/config"
	"github.com/ory/keto/internal/namespace"
	"github.com/ory/keto/internal/relationtuple"
)

type dsnT struct {
	name string
	conn string
	pre  func(*testing.T, *cmdx.CommandExecuter, []*namespace.Namespace)
}

func migrateEverythingUp(t *testing.T, c *cmdx.CommandExecuter, nn []*namespace.Namespace) {
	out := cmdx.ExecNoErrCtx(c.Ctx, t, c.New(), c.PersistentArgs[0], c.PersistentArgs[1], "migrate", "status")
	if strings.Contains(out, "Pending") {
		c.ExecNoErr(t, "migrate", "up", "--"+migrate.FlagYes)
	}

	for _, n := range nn {
		c.ExecNoErr(t, "namespace", "migrate", n.Name)
	}
}

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
				out = c.ExecNoErr(t, "namespace", "migrate", n.Name)
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

			// start the server
			serverCtx, serverCancel := context.WithCancel(ctx)
			serverDoneChannel := make(chan struct{})
			go func() {
				cmdx.ExecNoErrCtx(serverCtx, t, cmd.NewRootCmd(), append(c.PersistentArgs, "serve")...)
				close(serverDoneChannel)
			}()

			// wait for /health/ready
			for _, err := http.Get("http://127.0.0.1:4466" + healthx.ReadyCheckPath); err != nil; _, err = http.Get("http://127.0.0.1:4466" + healthx.ReadyCheckPath) {
				time.Sleep(10 * time.Millisecond)
			}

			lndTuple := &relationtuple.InternalRelationTuple{
				Namespace: nspaces[0].Name,
				Object:    "last nights dream",
				Relation:  "see",
				Subject:   &relationtuple.SubjectID{ID: "me"},
			}
			lndTupleEnc, err := json.Marshal(lndTuple)
			require.NoError(t, err)

			// create a relation tuple -- TODO use CLI commands instead
			relationTuple := &bytes.Buffer{}
			require.NoError(t, json.NewEncoder(relationTuple).Encode(lndTuple))

			//stdout, stderr, err := c.Exec(t, relationTuple, "relation-tuple", "create", "-", "--"+client.FlagRemoteURL, "127.0.0.1:4467")
			//require.NoError(t, err, "stdout: %s\nstderr: %s", stdout, stderr)
			//assert.Len(t, stderr, 0, stdout)

			req, err := http.NewRequest(http.MethodPut, "http://127.0.0.1:4466/relationtuple", bytes.NewBuffer(lndTupleEnc))
			require.NoError(t, err)
			resp, err := http.DefaultClient.Do(req)
			require.NoError(t, err)
			require.NoError(t, resp.Body.Close())
			assert.Equal(t, http.StatusCreated, resp.StatusCode)

			resp, err = http.Get(fmt.Sprintf("http://127.0.0.1:4466/relationtuple?namespace=%s", nspaces[0].Name))
			require.NoError(t, err)
			d, err := ioutil.ReadAll(resp.Body)
			require.NoError(t, err)
			require.NoError(t, resp.Body.Close())
			assert.Equal(t, string(lndTupleEnc), gjson.GetBytes(d, "relations.0").String())

			//relationTuple.Reset()
			//require.NoError(t, json.NewEncoder(relationTuple).Encode(&relationtuple.InternalRelationTuple{
			//	Namespace: namesp.Name,
			//	Object:    "last nights dream",
			//	Relation:  "see",
			//	Subject:   &relationtuple.SubjectID{ID: "nightmare"},
			//}))
			//
			//stdout, stderr, err = c.Exec(t, relationTuple, "relationtuple", "create", "-")
			//require.NoError(t, err)
			//assert.Len(t, stderr, 0, stdout)

			// try the check API to see whether the tuple is interpreted correctly
			resp, err = http.Get(fmt.Sprintf("http://127.0.0.1:4466/check?%s", lndTuple.ToURLQuery().Encode()))
			require.NoError(t, err)
			require.NoError(t, resp.Body.Close())
			assert.Equal(t, http.StatusOK, resp.StatusCode)

			// stop the server
			serverCancel()
			// wait for it to stop
			<-serverDoneChannel
		})
	}
}
