package check

import (
	"context"
	"fmt"
	"os"

	"github.com/ory/keto/relationtuple"
)

type (
	EngineProvider interface {
		PermissionEngine() *Engine
	}
	Engine struct {
		d engineDependencies
	}
	engineDependencies interface {
		relationtuple.ManagerProvider
	}
)

func NewEngine(d engineDependencies) *Engine {
	return &Engine{
		d: d,
	}
}

func equalRelation(a, b *relationtuple.InternalRelationTuple) bool {
	return a.Relation == b.Relation && a.Subject.Equals(b.Subject) && a.Object.Equals(b.Object)
}

func (e *Engine) subjectIsAllowed(ctx context.Context, requested *relationtuple.InternalRelationTuple, subjectRelations []*relationtuple.InternalRelationTuple) (bool, error) {
	// This is the same as the graph problem "can requested.ObjectID be reached from requested.SubjectID through the incoming edge requested.Name"
	//
	// recursive breadth-first search
	// TODO replace by more performant algorithm

	var res bool
	for _, sr := range subjectRelations {

		// we don't have to check SubjectID here as we know that sr was reached from requested.SubjectID through 0...n indirections
		if requested.Relation == sr.Relation && requested.Object.Equals(sr.Object) {
			// found the requested relation
			return true, nil
		}

		prevRelationsLen := len(subjectRelations)

		// compute one indirection
		indirect, err := e.d.RelationTupleManager().GetRelationTuples(ctx, &relationtuple.RelationQuery{Subject: sr.DeriveSubject()})
		if err != nil {
			// TODO fix error handling
			_, _ = fmt.Fprintf(os.Stderr, "%+v", err)
		}

		for _, maybeRel := range indirect {
			var found bool
			for _, knownRel := range subjectRelations {
				if equalRelation(knownRel, maybeRel) {
					found = true
					break
				}
			}
			if !found {
				subjectRelations = append(subjectRelations, maybeRel)
			}
		}

		if prevRelationsLen < len(subjectRelations) {
			is, err := e.subjectIsAllowed(ctx, requested, subjectRelations)
			if err != nil {
				// TODO fix error handling
				_, _ = fmt.Fprintf(os.Stderr, "%+v", err)
			}
			res = res || is
		}
	}

	return res, nil
}

func (e *Engine) SubjectIsAllowed(ctx context.Context, r *relationtuple.InternalRelationTuple) (bool, error) {
	subjectRelations, err := e.d.RelationTupleManager().GetRelationTuples(ctx, &relationtuple.RelationQuery{Subject: r.Subject})
	if err != nil {
		return false, err
	}

	return e.subjectIsAllowed(ctx, r, subjectRelations)
}
