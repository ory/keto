package storage

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ory/herodot"
	"github.com/ory/x/pagination"
)

func TestCRUD(t *testing.T) {
	for k, m := range map[string]Manager{
		"memory": NewMemoryManager(),
	} {
		t.Run(fmt.Sprintf("manager=%s", k), func(t *testing.T) {
			h := NewHandler(m, herodot.NewJSONWriter(nil))
			i := &mockHandler{c: "tests", sh: h}
			r := httprouter.New()
			i.Register(r)
			ts := httptest.NewServer(r)
			defer ts.Close()

			t.Run("case=404", func(t *testing.T) {
				res, err := ts.Client().Get(ts.URL + "/1234")
				require.NoError(t, err)
				res.Body.Close()
				assert.Equal(t, http.StatusNotFound, res.StatusCode)
			})

			t.Run("case=create", func(t *testing.T) {
				res, err := ts.Client().Post(ts.URL+"/?key=1234&value=bar", "application/json", bytes.NewBuffer(nil))
				require.NoError(t, err)
				res.Body.Close()
				assert.Equal(t, http.StatusOK, res.StatusCode)
			})

			t.Run("case=get", func(t *testing.T) {
				res, err := ts.Client().Get(ts.URL + "/1234")
				require.NoError(t, err)
				assert.Equal(t, http.StatusOK, res.StatusCode)
				b, err := ioutil.ReadAll(res.Body)
				require.NoError(t, err)
				res.Body.Close()
				assert.Equal(t, `"bar"`, string(b))
			})

			t.Run("case=list (all)", func(t *testing.T) {
				res, err := ts.Client().Get(ts.URL + "/")
				require.NoError(t, err)
				assert.Equal(t, http.StatusOK, res.StatusCode)
				b, err := ioutil.ReadAll(res.Body)
				require.NoError(t, err)
				res.Body.Close()
				assert.Equal(t, `["bar"]`, string(b))
			})

			t.Run("case=list (valid member)", func(t *testing.T) {
				res, err := ts.Client().Get(ts.URL + "/?member=1234")
				require.NoError(t, err)
				assert.Equal(t, http.StatusOK, res.StatusCode)
				b, err := ioutil.ReadAll(res.Body)
				require.NoError(t, err)
				res.Body.Close()
				assert.Equal(t, `["bar"]`, string(b))
			})

			t.Run("case=list (invalid member)", func(t *testing.T) {
				res, err := ts.Client().Get(ts.URL + "/?member=124")
				require.NoError(t, err)
				assert.Equal(t, http.StatusOK, res.StatusCode)
				b, err := ioutil.ReadAll(res.Body)
				require.NoError(t, err)
				res.Body.Close()
				assert.Equal(t, `[""]`, string(b))
			})

			t.Run("case=delete", func(t *testing.T) {
				req, err := http.NewRequest("DELETE", ts.URL+"/1234", nil)
				require.NoError(t, err)
				res, err := ts.Client().Do(req)
				require.NoError(t, err)
				res.Body.Close()
				assert.Equal(t, http.StatusNoContent, res.StatusCode)
			})

			t.Run("case=list", func(t *testing.T) {
				res, err := ts.Client().Get(ts.URL + "/")
				require.NoError(t, err)
				assert.Equal(t, http.StatusOK, res.StatusCode)
				b, err := ioutil.ReadAll(res.Body)
				require.NoError(t, err)
				res.Body.Close()
				assert.Equal(t, `[]`, string(b))
			})

		})
	}
}

type mockHandler struct {
	c  string
	sh *Handler
}

func (e *mockHandler) Register(r *httprouter.Router) {
	r.POST("/", e.sh.Upsert(e.create))
	r.GET("/", e.List(e.list))
	r.GET("/:id", e.sh.Get(e.get))
	r.DELETE("/:id", e.sh.Delete(e.delete))
}

func (e *mockHandler) create(ctx context.Context, r *http.Request, ps httprouter.Params) (*UpsertRequest, error) {
	return &UpsertRequest{
		Collection: e.c,
		Key:        r.URL.Query().Get("key"),
		Value:      r.URL.Query().Get("value"),
	}, nil
}

func (e *mockHandler) list(ctx context.Context, r *http.Request, ps httprouter.Params) (ListRequest, error) {
	var p []string
	var listReqeust ListRequest
	member := r.URL.Query().Get("member")
	if member != "" {
		listReqeust = &ListRequestByMember{
			Collection: e.c,
			Member:     member,
			Value:      &p,
		}
	} else {
		listReqeust = &ListRequestAllMembers{
			Collection: e.c,
			Value:      &p,
		}
	}
	return listReqeust.MakeRequest(), nil
}

func (e *mockHandler) List(factory func(context.Context, *http.Request, httprouter.Params) (ListRequest, error)) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		ctx := r.Context()
		l, err := factory(ctx, r, ps)
		if err != nil {
			e.sh.h.WriteError(w, r, err)
			return
		}
		limit, offset := pagination.Parse(r, 100, 0, 500)

		switch t := l.(type) {
		case *ListRequestByMember:
			i := t
			if err := e.sh.s.ListByMember(ctx, i.Collection, i.Value, i.Member, limit, offset); err != nil {
				e.sh.h.WriteError(w, r, err)
				return
			}
			e.sh.h.Write(w, r, i.Value)
		case *ListRequestAllMembers:
			i := t
			if err := e.sh.s.List(ctx, i.Collection, i.Value, limit, offset); err != nil {
				e.sh.h.WriteError(w, r, err)
				return
			}
			e.sh.h.Write(w, r, i.Value)
		}
	}
}

func (e *mockHandler) delete(ctx context.Context, r *http.Request, ps httprouter.Params) (*DeleteRequest, error) {
	return &DeleteRequest{
		Collection: e.c,
		Key:        ps.ByName("id"),
	}, nil
}

func (e *mockHandler) get(ctx context.Context, r *http.Request, ps httprouter.Params) (*GetRequest, error) {
	var p string
	return &GetRequest{
		Collection: e.c,
		Key:        ps.ByName("id"),
		Value:      &p,
	}, nil
}
