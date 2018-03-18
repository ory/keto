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

package roles_test

import (
	"flag"
	"fmt"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	. "github.com/ory/hades/roles"
	"github.com/jmoiron/sqlx"
)

var clientManagers = map[string]Manager{
	"memory": &MemoryManager{
		Roles: map[string]Role{},
	},
}

func TestMain(m *testing.M) {
	flag.Parse()
	if !testing.Short() {
		connectToPG()
		connectToMySQL()
	}

	os.Exit(m.Run())
}

func connectToMySQL() {
	db, err := sqlx.Open("postgres", os.Getenv("TEST_POSTGRES_URL"))
	if err != nil {
		panic(err)
	}
	s := &SQLManager{DB: db}
	if _, err := s.CreateSchemas(); err != nil {
		log.Fatalf("Could not create postgres schema: %v", err)
	}

	clientManagers["postgres"] = s
}

func connectToPG() {
	db, err := sqlx.Open("mysql", os.Getenv("TEST_MYSQL_URL"))
	s := &SQLManager{DB: db}

	if _, err := s.CreateSchemas(); err != nil {
		log.Fatalf("Could not create postgres schema: %v", err)
	}

	clientManagers["mysql"] = s
}

func TestManagers(t *testing.T) {
	for k, m := range clientManagers {
		t.Run(fmt.Sprintf("case=%s", k), TestHelperManagers(m))
	}
}
