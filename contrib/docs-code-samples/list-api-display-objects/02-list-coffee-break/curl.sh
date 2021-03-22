#!/bin/bash
set -euo pipefail

curl -G --silent \
     --data-urlencode "namespace=chats" \
     --data-urlencode "object=coffee-break" \
     --data-urlencode "relation=member" \
     http://127.0.0.1:4466/relationtuple | \
  jq ".relation_tuples[] | .subject" -r
