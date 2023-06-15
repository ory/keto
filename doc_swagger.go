// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package main

import "github.com/ory/herodot"

// JSON API Error Response
//
// The standard Ory JSON API error format.
//
// swagger:model errorGeneric
//
//lint:ignore U1000 Used to generate Swagger and OpenAPI definitions
type errorGeneric struct {
	// Contains error details
	//
	// required: true
	Error herodot.DefaultError `json:"error"`
}

// An empty response
//
// swagger:response emptyResponse
//
//lint:ignore U1000 Used to generate Swagger and OpenAPI definitions
type emptyResponse struct{}
