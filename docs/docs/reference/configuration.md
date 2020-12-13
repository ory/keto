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

This reference configuration documents all keys, also deprecated ones! It is a
reference for all possible configuration values.

If you are looking for an example configuration, it is better to try out the
quickstart.

To find out more about edge cases like setting string array values through
environmental variables head to the
[Configuring ORY services](https://www.ory.sh/docs/ecosystem/configuring)
section.

```yaml
## ORY Kratos Configuration
#

## HTTP REST API ##
#
serve:
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
  host: localhost

  ## Cross Origin Resource Sharing (CORS) ##
  #
  # Configure [Cross Origin Resource Sharing (CORS)](http://www.w3.org/TR/cors/) using the following options.
  #
  cors:
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
      - https://example.com
      - https://*.example.com
      - https://*.foo.example.com

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
      - GET

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
      - ''

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
      - ''

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
    allow_credentials: false

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
    max_age: -100000000

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

  ## HTTPS ##
  #
  # Configure HTTP over TLS (HTTPS). All options can also be set using environment variables by replacing dots (`.`) with underscores (`_`) and uppercasing the key. For example, `some.prefix.tls.key.path` becomes `export SOME_PREFIX_TLS_KEY_PATH`. If all keys are left undefined, TLS will be disabled.
  #
  tls:
    ## TLS Certificate (PEM) ##
    #
    cert:
      ## base64 ##
      #
      # Set this value using environment variables on
      # - Linux/macOS:
      #    $ export SERVE_TLS_CERT_BASE64=<value>
      # - Windows Command Line (CMD):
      #    > set SERVE_TLS_CERT_BASE64=<value>
      #
      base64: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tXG5NSUlEWlRDQ0FrMmdBd0lCQWdJRVY1eE90REFOQmdr...

      ## path ##
      #
      # Set this value using environment variables on
      # - Linux/macOS:
      #    $ export SERVE_TLS_CERT_PATH=<value>
      # - Windows Command Line (CMD):
      #    > set SERVE_TLS_CERT_PATH=<value>
      #
      path: path/to/file.pem

    ## Private Key (PEM) ##
    #
    key:
      ## base64 ##
      #
      # Set this value using environment variables on
      # - Linux/macOS:
      #    $ export SERVE_TLS_KEY_BASE64=<value>
      # - Windows Command Line (CMD):
      #    > set SERVE_TLS_KEY_BASE64=<value>
      #
      base64: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tXG5NSUlEWlRDQ0FrMmdBd0lCQWdJRVY1eE90REFOQmdr...

      ## path ##
      #
      # Set this value using environment variables on
      # - Linux/macOS:
      #    $ export SERVE_TLS_KEY_PATH=<value>
      # - Windows Command Line (CMD):
      #    > set SERVE_TLS_KEY_PATH=<value>
      #
      path: path/to/file.pem

  ## Port ##
  #
  # The port to listen on.
  #
  # Default value: 4456
  #
  # Minimum value: 1
  #
  # Maximum value: 65535
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

## Profiling ##
#
# Enables CPU or memory profiling if set. For more details on profiling Go programs read [Profiling Go Programs](https://blog.golang.org/profiling-go-programs).
#
# One of:
# - cpu
# - mem
# - ""
#
# Set this value using environment variables on
# - Linux/macOS:
#    $ export PROFILING=<value>
# - Windows Command Line (CMD):
#    > set PROFILING=<value>
#
profiling: cpu

## Log ##
#
# Configure logging using the following options. Logging will always be sent to stdout and stderr.
#
log:
  ## Format ##
  #
  # The log format can either be text or JSON.
  #
  # Default value: text
  #
  # One of:
  # - text
  # - json
  #
  # Set this value using environment variables on
  # - Linux/macOS:
  #    $ export LOG_FORMAT=<value>
  # - Windows Command Line (CMD):
  #    > set LOG_FORMAT=<value>
  #
  format: text

  ## Level ##
  #
  # Debug enables stack traces on errors. Can also be set using environment variable LOG_LEVEL.
  #
  # Default value: info
  #
  # One of:
  # - panic
  # - fatal
  # - error
  # - warn
  # - info
  # - debug
  #
  # Set this value using environment variables on
  # - Linux/macOS:
  #    $ export LOG_LEVEL=<value>
  # - Windows Command Line (CMD):
  #    > set LOG_LEVEL=<value>
  #
  level: panic

## tracing ##
#
# ORY Hydra supports distributed tracing.
#
tracing:
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
  service_name: ORY Keto

  ## providers ##
  #
  providers:
    ## jaeger ##
    #
    # Configures the jaeger tracing backend.
    #
    jaeger:
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

  ## provider ##
  #
  # Set this to the tracing backend you wish to use. Currently supports jaeger. If omitted or empty, tracing will be disabled.
  #
  # One of:
  # - jaeger
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

## The Keto version this config is written for. ##
#
# SemVer according to https://semver.org/ prefixed with `v` as in our releases.
#
# Set this value using environment variables on
# - Linux/macOS:
#    $ export VERSION=<value>
# - Windows Command Line (CMD):
#    > set VERSION=<value>
#
version: v0.0.0

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
```
