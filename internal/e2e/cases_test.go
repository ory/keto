package e2e

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/ory/keto/internal/x"

	"github.com/stretchr/testify/assert"

	"github.com/ory/keto/internal/expand"
	"github.com/ory/keto/internal/namespace"
	"github.com/ory/keto/internal/relationtuple"
)

func runCases(c client, nspaces []*namespace.Namespace) func(*testing.T) {
	return func(t *testing.T) {
		c.waitUntilLive(t)

		t.Run("case=creates tuple and uses it then", func(t *testing.T) {
			tuple := &relationtuple.InternalRelationTuple{
				Namespace: nspaces[0].Name,
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
			obj := fmt.Sprintf("tree for client %T", c)
			rel := "expand"

			subjects := []string{"s1", "s2"}
			expectedTree := &expand.Tree{
				Type: expand.Union,
				Subject: &relationtuple.SubjectSet{
					Namespace: nspaces[0].Name,
					Object:    obj,
					Relation:  rel,
				},
				Children: make([]*expand.Tree, len(subjects)),
			}

			for i, subjectID := range subjects {
				c.createTuple(t, &relationtuple.InternalRelationTuple{
					Namespace: nspaces[0].Name,
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

			rel := fmt.Sprintf("some unique relation %T", c)
			for i := 0; i < nTuples; i++ {
				c.createTuple(t, &relationtuple.InternalRelationTuple{
					Namespace: nspaces[0].Name,
					Object:    "o" + strconv.Itoa(i),
					Relation:  rel,
					Subject:   &relationtuple.SubjectID{ID: "s" + strconv.Itoa(i)},
				})
			}

			var (
				resp   relationtuple.GetResponse
				nPages int
			)
			for ; !resp.IsLastPage; nPages++ {
				resp = *c.queryTuple(t,
					&relationtuple.RelationQuery{
						Namespace: nspaces[0].Name,
						Relation:  rel,
					},
					x.WithToken(resp.NextPageToken),
					x.WithSize(1),
				)
				assert.Len(t, resp.RelationTuples, 1)
			}

			assert.Equal(t, nTuples, nPages)
		})
	}
}
