package sql

import (
	"context"

	"github.com/ory/keto/internal/namespace"
)

func (p *Persister) GetNamespaceByName(ctx context.Context, name string) (*namespace.Namespace, error) {
	nm, err := p.d.Config().NamespaceManager()
	if err != nil {
		return nil, err
	}
	return nm.GetNamespaceByName(ctx, name)
}

func (p *Persister) GetNamespaceByID(ctx context.Context, id int64) (*namespace.Namespace, error) {
	nm, err := p.d.Config().NamespaceManager()
	if err != nil {
		return nil, err
	}
	return nm.GetNamespaceByConfigID(ctx, id)
}
