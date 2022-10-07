// Copyright © 2022 Ory Corp

package schema

const (
	// tupleToSubjectSetTypeCheckMaxDepth controls the maximum number of recursions
	// for looking up the types of SubjectSet<Namespace, "relation">.
	tupleToSubjectSetTypeCheckMaxDepth = 10

	// expressionNestingMaxDepth is the maximum number of nested '(' and '!' in
	// a single 'permits'.
	expressionNestingMaxDepth = 10
)
