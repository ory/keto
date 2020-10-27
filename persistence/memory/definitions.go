package memory

import (
	"sync"

	"github.com/ory/keto/models"
)

type Persister struct {
	sync.RWMutex

	relations []*models.InternalRelationTuple
}

func NewPersister() *Persister {
	return &Persister{
		relations: []*models.InternalRelationTuple{},
	}
}
