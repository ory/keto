package keto

import "github.com/ory/keto/sdk/go/keto/swagger"

type WardenSDK interface {
	IsSubjectAuthorized(body swagger.WardenSubjectAuthorizationRequest) (*swagger.WardenSubjectAuthorizationResponse, *swagger.APIResponse, error)
	IsOAuth2AccessTokenAuthorized(body swagger.WardenOAuth2AuthorizationRequest) (*swagger.WardenOAuth2AuthorizationResponse, *swagger.APIResponse, error)
}
