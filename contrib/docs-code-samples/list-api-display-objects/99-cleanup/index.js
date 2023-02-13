// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

import grpc from "@ory/keto-grpc-client/node_modules/@grpc/grpc-js/build/src/index.js"
import { write, writeService, read, readService } from "@ory/keto-grpc-client"

const readClient = new readService.ReadServiceClient(
  "127.0.0.1:4466",
  grpc.credentials.createInsecure(),
)

const query = new read.ListRelationTuplesRequest.Query()
query.setNamespace("chats")

const readRequest = new read.ListRelationTuplesRequest()
readRequest.setQuery(query)

readClient.listRelationTuples(readRequest, (err, resp) => {
  const writeClient = new writeService.WriteServiceClient(
    "127.0.0.1:4467",
    grpc.credentials.createInsecure(),
  )

  const writeRequest = new write.TransactRelationTuplesRequest()

  resp.getRelationTuplesList().forEach((tuple) => {
    const tupleDelta = new write.RelationTupleDelta()
    tupleDelta.setAction(write.RelationTupleDelta.Action.ACTION_DELETE)
    tupleDelta.setRelationTuple(tuple)
    writeRequest.addRelationTupleDeltas(tupleDelta)
  })

  writeClient.transactRelationTuples(writeRequest, (err) => {
    if (err) {
      console.log("Unexpected err", err)
    }
  })
})
