package relationtuple

import (
	"fmt"
	"github.com/gofrs/uuid"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubject(t *testing.T) {
	t.Run("method=equals", func(t *testing.T) {
		a, b := uuid.Must(uuid.NewV4()), uuid.Must(uuid.NewV4())

		for i, tc := range []struct {
			a, b   Subject
			equals bool
		}{
			{
				a:      &SubjectID{ID: a},
				b:      &SubjectID{ID: a},
				equals: true,
			},
			{
				a:      &SubjectID{ID: a},
				b:      &SubjectID{ID: b},
				equals: false,
			},
			{
				a:      &SubjectSet{},
				b:      &SubjectID{},
				equals: false,
			},
			{
				a: &SubjectSet{
					Namespace: 1,
					Object:    a,
					Relation:  "R",
				},
				b: &SubjectSet{
					Namespace: 1,
					Object:    a,
					Relation:  "R",
				},
				equals: true,
			},
			{
				a: &SubjectSet{
					Object:   a,
					Relation: "R",
				},
				b: &SubjectSet{
					Namespace: 1,
					Object:    a,
					Relation:  "R",
				},
				equals: false,
			},
			{
				a: &SubjectSet{
					Namespace: 1,
					Relation:  "R",
				},
				b: &SubjectSet{
					Namespace: 1,
					Object:    a,
					Relation:  "R",
				},
				equals: false,
			},
			{
				a: &SubjectSet{
					Namespace: 1,
					Object:    a,
				},
				b: &SubjectSet{
					Namespace: 1,
					Object:    a,
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
}
