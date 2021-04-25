---
id: keto-status
title: keto status
description: keto status Get the status of the upstream Keto instance
---

<!--
This file is auto-generated.

To improve this file please make your change against the appropriate "./cmd/*.go" file.
-->
## keto status

Get the status of the upstream Keto instance

### Synopsis

Get a status report about the upstream Keto instance. Can also block until the service is healthy.

```
keto status [flags]
```

### Options

```
  -b, --block                 block until the service is healthy
      --endpoint string       which endpoint to use; one of {read, write} (default &#34;read&#34;)
  -h, --help                  help for status
  -q, --quiet                 Be quiet with output printing.
      --read-remote string    Remote URL of the read API endpoint. (default &#34;127.0.0.1:4466&#34;)
      --write-remote string   Remote URL of the write API endpoint. (default &#34;127.0.0.1:4467&#34;)
```

### Options inherited from parent commands

```
  -c, --config strings   Config files to load, overwriting in the order specified. (default [/Users/foobar/keto.yml])
```

### SEE ALSO

* [keto](keto)	 - Global and consistent permission and authorization server

