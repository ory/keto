package e2e

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tidwall/gjson"

	"github.com/ory/keto/internal/check"
	"github.com/ory/keto/internal/expand"
	"github.com/ory/keto/internal/relationtuple"
)

var _ client = &restClient{}

type restClient struct {
	basicURL, privilegedURL string
}

func (rc *restClient) makeRequest(t *testing.T, method, path, body string, privileged bool) (string, int) {
	var b io.Reader
	if body != "" {
		b = bytes.NewBufferString(body)
	}

	baseURL := rc.basicURL
	if privileged {
		baseURL = rc.privilegedURL
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

func (rc *restClient) createTuple(t *testing.T, r *relationtuple.InternalRelationTuple) {
	tEnc, err := json.Marshal(r)
	require.NoError(t, err)

	body, code := rc.makeRequest(t, http.MethodPut, relationtuple.RouteBase, string(tEnc), true)
	assert.Equal(t, http.StatusCreated, code, body)
}

func (rc *restClient) queryTuple(t *testing.T, q *relationtuple.RelationQuery) []*relationtuple.InternalRelationTuple {
	body, code := rc.makeRequest(t, http.MethodGet, fmt.Sprintf("%s?%s", relationtuple.RouteBase, q.ToURLQuery().Encode()), "", false)
	require.Equal(t, http.StatusOK, code, body)

	tuple := make([]*relationtuple.InternalRelationTuple, 0, gjson.Get(body, "relations.#").Int())
	require.NoError(t, json.Unmarshal([]byte(gjson.Get(body, "relations").Raw), &tuple))

	return tuple
}

func (rc *restClient) check(t *testing.T, r *relationtuple.InternalRelationTuple) bool {
	body, code := rc.makeRequest(t, http.MethodGet, fmt.Sprintf("%s?%s", check.RouteBase, r.ToURLQuery().Encode()), "", false)

	if code == http.StatusOK {
		assert.Equal(t, `"allowed"`, body) // JSON string, therefore quoted
		return true
	}

	assert.Equal(t, http.StatusForbidden, code)
	assert.Equal(t, `"rejected"`, body) // JSON string, therefore quoted
	return false
}

func (rc *restClient) expand(t *testing.T, r *relationtuple.SubjectSet, depth int) *expand.Tree {
	query := r.ToURLQuery()
	query.Set("depth", fmt.Sprintf("%d", depth))

	body, code := rc.makeRequest(t, http.MethodGet, fmt.Sprintf("%s?%s", expand.RouteBase, query.Encode()), "", false)
	require.Equal(t, http.StatusOK, code, body)

	tree := &expand.Tree{}
	require.NoError(t, json.Unmarshal([]byte(body), tree))

	return tree
}
