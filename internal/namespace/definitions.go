package namespace

import (
	"context"
	"encoding/json"

	"github.com/ory/keto/internal/namespace/ast"
)

type (
	Namespace struct {
		ID     int32           `json:"id" db:"-" toml:"id"`
		Name   string          `json:"name" db:"-" toml:"name"`
		Config json.RawMessage `json:"config,omitempty" db:"-" toml:"config,omitempty"`

		Relations []ast.Relation `json:"-" db:"-"`
	}
	Manager interface {
		GetNamespaceByName(ctx context.Context, name string) (*Namespace, error)
		GetNamespaceByConfigID(ctx context.Context, id int32) (*Namespace, error)
		Namespaces(ctx context.Context) ([]*Namespace, error)
		ShouldReload(newValue interface{}) bool
	}
	ManagerProvider interface {
		NamespaceManager() (Manager, error)
	}
)
