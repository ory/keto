package e2e

import (
	"net/http"
	"time"

	"github.com/ory/herodot"
	"github.com/ory/x/pointerx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ory/keto/internal/httpclient/client/health"

	"github.com/ory/keto/internal/httpclient/client/write"

	"github.com/ory/keto/internal/httpclient/models"

	httpclient "github.com/ory/keto/internal/httpclient/client"
	"github.com/ory/keto/internal/httpclient/client/read"

	"github.com/ory/keto/internal/expand"
	"github.com/ory/keto/internal/relationtuple"
	"github.com/ory/keto/internal/x"
)

type sdkClient struct {
	rc, wc *httpclient.OryKeto
	readRemote,
	writeRemote string
}

var _ client = (*sdkClient)(nil)

func (c *sdkClient) getReadClient() *httpclient.OryKeto {
	if c.rc == nil {
		c.rc = httpclient.NewHTTPClientWithConfig(nil, &httpclient.TransportConfig{
			Host:    c.readRemote,
			Schemes: []string{"http"},
		})
	}
	return c.rc
}

func (c *sdkClient) getWriteClient() *httpclient.OryKeto {
	if c.wc == nil {
		c.wc = httpclient.NewHTTPClientWithConfig(nil, &httpclient.TransportConfig{
			Host:    c.writeRemote,
			Schemes: []string{"http"},
		})
	}
	return c.wc
}

func (c *sdkClient) createTuple(t require.TestingT, r *relationtuple.InternalRelationTuple) {
	_, err := c.getWriteClient().Write.CreateRelationTuple(
		write.NewCreateRelationTupleParamsWithTimeout(time.Second).
			WithPayload(&models.InternalRelationTuple{
				Namespace: &r.Namespace,
				Object:    &r.Object,
				Relation:  &r.Relation,
				Subject:   (*models.Subject)(pointerx.String(r.Subject.String())),
			}),
	)
	require.NoError(t, err)
}

func (c *sdkClient) deleteTuple(t require.TestingT, r *relationtuple.InternalRelationTuple) {
	_, err := c.getWriteClient().Write.DeleteRelationTuple(
		write.NewDeleteRelationTupleParamsWithTimeout(time.Second).
			WithNamespace(r.Namespace).
			WithObject(r.Object).
			WithRelation(r.Relation).
			WithSubject(pointerx.String(r.Subject.String())),
	)
	require.NoError(t, err)
}

func compileParams(q *relationtuple.RelationQuery, opts []x.PaginationOptionSetter) *read.GetRelationTuplesParams {
	params := read.NewGetRelationTuplesParams().WithNamespace(q.Namespace)
	if q.Relation != "" {
		params = params.WithRelation(&q.Relation)
	}
	if q.Object != "" {
		params = params.WithObject(&q.Object)
	}
	if q.Subject != nil {
		params = params.WithSubject(pointerx.String(q.Subject.String()))
	}

	pagination := x.GetPaginationOptions(opts...)
	if pagination.Size != 0 {
		params = params.WithPageSize(pointerx.Int64(int64(pagination.Size)))
	}
	if pagination.Token != "" {
		params = params.WithPageToken(&pagination.Token)
	}

	return params
}

func (c *sdkClient) queryTuple(t require.TestingT, q *relationtuple.RelationQuery, opts ...x.PaginationOptionSetter) *relationtuple.GetResponse {
	resp, err := c.getReadClient().Read.GetRelationTuples(compileParams(q, opts).WithTimeout(time.Second))
	require.NoError(t, err)

	getResp := &relationtuple.GetResponse{
		RelationTuples: make([]*relationtuple.InternalRelationTuple, len(resp.Payload.RelationTuples)),
		NextPageToken:  resp.Payload.NextPageToken,
		IsLastPage:     resp.Payload.IsLastPage,
	}

	for i, rt := range resp.Payload.RelationTuples {
		sub, err := relationtuple.SubjectFromString(string(*rt.Subject))
		require.NoError(t, err)
		getResp.RelationTuples[i] = &relationtuple.InternalRelationTuple{
			Namespace: *rt.Namespace,
			Object:    *rt.Object,
			Relation:  *rt.Relation,
			Subject:   sub,
		}
	}

	return getResp
}

func (c *sdkClient) queryTupleErr(t require.TestingT, expected herodot.DefaultError, q *relationtuple.RelationQuery, opts ...x.PaginationOptionSetter) {
	_, err := c.getReadClient().Read.GetRelationTuples(compileParams(q, opts).WithTimeout(time.Second))

	switch err.(type) {
	case nil:
		require.FailNow(t, "expected error but got nil")
	case *read.GetRelationTuplesNotFound:
		assert.Equal(t, expected.CodeField, http.StatusNotFound)
	default:
		require.FailNow(t, "got unknown error %+v\nexpected %+v", err, expected)
	}
}

func (c *sdkClient) check(t require.TestingT, r *relationtuple.InternalRelationTuple) bool {
	resp, err := c.getReadClient().Read.GetCheck(
		read.NewGetCheckParamsWithTimeout(time.Second).
			WithNamespace(r.Namespace).
			WithObject(r.Object).
			WithRelation(r.Relation).
			WithSubject(pointerx.String(r.Subject.String())),
	)
	require.NoError(t, err)
	return *resp.Payload.Allowed
}

func buildTree(t require.TestingT, mt *models.ExpandTree) *expand.Tree {
	sub, err := relationtuple.SubjectFromString(string(*mt.Subject))
	require.NoError(t, err)
	et := &expand.Tree{
		Type:    expand.NodeType(*mt.Type),
		Subject: sub,
	}
	if et.Type != expand.Leaf && len(mt.Children) != 0 {
		et.Children = make([]*expand.Tree, len(mt.Children))
		for i, c := range mt.Children {
			et.Children[i] = buildTree(t, c)
		}
	}
	return et
}

func (c *sdkClient) expand(t require.TestingT, r *relationtuple.SubjectSet, depth int) *expand.Tree {
	resp, err := c.getReadClient().Read.GetExpand(
		read.NewGetExpandParamsWithTimeout(time.Second).
			WithNamespace(r.Namespace).
			WithObject(r.Object).
			WithRelation(r.Relation).
			WithDepth(pointerx.Int64(int64(depth))),
	)
	require.NoError(t, err)
	return buildTree(t, resp.Payload)
}

func (c *sdkClient) waitUntilLive(t require.TestingT) {
	resp, err := c.getReadClient().Health.IsInstanceAlive(health.NewIsInstanceAliveParams().WithTimeout(time.Second))
	for err != nil {
		resp, err = c.getReadClient().Health.IsInstanceAlive(health.NewIsInstanceAliveParams().WithTimeout(time.Second))
	}
	require.Equal(t, "ok", resp.Payload.Status)

	resp, err = c.getWriteClient().Health.IsInstanceAlive(health.NewIsInstanceAliveParams().WithTimeout(time.Second))
	for err != nil {
		resp, err = c.getWriteClient().Health.IsInstanceAlive(health.NewIsInstanceAliveParams().WithTimeout(time.Second))
	}
	require.Equal(t, "ok", resp.Payload.Status)
}
