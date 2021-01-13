#!/bin/bash
set -euxo pipefail

./keto serve -c contrib/small-example/keto.yml &
keto_server_pid=$!

function teardown() {
    kill $keto_server_pid || true
}
trap teardown EXIT

export KETO_GRPC_URL="127.0.0.1:4467"

for f in contrib/small-example/relation-tuples/*.json; do
  ./keto relation-tuple create "$f"
done

sleep 10d
