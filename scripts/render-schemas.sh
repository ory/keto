#!/bin/bash

set -euxo pipefail

ory_x_version="$(go list -f '{{.Version}}' -m github.com/ory/x)"

sed "s!ory://tracing-config!https://raw.githubusercontent.com/ory/x/$ory_x_version/tracing/config.schema.json!g;
s!ory://logging-config!https://raw.githubusercontent.com/ory/x/$ory_x_version/logrusx/config.schema.json!g" internal/driver/config/config.schema.json > .schema/config.schema.json

git config user.email "zepatrik@users.noreply.github.com"
git config user.name "zepatrik"

git add .schema/config.schema.json
git commit -m "autogen: render config schema"
