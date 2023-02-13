// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package namespace

import (
	"context"
	"encoding/json"

	"github.com/ory/keto/internal/namespace/ast"
)

type (
	Namespace struct {
		// Deprecated: Only use the Name instead.
		ID int32 `json:"id" db:"-" toml:"id"`
		// The unique name of the namespace.
		Name   string          `json:"name" db:"-" toml:"name"`
		Config json.RawMessage `json:"config,omitempty" db:"-" toml:"config,omitempty"`

		Relations []ast.Relation `json:"-" db:"-"`
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
