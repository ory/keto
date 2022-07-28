package ketoapi

import (
	"errors"

	"github.com/ory/herodot"
	"github.com/sirupsen/logrus"
)

var (
	ErrDroppedSubjectKey = herodot.ErrBadRequest.WithDebug(`provide "subject_id" or "subject_set.*"; support for "subject" was dropped`)
	ErrDuplicateSubject  = herodot.ErrBadRequest.WithError("exactly one of subject_set or subject_id has to be provided")
	ErrIncompleteSubject = herodot.ErrBadRequest.WithError(`incomplete subject, provide "subject_id" or a complete "subject_set.*"`)
	ErrNilSubject        = herodot.ErrBadRequest.WithError("subject is not allowed to be nil").WithDebug("Please provide a subject.")
	ErrIncompleteTuple   = herodot.ErrBadRequest.WithError(`incomplete tuple, provide "namespace", "object", "relation", and a subject`)
	ErrUnknownNodeType   = errors.New("unknown node type")
)

// swagger:model relationTuple
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

// swagger:model patchDelta
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
	SubjectIDKey           = "subject_id"
	SubjectSetNamespaceKey = "subject_set.namespace"
	SubjectSetObjectKey    = "subject_set.object"
	SubjectSetRelationKey  = "subject_set.relation"
)

// swagger:model getRelationTuplesResponse
type GetResponse struct {
	RelationTuples []*RelationTuple `json:"relation_tuples"`
	// The opaque token to provide in a subsequent request
	// to get the next page. It is the empty string iff this is
	// the last page.
	NextPageToken string `json:"next_page_token"`
}

func (r *RelationTuple) ToLoggerFields() logrus.Fields {
	fields := make(logrus.Fields, 7)
	q := r.ToURLQuery()
	for k := range q {
		fields[k] = q.Get(k)
	}
	return fields
}

// swagger:enum ExpandNodeType
type ExpandNodeType TreeNodeType
