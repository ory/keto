// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package e2e

import (
	"fmt"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ory/herodot"
	"github.com/ory/x/cmdx"
	prometheus "github.com/ory/x/prometheusx"

	"github.com/ory/keto/cmd"
	cliclient "github.com/ory/keto/cmd/client"
	"github.com/ory/keto/internal/relationtuple"
	"github.com/ory/keto/internal/x"
	"github.com/ory/keto/internal/x/dbx"
	"github.com/ory/keto/ketoapi"
)

type (
	transactClient interface {
		client
		transactTuples(t testing.TB, ins []*ketoapi.RelationTuple, del []*ketoapi.RelationTuple)
	}
	client interface {
		createTuple(t testing.TB, r *ketoapi.RelationTuple)
		deleteTuple(t testing.TB, r *ketoapi.RelationTuple)
		deleteAllTuples(t testing.TB, q *ketoapi.RelationQuery)
		queryTuple(t testing.TB, q *ketoapi.RelationQuery, opts ...x.PaginationOptionSetter) *ketoapi.GetResponse
		queryTupleErr(t testing.TB, expected herodot.DefaultError, q *ketoapi.RelationQuery, opts ...x.PaginationOptionSetter)
		check(t testing.TB, r *ketoapi.RelationTuple) bool
		batchCheck(t testing.TB, r []*ketoapi.RelationTuple) []checkResponse
		batchCheckErr(t testing.TB, requestTuples []*ketoapi.RelationTuple, expected herodot.DefaultError)
		expand(t testing.TB, r *ketoapi.SubjectSet, depth int) *ketoapi.Tree[*ketoapi.RelationTuple]
		oplCheckSyntax(t testing.TB, content []byte) []*ketoapi.ParseError
		waitUntilLive(t testing.TB)
		queryNamespaces(t testing.TB) ketoapi.GetNamespacesResponse
	}
)

const (
	promLogLine = "promhttp_metric_handler_requests_total"
)

func Test(t *testing.T) {
	t.Parallel()
	for _, dsn := range dbx.GetDSNs(t, false) {
		t.Run(fmt.Sprintf("dsn=%s", dsn.Name), func(t *testing.T) {
			t.Parallel()

			ctx, reg, namespaceTestMgr, getAddr := newInitializedReg(t, dsn, nil)

			closeServer := startServer(ctx, t, reg)
			t.Cleanup(closeServer)

			_, _, readAddr := getAddr(t, "read")
			_, _, writeAddr := getAddr(t, "write")
			_, _, oplAddr := getAddr(t, "opl")
			_, _, metricsAddr := getAddr(t, "metrics")

			// The test cases start here
			// We execute every test with all clients available
			for _, cl := range []client{
				newGrpcClient(t, ctx,
					readAddr,
					writeAddr,
					oplAddr,
				),
				&restClient{
					readURL:      "http://" + readAddr,
					writeURL:     "http://" + writeAddr,
					oplSyntaxURL: "http://" + oplAddr,
				},
				&cliClient{c: &cmdx.CommandExecuter{
					New: func() *cobra.Command {
						return cmd.NewRootCmd(nil)
					},
					Ctx: ctx,
					PersistentArgs: []string{
						"--" + cliclient.FlagReadRemote, readAddr,
						"--" + cliclient.FlagWriteRemote, writeAddr,
						"--insecure-disable-transport-security=true",
						"--" + cmdx.FlagFormat, string(cmdx.FormatJSON),
					},
				}},
				&sdkClient{
					readRemote:   readAddr,
					writeRemote:  writeAddr,
					syntaxRemote: oplAddr,
				},
			} {
				t.Run(fmt.Sprintf("client=%T", cl), runCases(cl, namespaceTestMgr))

				if tc, ok := cl.(transactClient); ok {
					t.Run(fmt.Sprintf("transactClient=%T", cl), runTransactionCases(tc, namespaceTestMgr))
				}
			}

			t.Run("case=metrics are served", func(t *testing.T) {
				t.Parallel()
				newGrpcClient(t, ctx,
					readAddr,
					writeAddr,
					oplAddr,
				).waitUntilLive(t)

				t.Run("case=on "+prometheus.MetricsPrometheusPath, func(t *testing.T) {
					t.Parallel()
					resp, err := http.Get(fmt.Sprintf("http://%s%s", metricsAddr, prometheus.MetricsPrometheusPath))
					require.NoError(t, err)
					require.Equal(t, resp.StatusCode, http.StatusOK)
					body, err := io.ReadAll(resp.Body)
					require.NoError(t, err)
					require.Contains(t, string(body), promLogLine)
				})

				t.Run("case=not on /", func(t *testing.T) {
					t.Parallel()
					resp, err := http.Get(fmt.Sprintf("http://%s", metricsAddr))
					require.NoError(t, err)
					require.Equal(t, resp.StatusCode, http.StatusNotFound)
				})
			})
		})
	}
}

func TestServeConfig(t *testing.T) {
	t.Parallel()

	ctx, reg, _, getAddr := newInitializedReg(t, dbx.GetSqlite(t, dbx.SQLiteMemory), map[string]interface{}{
		"serve.read.cors.enabled":         true,
		"serve.read.cors.debug":           true,
		"serve.read.cors.allowed_methods": []string{http.MethodGet},
		"serve.read.cors.allowed_origins": []string{"https://ory.sh"},
	})

	closeServer := startServer(ctx, t, reg)
	t.Cleanup(closeServer)

	_, _, readAddr := getAddr(t, "read")

	for !healthReady(t, "http://"+readAddr) {
		t.Log("Waiting for health check to be ready")
		time.Sleep(10 * time.Millisecond)
	}
	t.Log("Health check is ready")

	req, err := http.NewRequest(http.MethodOptions, "http://"+readAddr+relationtuple.ReadRouteBase, nil)
	require.NoError(t, err)
	req.Header.Set("Origin", "https://ory.sh")
	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "https://ory.sh", resp.Header.Get("Access-Control-Allow-Origin"), "%+v", resp.Header)
}
