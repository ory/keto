// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package expand

import (
	"context"
	"slices"

	"github.com/ory/x/logrusx"
	"github.com/ory/x/otelx"
	keysetpagination "github.com/ory/x/pagination/keysetpagination_v2"

	"github.com/ory/keto/internal/driver/config"
	"github.com/ory/keto/internal/namespace"
	"github.com/ory/keto/internal/namespace/ast"
	"github.com/ory/keto/internal/relationtuple"
	"github.com/ory/keto/internal/x"
	"github.com/ory/keto/ketoapi"
	"github.com/ory/keto/x/events"
)

type (
	EngineDependencies interface {
		relationtuple.ManagerProvider
		config.Provider
		logrusx.Provider
		otelx.Provider
		x.NetworkIDProvider
	}
	Engine struct {
		d EngineDependencies
	}
	EngineProvider interface {
		ExpandEngine() *Engine
	}
)

func NewEngine(d EngineDependencies) *Engine {
	return &Engine{
		d: d,
	}
}

// treeBuilder holds state shared across a single BuildTree call.
type treeBuilder struct {
	d                EngineDependencies
	namespaceManager namespace.Manager
	nodeCount        int
	maxNodes         int
	visited          map[relationtuple.SubjectSet]struct{}
	isStrict         bool
	expandRewrites   bool
}

func (e *Engine) newTreeBuilder(ctx context.Context, namespaceManager namespace.Manager) *treeBuilder {
	return &treeBuilder{
		d:                e.d,
		namespaceManager: namespaceManager,
		visited:          make(map[relationtuple.SubjectSet]struct{}),
		maxNodes:         e.d.Config(ctx).ExpandMaxTupleCount(),
		nodeCount:        0,
		isStrict:         e.d.Config(ctx).StrictMode(),
		expandRewrites:   e.d.Config(ctx).ExpandRewrites(),
	}
}

func (e *Engine) BuildTree(ctx context.Context, subSet *relationtuple.SubjectSet, restDepth int) (t *relationtuple.Tree, err error) {
	ctx, span := e.d.Tracer(ctx).Tracer().Start(ctx, "Engine.BuildTree")
	defer otelx.End(span, &err)

	namespaceManager, err := e.d.Config(ctx).NamespaceManager()
	if err != nil {
		return nil, err
	}

	// global max-depth takes precedence when it is the lesser or if the request max-depth is less than or equal to 0
	if globalMaxDepth := e.d.Config(ctx).MaxReadDepth(); restDepth <= 0 || globalMaxDepth < restDepth {
		restDepth = globalMaxDepth
	}

	b := e.newTreeBuilder(ctx, namespaceManager)

	if b.isStrict {
		ns, _ := b.namespaceManager.GetNamespaceByName(ctx, subSet.Namespace)
		rel := ns.FindRelation(subSet.Relation)
		if rel == nil || (len(rel.Types) == 0 && rel.SubjectSetRewrite == nil) {
			return nil, nil
		}
	}

	t, err = b.buildTreeRecursive(ctx, subSet, restDepth)
	if err == nil {
		span.AddEvent(events.NewPermissionsExpanded(ctx))
	}
	return
}

func (b *treeBuilder) buildTreeRecursive(ctx context.Context, subject relationtuple.Subject, restDepth int) (*relationtuple.Tree, error) {
	subSet, isSubjectSet := subject.(*relationtuple.SubjectSet)
	if !isSubjectSet || subSet.Relation == "" {
		b.nodeCount++
		return newNode(ketoapi.TreeNodeLeaf, subject), nil
	}

	// Cycle detection: if this subject is already being expanded in the current
	// call path (DFS ancestry), we have a cycle.
	if _, exists := b.visited[*subSet]; exists {
		return &relationtuple.Tree{
			Type:    ketoapi.TreeNodeUnion,
			Subject: subject,
			Truncation: &relationtuple.Truncation{
				Reason: relationtuple.TruncationReasonCycle,
			},
		}, nil
	}
	b.visited[*subSet] = struct{}{}
	defer delete(b.visited, *subSet)

	ns, _ := b.namespaceManager.GetNamespaceByName(ctx, subSet.Namespace)
	rel := ns.FindRelation(subSet.Relation)

	if rel == nil {
		if b.isStrict {
			return nil, nil
		}
		// unknown relation, but we're not in strict mode, so we'll try to build it directly
		return b.buildTreeDirect(ctx, subSet, nil, nil, restDepth)
	}

	if rel.SubjectSetRewrite != nil {
		if b.expandRewrites {
			return b.buildTreeForRewrite(ctx, subSet, rel.SubjectSetRewrite, restDepth)
		}
		// it is a rewrite, but we don't expand rewrites
		// checking direct doesn't make much sense, but that's what we currently do so we keep that behavior
		return b.buildTreeDirect(ctx, subSet, rel, nil, restDepth)
	}

	// rel is not a "permits" but "related", so we check directly
	return b.buildTreeDirect(ctx, subSet, rel, nil, restDepth)
}

