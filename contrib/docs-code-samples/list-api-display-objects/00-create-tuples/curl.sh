#!/bin/bash
set -euo pipefail

echo 'chats:memes#member@PM
chats:memes#member@Vincent
chats:memes#member@Julia

chats:cars#member@PM
chats:cars#member@Julia

chats:coffee-break#member@PM
chats:coffee-break#member@Vincent
chats:coffee-break#member@Julia
chats:coffee-break#member@Patrik' | \
  keto relation-tuple parse - --format json | \
    jq "[ .[] | { relation_tuple: . , action: \"insert\" } ]" -c | \
      curl -X PATCH --silent --fail \
        --data @- \
        http://127.0.0.1:4467/admin/relation-tuples

echo "Successfully created tuples"
