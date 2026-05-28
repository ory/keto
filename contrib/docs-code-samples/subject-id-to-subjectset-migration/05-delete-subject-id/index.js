// Copyright © 2026 Ory Corp
// SPDX-License-Identifier: Apache-2.0

import grpc from "@ory/keto-grpc-client/node_modules/@grpc/grpc-js/build/src/index.js"
import { relationTuples, write, writeService } from "@ory/keto-grpc-client"

const writeClient = new writeService.WriteServiceClient(
  "127.0.0.1:4467",
  grpc.credentials.createInsecure(),
)

// subjectIdTuples are the SubjectID tuples collected and persisted during the migration step.
// In production, load this list from the file or database written during migration so that
// you can target deletion of the tuples that were already migrated, and resume if interrupted.
const subjectIdTuples = [
  {
    namespace: "File",
    object: "data.txt",
    relation: "viewer",
    subjectId: "user_alice",
  },
  {
    namespace: "File",
    object: "data.txt",
    relation: "viewer",
    subjectId: "user_bob",
  },
  {
    namespace: "File",
    object: "data.txt",
    relation: "viewer",
    subjectId: "user_charlie",
  },
  {
    namespace: "File",
    object: "data.txt",
    relation: "viewer",
    subjectId: "apikey_ci-bot",
  },
]

const req = new write.TransactRelationTuplesRequest()
for (const { namespace, object, relation, subjectId } of subjectIdTuples) {
  const sub = new relationTuples.Subject()
  sub.setId(subjectId)

  const rt = new relationTuples.RelationTuple()
  rt.setNamespace(namespace)
    .setObject(object)
    .setRelation(relation)
    .setSubject(sub)

  const delta = new write.RelationTupleDelta()
  delta.setAction(write.RelationTupleDelta.Action.ACTION_DELETE)
  delta.setRelationTuple(rt)
  req.addRelationTupleDeltas(delta)
}

writeClient.transactRelationTuples(req, (error) => {
  if (error) {
    console.log("Encountered error", error)
  } else {
    console.log("Successfully deleted SubjectID tuples")
  }
})
