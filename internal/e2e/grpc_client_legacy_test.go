// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package e2e

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	grpcHealthV1 "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/status"

	rts "github.com/ory/keto/gen/go/ory/keto/relation_tuples/v1alpha2"

	"github.com/ory/herodot"

	"github.com/ory/keto/ketoapi"
	opllegacy "github.com/ory/keto/proto/ory/keto/opl/v1alpha1"
	rtslegacy "github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2"
)

type grpcClientLegacy struct {
	read, write, oplSyntax *grpc.ClientConn
	ctx                    context.Context
}

func newGrpcClientLegacy(t testing.TB, ctx context.Context, readRemote, writeRemote, oplSyntaxRemote string) *grpcClientLegacy {
	read, err := grpc.NewClient(readRemote, grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err)
	t.Cleanup(func() { require.NoError(t, read.Close()) })

	write, err := grpc.NewClient(writeRemote, grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err)
	t.Cleanup(func() { require.NoError(t, write.Close()) })

	oplSyntax, err := grpc.NewClient(oplSyntaxRemote, grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err)
	t.Cleanup(func() { require.NoError(t, oplSyntax.Close()) })

	ctx, cancel := context.WithCancel(ctx)
	t.Cleanup(cancel)

	return &grpcClientLegacy{
		read:      read,
		write:     write,
		oplSyntax: oplSyntax,
		ctx:       ctx,
	}
}

func (g *grpcClientLegacy) queryNamespaces(t testing.TB) (apiResponse ketoapi.GetNamespacesResponse) {
	client := rtslegacy.NewNamespacesServiceClient(g.read)
	res, err := client.ListNamespaces(g.ctx, &rtslegacy.ListNamespacesRequest{})
	require.NoError(t, err)
	require.NoError(t, convert(res, &apiResponse))

	return
}

var _ transactClient = (*grpcClientLegacy)(nil)

func (g *grpcClientLegacy) createTuple(t testing.TB, r *ketoapi.RelationTuple) {
	g.transactTuples(t, []*ketoapi.RelationTuple{r}, nil)
}

func (*grpcClientLegacy) createQuery(q *ketoapi.RelationQuery) *rtslegacy.RelationQuery {
	query := &rtslegacy.RelationQuery{
		Namespace: q.Namespace,
		Object:    q.Object,
		Relation:  q.Relation,
	}
	if q.SubjectID != nil {
		query.Subject = rtslegacy.NewSubjectID(*q.SubjectID)
	} else if q.SubjectSet != nil {
		query.Subject = rtslegacy.NewSubjectSet(q.SubjectSet.Namespace, q.SubjectSet.Object, q.SubjectSet.Relation)
	}
	return query
}

func (g *grpcClientLegacy) queryTuple(t testing.TB, q *ketoapi.RelationQuery, opts ...paginationOptionSetter) *ketoapi.GetResponse {
	c := rtslegacy.NewReadServiceClient(g.read)
	pagination := getPaginationOptions(opts...)

	resp, err := c.ListRelationTuples(g.ctx, &rtslegacy.ListRelationTuplesRequest{
		RelationQuery: g.createQuery(q),
		PageToken:     pagination.Token,
		PageSize:      requireInt32WithinBounds(t, pagination.Size),
	})
	require.NoError(t, err)

	irs := make([]*ketoapi.RelationTuple, len(resp.RelationTuples))
	for i := range irs {
		r := &ketoapi.RelationTuple{
			Object:    resp.RelationTuples[i].Object,
			Namespace: resp.RelationTuples[i].Namespace,
			Relation:  resp.RelationTuples[i].Relation,
		}
		switch s := resp.RelationTuples[i].GetSubject().GetRef().(type) {
		case nil:
			t.Fatalf("unexpected nil subject in relation tuple: %+v", resp.RelationTuples[i])
		case *rtslegacy.Subject_Set:
			r.SubjectSet = &ketoapi.SubjectSet{
				Namespace: s.Set.Namespace,
				Object:    s.Set.Object,
				Relation:  s.Set.Relation,
			}
		case *rtslegacy.Subject_Id:
			r.SubjectID = new(s.Id)
		}
		irs[i] = r
	}

	return &ketoapi.GetResponse{
		RelationTuples: irs,
		NextPageToken:  resp.NextPageToken,
	}
}

