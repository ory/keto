// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package relationtuple_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/julienschmidt/httprouter"
	keysetpagination "github.com/ory/x/pagination/keysetpagination_v2"
	"github.com/ory/x/pointerx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ory/keto/internal/driver"
	"github.com/ory/keto/internal/driver/config"
	"github.com/ory/keto/internal/namespace"
	"github.com/ory/keto/internal/relationtuple"
	"github.com/ory/keto/internal/x"
	"github.com/ory/keto/ketoapi"
)

func TestWriteHandlers(t *testing.T) {
	ctx := context.Background()
	r := httprouter.New()
	wr := &x.WriteRouter{Router: r}
	rr := &x.ReadRouter{Router: r}
	reg := driver.NewSqliteTestRegistry(t, false)

	var nspaces []*namespace.Namespace
	addNamespace := func(t *testing.T) *namespace.Namespace {
		n := &namespace.Namespace{
			Name: t.Name(),
		}
		nspaces = append(nspaces, n)

		require.NoError(t, reg.Config(ctx).Set(config.KeyNamespaces, nspaces))

		return n
	}

	h := relationtuple.NewHandler(reg)
	h.RegisterWriteRoutes(wr)
	h.RegisterReadRoutes(rr)
	ts := httptest.NewServer(r)
	defer ts.Close()

	t.Run("method=create", func(t *testing.T) {
		doCreate := func(raw []byte) *http.Response {
			req, err := http.NewRequest(http.MethodPut, ts.URL+relationtuple.WriteRouteBase, bytes.NewBuffer(raw))
			require.NoError(t, err)
			resp, err := ts.Client().Do(req)
			require.NoError(t, err)

			return resp
		}

		t.Run("case=creates tuple", func(t *testing.T) {
			nspace := addNamespace(t)

			rt := &ketoapi.RelationTuple{
				Namespace: nspace.Name,
				Object:    "obj",
				Relation:  "rel",
				SubjectID: pointerx.Ptr("subj"),
			}
			payload, err := json.Marshal(rt)
			require.NoError(t, err)

			resp := doCreate(payload)

			assert.Equal(t, http.StatusCreated, resp.StatusCode)

			body, err := io.ReadAll(resp.Body)
			require.NoError(t, err)

			assert.JSONEq(t, string(payload), string(body))

			t.Run("check=is contained in the manager", func(t *testing.T) {
				mapped, err := reg.Mapper().FromTuple(ctx, rt)
				require.NoError(t, err)
				// set a size > 1 just to make sure it gets all
				actualRTs, _, err := reg.RelationTupleManager().GetRelationTuples(ctx, mapped[0].ToQuery(), keysetpagination.WithSize(10))
				require.NoError(t, err)
				actual, err := reg.Mapper().ToTuple(ctx, actualRTs...)
				require.NoError(t, err)
				assert.Equalf(t, []*ketoapi.RelationTuple{rt}, actual, "want: %s\ngot:  %s", rt.String(), actual[0].String())
			})

			t.Run("check=is gettable with the returned URL", func(t *testing.T) {
				resp, err := ts.Client().Get(ts.URL + resp.Header.Get("Location"))
				require.NoError(t, err)
				require.Equal(t, http.StatusOK, resp.StatusCode)

				respDec := ketoapi.GetResponse{}
				require.NoError(t, json.NewDecoder(resp.Body).Decode(&respDec))
				assert.Equal(t, []*ketoapi.RelationTuple{rt}, respDec.RelationTuples)
			})
		})

		t.Run("case=returns bad request on JSON parse error", func(t *testing.T) {
			resp := doCreate([]byte("foo"))
			assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
		})

		t.Run("case=special chars", func(t *testing.T) {
			nspace := addNamespace(t)

			rts := []*ketoapi.RelationTuple{
				{
					Namespace: nspace.Name,
					Object:    "group:B",
					Relation:  "member",
					SubjectSet: &ketoapi.SubjectSet{
						Namespace: nspace.Name,
						Object:    "group:A",
						Relation:  "member",
					},
				},
				{
					Namespace: nspace.Name,
					Object:    "@all",
					Relation:  "member",
					SubjectID: pointerx.Ptr("this:could#be interpreted:as a@subject set"),
				},
			}

			for _, rt := range rts {
				payload, err := json.Marshal(rt)
				require.NoError(t, err)

				resp := doCreate(payload)
				assert.Equal(t, http.StatusCreated, resp.StatusCode)
			}

			actual, next, err := reg.RelationTupleManager().GetRelationTuples(ctx, &relationtuple.RelationQuery{
				Namespace: &nspace.Name,
			})
			require.NoError(t, err)
			actualMapped, err := reg.Mapper().ToTuple(ctx, actual...)
			require.NoError(t, err)
			assert.True(t, next.IsLast())
			assert.ElementsMatch(t, rts, actualMapped)
		})
	})

	t.Run("method=delete", func(t *testing.T) {
		t.Run("case=deletes a tuple", func(t *testing.T) {
			nspace := addNamespace(t)

			rt := &ketoapi.RelationTuple{
				Namespace: nspace.Name,
				Object:    "deleted obj",
				Relation:  "deleted rel",
				SubjectID: pointerx.Ptr("deleted subj"),
			}
			relationtuple.MapAndWriteTuples(t, reg, rt)

			req, err := http.NewRequest(http.MethodDelete, ts.URL+relationtuple.WriteRouteBase+"?"+rt.ToURLQuery().Encode(), nil)
			require.NoError(t, err)
			resp, err := ts.Client().Do(req)
			require.NoError(t, err)
			assert.Equal(t, http.StatusNoContent, resp.StatusCode)

			// set a size > 1 just to make sure it gets all
			actualRTs, _, err := reg.RelationTupleManager().GetRelationTuples(ctx, &relationtuple.RelationQuery{Namespace: &nspace.Name}, keysetpagination.WithSize(10))
			require.NoError(t, err)
			assert.Equal(t, []*relationtuple.RelationTuple{}, actualRTs)
		})

		t.Run("case=deletes multiple tuples", func(t *testing.T) {
			nspace := addNamespace(t)

			rts := []*ketoapi.RelationTuple{
				{
					Namespace: nspace.Name,
					Object:    "deleted obj",
					Relation:  "deleted rel",
					SubjectID: pointerx.Ptr("deleted subj 1"),
				},
				{
					Namespace: nspace.Name,
					Object:    "deleted obj",
					Relation:  "deleted rel",
					SubjectID: pointerx.Ptr("deleted subj 2"),
				},
			}

			relationtuple.MapAndWriteTuples(t, reg, rts...)

			q := url.Values{
				"namespace": {nspace.Name},
				"object":    {"deleted obj"},
				"relation":  {"deleted rel"},
			}
			req, err := http.NewRequest(http.MethodDelete, ts.URL+relationtuple.WriteRouteBase+"?"+q.Encode(), nil)
			require.NoError(t, err)
			resp, err := ts.Client().Do(req)
			require.NoError(t, err)
			assert.Equal(t, http.StatusNoContent, resp.StatusCode)

			query, err := (&ketoapi.RelationQuery{}).FromURLQuery(q)
			require.NoError(t, err)
			mappedQuery, err := reg.Mapper().FromQuery(ctx, query)
			require.NoError(t, err)

			actualRTs, _, err := reg.RelationTupleManager().GetRelationTuples(ctx, mappedQuery, keysetpagination.WithSize(10))
			require.NoError(t, err)
			assert.Equal(t, []*relationtuple.RelationTuple{}, actualRTs)
		})

		t.Run("suite=bad requests", func(t *testing.T) {
			nspace := addNamespace(t)

			rts := []*ketoapi.RelationTuple{
				{
					Namespace: nspace.Name,
					Object:    "deleted obj",
					Relation:  "deleted rel",
					SubjectID: pointerx.Ptr("deleted subj 1"),
				},
				{
					Namespace: nspace.Name,
					Object:    "deleted obj",
					Relation:  "deleted rel",
					SubjectID: pointerx.Ptr("deleted subj 2"),
				},
			}

			relationtuple.MapAndWriteTuples(t, reg, rts...)

			assertBadRequest := func(t *testing.T, req *http.Request) {
				resp, err := ts.Client().Do(req)
				require.NoError(t, err)
				assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
			}

			assertTuplesExist := func(t *testing.T) {
				mappedQuery, err := reg.Mapper().FromQuery(ctx, &ketoapi.RelationQuery{
					Namespace: &nspace.Name,
				})
				require.NoError(t, err)

				actualRTs, _, err := reg.RelationTupleManager().GetRelationTuples(ctx, mappedQuery, keysetpagination.WithSize(10))
				require.NoError(t, err)
				mappedRTs, err := reg.Mapper().ToTuple(ctx, actualRTs...)
				require.NoError(t, err)
				assert.ElementsMatch(t, rts, mappedRTs)
			}

			t.Run("case=bad request if body sent", func(t *testing.T) {
				q := url.Values{
					"namespace": {nspace.Name},
					"object":    {"deleted obj"},
					"relation":  {"deleted rel"},
				}
				req, err := http.NewRequest(
					http.MethodDelete,
					ts.URL+relationtuple.WriteRouteBase+"?"+q.Encode(),
					strings.NewReader("some body"))
				require.NoError(t, err)

				assertBadRequest(t, req)
				assertTuplesExist(t)
			})

			t.Run("case=bad request query param misspelled", func(t *testing.T) {
				req, err := http.NewRequest(
					http.MethodDelete,
					ts.URL+relationtuple.WriteRouteBase+"?invalid=param",
					nil)
				require.NoError(t, err)

				assertBadRequest(t, req)
				assertTuplesExist(t)
			})

			t.Run("case=bad request if query params misssing", func(t *testing.T) {
				req, err := http.NewRequest(
					http.MethodDelete,
					ts.URL+relationtuple.WriteRouteBase,
					nil)
				require.NoError(t, err)

				assertBadRequest(t, req)
				assertTuplesExist(t)
			})
		})

	})

	t.Run("method=patch", func(t *testing.T) {
		t.Run("case=create and delete", func(t *testing.T) {
			nspace := addNamespace(t)
			relation := t.Name()

			deltas := []*ketoapi.PatchDelta{
				{
					Action: ketoapi.ActionInsert,
					RelationTuple: &ketoapi.RelationTuple{
						Namespace: nspace.Name,
						Object:    "create obj",
						Relation:  relation,
						SubjectID: pointerx.Ptr("create sub"),
					},
				},
				{
					Action: ketoapi.ActionDelete,
					RelationTuple: &ketoapi.RelationTuple{
						Namespace: nspace.Name,
						Object:    "delete obj",
						Relation:  relation,
						SubjectID: pointerx.Ptr("delete sub"),
					},
				},
			}
			relationtuple.MapAndWriteTuples(t, reg, deltas[1].RelationTuple)

			body, err := json.Marshal(deltas)
			require.NoError(t, err)
			req, err := http.NewRequest(http.MethodPatch, ts.URL+relationtuple.WriteRouteBase, bytes.NewBuffer(body))
			require.NoError(t, err)
			resp, err := ts.Client().Do(req)
			require.NoError(t, err)
			assert.Equal(t, http.StatusNoContent, resp.StatusCode)

			actualRTs, _, err := reg.RelationTupleManager().GetRelationTuples(ctx, &relationtuple.RelationQuery{
				Namespace: &nspace.Name,
				Relation:  &relation,
			})
			require.NoError(t, err)
			mapped, err := reg.Mapper().ToTuple(ctx, actualRTs...)
			require.NoError(t, err)
			assert.Equal(t, []*ketoapi.RelationTuple{deltas[0].RelationTuple}, mapped)
		})

		t.Run("case=ignores rest on err", func(t *testing.T) {
			nspace := addNamespace(t)

			deltas := []*ketoapi.PatchDelta{
				{
					Action: ketoapi.ActionInsert,
					RelationTuple: &ketoapi.RelationTuple{
						Namespace: nspace.Name,
						Object:    "create obj",
						Relation:  t.Name(),
						SubjectID: pointerx.Ptr("create sub"),
					},
				},
				{
					Action: ketoapi.ActionDelete,
					RelationTuple: &ketoapi.RelationTuple{
						Namespace: "not " + nspace.Name,
						Object:    "o",
						Relation:  "r",
						SubjectID: pointerx.Ptr("s"),
					},
				},
			}

			body, err := json.Marshal(deltas)
			require.NoError(t, err)
			req, err := http.NewRequest(http.MethodPatch, ts.URL+relationtuple.WriteRouteBase, bytes.NewBuffer(body))
			require.NoError(t, err)
			resp, err := ts.Client().Do(req)
			require.NoError(t, err)
			assert.Equal(t, http.StatusNotFound, resp.StatusCode)

			// set a size > 1 just to make sure it gets all
			actualRTs, _, err := reg.RelationTupleManager().GetRelationTuples(ctx, &relationtuple.RelationQuery{Namespace: &nspace.Name}, keysetpagination.WithSize(10))
			require.NoError(t, err)
			assert.Len(t, actualRTs, 0)
		})

		t.Run("case=only create", func(t *testing.T) {
			nspace := addNamespace(t)

			deltas := []*ketoapi.PatchDelta{
				{
					Action: ketoapi.ActionInsert,
					RelationTuple: &ketoapi.RelationTuple{
						Namespace: nspace.Name,
						Object:    "create obj",
						Relation:  "rel",
						SubjectID: pointerx.Ptr("create sub"),
					},
				},
			}

			body, err := json.Marshal(deltas)
			require.NoError(t, err)
			req, err := http.NewRequest(http.MethodPatch, ts.URL+relationtuple.WriteRouteBase, bytes.NewBuffer(body))
			require.NoError(t, err)
			resp, err := ts.Client().Do(req)
			require.NoError(t, err)
			assert.Equal(t, http.StatusNoContent, resp.StatusCode)

			actualRTs, _, err := reg.RelationTupleManager().GetRelationTuples(ctx, &relationtuple.RelationQuery{
				Namespace: &nspace.Name,
			})
			require.NoError(t, err)
			mapped, err := reg.Mapper().ToTuple(ctx, actualRTs...)
			require.NoError(t, err)
			assert.Equal(t, []*ketoapi.RelationTuple{deltas[0].RelationTuple}, mapped)
		})

		t.Run("case=only delete", func(t *testing.T) {
			nspace := addNamespace(t)

			deltas := []*ketoapi.PatchDelta{
				{
					Action: ketoapi.ActionDelete,
					RelationTuple: &ketoapi.RelationTuple{
						Namespace: nspace.Name,
						Object:    "delete obj",
						Relation:  "rel",
						SubjectID: pointerx.Ptr("delete sub"),
					},
				},
			}

			body, err := json.Marshal(deltas)
			require.NoError(t, err)
			req, err := http.NewRequest(http.MethodPatch, ts.URL+relationtuple.WriteRouteBase, bytes.NewBuffer(body))
			require.NoError(t, err)
			resp, err := ts.Client().Do(req)
			require.NoError(t, err)
			assert.Equal(t, http.StatusNoContent, resp.StatusCode)

			actualRTs, _, err := reg.RelationTupleManager().GetRelationTuples(ctx, &relationtuple.RelationQuery{
				Namespace: &nspace.Name,
			})
			require.NoError(t, err)
			assert.Equal(t, []*relationtuple.RelationTuple{}, actualRTs)
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
			req, err := http.NewRequest(http.MethodPatch, ts.URL+relationtuple.WriteRouteBase, bytes.NewBufferString(rawJSON))
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
			req, err := http.NewRequest(http.MethodPatch, ts.URL+relationtuple.WriteRouteBase, bytes.NewBufferString(rawJSON))
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
