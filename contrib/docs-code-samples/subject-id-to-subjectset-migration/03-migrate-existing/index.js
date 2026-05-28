// Copyright © 2026 Ory Corp
// SPDX-License-Identifier: Apache-2.0

import grpc from "@ory/keto-grpc-client/node_modules/@grpc/grpc-js/build/src/index.js"
import {
  relationTuples,
  read,
  readService,
  write,
  writeService,
} from "@ory/keto-grpc-client"

const BATCH_SIZE = 100

const readClient = new readService.ReadServiceClient(
  "127.0.0.1:4466",
  grpc.credentials.createInsecure(),
)
const writeClient = new writeService.WriteServiceClient(
  "127.0.0.1:4467",
  grpc.credentials.createInsecure(),
)

function subjectIDToNamespacedSubject(id) {
  if (id.startsWith("user_"))
    return { namespace: "User", object: id.slice("user_".length) }
  if (id.startsWith("apikey_"))
    return { namespace: "ApiKey", object: id.slice("apikey_".length) }
  throw new Error("unknown subject ID prefix: " + id)
}

function listRelationTuples(req) {
  return new Promise((resolve, reject) => {
    readClient.listRelationTuples(req, (error, resp) => {
      if (error) reject(error)
      else resolve(resp)
    })
  })
}

function transactRelationTuples(req) {
  return new Promise((resolve, reject) => {
    writeClient.transactRelationTuples(req, (error, resp) => {
      if (error) reject(error)
      else resolve(resp)
    })
  })
}

// Step 1: Paginate all File#viewer tuples and collect those with a SubjectID.
// In production, persist these to a file or database before writing back,
// so the migration can be resumed if interrupted.
const toMigrate = []
let pageToken = ""
do {
  const req = new read.ListRelationTuplesRequest()
  const query = new read.ListRelationTuplesRequest.Query()
  query.setNamespace("File")
  query.setRelation("viewer")
  req.setQuery(query)
  req.setPageSize(BATCH_SIZE)
  if (pageToken) req.setPageToken(pageToken)

  const resp = await listRelationTuples(req)
  for (const t of resp.getRelationTuplesList()) {
    if (t.getSubject().getId()) {
      toMigrate.push(t)
    }
  }
  pageToken = resp.getNextPageToken()
} while (pageToken)

// Step 2: Write SubjectSet counterparts in batches.
for (let i = 0; i < toMigrate.length; i += BATCH_SIZE) {
  const batch = toMigrate.slice(i, i + BATCH_SIZE)
  const req = new write.TransactRelationTuplesRequest()

  for (const t of batch) {
    const { namespace, object } = subjectIDToNamespacedSubject(
      t.getSubject().getId(),
    )

    const ss = new relationTuples.SubjectSet()
    ss.setNamespace(namespace).setObject(object).setRelation("")
    const sub = new relationTuples.Subject()
    sub.setSet(ss)

    const rt = new relationTuples.RelationTuple()
    rt.setNamespace(t.getNamespace())
      .setObject(t.getObject())
      .setRelation(t.getRelation())
      .setSubject(sub)

    const delta = new write.RelationTupleDelta()
    delta.setAction(write.RelationTupleDelta.Action.ACTION_INSERT)
    delta.setRelationTuple(rt)
    req.addRelationTupleDeltas(delta)
  }

  await transactRelationTuples(req)
}

console.log("Migration complete")
