import grpc from '@ory/keto-grpc-client/node_modules/@grpc/grpc-js/build/src/index.js'
import writeService from '@ory/keto-grpc-client/write_service_grpc_pb.js'
import writeData from '@ory/keto-grpc-client/write_service_pb.js'
import readService from '@ory/keto-grpc-client/read_service_grpc_pb.js'
import readData from '@ory/keto-grpc-client/read_service_pb.js'

const readClient = new readService.ReadServiceClient(
  '127.0.0.1:4466',
  grpc.credentials.createInsecure()
)

const purgeNamespace = (namespace) => {
  const query = new readData.ListRelationTuplesRequest.Query()
  query.setNamespace(namespace)

  const readRequest = new readData.ListRelationTuplesRequest()
  readRequest.setQuery(query)

  readClient.listRelationTuples(readRequest, (err, resp) => {
    const writeClient = new writeService.WriteServiceClient(
      '127.0.0.1:4467',
      grpc.credentials.createInsecure()
    )

    const writeRequest = new writeData.TransactRelationTuplesRequest()

    resp.getRelationTuplesList().forEach((tuple) => {
      const tupleDelta = new writeData.RelationTupleDelta()
      tupleDelta.setAction(writeData.RelationTupleDelta.Action.DELETE)
      tupleDelta.setRelationTuple(tuple)
      writeRequest.addRelationTupleDeltas(tupleDelta)
    })

    writeClient.transactRelationTuples(writeRequest, (err) => {
      if (err) {
        console.log('Unexpected err', err)
        return 1
      }
    })
  })
}

purgeNamespace('files')
purgeNamespace('directories')
