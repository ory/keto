package namespace

import (
	"context"
	"encoding/json"
)

type (
	Namespace struct {
		ID     int64           `json:"id" db:"-" toml:"id"`
		Name   string          `json:"name" db:"-" toml:"name"`
		Config json.RawMessage `json:"config,omitempty" db:"-" toml:"config,omitempty"`
	}
	Manager interface {
		GetNamespaceByName(ctx context.Context, name string) (*Namespace, error)
		GetNamespaceByConfigID(ctx context.Context, id int64) (*Namespace, error)
		Namespaces(ctx context.Context) ([]*Namespace, error)
	}
	ManagerProvider interface {
		NamespaceManager() (Manager, error)
	}
)
