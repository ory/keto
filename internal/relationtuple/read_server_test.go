package relationtuple_test

import (
	"context"
	"encoding/json"
	"io"
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
	setup := func(t *testing.T) (ts *httptest.Server, reg driver.Registry) {
		r := &x.ReadRouter{Router: httprouter.New()}
		reg = driver.NewMemoryTestRegistry(t, []*namespace.Namespace{{Name: t.Name()}})
		h := relationtuple.NewHandler(reg)
		h.RegisterReadRoutes(r)
		ts = httptest.NewServer(r)
		t.Cleanup(ts.Close)

		return
	}

	t.Run("method=get", func(t *testing.T) {
		t.Run("case=empty response is not nil", func(t *testing.T) {
			ts, _ := setup(t)

			resp, err := ts.Client().Get(ts.URL + relationtuple.RouteBase + "?namespace=" + url.QueryEscape(t.Name()))
			require.NoError(t, err)

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

		t.Run("case=gets first page", func(t *testing.T) {
			ts, reg := setup(t)

			rts := []*relationtuple.InternalRelationTuple{
				{
					Namespace: t.Name(),
					Object:    "o1",
					Relation:  "r1",
					Subject:   &relationtuple.SubjectID{ID: "s1"},
				},
				{
					Namespace: t.Name(),
					Object:    "o2",
					Relation:  "r2",
					Subject: &relationtuple.SubjectSet{
						Namespace: t.Name(),
						Object:    "o1",
						Relation:  "r1",
					},
				},
			}

			require.NoError(t, reg.RelationTupleManager().WriteRelationTuples(context.Background(), rts...))

			resp, err := ts.Client().Get(ts.URL + relationtuple.RouteBase + "?namespace=" + url.QueryEscape(t.Name()))
			require.NoError(t, err)

			var respMsg relationtuple.GetResponse
			require.NoError(t, json.NewDecoder(resp.Body).Decode(&respMsg))

			for i := range rts {
				assert.Contains(t, respMsg.RelationTuples, rts[i])
			}
			assert.Equal(t, x.PageTokenEnd, respMsg.NextPageToken)
		})
	})
}
