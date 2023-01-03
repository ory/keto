// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package e2e

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/ory/herodot"
	"github.com/ory/x/pointerx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	httpclient "github.com/ory/keto/internal/httpclient"
	"github.com/ory/keto/internal/x"
	"github.com/ory/keto/ketoapi"
)

type (
	sdkClient struct {
		rc, wc, sc                            *httpclient.APIClient
		readRemote, writeRemote, syntaxRemote string
	}
)

var _ client = (*sdkClient)(nil)

var requestTimeout = 5 * time.Second

func (c *sdkClient) requestCtx() context.Context {
	ctx, _ := context.WithTimeout(context.Background(), requestTimeout)
	return ctx
}

func (c *sdkClient) oplCheckSyntax(t require.TestingT, content []byte) (parseErrors []*ketoapi.ParseError) {
	res, _, err := c.getOPLSyntaxClient().
		RelationshipApi.
		CheckOplSyntax(c.requestCtx()).
		Body(string(content)).
		Execute()
	require.NoError(t, err)

	raw, err := json.Marshal(res.Errors)
	require.NoError(t, err)
	err = json.Unmarshal(raw, &parseErrors)
	require.NoError(t, err)

	return
}

func (c *sdkClient) getReadClient() *httpclient.APIClient {
	if c.rc == nil {
		cfg := httpclient.NewConfiguration()
		cfg.Host = c.readRemote
		cfg.Scheme = "http"
		c.rc = httpclient.NewAPIClient(cfg)
	}
	return c.rc
}

func (c *sdkClient) getWriteClient() *httpclient.APIClient {
	if c.wc == nil {
		cfg := httpclient.NewConfiguration()
		cfg.Host = c.writeRemote
		cfg.Scheme = "http"
		c.wc = httpclient.NewAPIClient(cfg)
	}
	return c.wc
}

func (c *sdkClient) getOPLSyntaxClient() *httpclient.APIClient {
	if c.sc == nil {
		cfg := httpclient.NewConfiguration()
		cfg.Host = c.syntaxRemote
		cfg.Scheme = "http"
		c.sc = httpclient.NewAPIClient(cfg)
	}
	return c.sc
}

func (c *sdkClient) createTuple(t require.TestingT, r *ketoapi.RelationTuple) {
	payload := httpclient.CreateRelationshipBody{
		Namespace: pointerx.Ptr(r.Namespace),
		Object:    pointerx.Ptr(r.Object),
		Relation:  pointerx.Ptr(r.Relation),
		SubjectId: r.SubjectID,
	}
	if r.SubjectID == nil {
		payload.SubjectSet = &httpclient.SubjectSet{
			Namespace: r.SubjectSet.Namespace,
			Object:    r.SubjectSet.Object,
			Relation:  r.SubjectSet.Relation,
		}
	}

	_, _, err := c.getWriteClient().RelationshipApi.
		CreateRelationship(c.requestCtx()).
		CreateRelationshipBody(payload).
		Execute()
	require.NoError(t, err)
}

func withSubject[P interface {
	SubjectId(string) P
	SubjectSetNamespace(string) P
	SubjectSetObject(string) P
	SubjectSetRelation(string) P
}](params P, subID *string, subSet *ketoapi.SubjectSet) P {
	if subID != nil {
		return params.SubjectId(*subID)
	}
	if subSet != nil {
		return params.
			SubjectSetNamespace(subSet.Namespace).
			SubjectSetObject(subSet.Object).
			SubjectSetRelation(subSet.Relation)
	}
	return params
}

func (c *sdkClient) deleteTuple(t require.TestingT, r *ketoapi.RelationTuple) {
	request := c.getWriteClient().RelationshipApi.
		DeleteRelationships(c.requestCtx()).
		Namespace(r.Namespace).
		Object(r.Object).
		Relation(r.Relation)
	request = withSubject(request, r.SubjectID, r.SubjectSet)

	_, err := request.Execute()
	require.NoError(t, err)
}

func (c *sdkClient) deleteAllTuples(t require.TestingT, q *ketoapi.RelationQuery) {
	request := c.getWriteClient().RelationshipApi.DeleteRelationships(c.requestCtx())
	if q.Namespace != nil {
		request = request.Namespace(*q.Namespace)
	}
	if q.Object != nil {
		request = request.Object(*q.Object)
	}
	if q.Relation != nil {
		request = request.Relation(*q.Relation)
	}
	request = withSubject(request, q.SubjectID, q.SubjectSet)
	_, err := request.Execute()
	require.NoError(t, err)
}

func compileParams(req httpclient.RelationshipApiApiGetRelationshipsRequest, q *ketoapi.RelationQuery, opts []x.PaginationOptionSetter) httpclient.RelationshipApiApiGetRelationshipsRequest {
	if q.Namespace != nil {
		req = req.Namespace(*q.Namespace)
	}
	if q.Relation != nil {
		req = req.Relation(*q.Relation)
	}
	if q.Object != nil {
		req = req.Object(*q.Object)
	}
	if q.SubjectID != nil {
		req = req.SubjectId(*q.SubjectID)
	}
	if q.SubjectSet != nil {
		req = req.
			SubjectSetNamespace(q.SubjectSet.Namespace).
			SubjectSetObject(q.SubjectSet.Object).
			SubjectSetRelation(q.SubjectSet.Relation)
	}

	pagination := x.GetPaginationOptions(opts...)
	if pagination.Size != 0 {
		req = req.PageSize(int64(pagination.Size))
	}
	if pagination.Token != "" {
		req = req.PageToken(pagination.Token)
	}

	return req
}

