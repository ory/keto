// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package schema

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/ory/keto/ketoapi"
	opl "github.com/ory/keto/proto/ory/keto/opl/v1alpha1"
)

type (
	ParseError struct {
		msg  string
		item item
		p    *parser
	}
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (e *ParseError) Error() string {
	var s strings.Builder
	start := e.toSrcPos(e.item.Start)
	end := e.toSrcPos(e.item.End)
	rows := e.rows()
	startLineIdx := max(start.Line-2, 0)
	errorLineIdx := max(start.Line-1, 0)

	s.WriteString(fmt.Sprintf("error from %d:%d to %d:%d: %s\n\n",
		start.Line, start.Col,
		end.Line, end.Col,
		e.msg))

	if len(rows) < start.Line {
		s.WriteString("meta error: could not find source position in input\n")
		return s.String()
	}

	for line := startLineIdx; line <= errorLineIdx; line++ {
		s.WriteString(fmt.Sprintf("%4d | %s\n", line, rows[line]))
	}
	s.WriteString("     | ")
	for i, r := range rows[errorLineIdx] {
		switch {
		case start.Col == i:
			s.WriteRune('^')
		case start.Col <= i && i <= end.Col-1:
			s.WriteRune('~')
		case unicode.IsSpace(r):
			s.WriteRune(r)
		default:
			s.WriteRune(' ')
		}
	}
	s.WriteRune('\n')

	if errorLineIdx+1 < len(rows) {
		s.WriteString(fmt.Sprintf("%4d | %s\n", errorLineIdx+1, rows[errorLineIdx+1]))
		s.WriteRune('\n')
	}

	return s.String()
}

func (e *ParseError) ToAPI() *ketoapi.ParseError {
	return &ketoapi.ParseError{
		Message: e.msg,
		Start:   e.toSrcPos(e.item.Start),
		End:     e.toSrcPos(e.item.End),
	}
}
func (e *ParseError) ToProto() *opl.ParseError {
	start := e.toSrcPos(e.item.Start)
	end := e.toSrcPos(e.item.End)
	return &opl.ParseError{
		Message: e.msg,
		Start: &opl.SourcePosition{
			Line:   uint32(start.Line),
			Column: uint32(start.Col),
		},
		End: &opl.SourcePosition{
			Line:   uint32(end.Line),
			Column: uint32(end.Col),
		},
	}
}

// toSrcPos converts the given position in the input to a Line and column
// number.
func (e *ParseError) toSrcPos(pos int) (srcPos ketoapi.SourcePosition) {
	srcPos.Line = 1
	for _, c := range e.p.lexer.input {
		srcPos.Col++
		pos--
		if pos <= 0 {
			break
		}
		if c == '\n' {
			srcPos.Line++
			srcPos.Col = 0
		}
	}
	return
}

func (e *ParseError) rows() []string {
	return strings.Split(e.p.lexer.input, "\n")
}
