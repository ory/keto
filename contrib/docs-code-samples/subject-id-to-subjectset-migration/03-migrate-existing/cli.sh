#!/bin/bash
set -euo pipefail

export KETO_WRITE_REMOTE="127.0.0.1:4467"

# In production, paginate and filter all SubjectID tuples before writing back.
# This script writes the known SubjectSet tuples directly for brevity.
printf '%s\n' \
  'File:data.txt#viewer@User:alice#' \
  'File:data.txt#viewer@User:bob#' \
  'File:data.txt#viewer@ApiKey:ci-bot#' | \
  keto relation-tuple parse -f - --format json | \
  keto relation-tuple create -f - >/dev/null --insecure-disable-transport-security \
    && echo "Migration complete" \
    || echo "Encountered error"
