package graph

import (
	"context"

	"github.com/gofrs/uuid"
)

type contextKey string

const visitedMapKey = contextKey("visitedMap")

func CheckAndAddVisited(ctx context.Context, current uuid.UUID) (context.Context, bool) {
	visitedMap, ok := ctx.Value(visitedMapKey).(map[uuid.UUID]struct{})
	if !ok {
		// for the first time initialize the map
		visitedMap = make(map[uuid.UUID]struct{})
		visitedMap[current] = struct{}{}
		return context.WithValue(ctx, visitedMapKey, visitedMap), false
	}

	// check if current node was already visited
	if _, ok := visitedMap[current]; ok {
		return ctx, true
	}

	// set current entry to visited
	visitedMap[current] = struct{}{}

	return context.WithValue(
		ctx,
		visitedMapKey,
		visitedMap,
	), false
}
