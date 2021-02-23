#!/bin/bash
set -euo pipefail

relationtuple='
{
  "namespace": "messages",
  "object": "02y_15_4w350m3",
  "relation": "decypher",
  "subject": "john"
}'

keto relation-tuple delete <(echo "$relationtuple") -q > /dev/null
