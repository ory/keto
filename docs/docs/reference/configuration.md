---
id: configuration
title: Configuration
---

<!-- THIS FILE IS BEING AUTO-GENERATED. DO NOT MODIFY IT AS ALL CHANGES WILL BE OVERWRITTEN.
OPEN AN ISSUE IF YOU WOULD LIKE TO MAKE ADJUSTMENTS HERE AND MAINTAINERS WILL HELP YOU LOCATE THE RIGHT
FILE -->

If file `$HOME/.keto.yaml` exists, it will be used as a configuration file which
supports all configuration settings listed below.

You can load the config file from another source using the
`-c path/to/config.yaml` or `--config path/to/config.yaml` flag:
`keto --config path/to/config.yaml`.

Config files can be formatted as JSON, YAML and TOML. Some configuration values
support reloading without server restart. All configuration values can be set
using environment variables, as documented below.

To find out more about edge cases like setting string array values through
environmental variables head to the
[Configuring ORY services](https://www.ory.sh/docs/ecosystem/configuring)
section.

```yaml
## ORY Kratos Configuration
#

## Data Source Name ##
#
# Sets the data source name. This configures the backend where ORY Keto persists data. If dsn is "memory", data will be written to memory and is lost when you restart this instance. ORY Hydra supports popular SQL databases. For more detailed configuration information go to: https://www.ory.sh/docs/hydra/dependencies-environment#sql
#
# Examples:
# - postgres://user:password@host:123/database
# - mysql://user:password@tcp(host:123)/database
# - memory
#
# Set this value using environment variables on
# - Linux/macOS:
#    $ export DSN=<value>
# - Windows Command Line (CMD):
#    > set DSN=<value>
#
dsn: postgres://user:password@host:123/database

## HTTP REST API ##
#
serve:
  ## Port ##
  #
  # The port to listen on.
  #
  # Default value: 4456
  #
  # Examples:
  # - 4456
  #
  # Set this value using environment variables on
  # - Linux/macOS:
  #    $ export SERVE_PORT=<value>
  # - Windows Command Line (CMD):
  #    > set SERVE_PORT=<value>
  #
  port: 4456

  ## Host ##
  #
  # The network interface to listen on.
  #
  # Examples:
  # - localhost
  # - 127.0.0.1
  #
  # Set this value using environment variables on
  # - Linux/macOS:
  #    $ export SERVE_HOST=<value>
  # - Windows Command Line (CMD):
  #    > set SERVE_HOST=<value>
  #
  host: 127.0.0.1

  ## Cross Origin Resource Sharing (CORS) ##
  #
  # Configure [Cross Origin Resource Sharing (CORS)](http://www.w3.org/TR/cors/) using the following options.
  #
  cors:
    ## Enable CORS ##
    #
    # If set to true, CORS will be enabled and preflight-requests (OPTION) will be answered.
    #
    # Default value: false
    #
    # Set this value using environment variables on
    # - Linux/macOS:
    #    $ export SERVE_CORS_ENABLED=<value>
    # - Windows Command Line (CMD):
    #    > set SERVE_CORS_ENABLED=<value>
    #
    enabled: false

    ## Allowed Origins ##
    #
    # A list of origins a cross-domain request can be executed from. If the special * value is present in the list, all origins will be allowed. An origin may contain a wildcard (*) to replace 0 or more characters (i.e.: http://*.domain.com). Usage of wildcards implies a small performance penality. Only one wildcard can be used per origin.
    #
    # Default value: *
    #
    # Examples:
    # - - https://example.com
    #   - https://*.example.com
    #   - https://*.foo.example.com
    #
    # Set this value using environment variables on
    # - Linux/macOS:
    #    $ export SERVE_CORS_ALLOWED_ORIGINS=<value>
    # - Windows Command Line (CMD):
    #    > set SERVE_CORS_ALLOWED_ORIGINS=<value>
    #
    allowed_origins:
      - '*'

    ## Allowed HTTP Methods ##
    #
    # A list of methods the client is allowed to use with cross-domain requests.
    #
    # Default value: GET,POST,PUT,PATCH,DELETE
    #
    # Set this value using environment variables on
    # - Linux/macOS:
    #    $ export SERVE_CORS_ALLOWED_METHODS=<value>
    # - Windows Command Line (CMD):
    #    > set SERVE_CORS_ALLOWED_METHODS=<value>
    #
    allowed_methods:
      - POST
      - HEAD

    ## Allowed Request HTTP Headers ##
    #
    # A list of non simple headers the client is allowed to use with cross-domain requests.
    #
    # Default value: Authorization,Content-Type
    #
    # Set this value using environment variables on
    # - Linux/macOS:
    #    $ export SERVE_CORS_ALLOWED_HEADERS=<value>
    # - Windows Command Line (CMD):
    #    > set SERVE_CORS_ALLOWED_HEADERS=<value>
    #
    allowed_headers:
      - dolor mollit ipsum

    ## Allowed Response HTTP Headers ##
    #
    # Indicates which headers are safe to expose to the API of a CORS API specification
    #
    # Default value: Content-Type
    #
    # Set this value using environment variables on
    # - Linux/macOS:
    #    $ export SERVE_CORS_EXPOSED_HEADERS=<value>
    # - Windows Command Line (CMD):
    #    > set SERVE_CORS_EXPOSED_HEADERS=<value>
    #
    exposed_headers:
      - cupidatat

    ## Allow HTTP Credentials ##
    #
    # Indicates whether the request can include user credentials like cookies, HTTP authentication or client side SSL certificates.
    #
    # Default value: false
    #
    # Set this value using environment variables on
    # - Linux/macOS:
    #    $ export SERVE_CORS_ALLOW_CREDENTIALS=<value>
    # - Windows Command Line (CMD):
    #    > set SERVE_CORS_ALLOW_CREDENTIALS=<value>
    #
    allow_credentials: true

    ## Maximum Age ##
    #
    # Indicates how long (in seconds) the results of a preflight request can be cached. The default is 0 which stands for no max age.
    #
    # Set this value using environment variables on
    # - Linux/macOS:
    #    $ export SERVE_CORS_MAX_AGE=<value>
    # - Windows Command Line (CMD):
    #    > set SERVE_CORS_MAX_AGE=<value>
    #
    max_age: -57745217

    ## Enable Debugging ##
    #
    # Set to true to debug server side CORS issues.
    #
    # Default value: false
    #
    # Set this value using environment variables on
    # - Linux/macOS:
    #    $ export SERVE_CORS_DEBUG=<value>
    # - Windows Command Line (CMD):
    #    > set SERVE_CORS_DEBUG=<value>
    #
    debug: false

  ## HTTPS ##
  #
  # Configure HTTP over TLS (HTTPS). All options can also be set using environment variables by replacing dots (`.`) with underscores (`_`) and uppercasing the key. For example, `some.prefix.tls.key.path` becomes `export SOME_PREFIX_TLS_KEY_PATH`. If all keys are left undefined, TLS will be disabled.
  #
  tls:
    ## Private Key (PEM) ##
    #
    key:
      ## path ##
      #
      # Set this value using environment variables on
      # - Linux/macOS:
      #    $ export SERVE_TLS_KEY_PATH=<value>
      # - Windows Command Line (CMD):
      #    > set SERVE_TLS_KEY_PATH=<value>
      #
      path: path/to/file.pem

      ## base64 ##
      #
      # Set this value using environment variables on
      # - Linux/macOS:
      #    $ export SERVE_TLS_KEY_BASE64=<value>
      # - Windows Command Line (CMD):
      #    > set SERVE_TLS_KEY_BASE64=<value>
      #
      base64: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tXG5NSUlEWlRDQ0FrMmdBd0lCQWdJRVY1eE90REFOQmdr...

    ## TLS Certificate (PEM) ##
    #
    cert:
      ## path ##
      #
      # Set this value using environment variables on
      # - Linux/macOS:
      #    $ export SERVE_TLS_CERT_PATH=<value>
      # - Windows Command Line (CMD):
      #    > set SERVE_TLS_CERT_PATH=<value>
      #
      path: path/to/file.pem

      ## base64 ##
      #
      # Set this value using environment variables on
      # - Linux/macOS:
      #    $ export SERVE_TLS_CERT_BASE64=<value>
      # - Windows Command Line (CMD):
      #    > set SERVE_TLS_CERT_BASE64=<value>
      #
      base64: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tXG5NSUlEWlRDQ0FrMmdBd0lCQWdJRVY1eE90REFOQmdr...

## Profiling ##
#
# Enables CPU or memory profiling if set. For more details on profiling Go programs read [Profiling Go Programs](https://blog.golang.org/profiling-go-programs).
#
# Set this value using environment variables on
# - Linux/macOS:
#    $ export PROFILING=<value>
# - Windows Command Line (CMD):
#    > set PROFILING=<value>
#
profiling: mem

## Log ##
#
# Configure logging using the following options. Logging will always be sent to stdout and stderr.
#
log:
  ## Level ##
  #
  # Debug enables stack traces on errors. Can also be set using environment variable LOG_LEVEL.
  #
  # Default value: info
  #
  # Set this value using environment variables on
  # - Linux/macOS:
  #    $ export LOG_LEVEL=<value>
  # - Windows Command Line (CMD):
  #    > set LOG_LEVEL=<value>
  #
  level: debug

  ## Format ##
  #
  # The log format can either be text or JSON.
  #
  # Default value: text
  #
  # Set this value using environment variables on
  # - Linux/macOS:
  #    $ export LOG_FORMAT=<value>
  # - Windows Command Line (CMD):
  #    > set LOG_FORMAT=<value>
  #
  format: json

## tracing ##
#
# ORY Hydra supports distributed tracing.
#
tracing:
  ## provider ##
  #
  # Set this to the tracing backend you wish to use. Currently supports jaeger. If omitted or empty, tracing will be disabled.
  #
  # Examples:
  # - jaeger
  #
  # Set this value using environment variables on
  # - Linux/macOS:
  #    $ export TRACING_PROVIDER=<value>
  # - Windows Command Line (CMD):
  #    > set TRACING_PROVIDER=<value>
  #
  provider: jaeger

  ## service_name ##
  #
  # Specifies the service name to use on the tracer.
  #
  # Examples:
  # - ORY Hydra
  #
  # Set this value using environment variables on
  # - Linux/macOS:
  #    $ export TRACING_SERVICE_NAME=<value>
  # - Windows Command Line (CMD):
  #    > set TRACING_SERVICE_NAME=<value>
  #
  service_name: ORY Hydra

  ## providers ##
  #
  providers:
    ## jaeger ##
    #
    # Configures the jaeger tracing backend.
    #
    jaeger:
      ## local_agent_address ##
      #
      # The address of the jaeger-agent where spans should be sent to.
      #
      # Examples:
      # - 127.0.0.1:6831
      #
      # Set this value using environment variables on
      # - Linux/macOS:
      #    $ export TRACING_PROVIDERS_JAEGER_LOCAL_AGENT_ADDRESS=<value>
      # - Windows Command Line (CMD):
      #    > set TRACING_PROVIDERS_JAEGER_LOCAL_AGENT_ADDRESS=<value>
      #
      local_agent_address: 127.0.0.1:6831

      ## propagation ##
      #
      # The tracing header format
      #
      # Examples:
      # - jaeger
      #
      # Set this value using environment variables on
      # - Linux/macOS:
      #    $ export TRACING_PROVIDERS_JAEGER_PROPAGATION=<value>
      # - Windows Command Line (CMD):
      #    > set TRACING_PROVIDERS_JAEGER_PROPAGATION=<value>
      #
      propagation: jaeger

      ## sampling ##
      #
      # Examples:
      # - type: const
      #   value: 1
      #   server_url: http://localhost:5778/sampling
      #
      sampling:
        ## type ##
        #
        # Set this value using environment variables on
        # - Linux/macOS:
        #    $ export TRACING_PROVIDERS_JAEGER_SAMPLING_TYPE=<value>
        # - Windows Command Line (CMD):
        #    > set TRACING_PROVIDERS_JAEGER_SAMPLING_TYPE=<value>
        #
        type: const

        ## value ##
        #
        # Set this value using environment variables on
        # - Linux/macOS:
        #    $ export TRACING_PROVIDERS_JAEGER_SAMPLING_VALUE=<value>
        # - Windows Command Line (CMD):
        #    > set TRACING_PROVIDERS_JAEGER_SAMPLING_VALUE=<value>
        #
        value: 1

        ## server_url ##
        #
        # Set this value using environment variables on
        # - Linux/macOS:
        #    $ export TRACING_PROVIDERS_JAEGER_SAMPLING_SERVER_URL=<value>
        # - Windows Command Line (CMD):
        #    > set TRACING_PROVIDERS_JAEGER_SAMPLING_SERVER_URL=<value>
        #
        server_url: http://localhost:5778/sampling
```
