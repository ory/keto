package storage

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
	"github.com/stretchr/testify/require"
)

func TestMemoryManager(t *testing.T) {
	for k, m := range map[string]Manager{
		"memory": NewMemoryManager(),
	} {
		t.Run(fmt.Sprintf("manager=%s", k), func(t *testing.T) {
			require.Error(t, m.Get(nil, "test", "string", nil))

			t.Run("case=string", func(t *testing.T) {
				var vs string
				require.NoError(t, m.Upsert(nil, "test", "string", "foobar"))
				require.NoError(t, m.Get(nil, "test", "string", &vs))
				assert.EqualValues(t, "foobar", vs)
			})

			t.Run("case=int", func(t *testing.T) {
				var vs int
				require.NoError(t, m.Upsert(nil, "test", "int", 1234))
				require.NoError(t, m.Get(nil, "test", "int", &vs))
				assert.EqualValues(t, 1234, vs)
			})

			t.Run("case=list", func(t *testing.T) {
				for i := 0; i < 10; i++ {
					require.NoError(t, m.Upsert(nil, "test-list", fmt.Sprintf("list-%d", i), i))
				}

				var v int
				require.NoError(t, m.Get(nil, "test-list", "list-1", &v))
				assert.EqualValues(t, 1, v)

				var vs []int
				require.NoError(t, m.List(nil, "test-list", &vs, 10, 0))
				assert.Len(t, vs, 10)
				assert.EqualValues(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, vs)

				require.NoError(t, m.List(nil, "test-list", &vs, 5, 5))
				assert.Len(t, vs, 5)
				assert.EqualValues(t, []int{5, 6, 7, 8, 9}, vs)

			})
		})
	}
}
