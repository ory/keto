package keto

import "github.com/ory/keto/sdk/go/keto/swagger"

type WardenSDK interface {
	IsSubjectAuthorized(body swagger.WardenSubjectAuthorizationRequest) (*swagger.WardenSubjectAuthorizationResponse, *swagger.APIResponse, error)
	IsOAuth2AccessTokenAuthorized(body swagger.WardenOAuth2AccessTokenAuthorizationRequest) (*swagger.WardenOAuth2AccessTokenAuthorizationResponse, *swagger.APIResponse, error)
	IsOAuth2ClientAuthorized(body swagger.WardenOAuth2ClientAuthorizationRequest) (*swagger.WardenOAuth2ClientAuthorizationResponse, *swagger.APIResponse, error)
}
