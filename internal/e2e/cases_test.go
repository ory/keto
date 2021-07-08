package e2e

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/ory/herodot"

	"github.com/ory/keto/internal/x"

	"github.com/stretchr/testify/assert"

	"github.com/ory/keto/internal/expand"
	"github.com/ory/keto/internal/namespace"
	"github.com/ory/keto/internal/relationtuple"
)

func runCases(c client, addNamespace func(*testing.T, ...*namespace.Namespace)) func(*testing.T) {
	return func(t *testing.T) {
		c.waitUntilLive(t)

		t.Run("case=gets empty namespace", func(t *testing.T) {
			n := &namespace.Namespace{Name: t.Name()}
			addNamespace(t, n)

			resp := c.queryTuple(t, &relationtuple.RelationQuery{Namespace: n.Name})
			assert.Len(t, resp.RelationTuples, 0)
		})

		t.Run("case=creates tuple and uses it then", func(t *testing.T) {
			n := &namespace.Namespace{Name: t.Name()}
			addNamespace(t, n)

			tuple := &relationtuple.InternalRelationTuple{
				Namespace: n.Name,
				Object:    fmt.Sprintf("object for client %T", c),
				Relation:  "access",
				Subject:   &relationtuple.SubjectID{ID: "client"},
			}

			c.createTuple(t, tuple)

			resp := c.queryTuple(t, &relationtuple.RelationQuery{Namespace: tuple.Namespace})
			assert.Contains(t, resp.RelationTuples, tuple)

			// try the check API to see whether the tuple is interpreted correctly
			assert.True(t, c.check(t, tuple))
		})

		t.Run("case=expand API", func(t *testing.T) {
			n := &namespace.Namespace{Name: t.Name()}
			addNamespace(t, n)

			obj := fmt.Sprintf("tree for client %T", c)
			rel := "expand"

			subjects := []string{"s1", "s2"}
			expectedTree := &expand.Tree{
				Type: expand.Union,
				Subject: &relationtuple.SubjectSet{
					Namespace: n.Name,
					Object:    obj,
					Relation:  rel,
				},
				Children: make([]*expand.Tree, len(subjects)),
			}

			for i, subjectID := range subjects {
				c.createTuple(t, &relationtuple.InternalRelationTuple{
					Namespace: n.Name,
					Object:    obj,
					Relation:  rel,
					Subject:   &relationtuple.SubjectID{ID: subjectID},
				})
				expectedTree.Children[i] = &expand.Tree{
					Type:    expand.Leaf,
					Subject: &relationtuple.SubjectID{ID: subjectID},
				}
			}

			actualTree := c.expand(t, expectedTree.Subject.(*relationtuple.SubjectSet), 100)

			assert.Equal(t, expectedTree.Type, actualTree.Type)
			assert.Equal(t, expectedTree.Subject, actualTree.Subject)
			assert.Equal(t, len(expectedTree.Children), len(actualTree.Children), "expected: %+v; actual: %+v", expectedTree.Children, actualTree.Children)

			for _, child := range expectedTree.Children {
				assert.Contains(t, actualTree.Children, child)
			}
		})

		t.Run("case=gets result paginated", func(t *testing.T) {
			const nTuples = 10
			n := &namespace.Namespace{Name: t.Name()}
			addNamespace(t, n)

			rel := fmt.Sprintf("some unique relation %T", c)
			for i := 0; i < nTuples; i++ {
				c.createTuple(t, &relationtuple.InternalRelationTuple{
					Namespace: n.Name,
					Object:    "o" + strconv.Itoa(i),
					Relation:  rel,
					Subject:   &relationtuple.SubjectID{ID: "s" + strconv.Itoa(i)},
				})
			}

			var (
				resp   relationtuple.GetResponse
				nPages int
			)
			// do ... while resp.NextPageToken != ""
			for ok := true; ok; ok = resp.NextPageToken != "" {
				resp = *c.queryTuple(t,
					&relationtuple.RelationQuery{
						Namespace: n.Name,
						Relation:  rel,
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
			addNamespace(t, n)

			for _, s := range []relationtuple.Subject{
				&relationtuple.SubjectID{ID: "s"},
				&relationtuple.SubjectSet{
					Namespace: n.Name,
					Object:    "so",
					Relation:  "rel",
				},
			} {
				t.Run(fmt.Sprintf("subject_type=%T", s), func(t *testing.T) {
					rt := &relationtuple.InternalRelationTuple{
						Namespace: n.Name,
						Object:    "o",
						Relation:  "rel",
						Subject:   s,
					}
					c.createTuple(t, rt)

					resp := c.queryTuple(t, (*relationtuple.RelationQuery)(rt))
					assert.Equal(t, []*relationtuple.InternalRelationTuple{rt}, resp.RelationTuples)

					c.deleteTuple(t, rt)

					resp = c.queryTuple(t, (*relationtuple.RelationQuery)(rt))
					assert.Len(t, resp.RelationTuples, 0)
				})
			}
		})

		t.Run("case=returns error with status code on unknown namespace", func(t *testing.T) {
			c.queryTupleErr(t, herodot.ErrNotFound, &relationtuple.RelationQuery{Namespace: "unknown namespace"})
		})
	}
}
