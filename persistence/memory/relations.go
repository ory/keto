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

func (p *Persister) paginateRelations(rels []*models.InternalRelationTuple, page, perPage int32) []*models.InternalRelationTuple {
	if len(rels) == 0 {
		return rels
	}

	veryLast := int32(len(p.relations))
	start, end := page*perPage, (page+1)*perPage-1
	if veryLast < end {
		end = veryLast
	}
	return rels[start:end]
}

func buildRelationQueryFilter(query *models.RelationQuery) queryFilter {
	var filters []queryFilter

	if query.Object.ID != "" && query.Object.Namespace != "" {
		filters = append(filters, func(r *models.InternalRelationTuple) bool {
			return r.Object.ID == query.Object.ID &&
				r.Object.Namespace == query.Object.Namespace
		})
	}

	if query.Relation != "" {
		filters = append(filters, func(r *models.InternalRelationTuple) bool {
			return r.Relation == query.Relation
		})
	}

	if query.Subject != nil {
		switch s := query.Subject.(type) {
		case *models.UserID:
			filters = append(filters, func(r *models.InternalRelationTuple) bool {
				rUserId, ok := r.Subject.(*models.UserID)
				return ok &&
					r.Subject != nil &&
					rUserId.ID == s.ID
			})
		case *models.UserSet:
			filters = append(filters, func(r *models.InternalRelationTuple) bool {
				rUserSet, ok := r.Subject.(*models.UserSet)
				return ok &&
					rUserSet.Object.ID == s.Object.ID &&
					rUserSet.Object.Namespace == s.Object.Namespace &&
					rUserSet.Relation == s.Relation
			})
		}
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

func (p *Persister) GetRelationTuples(_ context.Context, queries []*models.RelationQuery, page, perPage int32) ([]*models.InternalRelationTuple, error) {
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

	return p.paginateRelations(res, page, perPage), nil
}

func (p *Persister) WriteRelationTuple(_ context.Context, r *models.InternalRelationTuple) error {
	p.Lock()
	defer p.Unlock()

	p.relations = append(p.relations, r)
	return nil
}
