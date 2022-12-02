#!/bin/bash
set -euo pipefail

keto serve -c contrib/cat-videos-example/keto.yml &
keto_server_pid=$!

function teardown() {
    kill $keto_server_pid || true
}
trap teardown EXIT

export KETO_WRITE_REMOTE="127.0.0.1:4467"

keto relation-tuple create contrib/cat-videos-example/relation-tuples --insecure-disable-transport-security

echo "

Created all relationships. Now you can use the Keto CLI client to play around:

export KETO_READ_REMOTE=\"127.0.0.1:4466\"
keto relation-tuple get --insecure-disable-transport-security
keto check \"*\" view videos /cats/1.mp4 --insecure-disable-transport-security
keto expand view videos /cats/2.mp4 --insecure-disable-transport-security
"

# sleep 10h; has to be defined like this because OSX does not know units https://www.unix.com/man-page/osx/1/sleep/
sleep 36000
