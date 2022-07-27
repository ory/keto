package check

import (
	"context"

	"github.com/pkg/errors"

	"github.com/ory/keto/internal/check/checkgroup"
	"github.com/ory/keto/internal/expand"
	"github.com/ory/keto/internal/namespace/ast"
	"github.com/ory/keto/internal/x"
)

func checkNotImplemented(_ context.Context, resultCh chan<- checkgroup.Result) {
	resultCh <- checkgroup.Result{Err: errors.WithStack(errors.New("not implemented"))}
}

func toExpandNodeType(op ast.Operator) expand.NodeType {
	switch op {
	case ast.OperatorOr:
		return expand.Union
	case ast.OperatorAnd:
		return expand.Intersection
	default:
		return expand.Union
	}
}

func (e *Engine) checkUsersetRewrite(
	ctx context.Context,
	tuple *RelationTuple,
	rewrite *ast.UsersetRewrite,
	restDepth int,
) checkgroup.CheckFunc {
	if restDepth < 0 {
		e.d.Logger().Debug("reached max-depth, therefore this query will not be further expanded")
		return checkgroup.UnknownMemberFunc
	}

	e.d.Logger().
		WithField("request", tuple.String()).
		Trace("check userset rewrite")

	var (
		op     binaryOperator
		checks []checkgroup.CheckFunc
	)
	switch rewrite.Operation {
	case ast.OperatorOr:
		op = or
	case ast.OperatorAnd:
		op = and
	default:
		return checkNotImplemented
	}

	for _, child := range rewrite.Children {
		switch c := child.(type) {

		case *ast.TupleToUserset:
			checks = append(checks, checkgroup.WithEdge(checkgroup.Edge{
				Tuple: *tuple,
				Type:  expand.TupeToUserset,
			}, e.checkTupleToUserset(tuple, c, restDepth)))

		case *ast.ComputedUserset:
			checks = append(checks, checkgroup.WithEdge(checkgroup.Edge{
				Tuple: *tuple,
				Type:  expand.ComputedUserset,
			}, e.checkComputedUserset(ctx, tuple, c, restDepth)))

		case *ast.UsersetRewrite:
			checks = append(checks, checkgroup.WithEdge(checkgroup.Edge{
				Tuple: *tuple,
				Type:  toExpandNodeType(c.Operation),
			}, e.checkUsersetRewrite(ctx, tuple, c, restDepth)))

		case *ast.InvertResult:
			checks = append(checks, checkgroup.WithEdge(checkgroup.Edge{
				Tuple: *tuple,
				Type:  expand.Not,
			}, e.checkInverted(ctx, tuple, c, restDepth)))

		default:
			return checkNotImplemented
		}
	}

	return func(ctx context.Context, resultCh chan<- checkgroup.Result) {
		resultCh <- op(ctx, checks)
	}
}

