package check

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ory/keto/driver"
	"github.com/ory/keto/models"
)

func TestEngine(t *testing.T) {
	t.Run("direct inclusion", func(t *testing.T) {
		rel := models.Relation{
			Name:      "access",
			ObjectID:  "object",
			SubjectID: "user",
		}

		reg := &driver.RegistryDefault{}
		require.NoError(t, reg.RelationManager().WriteRelation(context.Background(), &rel))

		e := NewEngine(reg)

		res, err := e.SubjectIsAllowed(context.Background(), &rel)
		require.NoError(t, err)
		assert.True(t, res)
	})

	t.Run("indirect inclusion level 1", func(t *testing.T) {
		// the set of users that are owners of "object" have access to it
		ownersUserSet := models.Relation{
			Name:     "owner",
			ObjectID: "object",
		}
		fileAccessRel := models.Relation{
			Name:      "access",
			ObjectID:  ownersUserSet.ObjectID,
			SubjectID: ownersUserSet.ToSubject(),
		}
		memberRel := models.Relation{
			Name:      "owner",
			ObjectID:  ownersUserSet.ObjectID,
			SubjectID: "user",
		}

		reg := &driver.RegistryDefault{}
		require.NoError(t, reg.RelationManager().WriteRelation(context.Background(), &ownersUserSet))
		require.NoError(t, reg.RelationManager().WriteRelation(context.Background(), &fileAccessRel))
		require.NoError(t, reg.RelationManager().WriteRelation(context.Background(), &memberRel))

		e := NewEngine(reg)

		res, err := e.SubjectIsAllowed(context.Background(), &models.Relation{
			Name:      fileAccessRel.Name,
			ObjectID:  ownersUserSet.ObjectID,
			SubjectID: memberRel.SubjectID,
		})
		require.NoError(t, err)
		assert.True(t, res)
	})

	t.Run("direct exclusion", func(t *testing.T) {
		rel := models.Relation{
			Name:      "relation",
			ObjectID:  "object",
			SubjectID: "user",
		}

		reg := &driver.RegistryDefault{}
		require.NoError(t, reg.RelationManager().WriteRelation(context.Background(), &rel))

		e := NewEngine(reg)

		rel.SubjectID = "not " + rel.SubjectID
		res, err := e.SubjectIsAllowed(context.Background(), &rel)
		require.NoError(t, err)
		assert.False(t, res)
	})

	t.Run("wrong object ID", func(t *testing.T) {
		owners := models.Relation{
			Name:     "owner",
			ObjectID: "object",
		}
		access := models.Relation{
			Name:      "access",
			ObjectID:  owners.ObjectID,
			SubjectID: owners.ToSubject(),
		}
		user := models.Relation{
			Name:      "owner",
			ObjectID:  "not " + owners.ObjectID,
			SubjectID: "user",
		}

		reg := &driver.RegistryDefault{}
		require.NoError(t, reg.RelationManager().WriteRelation(context.Background(), &owners))
		require.NoError(t, reg.RelationManager().WriteRelation(context.Background(), &access))
		require.NoError(t, reg.RelationManager().WriteRelation(context.Background(), &user))

		e := NewEngine(reg)

		res, err := e.SubjectIsAllowed(context.Background(), &models.Relation{
			Name:      access.Name,
			ObjectID:  owners.ObjectID,
			SubjectID: user.SubjectID,
		})
		require.NoError(t, err)
		assert.False(t, res)
	})

	t.Run("missing access relation", func(t *testing.T) {
		object := "test object"
		owners := models.Relation{
			Name:     "owner",
			ObjectID: object,
		}
		access := models.Relation{
			Name:      "access",
			ObjectID:  object,
			SubjectID: owners.ToSubject(),
		}
		user := models.Relation{
			Name:      "not " + owners.Name,
			ObjectID:  object,
			SubjectID: "user",
		}

		reg := &driver.RegistryDefault{}
		for _, r := range []*models.Relation{&owners, &access, &user} {
			require.NoError(t, reg.RelationManager().WriteRelation(context.Background(), r))
		}

		e := NewEngine(reg)

		res, err := e.SubjectIsAllowed(context.Background(), &models.Relation{
			Name:      access.Name,
			ObjectID:  object,
			SubjectID: user.SubjectID,
		})
		require.NoError(t, err)
		assert.False(t, res)
	})

	t.Run("indirect inclusion level 2", func(t *testing.T) {
		object := "some random object"
		user := "some user"
		organization := "some organization"

		ownerUserSet := models.Relation{
			Name:     "owner",
			ObjectID: object,
		}
		writeRel := models.Relation{
			Name:      "write",
			ObjectID:  object,
			SubjectID: ownerUserSet.ToSubject(),
		}
		orgMemberRel := models.Relation{
			Name:     "member",
			ObjectID: organization,
		}
		orgOwnerRel := models.Relation{
			Name:      ownerUserSet.Name,
			ObjectID:  object,
			SubjectID: orgMemberRel.ToSubject(),
		}
		userMembershipRel := models.Relation{
			Name:      orgMemberRel.Name,
			ObjectID:  orgMemberRel.ObjectID,
			SubjectID: user,
		}

		reg := &driver.RegistryDefault{}
		for _, r := range []*models.Relation{&ownerUserSet, &writeRel, &orgMemberRel, &orgOwnerRel, &userMembershipRel} {
			require.NoError(t, reg.RelationManager().WriteRelation(context.Background(), r))
		}

		e := NewEngine(reg)

		res, err := e.SubjectIsAllowed(context.Background(), &models.Relation{
			Name:      writeRel.Name,
			ObjectID:  object,
			SubjectID: user,
		})
		require.NoError(t, err)
		assert.True(t, res)

		res, err = e.SubjectIsAllowed(context.Background(), &models.Relation{
			Name:      orgMemberRel.Name,
			ObjectID:  organization,
			SubjectID: user,
		})
		require.NoError(t, err)
		assert.True(t, res)
	})

	t.Run("rejects transitive relation", func(t *testing.T) {
		// (file) <--parent-- (folder) <--access-- [user]
		//
		// as we don't know how to interpret the "parent" relation, there would have to be a userset rewrite to allow access
		// to files when you have access to the parent

		object := "file"
		user := "user"
		parentRel := models.Relation{
			Name:      "parent",
			ObjectID:  object,
			SubjectID: "folder", // <- this is an object, but this is allowed as a userset can have the "..." relation which means any relation
		}
		folderAccessRel := models.Relation{
			Name:      "access",
			ObjectID:  parentRel.SubjectID,
			SubjectID: user,
		}

		reg := &driver.RegistryDefault{}
		for _, r := range []*models.Relation{&parentRel, &folderAccessRel} {
			require.NoError(t, reg.RelationManager().WriteRelation(context.Background(), r))
		}

		e := NewEngine(reg)

		res, err := e.SubjectIsAllowed(context.Background(), &models.Relation{
			Name:      folderAccessRel.Name,
			ObjectID:  object,
			SubjectID: user,
		})
		require.NoError(t, err)
		assert.False(t, res)
	})
}
