package e2e

import (
	"fmt"
	"testing"

	"github.com/ory/herodot"

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
	transactClient interface {
		client
		transactTuples(t require.TestingT, ins []*relationtuple.InternalRelationTuple, del []*relationtuple.InternalRelationTuple)
	}
	client interface {
		createTuple(t require.TestingT, r *relationtuple.InternalRelationTuple)
		deleteTuple(t require.TestingT, r *relationtuple.InternalRelationTuple)
		queryTuple(t require.TestingT, q *relationtuple.RelationQuery, opts ...x.PaginationOptionSetter) *relationtuple.GetResponse
		queryTupleErr(t require.TestingT, expected herodot.DefaultError, q *relationtuple.RelationQuery, opts ...x.PaginationOptionSetter)
		check(t require.TestingT, r *relationtuple.InternalRelationTuple) bool
		expand(t require.TestingT, r *relationtuple.SubjectSet, depth int) *expand.Tree
		waitUntilLive(t require.TestingT)
	}
)

func Test(t *testing.T) {
	for _, dsn := range x.GetDSNs(t) {
		t.Run(fmt.Sprintf("dsn=%s", dsn.Name), func(t *testing.T) {
			nspaces := []*namespace.Namespace{
				{
					Name: "dreams0",
					ID:   0,
				},
				{
					Name: "dreams1",
					ID:   1,
				},
				{
					Name: "dreams2",
					ID:   2,
				},
				{
					Name: "dreams3",
					ID:   3,
				},
			}

			ctx, reg := newInitializedReg(t, dsn, nspaces)

			closeServer := startServer(ctx, t, reg)
			defer closeServer()

			// The test cases start here
			// We execute every test with the GRPC client (using the client commands) and REST client
			for ci, cl := range []client{
				&grpcClient{
					readRemote:  reg.Config().ReadAPIListenOn(),
					writeRemote: reg.Config().WriteAPIListenOn(),
					ctx:         ctx,
				},
				&restClient{
					readURL:  "http://" + reg.Config().ReadAPIListenOn(),
					writeURL: "http://" + reg.Config().WriteAPIListenOn(),
				},
				&cliClient{c: &cmdx.CommandExecuter{
					New:            cmd.NewRootCmd,
					Ctx:            ctx,
					PersistentArgs: []string{"--" + cliclient.FlagReadRemote, reg.Config().ReadAPIListenOn(), "--" + cliclient.FlagWriteRemote, reg.Config().WriteAPIListenOn(), "--" + cmdx.FlagFormat, string(cmdx.FormatJSON)},
				}},
				&sdkClient{
					readRemote:  reg.Config().ReadAPIListenOn(),
					writeRemote: reg.Config().WriteAPIListenOn(),
				},
			} {
				t.Run(fmt.Sprintf("client=%T", cl), runCases(cl, nspaces[ci:ci+1]))
			}
		})
	}
}
