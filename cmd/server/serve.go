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
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/gobuffalo/packr"
	"github.com/jmoiron/sqlx"
	"github.com/julienschmidt/httprouter"
	"github.com/meatballhat/negroni-logrus"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/urfave/negroni"

	"github.com/ory/go-convenience/stringslice"
	"github.com/ory/graceful"
	"github.com/ory/herodot"
	"github.com/ory/keto/engine"
	"github.com/ory/keto/engine/ladon"
	_ "github.com/ory/keto/engine/ladon/rego"
	"github.com/ory/keto/storage"
	"github.com/ory/x/cmdx"
	"github.com/ory/x/corsx"
	"github.com/ory/x/dbal"
	"github.com/ory/x/flagx"
	"github.com/ory/x/healthx"
	"github.com/ory/x/metricsx"
	"github.com/ory/x/tlsx"
)

// RunServe runs the Keto API HTTP server
func RunServe(
	logger *logrus.Logger,
	buildVersion, buildHash string, buildTime string,
) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		box := packr.NewBox("../../engine/ladon/rego")

		compiler, err := engine.NewCompiler(box, logger)
		cmdx.Must(err, "Unable to initialize compiler: %s", err)

		writer := herodot.NewJSONWriter(logger)
		//writer.ErrorEnhancer = nil

		var s storage.Manager
		checks := map[string]healthx.ReadyChecker{}

		dbal.Connect(viper.GetString("DATABASE_URL"), logger,
			func() error {
				s = storage.NewMemoryManager()
				checks["storage"] = healthx.NoopReadyChecker
				return nil
			},
			func(db *sqlx.DB) error {
				ss := storage.NewSQLManager(db)
				checks["storage"] = db.Ping
				s = ss
				return nil
			},
		)

		sh := storage.NewHandler(s, writer)
		e := engine.NewEngine(compiler, writer)

		router := httprouter.New()
		ladon.NewEngine(s, sh, e, writer).Register(router)
		healthx.NewHandler(writer, buildVersion, checks).SetRoutes(router)

		n := negroni.New()
		n.Use(negronilogrus.NewMiddlewareFromLogger(logger, "keto"))

		if flagx.MustGetBool(cmd, "disable-telemetry") {
			logger.Println("Transmission of telemetry data is enabled, to learn more go to: https://www.ory.sh/docs/ecosystem/sqa")

			m := metricsx.NewMetricsManager(
				metricsx.Hash("DATABASE_URL"),
				viper.GetString("DATABASE_URL") != "memory",
				"jk32cFATnj9GKbQdFL7fBB9qtKZdX9j7",
				stringslice.Merge(
					healthx.RoutesToObserve(),
					ladon.RoutesToObserve(),
				),
				logger,
				"ory-keto",
				100,
				"",
			)
			go m.RegisterSegment(buildVersion, buildHash, buildTime)
			go m.CommitMemoryStatistics()
			n.Use(m)
		}

		n.UseHandler(router)
		c := corsx.Initialize(n, logger)

		addr := fmt.Sprintf("%s:%s", viper.GetString("HOST"), viper.GetString("PORT"))
		server := graceful.WithDefaults(&http.Server{
			Addr:    addr,
			Handler: c,
		})

		cert, err := tlsx.HTTPSCertificate()
		if errors.Cause(err) == tlsx.ErrNoCertificatesConfigured {
			server.TLSConfig = &tls.Config{Certificates: cert}
		} else if err != nil {
			cmdx.Must(err, "Unable to load HTTP TLS certificate(s): %s", err)
		}

		if err := graceful.Graceful(func() error {
			if cert != nil {
				logger.Printf("Listening on https://%s", addr)
				return server.ListenAndServeTLS("", "")
			}
			logger.Printf("Listening on http://%s", addr)
			return server.ListenAndServe()
		}, server.Shutdown); err != nil {
			logger.Fatalf("Unable to gracefully shutdown HTTP(s) server because %v", err)
			return
		}
	}
}
