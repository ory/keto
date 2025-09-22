// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package relationtuple

import (
	"context"
	"fmt"
	"sync"

	"github.com/gofrs/uuid"
	keysetpagination "github.com/ory/x/pagination/keysetpagination_v2"

	"github.com/ory/keto/ketoapi"
	rts "github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2"
)

type (
	ManagerProvider interface {
		RelationTupleManager() Manager
	}
	Traverser interface {
		TraverseSubjectSetExpansion(ctx context.Context, tuple *RelationTuple) ([]*TraversalResult, error)
		TraverseSubjectSetRewrite(ctx context.Context, tuple *RelationTuple, computedSubjectSets []string) ([]*TraversalResult, error)
	}
	Manager interface {
		GetRelationTuples(ctx context.Context, query *RelationQuery, options ...keysetpagination.Option) ([]*RelationTuple, *keysetpagination.Paginator, error)
		ExistsRelationTuples(ctx context.Context, query *RelationQuery) (bool, error)
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
		String() string
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

	// TODO(hperl): Also use a ketoapi.Tree here.
	Tree struct {
		Type     ketoapi.TreeNodeType `json:"type"`
		Subject  Subject              `json:"subject"`
		Children []*Tree              `json:"children,omitempty"`
	}

	TraversalResult struct {
		From  *RelationTuple
		To    *RelationTuple
		Via   Traversal
		Found bool
	}

	Traversal string
)

const (
	TraversalUnknown          Traversal = "unknown"
	TraversalSubjectSetExpand Traversal = "subject set expand"
	TraversalComputedUserset  Traversal = "computed userset"
	TraversalTupleToUserset   Traversal = "tuple to userset"
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

func (s *SubjectID) UniqueID() uuid.UUID { return s.ID }
func (s *SubjectID) String() string      { return s.ID.String() }

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

func (s *SubjectSet) String() string {
	return fmt.Sprintf("%s:%s#%s", s.Namespace, s.Object, s.Relation)
}

func (t *RelationTuple) ToQuery() *RelationQuery {
	return &RelationQuery{
		Namespace: &t.Namespace,
		Object:    &t.Object,
		Relation:  &t.Relation,
		Subject:   t.Subject,
	}
}

func (t *RelationTuple) String() string {
	if t == nil {
		return ""
	}
	return fmt.Sprintf("%s:%s#%s@%s", t.Namespace, t.Object, t.Relation, t.Subject)
}

func (t *RelationTuple) FromProto(proto *rts.RelationTuple) *RelationTuple {
	// TODO(hperl)
	return t
}
func (t *RelationTuple) ToProto() *rts.RelationTuple {
	// TODO(hperl)
	return &rts.RelationTuple{}
}

type ManagerWrapper struct {
	Reg            ManagerProvider
	PageOpts       []keysetpagination.Option
	RequestedPages []keysetpagination.PageToken
	// lock is necessary so that GetRelationTuples() is safe for concurrency.
	requestedPagesLock sync.Mutex
}

var (
	_ Manager         = (*ManagerWrapper)(nil)
	_ ManagerProvider = (*ManagerWrapper)(nil)
)

func NewManagerWrapper(_ any, reg ManagerProvider, options ...keysetpagination.Option) *ManagerWrapper {
	return &ManagerWrapper{
		Reg:      reg,
		PageOpts: options,
	}
}

func (t *ManagerWrapper) GetRelationTuples(ctx context.Context, query *RelationQuery, options ...keysetpagination.Option) ([]*RelationTuple, *keysetpagination.Paginator, error) {
	p := keysetpagination.NewPaginator(options...)
	t.requestedPagesLock.Lock()
	defer t.requestedPagesLock.Unlock()
	t.RequestedPages = append(t.RequestedPages, p.PageToken())
	return t.Reg.RelationTupleManager().GetRelationTuples(ctx, query, append(t.PageOpts, options...)...)
}

func (t *ManagerWrapper) ExistsRelationTuples(ctx context.Context, query *RelationQuery) (bool, error) {
	return t.Reg.RelationTupleManager().ExistsRelationTuples(ctx, query)
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
