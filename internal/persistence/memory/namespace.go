package memory

import (
	"github.com/ory/keto/internal/namespace"
	"github.com/ory/keto/internal/persistence"
)

var _ namespace.Manager = &Persister{}

func (p *Persister) MigrateNamespaceUp(n *namespace.Namespace) error {
	p.Lock()
	defer p.Unlock()

	currStatus, ok := p.namespacesStatus[n.ID]
	if !ok {
		currStatus = &namespace.Status{}
		p.namespacesStatus[n.ID] = currStatus
		nc := *n
		p.namespaces[n.Name] = &nc
	}

	if currStatus.CurrentVersion < mostRecentNamespaceVersion {
		currStatus.CurrentVersion = mostRecentNamespaceVersion
	}

	return nil
}

func (p *Persister) NamespaceStatus(n *namespace.Namespace) (*namespace.Status, error) {
	p.RLock()
	defer p.RUnlock()

	s, ok := p.namespacesStatus[n.ID]
	if !ok {
		return nil, persistence.ErrNamespaceUnknown
	}

	sc := *s
	sc.NextVersion = mostRecentNamespaceVersion
	return &sc, nil
}
