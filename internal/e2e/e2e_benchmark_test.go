package e2e

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ory/keto/internal/namespace"
	"github.com/ory/keto/internal/x/dbx"
	"github.com/ory/keto/ketoapi"
	"github.com/ory/x/pointerx"
)

func BenchmarkE2E(b *testing.B) {

	dsn := dbx.GetSqlite(b, dbx.SQLiteMemory)
	ctx, reg, namespaceTestMgr, getAddr := newInitializedReg(b, dsn, map[string]interface{}{"log.level": "panic"})
	closeServer := startServer(ctx, b, reg)
	b.Cleanup(closeServer)

	_, _, readAddr := getAddr(b, "read")
	_, _, writeAddr := getAddr(b, "write")
	_, _, oplAddr := getAddr(b, "opl")

	for _, cl := range []client{
		newGrpcClient(b, ctx,
			readAddr,
			writeAddr,
			oplAddr,
		),
		&restClient{
			readURL:      "http://" + readAddr,
			writeURL:     "http://" + writeAddr,
			oplSyntaxURL: "http://" + oplAddr,
		},
		&sdkClient{
			readRemote:   readAddr,
			writeRemote:  writeAddr,
			syntaxRemote: oplAddr,
		},
	} {

		b.Run(fmt.Sprintf("client=%T", cl), func(b *testing.B) {
			n := &namespace.Namespace{Name: "test"}
			namespaceTestMgr.add(b, n)
			cl.waitUntilLive(b)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				tuple := &ketoapi.RelationTuple{
					Namespace: n.Name,
					Object:    fmt.Sprintf("object %d for client %T", i, cl),
					Relation:  "access",
					SubjectID: pointerx.Ptr("client"),
				}
				cl.createTuple(b, tuple)

				resp := cl.queryTuple(b, &ketoapi.RelationQuery{Namespace: &tuple.Namespace})
				require.Len(b, resp.RelationTuples, 1)
				assert.Equal(b, tuple, resp.RelationTuples[0])

				assert.True(b, cl.check(b, tuple))
				batchResult := cl.batchCheck(b, []*ketoapi.RelationTuple{tuple})
				require.Len(b, batchResult, 1)
				assert.True(b, batchResult[0].allowed)
				assert.Empty(b, batchResult[0].errorMessage)

				cl.deleteTuple(b, tuple)
				resp = cl.queryTuple(b, &ketoapi.RelationQuery{Namespace: &tuple.Namespace})
				require.Len(b, resp.RelationTuples, 0)
			}
		})
	}

}
