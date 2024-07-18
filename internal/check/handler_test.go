// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package check_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/test/bufconn"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tidwall/gjson"

	"github.com/ory/x/pointerx"

	"github.com/ory/keto/internal/check"
	"github.com/ory/keto/internal/driver"
	"github.com/ory/keto/internal/driver/config"
	client "github.com/ory/keto/internal/httpclient"
	"github.com/ory/keto/internal/namespace"
	"github.com/ory/keto/internal/relationtuple"
	"github.com/ory/keto/internal/x"
	"github.com/ory/keto/ketoapi"
	rts "github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2"
)

func assertAllowed(t *testing.T, resp *http.Response) {
	t.Helper()

	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, resp.StatusCode, "%s", body)
	assert.True(t, gjson.GetBytes(body, "allowed").Bool())
}

type responseAssertion func(t *testing.T, resp *http.Response)

func baseAssertDenied(t *testing.T, resp *http.Response) {
	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	assert.Equal(t, http.StatusForbidden, resp.StatusCode, "%s", body)
	assert.False(t, gjson.GetBytes(body, "allowed").Bool())
}

// For OpenAPI clients, we want to always return a 200 status code even if the
// check returned "denied" to not cause exceptions etc. in the generated clients.
func openAPIAssertDenied(t *testing.T, resp *http.Response) {
	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, resp.StatusCode, "%s", body)
	assert.False(t, gjson.GetBytes(body, "allowed").Bool())
}

func TestCheckRESTHandler(t *testing.T) {
	nspaces := []*namespace.Namespace{{
		Name: "check handler",
	}}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	reg := driver.NewSqliteTestRegistry(t, false)
	require.NoError(t, reg.Config(ctx).Set(config.KeyNamespaces, nspaces))
	h := check.NewHandler(reg)
	r := httprouter.New()
	h.RegisterReadRoutes(&x.ReadRouter{Router: r})
	ts := httptest.NewServer(r)
	defer ts.Close()

	for _, suite := range []struct {
		name         string
		base         string
		assertDenied responseAssertion
	}{
		{name: "base", base: check.RouteBase, assertDenied: baseAssertDenied},
		{name: "openapi", base: check.OpenAPIRouteBase, assertDenied: openAPIAssertDenied},
	} {
		t.Run("suite="+suite.name, func(t *testing.T) {
			assertDenied := suite.assertDenied

			t.Run("case=returns bad request on malformed int", func(t *testing.T) {
				resp, err := ts.Client().Get(ts.URL + suite.base + "?max-depth=foo")
				require.NoError(t, err)

				assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
				body, err := io.ReadAll(resp.Body)
				require.NoError(t, err)
				assert.Contains(t, string(body), "invalid syntax")
			})

			t.Run("case=returns bad request on malformed input", func(t *testing.T) {
				resp, err := ts.Client().Get(ts.URL + suite.base + "?" + url.Values{
					"subject": {"not#a valid userset rewrite"},
				}.Encode())
				require.NoError(t, err)

				assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
			})

			t.Run("case=returns bad request on missing subject", func(t *testing.T) {
				resp, err := ts.Client().Get(ts.URL + suite.base)
				require.NoError(t, err)

				assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
				body, err := io.ReadAll(resp.Body)
				require.NoError(t, err)
				assert.Contains(t, string(body), "subject is not allowed to be nil")
			})

			t.Run("case=returns denied on unknown namespace", func(t *testing.T) {
				resp, err := ts.Client().Get(ts.URL + suite.base + "?" + (&ketoapi.RelationTuple{
					Namespace: "not" + nspaces[0].Name,
					Object:    "o",
					Relation:  "r",
					SubjectID: pointerx.Ptr("s"),
				}).ToURLQuery().Encode())
				require.NoError(t, err)

				assertDenied(t, resp)
			})

			t.Run("case=returns allowed", func(t *testing.T) {
				rt := &ketoapi.RelationTuple{
					Namespace: nspaces[0].Name,
					Object:    "o",
					Relation:  "r",
					SubjectID: pointerx.Ptr("s"),
				}
				relationtuple.MapAndWriteTuples(t, reg, rt)

				q := rt.ToURLQuery()
				resp, err := ts.Client().Get(ts.URL + suite.base + "?" + q.Encode())
				require.NoError(t, err)

				assertAllowed(t, resp)
			})

			t.Run("case=returns denied", func(t *testing.T) {
				resp, err := ts.Client().Get(ts.URL + suite.base + "?" + (&ketoapi.RelationTuple{
					Namespace: nspaces[0].Name,
					Object:    "foo",
					Relation:  "r",
					SubjectID: pointerx.Ptr("s"),
				}).ToURLQuery().Encode())
				require.NoError(t, err)

				assertDenied(t, resp)
			})
		})
	}
}

