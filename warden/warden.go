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
 * @Copyright 	2015-2018 Aeneas Rekkas <aeneas+oss@aeneas.io>
 * @license 	Apache-2.0
 *
 */

// Package warden defines an API for validating access requests.
package warden

import (
	"context"
)

// AccessRequest is the warden's request object.
//
// swagger:model WardenSubjectAuthorizationRequest
type AccessRequest struct {
	// Resource is the resource that access is requested to.
	Resource string `json:"resource"`

	// Action is the action that is requested on the resource.
	Action string `json:"action"`

	// Subejct is the subject that is requesting access.
	Subject string `json:"subject"`

	// Context is the request's environmental context.
	Context map[string]interface{} `json:"context"`
}

// Firewall offers various validation strategies for access tokens.
type Firewall interface {
	// IsAllowed uses policies to return nil if the access request can be fulfilled or an error if not.
	//
	//  ctx, err := firewall.IsAllowed(context.Background(), &AccessRequest{
	//    Subject:  "alice",
	//    Resource: "matrix",
	//    Action:   "create",
	//    Context:  ladon.Context{},
	//  }, "photos", "files")
	//
	//  fmt.Sprintf("%s", ctx.Subject)
	IsAllowed(ctx context.Context, accessRequest *AccessRequest) error
}
