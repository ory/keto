#!/bin/bash
set -euo pipefail

curl -X PATCH --fail --silent \
     -H 'Content-Type: application/json' \
     --retry 7 --retry-connrefused \
     --data '[
       {"action":"delete","relation_tuple":{"namespace":"File","object":"data.txt","relation":"viewer","subject_id":"user_alice"}},
       {"action":"delete","relation_tuple":{"namespace":"File","object":"data.txt","relation":"viewer","subject_id":"user_bob"}},
       {"action":"delete","relation_tuple":{"namespace":"File","object":"data.txt","relation":"viewer","subject_id":"user_charlie"}},
       {"action":"delete","relation_tuple":{"namespace":"File","object":"data.txt","relation":"viewer","subject_id":"apikey_ci-bot"}}
     ]' \
     http://127.0.0.1:4467/admin/relation-tuples > /dev/null \
  && echo "Successfully deleted SubjectID tuples" \
  || echo "Encountered error"
