#!/bin/bash
set -euo pipefail

export KETO_WRITE_REMOTE="127.0.0.1:4467"

echo '// ownership
directories:/photos#owner@maureen
files:/photos/beach.jpg#owner@maureen
files:/photos/mountains.jpg#owner@laura

// maureen granted access to /photos to laura
directories:/photos#access@laura

// the following tuples are defined implicitly through subject set rewrites (not supported yet)
directories:/photos#access@(directories:/photos#owner)
files:/photos/beach.jpg#access@(files:/photos/beach.jpg#owner)
files:/photos/beach.jpg#access@(directories:/photos#access)
files:/photos/mountains.jpg#access@(files:/photos/mountains.jpg#owner)
files:/photos/mountains.jpg#access@(directories:/photos#access)' | \
  keto relation-tuple parse - --format json | \
    keto relation-tuple create --insecure-disable-transport-security - >/dev/null \
    && echo "Successfully created tuples" \
    || echo "Encountered error"
