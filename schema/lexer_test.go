// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package schema

import (
	"strings"
	"testing"

	"github.com/ory/x/snapshotx"
)

var lexingErrorTestCases = []struct{ name, input string }{
	{"open comment", "/*"},
	{"open string literal", "'"},
	{"non-token", "ü"},
	{"mixed quotes", `"invalid_literal'`},
}

var lexableTestCases = []struct{ name, input string }{
	{"empty", ""},
	{"import statement", "import { Foo, Bar, Baz } from '@ory/keto-namespace-types'"},
	{"single class", `
class name implements Namespace {
	metadata = {
		id: "123"
	}
}
`},
	{"comments", `
/**/

/** doc comment
 * content
 * more content
 */
class name implements Namespace {
	// Block comment
	/*
	äny ünicöde characterß???
	*/
}
`},
	{"two classes", `
class user implements Namespace { }

class document implements Namespace {

	related: {
		owners: user[]
		editors: user[]
		viewers: user[]
		parent: document[]
	}
}
`},
	{"full class", `
class File implements Namespace {
	related: {
	  parents: File[]
	  viewers: User[]
	  owners: User[]
	  siblings: File[]
	}
  
	permits = {
	  view: (ctx: Context): boolean =>
		this.related.parents.some(p => p.permits.view(ctx)) ||
		  this.related.viewers.includes(ctx.subject) ||
		  !this.related.owners.includes(ctx.subject),
  
	  edit: (ctx: Context) => this.related.owners.includes(ctx.subject),
  
	  rename: (ctx: Context) => this.related.siblings.some(s => s.permits.edit(ctx))
	}
}
`},
	{"subject sets", `
class X implements Namespace {
	related: {
		relation: SubjectSet<R, "someProp">[]
	}
}
`},
}

func TestLexer(t *testing.T) {
	t.Run("suite=snapshots", func(t *testing.T) {
		for _, tc := range lexableTestCases {
			t.Run(tc.name, func(t *testing.T) {
				l := Lex(tc.name, tc.input)
				var items []string
				for {
					item := l.nextItem()
					items = append(items, item.String())

					if item.Typ == itemError {
						t.Fail()
						break
					}
					if item.Typ == itemEOF {
						break
					}
				}
				t.Logf("Tokens:\n%s", strings.Join(items, "\n"))
				snapshotx.SnapshotT(t, items)
			})
		}
	})

	t.Run("suite=errors", func(t *testing.T) {
		for _, tc := range lexingErrorTestCases {
			t.Run(tc.name, func(t *testing.T) {
				l := Lex(tc.name, tc.input)
				var items []string
				for {
					item := l.nextItem()
					items = append(items, item.String())

					if item.Typ == itemError {
						break
					}
					if item.Typ == itemEOF {
						t.Fatal("reached EOF, but expected error")
					}
				}
				l.next()
				t.Logf("Tokens:\n%s", strings.Join(items, "\n"))
			})
		}

	})
}
