package e2e

import (
	"fmt"
	"github.com/ory/keto/ketoapi"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/ory/herodot"
	"github.com/ory/x/cmdx"
	prometheus "github.com/ory/x/prometheusx"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ory/keto/cmd"
	cliclient "github.com/ory/keto/cmd/client"
	"github.com/ory/keto/internal/expand"
	"github.com/ory/keto/internal/relationtuple"
	"github.com/ory/keto/internal/x"
	"github.com/ory/keto/internal/x/dbx"
)

type (
	transactClient interface {
		client
		transactTuples(t require.TestingT, ins []*ketoapi.RelationTuple, del []*ketoapi.RelationTuple)
	}
	client interface {
		createTuple(t require.TestingT, r *ketoapi.RelationTuple)
		deleteTuple(t require.TestingT, r *ketoapi.RelationTuple)
		deleteAllTuples(t require.TestingT, q *ketoapi.RelationQuery)
		queryTuple(t require.TestingT, q *ketoapi.RelationQuery, opts ...x.PaginationOptionSetter) *ketoapi.GetResponse
		queryTupleErr(t require.TestingT, expected herodot.DefaultError, q *ketoapi.RelationQuery, opts ...x.PaginationOptionSetter)
		check(t require.TestingT, r *ketoapi.RelationTuple) bool
		expand(t require.TestingT, r *ketoapi.SubjectSet, depth int) *expand.Tree
		waitUntilLive(t require.TestingT)
	}
)

const (
	promLogLine = "promhttp_metric_handler_requests_total"
)

func Test(t *testing.T) {
	t.Parallel()
	for _, dsn := range dbx.GetDSNs(t, false) {
		dsn := dsn
		t.Run(fmt.Sprintf("dsn=%s", dsn.Name), func(t *testing.T) {
			t.Parallel()

			ctx, reg, addNamespace := newInitializedReg(t, dsn, nil)

			closeServer := startServer(ctx, t, reg)
			t.Cleanup(closeServer)

			// The test cases start here
			// We execute every test with all clients available
			for _, cl := range []client{
				&grpcClient{
					readRemote:  reg.Config(ctx).ReadAPIListenOn(),
					writeRemote: reg.Config(ctx).WriteAPIListenOn(),
					ctx:         ctx,
				},
				&restClient{
					readURL:  "http://" + reg.Config(ctx).ReadAPIListenOn(),
					writeURL: "http://" + reg.Config(ctx).WriteAPIListenOn(),
				},
				&cliClient{c: &cmdx.CommandExecuter{
					New: func() *cobra.Command {
						return cmd.NewRootCmd(nil)
					},
					Ctx:            ctx,
					PersistentArgs: []string{"--" + cliclient.FlagReadRemote, reg.Config(ctx).ReadAPIListenOn(), "--" + cliclient.FlagWriteRemote, reg.Config(ctx).WriteAPIListenOn(), "--" + cmdx.FlagFormat, string(cmdx.FormatJSON)},
				}},
				&sdkClient{
					readRemote:  reg.Config(ctx).ReadAPIListenOn(),
					writeRemote: reg.Config(ctx).WriteAPIListenOn(),
				},
			} {
				t.Run(fmt.Sprintf("client=%T", cl), runCases(cl, addNamespace))

				if tc, ok := cl.(transactClient); ok {

					t.Run(fmt.Sprintf("transactClient=%T", cl), runTransactionCases(tc, addNamespace))
				}
			}

			t.Run("case=metrics are served", func(t *testing.T) {
				t.Parallel()
				(&grpcClient{
					readRemote:  reg.Config(ctx).ReadAPIListenOn(),
					writeRemote: reg.Config(ctx).WriteAPIListenOn(),
					ctx:         ctx,
				}).waitUntilLive(t)

				t.Run("case=on "+prometheus.MetricsPrometheusPath, func(t *testing.T) {
					t.Parallel()
					resp, err := http.Get(fmt.Sprintf("http://%s%s", reg.Config(ctx).MetricsListenOn(), prometheus.MetricsPrometheusPath))
					require.NoError(t, err)
					require.Equal(t, resp.StatusCode, http.StatusOK)
					body, err := ioutil.ReadAll(resp.Body)
					require.NoError(t, err)
					require.Contains(t, string(body), promLogLine)
				})

				t.Run("case=not on /", func(t *testing.T) {
					t.Parallel()
					resp, err := http.Get(fmt.Sprintf("http://%s", reg.Config(ctx).MetricsListenOn()))
					require.NoError(t, err)
					require.Equal(t, resp.StatusCode, http.StatusNotFound)
				})
			})
		})
	}
}

func TestServeConfig(t *testing.T) {
	t.Parallel()

	ctx, reg, _ := newInitializedReg(t, dbx.GetSqlite(t, dbx.SQLiteMemory), map[string]interface{}{
		"serve.read.cors.enabled":         true,
		"serve.read.cors.debug":           true,
		"serve.read.cors.allowed_methods": []string{http.MethodGet},
		"serve.read.cors.allowed_origins": []string{"https://ory.sh"},
	})

	closeServer := startServer(ctx, t, reg)
	t.Cleanup(closeServer)

	for !healthReady(t, "http://"+reg.Config(ctx).ReadAPIListenOn()) {
		t.Log("Waiting for health check to be ready")
		time.Sleep(10 * time.Millisecond)
	}

	req, err := http.NewRequest(http.MethodOptions, "http://"+reg.Config(ctx).ReadAPIListenOn()+relationtuple.ReadRouteBase, nil)
	require.NoError(t, err)
	req.Header.Set("Origin", "https://ory.sh")
	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "https://ory.sh", resp.Header.Get("Access-Control-Allow-Origin"), "%+v", resp.Header)
}
