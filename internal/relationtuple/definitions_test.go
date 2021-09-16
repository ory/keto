package relationtuple

import (
	"encoding/json"
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"testing"

	"github.com/pkg/errors"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	acl "github.com/ory/keto/proto/ory/keto/acl/v1alpha1"
)

func TestSubject(t *testing.T) {
	t.Run("case=string encoding-decoding", func(t *testing.T) {
		for i, sub := range []Subject{
			&SubjectSet{
				Namespace: "n",
				Object:    "o",
				Relation:  "r",
			},
			&SubjectSet{},
			&SubjectID{
				ID: "id",
			},
			&SubjectID{},
		} {
			t.Run(fmt.Sprintf("case=%d/type=%T", i, sub), func(t *testing.T) {
				enc := sub.String()
				dec, err := SubjectFromString(enc)
				require.NoError(t, err)
				assert.Equal(t, sub, dec)
			})
		}
	})

	t.Run("case=string decoding-encoding", func(t *testing.T) {
		for i, tc := range []struct {
			sub          string
			expectedType Subject
		}{
			{
				sub:          "",
				expectedType: &SubjectID{},
			},
			{
				sub:          "foobar",
				expectedType: &SubjectID{},
			},
			{
				sub:          "foo:bar#baz",
				expectedType: &SubjectSet{},
			},
			{
				sub:          ":#",
				expectedType: &SubjectSet{},
			},
		} {
			t.Run(fmt.Sprintf("case=%d", i), func(t *testing.T) {
				dec, err := SubjectFromString(tc.sub)
				require.NoError(t, err)
				assert.Equal(t, reflect.TypeOf(tc.expectedType), reflect.TypeOf(dec))
				assert.Equal(t, tc.sub, dec.String())
			})
		}
	})

	t.Run("case=proto decoding", func(t *testing.T) {
		for i, tc := range []struct {
			proto    *acl.Subject
			expected Subject
			err      error
		}{
			{
				proto: &acl.Subject{
					Ref: &acl.Subject_Id{Id: "foo"},
				},
				expected: &SubjectID{ID: "foo"},
			},
			{
				proto: nil,
				err:   ErrNilSubject,
			},
			{
				proto: &acl.Subject{
					Ref: &acl.Subject_Set{
						Set: &acl.SubjectSet{
							Namespace: "n",
							Object:    "o",
							Relation:  "r",
						},
					},
				},
				expected: &SubjectSet{
					Namespace: "n",
					Object:    "o",
					Relation:  "r",
				},
			},
		} {
			t.Run(fmt.Sprintf("case=%d", i), func(t *testing.T) {
				actual, err := SubjectFromProto(tc.proto)
				require.True(t, errors.Is(err, tc.err))
				assert.Equal(t, tc.expected, actual)
			})
		}
	})

	t.Run("method=equals", func(t *testing.T) {
		for i, tc := range []struct {
			a, b   Subject
			equals bool
		}{
			{
				a:      &SubjectID{ID: "foo"},
				b:      &SubjectID{ID: "foo"},
				equals: true,
			},
			{
				a:      &SubjectID{ID: "foo"},
				b:      &SubjectID{ID: "bar"},
				equals: false,
			},
			{
				a:      &SubjectSet{},
				b:      &SubjectID{},
				equals: false,
			},
			{
				a: &SubjectSet{
					Namespace: "N",
					Object:    "O",
					Relation:  "R",
				},
				b: &SubjectSet{
					Namespace: "N",
					Object:    "O",
					Relation:  "R",
				},
				equals: true,
			},
			{
				a: &SubjectSet{
					Object:   "O",
					Relation: "R",
				},
				b: &SubjectSet{
					Namespace: "N",
					Object:    "O",
					Relation:  "R",
				},
				equals: false,
			},
			{
				a: &SubjectSet{
					Namespace: "N",
					Relation:  "R",
				},
				b: &SubjectSet{
					Namespace: "N",
					Object:    "O",
					Relation:  "R",
				},
				equals: false,
			},
			{
				a: &SubjectSet{
					Namespace: "N",
					Object:    "O",
				},
				b: &SubjectSet{
					Namespace: "N",
					Object:    "O",
					Relation:  "R",
				},
				equals: false,
			},
		} {
			t.Run(fmt.Sprintf("case=%d", i), func(t *testing.T) {
				assert.Equal(t, tc.equals, tc.a.Equals(tc.b))
				assert.Equal(t, tc.equals, tc.b.Equals(tc.a))
			})
		}
	})

	t.Run("case=url encoding-decoding", func(t *testing.T) {
		for i, sub := range []*SubjectSet{
			{
				Namespace: "n",
				Object:    "o",
				Relation:  "r",
			},
			{},
		} {
			t.Run(fmt.Sprintf("case=%d", i), func(t *testing.T) {
				enc := sub.ToURLQuery()
				dec := (&SubjectSet{}).FromURLQuery(enc)
				assert.Equal(t, sub, dec)
			})
		}
	})

	t.Run("case=url decoding-encoding", func(t *testing.T) {
		for i, vals := range []url.Values{
			{
				"namespace": []string{"n"},
				"object":    []string{"o"},
				"relation":  []string{"r"},
			},
			{
				"namespace": []string{""},
				"object":    []string{""},
				"relation":  []string{""},
			},
		} {
			t.Run(fmt.Sprintf("case=%d", i), func(t *testing.T) {
				dec := (&SubjectSet{}).FromURLQuery(vals)
				assert.Equal(t, vals, dec.ToURLQuery())
			})
		}
	})

	t.Run("case=json encoding", func(t *testing.T) {
		for i, tc := range []struct {
			sub  Subject
			json string
		}{
			{
				sub: &SubjectSet{
					Namespace: "n",
					Object:    "o",
					Relation:  "r",
				},
				json: `
{
	"namespace": "n",
	"object": "o",
	"relation": "r"
}`,
			},
			{
				sub:  &SubjectID{ID: "foo"},
				json: "\"foo\"",
			},
		} {
			t.Run(fmt.Sprintf("case=%d", i), func(t *testing.T) {
				enc, err := json.Marshal(tc.sub)
				require.NoError(t, err)
				assert.JSONEq(t, tc.json, string(enc))
			})
		}
	})
}