func TestBatchCheckRESTHandler(t *testing.T) {
	nspaces := []*namespace.Namespace{{
		Name: "batch-check-handler",
	}}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	reg := driver.NewSqliteTestRegistry(t, false)
	require.NoError(t, reg.Config(ctx).Set(config.KeyNamespaces, nspaces))
	h := check.NewHandler(reg)
	r := httprouter.New()
	h.RegisterReadRoutes(&x.ReadRouter{Router: r})
	ts := httptest.NewServer(r)
	defer ts.Close()

	t.Run("case=returns bad request on non-int max depth", func(t *testing.T) {
		resp, err := ts.Client().Post(ts.URL+check.BatchRoute+"?max-depth=foo",
			"application/json", nil)
		require.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
		body, err := io.ReadAll(resp.Body)
		require.NoError(t, err)
		assert.Contains(t, string(body), "invalid syntax")
	})

	t.Run("case=returns bad request on invalid request body", func(t *testing.T) {
		resp, err := ts.Client().Post(buildBatchURL(ts.URL, "5"),
			"application/json", strings.NewReader("not-json"))
		require.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
		body, err := io.ReadAll(resp.Body)
		require.NoError(t, err)
		assert.Contains(t, string(body), "could not unmarshal json")
	})

	t.Run("case=returns bad request with too many tuples", func(t *testing.T) {
		tuples := make([]client.Relationship, 11)
		for i := 0; i < len(tuples); i++ {
			tuples[i] = client.Relationship{
				Namespace: "n",
				Object:    "o",
				Relation:  "r",
				SubjectId: pointerx.Ptr("s"),
			}
		}
		reqBody := client.BatchCheckPermissionBody{Tuples: tuples}
		bodyBytes, err := json.Marshal(reqBody)
		require.NoError(t, err)

		resp, err := ts.Client().Post(buildBatchURL(ts.URL, "5"),
			"application/json", bytes.NewReader(bodyBytes))
		require.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
		body, err := io.ReadAll(resp.Body)
		require.NoError(t, err)
		assert.Contains(t, string(body), "batch exceeds max size of 10")
	})

	t.Run("case=check tuples", func(t *testing.T) {
		rt := &ketoapi.RelationTuple{
			Namespace: nspaces[0].Name,
			Object:    "o",
			Relation:  "r",
			SubjectID: pointerx.Ptr("s"),
		}
		relationtuple.MapAndWriteTuples(t, reg, rt)

		reqBody := client.BatchCheckPermissionBody{Tuples: []client.Relationship{
			{ // Allowed
				Namespace: nspaces[0].Name,
				Object:    "o",
				Relation:  "r",
				SubjectId: pointerx.Ptr("s"),
			},
			{ // Not-allowed
				Namespace: nspaces[0].Name,
				Object:    "o2",
				Relation:  "r",
				SubjectId: pointerx.Ptr("s"),
			},
			{ // Unknown namespace
				Namespace: "n2",
				Object:    "o",
				Relation:  "r",
				SubjectId: pointerx.Ptr("s"),
			},
		}}
		bodyBytes, err := json.Marshal(reqBody)
		require.NoError(t, err)

		resp, err := ts.Client().Post(buildBatchURL(ts.URL, "5"),
			"application/json", bytes.NewReader(bodyBytes))
		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		body, err := io.ReadAll(resp.Body)
		require.NoError(t, err)
		var respBody client.BatchCheckPermissionResult
		require.NoError(t, json.Unmarshal(body, &respBody))
		require.Equal(t, respBody, client.BatchCheckPermissionResult{
			Results: []client.CheckPermissionResultWithError{
				{
					Allowed: true,
					Error:   nil,
				},
				{
					Allowed: false,
					Error:   nil,
				},
				{
					Allowed: false,
					Error:   pointerx.Ptr("The requested resource could not be found"),
				},
			},
		})

		// Check again with the default parallelization factor
		resp, err = ts.Client().Post(fmt.Sprintf("%s%s?max-depth=%v", ts.URL, check.BatchRoute, 5),
			"application/json", bytes.NewReader(bodyBytes))
		require.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		body, err = io.ReadAll(resp.Body)
		require.NoError(t, err)
		var defaultParallelizationRespBody client.BatchCheckPermissionResult
		require.NoError(t, json.Unmarshal(body, &defaultParallelizationRespBody))
		require.Equal(t, respBody, defaultParallelizationRespBody)
	})
}

