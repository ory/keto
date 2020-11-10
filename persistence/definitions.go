package persistence

import (
	"github.com/ory/keto/relationtuple"
)

type Persister interface {
	relationtuple.Manager
}
