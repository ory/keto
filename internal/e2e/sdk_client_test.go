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

var requestTimeout = 5 * time.Second

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
	payload := &models.RelationQuery{
		Namespace: r.Namespace,
		Object:    r.Object,
		Relation:  r.Relation,
	}
	switch s := r.Subject.(type) {
	case *relationtuple.SubjectID:
		payload.SubjectID = s.ID
	case *relationtuple.SubjectSet:
		payload.SubjectSet = &models.SubjectSet{
			Namespace: &s.Namespace,
			Object:    &s.Object,
			Relation:  &s.Relation,
		}
	}

	_, err := c.getWriteClient().Write.CreateRelationTuple(
		write.NewCreateRelationTupleParamsWithTimeout(requestTimeout).
			WithPayload(payload),
	)
	require.NoError(t, err)
}

func (c *sdkClient) deleteTuple(t require.TestingT, r *relationtuple.InternalRelationTuple) {
	params := write.NewDeleteRelationTuplesParamsWithTimeout(requestTimeout).
		WithNamespace(&r.Namespace).
		WithObject(&r.Object).
		WithRelation(&r.Relation)
	switch s := r.Subject.(type) {
	case *relationtuple.SubjectID:
		params = params.WithSubjectID(&s.ID)
	case *relationtuple.SubjectSet:
		params = params.
			WithSubjectSetNamespace(&s.Namespace).
			WithSubjectSetObject(&s.Object).
			WithSubjectSetRelation(&s.Relation)
	}

	_, err := c.getWriteClient().Write.DeleteRelationTuples(params)
	require.NoError(t, err)
}

func (c *sdkClient) deleteAllTuples(t require.TestingT, q *relationtuple.RelationQuery) {
	params := write.NewDeleteRelationTuplesParamsWithTimeout(requestTimeout).
		WithNamespace(&q.Namespace).
		WithObject(&q.Object).
		WithRelation(&q.Relation)

	if s := q.Subject(); s != nil {
		switch s.(type) {
		case *relationtuple.SubjectID:
			params = params.WithSubjectID(q.SubjectID)
		case *relationtuple.SubjectSet:
			params = params.
				WithSubjectSetNamespace(&s.SubjectSet().Namespace).
				WithObject(&s.SubjectSet().Object).
				WithRelation(&s.SubjectSet().Relation)
		}
	}

	_, err := c.getWriteClient().Write.DeleteRelationTuples(params)
	require.NoError(t, err)
}

func compileParams(q *relationtuple.RelationQuery, opts []x.PaginationOptionSetter) *read.GetRelationTuplesParams {
	params := read.NewGetRelationTuplesParams().WithNamespace(&q.Namespace)
	if q.Relation != "" {
		params = params.WithRelation(&q.Relation)
	}
	if q.Object != "" {
		params = params.WithObject(&q.Object)
	}
	if q.SubjectID != nil {
		params = params.WithSubjectID(q.SubjectID)
	}
	if q.SubjectSet != nil {
		params = params.
			WithSubjectSetNamespace(&q.SubjectSet.Namespace).
			WithSubjectSetObject(&q.SubjectSet.Object).
			WithSubjectSetRelation(&q.SubjectSet.Relation)
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
	resp, err := c.getReadClient().Read.GetRelationTuples(compileParams(q, opts).WithTimeout(requestTimeout))
	require.NoError(t, err)

	getResp := &relationtuple.GetResponse{
		RelationTuples: make([]*relationtuple.InternalRelationTuple, len(resp.Payload.RelationTuples)),
		NextPageToken:  resp.Payload.NextPageToken,
	}

	for i, rt := range resp.Payload.RelationTuples {
		getResp.RelationTuples[i] = &relationtuple.InternalRelationTuple{
			Namespace: *rt.Namespace,
			Object:    *rt.Object,
			Relation:  *rt.Relation,
		}
		if rt.SubjectSet != nil {
			getResp.RelationTuples[i].Subject = &relationtuple.SubjectSet{
				Namespace: *rt.SubjectSet.Namespace,
				Object:    *rt.SubjectSet.Object,
				Relation:  *rt.SubjectSet.Relation,
			}
		} else {
			getResp.RelationTuples[i].Subject = &relationtuple.SubjectID{ID: rt.SubjectID}
		}
	}

	return getResp
}

func (c *sdkClient) queryTupleErr(t require.TestingT, expected herodot.DefaultError, q *relationtuple.RelationQuery, opts ...x.PaginationOptionSetter) {
	_, err := c.getReadClient().Read.GetRelationTuples(compileParams(q, opts).WithTimeout(requestTimeout))

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
	params := read.NewGetCheckParamsWithTimeout(requestTimeout).
		WithNamespace(&r.Namespace).
		WithObject(&r.Object).
		WithRelation(&r.Relation)
	switch s := r.Subject.(type) {
	case *relationtuple.SubjectID:
		params = params.WithSubjectID(&s.ID)
	case *relationtuple.SubjectSet:
		params = params.
			WithSubjectSetNamespace(&s.Namespace).
			WithSubjectSetObject(&s.Object).
			WithSubjectSetRelation(&s.Relation)
	}
	resp, err := c.getReadClient().Read.GetCheck(params)
	require.NoError(t, err)
	return *resp.Payload.Allowed
}

func buildTree(t require.TestingT, mt *models.ExpandTree) *expand.Tree {
	et := &expand.Tree{
		Type: expand.NodeType(*mt.Type),
	}
	if mt.SubjectSet != nil {
		et.Subject = &relationtuple.SubjectSet{
			Namespace: *mt.SubjectSet.Namespace,
			Object:    *mt.SubjectSet.Object,
			Relation:  *mt.SubjectSet.Relation,
		}
	} else {
		et.Subject = &relationtuple.SubjectID{ID: mt.SubjectID}
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
		read.NewGetExpandParamsWithTimeout(requestTimeout).
			WithNamespace(r.Namespace).
			WithObject(r.Object).
			WithRelation(r.Relation).
			WithMaxDepth(pointerx.Int64(int64(depth))),
	)
	require.NoError(t, err)
	return buildTree(t, resp.Payload)
}

func (c *sdkClient) waitUntilLive(t require.TestingT) {
	resp, err := c.getReadClient().Health.IsInstanceAlive(health.NewIsInstanceAliveParams().WithTimeout(requestTimeout))
	for err != nil {
		resp, err = c.getReadClient().Health.IsInstanceAlive(health.NewIsInstanceAliveParams().WithTimeout(requestTimeout))
	}
	require.Equal(t, "ok", resp.Payload.Status)

	resp, err = c.getWriteClient().Health.IsInstanceAlive(health.NewIsInstanceAliveParams().WithTimeout(requestTimeout))
	for err != nil {
		resp, err = c.getWriteClient().Health.IsInstanceAlive(health.NewIsInstanceAliveParams().WithTimeout(requestTimeout))
	}
	require.Equal(t, "ok", resp.Payload.Status)
}
