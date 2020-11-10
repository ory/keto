package memory

import (
	"sync"

	"github.com/ory/keto/relationtuple"
)

type Persister struct {
	sync.RWMutex

	relations []*relationtuple.InternalRelationTuple
}

func NewPersister() *Persister {
	return &Persister{
		relations: []*relationtuple.InternalRelationTuple{},
	}
}
