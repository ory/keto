// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package ketoapi

import (
	"encoding/json"
	"fmt"

	"github.com/ory/herodot"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	rts "github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2"
)

var (
	ErrDroppedSubjectKey = herodot.ErrBadRequest.WithDebug(`provide "subject_id" or "subject_set.*"; support for "subject" was dropped`)
	ErrDuplicateSubject  = herodot.ErrBadRequest.WithError("exactly one of subject_set or subject_id has to be provided")
	ErrIncompleteSubject = herodot.ErrBadRequest.WithError(`incomplete subject, provide "subject_id" or a complete "subject_set.*"`)
	ErrNilSubject        = herodot.ErrBadRequest.WithError("subject is not allowed to be nil").WithDebug("Please provide a subject.")
	ErrIncompleteTuple   = herodot.ErrBadRequest.WithError(`incomplete tuple, provide "namespace", "object", "relation", and a subject`)
	ErrUnknownNodeType   = errors.New("unknown node type")
)

// swagger:model namespace
type Namespace struct {
	// Name of the namespace.
	Name string `json:"name"`
}

// Relationship
//
// swagger:model relationship
type RelationTuple struct {
	// Namespace of the Relation Tuple
	//
	// required: true
	Namespace string `json:"namespace"`

	// Object of the Relation Tuple
	//
	// required: true
	Object string `json:"object"`

	// Relation of the Relation Tuple
	//
	// required: true
	Relation string `json:"relation"`

	// SubjectID of the Relation Tuple
	//
	// Either SubjectSet or SubjectID can be provided.
	SubjectID *string `json:"subject_id,omitempty"`

	// SubjectSet of the Relation Tuple
	//
	// Either SubjectSet or SubjectID can be provided.
	//
	// swagger:allOf
	SubjectSet *SubjectSet `json:"subject_set,omitempty"`
}

// swagger:model subjectSet
type SubjectSet struct {
	// Namespace of the Subject Set
	//
	// required: true
	Namespace string `json:"namespace"`

	// Object of the Subject Set
	//
	// required: true
	Object string `json:"object"`

	// Relation of the Subject Set
	//
	// required: true
	Relation string `json:"relation"`
}

// Relation Query
//
// swagger:model relationQuery
type RelationQuery struct {
	// Namespace to query
	Namespace *string `json:"namespace"`

	// Object to query
	Object *string `json:"object"`

	// Relation to query
	Relation *string `json:"relation"`

	// SubjectID to query
	//
	// Either SubjectSet or SubjectID can be provided.
	SubjectID *string `json:"subject_id,omitempty"`
	// SubjectSet to query
	//
	// Either SubjectSet or SubjectID can be provided.
	//
	// swagger:allOf
	SubjectSet *SubjectSet `json:"subject_set,omitempty"`
}

// Payload for patching a relationship
//
// swagger:model relationshipPatch
type PatchDelta struct {
	Action        PatchAction    `json:"action"`
	RelationTuple *RelationTuple `json:"relation_tuple"`
}

// swagger:enum PatchAction
type PatchAction string

const (
	ActionInsert PatchAction = "insert"
	ActionDelete PatchAction = "delete"
)

const (
	NamespaceKey           = "namespace"
	ObjectKey              = "object"
	RelationKey            = "relation"
	SubjectIDKey           = "subject_id"
	SubjectSetNamespaceKey = "subject_set.namespace"
	SubjectSetObjectKey    = "subject_set.object"
	SubjectSetRelationKey  = "subject_set.relation"
)

var RelationQueryKeys = []string{
	NamespaceKey,
	ObjectKey,
	RelationKey,
	SubjectIDKey,
	SubjectSetNamespaceKey,
	SubjectSetObjectKey,
	SubjectSetRelationKey,
	"subject", // We have a more specific error message for this key.
}

