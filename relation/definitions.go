package relation

import (
	"context"
)

type (
	ManagerProvider interface {
		RelationManager() Manager
	}
	Manager interface {
		GetRelationsByUser(ctx context.Context, userID string, page, perPage int32) ([]Relation, error)
		GetRelationsByObject(ctx context.Context, objectID string, page, perPage int32) ([]Relation, error)
	}
	Relation struct {
		UserID   string
		Name     string
		ObjectID string
	}
)
