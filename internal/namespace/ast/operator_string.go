// Copyright © 2022 Ory Corp

// Code generated by "stringer -type=Operator -linecomment"; DO NOT EDIT.

package ast

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[OperatorOr-0]
	_ = x[OperatorAnd-1]
}

const _Operator_name = "orand"

var _Operator_index = [...]uint8{0, 2, 5}

func (i Operator) String() string {
	if i < 0 || i >= Operator(len(_Operator_index)-1) {
		return "Operator(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Operator_name[_Operator_index[i]:_Operator_index[i+1]]
}
