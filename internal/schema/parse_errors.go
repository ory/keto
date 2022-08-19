package schema

import (
	"fmt"
	"strings"
	"unicode"
)

type ParseError struct {
	msg  string
	item item
	p    *parser
}
type sourcePosition struct {
	line, col int
}

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
	startLineIdx := max(start.line-2, 0)
	errorLineIdx := max(start.line-1, 0)

	s.WriteString(fmt.Sprintf("error from %d:%d to %d:%d: %s\n\n",
		start.line, start.col,
		end.line, end.col,
		e.msg))

	if len(rows) < start.line {
		s.WriteString("meta error: could not find source position in input\n")
		return s.String()
	}

	for line := startLineIdx; line <= errorLineIdx; line++ {
		s.WriteString(fmt.Sprintf("%4d | %s\n", line, rows[line]))
	}
	s.WriteString("       ")
	for i, r := range rows[errorLineIdx] {
		switch {
		case start.col == i:
			s.WriteRune('^')
		case start.col <= i && i <= end.col-1:
			s.WriteRune('~')
		case unicode.IsSpace(r):
			s.WriteRune(r)
		default:
			s.WriteRune(' ')
		}
	}
	s.WriteRune('\n')

	if errorLineIdx+1 < len(rows) {
		s.WriteString(fmt.Sprintf("%4d | %s\n", errorLineIdx, rows[errorLineIdx+1]))
		s.WriteRune('\n')
	}

	return s.String()
}

// toSrcPos converts the given position in the input to a line and column
// number.
func (e *ParseError) toSrcPos(pos int) (srcPos sourcePosition) {
	srcPos.line = 1
	for _, c := range e.p.lexer.input {
		srcPos.col++
		pos--
		if pos == 0 {
			return
		}
		if c == '\n' {
			srcPos.line++
			srcPos.col = 0
		}
	}
	return sourcePosition{0, 0}
}

func (e *ParseError) rows() []string {
	return strings.Split(e.p.lexer.input, "\n")
}
