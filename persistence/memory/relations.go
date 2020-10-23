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

func (p *Persister) GetRelationsByUser(_ context.Context, userID string, page, perPage int32) ([]*models.Relation, error) {
	p.RLock()
	defer p.RUnlock()

	var res []*models.Relation
	for _, r := range p.relations {
		if r.UserID == userID {
			res = append(res, r)
		}
	}

	return p.paginateRelations(res, page, perPage), nil
}

func (p *Persister) GetRelationsByObject(_ context.Context, objectID string, page, perPage int32) ([]*models.Relation, error) {
	p.RLock()
	defer p.RUnlock()

	var res []*models.Relation
	for _, r := range p.relations {
		if r.ObjectID == objectID {
			res = append(res, r)
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
