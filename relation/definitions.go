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
		GetRelations(ctx context.Context, queries []*models.RelationQuery, page, perPage int32) ([]*models.Relation, error)
		WriteRelation(ctx context.Context, r *models.Relation) error
	}
)
