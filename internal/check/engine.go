package check

import (
	"context"
	"errors"
	"time"

	"github.com/ory/herodot"

	"github.com/ory/keto/internal/driver/config"
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
	RelationTuple = relationtuple.InternalRelationTuple
	Query         = relationtuple.RelationQuery
)

func NewEngine(d EngineDependencies) *Engine {
	return &Engine{
		d: d,
	}
}

func (e *Engine) checkDirect(
	ctx context.Context,
	requested *RelationTuple,
	rels []*RelationTuple,
	restDepth int,
) checkFn {
	// This is the same as the graph problem "can requested.Subject be reached from requested.Object through the first outgoing edge requested.Relation"
	//
	// We implement recursive depth-first search here.
	// TODO replace by more performant algorithm: https://github.com/ory/keto/issues/483

	var indirectChecks []checkFn

	for _, sr := range rels {
		ctx, wasAlreadyVisited := graph.CheckAndAddVisited(ctx, sr.Subject)
		if wasAlreadyVisited {
			continue
		}

		// we only have to check Subject here as we know that sr was reached from requested.ObjectID, requested.Relation through 0...n indirections
		if requested.Subject.Equals(sr.Subject) {
			// found the requested relation
			return isMemberCheckFn
		}

		sub, isSubjectSet := sr.Subject.(*relationtuple.SubjectSet)
		if !isSubjectSet {
			continue
		}

		// expand the set by one indirection; paginated

		// TODO(hperl): Convert everything to concurrent request.
		indirectChecks = append(indirectChecks, e.checkOneIndirectionFurther(
			ctx,
			requested,
			&Query{Object: sub.Object, Relation: sub.Relation, Namespace: sub.Namespace},
			restDepth-1,
		))
	}
	return unionCheckFn(ctx, indirectChecks)
}

func (e *Engine) checkOneIndirectionFurther(
	ctx context.Context,
	requested *RelationTuple,
	expandQuery *Query,
	restDepth int,
) checkFn {
	return func(ctx context.Context, resultCh chan<- checkResult) {
		e.d.Logger().
			WithField("request", requested.String()).
			WithField("query", expandQuery.String()).
			Trace("check one direction further")

		if restDepth < 0 {
			e.d.Logger().WithFields(requested.ToLoggerFields()).Debug("reached max-depth, therefore this query will not be further expanded")
			resultCh <- ResultNotMember
			return
		}

		// an empty page token denotes the first page (as tokens are opaque)
		var prevPage string

		subcheckCh := make(chan checkResult)

		// used to report to the number of subchecks (which we only know after all pages)
		totalNumOfSubchecksCh := make(chan int)
		consumer := func() {
			subCheckCount := 0
			totalNumOfSubchecks := 0
			for {
				select {
				case result := <-subcheckCh:
					subCheckCount++
					if result.Err != nil || result.Membership == IsMember || totalNumOfSubchecks == subCheckCount {
						resultCh <- result
						return
					}
				case totalNumOfSubchecks = <-totalNumOfSubchecksCh:
					if totalNumOfSubchecks == subCheckCount {
						resultCh <- ResultNotMember
						return
					}

				case <-ctx.Done():
					resultCh <- checkResult{Err: context.Canceled}
					return
				}
			}
		}
		go consumer()

		totalNumOfSubchecks := 0
		for {
			totalNumOfSubchecks++
			nextRels, nextPage, err := e.d.RelationTupleManager().GetRelationTuples(ctx, expandQuery, x.WithToken(prevPage))
			// herodot.ErrNotFound occurs when the namespace is unknown
			if errors.Is(err, herodot.ErrNotFound) {
				subcheckCh <- ResultNotMember
				break
			} else if err != nil {
				subcheckCh <- checkResult{Err: err}
				break
			}

			check := e.checkDirect(ctx, requested, nextRels, restDepth-1)
			go check(ctx, subcheckCh)

			// loop through pages until either allowed, end of pages, or an error occurred
			if nextPage == "" {
				break
			}
			prevPage = nextPage
		}
		// Now that we know the total amount of subchecks we need to once report it.
		totalNumOfSubchecksCh <- totalNumOfSubchecks
	}
}

type membership int

const (
	MembershipUnknown membership = iota
	IsMember
	NotMember
)

type checkResult struct {
	Membership membership
	Err        error
}

var (
	ResultIsMember  = checkResult{Membership: IsMember}
	ResultNotMember = checkResult{Membership: NotMember}
)

type checkFn func(ctx context.Context, resultCh chan<- checkResult)

func (e *Engine) SubjectIsAllowed(ctx context.Context, r *RelationTuple, restDepth int) (bool, error) {
	// global max-depth takes precedence when it is the lesser or if the request
	// max-depth is less than or equal to 0
	if globalMaxDepth := e.d.Config(ctx).MaxReadDepth(); restDepth <= 0 || globalMaxDepth < restDepth {
		restDepth = globalMaxDepth
	}
	result := union(ctx, []checkFn{e.checkIsAllowed(ctx, r, restDepth+1)})

	return result.Membership == IsMember, result.Err
}

func errorCheckFn(err error) checkFn {
	return func(_ context.Context, resultCh chan<- checkResult) {
		resultCh <- checkResult{Err: err}
	}
}

