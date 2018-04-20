#!/usr/bin/env bash

set -euo pipefail

cd "$( dirname "${BASH_SOURCE[0]}" )/.."

DATABASE_URL=memory keto serve --dangerous-auto-logon --dangerous-force-http --disable-telemetry &
while ! echo exit | nc 127.0.0.1 4444; do sleep 1; done

keto clients create --id foobar
keto clients delete foobar
curl --header "Authorization: bearer $(keto token client)" http://localhost:4444/clients
keto token validate $(keto token client)
