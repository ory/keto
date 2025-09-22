#!/bin/bash

set -euxo pipefail

schema_version="$(git rev-parse --short HEAD)"

sed "s!ory://tracing-config!https://raw.githubusercontent.com/ory/keto/$schema_version/oryx/otelx/config.schema.json!g;
s!ory://logging-config!https://raw.githubusercontent.com/ory/keto/$schema_version/oryx/logrusx/config.schema.json!g" embedx/config.schema.json > .schema/config.schema.json

git add .schema/config.schema.json

if ! git diff --exit-code .schema/config.schema.json
then
  git commit -m "autogen: render config schema"
  git push
fi
