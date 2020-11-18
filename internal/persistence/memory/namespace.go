package memory

import "github.com/ory/keto/internal/namespace"

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

	if currStatus.Version < mostRecentNamespaceVersion {
		currStatus.Version = mostRecentNamespaceVersion
	}

	return nil
}

func (p *Persister) NamespaceStatus(n *namespace.Namespace) (*namespace.Status, error) {
	p.RLock()
	defer p.RUnlock()

	s, ok := p.namespacesStatus[n.ID]
	if !ok {
		return nil, ErrNamespaceUnknown
	}

	sc := *s
	return &sc, nil
}
