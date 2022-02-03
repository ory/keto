---
id: keto-migrate
title: keto migrate
description: keto migrate Commands to migrate the database
---

<!--
This file is auto-generated.

To improve this file please make your change against the appropriate "./cmd/*.go" file.
-->

## keto migrate

Commands to migrate the database

### Synopsis

Commands to migrate the database. This does not affect namespaces. Use
`keto namespace migrate` for migrating namespaces.

### Options

```
  -h, --help   help for migrate
```

### Options inherited from parent commands

```
  -c, --config strings   Config files to load, overwriting in the order specified. (default [/home/runner/keto.yml])
```

### SEE ALSO

- [keto](keto) - Global and consistent permission and authorization server
- [keto migrate down](keto-migrate-down) - Migrate the database down
- [keto migrate status](keto-migrate-status) - Get the current migration status
- [keto migrate up](keto-migrate-up) - Migrate the database up
