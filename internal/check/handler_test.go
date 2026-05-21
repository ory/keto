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

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tidwall/gjson"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/test/bufconn"

	"github.com/ory/x/httprouterx"

	"github.com/ory/keto/internal/check"
	"github.com/ory/keto/internal/driver"
	"github.com/ory/keto/internal/driver/config"
	client "github.com/ory/keto/internal/httpclient"
	"github.com/ory/keto/internal/namespace"
	"github.com/ory/keto/internal/testhelpers"
	"github.com/ory/keto/ketoapi"
	rts "github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2"
)

func readBody(t *testing.T, resp *http.Response) []byte {
	t.Helper()
	defer func() { assert.NoError(t, resp.Body.Close()) }()
	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	return body
}

func assertAllowed(t *testing.T, resp *http.Response) {
	t.Helper()
	body := readBody(t, resp)
	require.Equal(t, http.StatusOK, resp.StatusCode, "%s", body)
	allowed := gjson.GetBytes(body, "allowed")
	require.True(t, allowed.Exists(), "response missing 'allowed' field: %s", body)
	require.True(t, allowed.Bool())
}

type responseAssertion func(t *testing.T, resp *http.Response)

func baseAssertDenied(t *testing.T, resp *http.Response) {
	t.Helper()
	body := readBody(t, resp)
	require.Equal(t, http.StatusForbidden, resp.StatusCode, "%s", body)
	allowed := gjson.GetBytes(body, "allowed")
	require.True(t, allowed.Exists(), "response missing 'allowed' field: %s", body)
	require.False(t, allowed.Bool())
}

// For OpenAPI clients, we want to always return a 200 status code even if the
// check returned "denied" to not cause exceptions etc. in the generated clients.
func openAPIAssertDenied(t *testing.T, resp *http.Response) {
	t.Helper()
	body := readBody(t, resp)
	require.Equal(t, http.StatusOK, resp.StatusCode, "%s", body)
	allowed := gjson.GetBytes(body, "allowed")
	require.True(t, allowed.Exists(), "response missing 'allowed' field: %s", body)
	require.False(t, allowed.Bool())
}

func newTestGRPCCheckClient(t *testing.T, h *check.Handler) rts.CheckServiceClient {
	t.Helper()
	l := bufconn.Listen(1024 * 1024)
	s := grpc.NewServer()
	h.RegisterReadGRPC(s)
	go func() {
		if err := s.Serve(l); err != nil {
			select {
			case <-t.Context().Done():
			default:
				t.Logf("gRPC server exited unexpectedly: %v", err)
			}
		}
	}()
	t.Cleanup(s.Stop)
	conn, err := grpc.NewClient("passthrough:///bufnet",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) { return l.Dial() }),
	)
	require.NoError(t, err)
	t.Cleanup(func() { require.NoError(t, conn.Close()) })
	return rts.NewCheckServiceClient(conn)
}

func buildBatchURL(baseURL, maxDepth string) string {
	return fmt.Sprintf("%s%s?max-depth=%s", baseURL, check.BatchRoute, maxDepth)
}