var isMemberCheckFn checkFn = func(_ context.Context, resultCh chan<- checkResult) {
	resultCh <- checkResult{Membership: IsMember}
}

func (e *Engine) checkIsAllowed(ctx context.Context, r *RelationTuple, restDepth int) checkFn {
	e.d.Logger().
		WithField("request", r.String()).
		Trace("check is allowed")

	directFn := e.checkOneIndirectionFurther(ctx, r,
		&Query{
			Object:    r.Object,
			Relation:  r.Relation,
			Namespace: r.Namespace,
		}, restDepth)

	checks := []checkFn{directFn}

	relation, err := e.astRelationFor(ctx, r)
	if err == nil && relation.UsersetRewrite != nil {
		checks = append(checks, e.checkUsersetRewrite(ctx, r, relation.UsersetRewrite))
	}

	return unionCheckFn(ctx, checks)
}

func unionCheckFn(ctx context.Context, checks []checkFn) checkFn {
	return func(ctx context.Context, resultCh chan<- checkResult) {
		resultCh <- union(ctx, checks)
	}
}

func checkNotImplemented(_ context.Context, resultCh chan<- checkResult) {
	resultCh <- checkResult{Err: errors.New("not implemented")}
}

type setOperation func(ctx context.Context, checks []checkFn) checkResult

func (e *Engine) checkUsersetRewrite(ctx context.Context, r *RelationTuple, rewrite *ast.UsersetRewrite) checkFn {
	e.d.Logger().
		WithField("request", r.String()).
		Trace("check userset rewrite")

	var (
		op     setOperation
		checks []checkFn
	)
	switch rewrite.Operation {
	case ast.SetOperationUnion:
		op = union
	case ast.SetOperationExclusion:
		return checkNotImplemented
	case ast.SetOperationIntersection:
		return checkNotImplemented
	default:
		return checkNotImplemented
	}

	for _, c := range rewrite.Children.ComputedUsersets {
		c := c
		checks = append(checks, e.checkComputedUserset(ctx, r, &c))
	}
	for _, c := range rewrite.Children.TupleToUsersets {
		c := c
		checks = append(checks, e.checkTupleToUserset(ctx, r, &c))
	}

	return func(ctx context.Context, resultCh chan<- checkResult) {
		resultCh <- op(ctx, checks)
	}
}

func (e *Engine) checkComputedUserset(ctx context.Context, r *RelationTuple, userset *ast.ComputedUserset) checkFn {
	e.d.Logger().
		WithField("request", r.String()).
		WithField("computed userset relation", userset.Relation).
		Trace("check computed userset")

	return e.checkIsAllowed(
		ctx,
		&RelationTuple{
			Namespace: r.Namespace,
			Object:    r.Object,
			Relation:  userset.Relation,
			Subject:   r.Subject,
		},
		100,
	)
}

func (e *Engine) checkTupleToUserset(ctx context.Context, r *RelationTuple, userset *ast.TupleToUserset) checkFn {
	e.d.Logger().
		WithField("request", r.String()).
		WithField("tuple to userset relation", userset.Relation).
		WithField("tuple to userset computed", userset.ComputedUsersetRelation).
		Trace("check tuple to userset")

	return func(ctx context.Context, resultCh chan<- checkResult) {
		var (
			prevPage, nextPage string
			rts, rtsPage       []*RelationTuple
			err                error
		)
		for nextPage = "x"; nextPage != ""; prevPage = nextPage {
			rtsPage, nextPage, err = e.d.RelationTupleManager().GetRelationTuples(
				ctx,
				&Query{
					Namespace: r.Namespace,
					Object:    r.Object,
					Relation:  userset.Relation,
				},
				x.WithToken(prevPage)) // TODO pagination
			if err != nil {
				resultCh <- checkResult{Err: err}
				return
			}
			rts = append(rts, rtsPage...)
		}
		checks := []checkFn{}
		for _, rt := range rts {
			if rt.Subject.SubjectSet() == nil {
				continue
			}
			checks = append(checks, e.checkIsAllowed(
				ctx,
				&RelationTuple{
					Namespace: rt.Subject.SubjectSet().Namespace,
					Object:    rt.Subject.SubjectSet().Object,
					Relation:  userset.ComputedUsersetRelation,
					Subject:   r.Subject,
				},
				100,
			))
		}
		resultCh <- union(ctx, checks)
	}
}

func (e *Engine) astRelationFor(ctx context.Context, r *RelationTuple) (*ast.Relation, error) {
	ns, err := e.namespaceFor(ctx, r)
	if err != nil {
		return nil, err
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

func union(ctx context.Context, checks []checkFn) checkResult {
	if len(checks) == 0 {
		return ResultNotMember
	}

	resultCh := make(chan checkResult, len(checks))
	childCtx, cancelFn := context.WithTimeout(ctx, 1*time.Second)
	defer cancelFn()

	for _, check := range checks {
		go check(childCtx, resultCh)
	}

	for i := 0; i < len(checks); i++ {
		select {
		case result := <-resultCh:
			// We return either the first error or the first success.
			if result.Err != nil || result.Membership == IsMember {
				return result
			}
		case <-ctx.Done():
			return checkResult{Err: context.Canceled}
		}
	}

	return ResultNotMember
}
