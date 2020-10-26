package memory

import (
	"context"

	"github.com/ory/keto/relation"

	"github.com/ory/keto/models"
)

var _ relation.Manager = &Persister{}

func (p *Persister) paginateRelations(rels []*models.Relation, page, perPage int32) []*models.Relation {
	if len(rels) == 0 {
		return rels
	}

	veryLast := int32(len(p.relations)) - 1
	start, end := page*perPage, (page+1)*perPage-1
	if veryLast < end {
		end = veryLast
	}
	return rels[start:end]
}

func BuildRelationQueryFilter(query *models.RelationQuery) func(r *models.Relation) bool {
	var filters []func(r *models.Relation) bool
	if query.Object.ID != "" && query.Object.Namespace != "" {
		filters = append(filters, func(r *models.Relation) bool {
			return r.Object.ID == query.Object.ID &&
				r.Object.Namespace == query.Object.Namespace
		})
	}
	if query.Relation != "" {
		filters = append(filters, func(r *models.Relation) bool {
			return r.Relation == query.Relation
		})
	}
	if query.User != nil {
		switch query.User.(type) {
		case models.UserID:
			filters = append(filters, func(r *models.Relation) bool {
				rUserId := r.User.(models.UserID)
				relationUserId := query.User.(models.UserID)
				return r.User != nil && rUserId.ID == relationUserId.ID
			})
		case models.UserSet:
			filters = append(filters, func(r *models.Relation) bool {
				rUserSet := r.User.(models.UserSet)
				relationUserSet := query.User.(models.UserSet)
				return rUserSet.Object.ID == relationUserSet.Object.ID &&
					rUserSet.Object.Namespace == relationUserSet.Object.Namespace &&
					rUserSet.Relation == relationUserSet.Relation
			})
		}
	}

	// Create composite filter
	return func(r *models.Relation) bool {
		for _, filter := range filters {
			if !filter(r) {
				return false
			}
		}
		return true
	}
}

func (p *Persister) GetRelations(_ context.Context, queries []*models.RelationQuery, page, perPage int32) ([]*models.Relation, error) {
	p.RLock()
	defer p.RUnlock()

	var filters []func(r *models.Relation) bool
	for _, q := range queries {
		filters = append(filters, BuildRelationQueryFilter(q))
	}

	var res []*models.Relation
	for _, r := range p.relations {
		for _, filter := range filters {
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

func (p *Persister) WriteRelation(_ context.Context, r *models.Relation) error {
	p.Lock()
	defer p.Unlock()

	p.relations = append(p.relations, r)
	return nil
}
