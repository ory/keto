// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package e2e

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/ory/keto/internal/schema"
	"github.com/ory/keto/ketoapi"

	"github.com/ory/herodot"
	"github.com/tidwall/gjson"

	"github.com/ory/x/healthx"

	"github.com/ory/keto/internal/x"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ory/keto/internal/check"
	"github.com/ory/keto/internal/expand"
	"github.com/ory/keto/internal/relationtuple"
)

var _ client = &restClient{}

type restClient struct {
	readURL, writeURL, oplSyntaxURL string
}

func (rc *restClient) queryNamespaces(t require.TestingT) (res ketoapi.GetNamespacesResponse) {
	body, code := rc.makeRequest(t, http.MethodGet, "/namespaces", "", rc.readURL)
	assert.Equal(t, http.StatusOK, code, body)
	require.NoError(t, json.Unmarshal([]byte(body), &res))

	return
}

func (rc *restClient) oplCheckSyntax(t require.TestingT, content []byte) []*ketoapi.ParseError {
	body, code := rc.makeRequest(t, http.MethodPost, schema.RouteBase, string(content), rc.oplSyntaxURL)
	assert.Equal(t, http.StatusOK, code, body)
	var response ketoapi.CheckOPLSyntaxResponse
	require.NoError(t, json.Unmarshal([]byte(body), &response))

	return response.Errors
}

func (rc *restClient) makeRequest(t require.TestingT, method, path, body string, baseURL string) (string, int) {
	var b io.Reader
	if body != "" {
		b = bytes.NewBufferString(body)
	}

	// t.Logf("Requesting %s %s%s with body %#v", method, baseURL, path, body)
	req, err := http.NewRequest(method, baseURL+path, b)
	require.NoError(t, err)
	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	respBody, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	require.NoError(t, resp.Body.Close())

	return string(respBody), resp.StatusCode
}

func (rc *restClient) createTuple(t require.TestingT, r *ketoapi.RelationTuple) {
	tEnc, err := json.Marshal(r)
	require.NoError(t, err)

	body, code := rc.makeRequest(t, http.MethodPut, relationtuple.WriteRouteBase, string(tEnc), rc.writeURL)
	assert.Equal(t, http.StatusCreated, code, body)
}

func (rc *restClient) deleteTuple(t require.TestingT, r *ketoapi.RelationTuple) {
	body, code := rc.makeRequest(t, http.MethodDelete, relationtuple.WriteRouteBase+"?"+r.ToURLQuery().Encode(), "", rc.writeURL)
	require.Equal(t, http.StatusNoContent, code, body)
}

func (rc *restClient) deleteAllTuples(t require.TestingT, q *ketoapi.RelationQuery) {
	body, code := rc.makeRequest(t, http.MethodDelete, relationtuple.WriteRouteBase+"?"+q.ToURLQuery().Encode(), "", rc.writeURL)
	require.Equal(t, http.StatusNoContent, code, body)
}

func (rc *restClient) queryTuple(t require.TestingT, q *ketoapi.RelationQuery, opts ...x.PaginationOptionSetter) *ketoapi.GetResponse {
	urlQuery := q.ToURLQuery()

	pagination := x.GetPaginationOptions(opts...)
	if pagination.Size != 0 {
		urlQuery.Set("page_size", strconv.Itoa(pagination.Size))
	}
	if pagination.Token != "" {
		urlQuery.Set("page_token", pagination.Token)
	}

	body, code := rc.makeRequest(t, http.MethodGet, fmt.Sprintf("%s?%s", relationtuple.ReadRouteBase, urlQuery.Encode()), "", rc.readURL)
	require.Equal(t, http.StatusOK, code, body)

	var dec ketoapi.GetResponse
	require.NoError(t, json.Unmarshal([]byte(body), &dec))

	return &dec
}

func (rc *restClient) queryTupleErr(t require.TestingT, expected herodot.DefaultError, q *ketoapi.RelationQuery, opts ...x.PaginationOptionSetter) {
	urlQuery := q.ToURLQuery()

	pagination := x.GetPaginationOptions(opts...)
	if pagination.Size != 0 {
		urlQuery.Set("page_size", strconv.Itoa(pagination.Size))
	}
	if pagination.Token != "" {
		urlQuery.Set("page_token", pagination.Token)
	}

	body, code := rc.makeRequest(t, http.MethodGet, fmt.Sprintf("%s?%s", relationtuple.ReadRouteBase, urlQuery.Encode()), "", rc.readURL)

	assert.Equal(t, expected.CodeField, code)
	assert.Equal(t, int64(expected.StatusCode()), gjson.Get(body, "error.code").Int())
	assert.Equal(t, expected.Status(), gjson.Get(body, "error.status").String())
	assert.Equal(t, expected.Error(), gjson.Get(body, "error.message").String(), body)
}

func (rc *restClient) check(t require.TestingT, r *ketoapi.RelationTuple) bool {
	q := r.ToURLQuery()
	bodyGet, codeGet := rc.makeRequest(t, http.MethodGet, fmt.Sprintf("%s?%s", check.RouteBase, q.Encode()), "", rc.readURL)

	var respGet check.CheckPermissionResult
	require.NoError(t, json.Unmarshal([]byte(bodyGet), &respGet))

	j, err := json.Marshal(r)
	require.NoError(t, err)
	bodyPost, codePost := rc.makeRequest(t, http.MethodPost, check.RouteBase, string(j), rc.readURL)

	var respPost check.CheckPermissionResult
	require.NoError(t, json.Unmarshal([]byte(bodyPost), &respPost))

	if codeGet == http.StatusOK && codePost == http.StatusOK {
		assert.Equal(t, true, respGet.Allowed, "%s", bodyGet)
		assert.Equal(t, true, respPost.Allowed, "%s", bodyPost)
		return true
	}

	assert.Equal(t, http.StatusForbidden, codeGet, bodyGet)
	assert.Equal(t, http.StatusForbidden, codePost, bodyPost)
	assert.Equal(t, false, respGet.Allowed)
	assert.Equal(t, false, respPost.Allowed)
	return false
}

func (rc *restClient) expand(t require.TestingT, r *ketoapi.SubjectSet, depth int) *ketoapi.Tree[*ketoapi.RelationTuple] {
	query := r.ToURLQuery()
	query.Set("max-depth", fmt.Sprintf("%d", depth))

	body, code := rc.makeRequest(t, http.MethodGet, fmt.Sprintf("%s?%s", expand.RouteBase, query.Encode()), "", rc.readURL)
	require.Equal(t, http.StatusOK, code, body)

	tree := &ketoapi.Tree[*ketoapi.RelationTuple]{}
	require.NoError(t, json.Unmarshal([]byte(body), tree))

	return tree
}

func healthReady(t require.TestingT, readURL string) bool {
	req, err := http.NewRequest("GET", readURL+healthx.ReadyCheckPath, nil)
	require.NoError(t, err)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return false
	}
	return resp.StatusCode == http.StatusOK
}

func (rc *restClient) waitUntilLive(t require.TestingT) {
	// wait for /health/ready
	for !healthReady(t, rc.readURL) {
		time.Sleep(10 * time.Millisecond)
	}
}
