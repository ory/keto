// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

import grpc from "@ory/keto-grpc-client/node_modules/@grpc/grpc-js/build/src/index.js"
import { read, readService } from "@ory/keto-grpc-client"

const readClient = new readService.ReadServiceClient(
  "127.0.0.1:4466",
  grpc.credentials.createInsecure(),
)

const readRequest = new read.ListRelationTuplesRequest()
const query = new read.ListRelationTuplesRequest.Query()
query.setNamespace("chats")
query.setObject("coffee-break")
query.setRelation("member")

readRequest.setQuery(query)

readClient.listRelationTuples(readRequest, (error, resp) => {
  if (error) {
    console.log("Encountered error:", error)
  } else {
    console.log(
      resp
        .getRelationTuplesList()
        .map((tuple) => tuple.getSubject().getId())
        .sort((a, b) => (a < b ? -1 : 1))
        .join("\n"),
    )
  }
})
