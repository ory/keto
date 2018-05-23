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

package role

// Role represents a group of users that share the same role. A role could be an administrator, a moderator, a regular
// user or some other sort of role.
//
// swagger:model role
type Role struct {
	// ID is the role's unique id.
	ID string `json:"id"`

	// Members is who belongs to the role.
	Members []string `json:"members"`
}

type Manager interface {
	CreateRole(*Role) error
	GetRole(id string) (*Role, error)
	DeleteRole(id string) error

	AddRoleMembers(role string, members []string) error
	RemoveRoleMembers(role string, members []string) error

	FindRolesByMember(member string, limit, offset int) ([]Role, error)
	ListRoles(limit, offset int) ([]Role, error)
	UpdateRole(role Role) error
}