func TestInternalRelationTuple(t *testing.T) {
	t.Run("method=string encoding", func(t *testing.T) {
		assert.Equal(t, "n:o#r@s", (&InternalRelationTuple{
			Namespace: "n",
			Object:    "o",
			Relation:  "r",
			Subject:   &SubjectID{ID: "s"},
		}).String())
	})

	t.Run("method=string decoding", func(t *testing.T) {
		for i, tc := range []struct {
			enc      string
			err      error
			expected *InternalRelationTuple
		}{
			{
				enc: "n:o#r@s",
				expected: &InternalRelationTuple{
					Namespace: "n",
					Object:    "o",
					Relation:  "r",
					Subject:   &SubjectID{ID: "s"},
				},
			},
			{
				enc: "n:o#r@n:o#r",
				expected: &InternalRelationTuple{
					Namespace: "n",
					Object:    "o",
					Relation:  "r",
					Subject: &SubjectSet{
						Namespace: "n",
						Object:    "o",
						Relation:  "r",
					},
				},
			},
			{
				enc: "n:o#r@(n:o#r)",
				expected: &InternalRelationTuple{
					Namespace: "n",
					Object:    "o",
					Relation:  "r",
					Subject: &SubjectSet{
						Namespace: "n",
						Object:    "o",
						Relation:  "r",
					},
				},
			},
			{
				enc: "#dev:@ory#:working:@projects:keto#awesome",
				expected: &InternalRelationTuple{
					Namespace: "#dev",
					Object:    "@ory",
					Relation:  ":working:",
					Subject: &SubjectSet{
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
				actual, err := (&InternalRelationTuple{}).FromString(tc.enc)
				assert.True(t, errors.Is(err, tc.err), "%+v", err)
				assert.Equal(t, tc.expected, actual)
			})
		}
	})

	t.Run("case=url encoding-decoding", func(t *testing.T) {
		for i, r := range []*InternalRelationTuple{
			{
				Namespace: "n",
				Object:    "o",
				Relation:  "r",
				Subject:   &SubjectID{ID: "s"},
			},
			{
				Namespace: "n",
				Object:    "o",
				Relation:  "r",
				Subject: &SubjectSet{
					Namespace: "sn",
					Object:    "so",
					Relation:  "sr",
				},
			},
			{},
		} {
			t.Run(fmt.Sprintf("case=%d", i), func(t *testing.T) {
				res, err := (&InternalRelationTuple{}).FromURLQuery(r.ToURLQuery())
				require.NoError(t, err)
				assert.Equal(t, r, res)
			})
		}
	})

	t.Run("case=url decoding-encoding", func(t *testing.T) {
		for i, v := range []url.Values{
			{
				"namespace": []string{"n"},
				"object":    []string{"o"},
				"relation":  []string{"r"},
				"subject":   []string{"foo"},
			},
			{
				"namespace": []string{"n"},
				"object":    []string{"o"},
				"relation":  []string{"r"},
				"subject":   []string{"sn:so#sr"},
			},
		} {
			t.Run(fmt.Sprintf("case=%d", i), func(t *testing.T) {
				rt, err := (&InternalRelationTuple{}).FromURLQuery(v)
				require.NoError(t, err)
				assert.Equal(t, v, rt.ToURLQuery())
			})
		}
	})

	t.Run("case=proto decoding", func(t *testing.T) {
		for i, tc := range []struct {
			proto    TupleData
			expected *InternalRelationTuple
			err      error
		}{
			{
				proto: &acl.RelationTuple{
					Namespace: "n",
					Object:    "o",
					Relation:  "r",
					Subject:   nil,
				},
				err: ErrNilSubject,
			},
			{
				proto: &acl.RelationTuple{
					Namespace: "n",
					Object:    "o",
					Relation:  "r",
					Subject: &acl.Subject{
						Ref: &acl.Subject_Set{
							Set: &acl.SubjectSet{
								Namespace: "n",
								Object:    "o",
								Relation:  "r",
							},
						},
					},
				},
				expected: &InternalRelationTuple{
					Namespace: "n",
					Object:    "o",
					Relation:  "r",
					Subject: &SubjectSet{
						Namespace: "n",
						Object:    "o",
						Relation:  "r",
					},
				},
			},
			{
				proto: &acl.RelationTuple{
					Namespace: "n",
					Object:    "o",
					Relation:  "r",
					Subject: &acl.Subject{
						Ref: &acl.Subject_Id{
							Id: "user",
						},
					},
				},
				expected: &InternalRelationTuple{
					Namespace: "n",
					Object:    "o",
					Relation:  "r",
					Subject: &SubjectID{
						ID: "user",
					},
				},
			},
		} {
			t.Run(fmt.Sprintf("case=%d", i), func(t *testing.T) {
				actual, err := (&InternalRelationTuple{}).FromDataProvider(tc.proto)
				require.True(t, errors.Is(err, tc.err))
				assert.Equal(t, tc.expected, actual)
			})
		}
	})

	t.Run("format=JSON", func(t *testing.T) {
		t.Run("direction=encoding-decoding", func(t *testing.T) {
			for _, tc := range []struct {
				name     string
				rt       *InternalRelationTuple
				expected string
			}{
				{
					name: "with subject ID",
					rt: &InternalRelationTuple{
						Namespace: "n",
						Object:    "o",
						Relation:  "r",
						Subject:   &SubjectID{ID: "s"},
					},
					expected: `
{
	"namespace": "n",
	"object": "o",
	"relation": "r",
	"subject": "s"
}`,
				},
				{
					name: "with subject set",
					rt: &InternalRelationTuple{
						Namespace: "n",
						Object:    "o",
						Relation:  "r",
						Subject: &SubjectSet{
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
	"subject": {
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

					var dec InternalRelationTuple
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
					"namespace": []string{"n"},
					"object":    []string{"o"},
					"relation":  []string{"r"},
					"subject":   []string{"foo"},
				},
				r: &RelationQuery{
					Namespace: "n",
					Object:    "o",
					Relation:  "r",
					Subject:   &SubjectID{ID: "foo"},
				},
			},
			{
				v: url.Values{
					"namespace": []string{"n"},
					"object":    []string{"o"},
					"relation":  []string{"r"},
					"subject":   []string{"sn:so#sr"},
				},
				r: &RelationQuery{
					Namespace: "n",
					Object:    "o",
					Relation:  "r",
					Subject: &SubjectSet{
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
					Namespace: "n",
					Relation:  "r",
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

func TestRelationCollection(t *testing.T) {
	t.Run("case=prints all", func(t *testing.T) {
		expected := make([]*InternalRelationTuple, 3)
		for i := range expected {
			expected[i] = &InternalRelationTuple{
				Namespace: "n" + strconv.Itoa(i),
				Object:    "o" + strconv.Itoa(i),
				Relation:  "r" + strconv.Itoa(i),
				Subject:   &SubjectID{ID: "s" + strconv.Itoa(i)},
			}
		}
		expected[2].Subject = &SubjectSet{
			Namespace: "sn",
			Object:    "so",
			Relation:  "sr",
		}

		proto := make([]*acl.RelationTuple, 3)
		for i := range expected {
			proto[i] = &acl.RelationTuple{
				Namespace: "n" + strconv.Itoa(i),
				Object:    "o" + strconv.Itoa(i),
				Relation:  "r" + strconv.Itoa(i),
				Subject:   (&SubjectID{ID: "s" + strconv.Itoa(i)}).ToProto(),
			}
		}
		proto[2].Subject = (&SubjectSet{
			Namespace: "sn",
			Object:    "so",
			Relation:  "sr",
		}).ToProto()

		NewRelationCollection([]*InternalRelationTuple{})
		NewProtoRelationCollection([]*acl.RelationTuple{})

		for i, c := range []*RelationCollection{
			NewRelationCollection(expected),
			NewProtoRelationCollection(proto),
		} {
			t.Run(fmt.Sprintf("case=%d", i), func(t *testing.T) {
				var vals []string
				for _, row := range c.Table() {
					vals = append(vals, row...)
				}

				ev := reflect.ValueOf(expected)
				for el := 0; el < ev.Len(); el++ {
					et := reflect.TypeOf(expected).Elem().Elem()

					for f := 0; f < et.NumField(); f++ {
						v := ev.Index(el).Elem().Field(f)
						// private field
						if !v.CanSet() {
							continue
						}

						switch v.Kind() {
						case reflect.String:
							assert.Contains(t, vals, v.String())
						default:
							str := v.MethodByName("String").Call(nil)[0].String()
							assert.Contains(t, vals, str)
						}
					}
				}
			})
		}
	})

	t.Run("func=toInternal", func(t *testing.T) {
		for i, tc := range []struct {
			collection *RelationCollection
			expected   []*InternalRelationTuple
			err        error
		}{
			{
				collection: NewProtoRelationCollection([]*acl.RelationTuple{{
					Namespace: "n",
					Object:    "o",
					Relation:  "r",
					Subject:   (&SubjectID{ID: "s"}).ToProto(),
				}}),
				expected: []*InternalRelationTuple{{
					Namespace: "n",
					Object:    "o",
					Relation:  "r",
					Subject:   &SubjectID{ID: "s"},
				}},
			},
			{
				collection: NewProtoRelationCollection([]*acl.RelationTuple{{ /*subject is nil*/ }}),
				err:        ErrNilSubject,
			},
		} {
			t.Run(fmt.Sprintf("case=%d", i), func(t *testing.T) {
				actual, err := tc.collection.ToInternal()
				require.True(t, errors.Is(err, tc.err))
				assert.Equal(t, tc.expected, actual)
			})
		}
	})
}
