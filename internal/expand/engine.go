// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package expand

import (
	"context"

	"github.com/ory/x/otelx"
	keysetpagination "github.com/ory/x/pagination/keysetpagination_v2"
	"go.opentelemetry.io/otel/trace"

	"github.com/ory/keto/internal/driver/config"
	"github.com/ory/keto/internal/relationtuple"
	"github.com/ory/keto/internal/x"
	"github.com/ory/keto/internal/x/graph"
	"github.com/ory/keto/ketoapi"
	"github.com/ory/keto/x/events"
)

type (
	EngineDependencies interface {
		relationtuple.ManagerProvider
		config.Provider
		x.LoggerProvider
		x.TracingProvider
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

func (e *Engine) BuildTree(ctx context.Context, subject relationtuple.Subject, restDepth int) (t *relationtuple.Tree, err error) {
	ctx, span := e.d.Tracer(ctx).Tracer().Start(ctx, "Engine.BuildTree")
	defer otelx.End(span, &err)

	t, err = e.buildTreeRecursive(ctx, subject, restDepth)
	if err != nil {
		trace.SpanFromContext(ctx).AddEvent(events.NewPermissionsExpanded(ctx))
	}
	return
}

func (e *Engine) buildTreeRecursive(ctx context.Context, subject relationtuple.Subject, restDepth int) (*relationtuple.Tree, error) {
	// global max-depth takes precedence when it is the lesser or if the request max-depth is less than or equal to 0
	if globalMaxDepth := e.d.Config(ctx).MaxReadDepth(); restDepth <= 0 || globalMaxDepth < restDepth {
		restDepth = globalMaxDepth
	}

	subSet, isSubjectSet := subject.(*relationtuple.SubjectSet)
	if !isSubjectSet {
		// is SubjectID
		return &relationtuple.Tree{
			Type:    ketoapi.TreeNodeLeaf,
			Subject: subject,
		}, nil
	}

	ctx, wasAlreadyVisited := graph.CheckAndAddVisited(ctx, subject)
	if wasAlreadyVisited {
		return nil, nil
	}

	subTree := &relationtuple.Tree{
		Type:    ketoapi.TreeNodeUnion,
		Subject: subject,
	}

	for nextPage := keysetpagination.NewPaginator(); !nextPage.IsLast(); {
		var rels []*relationtuple.RelationTuple
		var err error
		rels, nextPage, err = e.d.RelationTupleManager().GetRelationTuples(ctx, &relationtuple.RelationQuery{
			Relation:  &subSet.Relation,
			Object:    &subSet.Object,
			Namespace: &subSet.Namespace,
		},
			nextPage.ToOptions()...,
		)
		if err != nil {
			return nil, err
		} else if len(rels) == 0 {
			return nil, nil
		}

		if restDepth <= 1 {
			subTree.Type = ketoapi.TreeNodeLeaf
			return subTree, nil
		}

		children := make([]*relationtuple.Tree, len(rels))
		for ri, r := range rels {
			child, err := e.buildTreeRecursive(ctx, r.Subject, restDepth-1)
			if err != nil {
				return nil, err
			}
			if child == nil {
				child = &relationtuple.Tree{
					Type:    ketoapi.TreeNodeLeaf,
					Subject: r.Subject,
				}
			}
			children[ri] = child
		}
		subTree.Children = append(subTree.Children, children...)
	}

	return subTree, nil
}
