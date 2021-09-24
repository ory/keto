package namespace

import (
	"context"
	"io"
	"os"
	"path"
	"testing"

	"github.com/ory/keto/internal/relationtuple"

	"github.com/ory/x/cmdx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ory/keto/internal/driver"
	"github.com/ory/keto/internal/driver/config"
	"github.com/ory/keto/internal/namespace"
	"github.com/ory/keto/internal/x/dbx"
)

func TestMigrateLegacy(t *testing.T) {
	setup := func(t *testing.T) (*cmdx.CommandExecuter, *driver.RegistryDefault) {
		fp := path.Join(t.TempDir(), "db.sqlite")
		dst, err := os.Create(fp)
		require.NoError(t, err)
		defer dst.Close()
		src, err := os.Open("migrate_legacy_snapshot.sqlite")
		require.NoError(t, err)
		defer src.Close()
		_, err = io.Copy(dst, src)
		require.NoError(t, err)

		reg := driver.NewTestRegistry(t, &dbx.DsnT{
			Name:        "sqlite",
			Conn:        "sqlite://" + fp + "?_fk=true",
			MigrateUp:   true,
			MigrateDown: false,
		})
		nspaces := []*namespace.Namespace{
			{
				ID:   0,
				Name: "a",
			},
			{
				ID:   1,
				Name: "b",
			},
		}
		require.NoError(t, reg.Config().Set(config.KeyNamespaces, nspaces))

		c := &cmdx.CommandExecuter{
			New: NewMigrateLegacyCmd,
			Ctx: context.WithValue(context.Background(), driver.RegistryContextKey, reg),
		}

		return c, reg
	}

	t.Run("case=invalid subject", func(t *testing.T) {
		c, reg := setup(t)

		conn, err := reg.PopConnection()
		require.NoError(t, err)
		require.NoError(t, conn.RawQuery("insert into keto_0000000000_relation_tuples (shard_id, object, relation, subject, commit_time) values ('foo', 'obj', 'rel', 'invalid#subject', 'now')").Exec())

		stdErr := c.ExecExpectedErr(t, "-y", "a")
		assert.Contains(t, stdErr, "found non-deserializable relationtuples")
		assert.Contains(t, stdErr, "invalid#subject")

		assert.Contains(t, c.ExecNoErr(t, "-y", "--down-only", "a"), "Successfully migrated down")
	})

	t.Run("case=migrates down only", func(t *testing.T) {
		c, reg := setup(t)

		conn, err := reg.PopConnection()
		require.NoError(t, err)
		require.NoError(t, conn.RawQuery("insert into keto_0000000000_relation_tuples (shard_id, object, relation, subject, commit_time) values ('foo', 'obj', 'rel', 'sub', 'now')").Exec())

		c.ExecNoErr(t, "-y", "--down-only", "a")

		rts, _, err := reg.RelationTupleManager().GetRelationTuples(context.Background(), &relationtuple.RelationQuery{
			Namespace: "a",
			Object:    "obj",
		})
		require.NoError(t, err)
		assert.Len(t, rts, 0)
	})
}
