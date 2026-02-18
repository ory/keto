#!/bin/bash
set -euo pipefail

echo 'Chat:memes#member@User:PM
Chat:memes#member@User:Vincent
Chat:memes#member@User:Julia

Chat:cars#member@User:PM
Chat:cars#member@User:Julia

Chat:coffee-break#member@User:PM
Chat:coffee-break#member@User:Vincent
Chat:coffee-break#member@User:Julia
Chat:coffee-break#member@User:Patrik' | \
  keto relation-tuple parse -f - --format json | \
    jq "[ .[] | { relation_tuple: . , action: \"insert\" } ]" -c | \
      curl -X PATCH --silent --fail \
        -H 'Content-Type: application/json' \
        --retry 7 --retry-connrefused \
        --data @- \
        http://127.0.0.1:4467/admin/relation-tuples

echo "Successfully created tuples"
