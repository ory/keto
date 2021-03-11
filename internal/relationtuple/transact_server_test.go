package relationtuple_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/julienschmidt/httprouter"

	"github.com/ory/keto/internal/driver"
	"github.com/ory/keto/internal/namespace"
	"github.com/ory/keto/internal/relationtuple"
	"github.com/ory/keto/internal/x"
)

func TestWriteHandlers(t *testing.T) {
	r := httprouter.New()
	wr := &x.WriteRouter{Router: r}
	rr := &x.ReadRouter{Router: r}
	reg := driver.NewMemoryTestRegistry(t, []*namespace.Namespace{{Name: "handler test"}})
	h := relationtuple.NewHandler(reg)
	h.RegisterWriteRoutes(wr)
	h.RegisterReadRoutes(rr)
	ts := httptest.NewServer(r)
	defer ts.Close()

	t.Run("method=create", func(t *testing.T) {
		create := func(raw []byte) *http.Response {
			req, err := http.NewRequest(http.MethodPut, ts.URL+relationtuple.RouteBase, bytes.NewBuffer(raw))
			require.NoError(t, err)
			resp, err := ts.Client().Do(req)
			require.NoError(t, err)

			return resp
		}

		t.Run("case=creates tuple", func(t *testing.T) {
			rt := &relationtuple.InternalRelationTuple{
				Namespace: "handler test",
				Object:    "obj",
				Relation:  "rel",
				Subject:   &relationtuple.SubjectID{ID: "subj"},
			}
			payload, err := json.Marshal(rt)
			require.NoError(t, err)

			resp := create(payload)

			assert.Equal(t, http.StatusCreated, resp.StatusCode)

			body, err := io.ReadAll(resp.Body)
			require.NoError(t, err)

			assert.JSONEq(t, string(payload), string(body))

			t.Run("check=is contained in the manager", func(t *testing.T) {
				// set a size > 1 just to make sure it gets all
				actualRTs, _, err := reg.RelationTupleManager().GetRelationTuples(context.Background(), (*relationtuple.RelationQuery)(rt), x.WithSize(10))
				require.NoError(t, err)
				assert.Equal(t, []*relationtuple.InternalRelationTuple{rt}, actualRTs)
			})

			t.Run("check=is gettable with the returned URL", func(t *testing.T) {
				resp, err := ts.Client().Get(ts.URL + resp.Header.Get("Location"))
				require.NoError(t, err)
				require.Equal(t, http.StatusOK, resp.StatusCode)

				respDec := relationtuple.GetResponse{}
				require.NoError(t, json.NewDecoder(resp.Body).Decode(&respDec))
				assert.Equal(t, []*relationtuple.InternalRelationTuple{rt}, respDec.RelationTuples)
			})
		})

		t.Run("case=returns bad request on JSON parse error", func(t *testing.T) {
			resp := create([]byte("foo"))
			assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
		})
	})
}
