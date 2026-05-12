// Copyright © 2026 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package check

import (
	"context"
)

// StepKind identifies the execution model type of a Step.
type StepKind string

const (
	StepIsAllowed   StepKind = "is_allowed"
	StepDirect      StepKind = "direct"
	StepDirectMulti StepKind = "direct_multi"
	StepExpand      StepKind = "expand_subject"
	StepComputed    StepKind = "computed"
	StepTraverse    StepKind = "traverse"
	StepRewrite     StepKind = "rewrite"
	StepInvert      StepKind = "invert"
)

// Step is a single unit of execution within the permission check engine.
type Step interface {
	// Kind returns the execution-model kind of this step.
	Kind() StepKind

	// Execute performs the step's work and returns the membership result.
	// Child steps must be dispatched through ex, not called directly.
	Execute(ctx context.Context, req CheckRequest, ex Executor) Result
}
