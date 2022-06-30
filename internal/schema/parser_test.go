package schema

import (
	"testing"

	"github.com/ory/x/snapshotx"

	"github.com/ory/keto/internal/namespace/ast"
)

func TestParser(t *testing.T) {
	t.Run("suite=snapshots", func(t *testing.T) {
		cases := []struct {
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
	  viewers: (User | SubjectSet<Group, "members">)[]
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
		this.related.parents.traverse((p) =>
		  p.related.viewers.includes(ctx.subject),
		) ||
		this.related.parents.traverse(p => p.permits.view(ctx)) ||
		this.related.viewers.includes(ctx.subject) ||
		this.related.owners.includes(ctx.subject),
  
	  edit: (ctx: Context) => this.related.owners.includes(ctx.subject),
  
	  rename: (ctx: Context) =>
		this.related.siblings.traverse(s => s.permits.edit(ctx)),
	}
  }
`},
		}

		for _, tc := range cases {
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
