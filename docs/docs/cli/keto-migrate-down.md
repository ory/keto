---
id: keto-migrate-down
title: keto migrate down
description: keto migrate down Migrate the database down
---

<!--
This file is auto-generated.

To improve this file please make your change against the appropriate "./cmd/*.go" file.
-->
## keto migrate down

Migrate the database down

### Synopsis

Migrate the database down a specific amount of steps.
Pass 0 steps to fully migrate down.

```
keto migrate down &lt;steps&gt; [flags]
```

### Options

```
  -f, --format string   Set the output format. One of table, json, and json-pretty. (default &#34;default&#34;)
  -h, --help            help for down
  -q, --quiet           Be quiet with output printing.
  -y, --yes             yes to all questions, no user input required
```

### Options inherited from parent commands

```
  -c, --config strings   Config files to load, overwriting in the order specified. (default [/home/patrik/keto.yml])
```

### SEE ALSO

* [keto migrate](keto-migrate)	 - Commands to migrate the database

