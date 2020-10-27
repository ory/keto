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
		GetRelationTuples(ctx context.Context, queries []*models.RelationQuery, page, perPage int32) ([]*models.InternalRelationTuple, error)
		WriteRelationTuple(ctx context.Context, r *models.InternalRelationTuple) error
	}
)
