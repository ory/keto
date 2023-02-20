// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package internal

// Empty responses are sent when, for example, resources are deleted. The HTTP status code for empty responses is typically 204.
//
// swagger:response emptyResponse
// nolint:deadcode,unused
type emptyResponse struct{}
