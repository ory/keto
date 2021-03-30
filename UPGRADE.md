# Upgrading

The intent of this document is to make migration of breaking changes as easy as
possible. Please note that not all breaking changes might be included here.
Please check the [CHANGELOG.md](./CHANGELOG.md) for a full list of changes
before finalizing the upgrade process.

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->

- [0.4.0](#040)
- [0.3.0-sandbox](#030-sandbox)
  - [Configuration](#configuration)
  - [ORY Access Control Policies Allowed Endpoint](#ory-access-control-policies-allowed-endpoint)
  - [SDK](#sdk)
- [0.2.0-sandbox](#020-sandbox)
  - [Conceptual changes](#conceptual-changes)
    - [Deprecated](#deprecated)
    - [Changes](#changes)
    - [Additions](#additions)
    - [Untouched](#untouched)
  - [API Changes](#api-changes)
    - [Renamed Endpoints](#renamed-endpoints)
    - [Reworked Endpoints](#reworked-endpoints)
    - [New Endpoints](#new-endpoints)
  - [Migration](#migration)
    - [SQL](#sql)
- [0.0.1](#001)
  - [CORS is disabled by default](#cors-is-disabled-by-default)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## 0.4.0-sandbox

This release focuses on a rework of the SDK pipeline. First of all, we have
introduced new SDKs for all popular programming languages and published them on
their respective package repositories:

- [Python](https://pypi.org/project/ory-keto-client/)
- [PHP](https://packagist.org/packages/ory/keto-client)
- [Go](https://github.com/ory/keto-client-go)
- [NodeJS](https://www.npmjs.com/package/@oryd/keto-client) (with TypeScript)
- [Java](https://search.maven.org/artifact/sh.ory.keto/keto-client)
- [Ruby](https://rubygems.org/gems/ory-keto-client)

The SDKs hosted in this repository (under ./sdk/...) have been completely
removed. Please use only the SDKs from the above sources from now on as it will
also remove several issues that were caused by the previous SDK pipeline.

Unfortunately, there were breaking changes introduced by the new SDK generation:

- Several structs and fields have been renamed in the Go SDK. However, nothing
  else changed so upgrading should be a matter of half an hour if you made
  extensive use of the SDK, or several minutes if just one or two methods are
  being used.
- All other SDKs changed to `openapi-generator`, which is a better maintained
  generator that creates better code than the one previously used. This
  manifests in TypeScript definitions for the NodeJS SDK and several other
  goodies. We do not have a proper migration path for those, unfortunately.

If you have issues with upgrading the SDK, please let us know in an issue on
this repository!

## 0.3.0-sandbox

### Configuration

The configuration management was updated and now allows configuration via a
config file. Environment variables can still be used to configure ORY Keto but
have been updated. However, old env vars still work but will yield a warning.

An overview of an exemplary configuration file can be found in
[./docs/config.yml](https://github.com/ory/hydra/blob/master/docs/config.yaml).

### ORY Access Control Policies Allowed Endpoint

Endpoint `/engines/acp/ory/{flavor}/allowed` now returns a 403 error when the
request is disallowed.

### SDK

Generation of the Go SDK has moved from
[`swagger-codegen`](https://github.com/swagger-api/swagger-codegen) to
[`go-swagger`](https://github.com/go-swagger/go-swagger). If you wish to migrate
your existing SDK integration please open an issue.

## 0.2.0-sandbox

ORY Keto has been completely reworked. The major goals of this refactoring are:

1. To allow easy extension of existing access control mechanisms.
2. Improve stability and responsiveness.
3. Support more than one access control mechanism. Future mechanisms include:
   RBAC, ACL, AWS IAM Policies, ...

We know that these changes seem massive. They are, but they will benefit the
long-term use of this particular piece of software, and they will allow you to
build better systems.

If you relied on ORY Keto before this release and you are looking for a
migration path, don't hesitate to ask in [the forums](https://community.ory.sh/)
or open a [GitHub issue](https://github.com/ory/keto/issues/new/). Feel free to
do the same if you want the access control policy feature implemented in ORY
Hydra before version `1.0.0`.

### Conceptual changes

#### Deprecated

The following things have been completely deprecated:

1. Authorizers,
2. Previous storage mechanisms.

#### Changes

The following things have changed:

1. ORY Keto no longer uses ORY Ladon as the engine but instead relies on the
   [Open Policy Agent](http://openpolicyagent.org/). The concept of ORY Ladon
   Access Policies are working exactly like before, the internal logic however
   was rewritten in Rego.
2. The "Warden" concept has been deprecated and replaced.
3. The CLI commands have changed - apart from `serve`, `version`,
   `migrate sql` - entirely.
4. The API has changed (read the next section for information on this).
5. Environment variables changed or have been removed.

#### Additions

The following things have been added:

1. ORY (Ladon) Access Control Policies with `exact` string `matching-strategy`.
2. ORY (Ladon) Access Control Policies with `glob` string `matching-strategy`.

#### Untouched

The following things remain conceptually untouched:

1. ORY (Ladon) Access Control Policies with `regex` string `matching-strategy`.
   This is the logic that ORY Ladon and previous versions of ORY Keto implement.

### API Changes

#### Renamed Endpoints

- `GET,PUT,POST,DELETE /policies[/<id>]` moved to
  `/engines/acp/ory/<matching-strategy>/policies[/<id>]`.
  - `POST /policies` has been deprecated and merged with `PUT /policies/<id>`
    which is now available at
    `PUT /engines/acp/ory/<matching-strategy>/policies` and will upsert (insert
    or update) the policy identified by the `id` field in the JSON payload.
  - The request & response payloads **did not change** nor did any of the
    concepts.
- `GET,PUT,POST,DELETE /roles[/<id>]` moved to
  `/engines/acp/ory/<matching-strategy>/roles[/<id>]`.
  - `POST /roles` has been deprecated and merged with `PUT /roles/<id>` which is
    now available at `PUT /engines/acp/ory/<matching-strategy>/policies` and
    will upsert (insert or update) the role identified by the `id` field in the
    JSON payload.
  - The request & response payloads **did not change** nor did any of the
    concepts.
- `POST,GET /roles/<id>/members` move to
  `/engines/acp/ory/<matching-strategy>/roles/<id>/members`.
  - `POST /roles` has been moved to
    `PUT /engines/acp/ory/<matching-strategy>/policies/<id>/members` and will
    upsert (insert or update) the role identified by the `id` field in the URL
    path.
  - The request & response payloads **did not change** nor did any of the
    concepts.

#### Reworked Endpoints

The Warden concept has been deprecated. Previously, it was possible to send
credentials alongside requests for prior authentication. This concept interfered
with the clear boundary ORY Keto is focusing on, which is permissioning
concepts.

The Warden API featured endpoints such as:

- `/warden/oauth2/access-tokens/authorize`: Permformed OAuth 2.0 Token
  Introspection on the `token` field, took the `sub` value of the introspection
  and used that as input to ORY (Ladon) Access Control Policies.
- `/warden/oauth2/clients/authorize`: Validated the HTTP Basic Authorization
  Header using the OAuth 2.0 Client Credentials grant and took the `username`
  value of the HTTP Basic Authorization Header and used that as input to ORY
  (Ladon) Access Control Policies.

These endpoints have been deprecated without replacement. Another endpoint was
`/warden/subjects/authorize` which used the format
`{ "subject": "peter", "action": "delete", "resource": "something:valuable" }`
as syntax. This endpoint is available in the exact same format at
`/engines/acp/ory/<matching-strategy>/allowed`.

#### New Endpoints

- `GET /version`: Returns the running software version.
- `GET /health/ready`: Returns `{"status": "ok"}` with a 200 HTTP response if
  the service is ready to accept connections and handle data.
- `GET /health/alive`: Returns `{"status": "ok"}` with a 200 HTTP response if
  the service is ready to accept connections.

### Migration

If you relied on ORY Keto before this release and you are looking for a
migration path, don't hesitate to [contact us](mailto:hi@ory.sh). We will help
you migrate and improve this guide as we see more migration use cases.

#### SQL

The SQL schema changed completely and it is not possible to migrate from the
previous version to this version with just using `keto migrate sql`. Please ask
in [the forums](https://community.ory.sh/) or open a
[GitHub issue](https://github.com/ory/keto/issues/new/) if this affects you.

## 0.0.1

### CORS is disabled by default

A new environment variable `CORS_ENABLED` was introduced. It sets whether CORS
is enabled ("true") or not ("false")". Default is disabled.
