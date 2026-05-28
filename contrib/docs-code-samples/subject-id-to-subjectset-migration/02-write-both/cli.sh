#!/bin/bash
set -euo pipefail

export KETO_WRITE_REMOTE="127.0.0.1:4467"

# During the migration phase, write every new tuple as both SubjectID and SubjectSet.
printf '%s\n' \
  'File:data.txt#viewer@user_charlie' \
  'File:data.txt#viewer@User:charlie' | \
  keto relation-tuple parse -f - --format json | \
  keto relation-tuple create -f - >/dev/null --insecure-disable-transport-security \
    && echo "Successfully created tuples" \
    || echo "Encountered error"
