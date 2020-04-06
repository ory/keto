---
id: configure-deploy
title: Configure and Deploy
---

Like all other ORY services, ORY Keto is implemented following
[12factor principles](https://12factor.net) and completely stateless. To store
state, ORY Keto supports two types of storage adapters:

- **in-memory:** This adapter does not work with more than one instance
  ("cluster") and any state is lost after restarting the instance.
- **SQL:** This adapter works with more than one instance and state persists
  after restarts.

The SQL adapter supports two DBMS: PostgreSQL 9.6+ and MySQL 5.7+. Please note
that older MySQL versions may have issues with the database schema. We recommend
working with PostgreSQL as migrations will be faster.

This guide will:

1. Download and run a PostgreSQL container in Docker.
2. Download and run ORY Keto using Docker.

## Create a Network

As a first step, we create a network to which we connect all our Docker
containers. This enables the containers to communicate with each other.

```
$ docker network create ketoguide
```

## Start the PostgreSQL Container

For the purpose of this tutorial, we will use PostgreSQL as a database. As you
probably already know, don't run databases in Docker in production! For the sake
of this tutorial however, let's use Docker to quickly deploy the database.

```
$ docker run \
  --network ketoguide \
  --name ory-keto-example--postgres \
  -e POSTGRES_USER=keto \
  -e POSTGRES_PASSWORD=secret \
  -e POSTGRES_DB=keto \
  -d postgres:9.6
```

This command wil start a postgres instance with name
`ory-keto-example--postgres`, set up a database called `keto` and create a user
`keto` with password `secret`.

## Run the ORY Keto Service

```
# The database url points us at the postgres instance.
# This could also be an ephermal in-memory database (`export DSN=memory`)
# or a MySQL URI.
$ export DSN=postgres://keto:secret@ory-keto-example--postgres:5432/keto?sslmode=disable

# ORY Keto does not do magic, it requires conscious decisions.
# An example is running SQL migrations when setting up a new installation of ORY Keto
# or upgrading an existing one.
# This is equivalent to:
# DSN=postgres://keto:secret@ory-keto-example--postgres:5432/keto?sslmode=disable keto migrate sql`
$ docker run -it --rm \
  --network ketoguide \
  -e DSN=$DSN \
  oryd/keto:v0.5.3-alpha.3 \
  migrate sql -e

Applying `client` SQL migrations...
[...]
Migration successful!

# Let's run the server!
$ docker run -d \
  --name ory-keto-example--keto \
  --network ketoguide \
  -p 4466:4466 \
  -e DSN=$DSN \
  oryd/keto:v0.5.3-alpha.3 \
  serve
```

Great, the server is running now! Make sure to check the logs and see if there
were any errors or issues before moving on to the next step:

```
$ docker logs ory-keto-example--keto
```

You should see one line showing where the server is running:

```
time="2018-10-27T11:48:56Z" level=info msg="Listening on http://localhost:4466"
```

## Working with the CLI

Let's explore managing ORY Keto via the CLI. We will use the ORY Access Control
Policy Engine (`/engines/acp/ory`) with the `exact` matcher, define policies,
and check if particular users are allowed to do certain things. Let's create our
first policy:

```
$ mkdir policies

$ cat > policies/example-policy.json <<EOL
[{
    "id": "example-policy",
    "subjects": ["alice"],
    "resources": ["blog_posts:my-first-blog-post"],
    "actions": ["delete"],
    "effect": "allow"
}]
EOL

$ docker run -it --rm \
  --network ketoguide \
  -v $(pwd)/policies:/policies \
  -e KETO_URL=http://ory-keto-example--keto:4466/ \
  oryd/keto:v0.5.3-alpha.3 \
  engines acp ory policies import exact /policies/example-policy.json
```

Check if the policy has been created:

```
$ docker run -it --rm \
  --network ketoguide \
  -e KETO_URL=http://ory-keto-example--keto:4466/ \
  oryd/keto:v0.5.3-alpha.3 \
  engines acp ory policies get exact example-policy
{
  "actions": [
    "delete"
  ],
...
```

Check if Alice is allowed to delete the blog post:

```
$ docker run -it --rm \
  --network ketoguide \
  -e KETO_URL=http://ory-keto-example--keto:4466/ \
  oryd/keto:v0.5.3-alpha.3 \
  engines acp ory allowed exact alice blog_posts:my-first-blog-post delete
{
        "allowed": true
}
```

Other users like Bob can not delete it:

```
$ docker run -it --rm \
  --network ketoguide \
  -e KETO_URL=http://ory-keto-example--keto:4466/ \
  oryd/keto:v0.5.3-alpha.3 \
  engines acp ory allowed exact bob blog_posts:my-first-blog-post delete
{
        "allowed": false
}
```
