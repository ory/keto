---
id: keto-relation-tuple-parse
title: keto relation-tuple parse
description: keto relation-tuple parse Parse human readable relation tuples
---

<!--
This file is auto-generated.

To improve this file please make your change against the appropriate "./cmd/*.go" file.
-->

## keto relation-tuple parse

Parse human readable relation tuples

### Synopsis

Parse human readable relation tuples as used in the documentation. Supports
various output formats. Especially useful for piping into other commands by
using `--format json`. Ignores comments (starting with `//`) and blank lines.

```
keto relation-tuple parse [flags]
```

### Options

```
  -f, --format string   Set the output format. One of table, json, and json-pretty. (default &#34;default&#34;)
  -h, --help            help for parse
  -q, --quiet           Be quiet with output printing.
```

### Options inherited from parent commands

```
  -c, --config strings   Config files to load, overwriting in the order specified. (default [/home/circleci/keto.yml])
```

### SEE ALSO

- [keto relation-tuple](keto-relation-tuple) - Read and manipulate relation
  tuples
