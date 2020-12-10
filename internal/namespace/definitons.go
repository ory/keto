package namespace

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/ory/x/cmdx"
)

type (
	Namespace struct {
		ID     int             `json:"id" db:"id" toml:"id"`
		Name   string          `json:"name" db:"-" toml:"name"`
		Config json.RawMessage `json:"config,omitempty" db:"-" toml:"config,omitempty"`
	}
	Status struct {
		CurrentVersion int `json:"current_version" db:"-"`
		NextVersion    int `json:"next_version" db:"-"`
	}
	Migrator interface {
		MigrateNamespaceUp(ctx context.Context, n *Namespace) error
		NamespaceStatus(ctx context.Context, id int) (*Status, error)
	}
	Manager interface {
		GetNamespace(ctx context.Context, name string) (*Namespace, error)
		Namespaces(ctx context.Context) ([]*Namespace, error)
	}
	ManagerProvider interface {
		NamespaceManager() Manager
	}
	MigratorProvider interface {
		NamespaceMigrator() Migrator
	}
)

var (
	_ cmdx.OutputEntry = &Status{}
)

func (s *Status) Header() []string {
	return []string{
		"CURRENT VERSION",
		"NEXT VERSION",
	}
}

func (s *Status) Fields() []string {
	return []string{
		fmt.Sprintf("%d", s.CurrentVersion),
		fmt.Sprintf("%d", s.NextVersion),
	}
}

func (s *Status) Interface() interface{} {
	return s
}
