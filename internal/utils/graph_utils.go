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
)

func (*EngineUtilsProvider) CheckVisited(ctx context.Context, current string) (context.Context, bool) {
	visitedMap, ok := ctx.Value("visitedMap").(map[string]bool)
	if !ok {
		// for the first time initialize the map
		visitedMap = make(map[string]bool)
		visitedMap[current] = true
		return context.WithValue(ctx, "visitedMap", visitedMap), false
	}

	// check if current node was already visited
	if visitedMap[current] {
		return ctx, true
	}

	// set current entry to visited
	visitedMap[current] = true

	return context.WithValue(
		ctx,
		"visitedMap",
		visitedMap,
	), false
}
