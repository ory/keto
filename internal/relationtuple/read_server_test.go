// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package relationtuple_test

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/ory/x/pointerx"

	rts "github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2"

	"github.com/ory/keto/ketoapi"

	"github.com/ory/keto/internal/driver/config"

	"github.com/stretchr/testify/assert"
	"github.com/tidwall/gjson"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/require"

	"github.com/ory/keto/internal/driver"
	"github.com/ory/keto/internal/namespace"
	"github.com/ory/keto/internal/relationtuple"
	"github.com/ory/keto/internal/x"
)

func TestReadHandlers(t *testing.T) {
	ctx := context.Background()
	r := &x.ReadRouter{Router: httprouter.New()}
	reg := driver.NewSqliteTestRegistry(t, false)
	h := relationtuple.NewHandler(reg)
	h.RegisterReadRoutes(r)
	ts := httptest.NewServer(r)
	t.Cleanup(ts.Close)

	var newNamespace func(*testing.T) *namespace.Namespace
	{
		nspaces := 0
		newNamespace = func(t *testing.T) *namespace.Namespace {
			n := &namespace.Namespace{Name: fmt.Sprintf("relation tuple read test %d", nspaces)}
			nspaces++
			require.NoError(t, reg.Config(ctx).Set(config.KeyNamespaces, []*namespace.Namespace{n}))
			return n
		}
	}

	t.Run("method=get", func(t *testing.T) {
		t.Run("case=empty response is not nil", func(t *testing.T) {
			nspace := newNamespace(t)
			resp, err := ts.Client().Get(ts.URL + relationtuple.ReadRouteBase + "?" + url.Values{
				"namespace": {nspace.Name},
			}.Encode())
			require.NoError(t, err)
			require.Equal(t, http.StatusOK, resp.StatusCode)

			body, err := io.ReadAll(resp.Body)
			require.NoError(t, err)

			assert.Equal(t, "[]", gjson.GetBytes(body, "relation_tuples").Raw)

			var respMsg ketoapi.GetResponse
			require.NoError(t, json.Unmarshal(body, &respMsg))

			assert.Equal(t, ketoapi.GetResponse{
				RelationTuples: []*ketoapi.RelationTuple{},
				NextPageToken:  "",
			}, respMsg)
		})

		t.Run("case=returns tuples", func(t *testing.T) {
			nspace := newNamespace(t)
			tuples := []*ketoapi.RelationTuple{
				{
					Namespace: nspace.Name,
					Object:    "o1",
					Relation:  "r1",
					SubjectID: pointerx.Ptr("s1"),
				},
				{
					Namespace: nspace.Name,
					Object:    "o2",
					Relation:  "r2",
					SubjectSet: &ketoapi.SubjectSet{
						Namespace: nspace.Name,
						Object:    "o1",
						Relation:  "r1",
					},
				},
			}

			relationtuple.MapAndWriteTuples(t, reg, tuples...)

			resp, err := ts.Client().Get(ts.URL + relationtuple.ReadRouteBase + "?" + url.Values{
				"namespace": {nspace.Name},
			}.Encode())
			require.NoError(t, err)

			var respMsg ketoapi.GetResponse
			require.NoError(t, json.NewDecoder(resp.Body).Decode(&respMsg))

			assert.ElementsMatch(t, tuples, respMsg.RelationTuples)
			assert.Equal(t, "", respMsg.NextPageToken)
		})

		t.Run("case=return tuples without namespace", func(t *testing.T) {
			nspace := newNamespace(t)

			tuples := []*ketoapi.RelationTuple{
				{
					Namespace: nspace.Name,
					Object:    "obj",
					Relation:  "r1",
					SubjectID: pointerx.Ptr("s1"),
				},
			}

			relationtuple.MapAndWriteTuples(t, reg, tuples...)

			resp, err := ts.Client().Get(ts.URL + relationtuple.ReadRouteBase + "?" + url.Values{
				"object": {"obj"},
			}.Encode())
			require.NoError(t, err)
			assert.Equal(t, resp.StatusCode, http.StatusOK)

			var respMsg ketoapi.GetResponse
			require.NoError(t, json.NewDecoder(resp.Body).Decode(&respMsg))
			assert.Equal(t, 1, len(respMsg.RelationTuples))
			assert.Containsf(t, tuples, respMsg.RelationTuples[0], "expected to find %q in %q", respMsg.RelationTuples[0].String(), tuples)
			assert.Equal(t, "", respMsg.NextPageToken)
		})

		t.Run("case=returns bad request on malformed subject", func(t *testing.T) {
			resp, err := ts.Client().Get(ts.URL + relationtuple.ReadRouteBase + "?" + url.Values{
				"subject": {"not#a valid subject"},
			}.Encode())
			require.NoError(t, err)

			assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
		})

		t.Run("case=paginates", func(t *testing.T) {
			nspace := newNamespace(t)

			tuples := []*ketoapi.RelationTuple{
				{
					Namespace: nspace.Name,
					Object:    "o1",
					Relation:  "r1",
					SubjectID: pointerx.Ptr("s1"),
				},
				{
					Namespace: nspace.Name,
					Object:    "o2",
					Relation:  "r2",
					SubjectID: pointerx.Ptr("s2"),
				},
			}
			relationtuple.MapAndWriteTuples(t, reg, tuples...)

			var firstResp ketoapi.GetResponse
			resp, err := ts.Client().Get(ts.URL + relationtuple.ReadRouteBase + "?" + url.Values{
				"namespace": {nspace.Name},
				"page_size": {"1"},
			}.Encode())
			require.NoError(t, err)
			require.Equal(t, http.StatusOK, resp.StatusCode)

			require.NoError(t, json.NewDecoder(resp.Body).Decode(&firstResp))
			require.Len(t, firstResp.RelationTuples, 1)
			assert.Contains(t, tuples, firstResp.RelationTuples[0])
			assert.NotEqual(t, "", firstResp.NextPageToken)

			// second page
			resp, err = ts.Client().Get(ts.URL + relationtuple.ReadRouteBase + "?" + url.Values{
				"namespace":  {nspace.Name},
				"page_size":  {"1"},
				"page_token": {firstResp.NextPageToken},
			}.Encode())
			require.NoError(t, err)
			require.Equal(t, http.StatusOK, resp.StatusCode)

			secondResp := ketoapi.GetResponse{}
			require.NoError(t, json.NewDecoder(resp.Body).Decode(&secondResp))
			require.Len(t, secondResp.RelationTuples, 1)

			assert.NotEqual(t, firstResp.RelationTuples, secondResp.RelationTuples)
			assert.Contains(t, tuples, secondResp.RelationTuples[0])
			assert.Equal(t, "", secondResp.NextPageToken)
		})

		t.Run("case=returs bad request on invalid page size", func(t *testing.T) {
			resp, err := ts.Client().Get(ts.URL + relationtuple.ReadRouteBase + "?" + url.Values{
				"page_size": {"foo"},
			}.Encode())
			require.NoError(t, err)

			body, err := io.ReadAll(resp.Body)
			require.NoError(t, err)
			assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
			assert.Contains(t, string(body), "invalid syntax")
		})
	})

	t.Run("method=grpc", func(t *testing.T) {
		type requestEnhancer = func(req *rts.ListRelationTuplesRequest, query *ketoapi.RelationQuery)
		withRelationQuery := func(req *rts.ListRelationTuplesRequest, query *ketoapi.RelationQuery) {
			req.RelationQuery = query.ToProto()
		}
		withDeprecatedQuery := func(req *rts.ListRelationTuplesRequest, query *ketoapi.RelationQuery) {
			pq := query.ToProto()
			//nolint:staticcheck
			req.Query = &rts.ListRelationTuplesRequest_Query{ //lint:ignore SA1019 backwards compatibility
				Subject: pq.Subject,
			}
			if pq.Namespace != nil {
				//nolint:staticcheck
				req.Query.Namespace = *pq.Namespace //lint:ignore SA1019 backwards compatibility
			}
			if pq.Object != nil {
				//nolint:staticcheck
				req.Query.Object = *pq.Object //lint:ignore SA1019 backwards compatibility
			}
			if pq.Relation != nil {
				//nolint:staticcheck
				req.Query.Relation = *pq.Relation //lint:ignore SA1019 backwards compatibility
			}
		}
		apiTuplesFromProto := func(t *testing.T, pts ...*rts.RelationTuple) []*ketoapi.RelationTuple {
			actual := make([]*ketoapi.RelationTuple, len(pts))
			for i, rt := range pts {
				var err error
				actual[i], err = (&ketoapi.RelationTuple{}).FromDataProvider(rt)
				require.NoError(t, err)
			}
			return actual
		}
		soc, err := net.Listen("tcp", ":0") // nolint
		require.NoError(t, err)
		srv := grpc.NewServer()
		h.RegisterReadGRPC(srv)
		go srv.Serve(soc) // nolint
		t.Cleanup(srv.Stop)

		con, err := grpc.Dial(soc.Addr().String(), grpc.WithTransportCredentials(insecure.NewCredentials()))
		require.NoError(t, err)

		t.Run("method=list", func(t *testing.T) {
			client := rts.NewReadServiceClient(con)

			for key, enhancer := range map[string]requestEnhancer{"relation query": withRelationQuery, "deprecated query": withDeprecatedQuery} {
				t.Run("enhancer="+key, func(t *testing.T) {
					t.Run("case=returns empty list on no tuples", func(t *testing.T) {
						nspace := newNamespace(t)
						req := &rts.ListRelationTuplesRequest{}
						enhancer(req, &ketoapi.RelationQuery{
							Namespace: &nspace.Name,
						})
						resp, err := client.ListRelationTuples(ctx, req)
						require.NoError(t, err)
						assert.Len(t, resp.RelationTuples, 0)
					})

					t.Run("case=gets tuples", func(t *testing.T) {
						nspace := newNamespace(t)
						tuples := []*ketoapi.RelationTuple{
							{
								Namespace: nspace.Name,
								Object:    "o1",
								Relation:  "rel",
								SubjectID: pointerx.Ptr("s1"),
							},
							{
								Namespace: nspace.Name,
								Object:    "o2",
								Relation:  "rel",
								SubjectSet: &ketoapi.SubjectSet{
									Namespace: nspace.Name,
									Object:    "o1",
									Relation:  "r1",
								},
							},
						}
						relationtuple.MapAndWriteTuples(t, reg, tuples...)

						req := &rts.ListRelationTuplesRequest{}
						enhancer(req, &ketoapi.RelationQuery{
							Namespace: &nspace.Name,
						})

						resp, err := client.ListRelationTuples(ctx, req)
						require.NoError(t, err)
						assert.Len(t, resp.RelationTuples, 2)

						assert.ElementsMatch(t, tuples, apiTuplesFromProto(t, resp.RelationTuples...))
					})

					t.Run("case=paginates", func(t *testing.T) {
						nspace := newNamespace(t)
						tuples := []*ketoapi.RelationTuple{
							{
								Namespace: nspace.Name,
								Object:    "o1",
								Relation:  "rel",
								SubjectID: pointerx.Ptr("s1"),
							},
							{
								Namespace: nspace.Name,
								Object:    "o2",
								Relation:  "rel",
								SubjectID: pointerx.Ptr("s2"),
							},
							{
								Namespace: nspace.Name,
								Object:    "o3",
								Relation:  "rel",
								SubjectID: pointerx.Ptr("s3"),
							},
						}
						relationtuple.MapAndWriteTuples(t, reg, tuples...)

						query := &ketoapi.RelationQuery{
							Namespace: &nspace.Name,
						}
						firstReq := &rts.ListRelationTuplesRequest{}
						enhancer(firstReq, query)
						firstReq.PageSize = int32(2)

						firstResp, err := client.ListRelationTuples(ctx, firstReq)
						require.NoError(t, err)

						secondReq := &rts.ListRelationTuplesRequest{}
						enhancer(secondReq, query)
						secondReq.PageSize = int32(2)
						secondReq.PageToken = firstResp.NextPageToken

						secondResp, err := client.ListRelationTuples(ctx, secondReq)
						require.NoError(t, err)

						assert.Len(t, firstResp.RelationTuples, 2)
						assert.Len(t, secondResp.RelationTuples, 1)
						assert.Zero(t, secondResp.NextPageToken)

						assert.ElementsMatch(t, tuples, apiTuplesFromProto(t, append(firstResp.RelationTuples, secondResp.RelationTuples...)...))
					})
				})
			}
		})
	})
}
