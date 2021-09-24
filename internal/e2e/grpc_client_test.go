package e2e

import (
	"context"
	"time"

	"github.com/ory/herodot"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/status"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	grpcHealthV1 "google.golang.org/grpc/health/grpc_health_v1"

	"github.com/ory/keto/internal/expand"
	"github.com/ory/keto/internal/relationtuple"
	"github.com/ory/keto/internal/x"
	acl "github.com/ory/keto/proto/ory/keto/acl/v1alpha1"
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

	conn, err := grpc.DialContext(ctx, remote, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithDisableHealthCheck())
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

func (g *grpcClient) createTuple(t require.TestingT, r *relationtuple.InternalRelationTuple) {
	g.transactTuples(t, []*relationtuple.InternalRelationTuple{r}, nil)
}

func (g *grpcClient) queryTuple(t require.TestingT, q *relationtuple.RelationQuery, opts ...x.PaginationOptionSetter) *relationtuple.GetResponse {
	c := acl.NewReadServiceClient(g.readConn(t))

	query := &acl.ListRelationTuplesRequest_Query{
		Namespace: q.Namespace,
		Object:    q.Object,
		Relation:  q.Relation,
	}
	if s := q.Subject(); s != nil {
		query.Subject = s.ToProto()
	}

	pagination := x.GetPaginationOptions(opts...)

	resp, err := c.ListRelationTuples(g.ctx, &acl.ListRelationTuplesRequest{
		Query:     query,
		PageToken: pagination.Token,
		PageSize:  int32(pagination.Size),
	})
	require.NoError(t, err)

	irs := make([]*relationtuple.InternalRelationTuple, len(resp.RelationTuples))
	for i := range irs {
		irs[i], err = (&relationtuple.InternalRelationTuple{}).FromDataProvider(resp.RelationTuples[i])
		require.NoError(t, err)
	}

	return &relationtuple.GetResponse{
		RelationTuples: irs,
		NextPageToken:  resp.NextPageToken,
	}
}

func (g *grpcClient) queryTupleErr(t require.TestingT, expected herodot.DefaultError, q *relationtuple.RelationQuery, opts ...x.PaginationOptionSetter) {
	c := acl.NewReadServiceClient(g.readConn(t))

	query := &acl.ListRelationTuplesRequest_Query{
		Namespace: q.Namespace,
		Object:    q.Object,
		Relation:  q.Relation,
	}
	if s := q.Subject(); s != nil {
		query.Subject = s.ToProto()
	}

	pagination := x.GetPaginationOptions(opts...)

	_, err := c.ListRelationTuples(g.ctx, &acl.ListRelationTuplesRequest{
		Query:     query,
		PageToken: pagination.Token,
		PageSize:  int32(pagination.Size),
	})
	require.Error(t, err)
	s, ok := status.FromError(err)
	require.True(t, ok)
	assert.Equal(t, expected.GRPCCodeField, s.Code(), "%+v", err)
}

func (g *grpcClient) check(t require.TestingT, r *relationtuple.InternalRelationTuple) bool {
	c := acl.NewCheckServiceClient(g.readConn(t))

	resp, err := c.Check(g.ctx, &acl.CheckRequest{
		Namespace: r.Namespace,
		Object:    r.Object,
		Relation:  r.Relation,
		Subject:   r.Subject.ToProto(),
	})
	require.NoError(t, err)

	return resp.Allowed
}

func (g *grpcClient) expand(t require.TestingT, r *relationtuple.SubjectSet, depth int) *expand.Tree {
	c := acl.NewExpandServiceClient(g.readConn(t))

	resp, err := c.Expand(g.ctx, &acl.ExpandRequest{
		Subject:  r.ToProto(),
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

func (g *grpcClient) deleteTuple(t require.TestingT, r *relationtuple.InternalRelationTuple) {
	g.transactTuples(t, nil, []*relationtuple.InternalRelationTuple{r})
}

func (g *grpcClient) transactTuples(t require.TestingT, ins []*relationtuple.InternalRelationTuple, del []*relationtuple.InternalRelationTuple) {
	c := acl.NewWriteServiceClient(g.writeConn(t))

	deltas := make([]*acl.RelationTupleDelta, len(ins)+len(del))
	for i := range ins {
		deltas[i] = &acl.RelationTupleDelta{
			RelationTuple: ins[i].ToProto(),
			Action:        acl.RelationTupleDelta_INSERT,
		}
	}
	for i := range del {
		deltas[len(ins)+i] = &acl.RelationTupleDelta{
			RelationTuple: del[i].ToProto(),
			Action:        acl.RelationTupleDelta_DELETE,
		}
	}

	_, err := c.TransactRelationTuples(g.ctx, &acl.TransactRelationTuplesRequest{
		RelationTupleDeltas: deltas,
	})

	require.NoError(t, err)
}
