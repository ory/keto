package migrations_test

import (
	"context"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ory/keto/internal/driver"
	"github.com/ory/keto/internal/driver/config"
	"github.com/ory/keto/internal/namespace"
	"github.com/ory/keto/internal/persistence/sql/migrations"
	"github.com/ory/keto/internal/relationtuple"
	"github.com/ory/keto/internal/x"
	"github.com/ory/keto/internal/x/dbx"
)

func TestToSingleTableMigrator(t *testing.T) {
	t.Parallel()
	const debugOnDisk = false

	for _, dsn := range dbx.GetDSNs(t, debugOnDisk) {
		dsn := dsn
		t.Run("db="+dsn.Name, func(t *testing.T) {
			t.Parallel()

			r := driver.NewTestRegistry(t, dsn)
			ctx := context.Background()
			var nn []*namespace.Namespace
			m := migrations.NewToSingleTableMigrator(r)

			setup := func(t *testing.T) *namespace.Namespace {
				n := &namespace.Namespace{
					Name: t.Name(),
					ID:   int32(len(nn)),
				}

				nn = append(nn, n)

				mb, err := m.NamespaceMigrationBox(ctx, n)
				require.NoError(t, err)
				require.NoError(t, mb.Up(ctx))

				t.Cleanup(func() {
					if debugOnDisk {
						return
					}
					require.NoError(t, mb.Down(ctx, -1))
				})

				require.NoError(t, r.Config(ctx).Set(config.KeyNamespaces, nn))

				return n
			}

			t.Run("case=simple tuples", func(t *testing.T) {
				n := setup(t)

				// insert tuples into the old table
				sID := &relationtuple.InternalRelationTuple{
					Namespace: n.Name,
					Object:    "a",
					Relation:  "a",
					Subject:   &relationtuple.SubjectID{ID: "a"},
				}
				sSet := &relationtuple.InternalRelationTuple{
					Namespace: n.Name,
					Object:    "b",
					Relation:  "b",
					Subject: &relationtuple.SubjectSet{
						Namespace: n.Name,
						Object:    "b",
						Relation:  "b",
					},
				}
				require.NoError(t, m.InsertOldRelationTuples(ctx, n, sID, sSet))

				// get the tuple from the old table
				oldRts, next, err := m.GetOldRelationTuples(ctx, n, 0, 100)
				require.NoError(t, err)
				assert.False(t, next)
				require.Len(t, oldRts, 2)
				for i, r := range []*relationtuple.InternalRelationTuple{sID, sSet} {
					assert.Equal(t, r.Namespace, oldRts[i].Namespace.Name)
					assert.Equal(t, r.Object, oldRts[i].Object)
					assert.Equal(t, r.Relation, oldRts[i].Relation)
					assert.Equal(t, r.Subject.String(), oldRts[i].Subject)
				}

				// migrate to new table
				require.NoError(t, m.MigrateNamespace(ctx, n))

				// get the tuple from the new table
				rts, nextToken, err := r.RelationTupleManager().GetRelationTuples(ctx, &relationtuple.RelationQuery{Namespace: n.Name})
				require.NoError(t, err)
				assert.Equal(t, "", nextToken)
				require.Len(t, rts, 2)
				assert.Equal(t, sID, rts[0])
				assert.Equal(t, sSet, rts[1])
			})

			t.Run("case=paginates", func(t *testing.T) {
				n := setup(t)

				defer func(old int) {
					m.PerPage = old
				}(m.PerPage)
				m.PerPage = 1

				rts := make([]*relationtuple.InternalRelationTuple, 10)
				for i := range rts {
					rts[i] = &relationtuple.InternalRelationTuple{
						Namespace: n.Name,
						Object:    strconv.Itoa(i),
						Relation:  strconv.Itoa(i),
						Subject:   &relationtuple.SubjectID{ID: strconv.Itoa(i)},
					}
				}

				require.NoError(t, m.InsertOldRelationTuples(ctx, n, rts...))
				require.NoError(t, m.MigrateNamespace(ctx, n))

				migrated, nextToken, err := r.RelationTupleManager().GetRelationTuples(ctx, &relationtuple.RelationQuery{Namespace: n.Name}, x.WithSize(len(rts)))
				require.NoError(t, err)
				assert.Equal(t, "", nextToken)
				assert.Equal(t, rts, migrated)
			})

			t.Run("case=non-deserializable tuple", func(t *testing.T) {
				n := setup(t)

				valid := &relationtuple.InternalRelationTuple{
					Namespace: n.Name,
					Object:    "o1",
					Relation:  "r",
					Subject:   &relationtuple.SubjectID{ID: "s"},
				}
				require.NoError(t, m.InsertOldRelationTuples(ctx, n, &relationtuple.InternalRelationTuple{
					Namespace: n.Name,
					Object:    "o0",
					Relation:  "r",
					Subject:   &relationtuple.SubjectID{ID: "invalid#subject-id"},
				}, &relationtuple.InternalRelationTuple{
					Namespace: n.Name,
					Object:    "o1",
					Relation:  "r",
					Subject:   &relationtuple.SubjectID{ID: "invalid#subject-id"},
				}, valid))
				err := m.MigrateNamespace(ctx, n)
				require.Error(t, err)
				invalid, ok := err.(migrations.ErrInvalidTuples)
				require.True(t, ok)
				assert.Len(t, invalid, 2)

				rts, next, err := r.Persister().GetRelationTuples(ctx, &relationtuple.RelationQuery{
					Namespace: n.Name,
				})
				require.NoError(t, err)
				require.Equal(t, "", next)
				assert.Equal(t, []*relationtuple.InternalRelationTuple{valid}, rts)
			})
		})
	}
}

