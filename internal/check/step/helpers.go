// Copyright © 2026 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package step

import (
	"context"

	"github.com/ory/keto/internal/check"
	"github.com/ory/keto/internal/namespace"
	"github.com/ory/keto/internal/namespace/ast"
	"github.com/ory/keto/internal/relationtuple"
)

// astRelationFor looks up the AST Relation for the given tuple's namespace and
// relation. Returns nil (not an error) when no namespace config is present.
func astRelationFor(ctx context.Context, c check.EngineDependencies, r *relationTuple) (*ast.Relation, error) {
	namespaceManager, err := c.Config(ctx).NamespaceManager()
	if err != nil {
		return nil, err
	}
	return namespace.ASTRelationFor(ctx, namespaceManager, r.Namespace, r.Relation)
}

// maxDepthReached logs a debug message and returns MembershipUnknown.
// Call it when RestDepth is exhausted and the check cannot be expanded further.
func maxDepthReached(ex check.Executor, req check.CheckRequest) check.Result {
	ex.Deps().Logger().
		WithField("request", req.Tuple.String()).
		Debug("reached max-depth, therefore this query will not be further expanded")

	return check.Result{Membership: check.MembershipUnknown, Limitation: check.LimitationMaxDepthExceeded}
}

// subjectSetTypesFor returns the SubjectSet types declared for the relation in OPL.
func subjectSetTypesFor(relation *ast.Relation) []relationtuple.SubjectSetType {
	if relation == nil {
		return nil
	}
	subjectSets := make([]relationtuple.SubjectSetType, 0)
	for _, t := range relation.Types {
		if t.Relation != "" {
			subjectSets = append(subjectSets, relationtuple.SubjectSetType{
				Namespace: t.Namespace,
				Relation:  t.Relation,
			})
		}
	}
	return subjectSets
}
