package relationtuple

import (
	"context"
	"testing"

	"github.com/ory/keto/ketoapi"

	"github.com/gofrs/uuid"

	rts "github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2"

	"github.com/ory/keto/internal/x"
)

type (
	ManagerProvider interface {
		RelationTupleManager() Manager
	}
	Manager interface {
		GetRelationTuples(ctx context.Context, query *RelationQuery, options ...x.PaginationOptionSetter) ([]*RelationTuple, string, error)
		WriteRelationTuples(ctx context.Context, rs ...*RelationTuple) error
		DeleteRelationTuples(ctx context.Context, rs ...*RelationTuple) error
		DeleteAllRelationTuples(ctx context.Context, query *RelationQuery) error
		TransactRelationTuples(ctx context.Context, insert []*RelationTuple, delete []*RelationTuple) error
	}
	SubjectID struct {
		ID uuid.UUID `json:"id"`
	}
	RelationQuery struct {
		Namespace *string    `json:"namespace"`
		Object    *uuid.UUID `json:"object"`
		Relation  *string    `json:"relation"`
		Subject   Subject    `json:"subject_id,omitempty"`
	}
	TupleData interface {
		GetSubject() *rts.Subject
		GetObject() string
		GetNamespace() string
		GetRelation() string
	}
	Subject interface {
		Equals(Subject) bool
		UniqueID() uuid.UUID
	}
	RelationTuple struct {
		Namespace string    `json:"namespace"`
		Object    uuid.UUID `json:"object"`
		Relation  string    `json:"relation"`
		Subject   Subject   `json:"subject"`
	}
	InternalRelationTuples []*RelationTuple
	SubjectSet             struct {
		Namespace string    `json:"namespace"`
		Object    uuid.UUID `json:"object"`
		Relation  string    `json:"relation"`
	}
	Tree struct {
		Type     ketoapi.ExpandNodeType `json:"type"`
		Subject  Subject                `json:"subject"`
		Children []*Tree                `json:"children,omitempty"`
	}
)

var (
	_, _ Subject = (*SubjectID)(nil), (*SubjectSet)(nil)
)

func (s *SubjectID) Equals(other Subject) bool {
	uv, ok := other.(*SubjectID)
	if !ok {
		return false
	}
	return uv.ID == s.ID
}

func (s *SubjectID) UniqueID() uuid.UUID {
	return s.ID
}

func (s *SubjectSet) Equals(other Subject) bool {
	uv, ok := other.(*SubjectSet)
	if !ok {
		return false
	}
	return uv.Relation == s.Relation && uv.Object == s.Object && uv.Namespace == s.Namespace
}

func (s *SubjectSet) UniqueID() uuid.UUID {
	return uuid.NewV5(s.Object, s.Namespace+"-"+s.Relation)
}

func (t *RelationTuple) ToQuery() *RelationQuery {
	return &RelationQuery{
		Namespace: &t.Namespace,
		Object:    &t.Object,
		Relation:  &t.Relation,
		Subject:   t.Subject,
	}
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

func (t *ManagerWrapper) GetRelationTuples(ctx context.Context, query *RelationQuery, options ...x.PaginationOptionSetter) ([]*RelationTuple, string, error) {
	opts := x.GetPaginationOptions(options...)
	t.RequestedPages = append(t.RequestedPages, opts.Token)
	return t.Reg.RelationTupleManager().GetRelationTuples(ctx, query, append(t.PageOpts, options...)...)
}

func (t *ManagerWrapper) WriteRelationTuples(ctx context.Context, rs ...*RelationTuple) error {
	return t.Reg.RelationTupleManager().WriteRelationTuples(ctx, rs...)
}

func (t *ManagerWrapper) DeleteRelationTuples(ctx context.Context, rs ...*RelationTuple) error {
	return t.Reg.RelationTupleManager().DeleteRelationTuples(ctx, rs...)
}

func (t *ManagerWrapper) DeleteAllRelationTuples(ctx context.Context, query *RelationQuery) error {
	return t.Reg.RelationTupleManager().DeleteAllRelationTuples(ctx, query)
}

func (t *ManagerWrapper) TransactRelationTuples(ctx context.Context, insert []*RelationTuple, delete []*RelationTuple) error {
	return t.Reg.RelationTupleManager().TransactRelationTuples(ctx, insert, delete)
}

func (t *ManagerWrapper) RelationTupleManager() Manager {
	return t
}
