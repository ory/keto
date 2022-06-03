package relationtuple

import (
	"context"
	"github.com/gofrs/uuid"
	"github.com/ory/keto/ketoapi"
	"testing"

	rts "github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2"

	"github.com/ory/herodot"

	"github.com/ory/keto/internal/x"
)

type (
	ManagerProvider interface {
		RelationTupleManager() Manager
	}
	Manager interface {
		GetRelationTuples(ctx context.Context, query *RelationQuery, options ...x.PaginationOptionSetter) ([]*InternalRelationTuple, string, error)
		WriteRelationTuples(ctx context.Context, rs ...*InternalRelationTuple) error
		DeleteRelationTuples(ctx context.Context, rs ...*InternalRelationTuple) error
		DeleteAllRelationTuples(ctx context.Context, query *RelationQuery) error
		TransactRelationTuples(ctx context.Context, insert []*InternalRelationTuple, delete []*InternalRelationTuple) error
	}
	SubjectID struct {
		ID uuid.UUID `json:"id"`
	}
	RelationQuery struct {
		Namespace  *int32      `json:"namespace"`
		Object     *uuid.UUID  `json:"object"`
		Relation   *string     `json:"relation"`
		SubjectID  *uuid.UUID  `json:"subject_id,omitempty"`
		SubjectSet *SubjectSet `json:"subject_set,omitempty"`
	}
	TupleData interface {
		GetSubject() *rts.Subject
		GetObject() string
		GetNamespace() string
		GetRelation() string
	}
	Subject interface {
		Equals(Subject) bool
	}
	InternalRelationTuple struct {
		Namespace int32     `json:"namespace"`
		Object    uuid.UUID `json:"object"`
		Relation  string    `json:"relation"`
		Subject   Subject   `json:"subject"`
	}
	InternalRelationTuples []*InternalRelationTuple
	SubjectSet             struct {
		Namespace int32     `json:"namespace"`
		Object    uuid.UUID `json:"object"`
		Relation  string    `json:"relation"`
	}
)

var (
	_, _ Subject = &SubjectID{}, &SubjectSet{}

	ErrMalformedInput    = herodot.ErrBadRequest.WithError("malformed string input")
	ErrNilSubject        = herodot.ErrBadRequest.WithError("subject is not allowed to be nil")
	ErrDuplicateSubject  = herodot.ErrBadRequest.WithError("exactly one of subject_set or subject_id has to be provided")
	ErrDroppedSubjectKey = herodot.ErrBadRequest.WithDebug(`provide "subject_id" or "subject_set.*"; support for "subject" was dropped`)
	ErrIncompleteSubject = herodot.ErrBadRequest.WithError(`incomplete subject, provide "subject_id" or a complete "subject_set.*"`)
)

func (s *SubjectID) Equals(v interface{}) bool {
	uv, ok := v.(*SubjectID)
	if !ok {
		return false
	}
	return uv.ID == s.ID
}

func (s *SubjectSet) Equals(v interface{}) bool {
	uv, ok := v.(*SubjectSet)
	if !ok {
		return false
	}
	return uv.Relation == s.Relation && uv.Object == s.Object && uv.Namespace == s.Namespace
}

func (q *RelationQuery) Subject() Subject {
	if q.SubjectID != nil {
		return &SubjectID{ID: *q.SubjectID}
	} else if q.SubjectSet != nil {
		return q.SubjectSet
	}
	return nil
}

type ManagerWrapper struct {
	Reg            ManagerProvider
	PageOpts       []x.PaginationOptionSetter
	RequestedPages []string
}

var (
	_ Manager         = (*ManagerWrapper)(nil)
	_ ManagerProvider = (*ManagerWrapper)(nil)
)

func NewManagerWrapper(_ *testing.T, reg ManagerProvider, options ...x.PaginationOptionSetter) *ManagerWrapper {
	return &ManagerWrapper{
		Reg:      reg,
		PageOpts: options,
	}
}

func (t *ManagerWrapper) GetRelationTuples(ctx context.Context, query *RelationQuery, options ...x.PaginationOptionSetter) ([]*InternalRelationTuple, string, error) {
	opts := x.GetPaginationOptions(options...)
	t.RequestedPages = append(t.RequestedPages, opts.Token)
	return t.Reg.RelationTupleManager().GetRelationTuples(ctx, query, append(t.PageOpts, options...)...)
}

func (t *ManagerWrapper) WriteRelationTuples(ctx context.Context, rs ...*InternalRelationTuple) error {
	return t.Reg.RelationTupleManager().WriteRelationTuples(ctx, rs...)
}

func (t *ManagerWrapper) DeleteRelationTuples(ctx context.Context, rs ...*InternalRelationTuple) error {
	return t.Reg.RelationTupleManager().DeleteRelationTuples(ctx, rs...)
}

func (t *ManagerWrapper) DeleteAllRelationTuples(ctx context.Context, query *RelationQuery) error {
	return t.Reg.RelationTupleManager().DeleteAllRelationTuples(ctx, query)
}

func (t *ManagerWrapper) TransactRelationTuples(ctx context.Context, insert []*InternalRelationTuple, delete []*InternalRelationTuple) error {
	return t.Reg.RelationTupleManager().TransactRelationTuples(ctx, insert, delete)
}

func (t *ManagerWrapper) RelationTupleManager() Manager {
	return t
}

type Tree struct {
	Type     ketoapi.ExpandNodeType `json:"type"`
	Subject  Subject                `json:"subject"`
	Children []*Tree                `json:"children,omitempty"`
}
