package sql

import (
	"context"
	"fmt"
	"github.com/pkg/errors"

	"github.com/gobuffalo/pop/v5"

	"github.com/ory/keto/internal/namespace"
)

const namespaceCreateStatement = `
CREATE TABLE %[1]s
(
    shard_id    varchar(64),
    object_id   varchar(64),
    relation    varchar(64),
    subject     varchar(256), /* can be object_id:namespace#relation or user_id */
    commit_time timestamp,

	PRIMARY KEY (shard_id, object_id, relation, subject, commit_time)
);

CREATE INDEX %[1]s_object_id_idx ON %[1]s (object_id);

CREATE INDEX %[1]s_user_set_idx ON %[1]s (object_id, relation);
`

func tableFromNamespace(n *namespace.Namespace) string {
	return fmt.Sprintf("keto_%0.10d_relation_tuples", n.ID)
}

func createStmt(n *namespace.Namespace) string {
	return fmt.Sprintf(namespaceCreateStatement, tableFromNamespace(n))
}

func (p *Persister) NewNamespace(ctx context.Context, n *namespace.Namespace) error {
	return p.conn.Transaction(func(tx *pop.Connection) error {
		if err := tx.Create(n); err != nil {
			return err
		}

		return tx.RawQuery(createStmt(n)).Exec()
	})
}

func (p *Persister) NamespaceFromName(ctx context.Context, name string) (*namespace.Namespace, error) {
	var n namespace.Namespace

	return &n, errors.WithStack(
		p.conn.Where("name = ?", name).First(&n))
}
