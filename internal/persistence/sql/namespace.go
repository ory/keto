package sql

import (
	"context"
	"fmt"
	"io"

	"github.com/ory/x/pkgerx"

	"github.com/ory/keto/internal/namespace"
)

func tableFromNamespace(n *namespace.Namespace) string {
	return fmt.Sprintf("keto_%0.10d_relation_tuples", n.ID)
}

func (p *Persister) namespaceMigrationBox(n *namespace.Namespace) (*pkgerx.MigrationBox, error) {
	connDetails := *p.connDetails
	connDetails.Options["migration_table_name"] = fmt.Sprintf("keto_namespace_%0.10d_migrations", n.ID)

	return pkgerx.NewMigrationBox(namespaceMigrations, p.conn, p.l, pkgerx.WithTemplateValues(map[string]interface{}{
		"tableName": tableFromNamespace(n),
	}))
}

func (p *Persister) MigrateNamespaceUp(_ context.Context, n *namespace.Namespace) error {
	mb, err := p.namespaceMigrationBox(n)
	if err != nil {
		return err
	}

	return mb.Up()
}

func (p *Persister) MigrateNamespaceDown(_ context.Context, n *namespace.Namespace, steps int) error {
	mb, err := p.namespaceMigrationBox(n)
	if err != nil {
		return err
	}

	return mb.Down(steps)
}

func (p *Persister) NamespaceFromName(ctx context.Context, name string) (*namespace.Namespace, error) {
	return p.namespaces.GetNamespace(ctx, name)
}

func (p *Persister) NamespaceStatus(_ context.Context, w io.Writer, n *namespace.Namespace) error {
	mb, err := p.namespaceMigrationBox(n)
	if err != nil {
		return err
	}

	return mb.Status(w)
}
