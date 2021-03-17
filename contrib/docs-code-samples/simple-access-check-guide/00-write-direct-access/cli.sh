#!/bin/bash
set -euo pipefail

echo "messages:02y_15_4w350m3#decypher@john" | \
  keto relation-tuple parse - --format json | \
  keto relation-tuple create - >/dev/null \
    && echo "Successfully created tuple" \
    || echo "Encountered error"