func (c *sdkClient) queryTuple(t require.TestingT, q *ketoapi.RelationQuery, opts ...x.PaginationOptionSetter) *ketoapi.GetResponse {
	request := c.getReadClient().RelationshipApi.GetRelationships(c.requestCtx())
	request = compileParams(request, q, opts)

	resp, _, err := request.Execute()
	require.NoError(t, err)

	getResp := &ketoapi.GetResponse{
		RelationTuples: make([]*ketoapi.RelationTuple, len(resp.RelationTuples)),
		NextPageToken:  resp.GetNextPageToken(),
	}

	for i, rt := range resp.RelationTuples {
		getResp.RelationTuples[i] = &ketoapi.RelationTuple{
			Namespace: rt.Namespace,
			Object:    rt.Object,
			Relation:  rt.Relation,
		}
		if rt.SubjectSet != nil {
			getResp.RelationTuples[i].SubjectSet = &ketoapi.SubjectSet{
				Namespace: rt.SubjectSet.Namespace,
				Object:    rt.SubjectSet.Object,
				Relation:  rt.SubjectSet.Relation,
			}
		} else {
			getResp.RelationTuples[i].SubjectID = rt.SubjectId
		}
	}

	return getResp
}

func (c *sdkClient) queryTupleErr(t require.TestingT, expected herodot.DefaultError, q *ketoapi.RelationQuery, opts ...x.PaginationOptionSetter) {
	request := c.getReadClient().RelationshipApi.GetRelationships(c.requestCtx())
	request = compileParams(request, q, opts)
	_, _, err := request.Execute()

	switch err.(type) {
	case nil:
		require.FailNow(t, "expected error but got nil")
	case *httpclient.GenericOpenAPIError:
		assert.Equal(t, expected.CodeField, http.StatusNotFound)
	default:
		require.FailNow(t, "got unknown error %+v\nexpected %+v", err, expected)
	}
}

func (c *sdkClient) check(t require.TestingT, r *ketoapi.RelationTuple) bool {
	request := c.getReadClient().PermissionApi.CheckPermission(c.requestCtx()).
		Namespace(r.Namespace).
		Object(r.Object).
		Relation(r.Relation)
	request = withSubject(request, r.SubjectID, r.SubjectSet)

	resp, _, err := request.Execute()
	require.NoError(t, err)

	return resp.GetAllowed()
}

func buildTree(t require.TestingT, mt *httpclient.ExpandedPermissionTree) *ketoapi.Tree[*ketoapi.RelationTuple] {
	result := &ketoapi.Tree[*ketoapi.RelationTuple]{
		Type: ketoapi.TreeNodeType(mt.Type),
	}
	if mt.Tuple.SubjectSet != nil {
		result.Tuple = &ketoapi.RelationTuple{
			SubjectSet: &ketoapi.SubjectSet{
				Namespace: mt.Tuple.SubjectSet.Namespace,
				Object:    mt.Tuple.SubjectSet.Object,
				Relation:  mt.Tuple.SubjectSet.Relation,
			},
		}
	} else {
		result.Tuple = &ketoapi.RelationTuple{
			SubjectID: mt.Tuple.SubjectId,
		}
	}

	if result.Type != ketoapi.TreeNodeLeaf && len(mt.Children) != 0 {
		result.Children = make([]*ketoapi.Tree[*ketoapi.RelationTuple], len(mt.Children))
		for i, c := range mt.Children {
			c := c
			result.Children[i] = buildTree(t, &c)
		}
	}
	return result
}

func (c *sdkClient) expand(t require.TestingT, r *ketoapi.SubjectSet, depth int) *ketoapi.Tree[*ketoapi.RelationTuple] {
	request := c.getReadClient().PermissionApi.ExpandPermissions(c.requestCtx()).
		Namespace(r.Namespace).
		Object(r.Object).
		Relation(r.Relation).
		MaxDepth(int64(depth))

	resp, _, err := request.Execute()
	require.NoError(t, err)

	return buildTree(t, resp)
}

func (c *sdkClient) waitUntilLive(t require.TestingT) {
	resp, _, err := c.getReadClient().MetadataApi.IsReady(c.requestCtx()).Execute()
	for err != nil {
		resp, _, err = c.getReadClient().MetadataApi.IsReady(c.requestCtx()).Execute()
	}
	require.Equal(t, "ok", resp.Status)
}

func (c *sdkClient) queryNamespaces(t require.TestingT) (response ketoapi.GetNamespacesResponse) {
	res, _, err := c.getReadClient().RelationshipApi.ListRelationshipNamespaces(c.requestCtx()).Execute()
	require.NoError(t, err)
	require.NoError(t, convert(res, &response))

	return
}
