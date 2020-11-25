package memory

import (
	"sync"

	"github.com/ory/keto/internal/namespace"

	"github.com/ory/keto/internal/relationtuple"
)

const mostRecentNamespaceVersion = 1

type Persister struct {
	sync.RWMutex

	relations map[int][]*relationtuple.InternalRelationTuple

	namespacesStatus map[int]*namespace.Status
	namespaces       map[string]*namespace.Namespace
}

func NewPersister() *Persister {
	return &Persister{
		relations:        make(map[int][]*relationtuple.InternalRelationTuple),
		namespacesStatus: make(map[int]*namespace.Status),
		namespaces:       make(map[string]*namespace.Namespace),
	}
}