func (e *Engine) checkInverted(
	ctx context.Context,
	tuple *RelationTuple,
	inverted *ast.InvertResult,
	restDepth int,
) checkgroup.CheckFunc {
	if restDepth < 0 {
		e.d.Logger().Debug("reached max-depth, therefore this query will not be further expanded")
		return checkgroup.UnknownMemberFunc
	}

	e.d.Logger().
		WithField("request", tuple.String()).
		Trace("invert check")

	var check checkgroup.CheckFunc

	switch c := inverted.Child.(type) {

	case *ast.TupleToUserset:
		check = checkgroup.WithEdge(checkgroup.Edge{
			Tuple: *tuple,
			Type:  expand.TupeToUserset,
		}, e.checkTupleToUserset(tuple, c, restDepth))

	case *ast.ComputedUserset:
		check = checkgroup.WithEdge(checkgroup.Edge{
			Tuple: *tuple,
			Type:  expand.ComputedUserset,
		}, e.checkComputedUserset(ctx, tuple, c, restDepth))

	case *ast.UsersetRewrite:
		check = checkgroup.WithEdge(checkgroup.Edge{
			Tuple: *tuple,
			Type:  toExpandNodeType(c.Operation),
		}, e.checkUsersetRewrite(ctx, tuple, c, restDepth))

	case *ast.InvertResult:
		check = checkgroup.WithEdge(checkgroup.Edge{
			Tuple: *tuple,
			Type:  expand.Not,
		}, e.checkInverted(ctx, tuple, c, restDepth))

	default:
		return checkNotImplemented
	}

	return func(ctx context.Context, resultCh chan<- checkgroup.Result) {
		innerCh := make(chan checkgroup.Result)
		go check(ctx, innerCh)
		select {
		case result := <-innerCh:
			// invert result here
			switch result.Membership {
			case checkgroup.IsMember:
				result.Membership = checkgroup.NotMember
			case checkgroup.NotMember:
				result.Membership = checkgroup.IsMember
			}
			resultCh <- result
		case <-ctx.Done():
			resultCh <- checkgroup.Result{Err: errors.WithStack(ctx.Err())}
		}
	}
}

// checkComputedUserset rewrites the relation tuple to use the userset relation
// instead of the the relation from the tuple.
//
// A relation tuple n:obj#original_rel@user is rewritten to
// n:obj#userset@user, where the 'userset' relation is taken from the
// userset.Relation.
func (e *Engine) checkComputedUserset(
	ctx context.Context,
	r *RelationTuple,
	userset *ast.ComputedUserset,
	restDepth int,
) checkgroup.CheckFunc {
	if restDepth < 0 {
		e.d.Logger().Debug("reached max-depth, therefore this query will not be further expanded")
		return checkgroup.UnknownMemberFunc
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

// checkTupleToUserset rewrites the relation tuple to use the userset relation.
//
// Given a relation tuple like docs:readme#editor@user, and a tuple-to-userset
// rewrite with the relation "parent" and the computed userset relation
// "owner", the following checks will be performed:
//
// * query for all tuples like docs:readme#parent@??? to get a list of subjects
//   that have the parent relation on docs:readme
//
// * For each matching subject, then check if subject#owner@user.
func (e *Engine) checkTupleToUserset(
	tuple *RelationTuple,
	userset *ast.TupleToUserset,
	restDepth int,
) checkgroup.CheckFunc {
	if restDepth < 0 {
		e.d.Logger().Debug("reached max-depth, therefore this query will not be further expanded")
		return checkgroup.UnknownMemberFunc
	}

	e.d.Logger().
		WithField("request", tuple.String()).
		WithField("tuple to userset relation", userset.Relation).
		WithField("tuple to userset computed", userset.ComputedUsersetRelation).
		Trace("check tuple to userset")

	return func(ctx context.Context, resultCh chan<- checkgroup.Result) {
		var (
			prevPage, nextPage string
			tuples             []*RelationTuple
			err                error
		)
		g := checkgroup.New(ctx)
		for nextPage = "x"; nextPage != "" && !g.Done(); prevPage = nextPage {
			tuples, nextPage, err = e.d.RelationTupleManager().GetRelationTuples(
				ctx,
				&Query{
					Namespace: tuple.Namespace,
					Object:    tuple.Object,
					Relation:  userset.Relation,
				},
				x.WithToken(prevPage))
			if err != nil {
				g.Add(checkgroup.ErrorFunc(err))
				return
			}

			for _, t := range tuples {
				if t.Subject.SubjectSet() == nil {
					continue
				}
				g.Add(e.checkIsAllowed(
					ctx,
					&RelationTuple{
						Namespace: t.Subject.SubjectSet().Namespace,
						Object:    t.Subject.SubjectSet().Object,
						Relation:  userset.ComputedUsersetRelation,
						Subject:   tuple.Subject,
					},
					restDepth-1,
				))
			}
		}
		resultCh <- g.Result()
	}
}
