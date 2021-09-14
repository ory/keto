#!/bin/bash
set -euxo pipefail

rm migrate_e2e.sqlite || true

bash <(curl https://raw.githubusercontent.com/ory/keto/master/install.sh) -b . v0.6.0-alpha.3

export DSN="sqlite://./migrate_e2e.sqlite?_fk=true"
export KETO_READ_REMOTE="127.0.0.1:4466"
export KETO_WRITE_REMOTE="127.0.0.1:4467"

./keto migrate up -y -c ./contrib/cat-videos-example/keto.yml
./keto namespace migrate up -y -c ./contrib/cat-videos-example/keto.yml videos

./keto serve all -c ./contrib/cat-videos-example/keto.yml &
keto_server_pid=$!

function teardown() {
    kill $keto_server_pid || true
}
trap teardown EXIT

jq '[range(200)] | map({namespace: "videos", object: . | tostring, relation: "view", subject: "user"})' <(echo '{}') \
  | ./keto relation-tuple create -q -

kill $keto_server_pid

go build -tags sqlite -o keto_new .

./keto_new migrate up -y -c ./contrib/cat-videos-example/keto.yml
./keto_new namespace migrate legacy -y -c ./contrib/cat-videos-example/keto.yml videos

./keto_new serve all -c ./contrib/cat-videos-example/keto.yml &
keto_server_pid=$!

for i in {1..200..25} ; do
    diff <(echo 'Allowed') <(./keto_new check user view videos "$i")
done
