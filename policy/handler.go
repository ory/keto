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

package policy

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/ory/herodot"
	"github.com/ory/ladon"
	"github.com/ory/pagination"
	"github.com/pborman/uuid"
	"github.com/pkg/errors"
)

const (
	handlerBasePath = "/policies"
)

type Handler struct {
	Manager ladon.Manager
	H       herodot.Writer
}

func NewHandler(manager ladon.Manager, writer herodot.Writer) *Handler {
	return &Handler{
		H:       writer,
		Manager: manager,
	}
}

func (h *Handler) SetRoutes(r *httprouter.Router) {
	r.POST(handlerBasePath, h.Create)
	r.GET(handlerBasePath, h.List)
	r.GET(handlerBasePath+"/:id", h.Get)
	r.PUT(handlerBasePath+"/:id", h.Update)
	r.DELETE(handlerBasePath+"/:id", h.Delete)
}

// swagger:route GET /policies policy listPolicies
//
// List Access Control Policies
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
//       200: policyList
//       401: genericError
//       403: genericError
//       500: genericError
func (h *Handler) List(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	limit, offset := pagination.Parse(r, 500, 0, 1000)
	policies, err := h.Manager.GetAll(int64(limit), int64(offset))
	if err != nil {
		h.H.WriteError(w, r, errors.WithStack(err))
		return
	}
	h.H.Write(w, r, policies)
}

// swagger:route POST /policies policy createPolicy
//
// Create an Access Control Policy
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
//       201: policy
//       401: genericError
//       403: genericError
//       500: genericError
func (h *Handler) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var p = ladon.DefaultPolicy{
		Conditions: ladon.Conditions{},
	}

	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		h.H.WriteError(w, r, errors.WithStack(err))
		return
	}

	if p.ID == "" {
		p.ID = uuid.New()
	}

	if err := h.Manager.Create(&p); err != nil {
		h.H.WriteError(w, r, errors.WithStack(err))
		return
	}
	h.H.WriteCreated(w, r, "/policies/"+p.ID, &p)
}

// swagger:route GET /policies/{id} policy getPolicy
//
// Get an Access Control Policy
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
//       200: policy
//       401: genericError
//       403: genericError
//       500: genericError
func (h *Handler) Get(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	policy, err := h.Manager.Get(ps.ByName("id"))
	if err != nil {
		if err.Error() == "Not found" {
			h.H.WriteError(w, r, errors.WithStack(&herodot.ErrorNotFound))
			return
		}
		h.H.WriteError(w, r, errors.WithStack(err))
		return
	}
	h.H.Write(w, r, policy)
}

// swagger:route DELETE /policies/{id} policy deletePolicy
//
// Delete an Access Control Policy
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
func (h *Handler) Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	if err := h.Manager.Delete(id); err != nil {
		h.H.WriteError(w, r, errors.WithStack(err))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// swagger:route PUT /policies/{id} policy updatePolicy
//
// Update an Access Control Policy
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
//       200: policy
//       401: genericError
//       403: genericError
//       500: genericError
func (h *Handler) Update(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var id = ps.ByName("id")
	var p = ladon.DefaultPolicy{Conditions: ladon.Conditions{}}

	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		h.H.WriteError(w, r, errors.WithStack(err))
		return
	}

	if p.ID != id {
		h.H.WriteErrorCode(w, r, http.StatusBadRequest, errors.New("Payload ID does not match ID from URL"))
		return
	}

	if err := h.Manager.Update(&p); err != nil {
		h.H.WriteError(w, r, errors.WithStack(err))
		return
	}

	h.H.Write(w, r, p)
}
