package persistence

import (
	"context"
	"errors"

	"github.com/ory/x/popx"

	"github.com/gobuffalo/pop/v6"

	"github.com/ory/keto/internal/relationtuple"
)

type (
	Persister interface {
		relationtuple.Manager
		relationtuple.UUIDMappingManager

		Connection(ctx context.Context) *pop.Connection
	}
	Migrator interface {
		MigrationBox(ctx context.Context) (*popx.MigrationBox, error)
		MigrateUp(ctx context.Context) error
		MigrateDown(ctx context.Context) error
	}
	Provider interface {
		Persister() Persister
	}
)

var (
	ErrNamespaceUnknown         = errors.New("namespace unknown")
	ErrMalformedPageToken       = errors.New("malformed page token")
	ErrNetworkMigrationsMissing = errors.New("networkx migrations are not yet applied")
)
