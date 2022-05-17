package e2e

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ory/keto/internal/namespace"
	"github.com/ory/keto/internal/relationtuple"
)

func runTransactionCases(c transactClient, addNamespace func(*testing.T, ...*namespace.Namespace)) func(*testing.T) {
	return func(t *testing.T) {
		t.Parallel()
		t.Run("case=create and delete", func(t *testing.T) {
			n := &namespace.Namespace{Name: t.Name()}
			addNamespace(t, n)

			tuples := []*relationtuple.InternalRelationTuple{
				{
					Namespace: n.Name,
					Object:    "o",
					Relation:  "rel",
					Subject: &relationtuple.SubjectSet{
						Namespace: n.Name,
						Object:    "o",
						Relation:  "rel",
					},
				},
				{
					Namespace: n.Name,
					Object:    "o",
					Relation:  "rel",
					Subject:   &relationtuple.SubjectID{ID: "sid"},
				},
			}

			c.transactTuples(t, tuples, nil)

			resp := c.queryTuple(t, &relationtuple.RelationQuery{
				Namespace: n.Name,
			})
			for i := range tuples {
				assert.Contains(t, resp.RelationTuples, tuples[i])
			}

			c.transactTuples(t, nil, tuples)

			resp = c.queryTuple(t, &relationtuple.RelationQuery{
				Namespace: n.Name,
			})
			assert.Len(t, resp.RelationTuples, 0)
		})

		t.Run("case=expand-api-display-access docs code sample", func(t *testing.T) {
			files := &namespace.Namespace{Name: t.Name() + "files"}
			directories := &namespace.Namespace{Name: t.Name() + "directories"}
			addNamespace(t, files, directories)

			tuples := []*relationtuple.InternalRelationTuple{
				{
					Namespace: directories.Name,
					Object:    "/photos",
					Relation:  "owner",
					Subject: &relationtuple.SubjectID{
						ID: "maureen",
					},
				},
				{
					Namespace: files.Name,
					Object:    "/photos/beach.jpg",
					Relation:  "owner",
					Subject: &relationtuple.SubjectID{
						ID: "maureen",
					},
				},
				{
					Namespace: files.Name,
					Object:    "/photos/mountains.jpg",
					Relation:  "owner",
					Subject: &relationtuple.SubjectID{
						ID: "laura",
					},
				},
				{
					Namespace: directories.Name,
					Object:    "/photos",
					Relation:  "access",
					Subject: &relationtuple.SubjectID{
						ID: "laura",
					},
				},
			}
			for _, o := range []struct{ n, o string }{
				{files.Name, "/photos/beach.jpg"},
				{files.Name, "/photos/mountains.jpg"},
				{directories.Name, "/photos"},
			} {
				tuples = append(tuples, &relationtuple.InternalRelationTuple{
					Namespace: o.n,
					Object:    o.o,
					Relation:  "access",
					Subject: &relationtuple.SubjectSet{
						Namespace: o.n,
						Object:    o.o,
						Relation:  "owner",
					},
				})
			}
			for _, obj := range []string{"/photos/beach.jpg", "/photos/mountains.jpg"} {
				tuples = append(tuples, &relationtuple.InternalRelationTuple{
					Namespace: files.Name,
					Object:    obj,
					Relation:  "access",
					Subject: &relationtuple.SubjectSet{
						Namespace: directories.Name,
						Object:    "/photos",
						Relation:  "access",
					},
				})
			}

			c.transactTuples(t, tuples, nil)

			resp := c.queryTuple(t, &relationtuple.RelationQuery{})
			assert.Equal(t, len(tuples), len(resp.RelationTuples))
			for i := range tuples {
				assert.Contains(t, resp.RelationTuples, tuples[i])
			}

			c.transactTuples(t, nil, tuples)

			resp = c.queryTuple(t, &relationtuple.RelationQuery{})
			assert.Len(t, resp.RelationTuples, 0)
		})
	}
}
