// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

import grpc from "@ory/keto-grpc-client/node_modules/@grpc/grpc-js/build/src/index.js"
import { relationTuples, expand, expandService } from "@ory/keto-grpc-client"

const expandClient = new expandService.ExpandServiceClient(
  "127.0.0.1:4466",
  grpc.credentials.createInsecure(),
)

const subjectSet = new relationTuples.SubjectSet()
subjectSet.setNamespace("files")
subjectSet.setRelation("access")
subjectSet.setObject("/photos/beach.jpg")

const sub = new relationTuples.Subject()
sub.setSet(subjectSet)

const expandRequest = new expand.ExpandRequest()
expandRequest.setSubject(sub)
expandRequest.setMaxDepth(3)

// helper to get a nice result
const subjectJSON = (subject) => {
  if (subject.hasId()) {
    return { subject_id: subject.getId() }
  }
  const set = subject.getSet()
  return {
    subject_set: {
      namespace: set.getNamespace(),
      object: set.getObject(),
      relation: set.getRelation(),
    },
  }
}

// helper to get a nice result
const prettyTree = (tree) => {
  const [nodeType, tuple, children] = [
    tree.getNodeType(),
    {
      tuple: {
        namespace: "",
        object: "",
        relation: "",
        ...subjectJSON(tree.getSubject()),
      },
    },
    tree.getChildrenList(),
  ]
  switch (nodeType) {
    case expand.NodeType.NODE_TYPE_LEAF:
      return { type: "leaf", ...tuple }
    case expand.NodeType.NODE_TYPE_UNION:
      return { type: "union", children: children.map(prettyTree), ...tuple }
  }
}

expandClient.expand(expandRequest, (error, resp) => {
  if (error) {
    console.log("Encountered error:", error)
  } else {
    console.log(JSON.stringify(prettyTree(resp.getTree()), null, 2))
  }
})
