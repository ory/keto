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

func (m *Mapper) FromQuery(ctx context.Context, apiQuery *ketoapi.RelationQuery) (res *RelationQuery, err error) {
	onSuccess := newSuccess(&err)
	defer onSuccess.cleanup()

	var s []string
	var u []uuid.UUID
	res = &RelationQuery{
		Relation: apiQuery.Relation,
	}

	nm, err := m.D.Config(ctx).NamespaceManager()
	if err != nil {
		return nil, err
	}

	if apiQuery.Namespace != nil {
		n, err := nm.GetNamespaceByName(ctx, *apiQuery.Namespace)
		if err != nil {
			return nil, err
		}
		res.Namespace = x.Ptr(n.Name)
	}
	if apiQuery.Object != nil {
		s = append(s, *apiQuery.Object)
		onSuccess.do(func(i int) func() {
			return func() {
				res.Object = x.Ptr(u[i])
			}
		}(len(s) - 1))
	}
	if apiQuery.SubjectID != nil {
		s = append(s, *apiQuery.SubjectID)
		onSuccess.do(func(i int) func() {
			return func() {
				res.Subject = &SubjectID{u[i]}
			}
		}(len(s) - 1))
	}
	if apiQuery.SubjectSet != nil {
		s = append(s, apiQuery.SubjectSet.Object)
		n, err := nm.GetNamespaceByName(ctx, apiQuery.SubjectSet.Namespace)
		if err != nil {
			return nil, err
		}
		onSuccess.do(func(i int) func() {
			return func() {
				res.Subject = &SubjectSet{
					Namespace: n.Name,
					Object:    u[i],
					Relation:  apiQuery.SubjectSet.Relation,
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
	res = &ketoapi.RelationQuery{
		Relation: q.Relation,
	}

	nm, err := m.D.Config(ctx).NamespaceManager()
	if err != nil {
		return nil, err
	}

	if q.Namespace != nil {
		n, err := nm.GetNamespaceByName(ctx, *q.Namespace)
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
				res.SubjectID = x.Ptr(s[len(s)-1])
			})
		case *SubjectSet:
			u = append(u, sub.Object)
			n, err := nm.GetNamespaceByName(ctx, sub.Namespace)
			if err != nil {
				return nil, err
			}
			onSuccess.do(func() {
				res.SubjectSet = &ketoapi.SubjectSet{
					Namespace: n.Name,
					Object:    s[len(s)-1],
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

	res = make([]*RelationTuple, 0, len(ts))
	s := make([]string, 0, len(ts)*2)
	var u []uuid.UUID

	nm, err := m.D.Config(ctx).NamespaceManager()
	if err != nil {
		return nil, err
	}

	for _, t := range ts {
		t := t
		n, err := nm.GetNamespaceByName(ctx, t.Namespace)
		if err != nil {
			return nil, err
		}
		mt := RelationTuple{
			Namespace: n.Name,
			Relation:  t.Relation,
		}
		i := len(res)

		if t.SubjectID != nil {
			s = append(s, *t.SubjectID)
			onSuccess.do(func() {
				mt.Subject = &SubjectID{u[i*2]}
			})
		} else if t.SubjectSet != nil {
			n, err := nm.GetNamespaceByName(ctx, t.SubjectSet.Namespace)
			if err != nil {
				return nil, err
			}
			s = append(s, t.SubjectSet.Object)
			onSuccess.do(func() {
				mt.Subject = &SubjectSet{
					Namespace: n.Name,
					Object:    u[i*2],
					Relation:  t.SubjectSet.Relation,
				}
			})
		}

		s = append(s, t.Object)
		onSuccess.do(func() {
			mt.Object = u[i*2+1]
		})

		res = append(res, &mt)
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

	res = make([]*ketoapi.RelationTuple, 0, len(ts))
	u := make([]uuid.UUID, 0, len(ts)*2)
	var s []string

	for _, t := range ts {
		mt := ketoapi.RelationTuple{
			Namespace: t.Namespace,
			Relation:  t.Relation,
		}
		i := len(res)

		switch sub := t.Subject.(type) {
		case *SubjectID:
			u = append(u, sub.ID)
			onSuccess.do(func() {
				mt.SubjectID = x.Ptr(s[2*i])
			})
		case *SubjectSet:
			u = append(u, sub.Object)
			onSuccess.do(func() {
				mt.SubjectSet = &ketoapi.SubjectSet{
					Namespace: sub.Namespace,
					Object:    s[2*i],
					Relation:  sub.Relation,
				}
			})
		}

		u = append(u, t.Object)
		onSuccess.do(func() {
			mt.Object = s[2*i+1]
		})

		res = append(res, &mt)
	}

	s, err = m.D.MappingManager().MapUUIDsToStrings(ctx, u...)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (m *Mapper) FromSubjectSet(ctx context.Context, set *ketoapi.SubjectSet) (*SubjectSet, error) {
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
		Namespace: n.Name,
		Object:    u[0],
		Relation:  set.Relation,
	}, nil
}

func (m *Mapper) ToTree(ctx context.Context, tree *Tree) (res *ketoapi.ExpandTree, err error) {
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
		n, err := nm.GetNamespaceByName(ctx, sub.Namespace)
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
		mc, err := m.ToTree(ctx, c)
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
