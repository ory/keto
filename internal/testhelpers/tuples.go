package testhelpers

import (
	"testing"

	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/require"

	"github.com/ory/keto/internal/relationtuple"
	"github.com/ory/keto/ketoapi"
)

// TupleFactory parses human-readable tuples like "File:x#view@User:u" into
// internal relation tuples. It adds a randomized suffix to namespaces and maps
// object/subject IDs to UUIDs unique per factory.
type TupleFactory struct {
	suffix string
	seed   uuid.UUID
}

func NewTupleFactory() *TupleFactory {
	seed := uuid.Must(uuid.NewV4())
	return &TupleFactory{suffix: "_" + seed.String(), seed: seed}
}

// NS returns the namespace name with the factory's isolation suffix applied.
func (f *TupleFactory) NS(namespace string) string { return namespace + f.suffix }

// UUID maps an object or subject ID to a UUID unique to this factory.
func (f *TupleFactory) UUID(id string) uuid.UUID { return uuid.NewV5(f.seed, id) }

func (f *TupleFactory) Tuple(t testing.TB, s string) *relationtuple.RelationTuple {
	t.Helper()
	rt := APITupleFromString(t, s)
	result := &relationtuple.RelationTuple{
		Namespace: f.NS(rt.Namespace),
		Object:    f.UUID(rt.Object),
		Relation:  rt.Relation,
	}
	switch {
	case rt.SubjectID != nil:
		result.Subject = &relationtuple.SubjectID{ID: f.UUID(*rt.SubjectID)}
	case rt.SubjectSet != nil:
		result.Subject = &relationtuple.SubjectSet{
			Namespace: f.NS(rt.SubjectSet.Namespace),
			Object:    f.UUID(rt.SubjectSet.Object),
			Relation:  rt.SubjectSet.Relation,
		}
	default:
		t.Fatalf("tuple %q has no subject", s)
	}
	return result
}

func APITupleFromString(t testing.TB, s string) *ketoapi.RelationTuple {
	t.Helper()
	rt, err := (&ketoapi.RelationTuple{}).FromString(s)
	require.NoError(t, err)
	return rt
}

func TupleFromString(t testing.TB, s string) *relationtuple.RelationTuple {
	t.Helper()
	rt := APITupleFromString(t, s)
	result := &relationtuple.RelationTuple{
		Namespace: rt.Namespace,
		Object:    toUUID(rt.Object),
		Relation:  rt.Relation,
	}
	switch {
	case rt.SubjectID != nil:
		result.Subject = &relationtuple.SubjectID{ID: toUUID(*rt.SubjectID)}
	case rt.SubjectSet != nil:
		result.Subject = &relationtuple.SubjectSet{
			Namespace: rt.SubjectSet.Namespace,
			Object:    toUUID(rt.SubjectSet.Object),
			Relation:  rt.SubjectSet.Relation,
		}
	default:
		t.Fatal("invalid tuple")
	}
	return result
}

func SubjectFromString(t testing.TB, s string) relationtuple.Subject {
	t.Helper()
	rt, err := (&ketoapi.RelationTuple{}).FromString("	:object#relation@" + s)
	require.NoError(t, err)

	switch {
	case rt.SubjectID != nil:
		return &relationtuple.SubjectID{ID: toUUID(*rt.SubjectID)}
	case rt.SubjectSet != nil:
		return &relationtuple.SubjectSet{
			Namespace: rt.SubjectSet.Namespace,
			Object:    toUUID(rt.SubjectSet.Object),
			Relation:  rt.SubjectSet.Relation,
		}
	default:
		t.Fatal("invalid subject")
		return nil
	}
}

func SubjectSetFromString(t testing.TB, s string) *relationtuple.SubjectSet {
	t.Helper()
	subject := SubjectFromString(t, s)
	subjectSet, ok := subject.(*relationtuple.SubjectSet)
	require.True(t, ok, "expected subject to be a SubjectSet")
	return subjectSet
}

type deps interface {
	relationtuple.MapperProvider
	relationtuple.ManagerProvider
}

func MapAndInsertTuplesFromString(t testing.TB, d deps, tuples []string) {
	t.Helper()
	relationTuples := make([]*relationtuple.RelationTuple, len(tuples))
	for i, tuple := range tuples {
		relationTuples[i] = TupleFromString(t, tuple)
	}
	require.NoError(t, d.RelationTupleManager().WriteRelationTuples(t.Context(), relationTuples...))
}

func MapAndInsertTuples(t testing.TB, d deps, tuples ...*ketoapi.RelationTuple) {
	t.Helper()
	its, err := d.Mapper().FromTuple(t.Context(), tuples...)
	require.NoError(t, err)
	require.NoError(t, d.RelationTupleManager().WriteRelationTuples(t.Context(), its...))
}

func toUUID(s string) uuid.UUID {
	return uuid.NewV5(uuid.Nil, s)
}
