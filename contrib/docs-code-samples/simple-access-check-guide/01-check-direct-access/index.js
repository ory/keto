// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

import grpc from "@ory/keto-grpc-client/node_modules/@grpc/grpc-js/build/src/index.js"
import { relationTuples, check, checkService } from "@ory/keto-grpc-client"

const checkClient = new checkService.CheckServiceClient(
  "127.0.0.1:4466",
  grpc.credentials.createInsecure(),
)

const checkRequest = new check.CheckRequest()
checkRequest.setNamespace("messages")
checkRequest.setObject("02y_15_4w350m3")
checkRequest.setRelation("decypher")

const sub = new relationTuples.Subject()
sub.setId("john")
checkRequest.setSubject(sub)

checkClient.check(checkRequest, (error, resp) => {
  if (error) {
    console.log("Encountered error:", error)
  } else {
    console.log(resp.getAllowed() ? "Allowed" : "Denied")
  }
})
