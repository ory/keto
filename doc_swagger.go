// Copyright © 2022 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package main

import "github.com/ory/herodot"

// JSON API Error Response
//
// The standard Ory JSON API error format.
//
// swagger:model errorGeneric
type errorGeneric struct {
	// Contains error details
	//
	// required: true
	Error herodot.DefaultError `json:"error"`
}

// An empty response
//
// swagger:response emptyResponse
// nolint:deadcode,unused
type emptyResponse struct{}
