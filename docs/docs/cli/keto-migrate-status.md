---
id: keto-migrate-status
title: keto migrate status
description: keto migrate status Get the current migration status
---

<!--
This file is auto-generated.

To improve this file please make your change against the appropriate "./cmd/*.go" file.
-->
## keto migrate status

Get the current migration status

### Synopsis

Get the current migration status.
This does not affect namespaces. Use `keto namespace migrate status` for migrating namespaces.

```
keto migrate status [flags]
```

### Options

```
  -f, --format string   Set the output format. One of table, json, and json-pretty. (default &#34;default&#34;)
  -h, --help            help for status
  -q, --quiet           Be quiet with output printing.
```

### Options inherited from parent commands

```
  -c, --config strings   Config files to load, overwriting in the order specified. (default [/Users/foobar/keto.yml])
```

### SEE ALSO

* [keto migrate](keto-migrate)	 - Commands to migrate the database

