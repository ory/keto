// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package check

import (
	"context"

	"github.com/pkg/errors"
	"go.opentelemetry.io/otel/trace"

	"github.com/ory/herodot"
	"github.com/ory/x/otelx"

	"github.com/ory/keto/x/events"

	"github.com/ory/keto/internal/check/checkgroup"
	"github.com/ory/keto/internal/driver/config"
	"github.com/ory/keto/internal/namespace"
	"github.com/ory/keto/internal/namespace/ast"
	"github.com/ory/keto/internal/persistence"
	"github.com/ory/keto/internal/relationtuple"
	"github.com/ory/keto/internal/x"
	"github.com/ory/keto/internal/x/graph"
	"github.com/ory/keto/ketoapi"
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
		persistence.Provider
		config.Provider
		x.LoggerProvider
		x.TracingProvider
		x.NetworkIDProvider
	}

	EngineOpt func(*Engine)

	// Type aliases for shorter signatures
	relationTuple = relationtuple.RelationTuple
	query         = relationtuple.RelationQuery
)

const WildcardRelation = "..."

func NewEngine(d EngineDependencies, opts ...EngineOpt) *Engine {
	e := &Engine{d: d}
	for _, opt := range opts {
		opt(e)
	}

	return e
}

// CheckIsMember checks if the relation tuple's subject has the relation on the
// object in the namespace either directly or indirectly and returns a boolean
// result.
func (e *Engine) CheckIsMember(ctx context.Context, r *relationTuple, restDepth int) (bool, error) {
	result := e.CheckRelationTuple(ctx, r, restDepth)
	if result.Err != nil {
		return false, result.Err
	}
	return result.Membership == checkgroup.IsMember, nil
}

// CheckRelationTuple checks if the relation tuple's subject has the relation on
// the object in the namespace either directly or indirectly and returns a check
// result.
func (e *Engine) CheckRelationTuple(ctx context.Context, r *relationTuple, restDepth int) (res checkgroup.Result) {
	ctx, span := e.d.Tracer(ctx).Tracer().Start(ctx, "Engine.CheckRelationTuple")
	defer otelx.End(span, &res.Err)

	// global max-depth takes precedence when it is the lesser or if the request
	// max-depth is less than or equal to 0
	if globalMaxDepth := e.d.Config(ctx).MaxReadDepth(); restDepth <= 0 || globalMaxDepth < restDepth {
		restDepth = globalMaxDepth
	}

	resultCh := make(chan checkgroup.Result)
	go e.checkIsAllowed(ctx, r, restDepth, false)(ctx, resultCh)
	select {
	case result := <-resultCh:
		trace.SpanFromContext(ctx).AddEvent(events.NewPermissionsChecked(ctx))
		return result
	case <-ctx.Done():
		return checkgroup.Result{Err: errors.WithStack(ctx.Err())}
	}
}

// checkExpandSubject checks the expansions of the subject set of the tuple.
//
// For a relation tuple n:obj#rel@user, checkExpandSubject first queries for all
// tuples that match n:obj#rel@* (arbitrary subjects), and then for each
// subject set checks subject_set@user.
func (e *Engine) checkExpandSubject(r *relationTuple, restDepth int) checkgroup.CheckFunc {
	if restDepth <= 0 {
		e.d.Logger().
			WithField("request", r.String()).
			Debug("reached max-depth, therefore this query will not be further expanded")
		return checkgroup.UnknownMemberFunc
	}
	return func(ctx context.Context, resultCh chan<- checkgroup.Result) {
		e.d.Logger().
			WithField("request", r.String()).
			Trace("check expand subject")

		g := checkgroup.New(ctx)
		defer func() { resultCh <- g.Result() }()

		var (
			visited  bool
			innerCtx = graph.InitVisited(ctx)
		)

		results, err := e.d.Traverser().TraverseSubjectSetExpansion(ctx, r)

		if errors.Is(err, herodot.ErrNotFound) {
			g.Add(checkgroup.NotMemberFunc)
			return
		} else if err != nil {
			g.Add(checkgroup.ErrorFunc(err))
			return
		}

		// See if the current hop was already enough to answer the check
		for _, result := range results {
			if result.Found {
				g.Add(checkgroup.IsMemberFunc)
				return
			}
		}

		// If not, we must go another hop:
		maxWidth := e.d.Config(ctx).MaxReadWidth()
		if len(results) > maxWidth {
			e.d.Logger().
				WithField("method", "checkExpandSubject").
				WithField("request", r.String()).
				WithField("max_width", maxWidth).
				WithField("results", len(results)).
				Debug("too many results, truncating")
			results = results[:maxWidth-1]
		}
		for _, result := range results {
			sub := &relationtuple.SubjectSet{
				Namespace: result.To.Namespace,
				Object:    result.To.Object,
				Relation:  result.To.Relation,
			}
			innerCtx, visited = graph.CheckAndAddVisited(innerCtx, sub)
			if visited {
				continue
			}
			g.Add(e.checkIsAllowed(innerCtx, result.To, restDepth, true))
		}
	}
}

