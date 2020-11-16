package persistence

import (
	"context"

	"github.com/ory/keto/internal/relationtuple"
)

type Persister interface {
	relationtuple.Manager
}

type Migrator interface {
	MigrateUp(ctx context.Context) error
}
