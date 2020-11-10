package memory

import (
	"context"

	"github.com/ory/keto/relationtuple"

	"github.com/ory/keto/models"
)

type (
	queryFilter func(r *models.InternalRelationTuple) bool
)

var _ relationtuple.Manager = &Persister{}

func (p *Persister) paginateRelations(rels []*models.InternalRelationTuple, options ...relationtuple.PaginationOptionSetter) []*models.InternalRelationTuple {
	if len(rels) == 0 {
		return rels
	}

	pagination := relationtuple.GetPaginationOptions(options...)
	veryLast := len(rels)
	start, end := pagination.Page*pagination.PerPage, (pagination.Page+1)*pagination.PerPage-1
	if veryLast < end {
		end = veryLast
	}
	return rels[start:end]
}

func buildRelationQueryFilter(query *models.RelationQuery) queryFilter {
	var filters []queryFilter

	if query.Object != nil {
		filters = append(filters, func(r *models.InternalRelationTuple) bool {
			return query.Object.Equals(r.Object)
		})
	}

	if query.Relation != "" {
		filters = append(filters, func(r *models.InternalRelationTuple) bool {
			return r.Relation == query.Relation
		})
	}

	if query.Subject != nil {
		filters = append(filters, func(r *models.InternalRelationTuple) bool {
			return query.Subject.Equals(r.Subject)
		})
	}

	// Create composite filter
	return func(r *models.InternalRelationTuple) bool {
		// this is lazy-evaluating the AND of all filters
		for _, filter := range filters {
			if !filter(r) {
				return false
			}
		}
		return true
	}
}

func (p *Persister) GetRelationTuples(_ context.Context, queries []*models.RelationQuery, options ...relationtuple.PaginationOptionSetter) ([]*models.InternalRelationTuple, error) {
	p.RLock()
	defer p.RUnlock()

	filters := make([]queryFilter, len(queries))
	for i, q := range queries {
		filters[i] = buildRelationQueryFilter(q)
	}

	var res []*models.InternalRelationTuple
	for _, r := range p.relations {
		for _, filter := range filters {
			// this is lazy-evaluating the OR of all filters
			if filter(r) {
				// If one filter matches add relation to response and break inner loop
				// to check next relation
				res = append(res, r)
				break
			}
		}
	}

	return p.paginateRelations(res, options...), nil
}

func (p *Persister) WriteRelationTuples(_ context.Context, rs ...*models.InternalRelationTuple) error {
	p.Lock()
	defer p.Unlock()

	p.relations = append(p.relations, rs...)
	return nil
}