func buildBatchURL(baseURL, maxDepth string) string {
	return fmt.Sprintf("%s%s?max-depth=%s",
		baseURL, check.BatchRoute, maxDepth)
}

func TestBatchCheckGRPCHandler(t *testing.T) {
	ctx := context.Background()

	reg := driver.NewSqliteTestRegistry(t, false)
	h := check.NewHandler(reg)

	l := bufconn.Listen(1024 * 1024)
	s := grpc.NewServer()
	h.RegisterReadGRPC(s)
	go func() {
		if err := s.Serve(l); err != nil {
			t.Logf("Server exited with error: %v", err)
		}
	}()
	t.Cleanup(s.Stop)

	conn, err := grpc.Dial("bufnet",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) { return l.Dial() }),
	)
	require.NoError(t, err)
	t.Cleanup(func() { conn.Close() })

	nspaces := []*namespace.Namespace{{
		Name: "batch-check-grpc",
	}}
	require.NoError(t, reg.Config(ctx).Set(config.KeyNamespaces, nspaces))

	checkClient := rts.NewCheckServiceClient(conn)

	t.Run("case=returns bad request when batch too large", func(t *testing.T) {
		tuples := make([]*rts.RelationTuple, 11)
		for i := 0; i < len(tuples); i++ {
			tuples[i] = &rts.RelationTuple{
				Namespace: "n",
				Object:    "o",
				Relation:  "r",
				Subject: &rts.Subject{
					Ref: &rts.Subject_Id{
						Id: "s",
					},
				},
			}
		}
		_, err := checkClient.BatchCheck(ctx, &rts.BatchCheckRequest{
			Tuples:   tuples,
			MaxDepth: 5,
		})
		statusErr, ok := status.FromError(err)
		require.True(t, ok)
		require.Equal(t, codes.InvalidArgument, statusErr.Code())
		require.Equal(t, "batch exceeds max size of 10", statusErr.Message())
	})

	t.Run("case=batch check", func(t *testing.T) {
		rt := &ketoapi.RelationTuple{
			Namespace: nspaces[0].Name,
			Object:    "o",
			Relation:  "r",
			SubjectID: pointerx.Ptr("s"),
		}
		relationtuple.MapAndWriteTuples(t, reg, rt)

		resp, err := checkClient.BatchCheck(ctx, &rts.BatchCheckRequest{
			Tuples: []*rts.RelationTuple{
				{ // Allowed
					Namespace: nspaces[0].Name,
					Object:    "o",
					Relation:  "r",
					Subject: &rts.Subject{
						Ref: &rts.Subject_Id{
							Id: "s",
						},
					},
				},
				{ // Unknown namespace
					Namespace: "n2",
					Object:    "o",
					Relation:  "r",
					Subject: &rts.Subject{
						Ref: &rts.Subject_Id{
							Id: "s",
						},
					},
				},
				{ // Not allowed
					Namespace: nspaces[0].Name,
					Object:    "o2",
					Relation:  "r",
					Subject: &rts.Subject{
						Ref: &rts.Subject_Id{
							Id: "s",
						},
					},
				},
			},
			MaxDepth: 5,
		})
		require.NoError(t, err)
		require.Len(t, resp.Results, 3)
		require.True(t, resp.Results[0].Allowed)
		require.Empty(t, resp.Results[0].Error)
		require.False(t, resp.Results[1].Allowed)
		require.Equal(t, resp.Results[1].Error, "The requested resource could not be found")
		require.False(t, resp.Results[2].Allowed)
		require.Empty(t, resp.Results[2].Error)
	})
}
