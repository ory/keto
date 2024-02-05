// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

import {
  RelationTuple,
  RelationTupleDelta,
  RelationTupleDelta_Action,
  Subject,
  SubjectSet,
  TransactRelationTuplesRequest,
  WriteService,
} from "@ory/keto-grpc-client"
import { createGrpcTransport } from "@connectrpc/connect-node"
import { createPromiseClient } from "@connectrpc/connect"

const transport = createGrpcTransport({
  baseUrl: "http://127.0.0.1:4466/",
  httpVersion: "1.1",
  interceptors: [],
  nodeOptions: {
    rejectUnauthorized: false,
  },
})

const writeClient = createPromiseClient(WriteService, transport)

const writeRequest = new TransactRelationTuplesRequest()

const insert = (tuple) => {
  const tupleDelta = new RelationTupleDelta({
    action: RelationTupleDelta_Action.ACTION_INSERT,
    relationTuple: tuple,
  })

  writeRequest.relationTupleDeltas.push(tupleDelta)
}

const addSimpleTuple = (namespace, object, relation, user) => {
  const sub = new Subject({
    id: user,
  })
  const relationTuple = new RelationTuple({
    namespace,
    object,
    relation,
    subject: sub,
  })

  insert(relationTuple)
}

// ownership
addSimpleTuple("directories", "/photos", "owner", "maureen")
addSimpleTuple("files", "/photos/beach.jpg", "owner", "maureen")
addSimpleTuple("files", "/photos/mountains.jpg", "owner", "laura")
// granted access
addSimpleTuple("directories", "/photos", "access", "laura")

// should be subject set rewrite
// owners have access
;[
  ["files", "/photos/beach.jpg"],
  ["files", "/photos/mountains.jpg"],
  ["directories", "/photos"],
].forEach(([namespace, object]) => {
  const subjectSet = new SubjectSet({
    namespace: "directories",
    object: "/photos",
    relation: "owner",
  })

  const sub = new Subject({
    set: subjectSet,
  })

  const relationTuple = new RelationTuple({
    namespace,
    object,
    relation: "access",
    subject: sub,
  })

  insert(relationTuple)
})

// should be subject set rewrite
// access on parent means access on child
;["/photos/beach.jpg", "/photos/mountains.jpg"].forEach((file) => {
  const subjectSet = new SubjectSet({
    namespace: "directories",
    object: "/photos",
    relation: "access",
  })

  const sub = new Subject({
    set: subjectSet,
  })

  const relationTuple = new RelationTuple({
    namespace: "files",
    object: file,
    relation: "access",
    subject: sub,
  })

  insert(relationTuple)
})

writeClient.transactRelationTuples(writeRequest, (error) => {
  if (error) {
    console.log("Encountered error", error)
  } else {
    console.log("Successfully created tuples")
  }
})
