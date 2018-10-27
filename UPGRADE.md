# Upgrading

The intent of this document is to make migration of breaking changes as easy as possible. Please note that not all
breaking changes might be included here. Please check the [CHANGELOG.md](./CHANGELOG.md) for a full list of changes
before finalizing the upgrade process.

## 0.1.0-sandbox

### Rework

ORY Keto has been completely reworked. The major goals of this refactoring is:

1. To allow easy extension of existing access control mechanisms.
2. Improve stability and responsiveness.
3. Support more than one access control mechanism. Future mechanisms include: RBAC, ACL, AWS IAM Policies, ...

For this reason the following thing have been completely deprecated:

1. Authorizers

The following things have changed:

1. ORY Keto no longer uses ORY Ladon as the engine but instead relies on the [Open Policy Agent](http://openpolicyagent.org/).
2. The CLI commands have changed - apart from `serve`, `version`, `migrate sql` - entirely.
3. The API has changed (read the next section for information on this).

The following things have been added:

1. ORY (Ladon) Access Control Policies with exact string matching
2. ORY (Ladon) Access Control Policies with glob string matching

### API Changes

API locations

## pre-release

### CORS is disabled by default

A new environment variable `CORS_ENABLED` was introduced. It sets whether CORS is enabled ("true") or not ("false")".
Default is disabled.
