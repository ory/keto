package relationtuple_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ory/keto/internal/driver"
	"github.com/ory/keto/internal/namespace"
	"github.com/ory/keto/internal/relationtuple"
	"github.com/ory/keto/internal/x"
)

func TestWriteHandlers(t *testing.T) {
	t.Run("method=create", func(t *testing.T) {
		r := &x.WriteRouter{Router: httprouter.New()}
		reg := driver.NewMemoryTestRegistry(t, []*namespace.Namespace{{Name: "handler test"}})
		h := relationtuple.NewHandler(reg)
		h.RegisterWriteRoutes(r)
		ts := httptest.NewServer(r)
		defer ts.Close()

		c := ts.Client()

		rt := &relationtuple.InternalRelationTuple{
			Namespace: "handler test",
			Object:    "obj",
			Relation:  "rel",
			Subject:   &relationtuple.SubjectID{ID: "subj"},
		}
		payload, err := json.Marshal(rt)
		require.NoError(t, err)

		req, err := http.NewRequest(http.MethodPut, ts.URL+relationtuple.RouteBase, bytes.NewBuffer(payload))
		require.NoError(t, err)
		resp, err := c.Do(req)
		require.NoError(t, err)

		body, err := io.ReadAll(resp.Body)
		require.NoError(t, err)

		assert.JSONEq(t, string(payload), string(body))

		// set a size just to make sure it gets all
		actualRTs, _, err := reg.RelationTupleManager().GetRelationTuples(context.Background(), (*relationtuple.RelationQuery)(rt), x.WithSize(1000))
		require.NoError(t, err)
		assert.Equal(t, []*relationtuple.InternalRelationTuple{rt}, actualRTs)
	})
}
