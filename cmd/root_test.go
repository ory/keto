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

package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"net/http"

	"github.com/akutz/gotil"
	"github.com/pborman/uuid"
	"github.com/stretchr/testify/assert"
)

var port int

func init() {
	var osArgs = make([]string, len(os.Args))
	port = gotil.RandomTCPPort()
	os.Setenv("DATABASE_URL", "memory")
	os.Setenv("PORT", fmt.Sprintf("%d", port))
	copy(osArgs, os.Args)
}

func TestExecute(t *testing.T) {
	var path = filepath.Join(os.TempDir(), fmt.Sprintf("keto-%s.yml", uuid.New()))

	for _, c := range []struct {
		args      []string
		wait      func(t *testing.T) bool
		expectErr bool
	}{
		{
			args: []string{"serve", "--disable-telemetry"},
			wait: func(t *testing.T) bool {
				t.Logf("Trying to connect to port %d...", port)
				_, err := http.DefaultClient.Get(fmt.Sprintf("http://127.0.0.1:%d/", port))
				return err != nil
			},
		},
		{args: []string{"roles", "list", "--endpoint", fmt.Sprintf("http://127.0.0.1:%d", port)}},
		{args: []string{"roles", "create", "role-a", "--endpoint", fmt.Sprintf("http://127.0.0.1:%d", port)}},
		{args: []string{"roles", "get", "role-a", "--endpoint", fmt.Sprintf("http://127.0.0.1:%d", port)}},
		{args: []string{"roles", "members", "add", "role-a", "member-a", "member-b", "--endpoint", fmt.Sprintf("http://127.0.0.1:%d", port)}},
		{args: []string{"roles", "members", "remove", "role-a", "member-a", "--endpoint", fmt.Sprintf("http://127.0.0.1:%d", port)}},
		{args: []string{"roles", "find", "member-a", "--endpoint", fmt.Sprintf("http://127.0.0.1:%d", port)}},
		{args: []string{"roles", "delete", "role-a", "--endpoint", fmt.Sprintf("http://127.0.0.1:%d", port)}},
		{args: []string{"policies", "create", "-i", "foobar", "-s", "peter,max", "-r", "blog,users", "-a", "post,ban", "--allow", "--endpoint", fmt.Sprintf("http://127.0.0.1:%d", port)}},
		{args: []string{"policies", "actions", "add", "foobar", "update|create", "--endpoint", fmt.Sprintf("http://127.0.0.1:%d", port)}},
		{args: []string{"policies", "actions", "remove", "foobar", "update|create", "--endpoint", fmt.Sprintf("http://127.0.0.1:%d", port)}},
		{args: []string{"policies", "resources", "add", "foobar", "printer", "--endpoint", fmt.Sprintf("http://127.0.0.1:%d", port)}},
		{args: []string{"policies", "resources", "remove", "foobar", "printer", "--endpoint", fmt.Sprintf("http://127.0.0.1:%d", port)}},
		{args: []string{"policies", "subjects", "add", "foobar", "ken", "tracy", "--endpoint", fmt.Sprintf("http://127.0.0.1:%d", port)}},
		{args: []string{"policies", "subjects", "remove", "foobar", "ken", "tracy", "--endpoint", fmt.Sprintf("http://127.0.0.1:%d", port)}},
		{args: []string{"policies", "get", "foobar", "--endpoint", fmt.Sprintf("http://127.0.0.1:%d", port)}},
		{args: []string{"policies", "delete", "foobar", "--endpoint", fmt.Sprintf("http://127.0.0.1:%d", port)}},
		{args: []string{"warden", "authorize", "subject", "--subject", "foo", "--action", "bar", "--resource", "baz", "--endpoint", fmt.Sprintf("http://127.0.0.1:%d", port)}},
		{args: []string{"help", "migrate", "sql"}},
		{args: []string{"help", "migrate", "hydra"}},
		{args: []string{"version"}},
	} {
		c.args = append(c.args, []string{"--config", path}...)
		RootCmd.SetArgs(c.args)

		t.Run(fmt.Sprintf("command=%v", c.args), func(t *testing.T) {
			if c.wait != nil {
				go func() {
					assert.Nil(t, RootCmd.Execute())
				}()
			}

			if c.wait != nil {
				var count = 0
				for c.wait(t) {
					t.Logf("Port not open yet, retrying attempt #%d...", count)
					count++
					if count > 30 {
						t.FailNow()
					}
					time.Sleep(time.Second)
				}
			} else {
				err := RootCmd.Execute()
				if c.expectErr {
					assert.Error(t, err)
				} else {
					assert.NoError(t, err)
				}
			}
		})
	}
}
