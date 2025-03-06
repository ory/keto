// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package e2e

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ory/herodot"
	"github.com/ory/x/pointerx"

	rts "github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2"

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

func (c *sdkClient) requestCtx(t testing.TB) context.Context {
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	t.Cleanup(cancel)
	return ctx
}

func (c *sdkClient) oplCheckSyntax(t testing.TB, content []byte) (parseErrors []*ketoapi.ParseError) {
	res, _, err := c.getOPLSyntaxClient().
		RelationshipApi.
		CheckOplSyntax(c.requestCtx(t)).
		Body(string(content)).
		Execute()
	require.NoError(t, err)

	//enc, err := json.Marshal(content)
	//require.NoError(t, err)
	//res, _, err := c.getOPLSyntaxClient().
	//	RelationshipApi.
	//	CheckOplSyntax(c.requestCtx()).
	//	Body(string(enc)).
	//	Execute()
	//require.NoError(t, err)

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

func (c *sdkClient) createTuple(t testing.TB, r *ketoapi.RelationTuple) {
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
		CreateRelationship(c.requestCtx(t)).
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

func (c *sdkClient) deleteTuple(t testing.TB, r *ketoapi.RelationTuple) {
	request := c.getWriteClient().RelationshipApi.
		DeleteRelationships(c.requestCtx(t)).
		Namespace(r.Namespace).
		Object(r.Object).
		Relation(r.Relation)
	request = withSubject(request, r.SubjectID, r.SubjectSet)

	_, err := request.Execute()
	require.NoError(t, err)
}

func (c *sdkClient) deleteAllTuples(t testing.TB, q *ketoapi.RelationQuery) {
	request := c.getWriteClient().RelationshipApi.DeleteRelationships(c.requestCtx(t))
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
		req = req.PageSize(int32(pagination.Size))
	}
	if pagination.Token != "" {
		req = req.PageToken(pagination.Token)
	}

	return req
}

func (c *sdkClient) queryTuple(t testing.TB, q *ketoapi.RelationQuery, opts ...x.PaginationOptionSetter) *ketoapi.GetResponse {
	request := c.getReadClient().RelationshipApi.GetRelationships(c.requestCtx(t))
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

func (c *sdkClient) queryTupleErr(t testing.TB, expected herodot.DefaultError, q *ketoapi.RelationQuery, opts ...x.PaginationOptionSetter) {
	request := c.getReadClient().RelationshipApi.GetRelationships(c.requestCtx(t))
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

func (c *sdkClient) check(t testing.TB, r *ketoapi.RelationTuple) bool {
	request := c.getReadClient().PermissionApi.CheckPermission(c.requestCtx(t)).
		Namespace(r.Namespace).
		Object(r.Object).
		Relation(r.Relation)
	request = withSubject(request, r.SubjectID, r.SubjectSet)

	resp, _, err := request.Execute()
	require.NoError(t, err)

	return resp.GetAllowed()
}

func (c *sdkClient) batchCheckErr(t testing.TB, requestTuples []*ketoapi.RelationTuple, expected herodot.DefaultError) {
	request := c.getReadClient().PermissionApi.BatchCheckPermission(c.requestCtx(t)).
		BatchCheckPermissionBody(httpclient.BatchCheckPermissionBody{
			Tuples: tuplesToRelationships(requestTuples),
		})

	_, _, err := request.Execute()
	switch typedErr := err.(type) {
	case nil:
		require.FailNow(t, "expected error but got nil")
	case *httpclient.GenericOpenAPIError:
		assert.Contains(t, typedErr.Error(), expected.Reason())
	default:
		require.FailNow(t, "got unknown error %+v\nexpected %+v", err, expected)
	}
}

func (c *sdkClient) batchCheck(t testing.TB, requestTuples []*ketoapi.RelationTuple) []checkResponse {
	request := c.getReadClient().PermissionApi.BatchCheckPermission(c.requestCtx(t)).
		BatchCheckPermissionBody(httpclient.BatchCheckPermissionBody{
			Tuples: tuplesToRelationships(requestTuples),
		})

	resp, _, err := request.Execute()
	require.NoError(t, err)

	responses := make([]checkResponse, len(resp.Results))
	for i, result := range resp.Results {
		errMsg := ""
		if result.Error != nil {
			errMsg = *result.Error
		}
		responses[i] = checkResponse{
			allowed:      result.Allowed,
			errorMessage: errMsg,
		}
	}
	return responses
}

func tuplesToProto(tuples []*ketoapi.RelationTuple) []*rts.RelationTuple {
	relationships := make([]*rts.RelationTuple, len(tuples))
	for i, requestTuple := range tuples {
		relationships[i] = requestTuple.ToProto()
	}
	return relationships
}

func tuplesToRelationships(tuples []*ketoapi.RelationTuple) []httpclient.Relationship {
	relationships := make([]httpclient.Relationship, len(tuples))
	for i, requestTuple := range tuples {
		relationship := httpclient.Relationship{
			Namespace: requestTuple.Namespace,
			Object:    requestTuple.Object,
			Relation:  requestTuple.Relation,
			SubjectId: requestTuple.SubjectID,
		}
		if requestTuple.SubjectSet != nil {
			relationship.SubjectSet = &httpclient.SubjectSet{
				Namespace: requestTuple.SubjectSet.Namespace,
				Object:    requestTuple.SubjectSet.Object,
				Relation:  requestTuple.SubjectSet.Relation,
			}
		}
		relationships[i] = relationship
	}
	return relationships
}

func buildTree(t testing.TB, mt *httpclient.ExpandedPermissionTree) *ketoapi.Tree[*ketoapi.RelationTuple] {
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

func (c *sdkClient) expand(t testing.TB, r *ketoapi.SubjectSet, depth int) *ketoapi.Tree[*ketoapi.RelationTuple] {
	request := c.getReadClient().PermissionApi.ExpandPermissions(c.requestCtx(t)).
		Namespace(r.Namespace).
		Object(r.Object).
		Relation(r.Relation).
		MaxDepth(int32(depth))

	resp, _, err := request.Execute()
	require.NoError(t, err)

	return buildTree(t, resp)
}

func (c *sdkClient) waitUntilLive(t testing.TB) {
	resp, _, err := c.getReadClient().MetadataApi.IsReady(c.requestCtx(t)).Execute()
	for err != nil {
		resp, _, err = c.getReadClient().MetadataApi.IsReady(c.requestCtx(t)).Execute()
	}
	require.Equal(t, "ok", resp.Status)
}

func (c *sdkClient) queryNamespaces(t testing.TB) (response ketoapi.GetNamespacesResponse) {
	res, _, err := c.getReadClient().RelationshipApi.ListRelationshipNamespaces(c.requestCtx(t)).Execute()
	require.NoError(t, err)
	require.NoError(t, convert(res, &response))

	return
}
