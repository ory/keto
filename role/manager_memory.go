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
	"sync"

	"github.com/ory/herodot"
	"github.com/ory/pagination"
	"github.com/pborman/uuid"
	"github.com/pkg/errors"
)

func NewMemoryManager() *MemoryManager {
	return &MemoryManager{
		Roles: map[string]Role{},
	}
}

type MemoryManager struct {
	Roles map[string]Role
	sync.RWMutex
}

func (m *MemoryManager) CreateRole(r *Role) error {
	if r.ID == "" {
		r.ID = uuid.New()
	}
	if m.Roles == nil {
		m.Roles = map[string]Role{}
	}

	m.Roles[r.ID] = *r
	return nil
}

func (m *MemoryManager) GetRole(id string) (*Role, error) {
	if r, ok := m.Roles[id]; !ok {
		return nil, errors.WithStack(&herodot.ErrorNotFound)
	} else {
		return &r, nil
	}
}

func (m *MemoryManager) DeleteRole(id string) error {
	delete(m.Roles, id)
	return nil
}

func (m *MemoryManager) AddRoleMembers(role string, subjects []string) error {
	r, err := m.GetRole(role)
	if err != nil {
		return err
	}
	r.Members = append(r.Members, subjects...)
	return m.CreateRole(r)
}

func (m *MemoryManager) RemoveRoleMembers(role string, subjects []string) error {
	r, err := m.GetRole(role)
	if err != nil {
		return err
	}

	var subs []string
	for _, s := range r.Members {
		var remove bool
		for _, f := range subjects {
			if f == s {
				remove = true
				break
			}
		}
		if !remove {
			subs = append(subs, s)
		}
	}

	r.Members = subs
	return m.CreateRole(r)
}

func (m *MemoryManager) FindRolesByMember(member string, limit, offset int) ([]Role, error) {
	if m.Roles == nil {
		m.Roles = map[string]Role{}
	}

	res := make([]Role, 0)
	for _, r := range m.Roles {
		for _, s := range r.Members {
			if s == member {
				res = append(res, r)
				break
			}
		}
	}

	start, end := pagination.Index(limit, offset, len(res))
	return res[start:end], nil
}

func (m *MemoryManager) ListRoles(limit, offset int) ([]Role, error) {
	if m.Roles == nil {
		m.Roles = map[string]Role{}
	}

	i := 0
	res := make([]Role, len(m.Roles))
	for _, r := range m.Roles {
		res[i] = r
		i++
	}

	start, end := pagination.Index(limit, offset, len(res))
	return res[start:end], nil
}

func (m *MemoryManager) UpdateRole(role Role) error {
	if err := m.DeleteRole(role.ID); err != nil {
		return err
	}

	if err := m.CreateRole(&role); err != nil {
		return err
	}

	return nil
}
