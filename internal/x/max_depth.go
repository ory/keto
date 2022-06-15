package x

import (
	"net/url"
	"strconv"

	"github.com/ory/herodot"
)

func GetMaxDepthFromQuery(q url.Values) (int, error) {
	if !q.Has("max-depth") {
		return 0, nil
	}

	maxDepth, err := strconv.ParseInt(q.Get("max-depth"), 0, 0)
	if err != nil {
		return 0, herodot.ErrBadRequest.WithErrorf("unable to parse 'max-depth' query parameter to int: %s", err)
	}

	return int(maxDepth), err
}
