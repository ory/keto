package storage

import (
	"context"
	"flag"
	"fmt"
	"log"
	"sync"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx/v4"
	"github.com/open-policy-agent/opa/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ory/x/sqlcon/dockertest"
)

var managers = map[string]Manager{
	"memory": NewMemoryManager(),
}
var m sync.Mutex

func TestMain(m *testing.M) {
	runner := dockertest.Register()

	flag.Parse()
	if !testing.Short() {
		dockertest.Parallel([]func(){
			connectToPG,
			connectToMySQL,
		})
	}

	runner.Exit(m.Run())
}

func connectToMySQL() {
	db, err := dockertest.ConnectToTestMySQL()
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	s := NewSQLManager(db)
	m.Lock()
	managers["mysql"] = s
	m.Unlock()

	if _, err := s.CreateSchemas(db); err != nil {
		log.Fatalf("Unable to create schemas: %s", err)
	}
}

func connectToPG() {
	db, err := dockertest.ConnectToTestPostgreSQL()
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	s := NewSQLManager(db)
	m.Lock()
	managers["postgres"] = s
	m.Unlock()

	if _, err := s.CreateSchemas(db); err != nil {
		log.Fatalf("Unable to create schemas: %s", err)
	}
}

func TestMemoryManager(t *testing.T) {
	for k, m := range managers {
		t.Run(fmt.Sprintf("manager=%s", k), func(t *testing.T) {
			ctx := context.Background()

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

			t.Run("case=upsert", func(t *testing.T) {
				var v string
				require.NoError(t, m.Upsert(ctx, "test-upsert", "foo", "bar"))
				require.NoError(t, m.Get(ctx, "test-upsert", "foo", &v))
				assert.Equal(t, "bar", v)

				require.NoError(t, m.Upsert(ctx, "test-upsert", "foo", "baz"))
				require.NoError(t, m.Get(ctx, "test-upsert", "foo", &v))
				assert.Equal(t, "baz", v)

				var vs []string
				require.NoError(t, m.List(ctx, "test-upsert", &vs, 10, 0))
				assert.Equal(t, 1, len(vs))
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

			t.Run("case=listall", func(t *testing.T) {
				n := 102
				for i := 0; i < n; i++ {
					require.NoError(t, m.Upsert(ctx, "test-listall", fmt.Sprintf("list-%d", i), i))
				}
				expected := make([]int, n)
				// populate with 0 - 101
				for i := range expected {
					expected[i] = i
				}

				var v int
				require.NoError(t, m.Get(ctx, "test-listall", "list-1", &v))
				assert.EqualValues(t, 1, v)

				var vs []int
				require.NoError(t, m.ListAll(ctx, "test-listall", &vs))
				assert.Len(t, vs, 102)
				assert.EqualValues(t, expected, vs)

			})

			t.Run("case=delete", func(t *testing.T) {
				for i := 0; i < 10; i++ {
					require.NoError(t, m.Upsert(ctx, "test-delete", fmt.Sprintf("delete-%d", i), i))

					var v int
					require.NoError(t, m.Get(ctx, "test-delete", fmt.Sprintf("delete-%d", i), &v))
					assert.EqualValues(t, i, v)
					require.NoError(t, m.Delete(ctx, "test-delete", fmt.Sprintf("delete-%d", i)))
					require.Error(t, m.Get(ctx, "test-delete", fmt.Sprintf("delete-%d", i), &v))
				}
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
