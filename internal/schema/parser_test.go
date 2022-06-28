package schema

import (
	"testing"

	"github.com/ory/x/snapshotx"
)

func TestParser(t *testing.T) {
	t.Run("suite=snapshots", func(t *testing.T) {
		cases := []struct {
			name, input string
		}{
			{"full namespace", `
class User implements Namespace {
	metadata = {
		id: "1"
	}
}
class Document implements Namespace {
	metadata = {
		id: "2"
	}

	related: {
		owners: User[]
		editors: User[]
		viewers: User[]
		parent: Document[]
	}

	permits = {
	  view: (ctx: Context): boolean =>
		this.related.parents.some(p => p.permits.view(ctx)) ||
		  this.related.viewers.includes(ctx.subject) ||
		  this.related.owners.includes(ctx.subject),
  
	  edit: (ctx: Context) => this.related.owners.includes(ctx.subject),
  
	  rename: (ctx: Context) => this.related.siblings.some(s => s.permits.edit(ctx))
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
				snapshotx.SnapshotT(t, ns)
			})
		}
	})
}
