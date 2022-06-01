package check

import (
	"context"
	"errors"

	"github.com/ory/herodot"

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
) checkgroup.Func {
	if restDepth < 0 {
		e.d.Logger().
			Debug("reached max-depth, therefore this query will not be further expanded")
		return func(_ context.Context, resultCh chan<- checkgroup.Result) {
			resultCh <- checkgroup.Result{Membership: checkgroup.MembershipUnknown}
		}
	}

	// This is the same as the graph problem "can requested.Subject be reached
	// from requested.Object through the first outgoing edge requested.Relation"
	//
	// We implement recursive depth-first search here.
	// TODO replace by more performant algorithm:
	// https://github.com/ory/keto/issues/483

	g := checkgroup.New(ctx)

	for _, sr := range rels {
		ctx, wasAlreadyVisited := graph.CheckAndAddVisited(ctx, sr.Subject)
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

		g.Add(e.checkOneIndirectionFurther(
			ctx,
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

func (e *Engine) checkOneIndirectionFurther(
	ctx context.Context,
	requested *RelationTuple,
	expandQuery *Query,
	restDepth int,
) checkgroup.Func {
	if restDepth < 0 {
		e.d.Logger().
			WithFields(requested.ToLoggerFields()).
			Debug("reached max-depth, therefore this query will not be further expanded")
		return func(_ context.Context, resultCh chan<- checkgroup.Result) {
			resultCh <- checkgroup.Result{Membership: checkgroup.MembershipUnknown}
		}
	}

	return func(ctx context.Context, resultCh chan<- checkgroup.Result) {
		e.d.Logger().
			WithField("request", requested.String()).
			WithField("query", expandQuery.String()).
			Trace("check one indirection further")

		// an empty page token denotes the first page (as tokens are opaque)
		var prevPage string

		g := checkgroup.New(ctx)

		for {
			nextRels, nextPage, err := e.d.RelationTupleManager().GetRelationTuples(ctx, expandQuery, x.WithToken(prevPage))
			// herodot.ErrNotFound occurs when the namespace is unknown
			if errors.Is(err, herodot.ErrNotFound) {
				g.Add(checkgroup.NotMemberFunc)
				break
			} else if err != nil {
				g.Add(checkgroup.ErrorFunc(err))
				break
			}

			g.Add(e.checkDirect(ctx, requested, nextRels, restDepth-1))

			// loop through pages until either allowed, end of pages, or an error occurred
			if nextPage == "" || g.Done() {
				break
			}
			prevPage = nextPage
		}

		resultCh <- g.Result()
	}
}

func (e *Engine) SubjectIsAllowed(ctx context.Context, r *RelationTuple, restDepth int) (bool, error) {
	result := e.Check(ctx, r, restDepth)
	return result.Membership == checkgroup.IsMember, result.Err
}

func (e *Engine) Check(ctx context.Context, r *RelationTuple, restDepth int) checkgroup.Result {
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
		return checkgroup.Result{Err: context.Canceled}
	}
}

func (e *Engine) checkIsAllowed(ctx context.Context, r *RelationTuple, restDepth int) checkgroup.Func {
	if restDepth < 0 {
		e.d.Logger().
			Debug("reached max-depth, therefore this query will not be further expanded")
		return func(_ context.Context, resultCh chan<- checkgroup.Result) {
			resultCh <- checkgroup.Result{Membership: checkgroup.MembershipUnknown}
		}
	}

	e.d.Logger().
		WithField("request", r.String()).
		Trace("check is allowed")

	g := checkgroup.New(ctx)
	g.Add(e.checkOneIndirectionFurther(ctx, r,
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

func checkNotImplemented(_ context.Context, resultCh chan<- checkgroup.Result) {
	resultCh <- checkgroup.Result{Err: errors.New("not implemented")}
}

type setOperation func(ctx context.Context, checks []checkgroup.Func) checkgroup.Result

func (e *Engine) checkUsersetRewrite(
	ctx context.Context,
	r *RelationTuple,
	rewrite *ast.UsersetRewrite,
	restDepth int,
) checkgroup.Func {
	if restDepth < 0 {
		e.d.Logger().
			Debug("reached max-depth, therefore this query will not be further expanded")
		return func(_ context.Context, resultCh chan<- checkgroup.Result) {
			resultCh <- checkgroup.Result{Membership: checkgroup.MembershipUnknown}
		}
	}

	e.d.Logger().
		WithField("request", r.String()).
		Trace("check userset rewrite")

	var (
		op     setOperation
		checks []checkgroup.Func
	)
	switch rewrite.Operation {
	case ast.SetOperationUnion:
		op = or
	case ast.SetOperationDifference:
		op = butNot
	case ast.SetOperationIntersection:
		op = and
	default:
		return checkNotImplemented
	}

	for _, child := range rewrite.Children {
		switch c := child.(type) {
		case ast.TupleToUserset:
			checks = append(checks, checkgroup.WithEdge(checkgroup.Edge{
				Tuple: *r,
				Type:  expand.TupeToUserset,
			}, e.checkTupleToUserset(ctx, r, &c, restDepth)))
		case ast.ComputedUserset:
			checks = append(checks, checkgroup.WithEdge(checkgroup.Edge{
				Tuple: *r,
				Type:  expand.ComputedUserset,
			}, e.checkComputedUserset(ctx, r, &c, restDepth)))
		}
	}

	return func(ctx context.Context, resultCh chan<- checkgroup.Result) {
		resultCh <- op(ctx, checks)
	}
}

func (e *Engine) checkComputedUserset(
	ctx context.Context,
	r *RelationTuple,
	userset *ast.ComputedUserset,
	restDepth int,
) checkgroup.Func {
	if restDepth < 0 {
		e.d.Logger().
			Debug("reached max-depth, therefore this query will not be further expanded")
		return func(_ context.Context, resultCh chan<- checkgroup.Result) {
			resultCh <- checkgroup.Result{Membership: checkgroup.MembershipUnknown}
		}
	}

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
		restDepth,
	)
}

func (e *Engine) checkTupleToUserset(
	ctx context.Context,
	r *RelationTuple,
	userset *ast.TupleToUserset,
	restDepth int,
) checkgroup.Func {
	if restDepth < 0 {
		e.d.Logger().
			Debug("reached max-depth, therefore this query will not be further expanded")
		return func(_ context.Context, resultCh chan<- checkgroup.Result) {
			resultCh <- checkgroup.Result{Membership: checkgroup.MembershipUnknown}
		}
	}

	e.d.Logger().
		WithField("request", r.String()).
		WithField("tuple to userset relation", userset.Relation).
		WithField("tuple to userset computed", userset.ComputedUsersetRelation).
		Trace("check tuple to userset")

	return func(ctx context.Context, resultCh chan<- checkgroup.Result) {
		var (
			prevPage, nextPage string
			rts                []*RelationTuple
			err                error
		)
		g := checkgroup.New(ctx)
		for nextPage = "x"; nextPage != "" && !g.Done(); prevPage = nextPage {
			rts, nextPage, err = e.d.RelationTupleManager().GetRelationTuples(
				ctx,
				&Query{
					Namespace: r.Namespace,
					Object:    r.Object,
					Relation:  userset.Relation,
				},
				x.WithToken(prevPage))
			if err != nil {
				g.Add(checkgroup.ErrorFunc(err))
				return
			}

			for _, rt := range rts {
				if rt.Subject.SubjectSet() == nil {
					continue
				}
				g.Add(e.checkIsAllowed(
					ctx,
					&RelationTuple{
						Namespace: rt.Subject.SubjectSet().Namespace,
						Object:    rt.Subject.SubjectSet().Object,
						Relation:  userset.ComputedUsersetRelation,
						Subject:   r.Subject,
					},
					restDepth-1,
				))
			}
		}
		resultCh <- g.Result()
	}
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

func or(ctx context.Context, checks []checkgroup.Func) checkgroup.Result {
	if len(checks) == 0 {
		return checkgroup.ResultNotMember
	}

	resultCh := make(chan checkgroup.Result, len(checks))
	childCtx, cancelFn := context.WithCancel(ctx)
	defer cancelFn()

	for _, check := range checks {
		go check(childCtx, resultCh)
	}

	for i := 0; i < len(checks); i++ {
		select {
		case result := <-resultCh:
			// We return either the first error or the first success.
			if result.Err != nil || result.Membership == checkgroup.IsMember {
				return result
			}
		case <-ctx.Done():
			return checkgroup.Result{Err: context.Canceled}
		}
	}

	return checkgroup.ResultNotMember
}

func and(ctx context.Context, checks []checkgroup.Func) checkgroup.Result {
	if len(checks) == 0 {
		return checkgroup.ResultNotMember
	}

	resultCh := make(chan checkgroup.Result, len(checks))
	childCtx, cancelFn := context.WithCancel(ctx)
	defer cancelFn()

	for _, check := range checks {
		go check(childCtx, resultCh)
	}

	tree := &expand.Tree{
		Type:     expand.Intersection,
		Children: []*expand.Tree{},
	}

	for i := 0; i < len(checks); i++ {
		select {
		case result := <-resultCh:
			// We return fast on either an error or if a subcheck returns "not a
			// member".
			if result.Err != nil || result.Membership != checkgroup.IsMember {
				return checkgroup.Result{Err: result.Err, Membership: checkgroup.NotMember}
			} else {
				tree.Children = append(tree.Children, result.Tree)
			}
		case <-ctx.Done():
			return checkgroup.Result{Err: context.Canceled}
		}
	}

	return checkgroup.Result{
		Membership: checkgroup.IsMember,
		Tree:       tree,
	}
}

// butNot returns "is member" if and only if the first check returns "is member"
// and all subsequent checks return "not member".
func butNot(ctx context.Context, checks []checkgroup.Func) checkgroup.Result {
	if len(checks) < 2 {
		return checkgroup.ResultNotMember
	}

	expectMemberCh := make(chan checkgroup.Result, 1)
	expectNotMemberCh := make(chan checkgroup.Result, len(checks)-1)
	childCtx, cancelFn := context.WithCancel(ctx)
	defer cancelFn()

	go checks[0](childCtx, expectMemberCh)
	for _, check := range checks[1:] {
		go check(childCtx, expectNotMemberCh)
	}

	tree := &expand.Tree{
		Type:     expand.Exclusion,
		Children: []*expand.Tree{},
	}

	for i := 0; i < len(checks); i++ {
		select {
		case result := <-expectMemberCh:
			if result.Err != nil || result.Membership == checkgroup.NotMember {
				return checkgroup.Result{Err: result.Err, Membership: checkgroup.NotMember}
			} else {
				tree.Children = append(tree.Children, result.Tree)
			}
		case result := <-expectNotMemberCh:
			// We return fast on either an error or if a subcheck returns "not a
			// member".
			if result.Err != nil || result.Membership == checkgroup.IsMember {
				return checkgroup.Result{Err: result.Err, Membership: checkgroup.NotMember}
			} else {
				tree.Children = append(tree.Children, result.Tree)
			}
		case <-ctx.Done():
			return checkgroup.Result{Err: context.Canceled}
		}
	}

	return checkgroup.Result{
		Membership: checkgroup.IsMember,
		Tree:       tree,
	}
}
