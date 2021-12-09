---
id: keto-check
title: keto check
description: keto check Check whether a subject has a relation on an object
---

<!--
This file is auto-generated.

To improve this file please make your change against the appropriate "./cmd/*.go" file.
-->

## keto check

Check whether a subject has a relation on an object

### Synopsis

Check whether a subject has a relation on an object. This method resolves
subject sets and subject set rewrites.

```
keto check &lt;subject&gt; &lt;relation&gt; &lt;namespace&gt; &lt;object&gt; [flags]
```

### Options

```
  -f, --format string         Set the output format. One of table, json, and json-pretty. (default &#34;default&#34;)
  -h, --help                  help for check
  -d, --max-depth int32       Maximum depth of the search tree. If the value is less than 1 or greater than the global max-depth then the global max-depth will be used instead. (default 0)
  -q, --quiet                 Be quiet with output printing.
      --read-remote string    Remote URL of the read API endpoint. (default &#34;127.0.0.1:4466&#34;)
      --write-remote string   Remote URL of the write API endpoint. (default &#34;127.0.0.1:4467&#34;)
```

### Options inherited from parent commands

```
  -c, --config strings   Config files to load, overwriting in the order specified. (default [/home/circleci/keto.yml])
```

### SEE ALSO

- [keto](keto) - Global and consistent permission and authorization server
