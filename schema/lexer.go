// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

// The lexer is inspired by Rob Pike's talk at
// https://www.youtube.com/watch?v=HxaD_trXwRE.

package schema

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

type (
	item struct {
		Typ        itemType // Type of item.
		Val        string   // Value of item.
		Start, End int      // Start and end position of item.
	}
	itemType int

	// stateFn represents the state of the scanner as a function that returns
	// the next state.
	stateFn func(*lexer) stateFn

	// lexer holds the state of the scanner.
	lexer struct {
		name  string    // the name of the input; used only for error reports.
		input string    // the string being scanned.
		state stateFn   // the next lexing function to enter
		pos   int       // current position in the input.
		start int       // start position of this item.
		width int       // width of last rune read from input.
		items chan item // channel of scanned items.
	}
)

//go:generate stringer -type=itemType -trimprefix item -linecomment
const (
	// error occurred; value is text of error
	itemError itemType = iota
	// end of input
	itemEOF

	// identifier; value is string
	itemIdentifier
	// comment; value is string
	itemComment
	// string literal; value is string
	itemStringLiteral

	// keywords
	itemKeywordClass
	itemKeywordImplements
	itemKeywordThis
	itemKeywordCtx

	// operators
	itemOperatorAnd    // "&&"
	itemOperatorOr     // "||"
	itemOperatorNot    // "!"
	itemOperatorAssign // "="
	itemOperatorArrow  // "=>"
	itemOperatorDot    // "."
	itemOperatorColon  // ":"
	itemOperatorComma  // ","

	// misc characters
	itemSemicolon // ";"
	itemTypeUnion // "|"

	// brackets
	itemParenLeft    // "("
	itemParenRight   // ")"
	itemBraceLeft    // "{"
	itemBraceRight   // "}"
	itemBracketLeft  // "["
	itemBracketRight // "]"
	itemAngledLeft   // "<"
	itemAngledRight  // ">"
)

// string classes
const (
	spaces  = "\t\n\v\f\r "
	digits  = "0123456789"
	letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_"
)

const eof rune = -1 // rune indicating end of file

func (i item) String() string {
	switch i.Typ {
	case itemError:
		return "error: " + i.Val
	case itemEOF:
		return "EOF"
	case itemIdentifier, itemStringLiteral:
		if len(i.Val) > 10 {
			return fmt.Sprintf("'%.10s...'", i.Val)
		}
		return fmt.Sprintf("'%s'", i.Val)
	}
	return i.Val
}

// Lex creates a new scanner for the input string.
func Lex(name, input string) *lexer {
	l := &lexer{
		name:  name,
		input: input,
		state: lexCode,
		items: make(chan item, 20),
	}
	return l
}

// next returns the next rune in the input.
func (l *lexer) next() (r rune) {
	if l.pos >= len(l.input) {
		l.width = 0
		return eof
	}
	r, l.width = utf8.DecodeRuneInString(l.input[l.pos:])
	l.pos += l.width
	return r
}

// peek returns but does not consume the next rune in the input.
func (l *lexer) peek() rune {
	r := l.next()
	l.backup()
	return r
}

// backup steps back one rune. Can only be called once per call of next.
func (l *lexer) backup() {
	l.pos -= l.width
}

// emit passes an item back to the client.
func (l *lexer) emit(t itemType) {
	l.items <- item{t, l.input[l.start:l.pos], l.start, l.pos}
	l.start = l.pos
}

// ignore skips over the pending input before this point.
func (l *lexer) ignore() {
	l.start = l.pos
}

// accept consumes the next rune if it's from the valid set.
func (l *lexer) accept(valid string) bool {
	if strings.ContainsRune(valid, l.next()) {
		return true
	}
	l.backup()
	return false
}

// acceptRun consumes a run of runes from the valid set.
func (l *lexer) acceptRun(valid string) {
	for strings.ContainsRune(valid, l.next()) {
	}
	l.backup()
}

