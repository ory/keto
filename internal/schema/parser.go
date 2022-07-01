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

// maxExprNestingDepth is the maximum number of nested '(' in a single 'permits'.
const maxExprNestingDepth = 10

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
		// Get next non-comment token.
		for item = p.lexer.nextItem(); item.Typ == itemComment; item = p.lexer.nextItem() {
		}
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

			rewrite := simplifyExpression(p.parsePermissionExpressions(itemOperatorComma, maxExprNestingDepth))
			if rewrite == nil {
				return
			}
			p.namespace.Relations = append(p.namespace.Relations,
				ast.Relation{
					Name:           permission,
					UsersetRewrite: rewrite,
				})

		default:
			p.addFatal(item, "expected identifier or '}', got %q", item.Val)
			return
		}
	}
}

func (p *parser) parsePermissionExpressions(finalToken itemType, depth int) *ast.UsersetRewrite {
	var root *ast.UsersetRewrite
	lastParsedAnExpression := false
	for !p.fatal {
		switch item := p.peek(); {

		case item.Typ == itemParenLeft:
			p.next() // consume paren
			if depth <= 0 {
				p.addFatal(item, "expression nested too deeply; maximal nesting depth is %d", maxExprNestingDepth)
				return nil
			}
			child := p.parsePermissionExpressions(itemParenRight, depth-1)
			if child == nil {
				return nil
			}
			root = addChild(root, child)
			lastParsedAnExpression = false

		case item.Typ == finalToken:
			p.next() // consume final token
			return root

		case item.Typ == itemBraceRight:
			// We don't consume the '}' here, to allow `parsePermits` to consume
			// it.
			return root

		case item.Typ == itemOperatorAnd, item.Typ == itemOperatorOr:
			p.next() // consume operator
			newRoot := &ast.UsersetRewrite{
				Operation: setOperation(item.Typ),
				Children:  []ast.Child{root},
			}
			root = newRoot
			lastParsedAnExpression = false

		case lastParsedAnExpression:
			p.addFatal(item, "did not expect another expression")

		default:
			child := p.parsePermissionExpression()
			if child == nil {
				return nil
			}
			root = addChild(root, child)
			lastParsedAnExpression = true
		}
	}
	return nil
}

func addChild(root *ast.UsersetRewrite, child ast.Child) *ast.UsersetRewrite {
	if root == nil {
		return child.AsRewrite()
	} else {
		root.Children = append(root.Children, child)
		return root
	}
}

func setOperation(typ itemType) ast.SetOperation {
	switch typ {
	case itemOperatorAnd:
		return ast.SetOperationIntersection
	case itemOperatorOr:
		return ast.SetOperationUnion
	}
	panic("not reached")
}

func (p *parser) parsePermissionExpression() (child ast.Child) {
	var name string

	if !p.match("this", ".", "related", ".", &name, ".") {
		return
	}

	switch item := p.next(); item.Val {
	case "traverse":
		child = p.parseTupleToUserset(name)
	case "includes":
		child = p.parseComputedUserset(name)
	default:
		p.addFatal(item, "expected 'traverse' or 'includes', got %q", item.Val)
	}
	return
}

func (p *parser) parseTupleToUserset(relation string) (rewrite ast.Child) {
	var (
		usersetRel string
		arg, verb  item
	)
	if !p.match("(") {
		return nil
	}

	switch {
	case p.matchIf(is(itemParenLeft), "(", &arg, ")"):
	case p.match(&arg):
	default:
		return nil
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
		return nil
	}
	return &ast.TupleToUserset{
		Relation:                relation,
		ComputedUsersetRelation: usersetRel,
	}
}

func (p *parser) parseComputedUserset(relation string) (rewrite ast.Child) {
	if !p.match("(", "ctx", ".", "subject", ")") {
		return nil
	}
	return &ast.ComputedUserset{Relation: relation}
}

// simplifyExpression rewrites the expression to use n-ary set operations
// instead of binary ones.
func simplifyExpression(root *ast.UsersetRewrite) *ast.UsersetRewrite {
	if root == nil {
		return nil
	}
	var newChildren []ast.Child
	for _, child := range root.Children {
		if ch, ok := child.(*ast.UsersetRewrite); ok && ch.Operation == root.Operation {
			// merge child and root
			simplifyExpression(ch)
			newChildren = append(newChildren, ch.Children...)
		} else {
			// can't merge, just copy
			newChildren = append(newChildren, child)
		}
	}
	root.Children = newChildren

	return root
}
