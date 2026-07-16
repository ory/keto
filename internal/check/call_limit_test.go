// Copyright © 2026 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package check_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ory/keto/internal/check"
	"github.com/ory/keto/internal/check/trace"
	"github.com/ory/keto/internal/driver"
	"github.com/ory/keto/internal/namespace"
	"github.com/ory/keto/internal/testhelpers"
	"github.com/ory/keto/schema"
)

func newRegistryForOPL(t *testing.T, opl string, opts ...driver.TestRegistryOption) *driver.RegistryDefault {
	t.Helper()
	parsed, errs := schema.Parse(opl)
	require.Empty(t, errs)
	nss := make([]*namespace.Namespace, len(parsed))
	for i := range parsed {
		nss[i] = &parsed[i]
	}
	return driver.NewSqliteTestRegistry(t, append([]driver.TestRegistryOption{driver.WithNamespaces(nss)}, opts...)...)
}

// requireCyclicError asserts err is the call-limit backstop, not just any error
// (e.g. a context cancellation), by checking its reason.
func requireCyclicError(t *testing.T, err error) {
	t.Helper()
	var re interface{ Reason() string }
	require.ErrorAs(t, err, &re)
	require.Contains(t, re.Reason(), "cyclic permit reference")
}

func TestCallLimit(t *testing.T) {
	t.Parallel()

	type outcome int
	const (
		cyclicError outcome = iota
		member
		depthLimited
	)

	for _, tc := range []struct {
		name      string
		opl       string
		opts      []driver.TestRegistryOption
		tuples    []string
		check     string
		restDepth int
		want      outcome
	}{
		{
			name: "invert cycle aborts regardless of configured depth",
			opl: `
				class User implements Namespace {}
				class Resource implements Namespace {
					permits = {
						edit: (ctx: Context): boolean => !this.permits.view(ctx),
						view: (ctx: Context): boolean => !this.permits.edit(ctx),
					}
				}`,
			opts:      []driver.TestRegistryOption{driver.WithConfig("limit.max_read_depth", 65535)},
			check:     "Resource:doc#edit@User:alice",
			restDepth: 65535,
			want:      cyclicError,
		},
		{
			name: "branching cycle aborts decisively",
			opl: `
				class User implements Namespace {}
				class R implements Namespace {
					permits = {
						a: (ctx: Context): boolean => this.permits.b(ctx) && this.permits.c(ctx),
						b: (ctx: Context): boolean => this.permits.a(ctx) && this.permits.a(ctx),
						c: (ctx: Context): boolean => this.permits.a(ctx) && this.permits.a(ctx),
					}
				}`,
			check:     "R:x#a@User:u",
			restDepth: 5,
			want:      cyclicError,
		},
		{
			name: "self-loop is bounded by RestDepth",
			opl: `
				class User implements Namespace {}
				class Resource implements Namespace {
					permits = {
						edit: (ctx: Context): boolean => this.permits.edit(ctx),
					}
				}`,
			check:     "Resource:doc#edit@User:alice",
			restDepth: 5,
			want:      depthLimited,
		},
		{
			name: "legit role chain resolves",
			opl: `
				class User implements Namespace {}
				class Doc implements Namespace {
					related: {
						l0: User[]; l1: User[]; l2: User[]; l3: User[];
					}
					permits = {
						p0: (ctx: Context): boolean => this.permits.p1(ctx) || this.related.l0.includes(ctx.subject),
						p1: (ctx: Context): boolean => this.permits.p2(ctx) || this.related.l1.includes(ctx.subject),
						p2: (ctx: Context): boolean => this.permits.p3(ctx) || this.related.l2.includes(ctx.subject),
						p3: (ctx: Context): boolean => this.related.l3.includes(ctx.subject),
					}
				}`,
			tuples:    []string{"Doc:d#l3@User:alice"},
			check:     "Doc:d#p0@User:alice",
			restDepth: 5,
			want:      member,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			reg := newRegistryForOPL(t, tc.opl, tc.opts...)
			if len(tc.tuples) > 0 {
				testhelpers.MapAndInsertTuplesFromString(t, reg, tc.tuples)
			}
			rt := testhelpers.TupleFromString(t, tc.check)

			assertResult := func(t *testing.T, res check.Result) {
				switch tc.want {
				case cyclicError:
					requireCyclicError(t, res.Err)
					require.Equal(t, check.MembershipUnknown, res.Membership)
				case member:
					require.NoError(t, res.Err)
					require.Equal(t, check.IsMember, res.Membership)
				case depthLimited:
					require.NoError(t, res.Err)
					require.Equal(t, check.MembershipUnknown, res.Membership)
					require.Equal(t, check.LimitationMaxDepthExceeded, res.Limitation)
				}
			}

			// Both executors wire the middleware separately.
			t.Run("check engine", func(t *testing.T) {
				assertResult(t, check.NewEngine(reg).CheckRelationTuple(t.Context(), rt, tc.restDepth))
			})
			t.Run("trace engine", func(t *testing.T) {
				res, tree := trace.NewEngine(reg).CheckRelationTupleWithTrace(t.Context(), rt, tc.restDepth)
				require.NotNil(t, tree)
				assertResult(t, res)
			})
		})
	}
}
