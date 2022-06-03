package relationtuple

import (
	"context"
	"github.com/ory/keto/internal/expand"
	"github.com/ory/keto/internal/namespace"
	"github.com/ory/keto/internal/x"
	"github.com/ory/keto/ketoapi"
	"testing"

	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type (
	// UUIDMappable is an interface for objects that have fields that can be
	// mapped to and from UUIDs.
	UUIDMappable interface{ UUIDMappableFields() []*string }

	UUIDMappingManager interface {
		MapStringsToUUIDs(ctx context.Context, s ...string) ([]uuid.UUID, error)
		MapUUIDsToStrings(ctx context.Context, u ...uuid.UUID) ([]string, error)
	}

	mapperDependencies interface {
		MappingManagerProvider
		namespace.ManagerProvider
	}
	MappingManagerProvider interface {
		UUIDMappingManager() MappingManager
	}
	MappingManager interface {
		MapStringsToUUIDs(ctx context.Context, s ...string) ([]uuid.UUID, error)
		MapUUIDsToStrings(ctx context.Context, u ...uuid.UUID) ([]string, error)
	}
	MapperProvider interface {
		UUIDMapper() *Mapper
	}
	Mapper struct {
		d mapperDependencies
	}
)

type success struct {
	fs  []func()
	err *error
}

func newSuccess(err *error) *success {
	return &success{
		err: err,
	}
}

func (c *success) do(f func()) {
	c.fs = append(c.fs, f)
}

func (c *success) cleanup() {
	if *c.err != nil {
		return
	}
	for _, f := range c.fs {
		f()
	}
}

func (m *Mapper) FromQuery(ctx context.Context, q *ketoapi.RelationQuery) (res *RelationQuery, err error) {
	onSuccess := newSuccess(&err)
	defer onSuccess.cleanup()

	var s []string
	var u []uuid.UUID
	res = new(RelationQuery)

	nm, err := m.d.NamespaceManager()
	if err != nil {
		return nil, err
	}

	if q.Namespace != nil {
		n, err := nm.GetNamespaceByName(ctx, *q.Namespace)
		if err != nil {
			return nil, err
		}
		res.Namespace = x.Ptr(n.ID)
	}
	if q.Object != nil {
		s = append(s, *q.Object)
		onSuccess.do(func() {
			res.Object = x.Ptr(u[0])
		})
	}
	if q.SubjectID != nil {
		s = append(s, *q.SubjectID)
		onSuccess.do(func() {
			res.SubjectID = x.Ptr(u[1])
		})
	}
	if q.SubjectSet != nil {
		s = append(s, q.SubjectSet.Object)
		n, err := nm.GetNamespaceByName(ctx, q.SubjectSet.Namespace)
		if err != nil {
			return nil, err
		}
		onSuccess.do(func() {
			res.SubjectSet = &SubjectSet{
				Namespace: n.ID,
				Object:    u[1],
				Relation:  q.SubjectSet.Relation,
			}
		})
	}

	u, err = m.d.UUIDMappingManager().MapStringsToUUIDs(ctx, s...)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (m *Mapper) ToQuery(ctx context.Context, q *RelationQuery) (res *ketoapi.RelationQuery, err error) {
	onSuccess := newSuccess(&err)
	defer onSuccess.cleanup()

	var s []string
	var u []uuid.UUID
	res = new(ketoapi.RelationQuery)

	nm, err := m.d.NamespaceManager()
	if err != nil {
		return nil, err
	}

	if q.Namespace != nil {
		n, err := nm.GetNamespaceByConfigID(ctx, *q.Namespace)
		if err != nil {
			return nil, err
		}
		res.Namespace = x.Ptr(n.Name)
	}
	if q.Object != nil {
		u = append(u, *q.Object)
		onSuccess.do(func() {
			res.Object = x.Ptr(s[0])
		})
	}
	if q.SubjectID != nil {
		u = append(u, *q.SubjectID)
		onSuccess.do(func() {
			res.SubjectID = x.Ptr(s[1])
		})
	}
	if q.SubjectSet != nil {
		u = append(u, q.SubjectSet.Object)
		n, err := nm.GetNamespaceByConfigID(ctx, q.SubjectSet.Namespace)
		if err != nil {
			return nil, err
		}
		onSuccess.do(func() {
			res.SubjectSet = &ketoapi.SubjectSet{
				Namespace: n.Name,
				Object:    s[1],
				Relation:  q.SubjectSet.Relation,
			}
		})
	}

	s, err = m.d.UUIDMappingManager().MapUUIDsToStrings(ctx, u...)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (m *Mapper) FromTuple(ctx context.Context, ts ...*ketoapi.RelationTuple) (res []*InternalRelationTuple, err error) {
	onSuccess := newSuccess(&err)
	defer onSuccess.cleanup()

	res = make([]*InternalRelationTuple, len(ts))
	s := make([]string, 2)
	u := make([]uuid.UUID, 2)

	nm, err := m.d.NamespaceManager()
	if err != nil {
		return nil, err
	}

	for i, t := range ts {
		i := i
		n, err := nm.GetNamespaceByName(ctx, t.Namespace)
		if err != nil {
			return nil, err
		}
		res[i] = &InternalRelationTuple{
			Namespace: n.ID,
		}
		s[0] = t.Object
		onSuccess.do(func() {
			res[i].Object = u[0]
		})
		if t.SubjectID != nil {
			s[1] = *t.SubjectID
			onSuccess.do(func() {
				res[i].Subject = &SubjectID{u[1]}
			})
		} else if t.SubjectSet != nil {
			s[1] = t.SubjectSet.Object
			n, err := nm.GetNamespaceByName(ctx, t.SubjectSet.Namespace)
			if err != nil {
				return nil, err
			}
			onSuccess.do(func() {
				res[i].Subject = &SubjectSet{
					Namespace: n.ID,
					Object:    u[1],
					Relation:  t.SubjectSet.Relation,
				}
			})
		}
	}

	u, err = m.d.UUIDMappingManager().MapStringsToUUIDs(ctx, s...)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (m *Mapper) ToTuple(ctx context.Context, ts ...*ketoapi.RelationTuple) (res []*InternalRelationTuple, err error) {
	onSuccess := newSuccess(&err)
	defer onSuccess.cleanup()

	res = make([]*InternalRelationTuple, len(ts))
	s := make([]string, len(ts)*2)
	u := make([]uuid.UUID, 0, len(ts)*2)

	nm, err := m.d.NamespaceManager()
	if err != nil {
		return nil, err
	}

	for i, t := range ts {
		i := i
		n, err := nm.GetNamespaceByName(ctx, t.Namespace)
		if err != nil {
			return nil, err
		}
		res[i] = &InternalRelationTuple{
			Namespace: n.ID,
		}
		s[2*i] = t.Object
		onSuccess.do(func() {
			res[i].Object = u[2*i]
		})
		if t.SubjectID != nil {
			s[2*i+1] = *t.SubjectID
			onSuccess.do(func() {
				res[i].Subject = &SubjectID{u[2*i+1]}
			})
		} else if t.SubjectSet != nil {
			s[2*i+1] = t.SubjectSet.Object
			n, err := nm.GetNamespaceByName(ctx, t.SubjectSet.Namespace)
			if err != nil {
				return nil, err
			}
			onSuccess.do(func() {
				res[i].Subject = &SubjectSet{
					Namespace: n.ID,
					Object:    u[2*i+1],
					Relation:  t.SubjectSet.Relation,
				}
			})
		}
	}

	u, err = m.d.UUIDMappingManager().MapStringsToUUIDs(ctx, s...)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (m *Mapper) ToSubjectSet(ctx context.Context, set *ketoapi.SubjectSet) (*SubjectSet, error) {
	nm, err := m.d.NamespaceManager()
	if err != nil {
		return nil, err
	}
	n, err := nm.GetNamespaceByName(ctx, set.Namespace)
	if err != nil {
		return nil, err
	}
	u, err := m.d.UUIDMappingManager().MapStringsToUUIDs(ctx, set.Object)
	if err != nil {
		return nil, err
	}
	return &SubjectSet{
		Namespace: n.ID,
		Object:    u[0],
		Relation:  set.Relation,
	}, nil
}

func (m *Mapper) FromTree(ctx context.Context, tree *expand.Tree) (res *ketoapi.ExpandTree, err error) {
	onSuccess := newSuccess(&err)
	defer onSuccess.cleanup()

	var s []string
	var u []uuid.UUID
	res = &ketoapi.ExpandTree{
		Type: tree.Type,
	}

	nm, err := m.d.NamespaceManager()
	if err != nil {
		return nil, err
	}

	switch sub := tree.Subject.(type) {
	case *SubjectSet:
		u = append(u, sub.Object)
		n, err := nm.GetNamespaceByConfigID(ctx, sub.Namespace)
		if err != nil {
			return nil, err
		}
		onSuccess.do(func() {
			res.SubjectSet = &ketoapi.SubjectSet{
				Namespace: n.Name,
				Object:    s[0],
				Relation:  sub.Relation,
			}
		})
	case *SubjectID:
		u = append(u, sub.ID)
	}
	for _, c := range tree.Children {
		mc, err := m.FromTree(ctx, c)
		if err != nil {
			return nil, err
		}
		res.Children = append(res.Children, mc)
	}
	s, err = m.d.UUIDMappingManager().MapUUIDsToStrings(ctx, u...)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func UUIDMappingManagerTest(t *testing.T, m UUIDMappingManager) {
	ctx := context.Background()

	t.Run("case=str -> uuid -> str", func(t *testing.T) {
		const s = "rep1"
		u, err := m.MapStringsToUUIDs(ctx, s)
		require.NoError(t, err)

		actual, err := m.MapUUIDsToStrings(ctx, u[0])
		require.NoError(t, err)
		require.Len(t, actual, 1)
		assert.Equal(t, s, actual[0])

		t.Run("case=batch", func(t *testing.T) {
			s := []string{"rep1", "rep2", "rep3"}

			u, err := m.MapStringsToUUIDs(ctx, s...)
			require.NoError(t, err)
			require.Len(t, u, len(s))

			assert.NotContains(t, u, uuid.Nil)

			actual, err := m.MapUUIDsToStrings(ctx, u...)
			require.NoError(t, err)
			require.Len(t, actual, len(s))
			assert.Equal(t, s, actual)
		})
	})

	t.Run("case=deterministic MapStringsToUUIDs", func(t *testing.T) {
		const s = "some string"

		u0, err := m.MapStringsToUUIDs(ctx, s)
		require.NoError(t, err)
		u1, err := m.MapStringsToUUIDs(ctx, s)
		require.NoError(t, err)

		assert.Equal(t, u0, u1)
	})
}
