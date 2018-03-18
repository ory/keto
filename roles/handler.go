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

package roles

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/ory/herodot"
	"github.com/ory/pagination"
	"github.com/pkg/errors"
)

// swagger:model groupMembers
type membersRequest struct {
	Members []string `json:"members"`
}

type Handler struct {
	Manager Manager
	H       herodot.Writer
}

const (
	RolesPath = "/roles"
)

func (h *Handler) SetRoutes(r *httprouter.Router) {
	r.POST(RolesPath, h.CreateGroup)
	r.GET(RolesPath, h.ListGroupsHandler)
	r.GET(RolesPath+"/:id", h.GetGroup)
	r.DELETE(RolesPath+"/:id", h.DeleteGroup)
	r.POST(RolesPath+"/:id/members", h.AddGroupMembers)
	r.DELETE(RolesPath+"/:id/members", h.RemoveGroupMembers)
}

// swagger:route GET /warden/groups warden listGroups
//
// List groups
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
//       200: listGroupsResponse
//       401: genericError
//       403: genericError
//       500: genericError
func (h *Handler) ListGroupsHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	limit, offset := pagination.Parse(r, 100, 0, 500)
	if member := r.URL.Query().Get("member"); member != "" {
		h.FindGroupNames(w, r, member, limit, offset)
		return
	} else {
		h.ListGroups(w, r, limit, offset)
		return
	}
}

func (h *Handler) ListGroups(w http.ResponseWriter, r *http.Request, limit, offset int) {
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

// swagger:route POST /warden/groups warden createGroup
//
// Create a group
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
//       201: group
//       401: genericError
//       403: genericError
//       500: genericError
func (h *Handler) CreateGroup(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var g Role
	var ctx = r.Context()

	if err := json.NewDecoder(r.Body).Decode(&g); err != nil {
		h.H.WriteError(w, r, errors.WithStack(err))
		return
	}

	if err := h.Manager.CreateRole(&g); err != nil {
		h.H.WriteError(w, r, err)
		return
	}

	h.H.WriteCreated(w, r, RolesPath+"/"+g.ID, &g)
}

// swagger:route GET /warden/groups/{id} warden getGroup
//
// Get a group by id
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
//       201: group
//       401: genericError
//       403: genericError
//       500: genericError
func (h *Handler) GetGroup(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var ctx = r.Context()
	var id = ps.ByName("id")

	g, err := h.Manager.GetRole(id)
	if err != nil {
		h.H.WriteError(w, r, err)
		return
	}

	h.H.Write(w, r, g)
}

// swagger:route DELETE /warden/groups/{id} warden deleteGroup
//
// Delete a group by id
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
func (h *Handler) DeleteGroup(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var id = ps.ByName("id")

	if err := h.Manager.DeleteRole(id); err != nil {
		h.H.WriteError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// swagger:route POST /warden/groups/{id}/members warden addMembersToGroup
//
// Add members to a group
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
func (h *Handler) AddGroupMembers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

// swagger:route DELETE /warden/groups/{id}/members warden removeMembersFromGroup
//
// Remove members from a group
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
func (h *Handler) RemoveGroupMembers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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
