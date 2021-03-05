package e2e

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/ory/herodot"

	"github.com/ory/keto/internal/check"

	grpcHealthV1 "google.golang.org/grpc/health/grpc_health_v1"

	"github.com/ory/keto/cmd/status"

	"github.com/ory/keto/internal/x"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	cliexpand "github.com/ory/keto/cmd/expand"
	clirelationtuple "github.com/ory/keto/cmd/relationtuple"

	"github.com/ory/x/cmdx"

	"github.com/ory/keto/internal/expand"
	"github.com/ory/keto/internal/relationtuple"
)

type cliClient struct {
	c *cmdx.CommandExecuter
}

var _ client = (*cliClient)(nil)

func (g *cliClient) createTuple(t require.TestingT, r *relationtuple.InternalRelationTuple) {
	tupleEnc, err := json.Marshal(r)
	require.NoError(t, err)

	stdout, stderr, err := g.c.Exec(bytes.NewBuffer(tupleEnc), "relation-tuple", "create", "-")
	require.NoError(t, err, "stdout: %s\nstderr: %s", stdout, stderr)
	assert.Len(t, stderr, 0, stdout)
}

func (g *cliClient) queryTuple(t require.TestingT, q *relationtuple.RelationQuery, opts ...x.PaginationOptionSetter) *relationtuple.GetResponse {
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
	pagination := x.GetPaginationOptions(opts...)
	if pagination.Token != "" {
		flags = append(flags, "--"+clirelationtuple.FlagPageToken, pagination.Token)
	}
	if pagination.Size != 0 {
		flags = append(flags, "--"+clirelationtuple.FlagPageSize, strconv.Itoa(pagination.Size))
	}

	out := g.c.ExecNoErr(t, append(flags, "relation-tuple", "get", q.Namespace)...)

	var resp relationtuple.GetResponse
	require.NoError(t, json.Unmarshal([]byte(out), &resp), "%s", out)

	return &resp
}

func (g *cliClient) queryTupleErr(t require.TestingT, expected herodot.DefaultError, q *relationtuple.RelationQuery, opts ...x.PaginationOptionSetter) {
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
	pagination := x.GetPaginationOptions(opts...)
	if pagination.Token != "" {
		flags = append(flags, "--"+clirelationtuple.FlagPageToken, pagination.Token)
	}
	if pagination.Size != 0 {
		flags = append(flags, "--"+clirelationtuple.FlagPageSize, strconv.Itoa(pagination.Size))
	}

	stdErr := g.c.ExecExpectedErr(t, append(flags, "relation-tuple", "get", q.Namespace)...)
	assert.Contains(t, stdErr, expected.GRPCCodeField.String())
	assert.Contains(t, stdErr, expected.Error())
}

func (g *cliClient) check(t require.TestingT, r *relationtuple.InternalRelationTuple) bool {
	out := g.c.ExecNoErr(t, "check", r.Subject.String(), r.Relation, r.Namespace, r.Object)
	var res check.RESTResponse
	require.NoError(t, json.Unmarshal([]byte(out), &res))
	return res.Allowed
}

func (g *cliClient) expand(t require.TestingT, r *relationtuple.SubjectSet, depth int) *expand.Tree {
	out := g.c.ExecNoErr(t, "expand", r.Relation, r.Namespace, r.Object, "--"+cliexpand.FlagMaxDepth, fmt.Sprintf("%d", depth), "--"+cmdx.FlagFormat, string(cmdx.FormatJSON))
	res := expand.Tree{}
	require.NoError(t, json.Unmarshal([]byte(out), &res))
	return &res
}

func (g *cliClient) waitUntilLive(t require.TestingT) {
	flags := make([]string, len(g.c.PersistentArgs))
	copy(flags, g.c.PersistentArgs)

	for i, f := range flags {
		if f == "--"+cmdx.FlagFormat {
			flags = append(flags[:i], flags[i+2:]...)
			break
		}
	}

	ctx, cancel := context.WithTimeout(g.c.Ctx, time.Minute)
	defer cancel()

	out := cmdx.ExecNoErrCtx(ctx, t, g.c.New(), append(flags, "status", "--"+status.FlagBlock)...)
	require.Equal(t, grpcHealthV1.HealthCheckResponse_SERVING.String()+"\n", out)
}

func (g *cliClient) deleteTuple(t require.TestingT, r *relationtuple.InternalRelationTuple) {
	tupleEnc, err := json.Marshal(r)
	require.NoError(t, err)

	stdout, stderr, err := g.c.Exec(bytes.NewBuffer(tupleEnc), "relation-tuple", "delete", "-")
	require.NoError(t, err, "stdout: %s\nstderr: %s", stdout, stderr)
	assert.Len(t, stderr, 0, stdout)
}

func (g *cliClient) transactTuples(t require.TestingT, ins []*relationtuple.InternalRelationTuple, del []*relationtuple.InternalRelationTuple) {
	for _, i := range ins {
		g.createTuple(t, i)
	}
	for _, d := range del {
		g.deleteTuple(t, d)
	}
}
