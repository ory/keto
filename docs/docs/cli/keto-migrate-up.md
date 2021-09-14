---
id: keto-migrate-up
title: keto migrate up
description: keto migrate up Migrate the database up
---

<!--
This file is auto-generated.

To improve this file please make your change against the appropriate "./cmd/*.go" file.
-->

## keto migrate up

Migrate the database up

### Synopsis

Migrate the database up. This does not affect namespaces. Use
`keto namespace migrate up` for migrating namespaces.

```
keto migrate up [flags]
```

### Options

```
  -f, --format string   Set the output format. One of table, json, and json-pretty. (default &#34;default&#34;)
  -h, --help            help for up
  -q, --quiet           Be quiet with output printing.
  -y, --yes             yes to all questions, no user input required
```

### Options inherited from parent commands

```
  -c, --config strings   Config files to load, overwriting in the order specified. (default [/home/circleci/keto.yml])
```

### SEE ALSO

- [keto migrate](keto-migrate) - Commands to migrate the database
