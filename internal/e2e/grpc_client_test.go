package e2e

import (
	"testing"

	"github.com/ory/x/cmdx"

	"github.com/ory/keto/internal/expand"
	"github.com/ory/keto/internal/relationtuple"
)

type grpcClient struct {
	c *cmdx.CommandExecuter
}

var _ client = &grpcClient{}

func (g *grpcClient) createTuple(t *testing.T, r *relationtuple.InternalRelationTuple) {

	//stdout, stderr, err := c.Exec(t, relationTuple, "relation-tuple", "create", "-", "--"+client.FlagRemoteURL, "127.0.0.1:4467")
	//require.NoError(t, err, "stdout: %s\nstderr: %s", stdout, stderr)
	//assert.Len(t, stderr, 0, stdout)

	panic("implement me")
}

func (g *grpcClient) queryTuple(t *testing.T, q *relationtuple.RelationQuery) []*relationtuple.InternalRelationTuple {
	panic("implement me")
}

func (g *grpcClient) check(t *testing.T, r *relationtuple.InternalRelationTuple) bool {
	panic("implement me")
}

func (g *grpcClient) expand(t *testing.T, r *relationtuple.SubjectSet, depth int) *expand.Tree {
	panic("implement me")
}
