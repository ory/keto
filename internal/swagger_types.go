package internal

// Empty responses are sent when, for example, resources are deleted. The HTTP status code for empty responses is typically 201.
//
// swagger:response emptyResponse
// nolint:deadcode,unused
type emptyResponse struct{}
