package persistence

import (
	"context"

	"github.com/ory/keto/relationtuple"
)

type Persister interface {
	relationtuple.Manager
}

type Migrator interface {
	MigrateUp(ctx context.Context) error
}
