package e2e

import (
	"context"
	"github.com/ory/keto/ketoapi"
	"google.golang.org/grpc/credentials/insecure"
	"time"

	rts "github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2"

	"github.com/ory/herodot"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/status"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	grpcHealthV1 "google.golang.org/grpc/health/grpc_health_v1"

	"github.com/ory/keto/internal/expand"
	"github.com/ory/keto/internal/x"
)

type grpcClient struct {
	readRemote, writeRemote string
	wc, rc                  *grpc.ClientConn
	ctx                     context.Context
}

var _ transactClient = (*grpcClient)(nil)

func (g *grpcClient) conn(t require.TestingT, remote string) *grpc.ClientConn {
	ctx, cancel := context.WithTimeout(g.ctx, 3*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, remote, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock(), grpc.WithDisableHealthCheck())
	require.NoError(t, err)

	return conn
}

func (g *grpcClient) readConn(t require.TestingT) *grpc.ClientConn {
	if g.rc == nil {
		g.rc = g.conn(t, g.readRemote)
	}
	return g.rc
}

func (g *grpcClient) writeConn(t require.TestingT) *grpc.ClientConn {
	if g.wc == nil {
		g.wc = g.conn(t, g.writeRemote)
	}
	return g.wc
}

func (g *grpcClient) createTuple(t require.TestingT, r *ketoapi.RelationTuple) {
	g.transactTuples(t, []*ketoapi.RelationTuple{r}, nil)
}

func (g *grpcClient) queryTuple(t require.TestingT, q *ketoapi.RelationQuery, opts ...x.PaginationOptionSetter) *ketoapi.GetResponse {
	c := rts.NewReadServiceClient(g.readConn(t))

	query := &rts.ListRelationTuplesRequest_Query{
		Namespace: q.Namespace,
		Object:    q.Object,
		Relation:  q.Relation,
	}
	if q.SubjectID != nil {
		query.Subject = rts.NewSubjectID(*q.SubjectID)
	} else {
		query.Subject = rts.NewSubjectSet(q.SubjectSet.Namespace, q.SubjectSet.Object, q.SubjectSet.Relation)
	}

	pagination := x.GetPaginationOptions(opts...)

	resp, err := c.ListRelationTuples(g.ctx, &rts.ListRelationTuplesRequest{
		Query:     query,
		PageToken: pagination.Token,
		PageSize:  int32(pagination.Size),
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

func (g *grpcClient) queryTupleErr(t require.TestingT, expected herodot.DefaultError, q *ketoapi.RelationQuery, opts ...x.PaginationOptionSetter) {
	c := rts.NewReadServiceClient(g.readConn(t))

	query := &rts.ListRelationTuplesRequest_Query{
		Namespace: q.Namespace,
		Object:    q.Object,
		Relation:  q.Relation,
	}
	if q.SubjectID != nil {
		query.Subject = rts.NewSubjectID(*q.SubjectID)
	} else {
		query.Subject = rts.NewSubjectSet(q.SubjectSet.Namespace, q.SubjectSet.Object, q.SubjectSet.Relation)
	}

	pagination := x.GetPaginationOptions(opts...)

	_, err := c.ListRelationTuples(g.ctx, &rts.ListRelationTuplesRequest{
		Query:     query,
		PageToken: pagination.Token,
		PageSize:  int32(pagination.Size),
	})
	require.Error(t, err)
	s, ok := status.FromError(err)
	require.True(t, ok)
	assert.Equal(t, expected.GRPCCodeField, s.Code(), "%+v", err)
}

func (g *grpcClient) check(t require.TestingT, r *ketoapi.RelationTuple) bool {
	c := rts.NewCheckServiceClient(g.readConn(t))

	req := &rts.CheckRequest{
		Namespace: r.Namespace,
		Object:    r.Object,
		Relation:  r.Relation,
	}
	if r.SubjectID != nil {
		req.Subject = rts.NewSubjectID(*r.SubjectID)
	} else {
		req.Subject = rts.NewSubjectSet(r.SubjectSet.Namespace, r.SubjectSet.Object, r.SubjectSet.Relation)
	}
	resp, err := c.Check(g.ctx, req)
	require.NoError(t, err)

	return resp.Allowed
}

func (g *grpcClient) expand(t require.TestingT, r *ketoapi.SubjectSet, depth int) *expand.Tree {
	c := rts.NewExpandServiceClient(g.readConn(t))

	resp, err := c.Expand(g.ctx, &rts.ExpandRequest{
		Subject:  rts.NewSubjectSet(r.Namespace, r.Object, r.Relation),
		MaxDepth: int32(depth),
	})
	require.NoError(t, err)

	tree, err := expand.TreeFromProto(resp.Tree)
	require.NoError(t, err)
	return tree
}

func (g *grpcClient) waitUntilLive(t require.TestingT) {
	c := grpcHealthV1.NewHealthClient(g.readConn(t))

	ctx, cancel := context.WithCancel(g.ctx)
	defer cancel()

	cl, err := c.Watch(ctx, &grpcHealthV1.HealthCheckRequest{})
	require.NoError(t, err)
	require.NoError(t, cl.CloseSend())

	for {
		select {
		case <-g.ctx.Done():
			return
		default:
		}
		resp, err := cl.Recv()
		require.NoError(t, err)

		if resp.Status == grpcHealthV1.HealthCheckResponse_SERVING {
			return
		}
	}
}

func (g *grpcClient) deleteTuple(t require.TestingT, r *ketoapi.RelationTuple) {
	g.transactTuples(t, nil, []*ketoapi.RelationTuple{r})
}

func (g *grpcClient) deleteAllTuples(t require.TestingT, q *ketoapi.RelationQuery) {
	c := rts.NewWriteServiceClient(g.writeConn(t))
	query := &rts.DeleteRelationTuplesRequest_Query{
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
		Query: query,
	})
	require.NoError(t, err)
}

func (g *grpcClient) transactTuples(t require.TestingT, ins []*ketoapi.RelationTuple, del []*ketoapi.RelationTuple) {
	c := rts.NewWriteServiceClient(g.writeConn(t))

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

	_, err := c.TransactRelationTuples(g.ctx, &rts.TransactRelationTuplesRequest{
		RelationTupleDeltas: deltas,
	})

	require.NoError(t, err)
}
