package relationtuple

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/pkg/errors"

	"github.com/ory/herodot"

	"github.com/ory/keto/internal/x"
)

type (
	handlerDependencies interface {
		ManagerProvider
		x.LoggerProvider
		x.WriterProvider
	}
	handler struct {
		d handlerDependencies
	}
)

const (
	routeBase = "/relations"
)

func NewHandler(d handlerDependencies) *handler {
	return &handler{
		d: d,
	}
}

func (h *handler) RegisterPublicRoutes(router *httprouter.Router) {
	router.GET(routeBase+"/:namespace", h.getRelations)
	router.PUT(routeBase, h.createRelation)
}

func (h *handler) getRelations(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	qParams := r.URL.Query()
	query := &RelationQuery{
		Relation:  qParams.Get("relation"),
		ObjectID:  qParams.Get("object_id"),
		Namespace: params.ByName("namespace"),
	}
	if sub := qParams.Get("subject"); sub != "" {
		query.Subject = SubjectFromString(sub)
	}

	res, err := h.d.RelationTupleManager().GetRelationTuples(r.Context(), query)
	if err != nil {
		h.d.Writer().WriteError(w, r, err)
		return
	}

	h.d.Writer().Write(w, r, res)
}

func (h *handler) createRelation(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var rel InternalRelationTuple

	if err := json.NewDecoder(r.Body).Decode(&rel); err != nil {
		fmt.Printf("json decode error: %+v\n", err)
		h.d.Writer().WriteError(w, r, errors.WithStack(herodot.ErrBadRequest))
		return
	}

	if err := h.d.RelationTupleManager().WriteRelationTuples(r.Context(), &rel); err != nil {
		h.d.Writer().WriteError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
