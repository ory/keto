package e2e

import (
	"net/http"
	"time"

	"github.com/ory/keto/ketoapi"

	"github.com/ory/herodot"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ory/keto/internal/httpclient/client/health"

	"github.com/ory/keto/internal/httpclient/client/write"

	"github.com/ory/keto/internal/httpclient/models"

	httpclient "github.com/ory/keto/internal/httpclient/client"
	"github.com/ory/keto/internal/httpclient/client/read"

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

func (c *sdkClient) createTuple(t require.TestingT, r *ketoapi.RelationTuple) {
	payload := &models.RelationQuery{
		Namespace: r.Namespace,
		Object:    r.Object,
		Relation:  r.Relation,
	}
	if r.SubjectID != nil {
		payload.SubjectID = *r.SubjectID
	} else {
		payload.SubjectSet = &models.SubjectSet{
			Namespace: &r.SubjectSet.Namespace,
			Object:    &r.SubjectSet.Object,
			Relation:  &r.SubjectSet.Relation,
		}
	}

	_, err := c.getWriteClient().Write.CreateRelationTuple(
		write.NewCreateRelationTupleParamsWithTimeout(requestTimeout).
			WithPayload(payload),
	)
	require.NoError(t, err)
}

func withSubject[P interface {
	WithSubjectID(*string) P
	WithSubjectSetNamespace(*string) P
	WithSubjectSetObject(*string) P
	WithSubjectSetRelation(*string) P
	WithObject(*string) P
	WithRelation(*string) P
}](params P, subID *string, subSet *ketoapi.SubjectSet) P {
	if subID != nil {
		return params.WithSubjectID(subID)
	}
	if subSet != nil {
		return params.
			WithSubjectSetNamespace(&subSet.Namespace).
			WithSubjectSetObject(&subSet.Object).
			WithSubjectSetRelation(&subSet.Relation)
	}
	return params
}

func (c *sdkClient) deleteTuple(t require.TestingT, r *ketoapi.RelationTuple) {
	params := write.NewDeleteRelationTuplesParamsWithTimeout(time.Second).
		WithNamespace(&r.Namespace).
		WithObject(&r.Object).
		WithRelation(&r.Relation)
	params = withSubject(params, r.SubjectID, r.SubjectSet)

	_, err := c.getWriteClient().Write.DeleteRelationTuples(params)
	require.NoError(t, err)
}

func (c *sdkClient) deleteAllTuples(t require.TestingT, q *ketoapi.RelationQuery) {
	params := write.NewDeleteRelationTuplesParamsWithTimeout(time.Second).
		WithNamespace(q.Namespace).
		WithObject(q.Object).
		WithRelation(q.Relation)
	withSubject(params, q.SubjectID, q.SubjectSet)

	_, err := c.getWriteClient().Write.DeleteRelationTuples(params)
	require.NoError(t, err)
}

func compileParams(q *ketoapi.RelationQuery, opts []x.PaginationOptionSetter) *read.GetRelationTuplesParams {
	params := read.NewGetRelationTuplesParams().WithNamespace(q.Namespace)
	if q.Relation != nil {
		params = params.WithRelation(q.Relation)
	}
	if q.Object != nil {
		params = params.WithObject(q.Object)
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
		params = params.WithPageSize(x.Ptr(int64(pagination.Size)))
	}
	if pagination.Token != "" {
		params = params.WithPageToken(&pagination.Token)
	}

	return params
}

func (c *sdkClient) queryTuple(t require.TestingT, q *ketoapi.RelationQuery, opts ...x.PaginationOptionSetter) *ketoapi.GetResponse {
	resp, err := c.getReadClient().Read.GetRelationTuples(compileParams(q, opts).WithTimeout(time.Second))
	require.NoError(t, err)

	getResp := &ketoapi.GetResponse{
		RelationTuples: make([]*ketoapi.RelationTuple, len(resp.Payload.RelationTuples)),
		NextPageToken:  resp.Payload.NextPageToken,
	}

	for i, rt := range resp.Payload.RelationTuples {
		getResp.RelationTuples[i] = &ketoapi.RelationTuple{
			Namespace: *rt.Namespace,
			Object:    *rt.Object,
			Relation:  *rt.Relation,
		}
		if rt.SubjectSet != nil {
			getResp.RelationTuples[i].SubjectSet = &ketoapi.SubjectSet{
				Namespace: *rt.SubjectSet.Namespace,
				Object:    *rt.SubjectSet.Object,
				Relation:  *rt.SubjectSet.Relation,
			}
		} else {
			getResp.RelationTuples[i].SubjectID = &rt.SubjectID
		}
	}

	return getResp
}

func (c *sdkClient) queryTupleErr(t require.TestingT, expected herodot.DefaultError, q *ketoapi.RelationQuery, opts ...x.PaginationOptionSetter) {
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

func (c *sdkClient) check(t require.TestingT, r *ketoapi.RelationTuple) bool {
	params := read.NewGetCheckParamsWithTimeout(time.Second).
		WithNamespace(&r.Namespace).
		WithObject(&r.Object).
		WithRelation(&r.Relation)
	params = withSubject(params, r.SubjectID, r.SubjectSet)
	resp, err := c.getReadClient().Read.GetCheck(params)
	require.NoError(t, err)
	return *resp.Payload.Allowed
}

func buildTree(t require.TestingT, mt *models.ExpandTree) *ketoapi.ExpandTree {
	et := &ketoapi.ExpandTree{
		Type: ketoapi.ExpandNodeType(*mt.Type),
	}
	if mt.SubjectSet != nil {
		et.SubjectSet = &ketoapi.SubjectSet{
			Namespace: *mt.SubjectSet.Namespace,
			Object:    *mt.SubjectSet.Object,
			Relation:  *mt.SubjectSet.Relation,
		}
	} else {
		et.SubjectID = &mt.SubjectID
	}

	if et.Type != ketoapi.ExpandNodeLeaf && len(mt.Children) != 0 {
		et.Children = make([]*ketoapi.ExpandTree, len(mt.Children))
		for i, c := range mt.Children {
			et.Children[i] = buildTree(t, c)
		}
	}
	return et
}

func (c *sdkClient) expand(t require.TestingT, r *ketoapi.SubjectSet, depth int) *ketoapi.ExpandTree {
	resp, err := c.getReadClient().Read.GetExpand(
		read.NewGetExpandParamsWithTimeout(requestTimeout).
			WithNamespace(r.Namespace).
			WithObject(r.Object).
			WithRelation(r.Relation).
			WithMaxDepth(x.Ptr(int64(depth))),
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
