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
		GetRelationsByUser(ctx context.Context, userID string, page, perPage int32) ([]*models.Relation, error)
		GetRelationsByObject(ctx context.Context, objectID string, page, perPage int32) ([]*models.Relation, error)
	}
)
