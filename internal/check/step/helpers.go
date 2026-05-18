// Copyright © 2026 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package step

import (
	"context"

	"github.com/ory/keto/internal/check"
	"github.com/ory/keto/internal/namespace"
	"github.com/ory/keto/internal/namespace/ast"
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
	return check.Result{Membership: check.MembershipUnknown}
}

// containsSubjectSetExpand reports whether the relation's type list includes
// at least one SubjectSet type (i.e. a type with a non-empty Relation field).
func containsSubjectSetExpand(relation *ast.Relation) bool {
	for _, t := range relation.Types {
		if t.Relation != "" {
			return true
		}
	}
	return false
}
