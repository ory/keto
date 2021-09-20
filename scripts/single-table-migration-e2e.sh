#!/bin/bash
set -euxo pipefail

rm migrate_e2e.sqlite || true

bash <(curl https://raw.githubusercontent.com/ory/keto/master/install.sh) -b . v0.6.0-alpha.3

export DSN="sqlite://./migrate_e2e.sqlite?_fk=true"
export KETO_READ_REMOTE="127.0.0.1:4466"
export KETO_WRITE_REMOTE="127.0.0.1:4467"

config="$(mktemp --tmpdir keto.XXXXXX.yml)"
echo "
log:
  level: debug

namespaces:
  - id: 0
    name: a
  - id: 1
    name: b
" >> "$config"

./keto migrate up -y -c "$config" --all-namespaces

./keto serve all -c "$config" &
keto_server_pid=$!

function teardown() {
    kill $keto_server_pid || true
}
trap teardown EXIT

jq '[range(300)] | map({namespace: (if . % 2 == 0 then "a" else "b" end), object: . | tostring, relation: "view", subject: "user"})' <(echo '{}') \
  | ./keto relation-tuple create -q -

kill $keto_server_pid

go build -tags sqlite -o keto_new .

./keto_new migrate up -y -c "$config"
./keto_new namespace migrate legacy -y -c "$config"

./keto_new serve all -c "$config" &
keto_server_pid=$!

for i in {0..300..24} ; do
    diff <(echo 'Allowed') <(./keto_new check user view a "$i")
done

for i in {1..300..24} ; do
    diff <(echo 'Allowed') <(./keto_new check user view b "$i")
done
