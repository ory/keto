// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package e2e

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/ory/keto/ketoapi"
	opl "github.com/ory/keto/proto/ory/keto/opl/v1alpha1"

	rts "github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	grpcHealthV1 "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/status"

	"github.com/ory/herodot"

	"github.com/ory/keto/internal/x"
)

type grpcClient struct {
	read, write, oplSyntax *grpc.ClientConn
	ctx                    context.Context
}

func newGrpcClient(t testing.TB, ctx context.Context, readRemote, writeRemote, oplSyntaxRemote string) *grpcClient {
	read, err := grpc.NewClient(readRemote, grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err)
	t.Cleanup(func() { read.Close() })

	write, err := grpc.NewClient(writeRemote, grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err)
	t.Cleanup(func() { write.Close() })

	oplSyntax, err := grpc.NewClient(oplSyntaxRemote, grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err)
	t.Cleanup(func() { oplSyntax.Close() })

	ctx, cancel := context.WithCancel(ctx)
	t.Cleanup(cancel)

	return &grpcClient{
		read:      read,
		write:     write,
		oplSyntax: oplSyntax,
		ctx:       ctx,
	}
}

func (g *grpcClient) queryNamespaces(t testing.TB) (apiResponse ketoapi.GetNamespacesResponse) {
	client := rts.NewNamespacesServiceClient(g.read)
	res, err := client.ListNamespaces(g.ctx, &rts.ListNamespacesRequest{})
	require.NoError(t, err)
	require.NoError(t, convert(res, &apiResponse))

	return
}

var _ transactClient = (*grpcClient)(nil)

func (g *grpcClient) createTuple(t testing.TB, r *ketoapi.RelationTuple) {
	g.transactTuples(t, []*ketoapi.RelationTuple{r}, nil)
}

func (*grpcClient) createQuery(q *ketoapi.RelationQuery) *rts.RelationQuery {
	query := &rts.RelationQuery{
		Namespace: q.Namespace,
		Object:    q.Object,
		Relation:  q.Relation,
	}
	if q.SubjectID != nil {
		query.Subject = rts.NewSubjectID(*q.SubjectID)
	} else if q.SubjectSet != nil {
		query.Subject = rts.NewSubjectSet(q.SubjectSet.Namespace, q.SubjectSet.Object, q.SubjectSet.Relation)
	}
	return query
}

func (g *grpcClient) queryTuple(t testing.TB, q *ketoapi.RelationQuery, opts ...x.PaginationOptionSetter) *ketoapi.GetResponse {
	c := rts.NewReadServiceClient(g.read)
	pagination := x.GetPaginationOptions(opts...)

	resp, err := c.ListRelationTuples(g.ctx, &rts.ListRelationTuplesRequest{
		RelationQuery: g.createQuery(q),
		PageToken:     pagination.Token,
		PageSize:      int32(pagination.Size),
	})
	require.NoError(t, err)

	irs := make([]*ketoapi.RelationTuple, len(resp.RelationTuples))
	for i := range irs {
		irs[i], err = (&ketoapi.RelationTuple{}).FromDataProvider(resp.RelationTuples[i])
		require.NoError(t, err)
	}

	return &ketoapi.GetResponse{
		RelationTuples: irs,
		NextPageToken:  resp.NextPageToken,
	}
}

func (g *grpcClient) queryTupleErr(t testing.TB, expected herodot.DefaultError, q *ketoapi.RelationQuery, opts ...x.PaginationOptionSetter) {
	c := rts.NewReadServiceClient(g.read)
	pagination := x.GetPaginationOptions(opts...)

	_, err := c.ListRelationTuples(g.ctx, &rts.ListRelationTuplesRequest{
		RelationQuery: g.createQuery(q),
		PageToken:     pagination.Token,
		PageSize:      int32(pagination.Size),
	})
	require.Error(t, err)
	s, ok := status.FromError(err)
	require.True(t, ok)
	assert.Equal(t, expected.GRPCCodeField, s.Code(), "%+v", err)
}

func (g *grpcClient) check(t testing.TB, r *ketoapi.RelationTuple) bool {
	c := rts.NewCheckServiceClient(g.read)

	req := &rts.CheckRequest{
		Tuple: &rts.RelationTuple{
			Namespace: r.Namespace,
			Object:    r.Object,
			Relation:  r.Relation,
		},
	}
	if r.SubjectID != nil {
		req.Tuple.Subject = rts.NewSubjectID(*r.SubjectID)
	} else {
		req.Tuple.Subject = rts.NewSubjectSet(r.SubjectSet.Namespace, r.SubjectSet.Object, r.SubjectSet.Relation)
	}
	resp, err := c.Check(g.ctx, req)
	require.NoError(t, err)

	return resp.Allowed
}

type checkResponse struct {
	allowed      bool
	errorMessage string
}

func (g *grpcClient) batchCheckErr(t testing.TB, requestTuples []*ketoapi.RelationTuple, expected herodot.DefaultError) {

	_, err := g.doBatchCheck(t, requestTuples)
	require.Error(t, err)
	s, ok := status.FromError(err)
	require.True(t, ok)
	assert.Equal(t, expected.GRPCCodeField, s.Code(), "%+v", err)
	assert.Contains(t, s.Message(), expected.Reason())
}

