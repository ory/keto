package ketoapi

import (
	"errors"
	"github.com/ory/herodot"
	rts "github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2"
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

type PatchDelta struct {
	Action        PatchAction    `json:"action"`
	RelationTuple *RelationTuple `json:"relation_tuple"`
}

// swagger:model PatchDeltas
type PatchDeltas []*PatchDelta

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
type ExpandNodeType string

const (
	Union        ExpandNodeType = "union"
	Exclusion    ExpandNodeType = "exclusion"
	Intersection ExpandNodeType = "intersection"
	Leaf         ExpandNodeType = "leaf"
	Unspecified  ExpandNodeType = "unspecified"
)

// swagger:model ExpandTree
type ExpandTree struct {
	Type       ExpandNodeType `json:"type"`
	Children   []*ExpandTree  `json:"children,omitempty"`
	SubjectSet *SubjectSet    `json:"subject_set,omitempty"`
	SubjectID  *string        `json:"subject_id,omitempty"`
}

func (t ExpandNodeType) String() string {
	return string(t)
}

func (t *ExpandNodeType) UnmarshalJSON(v []byte) error {
	switch string(v) {
	case `"union"`:
		*t = Union
	case `"exclusion"`:
		*t = Exclusion
	case `"intersection"`:
		*t = Intersection
	case `"leaf"`:
		*t = Leaf
	default:
		return ErrUnknownNodeType
	}
	return nil
}

func (t ExpandNodeType) ToProto() rts.NodeType {
	switch t {
	case Leaf:
		return rts.NodeType_NODE_TYPE_LEAF
	case Union:
		return rts.NodeType_NODE_TYPE_UNION
	case Exclusion:
		return rts.NodeType_NODE_TYPE_EXCLUSION
	case Intersection:
		return rts.NodeType_NODE_TYPE_INTERSECTION
	}
	return rts.NodeType_NODE_TYPE_UNSPECIFIED
}

func (ExpandNodeType) FromProto(pt rts.NodeType) ExpandNodeType {
	switch pt {
	case rts.NodeType_NODE_TYPE_LEAF:
		return Leaf
	case rts.NodeType_NODE_TYPE_UNION:
		return Union
	case rts.NodeType_NODE_TYPE_EXCLUSION:
		return Exclusion
	case rts.NodeType_NODE_TYPE_INTERSECTION:
		return Intersection
	}
	return Unspecified
}
