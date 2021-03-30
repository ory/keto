---
id: keto-namespace-migrate-status
title: keto namespace migrate status
description:
  keto namespace migrate status Get the current namespace migration status
---

<!--
This file is auto-generated.

To improve this file please make your change against the appropriate "./cmd/*.go" file.
-->

## keto namespace migrate status

Get the current namespace migration status

### Synopsis

Get the current migration status of one specific namespace. Does not apply any
changes.

```
keto namespace migrate status &lt;namespace-name&gt; [flags]
```

### Options

```
  -f, --format string   Set the output format. One of table, json, and json-pretty. (default &#34;default&#34;)
  -h, --help            help for status
  -q, --quiet           Be quiet with output printing.
```

### Options inherited from parent commands

```
  -c, --config strings   Config files to load, overwriting in the order specified. (default [/home/circleci/keto.yml])
```

### SEE ALSO

- [keto namespace migrate](keto-namespace-migrate) - Migrate a namespace
