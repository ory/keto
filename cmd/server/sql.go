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
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"github.com/ory/keto/health"
	"github.com/ory/keto/role"
	"github.com/ory/ladon"
	"github.com/ory/ladon/manager/memory"
	"github.com/ory/ladon/manager/sql"
	"github.com/ory/sqlcon"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type managers struct {
	roleManager   role.Manager
	policyManager ladon.Manager
	readyCheckers map[string]health.ReadyChecker
}

func newManagers(db string, logger logrus.FieldLogger) (*managers, error) {
	if db == "memory" {
		return &managers{
			readyCheckers: map[string]health.ReadyChecker{
				"database": func() error {
					return nil
				},
			},
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
		fallthrough
	case "mysql":
		sdb, err := connectToSQL(db, logger)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		return &managers{
			readyCheckers: map[string]health.ReadyChecker{
				"database": func() error {
					return sdb.GetDatabase().Ping()
				},
			},
			roleManager:   role.NewSQLManager(sdb.GetDatabase()),
			policyManager: sql.NewSQLManager(sdb.GetDatabase(), nil),
		}, nil
	}

	return nil, errors.Errorf("The provided database URL %s can not be handled", db)
}

func retry(logger logrus.FieldLogger, maxWait time.Duration, failAfter time.Duration, f func() error) (err error) {
	var lastStart time.Time
	err = errors.New("Did not connect.")
	loopWait := time.Millisecond * 100
	retryStart := time.Now().UTC()
	for retryStart.Add(failAfter).After(time.Now().UTC()) {
		lastStart = time.Now().UTC()
		if err = f(); err == nil {
			return nil
		}

		if lastStart.Add(maxWait * 2).Before(time.Now().UTC()) {
			retryStart = time.Now().UTC()
		}

		logger.WithError(err).Infof("Retrying in %f seconds...", loopWait.Seconds())
		time.Sleep(loopWait)
		loopWait = loopWait * time.Duration(int64(2))
		if loopWait > maxWait {
			loopWait = maxWait
		}
	}
	return err
}

func connectToSQL(db string, logger logrus.FieldLogger) (sdb *sqlcon.SQLConnection, err error) {
	if err := retry(logger, time.Minute, time.Minute*15, func() error {
		var err error
		sdb, err = sqlcon.NewSQLConnection(db, logger)
		if err != nil {
			return errors.WithStack(err)
		}

		if err := sdb.GetDatabase().Ping(); err != nil {
			return errors.WithStack(err)
		}
		return nil
	}); err != nil {
		return nil, errors.WithStack(err)
	}

	return sdb, nil
}
