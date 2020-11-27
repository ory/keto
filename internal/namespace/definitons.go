package namespace

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/ory/x/cmdx"
)

type (
	Namespace struct {
		ID     int             `json:"id" db:"id"`
		Name   string          `json:"name" db:"name"`
		Config json.RawMessage `json:"config" db:"-"`
	}
	Status struct {
		CurrentVersion int `json:"current_version" db:"version"`
		NextVersion    int `json:"next_version" db:"-"`
	}
	Manager interface {
		MigrateNamespaceUp(ctx context.Context, n *Namespace) error
		NamespaceStatus(ctx context.Context, name string) (*Status, error)
	}
	ManagerProvider interface {
		NamespaceManager() Manager
	}
)

var _ cmdx.OutputEntry = &Status{}

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

func (n *Namespace) TableName() string {
	return "keto_namespace"
}
