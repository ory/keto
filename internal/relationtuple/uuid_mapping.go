package relationtuple

import (
	"context"
	"testing"

	"github.com/ory/keto/internal/driver/config"
	"github.com/ory/keto/internal/x"
	"github.com/ory/keto/ketoapi"

	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type (
	mapperDependencies interface {
		MappingManagerProvider
		config.Provider
	}
	MappingManagerProvider interface {
		MappingManager() MappingManager
	}
	MappingManager interface {
		MapStringsToUUIDs(ctx context.Context, s ...string) ([]uuid.UUID, error)
		MapUUIDsToStrings(ctx context.Context, u ...uuid.UUID) ([]string, error)
	}
	MapperProvider interface {
		Mapper() *Mapper
	}
	Mapper struct {
		D mapperDependencies
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

	nm, err := m.D.Config(ctx).NamespaceManager()
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
		onSuccess.do(func(i int) func() {
			return func() {
				res.Object = x.Ptr(u[i])
			}
		}(len(s) - 1))
	}
	if q.SubjectID != nil {
		s = append(s, *q.SubjectID)
		onSuccess.do(func(i int) func() {
			return func() {
				res.Subject = &SubjectID{u[i]}
			}
		}(len(s) - 1))
	}
	if q.SubjectSet != nil {
		s = append(s, q.SubjectSet.Object)
		n, err := nm.GetNamespaceByName(ctx, q.SubjectSet.Namespace)
		if err != nil {
			return nil, err
		}
		onSuccess.do(func(i int) func() {
			return func() {
				res.Subject = &SubjectSet{
					Namespace: n.ID,
					Object:    u[i],
					Relation:  q.SubjectSet.Relation,
				}
			}
		}(len(s) - 1))
	}

	u, err = m.D.MappingManager().MapStringsToUUIDs(ctx, s...)
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

	nm, err := m.D.Config(ctx).NamespaceManager()
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
	if q.Subject != nil {
		switch sub := q.Subject.(type) {
		case *SubjectID:
			u = append(u, sub.ID)
			onSuccess.do(func() {
				res.SubjectID = x.Ptr(s[1])
			})
		case *SubjectSet:
			u = append(u, sub.Object)
			n, err := nm.GetNamespaceByConfigID(ctx, sub.Namespace)
			if err != nil {
				return nil, err
			}
			onSuccess.do(func() {
				res.SubjectSet = &ketoapi.SubjectSet{
					Namespace: n.Name,
					Object:    s[1],
					Relation:  sub.Relation,
				}
			})
		}
	}

	s, err = m.D.MappingManager().MapUUIDsToStrings(ctx, u...)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (m *Mapper) FromTuple(ctx context.Context, ts ...*ketoapi.RelationTuple) (res []*RelationTuple, err error) {
	onSuccess := newSuccess(&err)
	defer onSuccess.cleanup()

	res = make([]*RelationTuple, len(ts))
	s := make([]string, len(ts)*2)
	u := make([]uuid.UUID, len(ts)*2)

	nm, err := m.D.Config(ctx).NamespaceManager()
	if err != nil {
		return nil, err
	}

	for i, t := range ts {
		i, t := i, t
		n, err := nm.GetNamespaceByName(ctx, t.Namespace)
		if err != nil {
			return nil, err
		}
		res[i] = &RelationTuple{
			Namespace: n.ID,
			Relation:  t.Relation,
		}
		s[i*2] = t.Object
		onSuccess.do(func() {
			res[i].Object = u[i*2]
		})
		if t.SubjectID != nil {
			s[i*2+1] = *t.SubjectID
			onSuccess.do(func() {
				res[i].Subject = &SubjectID{u[i*2+1]}
			})
		} else if t.SubjectSet != nil {
			s[i*2+1] = t.SubjectSet.Object
			n, err := nm.GetNamespaceByName(ctx, t.SubjectSet.Namespace)
			if err != nil {
				return nil, err
			}
			onSuccess.do(func() {
				res[i].Subject = &SubjectSet{
					Namespace: n.ID,
					Object:    u[i*2+1],
					Relation:  t.SubjectSet.Relation,
				}
			})
		}
	}

	u, err = m.D.MappingManager().MapStringsToUUIDs(ctx, s...)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (m *Mapper) ToTuple(ctx context.Context, ts ...*RelationTuple) (res []*ketoapi.RelationTuple, err error) {
	onSuccess := newSuccess(&err)
	defer onSuccess.cleanup()

	res = make([]*ketoapi.RelationTuple, len(ts))
	s := make([]string, 0, len(ts)*2)
	u := make([]uuid.UUID, len(ts)*2)

	nm, err := m.D.Config(ctx).NamespaceManager()
	if err != nil {
		return nil, err
	}

	for i, t := range ts {
		i := i
		n, err := nm.GetNamespaceByConfigID(ctx, t.Namespace)
		if err != nil {
			return nil, err
		}
		res[i] = &ketoapi.RelationTuple{
			Namespace: n.Name,
			Relation:  t.Relation,
		}
		u[2*i] = t.Object
		onSuccess.do(func() {
			res[i].Object = s[2*i]
		})
		switch sub := t.Subject.(type) {
		case *SubjectID:
			u[2*i+1] = sub.ID
			onSuccess.do(func() {
				res[i].SubjectID = x.Ptr(s[2*i+1])
			})
		case *SubjectSet:
			u[2*i+1] = sub.Object
			n, err := nm.GetNamespaceByConfigID(ctx, sub.Namespace)
			if err != nil {
				return nil, err
			}
			onSuccess.do(func() {
				res[i].SubjectSet = &ketoapi.SubjectSet{
					Namespace: n.Name,
					Object:    s[2*i+1],
					Relation:  sub.Relation,
				}
			})
		}
	}

	s, err = m.D.MappingManager().MapUUIDsToStrings(ctx, u...)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (m *Mapper) ToSubjectSet(ctx context.Context, set *ketoapi.SubjectSet) (*SubjectSet, error) {
	nm, err := m.D.Config(ctx).NamespaceManager()
	if err != nil {
		return nil, err
	}
	n, err := nm.GetNamespaceByName(ctx, set.Namespace)
	if err != nil {
		return nil, err
	}
	u, err := m.D.MappingManager().MapStringsToUUIDs(ctx, set.Object)
	if err != nil {
		return nil, err
	}
	return &SubjectSet{
		Namespace: n.ID,
		Object:    u[0],
		Relation:  set.Relation,
	}, nil
}

func (m *Mapper) FromTree(ctx context.Context, tree *Tree) (res *ketoapi.ExpandTree, err error) {
	onSuccess := newSuccess(&err)
	defer onSuccess.cleanup()

	var s []string
	var u []uuid.UUID
	res = &ketoapi.ExpandTree{
		Type: tree.Type,
	}

	nm, err := m.D.Config(ctx).NamespaceManager()
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
		onSuccess.do(func() {
			res.SubjectID = x.Ptr(s[0])
		})
	}
	for _, c := range tree.Children {
		mc, err := m.FromTree(ctx, c)
		if err != nil {
			return nil, err
		}
		res.Children = append(res.Children, mc)
	}
	s, err = m.D.MappingManager().MapUUIDsToStrings(ctx, u...)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func MappingManagerTest(t *testing.T, m MappingManager) {
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
