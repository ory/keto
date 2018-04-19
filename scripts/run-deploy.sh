#!/bin/bash

set -euo pipefail

gox -ldflags "-X github.com/ory/keto/cmd.Version=`git describe --tags` -X github.com/ory/keto/cmd.BuildTime=`TZ=UTC date -u '+%Y-%m-%dT%H:%M:%SZ'` -X github.com/ory/keto/cmd.GitHash=`git rev-parse HEAD`" -output "dist/{{.Dir}}-{{.OS}}-{{.Arch}}";
npm version -f --no-git-tag-version $(git describe --tag);
