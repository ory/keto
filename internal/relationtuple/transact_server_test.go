package relationtuple_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/ory/keto/internal/driver/config"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/julienschmidt/httprouter"

	"github.com/ory/keto/internal/driver"
	"github.com/ory/keto/internal/namespace"
	"github.com/ory/keto/internal/relationtuple"
	"github.com/ory/keto/internal/x"
)

func TestWriteHandlers(t *testing.T) {
	r := httprouter.New()
	wr := &x.WriteRouter{Router: r}
	rr := &x.ReadRouter{Router: r}
	reg := driver.NewSqliteTestRegistry(t, false)

	var nspaces []*namespace.Namespace
	addNamespace := func(t *testing.T) *namespace.Namespace {
		n := &namespace.Namespace{
			ID:   int32(len(nspaces)),
			Name: t.Name(),
		}
		nspaces = append(nspaces, n)

		require.NoError(t, reg.Config(context.Background()).Set(config.KeyNamespaces, nspaces))

		return n
	}

	h := relationtuple.NewHandler(reg)
	h.RegisterWriteRoutes(wr)
	h.RegisterReadRoutes(rr)
	ts := httptest.NewServer(r)
	defer ts.Close()

	t.Run("method=create", func(t *testing.T) {
		doCreate := func(raw []byte) *http.Response {
			req, err := http.NewRequest(http.MethodPut, ts.URL+relationtuple.RouteBase, bytes.NewBuffer(raw))
			require.NoError(t, err)
			resp, err := ts.Client().Do(req)
			require.NoError(t, err)

			return resp
		}

		t.Run("case=creates tuple", func(t *testing.T) {
			nspace := addNamespace(t)

			rt := &relationtuple.InternalRelationTuple{
				Namespace: nspace.Name,
				Object:    "obj",
				Relation:  "rel",
				Subject:   &relationtuple.SubjectID{ID: "subj"},
			}
			payload, err := json.Marshal(rt)
			require.NoError(t, err)

			resp := doCreate(payload)

			assert.Equal(t, http.StatusCreated, resp.StatusCode)

			body, err := io.ReadAll(resp.Body)
			require.NoError(t, err)

			assert.JSONEq(t, string(payload), string(body))

			t.Run("check=is contained in the manager", func(t *testing.T) {
				// set a size > 1 just to make sure it gets all
				actualRTs, _, err := reg.RelationTupleManager().GetRelationTuples(context.Background(), rt.ToQuery(), x.WithSize(10))
				require.NoError(t, err)
				assert.Equal(t, []*relationtuple.InternalRelationTuple{rt}, actualRTs)
			})

			t.Run("check=is gettable with the returned URL", func(t *testing.T) {
				resp, err := ts.Client().Get(ts.URL + resp.Header.Get("Location"))
				require.NoError(t, err)
				require.Equal(t, http.StatusOK, resp.StatusCode)

				respDec := relationtuple.GetResponse{}
				require.NoError(t, json.NewDecoder(resp.Body).Decode(&respDec))
				assert.Equal(t, []*relationtuple.InternalRelationTuple{rt}, respDec.RelationTuples)
			})
		})

		t.Run("case=returns bad request on JSON parse error", func(t *testing.T) {
			resp := doCreate([]byte("foo"))
			assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
		})

		t.Run("case=special chars", func(t *testing.T) {
			nspace := addNamespace(t)

			rts := []*relationtuple.InternalRelationTuple{
				{
					Namespace: nspace.Name,
					Object:    "group:B",
					Relation:  "member",
					Subject: &relationtuple.SubjectSet{
						Namespace: nspace.Name,
						Object:    "group:A",
						Relation:  "member",
					},
				},
				{
					Namespace: nspace.Name,
					Object:    "@all",
					Relation:  "member",
					Subject:   &relationtuple.SubjectID{ID: "this:could#be interpreted:as a@subject set"},
				},
			}

			for _, rt := range rts {
				payload, err := json.Marshal(rt)
				require.NoError(t, err)

				resp := doCreate(payload)
				assert.Equal(t, http.StatusCreated, resp.StatusCode)
			}

			actual, next, err := reg.RelationTupleManager().GetRelationTuples(context.Background(), &relationtuple.RelationQuery{
				Namespace: nspace.Name,
			})
			require.NoError(t, err)
			assert.Equal(t, "", next)
			assert.Len(t, actual, 2)
			for _, rt := range rts {
				assert.Contains(t, actual, rt)
			}
		})
	})

	t.Run("method=delete", func(t *testing.T) {
		t.Run("case=deletes a tuple", func(t *testing.T) {
			nspace := addNamespace(t)

			rt := &relationtuple.InternalRelationTuple{
				Namespace: nspace.Name,
				Object:    "deleted obj",
				Relation:  "deleted rel",
				Subject:   &relationtuple.SubjectID{ID: "deleted subj"},
			}
			require.NoError(t, reg.RelationTupleManager().WriteRelationTuples(context.Background(), rt))

			q, err := rt.ToURLQuery()
			require.NoError(t, err)
			req, err := http.NewRequest(http.MethodDelete, ts.URL+relationtuple.RouteBase+"?"+q.Encode(), nil)
			require.NoError(t, err)
			resp, err := ts.Client().Do(req)
			require.NoError(t, err)
			assert.Equal(t, http.StatusNoContent, resp.StatusCode)

			// set a size > 1 just to make sure it gets all
			actualRTs, _, err := reg.RelationTupleManager().GetRelationTuples(context.Background(), rt.ToQuery(), x.WithSize(10))
			require.NoError(t, err)
			assert.Equal(t, []*relationtuple.InternalRelationTuple{}, actualRTs)
		})

		t.Run("case=deletes multiple tuples", func(t *testing.T) {
			nspace := addNamespace(t)

			rts := []*relationtuple.InternalRelationTuple{
				{
					Namespace: nspace.Name,
					Object:    "deleted obj",
					Relation:  "deleted rel",
					Subject:   &relationtuple.SubjectID{ID: "deleted subj 1"},
				},
				{
					Namespace: nspace.Name,
					Object:    "deleted obj",
					Relation:  "deleted rel",
					Subject:   &relationtuple.SubjectID{ID: "deleted subj 2"},
				},
			}

			require.NoError(t, reg.RelationTupleManager().WriteRelationTuples(context.Background(), rts...))

			q := url.Values{
				"namespace": {nspace.Name},
				"object":    {"deleted obj"},
				"relation":  {"deleted rel"},
			}
			req, err := http.NewRequest(http.MethodDelete, ts.URL+relationtuple.RouteBase+"?"+q.Encode(), nil)
			require.NoError(t, err)
			resp, err := ts.Client().Do(req)
			require.NoError(t, err)
			assert.Equal(t, http.StatusNoContent, resp.StatusCode)

			query, err := (&relationtuple.RelationQuery{}).FromURLQuery(q)
			require.NoError(t, err)

			actualRTs, _, err := reg.RelationTupleManager().GetRelationTuples(context.Background(), query, x.WithSize(10))
			require.NoError(t, err)
			assert.Equal(t, []*relationtuple.InternalRelationTuple{}, actualRTs)
		})
	})

	t.Run("method=patch", func(t *testing.T) {
		t.Run("case=create and delete", func(t *testing.T) {
			nspace := addNamespace(t)

			deltas := []*relationtuple.PatchDelta{
				{
					Action: relationtuple.ActionInsert,
					RelationTuple: &relationtuple.InternalRelationTuple{
						Namespace: nspace.Name,
						Object:    "create obj",
						Relation:  t.Name(),
						Subject:   &relationtuple.SubjectID{ID: "create sub"},
					},
				},
				{
					Action: relationtuple.ActionDelete,
					RelationTuple: &relationtuple.InternalRelationTuple{
						Namespace: nspace.Name,
						Object:    "delete obj",
						Relation:  t.Name(),
						Subject:   &relationtuple.SubjectID{ID: "delete sub"},
					},
				},
			}
			require.NoError(t, reg.RelationTupleManager().WriteRelationTuples(context.Background(), deltas[1].RelationTuple))

			body, err := json.Marshal(deltas)
			require.NoError(t, err)
			req, err := http.NewRequest(http.MethodPatch, ts.URL+relationtuple.RouteBase, bytes.NewBuffer(body))
			require.NoError(t, err)
			resp, err := ts.Client().Do(req)
			require.NoError(t, err)
			assert.Equal(t, http.StatusNoContent, resp.StatusCode)

			actualRTs, _, err := reg.RelationTupleManager().GetRelationTuples(context.Background(), &relationtuple.RelationQuery{
				Namespace: nspace.Name,
				Relation:  t.Name(),
			})
			require.NoError(t, err)
			assert.Equal(t, []*relationtuple.InternalRelationTuple{deltas[0].RelationTuple}, actualRTs)
		})

		t.Run("case=ignores rest on err", func(t *testing.T) {
			nspace := addNamespace(t)

			deltas := []*relationtuple.PatchDelta{
				{
					Action: relationtuple.ActionInsert,
					RelationTuple: &relationtuple.InternalRelationTuple{
						Namespace: nspace.Name,
						Object:    "create obj",
						Relation:  t.Name(),
						Subject:   &relationtuple.SubjectID{ID: "create sub"},
					},
				},
				{
					Action: relationtuple.ActionDelete,
					RelationTuple: &relationtuple.InternalRelationTuple{
						Namespace: "not " + nspace.Name,
						Object:    "o",
						Relation:  "r",
						Subject:   &relationtuple.SubjectID{ID: "s"},
					},
				},
			}

			body, err := json.Marshal(deltas)
			require.NoError(t, err)
			req, err := http.NewRequest(http.MethodPatch, ts.URL+relationtuple.RouteBase, bytes.NewBuffer(body))
			require.NoError(t, err)
			resp, err := ts.Client().Do(req)
			require.NoError(t, err)
			assert.Equal(t, http.StatusNotFound, resp.StatusCode)

			// set a size > 1 just to make sure it gets all
			actualRTs, _, err := reg.RelationTupleManager().GetRelationTuples(context.Background(), deltas[0].RelationTuple.ToQuery(), x.WithSize(10))
			require.NoError(t, err)
			assert.Len(t, actualRTs, 0)
		})

		t.Run("case=only create", func(t *testing.T) {
			nspace := addNamespace(t)

			deltas := []*relationtuple.PatchDelta{
				{
					Action: relationtuple.ActionInsert,
					RelationTuple: &relationtuple.InternalRelationTuple{
						Namespace: nspace.Name,
						Object:    "create obj",
						Relation:  "rel",
						Subject:   &relationtuple.SubjectID{ID: "create sub"},
					},
				},
			}

			body, err := json.Marshal(deltas)
			require.NoError(t, err)
			req, err := http.NewRequest(http.MethodPatch, ts.URL+relationtuple.RouteBase, bytes.NewBuffer(body))
			require.NoError(t, err)
			resp, err := ts.Client().Do(req)
			require.NoError(t, err)
			assert.Equal(t, http.StatusNoContent, resp.StatusCode)

			actualRTs, _, err := reg.RelationTupleManager().GetRelationTuples(context.Background(), &relationtuple.RelationQuery{
				Namespace: nspace.Name,
			})
			require.NoError(t, err)
			assert.Equal(t, []*relationtuple.InternalRelationTuple{deltas[0].RelationTuple}, actualRTs)
		})

		t.Run("case=only delete", func(t *testing.T) {
			nspace := addNamespace(t)

			deltas := []*relationtuple.PatchDelta{
				{
					Action: relationtuple.ActionDelete,
					RelationTuple: &relationtuple.InternalRelationTuple{
						Namespace: nspace.Name,
						Object:    "delete obj",
						Relation:  "rel",
						Subject:   &relationtuple.SubjectID{ID: "delete sub"},
					},
				},
			}

			body, err := json.Marshal(deltas)
			require.NoError(t, err)
			req, err := http.NewRequest(http.MethodPatch, ts.URL+relationtuple.RouteBase, bytes.NewBuffer(body))
			require.NoError(t, err)
			resp, err := ts.Client().Do(req)
			require.NoError(t, err)
			assert.Equal(t, http.StatusNoContent, resp.StatusCode)

			actualRTs, _, err := reg.RelationTupleManager().GetRelationTuples(context.Background(), &relationtuple.RelationQuery{
				Namespace: nspace.Name,
			})
			require.NoError(t, err)
			assert.Equal(t, []*relationtuple.InternalRelationTuple{}, actualRTs)
		})

		t.Run("case=valid JSON, invalid content", func(t *testing.T) {
			rawJSON := `
[
    {
        "action": "insert",
        "namespace":"role",
        "object":"super-admin",
        "relation":"member",
        "subject":"role:company-admin"
    }
]`
			req, err := http.NewRequest(http.MethodPatch, ts.URL+relationtuple.RouteBase, bytes.NewBufferString(rawJSON))
			require.NoError(t, err)
			resp, err := ts.Client().Do(req)
			require.NoError(t, err)
			assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

			defer resp.Body.Close()
			errContent, err := io.ReadAll(resp.Body)
			require.NoError(t, err)
			assert.Contains(t, string(errContent), "relation_tuple is missing")
		})

		t.Run("case=unknown action", func(t *testing.T) {
			rawJSON := `
[
	{
		"action": "unknown_action_foo",
		"relation_tuple": {
			"namespace":"role",
			"object":"super-admin",
			"relation":"member",
			"subject_id":"role:company-admin"
		}
	}
]`
			req, err := http.NewRequest(http.MethodPatch, ts.URL+relationtuple.RouteBase, bytes.NewBufferString(rawJSON))
			require.NoError(t, err)
			resp, err := ts.Client().Do(req)
			require.NoError(t, err)
			assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

			defer resp.Body.Close()
			errContent, err := io.ReadAll(resp.Body)
			require.NoError(t, err)
			assert.Contains(t, string(errContent), "unknown_action_foo")
		})
	})
}
