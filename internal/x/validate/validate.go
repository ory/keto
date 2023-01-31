package validate

import (
	"io"
	"net/http"

	"github.com/ory/herodot"
)

func NoExtraQueryParams(req *http.Request, except ...string) error {
	allowed := make(map[string]struct{}, len(except))
	for _, e := range except {
		allowed[e] = struct{}{}
	}
	for key := range req.URL.Query() {
		if _, found := allowed[key]; !found {
			return herodot.ErrBadRequest.WithReasonf("query parameter key %q unknown", key)
		}
	}
	return nil
}

func HasEmptyBody(r *http.Request) error {
	_, err := r.Body.Read([]byte{})
	if err != io.EOF {
		return herodot.ErrBadRequest.WithReason("body is not empty")
	}
	return nil
}