func TestToSingleTableMigrator_HasLegacyTable(t *testing.T) {
	t.Parallel()
	const debugOnDisk = false

	for _, dsn := range dbx.GetDSNs(t, debugOnDisk) {
		dsn := dsn
		t.Run("db="+dsn.Name, func(t *testing.T) {
			t.Parallel()

			t.Run("case=simple detection", func(t *testing.T) {
				ctx := context.Background()
				reg := driver.NewTestRegistry(t, dsn)
				m := migrations.NewToSingleTableMigrator(reg)

				nspaces := []*namespace.Namespace{{
					ID:   3,
					Name: "foo",
				}}
				require.NoError(t, reg.Config(ctx).Set(config.KeyNamespaces, nspaces))

				// expect to not report legacy table
				legacyNamespaces, err := m.LegacyNamespaces(ctx)
				require.NoError(t, err)
				assert.Len(t, legacyNamespaces, 0)

				// migrate legacy table up
				mb, err := m.NamespaceMigrationBox(ctx, nspaces[0])
				require.NoError(t, err)
				require.NoError(t, mb.Up(ctx))

				// expect to report legacy table
				legacyNamespaces, err = m.LegacyNamespaces(ctx)
				require.NoError(t, err)
				assert.Equal(t, nspaces, legacyNamespaces)

				// migrate legacy down
				require.NoError(t, mb.Down(ctx, -1))

				// expect to not report legacy
				legacyNamespaces, err = m.LegacyNamespaces(ctx)
				require.NoError(t, err)
				assert.Len(t, legacyNamespaces, 0)
			})

			t.Run("case=multiple namespaces", func(t *testing.T) {
				ctx := context.Background()
				reg := driver.NewTestRegistry(t, dsn)
				m := migrations.NewToSingleTableMigrator(reg)

				nspaces := []*namespace.Namespace{{
					ID:   0,
					Name: "a",
				}, {
					ID:   1,
					Name: "b",
				}, {
					ID:   2,
					Name: "c",
				}}
				require.NoError(t, reg.Config(ctx).Set(config.KeyNamespaces, nspaces))

				for _, n := range nspaces {
					// migrate legacy table up
					mb, err := m.NamespaceMigrationBox(ctx, n)
					require.NoError(t, err)
					require.NoError(t, mb.Up(ctx))
				}

				ln, err := m.LegacyNamespaces(ctx)
				require.NoError(t, err)
				assert.ElementsMatch(t, nspaces, ln)
			})
		})
	}
}
