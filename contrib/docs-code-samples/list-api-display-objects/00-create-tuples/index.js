import grpc from '@ory/keto-acl/node_modules/@grpc/grpc-js/build/src/index.js'
import acl from '@ory/keto-acl/acl_pb.js'
import writeService from '@ory/keto-acl/write_service_grpc_pb.js'
import writeData from '@ory/keto-acl/write_service_pb.js'

const writeClient = new writeService.WriteServiceClient(
  '127.0.0.1:4467',
  grpc.credentials.createInsecure()
)

const writeRequest = new writeData.TransactRelationTuplesRequest()

const addToChat = (chatName) => (user) => {
  const relationTuple = new acl.RelationTuple()
  relationTuple.setNamespace('chats')
  relationTuple.setObject(chatName)
  relationTuple.setRelation('member')

  const sub = new acl.Subject()
  sub.setId(user)
  relationTuple.setSubject(sub)

  const tupleDelta = new writeData.RelationTupleDelta()
  tupleDelta.setAction(writeData.RelationTupleDelta.Action.INSERT)
  tupleDelta.setRelationTuple(relationTuple)

  writeRequest.addRelationTupleDeltas(tupleDelta)
}

;['PM', 'Vincent', 'Julia'].forEach(addToChat('memes'))
;['PM', 'Julia'].forEach(addToChat('cars'))
;['PM', 'Vincent', 'Julia', 'Patrik'].forEach(addToChat('coffee-break'))

writeClient.transactRelationTuples(writeRequest, (error) => {
  if (error) {
    console.log('Encountered error', error)
  } else {
    console.log('Successfully created tuples')
  }
})
