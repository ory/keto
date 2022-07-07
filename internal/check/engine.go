package check

import (
	"context"

	"github.com/ory/herodot"
	"github.com/pkg/errors"

	"github.com/ory/keto/internal/check/checkgroup"
	"github.com/ory/keto/internal/driver/config"
	"github.com/ory/keto/internal/expand"
	"github.com/ory/keto/internal/namespace"
	"github.com/ory/keto/internal/namespace/ast"
	"github.com/ory/keto/internal/relationtuple"
	"github.com/ory/keto/internal/x"
	"github.com/ory/keto/internal/x/graph"
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
		config.Provider
		x.LoggerProvider
	}

	// Type aliases for shorter signatures
	RelationTuple = relationtuple.RelationTuple
	Query         = relationtuple.RelationQuery
)

func NewEngine(d EngineDependencies) *Engine {
	return &Engine{
		d: d,
	}
}

func (e *Engine) isIncluded(
	ctx context.Context,
	requested *RelationTuple,
	rels []*RelationTuple,
	restDepth int,
) checkgroup.CheckFunc {
	if restDepth < 0 {
		e.d.Logger().Debug("reached max-depth, therefore this query will not be further expanded")
		return checkgroup.UnknownMemberFunc
	}

	// This is the same as the graph problem "can requested.Subject be reached
	// from requested.Object through the first outgoing edge requested.Relation"
	//
	// We implement recursive depth-first search here.
	// TODO replace by more performant algorithm:
	// https://github.com/ory/keto/issues/483

	ctx = graph.InitVisited(ctx)
	g := checkgroup.New(ctx)

	for _, sr := range rels {
		var wasAlreadyVisited bool
		ctx, wasAlreadyVisited = graph.CheckAndAddVisited(ctx, sr.Subject)
		if wasAlreadyVisited {
			continue
		}

		// we only have to check Subject here as we know that sr was reached
		// from requested.ObjectID, requested.Relation through 0...n
		// indirections
		if requested.Subject.Equals(sr.Subject) {
			// found the requested relation
			g.SetIsMember()
			break
		}

		sub, isSubjectSet := sr.Subject.(*relationtuple.SubjectSet)
		if !isSubjectSet {
			continue
		}

		g.Add(e.subQuery(
			requested,
			&Query{Object: sub.Object, Relation: sub.Relation, Namespace: sub.Namespace},
			restDepth,
		))
	}
	return checkgroup.WithEdge(checkgroup.Edge{
		Tuple: *requested,
		Type:  expand.Union,
	}, g.CheckFunc())
}

func (e *Engine) subQuery(
	requested *RelationTuple,
	query *Query,
	restDepth int,
) checkgroup.CheckFunc {
	if restDepth < 0 {
		e.d.Logger().
			WithFields(requested.ToLoggerFields()).
			Debug("reached max-depth, therefore this query will not be further expanded")
		return checkgroup.UnknownMemberFunc
	}

	return func(ctx context.Context, resultCh chan<- checkgroup.Result) {
		e.d.Logger().
			WithField("request", requested.String()).
			WithField("query", query.String()).
			Trace("check one indirection further")

		// an empty page token denotes the first page (as tokens are opaque)
		var prevPage string

		// Special case: check if we can find the subject id directly
		if rels, _, err := e.d.RelationTupleManager().GetRelationTuples(ctx, requested.ToQuery()); err == nil && len(rels) > 0 {
			resultCh <- checkgroup.Result{
				Membership: checkgroup.IsMember,
				Tree: &expand.Tree{
					Type:  expand.Leaf,
					Tuple: requested,
				},
			}
			return
		}

		g := checkgroup.New(ctx)

		for {
			nextRels, nextPage, err := e.d.RelationTupleManager().GetRelationTuples(ctx, query, x.WithToken(prevPage))
			// herodot.ErrNotFound occurs when the namespace is unknown
			if errors.Is(err, herodot.ErrNotFound) {
				g.Add(checkgroup.NotMemberFunc)
				break
			} else if err != nil {
				g.Add(checkgroup.ErrorFunc(err))
				break
			}

			g.Add(e.isIncluded(ctx, requested, nextRels, restDepth-1))

			// loop through pages until either allowed, end of pages, or an error occurred
			if nextPage == "" || g.Done() {
				break
			}
			prevPage = nextPage
		}

		resultCh <- g.Result()
	}
}

func (e *Engine) CheckIsMember(ctx context.Context, r *RelationTuple, restDepth int) (bool, error) {
	result := e.CheckRelationTuple(ctx, r, restDepth)
	if result.Err != nil {
		return false, result.Err
	}
	return result.Membership == checkgroup.IsMember, nil
}

func (e *Engine) CheckRelationTuple(ctx context.Context, r *RelationTuple, restDepth int) checkgroup.Result {
	// global max-depth takes precedence when it is the lesser or if the request
	// max-depth is less than or equal to 0
	if globalMaxDepth := e.d.Config(ctx).MaxReadDepth(); restDepth <= 0 || globalMaxDepth < restDepth {
		restDepth = globalMaxDepth
	}

	resultCh := make(chan checkgroup.Result)
	go e.checkIsAllowed(ctx, r, restDepth)(ctx, resultCh)
	select {
	case result := <-resultCh:
		return result
	case <-ctx.Done():
		return checkgroup.Result{Err: errors.WithStack(ctx.Err())}
	}
}

func (e *Engine) checkIsAllowed(ctx context.Context, r *RelationTuple, restDepth int) checkgroup.CheckFunc {
	if restDepth < 0 {
		e.d.Logger().Debug("reached max-depth, therefore this query will not be further expanded")
		return checkgroup.UnknownMemberFunc
	}

	e.d.Logger().
		WithField("request", r.String()).
		Trace("check is allowed")

	g := checkgroup.New(ctx)
	g.Add(e.subQuery(r,
		&Query{
			Object:    r.Object,
			Relation:  r.Relation,
			Namespace: r.Namespace,
		}, restDepth),
	)

	relation, err := e.astRelationFor(ctx, r)
	if err != nil {
		g.Add(checkgroup.ErrorFunc(err))
	} else if relation != nil && relation.UsersetRewrite != nil {
		g.Add(e.checkUsersetRewrite(ctx, r, relation.UsersetRewrite, restDepth))
	}

	return g.CheckFunc()
}

func (e *Engine) astRelationFor(ctx context.Context, r *RelationTuple) (*ast.Relation, error) {
	ns, err := e.namespaceFor(ctx, r)
	if err != nil {
		// On an unknown namespace the answer should be "not allowed", not "not
		// found". Therefore we don't return the error here.
		return nil, nil
	}

	// Special case: If Relations is empty, then there is no namespace
	// configuration, and it is not an error that the relation was not found.
	if len(ns.Relations) == 0 {
		return nil, nil
	}

	for _, rel := range ns.Relations {
		if rel.Name == r.Relation {
			return &rel, nil
		}
	}
	return nil, errors.New("relation not found")
}

func (e *Engine) namespaceFor(ctx context.Context, r *RelationTuple) (*namespace.Namespace, error) {
	namespaceManager, err := e.d.Config(ctx).NamespaceManager()
	if err != nil {
		return nil, err
	}
	ns, err := namespaceManager.GetNamespaceByName(ctx, r.Namespace)
	if err != nil {
		return nil, err
	}
	return ns, nil
}
