// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package schema

import (
	"testing"

	"github.com/ory/x/snapshotx"
	"github.com/stretchr/testify/assert"

	"github.com/ory/keto/internal/namespace/ast"
)

var parserErrorTestCases = []struct{ name, input string }{
	{"lexer error", "/* unclosed comment"},
	{"syntax and type errors",
		`
  class File implements Namespace {
	related: {
	  parents: (File | Folder)[]
	  viewers: (User | SubjectSet<Group, "members">)[]
	  owners: (User | SubjectSet<Group, "members">)[]
	  siblings: File[]
	}

	SYNTAX ERROR
  
	// Some comment
	permits = {
	  view: (ctx: Context): boolean =>
	    (
		this.related.parents.traverse((p) =>
		  p.related.viewers.includes(ctx.subject),
		) &&
		this.related.parents.traverse(p => p.permits.view(ctx)) ) ||
		(this.related.viewers.includes(ctx.subject) ||
		this.related.viewers.includes(ctx.subject) ||
		this.related.viewers.includes(ctx.subject) ) ||
		this.related.owners.includes(ctx.subject),
  
	  edit: (ctx: Context) => this.related.owners.includes(ctx.subject),

	  not: (ctx: Context) => !this.related.owners.includes(ctx.subject),
  
	  rename: (ctx: Context) =>
		this.related.siblings.traverse(s => s.permits.edit(ctx)),
	}
  }
`},
	{"parser error", `
class Resource implements Namespace {
  permits = {
    update: (ctx: Context) => ||
      this.related.annotators.traverse((role) => role.related.member.includes(ctx.subject)) ||
      this.related.supervisors.traverse((role) => role.related.member.includes(ctx.subject)),
`},
}

var parserTestCases = []struct {
	name, input string
}{
	{"full example", `
  import { Namespace, SubjectSet, FooBar, Anything } from '@ory/keto-namespace-types'

  class User implements Namespace {
	related: {
	  manager: User[];
	}
  }
  
  class Group implements Namespace {
	related: {
	  members: (User | Group)[];
	};
  }
  
  class Folder implements Namespace {
	related: {
	  parents: Array<File>
	  viewers: Array<SubjectSet<Group, "members">>
	}
  
	permits = {
	  view: (ctx: Context): boolean => this.related.viewers.includes(ctx.subject),
	}
  }
  
  class File implements Namespace {
	related: {
	  parents: Array<File | Folder>
	  viewers: (User | SubjectSet<Group, "members">)[]
	  "owners": (User | SubjectSet<Group, "members">)[]
	  siblings: File[]
	}
  
	// Some comment
	permits = {
	  view: (ctx: Context): boolean =>
	    (
		this.related.parents.traverse((p) /* comment */ =>
		  p.related.viewers.includes(ctx.subject),
		) && // comment
		this.related.parents.traverse(p => p.permits.view(ctx)) ) ||
		(this.related.viewers.includes(ctx.subject) || // some comment
		this.related.viewers.includes(ctx.subject) || /* another comment */
		this.related.viewers.includes(ctx.subject) ) ||
		this.related.owners.includes(ctx.subject),
  
	  'edit': (ctx: Context) => this.related.owners.includes(ctx.subject),

	  not: (ctx: Context) => !this.related.owners.includes(ctx.subject),
  
	  rename: (ctx: Context) =>
		this.related.siblings.traverse(s => s.permits.edit(ctx)),
	}
  }
`}, {"advanced typescript syntax",
		`
import { Namespace, SubjectSet, Context } from '@ory/keto-namespace-types';

class Role implements Namespace {
  related: {
    member: Role[]
  }
}

class Resource implements Namespace {
  related: {
    admins: SubjectSet<Role, 'member'>[],
    supervisors: SubjectSet<Role, 'member'>[],
    annotators: SubjectSet<Role, 'member'>[],
    medicalAnnotators: SubjectSet<Role, 'member'>[],
  };

  permits = {
    read: (ctx: Context) => this.related.admins.traverse((role) => role.related.member.includes(ctx.subject)) ||
      this.related.annotators.traverse((role) => role.related.member.includes(ctx.subject)) ||
      this.related.medicalAnnotators.traverse((role) => role.related.member.includes(ctx.subject)) ||
      this.related.supervisors.traverse((role) => role.related.member.includes(ctx.subject)),

    comment: (ctx: Context) => this.permits.read(ctx),

    update: (ctx: Context) => this.related.admins.traverse((role) => role.related.member.includes(ctx.subject)) ||
      this.related.annotators.traverse((role) => role.related.member.includes(ctx.subject)) ||
      this.related.medicalAnnotators.traverse((role) => role.related.member.includes(ctx.subject)) ||
      this.related.supervisors.traverse((role) => role.related.member.includes(ctx.subject)),

    create: (ctx: Context) => this.related.admins.traverse((role) => role.related.member.includes(ctx.subject)) ||
      this.related.annotators.traverse((role) => role.related.member.includes(ctx.subject)) ||
      this.related.supervisors.traverse((role) => role.related.member.includes(ctx.subject)),

    approve: (ctx: Context) => this.related.admins.traverse((role) => role.related.member.includes(ctx.subject)) ||
      this.related.supervisors.traverse((role) => role.related.member.includes(ctx.subject)),

    delete: (ctx: Context) => this.related.admins.traverse((role) => role.related.member.includes(ctx.subject)) ||
      this.related.supervisors.traverse((role) => role.related.member.includes(ctx.subject)),
  };
}
`}, {"quoted property names", `
class Resource implements Namespace {
  related: {
    "scope.relation": Resource[]
  }
  permits = {
    "scope.action_0": (ctx: Context) => this.related["scope.relation"].traverse((r) => r.permits["scope.action_1"](ctx)),
    "scope.action_1": (ctx: Context) => this.related["scope.relation"].traverse((r) => r.related["scope.relation"].includes(ctx.subject)),
    "scope.action_2": (ctx: Context) => this.permits["scope.action_0"](ctx),
  }
}`},
}

