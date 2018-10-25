package ladon

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ory/keto/sdk/go/keto/swagger"
	"github.com/ory/keto/x"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gobuffalo/packr"
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/urfave/negroni"

	"github.com/ory/herodot"
	"github.com/ory/keto/engine"
	"github.com/ory/keto/storage"
)

func base(ts *httptest.Server, f, path string) string {
	return ts.URL + fmt.Sprintf(strings.Replace(BasePath, ":flavor", "%s", 1), f) + path
}

func TestAllowed(t *testing.T) {
	box := packr.NewBox("./rego")
	compiler, err := engine.NewCompiler(box, logrus.New())
	require.NoError(t, err)

	s := storage.NewMemoryManager()
	sh := storage.NewHandler(s, herodot.NewJSONWriter(nil))
	e := engine.NewEngine(compiler, herodot.NewJSONWriter(nil))
	le := NewEngine(s, sh, e, herodot.NewJSONWriter(nil))

	n := negroni.Classic()
	r := httprouter.New()
	le.Register(r)
	n.UseHandler(r)

	ts := httptest.NewServer(n)
	defer ts.Close()

	for _, f := range []string{"regex", "exact"} {
		t.Run(fmt.Sprintf("flavor=%s", f), func(t *testing.T) {
			t.Run(fmt.Sprint("action=create"), func(t *testing.T) {
				for _, p := range policies[f] {
					t.Run(fmt.Sprintf("policy=%s", p.ID), func(t *testing.T) {
						var b bytes.Buffer
						require.NoError(t, json.NewEncoder(&b).Encode(&p))
						req, err := http.NewRequest("PUT", base(ts, f, "/policies"), &b)
						require.NoError(t, err)
						res, err := ts.Client().Do(req)
						require.NoError(t, err)
						assert.EqualValues(t, http.StatusOK, res.StatusCode)
						res.Body.Close()
					})
				}
				for _, r := range roles[f] {
					t.Run(fmt.Sprintf("role=%s", r.ID), func(t *testing.T) {
						var b bytes.Buffer
						require.NoError(t, json.NewEncoder(&b).Encode(&r))
						req, err := http.NewRequest("PUT", base(ts, f, "/roles"), &b)
						require.NoError(t, err)
						res, err := ts.Client().Do(req)
						require.NoError(t, err)
						assert.EqualValues(t, http.StatusOK, res.StatusCode)
						res.Body.Close()
					})
				}
			})

			t.Run("action=authorize", func(t *testing.T) {
				for k, c := range requests[f] {
					t.Run(fmt.Sprintf("case=%d", k), func(t *testing.T) {
						var b bytes.Buffer
						require.NoError(t, json.NewEncoder(&b).Encode(&c.req))
						res, err := ts.Client().Post(base(ts, f, "/allowed"), "application/json", &b)
						require.NoError(t, err)
						defer res.Body.Close()

						assert.EqualValues(t, http.StatusOK, res.StatusCode)
						body, err := ioutil.ReadAll(res.Body)
						require.NoError(t, err)

						var r engine.AuthorizationResult
						require.NoError(t, json.Unmarshal(body, &r))
						assert.Equal(t, c.allowed, r.Allowed, "%s", body)
					})
				}
			})
		})
	}
}

func TestValidatePolicy(t *testing.T) {
	_, err := validatePolicy(Policy{})
	require.Error(t, err)

	_, err = validatePolicy(Policy{Effect: "bar"})
	require.Error(t, err)

	p, err := validatePolicy(Policy{Effect: "allow"})
	require.NoError(t, err)
	assert.NotEmpty(t, p.ID)

	p, err = validatePolicy(Policy{Effect: "deny", ID: "foo"})
	require.NoError(t, err)
	assert.Equal(t, "foo", p.ID)
}

func crudts() *httptest.Server {
	s := storage.NewMemoryManager()
	sh := storage.NewHandler(s, herodot.NewJSONWriter(nil))
	e := NewEngine(s, sh, nil, herodot.NewJSONWriter(nil))
	r := httprouter.New()
	e.Register(r)
	return httptest.NewServer(r)
}

func toSwaggerPolicy(p Policy) swagger.OryAccessControlPolicy {
	return swagger.OryAccessControlPolicy{
		Actions:     p.Actions,
		Id:          p.ID,
		Resources:   p.Resources,
		Subjects:    p.Subjects,
		Effect:      p.Effect,
		Conditions:  p.Conditions,
		Description: p.Description,
	}
}

