package storage

import (
	"context"
	"fmt"
	"github.com/open-policy-agent/opa/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMemoryManager(t *testing.T) {
	for k, m := range map[string]Manager{
		"memory": NewMemoryManager(),
	} {
		t.Run(fmt.Sprintf("manager=%s", k), func(t *testing.T) {
			ctx := context.TODO()

			require.Error(t, m.Get(ctx, "test", "string", nil))

			t.Run("case=string", func(t *testing.T) {
				var vs string
				require.NoError(t, m.Upsert(ctx, "test", "string", "foobar"))
				require.NoError(t, m.Get(ctx, "test", "string", &vs))
				assert.EqualValues(t, "foobar", vs)
			})

			t.Run("case=int", func(t *testing.T) {
				var vs int
				require.NoError(t, m.Upsert(ctx, "test", "int", 1234))
				require.NoError(t, m.Get(ctx, "test", "int", &vs))
				assert.EqualValues(t, 1234, vs)
			})

			t.Run("case=list", func(t *testing.T) {
				for i := 0; i < 10; i++ {
					require.NoError(t, m.Upsert(ctx, "test-list", fmt.Sprintf("list-%d", i), i))
				}

				var v int
				require.NoError(t, m.Get(ctx, "test-list", "list-1", &v))
				assert.EqualValues(t, 1, v)

				var vs []int
				require.NoError(t, m.List(ctx, "test-list", &vs, 10, 0))
				assert.Len(t, vs, 10)
				assert.EqualValues(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, vs)

				require.NoError(t, m.List(ctx, "test-list", &vs, 5, 5))
				assert.Len(t, vs, 5)
				assert.EqualValues(t, []int{5, 6, 7, 8, 9}, vs)
			})

			t.Run("case=storage", func(t *testing.T) {
				for i := 0; i < 2; i++ {
					require.NoError(t, m.Upsert(ctx, "/tests/storage/bars", fmt.Sprintf("list-%d", i), fmt.Sprintf("a-%d", i)))
					require.NoError(t, m.Upsert(ctx, "/tests/storage/foos", fmt.Sprintf("list-%d", i), fmt.Sprintf("b-%d", i)))
				}

				s, err := m.Storage(ctx, `{"tests": {"storage": {"foos": [], "bars": []}}}`, []string{"/tests/storage/foos", "/tests/storage/bars"})
				require.NoError(t, err)

				tx, err := s.NewTransaction(ctx)
				require.NoError(t, err)

				res, err := s.Read(ctx, tx, storage.MustParsePath("/tests/storage/bars"))
				require.NoError(t, err)
				assert.Equal(t, `[a-0 a-1]`, fmt.Sprintf("%s", res))

				res, err = s.Read(ctx, tx, storage.MustParsePath("/tests/storage/foos"))
				require.NoError(t, err)
				assert.Equal(t, `[b-0 b-1]`, fmt.Sprintf("%s", res))
			})
		})
	}
}
