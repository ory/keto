/*
 * Copyright © 2016-2018 Aeneas Rekkas <aeneas+oss@aeneas.io>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * @author		Aeneas Rekkas <aeneas+oss@aeneas.io>
 * @copyright 	2015-2018 Aeneas Rekkas <aeneas+oss@aeneas.io>
 * @license 	Apache-2.0
 */

package ladon

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/pborman/uuid"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var TestManagerPolicies = []*DefaultPolicy{
	{
		ID:          uuid.New(),
		Description: "description",
		Subjects:    []string{"user", "anonymous"},
		Effect:      AllowAccess,
		Resources:   []string{"article", "user"},
		Actions:     []string{"create", "update"},
		Conditions:  Conditions{},
	},
	{
		ID:          uuid.New(),
		Description: "description",
		Subjects:    []string{},
		Effect:      AllowAccess,
		Resources:   []string{"<article|user>"},
		Actions:     []string{"view"},
		Conditions:  Conditions{},
	},
	{
		ID:          uuid.New(),
		Description: "description",
		Subjects:    []string{},
		Effect:      AllowAccess,
		Resources:   []string{},
		Actions:     []string{"view"},
		Conditions:  Conditions{},
	},
	{
		ID:          uuid.New(),
		Description: "description",
		Subjects:    []string{},
		Effect:      AllowAccess,
		Resources:   []string{},
		Actions:     []string{},
		Conditions:  Conditions{},
	},
	{
		ID:          uuid.New(),
		Description: "description",
		Subjects:    []string{},
		Effect:      AllowAccess,
		Resources:   []string{"foo"},
		Actions:     []string{},
		Conditions:  Conditions{},
	},
	{
		ID:          uuid.New(),
		Description: "description",
		Subjects:    []string{"foo"},
		Effect:      AllowAccess,
		Resources:   []string{"foo"},
		Actions:     []string{},
		Conditions:  Conditions{},
	},
	{
		ID:          uuid.New(),
		Description: "description",
		Subjects:    []string{"foo"},
		Effect:      AllowAccess,
		Resources:   []string{},
		Actions:     []string{},
		Conditions:  Conditions{},
	},
	{
		ID:          uuid.New(),
		Description: "description",
		Effect:      AllowAccess,
		Conditions:  Conditions{},
	},
	{
		ID:          uuid.New(),
		Description: "description",
		Subjects:    []string{"<peter|max>"},
		Effect:      DenyAccess,
		Resources:   []string{"article", "user"},
		Actions:     []string{"view"},
		Conditions: Conditions{
			"owner": &EqualsSubjectCondition{},
		},
	},
	{
		ID:          uuid.New(),
		Description: "description",
		Subjects:    []string{"<user|max|anonymous>", "peter"},
		Effect:      DenyAccess,
		Resources:   []string{".*"},
		Actions:     []string{"disable"},
		Conditions: Conditions{
			"ip": &CIDRCondition{
				CIDR: "1234",
			},
			"owner": &EqualsSubjectCondition{},
		},
	},
	{
		ID:          uuid.New(),
		Description: "description",
		Subjects:    []string{"<.*>"},
		Effect:      AllowAccess,
		Resources:   []string{"<article|user>"},
		Actions:     []string{"view"},
		Conditions: Conditions{
			"ip": &CIDRCondition{
				CIDR: "1234",
			},
			"owner": &EqualsSubjectCondition{},
		},
	},
	{
		ID:          uuid.New(),
		Description: "description",
		Subjects:    []string{"<us[er]+>"},
		Effect:      AllowAccess,
		Resources:   []string{"<article|user>"},
		Actions:     []string{"view"},
		Conditions: Conditions{
			"ip": &CIDRCondition{
				CIDR: "1234",
			},
			"owner": &EqualsSubjectCondition{},
		},
	},
	//Two new policies which do not persist in MySQL correctly
	{
		ID:          uuid.New(),
		Description: "A failed policy",
		Subjects:    []string{"supplier"},
		Effect:      AllowAccess,
		Resources:   []string{"product:<.*>"},
		Actions:     []string{"update"},
		Conditions:  Conditions{},
	},
	{
		ID:          uuid.New(),
		Description: "Another failed policy",
		Subjects:    []string{"buyer"},
		Effect:      AllowAccess,
		Resources:   []string{"products:attributeGroup:<.*>"},
		Actions:     []string{"create"},
		Conditions:  Conditions{},
	},
}

