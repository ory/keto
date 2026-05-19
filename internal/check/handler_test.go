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
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/test/bufconn"

	"github.com/ory/x/httprouterx"

	"github.com/ory/keto/internal/check"
	"github.com/ory/keto/internal/driver"
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
		Name: "check handler",
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
				resp, err := ts.Client().Get(ts.URL + suite.base + "?" + (&ketoapi.RelationTuple{
					Namespace: "not" + nspaces[0].Name,
					Object:    "o",
					Relation:  "r",
					SubjectID: new("s"),
				}).ToURLQuery().Encode())
				require.NoError(t, err)

				assertDenied(t, resp)
			})

			t.Run("case=returns allowed", func(t *testing.T) {
				rt := &ketoapi.RelationTuple{
					Namespace: nspaces[0].Name,
					Object:    "o",
					Relation:  "r",
					SubjectID: new("s"),
				}
				testhelpers.MapAndInsertTuples(t, reg, rt)

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
					SubjectID: new("s"),
				}).ToURLQuery().Encode())
				require.NoError(t, err)

				assertDenied(t, resp)
			})
		})
	}
}

func TestBatchCheckHandler(t *testing.T) {
	nspaces := []*namespace.Namespace{{Name: "batch-check"}}
	reg := driver.NewSqliteTestRegistry(t, driver.WithNamespaces(nspaces))
	h := check.NewHandler(reg)

	rt := &ketoapi.RelationTuple{
		Namespace: nspaces[0].Name,
		Object:    "o",
		Relation:  "r",
		SubjectID: new("s"),
	}
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
				tuples[i] = client.Relationship{Namespace: "n", Object: "o", Relation: "r", SubjectId: new("s")}
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
				tuples[i] = &rts.RelationTuple{
					Namespace: "n", Object: "o", Relation: "r",
					Subject: &rts.Subject{Ref: &rts.Subject_Id{Id: "s"}},
				}
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
			{Namespace: nspaces[0].Name, Object: "o", Relation: "r", SubjectId: new("s")},  // allowed
			{Namespace: nspaces[0].Name, Object: "o2", Relation: "r", SubjectId: new("s")}, // not allowed
			{Namespace: "n2", Object: "o", Relation: "r", SubjectId: new("s")},             // unknown namespace
		}
		grpcTuples := []*rts.RelationTuple{
			{Namespace: nspaces[0].Name, Object: "o", Relation: "r", Subject: &rts.Subject{Ref: &rts.Subject_Id{Id: "s"}}},  // allowed
			{Namespace: nspaces[0].Name, Object: "o2", Relation: "r", Subject: &rts.Subject{Ref: &rts.Subject_Id{Id: "s"}}}, // not allowed
			{Namespace: "n2", Object: "o", Relation: "r", Subject: &rts.Subject{Ref: &rts.Subject_Id{Id: "s"}}},             // unknown namespace
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
