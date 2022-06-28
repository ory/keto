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
func (p *parser) expect(typs ...itemType) (value string) {
	for _, typ := range typs {
		if p.fatal {
			return
		}
		item := p.lexer.nextItem()
		if item.Typ != typ {
			p.addFatal(fmt.Errorf("expected %d, got %d(%q)", typ, item.Typ, item.Val))
			return
		}
		value = item.Val
	}
	return
}

// parseClass parses a class. The "class" token was already consumed.
func (p *parser) parseClass() {
	name := p.expect(itemIdentifier)
	p.expect(itemKeywordImplements)
	if iface := p.expect(itemIdentifier); iface != "Namespace" {
		p.addErr(fmt.Errorf("unexpected 'implements' interface: %s", iface))
	}
	p.expect(itemBraceLeft)
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
	p.expect(itemOperatorAssign)
	p.expect(itemBraceLeft)
	p.expect(itemKeywordID)
	p.expect(itemOperatorColon)
	id := p.expect(itemStringLiteral)
	p.expect(itemBraceRight)
	if parsed, err := strconv.ParseUint(id, 10, 32); err != nil {
		p.addErr(fmt.Errorf("invalid ID: %w", err))
	} else {
		p.namespace.ID = int32(parsed)
	}
}

func (p *parser) parseRelated() {
	p.expect(itemOperatorColon)
	p.expect(itemBraceLeft)
	for {
		switch item := p.lexer.nextItem(); item.Typ {
		case itemBraceRight:
			return
		case itemIdentifier:
			relation := item.Val
			p.expect(itemOperatorColon)
			p.expect(itemIdentifier)
			p.expect(itemBracketLeft)
			p.expect(itemBracketRight)
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
	p.expect(itemOperatorAssign)
	p.expect(itemBraceLeft)
	for {
		switch item := p.lexer.nextItem(); item.Typ {

		case itemBraceRight:
			return

		case itemIdentifier:
			permission := item.Val
			p.expect(itemOperatorColon)
			p.expect(itemParenLeft)
			p.expect(itemKeywordCtx)

			// parse optional type annotation of 'ctx' argument
			switch item := p.lexer.nextItem(); item.Typ {
			case itemOperatorColon:
				if t := p.expect(itemIdentifier); t != "Context" {
					p.addErr(fmt.Errorf("unexpected type annotation: %s", t))
				}
				p.expect(itemParenRight)
			case itemParenRight:
			default:
				p.addFatal(fmt.Errorf("expected ':' or ')', got %q", item.Val))
			}

			// parse optional return type annotation
			switch item := p.lexer.nextItem(); item.Typ {
			case itemOperatorColon:
				if t := p.expect(itemIdentifier); t != "boolean" {
					p.addErr(fmt.Errorf("unexpected type annotation: %s", t))
				}
				p.expect(itemOperatorArrow)
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
				p.expect(itemKeywordThis)
				p.expect(itemOperatorDot)
				p.expect(itemKeywordRelated)
				p.expect(itemOperatorDot)
				relationName := p.expect(itemIdentifier)
				p.expect(itemOperatorDot)

				switch p.expect(itemIdentifier) {
				case "some": // tuple to userset
					p.expect(itemParenLeft)
					arg := p.expect(itemIdentifier)
					p.expect(itemOperatorArrow)
					obj := p.expect(itemIdentifier)
					if arg != obj {
						p.addErr(fmt.Errorf("unexpected object: %s", obj))
					}
					p.expect(itemOperatorDot)
					p.expect(itemKeywordPermits)
					p.expect(itemOperatorDot)
					computedUsersetRel := p.expect(itemIdentifier)
					p.expect(itemParenLeft)
					p.expect(itemKeywordCtx)
					p.expect(itemParenRight)
					p.expect(itemParenRight)
					addRewrite(&relation,
						ast.TupleToUserset{
							Relation:                relationName,
							ComputedUsersetRelation: computedUsersetRel,
						},
					)
				case "includes": // computed userset
					p.expect(itemParenLeft)
					p.expect(itemKeywordCtx)
					p.expect(itemOperatorDot)
					if s := p.expect(itemIdentifier); s != "subject" {
						p.addErr(fmt.Errorf("expected subject, got: %s", s))
					}
					p.expect(itemParenRight)
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
