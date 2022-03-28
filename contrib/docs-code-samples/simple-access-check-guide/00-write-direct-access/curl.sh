#!/bin/bash
set -euo pipefail

relationtuple='
{
  "namespace": "messages",
  "object": "02y_15_4w350m3",
  "relation": "decypher",
  "subject_id": "john"
}'

curl --fail --silent -X PUT \
     --data "$relationtuple" \
     http://127.0.0.1:4467/admin/relation-tuples > /dev/null \
  && echo "Successfully created tuple" \
  || echo "Encountered error"
