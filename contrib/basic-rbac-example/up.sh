#!/bin/bash
set -euo pipefail

keto serve -c contrib/basic-rbac-example/keto.yml &
keto_server_pid=$!

function teardown() {
    kill $keto_server_pid || true
}
trap teardown EXIT

export KETO_WRITE_REMOTE="127.0.0.1:4467"

keto relation-tuple create contrib/basic-rbac-example/relation-tuples

echo "
Created all relation tuples. Now you can use the Keto CLI client to play around:
export KETO_READ_REMOTE=\"127.0.0.1:4466\"
keto relation-tuple get --namespace docs
keto check user-1 view docs readme
keto expand view docs readme
"

# sleep 10h; has to be defined like this because OSX does not know units https://www.unix.com/man-page/osx/1/sleep/
sleep 36000
