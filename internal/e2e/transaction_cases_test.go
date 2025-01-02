// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package e2e

import (
	"cmp"
	"slices"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ory/x/pointerx"

	"github.com/ory/keto/internal/namespace"
	"github.com/ory/keto/internal/x"
	"github.com/ory/keto/ketoapi"
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

		t.Run("case=duplicate string representations", func(t *testing.T) {
			n := &namespace.Namespace{Name: t.Name()}
			m.add(t, n)
			c.transactTuples(t, []*ketoapi.RelationTuple{
				{
					Namespace: n.Name,
					Object:    "o",
					Relation:  "rel",
					SubjectID: pointerx.Ptr("sid"),
				},
				{
					Namespace: n.Name,
					Object:    "o",
					Relation:  "rel",
					SubjectID: pointerx.Ptr("sid"),
				},
			}, nil)
		})

		t.Run("case=large inserts and deletes", func(t *testing.T) {
			if !testing.Short() {
				t.Skip("This test is fairly expensive, especially the deletion.")
			}

			ns := []*namespace.Namespace{
				{Name: t.Name() + "1"},
				{Name: t.Name() + "2"},
			}
			m.add(t, ns...)

			var tuples []*ketoapi.RelationTuple
			//for i := range 12001 {
			for i := range 12001 {
				tuples = append(tuples,
					&ketoapi.RelationTuple{
						Namespace: ns[0].Name,
						Object:    "o" + strconv.Itoa(i),
						Relation:  "rela",
						SubjectSet: &ketoapi.SubjectSet{
							Namespace: ns[1].Name,
							Object:    "o" + strconv.Itoa(i),
							Relation:  "relx",
						},
					},
					&ketoapi.RelationTuple{
						Namespace: ns[0].Name,
						Object:    "o" + strconv.Itoa(i),
						Relation:  "relb",
						SubjectID: pointerx.Ptr("sid"),
					},
				)
			}

			t0 := time.Now()
			c.transactTuples(t, tuples, nil)
			t.Log("insert", time.Since(t0))

			t0 = time.Now()
			var resp []*ketoapi.RelationTuple
			var pt string
			for {
				r := c.queryTuple(t, &ketoapi.RelationQuery{
					Namespace: &ns[0].Name,
				}, x.WithSize(1000), x.WithToken(pt))
				resp = append(resp, r.RelationTuples...)
				pt = r.NextPageToken
				if pt == "" {
					break
				}
			}
			t.Log("query", time.Since(t0))

			sort := func(a, b *ketoapi.RelationTuple) int {
				return cmp.Or(
					cmp.Compare(a.Namespace, b.Namespace),
					cmp.Compare(a.Object, b.Object),
					cmp.Compare(a.Relation, b.Relation),
				)
			}
			t0 = time.Now()
			slices.SortFunc(resp, sort)
			slices.SortFunc(tuples, sort)
			t.Log("sort", time.Since(t0))

			t0 = time.Now()
			require.Equal(t, tuples, resp)
			t.Log("equal", time.Since(t0))

			t0 = time.Now()
			c.transactTuples(t, nil, tuples)
			t.Log(t.Name(), "delete took:", time.Since(t0))

			resp = c.queryTuple(t, &ketoapi.RelationQuery{
				Namespace: &ns[0].Name,
			}).RelationTuples
			assert.Len(t, resp, 0)
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
