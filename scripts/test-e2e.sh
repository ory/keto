#!/usr/bin/env bash

set -euo pipefail

cd "$( dirname "${BASH_SOURCE[0]}" )/.."

DATABASE_URL=memory hades host --dangerous-auto-logon --dangerous-force-http --disable-telemetry &
while ! echo exit | nc 127.0.0.1 4444; do sleep 1; done

hades clients create --id foobar
hades clients delete foobar
curl --header "Authorization: bearer $(hades token client)" http://localhost:4444/clients
hades token validate $(hades token client)
