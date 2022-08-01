#!/bin/bash
set -euo pipefail

(cd ../..; go install -tags sqlite .)
PATH="../../.bin/gobin:../../.bin/brew/bin:../../.bin/brew/sbin:$PATH"

function teardown() {
    kill "$keto_server_pid" || true
}
trap teardown EXIT

function compare() {
    local actual="$1"
    local expected="$2/expected_output.json"
    local d="jd -set"
    if [[ ! -f "$expected" ]]; then
        expected="$2/expected_output.txt"
        d="diff -U 100000"
    fi
    $d "$actual" "$expected"
}

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
        compare <(go run -tags docscodesamples "./$main" 2>&1) "$(dirname "$main")"
    done

    for index in "$suite"/*/index.js; do
        echo "Running $index"
        compare <(node "$index" 2>&1) "$(dirname "$index")"
    done

    for tc in "$suite"/*/curl.sh; do
        echo "Running $tc"
        compare <("$tc" 2>&1) "$(dirname "$tc")"
    done

    for tc in "$suite"/*/cli.sh; do
        echo "Running $tc"
        compare <("$tc" 2>&1) "$(dirname "$tc")"
    done

    kill "$keto_server_pid"
done

echo
echo "EVERYTHING PASSED"
echo
