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
	"sync"

	"github.com/ory/hydra/pkg"
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

func (m *MemoryManager) CreateRole(g *Role) error {
	if g.ID == "" {
		g.ID = uuid.New()
	}
	if m.Roles == nil {
		m.Roles = map[string]Role{}
	}

	m.Roles[g.ID] = *g
	return nil
}

func (m *MemoryManager) GetRole(id string) (*Role, error) {
	if g, ok := m.Roles[id]; !ok {
		return nil, errors.WithStack(pkg.ErrNotFound)
	} else {
		return &g, nil
	}
}

func (m *MemoryManager) DeleteRole(id string) error {
	delete(m.Roles, id)
	return nil
}

func (m *MemoryManager) AddRoleMembers(group string, subjects []string) error {
	g, err := m.GetRole(group)
	if err != nil {
		return err
	}
	g.Members = append(g.Members, subjects...)
	return m.CreateRole(g)
}

func (m *MemoryManager) RemoveRoleMembers(group string, subjects []string) error {
	g, err := m.GetRole(group)
	if err != nil {
		return err
	}

	var subs []string
	for _, s := range g.Members {
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

	g.Members = subs
	return m.CreateRole(g)
}

func (m *MemoryManager) FindRolesByMember(member string, limit, offset int) ([]Role, error) {
	if m.Roles == nil {
		m.Roles = map[string]Role{}
	}

	res := make([]Role, 0)
	for _, g := range m.Roles {
		for _, s := range g.Members {
			if s == member {
				res = append(res, g)
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
	for _, g := range m.Roles {
		res[i] = g
		i++
	}

	start, end := pagination.Index(limit, offset, len(res))
	return res[start:end], nil
}
