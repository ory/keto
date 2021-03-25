package check_test

import (
	"context"
	"testing"

	"github.com/ory/keto/internal/x"

	"github.com/ory/keto/internal/namespace"

	"github.com/ory/keto/internal/relationtuple"

	"github.com/ory/keto/internal/check"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ory/keto/internal/driver"
)

func newDepsProvider(t *testing.T, namespaces []*namespace.Namespace, pageOpts ...x.PaginationOptionSetter) *relationtuple.ManagerWrapper {
	reg := driver.NewMemoryTestRegistry(t, namespaces)
	t.Cleanup(func() {
		for _, n := range namespaces {
			mb, err := reg.NamespaceMigrator().NamespaceMigrationBox(context.Background(), n)
			require.NoError(t, err)
			require.NoError(t, mb.Down(context.Background(), 0))
		}
	})
	return relationtuple.NewManagerWrapper(t, reg, pageOpts...)
}

func TestEngine(t *testing.T) {
	t.Run("direct inclusion", func(t *testing.T) {
		rel := relationtuple.InternalRelationTuple{
			Relation:  "access",
			Object:    "object",
			Namespace: "test",
			Subject:   &relationtuple.SubjectID{ID: "user"},
		}

		reg := newDepsProvider(t, []*namespace.Namespace{
			{Name: rel.Namespace, ID: 1},
		})
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
		mark := relationtuple.SubjectID{
			ID: "Mark",
		}
		cleaningRelation := relationtuple.InternalRelationTuple{
			Namespace: sofaNamespace,
			Relation:  "have to remove",
			Object:    dust,
			Subject: &relationtuple.SubjectSet{
				Relation:  "producer",
				Object:    dust,
				Namespace: sofaNamespace,
			},
		}
		markProducesDust := relationtuple.InternalRelationTuple{
			Namespace: sofaNamespace,
			Relation:  "producer",
			Object:    dust,
			Subject:   &mark,
		}

		reg := newDepsProvider(t, []*namespace.Namespace{
			{Name: sofaNamespace, ID: 1},
		})
		require.NoError(t, reg.RelationTupleManager().WriteRelationTuples(context.Background(), &cleaningRelation, &markProducesDust))

		e := check.NewEngine(reg)

		res, err := e.SubjectIsAllowed(context.Background(), &relationtuple.InternalRelationTuple{
			Relation:  cleaningRelation.Relation,
			Object:    dust,
			Subject:   &mark,
			Namespace: sofaNamespace,
		})
		require.NoError(t, err)
		assert.True(t, res)
	})

	t.Run("direct exclusion", func(t *testing.T) {
		user := &relationtuple.SubjectID{
			ID: "user-id",
		}
		rel := relationtuple.InternalRelationTuple{
			Relation:  "relation",
			Object:    "object-id",
			Namespace: "object-namespace",
			Subject:   user,
		}

		reg := newDepsProvider(t, []*namespace.Namespace{
			{Name: rel.Namespace, ID: 10},
		})
		require.NoError(t, reg.RelationTupleManager().WriteRelationTuples(context.Background(), &rel))

		e := check.NewEngine(reg)

		res, err := e.SubjectIsAllowed(context.Background(), &relationtuple.InternalRelationTuple{
			Relation:  rel.Relation,
			Object:    rel.Object,
			Namespace: rel.Namespace,
			Subject:   &relationtuple.SubjectID{ID: "not " + user.ID},
		})
		require.NoError(t, err)
		assert.False(t, res)
	})

	t.Run("wrong object ID", func(t *testing.T) {
		object := "object"
		access := relationtuple.InternalRelationTuple{
			Relation: "access",
			Object:   object,
			Subject: &relationtuple.SubjectSet{
				Relation: "owner",
				Object:   object,
			},
		}
		user := relationtuple.InternalRelationTuple{
			Relation: "owner",
			Object:   "not " + object,
			Subject:  &relationtuple.SubjectID{ID: "user"},
		}

		reg := newDepsProvider(t, []*namespace.Namespace{
			{Name: "", ID: 1},
		})
		require.NoError(t, reg.RelationTupleManager().WriteRelationTuples(context.Background(), &access, &user))

		e := check.NewEngine(reg)

		res, err := e.SubjectIsAllowed(context.Background(), &relationtuple.InternalRelationTuple{
			Relation: access.Relation,
			Object:   object,
			Subject:  user.Subject,
		})
		require.NoError(t, err)
		assert.False(t, res)
	})

	t.Run("wrong relation name", func(t *testing.T) {
		diaryEntry := "entry for 6. Nov 2020"
		diaryNamespace := "diary"
		// this would be a userset rewrite
		readDiary := relationtuple.InternalRelationTuple{
			Namespace: diaryNamespace,
			Relation:  "read",
			Object:    diaryEntry,
			Subject: &relationtuple.SubjectSet{
				Relation:  "author",
				Object:    diaryEntry,
				Namespace: diaryNamespace,
			},
		}
		user := relationtuple.InternalRelationTuple{
			Namespace: diaryNamespace,
			Relation:  "not author",
			Object:    diaryEntry,
			Subject:   &relationtuple.SubjectID{ID: "your mother"},
		}

		reg := newDepsProvider(t, []*namespace.Namespace{
			{Name: diaryNamespace, ID: 1},
		})
		require.NoError(t, reg.RelationTupleManager().WriteRelationTuples(context.Background(), &readDiary, &user))

		e := check.NewEngine(reg)

		res, err := e.SubjectIsAllowed(context.Background(), &relationtuple.InternalRelationTuple{
			Relation:  readDiary.Relation,
			Object:    diaryEntry,
			Namespace: diaryNamespace,
			Subject:   user.Subject,
		})
		require.NoError(t, err)
		assert.False(t, res)
	})

	t.Run("indirect inclusion level 2", func(t *testing.T) {
		object := "some object"
		someNamespace := "some namespace"
		user := relationtuple.SubjectID{
			ID: "some user",
		}
		organization := "some organization"
		orgNamespace := "all organizations"

		ownerUserSet := relationtuple.SubjectSet{
			Namespace: someNamespace,
			Relation:  "owner",
			Object:    object,
		}
		orgMembers := relationtuple.SubjectSet{
			Namespace: orgNamespace,
			Relation:  "member",
			Object:    organization,
		}

		writeRel := relationtuple.InternalRelationTuple{
			Namespace: someNamespace,
			Relation:  "write",
			Object:    object,
			Subject:   &ownerUserSet,
		}
		orgOwnerRel := relationtuple.InternalRelationTuple{
			Namespace: someNamespace,
			Relation:  ownerUserSet.Relation,
			Object:    object,
			Subject:   &orgMembers,
		}
		userMembershipRel := relationtuple.InternalRelationTuple{
			Namespace: orgNamespace,
			Relation:  orgMembers.Relation,
			Object:    organization,
			Subject:   &user,
		}

		reg := newDepsProvider(t, []*namespace.Namespace{
			{Name: someNamespace, ID: 1},
			{Name: orgNamespace, ID: 2},
		})
		require.NoError(t, reg.RelationTupleManager().WriteRelationTuples(context.Background(), &writeRel, &orgOwnerRel, &userMembershipRel))

		e := check.NewEngine(reg)

		// user can write object
		res, err := e.SubjectIsAllowed(context.Background(), &relationtuple.InternalRelationTuple{
			Namespace: someNamespace,
			Relation:  writeRel.Relation,
			Object:    object,
			Subject:   &user,
		})
		require.NoError(t, err)
		assert.True(t, res)

		// user is member of the organization
		res, err = e.SubjectIsAllowed(context.Background(), &relationtuple.InternalRelationTuple{
			Namespace: orgNamespace,
			Relation:  orgMembers.Relation,
			Object:    organization,
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
		user := relationtuple.SubjectID{ID: "user"}

		parent := relationtuple.InternalRelationTuple{
			Relation: "parent",
			Object:   file,
			Subject: &relationtuple.SubjectSet{ // <- this is only an object, but this is allowed as a userset can have the "..." relation which means any relation
				Object: directory,
			},
		}
		directoryAccess := relationtuple.InternalRelationTuple{
			Relation: "access",
			Object:   directory,
			Subject:  &user,
		}

		reg := newDepsProvider(t, []*namespace.Namespace{
			{Name: "", ID: 2},
		})
		require.NoError(t, reg.RelationTupleManager().WriteRelationTuples(context.Background(), &parent, &directoryAccess))

		e := check.NewEngine(reg)

		res, err := e.SubjectIsAllowed(context.Background(), &relationtuple.InternalRelationTuple{
			Relation: directoryAccess.Relation,
			Object:   file,
			Subject:  &user,
		})
		require.NoError(t, err)
		assert.False(t, res)
	})

	t.Run("case=subject id next to subject set", func(t *testing.T) {
		namesp, obj, org, directOwner, indirectOwner, ownerRel, memberRel := "namesp", "obj", "org", "u1", "u2", "owner", "member"

		reg := newDepsProvider(t, []*namespace.Namespace{
			{Name: namesp, ID: 1},
		})
		require.NoError(t, reg.RelationTupleManager().WriteRelationTuples(
			context.Background(),
			&relationtuple.InternalRelationTuple{
				Namespace: namesp,
				Object:    obj,
				Relation:  ownerRel,
				Subject:   &relationtuple.SubjectID{ID: directOwner},
			},
			&relationtuple.InternalRelationTuple{
				Namespace: namesp,
				Object:    obj,
				Relation:  ownerRel,
				Subject: &relationtuple.SubjectSet{
					Namespace: namesp,
					Object:    org,
					Relation:  memberRel,
				},
			},
			&relationtuple.InternalRelationTuple{
				Namespace: namesp,
				Object:    org,
				Relation:  memberRel,
				Subject:   &relationtuple.SubjectID{ID: indirectOwner},
			},
		))

		e := check.NewEngine(reg)

		res, err := e.SubjectIsAllowed(context.Background(), &relationtuple.InternalRelationTuple{
			Namespace: namesp,
			Object:    obj,
			Relation:  ownerRel,
			Subject:   &relationtuple.SubjectID{ID: directOwner},
		})
		require.NoError(t, err)
		assert.True(t, res)

		res, err = e.SubjectIsAllowed(context.Background(), &relationtuple.InternalRelationTuple{
			Namespace: namesp,
			Object:    obj,
			Relation:  ownerRel,
			Subject:   &relationtuple.SubjectID{ID: indirectOwner},
		})
		require.NoError(t, err)
		assert.True(t, res)
	})

	t.Run("case=paginates", func(t *testing.T) {
		namesp, obj, access, users := "namesp", "obj", "access", []string{"u1", "u2", "u3", "u4"}
		pageSize := 2

		reg := newDepsProvider(
			t,
			[]*namespace.Namespace{{Name: namesp, ID: 1}},
			x.WithSize(pageSize),
		)

		for _, user := range users {
			require.NoError(t, reg.RelationTupleManager().WriteRelationTuples(context.Background(), &relationtuple.InternalRelationTuple{
				Namespace: namesp,
				Object:    obj,
				Relation:  access,
				Subject:   &relationtuple.SubjectID{ID: user},
			}))
		}

		e := check.NewEngine(reg)

		for i, user := range users {
			t.Run("user="+user, func(t *testing.T) {
				allowed, err := e.SubjectIsAllowed(context.Background(), &relationtuple.InternalRelationTuple{
					Namespace: namesp,
					Object:    obj,
					Relation:  access,
					Subject:   &relationtuple.SubjectID{ID: user},
				})
				require.NoError(t, err)
				assert.True(t, allowed)

				// pagination assertions
				if i >= pageSize {
					assert.Len(t, reg.RequestedPages, 2)
					// reset requested pages for next iteration
					reg.RequestedPages = nil
				} else {
					assert.Len(t, reg.RequestedPages, 1)
					// reset requested pages for next iteration
					reg.RequestedPages = nil
				}
			})
		}
	})

	t.Run("case=wide tuple graph", func(t *testing.T) {
		namesp, obj, access, member, users, orgs := "namesp", "obj", "access", "member", []string{"u1", "u2", "u3", "u4"}, []string{"o1", "o2"}

		reg := newDepsProvider(t, []*namespace.Namespace{{Name: namesp, ID: 1}})

		for _, org := range orgs {
			require.NoError(t, reg.RelationTupleManager().WriteRelationTuples(context.Background(), &relationtuple.InternalRelationTuple{
				Namespace: namesp,
				Object:    obj,
				Relation:  access,
				Subject: &relationtuple.SubjectSet{
					Namespace: namesp,
					Object:    org,
					Relation:  member,
				},
			}))
		}

		for i, user := range users {
			require.NoError(t, reg.RelationTupleManager().WriteRelationTuples(context.Background(), &relationtuple.InternalRelationTuple{
				Namespace: namesp,
				Object:    orgs[i%len(orgs)],
				Relation:  member,
				Subject:   &relationtuple.SubjectID{ID: user},
			}))
		}

		e := check.NewEngine(reg)

		for _, user := range users {
			req := &relationtuple.InternalRelationTuple{
				Namespace: namesp,
				Object:    obj,
				Relation:  access,
				Subject:   &relationtuple.SubjectID{ID: user},
			}
			allowed, err := e.SubjectIsAllowed(context.Background(), req)
			require.NoError(t, err)
			assert.True(t, allowed, req.String())
		}
	})
}
