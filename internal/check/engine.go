package check

import (
	"context"
	"fmt"

	"github.com/ory/keto/internal/persistence"
	"github.com/ory/keto/internal/relationtuple"
	"github.com/ory/keto/internal/x"
)

type (
	EngineProvider interface {
		PermissionEngine() *Engine
	}
	Engine struct {
		d EngineDependencies
	}
	EngineDependencies interface {
		relationtuple.ManagerProvider
	}
)

func NewEngine(d EngineDependencies) *Engine {
	return &Engine{
		d: d,
	}
}

func (e *Engine) subjectIsAllowed(ctx context.Context, requested *relationtuple.InternalRelationTuple, rels []*relationtuple.InternalRelationTuple) (bool, error) {
	// This is the same as the graph problem "can requested.Subject be reached from requested.Object through the first outgoing edge requested.Relation"
	//
	// recursive breadth-first search
	// TODO replace by more performant algorithm

	var allowed bool
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
			continue
		}

		var err error
		// expand the set by one indirection; paginated
		allowed, err = e.checkOneIndirectionFurther(ctx, requested, &relationtuple.RelationQuery{Object: sub.Object, Relation: sub.Relation, Namespace: sub.Namespace})
		if err != nil {
			return false, err
		}
	}

	return allowed, nil
}

func (e *Engine) checkOneIndirectionFurther(ctx context.Context, requested *relationtuple.InternalRelationTuple, expandQuery *relationtuple.RelationQuery) (allowed bool, err error) {
	var (
		nextRels []*relationtuple.InternalRelationTuple
		nextPage string
	)

	// loop through pages until either allowed, end of pages, or an error occurred
	for !allowed && nextPage != persistence.PageTokenEnd && err == nil {
		nextRels, nextPage, err = e.d.RelationTupleManager().GetRelationTuples(ctx, expandQuery, x.WithToken(nextPage))
		if err != nil {
			return
		}

		allowed, err = e.subjectIsAllowed(ctx, requested, nextRels)
	}
	return
}

func (e *Engine) SubjectIsAllowed(ctx context.Context, r *relationtuple.InternalRelationTuple) (bool, error) {
	return e.checkOneIndirectionFurther(ctx, r, &relationtuple.RelationQuery{Object: r.Object, Relation: r.Relation, Namespace: r.Namespace})
}
