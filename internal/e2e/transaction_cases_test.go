// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package e2e

import (
	"testing"

	"github.com/ory/x/pointerx"

	"github.com/ory/keto/ketoapi"

	"github.com/stretchr/testify/assert"

	"github.com/ory/keto/internal/namespace"
)

func runTransactionCases(c transactClient, m *namespaceTestManager) func(*testing.T) {
	return func(t *testing.T) {
		t.Parallel()
		t.Run("case=create and delete", func(t *testing.T) {
			n := &namespace.Namespace{Name: t.Name()}
			m.add(t, n)

			tuples := []*ketoapi.RelationTuple{
				{
					Namespace: n.Name,
					Object:    "o",
					Relation:  "rel",
					SubjectSet: &ketoapi.SubjectSet{
						Namespace: n.Name,
						Object:    "o",
						Relation:  "rel",
					},
				},
				{
					Namespace: n.Name,
					Object:    "o",
					Relation:  "rel",
					SubjectID: pointerx.Ptr("sid"),
				},
			}

			c.transactTuples(t, tuples, nil)

			resp := c.queryTuple(t, &ketoapi.RelationQuery{
				Namespace: &n.Name,
			})
			for i := range tuples {
				assert.Contains(t, resp.RelationTuples, tuples[i])
			}

			c.transactTuples(t, nil, tuples)

			resp = c.queryTuple(t, &ketoapi.RelationQuery{
				Namespace: &n.Name,
			})
			assert.Len(t, resp.RelationTuples, 0)
		})

		t.Run("case=expand-api-display-access docs code sample", func(t *testing.T) {
			files := &namespace.Namespace{Name: t.Name() + "files"}
			directories := &namespace.Namespace{Name: t.Name() + "directories"}
			m.add(t, files, directories)

			tuples := []*ketoapi.RelationTuple{
				{
					Namespace: directories.Name,
					Object:    "/photos",
					Relation:  "owner",
					SubjectID: pointerx.Ptr("maureen"),
				},
				{
					Namespace: files.Name,
					Object:    "/photos/beach.jpg",
					Relation:  "owner",
					SubjectID: pointerx.Ptr("maureen"),
				},
				{
					Namespace: files.Name,
					Object:    "/photos/mountains.jpg",
					Relation:  "owner",
					SubjectID: pointerx.Ptr("laura"),
				},
				{
					Namespace: directories.Name,
					Object:    "/photos",
					Relation:  "access",
					SubjectID: pointerx.Ptr("laura"),
				},
			}
			for _, o := range []struct{ n, o string }{
				{files.Name, "/photos/beach.jpg"},
				{files.Name, "/photos/mountains.jpg"},
				{directories.Name, "/photos"},
			} {
				tuples = append(tuples, &ketoapi.RelationTuple{
					Namespace: o.n,
					Object:    o.o,
					Relation:  "access",
					SubjectSet: &ketoapi.SubjectSet{
						Namespace: o.n,
						Object:    o.o,
						Relation:  "owner",
					},
				})
			}
			for _, obj := range []string{"/photos/beach.jpg", "/photos/mountains.jpg"} {
				tuples = append(tuples, &ketoapi.RelationTuple{
					Namespace: files.Name,
					Object:    obj,
					Relation:  "access",
					SubjectSet: &ketoapi.SubjectSet{
						Namespace: directories.Name,
						Object:    "/photos",
						Relation:  "access",
					},
				})
			}

			c.transactTuples(t, tuples, nil)

			resp := c.queryTuple(t, &ketoapi.RelationQuery{})
			assert.Equal(t, len(tuples), len(resp.RelationTuples))
			for i := range tuples {
				assert.Contains(t, resp.RelationTuples, tuples[i])
			}

			c.transactTuples(t, nil, tuples)

			resp = c.queryTuple(t, &ketoapi.RelationQuery{})
			assert.Len(t, resp.RelationTuples, 0)
		})
	}
}
