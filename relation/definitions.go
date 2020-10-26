package relation

import (
	"context"

	"github.com/ory/keto/models"
)

type (
	ManagerProvider interface {
		RelationManager() Manager
	}
	Manager interface {
		GetRelationsBySubject(ctx context.Context, subjectID string, options ...PaginationOptionSetter) ([]*models.Relation, error)
		GetRelationsByObject(ctx context.Context, objectID string, options ...PaginationOptionSetter) ([]*models.Relation, error)
		GetRelationsByObjectAndName(ctx context.Context, objectID, name string, options ...PaginationOptionSetter) ([]*models.Relation, error)
		WriteRelation(ctx context.Context, r *models.Relation) error
	}
	paginationOptions struct {
		Page, PerPage int
	}
	PaginationOptionSetter func(*paginationOptions) *paginationOptions
	PaginatedRelations     interface {
		Relations() []*models.Relation
		HasNext() bool
		Next() PaginatedRelations
	}
)

func WithPage(page int) PaginationOptionSetter {
	return func(opts *paginationOptions) *paginationOptions {
		opts.Page = page
		return opts
	}
}

func WithPerPage(perPage int) PaginationOptionSetter {
	return func(opts *paginationOptions) *paginationOptions {
		opts.PerPage = perPage
		return opts
	}
}

func GetPaginationOptions(modifiers ...PaginationOptionSetter) *paginationOptions {
	opts := &paginationOptions{
		Page:    0,
		PerPage: 100,
	}
	for _, f := range modifiers {
		opts = f(opts)
	}
	return opts
}
