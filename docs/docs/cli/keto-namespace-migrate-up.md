---
id: keto-namespace-migrate-up
title: keto namespace migrate up
description: keto namespace migrate up Migrate a namespace up
---

<!--
This file is auto-generated.

To improve this file please make your change against the appropriate "./cmd/*.go" file.
-->

## keto namespace migrate up

Migrate a namespace up

### Synopsis

Migrate a namespace up to the most recent migration.

```
keto namespace migrate up &lt;namespace-name&gt; [flags]
```

### Options

```
  -f, --format string         Set the output format. One of table, json, and json-pretty. (default &#34;default&#34;)
  -h, --help                  help for up
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
