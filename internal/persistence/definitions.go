package persistence

import (
	"context"
	"errors"

	"github.com/ory/x/popx"

	"github.com/ory/keto/internal/namespace"

	"github.com/ory/keto/internal/relationtuple"
)

type (
	Persister interface {
		relationtuple.Manager
		namespace.Migrator
	}
	Migrator interface {
		MigrateUp(context.Context) error
		MigrateDown(context.Context, int) error
		MigrationStatus(context.Context) (popx.MigrationStatuses, error)
	}
	MigratorProvider interface {
		Migrator() Migrator
	}
	Provider interface {
		Persister() Persister
	}
)

var (
	ErrNamespaceUnknown   = errors.New("namespace unknown")
	ErrMalformedPageToken = errors.New("malformed page token")
)
