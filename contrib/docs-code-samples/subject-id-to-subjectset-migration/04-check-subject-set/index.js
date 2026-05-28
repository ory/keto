// Copyright © 2026 Ory Corp
// SPDX-License-Identifier: Apache-2.0

import grpc from "@ory/keto-grpc-client/node_modules/@grpc/grpc-js/build/src/index.js"
import { relationTuples, check, checkService } from "@ory/keto-grpc-client"

const checkClient = new checkService.CheckServiceClient(
  "127.0.0.1:4466",
  grpc.credentials.createInsecure(),
)

const ss = new relationTuples.SubjectSet()
ss.setNamespace("User")
ss.setObject("alice")
ss.setRelation("")
const sub = new relationTuples.Subject()
sub.setSet(ss)

const req = new check.CheckRequest()
req.setNamespace("File")
req.setObject("data.txt")
req.setRelation("viewer")
req.setSubject(sub)

checkClient.check(req, (error, resp) => {
  if (error) {
    console.log("Encountered error:", error)
  } else {
    console.log(resp.getAllowed() ? "Allowed" : "Denied")
  }
})
