// Copyright © 2022 Ory Corp

package check

import (
	"context"
	"fmt"

	"github.com/ory/herodot"
	"github.com/pkg/errors"

	"github.com/ory/keto/internal/check/checkgroup"
	"github.com/ory/keto/internal/driver/config"
	"github.com/ory/keto/internal/namespace"
	"github.com/ory/keto/internal/namespace/ast"
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
		config.Provider
		x.LoggerProvider
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
func (e *Engine) CheckRelationTuple(ctx context.Context, r *relationTuple, restDepth int) checkgroup.Result {
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

// checkExpandSubject checks the expansions of the subject set of the tuple.
//
// For a relation tuple n:obj#rel@user, checkExpandSubject first queries for all
// subjects that match n:obj#rel@* (arbitrary subjects), and then for each
// subject set checks subject@user.
func (e *Engine) checkExpandSubject(r *relationTuple, restDepth int) checkgroup.CheckFunc {
	if restDepth < 0 {
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

		var (
			subjects  []*relationTuple
			pageToken string
			err       error
			visited   bool
			innerCtx  = graph.InitVisited(ctx)
			query     = &query{Namespace: &r.Namespace, Object: &r.Object, Relation: &r.Relation}
		)
		for {
			subjects, pageToken, err = e.d.RelationTupleManager().GetRelationTuples(innerCtx, query, x.WithToken(pageToken))
			if errors.Is(err, herodot.ErrNotFound) {
				g.Add(checkgroup.NotMemberFunc)
				break
			} else if err != nil {
				g.Add(checkgroup.ErrorFunc(err))
				break
			}
			for _, s := range subjects {
				innerCtx, visited = graph.CheckAndAddVisited(innerCtx, s.Subject)
				if visited {
					continue
				}
				subjectSet, ok := s.Subject.(*relationtuple.SubjectSet)
				if !ok || subjectSet.Relation == "" {
					continue
				}
				g.Add(e.checkIsAllowed(
					innerCtx,
					&relationTuple{
						Namespace: subjectSet.Namespace,
						Object:    subjectSet.Object,
						Relation:  subjectSet.Relation,
						Subject:   r.Subject,
					},
					restDepth-1,
				))
			}
			if pageToken == "" || g.Done() {
				break
			}
		}

		resultCh <- g.Result()
	}
}

// checkDirect checks if the relation tuple is in the database directly.
func (e *Engine) checkDirect(r *relationTuple, restDepth int) checkgroup.CheckFunc {
	if restDepth < 0 {
		e.d.Logger().
			WithField("method", "checkDirect").
			Debug("reached max-depth, therefore this query will not be further expanded")
		return checkgroup.UnknownMemberFunc
	}
	return func(ctx context.Context, resultCh chan<- checkgroup.Result) {
		e.d.Logger().
			WithField("request", r.String()).
			Trace("check direct")
		if rels, _, err := e.d.RelationTupleManager().GetRelationTuples(
			ctx,
			r.ToQuery(),
			x.WithSize(1),
		); err == nil && len(rels) > 0 {
			resultCh <- checkgroup.Result{
				Membership: checkgroup.IsMember,
				Tree: &ketoapi.Tree[*relationtuple.RelationTuple]{
					Type:  ketoapi.TreeNodeLeaf,
					Tuple: r,
				},
			}
		} else {
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
func (e *Engine) checkIsAllowed(ctx context.Context, r *relationTuple, restDepth int) checkgroup.CheckFunc {
	if restDepth < 0 {
		e.d.Logger().
			WithField("method", "checkIsAllowed").
			Debug("reached max-depth, therefore this query will not be further expanded")
		return checkgroup.UnknownMemberFunc
	}

	e.d.Logger().
		WithField("request", r.String()).
		Trace("check is allowed")

	g := checkgroup.New(ctx)
	g.Add(e.checkDirect(r, restDepth-1))
	g.Add(e.checkExpandSubject(r, restDepth))

	relation, err := e.astRelationFor(ctx, r)
	if err != nil {
		g.Add(checkgroup.ErrorFunc(err))
	} else if relation != nil && relation.SubjectSetRewrite != nil {
		g.Add(e.checkSubjectSetRewrite(ctx, r, relation.SubjectSetRewrite, restDepth))
	}

	return g.CheckFunc()
}

func (e *Engine) astRelationFor(ctx context.Context, r *relationTuple) (*ast.Relation, error) {
	// Special case: If the relationTuple's relation is empty, then it is not an
	// error that the relation was not found.
	if r.Relation == "" {
		return nil, nil
	}

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
	return nil, fmt.Errorf("relation %q not found", r.Relation)
}

func (e *Engine) namespaceFor(ctx context.Context, r *relationTuple) (*namespace.Namespace, error) {
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
