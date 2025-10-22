package rts

// RelationTupleToDeltas is a helper that converts a slice of RelationTuple to a slice of RelationTupleDelta
// with the specified RelationTupleDelta_Action. This allows you to conveniently assemble a request for the
// WriteServiceClient.TransactRelationTuples operation.
//
// Example:
//
//	 c.TransactRelationTuples(context.Background(), &rts.TransactRelationTuplesRequest{
//			RelationTupleDeltas: append(rts.RelationTupleToDeltas(insertTuples, rts.RelationTupleDelta_INSERT), rts.RelationTupleToDeltas(deleteTuples, rts.RelationTupleDelta_DELETE)...),
//	 })
func RelationTupleToDeltas(rs []*RelationTuple, action RelationTupleDelta_Action) []*RelationTupleDelta {
	deltas := make([]*RelationTupleDelta, len(rs))
	for i := range rs {
		deltas[i] = &RelationTupleDelta{
			RelationTuple: rs[i],
			Action:        action,
		}
	}
	return deltas
}

// NewSubjectSet returns a Subject with a SubjectSet ref.
func NewSubjectSet(namespace, object, relation string) *Subject {
	return &Subject{Ref: &Subject_Set{Set: &SubjectSet{
		Namespace: namespace,
		Object:    object,
		Relation:  relation,
	}}}
}

// NewSubjectID returns a Subject with a subject ID ref.
func NewSubjectID(id string) *Subject {
	return &Subject{Ref: &Subject_Id{Id: id}}
}

// NewSubjectQueryFromSet returns a subject query with a SubjectSet ref.
func NewSubjectQueryFromSet(namespace, object, relation string) *SubjectQuery {
	return &SubjectQuery{Ref: &SubjectQuery_Set{Set: &SubjectSetQuery{
		Namespace: namespace,
		Object:    object,
		Relation:  relation,
	}}}
}

// NewSubjectQueryFromID returns a subject query with a subject ID ref.
func NewSubjectQueryFromID(id string) *SubjectQuery {
	return &SubjectQuery{Ref: &SubjectQuery_Id{Id: id}}
}
