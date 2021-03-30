import grpc from '@ory/keto-grpc-client/node_modules/@grpc/grpc-js/build/src/index.js'
import acl from '@ory/keto-grpc-client/acl_pb.js'
import writeService from '@ory/keto-grpc-client/write_service_grpc_pb.js'
import writeData from '@ory/keto-grpc-client/write_service_pb.js'

const writeClient = new writeService.WriteServiceClient(
  '127.0.0.1:4467',
  grpc.credentials.createInsecure()
)

const relationTuple = new acl.RelationTuple()
relationTuple.setNamespace('messages')
relationTuple.setObject('02y_15_4w350m3')
relationTuple.setRelation('decypher')

const sub = new acl.Subject()
sub.setId('john')
relationTuple.setSubject(sub)

const tupleDelta = new writeData.RelationTupleDelta()
tupleDelta.setAction(writeData.RelationTupleDelta.Action.INSERT)
tupleDelta.setRelationTuple(relationTuple)

const writeRequest = new writeData.TransactRelationTuplesRequest()
writeRequest.addRelationTupleDeltas(tupleDelta)

writeClient.transactRelationTuples(writeRequest, (error) => {
  if (error) {
    console.log('Encountered error', error)
  } else {
    console.log('Successfully created tuple')
  }
})