// Paginated Relationship List
//
// swagger:model relationships
type GetResponse struct {
	RelationTuples []*RelationTuple `json:"relation_tuples"`
	// The opaque token to provide in a subsequent request
	// to get the next page. It is the empty string iff this is
	// the last page.
	NextPageToken string `json:"next_page_token"`
}

// Relationship Namespace List
//
// swagger:model relationshipNamespaces
type GetNamespacesResponse struct {
	Namespaces []Namespace `json:"namespaces"`
}

func (r *RelationTuple) ToLoggerFields() logrus.Fields {
	fields := make(logrus.Fields, 7)
	q := r.ToURLQuery()
	for k := range q {
		fields[k] = q.Get(k)
	}
	return fields
}

func (r *RelationTuple) Validate() error {
	if r.SubjectSet == nil && r.SubjectID == nil {
		return errors.WithStack(ErrNilSubject)
	}
	return nil
}

// swagger:enum ExpandNodeType
type ExpandNodeType TreeNodeType

// swagger:enum TreeNodeType
type TreeNodeType string

const (
	TreeNodeUnion              TreeNodeType = "union"
	TreeNodeExclusion          TreeNodeType = "exclusion"
	TreeNodeIntersection       TreeNodeType = "intersection"
	TreeNodeLeaf               TreeNodeType = "leaf"
	TreeNodeTupleToSubjectSet  TreeNodeType = "tuple_to_subject_set"
	TreeNodeComputedSubjectSet TreeNodeType = "computed_subject_set"
	TreeNodeNot                TreeNodeType = "not"
	TreeNodeUnspecified        TreeNodeType = "unspecified"
)

func (t *TreeNodeType) UnmarshalJSON(v []byte) error {
	var s string
	if err := json.Unmarshal(v, &s); err != nil {
		return err
	}
	switch nt := TreeNodeType(s); nt {
	case TreeNodeUnion, TreeNodeExclusion, TreeNodeIntersection, TreeNodeLeaf, TreeNodeTupleToSubjectSet, TreeNodeComputedSubjectSet, TreeNodeNot, TreeNodeUnspecified:
		*t = nt
	default:
		return ErrUnknownNodeType
	}
	return nil
}

type tuple[T any] interface {
	fmt.Stringer
	ToProto() *rts.RelationTuple
	FromProto(*rts.RelationTuple) T
}

// Tree is a generic tree of either internal relationships (with UUIDs for
// objects, etc.) or API relationships (with strings for objects, etc.).
type Tree[T tuple[T]] struct {
	// Propagate all struct changes to `swaggerOnlyExpandTree` as well.
	// The type of the node.
	//
	// required: true
	Type TreeNodeType `json:"type"`

	// The children of the node, possibly none.
	Children []*Tree[T] `json:"children,omitempty"`

	// The relation tuple the node represents.
	Tuple T `json:"tuple"`
}

// IMPORTANT: We need a manual instantiation of the generic Tree[T] for the
// OpenAPI spec, since go-swagger does not understand generics :(.
// This can be fixed by using grpc-gateway.

// swagger:model expandedPermissionTree
//
//lint:ignore U1000 Used to generate Swagger and OpenAPI definitions
type swaggerOnlyExpandTree struct {
	// The type of the node.
	//
	// required: true
	Type TreeNodeType `json:"type"`

	// The children of the node, possibly none.
	Children []*swaggerOnlyExpandTree `json:"children,omitempty"`

	// The relation tuple the node represents.
	Tuple *RelationTuple `json:"tuple"`
}

type ParseError struct {
	Message string         `json:"message"`
	Start   SourcePosition `json:"start"`
	End     SourcePosition `json:"end"`
}
type SourcePosition struct {
	Line int `json:"Line"`
	Col  int `json:"column"`
}

// CheckOPLSyntaxResponse represents the response for an OPL syntax check request.
//
// swagger:model checkOplSyntaxResult
type CheckOPLSyntaxResponse struct {
	// The list of syntax errors
	//
	// required: false
	Errors []*ParseError `json:"errors,omitempty"`
}
