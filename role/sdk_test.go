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

package role_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
	"github.com/ory/herodot"
	. "github.com/ory/keto/role"
	keto "github.com/ory/keto/sdk/go/keto/swagger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGroupSDK(t *testing.T) {
	clientManagers["memory"] = &MemoryManager{
		Roles: map[string]Role{},
	}

	handler := &Handler{
		Manager: &MemoryManager{
			Roles: map[string]Role{},
		},
		H: herodot.NewJSONWriter(nil),
	}

	router := httprouter.New()
	handler.SetRoutes(router)
	server := httptest.NewServer(router)

	client := keto.NewRoleApiWithBasePath(server.URL)

	t.Run("flows", func(*testing.T) {
		_, response, err := client.GetRole("4321")
		require.NoError(t, err)
		assert.Equal(t, http.StatusNotFound, response.StatusCode)

		firstGroup := keto.Role{Id: "1", Members: []string{"bar", "foo"}}
		result, response, err := client.CreateRole(firstGroup)
		require.NoError(t, err)
		assert.Equal(t, http.StatusCreated, response.StatusCode)
		assert.EqualValues(t, firstGroup, *result)

		secondGroup := keto.Role{Id: "2", Members: []string{"foo"}}
		result, response, err = client.CreateRole(secondGroup)
		require.NoError(t, err)
		assert.Equal(t, http.StatusCreated, response.StatusCode)
		assert.EqualValues(t, secondGroup, *result)

		result, response, err = client.GetRole("1")
		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, response.StatusCode)
		assert.EqualValues(t, firstGroup, *result)

		results, response, err := client.ListRoles("foo", 100, 0)
		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, response.StatusCode)
		assert.Len(t, results, 2)

		results, response, err = client.ListRoles("", 100, 0)
		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, response.StatusCode)
		assert.Len(t, results, 2)

		results, response, err = client.ListRoles("", 1, 0)
		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, response.StatusCode)
		assert.Len(t, results, 1)

		results, response, err = client.ListRoles("foo", 1, 0)
		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, response.StatusCode)
		assert.Len(t, results, 1)

		client.AddMembersToRole("1", keto.RoleMembers{Members: []string{"baz"}})

		results, response, err = client.ListRoles("baz", 100, 0)
		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, response.StatusCode)
		assert.Len(t, results, 1)

		response, err = client.RemoveMembersFromRole("1", keto.RoleMembers{Members: []string{"baz"}})
		require.NoError(t, err)
		assert.Equal(t, http.StatusNoContent, response.StatusCode)

		results, response, err = client.ListRoles("baz", 100, 0)
		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, response.StatusCode)
		assert.Len(t, results, 0)

		response, err = client.DeleteRole("1")
		require.NoError(t, err)
		assert.Equal(t, http.StatusNoContent, response.StatusCode)

		_, response, err = client.GetRole("4321")
		require.NoError(t, err)
		assert.Equal(t, http.StatusNotFound, response.StatusCode)
	})
}
