package memory

import (
	"context"

	"github.com/ory/keto/relation"

	"github.com/ory/keto/models"
)

var _ relation.Manager = &Persister{}

func (p *Persister) paginateRelations(rels []*models.Relation, options ...relation.PaginationOptionSetter) []*models.Relation {
	if len(rels) == 0 {
		return rels
	}

	pagination := relation.GetPaginationOptions(options...)
	veryLast := len(rels)
	start, end := pagination.Page*pagination.PerPage, (pagination.Page+1)*pagination.PerPage-1
	if veryLast < end {
		end = veryLast
	}
	return rels[start:end]
}

func (p *Persister) findRelations(filter func(*models.Relation) bool) (res []*models.Relation) {
	for _, r := range p.relations {
		if filter(r) {
			res = append(res, r.Copy())
		}
	}
	return
}

func (p *Persister) GetRelationsBySubject(_ context.Context, subjectID string, options ...relation.PaginationOptionSetter) ([]*models.Relation, error) {
	p.RLock()
	defer p.RUnlock()

	return p.paginateRelations(
		p.findRelations(
			func(r *models.Relation) bool {
				return r.SubjectID == subjectID
			},
		),
		options...,
	), nil
}

func (p *Persister) GetRelationsByObject(_ context.Context, objectID string, options ...relation.PaginationOptionSetter) ([]*models.Relation, error) {
	p.RLock()
	defer p.RUnlock()

	return p.paginateRelations(
		p.findRelations(
			func(r *models.Relation) bool {
				return r.ObjectID == objectID
			},
		),
		options...,
	), nil
}

func (p *Persister) GetRelationsByObjectAndName(ctx context.Context, objectID, name string, options ...relation.PaginationOptionSetter) ([]*models.Relation, error) {
	p.RLock()
	defer p.RUnlock()

	return p.paginateRelations(
		p.findRelations(
			func(r *models.Relation) bool {
				return r.ObjectID == objectID && r.Name == name
			},
		),
		options...,
	), nil
}

func (p *Persister) WriteRelation(_ context.Context, r *models.Relation) error {
	p.Lock()
	defer p.Unlock()

	p.relations = append(p.relations, r.Copy())
	return nil
}
