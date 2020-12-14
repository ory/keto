package check

import (
	"context"
	"fmt"
	"os"

	"github.com/ory/keto/internal/relationtuple"
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

func (e *Engine) subjectIsAllowed(ctx context.Context, requested *relationtuple.InternalRelationTuple, rels []*relationtuple.InternalRelationTuple) (bool, error) {
	// This is the same as the graph problem "can requested.SubjectID be reached from requested.Object through the first outgoing edge requested.Name"
	//
	// recursive breadth-first search
	// TODO replace by more performant algorithm

	var res bool
	for _, sr := range rels {
		// TODO move this to input validation
		if requested.Subject == nil {
			return false, fmt.Errorf("subject is unexpectedly nil for %+v", requested)
		}
		// we only have to check Subject here as we know that sr was reached from requested.ObjectID, requested.Relation through 0...n indirections
		if requested.Subject.Equals(sr.Subject) {
			// found the requested relation
			return true, nil
		}

		sub, isSubjectSet := sr.Subject.(*relationtuple.SubjectSet)
		if !isSubjectSet {
			return false, nil
		}

		// expand the set by one indirection
		// TODO handle pagination
		nextRels, _, err := e.d.RelationTupleManager().GetRelationTuples(ctx, &relationtuple.RelationQuery{Object: sub.Object, Relation: sub.Relation, Namespace: sub.Namespace})
		if err != nil {
			// TODO fix error handling
			_, _ = fmt.Fprintf(os.Stderr, "%+v", err)
			return false, err
		}

		is, err := e.subjectIsAllowed(ctx, requested, nextRels)
		if err != nil {
			// TODO fix error handling
			_, _ = fmt.Fprintf(os.Stderr, "%+v", err)
			return false, err
		}
		res = res || is
	}

	return res, nil
}

func (e *Engine) SubjectIsAllowed(ctx context.Context, r *relationtuple.InternalRelationTuple) (bool, error) {
	// TODO handle pagination
	subjectRelations, _, err := e.d.RelationTupleManager().GetRelationTuples(ctx, &relationtuple.RelationQuery{Object: r.Object, Relation: r.Relation, Namespace: r.Namespace})
	if err != nil {
		return false, err
	}

	return e.subjectIsAllowed(ctx, r, subjectRelations)
}
