#!/bin/bash
set -euo pipefail

export KETO_WRITE_REMOTE="127.0.0.1:4467"

echo "messages:02y_15_4w350m3#decypher@john" | \
  keto relation-tuple parse - --format json | \
  keto relation-tuple create - >/dev/null --insecure-disable-transport-security \
    && echo "Successfully created tuple" \
    || echo "Encountered error"
