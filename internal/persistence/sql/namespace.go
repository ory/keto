package sql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/pkg/errors"

	"github.com/gobuffalo/pop/v5"

	"github.com/ory/keto/internal/namespace"
	"github.com/ory/keto/internal/persistence"
)

type (
	namespaceRow struct {
		ID      int `db:"id"`
		Version int `db:"schema_version"`
	}
)

const (
	namespaceCreateStatement = `
CREATE TABLE %[1]s
(
    shard_id    varchar(64),
    object      varchar(64),
    relation    varchar(64),
    subject     varchar(256), /* can be <namespace:object#relation> or <user_id> */
    commit_time timestamp,

	PRIMARY KEY (shard_id, object, relation, subject, commit_time)
);

CREATE INDEX %[1]s_object_idx ON %[1]s (object);

CREATE INDEX %[1]s_user_set_idx ON %[1]s (object, relation);
`
	namespaceDropStatement = `
DROP INDEX %[1]s_user_set_idx;

DROP INDEX %[1]s_object_idx;

DROP TABLE %[1]s;
`

	mostRecentSchemaVersion = 1
)

func tableFromNamespace(n *namespace.Namespace) string {
	return fmt.Sprintf("keto_%0.10d_relation_tuples", n.ID)
}

func createStmt(n *namespace.Namespace) string {
	return fmt.Sprintf(namespaceCreateStatement, tableFromNamespace(n))
}

func dropStmt(n *namespace.Namespace) string {
	return fmt.Sprintf(namespaceDropStatement, tableFromNamespace(n))
}

func (p *Persister) MigrateNamespaceUp(ctx context.Context, n *namespace.Namespace) error {
	return p.transaction(ctx, func(ctx context.Context, c *pop.Connection) error {
		// TODO this is only creating new namespaces and not applying migrations
		nr := namespaceRow{
			ID:      n.ID,
			Version: mostRecentSchemaVersion,
		}

		// first create the table because of cockroach limitations, see https://github.com/cockroachdb/cockroach/issues/54477
		if err := c.RawQuery(createStmt(n)).Exec(); err != nil {
			return errors.WithStack(err)
		}

		return errors.WithStack(c.RawQuery(fmt.Sprintf("INSERT INTO %s (id, schema_version) VALUES (?, ?)", nr.TableName()), nr.ID, nr.Version).Exec())
	})
}

func (p *Persister) MigrateNamespaceDown(ctx context.Context, n *namespace.Namespace, _ int) error {
	return p.transaction(ctx, func(ctx context.Context, c *pop.Connection) error {
		if err := c.RawQuery(dropStmt(n)).Exec(); err != nil {
			return errors.WithStack(err)
		}

		return errors.WithStack(c.RawQuery(fmt.Sprintf("DELETE FROM %s WHERE id = ?", (&namespaceRow{}).TableName()), n.ID).Exec())
	})
}

func (p *Persister) NamespaceFromName(ctx context.Context, name string) (*namespace.Namespace, error) {
	return p.namespaces.GetNamespace(ctx, name)
}

func (p *Persister) NamespaceStatus(ctx context.Context, id int) (*namespace.Status, error) {
	var n namespaceRow
	if err := p.connection(ctx).Find(&n, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, persistence.ErrNamespaceUnknown
		}

		return nil, err
	}

	return &namespace.Status{
		CurrentVersion: n.Version,
		NextVersion:    mostRecentSchemaVersion,
	}, nil
}

func (n *namespaceRow) TableName() string {
	return "keto_namespace"
}
