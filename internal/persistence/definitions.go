package persistence

import (
	"context"
	"errors"

	"github.com/ory/keto/internal/namespace"

	"github.com/ory/keto/internal/relationtuple"
)

type Persister interface {
	relationtuple.Manager
	namespace.Manager
}

type Migrator interface {
	MigrateUp(ctx context.Context) error
}

var (
	ErrNamespaceUnknown   = errors.New("namespace unknown")
	ErrMalformedPageToken = errors.New("malformed page token")
)