// buildTreeDirect expands by querying the relation directly, without an OPL rewrites.
func (b *treeBuilder) buildTreeDirect(ctx context.Context, subSet *relationtuple.SubjectSet, rel *ast.Relation, resume *relationtuple.ExpandCursor, restDepth int) (*relationtuple.Tree, error) {
	tree := newNode(ketoapi.TreeNodeUnion, subSet)

	directResume := func(pageToken keysetpagination.PageToken) *relationtuple.ExpandCursor {
		return &relationtuple.ExpandCursor{Kind: relationtuple.ExpandCursorKindDirect, SubjectSet: subSet, NextPageToken: pageToken}
	}

	if b.isMaxNodeLimit() {
		tree.Truncation = tupleLimitTruncation(directResume(keysetpagination.PageToken{}))
		return tree, nil
	}
	if restDepth <= 1 {
		tree.Truncation = depthLimitTruncation(directResume(keysetpagination.PageToken{}))
		return tree, nil
	}

	for nextPage, _ := keysetpagination.NewPaginator(); !nextPage.IsLast(); {
		var tuples []*relationtuple.RelationTuple
		var err error
		tuples, nextPage, err = b.d.RelationTupleManager().GetRelationTuples(ctx, &relationtuple.RelationQuery{
			Namespace: &subSet.Namespace,
			Object:    &subSet.Object,
			Relation:  &subSet.Relation,
		},
			append(nextPage.ToOptions(), paginatorOptionsFromResume(resume)...)...,
		)
		if err != nil {
			return nil, err
		} else if len(tuples) == 0 {
			break
		}

		pageStart := len(tree.Children)
		for _, tuple := range tuples {
			if b.isStrict && rel != nil { // if b.isStrict then rel != nil as we check that before calling this function, but let's better not rely on that invariant
				ss, ok := tuple.Subject.(*relationtuple.SubjectSet)
				if !ok || !slices.ContainsFunc(rel.Types, func(rt ast.RelationType) bool {
					return rt.Namespace == ss.Namespace && rt.Relation == ss.Relation
				}) {
					continue
				}
			}

			childTree := newNode(ketoapi.TreeNodeLeaf, tuple.Subject)
			if tupleSubSet, ok := tuple.Subject.(*relationtuple.SubjectSet); ok && tupleSubSet.Relation != "" {
				childTree.Type = ketoapi.TreeNodeUnion
			}

			b.nodeCount++
			tree.Children = append(tree.Children, childTree)
		}

		if b.isMaxNodeLimit() && !nextPage.IsLast() {
			tree.Truncation = tupleLimitTruncation(directResume(nextPage.PageToken()))
		}

		// expand all the non-leaf nodes founds in this page
		for _, child := range tree.Children[pageStart:] {
			if child.Type != ketoapi.TreeNodeLeaf {
				updated, err := b.buildTreeRecursive(ctx, child.Subject, restDepth-1)
				if err != nil {
					return nil, err
				}
				*child = *updated
			}
		}

		if tree.Truncation != nil {
			break
		}
	}

	return tree, nil
}

// buildTreeForRewrite expands a subject set using an OPL-defined rewrite rule.
// OPL rewrite children are structural (never paginated); only data nodes (direct
// relation tuples fetched from the DB) can be paginated.
func (b *treeBuilder) buildTreeForRewrite(ctx context.Context, subSet *relationtuple.SubjectSet, rewrite ast.Child, restDepth int) (*relationtuple.Tree, error) {
	switch rewrite := rewrite.(type) {
	case *ast.SubjectSetRewrite:
		tree := newNode(ketoapi.TreeNodeUnion, subSet)
		if rewrite.Operation == ast.OperatorAnd {
			tree.Type = ketoapi.TreeNodeIntersection
		}

		children := slices.SortedFunc(slices.Values(rewrite.Children), func(a, b ast.Child) int {
			return rewriteCostOrder(a) - rewriteCostOrder(b)
		})

		for _, childRewrite := range children {
			childTree, err := b.buildTreeForRewrite(ctx, subSet, childRewrite, restDepth)
			if err != nil {
				return nil, err
			}

			// SubjectSetRewrite followed by another SubjectSetRewrite means compound expression;
			// The subject of the 2nd rewrite is not meaningful at that point.
			// ex: view = A AND (B OR C) -> (B OR C) has no meaningul subject
			if childTree != nil {
				if _, ok := childRewrite.(*ast.SubjectSetRewrite); ok {
					childTree.Subject = nil
				}
				tree.Children = append(tree.Children, childTree)
			}
		}
		return tree, nil

	case *ast.ComputedSubjectSet:
		return b.buildTreeRecursive(ctx, &relationtuple.SubjectSet{
			Namespace: subSet.Namespace,
			Object:    subSet.Object,
			Relation:  rewrite.Relation,
		}, restDepth-1)
	case *ast.TupleToSubjectSet:
		return b.buildTreeTupleToSubjectSet(ctx, subSet, rewrite, nil, restDepth)
	case *ast.InvertResult:
		// inversion never has a meaningful Subject
		invertNode := newNode(ketoapi.TreeNodeExclusion, nil)
		childTree, err := b.buildTreeForRewrite(ctx, subSet, rewrite.Child, restDepth)
		if err != nil {
			return nil, err
		}

		if childTree != nil {
			// any compound operation under an InvertResult also has no meaningful subject
			// ex: view = !(A OR B)
			if _, ok := rewrite.Child.(*ast.SubjectSetRewrite); ok {
				childTree.Subject = nil
			}
			invertNode.Children = []*relationtuple.Tree{childTree}
		}
		return invertNode, nil
	}
	return nil, nil
}

