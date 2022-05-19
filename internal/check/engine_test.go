package check_test

import (
	"context"
	"testing"

	"github.com/ory/keto/internal/driver/config"

	"github.com/ory/keto/internal/x"

	"github.com/ory/keto/internal/namespace"

	"github.com/ory/keto/internal/relationtuple"

	"github.com/ory/keto/internal/check"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ory/keto/internal/driver"
)

type configProvider = config.Provider
type loggerProvider = x.LoggerProvider

// deps is defined to capture engine dependencies in a single struct
type deps struct {
	*relationtuple.ManagerWrapper // managerProvider
	configProvider
	loggerProvider
}

func newDepsProvider(t *testing.T, namespaces []*namespace.Namespace, pageOpts ...x.PaginationOptionSetter) *deps {
	reg := driver.NewSqliteTestRegistry(t, false)
	require.NoError(t, reg.Config(context.Background()).Set(config.KeyNamespaces, namespaces))
	mr := relationtuple.NewManagerWrapper(t, reg, pageOpts...)

	return &deps{
		ManagerWrapper: mr,
		configProvider: reg,
		loggerProvider: reg,
	}
}

func TestEngine(t *testing.T) {
	ctx := context.Background()

	t.Run("respects max depth", func(t *testing.T) {
		reg := newDepsProvider(t, []*namespace.Namespace{
			{Name: "test", ID: 1},
		})

		// "user" has relation "access" through being an "owner" through being
		// an "admin" which requires at least 2 units of depth. If max-depth is
		// 2 then we hit max-depth
		insertFixtures(t, reg.RelationTupleManager(), []string{
			"test:object#admin@user",
			"test:object#owner@test:object#admin",
			"test:object#access@test:object#owner",
		})

		e := check.NewEngine(reg)

		userHasAccess, err := relationtuple.InternalFromString("test:object#access@user")
		require.NoError(t, err)

		// global max-depth defaults to 5
		assert.Equal(t, reg.Config(ctx).MaxReadDepth(), 5)

		// req max-depth takes precedence, max-depth=2 is not enough
		res, err := e.SubjectIsAllowed(ctx, userHasAccess, 2)
		require.NoError(t, err)
		assert.False(t, res)

		// req max-depth takes precedence, max-depth=3 is enough
		res, err = e.SubjectIsAllowed(ctx, userHasAccess, 3)
		require.NoError(t, err)
		assert.True(t, res)

		// global max-depth takes precedence and max-depth=2 is not enough
		require.NoError(t, reg.Config(ctx).Set(config.KeyLimitMaxReadDepth, 2))
		res, err = e.SubjectIsAllowed(ctx, userHasAccess, 3)
		require.NoError(t, err)
		assert.False(t, res)

		// global max-depth takes precedence and max-depth=3 is enough
		require.NoError(t, reg.Config(ctx).Set(config.KeyLimitMaxReadDepth, 3))
		res, err = e.SubjectIsAllowed(ctx, userHasAccess, 0)
		require.NoError(t, err)
		assert.True(t, res)
	})

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
		require.NoError(t, reg.RelationTupleManager().WriteRelationTuples(ctx, &rel))

		e := check.NewEngine(reg)

		res, err := e.SubjectIsAllowed(ctx, &rel, 0)
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
		require.NoError(t, reg.RelationTupleManager().WriteRelationTuples(ctx, &cleaningRelation, &markProducesDust))

		e := check.NewEngine(reg)

		res, err := e.SubjectIsAllowed(ctx, &relationtuple.InternalRelationTuple{
			Relation:  cleaningRelation.Relation,
			Object:    dust,
			Subject:   &mark,
			Namespace: sofaNamespace,
		}, 0)
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
		require.NoError(t, reg.RelationTupleManager().WriteRelationTuples(ctx, &rel))

		e := check.NewEngine(reg)

		res, err := e.SubjectIsAllowed(ctx, &relationtuple.InternalRelationTuple{
			Relation:  rel.Relation,
			Object:    rel.Object,
			Namespace: rel.Namespace,
			Subject:   &relationtuple.SubjectID{ID: "not " + user.ID},
		}, 0)
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
		require.NoError(t, reg.RelationTupleManager().WriteRelationTuples(ctx, &access, &user))

		e := check.NewEngine(reg)

		res, err := e.SubjectIsAllowed(ctx, &relationtuple.InternalRelationTuple{
			Relation: access.Relation,
			Object:   object,
			Subject:  user.Subject,
		}, 0)
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
		require.NoError(t, reg.RelationTupleManager().WriteRelationTuples(ctx, &readDiary, &user))

		e := check.NewEngine(reg)

		res, err := e.SubjectIsAllowed(ctx, &relationtuple.InternalRelationTuple{
			Relation:  readDiary.Relation,
			Object:    diaryEntry,
			Namespace: diaryNamespace,
			Subject:   user.Subject,
		}, 0)
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
		require.NoError(t, reg.RelationTupleManager().WriteRelationTuples(ctx, &writeRel, &orgOwnerRel, &userMembershipRel))

		e := check.NewEngine(reg)

		// user can write object
		res, err := e.SubjectIsAllowed(ctx, &relationtuple.InternalRelationTuple{
			Namespace: someNamespace,
			Relation:  writeRel.Relation,
			Object:    object,
			Subject:   &user,
		}, 0)
		require.NoError(t, err)
		assert.True(t, res)

		// user is member of the organization
		res, err = e.SubjectIsAllowed(ctx, &relationtuple.InternalRelationTuple{
			Namespace: orgNamespace,
			Relation:  orgMembers.Relation,
			Object:    organization,
			Subject:   &user,
		}, 0)
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
		require.NoError(t, reg.RelationTupleManager().WriteRelationTuples(ctx, &parent, &directoryAccess))

		e := check.NewEngine(reg)

		res, err := e.SubjectIsAllowed(ctx, &relationtuple.InternalRelationTuple{
			Relation: directoryAccess.Relation,
			Object:   file,
			Subject:  &user,
		}, 0)
		require.NoError(t, err)
		assert.False(t, res)
	})

	t.Run("case=subject id next to subject set", func(t *testing.T) {
		namesp, obj, org, directOwner, indirectOwner, ownerRel, memberRel := "namesp", "obj", "org", "u1", "u2", "owner", "member"

		reg := newDepsProvider(t, []*namespace.Namespace{
			{Name: namesp, ID: 1},
		})
		require.NoError(t, reg.RelationTupleManager().WriteRelationTuples(
			ctx,
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

		res, err := e.SubjectIsAllowed(ctx, &relationtuple.InternalRelationTuple{
			Namespace: namesp,
			Object:    obj,
			Relation:  ownerRel,
			Subject:   &relationtuple.SubjectID{ID: directOwner},
		}, 0)
		require.NoError(t, err)
		assert.True(t, res)

		res, err = e.SubjectIsAllowed(ctx, &relationtuple.InternalRelationTuple{
			Namespace: namesp,
			Object:    obj,
			Relation:  ownerRel,
			Subject:   &relationtuple.SubjectID{ID: indirectOwner},
		}, 0)
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
			require.NoError(t, reg.RelationTupleManager().WriteRelationTuples(ctx, &relationtuple.InternalRelationTuple{
				Namespace: namesp,
				Object:    obj,
				Relation:  access,
				Subject:   &relationtuple.SubjectID{ID: user},
			}))
		}

		e := check.NewEngine(reg)

		for i, user := range users {
			t.Run("user="+user, func(t *testing.T) {
				t.Skip() // TODO pagination
				allowed, err := e.SubjectIsAllowed(ctx, &relationtuple.InternalRelationTuple{
					Namespace: namesp,
					Object:    obj,
					Relation:  access,
					Subject:   &relationtuple.SubjectID{ID: user},
				}, 0)
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
			require.NoError(t, reg.RelationTupleManager().WriteRelationTuples(ctx, &relationtuple.InternalRelationTuple{
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
			require.NoError(t, reg.RelationTupleManager().WriteRelationTuples(ctx, &relationtuple.InternalRelationTuple{
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
			allowed, err := e.SubjectIsAllowed(ctx, req, 0)
			require.NoError(t, err)
			assert.True(t, allowed, req.String())
		}
	})

	t.Run("case=circular tuples", func(t *testing.T) {
		sendlingerTor, odeonsplatz, centralStation, connected, namesp := "Sendlinger Tor", "Odeonsplatz", "Central Station", "connected", "munich transport"

		reg := newDepsProvider(t, []*namespace.Namespace{{Name: namesp}})

		require.NoError(t, reg.RelationTupleManager().WriteRelationTuples(ctx, []*relationtuple.InternalRelationTuple{
			{
				Namespace: namesp,
				Object:    sendlingerTor,
				Relation:  connected,
				Subject: &relationtuple.SubjectSet{
					Namespace: namesp,
					Object:    odeonsplatz,
					Relation:  connected,
				},
			},
			{
				Namespace: namesp,
				Object:    odeonsplatz,
				Relation:  connected,
				Subject: &relationtuple.SubjectSet{
					Namespace: namesp,
					Object:    centralStation,
					Relation:  connected,
				},
			},
			{
				Namespace: namesp,
				Object:    centralStation,
				Relation:  connected,
				Subject: &relationtuple.SubjectSet{
					Namespace: namesp,
					Object:    sendlingerTor,
					Relation:  connected,
				},
			},
		}...))

		e := check.NewEngine(reg)

		stations := []string{sendlingerTor, odeonsplatz, centralStation}
		res, err := e.SubjectIsAllowed(ctx, &relationtuple.InternalRelationTuple{
			Namespace: namesp,
			Object:    stations[0],
			Relation:  connected,
			Subject: &relationtuple.SubjectID{
				ID: stations[2],
			},
		}, 0)
		require.NoError(t, err)
		assert.False(t, res)
	})
}
