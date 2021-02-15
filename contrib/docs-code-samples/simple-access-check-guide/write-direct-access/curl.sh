#!/bin/bash
set -euo pipefail

relationtuple='
{
  "namespace": "messages",
  "object": "02y_15_4w350m3",
  "relation": "decypher",
  "subject": "john"
}'

curl -X PUT \
     --data "$relationtuple" \
     -w "%{http_code}" \
     http://127.0.0.1:4467/relationtuple \
  && echo -e "\nCreated!"
