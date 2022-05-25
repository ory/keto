package sql

import (
	"context"

	"github.com/ory/keto/internal/namespace"
)

func (p *Persister) GetNamespaceByName(ctx context.Context, name string) (*namespace.Namespace, error) {
	ctx, span := p.d.Tracer(ctx).Tracer().Start(ctx, "persistence.sql.GetNamespaceByName")
	defer span.End()

	nm, err := p.d.Config(ctx).NamespaceManager()
	if err != nil {
		return nil, err
	}
	return nm.GetNamespaceByName(ctx, name)
}

func (p *Persister) GetNamespaceByID(ctx context.Context, id int32) (*namespace.Namespace, error) {
	ctx, span := p.d.Tracer(ctx).Tracer().Start(ctx, "persistence.sql.GetNamespaceByID")
	defer span.End()

	nm, err := p.d.Config(ctx).NamespaceManager()
	if err != nil {
		return nil, err
	}
	return nm.GetNamespaceByConfigID(ctx, id)
}
