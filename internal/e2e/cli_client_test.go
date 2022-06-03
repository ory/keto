package e2e

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/ory/keto/ketoapi"
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
)

type cliClient struct {
	c *cmdx.CommandExecuter
}

var _ client = (*cliClient)(nil)

func (g *cliClient) createTuple(t require.TestingT, r *ketoapi.RelationTuple) {
	tupleEnc, err := json.Marshal(r)
	require.NoError(t, err)

	stdout, stderr, err := g.c.Exec(bytes.NewBuffer(tupleEnc), "relation-tuple", "create", "-")
	require.NoError(t, err, "stdout: %s\nstderr: %s", stdout, stderr)
	assert.Len(t, stderr, 0, stdout)
}

func (g *cliClient) assembleQueryFlags(q *ketoapi.RelationQuery, opts []x.PaginationOptionSetter) []string {
	var flags []string
	if q.Namespace != nil {
		flags = append(flags, "--"+clirelationtuple.FlagNamespace, *q.Namespace)
	}
	if q.SubjectID != nil {
		flags = append(flags, "--"+clirelationtuple.FlagSubjectID, *q.SubjectID)
	}
	if q.SubjectSet != nil {
		flags = append(flags, "--"+clirelationtuple.FlagSubjectSet, q.SubjectSet.String())
	}
	if q.Relation != nil {
		flags = append(flags, "--"+clirelationtuple.FlagRelation, *q.Relation)
	}
	if q.Object != nil {
		flags = append(flags, "--"+clirelationtuple.FlagObject, *q.Object)
	}
	pagination := x.GetPaginationOptions(opts...)
	if pagination.Token != "" {
		flags = append(flags, "--"+clirelationtuple.FlagPageToken, pagination.Token)
	}
	if pagination.Size != 0 {
		flags = append(flags, "--"+clirelationtuple.FlagPageSize, strconv.Itoa(pagination.Size))
	}
	return flags
}

func (g *cliClient) queryTuple(t require.TestingT, q *ketoapi.RelationQuery, opts ...x.PaginationOptionSetter) *ketoapi.GetResponse {
	out := g.c.ExecNoErr(t, append(g.assembleQueryFlags(q, opts), "relation-tuple", "get")...)

	var resp ketoapi.GetResponse
	require.NoError(t, json.Unmarshal([]byte(out), &resp), "%s", out)

	return &resp
}

func (g *cliClient) queryTupleErr(t require.TestingT, expected herodot.DefaultError, q *ketoapi.RelationQuery, opts ...x.PaginationOptionSetter) {
	stdErr := g.c.ExecExpectedErr(t, append(g.assembleQueryFlags(q, opts), "relation-tuple", "get")...)
	assert.Contains(t, stdErr, expected.GRPCCodeField.String())
	assert.Contains(t, stdErr, expected.Error())
}

func (g *cliClient) check(t require.TestingT, r *ketoapi.RelationTuple) bool {
	var sub string
	if r.SubjectID != nil {
		sub = *r.SubjectID
	} else {
		sub = r.SubjectSet.String()
	}
	out := g.c.ExecNoErr(t, "check", sub, r.Relation, r.Namespace, r.Object)
	var res check.RESTResponse
	require.NoError(t, json.Unmarshal([]byte(out), &res))
	return res.Allowed
}

func (g *cliClient) expand(t require.TestingT, r *ketoapi.SubjectSet, depth int) *expand.Tree {
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

func (g *cliClient) deleteTuple(t require.TestingT, r *ketoapi.RelationTuple) {
	tupleEnc, err := json.Marshal(r)
	require.NoError(t, err)

	stdout, stderr, err := g.c.Exec(bytes.NewBuffer(tupleEnc), "relation-tuple", "delete", "-")
	require.NoError(t, err, "stdout: %s\nstderr: %s", stdout, stderr)
	assert.Len(t, stderr, 0, stdout)
}

func (g *cliClient) deleteAllTuples(t require.TestingT, q *ketoapi.RelationQuery) {
	_ = g.c.ExecNoErr(t, append(g.assembleQueryFlags(q, nil), "relation-tuple", "delete-all", "--force")...)
}
