package persistence

import (
	"github.com/ory/keto/internal/relationtuple"
)

type Persister interface {
	relationtuple.Manager
}
