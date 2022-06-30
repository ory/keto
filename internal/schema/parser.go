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

type matcher func(p *parser) (matched bool)

// optional optionally matches the first argument of tokens in the input. If
// matched, the tokens are consumed. If the first token matched, all other
// tokens must match as well.
func optional(tokens ...string) matcher {
	return func(p *parser) bool {
		if len(tokens) == 0 {
			return true
		}
		first := tokens[0]
		if p.peek().Val == first {
			p.next()
			for _, token := range tokens[1:] {
				i := p.next()
				if i.Val != token {
					p.addFatal(i, "expected %q, got %q", token, i.Val)
					return false
				}
			}
		}
		return true
	}
}

// match matches for the next tokens in the input.
//
// A token is matched depending on the type:
// For string arguments, the input token must match the given string exactly.
// For *string arguments, the input token must be an identifier, and the value
// of the identifier will be written to the *string.
// For *item arguments, the input token will be written to the pointer.
func (p *parser) match(tokens ...interface{}) (matched bool) {
	if p.fatal {
		return false
	}

	for _, token := range tokens {
		switch token := token.(type) {
		case string:
			i := p.next()
			if i.Val != token {
				p.addFatal(i, "expected %q, got %q", token, i.Val)
				return false
			}
		case *string:
			i := p.next()
			if i.Typ != itemIdentifier && i.Typ != itemStringLiteral {
				p.addFatal(i, "expected identifier, got %s", i.Typ)
				return false
			}
			*token = i.Val
		case *item:
			*token = p.next()
		case matcher:
			if !token(p) {
				return false
			}
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

// matchIf matches the tokens iff. the predicate is true.
func (p *parser) matchIf(predicate itemPredicate, tokens ...interface{}) (matched bool) {
	if p.fatal {
		return false
	}
	if !predicate(p.peek()) {
		return false
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
				if item.Val == "SubjectSet" {
					p.matchSubjectSet()
				}
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

func (p *parser) matchSubjectSet() {
	p.match("<", any, ",", any, ">")
}

func (p *parser) parseTypeUnion() {
	for !p.fatal {
		var identifier string
		p.match(&identifier)
		if identifier == "SubjectSet" {
			p.matchSubjectSet()
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
			p.match(
				":", "(", "ctx", optional(":", "Context"), ")",
				optional(":", "boolean"), "=>",
			)

			relation := ast.Relation{Name: permission}
			defer func() {
				p.namespace.Relations = append(p.namespace.Relations, relation)
			}()

		permissionLoop:
			for !p.fatal {
				p.parsePermissionExpression(&relation)
				if p.fatal {
					return
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

func (p *parser) parsePermissionExpression(relation *ast.Relation) {
	var (
		name string
		r    ast.Child
		ok   bool
	)
	if !p.match("this", ".", "related", ".", &name, ".") {
		return
	}

	switch item := p.next(); item.Val {
	case "traverse":
		ok, r = p.parseTupleToUserset(name)
	case "includes":
		ok, r = p.parseComputedUserset(name)
	default:
		p.addFatal(item, "expected 'traverse' or 'includes', got %q", item.Val)
		return
	}
	if !ok {
		return
	}
	addRewrite(relation, r)
}

func (p *parser) parseTupleToUserset(relation string) (ok bool, rewrite ast.Child) {
	var (
		usersetRel string
		arg, verb  item
	)
	if !p.match("(") {
		return false, nil
	}

	switch {
	case p.matchIf(is(itemParenLeft), "(", &arg, ")"):
	case p.match(&arg):
	default:
		return false, nil
	}
	p.match("=>", arg.Val, ".", &verb)

	switch verb.Val {
	case "related":
		p.match(
			".", &usersetRel, ".", "includes", "(", "ctx", ".", "subject",
			optional(","), ")", optional(","), ")",
		)
	case "permits":
		p.match(".", &usersetRel, "(", "ctx", ")", ")")
	default:
		p.addFatal(verb, "expected 'related' or 'permits', got %q", verb)
		return false, nil
	}
	return true, ast.TupleToUserset{
		Relation:                relation,
		ComputedUsersetRelation: usersetRel,
	}
}

func (p *parser) parseComputedUserset(relation string) (ok bool, rewrite ast.Child) {
	if !p.match("(", "ctx", ".", "subject", ")") {
		return false, nil
	}
	return true, ast.ComputedUserset{Relation: relation}
}

func addRewrite(r *ast.Relation, child ast.Child) {
	if r.UsersetRewrite == nil {
		r.UsersetRewrite = &ast.UsersetRewrite{}
	}
	r.UsersetRewrite.Children = append(r.UsersetRewrite.Children, child)
}
