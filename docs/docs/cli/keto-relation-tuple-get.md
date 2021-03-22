---
id: keto-relation-tuple-get
title: keto relation-tuple get
description: keto relation-tuple get
---

<!--
This file is auto-generated.

To improve this file please make your change against the appropriate "./cmd/*.go" file.
-->

## keto relation-tuple get

```
keto relation-tuple get <namespace> [flags]
```

### Options

```
  -f, --format string         Set the output format. One of table, json, and json-pretty. (default "default")
  -h, --help                  help for get
      --object string         Set the requested object
      --page-size int32       maximum number of items to return (default 100)
      --page-token string     page token acquired from a previous response
  -q, --quiet                 Be quiet with output printing.
      --read-remote string    Remote URL of the read API endpoint. (default "127.0.0.1:4466")
      --relation string       Set the requested relation
      --subject string        Set the requested subject
      --write-remote string   Remote URL of the write API endpoint. (default "127.0.0.1:4467")
```

### Options inherited from parent commands

```
  -c, --config strings   Config files to load, overwriting in the order specified. (default [/home/patrik/keto.yml])
```

### SEE ALSO

- [keto relation-tuple](keto-relation-tuple) -
