// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package check_test

import (
	"context"
	"testing"

	"github.com/gofrs/uuid"
	"github.com/ory/herodot"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ory/keto/internal/check"
	"github.com/ory/keto/internal/check/checkgroup"
	"github.com/ory/keto/internal/driver"
	"github.com/ory/keto/internal/driver/config"
	"github.com/ory/keto/internal/namespace"
	"github.com/ory/keto/internal/namespace/ast"
	"github.com/ory/keto/internal/relationtuple"
	"github.com/ory/keto/internal/testhelpers"
	"github.com/ory/keto/internal/x"
	"github.com/ory/keto/ketoapi"
)

func TestEngine(t *testing.T) {
	ctx := context.Background()

	t.Run("respects max depth", func(t *testing.T) {
		// "user" has relation "access" through being an "owner" through being an "admin"
		// which requires at least 2 units of depth. If max-depth is 2 then we hit max-depth
		reg := driver.NewSqliteTestRegistry(t, false)

		// "user" has relation "access" through being an "owner" through being
		// an "admin" which requires at least 2 units of depth. If max-depth is
		// 2 then we hit max-depth
		testhelpers.MapAndInsertTuplesFromString(t, reg, []string{
			"test:object#admin@user",
			"test:object#owner@test:object#admin",
			"test:object#access@test:object#owner",
		})

		e := check.NewEngine(reg)

		userHasAccess := testhelpers.TupleFromString(t, "test:object#access@user")

		// global max-depth defaults to 5
		assert.Equal(t, reg.Config(ctx).MaxReadDepth(), 5)

		// req max-depth takes precedence, max-depth=2 is not enough
		res, err := e.CheckIsMember(ctx, userHasAccess, 2)
		require.NoError(t, err)
		assert.False(t, res)

		// req max-depth takes precedence, max-depth=3 is enough
		res, err = e.CheckIsMember(ctx, userHasAccess, 3)
		require.NoError(t, err)
		assert.True(t, res)

		// global max-depth takes precedence and max-depth=2 is not enough
		require.NoError(t, reg.Config(ctx).Set(config.KeyLimitMaxReadDepth, 2))
		res, err = e.CheckIsMember(ctx, userHasAccess, 2)
		require.NoError(t, err)
		assert.False(t, res)

		// global max-depth takes precedence and max-depth=3 is enough
		require.NoError(t, reg.Config(ctx).Set(config.KeyLimitMaxReadDepth, 3))
		res, err = e.CheckIsMember(ctx, userHasAccess, 0)
		require.NoError(t, err)
		assert.True(t, res)
	})

	t.Run("direct inclusion", func(t *testing.T) {
		reg := driver.NewSqliteTestRegistry(t, false, driver.WithNamespaces([]*namespace.Namespace{{Name: "n"}, {Name: "u"}}))
		tuples := []string{
			`n:o#r@subject_id`,
			`n:o#r@u:with_relation#r`,
			`n:o#r@u:empty_relation#`,
			`n:o#r@u:missing_relation`,
		}

		testhelpers.MapAndInsertTuplesFromString(t, reg, tuples)
		e := check.NewEngine(reg)

		cases := []struct {
			tuple string
		}{
			{tuple: "n:o#r@subject_id"},
			{tuple: "n:o#r@u:with_relation#r"},

			{tuple: "n:o#r@u:empty_relation"},
			{tuple: "n:o#r@u:empty_relation#"},

			{tuple: "n:o#r@u:missing_relation"},
			{tuple: "n:o#r@u:missing_relation#"},
		}

		for _, tc := range cases {
			t.Run("case="+tc.tuple, func(t *testing.T) {
				res, err := e.CheckIsMember(ctx, testhelpers.TupleFromString(t, tc.tuple), 0)
				require.NoError(t, err)
				assert.True(t, res)
			})
		}
	})

	t.Run("indirect inclusion level 1", func(t *testing.T) {
		reg := driver.NewSqliteTestRegistry(t, false, driver.WithNamespaces([]*namespace.Namespace{{Name: "sofa"}}))

		testhelpers.MapAndInsertTuplesFromString(t, reg, []string{
			"sofa:dust#producer@mark",
			"sofa:dust#have_to_remove@sofa:dust#producer",
		})
		e := check.NewEngine(reg)

		res, err := e.CheckIsMember(ctx, testhelpers.TupleFromString(t, "sofa:dust#have_to_remove@mark"), 0)
		require.NoError(t, err)
		assert.True(t, res)
	})

	t.Run("direct exclusion", func(t *testing.T) {
		user := &relationtuple.SubjectID{ID: uuid.Must(uuid.NewV4())}
		rel := relationtuple.RelationTuple{
			Relation:  "relation",
			Object:    uuid.Must(uuid.NewV4()),
			Namespace: t.Name(),
			Subject:   user,
		}

		reg := driver.NewSqliteTestRegistry(t, false, driver.WithNamespaces([]*namespace.Namespace{{Name: rel.Namespace}}))
		require.NoError(t, reg.RelationTupleManager().WriteRelationTuples(ctx, &rel))

		e := check.NewEngine(reg)

		res, err := e.CheckIsMember(ctx, &relationtuple.RelationTuple{
			Namespace: rel.Namespace,
			Object:    rel.Object,
			Relation:  rel.Relation,
			Subject:   &relationtuple.SubjectID{ID: uuid.Must(uuid.NewV4())},
		}, 0)
		require.NoError(t, err)
		assert.False(t, res)
	})

	t.Run("subject expansion", func(t *testing.T) {
		reg := driver.NewSqliteTestRegistry(t, false, driver.WithNamespaces([]*namespace.Namespace{{
			Name: "n",
			Relations: []ast.Relation{
				{
					Name: "r",
					Types: []ast.RelationType{
						{
							Namespace: "n",
							Relation:  "r",
						},
					},
				},
			},
		}}))
		tuples := []string{
			`n:a#r@n:b#r`,
			`n:b#r@n:c#r`,
			`n:c#r@n:d#r`,
			`n:d#r@u`,
		}

		testhelpers.MapAndInsertTuplesFromString(t, reg, tuples)
		e := check.NewEngine(reg)
		require.NoError(t, reg.Config(ctx).Set(config.KeyLimitMaxReadDepth, 5))

		cases := []struct {
			tuple string
		}{
			{tuple: "n:d#r@u"},
			{tuple: "n:c#r@u"},
			{tuple: "n:b#r@u"},
			{tuple: "n:a#r@u"},
		}

		for _, tc := range cases {
			t.Run("case="+tc.tuple, func(t *testing.T) {
				res, err := e.CheckIsMember(ctx, testhelpers.TupleFromString(t, tc.tuple), 0)
				require.NoError(t, err)
				assert.True(t, res)
			})
		}
	})

	t.Run("wrong object ID", func(t *testing.T) {
		object := uuid.Must(uuid.NewV4())
		access := relationtuple.RelationTuple{
			Relation: "access",
			Object:   object,
			Subject: &relationtuple.SubjectSet{
				Relation: "owner",
				Object:   object,
			},
		}
		user := relationtuple.RelationTuple{
			Relation: "owner",
			Object:   uuid.Must(uuid.NewV4()),
			Subject:  &relationtuple.SubjectID{ID: uuid.Must(uuid.NewV4())},
		}

		reg := driver.NewSqliteTestRegistry(t, false)
		require.NoError(t, reg.RelationTupleManager().WriteRelationTuples(ctx, &access, &user))

		e := check.NewEngine(reg)

		res, err := e.CheckIsMember(ctx, &relationtuple.RelationTuple{
			Relation: access.Relation,
			Object:   object,
			Subject:  user.Subject,
		}, 0)
		require.NoError(t, err)
		assert.False(t, res)
	})

	t.Run("wrong relation name", func(t *testing.T) {
		diaryEntry := uuid.Must(uuid.NewV4())
		diaryNamespace := "diaries"
		// this would be a user-set rewrite
		readDiary := relationtuple.RelationTuple{
			Namespace: diaryNamespace,
			Relation:  "read",
			Object:    diaryEntry,
			Subject: &relationtuple.SubjectSet{
				Relation:  "author",
				Object:    diaryEntry,
				Namespace: diaryNamespace,
			},
		}
		user := relationtuple.RelationTuple{
			Namespace: diaryNamespace,
			Relation:  "not author",
			Object:    diaryEntry,
			Subject:   &relationtuple.SubjectID{ID: uuid.Must(uuid.NewV4())},
		}

		reg := driver.NewSqliteTestRegistry(t, false, driver.WithNamespaces([]*namespace.Namespace{{Name: diaryNamespace}}))
		require.NoError(t, reg.RelationTupleManager().WriteRelationTuples(ctx, &readDiary, &user))

		e := check.NewEngine(reg)

		res, err := e.CheckIsMember(ctx, &relationtuple.RelationTuple{
			Relation:  readDiary.Relation,
			Object:    diaryEntry,
			Namespace: diaryNamespace,
			Subject:   user.Subject,
		}, 0)
		require.NoError(t, err)
		assert.False(t, res)
	})

	t.Run("indirect inclusion level 2", func(t *testing.T) {
		reg := driver.NewSqliteTestRegistry(t, false, driver.WithNamespaces([]*namespace.Namespace{
			{Name: "obj"},
			{Name: "org"},
		}))

		testhelpers.MapAndInsertTuplesFromString(t, reg, []string{
			"obj:object#write@obj:object#owner",
			"obj:object#owner@org:organization#member",
			"org:organization#member@user",
		})

		e := check.NewEngine(reg)

		// user can write object
		res, err := e.CheckIsMember(ctx, testhelpers.TupleFromString(t, "obj:object#write@user"), 0)
		require.NoError(t, err)
		assert.True(t, res)

		// user is member of the organization
		res, err = e.CheckIsMember(ctx, testhelpers.TupleFromString(t, "org:organization#member@user"), 0)
		require.NoError(t, err)
		assert.True(t, res)
	})

	t.Run("rejects transitive relation", func(t *testing.T) {
		// (file) <--parent-- (directory) <--access-- [user]
		//
		// note the missing access relation from "users who have access to directory also have access to files inside of the directory"
		// as we don't know how to interpret the "parent" relation, there would have to be a userset rewrite to allow access
		// to files when you have access to the parent

		file := uuid.Must(uuid.NewV4())
		directory := uuid.Must(uuid.NewV4())
		user := relationtuple.SubjectID{ID: uuid.Must(uuid.NewV4())}

		parent := relationtuple.RelationTuple{
			Relation: "parent",
			Object:   file,
			Subject: &relationtuple.SubjectSet{ // <- this is only an object, but this is allowed as a userset can have the "..." relation which means any relation
				Object: directory,
			},
		}
		directoryAccess := relationtuple.RelationTuple{
			Relation: "access",
			Object:   directory,
			Subject:  &user,
		}

		reg := driver.NewSqliteTestRegistry(t, false, driver.WithNamespaces([]*namespace.Namespace{
			{Name: "2"},
		}))
		require.NoError(t, reg.RelationTupleManager().WriteRelationTuples(ctx, &parent, &directoryAccess))

		e := check.NewEngine(reg)

		res, err := e.CheckIsMember(ctx, &relationtuple.RelationTuple{
			Relation: directoryAccess.Relation,
			Object:   file,
			Subject:  &user,
		}, 0)
		require.NoError(t, err)
		assert.False(t, res)
	})

	t.Run("case=subject id next to subject set", func(t *testing.T) {
		namesp, obj, org, directOwner, indirectOwner, ownerRel, memberRel := "39231", uuid.Must(uuid.NewV4()), uuid.Must(uuid.NewV4()), uuid.Must(uuid.NewV4()), uuid.Must(uuid.NewV4()), "owner", "member"

		reg := driver.NewSqliteTestRegistry(t, false, driver.WithNamespaces([]*namespace.Namespace{{Name: namesp}}))
		require.NoError(t, reg.RelationTupleManager().WriteRelationTuples(
			ctx,
			&relationtuple.RelationTuple{
				Namespace: namesp,
				Object:    obj,
				Relation:  ownerRel,
				Subject:   &relationtuple.SubjectID{ID: directOwner},
			},
			&relationtuple.RelationTuple{
				Namespace: namesp,
				Object:    obj,
				Relation:  ownerRel,
				Subject: &relationtuple.SubjectSet{
					Namespace: namesp,
					Object:    org,
					Relation:  memberRel,
				},
			},
			&relationtuple.RelationTuple{
				Namespace: namesp,
				Object:    org,
				Relation:  memberRel,
				Subject:   &relationtuple.SubjectID{ID: indirectOwner},
			},
		))

		e := check.NewEngine(reg)

		res, err := e.CheckIsMember(ctx, &relationtuple.RelationTuple{
			Namespace: namesp,
			Object:    obj,
			Relation:  ownerRel,
			Subject:   &relationtuple.SubjectID{ID: directOwner},
		}, 0)
		require.NoError(t, err)
		assert.True(t, res)

		res, err = e.CheckIsMember(ctx, &relationtuple.RelationTuple{
			Namespace: namesp,
			Object:    obj,
			Relation:  ownerRel,
			Subject:   &relationtuple.SubjectID{ID: indirectOwner},
		}, 0)
		require.NoError(t, err)
		assert.True(t, res)
	})

	t.Run("case=wide tuple graph", func(t *testing.T) {
		namesp, obj, access, member, users, orgs := "9234", uuid.Must(uuid.NewV4()), "access", "member", x.UUIDs(4), x.UUIDs(2)

		reg := driver.NewSqliteTestRegistry(t, false, driver.WithNamespaces([]*namespace.Namespace{{Name: namesp}}))

		for _, org := range orgs {
			require.NoError(t, reg.RelationTupleManager().WriteRelationTuples(ctx, &relationtuple.RelationTuple{
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
			require.NoError(t, reg.RelationTupleManager().WriteRelationTuples(ctx, &relationtuple.RelationTuple{
				Namespace: namesp,
				Object:    orgs[i%len(orgs)],
				Relation:  member,
				Subject:   &relationtuple.SubjectID{ID: user},
			}))
		}

		e := check.NewEngine(reg)

		for _, user := range users {
			req := &relationtuple.RelationTuple{
				Namespace: namesp,
				Object:    obj,
				Relation:  access,
				Subject:   &relationtuple.SubjectID{ID: user},
			}
			allowed, err := e.CheckIsMember(ctx, req, 0)
			require.NoError(t, err)
			assert.Truef(t, allowed, "%+v", req)
		}
	})

	t.Run("case=circular tuples", func(t *testing.T) {
		sendlingerTor, odeonsplatz, centralStation, connected, namesp := uuid.NewV5(uuid.Nil, "Sendlinger Tor"), uuid.NewV5(uuid.Nil, "Odeonsplatz"), uuid.NewV5(uuid.Nil, "Central Station"), "connected", "7743"

		reg := driver.NewSqliteTestRegistry(t, false, driver.WithNamespaces([]*namespace.Namespace{{Name: namesp}}))

		require.NoError(t, reg.RelationTupleManager().WriteRelationTuples(ctx, []*relationtuple.RelationTuple{
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

		stations := []uuid.UUID{sendlingerTor, odeonsplatz, centralStation}
		res, err := e.CheckIsMember(ctx, &relationtuple.RelationTuple{
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

	t.Run("case=strict mode", func(t *testing.T) {
		reg := driver.NewSqliteTestRegistry(t, false,
			driver.WithOPL(ProjectOPLConfig),
			driver.WithConfig(config.KeyNamespacesExperimentalStrictMode, true))

		testhelpers.MapAndInsertTuplesFromString(t, reg, []string{
			"Project:abc#owner@User:1",
			"Project:abc#owner@User1",
			// The following tuples are ignored in strict mode
			"Project:abc#isOwner@User:isOwner",
			"Project:abc#readProject@readProjectUser",
			"Project:abc#readProject@User:ReadProject",
		})

		e := check.NewEngine(reg)

		for _, sub := range []string{"readProjectUser", "User:ReadProject", "User:isOwner"} {
			// These checks should return false, even though the exact tuple is in the db.
			res, err := e.CheckIsMember(ctx, testhelpers.TupleFromString(t, "Project:abc#readProject@"+sub), 10)
			require.NoError(t, err)
			assert.False(t, res)
		}

		for _, sub := range []string{"User:1", "User1"} {
			res, err := e.CheckIsMember(ctx, testhelpers.TupleFromString(t, "Project:abc#readProject@"+sub), 10)
			require.NoError(t, err)
			assert.True(t, res)
		}
	})

	t.Run("case=batch check", func(t *testing.T) {
		reg := driver.NewSqliteTestRegistry(t, false,
			driver.WithNamespaces([]*namespace.Namespace{{Name: "test"}}), driver.WithMapperNamespace(testhelpers.CustomMapperNamespace))

		testhelpers.MapAndInsertTuplesFromString(t, reg, []string{
			"test:object#admin@user",
			"test:object#owner@test:object#admin",
			"test:object#access@test:object#owner",
		})

		e := check.NewEngine(reg)

		targetTuples := []*ketoapi.RelationTuple{
			// direct relation
			testhelpers.APITupleFromString(t, "test:object#admin@user"),
			// indirect relation
			testhelpers.APITupleFromString(t, "test:object#owner@user"),
			// indirect relation, greater than max depth
			testhelpers.APITupleFromString(t, "test:object#access@user"),
			// non-existent namespace
			testhelpers.APITupleFromString(t, "test2:object#admin@user"),
			// unknown subject
			testhelpers.APITupleFromString(t, "test:object#admin@user2"),
			// relation via subject set
			testhelpers.APITupleFromString(t, "test:object#access@test:object#owner"),
		}

		// Batch check with low max depth
		results, err := e.BatchCheck(ctx, targetTuples, 2)
		require.NoError(t, err)

		require.Equal(t, checkgroup.IsMember, results[0].Membership)
		require.NoError(t, results[0].Err)
		require.Equal(t, checkgroup.IsMember, results[1].Membership)
		require.NoError(t, results[1].Err)
		require.Equal(t, checkgroup.NotMember, results[2].Membership)
		require.NoError(t, results[2].Err)
		require.Equal(t, checkgroup.MembershipUnknown, results[3].Membership)
		require.EqualError(t, results[3].Err, herodot.ErrNotFound().Error())
		require.Equal(t, checkgroup.NotMember, results[4].Membership)
		require.NoError(t, results[4].Err)
		require.Equal(t, checkgroup.IsMember, results[5].Membership)
		require.NoError(t, results[5].Err)

		// Check with higher max depth and verify the third tuple is now shown as a member
		results, err = e.BatchCheck(ctx, targetTuples, 3)
		require.NoError(t, err)
		require.Equal(t, checkgroup.IsMember, results[2].Membership)
	})
}
