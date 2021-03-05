package namespace

import (
	"context"
	"encoding/json"
	"io"
)

type (
	Namespace struct {
		ID        int             `json:"id" db:"id" toml:"id"`
		Name      string          `json:"name" db:"-" toml:"name"`
		Config    json.RawMessage `json:"config,omitempty" db:"-" toml:"config,omitempty"`
		Relations []struct {
			Name    string
			Rewrite []struct {
				Effect string     // union, intersection, ...
				Rules  []struct { // this will probably be an interface implementing all kinds of rules
					ComputedSubjectSetRelationName string
				}
			}
		}
	}
	Migrator interface {
		MigrateNamespaceUp(ctx context.Context, n *Namespace) error
		MigrateNamespaceDown(ctx context.Context, n *Namespace, steps int) error
		NamespaceStatus(ctx context.Context, w io.Writer, n *Namespace) error
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
