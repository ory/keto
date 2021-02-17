import grpc from '@ory/keto-acl/node_modules/@grpc/grpc-js/build/src/index.js'
import readService from '@ory/keto-acl/read_service_grpc_pb.js'
import readData from '@ory/keto-acl/read_service_pb.js'

it("works", () => {
    console.log = (...args) => {
        expect(args.length).toBe(1)
        expect(args[0]).toBe("Successfully created tuple.")
    }

    import("./index.js")

    const readClient = new readService.ReadServiceClient('127.0.0.1:4466', grpc.credentials.createInsecure())

    const listQuery = new readData.ListRelationTuplesRequest.Query()
    listQuery.setNamespace("messages")
    const listRequest = new readData.ListRelationTuplesRequest()
    listRequest.setQuery(listQuery)

    readClient.listRelationTuples(listRequest, (err, resp) => {
        expect(err).toBeFalsy()

        const tuples = resp.getRelationTuplesList()
        expect(tuples).toHaveLength(1)
        expect(tuples[0].getNamespace()).toBe("messages")
        expect(tuples[0].getObject()).toBe("02y_15_4w350m3")
        expect(tuples[0].getRelation()).toBe("decypher")
        expect(tuples[0].getSubject().getId()).toBe("john")
    })
})
