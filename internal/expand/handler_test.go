// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package expand_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/ory/x/pointerx"

	"github.com/ory/keto/ketoapi"

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
		rootSub := &ketoapi.SubjectSet{
			Namespace: nspace.Name,
			Object:    "root",
			Relation:  "parent of",
		}
		expectedTree := &ketoapi.Tree[*ketoapi.RelationTuple]{
			Type: ketoapi.TreeNodeUnion,
			Tuple: &ketoapi.RelationTuple{
				SubjectSet: rootSub,
			},
			Children: []*ketoapi.Tree[*ketoapi.RelationTuple]{
				{
					Type: ketoapi.TreeNodeLeaf,
					Tuple: &ketoapi.RelationTuple{
						SubjectID: pointerx.Ptr("child0"),
					},
				},
				{
					Type: ketoapi.TreeNodeLeaf,
					Tuple: &ketoapi.RelationTuple{
						SubjectID: pointerx.Ptr("child1"),
					},
				},
			},
		}

		relationtuple.MapAndWriteTuples(t, reg,
			&ketoapi.RelationTuple{
				Namespace: nspace.Name,
				Object:    rootSub.Object,
				Relation:  rootSub.Relation,
				SubjectID: expectedTree.Children[0].Tuple.SubjectID,
			},
			&ketoapi.RelationTuple{
				Namespace: nspace.Name,
				Object:    rootSub.Object,
				Relation:  rootSub.Relation,
				SubjectID: expectedTree.Children[1].Tuple.SubjectID,
			},
		)

		qs := rootSub.ToURLQuery()
		qs.Set("max-depth", "2")
		resp, err := ts.Client().Get(ts.URL + expand.RouteBase + "?" + qs.Encode())
		require.NoError(t, err)

		require.Equal(t, http.StatusOK, resp.StatusCode)

		actualTree := ketoapi.Tree[*ketoapi.RelationTuple]{}
		body, err := io.ReadAll(resp.Body)
		require.NoError(t, err)
		t.Logf("body: %s", string(body))
		require.NoError(t, json.NewDecoder(bytes.NewBuffer(body)).Decode(&actualTree))
		expand.AssertExternalTreesAreEqual(t, expectedTree, &actualTree)
	})
}