func (g *grpcClientLegacy) queryTupleErr(t testing.TB, expected *herodot.DefaultError, q *ketoapi.RelationQuery, opts ...paginationOptionSetter) {
	c := rtslegacy.NewReadServiceClient(g.read)
	pagination := getPaginationOptions(opts...)

	_, err := c.ListRelationTuples(g.ctx, &rtslegacy.ListRelationTuplesRequest{
		RelationQuery: g.createQuery(q),
		PageToken:     pagination.Token,
		PageSize:      requireInt32WithinBounds(t, pagination.Size),
	})
	require.Error(t, err)
	s, ok := status.FromError(err)
	require.True(t, ok)
	assert.Equal(t, expected.GRPCCodeField, s.Code(), "%+v", err)
}

func (g *grpcClientLegacy) check(t testing.TB, r *ketoapi.RelationTuple) bool {
	c := rtslegacy.NewCheckServiceClient(g.read)

	req := &rtslegacy.CheckRequest{
		Tuple: &rtslegacy.RelationTuple{
			Namespace: r.Namespace,
			Object:    r.Object,
			Relation:  r.Relation,
		},
	}
	if r.SubjectID != nil {
		req.Tuple.Subject = rtslegacy.NewSubjectID(*r.SubjectID)
	} else {
		req.Tuple.Subject = rtslegacy.NewSubjectSet(r.SubjectSet.Namespace, r.SubjectSet.Object, r.SubjectSet.Relation)
	}
	resp, err := c.Check(g.ctx, req)
	require.NoError(t, err)

	return resp.Allowed
}

func (g *grpcClientLegacy) batchCheckErr(t testing.TB, requestTuples []*ketoapi.RelationTuple, expected *herodot.DefaultError) {
	_, err := g.doBatchCheck(t, requestTuples)
	require.Error(t, err)
	s, ok := status.FromError(err)
	require.True(t, ok)
	assert.Equal(t, expected.GRPCCodeField, s.Code(), "%+v", err)
	assert.Contains(t, s.Message(), expected.Reason())
}

