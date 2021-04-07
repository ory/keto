package check

import (
	"context"
	"errors"

	"github.com/ory/herodot"

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
	// We implement recursive depth-first search here.
	// TODO replace by more performant algorithm: https://github.com/ory/keto/issues/483

	for _, sr := range rels {
		// we only have to check Subject here as we know that sr was reached from requested.ObjectID, requested.Relation through 0...n indirections
		if requested.Subject.Equals(sr.Subject) {
			// found the requested relation
			return true, nil
		}

		sub, isSubjectSet := sr.Subject.(*relationtuple.SubjectSet)
		if !isSubjectSet {
			continue
		}

		// expand the set by one indirection; paginated
		allowed, err := e.checkOneIndirectionFurther(ctx, requested, &relationtuple.RelationQuery{Object: sub.Object, Relation: sub.Relation, Namespace: sub.Namespace})
		if err != nil {
			return false, err
		}
		if allowed {
			return true, nil
		}
	}

	return false, nil
}

func (e *Engine) checkOneIndirectionFurther(ctx context.Context, requested *relationtuple.InternalRelationTuple, expandQuery *relationtuple.RelationQuery) (bool, error) {
	// an empty page token denotes the first page (as tokens are opaque)
	var prevPage string

	for {
		nextRels, nextPage, err := e.d.RelationTupleManager().GetRelationTuples(ctx, expandQuery, x.WithToken(prevPage))
		// herodot.ErrNotFound occurs when the namespace is unknown
		if errors.Is(err, herodot.ErrNotFound) {
			return false, nil
		} else if err != nil {
			return false, err
		}

		allowed, err := e.subjectIsAllowed(ctx, requested, nextRels)

		// loop through pages until either allowed, end of pages, or an error occurred
		if allowed || nextPage == "" || err != nil {
			return allowed, err
		}

		prevPage = nextPage
	}
}

func (e *Engine) SubjectIsAllowed(ctx context.Context, r *relationtuple.InternalRelationTuple) (bool, error) {
	return e.checkOneIndirectionFurther(ctx, r, &relationtuple.RelationQuery{Object: r.Object, Relation: r.Relation, Namespace: r.Namespace})
}
