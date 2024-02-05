module.exports.Write = 
module.exports.Write = require('./ory/keto/relation_tuples/v1alpha2/write_service_pb')

module.exports = {
  ...require('./ory/keto/relation_tuples/v1alpha2/write_service_pb'),
  ...require('./ory/keto/relation_tuples/v1alpha2/write_service_connect'),

  ...require('./ory/keto/relation_tuples/v1alpha2/relation_tuples_pb'),
}