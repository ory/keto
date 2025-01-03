// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package e2e

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	grpcHealthV1 "google.golang.org/grpc/health/grpc_health_v1"

	"github.com/ory/herodot"
	"github.com/ory/x/cmdx"

	rts "github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2"

	gprclient "github.com/ory/keto/cmd/client"
	cliexpand "github.com/ory/keto/cmd/expand"
	clirelationtuple "github.com/ory/keto/cmd/relationtuple"
	"github.com/ory/keto/internal/x"
	"github.com/ory/keto/ketoapi"
)

type cliClient struct {
	c *cmdx.CommandExecuter
}

func (g *cliClient) queryNamespaces(t *testing.T) (res ketoapi.GetNamespacesResponse) {
	t.Skip("not implemented for the CLI")
	return
}

var _ client = (*cliClient)(nil)

func (g *cliClient) oplCheckSyntax(t *testing.T, _ []byte) []*ketoapi.ParseError {
	t.Skip("not implemented as a command yet")
	return []*ketoapi.ParseError{}
}

func (g *cliClient) createTuple(t *testing.T, r *ketoapi.RelationTuple) {
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

func (g *cliClient) queryTuple(t *testing.T, q *ketoapi.RelationQuery, opts ...x.PaginationOptionSetter) *ketoapi.GetResponse {
	out := g.c.ExecNoErr(t, append(g.assembleQueryFlags(q, opts), "relation-tuple", "get")...)

	var resp ketoapi.GetResponse
	require.NoError(t, json.Unmarshal([]byte(out), &resp), "%s", out)

	return &resp
}

func (g *cliClient) queryTupleErr(t *testing.T, expected herodot.DefaultError, q *ketoapi.RelationQuery, opts ...x.PaginationOptionSetter) {
	stdErr := g.c.ExecExpectedErr(t, append(g.assembleQueryFlags(q, opts), "relation-tuple", "get")...)
	assert.Contains(t, stdErr, expected.GRPCCodeField.String())
	assert.Contains(t, stdErr, expected.Error())
}

func (g *cliClient) check(t *testing.T, r *ketoapi.RelationTuple) bool {
	var sub string
	if r.SubjectID != nil {
		sub = *r.SubjectID
	} else {
		sub = r.SubjectSet.String()
	}
	out := g.c.ExecNoErr(t, "check", sub, r.Relation, r.Namespace, r.Object)
	var res rts.CheckResponse
	require.NoError(t, json.Unmarshal([]byte(out), &res))
	return res.Allowed
}

func (g *cliClient) batchCheckErr(t *testing.T, requestTuples []*ketoapi.RelationTuple, expected herodot.DefaultError) {
	t.Skip("not implemented for the CLI")
}

func (g *cliClient) batchCheck(t *testing.T, requestTuples []*ketoapi.RelationTuple) []checkResponse {
	t.Skip("not implemented for the CLI")
	return nil
}

func (g *cliClient) expand(t *testing.T, r *ketoapi.SubjectSet, depth int) *ketoapi.Tree[*ketoapi.RelationTuple] {
	out := g.c.ExecNoErr(t, "expand", r.Relation, r.Namespace, r.Object, "--"+cliexpand.FlagMaxDepth, fmt.Sprintf("%d", depth), "--"+cmdx.FlagFormat, string(cmdx.FormatJSON))
	res := ketoapi.Tree[*ketoapi.RelationTuple]{}
	require.NoError(t, json.Unmarshal([]byte(out), &res))
	return &res
}

func (g *cliClient) waitUntilLive(t *testing.T) {
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

	out := cmdx.ExecNoErrCtx(ctx, t, g.c.New(), append(flags, "status", "--"+gprclient.FlagBlock)...)
	require.Equal(t, grpcHealthV1.HealthCheckResponse_SERVING.String()+"\n", out)
}

func (g *cliClient) deleteTuple(t *testing.T, r *ketoapi.RelationTuple) {
	tupleEnc, err := json.Marshal(r)
	require.NoError(t, err)

	stdout, stderr, err := g.c.Exec(bytes.NewBuffer(tupleEnc), "relation-tuple", "delete", "-")
	require.NoError(t, err, "stdout: %s\nstderr: %s", stdout, stderr)
	assert.Len(t, stderr, 0, stdout)
}

func (g *cliClient) deleteAllTuples(t *testing.T, q *ketoapi.RelationQuery) {
	_ = g.c.ExecNoErr(t, append(g.assembleQueryFlags(q, nil), "relation-tuple", "delete-all", "--force")...)
}
