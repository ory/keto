package check

import (
	"context"
	"errors"

	"github.com/ory/keto/internal/namespace"

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
		namespace.ManagerProvider
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
		// we only have to check Subject here as we know that sr was reached from requested.ObjectID, requested.Relation through 0...n indirections
		if requested.Subject.Equals(sr.Subject) {
			// found the requested relation
			return true, nil
		}

		sub, isSubjectSet := sr.Subject.(*relationtuple.SubjectSet)
		if !isSubjectSet {
			continue
		}

		nm, err := e.d.NamespaceManager()
		if err != nil {
			return false, err
		}
		n, err := nm.GetNamespace(ctx, sr.Namespace)
		if err != nil {
			return false, err
		}
		// use the namespace here to decide which bool operators to use in the loop, and how to adjust the query when moving on (e.g. different relation name)
		var strategy func(ctx context.Context, requested, current *relationtuple.InternalRelationTuple, checkFurther func(context.Context, *relationtuple.InternalRelationTuple, *relationtuple.RelationQuery) (bool, error)) (continueLoop, allowed bool, err error) = n.GetRelationRewriteStrategy(ctx, sr.Relation)

		// expand the set by one indirection; paginated
		continueLoop, allowed, err := strategy(ctx, requested, sr, e.checkOneIndirectionFurther)
		if err != nil {
			return false, err
		}
		if !continueLoop {
			return allowed, nil
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
	for !allowed && nextPage != x.PageTokenEnd && err == nil {
		nextRels, nextPage, err = e.d.RelationTupleManager().GetRelationTuples(ctx, expandQuery, x.WithToken(nextPage))
		if errors.Is(err, herodot.ErrNotFound) {
			allowed, err = false, nil
			return
		} else if err != nil {
			return
		}

		allowed, err = e.subjectIsAllowed(ctx, requested, nextRels)
	}
	return
}

func (e *Engine) SubjectIsAllowed(ctx context.Context, r *relationtuple.InternalRelationTuple) (bool, error) {
	return e.checkOneIndirectionFurther(ctx, r, &relationtuple.RelationQuery{Object: r.Object, Relation: r.Relation, Namespace: r.Namespace})
}
