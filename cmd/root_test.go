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

	"github.com/akutz/gotil"
	"github.com/pborman/uuid"
	"github.com/stretchr/testify/assert"
)

func TestExecute(t *testing.T) {
	var osArgs = make([]string, len(os.Args))
	var path = filepath.Join(os.TempDir(), fmt.Sprintf("keto-%s.yml", uuid.New()))
	port := gotil.RandomTCPPort()
	os.Setenv("DATABASE_URL", "memory")
	os.Setenv("PORT", fmt.Sprintf("%d", port))

	copy(osArgs, os.Args)

	for _, c := range []struct {
		args      []string
		wait      func() bool
		expectErr bool
	}{
		{
			args: []string{"serve"},
			wait: func() bool {
				time.Sleep(time.Second * 5)
				return !gotil.IsTCPPortAvailable(port)
			},
		},
		{args: []string{"roles", "list", "--url", fmt.Sprintf("http://127.0.0.1:%d", port)}},
		{args: []string{"roles", "create", "role-a", "--url", fmt.Sprintf("http://127.0.0.1:%d", port)}},
		{args: []string{"roles", "get", "role-a", "--url", fmt.Sprintf("http://127.0.0.1:%d", port)}},
		{args: []string{"roles", "members", "add", "role-a", "member-a", "member-b", "--url", fmt.Sprintf("http://127.0.0.1:%d", port)}},
		{args: []string{"roles", "members", "remove", "role-a", "member-a", "--url", fmt.Sprintf("http://127.0.0.1:%d", port)}},
		{args: []string{"roles", "find", "member-a", "--url", fmt.Sprintf("http://127.0.0.1:%d", port)}},
		{args: []string{"roles", "delete", "role-a", "--url", fmt.Sprintf("http://127.0.0.1:%d", port)}},
		{args: []string{"policies", "create", "-i", "foobar", "-s", "peter,max", "-r", "blog,users", "-a", "post,ban", "--allow", "--url", fmt.Sprintf("http://127.0.0.1:%d", port)}},
		{args: []string{"policies", "actions", "add", "foobar", "update|create", "--url", fmt.Sprintf("http://127.0.0.1:%d", port)}},
		{args: []string{"policies", "actions", "remove", "foobar", "update|create", "--url", fmt.Sprintf("http://127.0.0.1:%d", port)}},
		{args: []string{"policies", "resources", "add", "foobar", "printer", "--url", fmt.Sprintf("http://127.0.0.1:%d", port)}},
		{args: []string{"policies", "resources", "remove", "foobar", "printer", "--url", fmt.Sprintf("http://127.0.0.1:%d", port)}},
		{args: []string{"policies", "subjects", "add", "foobar", "ken", "tracy", "--url", fmt.Sprintf("http://127.0.0.1:%d", port)}},
		{args: []string{"policies", "subjects", "remove", "foobar", "ken", "tracy", "--url", fmt.Sprintf("http://127.0.0.1:%d", port)}},
		{args: []string{"policies", "get", "foobar", "--url", fmt.Sprintf("http://127.0.0.1:%d", port)}},
		{args: []string{"policies", "delete", "foobar", "--url", fmt.Sprintf("http://127.0.0.1:%d", port)}},
		{args: []string{"warden", "authorize", "subject", "--subject", "foo", "--action", "bar", "--resource", "baz", "--url", fmt.Sprintf("http://127.0.0.1:%d", port)}},
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
				for c.wait() {
					t.Logf("Port not open yet, retrying attempt #%d...", count)
					count++
					if count > 200 {
						t.FailNow()
					}
					time.Sleep(time.Second * 2)
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
