package expand_test

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ory/keto/internal/driver"
	"github.com/ory/keto/internal/expand"
	"github.com/ory/keto/internal/namespace"
	"github.com/ory/keto/internal/relationtuple"
	"github.com/ory/keto/internal/x"
)

func TestRESTHandler(t *testing.T) {
	nspace := &namespace.Namespace{
		Name: "expand handler",
	}

	reg := driver.NewMemoryTestRegistry(t, []*namespace.Namespace{nspace})
	h := expand.NewHandler(reg)
	r := httprouter.New()
	h.RegisterReadRoutes(&x.ReadRouter{Router: r})
	ts := httptest.NewServer(r)
	defer ts.Close()

	t.Run("case=returns bad request on malformed int", func(t *testing.T) {
		resp, err := ts.Client().Get(ts.URL + expand.RouteBase + "?depth=foo")
		require.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
		body, err := io.ReadAll(resp.Body)
		require.NoError(t, err)
		assert.Contains(t, string(body), "invalid syntax")
	})

	t.Run("case=returns not found on unknown namespace", func(t *testing.T) {
		resp, err := ts.Client().Get(ts.URL + expand.RouteBase + "?" + url.Values{
			"depth":     {"10"},
			"namespace": {"not " + nspace.Name},
		}.Encode())
		require.NoError(t, err)

		assert.Equal(t, http.StatusNotFound, resp.StatusCode)
		body, err := io.ReadAll(resp.Body)
		require.NoError(t, err)
		assert.Contains(t, string(body), "unknown namespace")
	})

	t.Run("case=returns expand tree", func(t *testing.T) {
		rootSub := &relationtuple.SubjectSet{
			Namespace: nspace.Name,
			Object:    "root",
			Relation:  "parent of",
		}
		expectedTree := &expand.Tree{
			Type:    expand.Union,
			Subject: rootSub,
			Children: []*expand.Tree{
				{
					Type:    expand.Leaf,
					Subject: &relationtuple.SubjectID{ID: "child0"},
				},
				{
					Type:    expand.Leaf,
					Subject: &relationtuple.SubjectID{ID: "child1"},
				},
			},
		}

		require.NoError(t, reg.RelationTupleManager().WriteRelationTuples(context.Background(), []*relationtuple.InternalRelationTuple{
			{
				Namespace: nspace.Name,
				Object:    rootSub.Object,
				Relation:  rootSub.Relation,
				Subject:   expectedTree.Children[0].Subject,
			},
			{
				Namespace: nspace.Name,
				Object:    rootSub.Object,
				Relation:  rootSub.Relation,
				Subject:   expectedTree.Children[1].Subject,
			},
		}...))

		qs := rootSub.ToURLQuery()
		qs.Set("depth", "2")
		resp, err := ts.Client().Get(ts.URL + expand.RouteBase + "?" + qs.Encode())
		require.NoError(t, err)

		require.Equal(t, http.StatusOK, resp.StatusCode)

		actualTree := expand.Tree{}
		require.NoError(t, json.NewDecoder(resp.Body).Decode(&actualTree))
		assert.Equal(t, expectedTree, &actualTree)
	})
}
