// Copyright © 2026 Ory Corp
// SPDX-License-Identifier: Apache-2.0

import grpc from "@ory/keto-grpc-client/node_modules/@grpc/grpc-js/build/src/index.js"
import { relationTuples, write, writeService } from "@ory/keto-grpc-client"

const writeClient = new writeService.WriteServiceClient(
  "127.0.0.1:4467",
  grpc.credentials.createInsecure(),
)

// During migration, write every new tuple as both SubjectID and SubjectSet.
function makeDelta(subject) {
  const rt = new relationTuples.RelationTuple()
  rt.setNamespace("File")
  rt.setObject("data.txt")
  rt.setRelation("viewer")
  rt.setSubject(subject)
  const delta = new write.RelationTupleDelta()
  delta.setAction(write.RelationTupleDelta.Action.ACTION_INSERT)
  delta.setRelationTuple(rt)
  return delta
}

const subjectID = new relationTuples.Subject()
subjectID.setId("user_charlie")

const ss = new relationTuples.SubjectSet()
ss.setNamespace("User")
ss.setObject("charlie")
ss.setRelation("")
const subjectSet = new relationTuples.Subject()
subjectSet.setSet(ss)

const req = new write.TransactRelationTuplesRequest()
req.addRelationTupleDeltas(makeDelta(subjectID))
req.addRelationTupleDeltas(makeDelta(subjectSet))

writeClient.transactRelationTuples(req, (error) => {
  if (error) {
    console.log("Encountered error", error)
  } else {
    console.log("Successfully created tuples")
  }
})
