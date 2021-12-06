package x

import (
	"fmt"
	"net/url"
	"strconv"
)

func GetMaxDepthFromQuery(q url.Values, required bool) (int, error) {
	if !q.Has("max-depth") {
		if required {
			return 0, fmt.Errorf("required query parameter 'max-depth' is missing")
		}
		return 0, nil
	}

	maxDepth, err := strconv.ParseInt(q.Get("max-depth"), 0, 0)
	if err != nil {
		return 0, fmt.Errorf("unable to parse 'max-depth' query parameter to int: %s", err)
	}

	return int(maxDepth), err
}
