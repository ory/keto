package persistence

import (
	"errors"
	"github.com/ory/keto/internal/namespace"
	"github.com/ory/keto/internal/relationtuple"
)

type Persister interface {
	relationtuple.Manager
	namespace.Manager
}

var (
	ErrNamespaceUnknown = errors.New("namespace unknown")
)
