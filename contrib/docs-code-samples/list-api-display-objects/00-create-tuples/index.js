// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

import grpc from "@ory/keto-grpc-client/node_modules/@grpc/grpc-js/build/src/index.js"
import { relationTuples, write, writeService } from "@ory/keto-grpc-client"

const writeClient = new writeService.WriteServiceClient(
  "127.0.0.1:4467",
  grpc.credentials.createInsecure(),
)

const writeRequest = new write.TransactRelationTuplesRequest()

const addToChat = (chatName) => (user) => {
  const relationTuple = new relationTuples.RelationTuple()
  relationTuple.setNamespace("chats")
  relationTuple.setObject(chatName)
  relationTuple.setRelation("member")

  const sub = new relationTuples.Subject()
  sub.setId(user)
  relationTuple.setSubject(sub)

  const tupleDelta = new write.RelationTupleDelta()
  tupleDelta.setAction(write.RelationTupleDelta.Action.ACTION_INSERT)
  tupleDelta.setRelationTuple(relationTuple)

  writeRequest.addRelationTupleDeltas(tupleDelta)
}

;["PM", "Vincent", "Julia"].forEach(addToChat("memes"))
;["PM", "Julia"].forEach(addToChat("cars"))
;["PM", "Vincent", "Julia", "Patrik"].forEach(addToChat("coffee-break"))

writeClient.transactRelationTuples(writeRequest, (error) => {
  if (error) {
    console.log("Encountered error", error)
  } else {
    console.log("Successfully created tuples")
  }
})
