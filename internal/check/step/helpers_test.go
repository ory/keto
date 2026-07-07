// Copyright © 2026 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package step_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ory/keto/internal/check/step"
	"github.com/ory/keto/internal/namespace/ast"
	"github.com/ory/keto/internal/testhelpers"
)

func TestAllowsDirectMember(t *testing.T) {
	t.Parallel()

	var (
		rewriteRel    = &ast.Relation{SubjectSetRewrite: &ast.SubjectSetRewrite{}}
		directRel     = &ast.Relation{Name: "users", Types: []ast.RelationType{{Namespace: "User"}}}                         // users: User[]
		anotherRel    = &ast.Relation{Name: "keys", Types: []ast.RelationType{{Namespace: "ApiKey"}}}                        // keys: ApiKey[]
		subjectSetRel = &ast.Relation{Name: "members", Types: []ast.RelationType{{Namespace: "Group", Relation: "members"}}} // members: SubjectSet<Group,"members">
	)

	tests := []struct {
		name       string
		relation   *ast.Relation
		subjectStr string
		want       bool
	}{
		{
			name:       "SubjectID always returns true (conservative)",
			relation:   directRel,
			subjectStr: "fixedID",
			want:       true,
		},
		{
			name:       "SubjectID with nil relation returns true (conservative)",
			relation:   nil,
			subjectStr: "fixedID",
			want:       true,
		},
		{
			name:       "SubjectID with SubjectSetRewrite returns false (computed relation has no direct tuples)",
			relation:   rewriteRel,
			subjectStr: "fixedID",
			want:       false,
		},
		{
			name:       "SubjectSet with nil relation returns false",
			relation:   nil,
			subjectStr: "User:u1",
			want:       false,
		},
		{
			name:       "SubjectSet with SubjectSetRewrite returns false",
			relation:   rewriteRel,
			subjectStr: "User:u1",
			want:       false,
		},
		{
			name:       "SubjectSet namespace matches direct type",
			relation:   directRel,
			subjectStr: "User:u1",
			want:       true,
		},
		{
			name:       "subject namespace does not match any direct type",
			relation:   anotherRel,
			subjectStr: "User:u1",
			want:       false,
		},
		{
			name:       "relation doesn't even have any User type",
			relation:   subjectSetRel,
			subjectStr: "User:u1",
			want:       false,
		},
		{
			name:       "subject is a single object, but relation is a SubjectSet<Group,'members'>, so no direct match allowed",
			relation:   subjectSetRel,
			subjectStr: "Group:g1",
			want:       false,
		},
		{
			name:       "subject is group#member, and relation is also SubjectSet<Group,'members'>, so direct match is allowed",
			relation:   subjectSetRel,
			subjectStr: "Group:g1#members",
			want:       true,
		},
		{
			name:       "subject is group#member, but relation is a Direct",
			relation:   &ast.Relation{Name: "groups", Types: []ast.RelationType{{Namespace: "Group"}}},
			subjectStr: "Group:g1#members",
			want:       false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			subject := testhelpers.SubjectFromString(t, tc.subjectStr)
			assert.Equal(t, tc.want, step.AllowsDirectMember(tc.relation, subject))
		})
	}
}
