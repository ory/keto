// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package internal

// Empty responses are sent when, for example, resources are deleted. The HTTP status code for empty responses is typically 204.
//
// swagger:response emptyResponse
//
//lint:ignore U1000 Used to generate Swagger and OpenAPI definitions
type emptyResponse struct{}
