package namespace

import (
	"context"
	"encoding/json"
)

type (
	Namespace struct {
		ID     int32           `json:"id" db:"-" toml:"id"`     // Deprecated: Use the Name instead.
		Name   string          `json:"name" db:"-" toml:"name"` // The unique name of the namespace.
		Config json.RawMessage `json:"config,omitempty" db:"-" toml:"config,omitempty"`
	}
	Manager interface {
		GetNamespaceByName(ctx context.Context, name string) (*Namespace, error)
		GetNamespaceByConfigID(ctx context.Context, id int32) (*Namespace, error) // Deprecated: Use GetNamespaceByName instead.
		Namespaces(ctx context.Context) ([]*Namespace, error)
		ShouldReload(newValue interface{}) bool
	}
	ManagerProvider interface {
		NamespaceManager() (Manager, error)
	}
)
