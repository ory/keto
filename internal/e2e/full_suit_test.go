package e2e

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	prometheus "github.com/ory/x/prometheusx"

	"github.com/stretchr/testify/assert"

	"github.com/ory/keto/internal/x/dbx"

	"github.com/ory/herodot"

	"github.com/ory/keto/internal/x"

	"github.com/stretchr/testify/require"

	cliclient "github.com/ory/keto/cmd/client"
	"github.com/ory/keto/internal/expand"

	"github.com/ory/x/cmdx"

	"github.com/ory/keto/cmd"
	"github.com/ory/keto/internal/relationtuple"
)

type (
	transactClient interface {
		client
		transactTuples(t require.TestingT, ins []*relationtuple.InternalRelationTuple, del []*relationtuple.InternalRelationTuple)
	}
	client interface {
		createTuple(t require.TestingT, r *relationtuple.InternalRelationTuple)
		deleteTuple(t require.TestingT, r *relationtuple.InternalRelationTuple)
		deleteAllTuples(t require.TestingT, q *relationtuple.RelationQuery)
		queryTuple(t require.TestingT, q *relationtuple.RelationQuery, opts ...x.PaginationOptionSetter) *relationtuple.GetResponse
		queryTupleErr(t require.TestingT, expected herodot.DefaultError, q *relationtuple.RelationQuery, opts ...x.PaginationOptionSetter)
		check(t require.TestingT, r *relationtuple.InternalRelationTuple) bool
		expand(t require.TestingT, r *relationtuple.SubjectSet, depth int) *expand.Tree
		waitUntilLive(t require.TestingT)
	}
)

const (
	promLogLine = "promhttp_metric_handler_requests_total"
)

func Test(t *testing.T) {
	for _, dsn := range dbx.GetDSNs(t, false) {
		t.Run(fmt.Sprintf("dsn=%s", dsn.Name), func(t *testing.T) {
			ctx, reg, addNamespace := newInitializedReg(t, dsn, nil)

			closeServer := startServer(ctx, t, reg)
			defer closeServer()

			// The test cases start here
			// We execute every test with all clients available
			for _, cl := range []client{
				&grpcClient{
					readRemote:  reg.Config().ReadAPIListenOn(),
					writeRemote: reg.Config().WriteAPIListenOn(),
					ctx:         ctx,
				},
				&restClient{
					readURL:  "http://" + reg.Config().ReadAPIListenOn(),
					writeURL: "http://" + reg.Config().WriteAPIListenOn(),
				},
				&cliClient{c: &cmdx.CommandExecuter{
					New:            cmd.NewRootCmd,
					Ctx:            ctx,
					PersistentArgs: []string{"--" + cliclient.FlagReadRemote, reg.Config().ReadAPIListenOn(), "--" + cliclient.FlagWriteRemote, reg.Config().WriteAPIListenOn(), "--" + cmdx.FlagFormat, string(cmdx.FormatJSON)},
				}},
				&sdkClient{
					readRemote:  reg.Config().ReadAPIListenOn(),
					writeRemote: reg.Config().WriteAPIListenOn(),
				},
			} {
				t.Run(fmt.Sprintf("client=%T", cl), runCases(cl, addNamespace))

				if tc, ok := cl.(transactClient); ok {
					t.Run(fmt.Sprintf("transactClient=%T", cl), runTransactionCases(tc, addNamespace))
				}

				t.Run("case=metrics are served", func(t *testing.T) {
					t.Run("case=are served on "+prometheus.MetricsPrometheusPath, func(t *testing.T) {
						resp, err := http.Get(fmt.Sprintf("http://%s%s", reg.Config().MetricsListenOn(), prometheus.MetricsPrometheusPath))
						require.NoError(t, err)
						require.Equal(t, resp.StatusCode, http.StatusOK)
						body, err := ioutil.ReadAll(resp.Body)
						require.NoError(t, err)
						require.Contains(t, string(body), promLogLine)
					})
					t.Run("case=are not served on /", func(t *testing.T) {
						resp, err := http.Get(fmt.Sprintf("http://%s", reg.Config().MetricsListenOn()))
						require.NoError(t, err)
						require.Equal(t, resp.StatusCode, http.StatusNotFound)
					})
				})
			}
		})
	}
}

func TestServeConfig(t *testing.T) {
	ctx, reg, _ := newInitializedReg(t, dbx.GetSqlite(t, dbx.SQLiteMemory), map[string]interface{}{
		"serve.read.cors.enabled":         true,
		"serve.read.cors.debug":           true,
		"serve.read.cors.allowed_methods": []string{http.MethodGet},
		"serve.read.cors.allowed_origins": []string{"https://ory.sh"},
	})

	closeServer := startServer(ctx, t, reg)
	defer closeServer()

	for !healthReady(t, "http://"+reg.Config().ReadAPIListenOn()) {
		t.Log("Waiting for health check to be ready")
		time.Sleep(10 * time.Millisecond)
	}

	req, err := http.NewRequest(http.MethodOptions, "http://"+reg.Config().ReadAPIListenOn()+relationtuple.RouteBase, nil)
	require.NoError(t, err)
	req.Header.Set("Origin", "https://ory.sh")
	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "https://ory.sh", resp.Header.Get("Access-Control-Allow-Origin"), "%+v", resp.Header)
}
