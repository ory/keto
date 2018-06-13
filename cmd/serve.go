// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"os"
	"strconv"

	"github.com/ory/keto/cmd/server"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts the server and serves the HTTP REST API",
	Long: `
This command exposes a variety of controls via environment variables. You can
set environments using "export KEY=VALUE" (Linux/macOS) or "set KEY=VALUE" (Windows). On Linux,
you can also set environments by pre-pending key value pairs: "KEY=VALUE KEY2=VALUE2 hydra"

All possible controls are listed below. The host process additionally exposes a few flags, which are listed below
the controls section.


CORE CONTROLS
=============

- DATABASE_URL: A URL to a persistent backend. ORY Keto supports various backends:
  - Memory: If DATABASE_URL is "memory", data will be written to memory and is lost when you restart this instance.
  	Example: DATABASE_URL=memory

  - Postgres: If DATABASE_URL is a DSN starting with postgres:// PostgreSQL will be used as storage backend.
	Example: DATABASE_URL=postgres://user:password@host:123/database

	If PostgreSQL is not serving TLS, append ?sslmode=disable to the url:
	DATABASE_URL=postgres://user:password@host:123/database?sslmode=disable

  - MySQL: If DATABASE_URL is a DSN starting with mysql:// MySQL will be used as storage backend.
	Example: DATABASE_URL=mysql://user:password@tcp(host:123)/database?parseTime=true

	Be aware that the ?parseTime=true parameter is mandatory, or timestamps will not work.

- PORT: The port ORY Keto should listen on.
	Defaults to PORT=4466

- HOST: The host interface ORY Keto should listen on. Leave empty to listen on all interfaces.
	Example: HOST=localhost

- LOG_LEVEL: Set the log level, supports "panic", "fatal", "error", "warn", "info" and "debug". Defaults to "info".
	Example: LOG_LEVEL=panic

- LOG_FORMAT: Leave empty for text based log format, or set to "json" for JSON formatting.
	Example: LOG_FORMAT="json"


AUTHENTICATORS
==============

- The OAuth 2.0 Token Introspection Authenticator is capable of resolving OAuth2 access tokens to a subject and a set
	of granted scopes using the OAuth 2.0 Introspection standard.

	- AUTHENTICATOR_OAUTH2_INTROSPECTION_CLIENT_ID: The client ID to be used when performing the OAuth 2.0 Introspection request.
		Example: AUTHENTICATOR_OAUTH2_INTROSPECTION_CLIENT_ID=my_client

	- AUTHENTICATOR_OAUTH2_INTROSPECTION_CLIENT_SECRET: The client secret to be used when performing the OAuth 2.0 Introspection request.
		Example: AUTHENTICATOR_OAUTH2_INTROSPECTION_CLIENT_SECRET=my_secret

	- AUTHENTICATOR_OAUTH2_INTROSPECTION_SCOPE: The scope(s) (comma separated) required to perform the introspection request. If no scopes are
		required, leave this value empty.
		Example: AUTHENTICATOR_OAUTH2_INTROSPECTION_SCOPE=scopeA,scopeB

	- AUTHENTICATOR_OAUTH2_INTROSPECTION_TOKEN_URL: The OAuth2 Token Endpoint URL of the server
		Example: AUTHENTICATOR_OAUTH2_INTROSPECTION_TOKEN_URL=https://my-server/oauth2/token

	- AUTHENTICATOR_OAUTH2_INTROSPECTION_INTROSPECTION_URL: The OAuth2 Introspection Endpoint URL of the server
		Example: AUTHENTICATOR_OAUTH2_INTROSPECTION_INTROSPECTION_URL=https://my-server/oauth2/introspect

- The OAuth 2.0 Client Credentials Authenticator is capable of authentication OAuth 2.0 clients using the client credentials
	grant.

	- AUTHENTICATOR_OAUTH2_CLIENT_CREDENTIALS_TOKEN_URL: The OAuth2 Token Endpoint URL of the server
		Example: AUTHENTICATOR_OAUTH2_CLIENT_CREDENTIALS_TOKEN_URL=https://my-server/oauth2/token

CORS CONTROLS
==============
- CORS_ALLOWED_ORIGINS: A list of origins (comma separated values) a cross-domain request can be executed from.
	If the special * value is present in the list, all origins will be allowed. An origin may contain a wildcard (*)
	to replace 0 or more characters (i.e.: http://*.domain.com). Usage of wildcards implies a small performance penality.
	Only one wildcard can be used per origin. The default value is *.
	Example: CORS_ALLOWED_ORIGINS=http://*.domain.com,http://*.domain2.com

- CORS_ALLOWED_METHODS: A list of methods  (comma separated values) the client is allowed to use with cross-domain
	requests. Default value is simple methods (GET and POST).
	Example: CORS_ALLOWED_METHODS=POST,GET,PUT

- CORS_ALLOWED_CREDENTIALS: Indicates whether the request can include user credentials like cookies, HTTP authentication
	or client side SSL certificates. The default is false.

- CORS_DEBUG: Debugging flag adds additional output to debug server side CORS issues.

- CORS_MAX_AGE: Indicates how long (in seconds) the results of a preflight request can be cached. The default is 0 which stands for no max age.

- CORS_ALLOWED_HEADERS: A list of non simple headers (comma separated values) the client is allowed to use with cross-domain requests.

- CORS_EXPOSED_HEADERS: Indicates which headers (comma separated values) are safe to expose to the API of a CORS API specification.


DEBUG CONTROLS
==============

- PROFILING: Set "PROFILING=cpu" to enable cpu profiling and "PROFILING=memory" to enable memory profiling.
	It is not possible to do both at the same time.
	Example: PROFILING=cpu
`,
	Run: server.RunServe(logger, Version, GitHash, BuildTime),
}

func init() {
	RootCmd.AddCommand(serveCmd)

	disableTelemetryEnv, _ := strconv.ParseBool(os.Getenv("DISABLE_TELEMETRY"))
	serveCmd.Flags().Bool("disable-telemetry", disableTelemetryEnv, "Disable anonymized telemetry reports - for more information please visit https://www.ory.sh/docs/guides/telemetry")
}