func (g *grpcClientLegacy) batchCheck(t testing.TB, requestTuples []*ketoapi.RelationTuple) []checkResponse {
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

func (g *grpcClientLegacy) doBatchCheck(_ testing.TB, requestTuples []*ketoapi.RelationTuple) (*rtslegacy.BatchCheckResponse, error) {
	c := rtslegacy.NewCheckServiceClient(g.read)

	tuples := make([]*rtslegacy.RelationTuple, len(requestTuples))
	for i, tuple := range requestTuples {
		var subject *rtslegacy.Subject
		if tuple.SubjectID != nil {
			subject = rtslegacy.NewSubjectID(*tuple.SubjectID)
		} else {
			subject = rtslegacy.NewSubjectSet(tuple.SubjectSet.Namespace, tuple.SubjectSet.Object, tuple.SubjectSet.Relation)
		}
		tuples[i] = &rtslegacy.RelationTuple{
			Namespace: tuple.Namespace,
			Object:    tuple.Object,
			Relation:  tuple.Relation,
			Subject:   subject,
		}
	}

	req := &rtslegacy.BatchCheckRequest{
		Tuples: tuples,
	}
	return c.BatchCheck(g.ctx, req)
}

func legacyTreeFromProto(pt *rtslegacy.SubjectTree) *ketoapi.Tree[*ketoapi.RelationTuple] {
	t := new(ketoapi.Tree[*ketoapi.RelationTuple])
	t.Type = ketoapi.TreeNodeType("").FromProto(rts.NodeType(pt.NodeType))

	var tuple ketoapi.RelationTuple
	var sub *rtslegacy.Subject
	if pt.Tuple != nil {
		tuple.Namespace = pt.Tuple.Namespace
		tuple.Object = pt.Tuple.Object
		tuple.Relation = pt.Tuple.Relation
		sub = pt.Tuple.Subject
	} else {
		//lint:ignore SA1019 backwards compatibility
		//nolint:staticcheck
		sub = pt.Subject
	}
	// legacy case: fetch from deprecated fields
	//lint:ignore SA1019 backwards compatibility
	//nolint:staticcheck
	switch sub := sub.Ref.(type) {
	case *rtslegacy.Subject_Id:
		tuple.SubjectID = new(sub.Id)
	case *rtslegacy.Subject_Set:
		tuple.SubjectSet = &ketoapi.SubjectSet{
			Namespace: sub.Set.Namespace,
			Object:    sub.Set.Object,
			Relation:  sub.Set.Relation,
		}
	default:
		panic(fmt.Sprintf("unexpected subject type in subject tree: %#v", pt.Subject))
	}
	t.Tuple = &tuple

	t.Children = make([]*ketoapi.Tree[*ketoapi.RelationTuple], len(pt.Children))
	for i := range pt.Children {
		t.Children[i] = legacyTreeFromProto(pt.Children[i])
	}

	return t
}

func (g *grpcClientLegacy) expand(t testing.TB, r *ketoapi.SubjectSet, depth int) *ketoapi.Tree[*ketoapi.RelationTuple] {
	c := rtslegacy.NewExpandServiceClient(g.read)

	resp, err := c.Expand(g.ctx, &rtslegacy.ExpandRequest{
		Subject:  rtslegacy.NewSubjectSet(r.Namespace, r.Object, r.Relation),
		MaxDepth: requireInt32WithinBounds(t, depth),
	})
	require.NoError(t, err)

	return legacyTreeFromProto(resp.Tree)
}

func (g *grpcClientLegacy) waitUntilLive(t testing.TB) {
	require.EventuallyWithT(t, func(t *assert.CollectT) {
		c := grpcHealthV1.NewHealthClient(g.read)

		res, err := c.Check(g.ctx, &grpcHealthV1.HealthCheckRequest{})
		require.NoError(t, err)
		assert.Equal(t, grpcHealthV1.HealthCheckResponse_SERVING, res.Status)
	}, 2*time.Second, 10*time.Millisecond)
}

func (g *grpcClientLegacy) deleteTuple(t testing.TB, r *ketoapi.RelationTuple) {
	g.transactTuples(t, nil, []*ketoapi.RelationTuple{r})
}

func (g *grpcClientLegacy) deleteAllTuples(t testing.TB, q *ketoapi.RelationQuery) {
	c := rtslegacy.NewWriteServiceClient(g.write)
	query := &rtslegacy.RelationQuery{
		Namespace: q.Namespace,
		Object:    q.Object,
		Relation:  q.Relation,
	}
	if q.SubjectID != nil {
		query.Subject = rtslegacy.NewSubjectID(*q.SubjectID)
	}
	if q.SubjectSet != nil {
		query.Subject = rtslegacy.NewSubjectSet(q.SubjectSet.Namespace, q.SubjectSet.Object, q.SubjectSet.Relation)
	}
	_, err := c.DeleteRelationTuples(g.ctx, &rtslegacy.DeleteRelationTuplesRequest{
		RelationQuery: query,
	})
	require.NoError(t, err)
}

func (g *grpcClientLegacy) transactTuples(t testing.TB, ins []*ketoapi.RelationTuple, del []*ketoapi.RelationTuple) {
	c := rtslegacy.NewWriteServiceClient(g.write)

	toLegacyProto := func(rt *ketoapi.RelationTuple) *rtslegacy.RelationTuple {
		return &rtslegacy.RelationTuple{
			Namespace: rt.Namespace,
			Object:    rt.Object,
			Relation:  rt.Relation,
			Subject: func() *rtslegacy.Subject {
				if rt.SubjectID != nil {
					return rtslegacy.NewSubjectID(*rt.SubjectID)
				}
				return rtslegacy.NewSubjectSet(rt.SubjectSet.Namespace, rt.SubjectSet.Object, rt.SubjectSet.Relation)
			}(),
		}
	}

	deltas := make([]*rtslegacy.RelationTupleDelta, len(ins)+len(del))
	for i := range ins {
		deltas[i] = &rtslegacy.RelationTupleDelta{
			RelationTuple: toLegacyProto(ins[i]),
			Action:        rtslegacy.RelationTupleDelta_ACTION_INSERT,
		}
	}
	for i := range del {
		deltas[len(ins)+i] = &rtslegacy.RelationTupleDelta{
			RelationTuple: toLegacyProto(del[i]),
			Action:        rtslegacy.RelationTupleDelta_ACTION_DELETE,
		}
	}

	_, err := c.TransactRelationTuples(g.ctx, &rtslegacy.TransactRelationTuplesRequest{
		RelationTupleDeltas: deltas,
	})

	require.NoError(t, err)
}

func (g *grpcClientLegacy) oplCheckSyntax(t testing.TB, content []byte) (parseErrors []*ketoapi.ParseError) {
	c := opllegacy.NewSyntaxServiceClient(g.oplSyntax)

	res, err := c.Check(g.ctx, &opllegacy.CheckRequest{Content: content})
	require.NoError(t, err)

	raw, err := json.Marshal(res.ParseErrors)
	require.NoError(t, err)
	err = json.Unmarshal(raw, &parseErrors)
	require.NoError(t, err)

	return
}