func fromSwaggerPolicy(p swagger.OryAccessControlPolicy) Policy {
	return Policy{
		Actions:     p.Actions,
		ID:          p.Id,
		Resources:   p.Resources,
		Subjects:    p.Subjects,
		Effect:      p.Effect,
		Conditions:  p.Conditions,
		Description: p.Description,
	}
}

func TestPolicyCRUD(t *testing.T) {
	ts := crudts()
	defer ts.Close()

	c := swagger.NewEnginesApiWithBasePath(ts.URL)
	for _, f := range []string{"exact", "regex"} {
		for l, p := range policies[f] {
			_, resp, err := c.GetOryAccessControlPolicy(f, p.ID)
			x.CheckResponseTest(t, err, http.StatusNotFound, resp)

			resp, err = c.UpsertOryAccessControlPolicy(f, p.ID, toSwaggerPolicy(p))
			x.CheckResponseTest(t, err, http.StatusOK, resp)

			o, resp, err := c.GetOryAccessControlPolicy(f, p.ID)
			x.CheckResponseTest(t, err, http.StatusOK, resp)
			assert.Equal(t, p, fromSwaggerPolicy(*o))

			os, resp, err := c.ListOryAccessControlPolicies(f, 100, 0)
			x.CheckResponseTest(t, err, http.StatusOK, resp)
			assert.Equal(t, p, os, policies[f][:l+1])
		}

		for _, p := range policies[f] {
			resp, err := c.DeleteOryAccessControlPolicy(f, p.ID)
			x.CheckResponseTest(t, err, http.StatusNoContent, resp)

			_, resp, err = c.GetOryAccessControlPolicy(f, p.ID)
			x.CheckResponseTest(t, err, http.StatusNotFound, resp)
		}
	}
}

func test404(t *testing.T, path string) {
	t.Run(fmt.Sprintf("action=404/path=%s", path), func(t *testing.T) {
		res, err := http.DefaultClient.Get(path)
		require.NoError(t, err)
		assert.Equal(t, http.StatusNotFound, res.StatusCode)
		require.NoError(t, res.Body.Close())
	})
}

func testDelete(t *testing.T, path string) {
	t.Run(fmt.Sprintf("action=delete/path=%s", path), func(t *testing.T) {
		req, err := http.NewRequest("DELETE", path, nil)
		require.NoError(t, err)
		res, err := http.DefaultClient.Do(req)
		require.NoError(t, err)
		assert.Equal(t, http.StatusNoContent, res.StatusCode)
		require.NoError(t, res.Body.Close())
	})
}

func testCreate(t *testing.T, path string, in, expect interface{}) {
	t.Run(fmt.Sprintf("action=create/path=%s", path), func(t *testing.T) {
		var b bytes.Buffer
		require.NoError(t, json.NewEncoder(&b).Encode(in))
		req, err := http.NewRequest("PUT", path, &b)
		require.NoError(t, err)
		res, err := http.DefaultClient.Do(req)
		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, res.StatusCode)

		body, err := ioutil.ReadAll(res.Body)
		require.NoError(t, err)

		var bb bytes.Buffer
		require.NoError(t, json.NewEncoder(&bb).Encode(expect))
		assert.Equal(t,
			strings.Replace(bb.String(), "\n", "", 1),
			string(body),
		)

		require.NoError(t, res.Body.Close())
	})
}

func testGet(t *testing.T, ty string, path string, expect interface{}) {
	t.Run(fmt.Sprintf("action=%s/path=%s", ty, path), func(t *testing.T) {
		res, err := http.DefaultClient.Get(path)
		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, res.StatusCode)

		body, err := ioutil.ReadAll(res.Body)
		require.NoError(t, err)

		var bb bytes.Buffer
		require.NoError(t, json.NewEncoder(&bb).Encode(expect))
		assert.Equal(t,
			strings.Replace(bb.String(), "\n", "", 1),
			string(body),
		)

		require.NoError(t, res.Body.Close())
	})
}

func TestRoleCRUD(t *testing.T) {
	ts := crudts()
	defer ts.Close()

	for _, f := range []string{"exact", "regex"} {
		for l, r := range roles[f] {
			test404(t, base(ts, f, "/roles/"+r.ID))
			testCreate(t, base(ts, f, "/roles"), r, r)
			testGet(t, "get", base(ts, f, "/roles/"+r.ID), r)
			testGet(t, "list", base(ts, f, "/roles"), roles[f][:l+1])
		}

		for _, r := range roles[f] {
			testDelete(t, base(ts, f, "/roles/"+r.ID))
			test404(t, base(ts, f, "/roles/"+r.ID))
		}
	}
}
