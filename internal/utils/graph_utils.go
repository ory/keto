package utils

import (
	"context"
)

type (
	EngineUtils interface {
		CheckVisited(ctx context.Context, current string) (context.Context, bool)
	}
	EngineUtilsProvider struct {
		EngineUtils
	}
	contextKey string
)

const visitedMapKey = contextKey("visitedMap")

func (*EngineUtilsProvider) CheckVisited(ctx context.Context, current string) (context.Context, bool) {
	visitedMap, ok := ctx.Value(visitedMapKey).(map[string]bool)
	if !ok {
		// for the first time initialize the map
		visitedMap = make(map[string]bool)
		visitedMap[current] = true
		return context.WithValue(ctx, visitedMapKey, visitedMap), false
	}

	// check if current node was already visited
	if visitedMap[current] {
		return ctx, true
	}

	// set current entry to visited
	visitedMap[current] = true

	return context.WithValue(
		ctx,
		visitedMapKey,
		visitedMap,
	), false
}
