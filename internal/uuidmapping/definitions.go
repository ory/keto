package uuidmapping

import (
	"context"
	"github.com/gofrs/uuid"
	"github.com/ory/keto/internal/namespace"
	"github.com/ory/keto/internal/relationtuple"
	"github.com/ory/keto/ketoapi"
)

type (
	dependencies interface {
		ManagerProvider
		namespace.ManagerProvider
	}
	ManagerProvider interface {
		GetUUIDMappingManager() Manager
	}
	Manager interface {
		MapStringsToUUIDs(ctx context.Context, s ...string) ([]uuid.UUID, error)
		MapUUIDsToStrings(ctx context.Context, u ...uuid.UUID) ([]string, error)
	}
	MapperProvider interface {
		GetUUIDMapper() Mapper
	}
	Mapper struct {
		d dependencies
	}
)

func ptr[T any](v T) *T {
	return &v
}

func (m *Mapper) FromQuery(ctx context.Context, q ketoapi.RelationQuery) (*relationtuple.RelationQuery, error) {
	var res relationtuple.RelationQuery
	var s []string
	var u []uuid.UUID

	nm, err := m.d.NamespaceManager()
	if err != nil {
		return nil, err
	}

	if q.Namespace != nil {
		n, err := nm.GetNamespaceByName(ctx, *q.Namespace)
		if err != nil {
			return nil, err
		}
		res.Namespace = ptr(n.ID)
	}
	if q.Object != nil {
		s = append(s, *q.Object)
		defer func() {
			res.Object = ptr(u[0])
		}()
	}
	if q.SubjectID != nil {
		s = append(s, *q.SubjectID)
		defer func() {
			res.SubjectID = ptr(u[1])
		}()
	}
	if q.SubjectSet != nil {
		s = append(s, q.SubjectSet.Object)
		n, err := nm.GetNamespaceByName(ctx, q.SubjectSet.Namespace)
		if err != nil {
			return nil, err
		}
		defer func() {
			res.SubjectSet = &relationtuple.SubjectSet{
				Namespace: n.ID,
				Object:    u[1],
				Relation:  q.SubjectSet.Relation,
			}
		}()
	}

	u, err = m.d.GetUUIDMappingManager().MapStringsToUUIDs(ctx, s...)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
