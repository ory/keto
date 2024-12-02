// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package persistence

import (
	"context"
	"errors"

	"github.com/gofrs/uuid"
	"github.com/ory/x/popx"

	"github.com/gobuffalo/pop/v6"

	"github.com/ory/keto/internal/relationtuple"
)

type (
	Persister interface {
		relationtuple.Manager
		relationtuple.MappingManager
		Migrator

		Connection(ctx context.Context) *pop.Connection
		NetworkID(ctx context.Context) uuid.UUID
		SetNetwork(nid uuid.UUID)
		Transaction(ctx context.Context, f func(ctx context.Context) error) error
	}
	Migrator interface {
		popx.MigrationProvider
	}
	Provider interface {
		Persister() Persister
		Traverser() relationtuple.Traverser
	}
)

var (
	ErrMalformedPageToken       = errors.New("malformed page token")
	ErrNetworkMigrationsMissing = errors.New("networkx migrations are not yet applied")
)
