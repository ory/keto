package namespace

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"testing"

	"github.com/ory/x/cmdx"
)

type (
	Namespace struct {
		ID     int             `json:"id" db:"id"`
		Name   string          `json:"name" db:"-"`
		Config json.RawMessage `json:"config" db:"-"`
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
		SetNamespaces(_ *testing.T, namespaces ...*Namespace)
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

	ErrNamespaceNotFound = errors.New("could not find namespace")
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
