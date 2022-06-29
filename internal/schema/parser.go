package schema

import (
	"fmt"

	internalNamespace "github.com/ory/keto/internal/namespace"
	"github.com/ory/keto/internal/namespace/ast"
)

type (
	ns = internalNamespace.Namespace

	namespace struct {
		ns
	}

	parser struct {
		lexer      *lexer      // lexer to get tokens from
		namespaces []namespace // list of parsed namespaces
		namespace  namespace   // current namespace
		errors     []error     // errors encountered during parsing
		fatal      bool        // parser encountered a fatal error
		lookahead  *item       // lookahead token
	}
)

var (
	empty = ""
	any   = &empty // helper to use in match() when we don't care which identifier.
)

func Parse(input string) ([]namespace, []error) {
	p := &parser{
		lexer: Lex("input", input),
	}
	return p.parse()
}

func (p *parser) next() (item item) {
	if p.lookahead != nil {
		item = *p.lookahead
		p.lookahead = nil
	} else {
		item = p.lexer.nextItem()
	}
	return
}

func (p *parser) peek() item {
	if p.lookahead == nil {
		i := p.lexer.nextItem()
		p.lookahead = &i
		return i
	}
	return *p.lookahead
}

func (p *parser) parse() ([]namespace, []error) {
loop:
	for !p.fatal {
		switch item := p.next(); item.Typ {
		case itemEOF:
			break loop
		case itemError:
			p.addFatal(item, "fatal: %s", item.Val)
		case itemKeywordClass:
			p.parseClass()
		}
	}

	return p.namespaces, p.errors
}

func (p *parser) addFatal(item item, format string, a ...interface{}) {
	p.addErr(item, format, a...)
	p.fatal = true
}
func (p *parser) addErr(item item, format string, a ...interface{}) {
	err := &ParseError{
		msg:  fmt.Sprintf(format, a...),
		item: item,
		p:    p,
	}
	p.errors = append(p.errors, err)
}
func (p *parser) expect(typ itemType) (value string) {
	if p.fatal {
		return
	}
	item := p.next()
	if item.Typ != typ {
		p.addFatal(item, "expected %d, got %d(%q)", typ, item.Typ, item.Val)
		return
	}
	return item.Val
}

// match matches for the next tokens in the input. For string arguments, the
// input token must match the given string exactly. For *string arguments, the
// input token must be an identifier, and the value of the identifier will be
// written to the *string.
func (p *parser) match(tokens ...interface{}) (ok bool) {
	if p.fatal {
		return false
	}

	for _, token := range tokens {
		item := p.next()
		switch token := token.(type) {
		case string:
			if item.Val != token {
				p.addFatal(item, "expected %q, got %q", token, item.Val)
				return false
			}
		case *string:
			if item.Typ != itemIdentifier && item.Typ != itemStringLiteral {
				p.addFatal(item, "expected identifier, got %s", item.Typ)
				return false
			}
			*token = item.Val
		}
	}
	return true
}

type itemPredicate func(item) bool

func is(typ itemType) itemPredicate {
	return func(item item) bool {
		return item.Typ == typ
	}
}
func hasValue(val string) itemPredicate {
	return func(item item) bool {
		return item.Val == val
	}
}

// matchIf matches the tokens iff. the predicate is true.
func (p *parser) matchIf(predicate itemPredicate, tokens ...interface{}) (ok bool) {
	if p.fatal {
		return false
	}
	if !predicate(p.peek()) {
		return true
	}
	return p.match(tokens...)
}

// parseClass parses a class. The "class" token was already consumed.
func (p *parser) parseClass() {
	var name string
	p.match(&name, "implements", "Namespace", "{")
	p.namespace = namespace{ns: ns{Name: name}}

	for !p.fatal {
		switch item := p.next(); {
		case item.Typ == itemBraceRight:
			p.namespaces = append(p.namespaces, p.namespace)
			return
		case item.Val == "related":
			p.parseRelated()
		case item.Val == "permits":
			p.parsePermits()
		default:
			p.addFatal(item, "expected 'permits' or 'related', got %q", item.Val)
			return
		}
	}
}

func (p *parser) parseRelated() {
	p.match(":", "{")
	for !p.fatal {
		switch item := p.next(); item.Typ {
		case itemBraceRight:
			return
		case itemIdentifier:
			relation := item.Val
			p.match(":")
			switch item := p.next(); item.Typ {
			case itemIdentifier:
			case itemParenLeft:
				p.parseTypeUnion()
			}
			p.namespace.Relations = append(p.namespace.Relations, ast.Relation{
				Name: relation,
			})
			p.match("[", "]")
		default:
			p.addFatal(item, "expected identifier or '}', got %q", item.Val)
			return
		}
	}
}

func (p *parser) parseTypeUnion() {
	for !p.fatal {
		var identifier string
		p.match(&identifier)
		if identifier == "SubjectSet" {
			p.match("<", any, ",", any, ">")
		}
		switch item := p.next(); item.Typ {
		case itemParenRight:
			return
		case itemTypeUnion:
		default:
			p.addFatal(item, "expected '|', got %q", item.Val)
		}
	}
}

func (p *parser) parsePermits() {
	p.match("=", "{")
	for !p.fatal {
		switch item := p.next(); item.Typ {

		case itemBraceRight:
			return

		case itemIdentifier:
			permission := item.Val
			p.match(":", "(", "ctx")
			p.matchIf(is(itemOperatorColon), ":", "Context")
			p.match(")")
			p.matchIf(is(itemOperatorColon), ":", "boolean")
			p.match("=>")

			relation := ast.Relation{Name: permission}
			defer func() {
				p.namespace.Relations = append(p.namespace.Relations, relation)
			}()

		permissionLoop:
			for !p.fatal {
				var relationName string
				p.match("this", ".", "related", ".", &relationName, ".")

				switch p.expect(itemIdentifier) {
				case "traverse": // tuple to userset
					var arg, obj, computedUsersetRel, verb string
					if !p.match("(", &arg, "=>", &obj, ".", &verb) {
						return
					}
					if arg != obj {
						p.addErr(item, "unexpected object: %s", obj)
					}
					switch verb {
					case "related":
						p.match(".", &computedUsersetRel, ".", "includes", "(", "ctx", ".", "subject", ")", ")")
					case "permits":
						p.match(".", &computedUsersetRel, "(", "ctx", ")", ")")
					default:
						p.addFatal(item, "expected 'related' or 'permits', got %q", verb)
						return
					}
					addRewrite(&relation,
						ast.TupleToUserset{
							Relation:                relationName,
							ComputedUsersetRelation: computedUsersetRel,
						},
					)
				case "includes": // computed userset
					if !p.match("(", "ctx", ".", "subject", ")") {
						return
					}
					addRewrite(&relation,
						ast.ComputedUserset{
							Relation: relationName,
						},
					)
				}

				switch item := p.next(); item.Typ {
				case itemOperatorOr:
					continue permissionLoop

				case itemOperatorComma:
					break permissionLoop

				case itemBraceRight:
					return

				default:
					p.addFatal(item, "expected ',' or '||', got %q", item.Val)
					return
				}
			}

		default:
			p.addFatal(item, "expected identifier or '}', got %q", item.Val)
			return
		}
	}
}

func addRewrite(r *ast.Relation, child ast.Child) {
	if r.UsersetRewrite == nil {
		r.UsersetRewrite = &ast.UsersetRewrite{}
	}
	r.UsersetRewrite.Children = append(r.UsersetRewrite.Children, child)
}
