package relationtuple_test

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tidwall/gjson"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/require"

	"github.com/ory/keto/internal/driver"
	"github.com/ory/keto/internal/namespace"
	"github.com/ory/keto/internal/relationtuple"
	"github.com/ory/keto/internal/x"
)

func TestReadHandlers(t *testing.T) {
	r := &x.ReadRouter{Router: httprouter.New()}
	nspace := &namespace.Namespace{Name: "relation tuple read test"}
	reg := driver.NewMemoryTestRegistry(t, []*namespace.Namespace{nspace})
	h := relationtuple.NewHandler(reg)
	h.RegisterReadRoutes(r)
	ts := httptest.NewServer(r)
	t.Cleanup(ts.Close)

	t.Run("method=get", func(t *testing.T) {
		t.Run("case=empty response is not nil", func(t *testing.T) {
			resp, err := ts.Client().Get(ts.URL + relationtuple.RouteBase + "?" + url.Values{
				"namespace": {nspace.Name},
			}.Encode())
			require.NoError(t, err)
			require.Equal(t, http.StatusOK, resp.StatusCode)

			body, err := io.ReadAll(resp.Body)
			require.NoError(t, err)

			assert.Equal(t, "[]", gjson.GetBytes(body, "relation_tuples").Raw)

			var respMsg relationtuple.GetResponse
			require.NoError(t, json.Unmarshal(body, &respMsg))

			assert.Equal(t, relationtuple.GetResponse{
				RelationTuples: []*relationtuple.InternalRelationTuple{},
				NextPageToken:  x.PageTokenEnd,
				IsLastPage:     true,
			}, respMsg)
		})

		t.Run("case=returns tuples", func(t *testing.T) {
			rts := []*relationtuple.InternalRelationTuple{
				{
					Namespace: nspace.Name,
					Object:    "o1",
					Relation:  "r1",
					Subject:   &relationtuple.SubjectID{ID: "s1"},
				},
				{
					Namespace: nspace.Name,
					Object:    "o2",
					Relation:  "r2",
					Subject: &relationtuple.SubjectSet{
						Namespace: nspace.Name,
						Object:    "o1",
						Relation:  "r1",
					},
				},
			}

			require.NoError(t, reg.RelationTupleManager().WriteRelationTuples(context.Background(), rts...))

			resp, err := ts.Client().Get(ts.URL + relationtuple.RouteBase + "?" + url.Values{
				"namespace": {nspace.Name},
			}.Encode())
			require.NoError(t, err)

			var respMsg relationtuple.GetResponse
			require.NoError(t, json.NewDecoder(resp.Body).Decode(&respMsg))

			for i := range rts {
				assert.Contains(t, respMsg.RelationTuples, rts[i])
			}
			assert.Equal(t, x.PageTokenEnd, respMsg.NextPageToken)
			assert.True(t, respMsg.IsLastPage)
		})

		t.Run("case=returns bad request on malformed subject", func(t *testing.T) {
			resp, err := ts.Client().Get(ts.URL + relationtuple.RouteBase + "?" + url.Values{
				"subject": {"not#a valid subject"},
			}.Encode())
			require.NoError(t, err)

			assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
		})

		t.Run("case=paginates", func(t *testing.T) {
			// this obj is used to filter out tuples from other cases
			obj := t.Name()

			rts := []*relationtuple.InternalRelationTuple{
				{
					Namespace: nspace.Name,
					Object:    obj,
					Relation:  "r1",
					Subject:   &relationtuple.SubjectID{ID: "s1"},
				},
				{
					Namespace: nspace.Name,
					Object:    obj,
					Relation:  "r2",
					Subject:   &relationtuple.SubjectID{ID: "s2"},
				},
			}
			require.NoError(t, reg.RelationTupleManager().WriteRelationTuples(context.Background(), rts...))

			var firstResp relationtuple.GetResponse
			t.Run("case=first page", func(t *testing.T) {
				resp, err := ts.Client().Get(ts.URL + relationtuple.RouteBase + "?" + url.Values{
					"namespace": {nspace.Name},
					"object":    {obj},
					"page_size": {"1"},
				}.Encode())
				require.NoError(t, err)
				require.Equal(t, http.StatusOK, resp.StatusCode)

				require.NoError(t, json.NewDecoder(resp.Body).Decode(&firstResp))
				require.Len(t, firstResp.RelationTuples, 1)
				assert.Contains(t, rts, firstResp.RelationTuples[0])
				assert.NotEqual(t, x.PageTokenEnd, firstResp.NextPageToken)
				assert.False(t, firstResp.IsLastPage)
			})

			t.Run("case=second page", func(t *testing.T) {
				resp, err := ts.Client().Get(ts.URL + relationtuple.RouteBase + "?" + url.Values{
					"namespace":  {nspace.Name},
					"object":     {obj},
					"page_size":  {"1"},
					"page_token": {firstResp.NextPageToken},
				}.Encode())
				require.NoError(t, err)
				require.Equal(t, http.StatusOK, resp.StatusCode)

				secondResp := relationtuple.GetResponse{}
				require.NoError(t, json.NewDecoder(resp.Body).Decode(&secondResp))
				require.Len(t, secondResp.RelationTuples, 1)

				assert.NotEqual(t, firstResp.RelationTuples, secondResp.RelationTuples)
				assert.Contains(t, rts, secondResp.RelationTuples[0])
				assert.Equal(t, x.PageTokenEnd, secondResp.NextPageToken)
				assert.True(t, secondResp.IsLastPage)
			})
		})

		t.Run("case=returs bad request on invalid page size", func(t *testing.T) {
			resp, err := ts.Client().Get(ts.URL + relationtuple.RouteBase + "?" + url.Values{
				"page_size": {"foo"},
			}.Encode())
			require.NoError(t, err)

			body, err := io.ReadAll(resp.Body)
			require.NoError(t, err)
			assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
			assert.Contains(t, string(body), "invalid syntax")
		})
	})
}
