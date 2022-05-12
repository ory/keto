#!/bin/bash
set -euo pipefail

curl -G --silent \
     --data-urlencode "namespace=chats" \
     --data-urlencode "relation=member" \
     --data-urlencode "subject_id=PM" \
     http://127.0.0.1:4466/relation-tuples | \
  jq ".relation_tuples[] | .object" -r | sort
