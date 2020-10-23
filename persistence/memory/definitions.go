package memory

import (
	"sync"

	"github.com/ory/keto/models"
)

type Persister struct {
	sync.RWMutex

	relations []*models.Relation
}

func NewPersister() *Persister {
	return &Persister{
		relations: []*models.Relation{
			{
				UserID:   "1",
				Name:     "testRelation",
				ObjectID: "2",
			},
		},
	}
}
