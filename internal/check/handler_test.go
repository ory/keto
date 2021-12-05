package check_test

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/ory/keto/internal/driver/config"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tidwall/gjson"

	"github.com/ory/keto/internal/check"
	"github.com/ory/keto/internal/driver"
	"github.com/ory/keto/internal/namespace"
	"github.com/ory/keto/internal/relationtuple"
	"github.com/ory/keto/internal/x"
)

func assertAllowed(t *testing.T, resp *http.Response) {
	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, resp.StatusCode, "%s", body)
	assert.True(t, gjson.GetBytes(body, "allowed").Bool())
}

func assertDenied(t *testing.T, resp *http.Response) {
	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	assert.Equal(t, http.StatusForbidden, resp.StatusCode, "%s", body)
	assert.False(t, gjson.GetBytes(body, "allowed").Bool())
}

func TestRESTHandler(t *testing.T) {
	nspaces := []*namespace.Namespace{{
		Name: "check handler",
	}}

	reg := driver.NewSqliteTestRegistry(t, false)
	require.NoError(t, reg.Config().Set(config.KeyNamespaces, nspaces))
	h := check.NewHandler(reg)
	r := httprouter.New()
	h.RegisterReadRoutes(&x.ReadRouter{Router: r})
	ts := httptest.NewServer(r)
	defer ts.Close()

	t.Run("case=returns required query parameter max-depth is missing", func(t *testing.T) {
		resp, err := ts.Client().Get(ts.URL + check.RouteBase)
		require.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
		body, err := io.ReadAll(resp.Body)
		require.NoError(t, err)
		assert.Contains(t, string(body), "required query parameter 'max-depth'")
	})

	t.Run("case=returns bad request on malformed int", func(t *testing.T) {
		resp, err := ts.Client().Get(ts.URL + check.RouteBase + "?max-depth=foo")
		require.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
		body, err := io.ReadAll(resp.Body)
		require.NoError(t, err)
		assert.Contains(t, string(body), "invalid syntax")
	})


	t.Run("case=returns bad request on malformed input", func(t *testing.T) {
		resp, err := ts.Client().Get(ts.URL + check.RouteBase + "?" + url.Values{
			"max-depth": {"10"},
			"subject": {"not#a valid userset rewrite"},
		}.Encode())
		require.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})

	t.Run("case=returns bad request on missing subject", func(t *testing.T) {
		resp, err := ts.Client().Get(ts.URL + check.RouteBase + "?" + url.Values{
			"max-depth": {"10"},
		}.Encode())
		require.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
		body, err := io.ReadAll(resp.Body)
		require.NoError(t, err)
		assert.Contains(t, string(body), "Subject has to be specified")
	})

	t.Run("case=returns denied on unknown namespace", func(t *testing.T) {
		resp, err := ts.Client().Get(ts.URL + check.RouteBase + "?" + url.Values{
			"max-depth": {"10"},
			"namespace":  {"not " + nspaces[0].Name},
			"subject_id": {"foo"},
		}.Encode())
		require.NoError(t, err)

		assertDenied(t, resp)
	})

	t.Run("case=returns allowed", func(t *testing.T) {
		rt := &relationtuple.InternalRelationTuple{
			Namespace: nspaces[0].Name,
			Object:    "o",
			Relation:  "r",
			Subject:   &relationtuple.SubjectID{ID: "s"},
		}
		require.NoError(t, reg.RelationTupleManager().WriteRelationTuples(context.Background(), rt))

		q, err := rt.ToURLQuery()
		require.NoError(t, err)
		q.Add("max-depth", "10")
		resp, err := ts.Client().Get(ts.URL + check.RouteBase + "?" + q.Encode())
		require.NoError(t, err)

		assertAllowed(t, resp)
	})

	t.Run("case=returns denied", func(t *testing.T) {
		resp, err := ts.Client().Get(ts.URL + check.RouteBase + "?" + url.Values{
			"namespace":  {nspaces[0].Name},
			"subject_id": {"foo"},
			"max-depth": {"10"},
		}.Encode())
		require.NoError(t, err)

		assertDenied(t, resp)
	})
}
