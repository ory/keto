---
id: production
title: Going to Production
---

:::warn

This document is still in development.

:::

## Database

ORY Keto requires a production-grade database such as PostgreSQL, MySQL,
CockroachDB. Do not use SQLite in production!

### Write API

Never expose the ORY Keto Write API to the internet unsecured. Always require
authorization. A good practice is to not expose the Write API at all to the
public internet and use a Zero Trust Networking Architecture within your
intranet.

## Scaling

There are no additional requirements for scaling ORY Keto, just spin up
another container!
