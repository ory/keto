import grpc from '@ory/keto-acl/node_modules/@grpc/grpc-js/build/src/index.js'
import acl from '@ory/keto-acl/acl_pb.js'
import readService from '@ory/keto-acl/read_service_grpc_pb.js'
import readData from '@ory/keto-acl/read_service_pb.js'

const readClient = new readService.ReadServiceClient(
  '127.0.0.1:4466',
  grpc.credentials.createInsecure()
)

const readRequest = new readData.ListRelationTuplesRequest()
const query = new readData.ListRelationTuplesRequest.Query()
query.setNamespace('chats')
query.setRelation('member')

const sub = new acl.Subject()
sub.setId('PM')
query.setSubject(sub)

readRequest.setQuery(query)

readClient.listRelationTuples(readRequest, (error, resp) => {
  if (error) {
    console.log('Encountered error:', error)
  } else {
    console.log(
      resp
        .getRelationTuplesList()
        .map((tuple) => tuple.getObject())
        .join('\n')
    )
  }
})
