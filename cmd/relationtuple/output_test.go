package relationtuple

import (
	"fmt"
	"github.com/ory/keto/internal/x"
	"github.com/ory/keto/ketoapi"
	rts "github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"reflect"
	"strconv"
	"testing"
)

func TestRelationCollection(t *testing.T) {
	t.Run("case=prints all", func(t *testing.T) {
		expected := make([]*ketoapi.RelationTuple, 3)
		for i := range expected {
			expected[i] = &ketoapi.RelationTuple{
				Namespace: "n" + strconv.Itoa(i),
				Object:    "o" + strconv.Itoa(i),
				Relation:  "r" + strconv.Itoa(i),
				SubjectID: x.Ptr("s" + strconv.Itoa(i)),
			}
		}
		expected[2].SubjectSet = &ketoapi.SubjectSet{
			Namespace: "sn",
			Object:    "so",
			Relation:  "sr",
		}
		expected[2].SubjectID = nil

		proto := make([]*rts.RelationTuple, 3)
		for i := range expected {
			proto[i] = &rts.RelationTuple{
				Namespace: "n" + strconv.Itoa(i),
				Object:    "o" + strconv.Itoa(i),
				Relation:  "r" + strconv.Itoa(i),
				Subject:   rts.NewSubjectID("s" + strconv.Itoa(i)),
			}
		}
		proto[2].Subject = rts.NewSubjectSet("sn", "so", "sr")

		NewAPICollection([]*ketoapi.RelationTuple{})
		NewProtoCollection([]*rts.RelationTuple{})

		for i, c := range []*Collection{
			NewAPICollection(expected),
			NewProtoCollection(proto),
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
			collection *Collection
			expected   []*ketoapi.RelationTuple
			err        error
		}{
			{
				collection: NewProtoCollection([]*rts.RelationTuple{{
					Namespace: "n",
					Object:    "o",
					Relation:  "r",
					Subject:   rts.NewSubjectID("s"),
				}}),
				expected: []*ketoapi.RelationTuple{{
					Namespace: "n",
					Object:    "o",
					Relation:  "r",
					SubjectID: x.Ptr("s"),
				}},
			},
			{
				collection: NewProtoCollection([]*rts.RelationTuple{{ /*subject is nil*/ }}),
				err:        ketoapi.ErrNilSubject,
			},
		} {
			t.Run(fmt.Sprintf("case=%d", i), func(t *testing.T) {
				actual, err := tc.collection.Normalize()
				require.ErrorIs(t, err, tc.err)
				assert.Equal(t, tc.expected, actual)
			})
		}
	})
}
