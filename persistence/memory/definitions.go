package memory

import (
	"github.com/ory/keto/models"
	"github.com/ory/keto/relation"
	"sync"
)

type Persister struct {
	sync.RWMutex

	relations []*models.Relation
}

var _ relation.Manager = &Persister{}

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
