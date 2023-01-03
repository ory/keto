// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package graph

import (
	"context"
	"fmt"
	"sync"

	"github.com/ory/keto/internal/relationtuple"
)

type contextKey string

const visitedMapKey = contextKey("visitedMap")

type stringSet struct {
	m map[string]struct{}
	l sync.Mutex
}

func newStringSet() *stringSet {
	return &stringSet{m: make(map[string]struct{})}
}

func (s *stringSet) addNoDuplicate(el fmt.Stringer) bool {
	s.l.Lock()
	defer s.l.Unlock()

	if _, found := s.m[el.String()]; found {
		return true
	}
	s.m[el.String()] = struct{}{}
	return false
}

func InitVisited(ctx context.Context) context.Context {
	if _, ok := ctx.Value(visitedMapKey).(*stringSet); !ok {
		ctx = context.WithValue(ctx, visitedMapKey, newStringSet())
	}
	return ctx
}

func CheckAndAddVisited(ctx context.Context, current relationtuple.Subject) (context.Context, bool) {
	set, ok := ctx.Value(visitedMapKey).(*stringSet)
	if !ok {
		set = newStringSet()
		ctx = context.WithValue(ctx, visitedMapKey, set)
	}

	return ctx, set.addNoDuplicate(current.UniqueID())
}
