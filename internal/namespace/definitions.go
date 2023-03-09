// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package namespace

import (
	"context"
	"encoding/json"

	"github.com/ory/herodot"

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

func ASTRelationFor(ctx context.Context, m Manager, namespace, relation string) (*ast.Relation, error) {
	// Special case: If the relationTuple's relation is empty, then it is not an
	// error that the relation was not found.
	if relation == "" {
		return nil, nil
	}
	ns, err := m.GetNamespaceByName(ctx, namespace)
	if err != nil {
		// On an unknown namespace the answer should be "not allowed", not "not
		// found". Therefore, we don't return the error here.
		return nil, nil
	}

	// Special case: If Relations is empty, then there is no namespace
	// configuration, and it is not an error that the relation was not found.
	if len(ns.Relations) == 0 {
		return nil, nil
	}

	for _, rel := range ns.Relations {
		if rel.Name == relation {
			return &rel, nil
		}
	}
	return nil, herodot.ErrBadRequest.WithReasonf("relation %q does not exist", relation)
}