var testPolicies = []*DefaultPolicy{
	{
		ID:          uuid.New(),
		Description: "description",
		Subjects:    []string{"sql<.*>match"},
		Effect:      AllowAccess,
		Resources:   []string{"master", "user", "article"},
		Actions:     []string{"create", "update", "delete"},
		Conditions: Conditions{
			"foo": &StringEqualCondition{
				Equals: "foo",
			},
		},
	},
	{
		ID:          uuid.New(),
		Description: "description",
		Subjects:    []string{"sqlmatch"},
		Effect:      AllowAccess,
		Resources:   []string{"master", "user", "article"},
		Actions:     []string{"create", "update", "delete"},
		Conditions: Conditions{
			"foo": &StringEqualCondition{
				Equals: "foo",
			},
		},
	},
	{
		ID:          uuid.New(),
		Description: "description",
		Subjects:    []string{},
		Effect:      AllowAccess,
		Resources:   []string{"master", "user", "article"},
		Actions:     []string{"create", "update", "delete"},
		Conditions: Conditions{
			"foo": &StringEqualCondition{
				Equals: "foo",
			},
		},
	},
	{
		ID:          uuid.New(),
		Description: "description",
		Effect:      AllowAccess,
		Resources:   []string{"master", "user", "article"},
		Actions:     []string{"create", "update", "delete"},
		Conditions: Conditions{
			"foo": &StringEqualCondition{
				Equals: "foo",
			},
		},
	},
}

func TestHelperFindPoliciesForSubject(k string, s Manager) func(t *testing.T) {
	return func(t *testing.T) {
		for _, c := range testPolicies {
			t.Run(fmt.Sprintf("create=%s", k), func(t *testing.T) {
				require.NoError(t, s.Create(c))
			})
		}

		res, err := s.FindRequestCandidates(&Request{
			Subject:  "sqlmatch",
			Resource: "article",
			Action:   "create",
		})
		require.NoError(t, err)
		require.Len(t, res, 2)

		if testPolicies[0].ID == res[0].GetID() {
			AssertPolicyEqual(t, testPolicies[0], res[0])
			AssertPolicyEqual(t, testPolicies[1], res[1])
		} else {
			AssertPolicyEqual(t, testPolicies[0], res[1])
			AssertPolicyEqual(t, testPolicies[1], res[0])
		}

		res, err = s.FindRequestCandidates(&Request{
			Subject:  "sqlamatch",
			Resource: "article",
			Action:   "create",
		})

		require.NoError(t, err)
		require.Len(t, res, 1)
		AssertPolicyEqual(t, testPolicies[0], res[0])
	}
}

func AssertPolicyEqual(t *testing.T, expected, got Policy) {
	assert.Equal(t, expected.GetID(), got.GetID())
	assert.Equal(t, expected.GetDescription(), got.GetDescription())
	assert.Equal(t, expected.GetEffect(), got.GetEffect())

	// This won't work in the memory manager
	//assert.NotNil(t, got.GetActions())
	//assert.NotNil(t, got.GetResources())
	//assert.NotNil(t, got.GetSubjects())

	assert.NoError(t, testEq(expected.GetActions(), got.GetActions()))
	assert.NoError(t, testEq(expected.GetResources(), got.GetResources()))
	assert.NoError(t, testEq(expected.GetSubjects(), got.GetSubjects()))
	assert.EqualValues(t, expected.GetConditions(), got.GetConditions())
}

