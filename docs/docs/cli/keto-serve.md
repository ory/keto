---
id: keto-serve
title: keto serve
description: keto serve Starts the server and serves the HTTP REST API
---

<!--
This file is auto-generated.

To improve this file please make your change against the appropriate "./cmd/*.go" file.
-->

## keto serve

Starts the server and serves the HTTP REST API

### Synopsis

This command opens a network port and listens to HTTP/2 API requests.

## Configuration

ORY Keto can be configured using environment variables as well as a
configuration file. For more information on configuration options, open the
configuration documentation:

> > https://github.com/ory/keto/blob/master/docs/config.yaml <<

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

- [keto](keto) -
