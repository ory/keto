// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package validate

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/ory/herodot"
)

type Validator func(r *http.Request) (ok bool, reason string)

// All runs all validators and returns an error if any of them fail. It returns
// a ErrBadRequest with all failed validation messages.
func All(r *http.Request, validator ...Validator) error {
	var reasons []string
	for _, v := range validator {
		if ok, msg := v(r); !ok {
			reasons = append(reasons, msg)
		}
	}
	if len(reasons) > 0 {
		return herodot.ErrBadRequest.WithReason(strings.Join(reasons, "; "))
	}
	return nil
}

// NoExtraQueryParams returns a validator that checks if the request has any
// query parameters that are not in the except list.
func NoExtraQueryParams(except ...string) Validator {
	return func(req *http.Request) (ok bool, reason string) {
		allowed := make(map[string]struct{}, len(except))
		for _, e := range except {
			allowed[e] = struct{}{}
		}
		for key := range req.URL.Query() {
			if _, found := allowed[key]; !found {
				return false, fmt.Sprintf("query parameter key %q unknown", key)
			}
		}
		return true, ""
	}
}

// QueryParamsContainsOneOf returns a validator that checks if the request has
// at least one of the specified query parameters.
func QueryParamsContainsOneOf(keys ...string) Validator {
	return func(req *http.Request) (ok bool, reason string) {
		oneOfKeys := make(map[string]struct{}, len(keys))
		for _, k := range keys {
			oneOfKeys[k] = struct{}{}
		}
		for key := range req.URL.Query() {
			if _, found := oneOfKeys[key]; found {
				return true, ""
			}
		}
		return false, fmt.Sprintf("query parameters must specify at least one of the following: %s", strings.Join(keys, ", "))
	}
}

// HasEmptyBody returns a validator that checks if the request body is empty.
func HasEmptyBody() Validator {
	return func(r *http.Request) (ok bool, reason string) {
		_, err := r.Body.Read([]byte{})
		if err != io.EOF {
			return false, "body is not empty"
		}
		return true, ""
	}
}
