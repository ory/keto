#!/bin/bash
set -euxo pipefail

bash <(curl https://raw.githubusercontent.com/ory/keto/master/install.sh) -b . v0.6.0-alpha.3

export DSN=sqlite://./migrate_e2e.sqlite

./keto migrate up -y -c ./contrib/cat-videos-example/keto.yml
