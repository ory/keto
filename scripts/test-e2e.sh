#!/usr/bin/env bash

set -euo pipefail

cd "$( dirname "${BASH_SOURCE[0]}" )/.."

killall keto || true

DATABASE_URL=memory keto serve --disable-telemetry &
while ! echo exit | nc 127.0.0.1 4466; do sleep 1; done

# Explicitly run without endpoint to see if that's working properly.
export KETO_URL=http://127.0.0.1:4466/
keto engines acp ory policies import regex ./tests/stubs/policies.json

# And check if it's working without trailing slash
export KETO_URL=http://127.0.0.1:4466
keto engines acp ory policies import exact ./tests/stubs/policies.json

# Now explicitly check if that works with the --endpoint flag
keto engines --endpoint http://localhost:4466 acp ory roles import regex ./tests/stubs/roles.json
# And with slash
keto engines --endpoint http://localhost:4466/ acp ory roles import exact ./tests/stubs/roles.json

# Importing data is done, let's perform some checks

exit $(keto engines --endpoint http://localhost:4466 acp ory allowed regex peter-1 resources-11 actions-11 | grep -c  '"allowed": false')
exit $(keto engines --endpoint http://localhost:4466 acp ory allowed regex maria-1 resources-11 actions-11 | grep -c  '"allowed": false')
exit $(keto engines --endpoint http://localhost:4466 acp ory allowed regex group-1 resources-11 actions-11 | grep -c  '"allowed": false')

exit $(keto engines --endpoint http://localhost:4466 acp ory allowed regex not-exist resources-11 actions-11 | grep -c  '"allowed": true')

exit $(keto engines --endpoint http://localhost:4466 acp ory allowed exact peter-1 resources-11 actions-11 | grep -c  '"allowed": false')
exit $(keto engines --endpoint http://localhost:4466 acp ory allowed exact maria-1 resources-11 actions-11 | grep -c  '"allowed": false')
exit $(keto engines --endpoint http://localhost:4466 acp ory allowed exact group-1 resources-11 actions-11 | grep -c  '"allowed": false')

exit $(keto engines --endpoint http://localhost:4466 acp ory allowed exact not-exist resources-11 actions-11 | grep -c  '"allowed": true')

kill %1
exit 0