func (g *grpcClient) batchCheck(t testing.TB, requestTuples []*ketoapi.RelationTuple) []checkResponse {
	resp, err := g.doBatchCheck(t, requestTuples)
	require.NoError(t, err)

	checkResponses := make([]checkResponse, len(resp.Results))
	for i, r := range resp.Results {
		checkResponses[i] = checkResponse{
			allowed:      r.Allowed,
			errorMessage: r.Error,
		}
	}

	return checkResponses
}

func (g *grpcClient) doBatchCheck(_ testing.TB, requestTuples []*ketoapi.RelationTuple) (*rts.BatchCheckResponse, error) {

	c := rts.NewCheckServiceClient(g.read)

	tuples := make([]*rts.RelationTuple, len(requestTuples))
	for i, tuple := range requestTuples {
		var subject *rts.Subject
		if tuple.SubjectID != nil {
			subject = rts.NewSubjectID(*tuple.SubjectID)
		} else {
			subject = rts.NewSubjectSet(tuple.SubjectSet.Namespace, tuple.SubjectSet.Object, tuple.SubjectSet.Relation)
		}
		tuples[i] = &rts.RelationTuple{
			Namespace: tuple.Namespace,
			Object:    tuple.Object,
			Relation:  tuple.Relation,
			Subject:   subject,
		}
	}

	req := &rts.BatchCheckRequest{
		Tuples: tuples,
	}
	return c.BatchCheck(g.ctx, req)
}

func (g *grpcClient) expand(t testing.TB, r *ketoapi.SubjectSet, depth int) *ketoapi.Tree[*ketoapi.RelationTuple] {
	c := rts.NewExpandServiceClient(g.read)

	resp, err := c.Expand(g.ctx, &rts.ExpandRequest{
		Subject:  rts.NewSubjectSet(r.Namespace, r.Object, r.Relation),
		MaxDepth: int32(depth),
	})
	require.NoError(t, err)

	return ketoapi.TreeFromProto[*ketoapi.RelationTuple](resp.Tree)
}

func (g *grpcClient) waitUntilLive(t testing.TB) {
	require.EventuallyWithT(t, func(t *assert.CollectT) {
		c := grpcHealthV1.NewHealthClient(g.read)

		res, err := c.Check(g.ctx, &grpcHealthV1.HealthCheckRequest{})
		require.NoError(t, err)
		assert.Equal(t, grpcHealthV1.HealthCheckResponse_SERVING, res.Status)
	}, 2*time.Second, 10*time.Millisecond)
}

func (g *grpcClient) deleteTuple(t testing.TB, r *ketoapi.RelationTuple) {
	g.transactTuples(t, nil, []*ketoapi.RelationTuple{r})
}

func (g *grpcClient) deleteAllTuples(t testing.TB, q *ketoapi.RelationQuery) {
	c := rts.NewWriteServiceClient(g.write)
	query := &rts.RelationQuery{
		Namespace: q.Namespace,
		Object:    q.Object,
		Relation:  q.Relation,
	}
	if q.SubjectID != nil {
		query.Subject = rts.NewSubjectID(*q.SubjectID)
	}
	if q.SubjectSet != nil {
		query.Subject = rts.NewSubjectSet(q.SubjectSet.Namespace, q.SubjectSet.Object, q.SubjectSet.Relation)
	}
	_, err := c.DeleteRelationTuples(g.ctx, &rts.DeleteRelationTuplesRequest{
		RelationQuery: query,
	})
	require.NoError(t, err)
}

func (g *grpcClient) transactTuples(t testing.TB, ins []*ketoapi.RelationTuple, del []*ketoapi.RelationTuple) {
	c := rts.NewWriteServiceClient(g.write)

	deltas := make([]*rts.RelationTupleDelta, len(ins)+len(del))
	for i := range ins {
		deltas[i] = &rts.RelationTupleDelta{
			RelationTuple: ins[i].ToProto(),
			Action:        rts.RelationTupleDelta_ACTION_INSERT,
		}
	}
	for i := range del {
		deltas[len(ins)+i] = &rts.RelationTupleDelta{
			RelationTuple: del[i].ToProto(),
			Action:        rts.RelationTupleDelta_ACTION_DELETE,
		}
	}

	_, err := c.TransactRelationTuples(g.ctx,
		&rts.TransactRelationTuplesRequest{RelationTupleDeltas: deltas},
	)

	require.NoError(t, err)
}

func (g *grpcClient) oplCheckSyntax(t testing.TB, content []byte) (parseErrors []*ketoapi.ParseError) {
	c := opl.NewSyntaxServiceClient(g.oplSyntax)

	res, err := c.Check(g.ctx, &opl.CheckRequest{Content: content})
	require.NoError(t, err)

	raw, err := json.Marshal(res.ParseErrors)
	require.NoError(t, err)
	err = json.Unmarshal(raw, &parseErrors)
	require.NoError(t, err)

	return
}
