import grpc from '@ory/keto-acl/node_modules/@grpc/grpc-js/build/src/index.js'
import acl from '@ory/keto-acl/acl_pb.js'
import checkService from '@ory/keto-acl/check_service_grpc_pb.js'
import checkData from '@ory/keto-acl/check_service_pb.js'

const checkClient = new checkService.CheckServiceClient(
  '127.0.0.1:4466',
  grpc.credentials.createInsecure()
)

const checkRequest = new checkData.CheckRequest()
checkRequest.setNamespace('messages')
checkRequest.setObject('02y_15_4w350m3')
checkRequest.setRelation('decypher')

const sub = new acl.Subject()
sub.setId('john')
checkRequest.setSubject(sub)

checkClient.check(checkRequest, (error, resp) => {
  if (error) {
    console.log('Encountered error:', error)
  } else {
    console.log(resp.getAllowed() ? 'Allowed' : 'Denied')
  }
})