func TestCheckRESTHandler(t *testing.T) {
	nspaces := []*namespace.Namespace{{
		Name: "ns",
	}}

	reg := driver.NewSqliteTestRegistry(t, driver.WithNamespaces(nspaces))
	h := check.NewHandler(reg)
	r := httprouterx.NewRouterPublic()
	h.RegisterReadRoutes(r)
	ts := httptest.NewServer(r)
	t.Cleanup(ts.Close)

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

				require.Equal(t, http.StatusBadRequest, resp.StatusCode)
				require.Contains(t, string(readBody(t, resp)), "invalid syntax")
			})

			t.Run("case=returns bad request on malformed input", func(t *testing.T) {
				resp, err := ts.Client().Get(ts.URL + suite.base + "?" + url.Values{
					"subject": {"not#a valid userset rewrite"},
				}.Encode())
				require.NoError(t, err)

				require.Equal(t, http.StatusBadRequest, resp.StatusCode)
			})

			t.Run("case=returns bad request on missing subject", func(t *testing.T) {
				resp, err := ts.Client().Get(ts.URL + suite.base)
				require.NoError(t, err)

				require.Equal(t, http.StatusBadRequest, resp.StatusCode)
				require.Contains(t, string(readBody(t, resp)), "subject is not allowed to be nil")
			})

			t.Run("case=returns denied on unknown namespace", func(t *testing.T) {
				rt := testhelpers.APITupleFromString(t, "not:o#r@s")
				resp, err := ts.Client().Get(ts.URL + suite.base + "?" + rt.ToURLQuery().Encode())
				require.NoError(t, err)

				assertDenied(t, resp)
			})

			t.Run("case=returns allowed", func(t *testing.T) {
				rt := testhelpers.APITupleFromString(t, "ns:o#r@s")
				testhelpers.MapAndInsertTuples(t, reg, rt)

				q := rt.ToURLQuery()
				resp, err := ts.Client().Get(ts.URL + suite.base + "?" + q.Encode())
				require.NoError(t, err)

				assertAllowed(t, resp)
			})

			t.Run("case=returns denied", func(t *testing.T) {
				rt := testhelpers.APITupleFromString(t, "ns:foo#r@s")
				resp, err := ts.Client().Get(ts.URL + suite.base + "?" + rt.ToURLQuery().Encode())
				require.NoError(t, err)

				assertDenied(t, resp)
			})
		})
	}
}

func TestBatchCheckHandler(t *testing.T) {
	nspaces := []*namespace.Namespace{{Name: "ns"}}
	reg := driver.NewSqliteTestRegistry(t, driver.WithNamespaces(nspaces))
	h := check.NewHandler(reg)

	rt := testhelpers.APITupleFromString(t, "ns:o#r@s")
	testhelpers.MapAndInsertTuples(t, reg, rt)

	// REST setup.
	router := httprouterx.NewRouterPublic()
	h.RegisterReadRoutes(router)
	ts := httptest.NewServer(router)
	t.Cleanup(ts.Close)

	// gRPC setup.
	checkClient := newTestGRPCCheckClient(t, h)

	t.Run("case=REST: validates request format", func(t *testing.T) {
		t.Run("non-int max depth", func(t *testing.T) {
			resp, err := ts.Client().Post(ts.URL+check.BatchRoute+"?max-depth=foo",
				"application/json", nil)
			require.NoError(t, err)
			require.Equal(t, http.StatusBadRequest, resp.StatusCode)
			require.Contains(t, string(readBody(t, resp)), "invalid syntax")
		})

		t.Run("invalid JSON body", func(t *testing.T) {
			resp, err := ts.Client().Post(buildBatchURL(ts.URL, "5"),
				"application/json", strings.NewReader("not-json"))
			require.NoError(t, err)
			require.Equal(t, http.StatusBadRequest, resp.StatusCode)
			require.Contains(t, string(readBody(t, resp)), "could not unmarshal json")
		})
	})

	t.Run("case=validates max batch size", func(t *testing.T) {
		t.Run("protocol=REST", func(t *testing.T) {
			tuples := make([]client.Relationship, 11)
			for i := range tuples {
				tuples[i] = relationshipFromAPITuple(testhelpers.APITupleFromString(t, "ns:o#r@s"))
			}
			bodyBytes, err := json.Marshal(client.BatchCheckPermissionBody{Tuples: tuples})
			require.NoError(t, err)
			resp, err := ts.Client().Post(buildBatchURL(ts.URL, "5"), "application/json", bytes.NewReader(bodyBytes))
			require.NoError(t, err)
			require.Equal(t, http.StatusBadRequest, resp.StatusCode)
			require.Contains(t, string(readBody(t, resp)), "batch exceeds max size of 10")
		})

		t.Run("protocol=gRPC", func(t *testing.T) {
			tuples := make([]*rts.RelationTuple, 11)
			for i := range tuples {
				tuples[i] = testhelpers.APITupleFromString(t, "ns:o#r@s").ToProto()
			}
			_, err := checkClient.BatchCheck(t.Context(), &rts.BatchCheckRequest{Tuples: tuples, MaxDepth: 5})
			st, ok := status.FromError(err)
			require.True(t, ok)
			require.Equal(t, codes.InvalidArgument, st.Code())
			require.Equal(t, "batch exceeds max size of 10", st.Message())
		})
	})

	t.Run("case=returns correct results per tuple", func(t *testing.T) {
		// Three tuples: allowed, not-allowed, unknown namespace.
		restTuples := []client.Relationship{
			relationshipFromAPITuple(testhelpers.APITupleFromString(t, "ns:o#r@s")),  // allowed
			relationshipFromAPITuple(testhelpers.APITupleFromString(t, "ns:o2#r@s")), // not allowed
			relationshipFromAPITuple(testhelpers.APITupleFromString(t, "n2:o#r@s")),  // unknown namespace
		}
		grpcTuples := []*rts.RelationTuple{
			testhelpers.APITupleFromString(t, "ns:o#r@s").ToProto(),  // allowed
			testhelpers.APITupleFromString(t, "ns:o2#r@s").ToProto(), // not allowed
			testhelpers.APITupleFromString(t, "n2:o#r@s").ToProto(),  // unknown namespace
		}

		t.Run("protocol=REST", func(t *testing.T) {
			bodyBytes, err := json.Marshal(client.BatchCheckPermissionBody{Tuples: restTuples})
			require.NoError(t, err)
			resp, err := ts.Client().Post(buildBatchURL(ts.URL, "5"), "application/json", bytes.NewReader(bodyBytes))
			require.NoError(t, err)
			require.Equal(t, http.StatusOK, resp.StatusCode)
			var result client.BatchCheckPermissionResult
			require.NoError(t, json.Unmarshal(readBody(t, resp), &result))
			require.Equal(t, client.BatchCheckPermissionResult{
				Results: []client.CheckPermissionResultWithError{
					{Allowed: true},
					{Allowed: false},
					{Allowed: false, Error: new("The requested resource could not be found")},
				},
			}, result)
		})

		t.Run("protocol=gRPC", func(t *testing.T) {
			resp, err := checkClient.BatchCheck(t.Context(), &rts.BatchCheckRequest{Tuples: grpcTuples, MaxDepth: 5})
			require.NoError(t, err)
			require.Len(t, resp.Results, 3)
			require.True(t, resp.Results[0].Allowed)
			require.Empty(t, resp.Results[0].Error)
			require.False(t, resp.Results[1].Allowed)
			require.Empty(t, resp.Results[1].Error)
			require.False(t, resp.Results[2].Allowed)
			require.Equal(t, "The requested resource could not be found", resp.Results[2].Error)
		})
	})
}

