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

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/ory/herodot"
	"github.com/ory/pagination"
	"github.com/pkg/errors"
)

// swagger:model roleMembers
type membersRequest struct {
	Members []string `json:"members"`
}

func NewHandler(manager Manager, writer herodot.Writer) *Handler {
	return &Handler{
		H:       writer,
		Manager: manager,
	}
}

type Handler struct {
	Manager Manager
	H       herodot.Writer
}

const (
	handlerBasePath = "/roles"
)

func (h *Handler) SetRoutes(r *httprouter.Router) {
	r.POST(handlerBasePath, h.CreateRole)
	r.GET(handlerBasePath, h.ListRoles)
	r.GET(handlerBasePath+"/:id", h.GetRole)
	r.DELETE(handlerBasePath+"/:id", h.DeleteRole)
	r.POST(handlerBasePath+"/:id/members", h.AddRoleMembers)
	r.DELETE(handlerBasePath+"/:id/members", h.DeleteRoleMembers)
	r.PUT(handlerBasePath+"/:id", h.UpdateRole)
}

// swagger:route GET /roles role listRoles
//
// List all roles
//
// A Role represents a group of users that share the same role and thus permissions. A role could be an administrator, a moderator, a regular
// user or some other sort of role.
//
// This endpoint allows you to retrieve all roles that are stored in the system.
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
//       200: listRolesResponse
//       401: genericError
//       403: genericError
//       500: genericError
func (h *Handler) ListRoles(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	limit, offset := pagination.Parse(r, 100, 0, 500)
	if member := r.URL.Query().Get("member"); member != "" {
		h.FindGroupNames(w, r, member, limit, offset)
		return
	} else {
		h.listAllRoles(w, r, limit, offset)
		return
	}
}

func (h *Handler) listAllRoles(w http.ResponseWriter, r *http.Request, limit, offset int) {
	groups, err := h.Manager.ListRoles(limit, offset)
	if err != nil {
		h.H.WriteError(w, r, err)
		return
	}

	h.H.Write(w, r, groups)
}

func (h *Handler) FindGroupNames(w http.ResponseWriter, r *http.Request, member string, limit, offset int) {
	groups, err := h.Manager.FindRolesByMember(member, limit, offset)
	if err != nil {
		h.H.WriteError(w, r, err)
		return
	}

	h.H.Write(w, r, groups)
}

// swagger:route POST /roles role createRole
//
// Create a role
//
// A Role represents a group of users that share the same role and thus permissions. A role could be an administrator, a moderator, a regular
// user or some other sort of role.
//
// This endpoint allows you to create a new role. You may define members as well but you don't have to.
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
//       201: role
//       401: genericError
//       403: genericError
//       500: genericError
func (h *Handler) CreateRole(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var g Role

	if err := json.NewDecoder(r.Body).Decode(&g); err != nil {
		h.H.WriteError(w, r, errors.WithStack(err))
		return
	}

	if err := h.Manager.CreateRole(&g); err != nil {
		h.H.WriteError(w, r, err)
		return
	}

	h.H.WriteCreated(w, r, handlerBasePath+"/"+g.ID, &g)
}

// swagger:route GET /roles/{id} role getRole
//
// Get a role by its ID
//
// A Role represents a group of users that share the same role and thus permissions. A role could be an administrator, a moderator, a regular
// user or some other sort of role.
//
// This endpoint allows you to retrieve an existing role. You have to know the role's ID.
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
//       201: role
//       401: genericError
//       403: genericError
//       500: genericError
func (h *Handler) GetRole(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var id = ps.ByName("id")

	g, err := h.Manager.GetRole(id)
	if err != nil {
		h.H.WriteError(w, r, err)
		return
	}

	h.H.Write(w, r, g)
}

// swagger:route DELETE /roles/{id} role deleteRole
//
// Get a role by its ID
//
// A Role represents a group of users that share the same role and thus permissions. A role could be an administrator, a moderator, a regular
// user or some other sort of role.
//
// This endpoint allows you to delete an existing role. You have to know the role's ID.
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
//       204: emptyResponse
//       401: genericError
//       403: genericError
//       500: genericError
func (h *Handler) DeleteRole(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var id = ps.ByName("id")

	if err := h.Manager.DeleteRole(id); err != nil {
		h.H.WriteError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// swagger:route POST /roles/{id}/members role addMembersToRole
//
// Add members to a role
//
// A Role represents a group of users that share the same role and thus permissions. A role could be an administrator, a moderator, a regular
// user or some other sort of role.
//
// This endpoint allows you to add members (users, applications, ...) to a specific role. You have to know the role's ID.
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
//       204: emptyResponse
//       401: genericError
//       403: genericError
//       500: genericError
func (h *Handler) AddRoleMembers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var id = ps.ByName("id")

	var m membersRequest
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		h.H.WriteError(w, r, errors.WithStack(err))
		return
	}

	if err := h.Manager.AddRoleMembers(id, m.Members); err != nil {
		h.H.WriteError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// swagger:route DELETE /roles/{id}/members role removeMembersFromRole
//
// Remove members from a role
//
// A Role represents a group of users that share the same role and thus permissions. A role could be an administrator, a moderator, a regular
// user or some other sort of role.
//
// This endpoint allows you to remove members (users, applications, ...) from a specific role. You have to know the role's ID.
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
//       204: emptyResponse
//       401: genericError
//       403: genericError
//       500: genericError
func (h *Handler) DeleteRoleMembers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var id = ps.ByName("id")

	var m membersRequest
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		h.H.WriteError(w, r, errors.WithStack(err))
		return
	}

	if err := h.Manager.RemoveRoleMembers(id, m.Members); err != nil {
		h.H.WriteError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// swagger:route PUT /roles/{id} role setRole
//
// A Role represents a group of users that share the same role and thus permissions. A role could be an administrator, a moderator, a regular
// user or some other sort of role.
//
// This endpoint allows you to overwrite a role. You have to know the role's ID.
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
//       204: emptyResponse
//       401: genericError
//       403: genericError
//       500: genericError
func (h *Handler) UpdateRole(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var id = ps.ByName("id")

	var m membersRequest
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		h.H.WriteError(w, r, errors.WithStack(err))
		return
	}

	if err := h.Manager.UpdateRole(Role{
		ID:      id,
		Members: m.Members,
	}); err != nil {
		h.H.WriteError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
