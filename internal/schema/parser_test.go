package schema

import (
	"testing"

	"github.com/ory/x/snapshotx"
	"github.com/stretchr/testify/assert"

	"github.com/ory/keto/internal/namespace/ast"
)

var parserTestCases = []struct {
	name, input string
}{
	{"full example", `
  class User implements Namespace {
	related: {
	  manager: User[]
	}
  }
  
  class Group implements Namespace {
	related: {
	  members: (User | Group)[]
	}
  }
  
  class Folder implements Namespace {
	related: {
	  parents: File[]
	  viewers: SubjectSet<Group, "members">[]
	}
  
	permits = {
	  view: (ctx: Context): boolean => this.related.viewers.includes(ctx.subject),
	}
  }
  
  class File implements Namespace {
	related: {
	  parents: (File | Folder)[]
	  viewers: (User | SubjectSet<Group, "members">)[]
	  owners: (User | SubjectSet<Group, "members">)[]
	  siblings: File[]
	}
  
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
  
	  rename: (ctx: Context) =>
		this.related.siblings.traverse(s => s.permits.edit(ctx)),
	}
  }
`},
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
}

func FuzzParser(f *testing.F) {
	for _, tc := range lexerTestCases {
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
		input, expected *ast.UsersetRewrite
	}{
		{"empty", nil, nil},
		{
			name: "merge all unions",
			input: &ast.UsersetRewrite{
				Operation: ast.SetOperationUnion,
				Children: ast.Children{
					&ast.ComputedUserset{Relation: "A"},
					&ast.UsersetRewrite{
						Children: ast.Children{
							&ast.ComputedUserset{Relation: "B"},
							&ast.UsersetRewrite{
								Children: ast.Children{
									&ast.ComputedUserset{Relation: "C"},
									&ast.UsersetRewrite{
										Children: ast.Children{
											&ast.ComputedUserset{Relation: "D"},
										},
									},
								},
							},
						},
					},
				},
			},
			expected: &ast.UsersetRewrite{
				Children: ast.Children{
					&ast.ComputedUserset{Relation: "A"},
					&ast.ComputedUserset{Relation: "B"},
					&ast.ComputedUserset{Relation: "C"},
					&ast.ComputedUserset{Relation: "D"},
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
