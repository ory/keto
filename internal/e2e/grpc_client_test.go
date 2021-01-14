package e2e

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	cliexpand "github.com/ory/keto/cmd/expand"
	clirelationtuple "github.com/ory/keto/cmd/relationtuple"

	"github.com/ory/x/cmdx"

	"github.com/ory/keto/internal/expand"
	"github.com/ory/keto/internal/relationtuple"
)

type grpcClient struct {
	c *cmdx.CommandExecuter
}

var _ client = &grpcClient{}

func (g *grpcClient) createTuple(t *testing.T, r *relationtuple.InternalRelationTuple) {
	tupleEnc, err := json.Marshal(r)
	require.NoError(t, err)

	stdout, stderr, err := g.c.Exec(t, bytes.NewBuffer(tupleEnc), "relation-tuple", "create", "-")
	require.NoError(t, err, "stdout: %s\nstderr: %s", stdout, stderr)
	assert.Len(t, stderr, 0, stdout)
}

func (g *grpcClient) queryTuple(t *testing.T, q *relationtuple.RelationQuery) []*relationtuple.InternalRelationTuple {
	var flags []string
	if q.Subject != nil {
		flags = append(flags, "--"+clirelationtuple.FlagSubject, q.Subject.String())
	}
	if q.Relation != "" {
		flags = append(flags, "--"+clirelationtuple.FlagRelation, q.Relation)
	}
	if q.Object != "" {
		flags = append(flags, "--"+clirelationtuple.FlagObject, q.Object)
	}

	out := g.c.ExecNoErr(t, append(flags, "relation-tuple", "get", q.Namespace)...)

	var rels []*relationtuple.InternalRelationTuple
	require.NoError(t, json.Unmarshal([]byte(out), &rels), "%s", out)

	return rels
}

func (g *grpcClient) check(t *testing.T, r *relationtuple.InternalRelationTuple) bool {
	out := g.c.ExecNoErr(t, "check", r.Subject.String(), r.Relation, r.Namespace, r.Object)
	res, err := strconv.ParseBool(strings.TrimSpace(out))
	require.NoError(t, err)
	return res
}

func (g *grpcClient) expand(t *testing.T, r *relationtuple.SubjectSet, depth int) *expand.Tree {
	out := g.c.ExecNoErr(t, "expand", r.Relation, r.Namespace, r.Object, "--"+cliexpand.FlagMaxDepth, fmt.Sprintf("%d", depth), "--"+cmdx.FlagFormat, string(cmdx.FormatJSON))
	res := expand.Tree{}
	require.NoError(t, json.Unmarshal([]byte(out), &res))
	return &res
}
