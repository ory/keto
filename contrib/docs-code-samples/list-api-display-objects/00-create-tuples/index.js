// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

import { createPromiseClient } from "@connectrpc/connect"
import { createGrpcTransport } from "@connectrpc/connect-node"
import { RelationTupleDelta_Action, WriteService } from "@ory/keto-grpc-client"

const transport = createGrpcTransport({
  baseUrl: "http://127.0.0.1:4467",
  httpVersion: "2",
  interceptors: [],
  nodeOptions: {
    rejectUnauthorized: false,
  },
})

const writeClient = createPromiseClient(WriteService, transport)

const relationTupleDeltas = []

const addToChat = (chatName) => (user) => {
  const relationTuple = {
    namespace: "chats",
    object: chatName,
    relation: "member",
    subject: {
      id: user,
    },
  }

  const tupleDelta = {
    action: RelationTupleDelta_Action.ACTION_INSERT,
    relationTuple,
  }

  relationTupleDeltas.push(tupleDelta)
}

;["PM", "Vincent", "Julia"].forEach(addToChat("memes"))
;["PM", "Julia"].forEach(addToChat("cars"))
;["PM", "Vincent", "Julia", "Patrik"].forEach(addToChat("coffee-break"))

writeClient.transactRelationTuples(
  {
    relationTupleDeltas,
  },
  (error) => {
    if (error) {
      console.log("Encountered error", error)
    } else {
      console.log("Successfully created tuples")
    }
  },
)
