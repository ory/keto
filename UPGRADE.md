# Upgrading

The intent of this document is to make migration of breaking changes as easy as possible. Please note that not all
breaking changes might be included here. Please check the [CHANGELOG.md](./CHANGELOG.md) for a full list of changes
before finalizing the upgrade process.

## 1.0.0-beta.9

### CORS is disabled by default

A new environment variable `CORS_ENABLED` was introduced. It sets whether CORS is enabled ("true") or not ("false")".
Default is disabled.
