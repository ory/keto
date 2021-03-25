package sql

import (
	"context"
	"fmt"

	"github.com/ory/x/popx"

	"github.com/ory/keto/internal/namespace"
)

func tableFromNamespace(n *namespace.Namespace) string {
	return fmt.Sprintf("keto_%0.10d_relation_tuples", n.ID)
}

func migrationTableFromNamespace(n *namespace.Namespace) string {
	return fmt.Sprintf("keto_namespace_%0.10d_migrations", n.ID)
}

func (p *Persister) namespaceMigrationBox(n *namespace.Namespace) (*popx.MigrationBox, error) {
	c, err := p.newConnection(map[string]string{
		"migration_table_name": migrationTableFromNamespace(n),
	})
	if err != nil {
		return nil, err
	}

	return popx.NewMigrationBox(
		namespaceMigrations,
		popx.NewMigrator(c, p.l, nil, 0),
		popx.WithTemplateValues(map[string]interface{}{
			"tableName": tableFromNamespace(n),
		}),
	)
}

func (p *Persister) NamespaceFromName(ctx context.Context, name string) (*namespace.Namespace, error) {
	return p.namespaces.GetNamespace(ctx, name)
}

func (p *Persister) NamespaceMigrationBox(_ context.Context, n *namespace.Namespace) (*popx.MigrationBox, error) {
	return p.namespaceMigrationBox(n)
}
