package persistence

import (
	"context"
	"errors"

	"github.com/ory/x/popx"

	"github.com/ory/keto/internal/relationtuple"
)

type (
	Persister interface {
		relationtuple.Manager
	}
	Migrator interface {
		MigrationBox(context.Context) (*popx.MigrationBox, error)
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
