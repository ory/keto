package check_test

import (
	"context"
	"testing"

	"github.com/ory/keto/check"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ory/keto/driver"
	"github.com/ory/keto/models"
)

func TestEngine(t *testing.T) {
	t.Run("direct inclusion", func(t *testing.T) {
		rel := models.InternalRelationTuple{
			Relation: "access",
			Object: &models.Object{
				ID:        "object",
				Namespace: "test",
			},
			Subject: &models.UserID{ID: "user"},
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
		dust := models.Object{
			ID:        "dust",
			Namespace: "under the sofa",
		}
		mark := models.UserID{
			ID: "Mark",
		}
		cleaningRelation := models.InternalRelationTuple{
			Relation: "have to remove",
			Object:   &dust,
			Subject: &models.UserSet{
				Relation: "producer",
				Object:   &dust,
			},
		}
		markProducesDust := models.InternalRelationTuple{
			Relation: "producer",
			Object:   &dust,
			Subject:  &mark,
		}

		reg := &driver.RegistryDefault{}
		require.NoError(t, reg.RelationTupleManager().WriteRelationTuples(context.Background(), &cleaningRelation, &markProducesDust))

		e := check.NewEngine(reg)

		res, err := e.SubjectIsAllowed(context.Background(), &models.InternalRelationTuple{
			Relation: cleaningRelation.Relation,
			Object:   &dust,
			Subject:  &mark,
		})
		require.NoError(t, err)
		assert.True(t, res)
	})

	t.Run("direct exclusion", func(t *testing.T) {
		user := &models.UserID{
			ID: "user-id",
		}
		rel := models.InternalRelationTuple{
			Relation: "relation",
			Object: &models.Object{
				ID:        "object-id",
				Namespace: "object-namespace",
			},
			Subject: user,
		}

		reg := &driver.RegistryDefault{}
		require.NoError(t, reg.RelationTupleManager().WriteRelationTuples(context.Background(), &rel))

		e := check.NewEngine(reg)

		res, err := e.SubjectIsAllowed(context.Background(), &models.InternalRelationTuple{
			Relation: rel.Relation,
			Object:   rel.Object,
			Subject:  &models.UserID{ID: "not " + user.ID},
		})
		require.NoError(t, err)
		assert.False(t, res)
	})

	t.Run("wrong object ID", func(t *testing.T) {
		object := models.Object{
			ID: "object",
		}
		access := models.InternalRelationTuple{
			Relation: "access",
			Object:   &object,
			Subject: &models.UserSet{
				Relation: "owner",
				Object:   &object,
			},
		}
		user := models.InternalRelationTuple{
			Relation: "owner",
			Object:   &models.Object{ID: "not " + object.ID},
			Subject:  &models.UserID{ID: "user"},
		}

		reg := &driver.RegistryDefault{}
		require.NoError(t, reg.RelationTupleManager().WriteRelationTuples(context.Background(), &access, &user))

		e := check.NewEngine(reg)

		res, err := e.SubjectIsAllowed(context.Background(), &models.InternalRelationTuple{
			Relation: access.Relation,
			Object:   &object,
			Subject:  user.Subject,
		})
		require.NoError(t, err)
		assert.False(t, res)
	})

	t.Run("wrong relation name", func(t *testing.T) {
		diaryEntry := &models.Object{
			ID:        "entry for 6. Nov 2020",
			Namespace: "diary",
		}
		// this would be a userset rewrite
		readDiary := models.InternalRelationTuple{
			Relation: "read",
			Object:   diaryEntry,
			Subject: &models.UserSet{
				Relation: "author",
				Object:   diaryEntry,
			},
		}
		user := models.InternalRelationTuple{
			Relation: "not author",
			Object:   diaryEntry,
			Subject:  &models.UserID{ID: "your mother"},
		}

		reg := &driver.RegistryDefault{}
		require.NoError(t, reg.RelationTupleManager().WriteRelationTuples(context.Background(), &readDiary, &user))

		e := check.NewEngine(reg)

		res, err := e.SubjectIsAllowed(context.Background(), &models.InternalRelationTuple{
			Relation: readDiary.Relation,
			Object:   diaryEntry,
			Subject:  user.Subject,
		})
		require.NoError(t, err)
		assert.False(t, res)
	})

	t.Run("indirect inclusion level 2", func(t *testing.T) {
		object := models.Object{
			ID:        "some object",
			Namespace: "some namespace",
		}
		user := models.UserID{
			ID: "some user",
		}
		organization := models.Object{
			ID:        "some organization",
			Namespace: "all organizations",
		}

		ownerUserSet := models.UserSet{
			Relation: "owner",
			Object:   &object,
		}
		orgMembers := models.UserSet{
			Relation: "member",
			Object:   &organization,
		}

		writeRel := models.InternalRelationTuple{
			Relation: "write",
			Object:   &object,
			Subject:  &ownerUserSet,
		}
		orgOwnerRel := models.InternalRelationTuple{
			Relation: ownerUserSet.Relation,
			Object:   &object,
			Subject:  &orgMembers,
		}
		userMembershipRel := models.InternalRelationTuple{
			Relation: orgMembers.Relation,
			Object:   orgMembers.Object,
			Subject:  &user,
		}

		reg := &driver.RegistryDefault{}
		require.NoError(t, reg.RelationTupleManager().WriteRelationTuples(context.Background(), &writeRel, &orgOwnerRel, &userMembershipRel))

		e := check.NewEngine(reg)

		// user can write object
		res, err := e.SubjectIsAllowed(context.Background(), &models.InternalRelationTuple{
			Relation: writeRel.Relation,
			Object:   &object,
			Subject:  &user,
		})
		require.NoError(t, err)
		assert.True(t, res)

		// user is member of the organization
		res, err = e.SubjectIsAllowed(context.Background(), &models.InternalRelationTuple{
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

		file := models.Object{ID: "file"}
		directory := models.Object{ID: "directory"}
		user := models.UserID{ID: "user"}

		parent := models.InternalRelationTuple{
			Relation: "parent",
			Object:   &file,
			Subject: &models.UserSet{ // <- this is only an object, but this is allowed as a userset can have the "..." relation which means any relation
				Object: &directory,
			},
		}
		directoryAccess := models.InternalRelationTuple{
			Relation: "access",
			Object:   &directory,
			Subject:  &user,
		}

		reg := &driver.RegistryDefault{}
		for _, r := range []*models.InternalRelationTuple{&parent, &directoryAccess} {
			require.NoError(t, reg.RelationTupleManager().WriteRelationTuples(context.Background(), r))
		}

		e := check.NewEngine(reg)

		res, err := e.SubjectIsAllowed(context.Background(), &models.InternalRelationTuple{
			Relation: directoryAccess.Relation,
			Object:   &file,
			Subject:  &user,
		})
		require.NoError(t, err)
		assert.False(t, res)
	})
}
