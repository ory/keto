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

// Package group offers capabilities for grouping subjects together, making policy management easier.
//
// This endpoint is **experimental**, use it at your own risk.

package roles

// A list of groups the member is belonging to
// swagger:response listGroupsResponse
type swaggerListGroupsResponse struct {
	// in: body
	// type: array
	Body []Role
}

// swagger:parameters listGroups
type swaggerListGroupsParameters struct {
	// The id of the member to look up.
	// in: query
	Member string `json:"member"`

	// The maximum amount of policies returned.
	// in: query
	Limit int `json:"limit"`

	// The offset from where to start looking.
	// in: query
	Offset int `json:"offset"`
}

// swagger:parameters createGroup
type swaggerCreateGroupParameters struct {
	// in: body
	Body Role
}

// swagger:parameters getGroup deleteGroup
type swaggerGetGroupParameters struct {
	// The id of the group to look up.
	// in: path
	ID string `json:"id"`
}

// swagger:parameters removeMembersFromGroup addMembersToGroup
type swaggerModifyMembersParameters struct {
	// The id of the group to modify.
	// in: path
	ID string `json:"id"`

	// in: body
	Body membersRequest
}
