// Copyright © 2022 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package check_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ory/keto/internal/check"
	"github.com/ory/keto/internal/check/checkgroup"
	"github.com/ory/keto/internal/driver/config"
	"github.com/ory/keto/internal/namespace"
	"github.com/ory/keto/internal/namespace/ast"
)

func wideNamespace(width int) *namespace.Namespace {
	wideNS := &namespace.Namespace{
		Name:      fmt.Sprintf("%d_wide", width),
		Relations: []ast.Relation{{Name: "editor"}},
	}
	viewerRelation := &ast.Relation{
		Name: "viewer",
		SubjectSetRewrite: &ast.SubjectSetRewrite{
			Operation: ast.OperatorOr,
			Children:  ast.Children{},
		},
	}
	for i := 0; i < width; i++ {
		relation := fmt.Sprintf("relation-%d", i)
		viewerRelation.SubjectSetRewrite.Children = append(
			viewerRelation.SubjectSetRewrite.Children,
			&ast.ComputedSubjectSet{Relation: relation},
		)
		wideNS.Relations = append(wideNS.Relations, ast.Relation{Name: relation})
	}
	viewerRelation.SubjectSetRewrite.Children = append(
		viewerRelation.SubjectSetRewrite.Children,
		&ast.ComputedSubjectSet{Relation: "editor"},
	)
	wideNS.Relations = append(wideNS.Relations, *viewerRelation)

	return wideNS
}

func BenchmarkCheckEngine(b *testing.B) {
	ctx := context.Background()
	var (
		depths   = []int{2, 4, 8, 16, 32}
		widths   = []int{10, 20, 40, 80, 100}
		maxDepth = depths[len(depths)-1]
	)

	var namespaces = []*namespace.Namespace{
		{Name: "deep",
			Relations: []ast.Relation{
				{Name: "owner"},
				{Name: "editor",
					SubjectSetRewrite: &ast.SubjectSetRewrite{
						Children: ast.Children{&ast.ComputedSubjectSet{
							Relation: "owner"}}}},
				{Name: "viewer",
					SubjectSetRewrite: &ast.SubjectSetRewrite{
						Children: ast.Children{
							&ast.ComputedSubjectSet{
								Relation: "editor"},
							&ast.TupleToSubjectSet{
								Relation:                   "parent",
								ComputedSubjectSetRelation: "viewer"}}}},
			}},
	}

	reg := newDepsProvider(b, namespaces)
	reg.Logger().Logger.SetLevel(logrus.InfoLevel)

	tuples := []string{
		"deep:deep_file#parent@deep:folder_1#...",
	}
	for i := 1; i < maxDepth; i++ {
		tuples = append(tuples, fmt.Sprintf("deep:folder_%d#parent@deep:folder_%d#...", i, i+1))
	}
	for _, d := range depths {
		tuples = append(tuples, fmt.Sprintf("deep:folder_%d#owner@user_%d", d, d))
	}
	for _, w := range widths {
		namespaces = append(namespaces, wideNamespace(w))
		tuples = append(tuples, fmt.Sprintf("%d-wide:wide_file#editor@user", w))
	}
	insertFixtures(b, reg.RelationTupleManager(), tuples)

	require.NoError(b, reg.Config(ctx).Set(config.KeyLimitMaxReadDepth, 100*maxDepth))
	e := check.NewEngine(reg)

	b.ResetTimer()
	b.Run("case=deep tree", func(b *testing.B) {
		for _, depth := range depths {
			b.Run(fmt.Sprintf("depth=%03d", depth), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					rt := tupleFromString(b, fmt.Sprintf("deep:deep_file#viewer@user_%d", depth))
					res := e.CheckRelationTuple(ctx, rt, 2*depth)
					assert.NoError(b, res.Err)
					if res.Membership != checkgroup.IsMember {
						b.Error("user should be able to view 'deep_file'")
					}
				}
			})
		}
	})

	b.Run("case=wide tree", func(b *testing.B) {
		for _, width := range widths {
			b.Run(fmt.Sprintf("width=%03d", width), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					rt := tupleFromString(b, fmt.Sprintf("%d-wide:wide_file#editor@user", width))
					res := e.CheckRelationTuple(ctx, rt, 2*width)
					assert.NoError(b, res.Err)
					if res.Membership != checkgroup.IsMember {
						b.Error("user should be able to view 'wide_file'")
					}
				}
			})
		}
	})
}