func TestParser(t *testing.T) {
	t.Run("suite=snapshots", func(t *testing.T) {
		for _, tc := range parserTestCases {
			t.Run(tc.name, func(t *testing.T) {
				ns, errs := Parse(tc.input)
				if len(errs) > 0 {
					for _, err := range errs {
						t.Error(err)
					}
				}
				t.Logf("namespaces:\n%+v", ns)
				nsMap := make(map[string][]ast.Relation)
				for _, n := range ns {
					nsMap[n.Name] = n.Relations
				}
				snapshotx.SnapshotT(t, nsMap)
			})
		}
	})

	t.Run("suite=errors", func(t *testing.T) {
		for _, tc := range parserErrorTestCases {
			t.Run(tc.name, func(t *testing.T) {
				_, errs := Parse(tc.input)
				assert.Len(t, errs, 1)
			})
		}
	})
}

func FuzzParser(f *testing.F) {
	for _, tc := range lexableTestCases {
		f.Add(tc.input)
	}
	for _, tc := range parserTestCases {
		f.Add(tc.input)
	}

	f.Fuzz(func(_ *testing.T, input string) {
		Parse(input)
	})
}

func Test_simplify(t *testing.T) {
	testCases := []struct {
		name            string
		input, expected *ast.SubjectSetRewrite
	}{
		{"empty", nil, nil},
		{
			name: "merge all unions",
			input: &ast.SubjectSetRewrite{
				Operation: ast.OperatorOr,
				Children: ast.Children{
					&ast.ComputedSubjectSet{Relation: "A"},
					&ast.SubjectSetRewrite{
						Children: ast.Children{
							&ast.ComputedSubjectSet{Relation: "B"},
							&ast.SubjectSetRewrite{
								Children: ast.Children{
									&ast.ComputedSubjectSet{Relation: "C"},
									&ast.SubjectSetRewrite{
										Children: ast.Children{
											&ast.ComputedSubjectSet{Relation: "D"},
										},
									},
								},
							},
						},
					},
				},
			},
			expected: &ast.SubjectSetRewrite{
				Children: ast.Children{
					&ast.ComputedSubjectSet{Relation: "A"},
					&ast.ComputedSubjectSet{Relation: "B"},
					&ast.ComputedSubjectSet{Relation: "C"},
					&ast.ComputedSubjectSet{Relation: "D"},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, simplifyExpression(tc.input))
		})
	}
}
