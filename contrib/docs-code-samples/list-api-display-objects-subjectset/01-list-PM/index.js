// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

import grpc from "@ory/keto-grpc-client/node_modules/@grpc/grpc-js/build/src/index.js"
import { relationTuples, read, readService } from "@ory/keto-grpc-client"

const readClient = new readService.ReadServiceClient(
  "127.0.0.1:4466",
  grpc.credentials.createInsecure(),
)

const readRequest = new read.ListRelationTuplesRequest()
const query = new read.ListRelationTuplesRequest.Query()
query.setNamespace("Chat")
query.setRelation("member")

const subSet = new relationTuples.SubjectSet()
subSet.setNamespace("User")
subSet.setObject("PM")

const sub = new relationTuples.Subject()
sub.setSet(subSet)
query.setSubject(sub)

readRequest.setQuery(query)

readClient.listRelationTuples(readRequest, (error, resp) => {
  if (error) {
    console.log("Encountered error:", error)
  } else {
    console.log(
      resp
        .getRelationTuplesList()
        .map((tuple) => tuple.getObject())
        .sort((a, b) => (a < b ? -1 : 1))
        .join("\n"),
    )
  }
})
