// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package check

import (
	"context"

	"github.com/ory/x/logrusx"
	"github.com/ory/x/otelx"
	"golang.org/x/sync/errgroup"

	"github.com/ory/keto/x/events"

	"github.com/ory/keto/internal/driver/config"
	"github.com/ory/keto/internal/persistence"
	"github.com/ory/keto/internal/relationtuple"
	"github.com/ory/keto/ketoapi"
)

type (
	EngineProvider interface {
		PermissionEngine() *Engine
	}

	Engine struct {
		dep EngineDependencies
	}

	EngineDependencies interface {
		relationtuple.ManagerProvider
		relationtuple.MapperProvider
		CheckerProvider
		persistence.Provider
		config.Provider
		logrusx.Provider
		otelx.Provider
	}

	// Type alias for shorter signatures.
	relationTuple = relationtuple.RelationTuple
)

// NewEngine creates an Engine that delegates checks to checker.
func NewEngine(d EngineDependencies) *Engine {
	return &Engine{dep: d}
}

// CheckIsMember checks if the relation tuple's subject has the relation on the
// object in the namespace either directly or indirectly and returns a boolean
// result.
func (e *Engine) CheckIsMember(ctx context.Context, r *relationTuple, restDepth int) (bool, error) {
	result := e.CheckRelationTuple(ctx, r, restDepth)
	if result.Err != nil {
		return false, result.Err
	}
	return result.Membership == IsMember, nil
}

// CheckRelationTuple checks if the relation tuple's subject has the relation on
// the object in the namespace either directly or indirectly and returns a check
// result.
func (e *Engine) CheckRelationTuple(ctx context.Context, r *relationTuple, restDepth int) (res Result) {
	ctx, span := e.dep.Tracer(ctx).Tracer().Start(ctx, "Engine.CheckRelationTuple")
	defer otelx.End(span, &res.Err)

	res = e.dep.Checker().CheckRelationTuple(ctx, r, restDepth)
	if res.Err == nil {
		span.AddEvent(events.NewPermissionsChecked(ctx))
	}
	return res
}

// BatchCheck makes parallelized check requests for tuples. The check results
// are returned as a slice where the result index matches the input tuple index.
func (e *Engine) BatchCheck(
	ctx context.Context,
	tuples []*ketoapi.RelationTuple,
	maxDepth int,
) ([]Result, error) {
	eg := &errgroup.Group{}
	eg.SetLimit(e.dep.Config(ctx).BatchCheckParallelizationLimit())

	mapper := e.dep.ReadOnlyMapper()
	results := make([]Result, len(tuples))
	for i, tuple := range tuples {
		eg.Go(func() error {
			internalTuple, err := mapper.FromTuple(ctx, tuple)
			if err != nil {
				results[i] = Result{
					Membership: MembershipUnknown,
					Err:        err,
				}
			} else {
				results[i] = e.CheckRelationTuple(ctx, internalTuple[0], maxDepth)
			}
			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		return nil, err
	}
	return results, nil
}
