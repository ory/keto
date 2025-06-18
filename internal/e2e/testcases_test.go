// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package e2e

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/ory/herodot"
	"github.com/ory/x/pointerx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ory/keto/internal/expand"
	"github.com/ory/keto/internal/namespace"
	"github.com/ory/keto/ketoapi"
)

func runCases(c client, m *namespaceTestManager) func(*testing.T) {
	return func(t *testing.T) {
		c.waitUntilLive(t)

		t.Run("case=list namespaces", func(t *testing.T) {
			first := namespace.Namespace{Name: "my namespace"}
			second := namespace.Namespace{Name: "my other namespace"}
			m.add(t, &first, &second)

			resp := c.queryNamespaces(t)
			assert.GreaterOrEqual(t, len(resp.Namespaces), 2)
			assertNamespacesContains(t, resp.Namespaces, "my namespace")
			assertNamespacesContains(t, resp.Namespaces, "my other namespace")
		})

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
				SubjectID: pointerx.Ptr("client"),
			}

			c.createTuple(t, tuple)

			resp := c.queryTuple(t, &ketoapi.RelationQuery{Namespace: &tuple.Namespace})
			require.Len(t, resp.RelationTuples, 1)
			assert.Equal(t, tuple, resp.RelationTuples[0])

			// try the check API to see whether the tuple is interpreted correctly
			assert.True(t, c.check(t, tuple))
			batchResult := c.batchCheck(t, []*ketoapi.RelationTuple{tuple})
			require.Len(t, batchResult, 1)
			assert.True(t, batchResult[0].allowed)
			assert.Empty(t, batchResult[0].errorMessage)
		})

		t.Run("case=creates tuple with empty IDs", func(t *testing.T) {
			n := &namespace.Namespace{Name: t.Name()}
			m.add(t, n)

			tuples := []*ketoapi.RelationTuple{{
				Namespace: n.Name,
				Object:    "",
				Relation:  "access",
				SubjectID: pointerx.Ptr(""),
			}, {
				Namespace: n.Name,
				Object:    "",
				Relation:  "access",
				SubjectSet: &ketoapi.SubjectSet{
					Namespace: n.Name,
					Object:    "",
					Relation:  "access",
				},
			}}

			for _, tp := range tuples {
				c.createTuple(t, tp)
				// try the check API to see whether the tuple is interpreted correctly
				assert.True(t, c.check(t, tp))
				batchResult := c.batchCheck(t, []*ketoapi.RelationTuple{tp})
				require.Len(t, batchResult, 1)
				assert.True(t, batchResult[0].allowed)
				assert.Empty(t, batchResult[0].errorMessage)
			}

			resp := c.queryTuple(t, &ketoapi.RelationQuery{Namespace: &n.Name})
			assert.ElementsMatch(t, tuples, resp.RelationTuples)
		})

		t.Run("case=check subjectSet relations", func(t *testing.T) {
			n := &namespace.Namespace{Name: t.Name()}
			m.add(t, n)

			obj := fmt.Sprintf("obj for client %T", c)
			rel := "check"

			rt := &ketoapi.RelationTuple{
				Namespace: n.Name,
				Object:    obj,
				Relation:  rel,
				SubjectSet: &ketoapi.SubjectSet{
					Namespace: n.Name,
					Object:    obj,
					Relation:  rel,
				},
			}
			c.createTuple(t, rt)

			assert.True(t, c.check(t, rt))
			batchResult := c.batchCheck(t, []*ketoapi.RelationTuple{rt})
			require.Len(t, batchResult, 1)
			assert.True(t, batchResult[0].allowed)
			assert.Empty(t, batchResult[0].errorMessage)
		})

		t.Run("case=batch check", func(t *testing.T) {
			/*
			 Test batch check with four cases:
			   - Allowed single subject tuple
			   - Allowed subject set tuple
			   - Tuple with unknown namespace (will return an error for this case)
			   - Dis-allowed tuple
			*/
			namespace1 := &namespace.Namespace{Name: t.Name()}
			namespace2 := &namespace.Namespace{Name: t.Name() + "-2"}
			m.add(t, namespace1)
			m.add(t, namespace2)

			obj1 := fmt.Sprintf("obj for client %T", c)
			obj2 := fmt.Sprintf("another obj for client %T", c)
			rel1 := "check"
			rel2 := "access"

			tupleSubjectSet := &ketoapi.RelationTuple{
				Namespace: namespace1.Name,
				Object:    obj1,
				Relation:  rel1,
				SubjectSet: &ketoapi.SubjectSet{
					Namespace: namespace1.Name,
					Object:    obj1,
					Relation:  rel1,
				},
			}
			c.createTuple(t, tupleSubjectSet)

			tupleSingleSubject := &ketoapi.RelationTuple{
				Namespace: namespace2.Name,
				Object:    obj2,
				Relation:  rel2,
				SubjectID: pointerx.Ptr("sub1"),
			}
			c.createTuple(t, tupleSingleSubject)

			unknownNamespaceTuple := &ketoapi.RelationTuple{
				Namespace: "unknown-namespace",
				Object:    obj1,
				Relation:  rel1,
				SubjectID: pointerx.Ptr("sub1"),
			}

			unknownSubjectTuple := &ketoapi.RelationTuple{
				Namespace: namespace1.Name,
				Object:    obj1,
				Relation:  rel1,
				SubjectID: pointerx.Ptr("unknown-sub"),
			}

			batchResult := c.batchCheck(t, []*ketoapi.RelationTuple{tupleSubjectSet, tupleSingleSubject,
				unknownNamespaceTuple, unknownSubjectTuple})
			require.Len(t, batchResult, 4)
			assert.True(t, batchResult[0].allowed)
			assert.Empty(t, batchResult[0].errorMessage)
			assert.True(t, batchResult[1].allowed)
			assert.Empty(t, batchResult[1].errorMessage)
			assert.False(t, batchResult[2].allowed)
			assert.Contains(t, batchResult[2].errorMessage, "The requested resource could not be found")
			assert.False(t, batchResult[3].allowed)
			assert.Empty(t, batchResult[3].errorMessage)

			// Verify a call with no tuples returns successfully with no results
			emptyResults := c.batchCheck(t, []*ketoapi.RelationTuple{})
			require.Empty(t, emptyResults)
		})

		t.Run("case=batch check validation errors", func(t *testing.T) {
			// Pass in 11 tuples to check, more than the default limit of 10, and verify the request is rejected
			tuples := make([]*ketoapi.RelationTuple, 11)
			for i := range tuples {
				tuples[i] = &ketoapi.RelationTuple{
					Namespace: "namespace-name",
					Object:    "obj",
					Relation:  "rel",
					SubjectID: pointerx.Ptr("sub"),
				}
			}
			c.batchCheckErr(t, tuples, herodot.ErrBadRequest)
		})

		t.Run("case=expand API", func(t *testing.T) {
			n := &namespace.Namespace{Name: t.Name()}
			m.add(t, n)

			obj := fmt.Sprintf("tree for client %T", c)
			rel := "expand"

			subjects := []string{"s1", "s2"}
			expectedTree := &ketoapi.Tree[*ketoapi.RelationTuple]{
				Type: ketoapi.TreeNodeUnion,
				Tuple: &ketoapi.RelationTuple{
					SubjectSet: &ketoapi.SubjectSet{
						Namespace: n.Name,
						Object:    obj,
						Relation:  rel,
					},
				},
				Children: make([]*ketoapi.Tree[*ketoapi.RelationTuple], len(subjects)),
			}

			for i, subjectID := range subjects {
				subjectID := subjectID
				c.createTuple(t, &ketoapi.RelationTuple{
					Namespace: n.Name,
					Object:    obj,
					Relation:  rel,
					SubjectID: &subjectID,
				})
				expectedTree.Children[i] = &ketoapi.Tree[*ketoapi.RelationTuple]{
					Type: ketoapi.TreeNodeLeaf,
					Tuple: &ketoapi.RelationTuple{
						SubjectID: &subjectID,
					},
				}
			}

			actualTree := c.expand(t, expectedTree.Tuple.SubjectSet, 100)

			assert.Equal(t, expectedTree.Type, actualTree.Type)
			assert.Equalf(t, expectedTree.Tuple, actualTree.Tuple,
				"want:\t%s\ngot:\t%s", expectedTree.Tuple, actualTree.Tuple)
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
					SubjectID: pointerx.Ptr("s" + strconv.Itoa(i)),
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
					withToken(resp.NextPageToken),
					withSize(1),
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
					SubjectID: pointerx.Ptr("s"),
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
					SubjectID: pointerx.Ptr("s1"),
				},
				{
					Namespace: n.Name,
					Object:    "o2",
					Relation:  "rel",
					SubjectID: pointerx.Ptr("s2"),
				},
			}
			for i := 0; i < len(rts); i++ {
				c.createTuple(t, rts[i])
			}

			q := &ketoapi.RelationQuery{
				Namespace: &n.Name,
				Relation:  pointerx.Ptr("rel"),
			}
			resp := c.queryTuple(t, q)
			require.ElementsMatch(t, resp.RelationTuples, rts)

			c.deleteAllTuples(t, q)
			resp = c.queryTuple(t, q)
			assert.Equal(t, 0, len(resp.RelationTuples))
		})

		t.Run("case=returns error with status code on unknown namespace", func(t *testing.T) {
			c.queryTupleErr(t, herodot.ErrNotFound, &ketoapi.RelationQuery{Namespace: pointerx.Ptr("unknown namespace")})
		})

		t.Run("case=still serves tuples from deleted namespace", func(t *testing.T) {
			n := &namespace.Namespace{Name: t.Name()}
			m.add(t, n)

			tuple := &ketoapi.RelationTuple{
				Namespace: n.Name,
				Object:    "o",
				Relation:  "rel",
				SubjectID: pointerx.Ptr("s"),
			}
			c.createTuple(t, tuple)

			m.remove(t, n.Name)

			resp := c.queryTuple(t, &ketoapi.RelationQuery{})
			assert.Equal(t, []*ketoapi.RelationTuple{tuple}, resp.RelationTuples)

			// Add the namespace again here, so that we can clean up properly.
			m.add(t, n)
		})

		t.Run("case=OPL syntax check", func(t *testing.T) {
			parseErrors := c.oplCheckSyntax(t, []byte("/* unclosed comment"))
			require.Len(t, parseErrors, 1)
			assert.Contains(t, parseErrors[0].Message, "unclosed comment")
		})
	}
}

func assertNamespacesContains(t *testing.T, list []ketoapi.Namespace, name string) {
	t.Helper()
	for _, n := range list {
		if n.Name == name {
			return
		}
	}
	t.Errorf("Could not find %q in %+v", name, list)
}
