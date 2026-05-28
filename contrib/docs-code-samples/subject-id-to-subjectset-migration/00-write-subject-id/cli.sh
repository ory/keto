#!/bin/bash
set -euo pipefail

export KETO_WRITE_REMOTE="127.0.0.1:4467"

printf '%s\n' \
  'File:data.txt#viewer@user_alice' \
  'File:data.txt#viewer@user_bob' \
  'File:data.txt#viewer@apikey_ci-bot' | \
  keto relation-tuple parse -f - --format json | \
  keto relation-tuple create -f - >/dev/null --insecure-disable-transport-security \
    && echo "Successfully created tuples" \
    || echo "Encountered error"
