package ketoapi

import (
	"github.com/ory/herodot"
	"github.com/ory/x/pointerx"
	"net/url"
	"time"
)

type RelationTuple struct {
	// Namespace of the Relation Tuple
	Namespace string `json:"namespace"`

	// Object of the Relation Tuple
	Object string `json:"object"`

	// Relation of the Relation Tuple
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

	CommitTime time.Time `json:"commit_time"`
}

// swagger:parameters getExpand
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

type RelationQuery struct {
	// Namespace of the Relation Tuple
	Namespace string `json:"namespace"`

	// Object of the Relation Tuple
	Object string `json:"object"`

	// Relation of the Relation Tuple
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

const (
	subjectIDKey           = "subject_id"
	subjectSetNamespaceKey = "subject_set.namespace"
	subjectSetObjectKey    = "subject_set.object"
	subjectSetRelationKey  = "subject_set.relation"
)

var (
	ErrDroppedSubjectKey = herodot.ErrBadRequest.WithDebug(`provide "subject_id" or "subject_set.*"; support for "subject" was dropped`)
	ErrDuplicateSubject  = herodot.ErrBadRequest.WithError("exactly one of subject_set or subject_id has to be provided")
	ErrIncompleteSubject = herodot.ErrBadRequest.WithError(`incomplete subject, provide "subject_id" or a complete "subject_set.*"`)
)

func (q *RelationQuery) FromURLQuery(query url.Values) (*RelationQuery, error) {
	if q == nil {
		q = &RelationQuery{}
	}

	if query.Has("subject") {
		return nil, ErrDroppedSubjectKey
	}

	// reset subject
	q.SubjectID = nil
	q.SubjectSet = nil

	switch {
	case !query.Has(subjectIDKey) && !query.Has(subjectSetNamespaceKey) && !query.Has(subjectSetObjectKey) && !query.Has(subjectSetRelationKey):
		// was not queried for the subject
	case query.Has(subjectIDKey) && query.Has(subjectSetNamespaceKey) && query.Has(subjectSetObjectKey) && query.Has(subjectSetRelationKey):
		return nil, ErrDuplicateSubject
	case query.Has(subjectIDKey):
		q.SubjectID = pointerx.String(query.Get(subjectIDKey))
	case query.Has(subjectSetNamespaceKey) && query.Has(subjectSetObjectKey) && query.Has(subjectSetRelationKey):
		q.SubjectSet = &SubjectSet{
			Namespace: query.Get(subjectSetNamespaceKey),
			Object:    query.Get(subjectSetObjectKey),
			Relation:  query.Get(subjectSetRelationKey),
		}
	default:
		return nil, ErrIncompleteSubject
	}

	q.Object = query.Get("object")
	q.Relation = query.Get("relation")
	q.Namespace = query.Get("namespace")

	return q, nil
}
