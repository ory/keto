#!/bin/bash
set -euo pipefail

relationtuples='chats:memes#member@PM
chats:memes#member@Vincent
chats:memes#member@Julia
chats:cars#member@PM
chats:cars#member@Julia
chats:coffee-break#member@PM
chats:coffee-break#member@Vincent
chats:coffee-break#member@Julia
chats:coffee-break#member@Patrik'

for tuple in $relationtuples; do
  curl --fail --silent -X PUT \
       --data "$(echo "$tuple" | keto relation-tuple parse - --format json)" \
       http://127.0.0.1:4467/relationtuple > /dev/null
done

echo "Successfully created tuples"
