package expand_test

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/ory/keto/internal/driver/config"

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

	reg := driver.NewSqliteTestRegistry(t, false)
	require.NoError(t, reg.Config(context.Background()).Set(config.KeyNamespaces, []*namespace.Namespace{nspace}))
	h := expand.NewHandler(reg)
	r := httprouter.New()
	h.RegisterReadRoutes(&x.ReadRouter{Router: r})
	ts := httptest.NewServer(r)
	defer ts.Close()

	t.Run("case=returns bad request on malformed int", func(t *testing.T) {
		resp, err := ts.Client().Get(ts.URL + expand.RouteBase + "?max-depth=foo")
		require.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
		body, err := io.ReadAll(resp.Body)
		require.NoError(t, err)
		assert.Contains(t, string(body), "invalid syntax")
	})

	t.Run("case=returns not found on unknown namespace", func(t *testing.T) {
		resp, err := ts.Client().Get(ts.URL + expand.RouteBase + "?" + url.Values{
			"max-depth": {"10"},
			"namespace": {"not " + nspace.Name},
		}.Encode())
		require.NoError(t, err)

		assert.Equal(t, http.StatusNotFound, resp.StatusCode)
		body, err := io.ReadAll(resp.Body)
		require.NoError(t, err)
		assert.Contains(t, string(body), "Unknown namespace")
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

		relationtuple.MapAndWriteTuples(t, reg,
			&relationtuple.InternalRelationTuple{
				Namespace: nspace.Name,
				Object:    rootSub.Object,
				Relation:  rootSub.Relation,
				Subject:   expectedTree.Children[0].Subject,
			},
			&relationtuple.InternalRelationTuple{
				Namespace: nspace.Name,
				Object:    rootSub.Object,
				Relation:  rootSub.Relation,
				Subject:   expectedTree.Children[1].Subject,
			},
		)

		qs := rootSub.ToURLQuery()
		qs.Set("max-depth", "2")
		resp, err := ts.Client().Get(ts.URL + expand.RouteBase + "?" + qs.Encode())
		require.NoError(t, err)

		require.Equal(t, http.StatusOK, resp.StatusCode)

		actualTree := expand.Tree{}
		require.NoError(t, json.NewDecoder(resp.Body).Decode(&actualTree))
		assertEqualTrees(t, expectedTree, &actualTree)
	})
}

func assertEqualTrees(t *testing.T, expected, actual *expand.Tree) {
	t.Helper()
	assert.Truef(t, treesAreEqual(t, expected, actual),
		"expected:\n%s\n\nactual:\n%s", expected.String(), actual.String())
}

func treesAreEqual(t *testing.T, expected, actual *expand.Tree) bool {
	if expected == nil || actual == nil {
		return expected == actual
	}

	if expected.Type != actual.Type {
		t.Logf("expected type %q, actual type %q", expected.Type, actual.Type)
		return false
	}
	if expected.Subject.String() != actual.Subject.String() {
		t.Logf("expected subject: %q, actual subject: %q", expected.Subject.String(), actual.Subject.String())
		return false
	}
	if len(expected.Children) != len(actual.Children) {
		t.Logf("expected len(children)=%d, actual len(children)=%d", len(expected.Children), len(actual.Children))
		return false
	}

	// For children, we check for equality disregarding the order
outer:
	for _, expectedChild := range expected.Children {
		for _, actualChild := range actual.Children {
			if treesAreEqual(t, expectedChild, actualChild) {
				continue outer
			}
		}
		t.Logf("expected child:\n%s\n\nactual child:\n%s", expectedChild.String(), actual.String())
		return false
	}
	return true
}