// error returns an error token and terminates the scan by passing
// back a nil pointer that will be the next state, terminating l.run.
func (l *lexer) errorf(format string, args ...interface{}) stateFn {
	location := fmt.Sprintf("at %q: ", l.input[l.pos:])
	l.items <- item{itemError, location + fmt.Sprintf(format, args...), l.start, l.pos}
	return nil
}

// nextItem returns the next item from the input.
func (l *lexer) nextItem() item {
	for {
		select {
		case item := <-l.items:
			return item
		default:
			if l.state == nil {
				return item{itemError, "broken state", 0, 0}
			}
			l.state = l.state(l)
		}
	}
}

// nextNonCommentItem returns the next item from the input that is not a comment.
func (l *lexer) nextNonCommentItem() (item item) {
	for item = l.nextItem(); item.Typ == itemComment; item = l.nextItem() {
	}
	return
}

// scanIdentifier scans an identifier.
func (l *lexer) scanIdentifier() bool {
	if !l.accept(letters) {
		return false
	}
	l.acceptRun(letters + digits)
	return true
}

func (l *lexer) scanCommentBegin() (bool, stateFn) {
	if strings.HasPrefix(l.input[l.pos:], "//") {
		l.pos += 2
		return true, lexLineComment
	}
	if strings.HasPrefix(l.input[l.pos:], "/*") {
		l.pos += 2
		return true, lexBlockComment
	}
	return false, nil
}

var oneRuneTokens = map[rune]itemType{
	':': itemOperatorColon,
	'.': itemOperatorDot,
	'(': itemParenLeft,
	')': itemParenRight,
	'[': itemBracketLeft,
	']': itemBracketRight,
	'{': itemBraceLeft,
	'}': itemBraceRight,
	'<': itemAngledLeft,
	'>': itemAngledRight,
	'=': itemOperatorAssign,
	',': itemOperatorComma,
	';': itemSemicolon,
	'|': itemTypeUnion,
	'!': itemOperatorNot,
}

var multiRuneTokens = map[string]itemType{
	"=>": itemOperatorArrow,
	"||": itemOperatorOr,
	"&&": itemOperatorAnd,
}

var keywords = map[string]itemType{
	"class":      itemKeywordClass,
	"implements": itemKeywordImplements,
	"this":       itemKeywordThis,
	"ctx":        itemKeywordCtx,
}

func lexCode(l *lexer) stateFn {
	l.acceptRun(spaces)
	l.ignore()

	r := l.peek()
	if r == eof {
		l.emit(itemEOF)
		return nil
	}

	// Two-rune tokens must be matched first
	for token, typ := range multiRuneTokens {
		if strings.HasPrefix(l.input[l.pos:], token) {
			l.pos += len(token)
			l.emit(typ)
			return lexCode
		}
	}

	if found, lexComment := l.scanCommentBegin(); found {
		return lexComment
	}

	if itemType, found := oneRuneTokens[r]; found {
		l.next()
		l.emit(itemType)
		return lexCode
	}

	if strings.ContainsRune(`'"`, r) {
		return lexStringLiteral
	}

	if l.scanIdentifier() {
		if kwType, found := keywords[l.input[l.start:l.pos]]; found {
			l.emit(kwType)
		} else {
			l.emit(itemIdentifier)
		}
		return lexCode
	}

	return l.errorf("unexpected token %c", r)
}

func lexLineComment(l *lexer) stateFn {
	for {
		switch l.next() {
		case '\n', eof:
			l.backup()
			l.emit(itemComment)
			return lexCode
		}
	}
}

func lexBlockComment(l *lexer) stateFn {
	for r := l.peek(); ; r = l.next() {
		if r == eof {
			return l.errorf("unclosed comment")
		}
		if strings.HasPrefix(l.input[l.pos:], "*/") {
			l.pos += 2
			l.emit(itemComment)
			return lexCode
		}
	}
}

// lexStringLiteral scans a string literal.
func lexStringLiteral(l *lexer) stateFn {
	r := l.next()
	l.ignore()

loop:
	for {
		switch l.next() {
		case eof:
			return l.errorf("unclosed string literal")
		case r:
			l.backup()
			break loop
		}
	}
	l.emit(itemStringLiteral)
	l.next()
	l.ignore()

	return lexCode
}
