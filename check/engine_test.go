package check_test

import (
	"context"
	"testing"

	"github.com/ory/keto/relationtuple"

	"github.com/ory/keto/check"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ory/keto/driver"
)

func TestEngine(t *testing.T) {
	t.Run("direct inclusion", func(t *testing.T) {
		rel := relationtuple.InternalRelationTuple{
			Relation: "access",
			Object: &relationtuple.Object{
				ID:        "object",
				Namespace: "test",
			},
			Subject: &relationtuple.UserID{ID: "user"},
		}

		reg := &driver.RegistryDefault{}
		require.NoError(t, reg.RelationTupleManager().WriteRelationTuples(context.Background(), &rel))

		e := check.NewEngine(reg)

		res, err := e.SubjectIsAllowed(context.Background(), &rel)
		require.NoError(t, err)
		assert.True(t, res)
	})

	t.Run("indirect inclusion level 1", func(t *testing.T) {
		// the set of users that are produces of "dust" have to remove it
		dust := relationtuple.Object{
			ID:        "dust",
			Namespace: "under the sofa",
		}
		mark := relationtuple.UserID{
			ID: "Mark",
		}
		cleaningRelation := relationtuple.InternalRelationTuple{
			Relation: "have to remove",
			Object:   &dust,
			Subject: &relationtuple.UserSet{
				Relation: "producer",
				Object:   &dust,
			},
		}
		markProducesDust := relationtuple.InternalRelationTuple{
			Relation: "producer",
			Object:   &dust,
			Subject:  &mark,
		}

		reg := &driver.RegistryDefault{}
		require.NoError(t, reg.RelationTupleManager().WriteRelationTuples(context.Background(), &cleaningRelation, &markProducesDust))

		e := check.NewEngine(reg)

		res, err := e.SubjectIsAllowed(context.Background(), &relationtuple.InternalRelationTuple{
			Relation: cleaningRelation.Relation,
			Object:   &dust,
			Subject:  &mark,
		})
		require.NoError(t, err)
		assert.True(t, res)
	})

	t.Run("direct exclusion", func(t *testing.T) {
		user := &relationtuple.UserID{
			ID: "user-id",
		}
		rel := relationtuple.InternalRelationTuple{
			Relation: "relation",
			Object: &relationtuple.Object{
				ID:        "object-id",
				Namespace: "object-namespace",
			},
			Subject: user,
		}

		reg := &driver.RegistryDefault{}
		require.NoError(t, reg.RelationTupleManager().WriteRelationTuples(context.Background(), &rel))

		e := check.NewEngine(reg)

		res, err := e.SubjectIsAllowed(context.Background(), &relationtuple.InternalRelationTuple{
			Relation: rel.Relation,
			Object:   rel.Object,
			Subject:  &relationtuple.UserID{ID: "not " + user.ID},
		})
		require.NoError(t, err)
		assert.False(t, res)
	})

	t.Run("wrong object ID", func(t *testing.T) {
		object := relationtuple.Object{
			ID: "object",
		}
		access := relationtuple.InternalRelationTuple{
			Relation: "access",
			Object:   &object,
			Subject: &relationtuple.UserSet{
				Relation: "owner",
				Object:   &object,
			},
		}
		user := relationtuple.InternalRelationTuple{
			Relation: "owner",
			Object:   &relationtuple.Object{ID: "not " + object.ID},
			Subject:  &relationtuple.UserID{ID: "user"},
		}

		reg := &driver.RegistryDefault{}
		require.NoError(t, reg.RelationTupleManager().WriteRelationTuples(context.Background(), &access, &user))

		e := check.NewEngine(reg)

		res, err := e.SubjectIsAllowed(context.Background(), &relationtuple.InternalRelationTuple{
			Relation: access.Relation,
			Object:   &object,
			Subject:  user.Subject,
		})
		require.NoError(t, err)
		assert.False(t, res)
	})

	t.Run("wrong relation name", func(t *testing.T) {
		diaryEntry := &relationtuple.Object{
			ID:        "entry for 6. Nov 2020",
			Namespace: "diary",
		}
		// this would be a userset rewrite
		readDiary := relationtuple.InternalRelationTuple{
			Relation: "read",
			Object:   diaryEntry,
			Subject: &relationtuple.UserSet{
				Relation: "author",
				Object:   diaryEntry,
			},
		}
		user := relationtuple.InternalRelationTuple{
			Relation: "not author",
			Object:   diaryEntry,
			Subject:  &relationtuple.UserID{ID: "your mother"},
		}

		reg := &driver.RegistryDefault{}
		require.NoError(t, reg.RelationTupleManager().WriteRelationTuples(context.Background(), &readDiary, &user))

		e := check.NewEngine(reg)

		res, err := e.SubjectIsAllowed(context.Background(), &relationtuple.InternalRelationTuple{
			Relation: readDiary.Relation,
			Object:   diaryEntry,
			Subject:  user.Subject,
		})
		require.NoError(t, err)
		assert.False(t, res)
	})

	t.Run("indirect inclusion level 2", func(t *testing.T) {
		object := relationtuple.Object{
			ID:        "some object",
			Namespace: "some namespace",
		}
		user := relationtuple.UserID{
			ID: "some user",
		}
		organization := relationtuple.Object{
			ID:        "some organization",
			Namespace: "all organizations",
		}

		ownerUserSet := relationtuple.UserSet{
			Relation: "owner",
			Object:   &object,
		}
		orgMembers := relationtuple.UserSet{
			Relation: "member",
			Object:   &organization,
		}

		writeRel := relationtuple.InternalRelationTuple{
			Relation: "write",
			Object:   &object,
			Subject:  &ownerUserSet,
		}
		orgOwnerRel := relationtuple.InternalRelationTuple{
			Relation: ownerUserSet.Relation,
			Object:   &object,
			Subject:  &orgMembers,
		}
		userMembershipRel := relationtuple.InternalRelationTuple{
			Relation: orgMembers.Relation,
			Object:   orgMembers.Object,
			Subject:  &user,
		}

		reg := &driver.RegistryDefault{}
		require.NoError(t, reg.RelationTupleManager().WriteRelationTuples(context.Background(), &writeRel, &orgOwnerRel, &userMembershipRel))

		e := check.NewEngine(reg)

		// user can write object
		res, err := e.SubjectIsAllowed(context.Background(), &relationtuple.InternalRelationTuple{
			Relation: writeRel.Relation,
			Object:   &object,
			Subject:  &user,
		})
		require.NoError(t, err)
		assert.True(t, res)

		// user is member of the organization
		res, err = e.SubjectIsAllowed(context.Background(), &relationtuple.InternalRelationTuple{
			Relation: orgMembers.Relation,
			Object:   &organization,
			Subject:  &user,
		})
		require.NoError(t, err)
		assert.True(t, res)
	})

	t.Run("rejects transitive relation", func(t *testing.T) {
		// (file) <--parent-- (directory) <--access-- [user]
		//
		// note the missing access relation from "users who have access to directory also have access to files inside of the directory"
		// as we don't know how to interpret the "parent" relation, there would have to be a userset rewrite to allow access
		// to files when you have access to the parent

		file := relationtuple.Object{ID: "file"}
		directory := relationtuple.Object{ID: "directory"}
		user := relationtuple.UserID{ID: "user"}

		parent := relationtuple.InternalRelationTuple{
			Relation: "parent",
			Object:   &file,
			Subject: &relationtuple.UserSet{ // <- this is only an object, but this is allowed as a userset can have the "..." relation which means any relation
				Object: &directory,
			},
		}
		directoryAccess := relationtuple.InternalRelationTuple{
			Relation: "access",
			Object:   &directory,
			Subject:  &user,
		}

		reg := &driver.RegistryDefault{}
		for _, r := range []*relationtuple.InternalRelationTuple{&parent, &directoryAccess} {
			require.NoError(t, reg.RelationTupleManager().WriteRelationTuples(context.Background(), r))
		}

		e := check.NewEngine(reg)

		res, err := e.SubjectIsAllowed(context.Background(), &relationtuple.InternalRelationTuple{
			Relation: directoryAccess.Relation,
			Object:   &file,
			Subject:  &user,
		})
		require.NoError(t, err)
		assert.False(t, res)
	})
}
