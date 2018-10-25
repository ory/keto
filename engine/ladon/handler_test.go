package ladon

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ory/keto/sdk/go/keto/swagger"
	"github.com/ory/keto/x"

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

	cl := swagger.NewEnginesApiWithBasePath(ts.URL)

	for _, f := range []string{"regex", "exact"} {
		t.Run(fmt.Sprintf("flavor=%s", f), func(t *testing.T) {
			t.Run(fmt.Sprint("action=create"), func(t *testing.T) {
				for _, p := range policies[f] {
					t.Run(fmt.Sprintf("policy=%s", p.ID), func(t *testing.T) {
						res, err := cl.UpsertOryAccessControlPolicy(f, toSwaggerPolicy(p))
						x.CheckResponseTest(t, err, http.StatusOK, res)
					})
				}
				for _, r := range roles[f] {
					t.Run(fmt.Sprintf("role=%s", r.ID), func(t *testing.T) {
						_, res, err := cl.UpsertOryAccessControlPolicyRole(f, "", toSwaggerRole(r))
						x.CheckResponseTest(t, err, http.StatusOK, res)
					})
				}
			})

			t.Run("action=authorize", func(t *testing.T) {
				for k, c := range requests[f] {
					t.Run(fmt.Sprintf("case=%d", k), func(t *testing.T) {
						d, res, err := cl.DoOryAccessControlPoliciesAllow(f, c.req)
						x.CheckResponseTest(t, err, http.StatusOK, res)
						assert.Equal(t, c.allowed, d.Allowed)
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

func toSwaggerRole(r Role) swagger.OryAccessControlPolicyRole {
	return swagger.OryAccessControlPolicyRole{
		Members: r.Members,
		Id:      r.ID,
	}
}

func fromSwaggerRole(r swagger.OryAccessControlPolicyRole) Role {
	return Role{
		Members: r.Members,
		ID:      r.Id,
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

			resp, err = c.UpsertOryAccessControlPolicy(f, toSwaggerPolicy(p))
			x.CheckResponseTest(t, err, http.StatusOK, resp)

			o, resp, err := c.GetOryAccessControlPolicy(f, p.ID)
			x.CheckResponseTest(t, err, http.StatusOK, resp)
			assert.Equal(t, p, fromSwaggerPolicy(*o))

			os, resp, err := c.ListOryAccessControlPolicies(f, 100, 0)
			x.CheckResponseTest(t, err, http.StatusOK, resp)

			var ps Policies
			for _, v := range os {
				ps = append(ps, fromSwaggerPolicy(v))
			}

			assert.Equal(t, ps, policies[f][:l+1])
		}

		for _, p := range policies[f] {
			resp, err := c.DeleteOryAccessControlPolicy(f, p.ID)
			x.CheckResponseTest(t, err, http.StatusNoContent, resp)

			_, resp, err = c.GetOryAccessControlPolicy(f, p.ID)
			x.CheckResponseTest(t, err, http.StatusNotFound, resp)
		}
	}
}

func TestRoleCRUD(t *testing.T) {
	ts := crudts()
	defer ts.Close()

	c := swagger.NewEnginesApiWithBasePath(ts.URL)
	for _, f := range []string{"exact", "regex"} {
		for l, r := range roles[f] {
			_, resp, err := c.GetOryAccessControlPolicyRole(f, r.ID)
			x.CheckResponseTest(t, err, http.StatusNotFound, resp)

			o, resp, err := c.UpsertOryAccessControlPolicyRole(f, r.ID, toSwaggerRole(r))
			x.CheckResponseTest(t, err, http.StatusOK, resp)
			require.EqualValues(t, r, fromSwaggerRole(*o))

			o, resp, err = c.GetOryAccessControlPolicyRole(f, r.ID)
			x.CheckResponseTest(t, err, http.StatusOK, resp)
			require.EqualValues(t, r, fromSwaggerRole(*o))

			os, resp, err := c.ListOryAccessControlPolicyRoles(f, 100, 0)
			x.CheckResponseTest(t, err, http.StatusOK, resp)

			var ps Roles
			for _, v := range os {
				ps = append(ps, fromSwaggerRole(v))
			}

			assert.Equal(t, ps, roles[f][:l+1])
		}

		for _, r := range roles[f] {
			resp, err := c.DeleteOryAccessControlPolicyRole(f, r.ID)
			x.CheckResponseTest(t, err, http.StatusNoContent, resp)

			_, resp, err = c.GetOryAccessControlPolicyRole(f, r.ID)
			x.CheckResponseTest(t, err, http.StatusNotFound, resp)
		}
	}
}
