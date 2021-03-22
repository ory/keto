#!/bin/bash
set -euo pipefail

(cd ../..; go install -tags sqlite .)

function teardown() {
    kill "$keto_server_pid" || true
}
trap teardown EXIT

for suite in */ ; do
    suite=$(basename "$suite")
    if [ "$suite" == "node_modules" ]; then
        continue
    fi

    keto serve -c "$suite/keto.yml" &> "serve_$suite.log" &
    keto_server_pid=$!

    until curl --output /dev/null --silent --fail http://127.0.0.1:4466/health/ready; do
        printf '.'
        sleep 0.2
    done
    echo

    for main in "$suite"/*/main.go; do
        echo "Running $main"
        diff <(go run -tags docscodesamples "./$main" 2>&1) "$(dirname "$main")/expected_output.txt"
    done

    for index in "$suite"/*/index.js; do
        echo "Running $index"
        diff <(node "$index" 2>&1) "$(dirname "$index")/expected_output.txt"
    done

    for tc in "$suite"/*/curl.sh; do
        echo "Running $tc"
        diff <("$tc" 2>&1) "$(dirname "$tc")/expected_output.txt"
    done

    for tc in "$suite"/*/cli.sh; do
        echo "Running $tc"
        diff <("$tc" 2>&1) "$(dirname "$tc")/expected_output.txt"
    done

    kill "$keto_server_pid"
done

echo
echo "EVERYTHING PASSED"
echo
