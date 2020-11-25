package namespace

import (
	"encoding/json"
	"fmt"
	"github.com/ory/x/cmdx"
)

type (
	Namespace struct {
		ID     int    `json:"id"`
		Name   string `json:"name"`
		Config json.RawMessage
	}
	Status struct {
		CurrentVersion int `json:"current_version" db:"version"`
		NextVersion    int `json:"next_version" db:"-"`
	}
	Manager interface {
		MigrateNamespaceUp(n *Namespace) error
		NamespaceStatus(n *Namespace) (*Status, error)
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
