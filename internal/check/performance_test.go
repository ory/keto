package check_test

//import (
//	"context"
//	"fmt"
//	"strconv"
//	"testing"
//
//	"github.com/stretchr/testify/assert"
//	"github.com/stretchr/testify/require"
//
//	"github.com/ory/keto/internal/driver"
//	"github.com/ory/keto/internal/e2e"
//	"github.com/ory/keto/internal/expand"
//	"github.com/ory/keto/internal/namespace"
//	"github.com/ory/keto/internal/relationtuple"
//)
//
//const (
//	defaultNamespace = "default"
//	indirectRelation = "indirect"
//	root             = "root"
//
//	tuplesPerLevel = 10
//)
//
//func createDataset(b testing.TB, reg driver.Registry, depth int, parent string) (created int) {
//	newTuples := make([]*relationtuple.InternalRelationTuple, tuplesPerLevel)
//
//	for p := 0; p < tuplesPerLevel; p++ {
//		me := parent + " " + strconv.Itoa(p)
//		newTuples[p] = &relationtuple.InternalRelationTuple{
//			Namespace: defaultNamespace,
//			Object:    parent,
//			Relation:  indirectRelation,
//			Subject: &relationtuple.SubjectSet{
//				Namespace: defaultNamespace,
//				Object:    me,
//				Relation:  indirectRelation,
//			},
//		}
//
//		if depth >= 1 {
//			created += createDataset(b, reg, depth-1, me)
//		}
//	}
//
//	require.NoError(b, reg.Persister().WriteRelationTuples(context.Background(), newTuples...))
//	return created + tuplesPerLevel
//}
//
//type noErrButResult struct {
//	res bool
//	cid int
//	o   string
//}
//
//func (e *noErrButResult) Error() string {
//	return "why did you call me, wtf?"
//}
//
//func XBenchmarkCheckEngine(b *testing.B) {
//	for _, dsn := range e2e.GetDSNs(b) {
//		func(dsn *e2e.DsnT) {
//			nspaces := []*namespace.Namespace{{Name: defaultNamespace, ID: 1}}
//			ctx, reg := e2e.NewInitializedReg(b, dsn, nspaces)
//
//			b.ResetTimer()
//
//			clientC := make([]chan int, 100)
//			clientErrs := make(chan error, 100)
//			for ci := range clientC {
//				clientC[ci] = make(chan int)
//
//				go func(ci int) {
//					for d := range clientC[ci] {
//						o := root
//						for i := 0; i < d; i++ {
//							o += " " + strconv.Itoa(ci%tuplesPerLevel)
//						}
//
//						res, err := reg.PermissionEngine().SubjectIsAllowed(ctx, &relationtuple.InternalRelationTuple{
//							Namespace: defaultNamespace,
//							Object:    root,
//							Relation:  indirectRelation,
//							Subject: &relationtuple.SubjectSet{
//								Namespace: defaultNamespace,
//								Object:    o,
//								Relation:  indirectRelation,
//							},
//						})
//						if err == nil {
//							err = &noErrButResult{res: res, cid: ci, o: o}
//						}
//						clientErrs <- err
//					}
//				}(ci)
//			}
//
//			defer func() {
//				for _, c := range clientC {
//					close(c)
//				}
//			}()
//
//			for depth := 2; depth <= 4; depth += 1 {
//				b.StopTimer()
//
//				b.Logf("created %d tuples", createDataset(b, reg, depth, root))
//
//				b.StartTimer()
//
//				for clients := 1; clients <= 5; clients++ {
//					b.Run(fmt.Sprintf("dsn=%s depth=%d clients=%d", dsn.Name, depth, clients), func(b *testing.B) {
//						for i := 0; i < b.N; i++ {
//							for ci := 0; ci < clients; ci++ {
//								clientC[ci] <- depth
//							}
//
//							for ci := 0; ci < clients; ci++ {
//								err := <-clientErrs
//								if res, ok := err.(*noErrButResult); ok {
//									require.True(b, res.res)
//									continue
//								}
//
//								b.Logf("got err %+v", err)
//								b.FailNow()
//							}
//						}
//					})
//				}
//			}
//		}(dsn)
//	}
//}
//
//func XTestCreateDataset(t *testing.T) {
//	for _, dsn := range e2e.GetDSNs(t) {
//		nspaces := []*namespace.Namespace{{Name: defaultNamespace, ID: 0}}
//
//		ctx, reg := e2e.NewInitializedReg(t, dsn, nspaces)
//		createDataset(t, reg, 4, "root")
//
//		tree, err := reg.ExpandEngine().BuildTree(ctx, &relationtuple.SubjectSet{
//			Namespace: defaultNamespace,
//			Object:    root,
//			Relation:  indirectRelation,
//		}, 5)
//		require.NoError(t, err)
//
//		for d := 0; d < 4; d++ {
//			assert.Equal(t, expand.Union, tree.Type)
//			assert.Len(t, tree.Children, tuplesPerLevel)
//			tree = tree.Children[0]
//		}
//	}
//}
