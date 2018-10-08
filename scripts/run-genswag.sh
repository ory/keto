#!/bin/bash

set -euo pipefail

cd "$( dirname "${BASH_SOURCE[0]}" )/.."

go build github.com/ory/keto
swagger generate spec -m -o ./docs/api.swagger.json
