package namespace

import (
	"context"
	"encoding/json"
)

type (
	Namespace struct {
		// Deprecated: Only use the Name instead.
		ID int32 `json:"id" db:"-" toml:"id"`
		// The unique name of the namespace.
		Name   string          `json:"name" db:"-" toml:"name"`
		Config json.RawMessage `json:"config,omitempty" db:"-" toml:"config,omitempty"`
	}
	Manager interface {
		GetNamespaceByName(ctx context.Context, name string) (*Namespace, error)
		// Deprecated: Use GetNamespaceByName instead.
		GetNamespaceByConfigID(ctx context.Context, id int32) (*Namespace, error)
		Namespaces(ctx context.Context) ([]*Namespace, error)
		ShouldReload(newValue interface{}) bool
	}
	ManagerProvider interface {
		NamespaceManager() (Manager, error)
	}
)
