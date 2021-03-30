---
id: keto-serve
title: keto serve
description: keto serve Starts the server and serves the HTTP REST and gRPC APIs
---

<!--
This file is auto-generated.

To improve this file please make your change against the appropriate "./cmd/*.go" file.
-->

## keto serve

Starts the server and serves the HTTP REST and gRPC APIs

### Synopsis

This command opens the network ports and listens to HTTP and gRPC API requests.

## Configuration

ORY Keto can be configured using environment variables as well as a
configuration file. For more information on configuration options, open the
configuration documentation:

&gt;&gt; https://www.ory.sh/keto/docs/reference/configuration &lt;&lt;

```
keto serve [flags]
```

### Options

```
      --disable-telemetry   Disable anonymized telemetry reports - for more information please visit https://www.ory.sh/docs/ecosystem/sqa (default true)
  -h, --help                help for serve
      --sqa-opt-out         Disable anonymized telemetry reports - for more information please visit https://www.ory.sh/docs/ecosystem/sqa (default true)
```

### Options inherited from parent commands

```
  -c, --config strings   Config files to load, overwriting in the order specified. (default [/home/circleci/keto.yml])
```

### SEE ALSO

- [keto](keto) - Global and consistent permission and authorization server
