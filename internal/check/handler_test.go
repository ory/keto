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

type responseAssertion func(t *testing.T, resp *http.Response)

func baseAssertDenied(t *testing.T, resp *http.Response) {
	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	assert.Equal(t, http.StatusForbidden, resp.StatusCode, "%s", body)
	assert.False(t, gjson.GetBytes(body, "allowed").Bool())
}

// For OpenAPI clients, we want to always regurn a 200 status code even if the
// check returned "denied".
func openAPIAssertDenied(t *testing.T, resp *http.Response) {
	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, resp.StatusCode, "%s", body)
	assert.False(t, gjson.GetBytes(body, "allowed").Bool())
}

func TestRESTHandler(t *testing.T) {
	nspaces := []*namespace.Namespace{{
		Name: "check handler",
	}}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	reg := driver.NewSqliteTestRegistry(t, false)
	require.NoError(t, reg.Config(ctx).Set(config.KeyNamespaces, nspaces))
	h := check.NewHandler(reg)
	r := httprouter.New()
	h.RegisterReadRoutes(&x.ReadRouter{Router: r})
	ts := httptest.NewServer(r)
	defer ts.Close()

	for _, suite := range []struct {
		name         string
		base         string
		assertDenied responseAssertion
	}{
		{name: "base", base: check.RouteBase, assertDenied: baseAssertDenied},
		{name: "openapi", base: check.OpenAPIRouteBase, assertDenied: openAPIAssertDenied},
	} {
		t.Run("suite="+suite.name, func(t *testing.T) {
			assertDenied := suite.assertDenied

			t.Run("case=returns bad request on malformed int", func(t *testing.T) {
				resp, err := ts.Client().Get(ts.URL + suite.base + "?max-depth=foo")
				require.NoError(t, err)

				assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
				body, err := io.ReadAll(resp.Body)
				require.NoError(t, err)
				assert.Contains(t, string(body), "invalid syntax")
			})

			t.Run("case=returns bad request on malformed input", func(t *testing.T) {
				resp, err := ts.Client().Get(ts.URL + suite.base + "?" + url.Values{
					"subject": {"not#a valid userset rewrite"},
				}.Encode())
				require.NoError(t, err)

				assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
			})

			t.Run("case=returns bad request on missing subject", func(t *testing.T) {
				resp, err := ts.Client().Get(ts.URL + suite.base)
				require.NoError(t, err)

				assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
				body, err := io.ReadAll(resp.Body)
				require.NoError(t, err)
				assert.Contains(t, string(body), "Please provide a subject")
			})

			t.Run("case=returns denied on unknown namespace", func(t *testing.T) {
				resp, err := ts.Client().Get(ts.URL + suite.base + "?" + url.Values{
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
				resp, err := ts.Client().Get(ts.URL + suite.base + "?" + q.Encode())
				require.NoError(t, err)

				assertAllowed(t, resp)
			})

			t.Run("case=returns denied", func(t *testing.T) {
				resp, err := ts.Client().Get(ts.URL + suite.base + "?" + url.Values{
					"namespace":  {nspaces[0].Name},
					"subject_id": {"foo"},
				}.Encode())
				require.NoError(t, err)

				assertDenied(t, resp)
			})
		})
	}
}
