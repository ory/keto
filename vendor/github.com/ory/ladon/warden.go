/*
 * Copyright © 2016-2018 Aeneas Rekkas <aeneas+oss@aeneas.io>
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

package ladon

// Request is the warden's request object.
type Request struct {
	// Resource is the resource that access is requested to.
	Resource string `json:"resource"`

	// Action is the action that is requested on the resource.
	Action string `json:"action"`

	// Subejct is the subject that is requesting access.
	Subject string `json:"subject"`

	// Context is the request's environmental context.
	Context Context `json:"context"`
}

// Warden is responsible for deciding if subject s can perform action a on resource r with context c.
type Warden interface {
	// IsAllowed returns nil if subject s can perform action a on resource r with context c or an error otherwise.
	//  if err := guard.IsAllowed(&Request{Resource: "article/1234", Action: "update", Subject: "peter"}); err != nil {
	//    return errors.New("Not allowed")
	//  }
	IsAllowed(r *Request) error
}
