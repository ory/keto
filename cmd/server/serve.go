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
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
	"github.com/meatballhat/negroni-logrus"
	"github.com/ory/fosite"
	"github.com/ory/go-convenience/corsx"
	"github.com/ory/graceful"
	"github.com/ory/herodot"
	"github.com/ory/keto/authentication"
	"github.com/ory/keto/policy"
	"github.com/ory/keto/role"
	"github.com/ory/keto/warden"
	"github.com/ory/ladon"
	"github.com/ory/metrics-middleware"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/urfave/negroni"
)

func RunServe(
	logger *logrus.Logger,
	buildVersion, buildHash string, buildTime string,
) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		router := httprouter.New()

		m, err := newManagers(viper.GetString("DATABASE_URL"), logger)
		if err != nil {
			logger.
				WithError(err).
				Fatal("Unable to initialise backends")
		}

		var strategy fosite.ScopeStrategy
		switch viper.GetString("AUTHENTICATOR_OAUTH2_INTROSPECTION_SCOPE_STRATEGY") {
		case "hierarchic":
			strategy = fosite.HierarchicScopeStrategy
			break
		case "exact":
			strategy = fosite.ExactScopeStrategy
			break
		case "wildcard":
			fallthrough
		default:
			strategy = fosite.WildcardScopeStrategy
		}

		authenticators := map[string]authentication.Authenticator{
			"subjects": authentication.NewPlaintextAuthentication(),
			"oauth2/access-tokens": authentication.NewOAuth2IntrospectionAuthentication(
				viper.GetString("AUTHENTICATOR_OAUTH2_INTROSPECTION_CLIENT_ID"),
				viper.GetString("AUTHENTICATOR_OAUTH2_INTROSPECTION_CLIENT_SECRET"),
				viper.GetString("AUTHENTICATOR_OAUTH2_INTROSPECTION_TOKEN_URL"),
				viper.GetString("AUTHENTICATOR_OAUTH2_INTROSPECTION_URL"),
				strings.Split(viper.GetString("AUTHENTICATOR_OAUTH2_INTROSPECTION_SCOPE"), ","),
				strategy,
			),
			"oauth2/clients": authentication.NewOAuth2ClientCredentialsAuthentication(
				viper.GetString("AUTHENTICATOR_OAUTH2_CLIENT_CREDENTIALS_TOKEN_URL"),
			),
		}

		decider := &ladon.Ladon{
			Manager:     m.policyManager,
			AuditLogger: &warden.AuditLoggerLogrus{Logger: logger},
			Matcher:     ladon.DefaultMatcher,
		}
		firewall := warden.NewWarden(decider, m.roleManager, logger)
		writer := herodot.NewJSONWriter(logger)
		roleHandler := role.NewHandler(m.roleManager, writer)
		policyHandler := policy.NewHandler(m.policyManager, writer)
		wardenHandler := warden.NewHandler(writer, firewall, authenticators)

		roleHandler.SetRoutes(router)
		policyHandler.SetRoutes(router)
		wardenHandler.SetRoutes(router)

		n := negroni.New()
		n.Use(negronilogrus.NewMiddlewareFromLogger(logger, "keto"))
		n.UseHandler(router)
		corsHandler := cors.New(corsx.ParseOptions()).Handler(n)

		if ok, _ := cmd.Flags().GetBool("disable-telemetry"); !ok {
			m := metrics.NewMetricsManager(
				metrics.Hash("DATABASE_URL"),
				viper.GetString("DATABASE_URL") != "memory",
				"jk32cFATnj9GKbQdFL7fBB9qtKZdX9j7",
				[]string{
					"/policies",
					"/roles",
					"/warden/subjects/authorize",
					"/warden/oauth2/access-tokens/authorize",
					"/warden/oauth2/clients/authorize",
				},
				logger,
			)
			go m.RegisterSegment(buildVersion, buildHash, buildTime)
			go m.CommitMemoryStatistics()
			n.Use(m)
		}

		address := fmt.Sprintf("%s:%s", viper.GetString("HOST"), viper.GetString("PORT"))
		var srv = graceful.WithDefaults(&http.Server{
			Addr:    address,
			Handler: context.ClearHandler(corsHandler),
		})

		if err := graceful.Graceful(func() error {
			logger.Infof("Setting up http server on %s", address)
			return srv.ListenAndServe()
		}, srv.Shutdown); err != nil {
			logger.
				WithError(err).
				Fatal("Could not gracefully run server")
		}

	}
}
