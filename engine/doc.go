// Package engine
package engine

// AuthorizationResult is the result of an access control decision. It contains the decision outcome.
// swagger:model authorizationResult
type AuthorizationResult struct {
	// Allowed is true if the request should be allowed and false otherwise.
	//
	// required: true
	Allowed bool `json:"allowed"`
}
