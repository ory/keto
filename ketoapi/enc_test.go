// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package ketoapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"testing"

	"github.com/ory/x/pointerx"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	rts "github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2"
)

func TestRelationTuple(t *testing.T) {
	t.Run("method=string encoding", func(t *testing.T) {
		for _, tc := range []struct {
			tuple    *RelationTuple
			expected string
		}{
			{ // full tuple
				tuple: &RelationTuple{
					Namespace: "n",
					Object:    "o",
					Relation:  "r",
					SubjectID: pointerx.Ptr("s"),
				},
				expected: "n:o#r@s",
			},
			{ // skip '#' in subject set when relation is empty
				tuple: &RelationTuple{
					Namespace: "groups",
					Object:    "dev",
					Relation:  "member",
					SubjectSet: &SubjectSet{
						Namespace: "users",
						Object:    "user",
					},
				},
				expected: "groups:dev#member@users:user",
			},
		} {
			t.Run("case="+tc.expected, func(t *testing.T) {
				assert.Equal(t, tc.expected, tc.tuple.String())
			})
		}
	})

	t.Run("method=string decoding", func(t *testing.T) {
		for _, tc := range []struct {
			enc      string
			err      error
			expected *RelationTuple
		}{
			{
				enc: "n:o#r@s",
				expected: &RelationTuple{
					Namespace: "n",
					Object:    "o",
					Relation:  "r",
					SubjectID: pointerx.Ptr("s"),
				},
			},
			{
				enc: "n:o#r@n:o#r",
				expected: &RelationTuple{
					Namespace: "n",
					Object:    "o",
					Relation:  "r",
					SubjectSet: &SubjectSet{
						Namespace: "n",
						Object:    "o",
						Relation:  "r",
					},
				},
			},
			{
				enc: "n:o#r@n:o#...",
				expected: &RelationTuple{
					Namespace: "n",
					Object:    "o",
					Relation:  "r",
					SubjectSet: &SubjectSet{
						Namespace: "n",
						Object:    "o",
						Relation:  "...",
					},
				},
			},
			{
				enc: "n:o#r@n:o#",
				expected: &RelationTuple{
					Namespace: "n",
					Object:    "o",
					Relation:  "r",
					SubjectSet: &SubjectSet{
						Namespace: "n",
						Object:    "o",
					},
				},
			},
			{
				enc: "n:o#r@n:o",
				expected: &RelationTuple{
					Namespace: "n",
					Object:    "o",
					Relation:  "r",
					SubjectSet: &SubjectSet{
						Namespace: "n",
						Object:    "o",
					},
				},
			},
			{
				enc: "n:o#r@(n:o#r)",
				expected: &RelationTuple{
					Namespace: "n",
					Object:    "o",
					Relation:  "r",
					SubjectSet: &SubjectSet{
						Namespace: "n",
						Object:    "o",
						Relation:  "r",
					},
				},
			},
			{
				enc: "#dev:@ory#:working:@projects:keto#awesome",
				expected: &RelationTuple{
					Namespace: "#dev",
					Object:    "@ory",
					Relation:  ":working:",
					SubjectSet: &SubjectSet{
						Namespace: "projects",
						Object:    "keto",
						Relation:  "awesome",
					},
				},
			},
			{
				enc: "no-colon#in@this",
				err: ErrMalformedInput,
			},
			{
				enc: "no:hash-in@this",
				err: ErrMalformedInput,
			},
			{
				enc: "no:at#in-this",
				err: ErrMalformedInput,
			},
		} {
			t.Run(fmt.Sprintf("string=%s", tc.enc), func(t *testing.T) {
				actual, err := (&RelationTuple{}).FromString(tc.enc)
				assert.True(t, errors.Is(err, tc.err), "%+v", err)
				assert.Equal(t, tc.expected, actual)
			})
		}
	})

	t.Run("case=url encoding-decoding", func(t *testing.T) {
		for i, r := range []*RelationTuple{
			{
				Namespace: "n",
				Object:    "o",
				Relation:  "r",
				SubjectID: pointerx.Ptr("s"),
			},
			{
				Namespace: "n",
				Object:    "o",
				Relation:  "r",
				SubjectSet: &SubjectSet{
					Namespace: "sn",
					Object:    "so",
					Relation:  "sr",
				},
			},
			{
				SubjectID: pointerx.Ptr(""),
			},
		} {
			t.Run(fmt.Sprintf("case=%d", i), func(t *testing.T) {
				vals := r.ToURLQuery()
				res, err := (&RelationTuple{}).FromURLQuery(vals)
				require.NoError(t, err, "raw: %+v, enc: %+v", r, vals)
				assert.Equal(t, r, res)
			})
		}
	})

	t.Run("case=url decoding-encoding", func(t *testing.T) {
		for i, v := range []url.Values{
			{
				"namespace":  []string{"n"},
				"object":     []string{"o"},
				"relation":   []string{"r"},
				"subject_id": []string{"foo"},
			},
			{
				"namespace":             []string{"n"},
				"object":                []string{"o"},
				"relation":              []string{"r"},
				"subject_set.namespace": []string{"sn"},
				"subject_set.object":    []string{"so"},
				"subject_set.relation":  []string{"sr"},
			},
		} {
			t.Run(fmt.Sprintf("case=%d", i), func(t *testing.T) {
				rt, err := (&RelationTuple{}).FromURLQuery(v)
				require.NoError(t, err)
				q := rt.ToURLQuery()
				assert.Equal(t, v, q)
			})
		}
	})

	t.Run("case=proto decoding", func(t *testing.T) {
		for i, tc := range []struct {
			proto    TupleData
			expected *RelationTuple
			err      error
		}{
			{
				proto: &rts.RelationTuple{
					Namespace: "n",
					Object:    "o",
					Relation:  "r",
					Subject:   nil,
				},
				err: ErrNilSubject,
			},
			{
				proto: &rts.RelationTuple{
					Namespace: "n",
					Object:    "o",
					Relation:  "r",
					Subject: &rts.Subject{
						Ref: &rts.Subject_Set{
							Set: &rts.SubjectSet{
								Namespace: "n",
								Object:    "o",
								Relation:  "r",
							},
						},
					},
				},
				expected: &RelationTuple{
					Namespace: "n",
					Object:    "o",
					Relation:  "r",
					SubjectSet: &SubjectSet{
						Namespace: "n",
						Object:    "o",
						Relation:  "r",
					},
				},
			},
			{
				proto: &rts.RelationTuple{
					Namespace: "n",
					Object:    "o",
					Relation:  "r",
					Subject: &rts.Subject{
						Ref: &rts.Subject_Id{
							Id: "user",
						},
					},
				},
				expected: &RelationTuple{
					Namespace: "n",
					Object:    "o",
					Relation:  "r",
					SubjectID: pointerx.Ptr("user"),
				},
			},
		} {
			t.Run(fmt.Sprintf("case=%d", i), func(t *testing.T) {
				actual, err := (&RelationTuple{}).FromDataProvider(tc.proto)
				require.ErrorIs(t, err, tc.err)
				assert.Equal(t, tc.expected, actual)
			})
		}
	})

	t.Run("format=JSON", func(t *testing.T) {
		t.Run("direction=encoding-decoding", func(t *testing.T) {
			for _, tc := range []struct {
				name     string
				rt       *RelationTuple
				expected string
			}{
				{
					name: "with subject ID",
					rt: &RelationTuple{
						Namespace: "n",
						Object:    "o",
						Relation:  "r",
						SubjectID: pointerx.Ptr("s"),
					},
					expected: `
{
	"namespace": "n",
	"object": "o",
	"relation": "r",
	"subject_id": "s"
}`,
				},
				{
					name: "with subject set",
					rt: &RelationTuple{
						Namespace: "n",
						Object:    "o",
						Relation:  "r",
						SubjectSet: &SubjectSet{
							Namespace: "sn",
							Object:    "so",
							Relation:  "sr",
						},
					},
					expected: `
{
	"namespace": "n",
	"object": "o",
	"relation": "r",
	"subject_set": {
		"namespace": "sn",
		"object": "so",
		"relation": "sr"
	}
}`,
				},
			} {
				t.Run("case="+tc.name, func(t *testing.T) {
					raw, err := json.Marshal(tc.rt)
					require.NoError(t, err)
					assert.JSONEq(t, tc.expected, string(raw))

					var dec RelationTuple
					require.NoError(t, json.Unmarshal(raw, &dec))
					assert.Equal(t, tc.rt, &dec)
				})
			}
		})
	})
}

