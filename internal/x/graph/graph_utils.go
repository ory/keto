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

func (s *stringSet) contains(el fmt.Stringer) (found bool) {
	s.l.Lock()
	defer s.l.Unlock()
	_, found = s.m[el.String()]
	return
}

func (s *stringSet) add(el fmt.Stringer) {
	s.l.Lock()
	defer s.l.Unlock()
	s.m[el.String()] = struct{}{}
}

func stringSetFromContext(ctx context.Context) (set *stringSet, ok bool) {
	set, ok = ctx.Value(visitedMapKey).(*stringSet)
	return
}

func contextWithStringSet(ctx context.Context) (context.Context, *stringSet) {
	set := newStringSet()
	return context.WithValue(ctx, visitedMapKey, set), set
}

func CheckAndAddVisited(ctx context.Context, current relationtuple.Subject) (context.Context, bool) {
	set, ok := stringSetFromContext(ctx)
	if !ok {
		ctx, set := contextWithStringSet(ctx)
		set.add(current)
		return ctx, false
	}

	if set.contains(current) {
		return ctx, true
	}
	set.add(current)

	return ctx, false
}
