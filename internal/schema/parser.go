package schema

import (
	"fmt"
	"strconv"

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
			p.addFatal(fmt.Errorf("fatal: %s", item.Val))
		case itemKeywordClass:
			p.parseClass()
		}
	}

	return p.namespaces, p.errors
}

func (p *parser) addFatal(err error) {
	p.errors = append(p.errors, err)
	p.fatal = true
}
func (p *parser) addErr(err error) {
	p.errors = append(p.errors, err)
}
func (p *parser) expect(typ itemType) (value string) {
	if p.fatal {
		return
	}
	item := p.lexer.nextItem()
	if item.Typ != typ {
		p.addFatal(fmt.Errorf("expected %d, got %d(%q)", typ, item.Typ, item.Val))
		return
	}
	return item.Val
}
func (p *parser) match(tokens ...interface{}) (matches map[string]string) {
	matches = make(map[string]string)
	for _, token := range tokens {
		item := p.lexer.nextItem()
		switch token := token.(type) {
		case string:
			if item.Val != token {
				p.addFatal(fmt.Errorf("expected %q, got %q", token, item.Val))
				return
			}
		case *string:
			if item.Typ != itemIdentifier && item.Typ != itemStringLiteral {
				p.addFatal(fmt.Errorf("expected identifier, got %d", item.Typ))
				return
			}
			*token = item.Val
		}
	}
	return
}

// parseClass parses a class. The "class" token was already consumed.
func (p *parser) parseClass() {
	var name string
	p.match(&name, "implements", "Namespace", "{")
	p.namespace = namespace{ns: ns{Name: name}}

	for {
		switch item := p.lexer.nextItem(); item.Typ {
		case itemBraceRight:
			p.namespaces = append(p.namespaces, p.namespace)
			return
		case itemKeywordMetadata:
			p.parseMetadata()
		case itemKeywordRelated:
			p.parseRelated()
		case itemKeywordPermits:
			p.parsePermits()
		default:
			p.addFatal(fmt.Errorf("expected 'metadata', 'permits' or 'related', got %q", item.Val))
			return
		}
	}
}

func (p *parser) parseMetadata() {
	var id string
	p.match("=", "{", "id", ":", &id, "}")
	if parsed, err := strconv.ParseUint(id, 10, 32); err != nil {
		p.addErr(fmt.Errorf("invalid ID: %w", err))
	} else {
		p.namespace.ID = int32(parsed)
	}
}

func (p *parser) parseRelated() {
	p.match(":", "{")
	for {
		switch item := p.lexer.nextItem(); item.Typ {
		case itemBraceRight:
			return
		case itemIdentifier:
			relation := item.Val
			p.match(":", any, "[", "]")
			p.namespace.Relations = append(p.namespace.Relations, ast.Relation{
				Name: relation,
			})
		default:
			p.addFatal(fmt.Errorf("expected identifier or '}', got %q", item.Val))
			return
		}
	}
}

func (p *parser) parsePermits() {
	p.match("=", "{")
	for {
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
				p.addFatal(fmt.Errorf("expected ':' or ')', got %q", item.Val))
			}

			// parse optional return type annotation
			switch item := p.lexer.nextItem(); item.Typ {
			case itemOperatorColon:
				p.match("boolean", "=>")
			case itemOperatorArrow:
			default:
				p.addFatal(fmt.Errorf("expected ':' or '=>', got %q", item.Val))
			}

			relation := ast.Relation{Name: permission}
			defer func() {
				p.namespace.Relations = append(p.namespace.Relations, relation)
			}()

		permissionLoop:
			for {
				var relationName string
				p.match("this", ".", "related", ".", &relationName, ".")

				switch p.expect(itemIdentifier) {
				case "some": // tuple to userset
					var arg, obj, computedUsersetRel string
					p.match(
						"(", &arg, "=>",
						&obj, ".", "permits", ".", &computedUsersetRel,
						"(", "ctx", ")", ")")
					if arg != obj {
						p.addErr(fmt.Errorf("unexpected object: %s", obj))
					}
					addRewrite(&relation,
						ast.TupleToUserset{
							Relation:                relationName,
							ComputedUsersetRelation: computedUsersetRel,
						},
					)
				case "includes": // computed userset
					p.match("(", "ctx", ".", "subject", ")")
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
					p.addFatal(fmt.Errorf("expected ',' or '||', got %q", item.Val))
					return
				}
			}

		default:
			p.addFatal(fmt.Errorf("expected identifier or '}', got %q", item.Val))
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
