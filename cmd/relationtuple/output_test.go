// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package relationtuple

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ory/x/pointerx"

	"github.com/ory/keto/ketoapi"
	rts "github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2"
)

func TestRelationCollection(t *testing.T) {
	t.Run("case=prints all", func(t *testing.T) {
		expected := make([]*ketoapi.RelationTuple, 3)
		for i := range expected {
			expected[i] = &ketoapi.RelationTuple{
				Namespace: "n" + strconv.Itoa(i),
				Object:    "o" + strconv.Itoa(i),
				Relation:  "r" + strconv.Itoa(i),
				SubjectID: pointerx.Ptr("s" + strconv.Itoa(i)),
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

		for i, c := range []*Collection{
			NewAPICollection(expected),
			MustNewProtoCollection(proto),
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
						// private field or nil
						if !v.CanSet() || (v.Kind() == reflect.Pointer && v.IsNil()) {
							continue
						}

						switch k := v.Kind(); {
						case k == reflect.String:
							assert.Contains(t, vals, v.String())
						case k == reflect.Ptr && v.Elem().Kind() == reflect.String:
							assert.Contains(t, vals, v.Elem().String())
						default:
							t.Logf("unhandled kind %s %T", v.Kind(), v.Interface())
							str := v.MethodByName("String").Call(nil)[0].String()
							assert.Contains(t, vals, str)
						}
					}
				}
			})
		}
	})

	t.Run("func=NewProtoCollection", func(t *testing.T) {
		_, err := NewProtoCollection([]*rts.RelationTuple{{ /*subject is nil*/ }})
		require.ErrorIs(t, err, ketoapi.ErrNilSubject)
	})

	t.Run("func=json", func(t *testing.T) {
		ts := []*ketoapi.RelationTuple{
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
				SubjectSet: &ketoapi.SubjectSet{
					Namespace: "sn",
					Object:    "so",
					Relation:  "sr",
				},
			},
		}
		expected, err := json.Marshal(ts)
		require.NoError(t, err)

		collection := NewAPICollection(ts)
		actual, err := json.Marshal(collection)
		require.NoError(t, err)
		assert.Equal(t, string(expected), string(actual))
	})
}
