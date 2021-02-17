#!/bin/bash
set -euo pipefail

testdir="contrib/docs-code-samples"

go install -tags sqlite .

function teardown() {
    kill "$keto_server_pid" || true
}
trap teardown EXIT

for suite in "$testdir"/*/ ; do
    suite=$(basename "$suite")
    if [ "$suite" == "node_modules" ]; then
        continue
    fi

    keto serve -c "$testdir/$suite/keto.yml" &> "$testdir/serve_$suite.log" &
    keto_server_pid=$!

    until curl --output /dev/null --silent --fail http://127.0.0.1:4466/health/ready; do
        printf '.'
        sleep 0.2
    done
    echo

    go test -tags docscodesamples -count=1 "./$testdir/$suite/..."
    echo

    node --experimental-vm-modules "$testdir/node_modules/.bin/jest" "$testdir/$suite"
    echo

    for tc in "$testdir"/"$suite"/*/curl_test.sh; do
        bash "$tc"
    done
    echo

    for tc in "$testdir"/"$suite"/*/cli_test.sh; do
        bash "$tc"
    done

    kill $keto_server_pid
done

echo
echo "EVERYTHING PASSED"
echo
