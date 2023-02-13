// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

import grpc from "@ory/keto-grpc-client/node_modules/@grpc/grpc-js/build/src/index.js"
import { relationTuples, write, writeService } from "@ory/keto-grpc-client"

const writeClient = new writeService.WriteServiceClient(
  "127.0.0.1:4467",
  grpc.credentials.createInsecure(),
)

const writeRequest = new write.TransactRelationTuplesRequest()

const insert = (tuple) => {
  const tupleDelta = new write.RelationTupleDelta()
  tupleDelta.setAction(write.RelationTupleDelta.Action.ACTION_INSERT)
  tupleDelta.setRelationTuple(tuple)

  writeRequest.addRelationTupleDeltas(tupleDelta)
}

const addSimpleTuple = (namespace, object, relation, user) => {
  const relationTuple = new relationTuples.RelationTuple()
  relationTuple.setNamespace(namespace)
  relationTuple.setObject(object)
  relationTuple.setRelation(relation)

  const sub = new relationTuples.Subject()
  sub.setId(user)
  relationTuple.setSubject(sub)

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
  const relationTuple = new relationTuples.RelationTuple()
  relationTuple.setNamespace(namespace)
  relationTuple.setObject(object)
  relationTuple.setRelation("access")

  const subjectSet = new relationTuples.SubjectSet()
  subjectSet.setNamespace(namespace)
  subjectSet.setObject(object)
  subjectSet.setRelation("owner")

  const sub = new relationTuples.Subject()
  sub.setSet(subjectSet)
  relationTuple.setSubject(sub)

  insert(relationTuple)
})

// should be subject set rewrite
// access on parent means access on child
;["/photos/beach.jpg", "/photos/mountains.jpg"].forEach((file) => {
  const relationTuple = new relationTuples.RelationTuple()
  relationTuple.setNamespace("files")
  relationTuple.setObject(file)
  relationTuple.setRelation("access")

  const subjectSet = new relationTuples.SubjectSet()
  subjectSet.setNamespace("directories")
  subjectSet.setObject("/photos")
  subjectSet.setRelation("access")

  const sub = new relationTuples.Subject()
  sub.setSet(subjectSet)
  relationTuple.setSubject(sub)

  insert(relationTuple)
})

writeClient.transactRelationTuples(writeRequest, (error) => {
  if (error) {
    console.log("Encountered error", error)
  } else {
    console.log("Successfully created tuples")
  }
})
