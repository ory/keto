#!/bin/bash
set -euo pipefail

curl -G --silent \
     --data-urlencode "namespace=chats" \
     --data-urlencode "relation=member" \
     --data-urlencode "subject=PM" \
     http://127.0.0.1:4466/relationtuple | \
  jq ".relation_tuples[] | .object" -r
