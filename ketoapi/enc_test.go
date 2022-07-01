package ketoapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ory/keto/internal/x"
	rts "github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/url"
	"testing"
)

func TestRelationTuple(t *testing.T) {
	t.Run("method=string encoding", func(t *testing.T) {
		assert.Equal(t, "n:o#r@s", (&RelationTuple{
			Namespace: "n",
			Object:    "o",
			Relation:  "r",
			SubjectID: x.Ptr("s"),
		}).String())
	})

	t.Run("method=string decoding", func(t *testing.T) {
		for i, tc := range []struct {
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
					SubjectID: x.Ptr("s"),
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
			t.Run(fmt.Sprintf("case=%d", i), func(t *testing.T) {
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
				SubjectID: x.Ptr("s"),
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
				SubjectID: x.Ptr(""),
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
					SubjectID: x.Ptr("user"),
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
						SubjectID: x.Ptr("s"),
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
					Namespace: x.Ptr("n"),
					Object:    x.Ptr("o"),
					Relation:  x.Ptr("r"),
					SubjectID: x.Ptr("foo"),
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
					Namespace: x.Ptr("n"),
					Object:    x.Ptr("o"),
					Relation:  x.Ptr("r"),
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
					Namespace: x.Ptr("n"),
					Relation:  x.Ptr("r"),
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
