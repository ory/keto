// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package e2e

import (
	"fmt"
	"io"
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
	"github.com/ory/keto/internal/relationtuple"
	"github.com/ory/keto/internal/x"
	"github.com/ory/keto/internal/x/dbx"
	"github.com/ory/keto/ketoapi"
)

type (
	transactClient interface {
		client
		transactTuples(t *testing.T, ins []*ketoapi.RelationTuple, del []*ketoapi.RelationTuple)
	}
	client interface {
		createTuple(t *testing.T, r *ketoapi.RelationTuple)
		deleteTuple(t *testing.T, r *ketoapi.RelationTuple)
		deleteAllTuples(t *testing.T, q *ketoapi.RelationQuery)
		queryTuple(t *testing.T, q *ketoapi.RelationQuery, opts ...x.PaginationOptionSetter) *ketoapi.GetResponse
		queryTupleErr(t *testing.T, expected herodot.DefaultError, q *ketoapi.RelationQuery, opts ...x.PaginationOptionSetter)
		check(t *testing.T, r *ketoapi.RelationTuple) bool
		batchCheck(t *testing.T, r []*ketoapi.RelationTuple) []checkResponse
		batchCheckErr(t *testing.T, requestTuples []*ketoapi.RelationTuple, expected herodot.DefaultError)
		expand(t *testing.T, r *ketoapi.SubjectSet, depth int) *ketoapi.Tree[*ketoapi.RelationTuple]
		oplCheckSyntax(t *testing.T, content []byte) []*ketoapi.ParseError
		waitUntilLive(t *testing.T)
		queryNamespaces(t *testing.T) ketoapi.GetNamespacesResponse
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

func TestServeCORS(t *testing.T) {
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
	req.Header.Set("Access-Control-Request-Method", http.MethodGet)
	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	require.Equal(t, http.StatusNoContent, resp.StatusCode)
	assert.Equal(t, "https://ory.sh", resp.Header.Get("Access-Control-Allow-Origin"), "%+v", resp.Header)
}
