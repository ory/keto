---
id: keto-namespace-migrate-down
title: keto namespace migrate down
description: keto namespace migrate down Migrate a namespace down
---

<!--
This file is auto-generated.

To improve this file please make your change against the appropriate "./cmd/*.go" file.
-->

## keto namespace migrate down

Migrate a namespace down

### Synopsis

Migrate a namespace down. Pass 0 steps to fully migrate down.

```
keto namespace migrate down <namespace-name> <steps> [flags]
```

### Options

```
  -f, --format string         Set the output format. One of table, json, and json-pretty. (default "default")
  -h, --help                  help for down
  -q, --quiet                 Be quiet with output printing.
      --read-remote string    Remote URL of the read API endpoint. (default "127.0.0.1:4466")
      --write-remote string   Remote URL of the write API endpoint. (default "127.0.0.1:4467")
  -y, --yes                   answer all questions with yes
```

### Options inherited from parent commands

```
  -c, --config strings   Config files to load, overwriting in the order specified. (default [/home/circleci/keto.yml])
```

### SEE ALSO

- [keto namespace migrate](keto-namespace-migrate) - Migrate a namespace
