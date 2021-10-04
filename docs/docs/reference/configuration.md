---
id: configuration
title: Configuration
---

<!-- THIS FILE IS BEING AUTO-GENERATED. DO NOT MODIFY IT AS ALL CHANGES WILL BE OVERWRITTEN.
OPEN AN ISSUE IF YOU WOULD LIKE TO MAKE ADJUSTMENTS HERE AND MAINTAINERS WILL HELP YOU LOCATE THE RIGHT
FILE -->

You can load the config file from another source using the
`-c path/to/config.yaml` or `--config path/to/config.yaml` flag:
`keto --config path/to/config.yaml`.

Config files can be formatted as JSON, YAML and TOML. Some configuration values
support reloading without server restart. All configuration values can be set
using environment variables, as documented below.

:::warning Disclaimer

This reference configuration documents all keys, also deprecated ones! It is a
reference for all possible configuration values.

If you are looking for an example configuration, it is better to try out the
quickstart.

:::

To find out more about edge cases like setting string array values through
environmental variables head to the
[Configuring ORY services](https://www.ory.sh/docs/ecosystem/configuring)
section.

```yaml
## ORY Keto Configuration
#

## serve ##
#
serve:
  ## Write API (http and gRPC) ##
  #
  write:
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
    #    $ export SERVE_WRITE_HOST=<value>
    # - Windows Command Line (CMD):
    #    > set SERVE_WRITE_HOST=<value>
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
      #    $ export SERVE_WRITE_CORS_ALLOWED_ORIGINS=<value>
      # - Windows Command Line (CMD):
      #    > set SERVE_WRITE_CORS_ALLOWED_ORIGINS=<value>
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
      #    $ export SERVE_WRITE_CORS_ALLOWED_METHODS=<value>
      # - Windows Command Line (CMD):
      #    > set SERVE_WRITE_CORS_ALLOWED_METHODS=<value>
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
      #    $ export SERVE_WRITE_CORS_ALLOWED_HEADERS=<value>
      # - Windows Command Line (CMD):
      #    > set SERVE_WRITE_CORS_ALLOWED_HEADERS=<value>
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
      #    $ export SERVE_WRITE_CORS_EXPOSED_HEADERS=<value>
      # - Windows Command Line (CMD):
      #    > set SERVE_WRITE_CORS_EXPOSED_HEADERS=<value>
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
      #    $ export SERVE_WRITE_CORS_ALLOW_CREDENTIALS=<value>
      # - Windows Command Line (CMD):
      #    > set SERVE_WRITE_CORS_ALLOW_CREDENTIALS=<value>
      #
      allow_credentials: false

      ## Maximum Age ##
      #
      # Indicates how long (in seconds) the results of a preflight request can be cached. The default is 0 which stands for no max age.
      #
      # Set this value using environment variables on
      # - Linux/macOS:
      #    $ export SERVE_WRITE_CORS_MAX_AGE=<value>
      # - Windows Command Line (CMD):
      #    > set SERVE_WRITE_CORS_MAX_AGE=<value>
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
      #    $ export SERVE_WRITE_CORS_DEBUG=<value>
      # - Windows Command Line (CMD):
      #    > set SERVE_WRITE_CORS_DEBUG=<value>
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
      #    $ export SERVE_WRITE_CORS_ENABLED=<value>
      # - Windows Command Line (CMD):
      #    > set SERVE_WRITE_CORS_ENABLED=<value>
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
        #    $ export SERVE_WRITE_TLS_CERT_BASE64=<value>
        # - Windows Command Line (CMD):
        #    > set SERVE_WRITE_TLS_CERT_BASE64=<value>
        #
        base64: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tXG5NSUlEWlRDQ0FrMmdBd0lCQWdJRVY1eE90REFOQmdr...

        ## path ##
        #
        # Set this value using environment variables on
        # - Linux/macOS:
        #    $ export SERVE_WRITE_TLS_CERT_PATH=<value>
        # - Windows Command Line (CMD):
        #    > set SERVE_WRITE_TLS_CERT_PATH=<value>
        #
        path: path/to/file.pem

      ## Private Key (PEM) ##
      #
      key:
        ## base64 ##
        #
        # Set this value using environment variables on
        # - Linux/macOS:
        #    $ export SERVE_WRITE_TLS_KEY_BASE64=<value>
        # - Windows Command Line (CMD):
        #    > set SERVE_WRITE_TLS_KEY_BASE64=<value>
        #
        base64: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tXG5NSUlEWlRDQ0FrMmdBd0lCQWdJRVY1eE90REFOQmdr...

        ## path ##
        #
        # Set this value using environment variables on
        # - Linux/macOS:
        #    $ export SERVE_WRITE_TLS_KEY_PATH=<value>
        # - Windows Command Line (CMD):
        #    > set SERVE_WRITE_TLS_KEY_PATH=<value>
        #
        path: path/to/file.pem

    ## Port ##
    #
    # The port to listen on.
    #
    # Default value: 4467
    #
    # Minimum value: 0
    #
    # Maximum value: 65535
    #
    # Set this value using environment variables on
    # - Linux/macOS:
    #    $ export SERVE_WRITE_PORT=<value>
    # - Windows Command Line (CMD):
    #    > set SERVE_WRITE_PORT=<value>
    #
    port: 0

  ## Read API (http and gRPC) ##
  #
  read:
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
    #    $ export SERVE_READ_HOST=<value>
    # - Windows Command Line (CMD):
    #    > set SERVE_READ_HOST=<value>
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
      #    $ export SERVE_READ_CORS_ALLOWED_ORIGINS=<value>
      # - Windows Command Line (CMD):
      #    > set SERVE_READ_CORS_ALLOWED_ORIGINS=<value>
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
      #    $ export SERVE_READ_CORS_ALLOWED_METHODS=<value>
      # - Windows Command Line (CMD):
      #    > set SERVE_READ_CORS_ALLOWED_METHODS=<value>
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
      #    $ export SERVE_READ_CORS_ALLOWED_HEADERS=<value>
      # - Windows Command Line (CMD):
      #    > set SERVE_READ_CORS_ALLOWED_HEADERS=<value>
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
      #    $ export SERVE_READ_CORS_EXPOSED_HEADERS=<value>
      # - Windows Command Line (CMD):
      #    > set SERVE_READ_CORS_EXPOSED_HEADERS=<value>
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
      #    $ export SERVE_READ_CORS_ALLOW_CREDENTIALS=<value>
      # - Windows Command Line (CMD):
      #    > set SERVE_READ_CORS_ALLOW_CREDENTIALS=<value>
      #
      allow_credentials: false

      ## Maximum Age ##
      #
      # Indicates how long (in seconds) the results of a preflight request can be cached. The default is 0 which stands for no max age.
      #
      # Set this value using environment variables on
      # - Linux/macOS:
      #    $ export SERVE_READ_CORS_MAX_AGE=<value>
      # - Windows Command Line (CMD):
      #    > set SERVE_READ_CORS_MAX_AGE=<value>
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
      #    $ export SERVE_READ_CORS_DEBUG=<value>
      # - Windows Command Line (CMD):
      #    > set SERVE_READ_CORS_DEBUG=<value>
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
      #    $ export SERVE_READ_CORS_ENABLED=<value>
      # - Windows Command Line (CMD):
      #    > set SERVE_READ_CORS_ENABLED=<value>
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
        #    $ export SERVE_READ_TLS_CERT_BASE64=<value>
        # - Windows Command Line (CMD):
        #    > set SERVE_READ_TLS_CERT_BASE64=<value>
        #
        base64: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tXG5NSUlEWlRDQ0FrMmdBd0lCQWdJRVY1eE90REFOQmdr...

        ## path ##
        #
        # Set this value using environment variables on
        # - Linux/macOS:
        #    $ export SERVE_READ_TLS_CERT_PATH=<value>
        # - Windows Command Line (CMD):
        #    > set SERVE_READ_TLS_CERT_PATH=<value>
        #
        path: path/to/file.pem

      ## Private Key (PEM) ##
      #
      key:
        ## base64 ##
        #
        # Set this value using environment variables on
        # - Linux/macOS:
        #    $ export SERVE_READ_TLS_KEY_BASE64=<value>
        # - Windows Command Line (CMD):
        #    > set SERVE_READ_TLS_KEY_BASE64=<value>
        #
        base64: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tXG5NSUlEWlRDQ0FrMmdBd0lCQWdJRVY1eE90REFOQmdr...

        ## path ##
        #
        # Set this value using environment variables on
        # - Linux/macOS:
        #    $ export SERVE_READ_TLS_KEY_PATH=<value>
        # - Windows Command Line (CMD):
        #    > set SERVE_READ_TLS_KEY_PATH=<value>
        #
        path: path/to/file.pem

    ## Port ##
    #
    # The port to listen on.
    #
    # Default value: 4466
    #
    # Minimum value: 0
    #
    # Maximum value: 65535
    #
    # Set this value using environment variables on
    # - Linux/macOS:
    #    $ export SERVE_READ_PORT=<value>
    # - Windows Command Line (CMD):
    #    > set SERVE_READ_PORT=<value>
    #
    port: 0

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
# Configure logging using the following options. Logs will always be sent to stdout and stderr.
#
log:
  ## Log Format ##
  #
  # The output format of log messages.
  #
  # Default value: text
  #
  # One of:
  # - json
  # - json_pretty
  # - gelf
  # - text
  #
  # Set this value using environment variables on
  # - Linux/macOS:
  #    $ export LOG_FORMAT=<value>
  # - Windows Command Line (CMD):
  #    > set LOG_FORMAT=<value>
  #
  format: json

  ## Leak Sensitive Log Values ##
  #
  # If set will leak sensitive values (e.g. emails) in the logs.
  #
  # Default value: false
  #
  # Set this value using environment variables on
  # - Linux/macOS:
  #    $ export LOG_LEAK_SENSITIVE_VALUES=<value>
  # - Windows Command Line (CMD):
  #    > set LOG_LEAK_SENSITIVE_VALUES=<value>
  #
  leak_sensitive_values: false

  ## Level ##
  #
  # The level of log entries to show. Debug enables stack traces on errors.
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
  # - trace
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
# Configure distributed tracing.
#
tracing:
  ## service_name ##
  #
  # Specifies the service name to use on the tracer.
  #
  # Examples:
  # - Ory Hydra
  # - Ory Kratos
  # - Ory Keto
  # - Ory Oathkeeper
  #
  # Set this value using environment variables on
  # - Linux/macOS:
  #    $ export TRACING_SERVICE_NAME=<value>
  # - Windows Command Line (CMD):
  #    > set TRACING_SERVICE_NAME=<value>
  #
  service_name: Ory Hydra

  ## providers ##
  #
  providers:
    ## zipkin ##
    #
    # Configures the zipkin tracing backend.
    #
    # Examples:
    # - server_url: http://localhost:9411/api/v2/spans
    #
    zipkin:
      ## server_url ##
      #
      # The address of Zipkin server where spans should be sent to.
      #
      # Set this value using environment variables on
      # - Linux/macOS:
      #    $ export TRACING_PROVIDERS_ZIPKIN_SERVER_URL=<value>
      # - Windows Command Line (CMD):
      #    > set TRACING_PROVIDERS_ZIPKIN_SERVER_URL=<value>
      #
      server_url: http://localhost:9411/api/v2/spans

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

      ## max_tag_value_length ##
      #
      # The value passed to the max tag value length that has been configured.
      #
      # Minimum value: 0
      #
      # Set this value using environment variables on
      # - Linux/macOS:
      #    $ export TRACING_PROVIDERS_JAEGER_MAX_TAG_VALUE_LENGTH=<value>
      # - Windows Command Line (CMD):
      #    > set TRACING_PROVIDERS_JAEGER_MAX_TAG_VALUE_LENGTH=<value>
      #
      max_tag_value_length: 0

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
  # Set this to the tracing backend you wish to use. Supports Jaeger, Zipkin DataDog, Elastic APM and Instana. If omitted or empty, tracing will be disabled. Use environment variables to configure DataDog (see https://docs.datadoghq.com/tracing/setup/go/#configuration).
  #
  # One of:
  # - jaeger
  # - zipkin
  # - datadog
  # - elastic-apm
  # - instana
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

## namespaces ##
#
# Namespace configuration or it's location.
#
# Default value: file://./keto_namespaces
#
# Set this value using environment variables on
# - Linux/macOS:
#    $ export NAMESPACES=<value>
# - Windows Command Line (CMD):
#    > set NAMESPACES=<value>
#
namespaces: http://a.aaa

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
version: v0.7.0-alpha.0.pre.1

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
