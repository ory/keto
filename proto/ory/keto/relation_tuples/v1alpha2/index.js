// This has to be as is because of the way named exports
// are supported in ESM from CommonJS packages. Don't ask, it just works.

const relationTuples = require('./relation_tuples_pb.js')
const write = require('./write_service_pb.js')
const writeService = require('./write_service_grpc_pb.js')
const check = require('./check_service_pb.js')
const checkService = require('./check_service_grpc_pb.js')
const expand = require('./expand_service_pb.js')
const expandService = require('./expand_service_grpc_pb.js')
const read = require('./read_service_pb.js')
const readService = require('./read_service_grpc_pb.js')

module.exports = {
    relationTuples,
    check,
    checkService,
    write,
    writeService,
    expand,
    expandService,
    read,
    readService
}