// relationshipFromAPITuple converts a ketoapi.RelationTuple to a client.Relationship
// for use in REST batch check requests.
func relationshipFromAPITuple(tuple *ketoapi.RelationTuple) client.Relationship {
	rel := client.Relationship{
		Namespace: tuple.Namespace,
		Object:    tuple.Object,
		Relation:  tuple.Relation,
	}
	if tuple.SubjectID != nil {
		rel.SubjectId = tuple.SubjectID
	} else if tuple.SubjectSet != nil {
		rel.SubjectSet = &client.SubjectSet{
			Namespace: tuple.SubjectSet.Namespace,
			Object:    tuple.SubjectSet.Object,
			Relation:  tuple.SubjectSet.Relation,
		}
	}
	return rel
}

func TestCheckStrictModeLimit(t *testing.T) {
	depthLimitScenario.WithStrict(true).Run(t, func(t *testing.T, reg driver.Registry) {
		require.NoError(t, reg.Config(t.Context()).Set(config.KeyLimitMaxReadDepth, 2))
		h := check.NewHandler(reg)

		// This tuple has a depth of 3, which exceeds the configured limit of 2, causing to hit the limitation.
		tupleQuery := testhelpers.APITupleFromString(t, "Resource:doc#union@User:Alice").ToURLQuery()

		t.Run("case=HTTP: returns 422 with limitation kind", func(t *testing.T) {
			router := httprouterx.NewRouterPublic()
			h.RegisterReadRoutes(router)
			ts := httptest.NewServer(router)
			t.Cleanup(ts.Close)

			resp, err := ts.Client().Get(ts.URL + check.RouteBase + "?max-depth=2&" + tupleQuery.Encode())
			require.NoError(t, err)
			body := readBody(t, resp)

			require.Equal(t, http.StatusUnprocessableEntity, resp.StatusCode, "%s", body)
			require.Equal(t, "The check could not be completed", gjson.GetBytes(body, "error.message").String())
			require.Contains(t, gjson.GetBytes(body, "error.reason").String(), string(check.LimitationMaxDepthExceeded))
		})

		t.Run("case=gRPC: returns FailedPrecondition with limitation kind", func(t *testing.T) {
			checkClient := newTestGRPCCheckClient(t, h)
			tuple := testhelpers.APITupleFromString(t, "Resource:doc#union@User:Alice")
			_, err := checkClient.Check(t.Context(), &rts.CheckRequest{
				Tuple:    tuple.ToProto(),
				MaxDepth: 2,
			})
			require.Error(t, err)

			st, ok := status.FromError(err)
			require.True(t, ok)
			require.Equal(t, codes.FailedPrecondition, st.Code())
			require.Equal(t, "The check could not be completed", st.Message())

			var ei *errdetails.ErrorInfo
			for _, d := range st.Details() {
				if info, ok := d.(*errdetails.ErrorInfo); ok {
					ei = info
					break
				}
			}
			require.NotNil(t, ei)
			require.Contains(t, ei.Reason, string(check.LimitationMaxDepthExceeded))
		})
	})
}

