---
id: keto-namespace-migrate-legacy
title: keto namespace migrate legacy
description:
  keto namespace migrate legacy Migrate a namespace from v0.6.x to v0.7.x and
  later.
---

<!--
This file is auto-generated.

To improve this file please make your change against the appropriate "./cmd/*.go" file.
-->

## keto namespace migrate legacy

Migrate a namespace from v0.6.x to v0.7.x and later.

### Synopsis

Migrate a legacy namespaces from v0.6.x to the v0.7.x and later. This step only
has to be executed once. If no namespace is specified, all legacy namespaces
will be migrated. Please ensure that namespace IDs did not change in the config
file and you have a backup in case something goes wrong!

```
keto namespace migrate legacy [&lt;namespace-name&gt;] [flags]
```

### Options

```
      --down-only             Migrate legacy namespace(s) only down.
  -f, --format string         Set the output format. One of table, json, and json-pretty. (default &#34;default&#34;)
  -h, --help                  help for legacy
  -q, --quiet                 Be quiet with output printing.
      --read-remote string    Remote URL of the read API endpoint. (default &#34;127.0.0.1:4466&#34;)
      --write-remote string   Remote URL of the write API endpoint. (default &#34;127.0.0.1:4467&#34;)
  -y, --yes                   yes to all questions, no user input required
```

### Options inherited from parent commands

```
  -c, --config strings   Config files to load, overwriting in the order specified. (default [/home/circleci/keto.yml])
```

### SEE ALSO

- [keto namespace migrate](keto-namespace-migrate) - Migrate a namespace
