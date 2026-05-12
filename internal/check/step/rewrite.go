// Copyright © 2026 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package step

import (
	"context"

	"github.com/pkg/errors"

	"github.com/ory/keto/internal/check"
	"github.com/ory/keto/internal/namespace"
	"github.com/ory/keto/internal/namespace/ast"
	"github.com/ory/keto/internal/relationtuple"
)

// RewriteStep evaluates a SubjectSetRewrite node (OR or AND).
// It translates the rewrite's children into PlannedSteps via planRewrite and
// then delegates to RunUnion (OR) or RunIntersection (AND).
type RewriteStep struct {
	Rewrite *ast.SubjectSetRewrite
}

func (RewriteStep) Kind() check.StepKind { return check.StepRewrite }

func (s RewriteStep) Execute(ctx context.Context, req check.CheckRequest, ex check.Executor) check.Result {
	if req.RestDepth <= 0 {
		ex.Deps().Logger().Debug("reached max-depth, therefore this query will not be further expanded")
		return check.Result{Membership: check.MembershipUnknown}
	}

	ex.Deps().Logger().
		WithField("request", req.Tuple.String()).
		Trace("check subject-set rewrite")

	steps, err := planRewrite(ctx, ex, req, s.Rewrite)
	if err != nil {
		return check.Result{Err: err}
	}

	if s.Rewrite.Operation == ast.OperatorAnd {
		return ex.RunIntersection(ctx, steps...)
	}
	return ex.RunUnion(ctx, steps...)
}

// planRewrite translates a SubjectSetRewrite AST node into a flat list of
// PlannedSteps ready for execution. It applies the OR-shortcut optimization:
// ComputedSubjectSet children in an OR are batched into a single
// DirectMultiStep followed by individual IsAllowedStep calls.
func planRewrite(ctx context.Context, ex check.Executor, req check.CheckRequest, rewrite *ast.SubjectSetRewrite) ([]check.PlannedStep, error) {
	var steps []check.PlannedStep
	handled := make(map[int]struct{}, len(rewrite.Children))

	// OR shortcut: batch all ComputedSubjectSet children into one DirectMulti
	// check plus individual IsAllowed calls (matching the current engine's
	// OR-of-computed-usersets optimization).
	if rewrite.Operation == ast.OperatorOr {
		var computedRelations []string
		for i, child := range rewrite.Children {
			if c, ok := child.(*ast.ComputedSubjectSet); ok {
				handled[i] = struct{}{}
				computedRelations = append(computedRelations, c.Relation)
			}
		}

		if len(computedRelations) > 0 {
			namespaceManager, err := ex.Deps().Config(ctx).NamespaceManager()
			if err != nil {
				return nil, err
			}

			// In strict mode, skip direct checks for relations that have their
			// own subject-set rewrite; the rewrite handles them in memory.
			var directRelations []string
			isStrictMode := ex.Deps().Config(ctx).StrictMode()
			for _, rel := range computedRelations {
				astRel, _ := namespace.ASTRelationFor(ctx, namespaceManager, req.Tuple.Namespace, rel)
				if isStrictMode && astRel != nil && astRel.SubjectSetRewrite != nil {
					continue
				}
				directRelations = append(directRelations, rel)
			}

			if len(directRelations) > 0 {
				steps = append(steps, check.PlannedStep{
					Step: DirectMultiStep{Relations: directRelations},
					Req:  req,
				})
			}

			// Each computed relation also gets a full IsAllowed check
			// (skipDirect=true since DirectMulti already covers direct hits).
			for _, rel := range computedRelations {
				steps = append(steps, check.PlannedStep{
					Step: IsAllowedStep{skipDirect: true},
					Req: check.CheckRequest{
						Tuple: &relationtuple.RelationTuple{
							Namespace: req.Tuple.Namespace,
							Object:    req.Tuple.Object,
							Relation:  rel,
							Subject:   req.Tuple.Subject,
						},
						RestDepth: req.RestDepth - 1,
					},
				})
			}
		}
	}

	// Process remaining (non-shortcut) children.
	for i, child := range rewrite.Children {
		if _, ok := handled[i]; ok {
			continue
		}
		step, err := rewriteChildToStep(child)
		if err != nil {
			return nil, err
		}

		childReq := req
		switch child.(type) {
		case *ast.SubjectSetRewrite:
			childReq = req.WithDepth(req.RestDepth - 1)
		}

		steps = append(steps, check.PlannedStep{Step: step, Req: childReq})
	}

	return steps, nil
}

// rewriteChildToStep converts a single rewrite AST child node into an
// executable Step. Returns an error for unrecognised node types.
func rewriteChildToStep(child ast.Child) (check.Step, error) {
	switch c := child.(type) {
	case *ast.TupleToSubjectSet:
		return TraverseStep{TTU: c}, nil
	case *ast.ComputedSubjectSet:
		return ComputedUsersetStep{Relation: c.Relation}, nil
	case *ast.SubjectSetRewrite:
		return RewriteStep{Rewrite: c}, nil
	case *ast.InvertResult:
		return InvertStep{Child: c.Child}, nil
	default:
		return nil, errors.New("not implemented: unknown rewrite child type")
	}
}