// buildTreeTupleToSubjectSet finds subject sets under subSet#child.Relation, and for each,
// recurses into that subject set's child.ComputedSubjectSetRelation.
func (b *treeBuilder) buildTreeTupleToSubjectSet(ctx context.Context, subSet *relationtuple.SubjectSet, rewrite *ast.TupleToSubjectSet, resume *relationtuple.ExpandCursor, restDepth int) (*relationtuple.Tree, error) {
	treeSubject := &relationtuple.SubjectSet{
		Namespace: subSet.Namespace,
		Object:    subSet.Object,
		Relation:  rewrite.Relation,
	}
	tree := newNode(ketoapi.TreeNodeUnion, treeSubject)
	ttuResume := func(pageToken keysetpagination.PageToken) *relationtuple.ExpandCursor {
		return &relationtuple.ExpandCursor{Kind: relationtuple.ExpandCursorKindTTU, SubjectSet: treeSubject, TraverseRelation: &rewrite.ComputedSubjectSetRelation, NextPageToken: pageToken}
	}

	if b.isMaxNodeLimit() {
		tree.Truncation = tupleLimitTruncation(ttuResume(keysetpagination.PageToken{}))
		return tree, nil
	}
	if restDepth <= 1 {
		tree.Truncation = depthLimitTruncation(ttuResume(keysetpagination.PageToken{}))
		return tree, nil
	}

	for nextPage, _ := keysetpagination.NewPaginator(); !nextPage.IsLast(); {
		var tuples []*relationtuple.RelationTuple
		var err error
		tuples, nextPage, err = b.d.RelationTupleManager().GetRelationTuples(ctx, &relationtuple.RelationQuery{
			Namespace: &subSet.Namespace,
			Object:    &subSet.Object,
			Relation:  &rewrite.Relation,
		},
			append(nextPage.ToOptions(), paginatorOptionsFromResume(resume)...)...,
		)
		if err != nil {
			return nil, err
		} else if len(tuples) == 0 {
			break
		}

		var toWorkOn []*relationtuple.RelationTuple
		for _, tuple := range tuples {
			_, ok := tuple.Subject.(*relationtuple.SubjectSet)
			if !ok {
				continue
			}

			b.nodeCount++
			toWorkOn = append(toWorkOn, tuple)
		}

		if b.isMaxNodeLimit() && !nextPage.IsLast() {
			tree.Truncation = tupleLimitTruncation(ttuResume(nextPage.PageToken()))
		}

		for _, tuple := range toWorkOn {
			tupleSubSet := tuple.Subject.(*relationtuple.SubjectSet)

			childTree, err := b.buildTreeRecursive(ctx, &relationtuple.SubjectSet{
				Namespace: tupleSubSet.Namespace,
				Object:    tupleSubSet.Object,
				Relation:  rewrite.ComputedSubjectSetRelation,
			}, restDepth-1)
			if err != nil {
				return nil, err
			}
			if childTree != nil {
				tree.Children = append(tree.Children, childTree)
			}
		}

		if tree.Truncation != nil {
			break
		}
	}
	return tree, nil
}

func tupleLimitTruncation(resume *relationtuple.ExpandCursor) *relationtuple.Truncation {
	return &relationtuple.Truncation{Reason: relationtuple.TruncationReasonTupleLimit, Cursor: resume}
}

func depthLimitTruncation(resume *relationtuple.ExpandCursor) *relationtuple.Truncation {
	return &relationtuple.Truncation{Reason: relationtuple.TruncationReasonDepthLimit, Cursor: resume}
}

func paginatorOptionsFromResume(resume *relationtuple.ExpandCursor) []keysetpagination.Option {
	if resume != nil {
		return []keysetpagination.Option{keysetpagination.WithToken(resume.NextPageToken)}
	}
	return nil
}

func newNode(typ ketoapi.TreeNodeType, subSet relationtuple.Subject) *relationtuple.Tree {
	return &relationtuple.Tree{
		Type:    typ,
		Subject: subSet,
	}
}

func (b *treeBuilder) isMaxNodeLimit() bool {
	return b.nodeCount >= b.maxNodes
}

// rewriteCostOrder returns a sort key for rewrite children so that cheaper,
// direct lookups are evaluated before more expensive traversals.
func rewriteCostOrder(c ast.Child) int {
	switch c.(type) {
	case *ast.ComputedSubjectSet:
		return 0
	case *ast.TupleToSubjectSet:
		return 1
	case *ast.SubjectSetRewrite:
		return 2
	case *ast.InvertResult:
		return 3
	default:
		return 4
	}
}
