import * as relationTuples from './relation_tuples_pb'
import * as write from './write_service_pb'
import * as writeService from './write_service_grpc_pb'
import * as check from './check_service_pb'
import * as checkService from './check_service_grpc_pb'
import * as expand from './expand_service_pb'
import * as expandService from './expand_service_grpc_pb'
import * as read from './read_service_pb'
import * as readService from './read_service_grpc_pb'
declare module '@ory/keto-grpc-client/ory/keto/relation_tuples/v1alpha2' {
  export {
    relationTuples,
    write,
    writeService,
    check,
    checkService,
    expand,
    expandService,
    read,
    readService
  }
}
