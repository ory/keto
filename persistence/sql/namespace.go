package sql

import (
	"context"
	"fmt"
)

const namespaceCreateStatement = `
CREATE TABLE %s
(
    shard_id    varchar(64),
    object_id   varchar(64),
    relation    varchar(64),
    subject     varchar(128), /* can be object_id#relation or user_id */
    commit_time timestamp,
    PRIMARY KEY (shard_id, object_id, relation, subject, commit_time)
);

CREATE INDEX object_id_idx ON %s (object_id);

CREATE INDEX user_set_idx ON %s (object_id, relation);
`

func sqlSafeTableFromNamespace(n string) string {
	// TODO AVOID SQL INJECTION
	return fmt.Sprintf("keto_%s_relation_tuples", n)
}

func createStmt(namespace string) string {
	tableName := sqlSafeTableFromNamespace(namespace)
	return fmt.Sprintf(namespaceCreateStatement, tableName, tableName, tableName)
}

func (p *Persister) NewNamespace(ctx context.Context, n string) error {
	return p.conn.RawQuery(createStmt(n)).Exec()
}
