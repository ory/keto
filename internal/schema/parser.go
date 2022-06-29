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

func (p *parser) parse() ([]namespace, []error) {
loop:
	for !p.fatal {
		switch item := p.lexer.nextItem(); item.Typ {
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
	item := p.lexer.nextItem()
	if item.Typ != typ {
		p.addFatal(item, "expected %d, got %d(%q)", typ, item.Typ, item.Val)
		return
	}
	return item.Val
}
func (p *parser) match(tokens ...interface{}) (ok bool) {
	if p.fatal {
		return false
	}

	for _, token := range tokens {
		item := p.lexer.nextItem()
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

// parseClass parses a class. The "class" token was already consumed.
func (p *parser) parseClass() {
	var name string
	p.match(&name, "implements", "Namespace", "{")
	p.namespace = namespace{ns: ns{Name: name}}

	for !p.fatal {
		switch item := p.lexer.nextItem(); {
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
		switch item := p.lexer.nextItem(); item.Typ {
		case itemBraceRight:
			return
		case itemIdentifier:
			relation := item.Val
			p.match(":")
			switch item := p.lexer.nextItem(); item.Typ {
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
		switch item := p.lexer.nextItem(); item.Typ {
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
		switch item := p.lexer.nextItem(); item.Typ {

		case itemBraceRight:
			return

		case itemIdentifier:
			permission := item.Val
			p.match(":", "(", "ctx")

			// parse optional type annotation of 'ctx' argument
			switch item := p.lexer.nextItem(); item.Typ {
			case itemOperatorColon:
				p.match("Context", ")")
			case itemParenRight:
			default:
				p.addFatal(item, "expected ':' or ')', got %q", item.Val)
			}

			// parse optional return type annotation
			switch item := p.lexer.nextItem(); item.Typ {
			case itemOperatorColon:
				p.match("boolean", "=>")
			case itemOperatorArrow:
			default:
				p.addFatal(item, "expected ':' or '=>', got %q", item.Val)
			}

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

				switch item := p.lexer.nextItem(); item.Typ {
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