func TestBatchCheckLimitHandler(t *testing.T) {
	for _, scenario := range []testhelpers.Scenario{
		widthLimitScenario,
		widthLimitScenario.WithStrict(true),
	} {
		scenario.Run(t, func(t *testing.T, reg driver.Registry) {
			require.NoError(t, reg.Config(t.Context()).Set(config.KeyLimitMaxReadWidth, 1))
			h := check.NewHandler(reg)
			strict := scenario.Strict

			// Bob is only in the truncated branch → MembershipUnknown at width=1.
			tuple := testhelpers.APITupleFromString(t, "Resource:doc#readers@User:Bob")

			t.Run("case=REST: strict mode surfaces limitation in error field", func(t *testing.T) {
				router := httprouterx.NewRouterPublic()
				h.RegisterReadRoutes(router)
				ts := httptest.NewServer(router)
				t.Cleanup(ts.Close)

				bodyBytes, err := json.Marshal(client.BatchCheckPermissionBody{
					Tuples: []client.Relationship{relationshipFromAPITuple(tuple)},
				})
				require.NoError(t, err)

				resp, err := ts.Client().Post(
					fmt.Sprintf("%s%s?max-depth=10", ts.URL, check.BatchRoute),
					"application/json", bytes.NewReader(bodyBytes))
				require.NoError(t, err)
				require.Equal(t, http.StatusOK, resp.StatusCode)

				var result client.BatchCheckPermissionResult
				require.NoError(t, json.Unmarshal(readBody(t, resp), &result))
				require.Len(t, result.Results, 1)
				require.False(t, result.Results[0].Allowed)
				if strict {
					require.NotNil(t, result.Results[0].Error)
					require.Contains(t, *result.Results[0].Error, string(check.LimitationMaxWidthExceeded))
				} else {
					require.Empty(t, result.Results[0].Error)
				}
			})

			t.Run("case=gRPC: strict mode surfaces limitation in error field", func(t *testing.T) {
				checkClient := newTestGRPCCheckClient(t, h)
				resp, err := checkClient.BatchCheck(t.Context(), &rts.BatchCheckRequest{
					Tuples:   []*rts.RelationTuple{tuple.ToProto()},
					MaxDepth: 10,
				})
				require.NoError(t, err)
				require.Len(t, resp.Results, 1)
				require.False(t, resp.Results[0].Allowed)
				if strict {
					require.Contains(t, resp.Results[0].Error, string(check.LimitationMaxWidthExceeded))
				} else {
					require.Empty(t, resp.Results[0].Error)
				}
			})
		})
	}
}
