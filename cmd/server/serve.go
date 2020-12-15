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

//
//// RunServe runs the Keto API HTTP server
//func RunServe(
//	logger *logrusx.Logger,
//	version, commit string, date string,
//) func(cmd *cobra.Command, args []string) {
//	return func(cmd *cobra.Command, args []string) {
//		d := driver.NewDefaultDriver(
//			logger,
//			version, commit, date,
//		)
//
//		router := httprouter.New()
//		d.Registry().HealthHandler().SetRoutes(router, true)
//
//		n := negroni.New()
//		n.Use(reqlog.NewMiddlewareFromLogger(logger, "keto").ExcludePaths(healthx.ReadyCheckPath, healthx.AliveCheckPath))
//
//		if tracer := d.Registry().Tracer(); tracer.IsLoaded() {
//			n.Use(tracer)
//		}
//
//		metrics := metricsx.New(cmd, logger,
//			&metricsx.Options{
//				Service:       "ory-keto",
//				ClusterID:     metricsx.Hash(viper.GetString("DATABASE_URL")),
//				IsDevelopment: viper.GetString("DATABASE_URL") != "memory",
//				WriteKey:      "jk32cFATnj9GKbQdFL7fBB9qtKZdX9j7",
//				WhitelistedPaths: stringslice.Merge(
//					healthx.RoutesToObserve(),
//				),
//				BuildVersion: version,
//				BuildTime:    date,
//				BuildHash:    commit,
//				Config: &analytics.Config{
//					Endpoint:             "https://sqa.ory.sh",
//					GzipCompressionLevel: 6,
//					BatchMaxSize:         500 * 1000,
//					BatchSize:            250,
//					Interval:             time.Hour * 24,
//				},
//			},
//		)
//		n.Use(metrics)
//
//		n.UseHandler(router)
//		c := corsx.Initialize(n, logger, "serve")
//
//		server := graceful.WithDefaults(&http.Server{
//			Addr:    d.Configuration().RESTListenOn(),
//			Handler: c,
//		})
//
//		cert, err := tlsx.Certificate(
//			viper.GetString("serve.tls.cert.base64"),
//			viper.GetString("serve.tls.key.base64"),
//			viper.GetString("serve.tls.cert.path"),
//			viper.GetString("serve.tls.key.path"),
//		)
//		if errors.Cause(err) == tlsx.ErrNoCertificatesConfigured {
//			// do nothing
//		} else if err != nil {
//			cmdx.Must(err, "Unable to load HTTP TLS certificate(s): %s", err)
//		} else {
//			server.TLSConfig = &tls.Config{
//				Certificates: cert,
//				MinVersion:   tls.VersionTLS13,
//			}
//		}
//
//		if d.Registry().Tracer().IsLoaded() {
//			server.RegisterOnShutdown(d.Registry().Tracer().Close)
//		}
//
//		if err := graceful.Graceful(func() error {
//			if cert != nil {
//				logger.Printf("Listening on https://%s", d.Configuration().RESTListenOn())
//				return server.ListenAndServeTLS("", "")
//			}
//			logger.Printf("Listening on http://%s", d.Configuration().RESTListenOn())
//			return server.ListenAndServe()
//		}, server.Shutdown); err != nil {
//			logger.Fatalf("Unable to gracefully shutdown HTTP(s) server because %v", err)
//			return
//		}
//	}
//}
