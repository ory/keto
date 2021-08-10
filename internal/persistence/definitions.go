package persistence

import (
	"context"
	"errors"

	"github.com/ory/x/popx"

	"github.com/gobuffalo/pop/v5"

	"github.com/ory/keto/internal/relationtuple"
)

type (
	Persister interface {
		relationtuple.Manager

		Connection(ctx context.Context) *pop.Connection
	}
	Migrator interface {
		MigrationBox() (*popx.MigrationBox, error)
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
