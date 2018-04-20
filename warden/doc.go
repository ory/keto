/*
 * Copyright © 2015-2018 Aeneas Rekkas <aeneas+oss@aeneas.io>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * @author		Aeneas Rekkas <aeneas+oss@aeneas.io>
 * @copyright 	2015-2018 Aeneas Rekkas <aeneas+oss@aeneas.io>
 * @license 	Apache-2.0
 */

// Package warden implements endpoints capable of making access control decisions based on Access Control Policies
package warden

import (
	"github.com/ory/keto/authentication"
)

// swagger:parameters isSubjectAuthorized
type swaggerDoesWardenAllowAccessRequestParameters struct {
	// in: body
	Body AccessRequest
}

// swagger:parameters isOAuth2AccessTokenAuthorized
type swaggerDoesWardenAllowTokenAccessRqeuestParameters struct {
	// in: body
	Body swaggerWardenTokenAccessRequest
}

// swager:model authorizedBaseRequest
type swaggerWardenBaseRequest struct {

	// Resource is the resource that access is requested to.
	Resource string `json:"resource"`

	// Action is the action that is requested on the resource.
	Action string `json:"action"`

	// Context is the request's environmental context.
	Context map[string]interface{} `json:"context"`
}

// swagger:model WardenOAuth2AuthorizationRequest
type swaggerWardenTokenAccessRequest struct {
	authentication.AuthenticationOAuth2IntrospectionRequest
	swaggerWardenBaseRequest
}

// swagger:model wardenOAuth2AuthorizationResponse
type oauth2Authorization struct {
	authentication.OAuth2Session
}

// swagger:model wardenSubjectAuthorizationResponse
type subjectAuthorization struct {
	authentication.DefaultSession
}

// swagger:route POST /warden/oauth2/authorize warden isOAuth2AccessTokenAuthorized
//
// Check if an OAuth 2.0 access token is authorized to access a resource
//
// Checks if a token is valid and if the token subject is allowed to perform an action on a resource.
// This endpoint requires a token, a scope, a resource name, an action name and a context.
//
//
// If a token is expired/invalid, has not been granted the requested scope or the subject is not allowed to
// perform the action on the resource, this endpoint returns a 200 response with `{ "allowed": false }`.
//
//
// This endpoint passes all data from the upstream OAuth 2.0 token introspection endpoint. If you use ORY Hydra as an
// upstream OAuth 2.0 provider, data set through the `accessTokenExtra` field in the consent flow will be included in this
// response as well.
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Schemes: http, https
//
//     Responses:
//       200: wardenOAuth2AuthorizationResponse
//       401: genericError
//       403: genericError
//       500: genericError
func swaggerOAuth2Mock() {}

// swagger:route POST /warden/subjects/authorize warden isSubjectAuthorized
//
// Check if a subject is authorized to access a resource
//
// Checks if a subject (e.g. user ID, API key, ...) is allowed to perform a certain action on a resource.
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Schemes: http, https
//
//     Responses:
//       200: wardenSubjectAuthorizationResponse
//       401: genericError
//       403: genericError
//       500: genericError
func swaggerSubjectMock() {}
