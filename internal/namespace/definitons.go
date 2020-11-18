package namespace

import "encoding/json"

type (
	Namespace struct {
		ID     int    `json:"id"`
		Name   string `json:"name"`
		Config json.RawMessage
	}
	Status struct {
		Version int `db:"version"`
	}
	Manager interface {
		MigrateNamespaceUp(n *Namespace) error
		NamespaceStatus(n *Namespace) (*Status, error)
	}
	ManagerProvider interface {
		NamespaceManager() Manager
	}
)