func TestRelationQuery(t *testing.T) {
	t.Run("case=url encoding-decoding-encoding", func(t *testing.T) {
		for i, tc := range []struct {
			v url.Values
			r *RelationQuery
		}{
			{
				v: url.Values{
					"namespace":  []string{"n"},
					"object":     []string{"o"},
					"relation":   []string{"r"},
					"subject_id": []string{"foo"},
				},
				r: &RelationQuery{
					Namespace: pointerx.Ptr("n"),
					Object:    pointerx.Ptr("o"),
					Relation:  pointerx.Ptr("r"),
					SubjectID: pointerx.Ptr("foo"),
				},
			},
			{
				v: url.Values{
					"namespace":             []string{"n"},
					"object":                []string{"o"},
					"relation":              []string{"r"},
					"subject_set.namespace": []string{"sn"},
					"subject_set.object":    []string{"so"},
					"subject_set.relation":  []string{"sr"},
				},
				r: &RelationQuery{
					Namespace: pointerx.Ptr("n"),
					Object:    pointerx.Ptr("o"),
					Relation:  pointerx.Ptr("r"),
					SubjectSet: &SubjectSet{
						Namespace: "sn",
						Object:    "so",
						Relation:  "sr",
					},
				},
			},
			{
				v: url.Values{
					"namespace": []string{"n"},
					"relation":  []string{"r"},
				},
				r: &RelationQuery{
					Namespace: pointerx.Ptr("n"),
					Relation:  pointerx.Ptr("r"),
				},
			},
		} {
			t.Run(fmt.Sprintf("case=%d", i), func(t *testing.T) {
				enc := tc.r.ToURLQuery()
				assert.Equal(t, tc.v, enc)

				dec, err := (&RelationQuery{}).FromURLQuery(tc.v)
				require.NoError(t, err)
				assert.Equal(t, tc.r, dec)
			})
		}
	})
}
