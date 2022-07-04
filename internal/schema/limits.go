package schema

const (
	// tupleToUsersetTypeCheckMaxDepth Controls the maximum number of recursions
	// for looking up the types of SubjectSet<Namespace, "relation">.
	tupleToUsersetTypeCheckMaxDepth = 10

	// expressionNestingMaxDepth is the maximum number of nested '(' in a single
	// 'permits'.
	expressionNestingMaxDepth = 10
)
