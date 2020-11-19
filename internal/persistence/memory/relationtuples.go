package memory

import (
	"context"

	"github.com/ory/keto/internal/x"

	"github.com/ory/keto/internal/relationtuple"
)

type (
	queryFilter func(r *relationtuple.InternalRelationTuple) bool
)

var _ relationtuple.Manager = &Persister{}

func (p *Persister) paginateRelations(rels []*relationtuple.InternalRelationTuple, options ...x.PaginationOptionSetter) []*relationtuple.InternalRelationTuple {
	if len(rels) == 0 {
		return rels
	}

	pagination := x.GetPaginationOptions(options...)
	veryLast := len(rels)
	start, end := pagination.Page*pagination.PerPage, (pagination.Page+1)*pagination.PerPage-1
	if veryLast < end {
		end = veryLast
	}
	return rels[start:end]
}

func buildRelationQueryFilter(query *relationtuple.RelationQuery) queryFilter {
	var filters []queryFilter

	if query.ObjectID != "" {
		filters = append(filters, func(r *relationtuple.InternalRelationTuple) bool {
			return query.ObjectID == r.ObjectID
		})
	}

	if query.Relation != "" {
		filters = append(filters, func(r *relationtuple.InternalRelationTuple) bool {
			return r.Relation == query.Relation
		})
	}

	if query.Subject != nil {
		filters = append(filters, func(r *relationtuple.InternalRelationTuple) bool {
			return query.Subject.Equals(r.Subject)
		})
	}

	// Create composite filter
	return func(r *relationtuple.InternalRelationTuple) bool {
		// this is lazy-evaluating the AND of all filters
		for _, filter := range filters {
			if !filter(r) {
				return false
			}
		}
		return true
	}
}

func (p *Persister) GetRelationTuples(_ context.Context, query *relationtuple.RelationQuery, options ...x.PaginationOptionSetter) ([]*relationtuple.InternalRelationTuple, error) {
	p.RLock()
	defer p.RUnlock()

	if query == nil {
		return nil, nil
	}

	filter := buildRelationQueryFilter(query)

	n, ok := p.namespaces[query.Namespace]
	if !ok {
		return nil, ErrNamespaceUnknown
	}

	var res []*relationtuple.InternalRelationTuple
	for _, r := range p.relations[n.ID] {
		if filter(r) {
			// If one filter matches add relation to response
			res = append(res, r)
		}
	}

	return p.paginateRelations(res, options...), nil
}

func (p *Persister) WriteRelationTuples(_ context.Context, rs ...*relationtuple.InternalRelationTuple) error {
	p.Lock()
	defer p.Unlock()

	for _, r := range rs {
		n, ok := p.namespaces[r.Namespace]
		if !ok {
			return ErrNamespaceUnknown
		}

		p.relations[n.ID] = append(p.relations[n.ID], r)
	}
	return nil
}