func testEq(a, b []string) error {
	// We don't care about nil types
	//if a == nil && b == nil {
	//	return true
	//}
	//
	//if a == nil || b == nil {
	//	return false
	//}

	if len(a) != len(b) {
		return errors.Errorf("Length not equal: %v (%d) != %v (%d)", a, len(a), b, len(b))
	}

	var found bool
	for i := range a {
		found = false

		for y := range b {
			if a[i] == b[y] {
				found = true
				break
			}
		}

		if !found {
			return errors.Errorf("No match found: %s from %v in %v", i, a, b)
		}
	}

	return nil
}

func TestHelperGetErrors(s Manager) func(t *testing.T) {
	return func(t *testing.T) {
		_, err := s.Get(uuid.New())
		assert.Error(t, err)

		_, err = s.Get("asdf")
		assert.Error(t, err)
	}
}

func TestHelperCreateGetDelete(s Manager) func(t *testing.T) {
	return func(t *testing.T) {
		for i, c := range TestManagerPolicies {
			t.Run(fmt.Sprintf("case=%d/id=%s/type=create", i, c.GetID()), func(t *testing.T) {
				_, err := s.Get(c.GetID())
				require.Error(t, err)
				require.NoError(t, s.Create(c))
			})

			t.Run(fmt.Sprintf("case=%d/id=%s/type=query", i, c.GetID()), func(t *testing.T) {
				get, err := s.Get(c.GetID())
				require.NoError(t, err)

				AssertPolicyEqual(t, c, get)
			})

			t.Run(fmt.Sprintf("case=%d/id=%s/type=update", i, c.GetID()), func(t *testing.T) {
				c.Description = c.Description + "_updated"
				require.NoError(t, s.Update(c))

				get, err := s.Get(c.GetID())
				require.NoError(t, err)

				AssertPolicyEqual(t, c, get)
			})

			t.Run(fmt.Sprintf("case=%d/id=%s/type=query", i, c.GetID()), func(t *testing.T) {
				get, err := s.Get(c.GetID())
				require.NoError(t, err)

				AssertPolicyEqual(t, c, get)
			})
		}

		t.Run("type=query-all", func(t *testing.T) {
			count := int64(len(TestManagerPolicies))

			pols, err := s.GetAll(100, 0)
			require.NoError(t, err)
			assert.Len(t, pols, len(TestManagerPolicies))

			pols4, err := s.GetAll(1, 0)
			require.NoError(t, err)
			assert.Len(t, pols4, 1)

			pols2, err := s.GetAll(100, count-1)
			require.NoError(t, err)
			assert.Len(t, pols2, 1)

			pols3, err := s.GetAll(100, count)
			require.NoError(t, err)
			assert.Len(t, pols3, 0)

			found := map[string]int{}
			for _, got := range pols {
				for _, expect := range TestManagerPolicies {
					if got.GetID() == expect.GetID() {
						assert.ObjectsAreEqualValues(t, reflect.DeepEqual(expect, got))
						found[got.GetID()]++
					}
				}
			}
			// for _, got := range pols {
			// 	for _, expect := range TestManagerPolicies {
			// 		//This is a modified equality check
			// 		if got.GetID() == expect.GetID() && reflect.DeepEqual(got.GetResources(), expect.GetResources()) && reflect.DeepEqual(got.GetActions(), expect.GetActions()) {
			// 			found[got.GetID()]++
			// 		}
			// 	}
			// }
			assert.Len(t, found, len(TestManagerPolicies))

			for _, f := range found {
				//This assert is supposed to pass
				assert.Equal(t, 1, f)
			}
		})

		for i, c := range TestManagerPolicies {
			t.Run(fmt.Sprintf("case=%d/id=%s/type=delete", i, c.GetID()), func(t *testing.T) {
				assert.NoError(t, s.Delete(c.ID))

				_, err := s.Get(c.GetID())
				assert.Error(t, err)
			})
		}
	}
}
