package e2e

import (
	"fmt"
	"testing"

	"github.com/ory/keto/internal/x"

	"github.com/stretchr/testify/require"

	cliclient "github.com/ory/keto/cmd/client"
	"github.com/ory/keto/internal/expand"

	"github.com/ory/x/cmdx"

	"github.com/ory/keto/cmd"
	"github.com/ory/keto/internal/namespace"
	"github.com/ory/keto/internal/relationtuple"
)

type (
	client interface {
		createTuple(t require.TestingT, r *relationtuple.InternalRelationTuple)
		queryTuple(t require.TestingT, q *relationtuple.RelationQuery, opts ...x.PaginationOptionSetter) *relationtuple.GetResponse
		check(t require.TestingT, r *relationtuple.InternalRelationTuple) bool
		expand(t require.TestingT, r *relationtuple.SubjectSet, depth int) *expand.Tree
	}
)

func Test(t *testing.T) {
	for _, dsn := range x.GetDSNs(t) {
		t.Run(fmt.Sprintf("dsn=%s", dsn.Name), func(t *testing.T) {
			nspaces := []*namespace.Namespace{{
				Name: "dreams",
				ID:   0,
			}}

			ctx, reg := newInitializedReg(t, dsn, nspaces)

			closeServer := startServer(ctx, t, reg)
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
