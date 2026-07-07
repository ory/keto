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
func astRelationFor(ctx context.Context, nm namespace.Manager, r *relationTuple) (*ast.Relation, error) {
	return namespace.ASTRelationFor(ctx, nm, r.Namespace, r.Relation)
}

// maxDepthReached logs a debug message and returns MembershipUnknown.
// Call it when RestDepth is exhausted and the check cannot be expanded further.
func maxDepthReached(ex check.Executor, req check.CheckRequest) check.Result {
	ex.Deps().Logger().
		WithField("request", req.Tuple.String()).
		Debug("reached max-depth, therefore this query will not be further expanded")

	return check.Result{Membership: check.MembershipUnknown, Limitation: check.LimitationMaxDepthExceeded}
}

// subjectSetTypesFor returns the OPL-declared subject-set types for the relation.
// It looks up the AST for each next-level (namespace, relation) to determine
// whether the Subject can be direct-match there.
func subjectSetTypesFor(ctx context.Context, nm namespace.Manager, subject relationtuple.Subject, relation *ast.Relation) ([]relationtuple.SubjectSetType, error) {
	if relation == nil {
		return nil, nil
	}
	var types []relationtuple.SubjectSetType
	for _, t := range relation.Types {
		// empty relation means direct membership, so not a SubjectSet<> declaration
		if t.Relation == "" {
			continue
		}
		nextRel, err := namespace.ASTRelationFor(ctx, nm, t.Namespace, t.Relation)
		if err != nil {
			return nil, err
		}
		types = append(types, relationtuple.SubjectSetType{
			Namespace:    t.Namespace,
			Relation:     t.Relation,
			AllowsDirect: AllowsDirectMember(nextRel, subject),
		})
	}
	return types, nil
}

// AllowsDirectMember reports whether OPL allows a direct subject match for the given relation.
// A relation with a SubjectSetRewrite is computed, so no direct tuple conforms to OPL,
// regardless of the subject kind.
// SubjectID subjects have no namespace, so declared types cannot constrain them and true
// is returned conservatively.
// For SubjectSet subjects, both namespace and relation must match a declared type exactly.
func AllowsDirectMember(relation *ast.Relation, subject relationtuple.Subject) bool {
	s, isSubjectSet := subject.(*relationtuple.SubjectSet)
	if !isSubjectSet {
		// For SubjectID, allow direct when the relation is not computed via a rewrite.
		return relation == nil || relation.SubjectSetRewrite == nil
	}
	if relation == nil || relation.SubjectSetRewrite != nil {
		return false
	}
	for _, t := range relation.Types {
		if t.Namespace == s.Namespace && t.Relation == s.Relation {
			return true
		}
	}
	return false
}