// checkDirect checks if the relation tuple is in the database directly.
func (e *Engine) checkDirect(r *relationTuple, restDepth int) checkgroup.CheckFunc {
	if restDepth <= 0 {
		e.d.Logger().
			WithField("method", "checkDirect").
			Debug("reached max-depth, therefore this query will not be further expanded")
		return checkgroup.UnknownMemberFunc
	}
	return func(ctx context.Context, resultCh chan<- checkgroup.Result) {
		e.d.Logger().
			WithField("request", r.String()).
			Trace("check direct")
		found, err := e.d.RelationTupleManager().ExistsRelationTuples(
			ctx,
			r.ToQuery(),
		)

		switch {
		case err != nil:
			e.d.Logger().
				WithField("method", "checkDirect").
				WithError(err).
				Error("failed to look up direct access in db")
			resultCh <- checkgroup.Result{
				Membership: checkgroup.NotMember,
			}

		case found:
			resultCh <- checkgroup.Result{
				Membership: checkgroup.IsMember,
				Tree: &ketoapi.Tree[*relationtuple.RelationTuple]{
					Type:  ketoapi.TreeNodeLeaf,
					Tuple: r,
				},
			}

		default:
			resultCh <- checkgroup.Result{
				Membership: checkgroup.NotMember,
			}
		}
	}
}

// checkIsAllowed checks if the relation tuple is allowed (there is a path from
// the relation tuple subject to the namespace, object and relation) either
// directly (in the database), or through subject-set expansions, or through
// user-set rewrites.
func (e *Engine) checkIsAllowed(ctx context.Context, r *relationTuple, restDepth int, skipDirect bool) checkgroup.CheckFunc {
	if restDepth <= 0 {
		e.d.Logger().
			WithField("method", "checkIsAllowed").
			Debug("reached max-depth, therefore this query will not be further expanded")
		return checkgroup.UnknownMemberFunc
	}

	e.d.Logger().
		WithField("request", r.String()).
		Trace("check is allowed")

	g := checkgroup.New(ctx)

	relation, err := e.astRelationFor(ctx, r)
	if err != nil {
		g.Add(checkgroup.ErrorFunc(err))
		return g.CheckFunc()
	}
	hasRewrite := relation != nil && relation.SubjectSetRewrite != nil
	strictMode := e.d.Config(ctx).StrictMode()
	canHaveSubjectSets := !strictMode || relation == nil || containsSubjectSetExpand(relation)
	if hasRewrite {
		g.Add(e.checkSubjectSetRewrite(ctx, r, relation.SubjectSetRewrite, restDepth))
	}
	if (!strictMode || !hasRewrite) && !skipDirect {
		// In strict mode, add a direct check only if there is no subject set rewrite for this relation.
		// Rewrites are added as 'permits'.
		g.Add(e.checkDirect(r, restDepth-1))
	}
	if canHaveSubjectSets {
		g.Add(e.checkExpandSubject(r, restDepth-1))
	}

	return g.CheckFunc()
}

func containsSubjectSetExpand(relation *ast.Relation) bool {
	for _, t := range relation.Types {
		if t.Relation != "" {
			return true
		}
	}
	return false
}

func (e *Engine) astRelationFor(ctx context.Context, r *relationTuple) (*ast.Relation, error) {
	namespaceManager, err := e.d.Config(ctx).NamespaceManager()
	if err != nil {
		return nil, err
	}
	return namespace.ASTRelationFor(ctx, namespaceManager, r.Namespace, r.Relation)
}
