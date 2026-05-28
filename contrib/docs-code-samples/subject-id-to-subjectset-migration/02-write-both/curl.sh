#!/bin/bash
set -euo pipefail

# During migration, write every new tuple as both SubjectID and SubjectSet.
curl -X PATCH --fail --silent \
     -H 'Content-Type: application/json' \
     --retry 7 --retry-connrefused \
     --data '[
       {"action":"insert","relation_tuple":{"namespace":"File","object":"data.txt","relation":"viewer","subject_id":"user_charlie"}},
       {"action":"insert","relation_tuple":{"namespace":"File","object":"data.txt","relation":"viewer","subject_set":{"namespace":"User","object":"charlie","relation":""}}}
     ]' \
     http://127.0.0.1:4467/admin/relation-tuples > /dev/null \
  && echo "Successfully created tuples" \
  || echo "Encountered error"
