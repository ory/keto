import grpc from '@ory/keto-grpc-client/node_modules/@grpc/grpc-js/build/src/index.js'
import acl from '@ory/keto-grpc-client/acl_pb.js'
import writeService from '@ory/keto-grpc-client/write_service_grpc_pb.js'
import writeData from '@ory/keto-grpc-client/write_service_pb.js'

const writeClient = new writeService.WriteServiceClient(
  '127.0.0.1:4467',
  grpc.credentials.createInsecure()
)

const writeRequest = new writeData.TransactRelationTuplesRequest()

const insert = (tuple) => {
  const tupleDelta = new writeData.RelationTupleDelta()
  tupleDelta.setAction(writeData.RelationTupleDelta.Action.INSERT)
  tupleDelta.setRelationTuple(tuple)

  writeRequest.addRelationTupleDeltas(tupleDelta)
}

const addSimpleTuple = (namespace, object, relation, user) => {
  const relationTuple = new acl.RelationTuple()
  relationTuple.setNamespace(namespace)
  relationTuple.setObject(object)
  relationTuple.setRelation(relation)

  const sub = new acl.Subject()
  sub.setId(user)
  relationTuple.setSubject(sub)

  insert(relationTuple)
}

// ownership
addSimpleTuple('directories', '/photos', 'owner', 'maureen')
addSimpleTuple('files', '/photos/beach.jpg', 'owner', 'maureen')
addSimpleTuple('files', '/photos/mountains.jpg', 'owner', 'laura')
// granted access
addSimpleTuple('directories', '/photos', 'access', 'laura')

// should be subject set rewrite
// owners have access
;[
  ['files', '/photos/beach.jpg'],
  ['files', '/photos/mountains.jpg'],
  ['directories', '/photos']
].forEach(([namespace, object]) => {
  const relationTuple = new acl.RelationTuple()
  relationTuple.setNamespace(namespace)
  relationTuple.setObject(object)
  relationTuple.setRelation('access')

  const subjectSet = new acl.SubjectSet()
  subjectSet.setNamespace(namespace)
  subjectSet.setObject(object)
  subjectSet.setRelation('owner')

  const sub = new acl.Subject()
  sub.setSet(subjectSet)
  relationTuple.setSubject(sub)

  insert(relationTuple)
})

// should be subject set rewrite
// access on parent means access on child
;['/photos/beach.jpg', '/photos/mountains.jpg'].forEach((file) => {
  const relationTuple = new acl.RelationTuple()
  relationTuple.setNamespace('files')
  relationTuple.setObject(file)
  relationTuple.setRelation('access')

  const subjectSet = new acl.SubjectSet()
  subjectSet.setNamespace('directories')
  subjectSet.setObject('/photos')
  subjectSet.setRelation('access')

  const sub = new acl.Subject()
  sub.setSet(subjectSet)
  relationTuple.setSubject(sub)

  insert(relationTuple)
})

writeClient.transactRelationTuples(writeRequest, (error) => {
  if (error) {
    console.log('Encountered error', error)
  } else {
    console.log('Successfully created tuples')
  }
})
