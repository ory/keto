#!/bin/bash
set -euo pipefail

./keto serve -c contrib/cat-videos-example/keto.yml &
keto_server_pid=$!

function teardown() {
    kill $keto_server_pid || true
}
trap teardown EXIT

export KETO_GRPC_URL="127.0.0.1:4467"

./keto relation-tuple create contrib/cat-videos-example/relation-tuples

echo "

Created all relation tuples. Now you can use the Keto CLI client like so:

export KETO_GRPC_URL=\"127.0.0.1:4467\"
./keto relation-tuple get videos
./keto check \"*\" view videos /cats/1.mp4
./keto expand view videos /cats/2.mp4
"

sleep 10d
