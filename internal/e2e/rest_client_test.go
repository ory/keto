package e2e

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/ory/keto/internal/x"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ory/keto/internal/check"
	"github.com/ory/keto/internal/expand"
	"github.com/ory/keto/internal/relationtuple"
)

var _ client = &restClient{}

type restClient struct {
	readURL, writeURL string
}

func (rc *restClient) makeRequest(t require.TestingT, method, path, body string, write bool) (string, int) {
	var b io.Reader
	if body != "" {
		b = bytes.NewBufferString(body)
	}

	baseURL := rc.readURL
	if write {
		baseURL = rc.writeURL
	}

	// t.Logf("Requesting %s %s%s with body %#v", method, baseURL, path, body)
	req, err := http.NewRequest(method, baseURL+path, b)
	require.NoError(t, err)
	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	respBody, err := ioutil.ReadAll(resp.Body)
	require.NoError(t, err)
	require.NoError(t, resp.Body.Close())

	return string(respBody), resp.StatusCode
}

func (rc *restClient) createTuple(t require.TestingT, r *relationtuple.InternalRelationTuple) {
	tEnc, err := json.Marshal(r)
	require.NoError(t, err)

	body, code := rc.makeRequest(t, http.MethodPut, relationtuple.RouteBase, string(tEnc), true)
	assert.Equal(t, http.StatusCreated, code, body)
}

func (rc *restClient) queryTuple(t require.TestingT, q *relationtuple.RelationQuery, opts ...x.PaginationOptionSetter) *relationtuple.GetResponse {
	urlQuery := q.ToURLQuery()

	pagination := x.GetPaginationOptions(opts...)
	if pagination.Size != 0 {
		urlQuery.Set("page_size", strconv.Itoa(pagination.Size))
	}
	if pagination.Token != "" {
		urlQuery.Set("page_token", pagination.Token)
	}

	body, code := rc.makeRequest(t, http.MethodGet, fmt.Sprintf("%s?%s", relationtuple.RouteBase, urlQuery.Encode()), "", false)
	require.Equal(t, http.StatusOK, code, body)

	var dec relationtuple.GetResponse
	require.NoError(t, json.Unmarshal([]byte(body), &dec))

	return &dec
}

func (rc *restClient) check(t require.TestingT, r *relationtuple.InternalRelationTuple) bool {
	body, code := rc.makeRequest(t, http.MethodGet, fmt.Sprintf("%s?%s", check.RouteBase, r.ToURLQuery().Encode()), "", false)

	if code == http.StatusOK {
		assert.Equal(t, `"allowed"`, body) // JSON string, therefore quoted
		return true
	}

	assert.Equal(t, http.StatusForbidden, code)
	assert.Equal(t, `"rejected"`, body) // JSON string, therefore quoted
	return false
}

func (rc *restClient) expand(t require.TestingT, r *relationtuple.SubjectSet, depth int) *expand.Tree {
	query := r.ToURLQuery()
	query.Set("depth", fmt.Sprintf("%d", depth))

	body, code := rc.makeRequest(t, http.MethodGet, fmt.Sprintf("%s?%s", expand.RouteBase, query.Encode()), "", false)
	require.Equal(t, http.StatusOK, code, body)

	tree := &expand.Tree{}
	require.NoError(t, json.Unmarshal([]byte(body), tree))

	return tree
}
