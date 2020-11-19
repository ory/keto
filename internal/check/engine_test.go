package check_test

import (
	"context"
	"testing"

	"github.com/ory/keto/internal/namespace"

	"github.com/ory/keto/internal/relationtuple"

	"github.com/ory/keto/internal/check"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ory/keto/internal/driver"
)

func TestEngine(t *testing.T) {
	t.Run("direct inclusion", func(t *testing.T) {
		rel := relationtuple.InternalRelationTuple{
			Relation:  "access",
			ObjectID:  "object",
			Namespace: "test",
			Subject:   &relationtuple.UserID{ID: "user"},
		}

		reg := &driver.RegistryDefault{}
		require.NoError(t, reg.NamespaceManager().MigrateNamespaceUp(&namespace.Namespace{Name: rel.Namespace}))
		require.NoError(t, reg.RelationTupleManager().WriteRelationTuples(context.Background(), &rel))

		e := check.NewEngine(reg)

		res, err := e.SubjectIsAllowed(context.Background(), &rel)
		require.NoError(t, err)
		assert.True(t, res)
	})

	t.Run("indirect inclusion level 1", func(t *testing.T) {
		// the set of users that are produces of "dust" have to remove it
		dust := "dust"
		sofaNamespace := "under the sofa"
		mark := relationtuple.UserID{
			ID: "Mark",
		}
		cleaningRelation := relationtuple.InternalRelationTuple{
			Relation:  "have to remove",
			ObjectID:  dust,
			Namespace: sofaNamespace,
			Subject: &relationtuple.UserSet{
				Relation:  "producer",
				ObjectID:  dust,
				Namespace: sofaNamespace,
			},
		}
		markProducesDust := relationtuple.InternalRelationTuple{
			Relation:  "producer",
			ObjectID:  dust,
			Namespace: sofaNamespace,
			Subject:   &mark,
		}

		reg := &driver.RegistryDefault{}
		require.NoError(t, reg.NamespaceManager().MigrateNamespaceUp(&namespace.Namespace{Name: sofaNamespace}))
		require.NoError(t, reg.RelationTupleManager().WriteRelationTuples(context.Background(), &cleaningRelation, &markProducesDust))

		e := check.NewEngine(reg)

		res, err := e.SubjectIsAllowed(context.Background(), &relationtuple.InternalRelationTuple{
			Relation:  cleaningRelation.Relation,
			ObjectID:  dust,
			Namespace: sofaNamespace,
			Subject:   &mark,
		})
		require.NoError(t, err)
		assert.True(t, res)
	})

	t.Run("direct exclusion", func(t *testing.T) {
		user := &relationtuple.UserID{
			ID: "user-id",
		}
		rel := relationtuple.InternalRelationTuple{
			Relation:  "relation",
			ObjectID:  "object-id",
			Namespace: "object-namespace",
			Subject:   user,
		}

		reg := &driver.RegistryDefault{}
		require.NoError(t, reg.NamespaceManager().MigrateNamespaceUp(&namespace.Namespace{Name: rel.Namespace}))
		require.NoError(t, reg.RelationTupleManager().WriteRelationTuples(context.Background(), &rel))

		e := check.NewEngine(reg)

		res, err := e.SubjectIsAllowed(context.Background(), &relationtuple.InternalRelationTuple{
			Relation:  rel.Relation,
			ObjectID:  rel.ObjectID,
			Namespace: rel.Namespace,
			Subject:   &relationtuple.UserID{ID: "not " + user.ID},
		})
		require.NoError(t, err)
		assert.False(t, res)
	})

	t.Run("wrong object ID", func(t *testing.T) {
		object := "object"
		defaultNamespace := "default"
		access := relationtuple.InternalRelationTuple{
			Relation:  "access",
			ObjectID:  object,
			Namespace: defaultNamespace,
			Subject: &relationtuple.UserSet{
				Relation:  "owner",
				ObjectID:  object,
				Namespace: defaultNamespace,
			},
		}
		user := relationtuple.InternalRelationTuple{
			Relation:  "owner",
			ObjectID:  "not " + object,
			Namespace: defaultNamespace,
			Subject:   &relationtuple.UserID{ID: "user"},
		}

		reg := &driver.RegistryDefault{}
		require.NoError(t, reg.NamespaceManager().MigrateNamespaceUp(&namespace.Namespace{Name: defaultNamespace}))
		require.NoError(t, reg.RelationTupleManager().WriteRelationTuples(context.Background(), &access, &user))

		e := check.NewEngine(reg)

		res, err := e.SubjectIsAllowed(context.Background(), &relationtuple.InternalRelationTuple{
			Relation:  access.Relation,
			ObjectID:  object,
			Subject:   user.Subject,
			Namespace: defaultNamespace,
		})
		require.NoError(t, err)
		assert.False(t, res)
	})

	t.Run("wrong relation name", func(t *testing.T) {
		diaryEntry := "entry for 6. Nov 2020"
		diaryNamespace := "diary"
		// this would be a userset rewrite
		readDiary := relationtuple.InternalRelationTuple{
			Relation:  "read",
			ObjectID:  diaryEntry,
			Namespace: diaryNamespace,
			Subject: &relationtuple.UserSet{
				Relation:  "author",
				ObjectID:  diaryEntry,
				Namespace: diaryNamespace,
			},
		}
		user := relationtuple.InternalRelationTuple{
			Relation:  "not author",
			ObjectID:  diaryEntry,
			Namespace: diaryNamespace,
			Subject:   &relationtuple.UserID{ID: "your mother"},
		}

		reg := &driver.RegistryDefault{}
		require.NoError(t, reg.NamespaceManager().MigrateNamespaceUp(&namespace.Namespace{Name: diaryNamespace}))
		require.NoError(t, reg.RelationTupleManager().WriteRelationTuples(context.Background(), &readDiary, &user))

		e := check.NewEngine(reg)

		res, err := e.SubjectIsAllowed(context.Background(), &relationtuple.InternalRelationTuple{
			Relation:  readDiary.Relation,
			ObjectID:  diaryEntry,
			Namespace: diaryNamespace,
			Subject:   user.Subject,
		})
		require.NoError(t, err)
		assert.False(t, res)
	})

	t.Run("indirect inclusion level 2", func(t *testing.T) {
		object := "some object"
		objNamespace := "some namespace"

		user := relationtuple.UserID{
			ID: "some user",
		}

		organization := "some organization"
		orgNamespace := "all organizations"

		ownerUserSet := relationtuple.UserSet{
			Relation:  "owner",
			ObjectID:  object,
			Namespace: objNamespace,
		}
		orgMembers := relationtuple.UserSet{
			Relation:  "member",
			ObjectID:  organization,
			Namespace: orgNamespace,
		}

		writeRel := relationtuple.InternalRelationTuple{
			Relation:  "write",
			ObjectID:  object,
			Namespace: objNamespace,
			Subject:   &ownerUserSet,
		}
		orgOwnerRel := relationtuple.InternalRelationTuple{
			Relation:  ownerUserSet.Relation,
			ObjectID:  object,
			Namespace: objNamespace,
			Subject:   &orgMembers,
		}
		userMembershipRel := relationtuple.InternalRelationTuple{
			Relation:  orgMembers.Relation,
			ObjectID:  organization,
			Namespace: orgNamespace,
			Subject:   &user,
		}

		reg := &driver.RegistryDefault{}
		require.NoError(t, reg.NamespaceManager().MigrateNamespaceUp(&namespace.Namespace{Name: objNamespace, ID: 0}))
		require.NoError(t, reg.NamespaceManager().MigrateNamespaceUp(&namespace.Namespace{Name: orgNamespace, ID: 1}))
		require.NoError(t, reg.RelationTupleManager().WriteRelationTuples(context.Background(), &writeRel, &orgOwnerRel, &userMembershipRel))

		e := check.NewEngine(reg)

		// user can write object
		res, err := e.SubjectIsAllowed(context.Background(), &relationtuple.InternalRelationTuple{
			Relation:  writeRel.Relation,
			ObjectID:  object,
			Namespace: objNamespace,
			Subject:   &user,
		})
		require.NoError(t, err)
		assert.True(t, res)

		// user is member of the organization
		res, err = e.SubjectIsAllowed(context.Background(), &relationtuple.InternalRelationTuple{
			Relation:  orgMembers.Relation,
			ObjectID:  organization,
			Namespace: orgNamespace,
			Subject:   &user,
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

		file := "file"
		directory := "directory"
		driveNamespace := "shared drive"

		user := relationtuple.UserID{ID: "user"}

		parent := relationtuple.InternalRelationTuple{
			Relation:  "parent",
			ObjectID:  file,
			Namespace: driveNamespace,
			Subject: &relationtuple.UserSet{ // <- this is only an object, but this is allowed as a userset can have the "..." relation which means any relation
				ObjectID:  directory,
				Namespace: driveNamespace,
			},
		}
		directoryAccess := relationtuple.InternalRelationTuple{
			Relation:  "access",
			ObjectID:  directory,
			Namespace: driveNamespace,
			Subject:   &user,
		}

		reg := &driver.RegistryDefault{}
		require.NoError(t, reg.NamespaceManager().MigrateNamespaceUp(&namespace.Namespace{Name: driveNamespace}))
		require.NoError(t, reg.RelationTupleManager().WriteRelationTuples(context.Background(), &parent, &directoryAccess))

		e := check.NewEngine(reg)

		res, err := e.SubjectIsAllowed(context.Background(), &relationtuple.InternalRelationTuple{
			Relation:  directoryAccess.Relation,
			ObjectID:  file,
			Namespace: driveNamespace,
			Subject:   &user,
		})
		require.NoError(t, err)
		assert.False(t, res)
	})
}
