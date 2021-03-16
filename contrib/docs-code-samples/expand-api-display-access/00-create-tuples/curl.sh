#!/bin/bash
set -euo pipefail

relationtuples='directories:/photos#owner@maureen
files:/photos/beach.jpg#owner@maureen
files:/photos/mountains.jpg#owner@laura
directories:/photos#access@laura
directories:/photos#access@(directories:/photos#owner)
files:/photos/beach.jpg#access@(files:/photos/beach.jpg#owner)
files:/photos/beach.jpg#access@(directories:/photos#access)
files:/photos/mountains.jpg#access@(files:/photos/mountains.jpg#owner)
files:/photos/mountains.jpg#access@(directories:/photos#access)'

for tuple in $relationtuples; do
  curl --fail --silent -X PUT \
       --data "$(echo "$tuple" | keto relation-tuple parse - --format json)" \
       http://127.0.0.1:4467/relationtuple > /dev/null
done

echo "Successfully created tuples"
