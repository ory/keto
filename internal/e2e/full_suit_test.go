package e2e

import (
	"bytes"
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	cliclient "github.com/ory/keto/cmd/client"
	"github.com/ory/keto/internal/driver"

	"github.com/ory/keto/internal/expand"

	"github.com/ory/x/cmdx"

	"github.com/ory/keto/cmd"
	"github.com/ory/keto/internal/namespace"
	"github.com/ory/keto/internal/relationtuple"
)

type (
	client interface {
		createTuple(t require.TestingT, r *relationtuple.InternalRelationTuple)
		queryTuple(t require.TestingT, q *relationtuple.RelationQuery) []*relationtuple.InternalRelationTuple
		check(t require.TestingT, r *relationtuple.InternalRelationTuple) bool
		expand(t require.TestingT, r *relationtuple.SubjectSet, depth int) *expand.Tree
	}
)

func Test(t *testing.T) {
	dsns := GetDSNs(t)
	dsns[0].Prepare = func(ctx context.Context, t testing.TB, r driver.Registry, nn []*namespace.Namespace) {
		// check if migrations are auto applied for dsn=memory
		status := &bytes.Buffer{}
		require.NoError(t, r.Migrator().MigrationStatus(ctx, status))
		assert.Contains(t, status.String(), "Applied")
		assert.NotContains(t, status.String(), "Pending")

		// TODO
		//nApplied := strings.Count(status.String(), "Applied")
		//t.Cleanup(func() {
		//	// migrate nApplied down
		//	c.ExecNoErr(t, "migrate", "down", fmt.Sprintf("%d", nApplied))
		//})

		for _, n := range nn {
			s, err := r.NamespaceMigrator().NamespaceStatus(ctx, n.ID)
			require.NoError(t, err)
			assert.Equal(t, s.NextVersion, s.CurrentVersion)

			// TODO
			//t.Cleanup(func() {
			//	c.ExecNoErr(t, "namespace", "migrate", "down", n.Name, "1")
			//})
		}
	}

	for _, dsn := range GetDSNs(t) {
		t.Run(fmt.Sprintf("dsn=%s", dsn.Name), func(t *testing.T) {
			nspaces := []*namespace.Namespace{{
				Name: "dreams",
				ID:   0,
			}}

			ctx, reg, closeServer := startServer(t, dsn, nspaces)
			defer closeServer()

			// The test cases start here
			// We execute every test with the GRPC client (using the client commands) and REST client
			for _, cl := range []client{
				&grpcClient{c: &cmdx.CommandExecuter{
					New:            cmd.NewRootCmd,
					Ctx:            ctx,
					PersistentArgs: []string{"--" + cliclient.FlagReadRemote, reg.Config().ReadAPIListenOn(), "--" + cliclient.FlagWriteRemote, reg.Config().WriteAPIListenOn(), "--" + cmdx.FlagFormat, string(cmdx.FormatJSON)},
				}},
				&restClient{
					readURL:  "http://" + reg.Config().ReadAPIListenOn(),
					writeURL: "http://" + reg.Config().WriteAPIListenOn(),
				},
			} {
				t.Run(fmt.Sprintf("client=%T", cl), runCases(cl, nspaces))
			}
		})
	}
}
