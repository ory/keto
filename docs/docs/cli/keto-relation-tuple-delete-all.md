---
id: keto-relation-tuple-delete-all
title: keto relation-tuple delete-all
description:
  keto relation-tuple delete-all Delete ALL relation tuples matching the
  relation query.
---

<!--
This file is auto-generated.

To improve this file please make your change against the appropriate "./cmd/*.go" file.
-->

## keto relation-tuple delete-all

Delete ALL relation tuples matching the relation query.

### Synopsis

Delete all relation tuples matching the relation query. It is recommended to
first run the command without the `--force` flag to verify that the operation is
safe.

```
keto relation-tuple delete-all [flags]
```

### Options

```
      --force                 Force the deletion of relation tuples
  -f, --format string         Set the output format. One of table, json, and json-pretty. (default &#34;default&#34;)
  -h, --help                  help for delete-all
      --namespace string      Set the requested namespace
      --object string         Set the requested object
  -q, --quiet                 Be quiet with output printing.
      --read-remote string    Remote URL of the read API endpoint. (default &#34;127.0.0.1:4466&#34;)
      --relation string       Set the requested relation
      --subject-id string     Set the requested subject ID
      --subject-set string    Set the requested subject set; format: &#34;namespace:object#relation&#34;
      --write-remote string   Remote URL of the write API endpoint. (default &#34;127.0.0.1:4467&#34;)
```

### Options inherited from parent commands

```
  -c, --config strings   Config files to load, overwriting in the order specified. (default [/home/patrik/keto.yml])
```

### SEE ALSO

- [keto relation-tuple](keto-relation-tuple) - Read and manipulate relation
  tuples
