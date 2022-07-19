package e2e

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/ory/keto/internal/expand"
	"github.com/ory/keto/ketoapi"

	"github.com/ory/herodot"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ory/keto/internal/namespace"
	"github.com/ory/keto/internal/x"
)

func runCases(c client, m *namespaceTestManager) func(*testing.T) {
	return func(t *testing.T) {
		c.waitUntilLive(t)

		t.Run("case=gets empty namespace", func(t *testing.T) {
			n := &namespace.Namespace{Name: t.Name()}
			m.add(t, n)

			resp := c.queryTuple(t, &ketoapi.RelationQuery{Namespace: &n.Name})
			assert.Len(t, resp.RelationTuples, 0)
		})

		t.Run("case=creates tuple and uses it then", func(t *testing.T) {
			n := &namespace.Namespace{Name: t.Name()}
			m.add(t, n)

			tuple := &ketoapi.RelationTuple{
				Namespace: n.Name,
				Object:    fmt.Sprintf("object for client %T", c),
				Relation:  "access",
				SubjectID: x.Ptr("client"),
			}

			c.createTuple(t, tuple)

			resp := c.queryTuple(t, &ketoapi.RelationQuery{Namespace: &tuple.Namespace})
			require.Len(t, resp.RelationTuples, 1)
			assert.Equal(t, tuple, resp.RelationTuples[0])

			// try the check API to see whether the tuple is interpreted correctly
			assert.True(t, c.check(t, tuple))
		})

		t.Run("case=expand API", func(t *testing.T) {
			n := &namespace.Namespace{Name: t.Name()}
			m.add(t, n)

			obj := fmt.Sprintf("tree for client %T", c)
			rel := "expand"

			subjects := []string{"s1", "s2"}
			expectedTree := &ketoapi.ExpandTree{
				Type: ketoapi.Union,
				SubjectSet: &ketoapi.SubjectSet{
					Namespace: n.Name,
					Object:    obj,
					Relation:  rel,
				},
				Children: make([]*ketoapi.ExpandTree, len(subjects)),
			}

			for i, subjectID := range subjects {
				c.createTuple(t, &ketoapi.RelationTuple{
					Namespace: n.Name,
					Object:    obj,
					Relation:  rel,
					SubjectID: &subjectID,
				})
				expectedTree.Children[i] = &ketoapi.ExpandTree{
					Type:      ketoapi.Leaf,
					SubjectID: &subjectID,
				}
			}

			actualTree := c.expand(t, expectedTree.SubjectSet, 100)

			assert.Equal(t, expectedTree.Type, actualTree.Type)
			assert.Equal(t, expectedTree.SubjectSet, actualTree.SubjectSet)
			assert.Equal(t, expectedTree.SubjectID, actualTree.SubjectID)
			assert.Equal(t, len(expectedTree.Children), len(actualTree.Children), "expected: %+v; actual: %+v", expectedTree.Children, actualTree.Children)

			expand.AssertExternalTreesAreEqual(t, expectedTree, actualTree)
		})

		t.Run("case=gets result paginated", func(t *testing.T) {
			const nTuples = 10
			n := &namespace.Namespace{Name: t.Name()}
			m.add(t, n)

			rel := fmt.Sprintf("some unique relation %T", c)
			for i := 0; i < nTuples; i++ {
				c.createTuple(t, &ketoapi.RelationTuple{
					Namespace: n.Name,
					Object:    "o" + strconv.Itoa(i),
					Relation:  rel,
					SubjectID: x.Ptr("s" + strconv.Itoa(i)),
				})
			}

			var (
				resp   ketoapi.GetResponse
				nPages int
			)
			// do ... while resp.NextPageToken != ""
			for ok := true; ok; ok = resp.NextPageToken != "" {
				resp = *c.queryTuple(t,
					&ketoapi.RelationQuery{
						Namespace: &n.Name,
						Relation:  &rel,
					},
					x.WithToken(resp.NextPageToken),
					x.WithSize(1),
				)
				nPages++
				assert.Len(t, resp.RelationTuples, 1)
			}

			assert.Equal(t, nTuples, nPages)
		})

		t.Run("case=deletes tuple", func(t *testing.T) {
			n := &namespace.Namespace{Name: t.Name()}
			m.add(t, n)

			for _, rt := range []*ketoapi.RelationTuple{
				{
					SubjectID: x.Ptr("s"),
				},
				{
					SubjectSet: &ketoapi.SubjectSet{
						Namespace: n.Name,
						Object:    "so",
						Relation:  "rel",
					},
				},
			} {
				t.Run(fmt.Sprintf("subject_id=%v", rt.SubjectID == nil), func(t *testing.T) {
					rt.Namespace = n.Name
					rt.Object = "o"
					rt.Relation = "rel"
					c.createTuple(t, rt)

					resp := c.queryTuple(t, &ketoapi.RelationQuery{Namespace: &n.Name})
					assert.Equal(t, []*ketoapi.RelationTuple{rt}, resp.RelationTuples)

					c.deleteTuple(t, rt)

					resp = c.queryTuple(t, &ketoapi.RelationQuery{Namespace: &n.Name})
					assert.Len(t, resp.RelationTuples, 0)
				})
			}
		})

		t.Run("case=deletes tuples based on relation query", func(t *testing.T) {
			n := &namespace.Namespace{Name: t.Name()}
			m.add(t, n)

			rts := []*ketoapi.RelationTuple{
				{
					Namespace: n.Name,
					Object:    "o1",
					Relation:  "rel",
					SubjectID: x.Ptr("s1"),
				},
				{
					Namespace: n.Name,
					Object:    "o2",
					Relation:  "rel",
					SubjectID: x.Ptr("s2"),
				},
			}
			for i := 0; i < len(rts); i++ {
				c.createTuple(t, rts[i])
			}

			q := &ketoapi.RelationQuery{
				Namespace: &n.Name,
				Relation:  x.Ptr("rel"),
			}
			resp := c.queryTuple(t, q)
			require.ElementsMatch(t, resp.RelationTuples, rts)

			c.deleteAllTuples(t, q)
			resp = c.queryTuple(t, q)
			assert.Equal(t, len(resp.RelationTuples), 0)
		})

		t.Run("case=returns error with status code on unknown namespace", func(t *testing.T) {
			c.queryTupleErr(t, herodot.ErrNotFound, &ketoapi.RelationQuery{Namespace: x.Ptr("unknown namespace")})
		})

		t.Run("case=hides tuples from deleted namespace", func(t *testing.T) {
			n := &namespace.Namespace{Name: t.Name()}
			m.add(t, n)

			c.createTuple(t, &relationtuple.InternalRelationTuple{
				Namespace: n.Name,
				Object:    "o",
				Relation:  "rel",
				Subject:   &relationtuple.SubjectID{ID: "s"},
			})

			m.remove(t, n.ID)

			resp := c.queryTuple(t, &relationtuple.RelationQuery{})
			assert.Equal(t, len(resp.RelationTuples), 0)

			// Add the namespace again here, so that we can clean up properly.
			m.add(t, n)
		})
	}
}
