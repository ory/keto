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

	"github.com/ory/keto/role"
	"github.com/ory/ladon/manager/sql"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"fmt"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"github.com/ory/keto/legacy"
	"github.com/rubenv/sql-migrate"
)

func getMigrationSql(cmd *cobra.Command, args []string, logger *logrus.Logger) (string, *url.URL) {
	var db string

	if a, b := cmd.Flags().GetBool("read-from-env"); a && b == nil {
		db = viper.GetString("DATABASE_URL")
	} else {
		if len(args) == 0 {
			fmt.Print(cmd.UsageString())
			logger.Fatalf("Argument 1 is missing")
		}
		db = args[0]
	}

	u, err := url.Parse(db)
	if err != nil {
		logger.WithError(err).WithField("database_url", db).Fatal("Unable to parse DATABASE_URL, make sure it has the right format")
	}

	return db, u
}

func RunMigrateSQL(logger *logrus.Logger) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		db, dbu := getMigrationSql(cmd, args, logger)
		if dbu.Scheme != "postgres" && dbu.Scheme != "mysql" {
			logger.WithField("database_url", dbu.Scheme+"://*:*@"+dbu.Host+dbu.Path+"?"+dbu.RawQuery).Fatal("Migrations can only be run against PostgreSQL or MySQL databases")
		}

		managers, err := newManagers(db, logger)
		if err != nil {
			logger.WithError(err).WithField("database_url", dbu.Scheme+"://*:*@"+dbu.Host+dbu.Path+"?"+dbu.RawQuery).Fatal("Unable to parse DATABASE_URL, make sure it has the right format")
		}

		logger.Info("Applying SQL migrations...")
		if n, err := managers.roleManager.(*role.SQLManager).CreateSchemas(); err != nil {
			logger.WithError(err).WithField("migrations", n).WithField("table", "policies").Fatal("An error occurred while trying to apply SQL migrations")
		} else {
			logger.WithField("migrations", n).WithField("table", "role").Print("Successfully applied SQL migrations")
		}

		if n, err := managers.policyManager.(*sql.SQLManager).CreateSchemas("", "keto_policy_migration"); err != nil {
			logger.WithError(err).WithField("migrations", n).WithField("table", "policies").Fatal("An error occurred while trying to apply SQL migrations")
		} else {
			logger.WithField("migrations", n).WithField("table", "policies").Print("Successfully applied SQL migrations")
		}

		logger.Info("Done applying SQL migrations")
	}
}

func RunMigrateHydra(logger *logrus.Logger) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		dbUrl, u := getMigrationSql(cmd, args, logger)

		db, err := connectToSQL(dbUrl, logger)
		if err != nil {
			logger.WithError(err).WithField("database_url", u.Scheme+"://*:*@"+u.Host+u.Path+"?"+u.RawQuery).Fatal("Unable to parse DATABASE_URL, make sure it has the right format")
		}

		migrate.SetTable("keto_legacy_hydra_migration")
		n, err := migrate.Exec(db.GetDatabase().DB, db.GetDatabase().DriverName(), legacy.HydraLegacyMigrations[db.GetDatabase().DriverName()], migrate.Up)
		if err != nil {
			logger.WithError(err).WithField("migrations", n).Fatal("An error occurred while trying to apply SQL migrations")
		}
		logger.WithField("migrations", n).Print("Successfully applied SQL migrations")
		logger.Info("Done applying SQL migrations")
	}
}
