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

package server

import (
	"net/url"
	"runtime"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/ory/keto/role"
	"github.com/ory/ladon"
	"github.com/ory/ladon/manager/memory"
	"github.com/ory/ladon/manager/sql"
	"github.com/ory/sqlcon"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func connectToSql(url string, dbt string) (*sqlx.DB, error) {
	db, err := sqlx.Open(dbt, url)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	maxConns := maxParallelism() * 2
	maxConnLifetime := time.Duration(0)
	maxIdleConns := maxParallelism()
	db.SetMaxOpenConns(maxConns)
	db.SetMaxIdleConns(maxIdleConns)
	db.SetConnMaxLifetime(maxConnLifetime)
	return db, nil
}

func maxParallelism() int {
	maxProcs := runtime.GOMAXPROCS(0)
	numCPU := runtime.NumCPU()
	if maxProcs < numCPU {
		return maxProcs
	}
	return numCPU
}

type managers struct {
	roleManager   role.Manager
	policyManager ladon.Manager
}

func newManagers(db string, logger logrus.FieldLogger) (*managers, error) {
	if db == "memory" {
		return &managers{
			roleManager:   role.NewMemoryManager(),
			policyManager: memory.NewMemoryManager(),
		}, nil
	} else if db == "" {
		return nil, errors.New("No database URL provided")
	}

	u, err := url.Parse(db)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	switch u.Scheme {
	case "postgres":
	case "mysql":
		sdb, err := sqlcon.NewSQLConnection(db, logger)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		return &managers{
			roleManager:   role.NewSQLManager(sdb.GetDatabase()),
			policyManager: sql.NewSQLManager(sdb.GetDatabase(), nil),
		}, nil
	}

	return nil, errors.Errorf("The provided database URL %s can not be handled", db)
}
