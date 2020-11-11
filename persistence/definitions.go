package persistence

import (
	"context"
	"github.com/ory/keto/relationtuple"
)

type Persister interface {
	relationtuple.Manager

	MigrateUp(ctx context.Context) error
}
