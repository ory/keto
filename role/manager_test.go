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
	"flag"
	"fmt"
	"log"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	. "github.com/ory/keto/role"
	"github.com/ory/sqlcon/dockertest"
)

var clientManagers = map[string]Manager{
	"memory": &MemoryManager{
		Roles: map[string]Role{},
	},
}

func TestMain(m *testing.M) {
	runner := dockertest.Register()

	flag.Parse()
	if !testing.Short() {
		dockertest.Parallel([]func(){
			connectToPostgres,
			connectToMySQL,
		})
	}

	runner.Exit(m.Run())
}

func connectToMySQL() {
	db, err := dockertest.ConnectToTestMySQL()
	if err != nil {
		panic(err)
	}

	s := NewSQLManager(db)
	if _, err := s.CreateSchemas(); err != nil {
		log.Fatalf("Could not create mysql schema: %v", err)
	}

	clientManagers["mysql"] = s
}

func connectToPostgres() {
	db, err := dockertest.ConnectToTestPostgreSQL()
	if err != nil {
		panic(err)
	}

	s := NewSQLManager(db)
	if _, err := s.CreateSchemas(); err != nil {
		log.Fatalf("Could not create postgres schema: %v", err)
	}

	clientManagers["postgres"] = s
}

func TestManagers(t *testing.T) {
	for k, m := range clientManagers {
		t.Run(fmt.Sprintf("case=%s", k), TestHelperManagers(m))
	}
}
