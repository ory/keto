#!/bin/bash
set -euo pipefail

export KETO_WRITE_REMOTE="127.0.0.1:4467"

relationtuple='
{
  "namespace": "messages",
  "object": "02y_15_4w350m3",
  "relation": "decypher",
  "subject_id": "john"
}'

keto relation-tuple delete <(echo "$relationtuple") -q > /dev/null --insecure-disable-transport-security
