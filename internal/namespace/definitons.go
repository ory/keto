package namespace

import (
	"context"
	"encoding/json"

	"github.com/ory/x/popx"
)

type (
	Namespace struct {
		ID     int             `json:"id" db:"id" toml:"id"`
		Name   string          `json:"name" db:"-" toml:"name"`
		Config json.RawMessage `json:"config,omitempty" db:"-" toml:"config,omitempty"`
	}
	Migrator interface {
		NamespaceMigrationBox(ctx context.Context, n *Namespace) (*popx.MigrationBox, error)
	}
	Manager interface {
		GetNamespace(ctx context.Context, name string) (*Namespace, error)
		Namespaces(ctx context.Context) ([]*Namespace, error)
	}
	ManagerProvider interface {
		NamespaceManager() (Manager, error)
	}
	MigratorProvider interface {
		NamespaceMigrator() Migrator
	}
)
