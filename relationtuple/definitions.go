package relationtuple

import (
	"context"

	"github.com/ory/keto/models"
)

type (
	ManagerProvider interface {
		RelationTupleManager() Manager
	}
	Manager interface {
		GetRelationTuples(ctx context.Context, queries []*models.RelationQuery, options ...PaginationOptionSetter) ([]*models.InternalRelationTuple, error)
		WriteRelationTuples(ctx context.Context, rs ...*models.InternalRelationTuple) error
	}
	paginationOptions struct {
		Page, PerPage int
	}
	PaginationOptionSetter func(*paginationOptions) *paginationOptions
	PaginatedRelations     interface {
		Relations() []*models.InternalRelationTuple
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
