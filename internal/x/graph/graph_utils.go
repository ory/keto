package graph

import (
	"context"

	"github.com/ory/keto/internal/relationtuple"
)

type contextKey string

const visitedMapKey = contextKey("visitedMap")

func CheckAndAddVisited(ctx context.Context, current relationtuple.Subject) (context.Context, bool) {
	visitedMap, ok := ctx.Value(visitedMapKey).(map[string]struct{})
	if !ok {
		// for the first time initialize the map
		visitedMap = make(map[string]struct{})
		visitedMap[current.String()] = struct{}{}
		return context.WithValue(ctx, visitedMapKey, visitedMap), false
	}

	// check if current node was already visited
	if _, ok := visitedMap[current.String()]; ok {
		return ctx, true
	}

	// set current entry to visited
	visitedMap[current.String()] = struct{}{}

	return context.WithValue(
		ctx,
		visitedMapKey,
		visitedMap,
	), false
}
