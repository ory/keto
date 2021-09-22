---
id: keto-relation-tuple-create
title: keto relation-tuple create
description: keto relation-tuple create Create relation tuples from JSON files
---

<!--
This file is auto-generated.

To improve this file please make your change against the appropriate "./cmd/*.go" file.
-->
## keto relation-tuple create

Create relation tuples from JSON files

### Synopsis

Create relation tuples from JSON files.
A directory will be traversed and all relation tuples will be created.
Pass the special filename `-` to read from STD_IN.

```
keto relation-tuple create &lt;relation-tuple.json&gt; [&lt;relation-tuple-dir&gt;] [flags]
```

### Options

```
  -f, --format string         Set the output format. One of table, json, and json-pretty. (default &#34;default&#34;)
  -h, --help                  help for create
  -q, --quiet                 Be quiet with output printing.
      --read-remote string    Remote URL of the read API endpoint. (default &#34;127.0.0.1:4466&#34;)
      --write-remote string   Remote URL of the write API endpoint. (default &#34;127.0.0.1:4467&#34;)
```

### Options inherited from parent commands

```
  -c, --config strings   Config files to load, overwriting in the order specified. (default [/home/patrik/keto.yml])
```

### SEE ALSO

* [keto relation-tuple](keto-relation-tuple)	 - Read and manipulate relation tuples

