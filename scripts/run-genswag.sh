#!/bin/bash

set -euo pipefail

cd "$( dirname "${BASH_SOURCE[0]}" )/.."

go get ./...
swagger generate spec -m -o ./docs/api.swagger.json
