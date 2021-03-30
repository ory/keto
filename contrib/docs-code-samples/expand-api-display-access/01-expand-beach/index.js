import grpc from '@ory/keto-grpc-client/node_modules/@grpc/grpc-js/build/src/index.js'
import expandService from '@ory/keto-grpc-client/expand_service_grpc_pb.js'
import expandData from '@ory/keto-grpc-client/expand_service_pb.js'
import acl from '@ory/keto-grpc-client/acl_pb.js'

const expandClient = new expandService.ExpandServiceClient(
  '127.0.0.1:4466',
  grpc.credentials.createInsecure()
)

const subjectSet = new acl.SubjectSet()
subjectSet.setNamespace('files')
subjectSet.setRelation('access')
subjectSet.setObject('/photos/beach.jpg')

const sub = new acl.Subject()
sub.setSet(subjectSet)

const expandRequest = new expandData.ExpandRequest()
expandRequest.setSubject(sub)
expandRequest.setMaxDepth(3)

// helper to get a nice result
const subjectString = (subject) => {
  if (subject.hasId()) {
    return subject.getId()
  }
  const set = subject.getSet()
  return set.getNamespace() + ':' + set.getObject() + '#' + set.getRelation()
}

// helper to get a nice result
const prettyTree = (tree) => {
  const [nodeType, subject, children] = [
    tree.getNodeType(),
    subjectString(tree.getSubject()),
    tree.getChildrenList()
  ]
  switch (nodeType) {
    case expandData.NodeType.NODE_TYPE_LEAF:
      return { type: 'leaf', subject }
    case expandData.NodeType.NODE_TYPE_UNION:
      return { type: 'union', subject, children: children.map(prettyTree) }
  }
}

expandClient.expand(expandRequest, (error, resp) => {
  if (error) {
    console.log('Encountered error:', error)
  } else {
    console.log(JSON.stringify(prettyTree(resp.getTree()), null, 2))
  }
})
