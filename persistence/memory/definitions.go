package memory

import (
	"github.com/ory/keto/relation"
	"sync"
)

type Persister struct {
	sync.RWMutex

	relations []relation.Relation
}

var _ relation.Manager = &Persister{}

func NewPersister() *Persister {
	return &Persister{
		relations: []relation.Relation{
			{
				UserID:   "1",
				Name:     "testRelation",
				ObjectID: "2",
			},
		},
	}
}
