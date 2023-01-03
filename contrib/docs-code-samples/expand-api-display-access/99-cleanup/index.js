// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

import grpc from "@ory/keto-grpc-client/node_modules/@grpc/grpc-js/build/src/index.js"
import { relationTuples, write, writeService } from "@ory/keto-grpc-client"

const purgeNamespace = (namespace) => {
  const query = new relationTuples.RelationQuery()
  query.setNamespace(namespace)
  const request = new write.DeleteRelationTuplesRequest()
  request.setRelationQuery(query)
  new writeService.WriteServiceClient(
    "127.0.0.1:4467",
    grpc.credentials.createInsecure(),
  ).deleteRelationTuples(request, (err) => {
    if (err) {
      console.error(err)
    }
  })
}

purgeNamespace("files")
purgeNamespace("directories")
