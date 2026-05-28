// Copyright © 2026 Ory Corp
// SPDX-License-Identifier: Apache-2.0

import grpc from "@ory/keto-grpc-client/node_modules/@grpc/grpc-js/build/src/index.js"
import { relationTuples, write, writeService } from "@ory/keto-grpc-client"

const writeClient = new writeService.WriteServiceClient(
  "127.0.0.1:4467",
  grpc.credentials.createInsecure(),
)

function insertSubjectID(subjectID) {
  const rt = new relationTuples.RelationTuple()
  rt.setNamespace("File")
  rt.setObject("data.txt")
  rt.setRelation("viewer")
  const sub = new relationTuples.Subject()
  sub.setId(subjectID)
  rt.setSubject(sub)
  const delta = new write.RelationTupleDelta()
  delta.setAction(write.RelationTupleDelta.Action.ACTION_INSERT)
  delta.setRelationTuple(rt)
  return delta
}

const req = new write.TransactRelationTuplesRequest()
for (const id of ["user_alice", "user_bob", "apikey_ci-bot"]) {
  req.addRelationTupleDeltas(insertSubjectID(id))
}

writeClient.transactRelationTuples(req, (error) => {
  if (error) {
    console.log("Encountered error", error)
  } else {
    console.log("Successfully created tuples")
  }
})
