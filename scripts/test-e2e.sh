#!/usr/bin/env bash

set -euo pipefail

cd "$( dirname "${BASH_SOURCE[0]}" )/.."

DATABASE_URL=memory keto serve --disable-telemetry &
while ! echo exit | nc 127.0.0.1 4466; do sleep 1; done

keto engines --endpoint http://localhost:4466 acp ory policies import regex ./tests/stubs/policies.json
keto engines --endpoint http://localhost:4466 acp ory policies import exact ./tests/stubs/policies.json

keto engines --endpoint http://localhost:4466 acp ory roles import regex ./tests/stubs/roles.json
keto engines --endpoint http://localhost:4466 acp ory roles import exact ./tests/stubs/roles.json

exit $(keto engines --endpoint http://localhost:4466 acp ory allowed regex peter-1 resources-11 actions-11 | grep -c  '"allowed": false')
exit $(keto engines --endpoint http://localhost:4466 acp ory allowed regex maria-1 resources-11 actions-11 | grep -c  '"allowed": false')
exit $(keto engines --endpoint http://localhost:4466 acp ory allowed regex group-1 resources-11 actions-11 | grep -c  '"allowed": false')

exit $(keto engines --endpoint http://localhost:4466 acp ory allowed regex not-exist resources-11 actions-11 | grep -c  '"allowed": true')

exit $(keto engines --endpoint http://localhost:4466 acp ory allowed exact peter-1 resources-11 actions-11 | grep -c  '"allowed": false')
exit $(keto engines --endpoint http://localhost:4466 acp ory allowed exact maria-1 resources-11 actions-11 | grep -c  '"allowed": false')
exit $(keto engines --endpoint http://localhost:4466 acp ory allowed exact group-1 resources-11 actions-11 | grep -c  '"allowed": false')

exit $(keto engines --endpoint http://localhost:4466 acp ory allowed exact not-exist resources-11 actions-11 | grep -c  '"allowed": true')
